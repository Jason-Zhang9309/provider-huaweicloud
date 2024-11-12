// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RabbitmqInstanceInitParameters struct {

	// Specifies a username. A username consists of 4 to 64 characters and
	// supports only letters, digits, and hyphens (-). Changing this creates a new instance resource.
	AccessUser *string `json:"accessUser,omitempty" tf:"access_user,omitempty"`

	// Specifies whether auto renew is enabled. Valid values are true and false.
	AutoRenew *string `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// Specifies the names of an AZ.
	// The parameter value can not be left blank or an empty array.
	// Changing this creates a new instance resource.
	// schema: Required
	// +listType=set
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	AvailableZones []*string `json:"availableZones,omitempty" tf:"available_zones,omitempty"`

	// Specifies the broker numbers.
	// It is required when creating a cluster instance with flavor_id.
	BrokerNum *float64 `json:"brokerNum,omitempty" tf:"broker_num,omitempty"`

	// Specifies the charging mode of the instance. Valid values are
	// prePaid and postPaid, defaults to postPaid. Changing this creates a new resource.
	ChargingMode *string `json:"chargingMode,omitempty" tf:"charging_mode,omitempty"`

	// Specifies the description of the DMS RabbitMQ instance.
	// It is a character string containing not more than 1,024 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to enable ACL. Only available when engine_version is AMQP-0-9-1.
	// Default to false.
	EnableACL *bool `json:"enableAcl,omitempty" tf:"enable_acl,omitempty"`

	// Specifies the version of the RabbitMQ engine. Default to "3.7.17".
	// Changing this creates a new instance resource.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Specifies the enterprise project ID of the RabbitMQ instance.
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies a flavor ID.
	// It is mandatory when the charging_mode is prePaid.
	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	// Specifies the time at which a maintenance time window starts. Format: HH:mm.
	// The start time and end time of a maintenance time window must indicate the time segment of a supported maintenance
	// time window.
	// The start time must be set to 22:00, 02:00, 06:00, 10:00, 14:00, or 18:00. Parameters maintain_begin
	// and maintain_end must be set in pairs. If parameter maintain_begin is left blank, parameter maintain_end is also
	// blank. In this case, the system automatically allocates the default start time 02:00.
	MaintainBegin *string `json:"maintainBegin,omitempty" tf:"maintain_begin,omitempty"`

	// Specifies the time at which a maintenance time window ends. Format: HH:mm.
	// The start time and end time of a maintenance time window must indicate the time segment of a supported maintenance
	// time window. The end time is four hours later than the start time.
	// For example, if the start time is 22:00, the end time is 02:00.
	// Parameters maintain_begin and maintain_end must be set in pairs.
	// If parameter maintain_end is left  blank, parameter maintain_begin is also blank.
	// In this case, the system automatically allocates the default end time 06:00.
	MaintainEnd *string `json:"maintainEnd,omitempty" tf:"maintain_end,omitempty"`

	// Specifies the name of the DMS RabbitMQ instance. An instance name starts with a letter,
	// consists of 4 to 64 characters, and supports only letters, digits, hyphens (-) and underscores (_).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the ID of a subnet. Changing this creates a new instance
	// resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Subnet in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Specifies the password of the DMS RabbitMQ instance. A password must meet
	// the following complexity requirements: Must be 8 to 32 characters long. Must contain at least 2 of the following
	// character types: lowercase letters, uppercase letters, digits,
	// and special characters (`~!@#$%^&*()-_=+\|[{}]:'",<.>/?).
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Specifies the charging period of the instance. If period_unit is set to
	// month, the value ranges from 1 to 9. If period_unit is set to year, the value ranges from 1 to 3.
	// This parameter is mandatory if charging_mode is set to prePaid. Changing this creates a new resource.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Specifies the charging period unit of the instance.
	// Valid values are month and year. This parameter is mandatory if charging_mode is set to prePaid.
	// Changing this creates a new resource.
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit,omitempty"`

	// Specifies a resource ID in UUID format.
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// Specifies the ID of the elastic IP address (EIP)
	// bound to the DMS RabbitMQ instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/eip/v1alpha1.VpcEip
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	PublicIPID *string `json:"publicIpId,omitempty" tf:"public_ip_id,omitempty"`

	// Reference to a VpcEip in eip to populate publicIpId.
	// +kubebuilder:validation:Optional
	PublicIPIDRef *v1.Reference `json:"publicIpIdRef,omitempty" tf:"-"`

	// Selector for a VpcEip in eip to populate publicIpId.
	// +kubebuilder:validation:Optional
	PublicIPIDSelector *v1.Selector `json:"publicIpIdSelector,omitempty" tf:"-"`

	// The region in which to create the DMS RabbitMQ instance resource. If omitted,
	// the provider-level region will be used. Changing this creates a new instance resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies whether to enable public access for the DMS RabbitMQ instance.
	// Changing this creates a new instance resource.
	SSLEnable *bool `json:"sslEnable,omitempty" tf:"ssl_enable,omitempty"`

	// Specifies the ID of a security group.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.Secgroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a Secgroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a Secgroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// , you need to manually modify the value of storage_space after changing the broker_num.
	StorageSpace *float64 `json:"storageSpace,omitempty" tf:"storage_space,omitempty"`

	// Specifies the storage I/O specification.
	// Valid values are dms.physical.storage.high.v2 and dms.physical.storage.ultra.v2.
	// Changing this creates a new instance resource.
	StorageSpecCode *string `json:"storageSpecCode,omitempty" tf:"storage_spec_code,omitempty"`

	// The key/value pairs to associate with the DMS RabbitMQ instance.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ID of a VPC. Changing this creates a new instance resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.VPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type RabbitmqInstanceObservation struct {

	// Specifies a username. A username consists of 4 to 64 characters and
	// supports only letters, digits, and hyphens (-). Changing this creates a new instance resource.
	AccessUser *string `json:"accessUser,omitempty" tf:"access_user,omitempty"`

	// Specifies whether auto renew is enabled. Valid values are true and false.
	AutoRenew *string `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// Specifies the names of an AZ.
	// The parameter value can not be left blank or an empty array.
	// Changing this creates a new instance resource.
	// schema: Required
	// +listType=set
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	AvailableZones []*string `json:"availableZones,omitempty" tf:"available_zones,omitempty"`

	// Specifies the broker numbers.
	// It is required when creating a cluster instance with flavor_id.
	BrokerNum *float64 `json:"brokerNum,omitempty" tf:"broker_num,omitempty"`

	// Specifies the charging mode of the instance. Valid values are
	// prePaid and postPaid, defaults to postPaid. Changing this creates a new resource.
	ChargingMode *string `json:"chargingMode,omitempty" tf:"charging_mode,omitempty"`

	// Indicates the IP address of the DMS RabbitMQ instance.
	ConnectAddress *string `json:"connectAddress,omitempty" tf:"connect_address,omitempty"`

	// Indicates the create time of the DMS RabbitMQ instance.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Specifies the description of the DMS RabbitMQ instance.
	// It is a character string containing not more than 1,024 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to enable ACL. Only available when engine_version is AMQP-0-9-1.
	// Default to false.
	EnableACL *bool `json:"enableAcl,omitempty" tf:"enable_acl,omitempty"`

	// Indicates whether public access to the DMS RabbitMQ instance is enabled.
	EnablePublicIP *bool `json:"enablePublicIp,omitempty" tf:"enable_public_ip,omitempty"`

	// Indicates the message engine.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Specifies the version of the RabbitMQ engine. Default to "3.7.17".
	// Changing this creates a new instance resource.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Specifies the enterprise project ID of the RabbitMQ instance.
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Indicates the extend times of the DMS RabbitMQ instance.
	ExtendTimes *float64 `json:"extendTimes,omitempty" tf:"extend_times,omitempty"`

	// Specifies a flavor ID.
	// It is mandatory when the charging_mode is prePaid.
	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	// Specifies a resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates whether the DMS RabbitMQ instance is logical volume.
	IsLogicalVolume *bool `json:"isLogicalVolume,omitempty" tf:"is_logical_volume,omitempty"`

	// Specifies the time at which a maintenance time window starts. Format: HH:mm.
	// The start time and end time of a maintenance time window must indicate the time segment of a supported maintenance
	// time window.
	// The start time must be set to 22:00, 02:00, 06:00, 10:00, 14:00, or 18:00. Parameters maintain_begin
	// and maintain_end must be set in pairs. If parameter maintain_begin is left blank, parameter maintain_end is also
	// blank. In this case, the system automatically allocates the default start time 02:00.
	MaintainBegin *string `json:"maintainBegin,omitempty" tf:"maintain_begin,omitempty"`

	// Specifies the time at which a maintenance time window ends. Format: HH:mm.
	// The start time and end time of a maintenance time window must indicate the time segment of a supported maintenance
	// time window. The end time is four hours later than the start time.
	// For example, if the start time is 22:00, the end time is 02:00.
	// Parameters maintain_begin and maintain_end must be set in pairs.
	// If parameter maintain_end is left  blank, parameter maintain_begin is also blank.
	// In this case, the system automatically allocates the default end time 06:00.
	MaintainEnd *string `json:"maintainEnd,omitempty" tf:"maintain_end,omitempty"`

	// Indicates the management address of the DMS RabbitMQ instance.
	ManagementConnectAddress *string `json:"managementConnectAddress,omitempty" tf:"management_connect_address,omitempty"`

	// Indicates the IP address of the DMS RabbitMQ instance.
	ManegementConnectAddress *string `json:"manegementConnectAddress,omitempty" tf:"manegement_connect_address,omitempty"`

	// Specifies the name of the DMS RabbitMQ instance. An instance name starts with a letter,
	// consists of 4 to 64 characters, and supports only letters, digits, hyphens (-) and underscores (_).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the ID of a subnet. Changing this creates a new instance
	// resource.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Specifies the charging period of the instance. If period_unit is set to
	// month, the value ranges from 1 to 9. If period_unit is set to year, the value ranges from 1 to 3.
	// This parameter is mandatory if charging_mode is set to prePaid. Changing this creates a new resource.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Specifies the charging period unit of the instance.
	// Valid values are month and year. This parameter is mandatory if charging_mode is set to prePaid.
	// Changing this creates a new resource.
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit,omitempty"`

	// Indicates the port number of the DMS RabbitMQ instance.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Specifies a resource ID in UUID format.
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// Indicates the public ip address of the DMS RabbitMQ instance.
	PublicIPAddress *string `json:"publicIpAddress,omitempty" tf:"public_ip_address,omitempty"`

	// Specifies the ID of the elastic IP address (EIP)
	// bound to the DMS RabbitMQ instance.
	PublicIPID *string `json:"publicIpId,omitempty" tf:"public_ip_id,omitempty"`

	// The region in which to create the DMS RabbitMQ instance resource. If omitted,
	// the provider-level region will be used. Changing this creates a new instance resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Indicates a resource specifications identifier.
	ResourceSpecCode *string `json:"resourceSpecCode,omitempty" tf:"resource_spec_code,omitempty"`

	// Specifies whether to enable public access for the DMS RabbitMQ instance.
	// Changing this creates a new instance resource.
	SSLEnable *bool `json:"sslEnable,omitempty" tf:"ssl_enable,omitempty"`

	// Specifies the ID of a security group.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Indicates the instance specification. For a single-node DMS RabbitMQ instance, VM specifications are
	// returned. For a cluster DMS RabbitMQ instance, VM specifications and the number of nodes are returned.
	Specification *string `json:"specification,omitempty" tf:"specification,omitempty"`

	// Indicates the status of the DMS RabbitMQ instance.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// , you need to manually modify the value of storage_space after changing the broker_num.
	StorageSpace *float64 `json:"storageSpace,omitempty" tf:"storage_space,omitempty"`

	// Specifies the storage I/O specification.
	// Valid values are dms.physical.storage.high.v2 and dms.physical.storage.ultra.v2.
	// Changing this creates a new instance resource.
	StorageSpecCode *string `json:"storageSpecCode,omitempty" tf:"storage_spec_code,omitempty"`

	// The key/value pairs to associate with the DMS RabbitMQ instance.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Indicates the DMS RabbitMQ instance type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Indicates the used message storage space. Unit: GB
	UsedStorageSpace *float64 `json:"usedStorageSpace,omitempty" tf:"used_storage_space,omitempty"`

	// Indicates the ID of the user who created the DMS RabbitMQ instance
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Indicates the name of the user who created the DMS RabbitMQ instance
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// Specifies the ID of a VPC. Changing this creates a new instance resource.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type RabbitmqInstanceParameters struct {

	// Specifies a username. A username consists of 4 to 64 characters and
	// supports only letters, digits, and hyphens (-). Changing this creates a new instance resource.
	// +kubebuilder:validation:Optional
	AccessUser *string `json:"accessUser,omitempty" tf:"access_user,omitempty"`

	// Specifies whether auto renew is enabled. Valid values are true and false.
	// +kubebuilder:validation:Optional
	AutoRenew *string `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// Specifies the names of an AZ.
	// The parameter value can not be left blank or an empty array.
	// Changing this creates a new instance resource.
	// schema: Required
	// +kubebuilder:validation:Optional
	// +listType=set
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// +kubebuilder:validation:Optional
	AvailableZones []*string `json:"availableZones,omitempty" tf:"available_zones,omitempty"`

	// Specifies the broker numbers.
	// It is required when creating a cluster instance with flavor_id.
	// +kubebuilder:validation:Optional
	BrokerNum *float64 `json:"brokerNum,omitempty" tf:"broker_num,omitempty"`

	// Specifies the charging mode of the instance. Valid values are
	// prePaid and postPaid, defaults to postPaid. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	ChargingMode *string `json:"chargingMode,omitempty" tf:"charging_mode,omitempty"`

	// Specifies the description of the DMS RabbitMQ instance.
	// It is a character string containing not more than 1,024 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to enable ACL. Only available when engine_version is AMQP-0-9-1.
	// Default to false.
	// +kubebuilder:validation:Optional
	EnableACL *bool `json:"enableAcl,omitempty" tf:"enable_acl,omitempty"`

	// Specifies the version of the RabbitMQ engine. Default to "3.7.17".
	// Changing this creates a new instance resource.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Specifies the enterprise project ID of the RabbitMQ instance.
	// +kubebuilder:validation:Optional
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies a flavor ID.
	// It is mandatory when the charging_mode is prePaid.
	// +kubebuilder:validation:Optional
	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	// Specifies the time at which a maintenance time window starts. Format: HH:mm.
	// The start time and end time of a maintenance time window must indicate the time segment of a supported maintenance
	// time window.
	// The start time must be set to 22:00, 02:00, 06:00, 10:00, 14:00, or 18:00. Parameters maintain_begin
	// and maintain_end must be set in pairs. If parameter maintain_begin is left blank, parameter maintain_end is also
	// blank. In this case, the system automatically allocates the default start time 02:00.
	// +kubebuilder:validation:Optional
	MaintainBegin *string `json:"maintainBegin,omitempty" tf:"maintain_begin,omitempty"`

	// Specifies the time at which a maintenance time window ends. Format: HH:mm.
	// The start time and end time of a maintenance time window must indicate the time segment of a supported maintenance
	// time window. The end time is four hours later than the start time.
	// For example, if the start time is 22:00, the end time is 02:00.
	// Parameters maintain_begin and maintain_end must be set in pairs.
	// If parameter maintain_end is left  blank, parameter maintain_begin is also blank.
	// In this case, the system automatically allocates the default end time 06:00.
	// +kubebuilder:validation:Optional
	MaintainEnd *string `json:"maintainEnd,omitempty" tf:"maintain_end,omitempty"`

	// Specifies the name of the DMS RabbitMQ instance. An instance name starts with a letter,
	// consists of 4 to 64 characters, and supports only letters, digits, hyphens (-) and underscores (_).
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the ID of a subnet. Changing this creates a new instance
	// resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Subnet in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Specifies the password of the DMS RabbitMQ instance. A password must meet
	// the following complexity requirements: Must be 8 to 32 characters long. Must contain at least 2 of the following
	// character types: lowercase letters, uppercase letters, digits,
	// and special characters (`~!@#$%^&*()-_=+\|[{}]:'",<.>/?).
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Specifies the charging period of the instance. If period_unit is set to
	// month, the value ranges from 1 to 9. If period_unit is set to year, the value ranges from 1 to 3.
	// This parameter is mandatory if charging_mode is set to prePaid. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Specifies the charging period unit of the instance.
	// Valid values are month and year. This parameter is mandatory if charging_mode is set to prePaid.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit,omitempty"`

	// Specifies a resource ID in UUID format.
	// +kubebuilder:validation:Optional
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// Specifies the ID of the elastic IP address (EIP)
	// bound to the DMS RabbitMQ instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/eip/v1alpha1.VpcEip
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PublicIPID *string `json:"publicIpId,omitempty" tf:"public_ip_id,omitempty"`

	// Reference to a VpcEip in eip to populate publicIpId.
	// +kubebuilder:validation:Optional
	PublicIPIDRef *v1.Reference `json:"publicIpIdRef,omitempty" tf:"-"`

	// Selector for a VpcEip in eip to populate publicIpId.
	// +kubebuilder:validation:Optional
	PublicIPIDSelector *v1.Selector `json:"publicIpIdSelector,omitempty" tf:"-"`

	// The region in which to create the DMS RabbitMQ instance resource. If omitted,
	// the provider-level region will be used. Changing this creates a new instance resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies whether to enable public access for the DMS RabbitMQ instance.
	// Changing this creates a new instance resource.
	// +kubebuilder:validation:Optional
	SSLEnable *bool `json:"sslEnable,omitempty" tf:"ssl_enable,omitempty"`

	// Specifies the ID of a security group.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.Secgroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a Secgroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a Secgroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// , you need to manually modify the value of storage_space after changing the broker_num.
	// +kubebuilder:validation:Optional
	StorageSpace *float64 `json:"storageSpace,omitempty" tf:"storage_space,omitempty"`

	// Specifies the storage I/O specification.
	// Valid values are dms.physical.storage.high.v2 and dms.physical.storage.ultra.v2.
	// Changing this creates a new instance resource.
	// +kubebuilder:validation:Optional
	StorageSpecCode *string `json:"storageSpecCode,omitempty" tf:"storage_spec_code,omitempty"`

	// The key/value pairs to associate with the DMS RabbitMQ instance.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ID of a VPC. Changing this creates a new instance resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.VPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// RabbitmqInstanceSpec defines the desired state of RabbitmqInstance
type RabbitmqInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RabbitmqInstanceParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RabbitmqInstanceInitParameters `json:"initProvider,omitempty"`
}

// RabbitmqInstanceStatus defines the observed state of RabbitmqInstance.
type RabbitmqInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RabbitmqInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RabbitmqInstance is the Schema for the RabbitmqInstances API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type RabbitmqInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageSpecCode) || (has(self.initProvider) && has(self.initProvider.storageSpecCode))",message="spec.forProvider.storageSpecCode is a required parameter"
	Spec   RabbitmqInstanceSpec   `json:"spec"`
	Status RabbitmqInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RabbitmqInstanceList contains a list of RabbitmqInstances
type RabbitmqInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RabbitmqInstance `json:"items"`
}

// Repository type metadata.
var (
	RabbitmqInstance_Kind             = "RabbitmqInstance"
	RabbitmqInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RabbitmqInstance_Kind}.String()
	RabbitmqInstance_KindAPIVersion   = RabbitmqInstance_Kind + "." + CRDGroupVersion.String()
	RabbitmqInstance_GroupVersionKind = CRDGroupVersion.WithKind(RabbitmqInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&RabbitmqInstance{}, &RabbitmqInstanceList{})
}
