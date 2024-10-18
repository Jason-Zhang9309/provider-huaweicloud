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

type AccountPermissionInitParameters struct {

	// Specifies the access to acl. Valid values are READ_ACP and WRITE_ACP.
	// Specifies the access to acl. Valid values are **READ_ACP** and **WRITE_ACP**.
	AccessToACL []*string `json:"accessToAcl,omitempty" tf:"access_to_acl,omitempty"`

	// Specifies the access to bucket. Valid values are READ and WRITE.
	// Specifies the access to bucket. Valid values are **READ** and **WRITE**.
	AccessToBucket []*string `json:"accessToBucket,omitempty" tf:"access_to_bucket,omitempty"`

	// Specifies the account id to authorize. The account id cannot be the bucket owner,
	// and must be unique.
	// Specifies the account id to authorize. The account id cannot be the bucket owner,
	// and must be unique.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`
}

type AccountPermissionObservation struct {

	// Specifies the access to acl. Valid values are READ_ACP and WRITE_ACP.
	// Specifies the access to acl. Valid values are **READ_ACP** and **WRITE_ACP**.
	AccessToACL []*string `json:"accessToAcl,omitempty" tf:"access_to_acl,omitempty"`

	// Specifies the access to bucket. Valid values are READ and WRITE.
	// Specifies the access to bucket. Valid values are **READ** and **WRITE**.
	AccessToBucket []*string `json:"accessToBucket,omitempty" tf:"access_to_bucket,omitempty"`

	// Specifies the account id to authorize. The account id cannot be the bucket owner,
	// and must be unique.
	// Specifies the account id to authorize. The account id cannot be the bucket owner,
	// and must be unique.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`
}

type AccountPermissionParameters struct {

	// Specifies the access to acl. Valid values are READ_ACP and WRITE_ACP.
	// Specifies the access to acl. Valid values are **READ_ACP** and **WRITE_ACP**.
	// +kubebuilder:validation:Optional
	AccessToACL []*string `json:"accessToAcl,omitempty" tf:"access_to_acl,omitempty"`

	// Specifies the access to bucket. Valid values are READ and WRITE.
	// Specifies the access to bucket. Valid values are **READ** and **WRITE**.
	// +kubebuilder:validation:Optional
	AccessToBucket []*string `json:"accessToBucket,omitempty" tf:"access_to_bucket,omitempty"`

	// Specifies the account id to authorize. The account id cannot be the bucket owner,
	// and must be unique.
	// Specifies the account id to authorize. The account id cannot be the bucket owner,
	// and must be unique.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId" tf:"account_id,omitempty"`
}

type BucketACLInitParameters struct {

	// Specifies the account permissions.
	// The account_permission_struct structure is documented below.
	// Specifies the account permissions.
	AccountPermission []AccountPermissionInitParameters `json:"accountPermission,omitempty" tf:"account_permission,omitempty"`

	// Specifies the name of the bucket to which to set the acl.
	// Specifies the name of the bucket to which to set the acl.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/obs/v1alpha1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("bucket",true)
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in obs to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in obs to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Specifies the log delivery user permission.
	// The permission_struct structure is documented below.
	// Specifies the log delivery user permission.
	LogDeliveryUserPermission []LogDeliveryUserPermissionInitParameters `json:"logDeliveryUserPermission,omitempty" tf:"log_delivery_user_permission,omitempty"`

	// Specifies the bucket owner permission. If omitted, the current obs bucket acl
	// owner permission will not be changed.
	// The permission_struct structure is documented below.
	// Specifies the bucket owner permission.
	OwnerPermission []OwnerPermissionInitParameters `json:"ownerPermission,omitempty" tf:"owner_permission,omitempty"`

	// Specifies the public permission.
	// The permission_struct structure is documented below.
	// Specifies the public permission.
	PublicPermission []PublicPermissionInitParameters `json:"publicPermission,omitempty" tf:"public_permission,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type BucketACLObservation struct {

	// Specifies the account permissions.
	// The account_permission_struct structure is documented below.
	// Specifies the account permissions.
	AccountPermission []AccountPermissionObservation `json:"accountPermission,omitempty" tf:"account_permission,omitempty"`

	// Specifies the name of the bucket to which to set the acl.
	// Specifies the name of the bucket to which to set the acl.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The name of the bucket.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the log delivery user permission.
	// The permission_struct structure is documented below.
	// Specifies the log delivery user permission.
	LogDeliveryUserPermission []LogDeliveryUserPermissionObservation `json:"logDeliveryUserPermission,omitempty" tf:"log_delivery_user_permission,omitempty"`

	// Specifies the bucket owner permission. If omitted, the current obs bucket acl
	// owner permission will not be changed.
	// The permission_struct structure is documented below.
	// Specifies the bucket owner permission.
	OwnerPermission []OwnerPermissionObservation `json:"ownerPermission,omitempty" tf:"owner_permission,omitempty"`

	// Specifies the public permission.
	// The permission_struct structure is documented below.
	// Specifies the public permission.
	PublicPermission []PublicPermissionObservation `json:"publicPermission,omitempty" tf:"public_permission,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type BucketACLParameters struct {

	// Specifies the account permissions.
	// The account_permission_struct structure is documented below.
	// Specifies the account permissions.
	// +kubebuilder:validation:Optional
	AccountPermission []AccountPermissionParameters `json:"accountPermission,omitempty" tf:"account_permission,omitempty"`

	// Specifies the name of the bucket to which to set the acl.
	// Specifies the name of the bucket to which to set the acl.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/obs/v1alpha1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("bucket",true)
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in obs to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in obs to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Specifies the log delivery user permission.
	// The permission_struct structure is documented below.
	// Specifies the log delivery user permission.
	// +kubebuilder:validation:Optional
	LogDeliveryUserPermission []LogDeliveryUserPermissionParameters `json:"logDeliveryUserPermission,omitempty" tf:"log_delivery_user_permission,omitempty"`

	// Specifies the bucket owner permission. If omitted, the current obs bucket acl
	// owner permission will not be changed.
	// The permission_struct structure is documented below.
	// Specifies the bucket owner permission.
	// +kubebuilder:validation:Optional
	OwnerPermission []OwnerPermissionParameters `json:"ownerPermission,omitempty" tf:"owner_permission,omitempty"`

	// Specifies the public permission.
	// The permission_struct structure is documented below.
	// Specifies the public permission.
	// +kubebuilder:validation:Optional
	PublicPermission []PublicPermissionParameters `json:"publicPermission,omitempty" tf:"public_permission,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type LogDeliveryUserPermissionInitParameters struct {

	// Specifies the access to acl. Valid values are READ_ACP and WRITE_ACP.
	// Specifies the access to acl. Valid values are **READ_ACP** and **WRITE_ACP**.
	AccessToACL []*string `json:"accessToAcl,omitempty" tf:"access_to_acl,omitempty"`

	// Specifies the access to bucket. Valid values are READ and WRITE.
	// Specifies the access to bucket. Valid values are **READ** and **WRITE**.
	AccessToBucket []*string `json:"accessToBucket,omitempty" tf:"access_to_bucket,omitempty"`
}

type LogDeliveryUserPermissionObservation struct {

	// Specifies the access to acl. Valid values are READ_ACP and WRITE_ACP.
	// Specifies the access to acl. Valid values are **READ_ACP** and **WRITE_ACP**.
	AccessToACL []*string `json:"accessToAcl,omitempty" tf:"access_to_acl,omitempty"`

	// Specifies the access to bucket. Valid values are READ and WRITE.
	// Specifies the access to bucket. Valid values are **READ** and **WRITE**.
	AccessToBucket []*string `json:"accessToBucket,omitempty" tf:"access_to_bucket,omitempty"`
}

type LogDeliveryUserPermissionParameters struct {

	// Specifies the access to acl. Valid values are READ_ACP and WRITE_ACP.
	// Specifies the access to acl. Valid values are **READ_ACP** and **WRITE_ACP**.
	// +kubebuilder:validation:Optional
	AccessToACL []*string `json:"accessToAcl,omitempty" tf:"access_to_acl,omitempty"`

	// Specifies the access to bucket. Valid values are READ and WRITE.
	// Specifies the access to bucket. Valid values are **READ** and **WRITE**.
	// +kubebuilder:validation:Optional
	AccessToBucket []*string `json:"accessToBucket,omitempty" tf:"access_to_bucket,omitempty"`
}

type OwnerPermissionInitParameters struct {

	// Specifies the access to acl. Valid values are READ_ACP and WRITE_ACP.
	// Specifies the access to acl. Valid values are **READ_ACP** and **WRITE_ACP**.
	AccessToACL []*string `json:"accessToAcl,omitempty" tf:"access_to_acl,omitempty"`

	// Specifies the access to bucket. Valid values are READ and WRITE.
	// Specifies the access to bucket. Valid values are **READ** and **WRITE**.
	AccessToBucket []*string `json:"accessToBucket,omitempty" tf:"access_to_bucket,omitempty"`
}

type OwnerPermissionObservation struct {

	// Specifies the access to acl. Valid values are READ_ACP and WRITE_ACP.
	// Specifies the access to acl. Valid values are **READ_ACP** and **WRITE_ACP**.
	AccessToACL []*string `json:"accessToAcl,omitempty" tf:"access_to_acl,omitempty"`

	// Specifies the access to bucket. Valid values are READ and WRITE.
	// Specifies the access to bucket. Valid values are **READ** and **WRITE**.
	AccessToBucket []*string `json:"accessToBucket,omitempty" tf:"access_to_bucket,omitempty"`
}

type OwnerPermissionParameters struct {

	// Specifies the access to acl. Valid values are READ_ACP and WRITE_ACP.
	// Specifies the access to acl. Valid values are **READ_ACP** and **WRITE_ACP**.
	// +kubebuilder:validation:Optional
	AccessToACL []*string `json:"accessToAcl,omitempty" tf:"access_to_acl,omitempty"`

	// Specifies the access to bucket. Valid values are READ and WRITE.
	// Specifies the access to bucket. Valid values are **READ** and **WRITE**.
	// +kubebuilder:validation:Optional
	AccessToBucket []*string `json:"accessToBucket,omitempty" tf:"access_to_bucket,omitempty"`
}

type PublicPermissionInitParameters struct {

	// Specifies the access to acl. Valid values are READ_ACP and WRITE_ACP.
	// Specifies the access to acl. Valid values are **READ_ACP** and **WRITE_ACP**.
	AccessToACL []*string `json:"accessToAcl,omitempty" tf:"access_to_acl,omitempty"`

	// Specifies the access to bucket. Valid values are READ and WRITE.
	// Specifies the access to bucket. Valid values are **READ** and **WRITE**.
	AccessToBucket []*string `json:"accessToBucket,omitempty" tf:"access_to_bucket,omitempty"`
}

type PublicPermissionObservation struct {

	// Specifies the access to acl. Valid values are READ_ACP and WRITE_ACP.
	// Specifies the access to acl. Valid values are **READ_ACP** and **WRITE_ACP**.
	AccessToACL []*string `json:"accessToAcl,omitempty" tf:"access_to_acl,omitempty"`

	// Specifies the access to bucket. Valid values are READ and WRITE.
	// Specifies the access to bucket. Valid values are **READ** and **WRITE**.
	AccessToBucket []*string `json:"accessToBucket,omitempty" tf:"access_to_bucket,omitempty"`
}

type PublicPermissionParameters struct {

	// Specifies the access to acl. Valid values are READ_ACP and WRITE_ACP.
	// Specifies the access to acl. Valid values are **READ_ACP** and **WRITE_ACP**.
	// +kubebuilder:validation:Optional
	AccessToACL []*string `json:"accessToAcl,omitempty" tf:"access_to_acl,omitempty"`

	// Specifies the access to bucket. Valid values are READ and WRITE.
	// Specifies the access to bucket. Valid values are **READ** and **WRITE**.
	// +kubebuilder:validation:Optional
	AccessToBucket []*string `json:"accessToBucket,omitempty" tf:"access_to_bucket,omitempty"`
}

// BucketACLSpec defines the desired state of BucketACL
type BucketACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketACLParameters `json:"forProvider"`
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
	InitProvider BucketACLInitParameters `json:"initProvider,omitempty"`
}

// BucketACLStatus defines the observed state of BucketACL.
type BucketACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BucketACL is the Schema for the BucketACLs API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type BucketACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketACLSpec   `json:"spec"`
	Status            BucketACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketACLList contains a list of BucketACLs
type BucketACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketACL `json:"items"`
}

// Repository type metadata.
var (
	BucketACL_Kind             = "BucketACL"
	BucketACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketACL_Kind}.String()
	BucketACL_KindAPIVersion   = BucketACL_Kind + "." + CRDGroupVersion.String()
	BucketACL_GroupVersionKind = CRDGroupVersion.WithKind(BucketACL_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketACL{}, &BucketACLList{})
}