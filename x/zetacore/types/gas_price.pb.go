// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetacore/gas_price.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type GasPrice struct {
	Creator     string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index       string   `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Chain       string   `protobuf:"bytes,3,opt,name=chain,proto3" json:"chain,omitempty"`
	Signers     []string `protobuf:"bytes,4,rep,name=signers,proto3" json:"signers,omitempty"`
	BlockNums   []uint64 `protobuf:"varint,5,rep,packed,name=blockNums,proto3" json:"blockNums,omitempty"`
	Prices      []uint64 `protobuf:"varint,6,rep,packed,name=prices,proto3" json:"prices,omitempty"`
	MedianIndex uint64   `protobuf:"varint,7,opt,name=medianIndex,proto3" json:"medianIndex,omitempty"`
}

func (m *GasPrice) Reset()         { *m = GasPrice{} }
func (m *GasPrice) String() string { return proto.CompactTextString(m) }
func (*GasPrice) ProtoMessage()    {}
func (*GasPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ad54f888769117a, []int{0}
}
func (m *GasPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GasPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GasPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GasPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasPrice.Merge(m, src)
}
func (m *GasPrice) XXX_Size() int {
	return m.Size()
}
func (m *GasPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_GasPrice.DiscardUnknown(m)
}

var xxx_messageInfo_GasPrice proto.InternalMessageInfo

func (m *GasPrice) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *GasPrice) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *GasPrice) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *GasPrice) GetSigners() []string {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (m *GasPrice) GetBlockNums() []uint64 {
	if m != nil {
		return m.BlockNums
	}
	return nil
}

func (m *GasPrice) GetPrices() []uint64 {
	if m != nil {
		return m.Prices
	}
	return nil
}

func (m *GasPrice) GetMedianIndex() uint64 {
	if m != nil {
		return m.MedianIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*GasPrice)(nil), "zetachain.zetacore.zetacore.GasPrice")
}

func init() { proto.RegisterFile("zetacore/gas_price.proto", fileDescriptor_9ad54f888769117a) }

var fileDescriptor_9ad54f888769117a = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x63, 0x92, 0xa6, 0xd4, 0x6c, 0x56, 0x85, 0x2c, 0x40, 0x56, 0xc4, 0x94, 0x85, 0x64,
	0xe0, 0x0d, 0x58, 0x50, 0x17, 0x84, 0x32, 0xb2, 0x20, 0xc7, 0xb5, 0x5c, 0x0b, 0x12, 0x47, 0xb6,
	0x2b, 0x15, 0x9e, 0x82, 0x67, 0x62, 0x62, 0xec, 0xc8, 0x88, 0x92, 0x17, 0x41, 0xbe, 0xe9, 0xdf,
	0x76, 0xce, 0x77, 0x6f, 0x4e, 0xae, 0x0f, 0xa6, 0x9f, 0xd2, 0x73, 0x61, 0xac, 0x2c, 0x15, 0x77,
	0xaf, 0x9d, 0xd5, 0x42, 0x16, 0x9d, 0x35, 0xde, 0x90, 0x6b, 0x98, 0xac, 0xb8, 0x6e, 0x8b, 0xfd,
	0xce, 0x41, 0x5c, 0xcd, 0x95, 0x51, 0x06, 0xf6, 0xca, 0xa0, 0xc6, 0x4f, 0x6e, 0xbf, 0x11, 0x3e,
	0x7f, 0xe4, 0xee, 0x39, 0xa4, 0x10, 0x8a, 0xa7, 0xc2, 0x4a, 0xee, 0x8d, 0xa5, 0x28, 0x43, 0xf9,
	0xac, 0xda, 0x5b, 0x32, 0xc7, 0x13, 0xdd, 0x2e, 0xe5, 0x86, 0x9e, 0x01, 0x1f, 0x4d, 0xa0, 0xf0,
	0x37, 0x1a, 0x8f, 0x14, 0x4c, 0x48, 0x71, 0x5a, 0xb5, 0xd2, 0x3a, 0x9a, 0x64, 0x71, 0x48, 0xd9,
	0x59, 0x72, 0x83, 0x67, 0xf5, 0xbb, 0x11, 0x6f, 0x4f, 0xeb, 0xc6, 0xd1, 0x49, 0x16, 0xe7, 0x49,
	0x75, 0x04, 0xe4, 0x12, 0xa7, 0xf0, 0x18, 0x47, 0x53, 0x18, 0xed, 0x1c, 0xc9, 0xf0, 0x45, 0x23,
	0x97, 0x9a, 0xb7, 0x0b, 0xb8, 0x60, 0x9a, 0xa1, 0x3c, 0xa9, 0x4e, 0xd1, 0xc3, 0xe2, 0xa7, 0x67,
	0x68, 0xdb, 0x33, 0xf4, 0xd7, 0x33, 0xf4, 0x35, 0xb0, 0x68, 0x3b, 0xb0, 0xe8, 0x77, 0x60, 0xd1,
	0x4b, 0xa9, 0xb4, 0x5f, 0xad, 0xeb, 0x42, 0x98, 0xa6, 0x0c, 0x4d, 0xdc, 0xc1, 0x89, 0xe5, 0xa1,
	0xc1, 0xcd, 0x51, 0xfa, 0x8f, 0x4e, 0xba, 0x3a, 0x85, 0x5a, 0xee, 0xff, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x69, 0x6d, 0x6a, 0x29, 0x65, 0x01, 0x00, 0x00,
}

func (m *GasPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MedianIndex != 0 {
		i = encodeVarintGasPrice(dAtA, i, uint64(m.MedianIndex))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Prices) > 0 {
		dAtA2 := make([]byte, len(m.Prices)*10)
		var j1 int
		for _, num := range m.Prices {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintGasPrice(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x32
	}
	if len(m.BlockNums) > 0 {
		dAtA4 := make([]byte, len(m.BlockNums)*10)
		var j3 int
		for _, num := range m.BlockNums {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintGasPrice(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintGasPrice(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintGasPrice(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintGasPrice(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintGasPrice(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGasPrice(dAtA []byte, offset int, v uint64) int {
	offset -= sovGasPrice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GasPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovGasPrice(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovGasPrice(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovGasPrice(uint64(l))
	}
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovGasPrice(uint64(l))
		}
	}
	if len(m.BlockNums) > 0 {
		l = 0
		for _, e := range m.BlockNums {
			l += sovGasPrice(uint64(e))
		}
		n += 1 + sovGasPrice(uint64(l)) + l
	}
	if len(m.Prices) > 0 {
		l = 0
		for _, e := range m.Prices {
			l += sovGasPrice(uint64(e))
		}
		n += 1 + sovGasPrice(uint64(l)) + l
	}
	if m.MedianIndex != 0 {
		n += 1 + sovGasPrice(uint64(m.MedianIndex))
	}
	return n
}

func sovGasPrice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGasPrice(x uint64) (n int) {
	return sovGasPrice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GasPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGasPrice
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
			return fmt.Errorf("proto: GasPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGasPrice
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
				return ErrInvalidLengthGasPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGasPrice
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
					return ErrIntOverflowGasPrice
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
				return ErrInvalidLengthGasPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGasPrice
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
					return ErrIntOverflowGasPrice
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
				return ErrInvalidLengthGasPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGasPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGasPrice
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
				return ErrInvalidLengthGasPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGasPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGasPrice
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.BlockNums = append(m.BlockNums, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGasPrice
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGasPrice
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGasPrice
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.BlockNums) == 0 {
					m.BlockNums = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGasPrice
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.BlockNums = append(m.BlockNums, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNums", wireType)
			}
		case 6:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGasPrice
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Prices = append(m.Prices, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGasPrice
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGasPrice
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGasPrice
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Prices) == 0 {
					m.Prices = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGasPrice
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Prices = append(m.Prices, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Prices", wireType)
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MedianIndex", wireType)
			}
			m.MedianIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGasPrice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MedianIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGasPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGasPrice
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
func skipGasPrice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGasPrice
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
					return 0, ErrIntOverflowGasPrice
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
					return 0, ErrIntOverflowGasPrice
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
				return 0, ErrInvalidLengthGasPrice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGasPrice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGasPrice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGasPrice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGasPrice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGasPrice = fmt.Errorf("proto: unexpected end of group")
)
