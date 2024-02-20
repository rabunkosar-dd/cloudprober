// Configuration proto for AWS provider.
// Example config:
// {
//
//   # EC2 instances
//   ec2_instances {}
//
//   # ElastiCache cluster
//   elasticache {}
//
//   # RDS clusters
//   rds {
//     identifier: "arn"
//   }
// }

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.5
// source: github.com/cloudprober/cloudprober/internal/rds/aws/proto/config.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EC2Instances struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How often resources should be refreshed.
	ReEvalSec *int32 `protobuf:"varint,98,opt,name=re_eval_sec,json=reEvalSec,def=600" json:"re_eval_sec,omitempty"` // default 10 mins
}

// Default values for EC2Instances fields.
const (
	Default_EC2Instances_ReEvalSec = int32(600)
)

func (x *EC2Instances) Reset() {
	*x = EC2Instances{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EC2Instances) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EC2Instances) ProtoMessage() {}

func (x *EC2Instances) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EC2Instances.ProtoReflect.Descriptor instead.
func (*EC2Instances) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *EC2Instances) GetReEvalSec() int32 {
	if x != nil && x.ReEvalSec != nil {
		return *x.ReEvalSec
	}
	return Default_EC2Instances_ReEvalSec
}

// ElastiCacheReplicationGroups discovery options.
type ElastiCacheReplicationGroups struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How often resources should be refreshed.
	ReEvalSec *int32 `protobuf:"varint,98,opt,name=re_eval_sec,json=reEvalSec,def=600" json:"re_eval_sec,omitempty"` // default 10 mins
}

// Default values for ElastiCacheReplicationGroups fields.
const (
	Default_ElastiCacheReplicationGroups_ReEvalSec = int32(600)
)

func (x *ElastiCacheReplicationGroups) Reset() {
	*x = ElastiCacheReplicationGroups{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElastiCacheReplicationGroups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElastiCacheReplicationGroups) ProtoMessage() {}

func (x *ElastiCacheReplicationGroups) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElastiCacheReplicationGroups.ProtoReflect.Descriptor instead.
func (*ElastiCacheReplicationGroups) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDescGZIP(), []int{1}
}

func (x *ElastiCacheReplicationGroups) GetReEvalSec() int32 {
	if x != nil && x.ReEvalSec != nil {
		return *x.ReEvalSec
	}
	return Default_ElastiCacheReplicationGroups_ReEvalSec
}

// ElastiCacheClusters discovery options.
type ElastiCacheClusters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How often resources should be refreshed.
	ReEvalSec *int32 `protobuf:"varint,98,opt,name=re_eval_sec,json=reEvalSec,def=600" json:"re_eval_sec,omitempty"` // default 10 mins
}

// Default values for ElastiCacheClusters fields.
const (
	Default_ElastiCacheClusters_ReEvalSec = int32(600)
)

func (x *ElastiCacheClusters) Reset() {
	*x = ElastiCacheClusters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElastiCacheClusters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElastiCacheClusters) ProtoMessage() {}

func (x *ElastiCacheClusters) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElastiCacheClusters.ProtoReflect.Descriptor instead.
func (*ElastiCacheClusters) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDescGZIP(), []int{2}
}

func (x *ElastiCacheClusters) GetReEvalSec() int32 {
	if x != nil && x.ReEvalSec != nil {
		return *x.ReEvalSec
	}
	return Default_ElastiCacheClusters_ReEvalSec
}

// RDS (Amazon Relational Databases) Clusters discovery options.
type RDSClusters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DB cluster identifier or the Amazon Resource Name (ARN) of the DB cluster
	// if specified, only the corresponding cluster information is returned.
	Identifier *string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	// Filters to be added to the discovery and search.
	Filter []string `protobuf:"bytes,2,rep,name=filter" json:"filter,omitempty"`
	// Whether to includes information about clusters shared from other AWS accounts.
	IncludeShared *bool  `protobuf:"varint,3,opt,name=include_shared,json=includeShared" json:"include_shared,omitempty"`
	ReEvalSec     *int32 `protobuf:"varint,98,opt,name=re_eval_sec,json=reEvalSec,def=600" json:"re_eval_sec,omitempty"` // default 10 mins
}

// Default values for RDSClusters fields.
const (
	Default_RDSClusters_ReEvalSec = int32(600)
)

func (x *RDSClusters) Reset() {
	*x = RDSClusters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RDSClusters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RDSClusters) ProtoMessage() {}

func (x *RDSClusters) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RDSClusters.ProtoReflect.Descriptor instead.
func (*RDSClusters) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDescGZIP(), []int{3}
}

func (x *RDSClusters) GetIdentifier() string {
	if x != nil && x.Identifier != nil {
		return *x.Identifier
	}
	return ""
}

func (x *RDSClusters) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *RDSClusters) GetIncludeShared() bool {
	if x != nil && x.IncludeShared != nil {
		return *x.IncludeShared
	}
	return false
}

func (x *RDSClusters) GetReEvalSec() int32 {
	if x != nil && x.ReEvalSec != nil {
		return *x.ReEvalSec
	}
	return Default_RDSClusters_ReEvalSec
}

// RDS (Amazon Relational Databases) Clusters discovery options.
type RDSInstances struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DB cluster identifier or the Amazon Resource Name (ARN) of the DB cluster
	// if specified, only the corresponding cluster information is returned.
	Identifier *string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	// Filters to be added to the discovery and search.
	Filter []string `protobuf:"bytes,2,rep,name=filter" json:"filter,omitempty"`
	// Whether to includes information about clusters shared from other AWS accounts.
	IncludeShared *bool  `protobuf:"varint,3,opt,name=include_shared,json=includeShared" json:"include_shared,omitempty"`
	ReEvalSec     *int32 `protobuf:"varint,98,opt,name=re_eval_sec,json=reEvalSec,def=600" json:"re_eval_sec,omitempty"` // default 10 mins
}

// Default values for RDSInstances fields.
const (
	Default_RDSInstances_ReEvalSec = int32(600)
)

func (x *RDSInstances) Reset() {
	*x = RDSInstances{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RDSInstances) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RDSInstances) ProtoMessage() {}

func (x *RDSInstances) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RDSInstances.ProtoReflect.Descriptor instead.
func (*RDSInstances) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDescGZIP(), []int{4}
}

func (x *RDSInstances) GetIdentifier() string {
	if x != nil && x.Identifier != nil {
		return *x.Identifier
	}
	return ""
}

func (x *RDSInstances) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *RDSInstances) GetIncludeShared() bool {
	if x != nil && x.IncludeShared != nil {
		return *x.IncludeShared
	}
	return false
}

func (x *RDSInstances) GetReEvalSec() int32 {
	if x != nil && x.ReEvalSec != nil {
		return *x.ReEvalSec
	}
	return Default_RDSInstances_ReEvalSec
}

// LoadBalancers discovery options.
type LoadBalancers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Amazon Resource Name (ARN) of the load balancer
	// if specified, only the corresponding load balancer information is returned.
	Name []string `protobuf:"bytes,1,rep,name=name" json:"name,omitempty"`
}

func (x *LoadBalancers) Reset() {
	*x = LoadBalancers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancers) ProtoMessage() {}

func (x *LoadBalancers) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancers.ProtoReflect.Descriptor instead.
func (*LoadBalancers) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDescGZIP(), []int{5}
}

func (x *LoadBalancers) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

// AWS provider config.
type ProviderConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Profile for the session.
	ProfileName *string `protobuf:"bytes,1,opt,name=profile_name,json=profileName" json:"profile_name,omitempty"`
	// AWS region
	Region *string `protobuf:"bytes,2,opt,name=region" json:"region,omitempty"`
	// ECS instances discovery options. This field should be declared for the AWS
	// instances discovery to be enabled.
	Ec2Instances *EC2Instances `protobuf:"bytes,3,opt,name=ec2_instances,json=ec2Instances" json:"ec2_instances,omitempty"`
	// ElastiCacheReplicationGroups discovery options. This field should be declared for the
	// elasticache replication groups discovery to be enabled.
	ElasticacheReplicationgroups *ElastiCacheReplicationGroups `protobuf:"bytes,4,opt,name=elasticache_replicationgroups,json=elasticacheReplicationgroups" json:"elasticache_replicationgroups,omitempty"`
	// ElastiCacheClusters discovery options. This field should be declared for the
	// elasticache cluster discovery to be enabled.
	ElasticacheClusters *ElastiCacheClusters `protobuf:"bytes,5,opt,name=elasticache_clusters,json=elasticacheClusters" json:"elasticache_clusters,omitempty"`
	// RDS instances discovery options.
	RdsInstances *RDSInstances `protobuf:"bytes,6,opt,name=rds_instances,json=rdsInstances" json:"rds_instances,omitempty"`
	// RDS clusters discovery options.
	RdsClusters *RDSClusters `protobuf:"bytes,7,opt,name=rds_clusters,json=rdsClusters" json:"rds_clusters,omitempty"`
}

func (x *ProviderConfig) Reset() {
	*x = ProviderConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderConfig) ProtoMessage() {}

func (x *ProviderConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderConfig.ProtoReflect.Descriptor instead.
func (*ProviderConfig) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDescGZIP(), []int{6}
}

func (x *ProviderConfig) GetProfileName() string {
	if x != nil && x.ProfileName != nil {
		return *x.ProfileName
	}
	return ""
}

func (x *ProviderConfig) GetRegion() string {
	if x != nil && x.Region != nil {
		return *x.Region
	}
	return ""
}

func (x *ProviderConfig) GetEc2Instances() *EC2Instances {
	if x != nil {
		return x.Ec2Instances
	}
	return nil
}

func (x *ProviderConfig) GetElasticacheReplicationgroups() *ElastiCacheReplicationGroups {
	if x != nil {
		return x.ElasticacheReplicationgroups
	}
	return nil
}

func (x *ProviderConfig) GetElasticacheClusters() *ElastiCacheClusters {
	if x != nil {
		return x.ElasticacheClusters
	}
	return nil
}

func (x *ProviderConfig) GetRdsInstances() *RDSInstances {
	if x != nil {
		return x.RdsInstances
	}
	return nil
}

func (x *ProviderConfig) GetRdsClusters() *RDSClusters {
	if x != nil {
		return x.RdsClusters
	}
	return nil
}

var File_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDesc = []byte{
	0x0a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x64,
	0x73, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x22, 0x33, 0x0a,
	0x0c, 0x45, 0x43, 0x32, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x23, 0x0a,
	0x0b, 0x72, 0x65, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x62, 0x20, 0x01,
	0x28, 0x05, 0x3a, 0x03, 0x36, 0x30, 0x30, 0x52, 0x09, 0x72, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x53,
	0x65, 0x63, 0x22, 0x43, 0x0a, 0x1c, 0x45, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x12, 0x23, 0x0a, 0x0b, 0x72, 0x65, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65,
	0x63, 0x18, 0x62, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x03, 0x36, 0x30, 0x30, 0x52, 0x09, 0x72, 0x65,
	0x45, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x22, 0x3a, 0x0a, 0x13, 0x45, 0x6c, 0x61, 0x73, 0x74,
	0x69, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x23,
	0x0a, 0x0b, 0x72, 0x65, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x62, 0x20,
	0x01, 0x28, 0x05, 0x3a, 0x03, 0x36, 0x30, 0x30, 0x52, 0x09, 0x72, 0x65, 0x45, 0x76, 0x61, 0x6c,
	0x53, 0x65, 0x63, 0x22, 0x91, 0x01, 0x0a, 0x0b, 0x52, 0x44, 0x53, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x69,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x12, 0x23, 0x0a, 0x0b, 0x72, 0x65, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65,
	0x63, 0x18, 0x62, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x03, 0x36, 0x30, 0x30, 0x52, 0x09, 0x72, 0x65,
	0x45, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x22, 0x92, 0x01, 0x0a, 0x0c, 0x52, 0x44, 0x53, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0b, 0x72, 0x65, 0x5f, 0x65, 0x76,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x62, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x03, 0x36, 0x30,
	0x30, 0x52, 0x09, 0x72, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x22, 0x23, 0x0a, 0x0d,
	0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0xf5, 0x03, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x46, 0x0a, 0x0d, 0x65, 0x63, 0x32, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x45, 0x43, 0x32,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x0c, 0x65, 0x63, 0x32, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x76, 0x0a, 0x1d, 0x65, 0x6c, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64, 0x73,
	0x2e, 0x61, 0x77, 0x73, 0x2e, 0x45, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x52, 0x1c, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12,
	0x5b, 0x0a, 0x14, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64, 0x73, 0x2e,
	0x61, 0x77, 0x73, 0x2e, 0x45, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x13, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x46, 0x0a, 0x0d,
	0x72, 0x64, 0x73, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x72, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x52, 0x44, 0x53, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x0c, 0x72, 0x64, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x0c, 0x72, 0x64, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x61, 0x77, 0x73,
	0x2e, 0x52, 0x44, 0x53, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0b, 0x72, 0x64,
	0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x64, 0x73, 0x2f, 0x61, 0x77, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDescData = file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_goTypes = []interface{}{
	(*EC2Instances)(nil),                 // 0: cloudprober.rds.aws.EC2Instances
	(*ElastiCacheReplicationGroups)(nil), // 1: cloudprober.rds.aws.ElastiCacheReplicationGroups
	(*ElastiCacheClusters)(nil),          // 2: cloudprober.rds.aws.ElastiCacheClusters
	(*RDSClusters)(nil),                  // 3: cloudprober.rds.aws.RDSClusters
	(*RDSInstances)(nil),                 // 4: cloudprober.rds.aws.RDSInstances
	(*LoadBalancers)(nil),                // 5: cloudprober.rds.aws.LoadBalancers
	(*ProviderConfig)(nil),               // 6: cloudprober.rds.aws.ProviderConfig
}
var file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_depIdxs = []int32{
	0, // 0: cloudprober.rds.aws.ProviderConfig.ec2_instances:type_name -> cloudprober.rds.aws.EC2Instances
	1, // 1: cloudprober.rds.aws.ProviderConfig.elasticache_replicationgroups:type_name -> cloudprober.rds.aws.ElastiCacheReplicationGroups
	2, // 2: cloudprober.rds.aws.ProviderConfig.elasticache_clusters:type_name -> cloudprober.rds.aws.ElastiCacheClusters
	4, // 3: cloudprober.rds.aws.ProviderConfig.rds_instances:type_name -> cloudprober.rds.aws.RDSInstances
	3, // 4: cloudprober.rds.aws.ProviderConfig.rds_clusters:type_name -> cloudprober.rds.aws.RDSClusters
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_init() }
func file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_init() {
	if File_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EC2Instances); i {
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
		file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElastiCacheReplicationGroups); i {
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
		file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElastiCacheClusters); i {
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
		file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RDSClusters); i {
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
		file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RDSInstances); i {
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
		file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancers); i {
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
		file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderConfig); i {
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
			RawDescriptor: file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_depIdxs,
		MessageInfos:      file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto = out.File
	file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_internal_rds_aws_proto_config_proto_depIdxs = nil
}