// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crosschain/last_block_height.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LastBlockHeight struct {
	Creator           string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index             string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Chain             string `protobuf:"bytes,3,opt,name=chain,proto3" json:"chain,omitempty"`
	LastSendHeight    uint64 `protobuf:"varint,4,opt,name=lastSendHeight,proto3" json:"lastSendHeight,omitempty"`
	LastReceiveHeight uint64 `protobuf:"varint,5,opt,name=lastReceiveHeight,proto3" json:"lastReceiveHeight,omitempty"`
}

func (m *LastBlockHeight) Reset()         { *m = LastBlockHeight{} }
func (m *LastBlockHeight) String() string { return proto.CompactTextString(m) }
func (*LastBlockHeight) ProtoMessage()    {}
func (*LastBlockHeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_a66188d8895bda91, []int{0}
}
func (m *LastBlockHeight) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LastBlockHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LastBlockHeight.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LastBlockHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastBlockHeight.Merge(m, src)
}
func (m *LastBlockHeight) XXX_Size() int {
	return m.Size()
}
func (m *LastBlockHeight) XXX_DiscardUnknown() {
	xxx_messageInfo_LastBlockHeight.DiscardUnknown(m)
}

var xxx_messageInfo_LastBlockHeight proto.InternalMessageInfo

func (m *LastBlockHeight) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *LastBlockHeight) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *LastBlockHeight) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *LastBlockHeight) GetLastSendHeight() uint64 {
	if m != nil {
		return m.LastSendHeight
	}
	return 0
}

func (m *LastBlockHeight) GetLastReceiveHeight() uint64 {
	if m != nil {
		return m.LastReceiveHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*LastBlockHeight)(nil), "zetachain.zetacore.crosschain.LastBlockHeight")
}

func init() {
	proto.RegisterFile("crosschain/last_block_height.proto", fileDescriptor_a66188d8895bda91)
}

var fileDescriptor_a66188d8895bda91 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0x2e, 0xca, 0x2f,
	0x2e, 0x4e, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xcf, 0x49, 0x2c, 0x2e, 0x89, 0x4f, 0xca, 0xc9, 0x4f,
	0xce, 0x8e, 0xcf, 0x48, 0xcd, 0x4c, 0xcf, 0x28, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92,
	0xad, 0x4a, 0x2d, 0x49, 0x04, 0x2b, 0xd1, 0x03, 0xb3, 0xf2, 0x8b, 0x52, 0xf5, 0x10, 0xda, 0x94,
	0xd6, 0x32, 0x72, 0xf1, 0xfb, 0x24, 0x16, 0x97, 0x38, 0x81, 0x74, 0x7a, 0x80, 0x35, 0x0a, 0x49,
	0x70, 0xb1, 0x27, 0x17, 0xa5, 0x26, 0x96, 0xe4, 0x17, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0xc1, 0xb8, 0x42, 0x22, 0x5c, 0xac, 0x99, 0x79, 0x29, 0xa9, 0x15, 0x12, 0x4c, 0x60, 0x71, 0x08,
	0x07, 0x24, 0x0a, 0x36, 0x4c, 0x82, 0x19, 0x22, 0x0a, 0xe6, 0x08, 0xa9, 0x71, 0xf1, 0x81, 0xdc,
	0x14, 0x9c, 0x9a, 0x97, 0x02, 0x31, 0x57, 0x82, 0x45, 0x81, 0x51, 0x83, 0x25, 0x08, 0x4d, 0x54,
	0x48, 0x87, 0x4b, 0x10, 0x24, 0x12, 0x94, 0x9a, 0x9c, 0x9a, 0x59, 0x96, 0x0a, 0x55, 0xca, 0x0a,
	0x56, 0x8a, 0x29, 0xe1, 0xe4, 0x7d, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e,
	0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51,
	0x86, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x20, 0x9f, 0xea, 0x42,
	0xc2, 0x05, 0xe6, 0x69, 0xfd, 0x0a, 0x7d, 0xa4, 0xd0, 0x2a, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62,
	0x03, 0x07, 0x91, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x07, 0x28, 0xe0, 0x48, 0x01, 0x00,
	0x00,
}

func (m *LastBlockHeight) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LastBlockHeight) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LastBlockHeight) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastReceiveHeight != 0 {
		i = encodeVarintLastBlockHeight(dAtA, i, uint64(m.LastReceiveHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.LastSendHeight != 0 {
		i = encodeVarintLastBlockHeight(dAtA, i, uint64(m.LastSendHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintLastBlockHeight(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintLastBlockHeight(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintLastBlockHeight(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLastBlockHeight(dAtA []byte, offset int, v uint64) int {
	offset -= sovLastBlockHeight(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LastBlockHeight) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovLastBlockHeight(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovLastBlockHeight(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovLastBlockHeight(uint64(l))
	}
	if m.LastSendHeight != 0 {
		n += 1 + sovLastBlockHeight(uint64(m.LastSendHeight))
	}
	if m.LastReceiveHeight != 0 {
		n += 1 + sovLastBlockHeight(uint64(m.LastReceiveHeight))
	}
	return n
}

func sovLastBlockHeight(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLastBlockHeight(x uint64) (n int) {
	return sovLastBlockHeight(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LastBlockHeight) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLastBlockHeight
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
			return fmt.Errorf("proto: LastBlockHeight: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LastBlockHeight: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLastBlockHeight
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
				return ErrInvalidLengthLastBlockHeight
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLastBlockHeight
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLastBlockHeight
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
				return ErrInvalidLengthLastBlockHeight
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLastBlockHeight
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLastBlockHeight
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
				return ErrInvalidLengthLastBlockHeight
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLastBlockHeight
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSendHeight", wireType)
			}
			m.LastSendHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLastBlockHeight
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastSendHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastReceiveHeight", wireType)
			}
			m.LastReceiveHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLastBlockHeight
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastReceiveHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLastBlockHeight(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLastBlockHeight
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
func skipLastBlockHeight(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLastBlockHeight
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
					return 0, ErrIntOverflowLastBlockHeight
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
					return 0, ErrIntOverflowLastBlockHeight
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
				return 0, ErrInvalidLengthLastBlockHeight
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLastBlockHeight
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLastBlockHeight
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLastBlockHeight        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLastBlockHeight          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLastBlockHeight = fmt.Errorf("proto: unexpected end of group")
)
