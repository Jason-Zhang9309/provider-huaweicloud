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

type SqlserverDatabaseInitParameters struct {

	// Specifies the ID of the RDS SQLServer instance.
	// Specifies the ID of the RDS SQLServer instance.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/rds/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in rds to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in rds to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the database name. The database name can contain 1 to 64 characters,
	// and can include letters, digits, hyphens (-), underscores (_), and periods (.). It cannot start or end with an RDS for
	// SQL Server system database name. RDS for SQL Server system databases include master, msdb, model,
	// tempdb, resource, and rdsadmin.
	// Specifies the database name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type SqlserverDatabaseObservation struct {

	// Indicates the character set used by the database.
	// Indicates the character set used by the database.
	CharacterSet *string `json:"characterSet,omitempty" tf:"character_set,omitempty"`

	// The resource ID of database which is formatted <instance_id>/<name>.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the RDS SQLServer instance.
	// Specifies the ID of the RDS SQLServer instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Specifies the database name. The database name can contain 1 to 64 characters,
	// and can include letters, digits, hyphens (-), underscores (_), and periods (.). It cannot start or end with an RDS for
	// SQL Server system database name. RDS for SQL Server system databases include master, msdb, model,
	// tempdb, resource, and rdsadmin.
	// Specifies the database name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Indicates the database status. Its value can be any of the following:
	// Indicates the database status.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type SqlserverDatabaseParameters struct {

	// Specifies the ID of the RDS SQLServer instance.
	// Specifies the ID of the RDS SQLServer instance.
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

	// Specifies the database name. The database name can contain 1 to 64 characters,
	// and can include letters, digits, hyphens (-), underscores (_), and periods (.). It cannot start or end with an RDS for
	// SQL Server system database name. RDS for SQL Server system databases include master, msdb, model,
	// tempdb, resource, and rdsadmin.
	// Specifies the database name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// SqlserverDatabaseSpec defines the desired state of SqlserverDatabase
type SqlserverDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SqlserverDatabaseParameters `json:"forProvider"`
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
	InitProvider SqlserverDatabaseInitParameters `json:"initProvider,omitempty"`
}

// SqlserverDatabaseStatus defines the observed state of SqlserverDatabase.
type SqlserverDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SqlserverDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SqlserverDatabase is the Schema for the SqlserverDatabases API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type SqlserverDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SqlserverDatabaseSpec   `json:"spec"`
	Status SqlserverDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqlserverDatabaseList contains a list of SqlserverDatabases
type SqlserverDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlserverDatabase `json:"items"`
}

// Repository type metadata.
var (
	SqlserverDatabase_Kind             = "SqlserverDatabase"
	SqlserverDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SqlserverDatabase_Kind}.String()
	SqlserverDatabase_KindAPIVersion   = SqlserverDatabase_Kind + "." + CRDGroupVersion.String()
	SqlserverDatabase_GroupVersionKind = CRDGroupVersion.WithKind(SqlserverDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&SqlserverDatabase{}, &SqlserverDatabaseList{})
}
