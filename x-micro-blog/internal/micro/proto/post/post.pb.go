// Code generated by protoc-gen-go. DO NOT EDIT.
// source: post.proto

package post

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

type Post struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	AuthorId             string   `protobuf:"bytes,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Tags                 []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Post) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Post) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Post) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Post) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type CreatePostRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	AuthorId             string   `protobuf:"bytes,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Tags                 []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostRequest) Reset()         { *m = CreatePostRequest{} }
func (m *CreatePostRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePostRequest) ProtoMessage()    {}
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{1}
}

func (m *CreatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostRequest.Unmarshal(m, b)
}
func (m *CreatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostRequest.Marshal(b, m, deterministic)
}
func (m *CreatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostRequest.Merge(m, src)
}
func (m *CreatePostRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePostRequest.Size(m)
}
func (m *CreatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostRequest proto.InternalMessageInfo

func (m *CreatePostRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreatePostRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CreatePostRequest) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *CreatePostRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type CreatePostResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Post                 *Post    `protobuf:"bytes,3,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostResponse) Reset()         { *m = CreatePostResponse{} }
func (m *CreatePostResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePostResponse) ProtoMessage()    {}
func (*CreatePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{2}
}

func (m *CreatePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostResponse.Unmarshal(m, b)
}
func (m *CreatePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostResponse.Marshal(b, m, deterministic)
}
func (m *CreatePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostResponse.Merge(m, src)
}
func (m *CreatePostResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePostResponse.Size(m)
}
func (m *CreatePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostResponse proto.InternalMessageInfo

func (m *CreatePostResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CreatePostResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreatePostResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type GetPostRequest struct {
	PostId               string   `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPostRequest) Reset()         { *m = GetPostRequest{} }
func (m *GetPostRequest) String() string { return proto.CompactTextString(m) }
func (*GetPostRequest) ProtoMessage()    {}
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{3}
}

func (m *GetPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPostRequest.Unmarshal(m, b)
}
func (m *GetPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPostRequest.Marshal(b, m, deterministic)
}
func (m *GetPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPostRequest.Merge(m, src)
}
func (m *GetPostRequest) XXX_Size() int {
	return xxx_messageInfo_GetPostRequest.Size(m)
}
func (m *GetPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPostRequest proto.InternalMessageInfo

func (m *GetPostRequest) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

type GetPostResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Post                 *Post    `protobuf:"bytes,3,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPostResponse) Reset()         { *m = GetPostResponse{} }
func (m *GetPostResponse) String() string { return proto.CompactTextString(m) }
func (*GetPostResponse) ProtoMessage()    {}
func (*GetPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{4}
}

func (m *GetPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPostResponse.Unmarshal(m, b)
}
func (m *GetPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPostResponse.Marshal(b, m, deterministic)
}
func (m *GetPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPostResponse.Merge(m, src)
}
func (m *GetPostResponse) XXX_Size() int {
	return xxx_messageInfo_GetPostResponse.Size(m)
}
func (m *GetPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPostResponse proto.InternalMessageInfo

func (m *GetPostResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *GetPostResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetPostResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type ListPostsRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	AuthorId             string   `protobuf:"bytes,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Tag                  string   `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostsRequest) Reset()         { *m = ListPostsRequest{} }
func (m *ListPostsRequest) String() string { return proto.CompactTextString(m) }
func (*ListPostsRequest) ProtoMessage()    {}
func (*ListPostsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{5}
}

func (m *ListPostsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostsRequest.Unmarshal(m, b)
}
func (m *ListPostsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostsRequest.Marshal(b, m, deterministic)
}
func (m *ListPostsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostsRequest.Merge(m, src)
}
func (m *ListPostsRequest) XXX_Size() int {
	return xxx_messageInfo_ListPostsRequest.Size(m)
}
func (m *ListPostsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostsRequest proto.InternalMessageInfo

func (m *ListPostsRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListPostsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListPostsRequest) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *ListPostsRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type ListPostsResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Posts                []*Post  `protobuf:"bytes,3,rep,name=posts,proto3" json:"posts,omitempty"`
	Total                int32    `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostsResponse) Reset()         { *m = ListPostsResponse{} }
func (m *ListPostsResponse) String() string { return proto.CompactTextString(m) }
func (*ListPostsResponse) ProtoMessage()    {}
func (*ListPostsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{6}
}

func (m *ListPostsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostsResponse.Unmarshal(m, b)
}
func (m *ListPostsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostsResponse.Marshal(b, m, deterministic)
}
func (m *ListPostsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostsResponse.Merge(m, src)
}
func (m *ListPostsResponse) XXX_Size() int {
	return xxx_messageInfo_ListPostsResponse.Size(m)
}
func (m *ListPostsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostsResponse proto.InternalMessageInfo

func (m *ListPostsResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ListPostsResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ListPostsResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

func (m *ListPostsResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type UpdatePostRequest struct {
	PostId               string   `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Tags                 []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePostRequest) Reset()         { *m = UpdatePostRequest{} }
func (m *UpdatePostRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePostRequest) ProtoMessage()    {}
func (*UpdatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{7}
}

func (m *UpdatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePostRequest.Unmarshal(m, b)
}
func (m *UpdatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePostRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePostRequest.Merge(m, src)
}
func (m *UpdatePostRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePostRequest.Size(m)
}
func (m *UpdatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePostRequest proto.InternalMessageInfo

func (m *UpdatePostRequest) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

func (m *UpdatePostRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdatePostRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *UpdatePostRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type UpdatePostResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Post                 *Post    `protobuf:"bytes,3,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePostResponse) Reset()         { *m = UpdatePostResponse{} }
func (m *UpdatePostResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePostResponse) ProtoMessage()    {}
func (*UpdatePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{8}
}

func (m *UpdatePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePostResponse.Unmarshal(m, b)
}
func (m *UpdatePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePostResponse.Marshal(b, m, deterministic)
}
func (m *UpdatePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePostResponse.Merge(m, src)
}
func (m *UpdatePostResponse) XXX_Size() int {
	return xxx_messageInfo_UpdatePostResponse.Size(m)
}
func (m *UpdatePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePostResponse proto.InternalMessageInfo

func (m *UpdatePostResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *UpdatePostResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UpdatePostResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type DeletePostRequest struct {
	PostId               string   `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostRequest) Reset()         { *m = DeletePostRequest{} }
func (m *DeletePostRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePostRequest) ProtoMessage()    {}
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{9}
}

func (m *DeletePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostRequest.Unmarshal(m, b)
}
func (m *DeletePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostRequest.Marshal(b, m, deterministic)
}
func (m *DeletePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostRequest.Merge(m, src)
}
func (m *DeletePostRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePostRequest.Size(m)
}
func (m *DeletePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostRequest proto.InternalMessageInfo

func (m *DeletePostRequest) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

type DeletePostResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostResponse) Reset()         { *m = DeletePostResponse{} }
func (m *DeletePostResponse) String() string { return proto.CompactTextString(m) }
func (*DeletePostResponse) ProtoMessage()    {}
func (*DeletePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{10}
}

func (m *DeletePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostResponse.Unmarshal(m, b)
}
func (m *DeletePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostResponse.Marshal(b, m, deterministic)
}
func (m *DeletePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostResponse.Merge(m, src)
}
func (m *DeletePostResponse) XXX_Size() int {
	return xxx_messageInfo_DeletePostResponse.Size(m)
}
func (m *DeletePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostResponse proto.InternalMessageInfo

func (m *DeletePostResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DeletePostResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Post)(nil), "post.Post")
	proto.RegisterType((*CreatePostRequest)(nil), "post.CreatePostRequest")
	proto.RegisterType((*CreatePostResponse)(nil), "post.CreatePostResponse")
	proto.RegisterType((*GetPostRequest)(nil), "post.GetPostRequest")
	proto.RegisterType((*GetPostResponse)(nil), "post.GetPostResponse")
	proto.RegisterType((*ListPostsRequest)(nil), "post.ListPostsRequest")
	proto.RegisterType((*ListPostsResponse)(nil), "post.ListPostsResponse")
	proto.RegisterType((*UpdatePostRequest)(nil), "post.UpdatePostRequest")
	proto.RegisterType((*UpdatePostResponse)(nil), "post.UpdatePostResponse")
	proto.RegisterType((*DeletePostRequest)(nil), "post.DeletePostRequest")
	proto.RegisterType((*DeletePostResponse)(nil), "post.DeletePostResponse")
}

func init() { proto.RegisterFile("post.proto", fileDescriptor_e114ad14deab1dd1) }

var fileDescriptor_e114ad14deab1dd1 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x6e, 0xfc, 0x91, 0xc4, 0xd3, 0x57, 0x79, 0x9b, 0x51, 0x21, 0xab, 0x20, 0x50, 0xb4, 0xa7,
	0x22, 0xa1, 0x54, 0x2a, 0x17, 0x4e, 0x48, 0x01, 0x24, 0xa8, 0xc4, 0x01, 0xb9, 0xe2, 0xc2, 0x25,
	0x32, 0xf6, 0x2a, 0xb5, 0x14, 0x62, 0xe3, 0x9d, 0x70, 0xe8, 0x85, 0x5f, 0xc5, 0xff, 0xe0, 0x27,
	0xa1, 0x9d, 0xcd, 0xda, 0x26, 0x56, 0xaa, 0x02, 0xea, 0xc9, 0xf3, 0xb5, 0xf3, 0x3c, 0xf3, 0x65,
	0x80, 0xb2, 0xd0, 0x34, 0x2f, 0xab, 0x82, 0x0a, 0x0c, 0x8c, 0x2c, 0x7f, 0xf4, 0x20, 0xf8, 0x50,
	0x68, 0xc2, 0x11, 0x78, 0x79, 0x26, 0x7a, 0xb3, 0xde, 0x59, 0x14, 0x7b, 0x79, 0x86, 0xa7, 0x10,
	0x52, 0x4e, 0x6b, 0x25, 0x3c, 0x36, 0x59, 0x05, 0x05, 0x0c, 0xd2, 0x62, 0x43, 0x6a, 0x43, 0xc2,
	0x67, 0xbb, 0x53, 0xf1, 0x11, 0x44, 0xc9, 0x96, 0xae, 0x8b, 0x6a, 0x99, 0x67, 0x22, 0x60, 0xdf,
	0xd0, 0x1a, 0x2e, 0x33, 0x44, 0x08, 0x28, 0x59, 0x69, 0x11, 0xce, 0xfc, 0xb3, 0x28, 0x66, 0x19,
	0x1f, 0x03, 0xa4, 0x95, 0x4a, 0x48, 0x65, 0xcb, 0x84, 0x44, 0x9f, 0x5f, 0x44, 0x3b, 0xcb, 0x82,
	0x8c, 0x7b, 0x5b, 0x66, 0xce, 0x3d, 0xb0, 0xee, 0x9d, 0x65, 0x41, 0x92, 0x60, 0xfc, 0x9a, 0x63,
	0x0d, 0xf9, 0x58, 0x7d, 0xdd, 0x2a, 0x4d, 0x0d, 0xe7, 0xde, 0x01, 0xce, 0xde, 0x2d, 0x9c, 0xfd,
	0x03, 0x9c, 0x83, 0x86, 0xb3, 0xbc, 0x06, 0x6c, 0xa3, 0xea, 0xb2, 0xd8, 0x68, 0x06, 0xd0, 0xdb,
	0x34, 0x55, 0x5a, 0x33, 0xf0, 0x30, 0x76, 0xaa, 0xf1, 0x7c, 0x51, 0x5a, 0x27, 0x2b, 0xd7, 0x46,
	0xa7, 0xe2, 0x13, 0xe0, 0xfe, 0x33, 0xea, 0xf1, 0x05, 0xcc, 0x79, 0x30, 0x9c, 0xd5, 0xce, 0xe5,
	0x29, 0x8c, 0xde, 0x2a, 0x6a, 0x17, 0x37, 0x81, 0x81, 0xf1, 0x2c, 0xeb, 0x29, 0xf5, 0x8d, 0x7a,
	0x99, 0x49, 0x05, 0xff, 0xd7, 0xa1, 0xf7, 0xc8, 0xa8, 0x82, 0x93, 0xf7, 0xb9, 0x66, 0x1c, 0xed,
	0x38, 0x21, 0x04, 0xa5, 0x49, 0x65, 0x40, 0xc2, 0x98, 0x65, 0xd3, 0x54, 0xf3, 0x5d, 0xea, 0xfc,
	0xc6, 0x62, 0x84, 0xf1, 0xd0, 0x18, 0xae, 0xf2, 0x1b, 0x75, 0x7b, 0xc7, 0x4f, 0xc0, 0xa7, 0x64,
	0xb5, 0x5b, 0x1e, 0x23, 0xca, 0xef, 0x30, 0x6e, 0x61, 0xfe, 0x43, 0x71, 0x33, 0x08, 0x4d, 0x11,
	0x5a, 0xf8, 0x33, 0x7f, 0xaf, 0x3a, 0xeb, 0xe0, 0xdd, 0x29, 0x28, 0x59, 0x33, 0x7c, 0x18, 0x5b,
	0x45, 0x96, 0x30, 0xfe, 0xc8, 0x3b, 0x77, 0x97, 0x49, 0xfc, 0xf1, 0xcd, 0x1c, 0x58, 0xb1, 0x36,
	0xe2, 0x3d, 0x0e, 0xf4, 0x19, 0x8c, 0xdf, 0xa8, 0xb5, 0xba, 0x5b, 0x6d, 0xf2, 0x1d, 0x60, 0x3b,
	0xfa, 0xef, 0x79, 0x5d, 0xfc, 0xf4, 0xe0, 0xd8, 0x24, 0xb9, 0x52, 0xd5, 0xb7, 0x3c, 0x55, 0xb8,
	0x00, 0x68, 0x8e, 0x0a, 0x27, 0x96, 0x67, 0xe7, 0xb8, 0xa7, 0xa2, 0xeb, 0xb0, 0x24, 0xe4, 0x11,
	0xbe, 0x80, 0xc1, 0xee, 0x04, 0xf0, 0xd4, 0x86, 0xfd, 0x7e, 0x3c, 0xd3, 0x07, 0x7b, 0xd6, 0xfa,
	0xe5, 0x4b, 0x88, 0xea, 0x0d, 0xc3, 0x87, 0x36, 0x6a, 0x7f, 0xcd, 0xa7, 0x93, 0x8e, 0xbd, 0x7e,
	0xbf, 0x00, 0x68, 0xc6, 0xe5, 0xc8, 0x77, 0x56, 0xc6, 0x91, 0xef, 0x4e, 0xd6, 0xa6, 0x68, 0x3a,
	0xeb, 0x52, 0x74, 0x26, 0xe3, 0x52, 0x74, 0x87, 0x20, 0x8f, 0x5e, 0x8d, 0x3e, 0xfd, 0x37, 0x3f,
	0xe7, 0xdf, 0xfa, 0xb9, 0x09, 0xfa, 0xdc, 0x67, 0xf9, 0xf9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x3d, 0xfc, 0x86, 0x1d, 0xf0, 0x05, 0x00, 0x00,
}