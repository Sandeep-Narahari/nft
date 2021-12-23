// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nft/v1beta1/market_place.proto

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

type MarketPlace struct {
	NftId   string `protobuf:"bytes,1,opt,name=nftId,proto3" json:"nftId,omitempty" yaml:"NFTID"`
	DenomID string `protobuf:"bytes,2,opt,name=denomID,proto3" json:"denomID,omitempty" yaml:"denom_id"`
	Price   string `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Seller  string `protobuf:"bytes,4,opt,name=seller,proto3" json:"seller,omitempty"`
	Buyer   string `protobuf:"bytes,5,opt,name=buyer,proto3" json:"buyer,omitempty"`
}

func (m *MarketPlace) Reset()         { *m = MarketPlace{} }
func (m *MarketPlace) String() string { return proto.CompactTextString(m) }
func (*MarketPlace) ProtoMessage()    {}
func (*MarketPlace) Descriptor() ([]byte, []int) {
	return fileDescriptor_79b3c3c94d423baa, []int{0}
}
func (m *MarketPlace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketPlace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketPlace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketPlace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketPlace.Merge(m, src)
}
func (m *MarketPlace) XXX_Size() int {
	return m.Size()
}
func (m *MarketPlace) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketPlace.DiscardUnknown(m)
}

var xxx_messageInfo_MarketPlace proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MarketPlace)(nil), "nft.v1beta1.MarketPlace")
}

func init() { proto.RegisterFile("nft/v1beta1/market_place.proto", fileDescriptor_79b3c3c94d423baa) }

var fileDescriptor_79b3c3c94d423baa = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0x4b, 0x2b, 0xd1,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0xcf, 0x4d, 0x2c, 0xca, 0x4e, 0x2d, 0x89, 0x2f,
	0xc8, 0x49, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0x4b, 0x2b, 0xd1,
	0x83, 0xca, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0xc5, 0xf5, 0x41, 0x2c, 0x88, 0x12, 0xa5,
	0x15, 0x8c, 0x5c, 0xdc, 0xbe, 0x60, 0x9d, 0x01, 0x20, 0x8d, 0x42, 0x6a, 0x5c, 0xac, 0x79, 0x69,
	0x25, 0x9e, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x4e, 0x02, 0x9f, 0xee, 0xc9, 0xf3, 0x54,
	0x26, 0xe6, 0xe6, 0x58, 0x29, 0xf9, 0xb9, 0x85, 0x78, 0xba, 0x28, 0x05, 0x41, 0xa4, 0x85, 0x74,
	0xb9, 0xd8, 0x53, 0x52, 0xf3, 0xf2, 0x73, 0x3d, 0x5d, 0x24, 0x98, 0xc0, 0x2a, 0x85, 0x3f, 0xdd,
	0x93, 0xe7, 0x87, 0xa8, 0x04, 0x4b, 0xc4, 0x67, 0xa6, 0x28, 0x05, 0xc1, 0xd4, 0x08, 0x89, 0x70,
	0xb1, 0x16, 0x14, 0x65, 0x26, 0xa7, 0x4a, 0x30, 0x83, 0x14, 0x07, 0x41, 0x38, 0x42, 0x62, 0x5c,
	0x6c, 0xc5, 0xa9, 0x39, 0x39, 0xa9, 0x45, 0x12, 0x2c, 0x60, 0x61, 0x28, 0x0f, 0xa4, 0x3a, 0xa9,
	0xb4, 0x32, 0xb5, 0x48, 0x82, 0x15, 0xa2, 0x1a, 0xcc, 0x71, 0x72, 0x3a, 0xf1, 0x50, 0x8e, 0xe1,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x54, 0xd2, 0x33, 0x4b, 0x32, 0x4a,
	0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x1d, 0x4b, 0x4b, 0xf2, 0xf3, 0xf2, 0x73, 0x2b, 0xfd, 0x52,
	0x4b, 0xca, 0xf3, 0x8b, 0xb2, 0xf5, 0x41, 0xc1, 0x54, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06,
	0xf6, 0xb5, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x9c, 0x0a, 0x5a, 0x3a, 0x01, 0x00, 0x00,
}

func (m *MarketPlace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketPlace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketPlace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Buyer) > 0 {
		i -= len(m.Buyer)
		copy(dAtA[i:], m.Buyer)
		i = encodeVarintMarketPlace(dAtA, i, uint64(len(m.Buyer)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Seller) > 0 {
		i -= len(m.Seller)
		copy(dAtA[i:], m.Seller)
		i = encodeVarintMarketPlace(dAtA, i, uint64(len(m.Seller)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintMarketPlace(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DenomID) > 0 {
		i -= len(m.DenomID)
		copy(dAtA[i:], m.DenomID)
		i = encodeVarintMarketPlace(dAtA, i, uint64(len(m.DenomID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.NftId) > 0 {
		i -= len(m.NftId)
		copy(dAtA[i:], m.NftId)
		i = encodeVarintMarketPlace(dAtA, i, uint64(len(m.NftId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMarketPlace(dAtA []byte, offset int, v uint64) int {
	offset -= sovMarketPlace(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MarketPlace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftId)
	if l > 0 {
		n += 1 + l + sovMarketPlace(uint64(l))
	}
	l = len(m.DenomID)
	if l > 0 {
		n += 1 + l + sovMarketPlace(uint64(l))
	}
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovMarketPlace(uint64(l))
	}
	l = len(m.Seller)
	if l > 0 {
		n += 1 + l + sovMarketPlace(uint64(l))
	}
	l = len(m.Buyer)
	if l > 0 {
		n += 1 + l + sovMarketPlace(uint64(l))
	}
	return n
}

func sovMarketPlace(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMarketPlace(x uint64) (n int) {
	return sovMarketPlace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MarketPlace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarketPlace
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
			return fmt.Errorf("proto: MarketPlace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketPlace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketPlace
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
				return ErrInvalidLengthMarketPlace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketPlace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketPlace
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
				return ErrInvalidLengthMarketPlace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketPlace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketPlace
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
				return ErrInvalidLengthMarketPlace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketPlace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketPlace
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
				return ErrInvalidLengthMarketPlace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketPlace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Buyer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketPlace
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
				return ErrInvalidLengthMarketPlace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketPlace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Buyer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarketPlace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarketPlace
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
func skipMarketPlace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMarketPlace
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
					return 0, ErrIntOverflowMarketPlace
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
					return 0, ErrIntOverflowMarketPlace
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
				return 0, ErrInvalidLengthMarketPlace
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMarketPlace
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMarketPlace
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMarketPlace        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMarketPlace          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMarketPlace = fmt.Errorf("proto: unexpected end of group")
)
