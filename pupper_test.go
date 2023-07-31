package pupper_test

import (
	"math"
	"testing"

	"github.com/sergeypdev/pupper"
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
	p.Int16LE(&d.i16)
	p.Uint16LE(&d.u16)
	p.Int32LE(&d.i32)
	p.Uint32LE(&d.u32)
	p.Int64LE(&d.i64)
	p.Uint64LE(&d.u64)
	p.Float32LE(&d.f32)
	p.Float64LE(&d.f64)
	return p.Len()
}

func TestCalcLength(t *testing.T) {
	data1 := data{
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

	encodedBytes := make([]byte, data1.Pup(pupper.Count()))
	// Encoding
	data1.Pup(pupper.Pack(encodedBytes))

	var result2 data
	// Decoding
	result2.Pup(pupper.Unpack(encodedBytes))

	if data1 != result2 {
		t.Errorf("Decoding the encoded value yields a different result\nEncoded:\t%v\nDecoded:\t%v\n", data1, result2)
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
	p.Uint32LE(&version)
	p.Int32LE(&vd.field)
	return p.Len()
}

func (vd *versionedData) PupV2(p pupper.P) int {
	// Specify current version
	version := uint32(2)
	p.Uint32LE(&version)
	p.Int32LE(&vd.field)
	if version >= 2 {
		p.Int8(&vd.fieldAddedV2)
	}
	return p.Len()
}

func (vd *versionedData) PupV3(p pupper.P) int {
	// Specify current version
	version := uint32(3)
	p.Uint32LE(&version)
	p.Int32LE(&vd.field)
	if version >= 2 {
		p.Int8(&vd.fieldAddedV2)
	}
	if version >= 3 {
		p.Int16LE(&vd.fieldAddedV3)
	}
	return p.Len()
}

func TestVersionUpgrade(t *testing.T) {
	dataV1 := versionedData{
		field: 1,
	}
	encodedDataV1 := make([]byte, dataV1.PupV1(pupper.Count()))
	dataV1.PupV1(pupper.Pack(encodedDataV1))

	var dataV2 versionedData
	dataV2.PupV2(pupper.Unpack(encodedDataV1))

	if dataV1 != dataV2 {
		t.Errorf("Decoding V1 data with V2 Pup produced a different result\nEncoded:\t%v\nDecoded:\t%v\n", dataV1, dataV2)
	}

	dataV2.fieldAddedV2 = 2
	encodedDataV2 := make([]byte, dataV2.PupV2(pupper.Count()))
	dataV2.PupV2(pupper.Pack(encodedDataV2))

	var dataV3 versionedData
	dataV3.PupV3(pupper.Unpack(encodedDataV2))

	if dataV2 != dataV3 {
		t.Errorf("Decoding V2 data with V3 Pup produced a different result\nEncoded:\t%v\nDecoded:\t%v\n", dataV2, dataV3)
	}

	dataV3.fieldAddedV3 = 3
	encodedDataV3 := make([]byte, dataV3.PupV3(pupper.Count()))
	dataV3.PupV3(pupper.Pack(encodedDataV3))

	// Checking backwards compat

	dataV2 = versionedData{}
	dataV2.PupV2(pupper.Unpack(encodedDataV3))

	expectedDataV2 := versionedData{field: 1, fieldAddedV2: 2}
	if dataV2 != expectedDataV2 {
		t.Errorf("Decoding V3 data with V2 Pup produced a different result\nEncoded:\t%v\nDecoded:\t%v\n", dataV2, expectedDataV2)
	}

  dataV1 = versionedData{}
  dataV1.PupV1(pupper.Unpack(encodedDataV3))

  expectedDataV1 := versionedData{field: 1}
  if dataV1 != expectedDataV1 {
		t.Errorf("Decoding V3 data with V1 Pup produced a different result\nEncoded:\t%v\nDecoded:\t%v\n", dataV1, expectedDataV1)
  }
}
