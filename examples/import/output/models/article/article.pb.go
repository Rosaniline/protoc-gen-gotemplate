// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/article.proto

/*
Package article is a generated protocol buffer package.

It is generated from these files:
	proto/article.proto

It has these top-level messages:
	GetArticleRequest
	GetArticleResponse
	Article
*/
package article

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/Rosaniline/protoc-gen-gotemplate/examples/import/output/models/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetArticleRequest struct {
	Getarticle *common.GetArticle `protobuf:"bytes,1,opt,name=getarticle" json:"getarticle,omitempty"`
}

func (m *GetArticleRequest) Reset()                    { *m = GetArticleRequest{} }
func (m *GetArticleRequest) String() string            { return proto.CompactTextString(m) }
func (*GetArticleRequest) ProtoMessage()               {}
func (*GetArticleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetArticleRequest) GetGetarticle() *common.GetArticle {
	if m != nil {
		return m.Getarticle
	}
	return nil
}

type GetArticleResponse struct {
	Article *Article `protobuf:"bytes,1,opt,name=article" json:"article,omitempty"`
	// The generated output should write []*GetArticleResponse_Storage.Storage for this field.
	Storages []*GetArticleResponse_Storage `protobuf:"bytes,2,rep,name=storages" json:"storages,omitempty"`
}

func (m *GetArticleResponse) Reset()                    { *m = GetArticleResponse{} }
func (m *GetArticleResponse) String() string            { return proto.CompactTextString(m) }
func (*GetArticleResponse) ProtoMessage()               {}
func (*GetArticleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetArticleResponse) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

func (m *GetArticleResponse) GetStorages() []*GetArticleResponse_Storage {
	if m != nil {
		return m.Storages
	}
	return nil
}

type GetArticleResponse_Storage struct {
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
}

func (m *GetArticleResponse_Storage) Reset()                    { *m = GetArticleResponse_Storage{} }
func (m *GetArticleResponse_Storage) String() string            { return proto.CompactTextString(m) }
func (*GetArticleResponse_Storage) ProtoMessage()               {}
func (*GetArticleResponse_Storage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *GetArticleResponse_Storage) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type Article struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Article) Reset()                    { *m = Article{} }
func (m *Article) String() string            { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()               {}
func (*Article) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Article) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Article) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GetArticleRequest)(nil), "company.GetArticleRequest")
	proto.RegisterType((*GetArticleResponse)(nil), "company.GetArticleResponse")
	proto.RegisterType((*GetArticleResponse_Storage)(nil), "company.GetArticleResponse.Storage")
	proto.RegisterType((*Article)(nil), "company.Article")
}

func init() { proto.RegisterFile("proto/article.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x6d, 0x15, 0xab, 0xb3, 0x20, 0x3a, 0x82, 0x94, 0x8a, 0xb0, 0xd4, 0xcb, 0x22, 0x18,
	0xa1, 0x1e, 0x3d, 0x88, 0x5e, 0x7a, 0xaf, 0x78, 0xf1, 0x56, 0xd3, 0x61, 0x29, 0x6c, 0x3a, 0xb5,
	0x89, 0x82, 0xff, 0xc6, 0x9f, 0x2a, 0x9b, 0x4c, 0xd7, 0x8a, 0xb2, 0xa7, 0x4e, 0xf3, 0xbe, 0xf7,
	0xf2, 0xc8, 0xc0, 0x69, 0x3f, 0xb0, 0xe3, 0x9b, 0x7a, 0x70, 0xad, 0x5e, 0x91, 0xf2, 0x7f, 0x98,
	0x68, 0x36, 0x7d, 0xdd, 0x7d, 0x66, 0x18, 0x54, 0xcd, 0xc6, 0x70, 0x17, 0xc4, 0xbc, 0x84, 0x93,
	0x92, 0xdc, 0x43, 0x30, 0x54, 0xf4, 0xf6, 0x4e, 0xd6, 0x61, 0x01, 0xb0, 0x24, 0x27, 0x29, 0x69,
	0x34, 0x8f, 0x16, 0xb3, 0x02, 0x95, 0xf8, 0x26, 0xf8, 0x84, 0xca, 0xbf, 0x22, 0xc0, 0x69, 0x92,
	0xed, 0xb9, 0xb3, 0x84, 0x57, 0x90, 0xfc, 0xce, 0x39, 0x56, 0x52, 0x47, 0x8d, 0xe8, 0x08, 0xe0,
	0x3d, 0x1c, 0x58, 0xc7, 0x43, 0xbd, 0x24, 0x9b, 0xc6, 0xf3, 0xdd, 0xc5, 0xac, 0xb8, 0xdc, 0xc0,
	0x7f, 0xa3, 0xd5, 0x53, 0x60, 0xab, 0x8d, 0x29, 0xbb, 0x80, 0x44, 0x0e, 0x11, 0x61, 0x4f, 0x73,
	0x13, 0x2e, 0x3d, 0xac, 0xfc, 0x9c, 0x5f, 0x43, 0x22, 0x19, 0x78, 0x04, 0x71, 0xdb, 0x88, 0x18,
	0xb7, 0xcd, 0x1a, 0xef, 0x6a, 0x43, 0x69, 0x1c, 0xf0, 0xf5, 0x5c, 0x3c, 0x03, 0x48, 0x33, 0xfb,
	0xa1, 0xb1, 0x04, 0xf8, 0xe9, 0x80, 0xd9, 0xbf, 0xc5, 0xfc, 0xeb, 0x65, 0xe7, 0x5b, 0x4a, 0xe7,
	0x3b, 0x8f, 0xe9, 0xcb, 0x99, 0xe1, 0x86, 0x56, 0x76, 0x5c, 0xd3, 0x9d, 0x7c, 0x5f, 0xf7, 0xfd,
	0x4a, 0x6e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x93, 0xb5, 0x4a, 0x95, 0xc6, 0x01, 0x00, 0x00,
}
