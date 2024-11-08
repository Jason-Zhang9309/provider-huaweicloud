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

type MySQLDatabaseInitParameters struct {

	// Specifies the character set used by the database, For example utf8,
	// gbk, ascii, etc. Changing this will create a new resource.
	// Specifies the character set used by the database.
	CharacterSet *string `json:"characterSet,omitempty" tf:"character_set,omitempty"`

	// Specifies the database description. The value can contain 0 to 512 characters.
	// This parameter takes effect only for DB instances whose kernel versions are at least 5.6.51.3, 5.7.33.1,
	// or 8.0.21.4.
	// Specifies the database description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the RDS instance ID. Changing this will create a new resource.
	// Specifies the ID of the RDS Mysql instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/rds/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in rds to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in rds to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the database name. The database name contains 1 to 64
	// characters. The name can only consist of lowercase letters, digits, hyphens (-), underscores (_) and dollar signs
	// ($). The total number of hyphens (-) and dollar signs ($) cannot exceed 10. RDS for MySQL 8.0 does not
	// support dollar signs ($). Changing this will create a new resource.
	// Specifies the database name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to create the RDS database resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type MySQLDatabaseObservation struct {

	// Specifies the character set used by the database, For example utf8,
	// gbk, ascii, etc. Changing this will create a new resource.
	// Specifies the character set used by the database.
	CharacterSet *string `json:"characterSet,omitempty" tf:"character_set,omitempty"`

	// Specifies the database description. The value can contain 0 to 512 characters.
	// This parameter takes effect only for DB instances whose kernel versions are at least 5.6.51.3, 5.7.33.1,
	// or 8.0.21.4.
	// Specifies the database description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The resource ID of database which is formatted <instance_id>/<name>.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the RDS instance ID. Changing this will create a new resource.
	// Specifies the ID of the RDS Mysql instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Specifies the database name. The database name contains 1 to 64
	// characters. The name can only consist of lowercase letters, digits, hyphens (-), underscores (_) and dollar signs
	// ($). The total number of hyphens (-) and dollar signs ($) cannot exceed 10. RDS for MySQL 8.0 does not
	// support dollar signs ($). Changing this will create a new resource.
	// Specifies the database name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to create the RDS database resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type MySQLDatabaseParameters struct {

	// Specifies the character set used by the database, For example utf8,
	// gbk, ascii, etc. Changing this will create a new resource.
	// Specifies the character set used by the database.
	// +kubebuilder:validation:Optional
	CharacterSet *string `json:"characterSet,omitempty" tf:"character_set,omitempty"`

	// Specifies the database description. The value can contain 0 to 512 characters.
	// This parameter takes effect only for DB instances whose kernel versions are at least 5.6.51.3, 5.7.33.1,
	// or 8.0.21.4.
	// Specifies the database description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the RDS instance ID. Changing this will create a new resource.
	// Specifies the ID of the RDS Mysql instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/rds/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in rds to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in rds to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the database name. The database name contains 1 to 64
	// characters. The name can only consist of lowercase letters, digits, hyphens (-), underscores (_) and dollar signs
	// ($). The total number of hyphens (-) and dollar signs ($) cannot exceed 10. RDS for MySQL 8.0 does not
	// support dollar signs ($). Changing this will create a new resource.
	// Specifies the database name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to create the RDS database resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// MySQLDatabaseSpec defines the desired state of MySQLDatabase
type MySQLDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MySQLDatabaseParameters `json:"forProvider"`
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
	InitProvider MySQLDatabaseInitParameters `json:"initProvider,omitempty"`
}

// MySQLDatabaseStatus defines the observed state of MySQLDatabase.
type MySQLDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MySQLDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MySQLDatabase is the Schema for the MySQLDatabases API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type MySQLDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.characterSet) || (has(self.initProvider) && has(self.initProvider.characterSet))",message="spec.forProvider.characterSet is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   MySQLDatabaseSpec   `json:"spec"`
	Status MySQLDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MySQLDatabaseList contains a list of MySQLDatabases
type MySQLDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLDatabase `json:"items"`
}

// Repository type metadata.
var (
	MySQLDatabase_Kind             = "MySQLDatabase"
	MySQLDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MySQLDatabase_Kind}.String()
	MySQLDatabase_KindAPIVersion   = MySQLDatabase_Kind + "." + CRDGroupVersion.String()
	MySQLDatabase_GroupVersionKind = CRDGroupVersion.WithKind(MySQLDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&MySQLDatabase{}, &MySQLDatabaseList{})
}
