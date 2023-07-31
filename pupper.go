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
	Unpack bool
	Data   []byte
	Cursor int
}

// Use this Pupper to count number of bytes
// you need to pack your data.
//
// Pupper will never allocate data for you, use pupper.Count()
// to check how many bytes you need and allocate them yourself.
func Count() P {
	return P{}
}

// Use this Pupper to pack your data.
//
// Pupper will never allocate data for you, use pupper.Count()
// to check how many bytes you need and allocate them yourself.
func Pack(data []byte) P {
	return P{Data: data}
}

// Use this Pupper to unpack your data.
func Unpack(data []byte) P {
	return P{Unpack: true, Data: data}
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

// Little Endian

func (p *P) Int16LE(value *int16) {
	if p.enoughLength(2) {
		if p.Unpack {
			*value = int16(binary.LittleEndian.Uint16(p.Data[p.Cursor:]))
		} else {
			binary.LittleEndian.PutUint16(p.Data[p.Cursor:], uint16(*value))
		}
	}
	p.Cursor += 2
}
func (p *P) Uint16LE(value *uint16) {
	if p.enoughLength(2) {
		if p.Unpack {
			*value = binary.LittleEndian.Uint16(p.Data[p.Cursor:])
		} else {
			binary.LittleEndian.PutUint16(p.Data[p.Cursor:], *value)
		}
	}
	p.Cursor += 2
}

func (p *P) Int32LE(value *int32) {
	if p.enoughLength(4) {
		if p.Unpack {
			*value = int32(binary.LittleEndian.Uint32(p.Data[p.Cursor:]))
		} else {
			binary.LittleEndian.PutUint32(p.Data[p.Cursor:], uint32(*value))
		}
	}
	p.Cursor += 4
}
func (p *P) Uint32LE(value *uint32) {
	if p.enoughLength(4) {
		if p.Unpack {
			*value = binary.LittleEndian.Uint32(p.Data[p.Cursor:])
		} else {
			binary.LittleEndian.PutUint32(p.Data[p.Cursor:], *value)
		}
	}
	p.Cursor += 4
}

func (p *P) Int64LE(value *int64) {
	if p.enoughLength(8) {
		if p.Unpack {
			*value = int64(binary.LittleEndian.Uint64(p.Data[p.Cursor:]))
		} else {
			binary.LittleEndian.PutUint64(p.Data[p.Cursor:], uint64(*value))
		}
	}
	p.Cursor += 8
}

func (p *P) Uint64LE(value *uint64) {
	if p.enoughLength(8) {
		if p.Unpack {
			*value = binary.LittleEndian.Uint64(p.Data[p.Cursor:])
		} else {
			binary.LittleEndian.PutUint64(p.Data[p.Cursor:], *value)
		}
	}
	p.Cursor += 8
}

func (p *P) Float32LE(value *float32) {
	if p.enoughLength(4) {
		if p.Unpack {
			*value = math.Float32frombits(binary.LittleEndian.Uint32(p.Data[p.Cursor:]))
		} else {
			binary.LittleEndian.PutUint32(p.Data[p.Cursor:], math.Float32bits(*value))
		}
	}
	p.Cursor += 4
}

func (p *P) Float64LE(value *float64) {
	if p.enoughLength(8) {
		if p.Unpack {
			*value = math.Float64frombits(binary.LittleEndian.Uint64(p.Data[p.Cursor:]))
		} else {
			binary.LittleEndian.PutUint64(p.Data[p.Cursor:], math.Float64bits(*value))
		}
	}
	p.Cursor += 8
}

// Big Endian

func (p *P) Int16BE(value *int16) {
	if p.enoughLength(2) {
		if p.Unpack {
			*value = int16(binary.BigEndian.Uint16(p.Data[p.Cursor:]))
		} else {
			binary.BigEndian.PutUint16(p.Data[p.Cursor:], uint16(*value))
		}
	}
	p.Cursor += 2
}
func (p *P) Uint16BE(value *uint16) {
	if p.enoughLength(2) {
		if p.Unpack {
			*value = binary.BigEndian.Uint16(p.Data[p.Cursor:])
		} else {
			binary.BigEndian.PutUint16(p.Data[p.Cursor:], *value)
		}
	}
	p.Cursor += 2
}

func (p *P) Int32BE(value *int32) {
	if p.enoughLength(4) {
		if p.Unpack {
			*value = int32(binary.BigEndian.Uint32(p.Data[p.Cursor:]))
		} else {
			binary.BigEndian.PutUint32(p.Data[p.Cursor:], uint32(*value))
		}
	}
	p.Cursor += 4
}
func (p *P) Uint32BE(value *uint32) {
	if p.enoughLength(4) {
		if p.Unpack {
			*value = binary.BigEndian.Uint32(p.Data[p.Cursor:])
		} else {
			binary.BigEndian.PutUint32(p.Data[p.Cursor:], *value)
		}
	}
	p.Cursor += 4
}

func (p *P) Int64BE(value *int64) {
	if p.enoughLength(8) {
		if p.Unpack {
			*value = int64(binary.BigEndian.Uint64(p.Data[p.Cursor:]))
		} else {
			binary.BigEndian.PutUint64(p.Data[p.Cursor:], uint64(*value))
		}
	}
	p.Cursor += 8
}

func (p *P) Uint64BE(value *uint64) {
	if p.enoughLength(8) {
		if p.Unpack {
			*value = binary.BigEndian.Uint64(p.Data[p.Cursor:])
		} else {
			binary.BigEndian.PutUint64(p.Data[p.Cursor:], *value)
		}
	}
	p.Cursor += 8
}

func (p *P) Float32BE(value *float32) {
	if p.enoughLength(4) {
		if p.Unpack {
			*value = math.Float32frombits(binary.BigEndian.Uint32(p.Data[p.Cursor:]))
		} else {
			binary.BigEndian.PutUint32(p.Data[p.Cursor:], math.Float32bits(*value))
		}
	}
	p.Cursor += 4
}

func (p *P) Float64BE(value *float64) {
	if p.enoughLength(8) {
		if p.Unpack {
			*value = math.Float64frombits(binary.BigEndian.Uint64(p.Data[p.Cursor:]))
		} else {
			binary.BigEndian.PutUint64(p.Data[p.Cursor:], math.Float64bits(*value))
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
