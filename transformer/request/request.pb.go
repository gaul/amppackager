// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transformer/request/request.proto

package request // import "github.com/ampproject/amppackager/transformer/request"

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

// This should be kept in sync with HtmlFormat.Code in
// github.com/ampproject/amphtml/validator/validator.proto.
type Request_HtmlFormat int32

const (
	Request_UNKNOWN_CODE Request_HtmlFormat = 0
	Request_AMP          Request_HtmlFormat = 1
	Request_AMP4ADS      Request_HtmlFormat = 2
	Request_AMP4EMAIL    Request_HtmlFormat = 3
	Request_EXPERIMENTAL Request_HtmlFormat = 4
)

var Request_HtmlFormat_name = map[int32]string{
	0: "UNKNOWN_CODE",
	1: "AMP",
	2: "AMP4ADS",
	3: "AMP4EMAIL",
	4: "EXPERIMENTAL",
}

var Request_HtmlFormat_value = map[string]int32{
	"UNKNOWN_CODE": 0,
	"AMP":          1,
	"AMP4ADS":      2,
	"AMP4EMAIL":    3,
	"EXPERIMENTAL": 4,
}

func (x Request_HtmlFormat) String() string {
	return proto.EnumName(Request_HtmlFormat_name, int32(x))
}

func (Request_HtmlFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_762cce2ac5f73405, []int{0, 0}
}

type Request_TransformersConfig int32

const (
	// Execute the default list of transformers. For packager production
	// environments, this should be the config used.
	Request_DEFAULT Request_TransformersConfig = 0
	// Execute none, and simply parse and re-emit. Some normalization will be
	// performed regardless, including, but not limited to:
	// - HTML normalization (e.g. closing all non-void tags).
	// - removal of all comments
	// - lowercase-ing of attribute keys
	// - lexical sort of attribute keys
	// - text is escaped
	//
	// WARNING. THIS IS FOR TESTING PURPOSES ONLY.
	// Use of this setting in a packager production environment could produce
	// invalid transformed AMP when ingested by AMP caches.
	Request_NONE Request_TransformersConfig = 1
	// Execute the minimum needed for verification/validation.
	//
	// WARNING. FOR AMP CACHE USE ONLY.
	// Use of this setting in a packager production environment could produce
	// invalid transformed AMP when ingested by AMP caches.
	Request_VALIDATION Request_TransformersConfig = 2
	// Execute a custom set of transformers.
	//
	// WARNING. THIS IS FOR TESTING PURPOSES ONLY.
	// Use of this setting in a packager production environment could produce
	// invalid transformed AMP when ingested by AMP caches.
	Request_CUSTOM Request_TransformersConfig = 3
)

var Request_TransformersConfig_name = map[int32]string{
	0: "DEFAULT",
	1: "NONE",
	2: "VALIDATION",
	3: "CUSTOM",
}

var Request_TransformersConfig_value = map[string]int32{
	"DEFAULT":    0,
	"NONE":       1,
	"VALIDATION": 2,
	"CUSTOM":     3,
}

func (x Request_TransformersConfig) String() string {
	return proto.EnumName(Request_TransformersConfig_name, int32(x))
}

func (Request_TransformersConfig) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_762cce2ac5f73405, []int{0, 1}
}

// A Request encapsulates input and contextual parameters for use by the
// local transformation library.
type Request struct {
	// The AMP HTML document to transform.
	Html string `protobuf:"bytes,1,opt,name=html,proto3" json:"html,omitempty"`
	// The public URL of the document, i.e. the location that should appear in
	// the browser URL bar.
	DocumentUrl string `protobuf:"bytes,2,opt,name=document_url,json=documentUrl,proto3" json:"document_url,omitempty"`
	// The AMP runtime version.
	Rtv string `protobuf:"bytes,4,opt,name=rtv,proto3" json:"rtv,omitempty"`
	// The CSS contents to inline into the transformed HTML
	Css string `protobuf:"bytes,5,opt,name=css,proto3" json:"css,omitempty"`
	// Transformations are only run if the HTML tag contains the attribute
	// specifying one of the provided formats. If allowed_formats is empty, then
	// all non-experimental AMP formats are allowed.
	AllowedFormats []Request_HtmlFormat       `protobuf:"varint,7,rep,packed,name=allowed_formats,json=allowedFormats,proto3,enum=amp.transform.Request_HtmlFormat" json:"allowed_formats,omitempty"`
	Config         Request_TransformersConfig `protobuf:"varint,6,opt,name=config,proto3,enum=amp.transform.Request_TransformersConfig" json:"config,omitempty"`
	// If config == CUSTOM, this is the list of custom transformers to execute,
	// in the order provided here. Otherwise, this is ignored.
	Transformers         []string `protobuf:"bytes,3,rep,name=transformers,proto3" json:"transformers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_762cce2ac5f73405, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetHtml() string {
	if m != nil {
		return m.Html
	}
	return ""
}

func (m *Request) GetDocumentUrl() string {
	if m != nil {
		return m.DocumentUrl
	}
	return ""
}

func (m *Request) GetRtv() string {
	if m != nil {
		return m.Rtv
	}
	return ""
}

func (m *Request) GetCss() string {
	if m != nil {
		return m.Css
	}
	return ""
}

func (m *Request) GetAllowedFormats() []Request_HtmlFormat {
	if m != nil {
		return m.AllowedFormats
	}
	return nil
}

func (m *Request) GetConfig() Request_TransformersConfig {
	if m != nil {
		return m.Config
	}
	return Request_DEFAULT
}

func (m *Request) GetTransformers() []string {
	if m != nil {
		return m.Transformers
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "amp.transform.Request")
	proto.RegisterEnum("amp.transform.Request_HtmlFormat", Request_HtmlFormat_name, Request_HtmlFormat_value)
	proto.RegisterEnum("amp.transform.Request_TransformersConfig", Request_TransformersConfig_name, Request_TransformersConfig_value)
}

func init() { proto.RegisterFile("transformer/request/request.proto", fileDescriptor_762cce2ac5f73405) }

var fileDescriptor_762cce2ac5f73405 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x5d, 0x73, 0x93, 0x40,
	0x14, 0x2d, 0xd9, 0x08, 0xe6, 0x36, 0xc5, 0x9d, 0xfb, 0xc4, 0x23, 0xe1, 0x09, 0x5f, 0xc8, 0x4c,
	0xd5, 0xf1, 0x79, 0x0d, 0x54, 0x51, 0x3e, 0x32, 0x14, 0xd4, 0xf1, 0x25, 0x43, 0x29, 0x4d, 0xab,
	0x90, 0xc5, 0xdd, 0x45, 0x7f, 0x9a, 0x7f, 0xcf, 0x81, 0xa4, 0xd6, 0x8c, 0xf6, 0x69, 0xcf, 0x3d,
	0xf7, 0x9c, 0xbb, 0x67, 0xee, 0x5c, 0x58, 0x28, 0x51, 0xee, 0xe4, 0x0d, 0x17, 0x6d, 0x2d, 0x96,
	0xa2, 0xfe, 0xde, 0xd7, 0x52, 0xdd, 0xbf, 0x5e, 0x27, 0xb8, 0xe2, 0x78, 0x56, 0xb6, 0x9d, 0xf7,
	0x47, 0xe6, 0xfc, 0x22, 0x60, 0x64, 0x7b, 0x01, 0x22, 0x4c, 0x6f, 0x55, 0xdb, 0x58, 0x9a, 0xad,
	0xb9, 0xb3, 0x6c, 0xc4, 0xb8, 0x80, 0xf9, 0x35, 0xaf, 0xfa, 0xb6, 0xde, 0xa9, 0x4d, 0x2f, 0x1a,
	0x6b, 0x32, 0xf6, 0x4e, 0xef, 0xb9, 0x42, 0x34, 0x48, 0x81, 0x08, 0xf5, 0xc3, 0x9a, 0x8e, 0x9d,
	0x01, 0x0e, 0x4c, 0x25, 0xa5, 0xf5, 0x64, 0xcf, 0x54, 0x52, 0xe2, 0x7b, 0x78, 0x56, 0x36, 0x0d,
	0xff, 0x59, 0x5f, 0x6f, 0x86, 0x6f, 0x4b, 0x25, 0x2d, 0xc3, 0x26, 0xae, 0x79, 0xbe, 0xf0, 0x8e,
	0xf2, 0x78, 0x87, 0x2c, 0xde, 0x3b, 0xd5, 0x36, 0x17, 0xa3, 0x32, 0x33, 0x0f, 0xce, 0x7d, 0x29,
	0x91, 0x81, 0x5e, 0xf1, 0xdd, 0xcd, 0xdd, 0xd6, 0xd2, 0x6d, 0xcd, 0x35, 0xcf, 0x9f, 0x3f, 0x32,
	0x22, 0x7f, 0xd8, 0x85, 0x5c, 0x8d, 0x86, 0xec, 0x60, 0x44, 0x07, 0xe6, 0x7f, 0x6d, 0x4a, 0x5a,
	0xc4, 0x26, 0xee, 0x2c, 0x3b, 0xe2, 0x9c, 0x02, 0xe0, 0x21, 0x04, 0x52, 0x98, 0x17, 0xc9, 0x87,
	0x24, 0xfd, 0x94, 0x6c, 0x56, 0xa9, 0x1f, 0xd0, 0x13, 0x34, 0x80, 0xb0, 0x78, 0x4d, 0x35, 0x3c,
	0x05, 0x83, 0xc5, 0xeb, 0x97, 0xcc, 0xbf, 0xa4, 0x13, 0x3c, 0x83, 0xd9, 0x50, 0x04, 0x31, 0x0b,
	0x23, 0x4a, 0x06, 0x5b, 0xf0, 0x79, 0x1d, 0x64, 0x61, 0x1c, 0x24, 0x39, 0x8b, 0xe8, 0xd4, 0x79,
	0x0b, 0xf8, 0x6f, 0xb0, 0x61, 0x86, 0x1f, 0x5c, 0xb0, 0x22, 0xca, 0xe9, 0x09, 0x3e, 0x85, 0x69,
	0x92, 0x26, 0x01, 0xd5, 0xd0, 0x04, 0xf8, 0xc8, 0xa2, 0xd0, 0x67, 0x79, 0x98, 0x26, 0x74, 0x82,
	0x00, 0xfa, 0xaa, 0xb8, 0xcc, 0xd3, 0x98, 0x92, 0x37, 0xaf, 0xbf, 0xbc, 0xda, 0xde, 0xa9, 0xdb,
	0xfe, 0xca, 0xab, 0x78, 0xbb, 0x2c, 0xdb, 0xae, 0x13, 0xfc, 0x6b, 0x5d, 0xa9, 0x11, 0x96, 0xd5,
	0xb7, 0x72, 0x5b, 0x8b, 0xe5, 0x7f, 0xee, 0xe1, 0x4a, 0x1f, 0x0f, 0xe1, 0xc5, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x8d, 0x9b, 0x92, 0x95, 0x2d, 0x02, 0x00, 0x00,
}