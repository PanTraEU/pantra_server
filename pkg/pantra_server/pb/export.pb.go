// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: internal/pb/export/export.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Data type representing why this key was published.
type TemporaryExposureKey_ReportType int32

const (
	TemporaryExposureKey_UNKNOWN                      TemporaryExposureKey_ReportType = 0 // Never returned by the client API.
	TemporaryExposureKey_CONFIRMED_TEST               TemporaryExposureKey_ReportType = 1
	TemporaryExposureKey_CONFIRMED_CLINICAL_DIAGNOSIS TemporaryExposureKey_ReportType = 2
	TemporaryExposureKey_SELF_REPORT                  TemporaryExposureKey_ReportType = 3
	TemporaryExposureKey_RECURSIVE                    TemporaryExposureKey_ReportType = 4
	TemporaryExposureKey_REVOKED                      TemporaryExposureKey_ReportType = 5 // Used to revoke a key, never returned by client API.
)

// Enum value maps for TemporaryExposureKey_ReportType.
var (
	TemporaryExposureKey_ReportType_name = map[int32]string{
		0: "UNKNOWN",
		1: "CONFIRMED_TEST",
		2: "CONFIRMED_CLINICAL_DIAGNOSIS",
		3: "SELF_REPORT",
		4: "RECURSIVE",
		5: "REVOKED",
	}
	TemporaryExposureKey_ReportType_value = map[string]int32{
		"UNKNOWN":                      0,
		"CONFIRMED_TEST":               1,
		"CONFIRMED_CLINICAL_DIAGNOSIS": 2,
		"SELF_REPORT":                  3,
		"RECURSIVE":                    4,
		"REVOKED":                      5,
	}
)

func (x TemporaryExposureKey_ReportType) Enum() *TemporaryExposureKey_ReportType {
	p := new(TemporaryExposureKey_ReportType)
	*p = x
	return p
}

func (x TemporaryExposureKey_ReportType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TemporaryExposureKey_ReportType) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_pb_export_export_proto_enumTypes[0].Descriptor()
}

func (TemporaryExposureKey_ReportType) Type() protoreflect.EnumType {
	return &file_internal_pb_export_export_proto_enumTypes[0]
}

func (x TemporaryExposureKey_ReportType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *TemporaryExposureKey_ReportType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = TemporaryExposureKey_ReportType(num)
	return nil
}

// Deprecated: Use TemporaryExposureKey_ReportType.Descriptor instead.
func (TemporaryExposureKey_ReportType) EnumDescriptor() ([]byte, []int) {
	return file_internal_pb_export_export_proto_rawDescGZIP(), []int{2, 0}
}

// Protobuf definition for exports of confirmed temporary exposure keys.
//
// The full file format is documented under "Exposure Key Export File Format
// and Verification" at https://www.google.com/covid19/exposurenotifications/
//
// These files have a 16-byte, space-padded header before the protobuf data
// starts. They will be contained in a zip archive, alongside a signature
// file verifying the contents.
type TemporaryExposureKeyExport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time window of keys in this file based on arrival to server, in UTC
	// seconds
	StartTimestamp *uint64 `protobuf:"fixed64,1,opt,name=start_timestamp,json=startTimestamp" json:"start_timestamp,omitempty"`
	EndTimestamp   *uint64 `protobuf:"fixed64,2,opt,name=end_timestamp,json=endTimestamp" json:"end_timestamp,omitempty"`
	// Region for which these keys came from (e.g., country)
	Region *string `protobuf:"bytes,3,opt,name=region" json:"region,omitempty"`
	// E.g., Batch 2 of 10. Ordinal, 1-based numbering.
	// Note: Not yet supported on iOS. Use values of 1 for both.
	BatchNum  *int32 `protobuf:"varint,4,opt,name=batch_num,json=batchNum" json:"batch_num,omitempty"`
	BatchSize *int32 `protobuf:"varint,5,opt,name=batch_size,json=batchSize" json:"batch_size,omitempty"`
	// Information about signatures
	// If there are multiple entries, they must be ordered in descending
	// time order by signing key effective time (most recent one first).
	// There is a limit of 10 signature infos per export file (mobile OS may
	// not check anything after that).
	SignatureInfos []*SignatureInfo `protobuf:"bytes,6,rep,name=signature_infos,json=signatureInfos" json:"signature_infos,omitempty"`
	// The TemporaryExposureKeys for initial release of keys.
	// Keys should be included in this list for initial release,
	// whereas revised or revoked keys should go in revised_keys.
	Keys []*TemporaryExposureKey `protobuf:"bytes,7,rep,name=keys" json:"keys,omitempty"`
	// TemporaryExposureKeys that have changed status.
	// Keys should be included in this list if they have changed status
	// or have been revoked.
	RevisedKeys []*TemporaryExposureKey `protobuf:"bytes,8,rep,name=revised_keys,json=revisedKeys" json:"revised_keys,omitempty"`
}

func (x *TemporaryExposureKeyExport) Reset() {
	*x = TemporaryExposureKeyExport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_export_export_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemporaryExposureKeyExport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemporaryExposureKeyExport) ProtoMessage() {}

func (x *TemporaryExposureKeyExport) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_export_export_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemporaryExposureKeyExport.ProtoReflect.Descriptor instead.
func (*TemporaryExposureKeyExport) Descriptor() ([]byte, []int) {
	return file_internal_pb_export_export_proto_rawDescGZIP(), []int{0}
}

func (x *TemporaryExposureKeyExport) GetStartTimestamp() uint64 {
	if x != nil && x.StartTimestamp != nil {
		return *x.StartTimestamp
	}
	return 0
}

func (x *TemporaryExposureKeyExport) GetEndTimestamp() uint64 {
	if x != nil && x.EndTimestamp != nil {
		return *x.EndTimestamp
	}
	return 0
}

func (x *TemporaryExposureKeyExport) GetRegion() string {
	if x != nil && x.Region != nil {
		return *x.Region
	}
	return ""
}

func (x *TemporaryExposureKeyExport) GetBatchNum() int32 {
	if x != nil && x.BatchNum != nil {
		return *x.BatchNum
	}
	return 0
}

func (x *TemporaryExposureKeyExport) GetBatchSize() int32 {
	if x != nil && x.BatchSize != nil {
		return *x.BatchSize
	}
	return 0
}

func (x *TemporaryExposureKeyExport) GetSignatureInfos() []*SignatureInfo {
	if x != nil {
		return x.SignatureInfos
	}
	return nil
}

func (x *TemporaryExposureKeyExport) GetKeys() []*TemporaryExposureKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *TemporaryExposureKeyExport) GetRevisedKeys() []*TemporaryExposureKey {
	if x != nil {
		return x.RevisedKeys
	}
	return nil
}

type SignatureInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key version for rollovers
	// Must be in character class [a-zA-Z0-9_]. E.g., 'v1'
	VerificationKeyVersion *string `protobuf:"bytes,3,opt,name=verification_key_version,json=verificationKeyVersion" json:"verification_key_version,omitempty"`
	// Alias with which to identify public key to be used for verification
	// Must be in character class [a-zA-Z0-9_]
	// For cross-compatibility with Apple, use MCC
	// (https://en.wikipedia.org/wiki/Mobile_country_code).
	VerificationKeyId *string `protobuf:"bytes,4,opt,name=verification_key_id,json=verificationKeyId" json:"verification_key_id,omitempty"`
	// ASN.1 OID for Algorithm Identifier. Supported algorithms are
	// either 1.2.840.10045.4.3.2 or 1.2.840.10045.4.3.4
	SignatureAlgorithm *string `protobuf:"bytes,5,opt,name=signature_algorithm,json=signatureAlgorithm" json:"signature_algorithm,omitempty"`
}

func (x *SignatureInfo) Reset() {
	*x = SignatureInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_export_export_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureInfo) ProtoMessage() {}

func (x *SignatureInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_export_export_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignatureInfo.ProtoReflect.Descriptor instead.
func (*SignatureInfo) Descriptor() ([]byte, []int) {
	return file_internal_pb_export_export_proto_rawDescGZIP(), []int{1}
}

func (x *SignatureInfo) GetVerificationKeyVersion() string {
	if x != nil && x.VerificationKeyVersion != nil {
		return *x.VerificationKeyVersion
	}
	return ""
}

func (x *SignatureInfo) GetVerificationKeyId() string {
	if x != nil && x.VerificationKeyId != nil {
		return *x.VerificationKeyId
	}
	return ""
}

func (x *SignatureInfo) GetSignatureAlgorithm() string {
	if x != nil && x.SignatureAlgorithm != nil {
		return *x.SignatureAlgorithm
	}
	return ""
}

type TemporaryExposureKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key of infected user
	KeyData []byte `protobuf:"bytes,1,opt,name=key_data,json=keyData" json:"key_data,omitempty"`
	// Varying risks associated with exposure depending on type of verification
	// Ignored by the v1.5 client API when report_type is set.
	//
	// Deprecated: Do not use.
	TransmissionRiskLevel *int32 `protobuf:"varint,2,opt,name=transmission_risk_level,json=transmissionRiskLevel" json:"transmission_risk_level,omitempty"`
	// The interval number since epoch for which a key starts
	RollingStartIntervalNumber *int32 `protobuf:"varint,3,opt,name=rolling_start_interval_number,json=rollingStartIntervalNumber" json:"rolling_start_interval_number,omitempty"`
	// Increments of 10 minutes describing how long a key is valid
	RollingPeriod *int32 `protobuf:"varint,4,opt,name=rolling_period,json=rollingPeriod,def=144" json:"rolling_period,omitempty"` // defaults to 24 hours
	// Type of diagnosis associated with a key.
	ReportType *TemporaryExposureKey_ReportType `protobuf:"varint,5,opt,name=report_type,json=reportType,enum=TemporaryExposureKey_ReportType" json:"report_type,omitempty"`
	// Number of days elapsed between symptom onset and the TEK being used.
	// E.g. 2 means TEK is 2 days after onset of symptoms.
	DaysSinceOnsetOfSymptoms *int32 `protobuf:"zigzag32,6,opt,name=days_since_onset_of_symptoms,json=daysSinceOnsetOfSymptoms" json:"days_since_onset_of_symptoms,omitempty"`
}

// Default values for TemporaryExposureKey fields.
const (
	Default_TemporaryExposureKey_RollingPeriod = int32(144)
)

func (x *TemporaryExposureKey) Reset() {
	*x = TemporaryExposureKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_export_export_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemporaryExposureKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemporaryExposureKey) ProtoMessage() {}

func (x *TemporaryExposureKey) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_export_export_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemporaryExposureKey.ProtoReflect.Descriptor instead.
func (*TemporaryExposureKey) Descriptor() ([]byte, []int) {
	return file_internal_pb_export_export_proto_rawDescGZIP(), []int{2}
}

func (x *TemporaryExposureKey) GetKeyData() []byte {
	if x != nil {
		return x.KeyData
	}
	return nil
}

// Deprecated: Do not use.
func (x *TemporaryExposureKey) GetTransmissionRiskLevel() int32 {
	if x != nil && x.TransmissionRiskLevel != nil {
		return *x.TransmissionRiskLevel
	}
	return 0
}

func (x *TemporaryExposureKey) GetRollingStartIntervalNumber() int32 {
	if x != nil && x.RollingStartIntervalNumber != nil {
		return *x.RollingStartIntervalNumber
	}
	return 0
}

func (x *TemporaryExposureKey) GetRollingPeriod() int32 {
	if x != nil && x.RollingPeriod != nil {
		return *x.RollingPeriod
	}
	return Default_TemporaryExposureKey_RollingPeriod
}

func (x *TemporaryExposureKey) GetReportType() TemporaryExposureKey_ReportType {
	if x != nil && x.ReportType != nil {
		return *x.ReportType
	}
	return TemporaryExposureKey_UNKNOWN
}

func (x *TemporaryExposureKey) GetDaysSinceOnsetOfSymptoms() int32 {
	if x != nil && x.DaysSinceOnsetOfSymptoms != nil {
		return *x.DaysSinceOnsetOfSymptoms
	}
	return 0
}

type TEKSignatureList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// When there are multiple signatures, they must be sorted in time order
	// by first effective date for the signing key in descending order.
	// The most recent effective signing key must appear first.
	// There is a limit of 10 signature infos per export file (mobile OS may
	// not check anything after that).
	Signatures []*TEKSignature `protobuf:"bytes,1,rep,name=signatures" json:"signatures,omitempty"`
}

func (x *TEKSignatureList) Reset() {
	*x = TEKSignatureList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_export_export_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TEKSignatureList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TEKSignatureList) ProtoMessage() {}

func (x *TEKSignatureList) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_export_export_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TEKSignatureList.ProtoReflect.Descriptor instead.
func (*TEKSignatureList) Descriptor() ([]byte, []int) {
	return file_internal_pb_export_export_proto_rawDescGZIP(), []int{3}
}

func (x *TEKSignatureList) GetSignatures() []*TEKSignature {
	if x != nil {
		return x.Signatures
	}
	return nil
}

type TEKSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Info about the signing key, version, algorithm, etc.
	SignatureInfo *SignatureInfo `protobuf:"bytes,1,opt,name=signature_info,json=signatureInfo" json:"signature_info,omitempty"`
	// E.g., Batch 2 of 10
	BatchNum  *int32 `protobuf:"varint,2,opt,name=batch_num,json=batchNum" json:"batch_num,omitempty"`
	BatchSize *int32 `protobuf:"varint,3,opt,name=batch_size,json=batchSize" json:"batch_size,omitempty"`
	// Signature in X9.62 format (ASN.1 SEQUENCE of two INTEGER fields)
	Signature []byte `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
}

func (x *TEKSignature) Reset() {
	*x = TEKSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_export_export_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TEKSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TEKSignature) ProtoMessage() {}

func (x *TEKSignature) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_export_export_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TEKSignature.ProtoReflect.Descriptor instead.
func (*TEKSignature) Descriptor() ([]byte, []int) {
	return file_internal_pb_export_export_proto_rawDescGZIP(), []int{4}
}

func (x *TEKSignature) GetSignatureInfo() *SignatureInfo {
	if x != nil {
		return x.SignatureInfo
	}
	return nil
}

func (x *TEKSignature) GetBatchNum() int32 {
	if x != nil && x.BatchNum != nil {
		return *x.BatchNum
	}
	return 0
}

func (x *TEKSignature) GetBatchSize() int32 {
	if x != nil && x.BatchSize != nil {
		return *x.BatchSize
	}
	return 0
}

func (x *TEKSignature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_internal_pb_export_export_proto protoreflect.FileDescriptor

var file_internal_pb_export_export_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdc, 0x02, 0x0a, 0x1a, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x45,
	0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06,
	0x52, 0x0c, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x37, 0x0a, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x6b,
	0x65, 0x79, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x54, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x72, 0x79, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79,
	0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x38, 0x0a, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x73, 0x65,
	0x64, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x54,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65,
	0x4b, 0x65, 0x79, 0x52, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x73, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73,
	0x22, 0xd6, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x38, 0x0a, 0x18, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x4a, 0x04, 0x08,
	0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x5f, 0x62,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x52, 0x0f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69,
	0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0xdd, 0x03, 0x0a, 0x14, 0x54, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a,
	0x17, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x69,
	0x73, 0x6b, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x69, 0x73, 0x6b, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x41, 0x0a, 0x1d, 0x72, 0x6f, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x1a, 0x72, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x0e,
	0x72, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x3a, 0x03, 0x31, 0x34, 0x34, 0x52, 0x0d, 0x72, 0x6f, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x41, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72,
	0x65, 0x4b, 0x65, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x1c, 0x64,
	0x61, 0x79, 0x73, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x6f, 0x6e, 0x73, 0x65, 0x74, 0x5f,
	0x6f, 0x66, 0x5f, 0x73, 0x79, 0x6d, 0x70, 0x74, 0x6f, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x18, 0x64, 0x61, 0x79, 0x73, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x4f, 0x6e, 0x73, 0x65,
	0x74, 0x4f, 0x66, 0x53, 0x79, 0x6d, 0x70, 0x74, 0x6f, 0x6d, 0x73, 0x22, 0x7c, 0x0a, 0x0a, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52,
	0x4d, 0x45, 0x44, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x4f,
	0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x5f, 0x43, 0x4c, 0x49, 0x4e, 0x49, 0x43, 0x41, 0x4c,
	0x5f, 0x44, 0x49, 0x41, 0x47, 0x4e, 0x4f, 0x53, 0x49, 0x53, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b,
	0x53, 0x45, 0x4c, 0x46, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a,
	0x09, 0x52, 0x45, 0x43, 0x55, 0x52, 0x53, 0x49, 0x56, 0x45, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07,
	0x52, 0x45, 0x56, 0x4f, 0x4b, 0x45, 0x44, 0x10, 0x05, 0x22, 0x41, 0x0a, 0x10, 0x54, 0x45, 0x4b,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x54, 0x45, 0x4b, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x9f, 0x01, 0x0a,
	0x0c, 0x54, 0x45, 0x4b, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x35, 0x0a,
	0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x75,
	0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x4b,
	0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x2d, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x3b, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74,
}

var (
	file_internal_pb_export_export_proto_rawDescOnce sync.Once
	file_internal_pb_export_export_proto_rawDescData = file_internal_pb_export_export_proto_rawDesc
)

func file_internal_pb_export_export_proto_rawDescGZIP() []byte {
	file_internal_pb_export_export_proto_rawDescOnce.Do(func() {
		file_internal_pb_export_export_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pb_export_export_proto_rawDescData)
	})
	return file_internal_pb_export_export_proto_rawDescData
}

var file_internal_pb_export_export_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_pb_export_export_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_pb_export_export_proto_goTypes = []interface{}{
	(TemporaryExposureKey_ReportType)(0), // 0: TemporaryExposureKey.ReportType
	(*TemporaryExposureKeyExport)(nil),   // 1: TemporaryExposureKeyExport
	(*SignatureInfo)(nil),                // 2: SignatureInfo
	(*TemporaryExposureKey)(nil),         // 3: TemporaryExposureKey
	(*TEKSignatureList)(nil),             // 4: TEKSignatureList
	(*TEKSignature)(nil),                 // 5: TEKSignature
}
var file_internal_pb_export_export_proto_depIdxs = []int32{
	2, // 0: TemporaryExposureKeyExport.signature_infos:type_name -> SignatureInfo
	3, // 1: TemporaryExposureKeyExport.keys:type_name -> TemporaryExposureKey
	3, // 2: TemporaryExposureKeyExport.revised_keys:type_name -> TemporaryExposureKey
	0, // 3: TemporaryExposureKey.report_type:type_name -> TemporaryExposureKey.ReportType
	5, // 4: TEKSignatureList.signatures:type_name -> TEKSignature
	2, // 5: TEKSignature.signature_info:type_name -> SignatureInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_internal_pb_export_export_proto_init() }
func file_internal_pb_export_export_proto_init() {
	if File_internal_pb_export_export_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pb_export_export_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemporaryExposureKeyExport); i {
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
		file_internal_pb_export_export_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignatureInfo); i {
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
		file_internal_pb_export_export_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemporaryExposureKey); i {
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
		file_internal_pb_export_export_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TEKSignatureList); i {
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
		file_internal_pb_export_export_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TEKSignature); i {
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
			RawDescriptor: file_internal_pb_export_export_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pb_export_export_proto_goTypes,
		DependencyIndexes: file_internal_pb_export_export_proto_depIdxs,
		EnumInfos:         file_internal_pb_export_export_proto_enumTypes,
		MessageInfos:      file_internal_pb_export_export_proto_msgTypes,
	}.Build()
	File_internal_pb_export_export_proto = out.File
	file_internal_pb_export_export_proto_rawDesc = nil
	file_internal_pb_export_export_proto_goTypes = nil
	file_internal_pb_export_export_proto_depIdxs = nil
}
