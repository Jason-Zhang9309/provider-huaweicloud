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

type AddressGroupInitParameters struct {

	// Specifies an array of one or more IP addresses. The address can be a single IP
	// address, IP address range or IP address CIDR. The maximum length is 20.
	// +listType=set
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// Specifies the supplementary information about the IP address group.
	// The value is a string of no more than 255 characters and cannot contain angle brackets (< or >).
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the enterprise project ID.
	// Changing this creates a new address group.
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies whether to forcibly destroy the address group if it is associated with
	// a security group rule, the address group and the associated security group rule will be deleted together.
	// The default value is false.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Specifies the IP version, either 4 (default) or 6.
	// Changing this creates a new address group.
	IPVersion *float64 `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Specifies the maximum number of addresses that an address group can contain.
	// The valid value is range from 1 to 20, the default value is 20.
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// Specifies the IP address group name.
	// The value is a string of 1 to 64 characters that can contain letters, digits, underscores (_), hyphens (-) and
	// periods (.).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the IP address group. If omitted, the
	// provider-level region will be used. Changing this creates a new address group.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type AddressGroupObservation struct {

	// Specifies an array of one or more IP addresses. The address can be a single IP
	// address, IP address range or IP address CIDR. The maximum length is 20.
	// +listType=set
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// Specifies the supplementary information about the IP address group.
	// The value is a string of no more than 255 characters and cannot contain angle brackets (< or >).
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the enterprise project ID.
	// Changing this creates a new address group.
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies whether to forcibly destroy the address group if it is associated with
	// a security group rule, the address group and the associated security group rule will be deleted together.
	// The default value is false.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the IP version, either 4 (default) or 6.
	// Changing this creates a new address group.
	IPVersion *float64 `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Specifies the maximum number of addresses that an address group can contain.
	// The valid value is range from 1 to 20, the default value is 20.
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// Specifies the IP address group name.
	// The value is a string of 1 to 64 characters that can contain letters, digits, underscores (_), hyphens (-) and
	// periods (.).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the IP address group. If omitted, the
	// provider-level region will be used. Changing this creates a new address group.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type AddressGroupParameters struct {

	// Specifies an array of one or more IP addresses. The address can be a single IP
	// address, IP address range or IP address CIDR. The maximum length is 20.
	// +kubebuilder:validation:Optional
	// +listType=set
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// Specifies the supplementary information about the IP address group.
	// The value is a string of no more than 255 characters and cannot contain angle brackets (< or >).
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the enterprise project ID.
	// Changing this creates a new address group.
	// +kubebuilder:validation:Optional
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies whether to forcibly destroy the address group if it is associated with
	// a security group rule, the address group and the associated security group rule will be deleted together.
	// The default value is false.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Specifies the IP version, either 4 (default) or 6.
	// Changing this creates a new address group.
	// +kubebuilder:validation:Optional
	IPVersion *float64 `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Specifies the maximum number of addresses that an address group can contain.
	// The valid value is range from 1 to 20, the default value is 20.
	// +kubebuilder:validation:Optional
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// Specifies the IP address group name.
	// The value is a string of 1 to 64 characters that can contain letters, digits, underscores (_), hyphens (-) and
	// periods (.).
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the IP address group. If omitted, the
	// provider-level region will be used. Changing this creates a new address group.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// AddressGroupSpec defines the desired state of AddressGroup
type AddressGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AddressGroupParameters `json:"forProvider"`
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
	InitProvider AddressGroupInitParameters `json:"initProvider,omitempty"`
}

// AddressGroupStatus defines the observed state of AddressGroup.
type AddressGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AddressGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AddressGroup is the Schema for the AddressGroups API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type AddressGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.addresses) || (has(self.initProvider) && has(self.initProvider.addresses))",message="spec.forProvider.addresses is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   AddressGroupSpec   `json:"spec"`
	Status AddressGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddressGroupList contains a list of AddressGroups
type AddressGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AddressGroup `json:"items"`
}

// Repository type metadata.
var (
	AddressGroup_Kind             = "AddressGroup"
	AddressGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AddressGroup_Kind}.String()
	AddressGroup_KindAPIVersion   = AddressGroup_Kind + "." + CRDGroupVersion.String()
	AddressGroup_GroupVersionKind = CRDGroupVersion.WithKind(AddressGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&AddressGroup{}, &AddressGroupList{})
}