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

type RabbitmqVhostInitParameters struct {

	// Specifies the DMS RabbitMQ instance ID.
	// Changing this creates a new resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/dms/v1alpha1.RabbitmqInstance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a RabbitmqInstance in dms to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a RabbitmqInstance in dms to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the vhost name. Changing this creates a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type RabbitmqVhostObservation struct {

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the DMS RabbitMQ instance ID.
	// Changing this creates a new resource.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Specifies the vhost name. Changing this creates a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Indicates whether the message tracing is enabled.
	Tracing *bool `json:"tracing,omitempty" tf:"tracing,omitempty"`
}

type RabbitmqVhostParameters struct {

	// Specifies the DMS RabbitMQ instance ID.
	// Changing this creates a new resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/dms/v1alpha1.RabbitmqInstance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a RabbitmqInstance in dms to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a RabbitmqInstance in dms to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the vhost name. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// RabbitmqVhostSpec defines the desired state of RabbitmqVhost
type RabbitmqVhostSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RabbitmqVhostParameters `json:"forProvider"`
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
	InitProvider RabbitmqVhostInitParameters `json:"initProvider,omitempty"`
}

// RabbitmqVhostStatus defines the observed state of RabbitmqVhost.
type RabbitmqVhostStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RabbitmqVhostObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RabbitmqVhost is the Schema for the RabbitmqVhosts API. Manages a DMS RabbitMQ vhost resource within HuaweiCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type RabbitmqVhost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   RabbitmqVhostSpec   `json:"spec"`
	Status RabbitmqVhostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RabbitmqVhostList contains a list of RabbitmqVhosts
type RabbitmqVhostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RabbitmqVhost `json:"items"`
}

// Repository type metadata.
var (
	RabbitmqVhost_Kind             = "RabbitmqVhost"
	RabbitmqVhost_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RabbitmqVhost_Kind}.String()
	RabbitmqVhost_KindAPIVersion   = RabbitmqVhost_Kind + "." + CRDGroupVersion.String()
	RabbitmqVhost_GroupVersionKind = CRDGroupVersion.WithKind(RabbitmqVhost_Kind)
)

func init() {
	SchemeBuilder.Register(&RabbitmqVhost{}, &RabbitmqVhostList{})
}