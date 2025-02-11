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

type AdviceIdsInitParameters struct {
}

type AdviceIdsObservation struct {

	// The resource ID. The value is the dianosis report ID.
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates the conclusion parameters.
	// +mapType=granular
	Params map[string]*string `json:"params,omitempty" tf:"params,omitempty"`
}

type AdviceIdsParameters struct {
}

type CauseIdsInitParameters struct {
}

type CauseIdsObservation struct {

	// The resource ID. The value is the dianosis report ID.
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates the conclusion parameters.
	// +mapType=granular
	Params map[string]*string `json:"params,omitempty" tf:"params,omitempty"`
}

type CauseIdsParameters struct {
}

type CommandListInitParameters struct {
}

type CommandListObservation struct {

	// Indicates the average duration of calls.
	AverageUsec *float64 `json:"averageUsec,omitempty" tf:"average_usec,omitempty"`

	// Indicates the number of calls.
	CallsSum *float64 `json:"callsSum,omitempty" tf:"calls_sum,omitempty"`

	// Indicates the command name.
	CommandName *string `json:"commandName,omitempty" tf:"command_name,omitempty"`

	// Indicates the duration percentage.
	PerUsec *string `json:"perUsec,omitempty" tf:"per_usec,omitempty"`

	// Indicates the total time consumed.
	UsecSum *float64 `json:"usecSum,omitempty" tf:"usec_sum,omitempty"`
}

type CommandListParameters struct {
}

type CommandTimeTakenListInitParameters struct {
}

type CommandTimeTakenListObservation struct {

	// Indicates the command execution latency statistics.
	// The command_list structure is documented below.
	CommandList []CommandListObservation `json:"commandList,omitempty" tf:"command_list,omitempty"`

	// Indicates the error code for the diagnosis item.
	ErrorCode *string `json:"errorCode,omitempty" tf:"error_code,omitempty"`

	// Indicates the diagnosis result. The value can be failed, abnormal or normal.
	Result *string `json:"result,omitempty" tf:"result,omitempty"`

	// Indicates the total number of times that commands are executed.
	TotalNum *float64 `json:"totalNum,omitempty" tf:"total_num,omitempty"`

	// Indicates the total duration of command execution.
	TotalUsecSum *float64 `json:"totalUsecSum,omitempty" tf:"total_usec_sum,omitempty"`
}

type CommandTimeTakenListParameters struct {
}

type DiagnosisDimensionListInitParameters struct {
}

type DiagnosisDimensionListObservation struct {

	// Indicates the total number of abnormal diagnosis items.
	AbnormalNum *float64 `json:"abnormalNum,omitempty" tf:"abnormal_num,omitempty"`

	// Indicates the diagnosis items.
	// The diagnosis_item_list structure is documented below.
	DiagnosisItemList []DiagnosisItemListObservation `json:"diagnosisItemList,omitempty" tf:"diagnosis_item_list,omitempty"`

	// Indicates the total number of failed diagnosis items.
	FailedNum *float64 `json:"failedNum,omitempty" tf:"failed_num,omitempty"`

	// Indicates the diagnosis item name.
	// The value can be connection_num, rx_controlled, persistence, centralized_expiration,
	// inner_memory_fragmentation, time_consuming_commands, hit_ratio, memory_usage or cpu_usage.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DiagnosisDimensionListParameters struct {
}

type DiagnosisItemListInitParameters struct {
}

type DiagnosisItemListObservation struct {

	// Indicates the list of suggestion IDs.
	// The advice_ids structure is documented below.
	AdviceIds []AdviceIdsObservation `json:"adviceIds,omitempty" tf:"advice_ids,omitempty"`

	// Indicates the list of cause IDs.
	// The cause_ids structure is documented below.
	CauseIds []CauseIdsObservation `json:"causeIds,omitempty" tf:"cause_ids,omitempty"`

	// Indicates the error code for the diagnosis item.
	ErrorCode *string `json:"errorCode,omitempty" tf:"error_code,omitempty"`

	// Indicates the list of impact IDs.
	// The impact_ids structure is documented below.
	ImpactIds []ImpactIdsObservation `json:"impactIds,omitempty" tf:"impact_ids,omitempty"`

	// Indicates the diagnosis item name.
	// The value can be connection_num, rx_controlled, persistence, centralized_expiration,
	// inner_memory_fragmentation, time_consuming_commands, hit_ratio, memory_usage or cpu_usage.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates the diagnosis result. The value can be failed, abnormal or normal.
	Result *string `json:"result,omitempty" tf:"result,omitempty"`
}

type DiagnosisItemListParameters struct {
}

type DiagnosisNodeReportListInitParameters struct {
}

type DiagnosisNodeReportListObservation struct {

	// Indicates the total number of abnormal diagnosis items.
	AbnormalSum *float64 `json:"abnormalSum,omitempty" tf:"abnormal_sum,omitempty"`

	// Indicates the code of the AZ where the node is.
	AzCode *string `json:"azCode,omitempty" tf:"az_code,omitempty"`

	// Indicates the command execution duration list.
	// The command_time_taken_list structure is documented below.
	CommandTimeTakenList []CommandTimeTakenListObservation `json:"commandTimeTakenList,omitempty" tf:"command_time_taken_list,omitempty"`

	// Indicates the diagnosis dimension list.
	// The diagnosis_dimension_list structure is documented below.
	DiagnosisDimensionList []DiagnosisDimensionListObservation `json:"diagnosisDimensionList,omitempty" tf:"diagnosis_dimension_list,omitempty"`

	// Indicates the total number of failed diagnosis items.
	FailedSum *float64 `json:"failedSum,omitempty" tf:"failed_sum,omitempty"`

	// Indicates the name of the shard where the node is.
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	// Indicates whether the node is faulted.
	IsFaulted *bool `json:"isFaulted,omitempty" tf:"is_faulted,omitempty"`

	// Indicates the IP address of the node diagnosed.
	NodeIP *string `json:"nodeIp,omitempty" tf:"node_ip,omitempty"`

	// Indicates the node role. The value can be master or slave.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type DiagnosisNodeReportListParameters struct {
}

type DiagnosisTaskInitParameters struct {

	// Specifies the start time of the diagnosis task, in RFC3339 format.
	// Changing this creates a new resource.
	// Specifies the start time of the diagnosis task, in RFC3339 format.
	BeginTime *string `json:"beginTime,omitempty" tf:"begin_time,omitempty"`

	// Specifies the end time of the diagnosis task, in RFC3339 format.
	// Changing this creates a new resource.
	// Specifies the end time of the diagnosis task, in RFC3339 format.
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies the ID of the DCS instance.
	// Changing this creates a new resource.
	// Specifies the ID of the DCS instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/dcs/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in dcs to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in dcs to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the IP addresses of diagnosed nodes.
	// By default, all nodes are diagnosed. Changing this creates a new resource.
	// Specifies the IP addresses of diagnosed nodes. By default, all nodes are diagnosed.
	// +listType=set
	NodeIPList []*string `json:"nodeIpList,omitempty" tf:"node_ip_list,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type DiagnosisTaskObservation struct {

	// Indicates the total number of abnormal diagnosis items.
	// Indicates the total number of abnormal diagnosis items.
	AbnormalItemSum *float64 `json:"abnormalItemSum,omitempty" tf:"abnormal_item_sum,omitempty"`

	// Specifies the start time of the diagnosis task, in RFC3339 format.
	// Changing this creates a new resource.
	// Specifies the start time of the diagnosis task, in RFC3339 format.
	BeginTime *string `json:"beginTime,omitempty" tf:"begin_time,omitempty"`

	// Indicates the list of node diagnosis report
	// The diagnosis_node_report_list structure is documented below.
	// Indicates the list of node diagnosis report
	DiagnosisNodeReportList []DiagnosisNodeReportListObservation `json:"diagnosisNodeReportList,omitempty" tf:"diagnosis_node_report_list,omitempty"`

	// Specifies the end time of the diagnosis task, in RFC3339 format.
	// Changing this creates a new resource.
	// Specifies the end time of the diagnosis task, in RFC3339 format.
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Indicates the total number of failed diagnosis items.
	// Indicates the total number of failed diagnosis items.
	FailedItemSum *float64 `json:"failedItemSum,omitempty" tf:"failed_item_sum,omitempty"`

	// The resource ID. The value is the dianosis report ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the DCS instance.
	// Changing this creates a new resource.
	// Specifies the ID of the DCS instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Specifies the IP addresses of diagnosed nodes.
	// By default, all nodes are diagnosed. Changing this creates a new resource.
	// Specifies the IP addresses of diagnosed nodes. By default, all nodes are diagnosed.
	// +listType=set
	NodeIPList []*string `json:"nodeIpList,omitempty" tf:"node_ip_list,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type DiagnosisTaskParameters struct {

	// Specifies the start time of the diagnosis task, in RFC3339 format.
	// Changing this creates a new resource.
	// Specifies the start time of the diagnosis task, in RFC3339 format.
	// +kubebuilder:validation:Optional
	BeginTime *string `json:"beginTime,omitempty" tf:"begin_time,omitempty"`

	// Specifies the end time of the diagnosis task, in RFC3339 format.
	// Changing this creates a new resource.
	// Specifies the end time of the diagnosis task, in RFC3339 format.
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies the ID of the DCS instance.
	// Changing this creates a new resource.
	// Specifies the ID of the DCS instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/dcs/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in dcs to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in dcs to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the IP addresses of diagnosed nodes.
	// By default, all nodes are diagnosed. Changing this creates a new resource.
	// Specifies the IP addresses of diagnosed nodes. By default, all nodes are diagnosed.
	// +kubebuilder:validation:Optional
	// +listType=set
	NodeIPList []*string `json:"nodeIpList,omitempty" tf:"node_ip_list,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type ImpactIdsInitParameters struct {
}

type ImpactIdsObservation struct {

	// The resource ID. The value is the dianosis report ID.
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates the conclusion parameters.
	// +mapType=granular
	Params map[string]*string `json:"params,omitempty" tf:"params,omitempty"`
}

type ImpactIdsParameters struct {
}

// DiagnosisTaskSpec defines the desired state of DiagnosisTask
type DiagnosisTaskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiagnosisTaskParameters `json:"forProvider"`
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
	InitProvider DiagnosisTaskInitParameters `json:"initProvider,omitempty"`
}

// DiagnosisTaskStatus defines the observed state of DiagnosisTask.
type DiagnosisTaskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiagnosisTaskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DiagnosisTask is the Schema for the DiagnosisTasks API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type DiagnosisTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.beginTime) || (has(self.initProvider) && has(self.initProvider.beginTime))",message="spec.forProvider.beginTime is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.endTime) || (has(self.initProvider) && has(self.initProvider.endTime))",message="spec.forProvider.endTime is a required parameter"
	Spec   DiagnosisTaskSpec   `json:"spec"`
	Status DiagnosisTaskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiagnosisTaskList contains a list of DiagnosisTasks
type DiagnosisTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiagnosisTask `json:"items"`
}

// Repository type metadata.
var (
	DiagnosisTask_Kind             = "DiagnosisTask"
	DiagnosisTask_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DiagnosisTask_Kind}.String()
	DiagnosisTask_KindAPIVersion   = DiagnosisTask_Kind + "." + CRDGroupVersion.String()
	DiagnosisTask_GroupVersionKind = CRDGroupVersion.WithKind(DiagnosisTask_Kind)
)

func init() {
	SchemeBuilder.Register(&DiagnosisTask{}, &DiagnosisTaskList{})
}
