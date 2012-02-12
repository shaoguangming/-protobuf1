// Code generated by protoc-gen-go from "google/protobuf/descriptor.proto"
// DO NOT EDIT!

package google_protobuf

import proto "code.google.com/p/goprotobuf/proto"
import "math"

// Reference proto, math & os imports to suppress error if they are not otherwise used.
var _ = proto.GetString
var _ = math.Inf

type FieldDescriptorProto_Type int32

const (
	FieldDescriptorProto_TYPE_DOUBLE   = 1
	FieldDescriptorProto_TYPE_FLOAT    = 2
	FieldDescriptorProto_TYPE_INT64    = 3
	FieldDescriptorProto_TYPE_UINT64   = 4
	FieldDescriptorProto_TYPE_INT32    = 5
	FieldDescriptorProto_TYPE_FIXED64  = 6
	FieldDescriptorProto_TYPE_FIXED32  = 7
	FieldDescriptorProto_TYPE_BOOL     = 8
	FieldDescriptorProto_TYPE_STRING   = 9
	FieldDescriptorProto_TYPE_GROUP    = 10
	FieldDescriptorProto_TYPE_MESSAGE  = 11
	FieldDescriptorProto_TYPE_BYTES    = 12
	FieldDescriptorProto_TYPE_UINT32   = 13
	FieldDescriptorProto_TYPE_ENUM     = 14
	FieldDescriptorProto_TYPE_SFIXED32 = 15
	FieldDescriptorProto_TYPE_SFIXED64 = 16
	FieldDescriptorProto_TYPE_SINT32   = 17
	FieldDescriptorProto_TYPE_SINT64   = 18
)

var FieldDescriptorProto_Type_name = map[int32]string{
	1:  "TYPE_DOUBLE",
	2:  "TYPE_FLOAT",
	3:  "TYPE_INT64",
	4:  "TYPE_UINT64",
	5:  "TYPE_INT32",
	6:  "TYPE_FIXED64",
	7:  "TYPE_FIXED32",
	8:  "TYPE_BOOL",
	9:  "TYPE_STRING",
	10: "TYPE_GROUP",
	11: "TYPE_MESSAGE",
	12: "TYPE_BYTES",
	13: "TYPE_UINT32",
	14: "TYPE_ENUM",
	15: "TYPE_SFIXED32",
	16: "TYPE_SFIXED64",
	17: "TYPE_SINT32",
	18: "TYPE_SINT64",
}
var FieldDescriptorProto_Type_value = map[string]int32{
	"TYPE_DOUBLE":   1,
	"TYPE_FLOAT":    2,
	"TYPE_INT64":    3,
	"TYPE_UINT64":   4,
	"TYPE_INT32":    5,
	"TYPE_FIXED64":  6,
	"TYPE_FIXED32":  7,
	"TYPE_BOOL":     8,
	"TYPE_STRING":   9,
	"TYPE_GROUP":    10,
	"TYPE_MESSAGE":  11,
	"TYPE_BYTES":    12,
	"TYPE_UINT32":   13,
	"TYPE_ENUM":     14,
	"TYPE_SFIXED32": 15,
	"TYPE_SFIXED64": 16,
	"TYPE_SINT32":   17,
	"TYPE_SINT64":   18,
}

func NewFieldDescriptorProto_Type(x int32) *FieldDescriptorProto_Type {
	e := FieldDescriptorProto_Type(x)
	return &e
}
func (x FieldDescriptorProto_Type) String() string {
	return proto.EnumName(FieldDescriptorProto_Type_name, int32(x))
}

type FieldDescriptorProto_Label int32

const (
	FieldDescriptorProto_LABEL_OPTIONAL = 1
	FieldDescriptorProto_LABEL_REQUIRED = 2
	FieldDescriptorProto_LABEL_REPEATED = 3
)

var FieldDescriptorProto_Label_name = map[int32]string{
	1: "LABEL_OPTIONAL",
	2: "LABEL_REQUIRED",
	3: "LABEL_REPEATED",
}
var FieldDescriptorProto_Label_value = map[string]int32{
	"LABEL_OPTIONAL": 1,
	"LABEL_REQUIRED": 2,
	"LABEL_REPEATED": 3,
}

func NewFieldDescriptorProto_Label(x int32) *FieldDescriptorProto_Label {
	e := FieldDescriptorProto_Label(x)
	return &e
}
func (x FieldDescriptorProto_Label) String() string {
	return proto.EnumName(FieldDescriptorProto_Label_name, int32(x))
}

type FileOptions_OptimizeMode int32

const (
	FileOptions_SPEED        = 1
	FileOptions_CODE_SIZE    = 2
	FileOptions_LITE_RUNTIME = 3
)

var FileOptions_OptimizeMode_name = map[int32]string{
	1: "SPEED",
	2: "CODE_SIZE",
	3: "LITE_RUNTIME",
}
var FileOptions_OptimizeMode_value = map[string]int32{
	"SPEED":        1,
	"CODE_SIZE":    2,
	"LITE_RUNTIME": 3,
}

func NewFileOptions_OptimizeMode(x int32) *FileOptions_OptimizeMode {
	e := FileOptions_OptimizeMode(x)
	return &e
}
func (x FileOptions_OptimizeMode) String() string {
	return proto.EnumName(FileOptions_OptimizeMode_name, int32(x))
}

type FieldOptions_CType int32

const (
	FieldOptions_STRING       = 0
	FieldOptions_CORD         = 1
	FieldOptions_STRING_PIECE = 2
)

var FieldOptions_CType_name = map[int32]string{
	0: "STRING",
	1: "CORD",
	2: "STRING_PIECE",
}
var FieldOptions_CType_value = map[string]int32{
	"STRING":       0,
	"CORD":         1,
	"STRING_PIECE": 2,
}

func NewFieldOptions_CType(x int32) *FieldOptions_CType {
	e := FieldOptions_CType(x)
	return &e
}
func (x FieldOptions_CType) String() string {
	return proto.EnumName(FieldOptions_CType_name, int32(x))
}

type FileDescriptorSet struct {
	File             []*FileDescriptorProto `protobuf:"bytes,1,rep,name=file"`
	XXX_unrecognized []byte
}

func (this *FileDescriptorSet) Reset()         { *this = FileDescriptorSet{} }
func (this *FileDescriptorSet) String() string { return proto.CompactTextString(this) }

type FileDescriptorProto struct {
	Name             *string                   `protobuf:"bytes,1,opt,name=name"`
	Package          *string                   `protobuf:"bytes,2,opt,name=package"`
	Dependency       []string                  `protobuf:"bytes,3,rep,name=dependency"`
	PublicDependency []int32                   `protobuf:"varint,10,rep,name=public_dependency"`
	WeakDependency   []int32                   `protobuf:"varint,11,rep,name=weak_dependency" json:"weak_dependency,omitempty"`
	MessageType      []*DescriptorProto        `protobuf:"bytes,4,rep,name=message_type"`
	EnumType         []*EnumDescriptorProto    `protobuf:"bytes,5,rep,name=enum_type"`
	Service          []*ServiceDescriptorProto `protobuf:"bytes,6,rep,name=service"`
	Extension        []*FieldDescriptorProto   `protobuf:"bytes,7,rep,name=extension"`
	Options          *FileOptions              `protobuf:"bytes,8,opt,name=options"`
	XXX_unrecognized []byte
}

func (this *FileDescriptorProto) Reset()         { *this = FileDescriptorProto{} }
func (this *FileDescriptorProto) String() string { return proto.CompactTextString(this) }

type DescriptorProto struct {
	Name             *string                           `protobuf:"bytes,1,opt,name=name"`
	Field            []*FieldDescriptorProto           `protobuf:"bytes,2,rep,name=field"`
	Extension        []*FieldDescriptorProto           `protobuf:"bytes,6,rep,name=extension"`
	NestedType       []*DescriptorProto                `protobuf:"bytes,3,rep,name=nested_type"`
	EnumType         []*EnumDescriptorProto            `protobuf:"bytes,4,rep,name=enum_type"`
	ExtensionRange   []*DescriptorProto_ExtensionRange `protobuf:"bytes,5,rep,name=extension_range"`
	Options          *MessageOptions                   `protobuf:"bytes,7,opt,name=options"`
	XXX_unrecognized []byte
}

func (this *DescriptorProto) Reset()         { *this = DescriptorProto{} }
func (this *DescriptorProto) String() string { return proto.CompactTextString(this) }

type DescriptorProto_ExtensionRange struct {
	Start            *int32 `protobuf:"varint,1,opt,name=start"`
	End              *int32 `protobuf:"varint,2,opt,name=end"`
	XXX_unrecognized []byte
}

func (this *DescriptorProto_ExtensionRange) Reset()         { *this = DescriptorProto_ExtensionRange{} }
func (this *DescriptorProto_ExtensionRange) String() string { return proto.CompactTextString(this) }

type FieldDescriptorProto struct {
	Name             *string                     `protobuf:"bytes,1,opt,name=name"`
	Number           *int32                      `protobuf:"varint,3,opt,name=number"`
	Label            *FieldDescriptorProto_Label `protobuf:"varint,4,opt,name=label,enum=google_protobuf.FieldDescriptorProto_Label"`
	Type             *FieldDescriptorProto_Type  `protobuf:"varint,5,opt,name=type,enum=google_protobuf.FieldDescriptorProto_Type"`
	TypeName         *string                     `protobuf:"bytes,6,opt,name=type_name"`
	Extendee         *string                     `protobuf:"bytes,2,opt,name=extendee"`
	DefaultValue     *string                     `protobuf:"bytes,7,opt,name=default_value"`
	Options          *FieldOptions               `protobuf:"bytes,8,opt,name=options"`
	XXX_unrecognized []byte
}

func (this *FieldDescriptorProto) Reset()         { *this = FieldDescriptorProto{} }
func (this *FieldDescriptorProto) String() string { return proto.CompactTextString(this) }

type EnumDescriptorProto struct {
	Name             *string                     `protobuf:"bytes,1,opt,name=name"`
	Value            []*EnumValueDescriptorProto `protobuf:"bytes,2,rep,name=value"`
	Options          *EnumOptions                `protobuf:"bytes,3,opt,name=options"`
	XXX_unrecognized []byte
}

func (this *EnumDescriptorProto) Reset()         { *this = EnumDescriptorProto{} }
func (this *EnumDescriptorProto) String() string { return proto.CompactTextString(this) }

type EnumValueDescriptorProto struct {
	Name             *string           `protobuf:"bytes,1,opt,name=name"`
	Number           *int32            `protobuf:"varint,2,opt,name=number"`
	Options          *EnumValueOptions `protobuf:"bytes,3,opt,name=options"`
	XXX_unrecognized []byte
}

func (this *EnumValueDescriptorProto) Reset()         { *this = EnumValueDescriptorProto{} }
func (this *EnumValueDescriptorProto) String() string { return proto.CompactTextString(this) }

type ServiceDescriptorProto struct {
	Name             *string                  `protobuf:"bytes,1,opt,name=name"`
	Method           []*MethodDescriptorProto `protobuf:"bytes,2,rep,name=method"`
	Options          *ServiceOptions          `protobuf:"bytes,3,opt,name=options"`
	XXX_unrecognized []byte
}

func (this *ServiceDescriptorProto) Reset()         { *this = ServiceDescriptorProto{} }
func (this *ServiceDescriptorProto) String() string { return proto.CompactTextString(this) }

type MethodDescriptorProto struct {
	Name             *string        `protobuf:"bytes,1,opt,name=name"`
	InputType        *string        `protobuf:"bytes,2,opt,name=input_type"`
	OutputType       *string        `protobuf:"bytes,3,opt,name=output_type"`
	Options          *MethodOptions `protobuf:"bytes,4,opt,name=options"`
	XXX_unrecognized []byte
}

func (this *MethodDescriptorProto) Reset()         { *this = MethodDescriptorProto{} }
func (this *MethodDescriptorProto) String() string { return proto.CompactTextString(this) }

type FileOptions struct {
	JavaPackage         *string                   `protobuf:"bytes,1,opt,name=java_package"`
	JavaOuterClassname  *string                   `protobuf:"bytes,8,opt,name=java_outer_classname"`
	JavaMultipleFiles   *bool                     `protobuf:"varint,10,opt,name=java_multiple_files,def=0"`
	OptimizeFor         *FileOptions_OptimizeMode `protobuf:"varint,9,opt,name=optimize_for,enum=google_protobuf.FileOptions_OptimizeMode,def=1"`
	CcGenericServices   *bool                     `protobuf:"varint,16,opt,name=cc_generic_services,def=1"`
	JavaGenericServices *bool                     `protobuf:"varint,17,opt,name=java_generic_services,def=1"`
	PyGenericServices   *bool                     `protobuf:"varint,18,opt,name=py_generic_services,def=1"`
	UninterpretedOption []*UninterpretedOption    `protobuf:"bytes,999,rep,name=uninterpreted_option"`
	XXX_extensions      map[int32][]byte
	XXX_unrecognized    []byte
}

func (this *FileOptions) Reset()         { *this = FileOptions{} }
func (this *FileOptions) String() string { return proto.CompactTextString(this) }

var extRange_FileOptions = []proto.ExtensionRange{
	proto.ExtensionRange{1000, 536870911},
}

func (*FileOptions) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_FileOptions
}
func (this *FileOptions) ExtensionMap() map[int32][]byte {
	if this.XXX_extensions == nil {
		this.XXX_extensions = make(map[int32][]byte)
	}
	return this.XXX_extensions
}

const Default_FileOptions_JavaMultipleFiles bool = false
const Default_FileOptions_OptimizeFor FileOptions_OptimizeMode = FileOptions_SPEED
const Default_FileOptions_CcGenericServices bool = true
const Default_FileOptions_JavaGenericServices bool = true
const Default_FileOptions_PyGenericServices bool = true

type MessageOptions struct {
	MessageSetWireFormat         *bool                  `protobuf:"varint,1,opt,name=message_set_wire_format,def=0"`
	NoStandardDescriptorAccessor *bool                  `protobuf:"varint,2,opt,name=no_standard_descriptor_accessor,def=0"`
	UninterpretedOption          []*UninterpretedOption `protobuf:"bytes,999,rep,name=uninterpreted_option"`
	XXX_extensions               map[int32][]byte
	XXX_unrecognized             []byte
}

func (this *MessageOptions) Reset()         { *this = MessageOptions{} }
func (this *MessageOptions) String() string { return proto.CompactTextString(this) }

var extRange_MessageOptions = []proto.ExtensionRange{
	proto.ExtensionRange{1000, 536870911},
}

func (*MessageOptions) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_MessageOptions
}
func (this *MessageOptions) ExtensionMap() map[int32][]byte {
	if this.XXX_extensions == nil {
		this.XXX_extensions = make(map[int32][]byte)
	}
	return this.XXX_extensions
}

const Default_MessageOptions_MessageSetWireFormat bool = false
const Default_MessageOptions_NoStandardDescriptorAccessor bool = false

type FieldOptions struct {
	Ctype               *FieldOptions_CType    `protobuf:"varint,1,opt,name=ctype,enum=google_protobuf.FieldOptions_CType,def=0"`
	Packed              *bool                  `protobuf:"varint,2,opt,name=packed"`
	Deprecated          *bool                  `protobuf:"varint,3,opt,name=deprecated,def=0"`
	ExperimentalMapKey  *string                `protobuf:"bytes,9,opt,name=experimental_map_key"`
	Weak                *bool                  `protobuf:"varint,10,opt,name=weak,def=0" json:"weak,omitempty"`
	UninterpretedOption []*UninterpretedOption `protobuf:"bytes,999,rep,name=uninterpreted_option"`
	XXX_extensions      map[int32][]byte
	XXX_unrecognized    []byte
}

func (this *FieldOptions) Reset()         { *this = FieldOptions{} }
func (this *FieldOptions) String() string { return proto.CompactTextString(this) }

var extRange_FieldOptions = []proto.ExtensionRange{
	proto.ExtensionRange{1000, 536870911},
}

func (*FieldOptions) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_FieldOptions
}
func (this *FieldOptions) ExtensionMap() map[int32][]byte {
	if this.XXX_extensions == nil {
		this.XXX_extensions = make(map[int32][]byte)
	}
	return this.XXX_extensions
}

const Default_FieldOptions_Ctype FieldOptions_CType = FieldOptions_STRING
const Default_FieldOptions_Deprecated bool = false

type EnumOptions struct {
	Proto1Name          *string                `protobuf:"bytes,1,opt,name=proto1_name" json:"proto1_name,omitempty"`
	AllowAlias          *bool                  `protobuf:"varint,2,opt,name=allow_alias,def=1" json:"allow_alias,omitempty"`
	UninterpretedOption []*UninterpretedOption `protobuf:"bytes,999,rep,name=uninterpreted_option"`
	XXX_extensions      map[int32][]byte
	XXX_unrecognized    []byte
}

func (this *EnumOptions) Reset()         { *this = EnumOptions{} }
func (this *EnumOptions) String() string { return proto.CompactTextString(this) }

var extRange_EnumOptions = []proto.ExtensionRange{
	proto.ExtensionRange{1000, 536870911},
}

func (*EnumOptions) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_EnumOptions
}
func (this *EnumOptions) ExtensionMap() map[int32][]byte {
	if this.XXX_extensions == nil {
		this.XXX_extensions = make(map[int32][]byte)
	}
	return this.XXX_extensions
}

const Default_EnumOptions_AllowAlias bool = true

type EnumValueOptions struct {
	UninterpretedOption []*UninterpretedOption `protobuf:"bytes,999,rep,name=uninterpreted_option"`
	XXX_extensions      map[int32][]byte
	XXX_unrecognized    []byte
}

func (this *EnumValueOptions) Reset()         { *this = EnumValueOptions{} }
func (this *EnumValueOptions) String() string { return proto.CompactTextString(this) }

var extRange_EnumValueOptions = []proto.ExtensionRange{
	proto.ExtensionRange{1000, 536870911},
}

func (*EnumValueOptions) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_EnumValueOptions
}
func (this *EnumValueOptions) ExtensionMap() map[int32][]byte {
	if this.XXX_extensions == nil {
		this.XXX_extensions = make(map[int32][]byte)
	}
	return this.XXX_extensions
}

type ServiceOptions struct {
	UninterpretedOption []*UninterpretedOption `protobuf:"bytes,999,rep,name=uninterpreted_option"`
	XXX_extensions      map[int32][]byte
	XXX_unrecognized    []byte
}

func (this *ServiceOptions) Reset()         { *this = ServiceOptions{} }
func (this *ServiceOptions) String() string { return proto.CompactTextString(this) }

var extRange_ServiceOptions = []proto.ExtensionRange{
	proto.ExtensionRange{1000, 536870911},
}

func (*ServiceOptions) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_ServiceOptions
}
func (this *ServiceOptions) ExtensionMap() map[int32][]byte {
	if this.XXX_extensions == nil {
		this.XXX_extensions = make(map[int32][]byte)
	}
	return this.XXX_extensions
}

type MethodOptions struct {
	UninterpretedOption []*UninterpretedOption `protobuf:"bytes,999,rep,name=uninterpreted_option"`
	XXX_extensions      map[int32][]byte
	XXX_unrecognized    []byte
}

func (this *MethodOptions) Reset()         { *this = MethodOptions{} }
func (this *MethodOptions) String() string { return proto.CompactTextString(this) }

var extRange_MethodOptions = []proto.ExtensionRange{
	proto.ExtensionRange{1000, 536870911},
}

func (*MethodOptions) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_MethodOptions
}
func (this *MethodOptions) ExtensionMap() map[int32][]byte {
	if this.XXX_extensions == nil {
		this.XXX_extensions = make(map[int32][]byte)
	}
	return this.XXX_extensions
}

type UninterpretedOption struct {
	Name             []*UninterpretedOption_NamePart `protobuf:"bytes,2,rep,name=name"`
	IdentifierValue  *string                         `protobuf:"bytes,3,opt,name=identifier_value"`
	PositiveIntValue *uint64                         `protobuf:"varint,4,opt,name=positive_int_value"`
	NegativeIntValue *int64                          `protobuf:"varint,5,opt,name=negative_int_value"`
	DoubleValue      *float64                        `protobuf:"fixed64,6,opt,name=double_value"`
	StringValue      []byte                          `protobuf:"bytes,7,opt,name=string_value"`
	XXX_unrecognized []byte
}

func (this *UninterpretedOption) Reset()         { *this = UninterpretedOption{} }
func (this *UninterpretedOption) String() string { return proto.CompactTextString(this) }

type UninterpretedOption_NamePart struct {
	NamePart         *string `protobuf:"bytes,1,req,name=name_part"`
	IsExtension      *bool   `protobuf:"varint,2,req,name=is_extension"`
	XXX_unrecognized []byte
}

func (this *UninterpretedOption_NamePart) Reset()         { *this = UninterpretedOption_NamePart{} }
func (this *UninterpretedOption_NamePart) String() string { return proto.CompactTextString(this) }

func init() {
	proto.RegisterEnum("google_protobuf.FieldDescriptorProto_Type", FieldDescriptorProto_Type_name, FieldDescriptorProto_Type_value)
	proto.RegisterEnum("google_protobuf.FieldDescriptorProto_Label", FieldDescriptorProto_Label_name, FieldDescriptorProto_Label_value)
	proto.RegisterEnum("google_protobuf.FileOptions_OptimizeMode", FileOptions_OptimizeMode_name, FileOptions_OptimizeMode_value)
	proto.RegisterEnum("google_protobuf.FieldOptions_CType", FieldOptions_CType_name, FieldOptions_CType_value)
}
