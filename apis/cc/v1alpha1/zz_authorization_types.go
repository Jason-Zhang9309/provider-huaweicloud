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

type AuthorizationInitParameters struct {

	// The peer account ID that you want to authorize.
	// The peer account ID that you want to authorize.
	CloudConnectionDomainID *string `json:"cloudConnectionDomainId,omitempty" tf:"cloud_connection_domain_id,omitempty"`

	// Peer cloud connection ID.
	// Peer cloud connection ID.
	CloudConnectionID *string `json:"cloudConnectionId,omitempty" tf:"cloud_connection_id,omitempty"`

	// The description of the cross-account authorization.
	// The description of the cross-account authorization.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The instance ID.
	// The instance ID.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The instance type.
	// The options are as follows:
	// The instance type.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The name of the cross-account authorization.
	// The name of the cross-account authorization.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type AuthorizationObservation struct {

	// The peer account ID that you want to authorize.
	// The peer account ID that you want to authorize.
	CloudConnectionDomainID *string `json:"cloudConnectionDomainId,omitempty" tf:"cloud_connection_domain_id,omitempty"`

	// Peer cloud connection ID.
	// Peer cloud connection ID.
	CloudConnectionID *string `json:"cloudConnectionId,omitempty" tf:"cloud_connection_id,omitempty"`

	// The description of the cross-account authorization.
	// The description of the cross-account authorization.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The instance ID.
	// The instance ID.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The instance type.
	// The options are as follows:
	// The instance type.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The name of the cross-account authorization.
	// The name of the cross-account authorization.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type AuthorizationParameters struct {

	// The peer account ID that you want to authorize.
	// The peer account ID that you want to authorize.
	// +kubebuilder:validation:Optional
	CloudConnectionDomainID *string `json:"cloudConnectionDomainId,omitempty" tf:"cloud_connection_domain_id,omitempty"`

	// Peer cloud connection ID.
	// Peer cloud connection ID.
	// +kubebuilder:validation:Optional
	CloudConnectionID *string `json:"cloudConnectionId,omitempty" tf:"cloud_connection_id,omitempty"`

	// The description of the cross-account authorization.
	// The description of the cross-account authorization.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The instance ID.
	// The instance ID.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The instance type.
	// The options are as follows:
	// The instance type.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// The name of the cross-account authorization.
	// The name of the cross-account authorization.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// AuthorizationSpec defines the desired state of Authorization
type AuthorizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthorizationParameters `json:"forProvider"`
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
	InitProvider AuthorizationInitParameters `json:"initProvider,omitempty"`
}

// AuthorizationStatus defines the observed state of Authorization.
type AuthorizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthorizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Authorization is the Schema for the Authorizations API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type Authorization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cloudConnectionDomainId) || (has(self.initProvider) && has(self.initProvider.cloudConnectionDomainId))",message="spec.forProvider.cloudConnectionDomainId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cloudConnectionId) || (has(self.initProvider) && has(self.initProvider.cloudConnectionId))",message="spec.forProvider.cloudConnectionId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceId) || (has(self.initProvider) && has(self.initProvider.instanceId))",message="spec.forProvider.instanceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceType) || (has(self.initProvider) && has(self.initProvider.instanceType))",message="spec.forProvider.instanceType is a required parameter"
	Spec   AuthorizationSpec   `json:"spec"`
	Status AuthorizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthorizationList contains a list of Authorizations
type AuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Authorization `json:"items"`
}

// Repository type metadata.
var (
	Authorization_Kind             = "Authorization"
	Authorization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Authorization_Kind}.String()
	Authorization_KindAPIVersion   = Authorization_Kind + "." + CRDGroupVersion.String()
	Authorization_GroupVersionKind = CRDGroupVersion.WithKind(Authorization_Kind)
)

func init() {
	SchemeBuilder.Register(&Authorization{}, &AuthorizationList{})
}
