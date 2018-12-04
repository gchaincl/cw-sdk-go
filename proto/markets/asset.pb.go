// Code generated by protoc-gen-go. DO NOT EDIT.
// source: markets/asset.proto

package ProtobufMarkets

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

type AssetUpdateMessage struct {
	Asset int32 `protobuf:"varint,1,opt,name=asset" json:"asset,omitempty"`
	// Types that are valid to be assigned to Update:
	//	*AssetUpdateMessage_UsdVolumeUpdate
	Update               isAssetUpdateMessage_Update `protobuf_oneof:"Update"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *AssetUpdateMessage) Reset()         { *m = AssetUpdateMessage{} }
func (m *AssetUpdateMessage) String() string { return proto.CompactTextString(m) }
func (*AssetUpdateMessage) ProtoMessage()    {}
func (*AssetUpdateMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_asset_11b1ebcf6470df0b, []int{0}
}
func (m *AssetUpdateMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetUpdateMessage.Unmarshal(m, b)
}
func (m *AssetUpdateMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetUpdateMessage.Marshal(b, m, deterministic)
}
func (dst *AssetUpdateMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetUpdateMessage.Merge(dst, src)
}
func (m *AssetUpdateMessage) XXX_Size() int {
	return xxx_messageInfo_AssetUpdateMessage.Size(m)
}
func (m *AssetUpdateMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetUpdateMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AssetUpdateMessage proto.InternalMessageInfo

type isAssetUpdateMessage_Update interface {
	isAssetUpdateMessage_Update()
}

type AssetUpdateMessage_UsdVolumeUpdate struct {
	UsdVolumeUpdate *AssetUSDVolumeUpdate `protobuf:"bytes,2,opt,name=usdVolumeUpdate,oneof"`
}

func (*AssetUpdateMessage_UsdVolumeUpdate) isAssetUpdateMessage_Update() {}

func (m *AssetUpdateMessage) GetUpdate() isAssetUpdateMessage_Update {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *AssetUpdateMessage) GetAsset() int32 {
	if m != nil {
		return m.Asset
	}
	return 0
}

func (m *AssetUpdateMessage) GetUsdVolumeUpdate() *AssetUSDVolumeUpdate {
	if x, ok := m.GetUpdate().(*AssetUpdateMessage_UsdVolumeUpdate); ok {
		return x.UsdVolumeUpdate
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AssetUpdateMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AssetUpdateMessage_OneofMarshaler, _AssetUpdateMessage_OneofUnmarshaler, _AssetUpdateMessage_OneofSizer, []interface{}{
		(*AssetUpdateMessage_UsdVolumeUpdate)(nil),
	}
}

func _AssetUpdateMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AssetUpdateMessage)
	// Update
	switch x := m.Update.(type) {
	case *AssetUpdateMessage_UsdVolumeUpdate:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UsdVolumeUpdate); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AssetUpdateMessage.Update has unexpected type %T", x)
	}
	return nil
}

func _AssetUpdateMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AssetUpdateMessage)
	switch tag {
	case 2: // Update.usdVolumeUpdate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AssetUSDVolumeUpdate)
		err := b.DecodeMessage(msg)
		m.Update = &AssetUpdateMessage_UsdVolumeUpdate{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AssetUpdateMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AssetUpdateMessage)
	// Update
	switch x := m.Update.(type) {
	case *AssetUpdateMessage_UsdVolumeUpdate:
		s := proto.Size(x.UsdVolumeUpdate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type AssetUSDVolumeUpdate struct {
	Volume               string   `protobuf:"bytes,1,opt,name=volume" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetUSDVolumeUpdate) Reset()         { *m = AssetUSDVolumeUpdate{} }
func (m *AssetUSDVolumeUpdate) String() string { return proto.CompactTextString(m) }
func (*AssetUSDVolumeUpdate) ProtoMessage()    {}
func (*AssetUSDVolumeUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_asset_11b1ebcf6470df0b, []int{1}
}
func (m *AssetUSDVolumeUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetUSDVolumeUpdate.Unmarshal(m, b)
}
func (m *AssetUSDVolumeUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetUSDVolumeUpdate.Marshal(b, m, deterministic)
}
func (dst *AssetUSDVolumeUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetUSDVolumeUpdate.Merge(dst, src)
}
func (m *AssetUSDVolumeUpdate) XXX_Size() int {
	return xxx_messageInfo_AssetUSDVolumeUpdate.Size(m)
}
func (m *AssetUSDVolumeUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetUSDVolumeUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_AssetUSDVolumeUpdate proto.InternalMessageInfo

func (m *AssetUSDVolumeUpdate) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

func init() {
	proto.RegisterType((*AssetUpdateMessage)(nil), "ProtobufMarkets.AssetUpdateMessage")
	proto.RegisterType((*AssetUSDVolumeUpdate)(nil), "ProtobufMarkets.AssetUSDVolumeUpdate")
}

func init() { proto.RegisterFile("markets/asset.proto", fileDescriptor_asset_11b1ebcf6470df0b) }

var fileDescriptor_asset_11b1ebcf6470df0b = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x4d, 0x2c, 0xca,
	0x4e, 0x2d, 0x29, 0xd6, 0x4f, 0x2c, 0x2e, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x0f, 0x00, 0x51, 0x49, 0xa5, 0x69, 0xbe, 0x10, 0x49, 0xa5, 0x76, 0x46, 0x2e, 0x21, 0x47,
	0x90, 0x82, 0xd0, 0x82, 0x94, 0xc4, 0x92, 0x54, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x21,
	0x11, 0x2e, 0x56, 0xb0, 0x36, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x08, 0x47, 0x28, 0x90,
	0x8b, 0xbf, 0xb4, 0x38, 0x25, 0x2c, 0x3f, 0xa7, 0x34, 0x37, 0x15, 0xa2, 0x5e, 0x82, 0x49, 0x81,
	0x51, 0x83, 0xdb, 0x48, 0x55, 0x0f, 0xcd, 0x5c, 0x3d, 0x88, 0x99, 0xc1, 0x2e, 0xc8, 0x8a, 0x3d,
	0x18, 0x82, 0xd0, 0xf5, 0x3b, 0x71, 0x70, 0xb1, 0x41, 0x58, 0x4a, 0x7a, 0x5c, 0x22, 0xd8, 0x34,
	0x09, 0x89, 0x71, 0xb1, 0x95, 0x81, 0xf9, 0x60, 0xb7, 0x70, 0x06, 0x41, 0x79, 0x49, 0x6c, 0x60,
	0x1f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x16, 0x95, 0x35, 0xf6, 0xe8, 0x00, 0x00, 0x00,
}
