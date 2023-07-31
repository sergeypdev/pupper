package pupper

import (
	"math"
	"testing"
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

func (d *data) Pup(p *Pupper) int {
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
	p := Pupper{}

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
  dataLen := result.Pup(&p)

  encodedBytes := make([]byte, dataLen)
  p = Pupper{
    Data: encodedBytes,
  }
  // Encoding
  result.Pup(&p)

  var result2 data
  p = Pupper{
    Unpack: true,
    Data: encodedBytes,
  }
  // Decoding
  result2.Pup(&p)

  if result != result2 {
    t.Errorf("Decoding the encoded value yields a different result\nEncoded:\t%v\nDecoded:\t%v\n", result, result2)
  }
}
