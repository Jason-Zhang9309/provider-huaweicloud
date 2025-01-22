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

type AlarmTemplateInitParameters struct {

	// Specifies whether delete the alarm rule which the alarm
	// template associated with. Default to false.
	// Specifies whether delete the alarm rule which the alarm template associated with.
	DeleteAssociateAlarm *bool `json:"deleteAssociateAlarm,omitempty" tf:"delete_associate_alarm,omitempty"`

	// Specifies the description of the CES alarm template.
	// The description can contain a maximum of 256 characters.
	// Specifies the description of the CES alarm template.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the name of the CES alarm template.
	// An alarm template name starts with a letter or Chinese, consists of 1 to 128 characters,
	// and can contain only letters, Chinese characters, digits, hyphens (-) and hyphens (-).
	// Specifies the name of the CES alarm template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the policy list of the CES alarm template.
	// The Policy structure is documented below.
	// Specifies the policy list of the CES alarm template.
	Policies []PoliciesInitParameters `json:"policies,omitempty" tf:"policies,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the type of the CES alarm template.
	// Default to 0. The valid values are as follows:
	// Specifies the type of the CES alarm template.
	Type *float64 `json:"type,omitempty" tf:"type,omitempty"`
}

type AlarmTemplateObservation struct {

	// Indicates the total num of the alarm that associated with the alarm template.
	// Indicates the total num of the alarm that associated with the alarm template.
	AssociationAlarmTotal *float64 `json:"associationAlarmTotal,omitempty" tf:"association_alarm_total,omitempty"`

	// Specifies whether delete the alarm rule which the alarm
	// template associated with. Default to false.
	// Specifies whether delete the alarm rule which the alarm template associated with.
	DeleteAssociateAlarm *bool `json:"deleteAssociateAlarm,omitempty" tf:"delete_associate_alarm,omitempty"`

	// Specifies the description of the CES alarm template.
	// The description can contain a maximum of 256 characters.
	// Specifies the description of the CES alarm template.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the CES alarm template.
	// An alarm template name starts with a letter or Chinese, consists of 1 to 128 characters,
	// and can contain only letters, Chinese characters, digits, hyphens (-) and hyphens (-).
	// Specifies the name of the CES alarm template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the policy list of the CES alarm template.
	// The Policy structure is documented below.
	// Specifies the policy list of the CES alarm template.
	Policies []PoliciesObservation `json:"policies,omitempty" tf:"policies,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the type of the CES alarm template.
	// Default to 0. The valid values are as follows:
	// Specifies the type of the CES alarm template.
	Type *float64 `json:"type,omitempty" tf:"type,omitempty"`
}

type AlarmTemplateParameters struct {

	// Specifies whether delete the alarm rule which the alarm
	// template associated with. Default to false.
	// Specifies whether delete the alarm rule which the alarm template associated with.
	// +kubebuilder:validation:Optional
	DeleteAssociateAlarm *bool `json:"deleteAssociateAlarm,omitempty" tf:"delete_associate_alarm,omitempty"`

	// Specifies the description of the CES alarm template.
	// The description can contain a maximum of 256 characters.
	// Specifies the description of the CES alarm template.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the name of the CES alarm template.
	// An alarm template name starts with a letter or Chinese, consists of 1 to 128 characters,
	// and can contain only letters, Chinese characters, digits, hyphens (-) and hyphens (-).
	// Specifies the name of the CES alarm template.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the policy list of the CES alarm template.
	// The Policy structure is documented below.
	// Specifies the policy list of the CES alarm template.
	// +kubebuilder:validation:Optional
	Policies []PoliciesParameters `json:"policies,omitempty" tf:"policies,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the type of the CES alarm template.
	// Default to 0. The valid values are as follows:
	// Specifies the type of the CES alarm template.
	// +kubebuilder:validation:Optional
	Type *float64 `json:"type,omitempty" tf:"type,omitempty"`
}

type PoliciesInitParameters struct {

	// Specifies the alarm level. It means no level if not set.
	// The valid values are as follows:
	// Specifies the alarm level.
	AlarmLevel *float64 `json:"alarmLevel,omitempty" tf:"alarm_level,omitempty"`

	// Specifies the comparison conditions for alarm threshold.
	// Value options: >, <, =, >=, <=.
	// Specifies the comparison conditions for alarm threshold.
	ComparisonOperator *string `json:"comparisonOperator,omitempty" tf:"comparison_operator,omitempty"`

	// Specifies the number of consecutive triggering of alarms. The value ranges from 1 to 5.
	// Specifies the number of consecutive triggering of alarms.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// Specifies the resource dimension.
	// The name starts with a letter and separated by commas(,) for multiple dimensions,
	// can contain only letters, digits, hyphens (-) and hyphens (-),
	// and contain a maximum of 32 characters for each dimension.
	// Specifies the resource dimension.
	DimensionName *string `json:"dimensionName,omitempty" tf:"dimension_name,omitempty"`

	// Specifies the data rollup methods.
	// Value options: average, variance, min, max, sum.
	// Specifies the data rollup methods.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// Specifies the alarm metric name.
	// Specifies the alarm metric name.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Specifies the namespace of the service.
	// Specifies the namespace of the service.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Specifies the judgment period of alarm condition.
	// Value options: 0, 1, 300, 1200, 3600, 14400, 86400.
	// Specifies the judgment period of alarm condition.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Specifies the alarm suppression cycle. Unit: second.
	// Only one alarm is sent when the alarm suppression period is 0.
	// Value options: 0, 300, 600, 900, 1800, 3600, 10800, 21600,
	// 43200, 86400.
	// Specifies the alarm suppression cycle.
	SuppressDuration *float64 `json:"suppressDuration,omitempty" tf:"suppress_duration,omitempty"`

	// Specifies the unit string of the alarm threshold.
	// The unit can contain a maximum of 32 characters.
	// Specifies the unit string of the alarm threshold.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// Specifies the alarm threshold.
	// Specifies the alarm threshold.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type PoliciesObservation struct {

	// Specifies the alarm level. It means no level if not set.
	// The valid values are as follows:
	// Specifies the alarm level.
	AlarmLevel *float64 `json:"alarmLevel,omitempty" tf:"alarm_level,omitempty"`

	// Specifies the comparison conditions for alarm threshold.
	// Value options: >, <, =, >=, <=.
	// Specifies the comparison conditions for alarm threshold.
	ComparisonOperator *string `json:"comparisonOperator,omitempty" tf:"comparison_operator,omitempty"`

	// Specifies the number of consecutive triggering of alarms. The value ranges from 1 to 5.
	// Specifies the number of consecutive triggering of alarms.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// Specifies the resource dimension.
	// The name starts with a letter and separated by commas(,) for multiple dimensions,
	// can contain only letters, digits, hyphens (-) and hyphens (-),
	// and contain a maximum of 32 characters for each dimension.
	// Specifies the resource dimension.
	DimensionName *string `json:"dimensionName,omitempty" tf:"dimension_name,omitempty"`

	// Specifies the data rollup methods.
	// Value options: average, variance, min, max, sum.
	// Specifies the data rollup methods.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// Specifies the alarm metric name.
	// Specifies the alarm metric name.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Specifies the namespace of the service.
	// Specifies the namespace of the service.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Specifies the judgment period of alarm condition.
	// Value options: 0, 1, 300, 1200, 3600, 14400, 86400.
	// Specifies the judgment period of alarm condition.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Specifies the alarm suppression cycle. Unit: second.
	// Only one alarm is sent when the alarm suppression period is 0.
	// Value options: 0, 300, 600, 900, 1800, 3600, 10800, 21600,
	// 43200, 86400.
	// Specifies the alarm suppression cycle.
	SuppressDuration *float64 `json:"suppressDuration,omitempty" tf:"suppress_duration,omitempty"`

	// Specifies the unit string of the alarm threshold.
	// The unit can contain a maximum of 32 characters.
	// Specifies the unit string of the alarm threshold.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// Specifies the alarm threshold.
	// Specifies the alarm threshold.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type PoliciesParameters struct {

	// Specifies the alarm level. It means no level if not set.
	// The valid values are as follows:
	// Specifies the alarm level.
	// +kubebuilder:validation:Optional
	AlarmLevel *float64 `json:"alarmLevel,omitempty" tf:"alarm_level,omitempty"`

	// Specifies the comparison conditions for alarm threshold.
	// Value options: >, <, =, >=, <=.
	// Specifies the comparison conditions for alarm threshold.
	// +kubebuilder:validation:Optional
	ComparisonOperator *string `json:"comparisonOperator" tf:"comparison_operator,omitempty"`

	// Specifies the number of consecutive triggering of alarms. The value ranges from 1 to 5.
	// Specifies the number of consecutive triggering of alarms.
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count" tf:"count,omitempty"`

	// Specifies the resource dimension.
	// The name starts with a letter and separated by commas(,) for multiple dimensions,
	// can contain only letters, digits, hyphens (-) and hyphens (-),
	// and contain a maximum of 32 characters for each dimension.
	// Specifies the resource dimension.
	// +kubebuilder:validation:Optional
	DimensionName *string `json:"dimensionName,omitempty" tf:"dimension_name,omitempty"`

	// Specifies the data rollup methods.
	// Value options: average, variance, min, max, sum.
	// Specifies the data rollup methods.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter" tf:"filter,omitempty"`

	// Specifies the alarm metric name.
	// Specifies the alarm metric name.
	// +kubebuilder:validation:Optional
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// Specifies the namespace of the service.
	// Specifies the namespace of the service.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`

	// Specifies the judgment period of alarm condition.
	// Value options: 0, 1, 300, 1200, 3600, 14400, 86400.
	// Specifies the judgment period of alarm condition.
	// +kubebuilder:validation:Optional
	Period *float64 `json:"period" tf:"period,omitempty"`

	// Specifies the alarm suppression cycle. Unit: second.
	// Only one alarm is sent when the alarm suppression period is 0.
	// Value options: 0, 300, 600, 900, 1800, 3600, 10800, 21600,
	// 43200, 86400.
	// Specifies the alarm suppression cycle.
	// +kubebuilder:validation:Optional
	SuppressDuration *float64 `json:"suppressDuration" tf:"suppress_duration,omitempty"`

	// Specifies the unit string of the alarm threshold.
	// The unit can contain a maximum of 32 characters.
	// Specifies the unit string of the alarm threshold.
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// Specifies the alarm threshold.
	// Specifies the alarm threshold.
	// +kubebuilder:validation:Optional
	Value *float64 `json:"value" tf:"value,omitempty"`
}

// AlarmTemplateSpec defines the desired state of AlarmTemplate
type AlarmTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlarmTemplateParameters `json:"forProvider"`
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
	InitProvider AlarmTemplateInitParameters `json:"initProvider,omitempty"`
}

// AlarmTemplateStatus defines the observed state of AlarmTemplate.
type AlarmTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlarmTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AlarmTemplate is the Schema for the AlarmTemplates API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type AlarmTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policies) || (has(self.initProvider) && has(self.initProvider.policies))",message="spec.forProvider.policies is a required parameter"
	Spec   AlarmTemplateSpec   `json:"spec"`
	Status AlarmTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlarmTemplateList contains a list of AlarmTemplates
type AlarmTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlarmTemplate `json:"items"`
}

// Repository type metadata.
var (
	AlarmTemplate_Kind             = "AlarmTemplate"
	AlarmTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AlarmTemplate_Kind}.String()
	AlarmTemplate_KindAPIVersion   = AlarmTemplate_Kind + "." + CRDGroupVersion.String()
	AlarmTemplate_GroupVersionKind = CRDGroupVersion.WithKind(AlarmTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&AlarmTemplate{}, &AlarmTemplateList{})
}
