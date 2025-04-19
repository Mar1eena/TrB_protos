// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: server/server.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Сообщения для unary-методов
type MessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Text          string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageRequest) Reset() {
	*x = MessageRequest{}
	mi := &file_server_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequest) ProtoMessage() {}

func (x *MessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRequest.ProtoReflect.Descriptor instead.
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{0}
}

func (x *MessageRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type MessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reply         string                 `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	mi := &file_server_server_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{1}
}

func (x *MessageResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type PoliceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PoliceRequest) Reset() {
	*x = PoliceRequest{}
	mi := &file_server_server_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PoliceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoliceRequest) ProtoMessage() {}

func (x *PoliceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoliceRequest.ProtoReflect.Descriptor instead.
func (*PoliceRequest) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{2}
}

func (x *PoliceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type PoliceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PoliceResponse) Reset() {
	*x = PoliceResponse{}
	mi := &file_server_server_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PoliceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoliceResponse) ProtoMessage() {}

func (x *PoliceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoliceResponse.ProtoReflect.Descriptor instead.
func (*PoliceResponse) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{3}
}

func (x *PoliceResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *PoliceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Сообщения для server-side streaming
type PoliceStreamResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Update        string                 `protobuf:"bytes,1,opt,name=update,proto3" json:"update,omitempty"`
	Severity      int32                  `protobuf:"varint,2,opt,name=severity,proto3" json:"severity,omitempty"`
	Timestamp     string                 `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PoliceStreamResponse) Reset() {
	*x = PoliceStreamResponse{}
	mi := &file_server_server_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PoliceStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoliceStreamResponse) ProtoMessage() {}

func (x *PoliceStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoliceStreamResponse.ProtoReflect.Descriptor instead.
func (*PoliceStreamResponse) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{4}
}

func (x *PoliceStreamResponse) GetUpdate() string {
	if x != nil {
		return x.Update
	}
	return ""
}

func (x *PoliceStreamResponse) GetSeverity() int32 {
	if x != nil {
		return x.Severity
	}
	return 0
}

func (x *PoliceStreamResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

// Сообщения для client-side streaming
type ClientStreamRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Payload       []byte                 `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientStreamRequest) Reset() {
	*x = ClientStreamRequest{}
	mi := &file_server_server_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamRequest) ProtoMessage() {}

func (x *ClientStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamRequest.ProtoReflect.Descriptor instead.
func (*ClientStreamRequest) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{5}
}

func (x *ClientStreamRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ClientStreamRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type StreamSummary struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TotalRequests int32                  `protobuf:"varint,1,opt,name=total_requests,json=totalRequests,proto3" json:"total_requests,omitempty"`
	TotalBytes    int64                  `protobuf:"varint,2,opt,name=total_bytes,json=totalBytes,proto3" json:"total_bytes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamSummary) Reset() {
	*x = StreamSummary{}
	mi := &file_server_server_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamSummary) ProtoMessage() {}

func (x *StreamSummary) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamSummary.ProtoReflect.Descriptor instead.
func (*StreamSummary) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{6}
}

func (x *StreamSummary) GetTotalRequests() int32 {
	if x != nil {
		return x.TotalRequests
	}
	return 0
}

func (x *StreamSummary) GetTotalBytes() int64 {
	if x != nil {
		return x.TotalBytes
	}
	return 0
}

// Сообщения для bidirectional streaming
type BidirectionalMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sender        string                 `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Sequence      int64                  `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BidirectionalMessage) Reset() {
	*x = BidirectionalMessage{}
	mi := &file_server_server_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BidirectionalMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidirectionalMessage) ProtoMessage() {}

func (x *BidirectionalMessage) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidirectionalMessage.ProtoReflect.Descriptor instead.
func (*BidirectionalMessage) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{7}
}

func (x *BidirectionalMessage) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *BidirectionalMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *BidirectionalMessage) GetSequence() int64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

var File_server_server_proto protoreflect.FileDescriptor

const file_server_server_proto_rawDesc = "" +
	"\n" +
	"\x13server/server.proto\x12\x06server\x1a\x1cgoogle/api/annotations.proto\"$\n" +
	"\x0eMessageRequest\x12\x12\n" +
	"\x04text\x18\x01 \x01(\tR\x04text\"'\n" +
	"\x0fMessageResponse\x12\x14\n" +
	"\x05reply\x18\x01 \x01(\tR\x05reply\"%\n" +
	"\rPoliceRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"@\n" +
	"\x0ePoliceResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"h\n" +
	"\x14PoliceStreamResponse\x12\x16\n" +
	"\x06update\x18\x01 \x01(\tR\x06update\x12\x1a\n" +
	"\bseverity\x18\x02 \x01(\x05R\bseverity\x12\x1c\n" +
	"\ttimestamp\x18\x03 \x01(\tR\ttimestamp\"L\n" +
	"\x13ClientStreamRequest\x12\x1b\n" +
	"\tclient_id\x18\x01 \x01(\tR\bclientId\x12\x18\n" +
	"\apayload\x18\x02 \x01(\fR\apayload\"W\n" +
	"\rStreamSummary\x12%\n" +
	"\x0etotal_requests\x18\x01 \x01(\x05R\rtotalRequests\x12\x1f\n" +
	"\vtotal_bytes\x18\x02 \x01(\x03R\n" +
	"totalBytes\"d\n" +
	"\x14BidirectionalMessage\x12\x16\n" +
	"\x06sender\x18\x01 \x01(\tR\x06sender\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontent\x12\x1a\n" +
	"\bsequence\x18\x03 \x01(\x03R\bsequence2\x93\x03\n" +
	"\x0eExampleService\x12O\n" +
	"\vSendMessage\x12\x16.server.MessageRequest\x1a\x17.server.MessageResponse\"\x0f\x82\xd3\xe4\x93\x02\t:\x01*\"\x04/say\x12;\n" +
	"\n" +
	"SendPolice\x12\x15.server.PoliceRequest\x1a\x16.server.PoliceResponse\x12L\n" +
	"\x13StreamPoliceUpdates\x12\x15.server.PoliceRequest\x1a\x1c.server.PoliceStreamResponse0\x01\x12N\n" +
	"\x16ClientStreamingExample\x12\x1b.server.ClientStreamRequest\x1a\x15.server.StreamSummary(\x01\x12U\n" +
	"\x13BidirectionalStream\x12\x1c.server.BidirectionalMessage\x1a\x1c.server.BidirectionalMessage(\x010\x01B\vZ\tsrc/protob\x06proto3"

var (
	file_server_server_proto_rawDescOnce sync.Once
	file_server_server_proto_rawDescData []byte
)

func file_server_server_proto_rawDescGZIP() []byte {
	file_server_server_proto_rawDescOnce.Do(func() {
		file_server_server_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_server_server_proto_rawDesc), len(file_server_server_proto_rawDesc)))
	})
	return file_server_server_proto_rawDescData
}

var file_server_server_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_server_server_proto_goTypes = []any{
	(*MessageRequest)(nil),       // 0: server.MessageRequest
	(*MessageResponse)(nil),      // 1: server.MessageResponse
	(*PoliceRequest)(nil),        // 2: server.PoliceRequest
	(*PoliceResponse)(nil),       // 3: server.PoliceResponse
	(*PoliceStreamResponse)(nil), // 4: server.PoliceStreamResponse
	(*ClientStreamRequest)(nil),  // 5: server.ClientStreamRequest
	(*StreamSummary)(nil),        // 6: server.StreamSummary
	(*BidirectionalMessage)(nil), // 7: server.BidirectionalMessage
}
var file_server_server_proto_depIdxs = []int32{
	0, // 0: server.ExampleService.SendMessage:input_type -> server.MessageRequest
	2, // 1: server.ExampleService.SendPolice:input_type -> server.PoliceRequest
	2, // 2: server.ExampleService.StreamPoliceUpdates:input_type -> server.PoliceRequest
	5, // 3: server.ExampleService.ClientStreamingExample:input_type -> server.ClientStreamRequest
	7, // 4: server.ExampleService.BidirectionalStream:input_type -> server.BidirectionalMessage
	1, // 5: server.ExampleService.SendMessage:output_type -> server.MessageResponse
	3, // 6: server.ExampleService.SendPolice:output_type -> server.PoliceResponse
	4, // 7: server.ExampleService.StreamPoliceUpdates:output_type -> server.PoliceStreamResponse
	6, // 8: server.ExampleService.ClientStreamingExample:output_type -> server.StreamSummary
	7, // 9: server.ExampleService.BidirectionalStream:output_type -> server.BidirectionalMessage
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_server_proto_init() }
func file_server_server_proto_init() {
	if File_server_server_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_server_server_proto_rawDesc), len(file_server_server_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_server_proto_goTypes,
		DependencyIndexes: file_server_server_proto_depIdxs,
		MessageInfos:      file_server_server_proto_msgTypes,
	}.Build()
	File_server_server_proto = out.File
	file_server_server_proto_goTypes = nil
	file_server_server_proto_depIdxs = nil
}
