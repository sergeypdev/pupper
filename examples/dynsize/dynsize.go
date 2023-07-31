/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package main

import (
	"fmt"

	"github.com/sergeypdev/pupper"
)

type Struct struct {
	SubStructs []SubStruct
}

func (s *Struct) Pup(p *pupper.P) int {
  version := uint32(1)
  p.Uint32LE(&version)

  length, lenChanged := p.SliceLenLE(len(s.SubStructs))
  
  if lenChanged {
    s.SubStructs = make([]SubStruct, length)
  }

  for i := 0; i < int(length); i++ {
    s.SubStructs[i].Pup(p)
  }

  return p.Len()
}

type SubStruct struct {
	field int32
}

func (s *SubStruct) Pup(p *pupper.P) int {
  p.Int32LE(&s.field)
  return p.Len()
}

func main() {
  data1 := Struct{
    SubStructs: []SubStruct{
      {field: 123},
      {field: 234},
    },
  }
  packed := make([]byte, data1.Pup(pupper.Count()))
  data1.Pup(pupper.Pack(packed))

  var data2 Struct
  data2.Pup(pupper.Unpack(packed))

  data2.Pup(pupper.Unpack(packed))

  fmt.Printf("Data1:\n%v\nData2:\n%v\n", data1, data2)
}
