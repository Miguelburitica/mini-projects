// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.15.8
// source: exampb/exam.proto

package exampb

import (
	studentpb "github.com/Miguelburitica/mini-projects/go/protobuffers/studentpb"
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

type Exam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Exam) Reset() {
	*x = Exam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampb_exam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exam) ProtoMessage() {}

func (x *Exam) ProtoReflect() protoreflect.Message {
	mi := &file_exampb_exam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exam.ProtoReflect.Descriptor instead.
func (*Exam) Descriptor() ([]byte, []int) {
	return file_exampb_exam_proto_rawDescGZIP(), []int{0}
}

func (x *Exam) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Exam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Answer   string `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	Question string `protobuf:"bytes,3,opt,name=question,proto3" json:"question,omitempty"`
	ExamId   string `protobuf:"bytes,4,opt,name=exam_id,json=examId,proto3" json:"exam_id,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampb_exam_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_exampb_exam_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_exampb_exam_proto_rawDescGZIP(), []int{1}
}

func (x *Question) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Question) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *Question) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *Question) GetExamId() string {
	if x != nil {
		return x.ExamId
	}
	return ""
}

type GetExamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetExamRequest) Reset() {
	*x = GetExamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampb_exam_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExamRequest) ProtoMessage() {}

func (x *GetExamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exampb_exam_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExamRequest.ProtoReflect.Descriptor instead.
func (*GetExamRequest) Descriptor() ([]byte, []int) {
	return file_exampb_exam_proto_rawDescGZIP(), []int{2}
}

func (x *GetExamRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SetExamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SetExamResponse) Reset() {
	*x = SetExamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampb_exam_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetExamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetExamResponse) ProtoMessage() {}

func (x *SetExamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exampb_exam_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetExamResponse.ProtoReflect.Descriptor instead.
func (*SetExamResponse) Descriptor() ([]byte, []int) {
	return file_exampb_exam_proto_rawDescGZIP(), []int{3}
}

func (x *SetExamResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetExamResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SetQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *SetQuestionResponse) Reset() {
	*x = SetQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampb_exam_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetQuestionResponse) ProtoMessage() {}

func (x *SetQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exampb_exam_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetQuestionResponse.ProtoReflect.Descriptor instead.
func (*SetQuestionResponse) Descriptor() ([]byte, []int) {
	return file_exampb_exam_proto_rawDescGZIP(), []int{4}
}

func (x *SetQuestionResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type EnrrolmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	ExamId    string `protobuf:"bytes,2,opt,name=exam_id,json=examId,proto3" json:"exam_id,omitempty"`
}

func (x *EnrrolmentRequest) Reset() {
	*x = EnrrolmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampb_exam_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrrolmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrrolmentRequest) ProtoMessage() {}

func (x *EnrrolmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exampb_exam_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrrolmentRequest.ProtoReflect.Descriptor instead.
func (*EnrrolmentRequest) Descriptor() ([]byte, []int) {
	return file_exampb_exam_proto_rawDescGZIP(), []int{5}
}

func (x *EnrrolmentRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *EnrrolmentRequest) GetExamId() string {
	if x != nil {
		return x.ExamId
	}
	return ""
}

type GetStudentsPerExamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExamId string `protobuf:"bytes,1,opt,name=exam_id,json=examId,proto3" json:"exam_id,omitempty"`
}

func (x *GetStudentsPerExamRequest) Reset() {
	*x = GetStudentsPerExamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exampb_exam_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentsPerExamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentsPerExamRequest) ProtoMessage() {}

func (x *GetStudentsPerExamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exampb_exam_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentsPerExamRequest.ProtoReflect.Descriptor instead.
func (*GetStudentsPerExamRequest) Descriptor() ([]byte, []int) {
	return file_exampb_exam_proto_rawDescGZIP(), []int{6}
}

func (x *GetStudentsPerExamRequest) GetExamId() string {
	if x != nil {
		return x.ExamId
	}
	return ""
}

var File_exampb_exam_proto protoreflect.FileDescriptor

var file_exampb_exam_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65, 0x78, 0x61, 0x6d, 0x1a, 0x17, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x04, 0x45, 0x78, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x67,
	0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17,
	0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x78, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x78,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x0f, 0x53, 0x65, 0x74,
	0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x25, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x4b, 0x0a, 0x11, 0x45, 0x6e, 0x72, 0x72, 0x6f,
	0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x65,
	0x78, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x78,
	0x61, 0x6d, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x50, 0x65, 0x72, 0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x78, 0x61, 0x6d, 0x49, 0x64, 0x32, 0xb8, 0x02, 0x0a, 0x0b, 0x45,
	0x78, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x45, 0x78, 0x61, 0x6d, 0x12, 0x14, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x12, 0x2c, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x45, 0x78,
	0x61, 0x6d, 0x12, 0x0a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x1a, 0x15,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x53, 0x65, 0x74,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x12, 0x46, 0x0a, 0x0e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x45, 0x6e, 0x72, 0x72,
	0x6f, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x2e, 0x53, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x49, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x65, 0x72, 0x45, 0x78, 0x61, 0x6d,
	0x12, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x50, 0x65, 0x72, 0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x30, 0x01, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x69, 0x67, 0x75, 0x65, 0x6c, 0x62, 0x75, 0x72, 0x69, 0x74, 0x69,
	0x63, 0x61, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exampb_exam_proto_rawDescOnce sync.Once
	file_exampb_exam_proto_rawDescData = file_exampb_exam_proto_rawDesc
)

func file_exampb_exam_proto_rawDescGZIP() []byte {
	file_exampb_exam_proto_rawDescOnce.Do(func() {
		file_exampb_exam_proto_rawDescData = protoimpl.X.CompressGZIP(file_exampb_exam_proto_rawDescData)
	})
	return file_exampb_exam_proto_rawDescData
}

var file_exampb_exam_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_exampb_exam_proto_goTypes = []interface{}{
	(*Exam)(nil),                      // 0: exam.Exam
	(*Question)(nil),                  // 1: exam.Question
	(*GetExamRequest)(nil),            // 2: exam.GetExamRequest
	(*SetExamResponse)(nil),           // 3: exam.SetExamResponse
	(*SetQuestionResponse)(nil),       // 4: exam.SetQuestionResponse
	(*EnrrolmentRequest)(nil),         // 5: exam.EnrrolmentRequest
	(*GetStudentsPerExamRequest)(nil), // 6: exam.GetStudentsPerExamRequest
	(*studentpb.Student)(nil),         // 7: student.Student
}
var file_exampb_exam_proto_depIdxs = []int32{
	2, // 0: exam.ExamService.GetExam:input_type -> exam.GetExamRequest
	0, // 1: exam.ExamService.SetExam:input_type -> exam.Exam
	1, // 2: exam.ExamService.SetQuestions:input_type -> exam.Question
	5, // 3: exam.ExamService.EnrollStudents:input_type -> exam.EnrrolmentRequest
	6, // 4: exam.ExamService.GetStudentsPerExam:input_type -> exam.GetStudentsPerExamRequest
	0, // 5: exam.ExamService.GetExam:output_type -> exam.Exam
	3, // 6: exam.ExamService.SetExam:output_type -> exam.SetExamResponse
	4, // 7: exam.ExamService.SetQuestions:output_type -> exam.SetQuestionResponse
	4, // 8: exam.ExamService.EnrollStudents:output_type -> exam.SetQuestionResponse
	7, // 9: exam.ExamService.GetStudentsPerExam:output_type -> student.Student
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_exampb_exam_proto_init() }
func file_exampb_exam_proto_init() {
	if File_exampb_exam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exampb_exam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Exam); i {
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
		file_exampb_exam_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
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
		file_exampb_exam_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExamRequest); i {
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
		file_exampb_exam_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetExamResponse); i {
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
		file_exampb_exam_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetQuestionResponse); i {
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
		file_exampb_exam_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrrolmentRequest); i {
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
		file_exampb_exam_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentsPerExamRequest); i {
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
			RawDescriptor: file_exampb_exam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exampb_exam_proto_goTypes,
		DependencyIndexes: file_exampb_exam_proto_depIdxs,
		MessageInfos:      file_exampb_exam_proto_msgTypes,
	}.Build()
	File_exampb_exam_proto = out.File
	file_exampb_exam_proto_rawDesc = nil
	file_exampb_exam_proto_goTypes = nil
	file_exampb_exam_proto_depIdxs = nil
}
