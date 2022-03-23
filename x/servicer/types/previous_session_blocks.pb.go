// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: servicer/previous_session_blocks.proto

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

type PreviousSessionBlocks struct {
	BlocksNum uint64 `protobuf:"varint,1,opt,name=blocksNum,proto3" json:"blocksNum,omitempty"`
}

func (m *PreviousSessionBlocks) Reset()         { *m = PreviousSessionBlocks{} }
func (m *PreviousSessionBlocks) String() string { return proto.CompactTextString(m) }
func (*PreviousSessionBlocks) ProtoMessage()    {}
func (*PreviousSessionBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_f29365c07641fad8, []int{0}
}
func (m *PreviousSessionBlocks) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PreviousSessionBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PreviousSessionBlocks.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PreviousSessionBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreviousSessionBlocks.Merge(m, src)
}
func (m *PreviousSessionBlocks) XXX_Size() int {
	return m.Size()
}
func (m *PreviousSessionBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_PreviousSessionBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_PreviousSessionBlocks proto.InternalMessageInfo

func (m *PreviousSessionBlocks) GetBlocksNum() uint64 {
	if m != nil {
		return m.BlocksNum
	}
	return 0
}

func init() {
	proto.RegisterType((*PreviousSessionBlocks)(nil), "lavanet.lava.servicer.PreviousSessionBlocks")
}

func init() {
	proto.RegisterFile("servicer/previous_session_blocks.proto", fileDescriptor_f29365c07641fad8)
}

var fileDescriptor_f29365c07641fad8 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd2, 0x2f, 0x28, 0x4a, 0x2d, 0xcb, 0xcc, 0x2f, 0x2d, 0x8e, 0x2f, 0x4e,
	0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0x8b, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0x2e, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0xcd, 0x49, 0x2c, 0x4b, 0xcc, 0x4b, 0x2d, 0xd1, 0x03, 0xd1, 0x7a, 0x30,
	0x4d, 0x4a, 0xa6, 0x5c, 0xa2, 0x01, 0x50, 0x7d, 0xc1, 0x10, 0x6d, 0x4e, 0x60, 0x5d, 0x42, 0x32,
	0x5c, 0x9c, 0x10, 0xfd, 0x7e, 0xa5, 0xb9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x08, 0x01,
	0x27, 0xa7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x48, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0x5a, 0x09, 0xa6, 0xf5, 0x2b, 0xf4, 0xe1,
	0x2e, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x3b, 0xcc, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x94, 0x23, 0x09, 0xee, 0xc2, 0x00, 0x00, 0x00,
}

func (m *PreviousSessionBlocks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PreviousSessionBlocks) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PreviousSessionBlocks) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlocksNum != 0 {
		i = encodeVarintPreviousSessionBlocks(dAtA, i, uint64(m.BlocksNum))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPreviousSessionBlocks(dAtA []byte, offset int, v uint64) int {
	offset -= sovPreviousSessionBlocks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PreviousSessionBlocks) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlocksNum != 0 {
		n += 1 + sovPreviousSessionBlocks(uint64(m.BlocksNum))
	}
	return n
}

func sovPreviousSessionBlocks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPreviousSessionBlocks(x uint64) (n int) {
	return sovPreviousSessionBlocks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PreviousSessionBlocks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPreviousSessionBlocks
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
			return fmt.Errorf("proto: PreviousSessionBlocks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PreviousSessionBlocks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksNum", wireType)
			}
			m.BlocksNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPreviousSessionBlocks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksNum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPreviousSessionBlocks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPreviousSessionBlocks
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
func skipPreviousSessionBlocks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPreviousSessionBlocks
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
					return 0, ErrIntOverflowPreviousSessionBlocks
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
					return 0, ErrIntOverflowPreviousSessionBlocks
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
				return 0, ErrInvalidLengthPreviousSessionBlocks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPreviousSessionBlocks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPreviousSessionBlocks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPreviousSessionBlocks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPreviousSessionBlocks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPreviousSessionBlocks = fmt.Errorf("proto: unexpected end of group")
)
