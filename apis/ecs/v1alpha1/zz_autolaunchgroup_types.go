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

type AutoLaunchGroupInitParameters struct {

	// Specifies the allocation strategy of the auto launch group.
	AllocationStrategy *string `json:"allocationStrategy,omitempty" tf:"allocation_strategy,omitempty"`

	// Specifies the interruption behavior of instances when deleting the auto launch
	// group.
	DeleteInstances *string `json:"deleteInstances,omitempty" tf:"delete_instances,omitempty"`

	// Specifies the interruption behavior of instances when target
	// capacity is exceeded or reduced. Valid values are terminate and noTermination. Default is terminate.
	ExcessFulfilledCapacityBehavior *string `json:"excessFulfilledCapacityBehavior,omitempty" tf:"excess_fulfilled_capacity_behavior,omitempty"`

	// Specifies the interruption behavior of running instances
	// when requests expire. Valid values are terminate and noTermination. Default is terminate.
	InstancesBehaviorWithExpiration *string `json:"instancesBehaviorWithExpiration,omitempty" tf:"instances_behavior_with_expiration,omitempty"`

	// Specifies the ID of launch template for instance.
	// Changing this creates a new resource.
	LaunchTemplateID *string `json:"launchTemplateId,omitempty" tf:"launch_template_id,omitempty"`

	// Specifies the version of launch template for instance.
	// Changing this creates a new resource.
	LaunchTemplateVersion *string `json:"launchTemplateVersion,omitempty" tf:"launch_template_version,omitempty"`

	// Specifies the name of the auto launch group. The valid length is limited
	// between 1 to 64, Only Chinese and English letters, digits, hyphens (-) and underscores (_) are allowed.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the instance details. Supporting mutiple overrides to create
	// instances of different specification. Changing this creates a new resource.
	// The overrides structure is documented below.
	Overrides []OverridesInitParameters `json:"overrides,omitempty" tf:"overrides,omitempty"`

	// Specifies the region in which to create the auto launch group.
	// If omitted, the provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the highest price a user is willing to pay per hour for a Spot instance.
	// If no price is provided in overrides, this price can be used.
	SpotPrice *float64 `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`

	// Specifies the target capacity to on-demand instance, the unit is the number of
	// vCPU, and the value must be less than or equal to target_capacity. There can be no on-demand instances in the auto
	// launch group.
	StableCapacity *float64 `json:"stableCapacity,omitempty" tf:"stable_capacity,omitempty"`

	// Specifies the selection strategies in resource supply.
	SupplyOption *string `json:"supplyOption,omitempty" tf:"supply_option,omitempty"`

	// Specifies the target capacity of the auto launch group, the unit is the number of
	// vCPU, and the value must be bigger than or equal to stable_capacity. The capacity of the spot instance euqals to
	// full capacity minus stable_capacity.
	TargetCapacity *float64 `json:"targetCapacity,omitempty" tf:"target_capacity,omitempty"`

	// Specifies the request type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the request start time, together with valid_since, determines
	// the validity period. In the format of yyyy-MM-dd'T'HH:mm:ssZ, an ISO 8601 time format. Default is starting now.
	// Changing this creates a new resource.
	ValidSince *string `json:"validSince,omitempty" tf:"valid_since,omitempty"`

	// Specifies the request end time, together with valid_since, determines
	// the validity period. In the format of yyyy-MM-dd'T'HH:mm:ssZ, an ISO 8601 time format. Default is never come to an
	// end. Changing this creates a new resource.
	ValidUntil *string `json:"validUntil,omitempty" tf:"valid_until,omitempty"`
}

type AutoLaunchGroupObservation struct {

	// Specifies the allocation strategy of the auto launch group.
	AllocationStrategy *string `json:"allocationStrategy,omitempty" tf:"allocation_strategy,omitempty"`

	// The create time of the auto launch group.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The total computing power that has been purchased successfully.
	CurrentCapacity *float64 `json:"currentCapacity,omitempty" tf:"current_capacity,omitempty"`

	// The on-demand computing power has been successfully purchased.
	CurrentStableCapacity *float64 `json:"currentStableCapacity,omitempty" tf:"current_stable_capacity,omitempty"`

	// Specifies the interruption behavior of instances when deleting the auto launch
	// group.
	DeleteInstances *string `json:"deleteInstances,omitempty" tf:"delete_instances,omitempty"`

	// Specifies the interruption behavior of instances when target
	// capacity is exceeded or reduced. Valid values are terminate and noTermination. Default is terminate.
	ExcessFulfilledCapacityBehavior *string `json:"excessFulfilledCapacityBehavior,omitempty" tf:"excess_fulfilled_capacity_behavior,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the interruption behavior of running instances
	// when requests expire. Valid values are terminate and noTermination. Default is terminate.
	InstancesBehaviorWithExpiration *string `json:"instancesBehaviorWithExpiration,omitempty" tf:"instances_behavior_with_expiration,omitempty"`

	// Specifies the ID of launch template for instance.
	// Changing this creates a new resource.
	LaunchTemplateID *string `json:"launchTemplateId,omitempty" tf:"launch_template_id,omitempty"`

	// Specifies the version of launch template for instance.
	// Changing this creates a new resource.
	LaunchTemplateVersion *string `json:"launchTemplateVersion,omitempty" tf:"launch_template_version,omitempty"`

	// Specifies the name of the auto launch group. The valid length is limited
	// between 1 to 64, Only Chinese and English letters, digits, hyphens (-) and underscores (_) are allowed.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the instance details. Supporting mutiple overrides to create
	// instances of different specification. Changing this creates a new resource.
	// The overrides structure is documented below.
	Overrides []OverridesObservation `json:"overrides,omitempty" tf:"overrides,omitempty"`

	// Specifies the region in which to create the auto launch group.
	// If omitted, the provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the highest price a user is willing to pay per hour for a Spot instance.
	// If no price is provided in overrides, this price can be used.
	SpotPrice *float64 `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`

	// Specifies the target capacity to on-demand instance, the unit is the number of
	// vCPU, and the value must be less than or equal to target_capacity. There can be no on-demand instances in the auto
	// launch group.
	StableCapacity *float64 `json:"stableCapacity,omitempty" tf:"stable_capacity,omitempty"`

	// The status of the auto launch group.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the selection strategies in resource supply.
	SupplyOption *string `json:"supplyOption,omitempty" tf:"supply_option,omitempty"`

	// Specifies the target capacity of the auto launch group, the unit is the number of
	// vCPU, and the value must be bigger than or equal to stable_capacity. The capacity of the spot instance euqals to
	// full capacity minus stable_capacity.
	TargetCapacity *float64 `json:"targetCapacity,omitempty" tf:"target_capacity,omitempty"`

	// The status of the auto launch group task. It can be:
	TaskState *string `json:"taskState,omitempty" tf:"task_state,omitempty"`

	// Specifies the request type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the request start time, together with valid_since, determines
	// the validity period. In the format of yyyy-MM-dd'T'HH:mm:ssZ, an ISO 8601 time format. Default is starting now.
	// Changing this creates a new resource.
	ValidSince *string `json:"validSince,omitempty" tf:"valid_since,omitempty"`

	// Specifies the request end time, together with valid_since, determines
	// the validity period. In the format of yyyy-MM-dd'T'HH:mm:ssZ, an ISO 8601 time format. Default is never come to an
	// end. Changing this creates a new resource.
	ValidUntil *string `json:"validUntil,omitempty" tf:"valid_until,omitempty"`
}

type AutoLaunchGroupParameters struct {

	// Specifies the allocation strategy of the auto launch group.
	// +kubebuilder:validation:Optional
	AllocationStrategy *string `json:"allocationStrategy,omitempty" tf:"allocation_strategy,omitempty"`

	// Specifies the interruption behavior of instances when deleting the auto launch
	// group.
	// +kubebuilder:validation:Optional
	DeleteInstances *string `json:"deleteInstances,omitempty" tf:"delete_instances,omitempty"`

	// Specifies the interruption behavior of instances when target
	// capacity is exceeded or reduced. Valid values are terminate and noTermination. Default is terminate.
	// +kubebuilder:validation:Optional
	ExcessFulfilledCapacityBehavior *string `json:"excessFulfilledCapacityBehavior,omitempty" tf:"excess_fulfilled_capacity_behavior,omitempty"`

	// Specifies the interruption behavior of running instances
	// when requests expire. Valid values are terminate and noTermination. Default is terminate.
	// +kubebuilder:validation:Optional
	InstancesBehaviorWithExpiration *string `json:"instancesBehaviorWithExpiration,omitempty" tf:"instances_behavior_with_expiration,omitempty"`

	// Specifies the ID of launch template for instance.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	LaunchTemplateID *string `json:"launchTemplateId,omitempty" tf:"launch_template_id,omitempty"`

	// Specifies the version of launch template for instance.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	LaunchTemplateVersion *string `json:"launchTemplateVersion,omitempty" tf:"launch_template_version,omitempty"`

	// Specifies the name of the auto launch group. The valid length is limited
	// between 1 to 64, Only Chinese and English letters, digits, hyphens (-) and underscores (_) are allowed.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the instance details. Supporting mutiple overrides to create
	// instances of different specification. Changing this creates a new resource.
	// The overrides structure is documented below.
	// +kubebuilder:validation:Optional
	Overrides []OverridesParameters `json:"overrides,omitempty" tf:"overrides,omitempty"`

	// Specifies the region in which to create the auto launch group.
	// If omitted, the provider-level region will be used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the highest price a user is willing to pay per hour for a Spot instance.
	// If no price is provided in overrides, this price can be used.
	// +kubebuilder:validation:Optional
	SpotPrice *float64 `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`

	// Specifies the target capacity to on-demand instance, the unit is the number of
	// vCPU, and the value must be less than or equal to target_capacity. There can be no on-demand instances in the auto
	// launch group.
	// +kubebuilder:validation:Optional
	StableCapacity *float64 `json:"stableCapacity,omitempty" tf:"stable_capacity,omitempty"`

	// Specifies the selection strategies in resource supply.
	// +kubebuilder:validation:Optional
	SupplyOption *string `json:"supplyOption,omitempty" tf:"supply_option,omitempty"`

	// Specifies the target capacity of the auto launch group, the unit is the number of
	// vCPU, and the value must be bigger than or equal to stable_capacity. The capacity of the spot instance euqals to
	// full capacity minus stable_capacity.
	// +kubebuilder:validation:Optional
	TargetCapacity *float64 `json:"targetCapacity,omitempty" tf:"target_capacity,omitempty"`

	// Specifies the request type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the request start time, together with valid_since, determines
	// the validity period. In the format of yyyy-MM-dd'T'HH:mm:ssZ, an ISO 8601 time format. Default is starting now.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	ValidSince *string `json:"validSince,omitempty" tf:"valid_since,omitempty"`

	// Specifies the request end time, together with valid_since, determines
	// the validity period. In the format of yyyy-MM-dd'T'HH:mm:ssZ, an ISO 8601 time format. Default is never come to an
	// end. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	ValidUntil *string `json:"validUntil,omitempty" tf:"valid_until,omitempty"`
}

type OverridesInitParameters struct {

	// Specifies the availability zone which the instance in.
	// Please refer to the document link reference for values.
	// Changing this creates a new resource.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Specifies the flavor ID of the instance. You can get available flavor id
	// through data source huaweicloud_compute_flavors.
	// Changing this creates a new resource.
	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	// Specifies the priority for launching. The smaller the value, the higher the
	// priority, and the launch will be given first. Valid value is from zero to max value of integer.
	// Changing this creates a new resource.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Specifies the highest price a user is willing to pay per hour for a Spot
	// instance. Changing this creates a new resource.
	SpotPrice *float64 `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`

	// Specifies the weight of the instance specification. The higher the
	// value, the greater the ability of a single instance to meet computing power requirements, and the smaller the number
	// of instances required. It must be bigger than zero. The weight value can be calculated based on the computing power of
	// the specified instance specification and the minimum computing power of a single node in the cluster.
	WeightedCapacity *float64 `json:"weightedCapacity,omitempty" tf:"weighted_capacity,omitempty"`
}

type OverridesObservation struct {

	// Specifies the availability zone which the instance in.
	// Please refer to the document link reference for values.
	// Changing this creates a new resource.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Specifies the flavor ID of the instance. You can get available flavor id
	// through data source huaweicloud_compute_flavors.
	// Changing this creates a new resource.
	FlavorID *string `json:"flavorId,omitempty" tf:"flavor_id,omitempty"`

	// Specifies the priority for launching. The smaller the value, the higher the
	// priority, and the launch will be given first. Valid value is from zero to max value of integer.
	// Changing this creates a new resource.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Specifies the highest price a user is willing to pay per hour for a Spot
	// instance. Changing this creates a new resource.
	SpotPrice *float64 `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`

	// Specifies the weight of the instance specification. The higher the
	// value, the greater the ability of a single instance to meet computing power requirements, and the smaller the number
	// of instances required. It must be bigger than zero. The weight value can be calculated based on the computing power of
	// the specified instance specification and the minimum computing power of a single node in the cluster.
	WeightedCapacity *float64 `json:"weightedCapacity,omitempty" tf:"weighted_capacity,omitempty"`
}

type OverridesParameters struct {

	// Specifies the availability zone which the instance in.
	// Please refer to the document link reference for values.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// Specifies the flavor ID of the instance. You can get available flavor id
	// through data source huaweicloud_compute_flavors.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	FlavorID *string `json:"flavorId" tf:"flavor_id,omitempty"`

	// Specifies the priority for launching. The smaller the value, the higher the
	// priority, and the launch will be given first. Valid value is from zero to max value of integer.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Specifies the highest price a user is willing to pay per hour for a Spot
	// instance. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	SpotPrice *float64 `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`

	// Specifies the weight of the instance specification. The higher the
	// value, the greater the ability of a single instance to meet computing power requirements, and the smaller the number
	// of instances required. It must be bigger than zero. The weight value can be calculated based on the computing power of
	// the specified instance specification and the minimum computing power of a single node in the cluster.
	// +kubebuilder:validation:Optional
	WeightedCapacity *float64 `json:"weightedCapacity,omitempty" tf:"weighted_capacity,omitempty"`
}

// AutoLaunchGroupSpec defines the desired state of AutoLaunchGroup
type AutoLaunchGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AutoLaunchGroupParameters `json:"forProvider"`
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
	InitProvider AutoLaunchGroupInitParameters `json:"initProvider,omitempty"`
}

// AutoLaunchGroupStatus defines the observed state of AutoLaunchGroup.
type AutoLaunchGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AutoLaunchGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AutoLaunchGroup is the Schema for the AutoLaunchGroups API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type AutoLaunchGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.launchTemplateId) || (has(self.initProvider) && has(self.initProvider.launchTemplateId))",message="spec.forProvider.launchTemplateId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.launchTemplateVersion) || (has(self.initProvider) && has(self.initProvider.launchTemplateVersion))",message="spec.forProvider.launchTemplateVersion is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.overrides) || (has(self.initProvider) && has(self.initProvider.overrides))",message="spec.forProvider.overrides is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetCapacity) || (has(self.initProvider) && has(self.initProvider.targetCapacity))",message="spec.forProvider.targetCapacity is a required parameter"
	Spec   AutoLaunchGroupSpec   `json:"spec"`
	Status AutoLaunchGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoLaunchGroupList contains a list of AutoLaunchGroups
type AutoLaunchGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoLaunchGroup `json:"items"`
}

// Repository type metadata.
var (
	AutoLaunchGroup_Kind             = "AutoLaunchGroup"
	AutoLaunchGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AutoLaunchGroup_Kind}.String()
	AutoLaunchGroup_KindAPIVersion   = AutoLaunchGroup_Kind + "." + CRDGroupVersion.String()
	AutoLaunchGroup_GroupVersionKind = CRDGroupVersion.WithKind(AutoLaunchGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&AutoLaunchGroup{}, &AutoLaunchGroupList{})
}
