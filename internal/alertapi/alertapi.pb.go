// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: alertapi/alertapi.proto

package alertapi

import (
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

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alertapi_alertapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_alertapi_alertapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_alertapi_alertapi_proto_rawDescGZIP(), []int{0}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alertapi_alertapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_alertapi_alertapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_alertapi_alertapi_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerName   string `protobuf:"bytes,1,opt,name=serverName,proto3" json:"serverName,omitempty"`
	MetricName   string `protobuf:"bytes,2,opt,name=metricName,proto3" json:"metricName,omitempty"`
	LogId        int64  `protobuf:"varint,3,opt,name=logId,proto3" json:"logId,omitempty"`
	Status       int32  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Subject      string `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	Content      string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp    string `protobuf:"bytes,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Resolved     bool   `protobuf:"varint,8,opt,name=resolved,proto3" json:"resolved,omitempty"`
	Disk         string `protobuf:"bytes,9,opt,name=disk,proto3" json:"disk,omitempty"`
	Service      string `protobuf:"bytes,10,opt,name=service,proto3" json:"service,omitempty"`
	Pagerduty    bool   `protobuf:"varint,11,opt,name=pagerduty,proto3" json:"pagerduty,omitempty"`
	Email        bool   `protobuf:"varint,12,opt,name=email,proto3" json:"email,omitempty"`
	Slack        bool   `protobuf:"varint,13,opt,name=slack,proto3" json:"slack,omitempty"`
	SlackChannel string `protobuf:"bytes,14,opt,name=slackChannel,proto3" json:"slackChannel,omitempty"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alertapi_alertapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_alertapi_alertapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_alertapi_alertapi_proto_rawDescGZIP(), []int{2}
}

func (x *Alert) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *Alert) GetMetricName() string {
	if x != nil {
		return x.MetricName
	}
	return ""
}

func (x *Alert) GetLogId() int64 {
	if x != nil {
		return x.LogId
	}
	return 0
}

func (x *Alert) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Alert) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Alert) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Alert) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Alert) GetResolved() bool {
	if x != nil {
		return x.Resolved
	}
	return false
}

func (x *Alert) GetDisk() string {
	if x != nil {
		return x.Disk
	}
	return ""
}

func (x *Alert) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Alert) GetPagerduty() bool {
	if x != nil {
		return x.Pagerduty
	}
	return false
}

func (x *Alert) GetEmail() bool {
	if x != nil {
		return x.Email
	}
	return false
}

func (x *Alert) GetSlack() bool {
	if x != nil {
		return x.Slack
	}
	return false
}

func (x *Alert) GetSlackChannel() string {
	if x != nil {
		return x.SlackChannel
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerName string `protobuf:"bytes,1,opt,name=serverName,proto3" json:"serverName,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alertapi_alertapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_alertapi_alertapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_alertapi_alertapi_proto_rawDescGZIP(), []int{3}
}

func (x *Request) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

type AlertArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alerts []*Alert `protobuf:"bytes,1,rep,name=alerts,proto3" json:"alerts,omitempty"`
}

func (x *AlertArray) Reset() {
	*x = AlertArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alertapi_alertapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertArray) ProtoMessage() {}

func (x *AlertArray) ProtoReflect() protoreflect.Message {
	mi := &file_alertapi_alertapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertArray.ProtoReflect.Descriptor instead.
func (*AlertArray) Descriptor() ([]byte, []int) {
	return file_alertapi_alertapi_proto_rawDescGZIP(), []int{4}
}

func (x *AlertArray) GetAlerts() []*Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

var File_alertapi_alertapi_proto protoreflect.FileDescriptor

var file_alertapi_alertapi_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x61, 0x70, 0x69, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0xff, 0x02, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x6f,
	0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x69, 0x73,
	0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x64, 0x75, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x64, 0x75, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x6c, 0x61, 0x63, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x6c, 0x61, 0x63,
	0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x29, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x35, 0x0a, 0x0a, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x27,
	0x0a, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52,
	0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x32, 0x80, 0x01, 0x0a, 0x0c, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x0f, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x39, 0x0a, 0x0c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x11, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alertapi_alertapi_proto_rawDescOnce sync.Once
	file_alertapi_alertapi_proto_rawDescData = file_alertapi_alertapi_proto_rawDesc
)

func file_alertapi_alertapi_proto_rawDescGZIP() []byte {
	file_alertapi_alertapi_proto_rawDescOnce.Do(func() {
		file_alertapi_alertapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_alertapi_alertapi_proto_rawDescData)
	})
	return file_alertapi_alertapi_proto_rawDescData
}

var file_alertapi_alertapi_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_alertapi_alertapi_proto_goTypes = []interface{}{
	(*Void)(nil),       // 0: alertapi.Void
	(*Response)(nil),   // 1: alertapi.Response
	(*Alert)(nil),      // 2: alertapi.Alert
	(*Request)(nil),    // 3: alertapi.Request
	(*AlertArray)(nil), // 4: alertapi.AlertArray
}
var file_alertapi_alertapi_proto_depIdxs = []int32{
	2, // 0: alertapi.AlertArray.alerts:type_name -> alertapi.Alert
	2, // 1: alertapi.AlertService.HandleAlerts:input_type -> alertapi.Alert
	3, // 2: alertapi.AlertService.AlertRequest:input_type -> alertapi.Request
	1, // 3: alertapi.AlertService.HandleAlerts:output_type -> alertapi.Response
	4, // 4: alertapi.AlertService.AlertRequest:output_type -> alertapi.AlertArray
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_alertapi_alertapi_proto_init() }
func file_alertapi_alertapi_proto_init() {
	if File_alertapi_alertapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alertapi_alertapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_alertapi_alertapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_alertapi_alertapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alert); i {
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
		file_alertapi_alertapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_alertapi_alertapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertArray); i {
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
			RawDescriptor: file_alertapi_alertapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_alertapi_alertapi_proto_goTypes,
		DependencyIndexes: file_alertapi_alertapi_proto_depIdxs,
		MessageInfos:      file_alertapi_alertapi_proto_msgTypes,
	}.Build()
	File_alertapi_alertapi_proto = out.File
	file_alertapi_alertapi_proto_rawDesc = nil
	file_alertapi_alertapi_proto_goTypes = nil
	file_alertapi_alertapi_proto_depIdxs = nil
}