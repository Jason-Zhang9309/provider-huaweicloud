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

type ImageAutoSyncInitParameters struct {

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

	// Specifies whether to overwrite.
	// Default to false, which indicates not to overwrite
	// any nonidentical image that has the same name in the target organization.
	// Specifies whether to overwrite.
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`

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

	// Specifies the target organization name.
	// Specifies the target organization name.
	TargetOrganization *string `json:"targetOrganization,omitempty" tf:"target_organization,omitempty"`

	// Specifies the target region name.
	// Specifies the target region name.
	TargetRegion *string `json:"targetRegion,omitempty" tf:"target_region,omitempty"`
}

type ImageAutoSyncObservation struct {

	// Indicates the creation time.
	// Indicates the creation time.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the organization.
	// Specifies the name of the organization.
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// Specifies whether to overwrite.
	// Default to false, which indicates not to overwrite
	// any nonidentical image that has the same name in the target organization.
	// Specifies whether to overwrite.
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`

	// Specifies the region in which to create the resource.
	// If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the name of the repository.
	// Specifies the name of the repository.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Specifies the target organization name.
	// Specifies the target organization name.
	TargetOrganization *string `json:"targetOrganization,omitempty" tf:"target_organization,omitempty"`

	// Specifies the target region name.
	// Specifies the target region name.
	TargetRegion *string `json:"targetRegion,omitempty" tf:"target_region,omitempty"`

	// Indicates the update time.
	// Indicates the update time.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ImageAutoSyncParameters struct {

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

	// Specifies whether to overwrite.
	// Default to false, which indicates not to overwrite
	// any nonidentical image that has the same name in the target organization.
	// Specifies whether to overwrite.
	// +kubebuilder:validation:Optional
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`

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

	// Specifies the target organization name.
	// Specifies the target organization name.
	// +kubebuilder:validation:Optional
	TargetOrganization *string `json:"targetOrganization,omitempty" tf:"target_organization,omitempty"`

	// Specifies the target region name.
	// Specifies the target region name.
	// +kubebuilder:validation:Optional
	TargetRegion *string `json:"targetRegion,omitempty" tf:"target_region,omitempty"`
}

// ImageAutoSyncSpec defines the desired state of ImageAutoSync
type ImageAutoSyncSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageAutoSyncParameters `json:"forProvider"`
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
	InitProvider ImageAutoSyncInitParameters `json:"initProvider,omitempty"`
}

// ImageAutoSyncStatus defines the observed state of ImageAutoSync.
type ImageAutoSyncStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageAutoSyncObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ImageAutoSync is the Schema for the ImageAutoSyncs API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type ImageAutoSync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetOrganization) || (has(self.initProvider) && has(self.initProvider.targetOrganization))",message="spec.forProvider.targetOrganization is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetRegion) || (has(self.initProvider) && has(self.initProvider.targetRegion))",message="spec.forProvider.targetRegion is a required parameter"
	Spec   ImageAutoSyncSpec   `json:"spec"`
	Status ImageAutoSyncStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageAutoSyncList contains a list of ImageAutoSyncs
type ImageAutoSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageAutoSync `json:"items"`
}

// Repository type metadata.
var (
	ImageAutoSync_Kind             = "ImageAutoSync"
	ImageAutoSync_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImageAutoSync_Kind}.String()
	ImageAutoSync_KindAPIVersion   = ImageAutoSync_Kind + "." + CRDGroupVersion.String()
	ImageAutoSync_GroupVersionKind = CRDGroupVersion.WithKind(ImageAutoSync_Kind)
)

func init() {
	SchemeBuilder.Register(&ImageAutoSync{}, &ImageAutoSyncList{})
}
