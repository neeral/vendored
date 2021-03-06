// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lightstep_carrier.proto

/*
Package lightsteppb is a generated protocol buffer package.

It is generated from these files:
	lightstep_carrier.proto

It has these top-level messages:
	BinaryCarrier
	BasicTracerCarrier
*/
package lightsteppb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The standard carrier for binary context propagation into LightStep.
type BinaryCarrier struct {
	// "text_ctx" was deprecated following lightstep-tracer-cpp-0.36
	DeprecatedTextCtx [][]byte `protobuf:"bytes,1,rep,name=deprecated_text_ctx,json=deprecatedTextCtx,proto3" json:"deprecated_text_ctx,omitempty"`
	// The Opentracing "basictracer" proto.
	BasicCtx *BasicTracerCarrier `protobuf:"bytes,2,opt,name=basic_ctx,json=basicCtx" json:"basic_ctx,omitempty"`
}

func (m *BinaryCarrier) Reset()                    { *m = BinaryCarrier{} }
func (m *BinaryCarrier) String() string            { return proto.CompactTextString(m) }
func (*BinaryCarrier) ProtoMessage()               {}
func (*BinaryCarrier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BinaryCarrier) GetDeprecatedTextCtx() [][]byte {
	if m != nil {
		return m.DeprecatedTextCtx
	}
	return nil
}

func (m *BinaryCarrier) GetBasicCtx() *BasicTracerCarrier {
	if m != nil {
		return m.BasicCtx
	}
	return nil
}

// Copy of https://github.com/opentracing/basictracer-go/blob/master/wire/wire.proto
type BasicTracerCarrier struct {
	TraceId      uint64            `protobuf:"fixed64,1,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	SpanId       uint64            `protobuf:"fixed64,2,opt,name=span_id,json=spanId" json:"span_id,omitempty"`
	Sampled      bool              `protobuf:"varint,3,opt,name=sampled" json:"sampled,omitempty"`
	BaggageItems map[string]string `protobuf:"bytes,4,rep,name=baggage_items,json=baggageItems" json:"baggage_items,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *BasicTracerCarrier) Reset()                    { *m = BasicTracerCarrier{} }
func (m *BasicTracerCarrier) String() string            { return proto.CompactTextString(m) }
func (*BasicTracerCarrier) ProtoMessage()               {}
func (*BasicTracerCarrier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BasicTracerCarrier) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

func (m *BasicTracerCarrier) GetSpanId() uint64 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *BasicTracerCarrier) GetSampled() bool {
	if m != nil {
		return m.Sampled
	}
	return false
}

func (m *BasicTracerCarrier) GetBaggageItems() map[string]string {
	if m != nil {
		return m.BaggageItems
	}
	return nil
}

func init() {
	proto.RegisterType((*BinaryCarrier)(nil), "lightstep.BinaryCarrier")
	proto.RegisterType((*BasicTracerCarrier)(nil), "lightstep.BasicTracerCarrier")
}

func init() { proto.RegisterFile("lightstep_carrier.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xc9, 0xaa, 0xdb, 0xfa, 0xb6, 0x81, 0x8b, 0xc2, 0xaa, 0x20, 0x94, 0x9d, 0x7a, 0xaa,
	0x30, 0x2f, 0xb2, 0x8b, 0xd0, 0xe1, 0x61, 0xd7, 0xb0, 0x93, 0x97, 0x92, 0x36, 0x8f, 0x1a, 0xdc,
	0x8f, 0x90, 0x3e, 0xa5, 0xc3, 0xbf, 0xdc, 0x9b, 0x24, 0xd5, 0x2a, 0x0c, 0xbc, 0xe5, 0xfb, 0x3e,
	0x9f, 0x97, 0x7c, 0x21, 0x30, 0xdb, 0xea, 0xea, 0x85, 0x6a, 0x42, 0x93, 0x97, 0xd2, 0x5a, 0x8d,
	0x36, 0x35, 0xf6, 0x40, 0x07, 0x1e, 0x76, 0x60, 0xfe, 0x01, 0x93, 0x4c, 0xef, 0xa5, 0x3d, 0xae,
	0x5a, 0x83, 0xa7, 0x70, 0xa9, 0xd0, 0x58, 0x2c, 0x25, 0xa1, 0xca, 0x09, 0x1b, 0xca, 0x4b, 0x6a,
	0x22, 0x16, 0x07, 0xc9, 0x58, 0x4c, 0x7f, 0xd1, 0x06, 0x1b, 0x5a, 0x51, 0xc3, 0x97, 0x10, 0x16,
	0xb2, 0xd6, 0xa5, 0xb7, 0x7a, 0x31, 0x4b, 0x46, 0x8b, 0xdb, 0xb4, 0xbb, 0x3f, 0xcd, 0x1c, 0xdb,
	0x58, 0x59, 0xa2, 0xfd, 0x7e, 0x41, 0x0c, 0xbd, 0xbf, 0xa2, 0x66, 0xfe, 0xc9, 0x80, 0x9f, 0x0a,
	0xfc, 0x1a, 0x86, 0xe4, 0x06, 0xb9, 0x56, 0x11, 0x8b, 0x59, 0xd2, 0x17, 0x03, 0x9f, 0xd7, 0x8a,
	0xcf, 0x60, 0x50, 0x1b, 0xb9, 0x77, 0xa4, 0xe7, 0x49, 0xdf, 0xc5, 0xb5, 0xe2, 0x11, 0x0c, 0x6a,
	0xb9, 0x33, 0x5b, 0x54, 0x51, 0x10, 0xb3, 0x64, 0x28, 0x7e, 0x22, 0xdf, 0xc0, 0xa4, 0x90, 0x55,
	0x25, 0x2b, 0xcc, 0x35, 0xe1, 0xae, 0x8e, 0xce, 0xe2, 0x20, 0x19, 0x2d, 0xee, 0xfe, 0x2d, 0x99,
	0x66, 0xed, 0xca, 0xda, 0x6d, 0x3c, 0xed, 0xc9, 0x1e, 0xc5, 0xb8, 0xf8, 0x33, 0xba, 0x79, 0x84,
	0xe9, 0x89, 0xc2, 0x2f, 0x20, 0x78, 0xc5, 0xa3, 0xef, 0x1c, 0x0a, 0x77, 0xe4, 0x57, 0x70, 0xfe,
	0x2e, 0xb7, 0x6f, 0xe8, 0xdb, 0x86, 0xa2, 0x0d, 0xcb, 0xde, 0x03, 0xcb, 0x26, 0xcf, 0xa3, 0xae,
	0x80, 0x29, 0x8a, 0xbe, 0xff, 0x99, 0xfb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0x62, 0x9a,
	0xfe, 0xb4, 0x01, 0x00, 0x00,
}
