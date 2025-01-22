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

type DetailDimensionsInitParameters struct {

	// Specifies the event name.
	// The resource dimension name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The resource dimension value.
	// The resource dimension value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DetailDimensionsObservation struct {

	// Specifies the event name.
	// The resource dimension name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The resource dimension value.
	// The resource dimension value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DetailDimensionsParameters struct {

	// Specifies the event name.
	// The resource dimension name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The resource dimension value.
	// The resource dimension value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type DetailInitParameters struct {

	// Specifies the event content.
	// Specifies the event content.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Specifies the resource dimensions.
	// The dimensions structure is documented below.
	// Specifies the resource dimensions.
	Dimensions []DetailDimensionsInitParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Specifies the group that the event belongs to.
	// Specifies the group that the event belongs to.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Specifies the event level.
	// The value can be Critical, Major, Minor, or Info.
	// Specifies the event level.
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// Specifies the resource ID.
	// Specifies the resource ID.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Specifies the resource name.
	// Specifies the resource name.
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`

	// Specifies the event status.
	// The value can be normal, warning, or incident.
	// Specifies the event status.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Specifies the event type.
	// The value can only be EVENT.CUSTOM.
	// Specifies the event type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the event user.
	// Specifies the event user.
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type DetailObservation struct {

	// Specifies the event content.
	// Specifies the event content.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Specifies the resource dimensions.
	// The dimensions structure is documented below.
	// Specifies the resource dimensions.
	Dimensions []DetailDimensionsObservation `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Specifies the group that the event belongs to.
	// Specifies the group that the event belongs to.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Specifies the event level.
	// The value can be Critical, Major, Minor, or Info.
	// Specifies the event level.
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// Specifies the resource ID.
	// Specifies the resource ID.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Specifies the resource name.
	// Specifies the resource name.
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`

	// Specifies the event status.
	// The value can be normal, warning, or incident.
	// Specifies the event status.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Specifies the event type.
	// The value can only be EVENT.CUSTOM.
	// Specifies the event type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the event user.
	// Specifies the event user.
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type DetailParameters struct {

	// Specifies the event content.
	// Specifies the event content.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Specifies the resource dimensions.
	// The dimensions structure is documented below.
	// Specifies the resource dimensions.
	// +kubebuilder:validation:Optional
	Dimensions []DetailDimensionsParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Specifies the group that the event belongs to.
	// Specifies the group that the event belongs to.
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Specifies the event level.
	// The value can be Critical, Major, Minor, or Info.
	// Specifies the event level.
	// +kubebuilder:validation:Optional
	Level *string `json:"level" tf:"level,omitempty"`

	// Specifies the resource ID.
	// Specifies the resource ID.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Specifies the resource name.
	// Specifies the resource name.
	// +kubebuilder:validation:Optional
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`

	// Specifies the event status.
	// The value can be normal, warning, or incident.
	// Specifies the event status.
	// +kubebuilder:validation:Optional
	State *string `json:"state" tf:"state,omitempty"`

	// Specifies the event type.
	// The value can only be EVENT.CUSTOM.
	// Specifies the event type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the event user.
	// Specifies the event user.
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type EventReportInitParameters struct {

	// Specifies the detail of the CES event.
	// The detail structure is documented below.
	// Specifies the detail of the CES event.
	Detail []DetailInitParameters `json:"detail,omitempty" tf:"detail,omitempty"`

	// Specifies the event name.
	// Specifies the CES event name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the event source.
	// Specifies the event source.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies the occurrence time of the event.
	// The time is in UTC. The format is yyyy-MM-dd HH:mm:ss.
	// Specifies the occurrence time of the event.
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type EventReportObservation struct {

	// Specifies the detail of the CES event.
	// The detail structure is documented below.
	// Specifies the detail of the CES event.
	Detail []DetailObservation `json:"detail,omitempty" tf:"detail,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the event name.
	// Specifies the CES event name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the event source.
	// Specifies the event source.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies the occurrence time of the event.
	// The time is in UTC. The format is yyyy-MM-dd HH:mm:ss.
	// Specifies the occurrence time of the event.
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type EventReportParameters struct {

	// Specifies the detail of the CES event.
	// The detail structure is documented below.
	// Specifies the detail of the CES event.
	// +kubebuilder:validation:Optional
	Detail []DetailParameters `json:"detail,omitempty" tf:"detail,omitempty"`

	// Specifies the event name.
	// Specifies the CES event name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the event source.
	// Specifies the event source.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies the occurrence time of the event.
	// The time is in UTC. The format is yyyy-MM-dd HH:mm:ss.
	// Specifies the occurrence time of the event.
	// +kubebuilder:validation:Optional
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

// EventReportSpec defines the desired state of EventReport
type EventReportSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventReportParameters `json:"forProvider"`
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
	InitProvider EventReportInitParameters `json:"initProvider,omitempty"`
}

// EventReportStatus defines the observed state of EventReport.
type EventReportStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventReportObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EventReport is the Schema for the EventReports API. Manages a CES event report resource within HuaweiCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type EventReport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.detail) || (has(self.initProvider) && has(self.initProvider.detail))",message="spec.forProvider.detail is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.source) || (has(self.initProvider) && has(self.initProvider.source))",message="spec.forProvider.source is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.time) || (has(self.initProvider) && has(self.initProvider.time))",message="spec.forProvider.time is a required parameter"
	Spec   EventReportSpec   `json:"spec"`
	Status EventReportStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventReportList contains a list of EventReports
type EventReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventReport `json:"items"`
}

// Repository type metadata.
var (
	EventReport_Kind             = "EventReport"
	EventReport_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventReport_Kind}.String()
	EventReport_KindAPIVersion   = EventReport_Kind + "." + CRDGroupVersion.String()
	EventReport_GroupVersionKind = CRDGroupVersion.WithKind(EventReport_Kind)
)

func init() {
	SchemeBuilder.Register(&EventReport{}, &EventReportList{})
}
