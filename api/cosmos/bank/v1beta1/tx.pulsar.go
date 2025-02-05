package bankv1beta1

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	_ "github.com/gogo/protobuf/gogoproto"
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

var _ protoreflect.List = (*_MsgSend_3_list)(nil)

type _MsgSend_3_list struct {
	list *[]*v1beta1.Coin
}

func (x *_MsgSend_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgSend_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgSend_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_MsgSend_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgSend_3_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgSend_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgSend_3_list) NewElement() protoreflect.Value {
	v := new(v1beta1.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgSend_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgSend              protoreflect.MessageDescriptor
	fd_MsgSend_from_address protoreflect.FieldDescriptor
	fd_MsgSend_to_address   protoreflect.FieldDescriptor
	fd_MsgSend_amount       protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_bank_v1beta1_tx_proto_init()
	md_MsgSend = File_cosmos_bank_v1beta1_tx_proto.Messages().ByName("MsgSend")
	fd_MsgSend_from_address = md_MsgSend.Fields().ByName("from_address")
	fd_MsgSend_to_address = md_MsgSend.Fields().ByName("to_address")
	fd_MsgSend_amount = md_MsgSend.Fields().ByName("amount")
}

var _ protoreflect.Message = (*fastReflection_MsgSend)(nil)

type fastReflection_MsgSend MsgSend

func (x *MsgSend) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSend)(x)
}

func (x *MsgSend) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_v1beta1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSend_messageType fastReflection_MsgSend_messageType
var _ protoreflect.MessageType = fastReflection_MsgSend_messageType{}

type fastReflection_MsgSend_messageType struct{}

func (x fastReflection_MsgSend_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSend)(nil)
}
func (x fastReflection_MsgSend_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSend)
}
func (x fastReflection_MsgSend_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSend
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSend) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSend
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSend) Type() protoreflect.MessageType {
	return _fastReflection_MsgSend_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSend) New() protoreflect.Message {
	return new(fastReflection_MsgSend)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSend) Interface() protoreflect.ProtoMessage {
	return (*MsgSend)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSend) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.FromAddress != "" {
		value := protoreflect.ValueOfString(x.FromAddress)
		if !f(fd_MsgSend_from_address, value) {
			return
		}
	}
	if x.ToAddress != "" {
		value := protoreflect.ValueOfString(x.ToAddress)
		if !f(fd_MsgSend_to_address, value) {
			return
		}
	}
	if len(x.Amount) != 0 {
		value := protoreflect.ValueOfList(&_MsgSend_3_list{list: &x.Amount})
		if !f(fd_MsgSend_amount, value) {
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
func (x *fastReflection_MsgSend) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.bank.v1beta1.MsgSend.from_address":
		return x.FromAddress != ""
	case "cosmos.bank.v1beta1.MsgSend.to_address":
		return x.ToAddress != ""
	case "cosmos.bank.v1beta1.MsgSend.amount":
		return len(x.Amount) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgSend"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgSend does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSend) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.bank.v1beta1.MsgSend.from_address":
		x.FromAddress = ""
	case "cosmos.bank.v1beta1.MsgSend.to_address":
		x.ToAddress = ""
	case "cosmos.bank.v1beta1.MsgSend.amount":
		x.Amount = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgSend"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgSend does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSend) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.bank.v1beta1.MsgSend.from_address":
		value := x.FromAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.bank.v1beta1.MsgSend.to_address":
		value := x.ToAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.bank.v1beta1.MsgSend.amount":
		if len(x.Amount) == 0 {
			return protoreflect.ValueOfList(&_MsgSend_3_list{})
		}
		listValue := &_MsgSend_3_list{list: &x.Amount}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgSend"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgSend does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSend) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.bank.v1beta1.MsgSend.from_address":
		x.FromAddress = value.Interface().(string)
	case "cosmos.bank.v1beta1.MsgSend.to_address":
		x.ToAddress = value.Interface().(string)
	case "cosmos.bank.v1beta1.MsgSend.amount":
		lv := value.List()
		clv := lv.(*_MsgSend_3_list)
		x.Amount = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgSend"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgSend does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSend) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.bank.v1beta1.MsgSend.amount":
		if x.Amount == nil {
			x.Amount = []*v1beta1.Coin{}
		}
		value := &_MsgSend_3_list{list: &x.Amount}
		return protoreflect.ValueOfList(value)
	case "cosmos.bank.v1beta1.MsgSend.from_address":
		panic(fmt.Errorf("field from_address of message cosmos.bank.v1beta1.MsgSend is not mutable"))
	case "cosmos.bank.v1beta1.MsgSend.to_address":
		panic(fmt.Errorf("field to_address of message cosmos.bank.v1beta1.MsgSend is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgSend"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgSend does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSend) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.bank.v1beta1.MsgSend.from_address":
		return protoreflect.ValueOfString("")
	case "cosmos.bank.v1beta1.MsgSend.to_address":
		return protoreflect.ValueOfString("")
	case "cosmos.bank.v1beta1.MsgSend.amount":
		list := []*v1beta1.Coin{}
		return protoreflect.ValueOfList(&_MsgSend_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgSend"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgSend does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSend) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.bank.v1beta1.MsgSend", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSend) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSend) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSend) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSend) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSend)
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
		l = len(x.FromAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ToAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Amount) > 0 {
			for _, e := range x.Amount {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*MsgSend)
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
		if len(x.Amount) > 0 {
			for iNdEx := len(x.Amount) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Amount[iNdEx])
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
		}
		if len(x.ToAddress) > 0 {
			i -= len(x.ToAddress)
			copy(dAtA[i:], x.ToAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ToAddress)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.FromAddress) > 0 {
			i -= len(x.FromAddress)
			copy(dAtA[i:], x.FromAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.FromAddress)))
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
		x := input.Message.Interface().(*MsgSend)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSend: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSend: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
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
				x.FromAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
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
				x.ToAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
				x.Amount = append(x.Amount, &v1beta1.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Amount[len(x.Amount)-1]); err != nil {
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
	md_MsgSendResponse protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_bank_v1beta1_tx_proto_init()
	md_MsgSendResponse = File_cosmos_bank_v1beta1_tx_proto.Messages().ByName("MsgSendResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgSendResponse)(nil)

type fastReflection_MsgSendResponse MsgSendResponse

func (x *MsgSendResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSendResponse)(x)
}

func (x *MsgSendResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_v1beta1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSendResponse_messageType fastReflection_MsgSendResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgSendResponse_messageType{}

type fastReflection_MsgSendResponse_messageType struct{}

func (x fastReflection_MsgSendResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSendResponse)(nil)
}
func (x fastReflection_MsgSendResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSendResponse)
}
func (x fastReflection_MsgSendResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSendResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSendResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSendResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSendResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgSendResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSendResponse) New() protoreflect.Message {
	return new(fastReflection_MsgSendResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSendResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgSendResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSendResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgSendResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgSendResponse"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgSendResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSendResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgSendResponse"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgSendResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSendResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgSendResponse"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgSendResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSendResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgSendResponse"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgSendResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSendResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgSendResponse"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgSendResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSendResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgSendResponse"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgSendResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSendResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.bank.v1beta1.MsgSendResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSendResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSendResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSendResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSendResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSendResponse)
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
		x := input.Message.Interface().(*MsgSendResponse)
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
		x := input.Message.Interface().(*MsgSendResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSendResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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

var _ protoreflect.List = (*_MsgMultiSend_1_list)(nil)

type _MsgMultiSend_1_list struct {
	list *[]*Input
}

func (x *_MsgMultiSend_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgMultiSend_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgMultiSend_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Input)
	(*x.list)[i] = concreteValue
}

func (x *_MsgMultiSend_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Input)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgMultiSend_1_list) AppendMutable() protoreflect.Value {
	v := new(Input)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgMultiSend_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgMultiSend_1_list) NewElement() protoreflect.Value {
	v := new(Input)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgMultiSend_1_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_MsgMultiSend_2_list)(nil)

type _MsgMultiSend_2_list struct {
	list *[]*Output
}

func (x *_MsgMultiSend_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgMultiSend_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgMultiSend_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Output)
	(*x.list)[i] = concreteValue
}

func (x *_MsgMultiSend_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Output)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgMultiSend_2_list) AppendMutable() protoreflect.Value {
	v := new(Output)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgMultiSend_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgMultiSend_2_list) NewElement() protoreflect.Value {
	v := new(Output)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgMultiSend_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgMultiSend         protoreflect.MessageDescriptor
	fd_MsgMultiSend_inputs  protoreflect.FieldDescriptor
	fd_MsgMultiSend_outputs protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_bank_v1beta1_tx_proto_init()
	md_MsgMultiSend = File_cosmos_bank_v1beta1_tx_proto.Messages().ByName("MsgMultiSend")
	fd_MsgMultiSend_inputs = md_MsgMultiSend.Fields().ByName("inputs")
	fd_MsgMultiSend_outputs = md_MsgMultiSend.Fields().ByName("outputs")
}

var _ protoreflect.Message = (*fastReflection_MsgMultiSend)(nil)

type fastReflection_MsgMultiSend MsgMultiSend

func (x *MsgMultiSend) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgMultiSend)(x)
}

func (x *MsgMultiSend) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_v1beta1_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgMultiSend_messageType fastReflection_MsgMultiSend_messageType
var _ protoreflect.MessageType = fastReflection_MsgMultiSend_messageType{}

type fastReflection_MsgMultiSend_messageType struct{}

func (x fastReflection_MsgMultiSend_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgMultiSend)(nil)
}
func (x fastReflection_MsgMultiSend_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgMultiSend)
}
func (x fastReflection_MsgMultiSend_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgMultiSend
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgMultiSend) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgMultiSend
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgMultiSend) Type() protoreflect.MessageType {
	return _fastReflection_MsgMultiSend_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgMultiSend) New() protoreflect.Message {
	return new(fastReflection_MsgMultiSend)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgMultiSend) Interface() protoreflect.ProtoMessage {
	return (*MsgMultiSend)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgMultiSend) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Inputs) != 0 {
		value := protoreflect.ValueOfList(&_MsgMultiSend_1_list{list: &x.Inputs})
		if !f(fd_MsgMultiSend_inputs, value) {
			return
		}
	}
	if len(x.Outputs) != 0 {
		value := protoreflect.ValueOfList(&_MsgMultiSend_2_list{list: &x.Outputs})
		if !f(fd_MsgMultiSend_outputs, value) {
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
func (x *fastReflection_MsgMultiSend) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.bank.v1beta1.MsgMultiSend.inputs":
		return len(x.Inputs) != 0
	case "cosmos.bank.v1beta1.MsgMultiSend.outputs":
		return len(x.Outputs) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgMultiSend"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgMultiSend does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgMultiSend) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.bank.v1beta1.MsgMultiSend.inputs":
		x.Inputs = nil
	case "cosmos.bank.v1beta1.MsgMultiSend.outputs":
		x.Outputs = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgMultiSend"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgMultiSend does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgMultiSend) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.bank.v1beta1.MsgMultiSend.inputs":
		if len(x.Inputs) == 0 {
			return protoreflect.ValueOfList(&_MsgMultiSend_1_list{})
		}
		listValue := &_MsgMultiSend_1_list{list: &x.Inputs}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.bank.v1beta1.MsgMultiSend.outputs":
		if len(x.Outputs) == 0 {
			return protoreflect.ValueOfList(&_MsgMultiSend_2_list{})
		}
		listValue := &_MsgMultiSend_2_list{list: &x.Outputs}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgMultiSend"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgMultiSend does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgMultiSend) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.bank.v1beta1.MsgMultiSend.inputs":
		lv := value.List()
		clv := lv.(*_MsgMultiSend_1_list)
		x.Inputs = *clv.list
	case "cosmos.bank.v1beta1.MsgMultiSend.outputs":
		lv := value.List()
		clv := lv.(*_MsgMultiSend_2_list)
		x.Outputs = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgMultiSend"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgMultiSend does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgMultiSend) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.bank.v1beta1.MsgMultiSend.inputs":
		if x.Inputs == nil {
			x.Inputs = []*Input{}
		}
		value := &_MsgMultiSend_1_list{list: &x.Inputs}
		return protoreflect.ValueOfList(value)
	case "cosmos.bank.v1beta1.MsgMultiSend.outputs":
		if x.Outputs == nil {
			x.Outputs = []*Output{}
		}
		value := &_MsgMultiSend_2_list{list: &x.Outputs}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgMultiSend"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgMultiSend does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgMultiSend) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.bank.v1beta1.MsgMultiSend.inputs":
		list := []*Input{}
		return protoreflect.ValueOfList(&_MsgMultiSend_1_list{list: &list})
	case "cosmos.bank.v1beta1.MsgMultiSend.outputs":
		list := []*Output{}
		return protoreflect.ValueOfList(&_MsgMultiSend_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgMultiSend"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgMultiSend does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgMultiSend) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.bank.v1beta1.MsgMultiSend", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgMultiSend) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgMultiSend) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgMultiSend) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgMultiSend) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgMultiSend)
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
		if len(x.Inputs) > 0 {
			for _, e := range x.Inputs {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.Outputs) > 0 {
			for _, e := range x.Outputs {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*MsgMultiSend)
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
		if len(x.Outputs) > 0 {
			for iNdEx := len(x.Outputs) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Outputs[iNdEx])
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
		if len(x.Inputs) > 0 {
			for iNdEx := len(x.Inputs) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Inputs[iNdEx])
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
		x := input.Message.Interface().(*MsgMultiSend)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgMultiSend: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgMultiSend: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
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
				x.Inputs = append(x.Inputs, &Input{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Inputs[len(x.Inputs)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Outputs", wireType)
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
				x.Outputs = append(x.Outputs, &Output{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Outputs[len(x.Outputs)-1]); err != nil {
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
	md_MsgMultiSendResponse protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_bank_v1beta1_tx_proto_init()
	md_MsgMultiSendResponse = File_cosmos_bank_v1beta1_tx_proto.Messages().ByName("MsgMultiSendResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgMultiSendResponse)(nil)

type fastReflection_MsgMultiSendResponse MsgMultiSendResponse

func (x *MsgMultiSendResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgMultiSendResponse)(x)
}

func (x *MsgMultiSendResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_v1beta1_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgMultiSendResponse_messageType fastReflection_MsgMultiSendResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgMultiSendResponse_messageType{}

type fastReflection_MsgMultiSendResponse_messageType struct{}

func (x fastReflection_MsgMultiSendResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgMultiSendResponse)(nil)
}
func (x fastReflection_MsgMultiSendResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgMultiSendResponse)
}
func (x fastReflection_MsgMultiSendResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgMultiSendResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgMultiSendResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgMultiSendResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgMultiSendResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgMultiSendResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgMultiSendResponse) New() protoreflect.Message {
	return new(fastReflection_MsgMultiSendResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgMultiSendResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgMultiSendResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgMultiSendResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgMultiSendResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgMultiSendResponse"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgMultiSendResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgMultiSendResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgMultiSendResponse"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgMultiSendResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgMultiSendResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgMultiSendResponse"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgMultiSendResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgMultiSendResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgMultiSendResponse"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgMultiSendResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgMultiSendResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgMultiSendResponse"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgMultiSendResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgMultiSendResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.MsgMultiSendResponse"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.MsgMultiSendResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgMultiSendResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.bank.v1beta1.MsgMultiSendResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgMultiSendResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgMultiSendResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgMultiSendResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgMultiSendResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgMultiSendResponse)
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
		x := input.Message.Interface().(*MsgMultiSendResponse)
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
		x := input.Message.Interface().(*MsgMultiSendResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgMultiSendResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgMultiSendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
	// Send defines a method for sending coins from one account to another account.
	Send(ctx context.Context, in *MsgSend, opts ...grpc.CallOption) (*MsgSendResponse, error)
	// MultiSend defines a method for sending coins from some accounts to other accounts.
	MultiSend(ctx context.Context, in *MsgMultiSend, opts ...grpc.CallOption) (*MsgMultiSendResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Send(ctx context.Context, in *MsgSend, opts ...grpc.CallOption) (*MsgSendResponse, error) {
	out := new(MsgSendResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Msg/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MultiSend(ctx context.Context, in *MsgMultiSend, opts ...grpc.CallOption) (*MsgMultiSendResponse, error) {
	out := new(MsgMultiSendResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Msg/MultiSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// Send defines a method for sending coins from one account to another account.
	Send(context.Context, *MsgSend) (*MsgSendResponse, error)
	// MultiSend defines a method for sending coins from some accounts to other accounts.
	MultiSend(context.Context, *MsgMultiSend) (*MsgMultiSendResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) Send(context.Context, *MsgSend) (*MsgSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedMsgServer) MultiSend(context.Context, *MsgMultiSend) (*MsgMultiSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiSend not implemented")
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

func _Msg_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Msg/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Send(ctx, req.(*MsgSend))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MultiSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMultiSend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MultiSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Msg/MultiSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MultiSend(ctx, req.(*MsgMultiSend))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.bank.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Msg_Send_Handler,
		},
		{
			MethodName: "MultiSend",
			Handler:    _Msg_MultiSend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/bank/v1beta1/tx.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/bank/v1beta1/tx.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MsgSend represents a message to send coins from one account to another.
type MsgSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAddress string          `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ToAddress   string          `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount      []*v1beta1.Coin `protobuf:"bytes,3,rep,name=amount,proto3" json:"amount,omitempty"`
}

func (x *MsgSend) Reset() {
	*x = MsgSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_v1beta1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSend) ProtoMessage() {}

// Deprecated: Use MsgSend.ProtoReflect.Descriptor instead.
func (*MsgSend) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_v1beta1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgSend) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *MsgSend) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *MsgSend) GetAmount() []*v1beta1.Coin {
	if x != nil {
		return x.Amount
	}
	return nil
}

// MsgSendResponse defines the Msg/Send response type.
type MsgSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgSendResponse) Reset() {
	*x = MsgSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_v1beta1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSendResponse) ProtoMessage() {}

// Deprecated: Use MsgSendResponse.ProtoReflect.Descriptor instead.
func (*MsgSendResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_v1beta1_tx_proto_rawDescGZIP(), []int{1}
}

// MsgMultiSend represents an arbitrary multi-in, multi-out send message.
type MsgMultiSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inputs  []*Input  `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs []*Output `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *MsgMultiSend) Reset() {
	*x = MsgMultiSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_v1beta1_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgMultiSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgMultiSend) ProtoMessage() {}

// Deprecated: Use MsgMultiSend.ProtoReflect.Descriptor instead.
func (*MsgMultiSend) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_v1beta1_tx_proto_rawDescGZIP(), []int{2}
}

func (x *MsgMultiSend) GetInputs() []*Input {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *MsgMultiSend) GetOutputs() []*Output {
	if x != nil {
		return x.Outputs
	}
	return nil
}

// MsgMultiSendResponse defines the Msg/MultiSend response type.
type MsgMultiSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgMultiSendResponse) Reset() {
	*x = MsgMultiSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_v1beta1_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgMultiSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgMultiSendResponse) ProtoMessage() {}

// Deprecated: Use MsgMultiSendResponse.ProtoReflect.Descriptor instead.
func (*MsgMultiSendResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_v1beta1_tx_proto_rawDescGZIP(), []int{3}
}

var File_cosmos_bank_v1beta1_tx_proto protoreflect.FileDescriptor

var file_cosmos_bank_v1beta1_tx_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63,
	0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64,
	0x12, 0x3b, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x37, 0x0a,
	0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x74, 0x6f, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x63, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69,
	0x6e, 0x42, 0x30, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x69, 0x6e, 0x73, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x08, 0x88, 0xa0, 0x1f,
	0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x11, 0x0a, 0x0f, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x0c, 0x4d, 0x73, 0x67,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x3a, 0x04, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x16, 0x0a, 0x14, 0x4d, 0x73, 0x67, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xac,
	0x01, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x4a, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x1c,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x1a, 0x24, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x59, 0x0a, 0x09, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x65, 0x6e, 0x64, 0x12,
	0x21, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x65,
	0x6e, 0x64, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xd2, 0x01,
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x07, 0x54, 0x78, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61,
	0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x62, 0x61, 0x6e, 0x6b, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x42, 0x58, 0xaa, 0x02, 0x13, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xca, 0x02, 0x13, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x42, 0x61, 0x6e, 0x6b,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1f, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5c, 0x42, 0x61, 0x6e, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x43, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x42, 0x61, 0x6e, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_bank_v1beta1_tx_proto_rawDescOnce sync.Once
	file_cosmos_bank_v1beta1_tx_proto_rawDescData = file_cosmos_bank_v1beta1_tx_proto_rawDesc
)

func file_cosmos_bank_v1beta1_tx_proto_rawDescGZIP() []byte {
	file_cosmos_bank_v1beta1_tx_proto_rawDescOnce.Do(func() {
		file_cosmos_bank_v1beta1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_bank_v1beta1_tx_proto_rawDescData)
	})
	return file_cosmos_bank_v1beta1_tx_proto_rawDescData
}

var file_cosmos_bank_v1beta1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cosmos_bank_v1beta1_tx_proto_goTypes = []interface{}{
	(*MsgSend)(nil),              // 0: cosmos.bank.v1beta1.MsgSend
	(*MsgSendResponse)(nil),      // 1: cosmos.bank.v1beta1.MsgSendResponse
	(*MsgMultiSend)(nil),         // 2: cosmos.bank.v1beta1.MsgMultiSend
	(*MsgMultiSendResponse)(nil), // 3: cosmos.bank.v1beta1.MsgMultiSendResponse
	(*v1beta1.Coin)(nil),         // 4: cosmos.base.v1beta1.Coin
	(*Input)(nil),                // 5: cosmos.bank.v1beta1.Input
	(*Output)(nil),               // 6: cosmos.bank.v1beta1.Output
}
var file_cosmos_bank_v1beta1_tx_proto_depIdxs = []int32{
	4, // 0: cosmos.bank.v1beta1.MsgSend.amount:type_name -> cosmos.base.v1beta1.Coin
	5, // 1: cosmos.bank.v1beta1.MsgMultiSend.inputs:type_name -> cosmos.bank.v1beta1.Input
	6, // 2: cosmos.bank.v1beta1.MsgMultiSend.outputs:type_name -> cosmos.bank.v1beta1.Output
	0, // 3: cosmos.bank.v1beta1.Msg.Send:input_type -> cosmos.bank.v1beta1.MsgSend
	2, // 4: cosmos.bank.v1beta1.Msg.MultiSend:input_type -> cosmos.bank.v1beta1.MsgMultiSend
	1, // 5: cosmos.bank.v1beta1.Msg.Send:output_type -> cosmos.bank.v1beta1.MsgSendResponse
	3, // 6: cosmos.bank.v1beta1.Msg.MultiSend:output_type -> cosmos.bank.v1beta1.MsgMultiSendResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cosmos_bank_v1beta1_tx_proto_init() }
func file_cosmos_bank_v1beta1_tx_proto_init() {
	if File_cosmos_bank_v1beta1_tx_proto != nil {
		return
	}
	file_cosmos_bank_v1beta1_bank_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_bank_v1beta1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSend); i {
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
		file_cosmos_bank_v1beta1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSendResponse); i {
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
		file_cosmos_bank_v1beta1_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgMultiSend); i {
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
		file_cosmos_bank_v1beta1_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgMultiSendResponse); i {
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
			RawDescriptor: file_cosmos_bank_v1beta1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_bank_v1beta1_tx_proto_goTypes,
		DependencyIndexes: file_cosmos_bank_v1beta1_tx_proto_depIdxs,
		MessageInfos:      file_cosmos_bank_v1beta1_tx_proto_msgTypes,
	}.Build()
	File_cosmos_bank_v1beta1_tx_proto = out.File
	file_cosmos_bank_v1beta1_tx_proto_rawDesc = nil
	file_cosmos_bank_v1beta1_tx_proto_goTypes = nil
	file_cosmos_bank_v1beta1_tx_proto_depIdxs = nil
}
