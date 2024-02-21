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

type GeoCityCreateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *GeoCityEntity     `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error *workspaces.IError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GeoCityCreateReply) Reset() {
	*x = GeoCityCreateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_geo_GeoCityDefinitions_dyno_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoCityCreateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoCityCreateReply) ProtoMessage() {}

func (x *GeoCityCreateReply) ProtoReflect() protoreflect.Message {
	mi := &file_modules_geo_GeoCityDefinitions_dyno_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GeoCityCreateReply) Descriptor() ([]byte, []int) {
	return file_modules_geo_GeoCityDefinitions_dyno_proto_rawDescGZIP(), []int{0}
}

func (x *GeoCityCreateReply) GetData() *GeoCityEntity {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GeoCityCreateReply) GetError() *workspaces.IError {
	if x != nil {
		return x.Error
	}
	return nil
}

type GeoCityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *GeoCityEntity     `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error *workspaces.IError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GeoCityReply) Reset() {
	*x = GeoCityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_geo_GeoCityDefinitions_dyno_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoCityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoCityReply) ProtoMessage() {}

func (x *GeoCityReply) ProtoReflect() protoreflect.Message {
	mi := &file_modules_geo_GeoCityDefinitions_dyno_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GeoCityReply) Descriptor() ([]byte, []int) {
	return file_modules_geo_GeoCityDefinitions_dyno_proto_rawDescGZIP(), []int{1}
}

func (x *GeoCityReply) GetData() *GeoCityEntity {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GeoCityReply) GetError() *workspaces.IError {
	if x != nil {
		return x.Error
	}
	return nil
}

type GeoCityQueryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items        []*GeoCityEntity   `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	TotalItems   int64              `protobuf:"varint,2,opt,name=totalItems,proto3" json:"totalItems,omitempty"`
	ItemsPerPage int64              `protobuf:"varint,3,opt,name=itemsPerPage,proto3" json:"itemsPerPage,omitempty"`
	StartIndex   int64              `protobuf:"varint,4,opt,name=startIndex,proto3" json:"startIndex,omitempty"`
	Error        *workspaces.IError `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GeoCityQueryReply) Reset() {
	*x = GeoCityQueryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_geo_GeoCityDefinitions_dyno_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoCityQueryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoCityQueryReply) ProtoMessage() {}

func (x *GeoCityQueryReply) ProtoReflect() protoreflect.Message {
	mi := &file_modules_geo_GeoCityDefinitions_dyno_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GeoCityQueryReply) Descriptor() ([]byte, []int) {
	return file_modules_geo_GeoCityDefinitions_dyno_proto_rawDescGZIP(), []int{2}
}

func (x *GeoCityQueryReply) GetItems() []*GeoCityEntity {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *GeoCityQueryReply) GetTotalItems() int64 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

func (x *GeoCityQueryReply) GetItemsPerPage() int64 {
	if x != nil {
		return x.ItemsPerPage
	}
	return 0
}

func (x *GeoCityQueryReply) GetStartIndex() int64 {
	if x != nil {
		return x.StartIndex
	}
	return 0
}

func (x *GeoCityQueryReply) GetError() *workspaces.IError {
	if x != nil {
		return x.Error
	}
	return nil
}

type GeoCityEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Visibility  *string        `protobuf:"bytes,1,opt,name=visibility,proto3,oneof" json:"visibility,omitempty" yaml:"visibility"`
	WorkspaceId *string        `protobuf:"bytes,2,opt,name=workspaceId,proto3,oneof" json:"workspaceId,omitempty" yaml:"workspaceId"`
	LinkerId    *string        `protobuf:"bytes,3,opt,name=linkerId,proto3,oneof" json:"linkerId,omitempty" yaml:"linkerId"`
	ParentId    *string        `protobuf:"bytes,4,opt,name=parentId,proto3,oneof" json:"parentId,omitempty" yaml:"parentId"`
	Parent      *GeoCityEntity `protobuf:"bytes,5,opt,name=parent,proto3,oneof" json:"parent,omitempty" yaml:"parent"`
	UniqueId    string         `protobuf:"bytes,6,opt,name=uniqueId,proto3" json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
	UserId      *string        `protobuf:"bytes,7,opt,name=userId,proto3,oneof" json:"userId,omitempty" yaml:"userId"`
	Name        *string        `protobuf:"bytes,9,opt,name=name,proto3,oneof" json:"name,omitempty"   yaml:"name"  `
	// One 2 one to external entity
	ProvinceId *string            `protobuf:"bytes,11,opt,name=provinceId,proto3,oneof" json:"provinceId,omitempty" yaml:"provinceId" `
	Province   *GeoProvinceEntity `protobuf:"bytes,12,opt,name=province,proto3" json:"province,omitempty" gorm:"foreignKey:ProvinceId;references:UniqueId" json:"province" yaml:"province" fbtype:"one"`
	// One 2 one to external entity
	StateId *string         `protobuf:"bytes,14,opt,name=stateId,proto3,oneof" json:"stateId,omitempty" yaml:"stateId" `
	State   *GeoStateEntity `protobuf:"bytes,15,opt,name=state,proto3" json:"state,omitempty" gorm:"foreignKey:StateId;references:UniqueId" json:"state" yaml:"state" fbtype:"one"`
	// One 2 one to external entity
	CountryId        *string           `protobuf:"bytes,17,opt,name=countryId,proto3,oneof" json:"countryId,omitempty" yaml:"countryId" `
	Country          *GeoCountryEntity `protobuf:"bytes,18,opt,name=country,proto3" json:"country,omitempty" gorm:"foreignKey:CountryId;references:UniqueId" json:"country" yaml:"country" fbtype:"one"`
	Rank             int64             `protobuf:"varint,19,opt,name=rank,proto3" json:"rank,omitempty" gorm:"type:int;name:rank"`
	Updated          int64             `protobuf:"varint,20,opt,name=updated,proto3" json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created          int64             `protobuf:"varint,21,opt,name=created,proto3" json:"created,omitempty" gorm:"autoUpdateTime:nano"`
	CreatedFormatted string            `protobuf:"bytes,22,opt,name=createdFormatted,proto3" json:"createdFormatted,omitempty" sql:"-"`
	UpdatedFormatted string            `protobuf:"bytes,23,opt,name=updatedFormatted,proto3" json:"updatedFormatted,omitempty" sql:"-"`
}

func (x *GeoCityEntity) Reset() {
	*x = GeoCityEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_geo_GeoCityDefinitions_dyno_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoCityEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoCityEntity) ProtoMessage() {}

func (x *GeoCityEntity) ProtoReflect() protoreflect.Message {
	mi := &file_modules_geo_GeoCityDefinitions_dyno_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GeoCityEntity) Descriptor() ([]byte, []int) {
	return file_modules_geo_GeoCityDefinitions_dyno_proto_rawDescGZIP(), []int{3}
}

func (x *GeoCityEntity) GetVisibility() string {
	if x != nil && x.Visibility != nil {
		return *x.Visibility
	}
	return ""
}

func (x *GeoCityEntity) GetWorkspaceId() string {
	if x != nil && x.WorkspaceId != nil {
		return *x.WorkspaceId
	}
	return ""
}

func (x *GeoCityEntity) GetLinkerId() string {
	if x != nil && x.LinkerId != nil {
		return *x.LinkerId
	}
	return ""
}

func (x *GeoCityEntity) GetParentId() string {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return ""
}

func (x *GeoCityEntity) GetParent() *GeoCityEntity {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *GeoCityEntity) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

func (x *GeoCityEntity) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *GeoCityEntity) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *GeoCityEntity) GetProvinceId() string {
	if x != nil && x.ProvinceId != nil {
		return *x.ProvinceId
	}
	return ""
}

func (x *GeoCityEntity) GetProvince() *GeoProvinceEntity {
	if x != nil {
		return x.Province
	}
	return nil
}

func (x *GeoCityEntity) GetStateId() string {
	if x != nil && x.StateId != nil {
		return *x.StateId
	}
	return ""
}

func (x *GeoCityEntity) GetState() *GeoStateEntity {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *GeoCityEntity) GetCountryId() string {
	if x != nil && x.CountryId != nil {
		return *x.CountryId
	}
	return ""
}

func (x *GeoCityEntity) GetCountry() *GeoCountryEntity {
	if x != nil {
		return x.Country
	}
	return nil
}

func (x *GeoCityEntity) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *GeoCityEntity) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *GeoCityEntity) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *GeoCityEntity) GetCreatedFormatted() string {
	if x != nil {
		return x.CreatedFormatted
	}
	return ""
}

func (x *GeoCityEntity) GetUpdatedFormatted() string {
	if x != nil {
		return x.UpdatedFormatted
	}
	return ""
}

var File_modules_geo_GeoCityDefinitions_dyno_proto protoreflect.FileDescriptor

var file_modules_geo_GeoCityDefinitions_dyno_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6f, 0x2f, 0x47, 0x65,
	0x6f, 0x43, 0x69, 0x74, 0x79, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x64, 0x79, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6f, 0x2f, 0x47, 0x65, 0x6f, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x64, 0x79, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6f, 0x2f, 0x47, 0x65, 0x6f, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x64,
	0x79, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x2f, 0x67, 0x65, 0x6f, 0x2f, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x64, 0x79, 0x6e,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x12, 0x47, 0x65, 0x6f, 0x43, 0x69,
	0x74, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x47, 0x65,
	0x6f, 0x43, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x49, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x51, 0x0a, 0x0c, 0x47, 0x65, 0x6f, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x47, 0x65, 0x6f, 0x43, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x49, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0xbc, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x6f, 0x43, 0x69, 0x74, 0x79, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x69,
	0x74, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x50, 0x65, 0x72, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x49, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0xa8, 0x06, 0x0a, 0x0d, 0x47, 0x65, 0x6f, 0x43, 0x69, 0x74, 0x79, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x1f, 0x0a, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x02, 0x52, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x2b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x48, 0x04, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x23, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x47, 0x65, 0x6f, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x07, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x47, 0x65, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x09,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x2b, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b,
	0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64,
	0x12, 0x2a, 0x0a, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x74, 0x65, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x49, 0x64, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x32, 0xc0, 0x02,
	0x0a, 0x08, 0x47, 0x65, 0x6f, 0x43, 0x69, 0x74, 0x79, 0x73, 0x12, 0x3c, 0x0a, 0x13, 0x47, 0x65,
	0x6f, 0x43, 0x69, 0x74, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x0e, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x69, 0x74, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x13, 0x47, 0x65, 0x6f, 0x43,
	0x69, 0x74, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x0e, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a,
	0x13, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x69, 0x74, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x12, 0x47, 0x65, 0x6f, 0x43, 0x69, 0x74,
	0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x13, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x69, 0x74, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x13, 0x47, 0x65, 0x6f, 0x43, 0x69,
	0x74, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x13,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x13, 0x47, 0x65, 0x6f, 0x43, 0x69, 0x74, 0x79, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x42, 0x24, 0x5a, 0x22, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x70, 0x6c, 0x75, 0x78, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x2f, 0x67, 0x65, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modules_geo_GeoCityDefinitions_dyno_proto_rawDescOnce sync.Once
	file_modules_geo_GeoCityDefinitions_dyno_proto_rawDescData = file_modules_geo_GeoCityDefinitions_dyno_proto_rawDesc
)

func file_modules_geo_GeoCityDefinitions_dyno_proto_rawDescGZIP() []byte {
	file_modules_geo_GeoCityDefinitions_dyno_proto_rawDescOnce.Do(func() {
		file_modules_geo_GeoCityDefinitions_dyno_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_geo_GeoCityDefinitions_dyno_proto_rawDescData)
	})
	return file_modules_geo_GeoCityDefinitions_dyno_proto_rawDescData
}

var file_modules_geo_GeoCityDefinitions_dyno_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_modules_geo_GeoCityDefinitions_dyno_proto_goTypes = []interface{}{
	(*GeoCityCreateReply)(nil),            // 0: GeoCityCreateReply
	(*GeoCityReply)(nil),                  // 1: GeoCityReply
	(*GeoCityQueryReply)(nil),             // 2: GeoCityQueryReply
	(*GeoCityEntity)(nil),                 // 3: GeoCityEntity
	(*workspaces.IError)(nil),             // 4: IError
	(*GeoProvinceEntity)(nil),             // 5: GeoProvinceEntity
	(*GeoStateEntity)(nil),                // 6: GeoStateEntity
	(*GeoCountryEntity)(nil),              // 7: GeoCountryEntity
	(*workspaces.QueryFilterRequest)(nil), // 8: QueryFilterRequest
	(*workspaces.RemoveReply)(nil),        // 9: RemoveReply
}
var file_modules_geo_GeoCityDefinitions_dyno_proto_depIdxs = []int32{
	3,  // 0: GeoCityCreateReply.data:type_name -> GeoCityEntity
	4,  // 1: GeoCityCreateReply.error:type_name -> IError
	3,  // 2: GeoCityReply.data:type_name -> GeoCityEntity
	4,  // 3: GeoCityReply.error:type_name -> IError
	3,  // 4: GeoCityQueryReply.items:type_name -> GeoCityEntity
	4,  // 5: GeoCityQueryReply.error:type_name -> IError
	3,  // 6: GeoCityEntity.parent:type_name -> GeoCityEntity
	5,  // 7: GeoCityEntity.province:type_name -> GeoProvinceEntity
	6,  // 8: GeoCityEntity.state:type_name -> GeoStateEntity
	7,  // 9: GeoCityEntity.country:type_name -> GeoCountryEntity
	3,  // 10: GeoCitys.GeoCityActionCreate:input_type -> GeoCityEntity
	3,  // 11: GeoCitys.GeoCityActionUpdate:input_type -> GeoCityEntity
	8,  // 12: GeoCitys.GeoCityActionQuery:input_type -> QueryFilterRequest
	8,  // 13: GeoCitys.GeoCityActionGetOne:input_type -> QueryFilterRequest
	8,  // 14: GeoCitys.GeoCityActionRemove:input_type -> QueryFilterRequest
	0,  // 15: GeoCitys.GeoCityActionCreate:output_type -> GeoCityCreateReply
	0,  // 16: GeoCitys.GeoCityActionUpdate:output_type -> GeoCityCreateReply
	2,  // 17: GeoCitys.GeoCityActionQuery:output_type -> GeoCityQueryReply
	1,  // 18: GeoCitys.GeoCityActionGetOne:output_type -> GeoCityReply
	9,  // 19: GeoCitys.GeoCityActionRemove:output_type -> RemoveReply
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_modules_geo_GeoCityDefinitions_dyno_proto_init() }
func file_modules_geo_GeoCityDefinitions_dyno_proto_init() {
	if File_modules_geo_GeoCityDefinitions_dyno_proto != nil {
		return
	}
	file_modules_geo_GeoStateDefinitions_dyno_proto_init()
	file_modules_geo_GeoProvinceDefinitions_dyno_proto_init()
	file_modules_geo_GeoCountryDefinitions_dyno_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_modules_geo_GeoCityDefinitions_dyno_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoCityCreateReply); i {
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
		file_modules_geo_GeoCityDefinitions_dyno_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoCityReply); i {
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
		file_modules_geo_GeoCityDefinitions_dyno_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoCityQueryReply); i {
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
		file_modules_geo_GeoCityDefinitions_dyno_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoCityEntity); i {
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
	file_modules_geo_GeoCityDefinitions_dyno_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_modules_geo_GeoCityDefinitions_dyno_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_geo_GeoCityDefinitions_dyno_proto_goTypes,
		DependencyIndexes: file_modules_geo_GeoCityDefinitions_dyno_proto_depIdxs,
		MessageInfos:      file_modules_geo_GeoCityDefinitions_dyno_proto_msgTypes,
	}.Build()
	File_modules_geo_GeoCityDefinitions_dyno_proto = out.File
	file_modules_geo_GeoCityDefinitions_dyno_proto_rawDesc = nil
	file_modules_geo_GeoCityDefinitions_dyno_proto_goTypes = nil
	file_modules_geo_GeoCityDefinitions_dyno_proto_depIdxs = nil
}
