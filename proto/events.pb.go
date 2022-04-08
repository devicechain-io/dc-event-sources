// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: proto/events.proto

package proto

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

//*
// Unresolved event that has been decoded by
// an event source.
type PUnresolvedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId      string  `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	AltId         *string `protobuf:"bytes,2,opt,name=alt_id,json=altId,proto3,oneof" json:"alt_id,omitempty"`
	Device        string  `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`
	Assignment    *string `protobuf:"bytes,4,opt,name=assignment,proto3,oneof" json:"assignment,omitempty"`
	Customer      *string `protobuf:"bytes,5,opt,name=customer,proto3,oneof" json:"customer,omitempty"`
	Area          *string `protobuf:"bytes,6,opt,name=area,proto3,oneof" json:"area,omitempty"`
	Asset         *string `protobuf:"bytes,7,opt,name=asset,proto3,oneof" json:"asset,omitempty"`
	OccurredTime  string  `protobuf:"bytes,8,opt,name=occurred_time,json=occurredTime,proto3" json:"occurred_time,omitempty"`
	ProcessedTime string  `protobuf:"bytes,9,opt,name=processed_time,json=processedTime,proto3" json:"processed_time,omitempty"`
	EventType     int64   `protobuf:"varint,10,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	Payload       []byte  `protobuf:"bytes,12,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *PUnresolvedEvent) Reset() {
	*x = PUnresolvedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PUnresolvedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PUnresolvedEvent) ProtoMessage() {}

func (x *PUnresolvedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PUnresolvedEvent.ProtoReflect.Descriptor instead.
func (*PUnresolvedEvent) Descriptor() ([]byte, []int) {
	return file_proto_events_proto_rawDescGZIP(), []int{0}
}

func (x *PUnresolvedEvent) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *PUnresolvedEvent) GetAltId() string {
	if x != nil && x.AltId != nil {
		return *x.AltId
	}
	return ""
}

func (x *PUnresolvedEvent) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *PUnresolvedEvent) GetAssignment() string {
	if x != nil && x.Assignment != nil {
		return *x.Assignment
	}
	return ""
}

func (x *PUnresolvedEvent) GetCustomer() string {
	if x != nil && x.Customer != nil {
		return *x.Customer
	}
	return ""
}

func (x *PUnresolvedEvent) GetArea() string {
	if x != nil && x.Area != nil {
		return *x.Area
	}
	return ""
}

func (x *PUnresolvedEvent) GetAsset() string {
	if x != nil && x.Asset != nil {
		return *x.Asset
	}
	return ""
}

func (x *PUnresolvedEvent) GetOccurredTime() string {
	if x != nil {
		return x.OccurredTime
	}
	return ""
}

func (x *PUnresolvedEvent) GetProcessedTime() string {
	if x != nil {
		return x.ProcessedTime
	}
	return ""
}

func (x *PUnresolvedEvent) GetEventType() int64 {
	if x != nil {
		return x.EventType
	}
	return 0
}

func (x *PUnresolvedEvent) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

//*
// Payload for a location event.
type PLocationPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  *string `protobuf:"bytes,1,opt,name=latitude,proto3,oneof" json:"latitude,omitempty"`
	Longitude *string `protobuf:"bytes,2,opt,name=longitude,proto3,oneof" json:"longitude,omitempty"`
	Elevation *string `protobuf:"bytes,3,opt,name=elevation,proto3,oneof" json:"elevation,omitempty"`
}

func (x *PLocationPayload) Reset() {
	*x = PLocationPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PLocationPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PLocationPayload) ProtoMessage() {}

func (x *PLocationPayload) ProtoReflect() protoreflect.Message {
	mi := &file_proto_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PLocationPayload.ProtoReflect.Descriptor instead.
func (*PLocationPayload) Descriptor() ([]byte, []int) {
	return file_proto_events_proto_rawDescGZIP(), []int{1}
}

func (x *PLocationPayload) GetLatitude() string {
	if x != nil && x.Latitude != nil {
		return *x.Latitude
	}
	return ""
}

func (x *PLocationPayload) GetLongitude() string {
	if x != nil && x.Longitude != nil {
		return *x.Longitude
	}
	return ""
}

func (x *PLocationPayload) GetElevation() string {
	if x != nil && x.Elevation != nil {
		return *x.Elevation
	}
	return ""
}

var File_proto_events_proto protoreflect.FileDescriptor

var file_proto_events_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x03, 0x0a, 0x10, 0x50, 0x55, 0x6e, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x06, 0x61, 0x6c, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x61, 0x6c, 0x74, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x1f, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03,
	0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x63, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x6c,
	0x74, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x22, 0xa2, 0x01, 0x0a, 0x10, 0x50, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1f, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09,
	0x65, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x09, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x65,
	0x6c, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_events_proto_rawDescOnce sync.Once
	file_proto_events_proto_rawDescData = file_proto_events_proto_rawDesc
)

func file_proto_events_proto_rawDescGZIP() []byte {
	file_proto_events_proto_rawDescOnce.Do(func() {
		file_proto_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_events_proto_rawDescData)
	})
	return file_proto_events_proto_rawDescData
}

var file_proto_events_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_events_proto_goTypes = []interface{}{
	(*PUnresolvedEvent)(nil), // 0: PUnresolvedEvent
	(*PLocationPayload)(nil), // 1: PLocationPayload
}
var file_proto_events_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_events_proto_init() }
func file_proto_events_proto_init() {
	if File_proto_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PUnresolvedEvent); i {
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
		file_proto_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PLocationPayload); i {
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
	file_proto_events_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_events_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_events_proto_goTypes,
		DependencyIndexes: file_proto_events_proto_depIdxs,
		MessageInfos:      file_proto_events_proto_msgTypes,
	}.Build()
	File_proto_events_proto = out.File
	file_proto_events_proto_rawDesc = nil
	file_proto_events_proto_goTypes = nil
	file_proto_events_proto_depIdxs = nil
}
