// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/evidence/types/types.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgSubmitEvidenceBase defines an sdk.Msg type that supports submitting arbitrary
// Evidence.
//
// Note, this message type provides the basis for which a true MsgSubmitEvidence
// can be constructed. Since the evidence submitted in the message can be arbitrary,
// assuming it fulfills the Evidence interface, it must be defined at the
// application-level and extend MsgSubmitEvidenceBase.
type MsgSubmitEvidenceBase struct {
	Submitter github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=submitter,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"submitter,omitempty"`
}

func (m *MsgSubmitEvidenceBase) Reset()         { *m = MsgSubmitEvidenceBase{} }
func (m *MsgSubmitEvidenceBase) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitEvidenceBase) ProtoMessage()    {}
func (*MsgSubmitEvidenceBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_72113e6a7b2536ae, []int{0}
}
func (m *MsgSubmitEvidenceBase) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitEvidenceBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitEvidenceBase.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitEvidenceBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitEvidenceBase.Merge(m, src)
}
func (m *MsgSubmitEvidenceBase) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitEvidenceBase) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitEvidenceBase.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitEvidenceBase proto.InternalMessageInfo

func (m *MsgSubmitEvidenceBase) GetSubmitter() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Submitter
	}
	return nil
}

func (*MsgSubmitEvidenceBase) XXX_MessageName() string {
	return "cosmos_sdk.x.evidence.v1.MsgSubmitEvidenceBase"
}

// Equivocation implements the Evidence interface and defines evidence of double
// signing misbehavior.
type Equivocation struct {
	Height           int64                                          `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Time             time.Time                                      `protobuf:"bytes,2,opt,name=time,proto3,stdtime" json:"time"`
	Power            int64                                          `protobuf:"varint,3,opt,name=power,proto3" json:"power,omitempty"`
	ConsensusAddress github_com_cosmos_cosmos_sdk_types.ConsAddress `protobuf:"bytes,4,opt,name=consensus_address,json=consensusAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ConsAddress" json:"consensus_address,omitempty" yaml:"consensus_address"`
}

func (m *Equivocation) Reset()      { *m = Equivocation{} }
func (*Equivocation) ProtoMessage() {}
func (*Equivocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_72113e6a7b2536ae, []int{1}
}
func (m *Equivocation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Equivocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Equivocation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Equivocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Equivocation.Merge(m, src)
}
func (m *Equivocation) XXX_Size() int {
	return m.Size()
}
func (m *Equivocation) XXX_DiscardUnknown() {
	xxx_messageInfo_Equivocation.DiscardUnknown(m)
}

var xxx_messageInfo_Equivocation proto.InternalMessageInfo

func (*Equivocation) XXX_MessageName() string {
	return "cosmos_sdk.x.evidence.v1.Equivocation"
}

// Params defines the total set of parameters for the evidence module
type Params struct {
	MaxEvidenceAge time.Duration `protobuf:"bytes,1,opt,name=max_evidence_age,json=maxEvidenceAge,proto3,stdduration" json:"max_evidence_age" yaml:"max_evidence_age"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_72113e6a7b2536ae, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxEvidenceAge() time.Duration {
	if m != nil {
		return m.MaxEvidenceAge
	}
	return 0
}

func (*Params) XXX_MessageName() string {
	return "cosmos_sdk.x.evidence.v1.Params"
}
func init() {
	proto.RegisterType((*MsgSubmitEvidenceBase)(nil), "cosmos_sdk.x.evidence.v1.MsgSubmitEvidenceBase")
	golang_proto.RegisterType((*MsgSubmitEvidenceBase)(nil), "cosmos_sdk.x.evidence.v1.MsgSubmitEvidenceBase")
	proto.RegisterType((*Equivocation)(nil), "cosmos_sdk.x.evidence.v1.Equivocation")
	golang_proto.RegisterType((*Equivocation)(nil), "cosmos_sdk.x.evidence.v1.Equivocation")
	proto.RegisterType((*Params)(nil), "cosmos_sdk.x.evidence.v1.Params")
	golang_proto.RegisterType((*Params)(nil), "cosmos_sdk.x.evidence.v1.Params")
}

func init() { proto.RegisterFile("x/evidence/types/types.proto", fileDescriptor_72113e6a7b2536ae) }
func init() {
	golang_proto.RegisterFile("x/evidence/types/types.proto", fileDescriptor_72113e6a7b2536ae)
}

var fileDescriptor_72113e6a7b2536ae = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3d, 0x6f, 0xd4, 0x30,
	0x1c, 0xc6, 0x63, 0x7a, 0x9c, 0x8a, 0x5b, 0xa1, 0x12, 0xf1, 0x12, 0x4e, 0xc8, 0xa9, 0x82, 0x84,
	0xba, 0xd4, 0x51, 0xcb, 0x82, 0x6e, 0xbb, 0x40, 0x47, 0x5e, 0x74, 0x65, 0x62, 0x89, 0x7c, 0x89,
	0x71, 0xa2, 0xd6, 0x71, 0x88, 0x9d, 0x23, 0x27, 0xbe, 0x00, 0x63, 0xc7, 0x8e, 0x1d, 0xf9, 0x18,
	0x8c, 0x37, 0x76, 0x64, 0x0a, 0xe8, 0xf2, 0x0d, 0x3a, 0x56, 0x42, 0x42, 0x67, 0x27, 0x54, 0xba,
	0x93, 0x10, 0x4b, 0x12, 0x3b, 0x4f, 0x9e, 0x3c, 0xbf, 0xff, 0x63, 0xf8, 0xa4, 0xf2, 0xe9, 0x34,
	0x8d, 0x69, 0x16, 0x51, 0x5f, 0xcd, 0x72, 0x2a, 0xcd, 0x15, 0xe7, 0x85, 0x50, 0xc2, 0x76, 0x22,
	0x21, 0xb9, 0x90, 0xa1, 0x8c, 0x4f, 0x70, 0x85, 0x3b, 0x21, 0x9e, 0x1e, 0x0c, 0x9e, 0xa9, 0x24,
	0x2d, 0xe2, 0x30, 0x27, 0x85, 0x9a, 0xf9, 0x5a, 0xec, 0x33, 0xc1, 0xc4, 0xcd, 0x93, 0x71, 0x18,
	0xb8, 0x4c, 0x08, 0x76, 0x4a, 0x8d, 0x64, 0x52, 0x7e, 0xf4, 0x55, 0xca, 0xa9, 0x54, 0x84, 0xe7,
	0xad, 0x00, 0xad, 0x0a, 0xe2, 0xb2, 0x20, 0x2a, 0x15, 0x99, 0x79, 0xef, 0x25, 0xf0, 0xc1, 0x6b,
	0xc9, 0x8e, 0xcb, 0x09, 0x4f, 0xd5, 0x51, 0x1b, 0x20, 0x20, 0x92, 0xda, 0x6f, 0xe1, 0x1d, 0xa9,
	0x77, 0x15, 0x2d, 0x1c, 0xb0, 0x0b, 0xf6, 0xb6, 0x83, 0x83, 0xeb, 0xda, 0xdd, 0x67, 0xa9, 0x4a,
	0xca, 0x09, 0x8e, 0x04, 0xf7, 0x4d, 0xfa, 0xf6, 0xb6, 0x2f, 0xe3, 0x93, 0x16, 0x6e, 0x14, 0x45,
	0xa3, 0x38, 0x2e, 0xa8, 0x94, 0xe3, 0x1b, 0x0f, 0xef, 0x37, 0x80, 0xdb, 0x47, 0x9f, 0xca, 0x74,
	0x2a, 0x22, 0x1d, 0xc0, 0x7e, 0x08, 0xfb, 0x09, 0x4d, 0x59, 0xa2, 0xb4, 0xfd, 0xc6, 0xb8, 0x5d,
	0xd9, 0x2f, 0x60, 0x6f, 0x49, 0xe1, 0xdc, 0xda, 0x05, 0x7b, 0x5b, 0x87, 0x03, 0x6c, 0x08, 0x70,
	0x47, 0x80, 0xdf, 0x77, 0x88, 0xc1, 0xe6, 0xbc, 0x76, 0xad, 0xb3, 0x9f, 0x2e, 0x18, 0xeb, 0x2f,
	0xec, 0xfb, 0xf0, 0x76, 0x2e, 0x3e, 0xd3, 0xc2, 0xd9, 0xd0, 0x86, 0x66, 0x61, 0x7f, 0x81, 0xf7,
	0x22, 0x91, 0x49, 0x9a, 0xc9, 0x52, 0x86, 0xc4, 0x04, 0x73, 0x7a, 0x9a, 0xe8, 0xcd, 0x55, 0xed,
	0x3a, 0x33, 0xc2, 0x4f, 0x87, 0xde, 0x9a, 0xc4, 0xbb, 0xae, 0x5d, 0xfc, 0x1f, 0xb4, 0x2f, 0x45,
	0x26, 0x3b, 0xdc, 0x9d, 0xbf, 0x2e, 0xed, 0xce, 0x70, 0xf3, 0xeb, 0x85, 0x6b, 0x9d, 0x5f, 0xb8,
	0x96, 0x57, 0xc1, 0xfe, 0x3b, 0x52, 0x10, 0x2e, 0xed, 0x04, 0xee, 0x70, 0x52, 0x85, 0x5d, 0xdf,
	0x21, 0x61, 0x54, 0x8f, 0x60, 0xeb, 0xf0, 0xf1, 0x1a, 0xec, 0xab, 0xb6, 0xae, 0xe0, 0xe9, 0x92,
	0xf5, 0xaa, 0x76, 0x1f, 0x99, 0xb8, 0xab, 0x06, 0xde, 0xf9, 0x72, 0x0c, 0x77, 0x39, 0xa9, 0xba,
	0x16, 0x47, 0x8c, 0x0e, 0x7b, 0xcb, 0x3f, 0x07, 0xc7, 0xdf, 0x16, 0x08, 0xcc, 0x17, 0x08, 0x5c,
	0x2e, 0x10, 0xf8, 0xb5, 0x40, 0xe0, 0xac, 0x41, 0xd6, 0xf7, 0x06, 0x81, 0x79, 0x83, 0xc0, 0x65,
	0x83, 0xac, 0x1f, 0x0d, 0xb2, 0x3e, 0xfc, 0xbb, 0xd9, 0xd5, 0x73, 0x3c, 0xe9, 0xeb, 0x88, 0xcf,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xbd, 0x24, 0x2d, 0xe2, 0x02, 0x00, 0x00,
}

func (this *MsgSubmitEvidenceBase) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgSubmitEvidenceBase)
	if !ok {
		that2, ok := that.(MsgSubmitEvidenceBase)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Submitter, that1.Submitter) {
		return false
	}
	return true
}
func (this *Equivocation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Equivocation)
	if !ok {
		that2, ok := that.(Equivocation)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Height != that1.Height {
		return false
	}
	if !this.Time.Equal(that1.Time) {
		return false
	}
	if this.Power != that1.Power {
		return false
	}
	if !bytes.Equal(this.ConsensusAddress, that1.ConsensusAddress) {
		return false
	}
	return true
}
func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MaxEvidenceAge != that1.MaxEvidenceAge {
		return false
	}
	return true
}
func (m *MsgSubmitEvidenceBase) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitEvidenceBase) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitEvidenceBase) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Submitter) > 0 {
		i -= len(m.Submitter)
		copy(dAtA[i:], m.Submitter)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Submitter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Equivocation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Equivocation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Equivocation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConsensusAddress) > 0 {
		i -= len(m.ConsensusAddress)
		copy(dAtA[i:], m.ConsensusAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ConsensusAddress)))
		i--
		dAtA[i] = 0x22
	}
	if m.Power != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x18
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTypes(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.MaxEvidenceAge, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxEvidenceAge):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTypes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSubmitEvidenceBase) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Submitter)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Equivocation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovTypes(uint64(l))
	if m.Power != 0 {
		n += 1 + sovTypes(uint64(m.Power))
	}
	l = len(m.ConsensusAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxEvidenceAge)
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSubmitEvidenceBase) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitEvidenceBase: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitEvidenceBase: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Submitter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Submitter = append(m.Submitter[:0], dAtA[iNdEx:postIndex]...)
			if m.Submitter == nil {
				m.Submitter = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Equivocation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Equivocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Equivocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsensusAddress = append(m.ConsensusAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ConsensusAddress == nil {
				m.ConsensusAddress = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEvidenceAge", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.MaxEvidenceAge, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
