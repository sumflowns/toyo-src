// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product_es.proto

package product_es

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

//信息
type Product struct {
	Id                         string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BrandId                    int64    `protobuf:"varint,2,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	ProductCategoryId          int64    `protobuf:"varint,3,opt,name=product_category_id,json=productCategoryId,proto3" json:"product_category_id,omitempty"`
	FeightTemplateId           int64    `protobuf:"varint,4,opt,name=feight_template_id,json=feightTemplateId,proto3" json:"feight_template_id,omitempty"`
	ProductAttributeCategoryId int64    `protobuf:"varint,5,opt,name=product_attribute_category_id,json=productAttributeCategoryId,proto3" json:"product_attribute_category_id,omitempty"`
	Name                       string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Pic                        string   `protobuf:"bytes,7,opt,name=pic,proto3" json:"pic,omitempty"`
	ProductSn                  string   `protobuf:"bytes,8,opt,name=product_sn,json=productSn,proto3" json:"product_sn,omitempty"`
	DeleteStatus               int64    `protobuf:"varint,9,opt,name=delete_status,json=deleteStatus,proto3" json:"delete_status,omitempty"`
	PublishStatus              int64    `protobuf:"varint,10,opt,name=publish_status,json=publishStatus,proto3" json:"publish_status,omitempty"`
	NewStatus                  int64    `protobuf:"varint,11,opt,name=new_status,json=newStatus,proto3" json:"new_status,omitempty"`
	RecommandStatus            int64    `protobuf:"varint,12,opt,name=recommand_status,json=recommandStatus,proto3" json:"recommand_status,omitempty"`
	VerifyStatus               int64    `protobuf:"varint,13,opt,name=verify_status,json=verifyStatus,proto3" json:"verify_status,omitempty"`
	Sort                       int64    `protobuf:"varint,14,opt,name=sort,proto3" json:"sort,omitempty"`
	Sale                       int64    `protobuf:"varint,15,opt,name=sale,proto3" json:"sale,omitempty"`
	Price                      float32  `protobuf:"fixed32,16,opt,name=price,proto3" json:"price,omitempty"`
	PromotionPrice             float32  `protobuf:"fixed32,17,opt,name=promotion_price,json=promotionPrice,proto3" json:"promotion_price,omitempty"`
	GiftGrowth                 int64    `protobuf:"varint,18,opt,name=gift_growth,json=giftGrowth,proto3" json:"gift_growth,omitempty"`
	GiftPoint                  int64    `protobuf:"varint,19,opt,name=gift_point,json=giftPoint,proto3" json:"gift_point,omitempty"`
	UsePointLimit              int64    `protobuf:"varint,20,opt,name=use_point_limit,json=usePointLimit,proto3" json:"use_point_limit,omitempty"`
	SubTitle                   string   `protobuf:"bytes,21,opt,name=sub_title,json=subTitle,proto3" json:"sub_title,omitempty"`
	Description                string   `protobuf:"bytes,22,opt,name=description,proto3" json:"description,omitempty"`
	OriginalPrice              float32  `protobuf:"fixed32,23,opt,name=original_price,json=originalPrice,proto3" json:"original_price,omitempty"`
	Stock                      int64    `protobuf:"varint,24,opt,name=stock,proto3" json:"stock,omitempty"`
	LowStock                   int64    `protobuf:"varint,25,opt,name=low_stock,json=lowStock,proto3" json:"low_stock,omitempty"`
	Unit                       string   `protobuf:"bytes,26,opt,name=unit,proto3" json:"unit,omitempty"`
	Weight                     float32  `protobuf:"fixed32,27,opt,name=weight,proto3" json:"weight,omitempty"`
	PreviewStatus              int64    `protobuf:"varint,28,opt,name=preview_status,json=previewStatus,proto3" json:"preview_status,omitempty"`
	ServiceIds                 string   `protobuf:"bytes,29,opt,name=service_ids,json=serviceIds,proto3" json:"service_ids,omitempty"`
	Keywords                   string   `protobuf:"bytes,30,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Note                       string   `protobuf:"bytes,31,opt,name=note,proto3" json:"note,omitempty"`
	AlbumPics                  string   `protobuf:"bytes,32,opt,name=album_pics,json=albumPics,proto3" json:"album_pics,omitempty"`
	DetailTitle                string   `protobuf:"bytes,33,opt,name=detail_title,json=detailTitle,proto3" json:"detail_title,omitempty"`
	DetailDesc                 string   `protobuf:"bytes,34,opt,name=detail_desc,json=detailDesc,proto3" json:"detail_desc,omitempty"`
	DetailHtml                 string   `protobuf:"bytes,35,opt,name=detail_html,json=detailHtml,proto3" json:"detail_html,omitempty"`
	DetailMobileHtml           string   `protobuf:"bytes,36,opt,name=detail_mobile_html,json=detailMobileHtml,proto3" json:"detail_mobile_html,omitempty"`
	PromotionStartTime         string   `protobuf:"bytes,37,opt,name=promotion_start_time,json=promotionStartTime,proto3" json:"promotion_start_time,omitempty"`
	PromotionEndTime           string   `protobuf:"bytes,38,opt,name=promotion_end_time,json=promotionEndTime,proto3" json:"promotion_end_time,omitempty"`
	PromotionPerLimit          int64    `protobuf:"varint,39,opt,name=promotion_per_limit,json=promotionPerLimit,proto3" json:"promotion_per_limit,omitempty"`
	PromotionType              int64    `protobuf:"varint,40,opt,name=promotion_type,json=promotionType,proto3" json:"promotion_type,omitempty"`
	BrandName                  string   `protobuf:"bytes,41,opt,name=brand_name,json=brandName,proto3" json:"brand_name,omitempty"`
	ProductCategoryName        string   `protobuf:"bytes,42,opt,name=product_category_name,json=productCategoryName,proto3" json:"product_category_name,omitempty"`
	CreatedTime                string   `protobuf:"bytes,43,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	UpdateTime                 string   `protobuf:"bytes,44,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6eac1b64d144ad, []int{0}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Product) GetBrandId() int64 {
	if m != nil {
		return m.BrandId
	}
	return 0
}

func (m *Product) GetProductCategoryId() int64 {
	if m != nil {
		return m.ProductCategoryId
	}
	return 0
}

func (m *Product) GetFeightTemplateId() int64 {
	if m != nil {
		return m.FeightTemplateId
	}
	return 0
}

func (m *Product) GetProductAttributeCategoryId() int64 {
	if m != nil {
		return m.ProductAttributeCategoryId
	}
	return 0
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetPic() string {
	if m != nil {
		return m.Pic
	}
	return ""
}

func (m *Product) GetProductSn() string {
	if m != nil {
		return m.ProductSn
	}
	return ""
}

func (m *Product) GetDeleteStatus() int64 {
	if m != nil {
		return m.DeleteStatus
	}
	return 0
}

func (m *Product) GetPublishStatus() int64 {
	if m != nil {
		return m.PublishStatus
	}
	return 0
}

func (m *Product) GetNewStatus() int64 {
	if m != nil {
		return m.NewStatus
	}
	return 0
}

func (m *Product) GetRecommandStatus() int64 {
	if m != nil {
		return m.RecommandStatus
	}
	return 0
}

func (m *Product) GetVerifyStatus() int64 {
	if m != nil {
		return m.VerifyStatus
	}
	return 0
}

func (m *Product) GetSort() int64 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *Product) GetSale() int64 {
	if m != nil {
		return m.Sale
	}
	return 0
}

func (m *Product) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Product) GetPromotionPrice() float32 {
	if m != nil {
		return m.PromotionPrice
	}
	return 0
}

func (m *Product) GetGiftGrowth() int64 {
	if m != nil {
		return m.GiftGrowth
	}
	return 0
}

func (m *Product) GetGiftPoint() int64 {
	if m != nil {
		return m.GiftPoint
	}
	return 0
}

func (m *Product) GetUsePointLimit() int64 {
	if m != nil {
		return m.UsePointLimit
	}
	return 0
}

func (m *Product) GetSubTitle() string {
	if m != nil {
		return m.SubTitle
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Product) GetOriginalPrice() float32 {
	if m != nil {
		return m.OriginalPrice
	}
	return 0
}

func (m *Product) GetStock() int64 {
	if m != nil {
		return m.Stock
	}
	return 0
}

func (m *Product) GetLowStock() int64 {
	if m != nil {
		return m.LowStock
	}
	return 0
}

func (m *Product) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Product) GetWeight() float32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Product) GetPreviewStatus() int64 {
	if m != nil {
		return m.PreviewStatus
	}
	return 0
}

func (m *Product) GetServiceIds() string {
	if m != nil {
		return m.ServiceIds
	}
	return ""
}

func (m *Product) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *Product) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *Product) GetAlbumPics() string {
	if m != nil {
		return m.AlbumPics
	}
	return ""
}

func (m *Product) GetDetailTitle() string {
	if m != nil {
		return m.DetailTitle
	}
	return ""
}

func (m *Product) GetDetailDesc() string {
	if m != nil {
		return m.DetailDesc
	}
	return ""
}

func (m *Product) GetDetailHtml() string {
	if m != nil {
		return m.DetailHtml
	}
	return ""
}

func (m *Product) GetDetailMobileHtml() string {
	if m != nil {
		return m.DetailMobileHtml
	}
	return ""
}

func (m *Product) GetPromotionStartTime() string {
	if m != nil {
		return m.PromotionStartTime
	}
	return ""
}

func (m *Product) GetPromotionEndTime() string {
	if m != nil {
		return m.PromotionEndTime
	}
	return ""
}

func (m *Product) GetPromotionPerLimit() int64 {
	if m != nil {
		return m.PromotionPerLimit
	}
	return 0
}

func (m *Product) GetPromotionType() int64 {
	if m != nil {
		return m.PromotionType
	}
	return 0
}

func (m *Product) GetBrandName() string {
	if m != nil {
		return m.BrandName
	}
	return ""
}

func (m *Product) GetProductCategoryName() string {
	if m != nil {
		return m.ProductCategoryName
	}
	return ""
}

func (m *Product) GetCreatedTime() string {
	if m != nil {
		return m.CreatedTime
	}
	return ""
}

func (m *Product) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

type ProductEs struct {
	XIndex               string   `protobuf:"bytes,1,opt,name=_index,json=Index,proto3" json:"_index,omitempty"`
	XType                string   `protobuf:"bytes,2,opt,name=_type,json=Type,proto3" json:"_type,omitempty"`
	XId                  string   `protobuf:"bytes,3,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	XScore               int64    `protobuf:"varint,4,opt,name=_score,json=Score,proto3" json:"_score,omitempty"`
	XSource              *Product `protobuf:"bytes,5,opt,name=_source,json=Source,proto3" json:"_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductEs) Reset()         { *m = ProductEs{} }
func (m *ProductEs) String() string { return proto.CompactTextString(m) }
func (*ProductEs) ProtoMessage()    {}
func (*ProductEs) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6eac1b64d144ad, []int{1}
}

func (m *ProductEs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductEs.Unmarshal(m, b)
}
func (m *ProductEs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductEs.Marshal(b, m, deterministic)
}
func (m *ProductEs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductEs.Merge(m, src)
}
func (m *ProductEs) XXX_Size() int {
	return xxx_messageInfo_ProductEs.Size(m)
}
func (m *ProductEs) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductEs.DiscardUnknown(m)
}

var xxx_messageInfo_ProductEs proto.InternalMessageInfo

func (m *ProductEs) GetXIndex() string {
	if m != nil {
		return m.XIndex
	}
	return ""
}

func (m *ProductEs) GetXType() string {
	if m != nil {
		return m.XType
	}
	return ""
}

func (m *ProductEs) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *ProductEs) GetXScore() int64 {
	if m != nil {
		return m.XScore
	}
	return 0
}

func (m *ProductEs) GetXSource() *Product {
	if m != nil {
		return m.XSource
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6eac1b64d144ad, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type In_GetProductById struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *In_GetProductById) Reset()         { *m = In_GetProductById{} }
func (m *In_GetProductById) String() string { return proto.CompactTextString(m) }
func (*In_GetProductById) ProtoMessage()    {}
func (*In_GetProductById) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6eac1b64d144ad, []int{3}
}

func (m *In_GetProductById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In_GetProductById.Unmarshal(m, b)
}
func (m *In_GetProductById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In_GetProductById.Marshal(b, m, deterministic)
}
func (m *In_GetProductById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In_GetProductById.Merge(m, src)
}
func (m *In_GetProductById) XXX_Size() int {
	return xxx_messageInfo_In_GetProductById.Size(m)
}
func (m *In_GetProductById) XXX_DiscardUnknown() {
	xxx_messageInfo_In_GetProductById.DiscardUnknown(m)
}

var xxx_messageInfo_In_GetProductById proto.InternalMessageInfo

func (m *In_GetProductById) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Out_GetProductById struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Product              *Product `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Out_GetProductById) Reset()         { *m = Out_GetProductById{} }
func (m *Out_GetProductById) String() string { return proto.CompactTextString(m) }
func (*Out_GetProductById) ProtoMessage()    {}
func (*Out_GetProductById) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6eac1b64d144ad, []int{4}
}

func (m *Out_GetProductById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out_GetProductById.Unmarshal(m, b)
}
func (m *Out_GetProductById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out_GetProductById.Marshal(b, m, deterministic)
}
func (m *Out_GetProductById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out_GetProductById.Merge(m, src)
}
func (m *Out_GetProductById) XXX_Size() int {
	return xxx_messageInfo_Out_GetProductById.Size(m)
}
func (m *Out_GetProductById) XXX_DiscardUnknown() {
	xxx_messageInfo_Out_GetProductById.DiscardUnknown(m)
}

var xxx_messageInfo_Out_GetProductById proto.InternalMessageInfo

func (m *Out_GetProductById) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Out_GetProductById) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type Out_UpdateProductInfo struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Product              *Product `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Out_UpdateProductInfo) Reset()         { *m = Out_UpdateProductInfo{} }
func (m *Out_UpdateProductInfo) String() string { return proto.CompactTextString(m) }
func (*Out_UpdateProductInfo) ProtoMessage()    {}
func (*Out_UpdateProductInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6eac1b64d144ad, []int{5}
}

func (m *Out_UpdateProductInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out_UpdateProductInfo.Unmarshal(m, b)
}
func (m *Out_UpdateProductInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out_UpdateProductInfo.Marshal(b, m, deterministic)
}
func (m *Out_UpdateProductInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out_UpdateProductInfo.Merge(m, src)
}
func (m *Out_UpdateProductInfo) XXX_Size() int {
	return xxx_messageInfo_Out_UpdateProductInfo.Size(m)
}
func (m *Out_UpdateProductInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Out_UpdateProductInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Out_UpdateProductInfo proto.InternalMessageInfo

func (m *Out_UpdateProductInfo) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Out_UpdateProductInfo) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type In_UpdateProductInfo struct {
	Product              *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *In_UpdateProductInfo) Reset()         { *m = In_UpdateProductInfo{} }
func (m *In_UpdateProductInfo) String() string { return proto.CompactTextString(m) }
func (*In_UpdateProductInfo) ProtoMessage()    {}
func (*In_UpdateProductInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6eac1b64d144ad, []int{6}
}

func (m *In_UpdateProductInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In_UpdateProductInfo.Unmarshal(m, b)
}
func (m *In_UpdateProductInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In_UpdateProductInfo.Marshal(b, m, deterministic)
}
func (m *In_UpdateProductInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In_UpdateProductInfo.Merge(m, src)
}
func (m *In_UpdateProductInfo) XXX_Size() int {
	return xxx_messageInfo_In_UpdateProductInfo.Size(m)
}
func (m *In_UpdateProductInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_In_UpdateProductInfo.DiscardUnknown(m)
}

var xxx_messageInfo_In_UpdateProductInfo proto.InternalMessageInfo

func (m *In_UpdateProductInfo) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type In_GetProducts struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Pages                int64    `protobuf:"varint,2,opt,name=pages,proto3" json:"pages,omitempty"`
	SearchKey            string   `protobuf:"bytes,3,opt,name=search_key,json=searchKey,proto3" json:"search_key,omitempty"`
	StartTime            string   `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              string   `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *In_GetProducts) Reset()         { *m = In_GetProducts{} }
func (m *In_GetProducts) String() string { return proto.CompactTextString(m) }
func (*In_GetProducts) ProtoMessage()    {}
func (*In_GetProducts) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6eac1b64d144ad, []int{7}
}

func (m *In_GetProducts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In_GetProducts.Unmarshal(m, b)
}
func (m *In_GetProducts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In_GetProducts.Marshal(b, m, deterministic)
}
func (m *In_GetProducts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In_GetProducts.Merge(m, src)
}
func (m *In_GetProducts) XXX_Size() int {
	return xxx_messageInfo_In_GetProducts.Size(m)
}
func (m *In_GetProducts) XXX_DiscardUnknown() {
	xxx_messageInfo_In_GetProducts.DiscardUnknown(m)
}

var xxx_messageInfo_In_GetProducts proto.InternalMessageInfo

func (m *In_GetProducts) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *In_GetProducts) GetPages() int64 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *In_GetProducts) GetSearchKey() string {
	if m != nil {
		return m.SearchKey
	}
	return ""
}

func (m *In_GetProducts) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *In_GetProducts) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

type Out_GetProducts struct {
	Error                *Error     `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Limit                int64      `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Pages                int64      `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Total                int64      `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	ProductList          []*Product `protobuf:"bytes,5,rep,name=product_list,json=productList,proto3" json:"product_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Out_GetProducts) Reset()         { *m = Out_GetProducts{} }
func (m *Out_GetProducts) String() string { return proto.CompactTextString(m) }
func (*Out_GetProducts) ProtoMessage()    {}
func (*Out_GetProducts) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6eac1b64d144ad, []int{8}
}

func (m *Out_GetProducts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out_GetProducts.Unmarshal(m, b)
}
func (m *Out_GetProducts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out_GetProducts.Marshal(b, m, deterministic)
}
func (m *Out_GetProducts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out_GetProducts.Merge(m, src)
}
func (m *Out_GetProducts) XXX_Size() int {
	return xxx_messageInfo_Out_GetProducts.Size(m)
}
func (m *Out_GetProducts) XXX_DiscardUnknown() {
	xxx_messageInfo_Out_GetProducts.DiscardUnknown(m)
}

var xxx_messageInfo_Out_GetProducts proto.InternalMessageInfo

func (m *Out_GetProducts) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Out_GetProducts) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Out_GetProducts) GetPages() int64 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *Out_GetProducts) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Out_GetProducts) GetProductList() []*Product {
	if m != nil {
		return m.ProductList
	}
	return nil
}

type In_DeleteProducts struct {
	ProductList          []string `protobuf:"bytes,1,rep,name=product_list,json=productList,proto3" json:"product_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *In_DeleteProducts) Reset()         { *m = In_DeleteProducts{} }
func (m *In_DeleteProducts) String() string { return proto.CompactTextString(m) }
func (*In_DeleteProducts) ProtoMessage()    {}
func (*In_DeleteProducts) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6eac1b64d144ad, []int{9}
}

func (m *In_DeleteProducts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In_DeleteProducts.Unmarshal(m, b)
}
func (m *In_DeleteProducts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In_DeleteProducts.Marshal(b, m, deterministic)
}
func (m *In_DeleteProducts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In_DeleteProducts.Merge(m, src)
}
func (m *In_DeleteProducts) XXX_Size() int {
	return xxx_messageInfo_In_DeleteProducts.Size(m)
}
func (m *In_DeleteProducts) XXX_DiscardUnknown() {
	xxx_messageInfo_In_DeleteProducts.DiscardUnknown(m)
}

var xxx_messageInfo_In_DeleteProducts proto.InternalMessageInfo

func (m *In_DeleteProducts) GetProductList() []string {
	if m != nil {
		return m.ProductList
	}
	return nil
}

type Out_DeleteProducts struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Out_DeleteProducts) Reset()         { *m = Out_DeleteProducts{} }
func (m *Out_DeleteProducts) String() string { return proto.CompactTextString(m) }
func (*Out_DeleteProducts) ProtoMessage()    {}
func (*Out_DeleteProducts) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6eac1b64d144ad, []int{10}
}

func (m *Out_DeleteProducts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out_DeleteProducts.Unmarshal(m, b)
}
func (m *Out_DeleteProducts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out_DeleteProducts.Marshal(b, m, deterministic)
}
func (m *Out_DeleteProducts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out_DeleteProducts.Merge(m, src)
}
func (m *Out_DeleteProducts) XXX_Size() int {
	return xxx_messageInfo_Out_DeleteProducts.Size(m)
}
func (m *Out_DeleteProducts) XXX_DiscardUnknown() {
	xxx_messageInfo_Out_DeleteProducts.DiscardUnknown(m)
}

var xxx_messageInfo_Out_DeleteProducts proto.InternalMessageInfo

func (m *Out_DeleteProducts) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type In_CreateProduct struct {
	Product              *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *In_CreateProduct) Reset()         { *m = In_CreateProduct{} }
func (m *In_CreateProduct) String() string { return proto.CompactTextString(m) }
func (*In_CreateProduct) ProtoMessage()    {}
func (*In_CreateProduct) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6eac1b64d144ad, []int{11}
}

func (m *In_CreateProduct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In_CreateProduct.Unmarshal(m, b)
}
func (m *In_CreateProduct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In_CreateProduct.Marshal(b, m, deterministic)
}
func (m *In_CreateProduct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In_CreateProduct.Merge(m, src)
}
func (m *In_CreateProduct) XXX_Size() int {
	return xxx_messageInfo_In_CreateProduct.Size(m)
}
func (m *In_CreateProduct) XXX_DiscardUnknown() {
	xxx_messageInfo_In_CreateProduct.DiscardUnknown(m)
}

var xxx_messageInfo_In_CreateProduct proto.InternalMessageInfo

func (m *In_CreateProduct) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type Out_CreateProduct struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Product              *Product `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Out_CreateProduct) Reset()         { *m = Out_CreateProduct{} }
func (m *Out_CreateProduct) String() string { return proto.CompactTextString(m) }
func (*Out_CreateProduct) ProtoMessage()    {}
func (*Out_CreateProduct) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6eac1b64d144ad, []int{12}
}

func (m *Out_CreateProduct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out_CreateProduct.Unmarshal(m, b)
}
func (m *Out_CreateProduct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out_CreateProduct.Marshal(b, m, deterministic)
}
func (m *Out_CreateProduct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out_CreateProduct.Merge(m, src)
}
func (m *Out_CreateProduct) XXX_Size() int {
	return xxx_messageInfo_Out_CreateProduct.Size(m)
}
func (m *Out_CreateProduct) XXX_DiscardUnknown() {
	xxx_messageInfo_Out_CreateProduct.DiscardUnknown(m)
}

var xxx_messageInfo_Out_CreateProduct proto.InternalMessageInfo

func (m *Out_CreateProduct) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Out_CreateProduct) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func init() {
	proto.RegisterType((*Product)(nil), "product_es.Product")
	proto.RegisterType((*ProductEs)(nil), "product_es.ProductEs")
	proto.RegisterType((*Error)(nil), "product_es.Error")
	proto.RegisterType((*In_GetProductById)(nil), "product_es.In_GetProductById")
	proto.RegisterType((*Out_GetProductById)(nil), "product_es.Out_GetProductById")
	proto.RegisterType((*Out_UpdateProductInfo)(nil), "product_es.Out_UpdateProductInfo")
	proto.RegisterType((*In_UpdateProductInfo)(nil), "product_es.In_UpdateProductInfo")
	proto.RegisterType((*In_GetProducts)(nil), "product_es.In_GetProducts")
	proto.RegisterType((*Out_GetProducts)(nil), "product_es.Out_GetProducts")
	proto.RegisterType((*In_DeleteProducts)(nil), "product_es.In_DeleteProducts")
	proto.RegisterType((*Out_DeleteProducts)(nil), "product_es.Out_DeleteProducts")
	proto.RegisterType((*In_CreateProduct)(nil), "product_es.In_CreateProduct")
	proto.RegisterType((*Out_CreateProduct)(nil), "product_es.Out_CreateProduct")
}

func init() { proto.RegisterFile("product_es.proto", fileDescriptor_cf6eac1b64d144ad) }

var fileDescriptor_cf6eac1b64d144ad = []byte{
	// 1234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xcb, 0x72, 0xdb, 0xc6,
	0x12, 0x15, 0x49, 0x51, 0x12, 0x9b, 0x7a, 0x50, 0x23, 0xc9, 0x77, 0x4c, 0x5b, 0x16, 0x05, 0x5d,
	0xdb, 0xf2, 0xbd, 0x8e, 0x2b, 0xa5, 0x54, 0xbc, 0xcb, 0xc2, 0xb1, 0x55, 0x36, 0x2b, 0x8e, 0xa3,
	0x50, 0xca, 0x26, 0x1b, 0x14, 0x08, 0xb4, 0xa8, 0x29, 0x01, 0x18, 0xd4, 0xcc, 0xc0, 0x0a, 0x3f,
	0x22, 0xcb, 0xfc, 0x46, 0x56, 0xf9, 0xa1, 0xfc, 0x49, 0x6a, 0x7a, 0x06, 0x14, 0x1f, 0x56, 0x95,
	0xb3, 0xf0, 0x0e, 0x7d, 0xce, 0x41, 0x4f, 0xbf, 0xa6, 0x41, 0x42, 0xa7, 0x50, 0x32, 0x29, 0x63,
	0x13, 0xa2, 0x7e, 0x51, 0x28, 0x69, 0x24, 0x83, 0x5b, 0x24, 0xf8, 0xbb, 0x0d, 0xab, 0x67, 0xce,
	0x64, 0x9b, 0x50, 0x17, 0x09, 0xaf, 0xf5, 0x6a, 0xc7, 0xad, 0x41, 0x5d, 0x24, 0xec, 0x3e, 0xac,
	0x0d, 0x55, 0x94, 0x27, 0xa1, 0x48, 0x78, 0xbd, 0x57, 0x3b, 0x6e, 0x0c, 0x56, 0xc9, 0xee, 0x27,
	0xec, 0x05, 0xec, 0x54, 0x4e, 0xe2, 0xc8, 0xe0, 0x48, 0xaa, 0xb1, 0x55, 0x35, 0x48, 0xb5, 0xed,
	0xa9, 0xd7, 0x9e, 0xe9, 0x27, 0xec, 0x39, 0xb0, 0x4b, 0x14, 0xa3, 0x2b, 0x13, 0x1a, 0xcc, 0x8a,
	0x34, 0x32, 0x68, 0xe5, 0xcb, 0x24, 0xef, 0x38, 0xe6, 0xc2, 0x13, 0xfd, 0x84, 0xbd, 0x82, 0xfd,
	0xca, 0x7b, 0x64, 0x8c, 0x12, 0xc3, 0xd2, 0xe0, 0xcc, 0x39, 0x4d, 0x7a, 0xb1, 0xeb, 0x45, 0xaf,
	0x2a, 0xcd, 0xd4, 0x81, 0x0c, 0x96, 0xf3, 0x28, 0x43, 0xbe, 0x42, 0xd9, 0xd0, 0x33, 0xeb, 0x40,
	0xa3, 0x10, 0x31, 0x5f, 0x25, 0xc8, 0x3e, 0xb2, 0x7d, 0x98, 0xd4, 0x42, 0xe7, 0x7c, 0x8d, 0x88,
	0x96, 0x47, 0xce, 0x73, 0x76, 0x04, 0x1b, 0x09, 0xa6, 0x68, 0x30, 0xd4, 0x26, 0x32, 0xa5, 0xe6,
	0x2d, 0x3a, 0x77, 0xdd, 0x81, 0xe7, 0x84, 0xb1, 0xc7, 0xb0, 0x59, 0x94, 0xc3, 0x54, 0xe8, 0xab,
	0x4a, 0x05, 0xa4, 0xda, 0xf0, 0xa8, 0x97, 0xed, 0x03, 0xe4, 0x78, 0x53, 0x49, 0xda, 0x24, 0x69,
	0xe5, 0x78, 0xe3, 0xe9, 0x67, 0xd0, 0x51, 0x18, 0xcb, 0x2c, 0xb3, 0xf5, 0xf6, 0xa2, 0x75, 0x12,
	0x6d, 0x4d, 0x70, 0x2f, 0x3d, 0x82, 0x8d, 0x8f, 0xa8, 0xc4, 0xe5, 0xb8, 0xd2, 0x6d, 0xb8, 0xa8,
	0x1c, 0xe8, 0x45, 0x0c, 0x96, 0xb5, 0x54, 0x86, 0x6f, 0x12, 0x47, 0xcf, 0x84, 0x45, 0x29, 0xf2,
	0x2d, 0x8f, 0x45, 0x29, 0xb2, 0x5d, 0x68, 0x16, 0x4a, 0xc4, 0xc8, 0x3b, 0xbd, 0xda, 0x71, 0x7d,
	0xe0, 0x0c, 0xf6, 0x14, 0xb6, 0x0a, 0x25, 0x33, 0x69, 0x84, 0xcc, 0x43, 0xc7, 0x6f, 0x13, 0xbf,
	0x39, 0x81, 0xcf, 0x48, 0x78, 0x00, 0xed, 0x91, 0xb8, 0x34, 0xe1, 0x48, 0xc9, 0x1b, 0x73, 0xc5,
	0x19, 0x79, 0x06, 0x0b, 0xbd, 0x25, 0xc4, 0xa6, 0x4d, 0x82, 0x42, 0x8a, 0xdc, 0xf0, 0x1d, 0x97,
	0xb6, 0x45, 0xce, 0x2c, 0xc0, 0x9e, 0xc0, 0x56, 0xa9, 0xd1, 0xb1, 0x61, 0x2a, 0x32, 0x61, 0xf8,
	0xae, 0xab, 0x5e, 0xa9, 0x91, 0x24, 0xef, 0x2d, 0xc8, 0x1e, 0x40, 0x4b, 0x97, 0xc3, 0xd0, 0x08,
	0x93, 0x22, 0xdf, 0xa3, 0x3e, 0xad, 0xe9, 0x72, 0x78, 0x61, 0x6d, 0xd6, 0x83, 0x76, 0x82, 0x3a,
	0x56, 0xa2, 0xb0, 0x81, 0xf1, 0x7b, 0x44, 0x4f, 0x43, 0xb6, 0x47, 0x52, 0x89, 0x91, 0xc8, 0xa3,
	0xd4, 0xa7, 0xf3, 0x1f, 0x4a, 0x67, 0xa3, 0x42, 0x5d, 0x36, 0xbb, 0xd0, 0xd4, 0x46, 0xc6, 0xd7,
	0x9c, 0x53, 0x0c, 0xce, 0xb0, 0x67, 0xa7, 0xd2, 0x76, 0xce, 0x32, 0xf7, 0x89, 0x59, 0x4b, 0xe5,
	0xcd, 0x39, 0x91, 0x0c, 0x96, 0xcb, 0x5c, 0x18, 0xde, 0x75, 0x73, 0x66, 0x9f, 0xd9, 0x3d, 0x58,
	0xb9, 0xa1, 0x91, 0xe6, 0x0f, 0xe8, 0x14, 0x6f, 0xd1, 0xa4, 0x28, 0xfc, 0x28, 0x6e, 0xc7, 0xe0,
	0xa1, 0x9f, 0x14, 0x87, 0xfa, 0xd6, 0x1d, 0x40, 0x5b, 0xa3, 0xfa, 0x28, 0x62, 0x7b, 0x47, 0x34,
	0xdf, 0x27, 0xcf, 0xe0, 0xa1, 0x7e, 0xa2, 0x59, 0x17, 0xd6, 0xae, 0x71, 0x7c, 0x23, 0x55, 0xa2,
	0xf9, 0x23, 0x57, 0x8b, 0xca, 0xa6, 0xb9, 0x97, 0x06, 0xf9, 0x81, 0x9f, 0x7b, 0x69, 0xd0, 0xf6,
	0x20, 0x4a, 0x87, 0x65, 0x16, 0x16, 0x22, 0xd6, 0xbc, 0xe7, 0xa6, 0x9c, 0x90, 0x33, 0x11, 0x6b,
	0x76, 0x08, 0xeb, 0x09, 0x9a, 0x48, 0xa4, 0xbe, 0xbc, 0x87, 0x55, 0xfd, 0x2c, 0xe6, 0x2a, 0x7c,
	0x00, 0xde, 0x0c, 0x6d, 0x55, 0x79, 0xe0, 0x42, 0x72, 0xd0, 0x1b, 0xd4, 0xf1, 0x94, 0xe0, 0xca,
	0x64, 0x29, 0x3f, 0x9a, 0x16, 0xbc, 0x33, 0x59, 0x6a, 0x17, 0x80, 0x17, 0x64, 0x72, 0x28, 0x52,
	0x74, 0xba, 0xff, 0x92, 0xae, 0xe3, 0x98, 0x1f, 0x89, 0x20, 0xf5, 0xd7, 0xb0, 0x7b, 0x3b, 0x7f,
	0xda, 0x44, 0xca, 0x84, 0x46, 0x64, 0xc8, 0x1f, 0x93, 0x9e, 0x4d, 0xb8, 0x73, 0x4b, 0x5d, 0x88,
	0x0c, 0xad, 0xff, 0xdb, 0x37, 0x30, 0x4f, 0x9c, 0xfe, 0x89, 0xf3, 0x3f, 0x61, 0x4e, 0xf3, 0x84,
	0xd4, 0x6e, 0x7d, 0x55, 0xf3, 0x8d, 0xca, 0x8f, 0xde, 0xd3, 0xc9, 0xfa, 0xf2, 0x33, 0x8e, 0xca,
	0x8d, 0x1f, 0x75, 0xae, 0xd2, 0x9b, 0x71, 0x81, 0xfc, 0xb8, 0xea, 0x9c, 0x47, 0x2f, 0xc6, 0x05,
	0x15, 0xda, 0x2d, 0x4c, 0x5a, 0x3d, 0xcf, 0x5c, 0xa1, 0x09, 0xf9, 0x60, 0xf7, 0xcf, 0x09, 0xec,
	0x2d, 0x2c, 0x4d, 0x52, 0xfe, 0x8f, 0x94, 0x3b, 0x73, 0x6b, 0x93, 0xde, 0x39, 0x84, 0xf5, 0x58,
	0x61, 0x64, 0xd0, 0x67, 0xf4, 0x7f, 0xd7, 0x1c, 0x8f, 0x51, 0x32, 0x07, 0xd0, 0x2e, 0x8b, 0xc4,
	0xae, 0x54, 0x52, 0x3c, 0x77, 0xb5, 0x77, 0x90, 0x15, 0x04, 0xbf, 0xd7, 0xa0, 0xe5, 0x77, 0xfc,
	0xa9, 0x66, 0x7b, 0xb0, 0x12, 0x8a, 0x3c, 0xc1, 0xdf, 0xfc, 0xa6, 0x6f, 0xf6, 0xad, 0xc1, 0x76,
	0xa0, 0xe9, 0x32, 0xab, 0xbb, 0xc9, 0xa1, 0x84, 0xb6, 0xa0, 0x51, 0xad, 0xf5, 0xd6, 0xa0, 0xde,
	0x4f, 0xe8, 0x65, 0x1d, 0x4b, 0x85, 0x7e, 0x77, 0x37, 0xcf, 0xad, 0xc1, 0x9e, 0xc3, 0x6a, 0xa8,
	0x65, 0xa9, 0x62, 0xa4, 0xd5, 0xdc, 0x3e, 0xd9, 0x79, 0x31, 0xf5, 0xd5, 0xf1, 0x67, 0x0f, 0x56,
	0xce, 0x49, 0x12, 0x7c, 0x0b, 0xcd, 0x53, 0xa5, 0xa4, 0xb2, 0xc3, 0x1a, 0xcb, 0x04, 0x29, 0x90,
	0xe6, 0x80, 0x9e, 0x19, 0x87, 0xd5, 0x0c, 0xb5, 0x8e, 0x46, 0x55, 0x24, 0x95, 0x19, 0x1c, 0xc1,
	0x76, 0x3f, 0x0f, 0xdf, 0xa2, 0xf1, 0xfe, 0xbe, 0xb7, 0x7b, 0x7e, 0xee, 0x9b, 0x15, 0xa4, 0xc0,
	0x7e, 0x2a, 0xcd, 0xbc, 0xea, 0x29, 0x34, 0xd1, 0x9e, 0x48, 0xc2, 0xf6, 0xc9, 0xf6, 0x74, 0x74,
	0x14, 0xca, 0xc0, 0xf1, 0xec, 0x2b, 0x58, 0xf5, 0x14, 0x9d, 0x7e, 0x47, 0x22, 0x95, 0x26, 0x90,
	0xb0, 0x67, 0x4f, 0xfb, 0x85, 0x6a, 0xed, 0xd9, 0x7e, 0x7e, 0x29, 0xbf, 0xd8, 0x81, 0xa7, 0xb0,
	0xdb, 0xcf, 0x3f, 0x71, 0xde, 0x94, 0x9b, 0xda, 0x67, 0xb8, 0xf9, 0xa3, 0x06, 0x9b, 0x33, 0xb5,
	0xd4, 0x76, 0xf7, 0xb9, 0x4b, 0x50, 0x73, 0x8d, 0x25, 0x83, 0x3e, 0x0f, 0xd1, 0x08, 0xb5, 0xff,
	0xfe, 0x3b, 0xc3, 0xce, 0xb9, 0xc6, 0x48, 0xc5, 0x57, 0xe1, 0x35, 0x8e, 0xfd, 0x74, 0xb4, 0x1c,
	0xf2, 0x03, 0x8e, 0x89, 0xbe, 0xbd, 0xb3, 0xcb, 0x9e, 0x9e, 0x5c, 0xd5, 0xfb, 0xb0, 0x36, 0xb9,
	0xa0, 0x4d, 0xd7, 0x62, 0x74, 0xf7, 0x32, 0xf8, 0xab, 0x06, 0x5b, 0xb3, 0xed, 0xd3, 0x9f, 0x5f,
	0xca, 0x49, 0x06, 0xf5, 0x4f, 0x66, 0xd0, 0x98, 0xce, 0x60, 0x17, 0x9a, 0x46, 0x9a, 0x28, 0xad,
	0xc6, 0x98, 0x0c, 0xf6, 0x12, 0xd6, 0x2b, 0xe7, 0xa9, 0xd0, 0x86, 0x37, 0x7b, 0x8d, 0xbb, 0x4a,
	0xd9, 0xf6, 0xd8, 0x7b, 0xa1, 0x4d, 0xf0, 0x92, 0x26, 0xf3, 0x0d, 0xfd, 0x2a, 0x98, 0xc4, 0x7d,
	0x38, 0xe7, 0xac, 0xd6, 0x6b, 0xd8, 0x9b, 0x3b, 0xfd, 0xde, 0x77, 0x6e, 0x58, 0xe7, 0x5e, 0xfc,
	0xdc, 0x84, 0x83, 0x57, 0xd0, 0xe9, 0xe7, 0xe1, 0x6b, 0x5a, 0x05, 0xd5, 0x6f, 0xb8, 0x7f, 0x39,
	0x08, 0xd7, 0xb0, 0x6d, 0x23, 0x98, 0xf5, 0xf1, 0x85, 0x86, 0xf7, 0xe4, 0xcf, 0x06, 0x74, 0x26,
	0x7b, 0xe8, 0x5d, 0x94, 0x27, 0x29, 0x2a, 0xf6, 0x33, 0x6c, 0xce, 0x5d, 0xd6, 0xfd, 0x69, 0x27,
	0x0b, 0x37, 0xbe, 0xfb, 0x68, 0x9a, 0x5e, 0xbc, 0xeb, 0xc1, 0x12, 0xfb, 0x15, 0xb6, 0x17, 0x6f,
	0x48, 0x6f, 0xce, 0xeb, 0x82, 0xa2, 0x7b, 0x38, 0xef, 0x78, 0x41, 0x12, 0x2c, 0xb1, 0x77, 0xd0,
	0x9e, 0x1e, 0xce, 0xee, 0x9d, 0xb1, 0xea, 0xee, 0x83, 0xbb, 0x03, 0xd5, 0xc1, 0x92, 0x4d, 0x7c,
	0xae, 0xf1, 0xf3, 0x89, 0xcf, 0xd2, 0x8b, 0x89, 0xcf, 0xf2, 0xc1, 0x12, 0xfb, 0x00, 0x1b, 0xb3,
	0x9d, 0x7c, 0x38, 0xe7, 0x71, 0x86, 0xed, 0xee, 0xcf, 0x3b, 0x9c, 0xa1, 0x83, 0xa5, 0xe1, 0x0a,
	0xfd, 0x5f, 0xf8, 0xe6, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x2f, 0x6b, 0x7b, 0x43, 0x0c,
	0x00, 0x00,
}