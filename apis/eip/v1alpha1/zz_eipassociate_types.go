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

type EIPAssociateInitParameters_2 struct {

	// Specifies a private IP address to associate with the EIP.
	// Changing this creates a new resource.
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// Specifies the ID of the network to which the fixed_ip belongs.
	// It is mandatory when fixed_ip is set. Changing this creates a new resource.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Specifies an existing port ID to associate with the EIP.
	// This parameter and fixed_ip are alternative. Changing this creates a new resource.
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// Specifies the EIP address to associate. Changing this creates a new resource.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// Specifies the region in which to associate the EIP. If omitted, the provider-level
	// region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type EIPAssociateObservation_2 struct {

	// Specifies a private IP address to associate with the EIP.
	// Changing this creates a new resource.
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// The resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The MAC address of the private IP.
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	// Specifies the ID of the network to which the fixed_ip belongs.
	// It is mandatory when fixed_ip is set. Changing this creates a new resource.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Specifies an existing port ID to associate with the EIP.
	// This parameter and fixed_ip are alternative. Changing this creates a new resource.
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// Specifies the EIP address to associate. Changing this creates a new resource.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// The IPv6 address of the private IP.
	PublicIPv6 *string `json:"publicIpv6,omitempty" tf:"public_ipv6,omitempty"`

	// Specifies the region in which to associate the EIP. If omitted, the provider-level
	// region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The status of EIP, should be BOUND.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type EIPAssociateParameters_2 struct {

	// Specifies a private IP address to associate with the EIP.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// Specifies the ID of the network to which the fixed_ip belongs.
	// It is mandatory when fixed_ip is set. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Specifies an existing port ID to associate with the EIP.
	// This parameter and fixed_ip are alternative. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// Specifies the EIP address to associate. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// Specifies the region in which to associate the EIP. If omitted, the provider-level
	// region will be used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// EIPAssociateSpec defines the desired state of EIPAssociate
type EIPAssociateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EIPAssociateParameters_2 `json:"forProvider"`
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
	InitProvider EIPAssociateInitParameters_2 `json:"initProvider,omitempty"`
}

// EIPAssociateStatus defines the observed state of EIPAssociate.
type EIPAssociateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EIPAssociateObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EIPAssociate is the Schema for the EIPAssociates API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type EIPAssociate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.publicIp) || (has(self.initProvider) && has(self.initProvider.publicIp))",message="spec.forProvider.publicIp is a required parameter"
	Spec   EIPAssociateSpec   `json:"spec"`
	Status EIPAssociateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EIPAssociateList contains a list of EIPAssociates
type EIPAssociateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EIPAssociate `json:"items"`
}

// Repository type metadata.
var (
	EIPAssociate_Kind             = "EIPAssociate"
	EIPAssociate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EIPAssociate_Kind}.String()
	EIPAssociate_KindAPIVersion   = EIPAssociate_Kind + "." + CRDGroupVersion.String()
	EIPAssociate_GroupVersionKind = CRDGroupVersion.WithKind(EIPAssociate_Kind)
)

func init() {
	SchemeBuilder.Register(&EIPAssociate{}, &EIPAssociateList{})
}
