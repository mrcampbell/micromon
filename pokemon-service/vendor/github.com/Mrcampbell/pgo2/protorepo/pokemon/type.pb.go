// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pokemon/type.proto

package pokemon

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Type int32

const (
	Type_UNKNOWN  Type = 0
	Type_BUG      Type = 1
	Type_DARK     Type = 2
	Type_DRAGON   Type = 3
	Type_ELECTRIC Type = 4
	Type_FAIRY    Type = 5
	Type_FIGHTING Type = 6
	Type_FIRE     Type = 7
	Type_FLYING   Type = 8
	Type_GHOST    Type = 9
	Type_GRASS    Type = 10
	Type_GROUND   Type = 11
	Type_ICE      Type = 12
	Type_NORMAL   Type = 13
	Type_POISON   Type = 14
	Type_PSYCHIC  Type = 15
	Type_ROCK     Type = 16
	Type_STEEL    Type = 17
	Type_WATER    Type = 18
)

var Type_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "BUG",
	2:  "DARK",
	3:  "DRAGON",
	4:  "ELECTRIC",
	5:  "FAIRY",
	6:  "FIGHTING",
	7:  "FIRE",
	8:  "FLYING",
	9:  "GHOST",
	10: "GRASS",
	11: "GROUND",
	12: "ICE",
	13: "NORMAL",
	14: "POISON",
	15: "PSYCHIC",
	16: "ROCK",
	17: "STEEL",
	18: "WATER",
}

var Type_value = map[string]int32{
	"UNKNOWN":  0,
	"BUG":      1,
	"DARK":     2,
	"DRAGON":   3,
	"ELECTRIC": 4,
	"FAIRY":    5,
	"FIGHTING": 6,
	"FIRE":     7,
	"FLYING":   8,
	"GHOST":    9,
	"GRASS":    10,
	"GROUND":   11,
	"ICE":      12,
	"NORMAL":   13,
	"POISON":   14,
	"PSYCHIC":  15,
	"ROCK":     16,
	"STEEL":    17,
	"WATER":    18,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0985705e2f445762, []int{0}
}

func init() {
	proto.RegisterEnum("pokemon.Type", Type_name, Type_value)
}

func init() { proto.RegisterFile("pokemon/type.proto", fileDescriptor_0985705e2f445762) }

var fileDescriptor_0985705e2f445762 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0x4d, 0x4e, 0xc3, 0x30,
	0x14, 0x84, 0x29, 0x4d, 0xf3, 0xe3, 0x16, 0x18, 0x7c, 0x0c, 0x16, 0xb0, 0xe0, 0x04, 0xae, 0xe3,
	0x38, 0x56, 0x82, 0x5d, 0x3d, 0x3b, 0xaa, 0xc2, 0x12, 0x29, 0x2b, 0x04, 0x89, 0x50, 0x37, 0x3d,
	0x2d, 0x57, 0x41, 0xaf, 0xa2, 0xbb, 0xd1, 0xe7, 0x19, 0xcf, 0x1b, 0x21, 0x97, 0xf9, 0x73, 0xfa,
	0x9a, 0xbf, 0x5f, 0x4e, 0xe7, 0x65, 0x7a, 0x5e, 0x7e, 0xe6, 0xd3, 0x2c, 0x8b, 0x7f, 0xf6, 0xf4,
	0xbb, 0x12, 0x59, 0x3a, 0x2f, 0x93, 0xdc, 0x8a, 0x62, 0xf0, 0x9d, 0x0f, 0x47, 0x8f, 0x1b, 0x59,
	0x88, 0xf5, 0x7e, 0xb0, 0x58, 0xc9, 0x52, 0x64, 0xb5, 0xa2, 0x0e, 0xb7, 0x52, 0x88, 0xbc, 0x26,
	0x65, 0x83, 0xc7, 0x5a, 0xee, 0x44, 0x69, 0x7a, 0xa3, 0x13, 0x39, 0x8d, 0x4c, 0x56, 0x62, 0xd3,
	0x28, 0x47, 0x23, 0x36, 0xfc, 0xd0, 0x38, 0xdb, 0x26, 0xe7, 0x2d, 0x72, 0x0e, 0x37, 0x8e, 0x0c,
	0x0a, 0x0e, 0x37, 0xfd, 0xc8, 0xb4, 0x64, 0xbb, 0x6d, 0x43, 0x4c, 0xa8, 0x2e, 0x92, 0x54, 0x8c,
	0x10, 0xec, 0xb0, 0x14, 0x06, 0x5f, 0x63, 0xcb, 0xed, 0x4e, 0x1b, 0xec, 0x18, 0xfa, 0x40, 0x6f,
	0xaa, 0xc7, 0x1d, 0xeb, 0x43, 0x70, 0x31, 0x78, 0xdc, 0xf3, 0xad, 0x87, 0x38, 0xea, 0xd6, 0x69,
	0x3c, 0x70, 0x0b, 0x05, 0xdd, 0x01, 0xfc, 0x5d, 0x4c, 0xc6, 0xf4, 0x78, 0x64, 0x79, 0x54, 0xc9,
	0x10, 0xe4, 0xbe, 0x7a, 0xbf, 0x8e, 0xfd, 0xc8, 0x2f, 0xe3, 0x5f, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xe3, 0x92, 0x97, 0x09, 0x12, 0x01, 0x00, 0x00,
}
