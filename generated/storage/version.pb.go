// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/version.proto

package storage

import (
	fmt "fmt"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Version struct {
	// This is a strictly incrementing integer describing the DB version.
	SeqNum int32 `protobuf:"varint,1,opt,name=seq_num,json=seqNum,proto3" json:"seq_num,omitempty"`
	// Associated version metadata. (For example, the corresponding product version.)
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Last time version was updated
	LastPersisted        *types.Timestamp `protobuf:"bytes,3,opt,name=last_persisted,json=lastPersisted,proto3" json:"last_persisted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6abc271df6a8e78, []int{0}
}
func (m *Version) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Version.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return m.Size()
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetSeqNum() int32 {
	if m != nil {
		return m.SeqNum
	}
	return 0
}

func (m *Version) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Version) GetLastPersisted() *types.Timestamp {
	if m != nil {
		return m.LastPersisted
	}
	return nil
}

func (m *Version) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Version) Clone() *Version {
	if m == nil {
		return nil
	}
	cloned := new(Version)
	*cloned = *m

	cloned.LastPersisted = m.LastPersisted.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*Version)(nil), "storage.Version")
}

func init() { proto.RegisterFile("storage/version.proto", fileDescriptor_e6abc271df6a8e78) }

var fileDescriptor_e6abc271df6a8e78 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x2f, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x87, 0x0a, 0x4b, 0xc9, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0x85,
	0x93, 0x4a, 0xd3, 0xf4, 0x4b, 0x32, 0x73, 0x53, 0x8b, 0x4b, 0x12, 0x73, 0x0b, 0x20, 0x2a, 0x95,
	0xea, 0xb9, 0xd8, 0xc3, 0x20, 0x5a, 0x85, 0xc4, 0xb9, 0xd8, 0x8b, 0x53, 0x0b, 0xe3, 0xf3, 0x4a,
	0x73, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xd8, 0x8a, 0x53, 0x0b, 0xfd, 0x4a, 0x73, 0x85,
	0x24, 0xb8, 0xd8, 0xa1, 0xc6, 0x4b, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42, 0x8e,
	0x5c, 0x7c, 0x39, 0x89, 0xc5, 0x25, 0xf1, 0x05, 0x20, 0x7e, 0x71, 0x49, 0x6a, 0x8a, 0x04, 0xb3,
	0x02, 0xa3, 0x06, 0xb7, 0x91, 0x94, 0x1e, 0xc4, 0x5e, 0x3d, 0x98, 0xbd, 0x7a, 0x21, 0x30, 0x7b,
	0x83, 0x78, 0x41, 0x3a, 0x02, 0x60, 0x1a, 0x9c, 0x4c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48,
	0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0xb8, 0x24, 0x33, 0xf3, 0xf5, 0x8a,
	0x4b, 0x12, 0x93, 0xb3, 0x8b, 0xf2, 0x2b, 0x20, 0x06, 0xe8, 0x41, 0xbd, 0x13, 0x05, 0xf3, 0x57,
	0x12, 0x1b, 0x58, 0xdc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x45, 0x11, 0x3d, 0x54, 0x00, 0x01,
	0x00, 0x00,
}

func (m *Version) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Version) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Version) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.LastPersisted != nil {
		{
			size, err := m.LastPersisted.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVersion(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintVersion(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if m.SeqNum != 0 {
		i = encodeVarintVersion(dAtA, i, uint64(m.SeqNum))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVersion(dAtA []byte, offset int, v uint64) int {
	offset -= sovVersion(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Version) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SeqNum != 0 {
		n += 1 + sovVersion(uint64(m.SeqNum))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovVersion(uint64(l))
	}
	if m.LastPersisted != nil {
		l = m.LastPersisted.Size()
		n += 1 + l + sovVersion(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovVersion(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVersion(x uint64) (n int) {
	return sovVersion(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Version) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVersion
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
			return fmt.Errorf("proto: Version: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Version: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeqNum", wireType)
			}
			m.SeqNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersion
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SeqNum |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersion
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
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastPersisted", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersion
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
				return ErrInvalidLengthVersion
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastPersisted == nil {
				m.LastPersisted = &types.Timestamp{}
			}
			if err := m.LastPersisted.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVersion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVersion
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipVersion(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVersion
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
					return 0, ErrIntOverflowVersion
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
					return 0, ErrIntOverflowVersion
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
				return 0, ErrInvalidLengthVersion
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVersion
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVersion
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVersion        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVersion          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVersion = fmt.Errorf("proto: unexpected end of group")
)
