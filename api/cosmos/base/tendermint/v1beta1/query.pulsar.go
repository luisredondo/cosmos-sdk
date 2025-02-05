package tendermintv1beta1

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/query/v1beta1"
	p2p "github.com/cosmos/cosmos-sdk/api/tendermint/p2p"
	types "github.com/cosmos/cosmos-sdk/api/tendermint/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_GetValidatorSetByHeightRequest            protoreflect.MessageDescriptor
	fd_GetValidatorSetByHeightRequest_height     protoreflect.FieldDescriptor
	fd_GetValidatorSetByHeightRequest_pagination protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_tendermint_v1beta1_query_proto_init()
	md_GetValidatorSetByHeightRequest = File_cosmos_base_tendermint_v1beta1_query_proto.Messages().ByName("GetValidatorSetByHeightRequest")
	fd_GetValidatorSetByHeightRequest_height = md_GetValidatorSetByHeightRequest.Fields().ByName("height")
	fd_GetValidatorSetByHeightRequest_pagination = md_GetValidatorSetByHeightRequest.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_GetValidatorSetByHeightRequest)(nil)

type fastReflection_GetValidatorSetByHeightRequest GetValidatorSetByHeightRequest

func (x *GetValidatorSetByHeightRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetValidatorSetByHeightRequest)(x)
}

func (x *GetValidatorSetByHeightRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetValidatorSetByHeightRequest_messageType fastReflection_GetValidatorSetByHeightRequest_messageType
var _ protoreflect.MessageType = fastReflection_GetValidatorSetByHeightRequest_messageType{}

type fastReflection_GetValidatorSetByHeightRequest_messageType struct{}

func (x fastReflection_GetValidatorSetByHeightRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetValidatorSetByHeightRequest)(nil)
}
func (x fastReflection_GetValidatorSetByHeightRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_GetValidatorSetByHeightRequest)
}
func (x fastReflection_GetValidatorSetByHeightRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetValidatorSetByHeightRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetValidatorSetByHeightRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_GetValidatorSetByHeightRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetValidatorSetByHeightRequest) Type() protoreflect.MessageType {
	return _fastReflection_GetValidatorSetByHeightRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetValidatorSetByHeightRequest) New() protoreflect.Message {
	return new(fastReflection_GetValidatorSetByHeightRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetValidatorSetByHeightRequest) Interface() protoreflect.ProtoMessage {
	return (*GetValidatorSetByHeightRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetValidatorSetByHeightRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Height != int64(0) {
		value := protoreflect.ValueOfInt64(x.Height)
		if !f(fd_GetValidatorSetByHeightRequest_height, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_GetValidatorSetByHeightRequest_pagination, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GetValidatorSetByHeightRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest.height":
		return x.Height != int64(0)
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetValidatorSetByHeightRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest.height":
		x.Height = int64(0)
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetValidatorSetByHeightRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest.height":
		value := x.Height
		return protoreflect.ValueOfInt64(value)
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetValidatorSetByHeightRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest.height":
		x.Height = value.Int()
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageRequest)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetValidatorSetByHeightRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageRequest)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest.height":
		panic(fmt.Errorf("field height of message cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetValidatorSetByHeightRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest.height":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest.pagination":
		m := new(v1beta1.PageRequest)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetValidatorSetByHeightRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetValidatorSetByHeightRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetValidatorSetByHeightRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GetValidatorSetByHeightRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetValidatorSetByHeightRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetValidatorSetByHeightRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Height != 0 {
			n += 1 + runtime.Sov(uint64(x.Height))
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetValidatorSetByHeightRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.Height != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Height))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GetValidatorSetByHeightRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetValidatorSetByHeightRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetValidatorSetByHeightRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
				}
				x.Height = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Height |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageRequest{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_GetValidatorSetByHeightResponse_2_list)(nil)

type _GetValidatorSetByHeightResponse_2_list struct {
	list *[]*Validator
}

func (x *_GetValidatorSetByHeightResponse_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GetValidatorSetByHeightResponse_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GetValidatorSetByHeightResponse_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Validator)
	(*x.list)[i] = concreteValue
}

func (x *_GetValidatorSetByHeightResponse_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Validator)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GetValidatorSetByHeightResponse_2_list) AppendMutable() protoreflect.Value {
	v := new(Validator)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GetValidatorSetByHeightResponse_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GetValidatorSetByHeightResponse_2_list) NewElement() protoreflect.Value {
	v := new(Validator)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GetValidatorSetByHeightResponse_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_GetValidatorSetByHeightResponse              protoreflect.MessageDescriptor
	fd_GetValidatorSetByHeightResponse_block_height protoreflect.FieldDescriptor
	fd_GetValidatorSetByHeightResponse_validators   protoreflect.FieldDescriptor
	fd_GetValidatorSetByHeightResponse_pagination   protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_tendermint_v1beta1_query_proto_init()
	md_GetValidatorSetByHeightResponse = File_cosmos_base_tendermint_v1beta1_query_proto.Messages().ByName("GetValidatorSetByHeightResponse")
	fd_GetValidatorSetByHeightResponse_block_height = md_GetValidatorSetByHeightResponse.Fields().ByName("block_height")
	fd_GetValidatorSetByHeightResponse_validators = md_GetValidatorSetByHeightResponse.Fields().ByName("validators")
	fd_GetValidatorSetByHeightResponse_pagination = md_GetValidatorSetByHeightResponse.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_GetValidatorSetByHeightResponse)(nil)

type fastReflection_GetValidatorSetByHeightResponse GetValidatorSetByHeightResponse

func (x *GetValidatorSetByHeightResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetValidatorSetByHeightResponse)(x)
}

func (x *GetValidatorSetByHeightResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetValidatorSetByHeightResponse_messageType fastReflection_GetValidatorSetByHeightResponse_messageType
var _ protoreflect.MessageType = fastReflection_GetValidatorSetByHeightResponse_messageType{}

type fastReflection_GetValidatorSetByHeightResponse_messageType struct{}

func (x fastReflection_GetValidatorSetByHeightResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetValidatorSetByHeightResponse)(nil)
}
func (x fastReflection_GetValidatorSetByHeightResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_GetValidatorSetByHeightResponse)
}
func (x fastReflection_GetValidatorSetByHeightResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetValidatorSetByHeightResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetValidatorSetByHeightResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_GetValidatorSetByHeightResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetValidatorSetByHeightResponse) Type() protoreflect.MessageType {
	return _fastReflection_GetValidatorSetByHeightResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetValidatorSetByHeightResponse) New() protoreflect.Message {
	return new(fastReflection_GetValidatorSetByHeightResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetValidatorSetByHeightResponse) Interface() protoreflect.ProtoMessage {
	return (*GetValidatorSetByHeightResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetValidatorSetByHeightResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BlockHeight != int64(0) {
		value := protoreflect.ValueOfInt64(x.BlockHeight)
		if !f(fd_GetValidatorSetByHeightResponse_block_height, value) {
			return
		}
	}
	if len(x.Validators) != 0 {
		value := protoreflect.ValueOfList(&_GetValidatorSetByHeightResponse_2_list{list: &x.Validators})
		if !f(fd_GetValidatorSetByHeightResponse_validators, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_GetValidatorSetByHeightResponse_pagination, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GetValidatorSetByHeightResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.block_height":
		return x.BlockHeight != int64(0)
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.validators":
		return len(x.Validators) != 0
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetValidatorSetByHeightResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.block_height":
		x.BlockHeight = int64(0)
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.validators":
		x.Validators = nil
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetValidatorSetByHeightResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.block_height":
		value := x.BlockHeight
		return protoreflect.ValueOfInt64(value)
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.validators":
		if len(x.Validators) == 0 {
			return protoreflect.ValueOfList(&_GetValidatorSetByHeightResponse_2_list{})
		}
		listValue := &_GetValidatorSetByHeightResponse_2_list{list: &x.Validators}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetValidatorSetByHeightResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.block_height":
		x.BlockHeight = value.Int()
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.validators":
		lv := value.List()
		clv := lv.(*_GetValidatorSetByHeightResponse_2_list)
		x.Validators = *clv.list
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetValidatorSetByHeightResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.validators":
		if x.Validators == nil {
			x.Validators = []*Validator{}
		}
		value := &_GetValidatorSetByHeightResponse_2_list{list: &x.Validators}
		return protoreflect.ValueOfList(value)
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageResponse)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.block_height":
		panic(fmt.Errorf("field block_height of message cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetValidatorSetByHeightResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.block_height":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.validators":
		list := []*Validator{}
		return protoreflect.ValueOfList(&_GetValidatorSetByHeightResponse_2_list{list: &list})
	case "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.pagination":
		m := new(v1beta1.PageResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetValidatorSetByHeightResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetValidatorSetByHeightResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetValidatorSetByHeightResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GetValidatorSetByHeightResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetValidatorSetByHeightResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetValidatorSetByHeightResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BlockHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.BlockHeight))
		}
		if len(x.Validators) > 0 {
			for _, e := range x.Validators {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetValidatorSetByHeightResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Validators) > 0 {
			for iNdEx := len(x.Validators) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Validators[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
			}
		}
		if x.BlockHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BlockHeight))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GetValidatorSetByHeightResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetValidatorSetByHeightResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetValidatorSetByHeightResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
				}
				x.BlockHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BlockHeight |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Validators = append(x.Validators, &Validator{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Validators[len(x.Validators)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_GetLatestValidatorSetRequest            protoreflect.MessageDescriptor
	fd_GetLatestValidatorSetRequest_pagination protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_tendermint_v1beta1_query_proto_init()
	md_GetLatestValidatorSetRequest = File_cosmos_base_tendermint_v1beta1_query_proto.Messages().ByName("GetLatestValidatorSetRequest")
	fd_GetLatestValidatorSetRequest_pagination = md_GetLatestValidatorSetRequest.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_GetLatestValidatorSetRequest)(nil)

type fastReflection_GetLatestValidatorSetRequest GetLatestValidatorSetRequest

func (x *GetLatestValidatorSetRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetLatestValidatorSetRequest)(x)
}

func (x *GetLatestValidatorSetRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetLatestValidatorSetRequest_messageType fastReflection_GetLatestValidatorSetRequest_messageType
var _ protoreflect.MessageType = fastReflection_GetLatestValidatorSetRequest_messageType{}

type fastReflection_GetLatestValidatorSetRequest_messageType struct{}

func (x fastReflection_GetLatestValidatorSetRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetLatestValidatorSetRequest)(nil)
}
func (x fastReflection_GetLatestValidatorSetRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_GetLatestValidatorSetRequest)
}
func (x fastReflection_GetLatestValidatorSetRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetLatestValidatorSetRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetLatestValidatorSetRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_GetLatestValidatorSetRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetLatestValidatorSetRequest) Type() protoreflect.MessageType {
	return _fastReflection_GetLatestValidatorSetRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetLatestValidatorSetRequest) New() protoreflect.Message {
	return new(fastReflection_GetLatestValidatorSetRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetLatestValidatorSetRequest) Interface() protoreflect.ProtoMessage {
	return (*GetLatestValidatorSetRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetLatestValidatorSetRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_GetLatestValidatorSetRequest_pagination, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GetLatestValidatorSetRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetLatestValidatorSetRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetLatestValidatorSetRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetLatestValidatorSetRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageRequest)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetLatestValidatorSetRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageRequest)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetLatestValidatorSetRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest.pagination":
		m := new(v1beta1.PageRequest)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetLatestValidatorSetRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetLatestValidatorSetRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetLatestValidatorSetRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GetLatestValidatorSetRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetLatestValidatorSetRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetLatestValidatorSetRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetLatestValidatorSetRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GetLatestValidatorSetRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetLatestValidatorSetRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetLatestValidatorSetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageRequest{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_GetLatestValidatorSetResponse_2_list)(nil)

type _GetLatestValidatorSetResponse_2_list struct {
	list *[]*Validator
}

func (x *_GetLatestValidatorSetResponse_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GetLatestValidatorSetResponse_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GetLatestValidatorSetResponse_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Validator)
	(*x.list)[i] = concreteValue
}

func (x *_GetLatestValidatorSetResponse_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Validator)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GetLatestValidatorSetResponse_2_list) AppendMutable() protoreflect.Value {
	v := new(Validator)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GetLatestValidatorSetResponse_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GetLatestValidatorSetResponse_2_list) NewElement() protoreflect.Value {
	v := new(Validator)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GetLatestValidatorSetResponse_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_GetLatestValidatorSetResponse              protoreflect.MessageDescriptor
	fd_GetLatestValidatorSetResponse_block_height protoreflect.FieldDescriptor
	fd_GetLatestValidatorSetResponse_validators   protoreflect.FieldDescriptor
	fd_GetLatestValidatorSetResponse_pagination   protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_tendermint_v1beta1_query_proto_init()
	md_GetLatestValidatorSetResponse = File_cosmos_base_tendermint_v1beta1_query_proto.Messages().ByName("GetLatestValidatorSetResponse")
	fd_GetLatestValidatorSetResponse_block_height = md_GetLatestValidatorSetResponse.Fields().ByName("block_height")
	fd_GetLatestValidatorSetResponse_validators = md_GetLatestValidatorSetResponse.Fields().ByName("validators")
	fd_GetLatestValidatorSetResponse_pagination = md_GetLatestValidatorSetResponse.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_GetLatestValidatorSetResponse)(nil)

type fastReflection_GetLatestValidatorSetResponse GetLatestValidatorSetResponse

func (x *GetLatestValidatorSetResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetLatestValidatorSetResponse)(x)
}

func (x *GetLatestValidatorSetResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetLatestValidatorSetResponse_messageType fastReflection_GetLatestValidatorSetResponse_messageType
var _ protoreflect.MessageType = fastReflection_GetLatestValidatorSetResponse_messageType{}

type fastReflection_GetLatestValidatorSetResponse_messageType struct{}

func (x fastReflection_GetLatestValidatorSetResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetLatestValidatorSetResponse)(nil)
}
func (x fastReflection_GetLatestValidatorSetResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_GetLatestValidatorSetResponse)
}
func (x fastReflection_GetLatestValidatorSetResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetLatestValidatorSetResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetLatestValidatorSetResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_GetLatestValidatorSetResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetLatestValidatorSetResponse) Type() protoreflect.MessageType {
	return _fastReflection_GetLatestValidatorSetResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetLatestValidatorSetResponse) New() protoreflect.Message {
	return new(fastReflection_GetLatestValidatorSetResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetLatestValidatorSetResponse) Interface() protoreflect.ProtoMessage {
	return (*GetLatestValidatorSetResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetLatestValidatorSetResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BlockHeight != int64(0) {
		value := protoreflect.ValueOfInt64(x.BlockHeight)
		if !f(fd_GetLatestValidatorSetResponse_block_height, value) {
			return
		}
	}
	if len(x.Validators) != 0 {
		value := protoreflect.ValueOfList(&_GetLatestValidatorSetResponse_2_list{list: &x.Validators})
		if !f(fd_GetLatestValidatorSetResponse_validators, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_GetLatestValidatorSetResponse_pagination, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GetLatestValidatorSetResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.block_height":
		return x.BlockHeight != int64(0)
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.validators":
		return len(x.Validators) != 0
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetLatestValidatorSetResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.block_height":
		x.BlockHeight = int64(0)
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.validators":
		x.Validators = nil
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetLatestValidatorSetResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.block_height":
		value := x.BlockHeight
		return protoreflect.ValueOfInt64(value)
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.validators":
		if len(x.Validators) == 0 {
			return protoreflect.ValueOfList(&_GetLatestValidatorSetResponse_2_list{})
		}
		listValue := &_GetLatestValidatorSetResponse_2_list{list: &x.Validators}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetLatestValidatorSetResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.block_height":
		x.BlockHeight = value.Int()
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.validators":
		lv := value.List()
		clv := lv.(*_GetLatestValidatorSetResponse_2_list)
		x.Validators = *clv.list
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetLatestValidatorSetResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.validators":
		if x.Validators == nil {
			x.Validators = []*Validator{}
		}
		value := &_GetLatestValidatorSetResponse_2_list{list: &x.Validators}
		return protoreflect.ValueOfList(value)
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageResponse)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.block_height":
		panic(fmt.Errorf("field block_height of message cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetLatestValidatorSetResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.block_height":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.validators":
		list := []*Validator{}
		return protoreflect.ValueOfList(&_GetLatestValidatorSetResponse_2_list{list: &list})
	case "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.pagination":
		m := new(v1beta1.PageResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetLatestValidatorSetResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetLatestValidatorSetResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetLatestValidatorSetResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GetLatestValidatorSetResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetLatestValidatorSetResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetLatestValidatorSetResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BlockHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.BlockHeight))
		}
		if len(x.Validators) > 0 {
			for _, e := range x.Validators {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetLatestValidatorSetResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Validators) > 0 {
			for iNdEx := len(x.Validators) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Validators[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
			}
		}
		if x.BlockHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BlockHeight))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GetLatestValidatorSetResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetLatestValidatorSetResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetLatestValidatorSetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
				}
				x.BlockHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BlockHeight |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Validators = append(x.Validators, &Validator{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Validators[len(x.Validators)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_Validator                   protoreflect.MessageDescriptor
	fd_Validator_address           protoreflect.FieldDescriptor
	fd_Validator_pub_key           protoreflect.FieldDescriptor
	fd_Validator_voting_power      protoreflect.FieldDescriptor
	fd_Validator_proposer_priority protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_tendermint_v1beta1_query_proto_init()
	md_Validator = File_cosmos_base_tendermint_v1beta1_query_proto.Messages().ByName("Validator")
	fd_Validator_address = md_Validator.Fields().ByName("address")
	fd_Validator_pub_key = md_Validator.Fields().ByName("pub_key")
	fd_Validator_voting_power = md_Validator.Fields().ByName("voting_power")
	fd_Validator_proposer_priority = md_Validator.Fields().ByName("proposer_priority")
}

var _ protoreflect.Message = (*fastReflection_Validator)(nil)

type fastReflection_Validator Validator

func (x *Validator) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Validator)(x)
}

func (x *Validator) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Validator_messageType fastReflection_Validator_messageType
var _ protoreflect.MessageType = fastReflection_Validator_messageType{}

type fastReflection_Validator_messageType struct{}

func (x fastReflection_Validator_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Validator)(nil)
}
func (x fastReflection_Validator_messageType) New() protoreflect.Message {
	return new(fastReflection_Validator)
}
func (x fastReflection_Validator_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Validator
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Validator) Descriptor() protoreflect.MessageDescriptor {
	return md_Validator
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Validator) Type() protoreflect.MessageType {
	return _fastReflection_Validator_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Validator) New() protoreflect.Message {
	return new(fastReflection_Validator)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Validator) Interface() protoreflect.ProtoMessage {
	return (*Validator)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Validator) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_Validator_address, value) {
			return
		}
	}
	if x.PubKey != nil {
		value := protoreflect.ValueOfMessage(x.PubKey.ProtoReflect())
		if !f(fd_Validator_pub_key, value) {
			return
		}
	}
	if x.VotingPower != int64(0) {
		value := protoreflect.ValueOfInt64(x.VotingPower)
		if !f(fd_Validator_voting_power, value) {
			return
		}
	}
	if x.ProposerPriority != int64(0) {
		value := protoreflect.ValueOfInt64(x.ProposerPriority)
		if !f(fd_Validator_proposer_priority, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Validator) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.Validator.address":
		return x.Address != ""
	case "cosmos.base.tendermint.v1beta1.Validator.pub_key":
		return x.PubKey != nil
	case "cosmos.base.tendermint.v1beta1.Validator.voting_power":
		return x.VotingPower != int64(0)
	case "cosmos.base.tendermint.v1beta1.Validator.proposer_priority":
		return x.ProposerPriority != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.Validator"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.Validator does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Validator) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.Validator.address":
		x.Address = ""
	case "cosmos.base.tendermint.v1beta1.Validator.pub_key":
		x.PubKey = nil
	case "cosmos.base.tendermint.v1beta1.Validator.voting_power":
		x.VotingPower = int64(0)
	case "cosmos.base.tendermint.v1beta1.Validator.proposer_priority":
		x.ProposerPriority = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.Validator"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.Validator does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Validator) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.tendermint.v1beta1.Validator.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "cosmos.base.tendermint.v1beta1.Validator.pub_key":
		value := x.PubKey
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.base.tendermint.v1beta1.Validator.voting_power":
		value := x.VotingPower
		return protoreflect.ValueOfInt64(value)
	case "cosmos.base.tendermint.v1beta1.Validator.proposer_priority":
		value := x.ProposerPriority
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.Validator"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.Validator does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Validator) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.Validator.address":
		x.Address = value.Interface().(string)
	case "cosmos.base.tendermint.v1beta1.Validator.pub_key":
		x.PubKey = value.Message().Interface().(*anypb.Any)
	case "cosmos.base.tendermint.v1beta1.Validator.voting_power":
		x.VotingPower = value.Int()
	case "cosmos.base.tendermint.v1beta1.Validator.proposer_priority":
		x.ProposerPriority = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.Validator"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.Validator does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Validator) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.Validator.pub_key":
		if x.PubKey == nil {
			x.PubKey = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.PubKey.ProtoReflect())
	case "cosmos.base.tendermint.v1beta1.Validator.address":
		panic(fmt.Errorf("field address of message cosmos.base.tendermint.v1beta1.Validator is not mutable"))
	case "cosmos.base.tendermint.v1beta1.Validator.voting_power":
		panic(fmt.Errorf("field voting_power of message cosmos.base.tendermint.v1beta1.Validator is not mutable"))
	case "cosmos.base.tendermint.v1beta1.Validator.proposer_priority":
		panic(fmt.Errorf("field proposer_priority of message cosmos.base.tendermint.v1beta1.Validator is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.Validator"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.Validator does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Validator) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.Validator.address":
		return protoreflect.ValueOfString("")
	case "cosmos.base.tendermint.v1beta1.Validator.pub_key":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.base.tendermint.v1beta1.Validator.voting_power":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.base.tendermint.v1beta1.Validator.proposer_priority":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.Validator"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.Validator does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Validator) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.tendermint.v1beta1.Validator", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Validator) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Validator) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Validator) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Validator) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Validator)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.PubKey != nil {
			l = options.Size(x.PubKey)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.VotingPower != 0 {
			n += 1 + runtime.Sov(uint64(x.VotingPower))
		}
		if x.ProposerPriority != 0 {
			n += 1 + runtime.Sov(uint64(x.ProposerPriority))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Validator)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.ProposerPriority != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProposerPriority))
			i--
			dAtA[i] = 0x20
		}
		if x.VotingPower != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.VotingPower))
			i--
			dAtA[i] = 0x18
		}
		if x.PubKey != nil {
			encoded, err := options.Marshal(x.PubKey)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Validator)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Validator: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.PubKey == nil {
					x.PubKey = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PubKey); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
				}
				x.VotingPower = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.VotingPower |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProposerPriority", wireType)
				}
				x.ProposerPriority = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProposerPriority |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_GetBlockByHeightRequest        protoreflect.MessageDescriptor
	fd_GetBlockByHeightRequest_height protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_tendermint_v1beta1_query_proto_init()
	md_GetBlockByHeightRequest = File_cosmos_base_tendermint_v1beta1_query_proto.Messages().ByName("GetBlockByHeightRequest")
	fd_GetBlockByHeightRequest_height = md_GetBlockByHeightRequest.Fields().ByName("height")
}

var _ protoreflect.Message = (*fastReflection_GetBlockByHeightRequest)(nil)

type fastReflection_GetBlockByHeightRequest GetBlockByHeightRequest

func (x *GetBlockByHeightRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetBlockByHeightRequest)(x)
}

func (x *GetBlockByHeightRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetBlockByHeightRequest_messageType fastReflection_GetBlockByHeightRequest_messageType
var _ protoreflect.MessageType = fastReflection_GetBlockByHeightRequest_messageType{}

type fastReflection_GetBlockByHeightRequest_messageType struct{}

func (x fastReflection_GetBlockByHeightRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetBlockByHeightRequest)(nil)
}
func (x fastReflection_GetBlockByHeightRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_GetBlockByHeightRequest)
}
func (x fastReflection_GetBlockByHeightRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetBlockByHeightRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetBlockByHeightRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_GetBlockByHeightRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetBlockByHeightRequest) Type() protoreflect.MessageType {
	return _fastReflection_GetBlockByHeightRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetBlockByHeightRequest) New() protoreflect.Message {
	return new(fastReflection_GetBlockByHeightRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetBlockByHeightRequest) Interface() protoreflect.ProtoMessage {
	return (*GetBlockByHeightRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetBlockByHeightRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Height != int64(0) {
		value := protoreflect.ValueOfInt64(x.Height)
		if !f(fd_GetBlockByHeightRequest_height, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GetBlockByHeightRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest.height":
		return x.Height != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetBlockByHeightRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest.height":
		x.Height = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetBlockByHeightRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest.height":
		value := x.Height
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetBlockByHeightRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest.height":
		x.Height = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetBlockByHeightRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest.height":
		panic(fmt.Errorf("field height of message cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetBlockByHeightRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest.height":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetBlockByHeightRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetBlockByHeightRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetBlockByHeightRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GetBlockByHeightRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetBlockByHeightRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetBlockByHeightRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Height != 0 {
			n += 1 + runtime.Sov(uint64(x.Height))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetBlockByHeightRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Height != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Height))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GetBlockByHeightRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetBlockByHeightRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetBlockByHeightRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
				}
				x.Height = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Height |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_GetBlockByHeightResponse          protoreflect.MessageDescriptor
	fd_GetBlockByHeightResponse_block_id protoreflect.FieldDescriptor
	fd_GetBlockByHeightResponse_block    protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_tendermint_v1beta1_query_proto_init()
	md_GetBlockByHeightResponse = File_cosmos_base_tendermint_v1beta1_query_proto.Messages().ByName("GetBlockByHeightResponse")
	fd_GetBlockByHeightResponse_block_id = md_GetBlockByHeightResponse.Fields().ByName("block_id")
	fd_GetBlockByHeightResponse_block = md_GetBlockByHeightResponse.Fields().ByName("block")
}

var _ protoreflect.Message = (*fastReflection_GetBlockByHeightResponse)(nil)

type fastReflection_GetBlockByHeightResponse GetBlockByHeightResponse

func (x *GetBlockByHeightResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetBlockByHeightResponse)(x)
}

func (x *GetBlockByHeightResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetBlockByHeightResponse_messageType fastReflection_GetBlockByHeightResponse_messageType
var _ protoreflect.MessageType = fastReflection_GetBlockByHeightResponse_messageType{}

type fastReflection_GetBlockByHeightResponse_messageType struct{}

func (x fastReflection_GetBlockByHeightResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetBlockByHeightResponse)(nil)
}
func (x fastReflection_GetBlockByHeightResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_GetBlockByHeightResponse)
}
func (x fastReflection_GetBlockByHeightResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetBlockByHeightResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetBlockByHeightResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_GetBlockByHeightResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetBlockByHeightResponse) Type() protoreflect.MessageType {
	return _fastReflection_GetBlockByHeightResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetBlockByHeightResponse) New() protoreflect.Message {
	return new(fastReflection_GetBlockByHeightResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetBlockByHeightResponse) Interface() protoreflect.ProtoMessage {
	return (*GetBlockByHeightResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetBlockByHeightResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BlockId != nil {
		value := protoreflect.ValueOfMessage(x.BlockId.ProtoReflect())
		if !f(fd_GetBlockByHeightResponse_block_id, value) {
			return
		}
	}
	if x.Block != nil {
		value := protoreflect.ValueOfMessage(x.Block.ProtoReflect())
		if !f(fd_GetBlockByHeightResponse_block, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GetBlockByHeightResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse.block_id":
		return x.BlockId != nil
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse.block":
		return x.Block != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetBlockByHeightResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse.block_id":
		x.BlockId = nil
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse.block":
		x.Block = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetBlockByHeightResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse.block_id":
		value := x.BlockId
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse.block":
		value := x.Block
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetBlockByHeightResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse.block_id":
		x.BlockId = value.Message().Interface().(*types.BlockID)
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse.block":
		x.Block = value.Message().Interface().(*types.Block)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetBlockByHeightResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse.block_id":
		if x.BlockId == nil {
			x.BlockId = new(types.BlockID)
		}
		return protoreflect.ValueOfMessage(x.BlockId.ProtoReflect())
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse.block":
		if x.Block == nil {
			x.Block = new(types.Block)
		}
		return protoreflect.ValueOfMessage(x.Block.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetBlockByHeightResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse.block_id":
		m := new(types.BlockID)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse.block":
		m := new(types.Block)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetBlockByHeightResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetBlockByHeightResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetBlockByHeightResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GetBlockByHeightResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetBlockByHeightResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetBlockByHeightResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BlockId != nil {
			l = options.Size(x.BlockId)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Block != nil {
			l = options.Size(x.Block)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetBlockByHeightResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Block != nil {
			encoded, err := options.Marshal(x.Block)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.BlockId != nil {
			encoded, err := options.Marshal(x.BlockId)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GetBlockByHeightResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetBlockByHeightResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetBlockByHeightResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.BlockId == nil {
					x.BlockId = &types.BlockID{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BlockId); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Block == nil {
					x.Block = &types.Block{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Block); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_GetLatestBlockRequest protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_base_tendermint_v1beta1_query_proto_init()
	md_GetLatestBlockRequest = File_cosmos_base_tendermint_v1beta1_query_proto.Messages().ByName("GetLatestBlockRequest")
}

var _ protoreflect.Message = (*fastReflection_GetLatestBlockRequest)(nil)

type fastReflection_GetLatestBlockRequest GetLatestBlockRequest

func (x *GetLatestBlockRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetLatestBlockRequest)(x)
}

func (x *GetLatestBlockRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetLatestBlockRequest_messageType fastReflection_GetLatestBlockRequest_messageType
var _ protoreflect.MessageType = fastReflection_GetLatestBlockRequest_messageType{}

type fastReflection_GetLatestBlockRequest_messageType struct{}

func (x fastReflection_GetLatestBlockRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetLatestBlockRequest)(nil)
}
func (x fastReflection_GetLatestBlockRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_GetLatestBlockRequest)
}
func (x fastReflection_GetLatestBlockRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetLatestBlockRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetLatestBlockRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_GetLatestBlockRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetLatestBlockRequest) Type() protoreflect.MessageType {
	return _fastReflection_GetLatestBlockRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetLatestBlockRequest) New() protoreflect.Message {
	return new(fastReflection_GetLatestBlockRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetLatestBlockRequest) Interface() protoreflect.ProtoMessage {
	return (*GetLatestBlockRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetLatestBlockRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GetLatestBlockRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestBlockRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestBlockRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetLatestBlockRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestBlockRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestBlockRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetLatestBlockRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestBlockRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestBlockRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetLatestBlockRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestBlockRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestBlockRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetLatestBlockRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestBlockRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestBlockRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetLatestBlockRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestBlockRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestBlockRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetLatestBlockRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.tendermint.v1beta1.GetLatestBlockRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetLatestBlockRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetLatestBlockRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GetLatestBlockRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetLatestBlockRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetLatestBlockRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetLatestBlockRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GetLatestBlockRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetLatestBlockRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetLatestBlockRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_GetLatestBlockResponse          protoreflect.MessageDescriptor
	fd_GetLatestBlockResponse_block_id protoreflect.FieldDescriptor
	fd_GetLatestBlockResponse_block    protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_tendermint_v1beta1_query_proto_init()
	md_GetLatestBlockResponse = File_cosmos_base_tendermint_v1beta1_query_proto.Messages().ByName("GetLatestBlockResponse")
	fd_GetLatestBlockResponse_block_id = md_GetLatestBlockResponse.Fields().ByName("block_id")
	fd_GetLatestBlockResponse_block = md_GetLatestBlockResponse.Fields().ByName("block")
}

var _ protoreflect.Message = (*fastReflection_GetLatestBlockResponse)(nil)

type fastReflection_GetLatestBlockResponse GetLatestBlockResponse

func (x *GetLatestBlockResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetLatestBlockResponse)(x)
}

func (x *GetLatestBlockResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetLatestBlockResponse_messageType fastReflection_GetLatestBlockResponse_messageType
var _ protoreflect.MessageType = fastReflection_GetLatestBlockResponse_messageType{}

type fastReflection_GetLatestBlockResponse_messageType struct{}

func (x fastReflection_GetLatestBlockResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetLatestBlockResponse)(nil)
}
func (x fastReflection_GetLatestBlockResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_GetLatestBlockResponse)
}
func (x fastReflection_GetLatestBlockResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetLatestBlockResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetLatestBlockResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_GetLatestBlockResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetLatestBlockResponse) Type() protoreflect.MessageType {
	return _fastReflection_GetLatestBlockResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetLatestBlockResponse) New() protoreflect.Message {
	return new(fastReflection_GetLatestBlockResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetLatestBlockResponse) Interface() protoreflect.ProtoMessage {
	return (*GetLatestBlockResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetLatestBlockResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BlockId != nil {
		value := protoreflect.ValueOfMessage(x.BlockId.ProtoReflect())
		if !f(fd_GetLatestBlockResponse_block_id, value) {
			return
		}
	}
	if x.Block != nil {
		value := protoreflect.ValueOfMessage(x.Block.ProtoReflect())
		if !f(fd_GetLatestBlockResponse_block, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GetLatestBlockResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestBlockResponse.block_id":
		return x.BlockId != nil
	case "cosmos.base.tendermint.v1beta1.GetLatestBlockResponse.block":
		return x.Block != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestBlockResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestBlockResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetLatestBlockResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestBlockResponse.block_id":
		x.BlockId = nil
	case "cosmos.base.tendermint.v1beta1.GetLatestBlockResponse.block":
		x.Block = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestBlockResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestBlockResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetLatestBlockResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestBlockResponse.block_id":
		value := x.BlockId
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.base.tendermint.v1beta1.GetLatestBlockResponse.block":
		value := x.Block
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestBlockResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestBlockResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetLatestBlockResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestBlockResponse.block_id":
		x.BlockId = value.Message().Interface().(*types.BlockID)
	case "cosmos.base.tendermint.v1beta1.GetLatestBlockResponse.block":
		x.Block = value.Message().Interface().(*types.Block)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestBlockResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestBlockResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetLatestBlockResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestBlockResponse.block_id":
		if x.BlockId == nil {
			x.BlockId = new(types.BlockID)
		}
		return protoreflect.ValueOfMessage(x.BlockId.ProtoReflect())
	case "cosmos.base.tendermint.v1beta1.GetLatestBlockResponse.block":
		if x.Block == nil {
			x.Block = new(types.Block)
		}
		return protoreflect.ValueOfMessage(x.Block.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestBlockResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestBlockResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetLatestBlockResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetLatestBlockResponse.block_id":
		m := new(types.BlockID)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.base.tendermint.v1beta1.GetLatestBlockResponse.block":
		m := new(types.Block)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetLatestBlockResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetLatestBlockResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetLatestBlockResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.tendermint.v1beta1.GetLatestBlockResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetLatestBlockResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetLatestBlockResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GetLatestBlockResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetLatestBlockResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetLatestBlockResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BlockId != nil {
			l = options.Size(x.BlockId)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Block != nil {
			l = options.Size(x.Block)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetLatestBlockResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Block != nil {
			encoded, err := options.Marshal(x.Block)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.BlockId != nil {
			encoded, err := options.Marshal(x.BlockId)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GetLatestBlockResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetLatestBlockResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetLatestBlockResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.BlockId == nil {
					x.BlockId = &types.BlockID{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BlockId); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Block == nil {
					x.Block = &types.Block{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Block); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_GetSyncingRequest protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_base_tendermint_v1beta1_query_proto_init()
	md_GetSyncingRequest = File_cosmos_base_tendermint_v1beta1_query_proto.Messages().ByName("GetSyncingRequest")
}

var _ protoreflect.Message = (*fastReflection_GetSyncingRequest)(nil)

type fastReflection_GetSyncingRequest GetSyncingRequest

func (x *GetSyncingRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetSyncingRequest)(x)
}

func (x *GetSyncingRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetSyncingRequest_messageType fastReflection_GetSyncingRequest_messageType
var _ protoreflect.MessageType = fastReflection_GetSyncingRequest_messageType{}

type fastReflection_GetSyncingRequest_messageType struct{}

func (x fastReflection_GetSyncingRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetSyncingRequest)(nil)
}
func (x fastReflection_GetSyncingRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_GetSyncingRequest)
}
func (x fastReflection_GetSyncingRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetSyncingRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetSyncingRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_GetSyncingRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetSyncingRequest) Type() protoreflect.MessageType {
	return _fastReflection_GetSyncingRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetSyncingRequest) New() protoreflect.Message {
	return new(fastReflection_GetSyncingRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetSyncingRequest) Interface() protoreflect.ProtoMessage {
	return (*GetSyncingRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetSyncingRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GetSyncingRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetSyncingRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetSyncingRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetSyncingRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetSyncingRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetSyncingRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetSyncingRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetSyncingRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetSyncingRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetSyncingRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetSyncingRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetSyncingRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetSyncingRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetSyncingRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetSyncingRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetSyncingRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetSyncingRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetSyncingRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetSyncingRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.tendermint.v1beta1.GetSyncingRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetSyncingRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetSyncingRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GetSyncingRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetSyncingRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetSyncingRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetSyncingRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GetSyncingRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetSyncingRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetSyncingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_GetSyncingResponse         protoreflect.MessageDescriptor
	fd_GetSyncingResponse_syncing protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_tendermint_v1beta1_query_proto_init()
	md_GetSyncingResponse = File_cosmos_base_tendermint_v1beta1_query_proto.Messages().ByName("GetSyncingResponse")
	fd_GetSyncingResponse_syncing = md_GetSyncingResponse.Fields().ByName("syncing")
}

var _ protoreflect.Message = (*fastReflection_GetSyncingResponse)(nil)

type fastReflection_GetSyncingResponse GetSyncingResponse

func (x *GetSyncingResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetSyncingResponse)(x)
}

func (x *GetSyncingResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetSyncingResponse_messageType fastReflection_GetSyncingResponse_messageType
var _ protoreflect.MessageType = fastReflection_GetSyncingResponse_messageType{}

type fastReflection_GetSyncingResponse_messageType struct{}

func (x fastReflection_GetSyncingResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetSyncingResponse)(nil)
}
func (x fastReflection_GetSyncingResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_GetSyncingResponse)
}
func (x fastReflection_GetSyncingResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetSyncingResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetSyncingResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_GetSyncingResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetSyncingResponse) Type() protoreflect.MessageType {
	return _fastReflection_GetSyncingResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetSyncingResponse) New() protoreflect.Message {
	return new(fastReflection_GetSyncingResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetSyncingResponse) Interface() protoreflect.ProtoMessage {
	return (*GetSyncingResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetSyncingResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Syncing != false {
		value := protoreflect.ValueOfBool(x.Syncing)
		if !f(fd_GetSyncingResponse_syncing, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GetSyncingResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetSyncingResponse.syncing":
		return x.Syncing != false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetSyncingResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetSyncingResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetSyncingResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetSyncingResponse.syncing":
		x.Syncing = false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetSyncingResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetSyncingResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetSyncingResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetSyncingResponse.syncing":
		value := x.Syncing
		return protoreflect.ValueOfBool(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetSyncingResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetSyncingResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetSyncingResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetSyncingResponse.syncing":
		x.Syncing = value.Bool()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetSyncingResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetSyncingResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetSyncingResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetSyncingResponse.syncing":
		panic(fmt.Errorf("field syncing of message cosmos.base.tendermint.v1beta1.GetSyncingResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetSyncingResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetSyncingResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetSyncingResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetSyncingResponse.syncing":
		return protoreflect.ValueOfBool(false)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetSyncingResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetSyncingResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetSyncingResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.tendermint.v1beta1.GetSyncingResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetSyncingResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetSyncingResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GetSyncingResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetSyncingResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetSyncingResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Syncing {
			n += 2
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetSyncingResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Syncing {
			i--
			if x.Syncing {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GetSyncingResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetSyncingResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetSyncingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Syncing", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Syncing = bool(v != 0)
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_GetNodeInfoRequest protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_base_tendermint_v1beta1_query_proto_init()
	md_GetNodeInfoRequest = File_cosmos_base_tendermint_v1beta1_query_proto.Messages().ByName("GetNodeInfoRequest")
}

var _ protoreflect.Message = (*fastReflection_GetNodeInfoRequest)(nil)

type fastReflection_GetNodeInfoRequest GetNodeInfoRequest

func (x *GetNodeInfoRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetNodeInfoRequest)(x)
}

func (x *GetNodeInfoRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetNodeInfoRequest_messageType fastReflection_GetNodeInfoRequest_messageType
var _ protoreflect.MessageType = fastReflection_GetNodeInfoRequest_messageType{}

type fastReflection_GetNodeInfoRequest_messageType struct{}

func (x fastReflection_GetNodeInfoRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetNodeInfoRequest)(nil)
}
func (x fastReflection_GetNodeInfoRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_GetNodeInfoRequest)
}
func (x fastReflection_GetNodeInfoRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetNodeInfoRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetNodeInfoRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_GetNodeInfoRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetNodeInfoRequest) Type() protoreflect.MessageType {
	return _fastReflection_GetNodeInfoRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetNodeInfoRequest) New() protoreflect.Message {
	return new(fastReflection_GetNodeInfoRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetNodeInfoRequest) Interface() protoreflect.ProtoMessage {
	return (*GetNodeInfoRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetNodeInfoRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GetNodeInfoRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetNodeInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetNodeInfoRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetNodeInfoRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetNodeInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetNodeInfoRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetNodeInfoRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetNodeInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetNodeInfoRequest does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetNodeInfoRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetNodeInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetNodeInfoRequest does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetNodeInfoRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetNodeInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetNodeInfoRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetNodeInfoRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetNodeInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetNodeInfoRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetNodeInfoRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.tendermint.v1beta1.GetNodeInfoRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetNodeInfoRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetNodeInfoRequest) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GetNodeInfoRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetNodeInfoRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetNodeInfoRequest)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetNodeInfoRequest)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GetNodeInfoRequest)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetNodeInfoRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetNodeInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_GetNodeInfoResponse                     protoreflect.MessageDescriptor
	fd_GetNodeInfoResponse_node_info           protoreflect.FieldDescriptor
	fd_GetNodeInfoResponse_application_version protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_tendermint_v1beta1_query_proto_init()
	md_GetNodeInfoResponse = File_cosmos_base_tendermint_v1beta1_query_proto.Messages().ByName("GetNodeInfoResponse")
	fd_GetNodeInfoResponse_node_info = md_GetNodeInfoResponse.Fields().ByName("node_info")
	fd_GetNodeInfoResponse_application_version = md_GetNodeInfoResponse.Fields().ByName("application_version")
}

var _ protoreflect.Message = (*fastReflection_GetNodeInfoResponse)(nil)

type fastReflection_GetNodeInfoResponse GetNodeInfoResponse

func (x *GetNodeInfoResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetNodeInfoResponse)(x)
}

func (x *GetNodeInfoResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetNodeInfoResponse_messageType fastReflection_GetNodeInfoResponse_messageType
var _ protoreflect.MessageType = fastReflection_GetNodeInfoResponse_messageType{}

type fastReflection_GetNodeInfoResponse_messageType struct{}

func (x fastReflection_GetNodeInfoResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetNodeInfoResponse)(nil)
}
func (x fastReflection_GetNodeInfoResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_GetNodeInfoResponse)
}
func (x fastReflection_GetNodeInfoResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetNodeInfoResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetNodeInfoResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_GetNodeInfoResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetNodeInfoResponse) Type() protoreflect.MessageType {
	return _fastReflection_GetNodeInfoResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetNodeInfoResponse) New() protoreflect.Message {
	return new(fastReflection_GetNodeInfoResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetNodeInfoResponse) Interface() protoreflect.ProtoMessage {
	return (*GetNodeInfoResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetNodeInfoResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.NodeInfo != nil {
		value := protoreflect.ValueOfMessage(x.NodeInfo.ProtoReflect())
		if !f(fd_GetNodeInfoResponse_node_info, value) {
			return
		}
	}
	if x.ApplicationVersion != nil {
		value := protoreflect.ValueOfMessage(x.ApplicationVersion.ProtoReflect())
		if !f(fd_GetNodeInfoResponse_application_version, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GetNodeInfoResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetNodeInfoResponse.node_info":
		return x.NodeInfo != nil
	case "cosmos.base.tendermint.v1beta1.GetNodeInfoResponse.application_version":
		return x.ApplicationVersion != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetNodeInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetNodeInfoResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetNodeInfoResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetNodeInfoResponse.node_info":
		x.NodeInfo = nil
	case "cosmos.base.tendermint.v1beta1.GetNodeInfoResponse.application_version":
		x.ApplicationVersion = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetNodeInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetNodeInfoResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetNodeInfoResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetNodeInfoResponse.node_info":
		value := x.NodeInfo
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.base.tendermint.v1beta1.GetNodeInfoResponse.application_version":
		value := x.ApplicationVersion
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetNodeInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetNodeInfoResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetNodeInfoResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetNodeInfoResponse.node_info":
		x.NodeInfo = value.Message().Interface().(*p2p.NodeInfo)
	case "cosmos.base.tendermint.v1beta1.GetNodeInfoResponse.application_version":
		x.ApplicationVersion = value.Message().Interface().(*VersionInfo)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetNodeInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetNodeInfoResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetNodeInfoResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetNodeInfoResponse.node_info":
		if x.NodeInfo == nil {
			x.NodeInfo = new(p2p.NodeInfo)
		}
		return protoreflect.ValueOfMessage(x.NodeInfo.ProtoReflect())
	case "cosmos.base.tendermint.v1beta1.GetNodeInfoResponse.application_version":
		if x.ApplicationVersion == nil {
			x.ApplicationVersion = new(VersionInfo)
		}
		return protoreflect.ValueOfMessage(x.ApplicationVersion.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetNodeInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetNodeInfoResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetNodeInfoResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.GetNodeInfoResponse.node_info":
		m := new(p2p.NodeInfo)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.base.tendermint.v1beta1.GetNodeInfoResponse.application_version":
		m := new(VersionInfo)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.GetNodeInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.GetNodeInfoResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetNodeInfoResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.tendermint.v1beta1.GetNodeInfoResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetNodeInfoResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetNodeInfoResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GetNodeInfoResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetNodeInfoResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetNodeInfoResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.NodeInfo != nil {
			l = options.Size(x.NodeInfo)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.ApplicationVersion != nil {
			l = options.Size(x.ApplicationVersion)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetNodeInfoResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.ApplicationVersion != nil {
			encoded, err := options.Marshal(x.ApplicationVersion)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.NodeInfo != nil {
			encoded, err := options.Marshal(x.NodeInfo)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GetNodeInfoResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetNodeInfoResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetNodeInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NodeInfo", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.NodeInfo == nil {
					x.NodeInfo = &p2p.NodeInfo{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.NodeInfo); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ApplicationVersion", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.ApplicationVersion == nil {
					x.ApplicationVersion = &VersionInfo{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ApplicationVersion); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_VersionInfo_7_list)(nil)

type _VersionInfo_7_list struct {
	list *[]*Module
}

func (x *_VersionInfo_7_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_VersionInfo_7_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_VersionInfo_7_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Module)
	(*x.list)[i] = concreteValue
}

func (x *_VersionInfo_7_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Module)
	*x.list = append(*x.list, concreteValue)
}

func (x *_VersionInfo_7_list) AppendMutable() protoreflect.Value {
	v := new(Module)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_VersionInfo_7_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_VersionInfo_7_list) NewElement() protoreflect.Value {
	v := new(Module)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_VersionInfo_7_list) IsValid() bool {
	return x.list != nil
}

var (
	md_VersionInfo                    protoreflect.MessageDescriptor
	fd_VersionInfo_name               protoreflect.FieldDescriptor
	fd_VersionInfo_app_name           protoreflect.FieldDescriptor
	fd_VersionInfo_version            protoreflect.FieldDescriptor
	fd_VersionInfo_git_commit         protoreflect.FieldDescriptor
	fd_VersionInfo_build_tags         protoreflect.FieldDescriptor
	fd_VersionInfo_go_version         protoreflect.FieldDescriptor
	fd_VersionInfo_build_deps         protoreflect.FieldDescriptor
	fd_VersionInfo_cosmos_sdk_version protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_tendermint_v1beta1_query_proto_init()
	md_VersionInfo = File_cosmos_base_tendermint_v1beta1_query_proto.Messages().ByName("VersionInfo")
	fd_VersionInfo_name = md_VersionInfo.Fields().ByName("name")
	fd_VersionInfo_app_name = md_VersionInfo.Fields().ByName("app_name")
	fd_VersionInfo_version = md_VersionInfo.Fields().ByName("version")
	fd_VersionInfo_git_commit = md_VersionInfo.Fields().ByName("git_commit")
	fd_VersionInfo_build_tags = md_VersionInfo.Fields().ByName("build_tags")
	fd_VersionInfo_go_version = md_VersionInfo.Fields().ByName("go_version")
	fd_VersionInfo_build_deps = md_VersionInfo.Fields().ByName("build_deps")
	fd_VersionInfo_cosmos_sdk_version = md_VersionInfo.Fields().ByName("cosmos_sdk_version")
}

var _ protoreflect.Message = (*fastReflection_VersionInfo)(nil)

type fastReflection_VersionInfo VersionInfo

func (x *VersionInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_VersionInfo)(x)
}

func (x *VersionInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_VersionInfo_messageType fastReflection_VersionInfo_messageType
var _ protoreflect.MessageType = fastReflection_VersionInfo_messageType{}

type fastReflection_VersionInfo_messageType struct{}

func (x fastReflection_VersionInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_VersionInfo)(nil)
}
func (x fastReflection_VersionInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_VersionInfo)
}
func (x fastReflection_VersionInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_VersionInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_VersionInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_VersionInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_VersionInfo) Type() protoreflect.MessageType {
	return _fastReflection_VersionInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_VersionInfo) New() protoreflect.Message {
	return new(fastReflection_VersionInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_VersionInfo) Interface() protoreflect.ProtoMessage {
	return (*VersionInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_VersionInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_VersionInfo_name, value) {
			return
		}
	}
	if x.AppName != "" {
		value := protoreflect.ValueOfString(x.AppName)
		if !f(fd_VersionInfo_app_name, value) {
			return
		}
	}
	if x.Version != "" {
		value := protoreflect.ValueOfString(x.Version)
		if !f(fd_VersionInfo_version, value) {
			return
		}
	}
	if x.GitCommit != "" {
		value := protoreflect.ValueOfString(x.GitCommit)
		if !f(fd_VersionInfo_git_commit, value) {
			return
		}
	}
	if x.BuildTags != "" {
		value := protoreflect.ValueOfString(x.BuildTags)
		if !f(fd_VersionInfo_build_tags, value) {
			return
		}
	}
	if x.GoVersion != "" {
		value := protoreflect.ValueOfString(x.GoVersion)
		if !f(fd_VersionInfo_go_version, value) {
			return
		}
	}
	if len(x.BuildDeps) != 0 {
		value := protoreflect.ValueOfList(&_VersionInfo_7_list{list: &x.BuildDeps})
		if !f(fd_VersionInfo_build_deps, value) {
			return
		}
	}
	if x.CosmosSdkVersion != "" {
		value := protoreflect.ValueOfString(x.CosmosSdkVersion)
		if !f(fd_VersionInfo_cosmos_sdk_version, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_VersionInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.VersionInfo.name":
		return x.Name != ""
	case "cosmos.base.tendermint.v1beta1.VersionInfo.app_name":
		return x.AppName != ""
	case "cosmos.base.tendermint.v1beta1.VersionInfo.version":
		return x.Version != ""
	case "cosmos.base.tendermint.v1beta1.VersionInfo.git_commit":
		return x.GitCommit != ""
	case "cosmos.base.tendermint.v1beta1.VersionInfo.build_tags":
		return x.BuildTags != ""
	case "cosmos.base.tendermint.v1beta1.VersionInfo.go_version":
		return x.GoVersion != ""
	case "cosmos.base.tendermint.v1beta1.VersionInfo.build_deps":
		return len(x.BuildDeps) != 0
	case "cosmos.base.tendermint.v1beta1.VersionInfo.cosmos_sdk_version":
		return x.CosmosSdkVersion != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.VersionInfo"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.VersionInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VersionInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.VersionInfo.name":
		x.Name = ""
	case "cosmos.base.tendermint.v1beta1.VersionInfo.app_name":
		x.AppName = ""
	case "cosmos.base.tendermint.v1beta1.VersionInfo.version":
		x.Version = ""
	case "cosmos.base.tendermint.v1beta1.VersionInfo.git_commit":
		x.GitCommit = ""
	case "cosmos.base.tendermint.v1beta1.VersionInfo.build_tags":
		x.BuildTags = ""
	case "cosmos.base.tendermint.v1beta1.VersionInfo.go_version":
		x.GoVersion = ""
	case "cosmos.base.tendermint.v1beta1.VersionInfo.build_deps":
		x.BuildDeps = nil
	case "cosmos.base.tendermint.v1beta1.VersionInfo.cosmos_sdk_version":
		x.CosmosSdkVersion = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.VersionInfo"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.VersionInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_VersionInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.tendermint.v1beta1.VersionInfo.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "cosmos.base.tendermint.v1beta1.VersionInfo.app_name":
		value := x.AppName
		return protoreflect.ValueOfString(value)
	case "cosmos.base.tendermint.v1beta1.VersionInfo.version":
		value := x.Version
		return protoreflect.ValueOfString(value)
	case "cosmos.base.tendermint.v1beta1.VersionInfo.git_commit":
		value := x.GitCommit
		return protoreflect.ValueOfString(value)
	case "cosmos.base.tendermint.v1beta1.VersionInfo.build_tags":
		value := x.BuildTags
		return protoreflect.ValueOfString(value)
	case "cosmos.base.tendermint.v1beta1.VersionInfo.go_version":
		value := x.GoVersion
		return protoreflect.ValueOfString(value)
	case "cosmos.base.tendermint.v1beta1.VersionInfo.build_deps":
		if len(x.BuildDeps) == 0 {
			return protoreflect.ValueOfList(&_VersionInfo_7_list{})
		}
		listValue := &_VersionInfo_7_list{list: &x.BuildDeps}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.base.tendermint.v1beta1.VersionInfo.cosmos_sdk_version":
		value := x.CosmosSdkVersion
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.VersionInfo"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.VersionInfo does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VersionInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.VersionInfo.name":
		x.Name = value.Interface().(string)
	case "cosmos.base.tendermint.v1beta1.VersionInfo.app_name":
		x.AppName = value.Interface().(string)
	case "cosmos.base.tendermint.v1beta1.VersionInfo.version":
		x.Version = value.Interface().(string)
	case "cosmos.base.tendermint.v1beta1.VersionInfo.git_commit":
		x.GitCommit = value.Interface().(string)
	case "cosmos.base.tendermint.v1beta1.VersionInfo.build_tags":
		x.BuildTags = value.Interface().(string)
	case "cosmos.base.tendermint.v1beta1.VersionInfo.go_version":
		x.GoVersion = value.Interface().(string)
	case "cosmos.base.tendermint.v1beta1.VersionInfo.build_deps":
		lv := value.List()
		clv := lv.(*_VersionInfo_7_list)
		x.BuildDeps = *clv.list
	case "cosmos.base.tendermint.v1beta1.VersionInfo.cosmos_sdk_version":
		x.CosmosSdkVersion = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.VersionInfo"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.VersionInfo does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VersionInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.VersionInfo.build_deps":
		if x.BuildDeps == nil {
			x.BuildDeps = []*Module{}
		}
		value := &_VersionInfo_7_list{list: &x.BuildDeps}
		return protoreflect.ValueOfList(value)
	case "cosmos.base.tendermint.v1beta1.VersionInfo.name":
		panic(fmt.Errorf("field name of message cosmos.base.tendermint.v1beta1.VersionInfo is not mutable"))
	case "cosmos.base.tendermint.v1beta1.VersionInfo.app_name":
		panic(fmt.Errorf("field app_name of message cosmos.base.tendermint.v1beta1.VersionInfo is not mutable"))
	case "cosmos.base.tendermint.v1beta1.VersionInfo.version":
		panic(fmt.Errorf("field version of message cosmos.base.tendermint.v1beta1.VersionInfo is not mutable"))
	case "cosmos.base.tendermint.v1beta1.VersionInfo.git_commit":
		panic(fmt.Errorf("field git_commit of message cosmos.base.tendermint.v1beta1.VersionInfo is not mutable"))
	case "cosmos.base.tendermint.v1beta1.VersionInfo.build_tags":
		panic(fmt.Errorf("field build_tags of message cosmos.base.tendermint.v1beta1.VersionInfo is not mutable"))
	case "cosmos.base.tendermint.v1beta1.VersionInfo.go_version":
		panic(fmt.Errorf("field go_version of message cosmos.base.tendermint.v1beta1.VersionInfo is not mutable"))
	case "cosmos.base.tendermint.v1beta1.VersionInfo.cosmos_sdk_version":
		panic(fmt.Errorf("field cosmos_sdk_version of message cosmos.base.tendermint.v1beta1.VersionInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.VersionInfo"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.VersionInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_VersionInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.VersionInfo.name":
		return protoreflect.ValueOfString("")
	case "cosmos.base.tendermint.v1beta1.VersionInfo.app_name":
		return protoreflect.ValueOfString("")
	case "cosmos.base.tendermint.v1beta1.VersionInfo.version":
		return protoreflect.ValueOfString("")
	case "cosmos.base.tendermint.v1beta1.VersionInfo.git_commit":
		return protoreflect.ValueOfString("")
	case "cosmos.base.tendermint.v1beta1.VersionInfo.build_tags":
		return protoreflect.ValueOfString("")
	case "cosmos.base.tendermint.v1beta1.VersionInfo.go_version":
		return protoreflect.ValueOfString("")
	case "cosmos.base.tendermint.v1beta1.VersionInfo.build_deps":
		list := []*Module{}
		return protoreflect.ValueOfList(&_VersionInfo_7_list{list: &list})
	case "cosmos.base.tendermint.v1beta1.VersionInfo.cosmos_sdk_version":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.VersionInfo"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.VersionInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_VersionInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.tendermint.v1beta1.VersionInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_VersionInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VersionInfo) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_VersionInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_VersionInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*VersionInfo)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Name)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.AppName)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Version)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.GitCommit)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BuildTags)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.GoVersion)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.BuildDeps) > 0 {
			for _, e := range x.BuildDeps {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.CosmosSdkVersion)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*VersionInfo)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.CosmosSdkVersion) > 0 {
			i -= len(x.CosmosSdkVersion)
			copy(dAtA[i:], x.CosmosSdkVersion)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.CosmosSdkVersion)))
			i--
			dAtA[i] = 0x42
		}
		if len(x.BuildDeps) > 0 {
			for iNdEx := len(x.BuildDeps) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.BuildDeps[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x3a
			}
		}
		if len(x.GoVersion) > 0 {
			i -= len(x.GoVersion)
			copy(dAtA[i:], x.GoVersion)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.GoVersion)))
			i--
			dAtA[i] = 0x32
		}
		if len(x.BuildTags) > 0 {
			i -= len(x.BuildTags)
			copy(dAtA[i:], x.BuildTags)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BuildTags)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.GitCommit) > 0 {
			i -= len(x.GitCommit)
			copy(dAtA[i:], x.GitCommit)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.GitCommit)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Version) > 0 {
			i -= len(x.Version)
			copy(dAtA[i:], x.Version)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Version)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.AppName) > 0 {
			i -= len(x.AppName)
			copy(dAtA[i:], x.AppName)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AppName)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*VersionInfo)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: VersionInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: VersionInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Name = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AppName", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.AppName = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Version = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GitCommit", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.GitCommit = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BuildTags", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BuildTags = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GoVersion", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.GoVersion = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BuildDeps", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BuildDeps = append(x.BuildDeps, &Module{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BuildDeps[len(x.BuildDeps)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CosmosSdkVersion", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.CosmosSdkVersion = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_Module         protoreflect.MessageDescriptor
	fd_Module_path    protoreflect.FieldDescriptor
	fd_Module_version protoreflect.FieldDescriptor
	fd_Module_sum     protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_tendermint_v1beta1_query_proto_init()
	md_Module = File_cosmos_base_tendermint_v1beta1_query_proto.Messages().ByName("Module")
	fd_Module_path = md_Module.Fields().ByName("path")
	fd_Module_version = md_Module.Fields().ByName("version")
	fd_Module_sum = md_Module.Fields().ByName("sum")
}

var _ protoreflect.Message = (*fastReflection_Module)(nil)

type fastReflection_Module Module

func (x *Module) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Module)(x)
}

func (x *Module) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Module_messageType fastReflection_Module_messageType
var _ protoreflect.MessageType = fastReflection_Module_messageType{}

type fastReflection_Module_messageType struct{}

func (x fastReflection_Module_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Module)(nil)
}
func (x fastReflection_Module_messageType) New() protoreflect.Message {
	return new(fastReflection_Module)
}
func (x fastReflection_Module_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Module
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Module) Descriptor() protoreflect.MessageDescriptor {
	return md_Module
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Module) Type() protoreflect.MessageType {
	return _fastReflection_Module_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Module) New() protoreflect.Message {
	return new(fastReflection_Module)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Module) Interface() protoreflect.ProtoMessage {
	return (*Module)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Module) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Path != "" {
		value := protoreflect.ValueOfString(x.Path)
		if !f(fd_Module_path, value) {
			return
		}
	}
	if x.Version != "" {
		value := protoreflect.ValueOfString(x.Version)
		if !f(fd_Module_version, value) {
			return
		}
	}
	if x.Sum != "" {
		value := protoreflect.ValueOfString(x.Sum)
		if !f(fd_Module_sum, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Module) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.Module.path":
		return x.Path != ""
	case "cosmos.base.tendermint.v1beta1.Module.version":
		return x.Version != ""
	case "cosmos.base.tendermint.v1beta1.Module.sum":
		return x.Sum != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.Module"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.Module does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Module) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.Module.path":
		x.Path = ""
	case "cosmos.base.tendermint.v1beta1.Module.version":
		x.Version = ""
	case "cosmos.base.tendermint.v1beta1.Module.sum":
		x.Sum = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.Module"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.Module does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Module) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.tendermint.v1beta1.Module.path":
		value := x.Path
		return protoreflect.ValueOfString(value)
	case "cosmos.base.tendermint.v1beta1.Module.version":
		value := x.Version
		return protoreflect.ValueOfString(value)
	case "cosmos.base.tendermint.v1beta1.Module.sum":
		value := x.Sum
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.Module"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.Module does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Module) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.Module.path":
		x.Path = value.Interface().(string)
	case "cosmos.base.tendermint.v1beta1.Module.version":
		x.Version = value.Interface().(string)
	case "cosmos.base.tendermint.v1beta1.Module.sum":
		x.Sum = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.Module"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.Module does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Module) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.Module.path":
		panic(fmt.Errorf("field path of message cosmos.base.tendermint.v1beta1.Module is not mutable"))
	case "cosmos.base.tendermint.v1beta1.Module.version":
		panic(fmt.Errorf("field version of message cosmos.base.tendermint.v1beta1.Module is not mutable"))
	case "cosmos.base.tendermint.v1beta1.Module.sum":
		panic(fmt.Errorf("field sum of message cosmos.base.tendermint.v1beta1.Module is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.Module"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.Module does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Module) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.tendermint.v1beta1.Module.path":
		return protoreflect.ValueOfString("")
	case "cosmos.base.tendermint.v1beta1.Module.version":
		return protoreflect.ValueOfString("")
	case "cosmos.base.tendermint.v1beta1.Module.sum":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.tendermint.v1beta1.Module"))
		}
		panic(fmt.Errorf("message cosmos.base.tendermint.v1beta1.Module does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Module) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.tendermint.v1beta1.Module", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Module) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Module) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Module) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Module) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Module)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Path)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Version)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Sum)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Module)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Sum) > 0 {
			i -= len(x.Sum)
			copy(dAtA[i:], x.Sum)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Sum)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Version) > 0 {
			i -= len(x.Version)
			copy(dAtA[i:], x.Version)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Version)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Path) > 0 {
			i -= len(x.Path)
			copy(dAtA[i:], x.Path)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Path)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Module)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Module: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Module: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Path = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Version = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Sum = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// GetNodeInfo queries the current node info.
	GetNodeInfo(ctx context.Context, in *GetNodeInfoRequest, opts ...grpc.CallOption) (*GetNodeInfoResponse, error)
	// GetSyncing queries node syncing.
	GetSyncing(ctx context.Context, in *GetSyncingRequest, opts ...grpc.CallOption) (*GetSyncingResponse, error)
	// GetLatestBlock returns the latest block.
	GetLatestBlock(ctx context.Context, in *GetLatestBlockRequest, opts ...grpc.CallOption) (*GetLatestBlockResponse, error)
	// GetBlockByHeight queries block for given height.
	GetBlockByHeight(ctx context.Context, in *GetBlockByHeightRequest, opts ...grpc.CallOption) (*GetBlockByHeightResponse, error)
	// GetLatestValidatorSet queries latest validator-set.
	GetLatestValidatorSet(ctx context.Context, in *GetLatestValidatorSetRequest, opts ...grpc.CallOption) (*GetLatestValidatorSetResponse, error)
	// GetValidatorSetByHeight queries validator-set at a given height.
	GetValidatorSetByHeight(ctx context.Context, in *GetValidatorSetByHeightRequest, opts ...grpc.CallOption) (*GetValidatorSetByHeightResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) GetNodeInfo(ctx context.Context, in *GetNodeInfoRequest, opts ...grpc.CallOption) (*GetNodeInfoResponse, error) {
	out := new(GetNodeInfoResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetNodeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetSyncing(ctx context.Context, in *GetSyncingRequest, opts ...grpc.CallOption) (*GetSyncingResponse, error) {
	out := new(GetSyncingResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetSyncing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetLatestBlock(ctx context.Context, in *GetLatestBlockRequest, opts ...grpc.CallOption) (*GetLatestBlockResponse, error) {
	out := new(GetLatestBlockResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetLatestBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetBlockByHeight(ctx context.Context, in *GetBlockByHeightRequest, opts ...grpc.CallOption) (*GetBlockByHeightResponse, error) {
	out := new(GetBlockByHeightResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetBlockByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetLatestValidatorSet(ctx context.Context, in *GetLatestValidatorSetRequest, opts ...grpc.CallOption) (*GetLatestValidatorSetResponse, error) {
	out := new(GetLatestValidatorSetResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetLatestValidatorSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetValidatorSetByHeight(ctx context.Context, in *GetValidatorSetByHeightRequest, opts ...grpc.CallOption) (*GetValidatorSetByHeightResponse, error) {
	out := new(GetValidatorSetByHeightResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetValidatorSetByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// GetNodeInfo queries the current node info.
	GetNodeInfo(context.Context, *GetNodeInfoRequest) (*GetNodeInfoResponse, error)
	// GetSyncing queries node syncing.
	GetSyncing(context.Context, *GetSyncingRequest) (*GetSyncingResponse, error)
	// GetLatestBlock returns the latest block.
	GetLatestBlock(context.Context, *GetLatestBlockRequest) (*GetLatestBlockResponse, error)
	// GetBlockByHeight queries block for given height.
	GetBlockByHeight(context.Context, *GetBlockByHeightRequest) (*GetBlockByHeightResponse, error)
	// GetLatestValidatorSet queries latest validator-set.
	GetLatestValidatorSet(context.Context, *GetLatestValidatorSetRequest) (*GetLatestValidatorSetResponse, error)
	// GetValidatorSetByHeight queries validator-set at a given height.
	GetValidatorSetByHeight(context.Context, *GetValidatorSetByHeightRequest) (*GetValidatorSetByHeightResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) GetNodeInfo(context.Context, *GetNodeInfoRequest) (*GetNodeInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeInfo not implemented")
}
func (UnimplementedServiceServer) GetSyncing(context.Context, *GetSyncingRequest) (*GetSyncingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSyncing not implemented")
}
func (UnimplementedServiceServer) GetLatestBlock(context.Context, *GetLatestBlockRequest) (*GetLatestBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestBlock not implemented")
}
func (UnimplementedServiceServer) GetBlockByHeight(context.Context, *GetBlockByHeightRequest) (*GetBlockByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockByHeight not implemented")
}
func (UnimplementedServiceServer) GetLatestValidatorSet(context.Context, *GetLatestValidatorSetRequest) (*GetLatestValidatorSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestValidatorSet not implemented")
}
func (UnimplementedServiceServer) GetValidatorSetByHeight(context.Context, *GetValidatorSetByHeightRequest) (*GetValidatorSetByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidatorSetByHeight not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_GetNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetNodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetNodeInfo(ctx, req.(*GetNodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetSyncing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSyncingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetSyncing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetSyncing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetSyncing(ctx, req.(*GetSyncingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetLatestBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetLatestBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetLatestBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetLatestBlock(ctx, req.(*GetLatestBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetBlockByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetBlockByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetBlockByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetBlockByHeight(ctx, req.(*GetBlockByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetLatestValidatorSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestValidatorSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetLatestValidatorSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetLatestValidatorSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetLatestValidatorSet(ctx, req.(*GetLatestValidatorSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetValidatorSetByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValidatorSetByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetValidatorSetByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetValidatorSetByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetValidatorSetByHeight(ctx, req.(*GetValidatorSetByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.base.tendermint.v1beta1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeInfo",
			Handler:    _Service_GetNodeInfo_Handler,
		},
		{
			MethodName: "GetSyncing",
			Handler:    _Service_GetSyncing_Handler,
		},
		{
			MethodName: "GetLatestBlock",
			Handler:    _Service_GetLatestBlock_Handler,
		},
		{
			MethodName: "GetBlockByHeight",
			Handler:    _Service_GetBlockByHeight_Handler,
		},
		{
			MethodName: "GetLatestValidatorSet",
			Handler:    _Service_GetLatestValidatorSet_Handler,
		},
		{
			MethodName: "GetValidatorSetByHeight",
			Handler:    _Service_GetValidatorSetByHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/base/tendermint/v1beta1/query.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/base/tendermint/v1beta1/query.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetValidatorSetByHeightRequest is the request type for the Query/GetValidatorSetByHeight RPC method.
type GetValidatorSetByHeightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// pagination defines an pagination for the request.
	Pagination *v1beta1.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetValidatorSetByHeightRequest) Reset() {
	*x = GetValidatorSetByHeightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValidatorSetByHeightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValidatorSetByHeightRequest) ProtoMessage() {}

// Deprecated: Use GetValidatorSetByHeightRequest.ProtoReflect.Descriptor instead.
func (*GetValidatorSetByHeightRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_base_tendermint_v1beta1_query_proto_rawDescGZIP(), []int{0}
}

func (x *GetValidatorSetByHeightRequest) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *GetValidatorSetByHeightRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// GetValidatorSetByHeightResponse is the response type for the Query/GetValidatorSetByHeight RPC method.
type GetValidatorSetByHeightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHeight int64        `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Validators  []*Validator `protobuf:"bytes,2,rep,name=validators,proto3" json:"validators,omitempty"`
	// pagination defines an pagination for the response.
	Pagination *v1beta1.PageResponse `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetValidatorSetByHeightResponse) Reset() {
	*x = GetValidatorSetByHeightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValidatorSetByHeightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValidatorSetByHeightResponse) ProtoMessage() {}

// Deprecated: Use GetValidatorSetByHeightResponse.ProtoReflect.Descriptor instead.
func (*GetValidatorSetByHeightResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_base_tendermint_v1beta1_query_proto_rawDescGZIP(), []int{1}
}

func (x *GetValidatorSetByHeightResponse) GetBlockHeight() int64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *GetValidatorSetByHeightResponse) GetValidators() []*Validator {
	if x != nil {
		return x.Validators
	}
	return nil
}

func (x *GetValidatorSetByHeightResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// GetLatestValidatorSetRequest is the request type for the Query/GetValidatorSetByHeight RPC method.
type GetLatestValidatorSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pagination defines an pagination for the request.
	Pagination *v1beta1.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetLatestValidatorSetRequest) Reset() {
	*x = GetLatestValidatorSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestValidatorSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestValidatorSetRequest) ProtoMessage() {}

// Deprecated: Use GetLatestValidatorSetRequest.ProtoReflect.Descriptor instead.
func (*GetLatestValidatorSetRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_base_tendermint_v1beta1_query_proto_rawDescGZIP(), []int{2}
}

func (x *GetLatestValidatorSetRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// GetLatestValidatorSetResponse is the response type for the Query/GetValidatorSetByHeight RPC method.
type GetLatestValidatorSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHeight int64        `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Validators  []*Validator `protobuf:"bytes,2,rep,name=validators,proto3" json:"validators,omitempty"`
	// pagination defines an pagination for the response.
	Pagination *v1beta1.PageResponse `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetLatestValidatorSetResponse) Reset() {
	*x = GetLatestValidatorSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestValidatorSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestValidatorSetResponse) ProtoMessage() {}

// Deprecated: Use GetLatestValidatorSetResponse.ProtoReflect.Descriptor instead.
func (*GetLatestValidatorSetResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_base_tendermint_v1beta1_query_proto_rawDescGZIP(), []int{3}
}

func (x *GetLatestValidatorSetResponse) GetBlockHeight() int64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *GetLatestValidatorSetResponse) GetValidators() []*Validator {
	if x != nil {
		return x.Validators
	}
	return nil
}

func (x *GetLatestValidatorSetResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// Validator is the type for the validator-set.
type Validator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address          string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey           *anypb.Any `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	VotingPower      int64      `protobuf:"varint,3,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
	ProposerPriority int64      `protobuf:"varint,4,opt,name=proposer_priority,json=proposerPriority,proto3" json:"proposer_priority,omitempty"`
}

func (x *Validator) Reset() {
	*x = Validator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validator) ProtoMessage() {}

// Deprecated: Use Validator.ProtoReflect.Descriptor instead.
func (*Validator) Descriptor() ([]byte, []int) {
	return file_cosmos_base_tendermint_v1beta1_query_proto_rawDescGZIP(), []int{4}
}

func (x *Validator) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Validator) GetPubKey() *anypb.Any {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *Validator) GetVotingPower() int64 {
	if x != nil {
		return x.VotingPower
	}
	return 0
}

func (x *Validator) GetProposerPriority() int64 {
	if x != nil {
		return x.ProposerPriority
	}
	return 0
}

// GetBlockByHeightRequest is the request type for the Query/GetBlockByHeight RPC method.
type GetBlockByHeightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *GetBlockByHeightRequest) Reset() {
	*x = GetBlockByHeightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockByHeightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockByHeightRequest) ProtoMessage() {}

// Deprecated: Use GetBlockByHeightRequest.ProtoReflect.Descriptor instead.
func (*GetBlockByHeightRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_base_tendermint_v1beta1_query_proto_rawDescGZIP(), []int{5}
}

func (x *GetBlockByHeightRequest) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

// GetBlockByHeightResponse is the response type for the Query/GetBlockByHeight RPC method.
type GetBlockByHeightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId *types.BlockID `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Block   *types.Block   `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *GetBlockByHeightResponse) Reset() {
	*x = GetBlockByHeightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockByHeightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockByHeightResponse) ProtoMessage() {}

// Deprecated: Use GetBlockByHeightResponse.ProtoReflect.Descriptor instead.
func (*GetBlockByHeightResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_base_tendermint_v1beta1_query_proto_rawDescGZIP(), []int{6}
}

func (x *GetBlockByHeightResponse) GetBlockId() *types.BlockID {
	if x != nil {
		return x.BlockId
	}
	return nil
}

func (x *GetBlockByHeightResponse) GetBlock() *types.Block {
	if x != nil {
		return x.Block
	}
	return nil
}

// GetLatestBlockRequest is the request type for the Query/GetLatestBlock RPC method.
type GetLatestBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLatestBlockRequest) Reset() {
	*x = GetLatestBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestBlockRequest) ProtoMessage() {}

// Deprecated: Use GetLatestBlockRequest.ProtoReflect.Descriptor instead.
func (*GetLatestBlockRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_base_tendermint_v1beta1_query_proto_rawDescGZIP(), []int{7}
}

// GetLatestBlockResponse is the response type for the Query/GetLatestBlock RPC method.
type GetLatestBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId *types.BlockID `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Block   *types.Block   `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *GetLatestBlockResponse) Reset() {
	*x = GetLatestBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestBlockResponse) ProtoMessage() {}

// Deprecated: Use GetLatestBlockResponse.ProtoReflect.Descriptor instead.
func (*GetLatestBlockResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_base_tendermint_v1beta1_query_proto_rawDescGZIP(), []int{8}
}

func (x *GetLatestBlockResponse) GetBlockId() *types.BlockID {
	if x != nil {
		return x.BlockId
	}
	return nil
}

func (x *GetLatestBlockResponse) GetBlock() *types.Block {
	if x != nil {
		return x.Block
	}
	return nil
}

// GetSyncingRequest is the request type for the Query/GetSyncing RPC method.
type GetSyncingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSyncingRequest) Reset() {
	*x = GetSyncingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSyncingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSyncingRequest) ProtoMessage() {}

// Deprecated: Use GetSyncingRequest.ProtoReflect.Descriptor instead.
func (*GetSyncingRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_base_tendermint_v1beta1_query_proto_rawDescGZIP(), []int{9}
}

// GetSyncingResponse is the response type for the Query/GetSyncing RPC method.
type GetSyncingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Syncing bool `protobuf:"varint,1,opt,name=syncing,proto3" json:"syncing,omitempty"`
}

func (x *GetSyncingResponse) Reset() {
	*x = GetSyncingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSyncingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSyncingResponse) ProtoMessage() {}

// Deprecated: Use GetSyncingResponse.ProtoReflect.Descriptor instead.
func (*GetSyncingResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_base_tendermint_v1beta1_query_proto_rawDescGZIP(), []int{10}
}

func (x *GetSyncingResponse) GetSyncing() bool {
	if x != nil {
		return x.Syncing
	}
	return false
}

// GetNodeInfoRequest is the request type for the Query/GetNodeInfo RPC method.
type GetNodeInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNodeInfoRequest) Reset() {
	*x = GetNodeInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeInfoRequest) ProtoMessage() {}

// Deprecated: Use GetNodeInfoRequest.ProtoReflect.Descriptor instead.
func (*GetNodeInfoRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_base_tendermint_v1beta1_query_proto_rawDescGZIP(), []int{11}
}

// GetNodeInfoResponse is the response type for the Query/GetNodeInfo RPC method.
type GetNodeInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeInfo           *p2p.NodeInfo `protobuf:"bytes,1,opt,name=node_info,json=nodeInfo,proto3" json:"node_info,omitempty"`
	ApplicationVersion *VersionInfo  `protobuf:"bytes,2,opt,name=application_version,json=applicationVersion,proto3" json:"application_version,omitempty"`
}

func (x *GetNodeInfoResponse) Reset() {
	*x = GetNodeInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeInfoResponse) ProtoMessage() {}

// Deprecated: Use GetNodeInfoResponse.ProtoReflect.Descriptor instead.
func (*GetNodeInfoResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_base_tendermint_v1beta1_query_proto_rawDescGZIP(), []int{12}
}

func (x *GetNodeInfoResponse) GetNodeInfo() *p2p.NodeInfo {
	if x != nil {
		return x.NodeInfo
	}
	return nil
}

func (x *GetNodeInfoResponse) GetApplicationVersion() *VersionInfo {
	if x != nil {
		return x.ApplicationVersion
	}
	return nil
}

// VersionInfo is the type for the GetNodeInfoResponse message.
type VersionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AppName   string    `protobuf:"bytes,2,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	Version   string    `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	GitCommit string    `protobuf:"bytes,4,opt,name=git_commit,json=gitCommit,proto3" json:"git_commit,omitempty"`
	BuildTags string    `protobuf:"bytes,5,opt,name=build_tags,json=buildTags,proto3" json:"build_tags,omitempty"`
	GoVersion string    `protobuf:"bytes,6,opt,name=go_version,json=goVersion,proto3" json:"go_version,omitempty"`
	BuildDeps []*Module `protobuf:"bytes,7,rep,name=build_deps,json=buildDeps,proto3" json:"build_deps,omitempty"`
	// Since: cosmos-sdk 0.43
	CosmosSdkVersion string `protobuf:"bytes,8,opt,name=cosmos_sdk_version,json=cosmosSdkVersion,proto3" json:"cosmos_sdk_version,omitempty"`
}

func (x *VersionInfo) Reset() {
	*x = VersionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionInfo) ProtoMessage() {}

// Deprecated: Use VersionInfo.ProtoReflect.Descriptor instead.
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return file_cosmos_base_tendermint_v1beta1_query_proto_rawDescGZIP(), []int{13}
}

func (x *VersionInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VersionInfo) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *VersionInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *VersionInfo) GetGitCommit() string {
	if x != nil {
		return x.GitCommit
	}
	return ""
}

func (x *VersionInfo) GetBuildTags() string {
	if x != nil {
		return x.BuildTags
	}
	return ""
}

func (x *VersionInfo) GetGoVersion() string {
	if x != nil {
		return x.GoVersion
	}
	return ""
}

func (x *VersionInfo) GetBuildDeps() []*Module {
	if x != nil {
		return x.BuildDeps
	}
	return nil
}

func (x *VersionInfo) GetCosmosSdkVersion() string {
	if x != nil {
		return x.CosmosSdkVersion
	}
	return ""
}

// Module is the type for VersionInfo
type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// module path
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// module version
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// checksum
	Sum string `protobuf:"bytes,3,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *Module) Reset() {
	*x = Module{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_cosmos_base_tendermint_v1beta1_query_proto_rawDescGZIP(), []int{14}
}

func (x *Module) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Module) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Module) GetSum() string {
	if x != nil {
		return x.Sum
	}
	return ""
}

var File_cosmos_base_tendermint_v1beta1_query_proto protoreflect.FileDescriptor

var file_cosmos_base_tendermint_v1beta1_query_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x74, 0x2f, 0x70, 0x32, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd8, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x42, 0x79, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x49,
	0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x66, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd6, 0x01, 0x0a, 0x1d, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x49, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0a,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xbe, 0x01, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x70, 0x75,
	0x62, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x69,
	0x6e, 0x67, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x22, 0x31, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x7f, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44,
	0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x05, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x7d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x64, 0x12, 0x2d, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x79, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x79,
	0x6e, 0x63, 0x69, 0x6e, 0x67, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xaa, 0x01, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5c, 0x0a, 0x13, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x12, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xa8, 0x02, 0x0a, 0x0b, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x67, 0x6f, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x6f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x45,
	0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x64, 0x65, 0x70, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x09, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x44, 0x65, 0x70, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f,
	0x73, 0x64, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x53, 0x64, 0x6b, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x48, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x32, 0x88, 0x09,
	0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa9, 0x01, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0xa4, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6e,
	0x63, 0x69, 0x6e, 0x67, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x29, 0x12, 0x27, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x12, 0xb6, 0x01, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x35, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x12, 0x2d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x2f, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0xbe, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x37, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x48,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x31, 0x12, 0x2f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x2f, 0x7b, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x7d, 0x12, 0xd2, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74,
	0x12, 0x3c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x36, 0x12, 0x34, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x65, 0x74, 0x73, 0x2f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0xda, 0x01, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x42,
	0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x3e, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x38,
	0x12, 0x36, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x65, 0x74, 0x73, 0x2f,
	0x7b, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x7d, 0x42, 0x9e, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42,
	0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xa2, 0x02, 0x03, 0x43, 0x42, 0x54, 0xaa, 0x02, 0x1e, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1e, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x5c, 0x42, 0x61, 0x73, 0x65, 0x5c, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x2a, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5c, 0x42, 0x61, 0x73, 0x65, 0x5c, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x74, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a,
	0x42, 0x61, 0x73, 0x65, 0x3a, 0x3a, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cosmos_base_tendermint_v1beta1_query_proto_rawDescOnce sync.Once
	file_cosmos_base_tendermint_v1beta1_query_proto_rawDescData = file_cosmos_base_tendermint_v1beta1_query_proto_rawDesc
)

func file_cosmos_base_tendermint_v1beta1_query_proto_rawDescGZIP() []byte {
	file_cosmos_base_tendermint_v1beta1_query_proto_rawDescOnce.Do(func() {
		file_cosmos_base_tendermint_v1beta1_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_base_tendermint_v1beta1_query_proto_rawDescData)
	})
	return file_cosmos_base_tendermint_v1beta1_query_proto_rawDescData
}

var file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_cosmos_base_tendermint_v1beta1_query_proto_goTypes = []interface{}{
	(*GetValidatorSetByHeightRequest)(nil),  // 0: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest
	(*GetValidatorSetByHeightResponse)(nil), // 1: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse
	(*GetLatestValidatorSetRequest)(nil),    // 2: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest
	(*GetLatestValidatorSetResponse)(nil),   // 3: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse
	(*Validator)(nil),                       // 4: cosmos.base.tendermint.v1beta1.Validator
	(*GetBlockByHeightRequest)(nil),         // 5: cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest
	(*GetBlockByHeightResponse)(nil),        // 6: cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse
	(*GetLatestBlockRequest)(nil),           // 7: cosmos.base.tendermint.v1beta1.GetLatestBlockRequest
	(*GetLatestBlockResponse)(nil),          // 8: cosmos.base.tendermint.v1beta1.GetLatestBlockResponse
	(*GetSyncingRequest)(nil),               // 9: cosmos.base.tendermint.v1beta1.GetSyncingRequest
	(*GetSyncingResponse)(nil),              // 10: cosmos.base.tendermint.v1beta1.GetSyncingResponse
	(*GetNodeInfoRequest)(nil),              // 11: cosmos.base.tendermint.v1beta1.GetNodeInfoRequest
	(*GetNodeInfoResponse)(nil),             // 12: cosmos.base.tendermint.v1beta1.GetNodeInfoResponse
	(*VersionInfo)(nil),                     // 13: cosmos.base.tendermint.v1beta1.VersionInfo
	(*Module)(nil),                          // 14: cosmos.base.tendermint.v1beta1.Module
	(*v1beta1.PageRequest)(nil),             // 15: cosmos.base.query.v1beta1.PageRequest
	(*v1beta1.PageResponse)(nil),            // 16: cosmos.base.query.v1beta1.PageResponse
	(*anypb.Any)(nil),                       // 17: google.protobuf.Any
	(*types.BlockID)(nil),                   // 18: tendermint.types.BlockID
	(*types.Block)(nil),                     // 19: tendermint.types.Block
	(*p2p.NodeInfo)(nil),                    // 20: tendermint.p2p.NodeInfo
}
var file_cosmos_base_tendermint_v1beta1_query_proto_depIdxs = []int32{
	15, // 0: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	4,  // 1: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.validators:type_name -> cosmos.base.tendermint.v1beta1.Validator
	16, // 2: cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	15, // 3: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	4,  // 4: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.validators:type_name -> cosmos.base.tendermint.v1beta1.Validator
	16, // 5: cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	17, // 6: cosmos.base.tendermint.v1beta1.Validator.pub_key:type_name -> google.protobuf.Any
	18, // 7: cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse.block_id:type_name -> tendermint.types.BlockID
	19, // 8: cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse.block:type_name -> tendermint.types.Block
	18, // 9: cosmos.base.tendermint.v1beta1.GetLatestBlockResponse.block_id:type_name -> tendermint.types.BlockID
	19, // 10: cosmos.base.tendermint.v1beta1.GetLatestBlockResponse.block:type_name -> tendermint.types.Block
	20, // 11: cosmos.base.tendermint.v1beta1.GetNodeInfoResponse.node_info:type_name -> tendermint.p2p.NodeInfo
	13, // 12: cosmos.base.tendermint.v1beta1.GetNodeInfoResponse.application_version:type_name -> cosmos.base.tendermint.v1beta1.VersionInfo
	14, // 13: cosmos.base.tendermint.v1beta1.VersionInfo.build_deps:type_name -> cosmos.base.tendermint.v1beta1.Module
	11, // 14: cosmos.base.tendermint.v1beta1.Service.GetNodeInfo:input_type -> cosmos.base.tendermint.v1beta1.GetNodeInfoRequest
	9,  // 15: cosmos.base.tendermint.v1beta1.Service.GetSyncing:input_type -> cosmos.base.tendermint.v1beta1.GetSyncingRequest
	7,  // 16: cosmos.base.tendermint.v1beta1.Service.GetLatestBlock:input_type -> cosmos.base.tendermint.v1beta1.GetLatestBlockRequest
	5,  // 17: cosmos.base.tendermint.v1beta1.Service.GetBlockByHeight:input_type -> cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest
	2,  // 18: cosmos.base.tendermint.v1beta1.Service.GetLatestValidatorSet:input_type -> cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest
	0,  // 19: cosmos.base.tendermint.v1beta1.Service.GetValidatorSetByHeight:input_type -> cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest
	12, // 20: cosmos.base.tendermint.v1beta1.Service.GetNodeInfo:output_type -> cosmos.base.tendermint.v1beta1.GetNodeInfoResponse
	10, // 21: cosmos.base.tendermint.v1beta1.Service.GetSyncing:output_type -> cosmos.base.tendermint.v1beta1.GetSyncingResponse
	8,  // 22: cosmos.base.tendermint.v1beta1.Service.GetLatestBlock:output_type -> cosmos.base.tendermint.v1beta1.GetLatestBlockResponse
	6,  // 23: cosmos.base.tendermint.v1beta1.Service.GetBlockByHeight:output_type -> cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse
	3,  // 24: cosmos.base.tendermint.v1beta1.Service.GetLatestValidatorSet:output_type -> cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse
	1,  // 25: cosmos.base.tendermint.v1beta1.Service.GetValidatorSetByHeight:output_type -> cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse
	20, // [20:26] is the sub-list for method output_type
	14, // [14:20] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_cosmos_base_tendermint_v1beta1_query_proto_init() }
func file_cosmos_base_tendermint_v1beta1_query_proto_init() {
	if File_cosmos_base_tendermint_v1beta1_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValidatorSetByHeightRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValidatorSetByHeightResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestValidatorSetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestValidatorSetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Validator); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockByHeightRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockByHeightResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestBlockRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestBlockResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSyncingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSyncingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Module); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_base_tendermint_v1beta1_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_base_tendermint_v1beta1_query_proto_goTypes,
		DependencyIndexes: file_cosmos_base_tendermint_v1beta1_query_proto_depIdxs,
		MessageInfos:      file_cosmos_base_tendermint_v1beta1_query_proto_msgTypes,
	}.Build()
	File_cosmos_base_tendermint_v1beta1_query_proto = out.File
	file_cosmos_base_tendermint_v1beta1_query_proto_rawDesc = nil
	file_cosmos_base_tendermint_v1beta1_query_proto_goTypes = nil
	file_cosmos_base_tendermint_v1beta1_query_proto_depIdxs = nil
}
