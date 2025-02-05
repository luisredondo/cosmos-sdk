package feegrantv1beta1

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
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
	md_MsgGrantAllowance           protoreflect.MessageDescriptor
	fd_MsgGrantAllowance_granter   protoreflect.FieldDescriptor
	fd_MsgGrantAllowance_grantee   protoreflect.FieldDescriptor
	fd_MsgGrantAllowance_allowance protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_feegrant_v1beta1_tx_proto_init()
	md_MsgGrantAllowance = File_cosmos_feegrant_v1beta1_tx_proto.Messages().ByName("MsgGrantAllowance")
	fd_MsgGrantAllowance_granter = md_MsgGrantAllowance.Fields().ByName("granter")
	fd_MsgGrantAllowance_grantee = md_MsgGrantAllowance.Fields().ByName("grantee")
	fd_MsgGrantAllowance_allowance = md_MsgGrantAllowance.Fields().ByName("allowance")
}

var _ protoreflect.Message = (*fastReflection_MsgGrantAllowance)(nil)

type fastReflection_MsgGrantAllowance MsgGrantAllowance

func (x *MsgGrantAllowance) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgGrantAllowance)(x)
}

func (x *MsgGrantAllowance) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_feegrant_v1beta1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgGrantAllowance_messageType fastReflection_MsgGrantAllowance_messageType
var _ protoreflect.MessageType = fastReflection_MsgGrantAllowance_messageType{}

type fastReflection_MsgGrantAllowance_messageType struct{}

func (x fastReflection_MsgGrantAllowance_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgGrantAllowance)(nil)
}
func (x fastReflection_MsgGrantAllowance_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgGrantAllowance)
}
func (x fastReflection_MsgGrantAllowance_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgGrantAllowance
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgGrantAllowance) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgGrantAllowance
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgGrantAllowance) Type() protoreflect.MessageType {
	return _fastReflection_MsgGrantAllowance_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgGrantAllowance) New() protoreflect.Message {
	return new(fastReflection_MsgGrantAllowance)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgGrantAllowance) Interface() protoreflect.ProtoMessage {
	return (*MsgGrantAllowance)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgGrantAllowance) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Granter != "" {
		value := protoreflect.ValueOfString(x.Granter)
		if !f(fd_MsgGrantAllowance_granter, value) {
			return
		}
	}
	if x.Grantee != "" {
		value := protoreflect.ValueOfString(x.Grantee)
		if !f(fd_MsgGrantAllowance_grantee, value) {
			return
		}
	}
	if x.Allowance != nil {
		value := protoreflect.ValueOfMessage(x.Allowance.ProtoReflect())
		if !f(fd_MsgGrantAllowance_allowance, value) {
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
func (x *fastReflection_MsgGrantAllowance) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.granter":
		return x.Granter != ""
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.grantee":
		return x.Grantee != ""
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.allowance":
		return x.Allowance != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgGrantAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgGrantAllowance does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgGrantAllowance) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.granter":
		x.Granter = ""
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.grantee":
		x.Grantee = ""
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.allowance":
		x.Allowance = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgGrantAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgGrantAllowance does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgGrantAllowance) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.granter":
		value := x.Granter
		return protoreflect.ValueOfString(value)
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.grantee":
		value := x.Grantee
		return protoreflect.ValueOfString(value)
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.allowance":
		value := x.Allowance
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgGrantAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgGrantAllowance does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgGrantAllowance) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.granter":
		x.Granter = value.Interface().(string)
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.grantee":
		x.Grantee = value.Interface().(string)
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.allowance":
		x.Allowance = value.Message().Interface().(*anypb.Any)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgGrantAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgGrantAllowance does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgGrantAllowance) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.allowance":
		if x.Allowance == nil {
			x.Allowance = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.Allowance.ProtoReflect())
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.granter":
		panic(fmt.Errorf("field granter of message cosmos.feegrant.v1beta1.MsgGrantAllowance is not mutable"))
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.grantee":
		panic(fmt.Errorf("field grantee of message cosmos.feegrant.v1beta1.MsgGrantAllowance is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgGrantAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgGrantAllowance does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgGrantAllowance) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.granter":
		return protoreflect.ValueOfString("")
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.grantee":
		return protoreflect.ValueOfString("")
	case "cosmos.feegrant.v1beta1.MsgGrantAllowance.allowance":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgGrantAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgGrantAllowance does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgGrantAllowance) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.feegrant.v1beta1.MsgGrantAllowance", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgGrantAllowance) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgGrantAllowance) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgGrantAllowance) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgGrantAllowance) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgGrantAllowance)
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
		l = len(x.Granter)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Grantee)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Allowance != nil {
			l = options.Size(x.Allowance)
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
		x := input.Message.Interface().(*MsgGrantAllowance)
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
		if x.Allowance != nil {
			encoded, err := options.Marshal(x.Allowance)
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
		if len(x.Grantee) > 0 {
			i -= len(x.Grantee)
			copy(dAtA[i:], x.Grantee)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Grantee)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Granter) > 0 {
			i -= len(x.Granter)
			copy(dAtA[i:], x.Granter)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Granter)))
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
		x := input.Message.Interface().(*MsgGrantAllowance)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgGrantAllowance: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgGrantAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
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
				x.Granter = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
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
				x.Grantee = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
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
				if x.Allowance == nil {
					x.Allowance = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Allowance); err != nil {
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
	md_MsgGrantAllowanceResponse protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_feegrant_v1beta1_tx_proto_init()
	md_MsgGrantAllowanceResponse = File_cosmos_feegrant_v1beta1_tx_proto.Messages().ByName("MsgGrantAllowanceResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgGrantAllowanceResponse)(nil)

type fastReflection_MsgGrantAllowanceResponse MsgGrantAllowanceResponse

func (x *MsgGrantAllowanceResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgGrantAllowanceResponse)(x)
}

func (x *MsgGrantAllowanceResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_feegrant_v1beta1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgGrantAllowanceResponse_messageType fastReflection_MsgGrantAllowanceResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgGrantAllowanceResponse_messageType{}

type fastReflection_MsgGrantAllowanceResponse_messageType struct{}

func (x fastReflection_MsgGrantAllowanceResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgGrantAllowanceResponse)(nil)
}
func (x fastReflection_MsgGrantAllowanceResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgGrantAllowanceResponse)
}
func (x fastReflection_MsgGrantAllowanceResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgGrantAllowanceResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgGrantAllowanceResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgGrantAllowanceResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgGrantAllowanceResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgGrantAllowanceResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgGrantAllowanceResponse) New() protoreflect.Message {
	return new(fastReflection_MsgGrantAllowanceResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgGrantAllowanceResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgGrantAllowanceResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgGrantAllowanceResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgGrantAllowanceResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgGrantAllowanceResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgGrantAllowanceResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgGrantAllowanceResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgGrantAllowanceResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgGrantAllowanceResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgGrantAllowanceResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgGrantAllowanceResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgGrantAllowanceResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgGrantAllowanceResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgGrantAllowanceResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgGrantAllowanceResponse)
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
		x := input.Message.Interface().(*MsgGrantAllowanceResponse)
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
		x := input.Message.Interface().(*MsgGrantAllowanceResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgGrantAllowanceResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgGrantAllowanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
	md_MsgRevokeAllowance         protoreflect.MessageDescriptor
	fd_MsgRevokeAllowance_granter protoreflect.FieldDescriptor
	fd_MsgRevokeAllowance_grantee protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_feegrant_v1beta1_tx_proto_init()
	md_MsgRevokeAllowance = File_cosmos_feegrant_v1beta1_tx_proto.Messages().ByName("MsgRevokeAllowance")
	fd_MsgRevokeAllowance_granter = md_MsgRevokeAllowance.Fields().ByName("granter")
	fd_MsgRevokeAllowance_grantee = md_MsgRevokeAllowance.Fields().ByName("grantee")
}

var _ protoreflect.Message = (*fastReflection_MsgRevokeAllowance)(nil)

type fastReflection_MsgRevokeAllowance MsgRevokeAllowance

func (x *MsgRevokeAllowance) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgRevokeAllowance)(x)
}

func (x *MsgRevokeAllowance) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_feegrant_v1beta1_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgRevokeAllowance_messageType fastReflection_MsgRevokeAllowance_messageType
var _ protoreflect.MessageType = fastReflection_MsgRevokeAllowance_messageType{}

type fastReflection_MsgRevokeAllowance_messageType struct{}

func (x fastReflection_MsgRevokeAllowance_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgRevokeAllowance)(nil)
}
func (x fastReflection_MsgRevokeAllowance_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgRevokeAllowance)
}
func (x fastReflection_MsgRevokeAllowance_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgRevokeAllowance
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgRevokeAllowance) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgRevokeAllowance
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgRevokeAllowance) Type() protoreflect.MessageType {
	return _fastReflection_MsgRevokeAllowance_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgRevokeAllowance) New() protoreflect.Message {
	return new(fastReflection_MsgRevokeAllowance)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgRevokeAllowance) Interface() protoreflect.ProtoMessage {
	return (*MsgRevokeAllowance)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgRevokeAllowance) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Granter != "" {
		value := protoreflect.ValueOfString(x.Granter)
		if !f(fd_MsgRevokeAllowance_granter, value) {
			return
		}
	}
	if x.Grantee != "" {
		value := protoreflect.ValueOfString(x.Grantee)
		if !f(fd_MsgRevokeAllowance_grantee, value) {
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
func (x *fastReflection_MsgRevokeAllowance) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.MsgRevokeAllowance.granter":
		return x.Granter != ""
	case "cosmos.feegrant.v1beta1.MsgRevokeAllowance.grantee":
		return x.Grantee != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgRevokeAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgRevokeAllowance does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgRevokeAllowance) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.MsgRevokeAllowance.granter":
		x.Granter = ""
	case "cosmos.feegrant.v1beta1.MsgRevokeAllowance.grantee":
		x.Grantee = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgRevokeAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgRevokeAllowance does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgRevokeAllowance) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.feegrant.v1beta1.MsgRevokeAllowance.granter":
		value := x.Granter
		return protoreflect.ValueOfString(value)
	case "cosmos.feegrant.v1beta1.MsgRevokeAllowance.grantee":
		value := x.Grantee
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgRevokeAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgRevokeAllowance does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgRevokeAllowance) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.MsgRevokeAllowance.granter":
		x.Granter = value.Interface().(string)
	case "cosmos.feegrant.v1beta1.MsgRevokeAllowance.grantee":
		x.Grantee = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgRevokeAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgRevokeAllowance does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgRevokeAllowance) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.MsgRevokeAllowance.granter":
		panic(fmt.Errorf("field granter of message cosmos.feegrant.v1beta1.MsgRevokeAllowance is not mutable"))
	case "cosmos.feegrant.v1beta1.MsgRevokeAllowance.grantee":
		panic(fmt.Errorf("field grantee of message cosmos.feegrant.v1beta1.MsgRevokeAllowance is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgRevokeAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgRevokeAllowance does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgRevokeAllowance) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.MsgRevokeAllowance.granter":
		return protoreflect.ValueOfString("")
	case "cosmos.feegrant.v1beta1.MsgRevokeAllowance.grantee":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgRevokeAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgRevokeAllowance does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgRevokeAllowance) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.feegrant.v1beta1.MsgRevokeAllowance", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgRevokeAllowance) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgRevokeAllowance) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgRevokeAllowance) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgRevokeAllowance) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgRevokeAllowance)
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
		l = len(x.Granter)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Grantee)
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
		x := input.Message.Interface().(*MsgRevokeAllowance)
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
		if len(x.Grantee) > 0 {
			i -= len(x.Grantee)
			copy(dAtA[i:], x.Grantee)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Grantee)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Granter) > 0 {
			i -= len(x.Granter)
			copy(dAtA[i:], x.Granter)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Granter)))
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
		x := input.Message.Interface().(*MsgRevokeAllowance)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgRevokeAllowance: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgRevokeAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
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
				x.Granter = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
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
				x.Grantee = string(dAtA[iNdEx:postIndex])
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
	md_MsgRevokeAllowanceResponse protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_feegrant_v1beta1_tx_proto_init()
	md_MsgRevokeAllowanceResponse = File_cosmos_feegrant_v1beta1_tx_proto.Messages().ByName("MsgRevokeAllowanceResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgRevokeAllowanceResponse)(nil)

type fastReflection_MsgRevokeAllowanceResponse MsgRevokeAllowanceResponse

func (x *MsgRevokeAllowanceResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgRevokeAllowanceResponse)(x)
}

func (x *MsgRevokeAllowanceResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_feegrant_v1beta1_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgRevokeAllowanceResponse_messageType fastReflection_MsgRevokeAllowanceResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgRevokeAllowanceResponse_messageType{}

type fastReflection_MsgRevokeAllowanceResponse_messageType struct{}

func (x fastReflection_MsgRevokeAllowanceResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgRevokeAllowanceResponse)(nil)
}
func (x fastReflection_MsgRevokeAllowanceResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgRevokeAllowanceResponse)
}
func (x fastReflection_MsgRevokeAllowanceResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgRevokeAllowanceResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgRevokeAllowanceResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgRevokeAllowanceResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgRevokeAllowanceResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgRevokeAllowanceResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgRevokeAllowanceResponse) New() protoreflect.Message {
	return new(fastReflection_MsgRevokeAllowanceResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgRevokeAllowanceResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgRevokeAllowanceResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgRevokeAllowanceResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgRevokeAllowanceResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgRevokeAllowanceResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgRevokeAllowanceResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgRevokeAllowanceResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgRevokeAllowanceResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgRevokeAllowanceResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgRevokeAllowanceResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgRevokeAllowanceResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgRevokeAllowanceResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgRevokeAllowanceResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgRevokeAllowanceResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgRevokeAllowanceResponse)
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
		x := input.Message.Interface().(*MsgRevokeAllowanceResponse)
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
		x := input.Message.Interface().(*MsgRevokeAllowanceResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgRevokeAllowanceResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgRevokeAllowanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// GrantAllowance grants fee allowance to the grantee on the granter's
	// account with the provided expiration time.
	GrantAllowance(ctx context.Context, in *MsgGrantAllowance, opts ...grpc.CallOption) (*MsgGrantAllowanceResponse, error)
	// RevokeAllowance revokes any fee allowance of granter's account that
	// has been granted to the grantee.
	RevokeAllowance(ctx context.Context, in *MsgRevokeAllowance, opts ...grpc.CallOption) (*MsgRevokeAllowanceResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) GrantAllowance(ctx context.Context, in *MsgGrantAllowance, opts ...grpc.CallOption) (*MsgGrantAllowanceResponse, error) {
	out := new(MsgGrantAllowanceResponse)
	err := c.cc.Invoke(ctx, "/cosmos.feegrant.v1beta1.Msg/GrantAllowance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RevokeAllowance(ctx context.Context, in *MsgRevokeAllowance, opts ...grpc.CallOption) (*MsgRevokeAllowanceResponse, error) {
	out := new(MsgRevokeAllowanceResponse)
	err := c.cc.Invoke(ctx, "/cosmos.feegrant.v1beta1.Msg/RevokeAllowance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// GrantAllowance grants fee allowance to the grantee on the granter's
	// account with the provided expiration time.
	GrantAllowance(context.Context, *MsgGrantAllowance) (*MsgGrantAllowanceResponse, error)
	// RevokeAllowance revokes any fee allowance of granter's account that
	// has been granted to the grantee.
	RevokeAllowance(context.Context, *MsgRevokeAllowance) (*MsgRevokeAllowanceResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) GrantAllowance(context.Context, *MsgGrantAllowance) (*MsgGrantAllowanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantAllowance not implemented")
}
func (UnimplementedMsgServer) RevokeAllowance(context.Context, *MsgRevokeAllowance) (*MsgRevokeAllowanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeAllowance not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_GrantAllowance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGrantAllowance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GrantAllowance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.feegrant.v1beta1.Msg/GrantAllowance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GrantAllowance(ctx, req.(*MsgGrantAllowance))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RevokeAllowance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRevokeAllowance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RevokeAllowance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.feegrant.v1beta1.Msg/RevokeAllowance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RevokeAllowance(ctx, req.(*MsgRevokeAllowance))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.feegrant.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GrantAllowance",
			Handler:    _Msg_GrantAllowance_Handler,
		},
		{
			MethodName: "RevokeAllowance",
			Handler:    _Msg_RevokeAllowance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/feegrant/v1beta1/tx.proto",
}

// Since: cosmos-sdk 0.43

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/feegrant/v1beta1/tx.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MsgGrantAllowance adds permission for Grantee to spend up to Allowance
// of fees from the account of Granter.
type MsgGrantAllowance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// granter is the address of the user granting an allowance of their funds.
	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	// grantee is the address of the user being granted an allowance of another user's funds.
	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// allowance can be any of basic and filtered fee allowance.
	Allowance *anypb.Any `protobuf:"bytes,3,opt,name=allowance,proto3" json:"allowance,omitempty"`
}

func (x *MsgGrantAllowance) Reset() {
	*x = MsgGrantAllowance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_feegrant_v1beta1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgGrantAllowance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgGrantAllowance) ProtoMessage() {}

// Deprecated: Use MsgGrantAllowance.ProtoReflect.Descriptor instead.
func (*MsgGrantAllowance) Descriptor() ([]byte, []int) {
	return file_cosmos_feegrant_v1beta1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgGrantAllowance) GetGranter() string {
	if x != nil {
		return x.Granter
	}
	return ""
}

func (x *MsgGrantAllowance) GetGrantee() string {
	if x != nil {
		return x.Grantee
	}
	return ""
}

func (x *MsgGrantAllowance) GetAllowance() *anypb.Any {
	if x != nil {
		return x.Allowance
	}
	return nil
}

// MsgGrantAllowanceResponse defines the Msg/GrantAllowanceResponse response type.
type MsgGrantAllowanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgGrantAllowanceResponse) Reset() {
	*x = MsgGrantAllowanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_feegrant_v1beta1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgGrantAllowanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgGrantAllowanceResponse) ProtoMessage() {}

// Deprecated: Use MsgGrantAllowanceResponse.ProtoReflect.Descriptor instead.
func (*MsgGrantAllowanceResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_feegrant_v1beta1_tx_proto_rawDescGZIP(), []int{1}
}

// MsgRevokeAllowance removes any existing Allowance from Granter to Grantee.
type MsgRevokeAllowance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// granter is the address of the user granting an allowance of their funds.
	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	// grantee is the address of the user being granted an allowance of another user's funds.
	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
}

func (x *MsgRevokeAllowance) Reset() {
	*x = MsgRevokeAllowance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_feegrant_v1beta1_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRevokeAllowance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRevokeAllowance) ProtoMessage() {}

// Deprecated: Use MsgRevokeAllowance.ProtoReflect.Descriptor instead.
func (*MsgRevokeAllowance) Descriptor() ([]byte, []int) {
	return file_cosmos_feegrant_v1beta1_tx_proto_rawDescGZIP(), []int{2}
}

func (x *MsgRevokeAllowance) GetGranter() string {
	if x != nil {
		return x.Granter
	}
	return ""
}

func (x *MsgRevokeAllowance) GetGrantee() string {
	if x != nil {
		return x.Grantee
	}
	return ""
}

// MsgRevokeAllowanceResponse defines the Msg/RevokeAllowanceResponse response type.
type MsgRevokeAllowanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgRevokeAllowanceResponse) Reset() {
	*x = MsgRevokeAllowanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_feegrant_v1beta1_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRevokeAllowanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRevokeAllowanceResponse) ProtoMessage() {}

// Deprecated: Use MsgRevokeAllowanceResponse.ProtoReflect.Descriptor instead.
func (*MsgRevokeAllowanceResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_feegrant_v1beta1_tx_proto_rawDescGZIP(), []int{3}
}

var File_cosmos_feegrant_v1beta1_tx_proto protoreflect.FileDescriptor

var file_cosmos_feegrant_v1beta1_tx_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x66, 0x65, 0x65, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x11, 0x4d, 0x73, 0x67, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x07, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4,
	0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x12,
	0x45, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x11, 0xca, 0xb4, 0x2d, 0x0d, 0x46, 0x65,
	0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x52, 0x09, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x1b, 0x0a, 0x19, 0x4d, 0x73, 0x67, 0x47, 0x72, 0x61,
	0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x7c, 0x0a, 0x12, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a,
	0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18,
	0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65,
	0x65, 0x22, 0x1c, 0x0a, 0x1a, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xec, 0x01, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x70, 0x0a, 0x0e, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2a, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x66,
	0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x4d, 0x73, 0x67, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x0f, 0x52, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2b, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x6c, 0x6c,
	0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xee,
	0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x66, 0x65,
	0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x07,
	0x54, 0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x3b, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x46, 0x58, 0xaa, 0x02, 0x17, 0x43, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xca, 0x02, 0x17, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x46, 0x65, 0x65,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x23,
	0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x46, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5c,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x46, 0x65,
	0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_feegrant_v1beta1_tx_proto_rawDescOnce sync.Once
	file_cosmos_feegrant_v1beta1_tx_proto_rawDescData = file_cosmos_feegrant_v1beta1_tx_proto_rawDesc
)

func file_cosmos_feegrant_v1beta1_tx_proto_rawDescGZIP() []byte {
	file_cosmos_feegrant_v1beta1_tx_proto_rawDescOnce.Do(func() {
		file_cosmos_feegrant_v1beta1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_feegrant_v1beta1_tx_proto_rawDescData)
	})
	return file_cosmos_feegrant_v1beta1_tx_proto_rawDescData
}

var file_cosmos_feegrant_v1beta1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cosmos_feegrant_v1beta1_tx_proto_goTypes = []interface{}{
	(*MsgGrantAllowance)(nil),          // 0: cosmos.feegrant.v1beta1.MsgGrantAllowance
	(*MsgGrantAllowanceResponse)(nil),  // 1: cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse
	(*MsgRevokeAllowance)(nil),         // 2: cosmos.feegrant.v1beta1.MsgRevokeAllowance
	(*MsgRevokeAllowanceResponse)(nil), // 3: cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse
	(*anypb.Any)(nil),                  // 4: google.protobuf.Any
}
var file_cosmos_feegrant_v1beta1_tx_proto_depIdxs = []int32{
	4, // 0: cosmos.feegrant.v1beta1.MsgGrantAllowance.allowance:type_name -> google.protobuf.Any
	0, // 1: cosmos.feegrant.v1beta1.Msg.GrantAllowance:input_type -> cosmos.feegrant.v1beta1.MsgGrantAllowance
	2, // 2: cosmos.feegrant.v1beta1.Msg.RevokeAllowance:input_type -> cosmos.feegrant.v1beta1.MsgRevokeAllowance
	1, // 3: cosmos.feegrant.v1beta1.Msg.GrantAllowance:output_type -> cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse
	3, // 4: cosmos.feegrant.v1beta1.Msg.RevokeAllowance:output_type -> cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cosmos_feegrant_v1beta1_tx_proto_init() }
func file_cosmos_feegrant_v1beta1_tx_proto_init() {
	if File_cosmos_feegrant_v1beta1_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_feegrant_v1beta1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgGrantAllowance); i {
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
		file_cosmos_feegrant_v1beta1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgGrantAllowanceResponse); i {
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
		file_cosmos_feegrant_v1beta1_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRevokeAllowance); i {
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
		file_cosmos_feegrant_v1beta1_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRevokeAllowanceResponse); i {
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
			RawDescriptor: file_cosmos_feegrant_v1beta1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_feegrant_v1beta1_tx_proto_goTypes,
		DependencyIndexes: file_cosmos_feegrant_v1beta1_tx_proto_depIdxs,
		MessageInfos:      file_cosmos_feegrant_v1beta1_tx_proto_msgTypes,
	}.Build()
	File_cosmos_feegrant_v1beta1_tx_proto = out.File
	file_cosmos_feegrant_v1beta1_tx_proto_rawDesc = nil
	file_cosmos_feegrant_v1beta1_tx_proto_goTypes = nil
	file_cosmos_feegrant_v1beta1_tx_proto_depIdxs = nil
}
