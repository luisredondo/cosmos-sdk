package slashingv1beta1

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/query/v1beta1"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_QueryParamsRequest protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_slashing_v1beta1_query_proto_init()
	md_QueryParamsRequest = File_cosmos_slashing_v1beta1_query_proto.Messages().ByName("QueryParamsRequest")
}

var _ protoreflect.Message = (*fastReflection_QueryParamsRequest)(nil)

type fastReflection_QueryParamsRequest QueryParamsRequest

func (x *QueryParamsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryParamsRequest)(x)
}

func (x *QueryParamsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_slashing_v1beta1_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryParamsRequest_messageType fastReflection_QueryParamsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryParamsRequest_messageType{}

type fastReflection_QueryParamsRequest_messageType struct{}

func (x fastReflection_QueryParamsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryParamsRequest)(nil)
}
func (x fastReflection_QueryParamsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryParamsRequest)
}
func (x fastReflection_QueryParamsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryParamsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryParamsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryParamsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryParamsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryParamsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryParamsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryParamsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryParamsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryParamsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryParamsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_QueryParamsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QueryParamsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryParamsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QueryParamsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryParamsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QueryParamsRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryParamsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QueryParamsRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryParamsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QueryParamsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryParamsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QueryParamsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryParamsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.slashing.v1beta1.QueryParamsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryParamsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryParamsRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryParamsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryParamsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryParamsRequest)
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
		x := input.Message.Interface().(*QueryParamsRequest)
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
		x := input.Message.Interface().(*QueryParamsRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
	md_QueryParamsResponse        protoreflect.MessageDescriptor
	fd_QueryParamsResponse_params protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_slashing_v1beta1_query_proto_init()
	md_QueryParamsResponse = File_cosmos_slashing_v1beta1_query_proto.Messages().ByName("QueryParamsResponse")
	fd_QueryParamsResponse_params = md_QueryParamsResponse.Fields().ByName("params")
}

var _ protoreflect.Message = (*fastReflection_QueryParamsResponse)(nil)

type fastReflection_QueryParamsResponse QueryParamsResponse

func (x *QueryParamsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryParamsResponse)(x)
}

func (x *QueryParamsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_slashing_v1beta1_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryParamsResponse_messageType fastReflection_QueryParamsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryParamsResponse_messageType{}

type fastReflection_QueryParamsResponse_messageType struct{}

func (x fastReflection_QueryParamsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryParamsResponse)(nil)
}
func (x fastReflection_QueryParamsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryParamsResponse)
}
func (x fastReflection_QueryParamsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryParamsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryParamsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryParamsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryParamsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryParamsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryParamsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryParamsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryParamsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryParamsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryParamsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Params != nil {
		value := protoreflect.ValueOfMessage(x.Params.ProtoReflect())
		if !f(fd_QueryParamsResponse_params, value) {
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
func (x *fastReflection_QueryParamsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QueryParamsResponse.params":
		return x.Params != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QueryParamsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryParamsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QueryParamsResponse.params":
		x.Params = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QueryParamsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryParamsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.slashing.v1beta1.QueryParamsResponse.params":
		value := x.Params
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QueryParamsResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryParamsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QueryParamsResponse.params":
		x.Params = value.Message().Interface().(*Params)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QueryParamsResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryParamsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QueryParamsResponse.params":
		if x.Params == nil {
			x.Params = new(Params)
		}
		return protoreflect.ValueOfMessage(x.Params.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QueryParamsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryParamsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QueryParamsResponse.params":
		m := new(Params)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QueryParamsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryParamsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.slashing.v1beta1.QueryParamsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryParamsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryParamsResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryParamsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryParamsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryParamsResponse)
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
		if x.Params != nil {
			l = options.Size(x.Params)
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
		x := input.Message.Interface().(*QueryParamsResponse)
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
		if x.Params != nil {
			encoded, err := options.Marshal(x.Params)
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
		x := input.Message.Interface().(*QueryParamsResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
				if x.Params == nil {
					x.Params = &Params{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Params); err != nil {
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
	md_QuerySigningInfoRequest              protoreflect.MessageDescriptor
	fd_QuerySigningInfoRequest_cons_address protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_slashing_v1beta1_query_proto_init()
	md_QuerySigningInfoRequest = File_cosmos_slashing_v1beta1_query_proto.Messages().ByName("QuerySigningInfoRequest")
	fd_QuerySigningInfoRequest_cons_address = md_QuerySigningInfoRequest.Fields().ByName("cons_address")
}

var _ protoreflect.Message = (*fastReflection_QuerySigningInfoRequest)(nil)

type fastReflection_QuerySigningInfoRequest QuerySigningInfoRequest

func (x *QuerySigningInfoRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QuerySigningInfoRequest)(x)
}

func (x *QuerySigningInfoRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_slashing_v1beta1_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QuerySigningInfoRequest_messageType fastReflection_QuerySigningInfoRequest_messageType
var _ protoreflect.MessageType = fastReflection_QuerySigningInfoRequest_messageType{}

type fastReflection_QuerySigningInfoRequest_messageType struct{}

func (x fastReflection_QuerySigningInfoRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QuerySigningInfoRequest)(nil)
}
func (x fastReflection_QuerySigningInfoRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QuerySigningInfoRequest)
}
func (x fastReflection_QuerySigningInfoRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QuerySigningInfoRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QuerySigningInfoRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QuerySigningInfoRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QuerySigningInfoRequest) Type() protoreflect.MessageType {
	return _fastReflection_QuerySigningInfoRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QuerySigningInfoRequest) New() protoreflect.Message {
	return new(fastReflection_QuerySigningInfoRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QuerySigningInfoRequest) Interface() protoreflect.ProtoMessage {
	return (*QuerySigningInfoRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QuerySigningInfoRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ConsAddress != "" {
		value := protoreflect.ValueOfString(x.ConsAddress)
		if !f(fd_QuerySigningInfoRequest_cons_address, value) {
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
func (x *fastReflection_QuerySigningInfoRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfoRequest.cons_address":
		return x.ConsAddress != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfoRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QuerySigningInfoRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfoRequest.cons_address":
		x.ConsAddress = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfoRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QuerySigningInfoRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfoRequest.cons_address":
		value := x.ConsAddress
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfoRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QuerySigningInfoRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfoRequest.cons_address":
		x.ConsAddress = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfoRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QuerySigningInfoRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfoRequest.cons_address":
		panic(fmt.Errorf("field cons_address of message cosmos.slashing.v1beta1.QuerySigningInfoRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfoRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QuerySigningInfoRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfoRequest.cons_address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfoRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QuerySigningInfoRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.slashing.v1beta1.QuerySigningInfoRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QuerySigningInfoRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QuerySigningInfoRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QuerySigningInfoRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QuerySigningInfoRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QuerySigningInfoRequest)
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
		l = len(x.ConsAddress)
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
		x := input.Message.Interface().(*QuerySigningInfoRequest)
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
		if len(x.ConsAddress) > 0 {
			i -= len(x.ConsAddress)
			copy(dAtA[i:], x.ConsAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ConsAddress)))
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
		x := input.Message.Interface().(*QuerySigningInfoRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QuerySigningInfoRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QuerySigningInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ConsAddress", wireType)
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
				x.ConsAddress = string(dAtA[iNdEx:postIndex])
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
	md_QuerySigningInfoResponse                  protoreflect.MessageDescriptor
	fd_QuerySigningInfoResponse_val_signing_info protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_slashing_v1beta1_query_proto_init()
	md_QuerySigningInfoResponse = File_cosmos_slashing_v1beta1_query_proto.Messages().ByName("QuerySigningInfoResponse")
	fd_QuerySigningInfoResponse_val_signing_info = md_QuerySigningInfoResponse.Fields().ByName("val_signing_info")
}

var _ protoreflect.Message = (*fastReflection_QuerySigningInfoResponse)(nil)

type fastReflection_QuerySigningInfoResponse QuerySigningInfoResponse

func (x *QuerySigningInfoResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QuerySigningInfoResponse)(x)
}

func (x *QuerySigningInfoResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_slashing_v1beta1_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QuerySigningInfoResponse_messageType fastReflection_QuerySigningInfoResponse_messageType
var _ protoreflect.MessageType = fastReflection_QuerySigningInfoResponse_messageType{}

type fastReflection_QuerySigningInfoResponse_messageType struct{}

func (x fastReflection_QuerySigningInfoResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QuerySigningInfoResponse)(nil)
}
func (x fastReflection_QuerySigningInfoResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QuerySigningInfoResponse)
}
func (x fastReflection_QuerySigningInfoResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QuerySigningInfoResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QuerySigningInfoResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QuerySigningInfoResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QuerySigningInfoResponse) Type() protoreflect.MessageType {
	return _fastReflection_QuerySigningInfoResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QuerySigningInfoResponse) New() protoreflect.Message {
	return new(fastReflection_QuerySigningInfoResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QuerySigningInfoResponse) Interface() protoreflect.ProtoMessage {
	return (*QuerySigningInfoResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QuerySigningInfoResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ValSigningInfo != nil {
		value := protoreflect.ValueOfMessage(x.ValSigningInfo.ProtoReflect())
		if !f(fd_QuerySigningInfoResponse_val_signing_info, value) {
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
func (x *fastReflection_QuerySigningInfoResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfoResponse.val_signing_info":
		return x.ValSigningInfo != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfoResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QuerySigningInfoResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfoResponse.val_signing_info":
		x.ValSigningInfo = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfoResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QuerySigningInfoResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfoResponse.val_signing_info":
		value := x.ValSigningInfo
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfoResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QuerySigningInfoResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfoResponse.val_signing_info":
		x.ValSigningInfo = value.Message().Interface().(*ValidatorSigningInfo)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfoResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QuerySigningInfoResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfoResponse.val_signing_info":
		if x.ValSigningInfo == nil {
			x.ValSigningInfo = new(ValidatorSigningInfo)
		}
		return protoreflect.ValueOfMessage(x.ValSigningInfo.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfoResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QuerySigningInfoResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfoResponse.val_signing_info":
		m := new(ValidatorSigningInfo)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfoResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QuerySigningInfoResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.slashing.v1beta1.QuerySigningInfoResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QuerySigningInfoResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QuerySigningInfoResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QuerySigningInfoResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QuerySigningInfoResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QuerySigningInfoResponse)
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
		if x.ValSigningInfo != nil {
			l = options.Size(x.ValSigningInfo)
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
		x := input.Message.Interface().(*QuerySigningInfoResponse)
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
		if x.ValSigningInfo != nil {
			encoded, err := options.Marshal(x.ValSigningInfo)
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
		x := input.Message.Interface().(*QuerySigningInfoResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QuerySigningInfoResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QuerySigningInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValSigningInfo", wireType)
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
				if x.ValSigningInfo == nil {
					x.ValSigningInfo = &ValidatorSigningInfo{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ValSigningInfo); err != nil {
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
	md_QuerySigningInfosRequest            protoreflect.MessageDescriptor
	fd_QuerySigningInfosRequest_pagination protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_slashing_v1beta1_query_proto_init()
	md_QuerySigningInfosRequest = File_cosmos_slashing_v1beta1_query_proto.Messages().ByName("QuerySigningInfosRequest")
	fd_QuerySigningInfosRequest_pagination = md_QuerySigningInfosRequest.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QuerySigningInfosRequest)(nil)

type fastReflection_QuerySigningInfosRequest QuerySigningInfosRequest

func (x *QuerySigningInfosRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QuerySigningInfosRequest)(x)
}

func (x *QuerySigningInfosRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_slashing_v1beta1_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QuerySigningInfosRequest_messageType fastReflection_QuerySigningInfosRequest_messageType
var _ protoreflect.MessageType = fastReflection_QuerySigningInfosRequest_messageType{}

type fastReflection_QuerySigningInfosRequest_messageType struct{}

func (x fastReflection_QuerySigningInfosRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QuerySigningInfosRequest)(nil)
}
func (x fastReflection_QuerySigningInfosRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QuerySigningInfosRequest)
}
func (x fastReflection_QuerySigningInfosRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QuerySigningInfosRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QuerySigningInfosRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QuerySigningInfosRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QuerySigningInfosRequest) Type() protoreflect.MessageType {
	return _fastReflection_QuerySigningInfosRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QuerySigningInfosRequest) New() protoreflect.Message {
	return new(fastReflection_QuerySigningInfosRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QuerySigningInfosRequest) Interface() protoreflect.ProtoMessage {
	return (*QuerySigningInfosRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QuerySigningInfosRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QuerySigningInfosRequest_pagination, value) {
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
func (x *fastReflection_QuerySigningInfosRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfosRequest.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfosRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfosRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QuerySigningInfosRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfosRequest.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfosRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfosRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QuerySigningInfosRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfosRequest.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfosRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfosRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QuerySigningInfosRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfosRequest.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageRequest)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfosRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfosRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QuerySigningInfosRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfosRequest.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageRequest)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfosRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfosRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QuerySigningInfosRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfosRequest.pagination":
		m := new(v1beta1.PageRequest)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfosRequest"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfosRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QuerySigningInfosRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.slashing.v1beta1.QuerySigningInfosRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QuerySigningInfosRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QuerySigningInfosRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QuerySigningInfosRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QuerySigningInfosRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QuerySigningInfosRequest)
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
		x := input.Message.Interface().(*QuerySigningInfosRequest)
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
		x := input.Message.Interface().(*QuerySigningInfosRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QuerySigningInfosRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QuerySigningInfosRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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

var _ protoreflect.List = (*_QuerySigningInfosResponse_1_list)(nil)

type _QuerySigningInfosResponse_1_list struct {
	list *[]*ValidatorSigningInfo
}

func (x *_QuerySigningInfosResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QuerySigningInfosResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QuerySigningInfosResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ValidatorSigningInfo)
	(*x.list)[i] = concreteValue
}

func (x *_QuerySigningInfosResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ValidatorSigningInfo)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QuerySigningInfosResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(ValidatorSigningInfo)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QuerySigningInfosResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QuerySigningInfosResponse_1_list) NewElement() protoreflect.Value {
	v := new(ValidatorSigningInfo)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QuerySigningInfosResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QuerySigningInfosResponse            protoreflect.MessageDescriptor
	fd_QuerySigningInfosResponse_info       protoreflect.FieldDescriptor
	fd_QuerySigningInfosResponse_pagination protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_slashing_v1beta1_query_proto_init()
	md_QuerySigningInfosResponse = File_cosmos_slashing_v1beta1_query_proto.Messages().ByName("QuerySigningInfosResponse")
	fd_QuerySigningInfosResponse_info = md_QuerySigningInfosResponse.Fields().ByName("info")
	fd_QuerySigningInfosResponse_pagination = md_QuerySigningInfosResponse.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QuerySigningInfosResponse)(nil)

type fastReflection_QuerySigningInfosResponse QuerySigningInfosResponse

func (x *QuerySigningInfosResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QuerySigningInfosResponse)(x)
}

func (x *QuerySigningInfosResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_slashing_v1beta1_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QuerySigningInfosResponse_messageType fastReflection_QuerySigningInfosResponse_messageType
var _ protoreflect.MessageType = fastReflection_QuerySigningInfosResponse_messageType{}

type fastReflection_QuerySigningInfosResponse_messageType struct{}

func (x fastReflection_QuerySigningInfosResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QuerySigningInfosResponse)(nil)
}
func (x fastReflection_QuerySigningInfosResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QuerySigningInfosResponse)
}
func (x fastReflection_QuerySigningInfosResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QuerySigningInfosResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QuerySigningInfosResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QuerySigningInfosResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QuerySigningInfosResponse) Type() protoreflect.MessageType {
	return _fastReflection_QuerySigningInfosResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QuerySigningInfosResponse) New() protoreflect.Message {
	return new(fastReflection_QuerySigningInfosResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QuerySigningInfosResponse) Interface() protoreflect.ProtoMessage {
	return (*QuerySigningInfosResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QuerySigningInfosResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Info) != 0 {
		value := protoreflect.ValueOfList(&_QuerySigningInfosResponse_1_list{list: &x.Info})
		if !f(fd_QuerySigningInfosResponse_info, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QuerySigningInfosResponse_pagination, value) {
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
func (x *fastReflection_QuerySigningInfosResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfosResponse.info":
		return len(x.Info) != 0
	case "cosmos.slashing.v1beta1.QuerySigningInfosResponse.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfosResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfosResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QuerySigningInfosResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfosResponse.info":
		x.Info = nil
	case "cosmos.slashing.v1beta1.QuerySigningInfosResponse.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfosResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfosResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QuerySigningInfosResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfosResponse.info":
		if len(x.Info) == 0 {
			return protoreflect.ValueOfList(&_QuerySigningInfosResponse_1_list{})
		}
		listValue := &_QuerySigningInfosResponse_1_list{list: &x.Info}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.slashing.v1beta1.QuerySigningInfosResponse.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfosResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfosResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QuerySigningInfosResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfosResponse.info":
		lv := value.List()
		clv := lv.(*_QuerySigningInfosResponse_1_list)
		x.Info = *clv.list
	case "cosmos.slashing.v1beta1.QuerySigningInfosResponse.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfosResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfosResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QuerySigningInfosResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfosResponse.info":
		if x.Info == nil {
			x.Info = []*ValidatorSigningInfo{}
		}
		value := &_QuerySigningInfosResponse_1_list{list: &x.Info}
		return protoreflect.ValueOfList(value)
	case "cosmos.slashing.v1beta1.QuerySigningInfosResponse.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageResponse)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfosResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfosResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QuerySigningInfosResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.QuerySigningInfosResponse.info":
		list := []*ValidatorSigningInfo{}
		return protoreflect.ValueOfList(&_QuerySigningInfosResponse_1_list{list: &list})
	case "cosmos.slashing.v1beta1.QuerySigningInfosResponse.pagination":
		m := new(v1beta1.PageResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.QuerySigningInfosResponse"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.QuerySigningInfosResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QuerySigningInfosResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.slashing.v1beta1.QuerySigningInfosResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QuerySigningInfosResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QuerySigningInfosResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QuerySigningInfosResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QuerySigningInfosResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QuerySigningInfosResponse)
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
		if len(x.Info) > 0 {
			for _, e := range x.Info {
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
		x := input.Message.Interface().(*QuerySigningInfosResponse)
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
		if len(x.Info) > 0 {
			for iNdEx := len(x.Info) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Info[iNdEx])
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
		x := input.Message.Interface().(*QuerySigningInfosResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QuerySigningInfosResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QuerySigningInfosResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
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
				x.Info = append(x.Info, &ValidatorSigningInfo{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Info[len(x.Info)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Params queries the parameters of slashing module
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// SigningInfo queries the signing info of given cons address
	SigningInfo(ctx context.Context, in *QuerySigningInfoRequest, opts ...grpc.CallOption) (*QuerySigningInfoResponse, error)
	// SigningInfos queries signing info of all validators
	SigningInfos(ctx context.Context, in *QuerySigningInfosRequest, opts ...grpc.CallOption) (*QuerySigningInfosResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.slashing.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SigningInfo(ctx context.Context, in *QuerySigningInfoRequest, opts ...grpc.CallOption) (*QuerySigningInfoResponse, error) {
	out := new(QuerySigningInfoResponse)
	err := c.cc.Invoke(ctx, "/cosmos.slashing.v1beta1.Query/SigningInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SigningInfos(ctx context.Context, in *QuerySigningInfosRequest, opts ...grpc.CallOption) (*QuerySigningInfosResponse, error) {
	out := new(QuerySigningInfosResponse)
	err := c.cc.Invoke(ctx, "/cosmos.slashing.v1beta1.Query/SigningInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Params queries the parameters of slashing module
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// SigningInfo queries the signing info of given cons address
	SigningInfo(context.Context, *QuerySigningInfoRequest) (*QuerySigningInfoResponse, error)
	// SigningInfos queries signing info of all validators
	SigningInfos(context.Context, *QuerySigningInfosRequest) (*QuerySigningInfosResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) SigningInfo(context.Context, *QuerySigningInfoRequest) (*QuerySigningInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SigningInfo not implemented")
}
func (UnimplementedQueryServer) SigningInfos(context.Context, *QuerySigningInfosRequest) (*QuerySigningInfosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SigningInfos not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.slashing.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SigningInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySigningInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SigningInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.slashing.v1beta1.Query/SigningInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SigningInfo(ctx, req.(*QuerySigningInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SigningInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySigningInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SigningInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.slashing.v1beta1.Query/SigningInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SigningInfos(ctx, req.(*QuerySigningInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.slashing.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "SigningInfo",
			Handler:    _Query_SigningInfo_Handler,
		},
		{
			MethodName: "SigningInfos",
			Handler:    _Query_SigningInfos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/slashing/v1beta1/query.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/slashing/v1beta1/query.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// QueryParamsRequest is the request type for the Query/Params RPC method
type QueryParamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryParamsRequest) Reset() {
	*x = QueryParamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_slashing_v1beta1_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParamsRequest) ProtoMessage() {}

// Deprecated: Use QueryParamsRequest.ProtoReflect.Descriptor instead.
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_slashing_v1beta1_query_proto_rawDescGZIP(), []int{0}
}

// QueryParamsResponse is the response type for the Query/Params RPC method
type QueryParamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *QueryParamsResponse) Reset() {
	*x = QueryParamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_slashing_v1beta1_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParamsResponse) ProtoMessage() {}

// Deprecated: Use QueryParamsResponse.ProtoReflect.Descriptor instead.
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_slashing_v1beta1_query_proto_rawDescGZIP(), []int{1}
}

func (x *QueryParamsResponse) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

// QuerySigningInfoRequest is the request type for the Query/SigningInfo RPC
// method
type QuerySigningInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cons_address is the address to query signing info of
	ConsAddress string `protobuf:"bytes,1,opt,name=cons_address,json=consAddress,proto3" json:"cons_address,omitempty"`
}

func (x *QuerySigningInfoRequest) Reset() {
	*x = QuerySigningInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_slashing_v1beta1_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySigningInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySigningInfoRequest) ProtoMessage() {}

// Deprecated: Use QuerySigningInfoRequest.ProtoReflect.Descriptor instead.
func (*QuerySigningInfoRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_slashing_v1beta1_query_proto_rawDescGZIP(), []int{2}
}

func (x *QuerySigningInfoRequest) GetConsAddress() string {
	if x != nil {
		return x.ConsAddress
	}
	return ""
}

// QuerySigningInfoResponse is the response type for the Query/SigningInfo RPC
// method
type QuerySigningInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// val_signing_info is the signing info of requested val cons address
	ValSigningInfo *ValidatorSigningInfo `protobuf:"bytes,1,opt,name=val_signing_info,json=valSigningInfo,proto3" json:"val_signing_info,omitempty"`
}

func (x *QuerySigningInfoResponse) Reset() {
	*x = QuerySigningInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_slashing_v1beta1_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySigningInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySigningInfoResponse) ProtoMessage() {}

// Deprecated: Use QuerySigningInfoResponse.ProtoReflect.Descriptor instead.
func (*QuerySigningInfoResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_slashing_v1beta1_query_proto_rawDescGZIP(), []int{3}
}

func (x *QuerySigningInfoResponse) GetValSigningInfo() *ValidatorSigningInfo {
	if x != nil {
		return x.ValSigningInfo
	}
	return nil
}

// QuerySigningInfosRequest is the request type for the Query/SigningInfos RPC
// method
type QuerySigningInfosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *v1beta1.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QuerySigningInfosRequest) Reset() {
	*x = QuerySigningInfosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_slashing_v1beta1_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySigningInfosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySigningInfosRequest) ProtoMessage() {}

// Deprecated: Use QuerySigningInfosRequest.ProtoReflect.Descriptor instead.
func (*QuerySigningInfosRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_slashing_v1beta1_query_proto_rawDescGZIP(), []int{4}
}

func (x *QuerySigningInfosRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QuerySigningInfosResponse is the response type for the Query/SigningInfos RPC
// method
type QuerySigningInfosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// info is the signing info of all validators
	Info       []*ValidatorSigningInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
	Pagination *v1beta1.PageResponse   `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QuerySigningInfosResponse) Reset() {
	*x = QuerySigningInfosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_slashing_v1beta1_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySigningInfosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySigningInfosResponse) ProtoMessage() {}

// Deprecated: Use QuerySigningInfosResponse.ProtoReflect.Descriptor instead.
func (*QuerySigningInfosResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_slashing_v1beta1_query_proto_rawDescGZIP(), []int{5}
}

func (x *QuerySigningInfosResponse) GetInfo() []*ValidatorSigningInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *QuerySigningInfosResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_cosmos_slashing_v1beta1_query_proto protoreflect.FileDescriptor

var file_cosmos_slashing_v1beta1_query_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x6c,
	0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x2a,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x14, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x54, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42,
	0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x56, 0x0a,
	0x17, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x73,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18,
	0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x79, 0x0a, 0x18, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5d, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00,
	0x52, 0x0e, 0x76, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x62, 0x0a, 0x18, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xad, 0x01, 0x0a, 0x19, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x47, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x42,
	0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x47, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x32, 0xf2, 0x03, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x8c,
	0x01, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0xb1, 0x01,
	0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x30, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x69, 0x67,
	0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x3d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x12, 0x35, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x7d, 0x12, 0xa5, 0x01, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x6c, 0x61, 0x73,
	0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73,
	0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x28, 0x12, 0x26, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x42, 0xf1, 0x01, 0x0a, 0x1b, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x3b, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x17, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xca, 0x02, 0x17, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x6c, 0x61, 0x73, 0x68,
	0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x23, 0x43, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x19, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x53, 0x6c, 0x61, 0x73,
	0x68, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_slashing_v1beta1_query_proto_rawDescOnce sync.Once
	file_cosmos_slashing_v1beta1_query_proto_rawDescData = file_cosmos_slashing_v1beta1_query_proto_rawDesc
)

func file_cosmos_slashing_v1beta1_query_proto_rawDescGZIP() []byte {
	file_cosmos_slashing_v1beta1_query_proto_rawDescOnce.Do(func() {
		file_cosmos_slashing_v1beta1_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_slashing_v1beta1_query_proto_rawDescData)
	})
	return file_cosmos_slashing_v1beta1_query_proto_rawDescData
}

var file_cosmos_slashing_v1beta1_query_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cosmos_slashing_v1beta1_query_proto_goTypes = []interface{}{
	(*QueryParamsRequest)(nil),        // 0: cosmos.slashing.v1beta1.QueryParamsRequest
	(*QueryParamsResponse)(nil),       // 1: cosmos.slashing.v1beta1.QueryParamsResponse
	(*QuerySigningInfoRequest)(nil),   // 2: cosmos.slashing.v1beta1.QuerySigningInfoRequest
	(*QuerySigningInfoResponse)(nil),  // 3: cosmos.slashing.v1beta1.QuerySigningInfoResponse
	(*QuerySigningInfosRequest)(nil),  // 4: cosmos.slashing.v1beta1.QuerySigningInfosRequest
	(*QuerySigningInfosResponse)(nil), // 5: cosmos.slashing.v1beta1.QuerySigningInfosResponse
	(*Params)(nil),                    // 6: cosmos.slashing.v1beta1.Params
	(*ValidatorSigningInfo)(nil),      // 7: cosmos.slashing.v1beta1.ValidatorSigningInfo
	(*v1beta1.PageRequest)(nil),       // 8: cosmos.base.query.v1beta1.PageRequest
	(*v1beta1.PageResponse)(nil),      // 9: cosmos.base.query.v1beta1.PageResponse
}
var file_cosmos_slashing_v1beta1_query_proto_depIdxs = []int32{
	6, // 0: cosmos.slashing.v1beta1.QueryParamsResponse.params:type_name -> cosmos.slashing.v1beta1.Params
	7, // 1: cosmos.slashing.v1beta1.QuerySigningInfoResponse.val_signing_info:type_name -> cosmos.slashing.v1beta1.ValidatorSigningInfo
	8, // 2: cosmos.slashing.v1beta1.QuerySigningInfosRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	7, // 3: cosmos.slashing.v1beta1.QuerySigningInfosResponse.info:type_name -> cosmos.slashing.v1beta1.ValidatorSigningInfo
	9, // 4: cosmos.slashing.v1beta1.QuerySigningInfosResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	0, // 5: cosmos.slashing.v1beta1.Query.Params:input_type -> cosmos.slashing.v1beta1.QueryParamsRequest
	2, // 6: cosmos.slashing.v1beta1.Query.SigningInfo:input_type -> cosmos.slashing.v1beta1.QuerySigningInfoRequest
	4, // 7: cosmos.slashing.v1beta1.Query.SigningInfos:input_type -> cosmos.slashing.v1beta1.QuerySigningInfosRequest
	1, // 8: cosmos.slashing.v1beta1.Query.Params:output_type -> cosmos.slashing.v1beta1.QueryParamsResponse
	3, // 9: cosmos.slashing.v1beta1.Query.SigningInfo:output_type -> cosmos.slashing.v1beta1.QuerySigningInfoResponse
	5, // 10: cosmos.slashing.v1beta1.Query.SigningInfos:output_type -> cosmos.slashing.v1beta1.QuerySigningInfosResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cosmos_slashing_v1beta1_query_proto_init() }
func file_cosmos_slashing_v1beta1_query_proto_init() {
	if File_cosmos_slashing_v1beta1_query_proto != nil {
		return
	}
	file_cosmos_slashing_v1beta1_slashing_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_slashing_v1beta1_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParamsRequest); i {
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
		file_cosmos_slashing_v1beta1_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParamsResponse); i {
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
		file_cosmos_slashing_v1beta1_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySigningInfoRequest); i {
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
		file_cosmos_slashing_v1beta1_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySigningInfoResponse); i {
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
		file_cosmos_slashing_v1beta1_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySigningInfosRequest); i {
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
		file_cosmos_slashing_v1beta1_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySigningInfosResponse); i {
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
			RawDescriptor: file_cosmos_slashing_v1beta1_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_slashing_v1beta1_query_proto_goTypes,
		DependencyIndexes: file_cosmos_slashing_v1beta1_query_proto_depIdxs,
		MessageInfos:      file_cosmos_slashing_v1beta1_query_proto_msgTypes,
	}.Build()
	File_cosmos_slashing_v1beta1_query_proto = out.File
	file_cosmos_slashing_v1beta1_query_proto_rawDesc = nil
	file_cosmos_slashing_v1beta1_query_proto_goTypes = nil
	file_cosmos_slashing_v1beta1_query_proto_depIdxs = nil
}
