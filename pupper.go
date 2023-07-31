// Pupper! (Packer/Unpacker)
// You can use Pupper to pack or unpack your
// data, by writing a single Pup function.
package pupper

import (
	"encoding/binary"
	"math"
)

// Main struct used by this module.
// It contains data and configuration required
// to pack or unpack your data.
type P struct {
	Unpack    bool
	BigEndian bool
	Data      []byte
	Cursor    int
}

// Use this Pupper to count number of bytes
// you need to pack your data.
//
// Pupper will never allocate data for you, use pupper.Count()
// to check how many bytes you need and allocate them yourself.
func Count() P {
  return P{}
}

// Use this Pupper to pack your data using little endian.
//
// Pupper will never allocate data for you, use pupper.Count()
// to check how many bytes you need and allocate them yourself.
func PackLittleEndian(data []byte) P {
  return P{Data: data}
}

// Use this Pupper to unpack your data packed as little endian.
func UnpackLittleEndian(data []byte) P {
  return P{Unpack: true, Data: data}
}

// Use this Pupper to pack your data using big endian.
//
// Pupper will never allocate data for you, use pupper.Count()
// to check how many bytes you need and allocate them yourself.
func PackBigEndian(data []byte) P {
  return P{Data: data, BigEndian: true}
}

// Use this Pupper to unpack your data packed as big endian.
func UnpackBigEndian(data []byte) P {
  return P{Unpack: true, Data: data, BigEndian: true}
}

func (p *P) Len() int {
	return p.Cursor
}

func (p *P) enoughLength(requiredLength int) bool {
	return p.Data != nil && len(p.Data[p.Cursor:]) >= requiredLength
}

func (p *P) Int8(value *int8) {
	if p.enoughLength(1) {
		if p.Unpack {
			*value = int8(p.Data[p.Cursor])
		} else {
			p.Data[p.Cursor] = uint8(*value)
		}
	}
	p.Cursor += 1
}

func (p *P) Uint8(value *uint8) {
	if p.enoughLength(1) {
		if p.Unpack {
			*value = p.Data[p.Cursor]
		} else {
			p.Data[p.Cursor] = *value
		}
	}
	p.Cursor += 1
}

func (p *P) Int16(value *int16) {
	if p.enoughLength(2) {
		if p.Unpack {
			if p.BigEndian {
				*value = int16(binary.BigEndian.Uint16(p.Data[p.Cursor:]))
			} else {
				*value = int16(binary.LittleEndian.Uint16(p.Data[p.Cursor:]))
			}
		} else {
			if p.BigEndian {
				binary.BigEndian.PutUint16(p.Data[p.Cursor:], uint16(*value))
			} else {
				binary.LittleEndian.PutUint16(p.Data[p.Cursor:], uint16(*value))
			}
		}
	}
	p.Cursor += 2
}
func (p *P) Uint16(value *uint16) {
	if p.enoughLength(2) {
		if p.Unpack {
			if p.BigEndian {
				*value = binary.BigEndian.Uint16(p.Data[p.Cursor:])
			} else {
				*value = binary.LittleEndian.Uint16(p.Data[p.Cursor:])
			}
		} else {
			if p.BigEndian {
				binary.BigEndian.PutUint16(p.Data[p.Cursor:], *value)
			} else {
				binary.LittleEndian.PutUint16(p.Data[p.Cursor:], *value)
			}
		}
	}
	p.Cursor += 2
}

func (p *P) Int32(value *int32) {
	if p.enoughLength(4) {
		if p.Unpack {
			if p.BigEndian {
				*value = int32(binary.BigEndian.Uint32(p.Data[p.Cursor:]))
			} else {
				*value = int32(binary.LittleEndian.Uint32(p.Data[p.Cursor:]))
			}
		} else {
			if p.BigEndian {
				binary.BigEndian.PutUint32(p.Data[p.Cursor:], uint32(*value))
			} else {
				binary.LittleEndian.PutUint32(p.Data[p.Cursor:], uint32(*value))
			}
		}
	}
	p.Cursor += 4
}
func (p *P) Uint32(value *uint32) {
	if p.enoughLength(4) {
		if p.Unpack {
			if p.BigEndian {
				*value = binary.BigEndian.Uint32(p.Data[p.Cursor:])
			} else {
				*value = binary.LittleEndian.Uint32(p.Data[p.Cursor:])
			}
		} else {
			if p.BigEndian {
				binary.BigEndian.PutUint32(p.Data[p.Cursor:], *value)
			} else {
				binary.LittleEndian.PutUint32(p.Data[p.Cursor:], *value)
			}
		}
	}
	p.Cursor += 4
}

func (p *P) Int64(value *int64) {
	if p.enoughLength(8) {
		if p.Unpack {
			if p.BigEndian {
				*value = int64(binary.BigEndian.Uint64(p.Data[p.Cursor:]))
			} else {
				*value = int64(binary.LittleEndian.Uint64(p.Data[p.Cursor:]))
			}
		} else {
			if p.BigEndian {
				binary.BigEndian.PutUint64(p.Data[p.Cursor:], uint64(*value))
			} else {
				binary.LittleEndian.PutUint64(p.Data[p.Cursor:], uint64(*value))
			}
		}
	}
	p.Cursor += 8
}
func (p *P) Uint64(value *uint64) {
	if p.enoughLength(8) {
		if p.Unpack {
			if p.BigEndian {
				*value = binary.BigEndian.Uint64(p.Data[p.Cursor:])
			} else {
				*value = binary.LittleEndian.Uint64(p.Data[p.Cursor:])
			}
		} else {
			if p.BigEndian {
				binary.BigEndian.PutUint64(p.Data[p.Cursor:], *value)
			} else {
				binary.LittleEndian.PutUint64(p.Data[p.Cursor:], *value)
			}
		}
	}
	p.Cursor += 8
}

func (p *P) Float32(value *float32) {
	if p.enoughLength(4) {
		if p.Unpack {
			if p.BigEndian {
				*value = math.Float32frombits(binary.BigEndian.Uint32(p.Data[p.Cursor:]))
			} else {
				*value = math.Float32frombits(binary.LittleEndian.Uint32(p.Data[p.Cursor:]))
			}
		} else {
			if p.BigEndian {
				binary.BigEndian.PutUint32(p.Data[p.Cursor:], math.Float32bits(*value))
			} else {
				binary.LittleEndian.PutUint32(p.Data[p.Cursor:], math.Float32bits(*value))
			}
		}
	}
	p.Cursor += 4
}

func (p *P) Float64(value *float64) {
	if p.enoughLength(8) {
		if p.Unpack {
			if p.BigEndian {
				*value = math.Float64frombits(binary.BigEndian.Uint64(p.Data[p.Cursor:]))
			} else {
				*value = math.Float64frombits(binary.LittleEndian.Uint64(p.Data[p.Cursor:]))
			}
    } else {
			if p.BigEndian {
				binary.BigEndian.PutUint64(p.Data[p.Cursor:], math.Float64bits(*value))
			} else {
				binary.LittleEndian.PutUint64(p.Data[p.Cursor:], math.Float64bits(*value))
			}
		}
	}
	p.Cursor += 8
}

func (p *P) Bytes(value []byte) {
	if p.Unpack {
		p.Cursor += copy(value, p.Data)
	} else {
		// Special case for counting how many bytes are required
		if len(p.Data) == 0 {
			p.Cursor += len(value)
		} else {
			p.Cursor += copy(p.Data[p.Cursor:], value)
		}
	}
}
