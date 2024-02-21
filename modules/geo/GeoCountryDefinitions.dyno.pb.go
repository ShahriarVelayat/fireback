//
//
//
//
//

package geo

import (
	reflect "reflect"
	sync "sync"

	"github.com/torabian/fireback/modules/workspaces"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GeoCountryCreateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *GeoCountryEntity  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error *workspaces.IError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GeoCountryCreateReply) Reset() {
	*x = GeoCountryCreateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoCountryCreateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoCountryCreateReply) ProtoMessage() {}

func (x *GeoCountryCreateReply) ProtoReflect() protoreflect.Message {
	mi := &file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GeoCountryCreateReply) Descriptor() ([]byte, []int) {
	return file_modules_geo_GeoCountryDefinitions_dyno_proto_rawDescGZIP(), []int{0}
}

func (x *GeoCountryCreateReply) GetData() *GeoCountryEntity {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GeoCountryCreateReply) GetError() *workspaces.IError {
	if x != nil {
		return x.Error
	}
	return nil
}

type GeoCountryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *GeoCountryEntity  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error *workspaces.IError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GeoCountryReply) Reset() {
	*x = GeoCountryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoCountryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoCountryReply) ProtoMessage() {}

func (x *GeoCountryReply) ProtoReflect() protoreflect.Message {
	mi := &file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GeoCountryReply) Descriptor() ([]byte, []int) {
	return file_modules_geo_GeoCountryDefinitions_dyno_proto_rawDescGZIP(), []int{1}
}

func (x *GeoCountryReply) GetData() *GeoCountryEntity {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GeoCountryReply) GetError() *workspaces.IError {
	if x != nil {
		return x.Error
	}
	return nil
}

type GeoCountryQueryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items        []*GeoCountryEntity `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	TotalItems   int64               `protobuf:"varint,2,opt,name=totalItems,proto3" json:"totalItems,omitempty"`
	ItemsPerPage int64               `protobuf:"varint,3,opt,name=itemsPerPage,proto3" json:"itemsPerPage,omitempty"`
	StartIndex   int64               `protobuf:"varint,4,opt,name=startIndex,proto3" json:"startIndex,omitempty"`
	Error        *workspaces.IError  `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GeoCountryQueryReply) Reset() {
	*x = GeoCountryQueryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoCountryQueryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoCountryQueryReply) ProtoMessage() {}

func (x *GeoCountryQueryReply) ProtoReflect() protoreflect.Message {
	mi := &file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GeoCountryQueryReply) Descriptor() ([]byte, []int) {
	return file_modules_geo_GeoCountryDefinitions_dyno_proto_rawDescGZIP(), []int{2}
}

func (x *GeoCountryQueryReply) GetItems() []*GeoCountryEntity {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *GeoCountryQueryReply) GetTotalItems() int64 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

func (x *GeoCountryQueryReply) GetItemsPerPage() int64 {
	if x != nil {
		return x.ItemsPerPage
	}
	return 0
}

func (x *GeoCountryQueryReply) GetStartIndex() int64 {
	if x != nil {
		return x.StartIndex
	}
	return 0
}

func (x *GeoCountryQueryReply) GetError() *workspaces.IError {
	if x != nil {
		return x.Error
	}
	return nil
}

type GeoCountryEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Visibility       *string                     `protobuf:"bytes,1,opt,name=visibility,proto3,oneof" json:"visibility,omitempty" yaml:"visibility"`
	WorkspaceId      *string                     `protobuf:"bytes,2,opt,name=workspaceId,proto3,oneof" json:"workspaceId,omitempty" yaml:"workspaceId"`
	LinkerId         *string                     `protobuf:"bytes,3,opt,name=linkerId,proto3,oneof" json:"linkerId,omitempty" yaml:"linkerId"`
	ParentId         *string                     `protobuf:"bytes,4,opt,name=parentId,proto3,oneof" json:"parentId,omitempty" yaml:"parentId"`
	Parent           *GeoCountryEntity           `protobuf:"bytes,5,opt,name=parent,proto3,oneof" json:"parent,omitempty" yaml:"parent"`
	UniqueId         string                      `protobuf:"bytes,6,opt,name=uniqueId,proto3" json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
	UserId           *string                     `protobuf:"bytes,7,opt,name=userId,proto3,oneof" json:"userId,omitempty" yaml:"userId"`
	Translations     []*GeoCountryEntityPolyglot `protobuf:"bytes,8,rep,name=translations,proto3" json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId" json:"translations"`
	Status           *string                     `protobuf:"bytes,10,opt,name=status,proto3,oneof" json:"status,omitempty"   yaml:"status"  `
	Flag             *string                     `protobuf:"bytes,11,opt,name=flag,proto3,oneof" json:"flag,omitempty"   yaml:"flag"  `
	CommonName       *string                     `protobuf:"bytes,12,opt,name=commonName,proto3,oneof" json:"commonName,omitempty" translate:"true"  yaml:"commonName"  `
	OfficialName     *string                     `protobuf:"bytes,13,opt,name=officialName,proto3,oneof" json:"officialName,omitempty" translate:"true"  yaml:"officialName"  `
	Rank             int64                       `protobuf:"varint,14,opt,name=rank,proto3" json:"rank,omitempty" gorm:"type:int;name:rank"`
	Updated          int64                       `protobuf:"varint,15,opt,name=updated,proto3" json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created          int64                       `protobuf:"varint,16,opt,name=created,proto3" json:"created,omitempty" gorm:"autoUpdateTime:nano"`
	CreatedFormatted string                      `protobuf:"bytes,17,opt,name=createdFormatted,proto3" json:"createdFormatted,omitempty" sql:"-"`
	UpdatedFormatted string                      `protobuf:"bytes,18,opt,name=updatedFormatted,proto3" json:"updatedFormatted,omitempty" sql:"-"`
}

func (x *GeoCountryEntity) Reset() {
	*x = GeoCountryEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoCountryEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoCountryEntity) ProtoMessage() {}

func (x *GeoCountryEntity) ProtoReflect() protoreflect.Message {
	mi := &file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GeoCountryEntity) Descriptor() ([]byte, []int) {
	return file_modules_geo_GeoCountryDefinitions_dyno_proto_rawDescGZIP(), []int{3}
}

func (x *GeoCountryEntity) GetVisibility() string {
	if x != nil && x.Visibility != nil {
		return *x.Visibility
	}
	return ""
}

func (x *GeoCountryEntity) GetWorkspaceId() string {
	if x != nil && x.WorkspaceId != nil {
		return *x.WorkspaceId
	}
	return ""
}

func (x *GeoCountryEntity) GetLinkerId() string {
	if x != nil && x.LinkerId != nil {
		return *x.LinkerId
	}
	return ""
}

func (x *GeoCountryEntity) GetParentId() string {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return ""
}

func (x *GeoCountryEntity) GetParent() *GeoCountryEntity {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *GeoCountryEntity) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

func (x *GeoCountryEntity) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *GeoCountryEntity) GetTranslations() []*GeoCountryEntityPolyglot {
	if x != nil {
		return x.Translations
	}
	return nil
}

func (x *GeoCountryEntity) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *GeoCountryEntity) GetFlag() string {
	if x != nil && x.Flag != nil {
		return *x.Flag
	}
	return ""
}

func (x *GeoCountryEntity) GetCommonName() string {
	if x != nil && x.CommonName != nil {
		return *x.CommonName
	}
	return ""
}

func (x *GeoCountryEntity) GetOfficialName() string {
	if x != nil && x.OfficialName != nil {
		return *x.OfficialName
	}
	return ""
}

func (x *GeoCountryEntity) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *GeoCountryEntity) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *GeoCountryEntity) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *GeoCountryEntity) GetCreatedFormatted() string {
	if x != nil {
		return x.CreatedFormatted
	}
	return ""
}

func (x *GeoCountryEntity) GetUpdatedFormatted() string {
	if x != nil {
		return x.UpdatedFormatted
	}
	return ""
}

type GeoCountryEntityPolyglot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LinkerId     string `protobuf:"bytes,1,opt,name=linkerId,proto3" json:"linkerId,omitempty" gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
	LanguageId   string `protobuf:"bytes,2,opt,name=languageId,proto3" json:"languageId,omitempty" gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
	CommonName   string `protobuf:"bytes,3,opt,name=commonName,proto3" json:"commonName,omitempty" yaml:"commonName" json:"commonName"`
	OfficialName string `protobuf:"bytes,4,opt,name=officialName,proto3" json:"officialName,omitempty" yaml:"officialName" json:"officialName"`
}

func (x *GeoCountryEntityPolyglot) Reset() {
	*x = GeoCountryEntityPolyglot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoCountryEntityPolyglot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoCountryEntityPolyglot) ProtoMessage() {}

func (x *GeoCountryEntityPolyglot) ProtoReflect() protoreflect.Message {
	mi := &file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GeoCountryEntityPolyglot) Descriptor() ([]byte, []int) {
	return file_modules_geo_GeoCountryDefinitions_dyno_proto_rawDescGZIP(), []int{4}
}

func (x *GeoCountryEntityPolyglot) GetLinkerId() string {
	if x != nil {
		return x.LinkerId
	}
	return ""
}

func (x *GeoCountryEntityPolyglot) GetLanguageId() string {
	if x != nil {
		return x.LanguageId
	}
	return ""
}

func (x *GeoCountryEntityPolyglot) GetCommonName() string {
	if x != nil {
		return x.CommonName
	}
	return ""
}

func (x *GeoCountryEntityPolyglot) GetOfficialName() string {
	if x != nil {
		return x.OfficialName
	}
	return ""
}

var File_modules_geo_GeoCountryDefinitions_dyno_proto protoreflect.FileDescriptor

var file_modules_geo_GeoCountryDefinitions_dyno_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6f, 0x2f, 0x47, 0x65,
	0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x64, 0x79, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5d, 0x0a, 0x15, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x49, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x57, 0x0a, 0x0f, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x49, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xc2, 0x01, 0x0a, 0x14, 0x47, 0x65,
	0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x49, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xef,
	0x05, 0x0a, 0x10, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x1f, 0x0a, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x48, 0x04, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6c, 0x6f, 0x74, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x88, 0x01, 0x01, 0x12,
	0x23, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x0c, 0x6f, 0x66,
	0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e,
	0x6b, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65,
	0x64, 0x12, 0x2a, 0x0a, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x74, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x66, 0x6c, 0x61, 0x67,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x42,
	0x0f, 0x0a, 0x0d, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x9a, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6c, 0x6f, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x66, 0x66,
	0x69, 0x63, 0x69, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xe4, 0x02,
	0x0a, 0x0b, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x73, 0x12, 0x45, 0x0a,
	0x16, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x6f,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x16, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11,
	0x2e, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x15, 0x47,
	0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x13, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x6f, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x41, 0x0a, 0x16, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x13, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x16, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12,
	0x13, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x70, 0x6c, 0x75,
	0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_modules_geo_GeoCountryDefinitions_dyno_proto_rawDescOnce sync.Once
	file_modules_geo_GeoCountryDefinitions_dyno_proto_rawDescData = file_modules_geo_GeoCountryDefinitions_dyno_proto_rawDesc
)

func file_modules_geo_GeoCountryDefinitions_dyno_proto_rawDescGZIP() []byte {
	file_modules_geo_GeoCountryDefinitions_dyno_proto_rawDescOnce.Do(func() {
		file_modules_geo_GeoCountryDefinitions_dyno_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_geo_GeoCountryDefinitions_dyno_proto_rawDescData)
	})
	return file_modules_geo_GeoCountryDefinitions_dyno_proto_rawDescData
}

var file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_modules_geo_GeoCountryDefinitions_dyno_proto_goTypes = []interface{}{
	(*GeoCountryCreateReply)(nil),         // 0: GeoCountryCreateReply
	(*GeoCountryReply)(nil),               // 1: GeoCountryReply
	(*GeoCountryQueryReply)(nil),          // 2: GeoCountryQueryReply
	(*GeoCountryEntity)(nil),              // 3: GeoCountryEntity
	(*GeoCountryEntityPolyglot)(nil),      // 4: GeoCountryEntityPolyglot
	(*workspaces.IError)(nil),             // 5: IError
	(*workspaces.QueryFilterRequest)(nil), // 6: QueryFilterRequest
	(*workspaces.RemoveReply)(nil),        // 7: RemoveReply
}
var file_modules_geo_GeoCountryDefinitions_dyno_proto_depIdxs = []int32{
	3,  // 0: GeoCountryCreateReply.data:type_name -> GeoCountryEntity
	5,  // 1: GeoCountryCreateReply.error:type_name -> IError
	3,  // 2: GeoCountryReply.data:type_name -> GeoCountryEntity
	5,  // 3: GeoCountryReply.error:type_name -> IError
	3,  // 4: GeoCountryQueryReply.items:type_name -> GeoCountryEntity
	5,  // 5: GeoCountryQueryReply.error:type_name -> IError
	3,  // 6: GeoCountryEntity.parent:type_name -> GeoCountryEntity
	4,  // 7: GeoCountryEntity.translations:type_name -> GeoCountryEntityPolyglot
	3,  // 8: GeoCountrys.GeoCountryActionCreate:input_type -> GeoCountryEntity
	3,  // 9: GeoCountrys.GeoCountryActionUpdate:input_type -> GeoCountryEntity
	6,  // 10: GeoCountrys.GeoCountryActionQuery:input_type -> QueryFilterRequest
	6,  // 11: GeoCountrys.GeoCountryActionGetOne:input_type -> QueryFilterRequest
	6,  // 12: GeoCountrys.GeoCountryActionRemove:input_type -> QueryFilterRequest
	0,  // 13: GeoCountrys.GeoCountryActionCreate:output_type -> GeoCountryCreateReply
	0,  // 14: GeoCountrys.GeoCountryActionUpdate:output_type -> GeoCountryCreateReply
	2,  // 15: GeoCountrys.GeoCountryActionQuery:output_type -> GeoCountryQueryReply
	1,  // 16: GeoCountrys.GeoCountryActionGetOne:output_type -> GeoCountryReply
	7,  // 17: GeoCountrys.GeoCountryActionRemove:output_type -> RemoveReply
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_modules_geo_GeoCountryDefinitions_dyno_proto_init() }
func file_modules_geo_GeoCountryDefinitions_dyno_proto_init() {
	if File_modules_geo_GeoCountryDefinitions_dyno_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoCountryCreateReply); i {
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
		file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoCountryReply); i {
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
		file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoCountryQueryReply); i {
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
		file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoCountryEntity); i {
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
		file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoCountryEntityPolyglot); i {
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
	file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_modules_geo_GeoCountryDefinitions_dyno_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_geo_GeoCountryDefinitions_dyno_proto_goTypes,
		DependencyIndexes: file_modules_geo_GeoCountryDefinitions_dyno_proto_depIdxs,
		MessageInfos:      file_modules_geo_GeoCountryDefinitions_dyno_proto_msgTypes,
	}.Build()
	File_modules_geo_GeoCountryDefinitions_dyno_proto = out.File
	file_modules_geo_GeoCountryDefinitions_dyno_proto_rawDesc = nil
	file_modules_geo_GeoCountryDefinitions_dyno_proto_goTypes = nil
	file_modules_geo_GeoCountryDefinitions_dyno_proto_depIdxs = nil
}
