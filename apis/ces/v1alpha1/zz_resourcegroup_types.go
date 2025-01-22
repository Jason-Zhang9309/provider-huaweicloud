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

type ResourceGroupInitParameters struct {

	// Specifies the enterprise project IDs where the resources from.
	// It's required if the value of type is EPS.
	// Specifies the enterprise project IDs where the resources from.
	AssociatedEpsIds []*string `json:"associatedEpsIds,omitempty" tf:"associated_eps_ids,omitempty"`

	// Specifies the enterprise project ID of the resource group.
	// Specifies the enterprise project ID of the resource group.
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies the resource group name.
	// This parameter can contain a maximum of 128 characters, which may consist of letters,
	// digits, hyphens (-), underscore (_) and Chinese characters.
	// Specifies the resource group name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the list of resources to add into the group.
	// The ResourcesOpts structure is documented below.
	// Specifies the list of resources to add into the group.
	Resources []ResourceGroupResourcesInitParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// Specifies the key/value to match resources.
	// It's required if the value of type is TAG.
	// Specifies the key/value to match resources.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the resource group type.
	// The value can be EPS and TAG. If not specified, that means add resources manually.
	// Specifies the resource group type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResourceGroupObservation struct {

	// Specifies the enterprise project IDs where the resources from.
	// It's required if the value of type is EPS.
	// Specifies the enterprise project IDs where the resources from.
	AssociatedEpsIds []*string `json:"associatedEpsIds,omitempty" tf:"associated_eps_ids,omitempty"`

	// The creation time.
	// The creation time.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Specifies the enterprise project ID of the resource group.
	// Specifies the enterprise project ID of the resource group.
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the resource group name.
	// This parameter can contain a maximum of 128 characters, which may consist of letters,
	// digits, hyphens (-), underscore (_) and Chinese characters.
	// Specifies the resource group name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the list of resources to add into the group.
	// The ResourcesOpts structure is documented below.
	// Specifies the list of resources to add into the group.
	Resources []ResourceGroupResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	// Specifies the key/value to match resources.
	// It's required if the value of type is TAG.
	// Specifies the key/value to match resources.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the resource group type.
	// The value can be EPS and TAG. If not specified, that means add resources manually.
	// Specifies the resource group type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResourceGroupParameters struct {

	// Specifies the enterprise project IDs where the resources from.
	// It's required if the value of type is EPS.
	// Specifies the enterprise project IDs where the resources from.
	// +kubebuilder:validation:Optional
	AssociatedEpsIds []*string `json:"associatedEpsIds,omitempty" tf:"associated_eps_ids,omitempty"`

	// Specifies the enterprise project ID of the resource group.
	// Specifies the enterprise project ID of the resource group.
	// +kubebuilder:validation:Optional
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies the resource group name.
	// This parameter can contain a maximum of 128 characters, which may consist of letters,
	// digits, hyphens (-), underscore (_) and Chinese characters.
	// Specifies the resource group name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the list of resources to add into the group.
	// The ResourcesOpts structure is documented below.
	// Specifies the list of resources to add into the group.
	// +kubebuilder:validation:Optional
	Resources []ResourceGroupResourcesParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// Specifies the key/value to match resources.
	// It's required if the value of type is TAG.
	// Specifies the key/value to match resources.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the resource group type.
	// The value can be EPS and TAG. If not specified, that means add resources manually.
	// Specifies the resource group type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResourceGroupResourcesDimensionsInitParameters struct {

	// Specifies the resource group name.
	// This parameter can contain a maximum of 128 characters, which may consist of letters,
	// digits, hyphens (-), underscore (_) and Chinese characters.
	// Specifies the dimension name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the dimension value.
	// The value can be a string of 1 to 64 characters that must start with a letter or a number
	// and contain only letters, digits, underscores (_), and hyphens (-).
	// Specifies the dimension value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ResourceGroupResourcesDimensionsObservation struct {

	// Specifies the resource group name.
	// This parameter can contain a maximum of 128 characters, which may consist of letters,
	// digits, hyphens (-), underscore (_) and Chinese characters.
	// Specifies the dimension name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the dimension value.
	// The value can be a string of 1 to 64 characters that must start with a letter or a number
	// and contain only letters, digits, underscores (_), and hyphens (-).
	// Specifies the dimension value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ResourceGroupResourcesDimensionsParameters struct {

	// Specifies the resource group name.
	// This parameter can contain a maximum of 128 characters, which may consist of letters,
	// digits, hyphens (-), underscore (_) and Chinese characters.
	// Specifies the dimension name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the dimension value.
	// The value can be a string of 1 to 64 characters that must start with a letter or a number
	// and contain only letters, digits, underscores (_), and hyphens (-).
	// Specifies the dimension value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type ResourceGroupResourcesInitParameters struct {

	// Specifies the list of dimensions.
	// The DimensionOpts structure is documented below.
	// Specifies the list of dimensions.
	Dimensions []ResourceGroupResourcesDimensionsInitParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Specifies the namespace in service.item format.
	// service and item each must be a string that starts with a letter and contains only letters, digits, and
	// underscores (_). For details,
	// see Services Interconnected with Cloud Eye.
	// Specifies the namespace in **service.item** format.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type ResourceGroupResourcesObservation struct {

	// Specifies the list of dimensions.
	// The DimensionOpts structure is documented below.
	// Specifies the list of dimensions.
	Dimensions []ResourceGroupResourcesDimensionsObservation `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Specifies the namespace in service.item format.
	// service and item each must be a string that starts with a letter and contains only letters, digits, and
	// underscores (_). For details,
	// see Services Interconnected with Cloud Eye.
	// Specifies the namespace in **service.item** format.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type ResourceGroupResourcesParameters struct {

	// Specifies the list of dimensions.
	// The DimensionOpts structure is documented below.
	// Specifies the list of dimensions.
	// +kubebuilder:validation:Optional
	Dimensions []ResourceGroupResourcesDimensionsParameters `json:"dimensions" tf:"dimensions,omitempty"`

	// Specifies the namespace in service.item format.
	// service and item each must be a string that starts with a letter and contains only letters, digits, and
	// underscores (_). For details,
	// see Services Interconnected with Cloud Eye.
	// Specifies the namespace in **service.item** format.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

// ResourceGroupSpec defines the desired state of ResourceGroup
type ResourceGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceGroupParameters `json:"forProvider"`
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
	InitProvider ResourceGroupInitParameters `json:"initProvider,omitempty"`
}

// ResourceGroupStatus defines the observed state of ResourceGroup.
type ResourceGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ResourceGroup is the Schema for the ResourceGroups API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type ResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ResourceGroupSpec   `json:"spec"`
	Status ResourceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupList contains a list of ResourceGroups
type ResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroup `json:"items"`
}

// Repository type metadata.
var (
	ResourceGroup_Kind             = "ResourceGroup"
	ResourceGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceGroup_Kind}.String()
	ResourceGroup_KindAPIVersion   = ResourceGroup_Kind + "." + CRDGroupVersion.String()
	ResourceGroup_GroupVersionKind = CRDGroupVersion.WithKind(ResourceGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceGroup{}, &ResourceGroupList{})
}
