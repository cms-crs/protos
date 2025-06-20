// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: task_service.proto

package taskv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ListId        string                 `protobuf:"bytes,2,opt,name=list_id,json=listId,proto3" json:"list_id,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Position      int32                  `protobuf:"varint,5,opt,name=position,proto3" json:"position,omitempty"`
	DueDate       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
	CreatedBy     string                 `protobuf:"bytes,7,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	AssignedUsers []string               `protobuf:"bytes,10,rep,name=assigned_users,json=assignedUsers,proto3" json:"assigned_users,omitempty"`
	Labels        []*Label               `protobuf:"bytes,11,rep,name=labels,proto3" json:"labels,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_task_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetListId() string {
	if x != nil {
		return x.ListId
	}
	return ""
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *Task) GetDueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DueDate
	}
	return nil
}

func (x *Task) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Task) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Task) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Task) GetAssignedUsers() []string {
	if x != nil {
		return x.AssignedUsers
	}
	return nil
}

func (x *Task) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

type Label struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BoardId       string                 `protobuf:"bytes,2,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Color         string                 `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Label) Reset() {
	*x = Label{}
	mi := &file_task_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{1}
}

func (x *Label) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Label) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

func (x *Label) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Label) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type CreateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ListId        string                 `protobuf:"bytes,1,opt,name=list_id,json=listId,proto3" json:"list_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Position      int32                  `protobuf:"varint,4,opt,name=position,proto3" json:"position,omitempty"`
	DueDate       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
	CreatedBy     string                 `protobuf:"bytes,6,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	mi := &file_task_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTaskRequest) GetListId() string {
	if x != nil {
		return x.ListId
	}
	return ""
}

func (x *CreateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTaskRequest) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *CreateTaskRequest) GetDueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DueDate
	}
	return nil
}

func (x *CreateTaskRequest) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

type GetTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	mi := &file_task_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DueDate       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	mi := &file_task_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateTaskRequest) GetDueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DueDate
	}
	return nil
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	mi := &file_task_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MoveTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	ToListId      string                 `protobuf:"bytes,2,opt,name=to_list_id,json=toListId,proto3" json:"to_list_id,omitempty"`
	Position      int32                  `protobuf:"varint,3,opt,name=position,proto3" json:"position,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MoveTaskRequest) Reset() {
	*x = MoveTaskRequest{}
	mi := &file_task_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MoveTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveTaskRequest) ProtoMessage() {}

func (x *MoveTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveTaskRequest.ProtoReflect.Descriptor instead.
func (*MoveTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{6}
}

func (x *MoveTaskRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *MoveTaskRequest) GetToListId() string {
	if x != nil {
		return x.ToListId
	}
	return ""
}

func (x *MoveTaskRequest) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

type AssignUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssignUserRequest) Reset() {
	*x = AssignUserRequest{}
	mi := &file_task_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignUserRequest) ProtoMessage() {}

func (x *AssignUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignUserRequest.ProtoReflect.Descriptor instead.
func (*AssignUserRequest) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{7}
}

func (x *AssignUserRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *AssignUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UnassignUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnassignUserRequest) Reset() {
	*x = UnassignUserRequest{}
	mi := &file_task_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnassignUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnassignUserRequest) ProtoMessage() {}

func (x *UnassignUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnassignUserRequest.ProtoReflect.Descriptor instead.
func (*UnassignUserRequest) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{8}
}

func (x *UnassignUserRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *UnassignUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetTasksForUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTasksForUserRequest) Reset() {
	*x = GetTasksForUserRequest{}
	mi := &file_task_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTasksForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksForUserRequest) ProtoMessage() {}

func (x *GetTasksForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksForUserRequest.ProtoReflect.Descriptor instead.
func (*GetTasksForUserRequest) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetTasksForUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetTasksForUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*Task                `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTasksForUserResponse) Reset() {
	*x = GetTasksForUserResponse{}
	mi := &file_task_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTasksForUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksForUserResponse) ProtoMessage() {}

func (x *GetTasksForUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksForUserResponse.ProtoReflect.Descriptor instead.
func (*GetTasksForUserResponse) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{10}
}

func (x *GetTasksForUserResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type CreateLabelRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BoardId       string                 `protobuf:"bytes,1,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Color         string                 `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateLabelRequest) Reset() {
	*x = CreateLabelRequest{}
	mi := &file_task_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLabelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLabelRequest) ProtoMessage() {}

func (x *CreateLabelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLabelRequest.ProtoReflect.Descriptor instead.
func (*CreateLabelRequest) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{11}
}

func (x *CreateLabelRequest) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

func (x *CreateLabelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateLabelRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type UpdateLabelRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Color         string                 `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateLabelRequest) Reset() {
	*x = UpdateLabelRequest{}
	mi := &file_task_service_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateLabelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLabelRequest) ProtoMessage() {}

func (x *UpdateLabelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLabelRequest.ProtoReflect.Descriptor instead.
func (*UpdateLabelRequest) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateLabelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateLabelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateLabelRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type DeleteLabelRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteLabelRequest) Reset() {
	*x = DeleteLabelRequest{}
	mi := &file_task_service_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteLabelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLabelRequest) ProtoMessage() {}

func (x *DeleteLabelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLabelRequest.ProtoReflect.Descriptor instead.
func (*DeleteLabelRequest) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteLabelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AddLabelToTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	LabelId       string                 `protobuf:"bytes,2,opt,name=label_id,json=labelId,proto3" json:"label_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddLabelToTaskRequest) Reset() {
	*x = AddLabelToTaskRequest{}
	mi := &file_task_service_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddLabelToTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLabelToTaskRequest) ProtoMessage() {}

func (x *AddLabelToTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLabelToTaskRequest.ProtoReflect.Descriptor instead.
func (*AddLabelToTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{14}
}

func (x *AddLabelToTaskRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *AddLabelToTaskRequest) GetLabelId() string {
	if x != nil {
		return x.LabelId
	}
	return ""
}

type RemoveLabelFromTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	LabelId       string                 `protobuf:"bytes,2,opt,name=label_id,json=labelId,proto3" json:"label_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveLabelFromTaskRequest) Reset() {
	*x = RemoveLabelFromTaskRequest{}
	mi := &file_task_service_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveLabelFromTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveLabelFromTaskRequest) ProtoMessage() {}

func (x *RemoveLabelFromTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveLabelFromTaskRequest.ProtoReflect.Descriptor instead.
func (*RemoveLabelFromTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{15}
}

func (x *RemoveLabelFromTaskRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *RemoveLabelFromTaskRequest) GetLabelId() string {
	if x != nil {
		return x.LabelId
	}
	return ""
}

type GetTasksForListsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ListIds       []string               `protobuf:"bytes,1,rep,name=list_ids,json=listIds,proto3" json:"list_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTasksForListsRequest) Reset() {
	*x = GetTasksForListsRequest{}
	mi := &file_task_service_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTasksForListsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksForListsRequest) ProtoMessage() {}

func (x *GetTasksForListsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksForListsRequest.ProtoReflect.Descriptor instead.
func (*GetTasksForListsRequest) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{16}
}

func (x *GetTasksForListsRequest) GetListIds() []string {
	if x != nil {
		return x.ListIds
	}
	return nil
}

type GetTasksForListsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*Task                `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTasksForListsResponse) Reset() {
	*x = GetTasksForListsResponse{}
	mi := &file_task_service_proto_msgTypes[17]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTasksForListsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksForListsResponse) ProtoMessage() {}

func (x *GetTasksForListsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_service_proto_msgTypes[17]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksForListsResponse.ProtoReflect.Descriptor instead.
func (*GetTasksForListsResponse) Descriptor() ([]byte, []int) {
	return file_task_service_proto_rawDescGZIP(), []int{17}
}

func (x *GetTasksForListsResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_task_service_proto protoreflect.FileDescriptor

const file_task_service_proto_rawDesc = "" +
	"\n" +
	"\x12task_service.proto\x12\atask.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1bgoogle/protobuf/empty.proto\"\x9e\x03\n" +
	"\x04Task\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\alist_id\x18\x02 \x01(\tR\x06listId\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12\x1a\n" +
	"\bposition\x18\x05 \x01(\x05R\bposition\x125\n" +
	"\bdue_date\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\adueDate\x12\x1d\n" +
	"\n" +
	"created_by\x18\a \x01(\tR\tcreatedBy\x129\n" +
	"\n" +
	"created_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\t \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12%\n" +
	"\x0eassigned_users\x18\n" +
	" \x03(\tR\rassignedUsers\x12&\n" +
	"\x06labels\x18\v \x03(\v2\x0e.task.v1.LabelR\x06labels\"\\\n" +
	"\x05Label\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x19\n" +
	"\bboard_id\x18\x02 \x01(\tR\aboardId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x14\n" +
	"\x05color\x18\x04 \x01(\tR\x05color\"\xd6\x01\n" +
	"\x11CreateTaskRequest\x12\x17\n" +
	"\alist_id\x18\x01 \x01(\tR\x06listId\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1a\n" +
	"\bposition\x18\x04 \x01(\x05R\bposition\x125\n" +
	"\bdue_date\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\adueDate\x12\x1d\n" +
	"\n" +
	"created_by\x18\x06 \x01(\tR\tcreatedBy\" \n" +
	"\x0eGetTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"\x92\x01\n" +
	"\x11UpdateTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x125\n" +
	"\bdue_date\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\adueDate\"#\n" +
	"\x11DeleteTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"d\n" +
	"\x0fMoveTaskRequest\x12\x17\n" +
	"\atask_id\x18\x01 \x01(\tR\x06taskId\x12\x1c\n" +
	"\n" +
	"to_list_id\x18\x02 \x01(\tR\btoListId\x12\x1a\n" +
	"\bposition\x18\x03 \x01(\x05R\bposition\"E\n" +
	"\x11AssignUserRequest\x12\x17\n" +
	"\atask_id\x18\x01 \x01(\tR\x06taskId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\"G\n" +
	"\x13UnassignUserRequest\x12\x17\n" +
	"\atask_id\x18\x01 \x01(\tR\x06taskId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\"1\n" +
	"\x16GetTasksForUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\">\n" +
	"\x17GetTasksForUserResponse\x12#\n" +
	"\x05tasks\x18\x01 \x03(\v2\r.task.v1.TaskR\x05tasks\"Y\n" +
	"\x12CreateLabelRequest\x12\x19\n" +
	"\bboard_id\x18\x01 \x01(\tR\aboardId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05color\x18\x03 \x01(\tR\x05color\"N\n" +
	"\x12UpdateLabelRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05color\x18\x03 \x01(\tR\x05color\"$\n" +
	"\x12DeleteLabelRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"K\n" +
	"\x15AddLabelToTaskRequest\x12\x17\n" +
	"\atask_id\x18\x01 \x01(\tR\x06taskId\x12\x19\n" +
	"\blabel_id\x18\x02 \x01(\tR\alabelId\"P\n" +
	"\x1aRemoveLabelFromTaskRequest\x12\x17\n" +
	"\atask_id\x18\x01 \x01(\tR\x06taskId\x12\x19\n" +
	"\blabel_id\x18\x02 \x01(\tR\alabelId\"4\n" +
	"\x17GetTasksForListsRequest\x12\x19\n" +
	"\blist_ids\x18\x01 \x03(\tR\alistIds\"?\n" +
	"\x18GetTasksForListsResponse\x12#\n" +
	"\x05tasks\x18\x01 \x03(\v2\r.task.v1.TaskR\x05tasks2\xa8\a\n" +
	"\vTaskService\x127\n" +
	"\n" +
	"CreateTask\x12\x1a.task.v1.CreateTaskRequest\x1a\r.task.v1.Task\x121\n" +
	"\aGetTask\x12\x17.task.v1.GetTaskRequest\x1a\r.task.v1.Task\x127\n" +
	"\n" +
	"UpdateTask\x12\x1a.task.v1.UpdateTaskRequest\x1a\r.task.v1.Task\x12@\n" +
	"\n" +
	"DeleteTask\x12\x1a.task.v1.DeleteTaskRequest\x1a\x16.google.protobuf.Empty\x123\n" +
	"\bMoveTask\x12\x18.task.v1.MoveTaskRequest\x1a\r.task.v1.Task\x127\n" +
	"\n" +
	"AssignUser\x12\x1a.task.v1.AssignUserRequest\x1a\r.task.v1.Task\x12;\n" +
	"\fUnassignUser\x12\x1c.task.v1.UnassignUserRequest\x1a\r.task.v1.Task\x12W\n" +
	"\x10GetTasksForLists\x12 .task.v1.GetTasksForListsRequest\x1a!.task.v1.GetTasksForListsResponse\x12T\n" +
	"\x0fGetTasksForUser\x12\x1f.task.v1.GetTasksForUserRequest\x1a .task.v1.GetTasksForUserResponse\x12:\n" +
	"\vCreateLabel\x12\x1b.task.v1.CreateLabelRequest\x1a\x0e.task.v1.Label\x12:\n" +
	"\vUpdateLabel\x12\x1b.task.v1.UpdateLabelRequest\x1a\x0e.task.v1.Label\x12B\n" +
	"\vDeleteLabel\x12\x1b.task.v1.DeleteLabelRequest\x1a\x16.google.protobuf.Empty\x12H\n" +
	"\x0eAddLabelToTask\x12\x1e.task.v1.AddLabelToTaskRequest\x1a\x16.google.protobuf.Empty\x12R\n" +
	"\x13RemoveLabelFromTask\x12#.task.v1.RemoveLabelFromTaskRequest\x1a\x16.google.protobuf.EmptyB\x1aZ\x18shiroyama.task.v1;taskv1b\x06proto3"

var (
	file_task_service_proto_rawDescOnce sync.Once
	file_task_service_proto_rawDescData []byte
)

func file_task_service_proto_rawDescGZIP() []byte {
	file_task_service_proto_rawDescOnce.Do(func() {
		file_task_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_task_service_proto_rawDesc), len(file_task_service_proto_rawDesc)))
	})
	return file_task_service_proto_rawDescData
}

var file_task_service_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_task_service_proto_goTypes = []any{
	(*Task)(nil),                       // 0: task.v1.Task
	(*Label)(nil),                      // 1: task.v1.Label
	(*CreateTaskRequest)(nil),          // 2: task.v1.CreateTaskRequest
	(*GetTaskRequest)(nil),             // 3: task.v1.GetTaskRequest
	(*UpdateTaskRequest)(nil),          // 4: task.v1.UpdateTaskRequest
	(*DeleteTaskRequest)(nil),          // 5: task.v1.DeleteTaskRequest
	(*MoveTaskRequest)(nil),            // 6: task.v1.MoveTaskRequest
	(*AssignUserRequest)(nil),          // 7: task.v1.AssignUserRequest
	(*UnassignUserRequest)(nil),        // 8: task.v1.UnassignUserRequest
	(*GetTasksForUserRequest)(nil),     // 9: task.v1.GetTasksForUserRequest
	(*GetTasksForUserResponse)(nil),    // 10: task.v1.GetTasksForUserResponse
	(*CreateLabelRequest)(nil),         // 11: task.v1.CreateLabelRequest
	(*UpdateLabelRequest)(nil),         // 12: task.v1.UpdateLabelRequest
	(*DeleteLabelRequest)(nil),         // 13: task.v1.DeleteLabelRequest
	(*AddLabelToTaskRequest)(nil),      // 14: task.v1.AddLabelToTaskRequest
	(*RemoveLabelFromTaskRequest)(nil), // 15: task.v1.RemoveLabelFromTaskRequest
	(*GetTasksForListsRequest)(nil),    // 16: task.v1.GetTasksForListsRequest
	(*GetTasksForListsResponse)(nil),   // 17: task.v1.GetTasksForListsResponse
	(*timestamppb.Timestamp)(nil),      // 18: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),              // 19: google.protobuf.Empty
}
var file_task_service_proto_depIdxs = []int32{
	18, // 0: task.v1.Task.due_date:type_name -> google.protobuf.Timestamp
	18, // 1: task.v1.Task.created_at:type_name -> google.protobuf.Timestamp
	18, // 2: task.v1.Task.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 3: task.v1.Task.labels:type_name -> task.v1.Label
	18, // 4: task.v1.CreateTaskRequest.due_date:type_name -> google.protobuf.Timestamp
	18, // 5: task.v1.UpdateTaskRequest.due_date:type_name -> google.protobuf.Timestamp
	0,  // 6: task.v1.GetTasksForUserResponse.tasks:type_name -> task.v1.Task
	0,  // 7: task.v1.GetTasksForListsResponse.tasks:type_name -> task.v1.Task
	2,  // 8: task.v1.TaskService.CreateTask:input_type -> task.v1.CreateTaskRequest
	3,  // 9: task.v1.TaskService.GetTask:input_type -> task.v1.GetTaskRequest
	4,  // 10: task.v1.TaskService.UpdateTask:input_type -> task.v1.UpdateTaskRequest
	5,  // 11: task.v1.TaskService.DeleteTask:input_type -> task.v1.DeleteTaskRequest
	6,  // 12: task.v1.TaskService.MoveTask:input_type -> task.v1.MoveTaskRequest
	7,  // 13: task.v1.TaskService.AssignUser:input_type -> task.v1.AssignUserRequest
	8,  // 14: task.v1.TaskService.UnassignUser:input_type -> task.v1.UnassignUserRequest
	16, // 15: task.v1.TaskService.GetTasksForLists:input_type -> task.v1.GetTasksForListsRequest
	9,  // 16: task.v1.TaskService.GetTasksForUser:input_type -> task.v1.GetTasksForUserRequest
	11, // 17: task.v1.TaskService.CreateLabel:input_type -> task.v1.CreateLabelRequest
	12, // 18: task.v1.TaskService.UpdateLabel:input_type -> task.v1.UpdateLabelRequest
	13, // 19: task.v1.TaskService.DeleteLabel:input_type -> task.v1.DeleteLabelRequest
	14, // 20: task.v1.TaskService.AddLabelToTask:input_type -> task.v1.AddLabelToTaskRequest
	15, // 21: task.v1.TaskService.RemoveLabelFromTask:input_type -> task.v1.RemoveLabelFromTaskRequest
	0,  // 22: task.v1.TaskService.CreateTask:output_type -> task.v1.Task
	0,  // 23: task.v1.TaskService.GetTask:output_type -> task.v1.Task
	0,  // 24: task.v1.TaskService.UpdateTask:output_type -> task.v1.Task
	19, // 25: task.v1.TaskService.DeleteTask:output_type -> google.protobuf.Empty
	0,  // 26: task.v1.TaskService.MoveTask:output_type -> task.v1.Task
	0,  // 27: task.v1.TaskService.AssignUser:output_type -> task.v1.Task
	0,  // 28: task.v1.TaskService.UnassignUser:output_type -> task.v1.Task
	17, // 29: task.v1.TaskService.GetTasksForLists:output_type -> task.v1.GetTasksForListsResponse
	10, // 30: task.v1.TaskService.GetTasksForUser:output_type -> task.v1.GetTasksForUserResponse
	1,  // 31: task.v1.TaskService.CreateLabel:output_type -> task.v1.Label
	1,  // 32: task.v1.TaskService.UpdateLabel:output_type -> task.v1.Label
	19, // 33: task.v1.TaskService.DeleteLabel:output_type -> google.protobuf.Empty
	19, // 34: task.v1.TaskService.AddLabelToTask:output_type -> google.protobuf.Empty
	19, // 35: task.v1.TaskService.RemoveLabelFromTask:output_type -> google.protobuf.Empty
	22, // [22:36] is the sub-list for method output_type
	8,  // [8:22] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_task_service_proto_init() }
func file_task_service_proto_init() {
	if File_task_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_task_service_proto_rawDesc), len(file_task_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_service_proto_goTypes,
		DependencyIndexes: file_task_service_proto_depIdxs,
		MessageInfos:      file_task_service_proto_msgTypes,
	}.Build()
	File_task_service_proto = out.File
	file_task_service_proto_goTypes = nil
	file_task_service_proto_depIdxs = nil
}
