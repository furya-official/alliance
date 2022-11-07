// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: alliance/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	RewardDelayTime time.Duration `protobuf:"bytes,1,opt,name=reward_delay_time,json=rewardDelayTime,proto3,stdduration" json:"reward_delay_time"`
	// Time interval between consecutive applications of `take_rate`
	RewardClaimInterval time.Duration `protobuf:"bytes,2,opt,name=reward_claim_interval,json=rewardClaimInterval,proto3,stdduration" json:"reward_claim_interval"`
	// Last application of `take_rate` on assets
	LastRewardClaimTime time.Time `protobuf:"bytes,3,opt,name=last_reward_claim_time,json=lastRewardClaimTime,proto3,stdtime" json:"last_reward_claim_time"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dc4a5b6d277cc53, []int{0}
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

func (m *Params) GetRewardDelayTime() time.Duration {
	if m != nil {
		return m.RewardDelayTime
	}
	return 0
}

func (m *Params) GetRewardClaimInterval() time.Duration {
	if m != nil {
		return m.RewardClaimInterval
	}
	return 0
}

func (m *Params) GetLastRewardClaimTime() time.Time {
	if m != nil {
		return m.LastRewardClaimTime
	}
	return time.Time{}
}

type RewardHistory struct {
	Denom string                                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Index github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=index,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"index"`
}

func (m *RewardHistory) Reset()         { *m = RewardHistory{} }
func (m *RewardHistory) String() string { return proto.CompactTextString(m) }
func (*RewardHistory) ProtoMessage()    {}
func (*RewardHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dc4a5b6d277cc53, []int{1}
}
func (m *RewardHistory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardHistory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardHistory.Merge(m, src)
}
func (m *RewardHistory) XXX_Size() int {
	return m.Size()
}
func (m *RewardHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardHistory.DiscardUnknown(m)
}

var xxx_messageInfo_RewardHistory proto.InternalMessageInfo

func (m *RewardHistory) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "alliance.alliance.Params")
	proto.RegisterType((*RewardHistory)(nil), "alliance.alliance.RewardHistory")
}

func init() { proto.RegisterFile("alliance/params.proto", fileDescriptor_3dc4a5b6d277cc53) }

var fileDescriptor_3dc4a5b6d277cc53 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x4f, 0xe3, 0x30,
	0x14, 0xc7, 0x93, 0xde, 0xb5, 0xba, 0xfa, 0x74, 0x3a, 0x35, 0xd7, 0x9e, 0xda, 0x0e, 0x09, 0xea,
	0x80, 0x58, 0x9a, 0x48, 0x74, 0x43, 0x4c, 0x25, 0x03, 0x4c, 0xa0, 0xa8, 0x12, 0x82, 0x25, 0x72,
	0x13, 0x13, 0x2c, 0xe2, 0x38, 0x8a, 0x5d, 0x68, 0x27, 0xbe, 0x42, 0x25, 0x16, 0x46, 0x3e, 0x04,
	0x1f, 0xa2, 0x63, 0xc5, 0x84, 0x18, 0x0a, 0x6a, 0x17, 0x3e, 0x06, 0x8a, 0xed, 0x14, 0x01, 0x0b,
	0x53, 0xde, 0xcb, 0xff, 0xff, 0x7e, 0x7e, 0xcf, 0xcf, 0xa0, 0x01, 0xe3, 0x18, 0xc3, 0x24, 0x40,
	0x4e, 0x0a, 0x33, 0x48, 0x98, 0x9d, 0x66, 0x94, 0x53, 0xa3, 0x56, 0xfc, 0xb6, 0x8b, 0xa0, 0x5d,
	0x8f, 0x68, 0x44, 0x85, 0xea, 0xe4, 0x91, 0x34, 0xb6, 0x5b, 0x01, 0x65, 0x84, 0x32, 0x5f, 0x0a,
	0x32, 0x51, 0x92, 0x19, 0x51, 0x1a, 0xc5, 0xc8, 0x11, 0xd9, 0x70, 0x74, 0xe6, 0x84, 0xa3, 0x0c,
	0x72, 0x4c, 0x13, 0xa5, 0x5b, 0x9f, 0x75, 0x8e, 0x09, 0x62, 0x1c, 0x92, 0x54, 0x1a, 0x3a, 0x37,
	0x25, 0x50, 0x39, 0x12, 0x5d, 0x19, 0x87, 0xa0, 0x96, 0xa1, 0x2b, 0x98, 0x85, 0x7e, 0x88, 0x62,
	0x38, 0xf1, 0x73, 0x6b, 0x53, 0xdf, 0xd0, 0xb7, 0x7e, 0x6f, 0xb7, 0x6c, 0xc9, 0xb1, 0x0b, 0x8e,
	0xed, 0xaa, 0x73, 0xfa, 0xbf, 0x66, 0x0b, 0x4b, 0xbb, 0x7d, 0xb6, 0x74, 0xef, 0xaf, 0xac, 0x76,
	0xf3, 0xe2, 0x01, 0x26, 0xc8, 0x38, 0x06, 0x0d, 0x05, 0x0c, 0x62, 0x88, 0x89, 0x8f, 0x13, 0x8e,
	0xb2, 0x4b, 0x18, 0x37, 0x4b, 0xdf, 0x87, 0xfe, 0x93, 0x84, 0xbd, 0x1c, 0x70, 0xa0, 0xea, 0x8d,
	0x13, 0xf0, 0x3f, 0x86, 0x8c, 0xfb, 0x1f, 0xe8, 0xa2, 0xdd, 0x1f, 0x82, 0xdc, 0xfe, 0x42, 0x1e,
	0x14, 0x63, 0x4b, 0xf4, 0x54, 0xa0, 0x73, 0x86, 0xf7, 0x8e, 0xcf, 0x3d, 0x3b, 0x3f, 0x5f, 0xef,
	0x2c, 0xbd, 0x73, 0x0d, 0xfe, 0x48, 0x61, 0x1f, 0x33, 0x4e, 0xb3, 0x89, 0x51, 0x07, 0xe5, 0x10,
	0x25, 0x94, 0x88, 0xfb, 0xa8, 0x7a, 0x32, 0x31, 0x3c, 0x50, 0xc6, 0x49, 0x88, 0xc6, 0x62, 0xa0,
	0x6a, 0x7f, 0x37, 0x47, 0x3f, 0x2d, 0xac, 0xcd, 0x08, 0xf3, 0xf3, 0xd1, 0xd0, 0x0e, 0x28, 0x51,
	0xdb, 0x52, 0x9f, 0x2e, 0x0b, 0x2f, 0x1c, 0x3e, 0x49, 0x11, 0xb3, 0x5d, 0x14, 0x3c, 0xdc, 0x77,
	0x81, 0x5a, 0xa6, 0x8b, 0x02, 0x4f, 0xa2, 0x64, 0x03, 0xfd, 0xde, 0x6c, 0x69, 0xea, 0xf3, 0xa5,
	0xa9, 0xbf, 0x2c, 0x4d, 0x7d, 0xba, 0x32, 0xb5, 0xf9, 0xca, 0xd4, 0x1e, 0x57, 0xa6, 0x76, 0xda,
	0x5a, 0x3f, 0xa6, 0xb1, 0xb3, 0x0e, 0x05, 0x73, 0x58, 0x11, 0xe3, 0xf6, 0xde, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x28, 0x6d, 0xed, 0x34, 0x70, 0x02, 0x00, 0x00,
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
	if this.RewardDelayTime != that1.RewardDelayTime {
		return false
	}
	if this.RewardClaimInterval != that1.RewardClaimInterval {
		return false
	}
	if !this.LastRewardClaimTime.Equal(that1.LastRewardClaimTime) {
		return false
	}
	return true
}
func (this *RewardHistory) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RewardHistory)
	if !ok {
		that2, ok := that.(RewardHistory)
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
	if this.Denom != that1.Denom {
		return false
	}
	if !this.Index.Equal(that1.Index) {
		return false
	}
	return true
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
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastRewardClaimTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastRewardClaimTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.RewardClaimInterval, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.RewardClaimInterval):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintParams(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.RewardDelayTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.RewardDelayTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintParams(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *RewardHistory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardHistory) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardHistory) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Index.Size()
		i -= size
		if _, err := m.Index.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.RewardDelayTime)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.RewardClaimInterval)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastRewardClaimTime)
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *RewardHistory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.Index.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
				return fmt.Errorf("proto: wrong wireType = %d for field RewardDelayTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.RewardDelayTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardClaimInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.RewardClaimInterval, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastRewardClaimTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastRewardClaimTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *RewardHistory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: RewardHistory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardHistory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Index.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)