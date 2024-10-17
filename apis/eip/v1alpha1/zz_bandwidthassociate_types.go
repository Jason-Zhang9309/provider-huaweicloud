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

type BandwidthAssociateInitParameters struct {

	// Specifies the billing mode of the dedicated bandwidth used by the EIP that
	// has been removed from a shared bandwidth. The value can be bandwidth or traffic. If not specified, the dedicated
	// bandwidth will be billed by bandwidth.
	// The charge mode after removal bandwidth.
	BandwidthChargeMode *string `json:"bandwidthChargeMode,omitempty" tf:"bandwidth_charge_mode,omitempty"`

	// Specifies the shared bandwidth ID to associate.
	// Changing this creates a new resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/eip/v1alpha1.Bandwidth
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	BandwidthID *string `json:"bandwidthId,omitempty" tf:"bandwidth_id,omitempty"`

	// Reference to a Bandwidth in eip to populate bandwidthId.
	// +kubebuilder:validation:Optional
	BandwidthIDRef *v1.Reference `json:"bandwidthIdRef,omitempty" tf:"-"`

	// Selector for a Bandwidth in eip to populate bandwidthId.
	// +kubebuilder:validation:Optional
	BandwidthIDSelector *v1.Selector `json:"bandwidthIdSelector,omitempty" tf:"-"`

	// Specifies the size (Mbit/s) of the dedicated bandwidth used by the EIP that
	// has been removed from a shared bandwidth. The default bandwidth size is 5 Mbit/s.
	// The size (Mbits/s) after removal bandwidth.
	BandwidthSize *float64 `json:"bandwidthSize,omitempty" tf:"bandwidth_size,omitempty"`

	// Specifies the ID of the EIP that uses the bandwidth.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/eip/v1alpha1.EIP
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	EIPID *string `json:"eipId,omitempty" tf:"eip_id,omitempty"`

	// Reference to a EIP in eip to populate eipId.
	// +kubebuilder:validation:Optional
	EIPIDRef *v1.Reference `json:"eipIdRef,omitempty" tf:"-"`

	// Selector for a EIP in eip to populate eipId.
	// +kubebuilder:validation:Optional
	EIPIDSelector *v1.Selector `json:"eipIdSelector,omitempty" tf:"-"`

	// Specifies the region in which to associate the bandwidth. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type BandwidthAssociateObservation struct {

	// Specifies the billing mode of the dedicated bandwidth used by the EIP that
	// has been removed from a shared bandwidth. The value can be bandwidth or traffic. If not specified, the dedicated
	// bandwidth will be billed by bandwidth.
	// The charge mode after removal bandwidth.
	BandwidthChargeMode *string `json:"bandwidthChargeMode,omitempty" tf:"bandwidth_charge_mode,omitempty"`

	// Specifies the shared bandwidth ID to associate.
	// Changing this creates a new resource.
	BandwidthID *string `json:"bandwidthId,omitempty" tf:"bandwidth_id,omitempty"`

	// The shared bandwidth name.
	BandwidthName *string `json:"bandwidthName,omitempty" tf:"bandwidth_name,omitempty"`

	// Specifies the size (Mbit/s) of the dedicated bandwidth used by the EIP that
	// has been removed from a shared bandwidth. The default bandwidth size is 5 Mbit/s.
	// The size (Mbits/s) after removal bandwidth.
	BandwidthSize *float64 `json:"bandwidthSize,omitempty" tf:"bandwidth_size,omitempty"`

	// Specifies the ID of the EIP that uses the bandwidth.
	EIPID *string `json:"eipId,omitempty" tf:"eip_id,omitempty"`

	// The resource ID in format of <bandwidth_id>/<eip_id>.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The EIP address.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// Specifies the region in which to associate the bandwidth. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type BandwidthAssociateParameters struct {

	// Specifies the billing mode of the dedicated bandwidth used by the EIP that
	// has been removed from a shared bandwidth. The value can be bandwidth or traffic. If not specified, the dedicated
	// bandwidth will be billed by bandwidth.
	// The charge mode after removal bandwidth.
	// +kubebuilder:validation:Optional
	BandwidthChargeMode *string `json:"bandwidthChargeMode,omitempty" tf:"bandwidth_charge_mode,omitempty"`

	// Specifies the shared bandwidth ID to associate.
	// Changing this creates a new resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/eip/v1alpha1.Bandwidth
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	BandwidthID *string `json:"bandwidthId,omitempty" tf:"bandwidth_id,omitempty"`

	// Reference to a Bandwidth in eip to populate bandwidthId.
	// +kubebuilder:validation:Optional
	BandwidthIDRef *v1.Reference `json:"bandwidthIdRef,omitempty" tf:"-"`

	// Selector for a Bandwidth in eip to populate bandwidthId.
	// +kubebuilder:validation:Optional
	BandwidthIDSelector *v1.Selector `json:"bandwidthIdSelector,omitempty" tf:"-"`

	// Specifies the size (Mbit/s) of the dedicated bandwidth used by the EIP that
	// has been removed from a shared bandwidth. The default bandwidth size is 5 Mbit/s.
	// The size (Mbits/s) after removal bandwidth.
	// +kubebuilder:validation:Optional
	BandwidthSize *float64 `json:"bandwidthSize,omitempty" tf:"bandwidth_size,omitempty"`

	// Specifies the ID of the EIP that uses the bandwidth.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/eip/v1alpha1.EIP
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	EIPID *string `json:"eipId,omitempty" tf:"eip_id,omitempty"`

	// Reference to a EIP in eip to populate eipId.
	// +kubebuilder:validation:Optional
	EIPIDRef *v1.Reference `json:"eipIdRef,omitempty" tf:"-"`

	// Selector for a EIP in eip to populate eipId.
	// +kubebuilder:validation:Optional
	EIPIDSelector *v1.Selector `json:"eipIdSelector,omitempty" tf:"-"`

	// Specifies the region in which to associate the bandwidth. If omitted,
	// the provider-level region will be used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// BandwidthAssociateSpec defines the desired state of BandwidthAssociate
type BandwidthAssociateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BandwidthAssociateParameters `json:"forProvider"`
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
	InitProvider BandwidthAssociateInitParameters `json:"initProvider,omitempty"`
}

// BandwidthAssociateStatus defines the observed state of BandwidthAssociate.
type BandwidthAssociateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BandwidthAssociateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BandwidthAssociate is the Schema for the BandwidthAssociates API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type BandwidthAssociate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BandwidthAssociateSpec   `json:"spec"`
	Status            BandwidthAssociateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BandwidthAssociateList contains a list of BandwidthAssociates
type BandwidthAssociateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BandwidthAssociate `json:"items"`
}

// Repository type metadata.
var (
	BandwidthAssociate_Kind             = "BandwidthAssociate"
	BandwidthAssociate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BandwidthAssociate_Kind}.String()
	BandwidthAssociate_KindAPIVersion   = BandwidthAssociate_Kind + "." + CRDGroupVersion.String()
	BandwidthAssociate_GroupVersionKind = CRDGroupVersion.WithKind(BandwidthAssociate_Kind)
)

func init() {
	SchemeBuilder.Register(&BandwidthAssociate{}, &BandwidthAssociateList{})
}
