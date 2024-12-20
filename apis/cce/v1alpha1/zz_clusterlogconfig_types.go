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

type ClusterLogConfigInitParameters struct {

	// Specifies the cluster ID.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/cce/v1alpha1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a Cluster in cce to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in cce to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Specifies the list of log configs.
	// The log_configs structure is documented below.
	LogConfigs []LogConfigsInitParameters `json:"logConfigs,omitempty" tf:"log_configs,omitempty"`

	// Specifies the region in which to create the cluster log config resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the log keeping days, default to 7.
	TTLInDays *float64 `json:"ttlInDays,omitempty" tf:"ttl_in_days,omitempty"`
}

type ClusterLogConfigObservation struct {

	// Specifies the cluster ID.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// The resource ID, the value is the cluster ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the list of log configs.
	// The log_configs structure is documented below.
	LogConfigs []LogConfigsObservation `json:"logConfigs,omitempty" tf:"log_configs,omitempty"`

	// Specifies the region in which to create the cluster log config resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the log keeping days, default to 7.
	TTLInDays *float64 `json:"ttlInDays,omitempty" tf:"ttl_in_days,omitempty"`
}

type ClusterLogConfigParameters struct {

	// Specifies the cluster ID.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/cce/v1alpha1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a Cluster in cce to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in cce to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Specifies the list of log configs.
	// The log_configs structure is documented below.
	// +kubebuilder:validation:Optional
	LogConfigs []LogConfigsParameters `json:"logConfigs,omitempty" tf:"log_configs,omitempty"`

	// Specifies the region in which to create the cluster log config resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the log keeping days, default to 7.
	// +kubebuilder:validation:Optional
	TTLInDays *float64 `json:"ttlInDays,omitempty" tf:"ttl_in_days,omitempty"`
}

type LogConfigsInitParameters struct {

	// Specifies whether to collect the log.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Specifies the log type.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type LogConfigsObservation struct {

	// Specifies whether to collect the log.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Specifies the log type.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type LogConfigsParameters struct {

	// Specifies whether to collect the log.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Specifies the log type.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// ClusterLogConfigSpec defines the desired state of ClusterLogConfig
type ClusterLogConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterLogConfigParameters `json:"forProvider"`
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
	InitProvider ClusterLogConfigInitParameters `json:"initProvider,omitempty"`
}

// ClusterLogConfigStatus defines the observed state of ClusterLogConfig.
type ClusterLogConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterLogConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ClusterLogConfig is the Schema for the ClusterLogConfigs API. Use this resource to manage the log config of a CCE cluster within HuaweiCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type ClusterLogConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterLogConfigSpec   `json:"spec"`
	Status            ClusterLogConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterLogConfigList contains a list of ClusterLogConfigs
type ClusterLogConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterLogConfig `json:"items"`
}

// Repository type metadata.
var (
	ClusterLogConfig_Kind             = "ClusterLogConfig"
	ClusterLogConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterLogConfig_Kind}.String()
	ClusterLogConfig_KindAPIVersion   = ClusterLogConfig_Kind + "." + CRDGroupVersion.String()
	ClusterLogConfig_GroupVersionKind = CRDGroupVersion.WithKind(ClusterLogConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterLogConfig{}, &ClusterLogConfigList{})
}
