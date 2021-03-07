// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: submitterpb/service.proto

package submitterpb

import (
	base "github.com/ecnuvj/vhoj_submitter/pkg/sdk/base"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ListSubmissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	ProblemId uint64 `protobuf:"varint,2,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Result    int32  `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`
	Language  int32  `protobuf:"varint,4,opt,name=language,proto3" json:"language,omitempty"`
	PageNo    int32  `protobuf:"varint,5,opt,name=page_no,json=pageNo,proto3" json:"page_no,omitempty"`
	PageSize  int32  `protobuf:"varint,6,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListSubmissionsRequest) Reset() {
	*x = ListSubmissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submitterpb_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubmissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubmissionsRequest) ProtoMessage() {}

func (x *ListSubmissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submitterpb_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubmissionsRequest.ProtoReflect.Descriptor instead.
func (*ListSubmissionsRequest) Descriptor() ([]byte, []int) {
	return file_submitterpb_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListSubmissionsRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ListSubmissionsRequest) GetProblemId() uint64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

func (x *ListSubmissionsRequest) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *ListSubmissionsRequest) GetLanguage() int32 {
	if x != nil {
		return x.Language
	}
	return 0
}

func (x *ListSubmissionsRequest) GetPageNo() int32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *ListSubmissionsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListSubmissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Submissions  []*Submission      `protobuf:"bytes,1,rep,name=submissions,proto3" json:"submissions,omitempty"`
	TotalPages   int32              `protobuf:"varint,2,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	TotalCount   int32              `protobuf:"varint,3,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	BaseResponse *base.BaseResponse `protobuf:"bytes,4,opt,name=base_response,json=baseResponse,proto3" json:"base_response,omitempty"`
}

func (x *ListSubmissionsResponse) Reset() {
	*x = ListSubmissionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submitterpb_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubmissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubmissionsResponse) ProtoMessage() {}

func (x *ListSubmissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submitterpb_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubmissionsResponse.ProtoReflect.Descriptor instead.
func (*ListSubmissionsResponse) Descriptor() ([]byte, []int) {
	return file_submitterpb_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListSubmissionsResponse) GetSubmissions() []*Submission {
	if x != nil {
		return x.Submissions
	}
	return nil
}

func (x *ListSubmissionsResponse) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *ListSubmissionsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListSubmissionsResponse) GetBaseResponse() *base.BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

type SubmitCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId  uint64 `protobuf:"varint,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	UserId     uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username   string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Language   int32  `protobuf:"varint,4,opt,name=language,proto3" json:"language,omitempty"`
	ContestId  uint64 `protobuf:"varint,5,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	SourceCode string `protobuf:"bytes,6,opt,name=source_code,json=sourceCode,proto3" json:"source_code,omitempty"`
}

func (x *SubmitCodeRequest) Reset() {
	*x = SubmitCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submitterpb_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitCodeRequest) ProtoMessage() {}

func (x *SubmitCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submitterpb_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitCodeRequest.ProtoReflect.Descriptor instead.
func (*SubmitCodeRequest) Descriptor() ([]byte, []int) {
	return file_submitterpb_service_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitCodeRequest) GetProblemId() uint64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

func (x *SubmitCodeRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SubmitCodeRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SubmitCodeRequest) GetLanguage() int32 {
	if x != nil {
		return x.Language
	}
	return 0
}

func (x *SubmitCodeRequest) GetContestId() uint64 {
	if x != nil {
		return x.ContestId
	}
	return 0
}

func (x *SubmitCodeRequest) GetSourceCode() string {
	if x != nil {
		return x.SourceCode
	}
	return ""
}

type SubmitCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubmissionId uint64             `protobuf:"varint,1,opt,name=submission_id,json=submissionId,proto3" json:"submission_id,omitempty"`
	BaseResponse *base.BaseResponse `protobuf:"bytes,2,opt,name=base_response,json=baseResponse,proto3" json:"base_response,omitempty"`
}

func (x *SubmitCodeResponse) Reset() {
	*x = SubmitCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submitterpb_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitCodeResponse) ProtoMessage() {}

func (x *SubmitCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submitterpb_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitCodeResponse.ProtoReflect.Descriptor instead.
func (*SubmitCodeResponse) Descriptor() ([]byte, []int) {
	return file_submitterpb_service_proto_rawDescGZIP(), []int{3}
}

func (x *SubmitCodeResponse) GetSubmissionId() uint64 {
	if x != nil {
		return x.SubmissionId
	}
	return 0
}

func (x *SubmitCodeResponse) GetBaseResponse() *base.BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

type ReSubmitCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubmissionId uint64 `protobuf:"varint,1,opt,name=submission_id,json=submissionId,proto3" json:"submission_id,omitempty"`
}

func (x *ReSubmitCodeRequest) Reset() {
	*x = ReSubmitCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submitterpb_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReSubmitCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReSubmitCodeRequest) ProtoMessage() {}

func (x *ReSubmitCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submitterpb_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReSubmitCodeRequest.ProtoReflect.Descriptor instead.
func (*ReSubmitCodeRequest) Descriptor() ([]byte, []int) {
	return file_submitterpb_service_proto_rawDescGZIP(), []int{4}
}

func (x *ReSubmitCodeRequest) GetSubmissionId() uint64 {
	if x != nil {
		return x.SubmissionId
	}
	return 0
}

type ReSubmitCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *base.BaseResponse `protobuf:"bytes,1,opt,name=base_response,json=baseResponse,proto3" json:"base_response,omitempty"`
}

func (x *ReSubmitCodeResponse) Reset() {
	*x = ReSubmitCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submitterpb_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReSubmitCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReSubmitCodeResponse) ProtoMessage() {}

func (x *ReSubmitCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submitterpb_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReSubmitCodeResponse.ProtoReflect.Descriptor instead.
func (*ReSubmitCodeResponse) Descriptor() ([]byte, []int) {
	return file_submitterpb_service_proto_rawDescGZIP(), []int{5}
}

func (x *ReSubmitCodeResponse) GetBaseResponse() *base.BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

var File_submitterpb_service_proto protoreflect.FileDescriptor

var file_submitterpb_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x64, 0x6b,
	0x1a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbd, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0xc7, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x37, 0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x11, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x72, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x73, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0d, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x3a, 0x0a, 0x13, 0x52, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x4f, 0x0a, 0x14, 0x52, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xe7, 0x01, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x52, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x52, 0x65, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x52, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x63, 0x6e, 0x75, 0x76, 0x6a, 0x2f,
	0x76, 0x68, 0x6f, 0x6a, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_submitterpb_service_proto_rawDescOnce sync.Once
	file_submitterpb_service_proto_rawDescData = file_submitterpb_service_proto_rawDesc
)

func file_submitterpb_service_proto_rawDescGZIP() []byte {
	file_submitterpb_service_proto_rawDescOnce.Do(func() {
		file_submitterpb_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_submitterpb_service_proto_rawDescData)
	})
	return file_submitterpb_service_proto_rawDescData
}

var file_submitterpb_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_submitterpb_service_proto_goTypes = []interface{}{
	(*ListSubmissionsRequest)(nil),  // 0: sdk.ListSubmissionsRequest
	(*ListSubmissionsResponse)(nil), // 1: sdk.ListSubmissionsResponse
	(*SubmitCodeRequest)(nil),       // 2: sdk.SubmitCodeRequest
	(*SubmitCodeResponse)(nil),      // 3: sdk.SubmitCodeResponse
	(*ReSubmitCodeRequest)(nil),     // 4: sdk.ReSubmitCodeRequest
	(*ReSubmitCodeResponse)(nil),    // 5: sdk.ReSubmitCodeResponse
	(*Submission)(nil),              // 6: sdk.Submission
	(*base.BaseResponse)(nil),       // 7: base.BaseResponse
}
var file_submitterpb_service_proto_depIdxs = []int32{
	6, // 0: sdk.ListSubmissionsResponse.submissions:type_name -> sdk.Submission
	7, // 1: sdk.ListSubmissionsResponse.base_response:type_name -> base.BaseResponse
	7, // 2: sdk.SubmitCodeResponse.base_response:type_name -> base.BaseResponse
	7, // 3: sdk.ReSubmitCodeResponse.base_response:type_name -> base.BaseResponse
	2, // 4: sdk.SubmitService.SubmitCode:input_type -> sdk.SubmitCodeRequest
	4, // 5: sdk.SubmitService.ReSubmitCode:input_type -> sdk.ReSubmitCodeRequest
	0, // 6: sdk.SubmitService.ListSubmissions:input_type -> sdk.ListSubmissionsRequest
	3, // 7: sdk.SubmitService.SubmitCode:output_type -> sdk.SubmitCodeResponse
	5, // 8: sdk.SubmitService.ReSubmitCode:output_type -> sdk.ReSubmitCodeResponse
	1, // 9: sdk.SubmitService.ListSubmissions:output_type -> sdk.ListSubmissionsResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_submitterpb_service_proto_init() }
func file_submitterpb_service_proto_init() {
	if File_submitterpb_service_proto != nil {
		return
	}
	file_submitterpb_submission_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_submitterpb_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubmissionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_submitterpb_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubmissionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_submitterpb_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitCodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_submitterpb_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitCodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_submitterpb_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReSubmitCodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_submitterpb_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReSubmitCodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_submitterpb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_submitterpb_service_proto_goTypes,
		DependencyIndexes: file_submitterpb_service_proto_depIdxs,
		MessageInfos:      file_submitterpb_service_proto_msgTypes,
	}.Build()
	File_submitterpb_service_proto = out.File
	file_submitterpb_service_proto_rawDesc = nil
	file_submitterpb_service_proto_goTypes = nil
	file_submitterpb_service_proto_depIdxs = nil
}
