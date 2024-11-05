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

type TrafficMirrorFilterRuleInitParameters struct {

	// Specifies the policy of in the traffic mirror filter rule.
	// Valid values are accept or reject.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Specifies the description of the traffic mirror filter rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the destination IP address of the traffic traffic mirror filter
	// rule.
	DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty" tf:"destination_cidr_block,omitempty"`

	// Specifies the destination port number range of the traffic mirror filter
	// rule. The value ranges from 1 to 65,535, enter two port numbers connected by a hyphen (-). For example, 80-200.
	DestinationPortRange *string `json:"destinationPortRange,omitempty" tf:"destination_port_range,omitempty"`

	// Specifies the direction of the traffic mirror filter rule.
	// Valid values are ingress or egress. Changing this creates a new resource.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Specifies the IP address protocol type of the traffic mirror filter rule.
	// Valid values are IPv4 or IPv6.
	Ethertype *string `json:"ethertype,omitempty" tf:"ethertype,omitempty"`

	// Specifies the priority number of the traffic mirror filter rule.
	// Valid value ranges from 1 to 65,535.
	// The smaller the priority number, the higher the priority.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Specifies the protocol of the traffic mirror filter rule.
	// Valid value are tcp, udp, icmp, icmpv6, all.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the source IP address of the traffic traffic mirror filter rule.
	SourceCidrBlock *string `json:"sourceCidrBlock,omitempty" tf:"source_cidr_block,omitempty"`

	// Specifies the source port number range of the traffic mirror filter rule.
	// The value ranges from 1 to 65,535, enter two port numbers connected by a hyphen (-). For example, 80-200.
	SourcePortRange *string `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`

	// Specifies an ID of the traffic mirror filter to which
	// the rule belongs. Changing this creates a new resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.TrafficMirrorFilter
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	TrafficMirrorFilterID *string `json:"trafficMirrorFilterId,omitempty" tf:"traffic_mirror_filter_id,omitempty"`

	// Reference to a TrafficMirrorFilter in vpc to populate trafficMirrorFilterId.
	// +kubebuilder:validation:Optional
	TrafficMirrorFilterIDRef *v1.Reference `json:"trafficMirrorFilterIdRef,omitempty" tf:"-"`

	// Selector for a TrafficMirrorFilter in vpc to populate trafficMirrorFilterId.
	// +kubebuilder:validation:Optional
	TrafficMirrorFilterIDSelector *v1.Selector `json:"trafficMirrorFilterIdSelector,omitempty" tf:"-"`
}

type TrafficMirrorFilterRuleObservation struct {

	// Specifies the policy of in the traffic mirror filter rule.
	// Valid values are accept or reject.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The creation time of the traffic mirror filter rule.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Specifies the description of the traffic mirror filter rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the destination IP address of the traffic traffic mirror filter
	// rule.
	DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty" tf:"destination_cidr_block,omitempty"`

	// Specifies the destination port number range of the traffic mirror filter
	// rule. The value ranges from 1 to 65,535, enter two port numbers connected by a hyphen (-). For example, 80-200.
	DestinationPortRange *string `json:"destinationPortRange,omitempty" tf:"destination_port_range,omitempty"`

	// Specifies the direction of the traffic mirror filter rule.
	// Valid values are ingress or egress. Changing this creates a new resource.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Specifies the IP address protocol type of the traffic mirror filter rule.
	// Valid values are IPv4 or IPv6.
	Ethertype *string `json:"ethertype,omitempty" tf:"ethertype,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the priority number of the traffic mirror filter rule.
	// Valid value ranges from 1 to 65,535.
	// The smaller the priority number, the higher the priority.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Specifies the protocol of the traffic mirror filter rule.
	// Valid value are tcp, udp, icmp, icmpv6, all.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the source IP address of the traffic traffic mirror filter rule.
	SourceCidrBlock *string `json:"sourceCidrBlock,omitempty" tf:"source_cidr_block,omitempty"`

	// Specifies the source port number range of the traffic mirror filter rule.
	// The value ranges from 1 to 65,535, enter two port numbers connected by a hyphen (-). For example, 80-200.
	SourcePortRange *string `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`

	// Specifies an ID of the traffic mirror filter to which
	// the rule belongs. Changing this creates a new resource.
	TrafficMirrorFilterID *string `json:"trafficMirrorFilterId,omitempty" tf:"traffic_mirror_filter_id,omitempty"`

	// The latest update time of the traffic mirror filter rule.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type TrafficMirrorFilterRuleParameters struct {

	// Specifies the policy of in the traffic mirror filter rule.
	// Valid values are accept or reject.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Specifies the description of the traffic mirror filter rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the destination IP address of the traffic traffic mirror filter
	// rule.
	// +kubebuilder:validation:Optional
	DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty" tf:"destination_cidr_block,omitempty"`

	// Specifies the destination port number range of the traffic mirror filter
	// rule. The value ranges from 1 to 65,535, enter two port numbers connected by a hyphen (-). For example, 80-200.
	// +kubebuilder:validation:Optional
	DestinationPortRange *string `json:"destinationPortRange,omitempty" tf:"destination_port_range,omitempty"`

	// Specifies the direction of the traffic mirror filter rule.
	// Valid values are ingress or egress. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Specifies the IP address protocol type of the traffic mirror filter rule.
	// Valid values are IPv4 or IPv6.
	// +kubebuilder:validation:Optional
	Ethertype *string `json:"ethertype,omitempty" tf:"ethertype,omitempty"`

	// Specifies the priority number of the traffic mirror filter rule.
	// Valid value ranges from 1 to 65,535.
	// The smaller the priority number, the higher the priority.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Specifies the protocol of the traffic mirror filter rule.
	// Valid value are tcp, udp, icmp, icmpv6, all.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the source IP address of the traffic traffic mirror filter rule.
	// +kubebuilder:validation:Optional
	SourceCidrBlock *string `json:"sourceCidrBlock,omitempty" tf:"source_cidr_block,omitempty"`

	// Specifies the source port number range of the traffic mirror filter rule.
	// The value ranges from 1 to 65,535, enter two port numbers connected by a hyphen (-). For example, 80-200.
	// +kubebuilder:validation:Optional
	SourcePortRange *string `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`

	// Specifies an ID of the traffic mirror filter to which
	// the rule belongs. Changing this creates a new resource.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/vpc/v1alpha1.TrafficMirrorFilter
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TrafficMirrorFilterID *string `json:"trafficMirrorFilterId,omitempty" tf:"traffic_mirror_filter_id,omitempty"`

	// Reference to a TrafficMirrorFilter in vpc to populate trafficMirrorFilterId.
	// +kubebuilder:validation:Optional
	TrafficMirrorFilterIDRef *v1.Reference `json:"trafficMirrorFilterIdRef,omitempty" tf:"-"`

	// Selector for a TrafficMirrorFilter in vpc to populate trafficMirrorFilterId.
	// +kubebuilder:validation:Optional
	TrafficMirrorFilterIDSelector *v1.Selector `json:"trafficMirrorFilterIdSelector,omitempty" tf:"-"`
}

// TrafficMirrorFilterRuleSpec defines the desired state of TrafficMirrorFilterRule
type TrafficMirrorFilterRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrafficMirrorFilterRuleParameters `json:"forProvider"`
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
	InitProvider TrafficMirrorFilterRuleInitParameters `json:"initProvider,omitempty"`
}

// TrafficMirrorFilterRuleStatus defines the observed state of TrafficMirrorFilterRule.
type TrafficMirrorFilterRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrafficMirrorFilterRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TrafficMirrorFilterRule is the Schema for the TrafficMirrorFilterRules API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type TrafficMirrorFilterRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || (has(self.initProvider) && has(self.initProvider.action))",message="spec.forProvider.action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.direction) || (has(self.initProvider) && has(self.initProvider.direction))",message="spec.forProvider.direction is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ethertype) || (has(self.initProvider) && has(self.initProvider.ethertype))",message="spec.forProvider.ethertype is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.priority) || (has(self.initProvider) && has(self.initProvider.priority))",message="spec.forProvider.priority is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	Spec   TrafficMirrorFilterRuleSpec   `json:"spec"`
	Status TrafficMirrorFilterRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficMirrorFilterRuleList contains a list of TrafficMirrorFilterRules
type TrafficMirrorFilterRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficMirrorFilterRule `json:"items"`
}

// Repository type metadata.
var (
	TrafficMirrorFilterRule_Kind             = "TrafficMirrorFilterRule"
	TrafficMirrorFilterRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrafficMirrorFilterRule_Kind}.String()
	TrafficMirrorFilterRule_KindAPIVersion   = TrafficMirrorFilterRule_Kind + "." + CRDGroupVersion.String()
	TrafficMirrorFilterRule_GroupVersionKind = CRDGroupVersion.WithKind(TrafficMirrorFilterRule_Kind)
)

func init() {
	SchemeBuilder.Register(&TrafficMirrorFilterRule{}, &TrafficMirrorFilterRuleList{})
}
