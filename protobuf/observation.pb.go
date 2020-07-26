// Code generated by protoc-gen-go. DO NOT EDIT.
// source: observation.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Constellation int32

const (
	Constellation_GPS     Constellation = 0
	Constellation_GLONASS Constellation = 1
	Constellation_GALILEO Constellation = 2
	Constellation_SBAS    Constellation = 3
	Constellation_QZSS    Constellation = 4
	Constellation_BEIDOU  Constellation = 5
)

var Constellation_name = map[int32]string{
	0: "GPS",
	1: "GLONASS",
	2: "GALILEO",
	3: "SBAS",
	4: "QZSS",
	5: "BEIDOU",
}

var Constellation_value = map[string]int32{
	"GPS":     0,
	"GLONASS": 1,
	"GALILEO": 2,
	"SBAS":    3,
	"QZSS":    4,
	"BEIDOU":  5,
}

func (x Constellation) String() string {
	return proto.EnumName(Constellation_name, int32(x))
}

func (Constellation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5ed04ec06c6de2a5, []int{0}
}

type SignalData struct {
	Band                 string   `protobuf:"bytes,1,opt,name=band,proto3" json:"band,omitempty"`
	Frequency            string   `protobuf:"bytes,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Pseudorange          float32  `protobuf:"fixed32,3,opt,name=pseudorange,proto3" json:"pseudorange,omitempty"`
	PhaseRange           float32  `protobuf:"fixed32,4,opt,name=phaseRange,proto3" json:"phaseRange,omitempty"`
	PhaseRangeLock       uint32   `protobuf:"varint,5,opt,name=phaseRangeLock,proto3" json:"phaseRangeLock,omitempty"`
	HalfCycle            bool     `protobuf:"varint,6,opt,name=halfCycle,proto3" json:"halfCycle,omitempty"`
	Snr                  float32  `protobuf:"fixed32,7,opt,name=snr,proto3" json:"snr,omitempty"`
	PhaseRangeRate       float32  `protobuf:"fixed32,8,opt,name=phaseRangeRate,proto3" json:"phaseRangeRate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignalData) Reset()         { *m = SignalData{} }
func (m *SignalData) String() string { return proto.CompactTextString(m) }
func (*SignalData) ProtoMessage()    {}
func (*SignalData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ed04ec06c6de2a5, []int{0}
}

func (m *SignalData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalData.Unmarshal(m, b)
}
func (m *SignalData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalData.Marshal(b, m, deterministic)
}
func (m *SignalData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalData.Merge(m, src)
}
func (m *SignalData) XXX_Size() int {
	return xxx_messageInfo_SignalData.Size(m)
}
func (m *SignalData) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalData.DiscardUnknown(m)
}

var xxx_messageInfo_SignalData proto.InternalMessageInfo

func (m *SignalData) GetBand() string {
	if m != nil {
		return m.Band
	}
	return ""
}

func (m *SignalData) GetFrequency() string {
	if m != nil {
		return m.Frequency
	}
	return ""
}

func (m *SignalData) GetPseudorange() float32 {
	if m != nil {
		return m.Pseudorange
	}
	return 0
}

func (m *SignalData) GetPhaseRange() float32 {
	if m != nil {
		return m.PhaseRange
	}
	return 0
}

func (m *SignalData) GetPhaseRangeLock() uint32 {
	if m != nil {
		return m.PhaseRangeLock
	}
	return 0
}

func (m *SignalData) GetHalfCycle() bool {
	if m != nil {
		return m.HalfCycle
	}
	return false
}

func (m *SignalData) GetSnr() float32 {
	if m != nil {
		return m.Snr
	}
	return 0
}

func (m *SignalData) GetPhaseRangeRate() float32 {
	if m != nil {
		return m.PhaseRangeRate
	}
	return 0
}

type SatelliteData struct {
	SatelliteId uint32 `protobuf:"varint,1,opt,name=satelliteId,proto3" json:"satelliteId,omitempty"`
	// TODO: Consider putting satelliteID in SignalData and removing SatelliteData
	SignalData           []*SignalData `protobuf:"bytes,2,rep,name=signalData,proto3" json:"signalData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SatelliteData) Reset()         { *m = SatelliteData{} }
func (m *SatelliteData) String() string { return proto.CompactTextString(m) }
func (*SatelliteData) ProtoMessage()    {}
func (*SatelliteData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ed04ec06c6de2a5, []int{1}
}

func (m *SatelliteData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SatelliteData.Unmarshal(m, b)
}
func (m *SatelliteData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SatelliteData.Marshal(b, m, deterministic)
}
func (m *SatelliteData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SatelliteData.Merge(m, src)
}
func (m *SatelliteData) XXX_Size() int {
	return xxx_messageInfo_SatelliteData.Size(m)
}
func (m *SatelliteData) XXX_DiscardUnknown() {
	xxx_messageInfo_SatelliteData.DiscardUnknown(m)
}

var xxx_messageInfo_SatelliteData proto.InternalMessageInfo

func (m *SatelliteData) GetSatelliteId() uint32 {
	if m != nil {
		return m.SatelliteId
	}
	return 0
}

func (m *SatelliteData) GetSignalData() []*SignalData {
	if m != nil {
		return m.SignalData
	}
	return nil
}

type ObservationSet struct {
	Epoch              *timestamp.Timestamp `protobuf:"bytes,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	ReferenceStationId string               `protobuf:"bytes,2,opt,name=referenceStationId,proto3" json:"referenceStationId,omitempty"`
	Constellation      Constellation        `protobuf:"varint,3,opt,name=constellation,proto3,enum=data.Constellation" json:"constellation,omitempty"`
	// Ignoring Clock Indicators
	SatelliteData        []*SatelliteData `protobuf:"bytes,4,rep,name=satelliteData,proto3" json:"satelliteData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ObservationSet) Reset()         { *m = ObservationSet{} }
func (m *ObservationSet) String() string { return proto.CompactTextString(m) }
func (*ObservationSet) ProtoMessage()    {}
func (*ObservationSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ed04ec06c6de2a5, []int{2}
}

func (m *ObservationSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObservationSet.Unmarshal(m, b)
}
func (m *ObservationSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObservationSet.Marshal(b, m, deterministic)
}
func (m *ObservationSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObservationSet.Merge(m, src)
}
func (m *ObservationSet) XXX_Size() int {
	return xxx_messageInfo_ObservationSet.Size(m)
}
func (m *ObservationSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ObservationSet.DiscardUnknown(m)
}

var xxx_messageInfo_ObservationSet proto.InternalMessageInfo

func (m *ObservationSet) GetEpoch() *timestamp.Timestamp {
	if m != nil {
		return m.Epoch
	}
	return nil
}

func (m *ObservationSet) GetReferenceStationId() string {
	if m != nil {
		return m.ReferenceStationId
	}
	return ""
}

func (m *ObservationSet) GetConstellation() Constellation {
	if m != nil {
		return m.Constellation
	}
	return Constellation_GPS
}

func (m *ObservationSet) GetSatelliteData() []*SatelliteData {
	if m != nil {
		return m.SatelliteData
	}
	return nil
}

func init() {
	proto.RegisterEnum("data.Constellation", Constellation_name, Constellation_value)
	proto.RegisterType((*SignalData)(nil), "data.SignalData")
	proto.RegisterType((*SatelliteData)(nil), "data.SatelliteData")
	proto.RegisterType((*ObservationSet)(nil), "data.ObservationSet")
}

func init() { proto.RegisterFile("observation.proto", fileDescriptor_5ed04ec06c6de2a5) }

var fileDescriptor_5ed04ec06c6de2a5 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6f, 0x9b, 0x40,
	0x10, 0xc5, 0x8b, 0xc1, 0x7f, 0x32, 0x08, 0x8b, 0x6e, 0x2f, 0x28, 0xaa, 0x5a, 0xe4, 0x43, 0x85,
	0x2a, 0x15, 0x47, 0xee, 0xa9, 0x47, 0x3b, 0x89, 0x22, 0x4b, 0x56, 0xdd, 0xee, 0xb6, 0x97, 0xdc,
	0x16, 0x3c, 0xc6, 0xa8, 0x78, 0x97, 0xb2, 0xeb, 0x4a, 0xb9, 0xf5, 0xe3, 0xf6, 0x63, 0x54, 0x2c,
	0x71, 0x80, 0x28, 0xb7, 0xd9, 0xdf, 0xbc, 0x1d, 0x86, 0xf7, 0x16, 0x5e, 0xcb, 0x44, 0x61, 0xf5,
	0x87, 0xeb, 0x5c, 0x8a, 0xb8, 0xac, 0xa4, 0x96, 0xc4, 0xd9, 0x71, 0xcd, 0x2f, 0xdf, 0x67, 0x52,
	0x66, 0x05, 0xce, 0x0d, 0x4b, 0x4e, 0xfb, 0xb9, 0xce, 0x8f, 0xa8, 0x34, 0x3f, 0x96, 0x8d, 0x6c,
	0xf6, 0x77, 0x00, 0xc0, 0xf2, 0x4c, 0xf0, 0xe2, 0x86, 0x6b, 0x4e, 0x08, 0x38, 0x09, 0x17, 0xbb,
	0xc0, 0x0a, 0xad, 0xe8, 0x82, 0x9a, 0x9a, 0xbc, 0x85, 0x8b, 0x7d, 0x85, 0xbf, 0x4f, 0x28, 0xd2,
	0x87, 0x60, 0x60, 0x1a, 0x2d, 0x20, 0x21, 0xb8, 0xa5, 0xc2, 0xd3, 0x4e, 0x56, 0x5c, 0x64, 0x18,
	0xd8, 0xa1, 0x15, 0x0d, 0x68, 0x17, 0x91, 0x77, 0x00, 0xe5, 0x81, 0x2b, 0xa4, 0x46, 0xe0, 0x18,
	0x41, 0x87, 0x90, 0x0f, 0x30, 0x6d, 0x4f, 0x1b, 0x99, 0xfe, 0x0a, 0x86, 0xa1, 0x15, 0x79, 0xf4,
	0x19, 0xad, 0xf7, 0x38, 0xf0, 0x62, 0x7f, 0xfd, 0x90, 0x16, 0x18, 0x8c, 0x42, 0x2b, 0x9a, 0xd0,
	0x16, 0x10, 0x1f, 0x6c, 0x25, 0xaa, 0x60, 0x6c, 0xc6, 0xd7, 0x65, 0x7f, 0x2e, 0xe5, 0x1a, 0x83,
	0x89, 0x69, 0x3e, 0xa3, 0xb3, 0x14, 0x3c, 0xc6, 0x35, 0x16, 0x45, 0xae, 0xd1, 0x98, 0x10, 0x82,
	0xab, 0xce, 0x60, 0xdd, 0x78, 0xe1, 0xd1, 0x2e, 0x22, 0x57, 0x00, 0xea, 0xc9, 0xb4, 0x60, 0x10,
	0xda, 0x91, 0xbb, 0xf0, 0xe3, 0xda, 0xf1, 0xb8, 0x35, 0x93, 0x76, 0x34, 0xb3, 0x7f, 0x16, 0x4c,
	0xb7, 0x6d, 0x48, 0x0c, 0x35, 0xb9, 0x82, 0x21, 0x96, 0x32, 0x3d, 0x98, 0x0f, 0xb8, 0x8b, 0xcb,
	0xb8, 0xc9, 0x2a, 0x3e, 0x67, 0x15, 0xff, 0x38, 0x67, 0x45, 0x1b, 0x21, 0x89, 0x81, 0x54, 0xb8,
	0xc7, 0x0a, 0x45, 0x8a, 0x4c, 0x9b, 0x41, 0xeb, 0xdd, 0x63, 0x24, 0x2f, 0x74, 0xc8, 0x17, 0xf0,
	0x52, 0x29, 0x54, 0xbd, 0xb7, 0x41, 0x26, 0x9d, 0xe9, 0xe2, 0x4d, 0xb3, 0xe9, 0x75, 0xb7, 0x45,
	0xfb, 0xca, 0xfa, 0xaa, 0xea, 0x9a, 0x12, 0x38, 0xe6, 0x27, 0x1f, 0xaf, 0xf6, 0xfc, 0xa2, 0x7d,
	0xe5, 0x47, 0x0a, 0x5e, 0x6f, 0x34, 0x19, 0x83, 0x7d, 0xf7, 0x8d, 0xf9, 0xaf, 0x88, 0x0b, 0xe3,
	0xbb, 0xcd, 0xf6, 0xeb, 0x92, 0x31, 0xdf, 0x32, 0x87, 0xe5, 0x66, 0xbd, 0xb9, 0xdd, 0xfa, 0x03,
	0x32, 0x01, 0x87, 0xad, 0x96, 0xcc, 0xb7, 0xeb, 0xea, 0xfb, 0x3d, 0x63, 0xbe, 0x43, 0x00, 0x46,
	0xab, 0xdb, 0xf5, 0xcd, 0xf6, 0xa7, 0x3f, 0x5c, 0xcd, 0xee, 0xc3, 0x2c, 0xd7, 0x87, 0x53, 0x12,
	0xa7, 0xf2, 0x38, 0xcf, 0xe4, 0xa7, 0x4c, 0x28, 0x35, 0xaf, 0x77, 0x79, 0x7a, 0xda, 0xc9, 0xc8,
	0x54, 0x9f, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xce, 0xad, 0xe7, 0xfb, 0x0d, 0x03, 0x00, 0x00,
}
