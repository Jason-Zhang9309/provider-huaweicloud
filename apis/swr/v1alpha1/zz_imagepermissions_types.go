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

type ImagePermissionsInitParameters struct {

	// Specifies the name of the organization.
	// Specifies the name of the organization.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/swr/v1alpha1.Organization
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// Reference to a Organization in swr to populate organization.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// Selector for a Organization in swr to populate organization.
	// +kubebuilder:validation:Optional
	OrganizationSelector *v1.Selector `json:"organizationSelector,omitempty" tf:"-"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the name of the repository.
	// Specifies the name of the repository.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/swr/v1alpha1.Repository
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in swr to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in swr to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`

	// Specifies the users to access to the image (repository).
	// The User structure is documented below.
	// Specifies the users to access to the image (repository).
	Users []UsersInitParameters `json:"users,omitempty" tf:"users,omitempty"`
}

type ImagePermissionsObservation struct {

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the organization.
	// Specifies the name of the organization.
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the name of the repository.
	// Specifies the name of the repository.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Indicates the permission information of current user.
	// The SelfPermission structure is documented below.
	// Indicates the permission information of current user.
	SelfPermission []SelfPermissionObservation `json:"selfPermission,omitempty" tf:"self_permission,omitempty"`

	// Specifies the users to access to the image (repository).
	// The User structure is documented below.
	// Specifies the users to access to the image (repository).
	Users []UsersObservation `json:"users,omitempty" tf:"users,omitempty"`
}

type ImagePermissionsParameters struct {

	// Specifies the name of the organization.
	// Specifies the name of the organization.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/swr/v1alpha1.Organization
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// Reference to a Organization in swr to populate organization.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// Selector for a Organization in swr to populate organization.
	// +kubebuilder:validation:Optional
	OrganizationSelector *v1.Selector `json:"organizationSelector,omitempty" tf:"-"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the name of the repository.
	// Specifies the name of the repository.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/swr/v1alpha1.Repository
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in swr to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in swr to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`

	// Specifies the users to access to the image (repository).
	// The User structure is documented below.
	// Specifies the users to access to the image (repository).
	// +kubebuilder:validation:Optional
	Users []UsersParameters `json:"users,omitempty" tf:"users,omitempty"`
}

type SelfPermissionInitParameters struct {
}

type SelfPermissionObservation struct {

	// Specifies the user permission of the existing HuaweiCloud user.
	// The values can be Manage, Write and Read.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// Specifies the ID of the existing HuaweiCloud user.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Specifies the name of the existing HuaweiCloud user.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type SelfPermissionParameters struct {
}

type UsersInitParameters struct {

	// Specifies the user permission of the existing HuaweiCloud user.
	// The values can be Manage, Write and Read.
	// Specifies the user permission of the existing HuaweiCloud user.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// Specifies the ID of the existing HuaweiCloud user.
	// Specifies the ID of the existing HuaweiCloud user.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/iam/v1alpha1.User
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User in iam to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User in iam to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`

	// Specifies the name of the existing HuaweiCloud user.
	// Specifies the name of the existing HuaweiCloud user.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type UsersObservation struct {

	// Specifies the user permission of the existing HuaweiCloud user.
	// The values can be Manage, Write and Read.
	// Specifies the user permission of the existing HuaweiCloud user.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// Specifies the ID of the existing HuaweiCloud user.
	// Specifies the ID of the existing HuaweiCloud user.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Specifies the name of the existing HuaweiCloud user.
	// Specifies the name of the existing HuaweiCloud user.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type UsersParameters struct {

	// Specifies the user permission of the existing HuaweiCloud user.
	// The values can be Manage, Write and Read.
	// Specifies the user permission of the existing HuaweiCloud user.
	// +kubebuilder:validation:Optional
	Permission *string `json:"permission" tf:"permission,omitempty"`

	// Specifies the ID of the existing HuaweiCloud user.
	// Specifies the ID of the existing HuaweiCloud user.
	// +crossplane:generate:reference:type=github.com/huaweicloud/provider-huaweicloud/apis/iam/v1alpha1.User
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User in iam to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User in iam to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`

	// Specifies the name of the existing HuaweiCloud user.
	// Specifies the name of the existing HuaweiCloud user.
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

// ImagePermissionsSpec defines the desired state of ImagePermissions
type ImagePermissionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImagePermissionsParameters `json:"forProvider"`
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
	InitProvider ImagePermissionsInitParameters `json:"initProvider,omitempty"`
}

// ImagePermissionsStatus defines the observed state of ImagePermissions.
type ImagePermissionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImagePermissionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ImagePermissions is the Schema for the ImagePermissionss API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type ImagePermissions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.users) || (has(self.initProvider) && has(self.initProvider.users))",message="spec.forProvider.users is a required parameter"
	Spec   ImagePermissionsSpec   `json:"spec"`
	Status ImagePermissionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImagePermissionsList contains a list of ImagePermissionss
type ImagePermissionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImagePermissions `json:"items"`
}

// Repository type metadata.
var (
	ImagePermissions_Kind             = "ImagePermissions"
	ImagePermissions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImagePermissions_Kind}.String()
	ImagePermissions_KindAPIVersion   = ImagePermissions_Kind + "." + CRDGroupVersion.String()
	ImagePermissions_GroupVersionKind = CRDGroupVersion.WithKind(ImagePermissions_Kind)
)

func init() {
	SchemeBuilder.Register(&ImagePermissions{}, &ImagePermissionsList{})
}
