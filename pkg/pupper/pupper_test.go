package pupper_test

import (
	"math"
	"testing"

	"github.com/sergeypdev/pupper/pkg/pupper"
)

type data struct {
	i8  int8
	u8  uint8
	i16 int16
	u16 uint16
	i32 int32
	u32 uint32
	i64 int64
	u64 uint64
	f32 float32
	f64 float64
}

func (d *data) Pup(p pupper.P) int {
	p.Int8(&d.i8)
	p.Uint8(&d.u8)
	p.Int16(&d.i16)
	p.Uint16(&d.u16)
	p.Int32(&d.i32)
	p.Uint32(&d.u32)
	p.Int64(&d.i64)
	p.Uint64(&d.u64)
	p.Float32(&d.f32)
	p.Float64(&d.f64)
	return p.Len()
}

func TestCalcLength(t *testing.T) {
	result := data{
		i8:  -50,
		u8:  math.MaxUint8,
		i16: -1024,
		u16: math.MaxUint16,
		i32: -128000,
		u32: math.MaxUint32,
		i64: -4294967296,
		u64: math.MaxUint64,
		f32: math.MaxFloat32,
		f64: math.MaxFloat64,
	}

	// Calculate size, not writing anything yet
	dataLen := result.Pup(pupper.Count())

	encodedBytes := make([]byte, dataLen)
  p := pupper.PackLittleEndian(encodedBytes)
	// Encoding
	result.Pup(p)

	var result2 data
	p = pupper.UnpackLittleEndian(encodedBytes)
	// Decoding
	result2.Pup(p)

	if result != result2 {
		t.Errorf("Decoding the encoded value yields a different result\nEncoded:\t%v\nDecoded:\t%v\n", result, result2)
	}
}

type versionedData struct {
	field        int32
	fieldAddedV2 int8
	fieldAddedV3 int16
}

// These will not be separate functions, simulating data format
// evolution over time. In real projects you would have only one Pup()
// function that is the latest version
func (vd *versionedData) PupV1(p pupper.P) int {
	// Specify current version
	version := uint32(1)
	p.Uint32(&version)
	p.Int32(&vd.field)
	return p.Len()
}

func (vd *versionedData) PupV2(p pupper.P) int {
	// Specify current version
	version := uint32(2)
	p.Uint32(&version)
	p.Int32(&vd.field)
	if version >= 2 {
		p.Int8(&vd.fieldAddedV2)
	}
	return p.Len()
}

func (vd *versionedData) PupV3(p pupper.P) int {
	// Specify current version
	version := uint32(3)
	p.Uint32(&version)
	p.Int32(&vd.field)
	if version >= 2 {
		p.Int8(&vd.fieldAddedV2)
	}
	if version >= 3 {
		p.Int16(&vd.fieldAddedV3)
	}
	return p.Len()
}

func TestVersionUpgrade(t *testing.T) {
	dataV1 := versionedData{
		field: 1,
	}
	encodedDataV1 := make([]byte, dataV1.PupV1(pupper.Count()))
	dataV1.PupV1(pupper.PackLittleEndian(encodedDataV1))

	var dataV2 versionedData
	dataV2.PupV2(pupper.UnpackLittleEndian(encodedDataV1))

	if dataV1 != dataV2 {
		t.Errorf("Decoding V1 data with V2 Pup produced a different result\nEncoded:\t%v\nDecoded:\t%v\n", dataV1, dataV2)
	}

	dataV2.fieldAddedV2 = 2
	encodedDataV2 := make([]byte, dataV2.PupV2(pupper.Count()))
	dataV2.PupV2(pupper.PackLittleEndian(encodedDataV2))

	var dataV3 versionedData
	dataV3.PupV3(pupper.UnpackLittleEndian(encodedDataV2))

	if dataV2 != dataV3 {
		t.Errorf("Decoding V2 data with V3 Pup produced a different result\nEncoded:\t%v\nDecoded:\t%v\n", dataV2, dataV3)
	}

	dataV3.fieldAddedV3 = 3
	encodedDataV3 := make([]byte, dataV3.PupV3(pupper.Count()))
	dataV3.PupV3(pupper.PackLittleEndian(encodedDataV3))

	// Checking backwards compat

	dataV2 = versionedData{}
	dataV2.PupV2(pupper.UnpackLittleEndian(encodedDataV3))

	expectedDataV2 := versionedData{field: 1, fieldAddedV2: 2}
	if dataV2 != expectedDataV2 {
		t.Errorf("Decoding V3 data with V2 Pup produced a different result\nEncoded:\t%v\nDecoded:\t%v\n", dataV2, expectedDataV2)
	}

  dataV1 = versionedData{}
  dataV1.PupV1(pupper.UnpackLittleEndian(encodedDataV3))

  expectedDataV1 := versionedData{field: 1}
  if dataV1 != expectedDataV1 {
		t.Errorf("Decoding V3 data with V1 Pup produced a different result\nEncoded:\t%v\nDecoded:\t%v\n", dataV1, expectedDataV1)
  }
}
