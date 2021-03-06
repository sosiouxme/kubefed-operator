// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/vision/v1p4beta1/geometry.proto

package vision // import "google.golang.org/genproto/googleapis/cloud/vision/v1p4beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A vertex represents a 2D point in the image.
// NOTE: the vertex coordinates are in the same scale as the original image.
type Vertex struct {
	// X coordinate.
	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	// Y coordinate.
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vertex) Reset()         { *m = Vertex{} }
func (m *Vertex) String() string { return proto.CompactTextString(m) }
func (*Vertex) ProtoMessage()    {}
func (*Vertex) Descriptor() ([]byte, []int) {
	return fileDescriptor_geometry_d92abf8dc265831c, []int{0}
}
func (m *Vertex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vertex.Unmarshal(m, b)
}
func (m *Vertex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vertex.Marshal(b, m, deterministic)
}
func (dst *Vertex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vertex.Merge(dst, src)
}
func (m *Vertex) XXX_Size() int {
	return xxx_messageInfo_Vertex.Size(m)
}
func (m *Vertex) XXX_DiscardUnknown() {
	xxx_messageInfo_Vertex.DiscardUnknown(m)
}

var xxx_messageInfo_Vertex proto.InternalMessageInfo

func (m *Vertex) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vertex) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

// A vertex represents a 2D point in the image.
// NOTE: the normalized vertex coordinates are relative to the original image
// and range from 0 to 1.
type NormalizedVertex struct {
	// X coordinate.
	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	// Y coordinate.
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NormalizedVertex) Reset()         { *m = NormalizedVertex{} }
func (m *NormalizedVertex) String() string { return proto.CompactTextString(m) }
func (*NormalizedVertex) ProtoMessage()    {}
func (*NormalizedVertex) Descriptor() ([]byte, []int) {
	return fileDescriptor_geometry_d92abf8dc265831c, []int{1}
}
func (m *NormalizedVertex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NormalizedVertex.Unmarshal(m, b)
}
func (m *NormalizedVertex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NormalizedVertex.Marshal(b, m, deterministic)
}
func (dst *NormalizedVertex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NormalizedVertex.Merge(dst, src)
}
func (m *NormalizedVertex) XXX_Size() int {
	return xxx_messageInfo_NormalizedVertex.Size(m)
}
func (m *NormalizedVertex) XXX_DiscardUnknown() {
	xxx_messageInfo_NormalizedVertex.DiscardUnknown(m)
}

var xxx_messageInfo_NormalizedVertex proto.InternalMessageInfo

func (m *NormalizedVertex) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *NormalizedVertex) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

// A bounding polygon for the detected image annotation.
type BoundingPoly struct {
	// The bounding polygon vertices.
	Vertices []*Vertex `protobuf:"bytes,1,rep,name=vertices,proto3" json:"vertices,omitempty"`
	// The bounding polygon normalized vertices.
	NormalizedVertices   []*NormalizedVertex `protobuf:"bytes,2,rep,name=normalized_vertices,json=normalizedVertices,proto3" json:"normalized_vertices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BoundingPoly) Reset()         { *m = BoundingPoly{} }
func (m *BoundingPoly) String() string { return proto.CompactTextString(m) }
func (*BoundingPoly) ProtoMessage()    {}
func (*BoundingPoly) Descriptor() ([]byte, []int) {
	return fileDescriptor_geometry_d92abf8dc265831c, []int{2}
}
func (m *BoundingPoly) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoundingPoly.Unmarshal(m, b)
}
func (m *BoundingPoly) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoundingPoly.Marshal(b, m, deterministic)
}
func (dst *BoundingPoly) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoundingPoly.Merge(dst, src)
}
func (m *BoundingPoly) XXX_Size() int {
	return xxx_messageInfo_BoundingPoly.Size(m)
}
func (m *BoundingPoly) XXX_DiscardUnknown() {
	xxx_messageInfo_BoundingPoly.DiscardUnknown(m)
}

var xxx_messageInfo_BoundingPoly proto.InternalMessageInfo

func (m *BoundingPoly) GetVertices() []*Vertex {
	if m != nil {
		return m.Vertices
	}
	return nil
}

func (m *BoundingPoly) GetNormalizedVertices() []*NormalizedVertex {
	if m != nil {
		return m.NormalizedVertices
	}
	return nil
}

// A 3D position in the image, used primarily for Face detection landmarks.
// A valid Position must have both x and y coordinates.
// The position coordinates are in the same scale as the original image.
type Position struct {
	// X coordinate.
	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	// Y coordinate.
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	// Z coordinate (or depth).
	Z                    float32  `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_geometry_d92abf8dc265831c, []int{3}
}
func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (dst *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(dst, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Position) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Position) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func init() {
	proto.RegisterType((*Vertex)(nil), "google.cloud.vision.v1p4beta1.Vertex")
	proto.RegisterType((*NormalizedVertex)(nil), "google.cloud.vision.v1p4beta1.NormalizedVertex")
	proto.RegisterType((*BoundingPoly)(nil), "google.cloud.vision.v1p4beta1.BoundingPoly")
	proto.RegisterType((*Position)(nil), "google.cloud.vision.v1p4beta1.Position")
}

func init() {
	proto.RegisterFile("google/cloud/vision/v1p4beta1/geometry.proto", fileDescriptor_geometry_d92abf8dc265831c)
}

var fileDescriptor_geometry_d92abf8dc265831c = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xb1, 0x6a, 0xf3, 0x30,
	0x14, 0x85, 0x91, 0xf3, 0xff, 0x21, 0xa8, 0x29, 0x14, 0x77, 0x31, 0xa5, 0x85, 0xd4, 0xb4, 0x90,
	0xa1, 0x48, 0xa4, 0xcd, 0xd6, 0xa9, 0xce, 0x90, 0x2d, 0x18, 0x0f, 0x19, 0xba, 0xb4, 0x8a, 0x2d,
	0x84, 0xc0, 0xd6, 0x35, 0xb2, 0x62, 0x62, 0xaf, 0x7d, 0x93, 0xbe, 0x42, 0x5f, 0xae, 0x63, 0xb1,
	0x65, 0x0c, 0x09, 0xd4, 0x1d, 0xcf, 0xd5, 0x77, 0x8f, 0x8e, 0xee, 0x15, 0x7e, 0x10, 0x00, 0x22,
	0xe5, 0x34, 0x4e, 0x61, 0x9f, 0xd0, 0x52, 0x16, 0x12, 0x14, 0x2d, 0x17, 0xf9, 0x72, 0xc7, 0x0d,
	0x5b, 0x50, 0xc1, 0x21, 0xe3, 0x46, 0x57, 0x24, 0xd7, 0x60, 0xc0, 0xbd, 0xb1, 0x34, 0x69, 0x69,
	0x62, 0x69, 0xd2, 0xd3, 0x57, 0xd7, 0x9d, 0x19, 0xcb, 0x25, 0x65, 0x4a, 0x81, 0x61, 0x46, 0x82,
	0x2a, 0x6c, 0xb3, 0x7f, 0x87, 0xc7, 0x5b, 0xae, 0x0d, 0x3f, 0xb8, 0x53, 0x8c, 0x0e, 0x1e, 0x9a,
	0xa1, 0xf9, 0xff, 0x08, 0xb5, 0xaa, 0xf2, 0x1c, 0xab, 0x2a, 0x9f, 0xe0, 0x8b, 0x0d, 0xe8, 0x8c,
	0xa5, 0xb2, 0xe6, 0xc9, 0x29, 0xef, 0x1c, 0xf1, 0x4e, 0xc3, 0x7f, 0x21, 0x3c, 0x0d, 0x60, 0xaf,
	0x12, 0xa9, 0x44, 0x08, 0x69, 0xe5, 0xbe, 0xe0, 0x49, 0xc9, 0xb5, 0x91, 0x31, 0x2f, 0x3c, 0x34,
	0x1b, 0xcd, 0xcf, 0x1e, 0xef, 0xc9, 0x60, 0x6c, 0x62, 0x6f, 0x89, 0xfa, 0x36, 0xf7, 0x1d, 0x5f,
	0xaa, 0x3e, 0xc3, 0x5b, 0xef, 0xe6, 0xb4, 0x6e, 0xf4, 0x0f, 0xb7, 0xd3, 0xf4, 0x91, 0xab, 0x8e,
	0x2a, 0x8d, 0x95, 0xbf, 0xc4, 0x93, 0x10, 0x0a, 0xd9, 0x8c, 0x67, 0xe8, 0x75, 0x8d, 0xaa, 0xbd,
	0x91, 0x55, 0x75, 0xf0, 0x81, 0xf0, 0x6d, 0x0c, 0xd9, 0x70, 0x80, 0xe0, 0x7c, 0xdd, 0x2d, 0x2d,
	0x6c, 0xc6, 0x1e, 0xa2, 0xd7, 0x55, 0xc7, 0x0b, 0x48, 0x99, 0x12, 0x04, 0xb4, 0xa0, 0x82, 0xab,
	0x76, 0x29, 0xd4, 0x1e, 0xb1, 0x5c, 0x16, 0xbf, 0x7c, 0x81, 0x67, 0x5b, 0xf8, 0x46, 0xe8, 0xd3,
	0xf9, 0xb7, 0x5e, 0x6d, 0x37, 0xbb, 0x71, 0xdb, 0xf9, 0xf4, 0x13, 0x00, 0x00, 0xff, 0xff, 0xef,
	0xec, 0x6e, 0xa2, 0x3b, 0x02, 0x00, 0x00,
}
