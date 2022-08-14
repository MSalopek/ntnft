// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ntnft/owner.proto

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

type OwnerCollection struct {
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	Token   *NtNft `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *OwnerCollection) Reset()         { *m = OwnerCollection{} }
func (m *OwnerCollection) String() string { return proto.CompactTextString(m) }
func (*OwnerCollection) ProtoMessage()    {}
func (*OwnerCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_023d566f9fd62b98, []int{0}
}
func (m *OwnerCollection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OwnerCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OwnerCollection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OwnerCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerCollection.Merge(m, src)
}
func (m *OwnerCollection) XXX_Size() int {
	return m.Size()
}
func (m *OwnerCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerCollection.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerCollection proto.InternalMessageInfo

func (m *OwnerCollection) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *OwnerCollection) GetToken() *NtNft {
	if m != nil {
		return m.Token
	}
	return nil
}

type Owner struct {
	Index      string             `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Address    string             `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Collection []*OwnerCollection `protobuf:"bytes,3,rep,name=collection,proto3" json:"collection,omitempty"`
}

func (m *Owner) Reset()         { *m = Owner{} }
func (m *Owner) String() string { return proto.CompactTextString(m) }
func (*Owner) ProtoMessage()    {}
func (*Owner) Descriptor() ([]byte, []int) {
	return fileDescriptor_023d566f9fd62b98, []int{1}
}
func (m *Owner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Owner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Owner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Owner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Owner.Merge(m, src)
}
func (m *Owner) XXX_Size() int {
	return m.Size()
}
func (m *Owner) XXX_DiscardUnknown() {
	xxx_messageInfo_Owner.DiscardUnknown(m)
}

var xxx_messageInfo_Owner proto.InternalMessageInfo

func (m *Owner) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Owner) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Owner) GetCollection() []*OwnerCollection {
	if m != nil {
		return m.Collection
	}
	return nil
}

func init() {
	proto.RegisterType((*OwnerCollection)(nil), "ntnft.ntnft.OwnerCollection")
	proto.RegisterType((*Owner)(nil), "ntnft.ntnft.Owner")
}

func init() { proto.RegisterFile("ntnft/owner.proto", fileDescriptor_023d566f9fd62b98) }

var fileDescriptor_023d566f9fd62b98 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x2b, 0xc9, 0x4b,
	0x2b, 0xd1, 0xcf, 0x2f, 0xcf, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06,
	0x0b, 0xe9, 0x81, 0x49, 0x29, 0x21, 0x88, 0x7c, 0x5e, 0x49, 0x3c, 0x48, 0x0c, 0xac, 0x40, 0x29,
	0x8c, 0x8b, 0xdf, 0x1f, 0xa4, 0xde, 0x39, 0x3f, 0x27, 0x27, 0x35, 0xb9, 0x24, 0x33, 0x3f, 0x4f,
	0x48, 0x92, 0x8b, 0x23, 0x39, 0x27, 0xb1, 0xb8, 0x38, 0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51,
	0x83, 0x33, 0x88, 0x1d, 0xcc, 0xf7, 0x4c, 0x11, 0xd2, 0xe0, 0x62, 0x2d, 0xc9, 0xcf, 0x4e, 0xcd,
	0x93, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0xd2, 0x43, 0x32, 0x5e, 0xcf, 0xaf, 0xc4, 0x2f,
	0xad, 0x24, 0x08, 0xa2, 0x40, 0xa9, 0x94, 0x8b, 0x15, 0x6c, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x66,
	0x5e, 0x4a, 0x6a, 0x05, 0xd4, 0x28, 0x08, 0x47, 0x48, 0x82, 0x8b, 0x3d, 0x31, 0x25, 0xa5, 0x28,
	0xb5, 0xb8, 0x18, 0x6c, 0x14, 0x67, 0x10, 0x8c, 0x2b, 0x64, 0xc3, 0xc5, 0x95, 0x0c, 0x77, 0x8b,
	0x04, 0xb3, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x0c, 0x8a, 0x3d, 0x68, 0xee, 0x0d, 0x42, 0x52, 0xef,
	0xa4, 0x77, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x22, 0x79, 0x25, 0xba,
	0x20, 0xdf, 0x57, 0xe8, 0x43, 0x42, 0xa1, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x1c, 0x0a,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x94, 0xd2, 0xa3, 0x3b, 0x01, 0x00, 0x00,
}

func (m *OwnerCollection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OwnerCollection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OwnerCollection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Token != nil {
		{
			size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOwner(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintOwner(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Owner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Owner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Owner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Collection) > 0 {
		for iNdEx := len(m.Collection) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Collection[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOwner(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintOwner(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintOwner(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOwner(dAtA []byte, offset int, v uint64) int {
	offset -= sovOwner(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OwnerCollection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovOwner(uint64(l))
	}
	if m.Token != nil {
		l = m.Token.Size()
		n += 1 + l + sovOwner(uint64(l))
	}
	return n
}

func (m *Owner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovOwner(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovOwner(uint64(l))
	}
	if len(m.Collection) > 0 {
		for _, e := range m.Collection {
			l = e.Size()
			n += 1 + l + sovOwner(uint64(l))
		}
	}
	return n
}

func sovOwner(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOwner(x uint64) (n int) {
	return sovOwner(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OwnerCollection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOwner
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
			return fmt.Errorf("proto: OwnerCollection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OwnerCollection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOwner
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
				return ErrInvalidLengthOwner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOwner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOwner
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
				return ErrInvalidLengthOwner
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOwner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Token == nil {
				m.Token = &NtNft{}
			}
			if err := m.Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOwner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOwner
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
func (m *Owner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOwner
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
			return fmt.Errorf("proto: Owner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Owner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOwner
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
				return ErrInvalidLengthOwner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOwner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOwner
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
				return ErrInvalidLengthOwner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOwner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collection", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOwner
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
				return ErrInvalidLengthOwner
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOwner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collection = append(m.Collection, &OwnerCollection{})
			if err := m.Collection[len(m.Collection)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOwner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOwner
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
func skipOwner(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOwner
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
					return 0, ErrIntOverflowOwner
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
					return 0, ErrIntOverflowOwner
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
				return 0, ErrInvalidLengthOwner
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOwner
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOwner
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOwner        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOwner          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOwner = fmt.Errorf("proto: unexpected end of group")
)
