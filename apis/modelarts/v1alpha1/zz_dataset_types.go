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

type DataSourceInitParameters struct {

	// Specifies the cluster ID of DWS/MRS when data_type is 1 or 4.
	// Changing this parameter will create a new resource.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Specifies the type of data source. The options are as follows:
	DataType *float64 `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// Specifies the database name of DWS/DLI when data_type is 1 or 2.
	// Changing this parameter will create a new resource.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Specifies the password of database when data_type is 1.
	// Changing this parameter will create a new resource.
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Specifies the OBS path when data_type is 0
	// or the hdsf path when data_type is 4. All the file in this directory and subdirectories will be which be imported
	// to the dataset. Changing this parameter will create a new resource.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Specifies the queue name of DLI when data_type is 2.
	// Changing this parameter will create a new resource.
	QueueName *string `json:"queueName,omitempty" tf:"queue_name,omitempty"`

	// Specifies the table name of DWS/DLI when data_type is 1 or 2.
	// Changing this parameter will create a new resource.
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`

	// Specifies the user name of database when data_type is 1.
	// Changing this parameter will create a new resource.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// Specifies whether the data contains table header when the type
	// of dataset is 400(Table type). Default value is true. Changing this parameter will create a new resource.
	WithColumnHeader *bool `json:"withColumnHeader,omitempty" tf:"with_column_header,omitempty"`
}

type DataSourceObservation struct {

	// Specifies the cluster ID of DWS/MRS when data_type is 1 or 4.
	// Changing this parameter will create a new resource.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Specifies the type of data source. The options are as follows:
	DataType *float64 `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// Specifies the database name of DWS/DLI when data_type is 1 or 2.
	// Changing this parameter will create a new resource.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Specifies the OBS path when data_type is 0
	// or the hdsf path when data_type is 4. All the file in this directory and subdirectories will be which be imported
	// to the dataset. Changing this parameter will create a new resource.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Specifies the queue name of DLI when data_type is 2.
	// Changing this parameter will create a new resource.
	QueueName *string `json:"queueName,omitempty" tf:"queue_name,omitempty"`

	// Specifies the table name of DWS/DLI when data_type is 1 or 2.
	// Changing this parameter will create a new resource.
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`

	// Specifies the user name of database when data_type is 1.
	// Changing this parameter will create a new resource.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// Specifies whether the data contains table header when the type
	// of dataset is 400(Table type). Default value is true. Changing this parameter will create a new resource.
	WithColumnHeader *bool `json:"withColumnHeader,omitempty" tf:"with_column_header,omitempty"`
}

type DataSourceParameters struct {

	// Specifies the cluster ID of DWS/MRS when data_type is 1 or 4.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Specifies the type of data source. The options are as follows:
	// +kubebuilder:validation:Optional
	DataType *float64 `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// Specifies the database name of DWS/DLI when data_type is 1 or 2.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Specifies the password of database when data_type is 1.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Specifies the OBS path when data_type is 0
	// or the hdsf path when data_type is 4. All the file in this directory and subdirectories will be which be imported
	// to the dataset. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Specifies the queue name of DLI when data_type is 2.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	QueueName *string `json:"queueName,omitempty" tf:"queue_name,omitempty"`

	// Specifies the table name of DWS/DLI when data_type is 1 or 2.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`

	// Specifies the user name of database when data_type is 1.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// Specifies whether the data contains table header when the type
	// of dataset is 400(Table type). Default value is true. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	WithColumnHeader *bool `json:"withColumnHeader,omitempty" tf:"with_column_header,omitempty"`
}

type DatasetInitParameters struct {

	// Specifies the data sources which be used to imported the source data (such
	// as pictures/files/audio, etc.) in this directory and subdirectories to the dataset. Structure is documented below.
	// Changing this parameter will create a new resource.
	DataSource []DataSourceInitParameters `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// Specifies the description of dataset. It contains a maximum of 256 characters and
	// cannot contain special characters !<>=&"'.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether to import labeled files.
	// Default value is true. Changing this parameter will create a new resource.
	ImportLabeledEnabled *bool `json:"importLabeledEnabled,omitempty" tf:"import_labeled_enabled,omitempty"`

	// Specifies the custom format information of labeled files when import
	// labeled files for Text classification. Structure is documented below.
	// Changing this parameter will create a new resource.
	// It is required only the dataType=100
	LabelFormat []LabelFormatInitParameters `json:"labelFormat,omitempty" tf:"label_format,omitempty"`

	// Specifies labels information. Structure is documented below.
	Labels []LabelsInitParameters `json:"labels,omitempty" tf:"labels,omitempty"`

	// Specifies the name of the dataset. The name consists of 1 to 100 characters,
	// starting with a letter. Only letters, chinese characters, digits underscores (_) and hyphens (-) are allowed.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the OBS path for storing output files such as labeled files.
	// The path cannot be the same as the import path or subdirectory of the import path.
	// Changing this parameter will create a new resource.
	OutputPath *string `json:"outputPath,omitempty" tf:"output_path,omitempty"`

	// The region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the schema information of source data when type is 400.
	// Structure is documented below. Changing this parameter will create a new resource.
	Schemas []SchemasInitParameters `json:"schemas,omitempty" tf:"schemas,omitempty"`

	// Specifies the type of dataset. The options are as follows:
	Type *float64 `json:"type,omitempty" tf:"type,omitempty"`
}

type DatasetObservation struct {

	// The dataset creation time.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// dataset format. Valid values include: Default, CarbonData: Carbon format(Supported only for
	// table type datasets).
	DataFormat *string `json:"dataFormat,omitempty" tf:"data_format,omitempty"`

	// Specifies the data sources which be used to imported the source data (such
	// as pictures/files/audio, etc.) in this directory and subdirectories to the dataset. Structure is documented below.
	// Changing this parameter will create a new resource.
	DataSource []DataSourceObservation `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// Specifies the description of dataset. It contains a maximum of 256 characters and
	// cannot contain special characters !<>=&"'.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The resource ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies whether to import labeled files.
	// Default value is true. Changing this parameter will create a new resource.
	ImportLabeledEnabled *bool `json:"importLabeledEnabled,omitempty" tf:"import_labeled_enabled,omitempty"`

	// Specifies the custom format information of labeled files when import
	// labeled files for Text classification. Structure is documented below.
	// Changing this parameter will create a new resource.
	// It is required only the dataType=100
	LabelFormat []LabelFormatObservation `json:"labelFormat,omitempty" tf:"label_format,omitempty"`

	// Specifies labels information. Structure is documented below.
	Labels []LabelsObservation `json:"labels,omitempty" tf:"labels,omitempty"`

	// Specifies the name of the dataset. The name consists of 1 to 100 characters,
	// starting with a letter. Only letters, chinese characters, digits underscores (_) and hyphens (-) are allowed.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the OBS path for storing output files such as labeled files.
	// The path cannot be the same as the import path or subdirectory of the import path.
	// Changing this parameter will create a new resource.
	OutputPath *string `json:"outputPath,omitempty" tf:"output_path,omitempty"`

	// The region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the schema information of source data when type is 400.
	// Structure is documented below. Changing this parameter will create a new resource.
	Schemas []SchemasObservation `json:"schemas,omitempty" tf:"schemas,omitempty"`

	// Dataset status. Valid values are as follows:
	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the type of dataset. The options are as follows:
	Type *float64 `json:"type,omitempty" tf:"type,omitempty"`
}

type DatasetParameters struct {

	// Specifies the data sources which be used to imported the source data (such
	// as pictures/files/audio, etc.) in this directory and subdirectories to the dataset. Structure is documented below.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	DataSource []DataSourceParameters `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// Specifies the description of dataset. It contains a maximum of 256 characters and
	// cannot contain special characters !<>=&"'.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether to import labeled files.
	// Default value is true. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	ImportLabeledEnabled *bool `json:"importLabeledEnabled,omitempty" tf:"import_labeled_enabled,omitempty"`

	// Specifies the custom format information of labeled files when import
	// labeled files for Text classification. Structure is documented below.
	// Changing this parameter will create a new resource.
	// It is required only the dataType=100
	// +kubebuilder:validation:Optional
	LabelFormat []LabelFormatParameters `json:"labelFormat,omitempty" tf:"label_format,omitempty"`

	// Specifies labels information. Structure is documented below.
	// +kubebuilder:validation:Optional
	Labels []LabelsParameters `json:"labels,omitempty" tf:"labels,omitempty"`

	// Specifies the name of the dataset. The name consists of 1 to 100 characters,
	// starting with a letter. Only letters, chinese characters, digits underscores (_) and hyphens (-) are allowed.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the OBS path for storing output files such as labeled files.
	// The path cannot be the same as the import path or subdirectory of the import path.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	OutputPath *string `json:"outputPath,omitempty" tf:"output_path,omitempty"`

	// The region in which to create the resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the schema information of source data when type is 400.
	// Structure is documented below. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Schemas []SchemasParameters `json:"schemas,omitempty" tf:"schemas,omitempty"`

	// Specifies the type of dataset. The options are as follows:
	// +kubebuilder:validation:Optional
	Type *float64 `json:"type,omitempty" tf:"type,omitempty"`
}

type LabelFormatInitParameters struct {

	// Specifies the separator between label and label.
	// Changing this parameter will create a new resource.
	LabelSeparator *string `json:"labelSeparator,omitempty" tf:"label_separator,omitempty"`

	// Specifies the separator between text and label.
	// Changing this parameter will create a new resource.
	TextLabelSeparator *string `json:"textLabelSeparator,omitempty" tf:"text_label_separator,omitempty"`

	// Specifies Label type for text classification.
	// The optional values are as follows:
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LabelFormatObservation struct {

	// Specifies the separator between label and label.
	// Changing this parameter will create a new resource.
	LabelSeparator *string `json:"labelSeparator,omitempty" tf:"label_separator,omitempty"`

	// Specifies the separator between text and label.
	// Changing this parameter will create a new resource.
	TextLabelSeparator *string `json:"textLabelSeparator,omitempty" tf:"text_label_separator,omitempty"`

	// Specifies Label type for text classification.
	// The optional values are as follows:
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LabelFormatParameters struct {

	// Specifies the separator between label and label.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	LabelSeparator *string `json:"labelSeparator,omitempty" tf:"label_separator,omitempty"`

	// Specifies the separator between text and label.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	TextLabelSeparator *string `json:"textLabelSeparator,omitempty" tf:"text_label_separator,omitempty"`

	// Specifies Label type for text classification.
	// The optional values are as follows:
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LabelsInitParameters struct {

	// Specifies the name of label.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies color of label.
	PropertyColor *string `json:"propertyColor,omitempty" tf:"property_color,omitempty"`

	// Specifies shape of label. Valid values include: bndbox, polygon,
	// circle, line, dashed, point, polyline.
	PropertyShape *string `json:"propertyShape,omitempty" tf:"property_shape,omitempty"`

	// Specifies shortcut of label.
	PropertyShortcut *string `json:"propertyShortcut,omitempty" tf:"property_shortcut,omitempty"`
}

type LabelsObservation struct {

	// Specifies the name of label.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies color of label.
	PropertyColor *string `json:"propertyColor,omitempty" tf:"property_color,omitempty"`

	// Specifies shape of label. Valid values include: bndbox, polygon,
	// circle, line, dashed, point, polyline.
	PropertyShape *string `json:"propertyShape,omitempty" tf:"property_shape,omitempty"`

	// Specifies shortcut of label.
	PropertyShortcut *string `json:"propertyShortcut,omitempty" tf:"property_shortcut,omitempty"`
}

type LabelsParameters struct {

	// Specifies the name of label.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies color of label.
	// +kubebuilder:validation:Optional
	PropertyColor *string `json:"propertyColor,omitempty" tf:"property_color,omitempty"`

	// Specifies shape of label. Valid values include: bndbox, polygon,
	// circle, line, dashed, point, polyline.
	// +kubebuilder:validation:Optional
	PropertyShape *string `json:"propertyShape,omitempty" tf:"property_shape,omitempty"`

	// Specifies shortcut of label.
	// +kubebuilder:validation:Optional
	PropertyShortcut *string `json:"propertyShortcut,omitempty" tf:"property_shortcut,omitempty"`
}

type SchemasInitParameters struct {

	// Specifies the field name. Changing this parameter will create a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the field type. Valid values include: String, Short, Int,
	// Long, Double, Float, Byte, Date, Timestamp, Bool. Changing this parameter will create a new resource.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SchemasObservation struct {

	// Specifies the field name. Changing this parameter will create a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the field type. Valid values include: String, Short, Int,
	// Long, Double, Float, Byte, Date, Timestamp, Bool. Changing this parameter will create a new resource.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SchemasParameters struct {

	// Specifies the field name. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the field type. Valid values include: String, Short, Int,
	// Long, Double, Float, Byte, Date, Timestamp, Bool. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// DatasetSpec defines the desired state of Dataset
type DatasetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatasetParameters `json:"forProvider"`
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
	InitProvider DatasetInitParameters `json:"initProvider,omitempty"`
}

// DatasetStatus defines the observed state of Dataset.
type DatasetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatasetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Dataset is the Schema for the Datasets API. ""
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,huaweicloud}
type Dataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataSource) || (has(self.initProvider) && has(self.initProvider.dataSource))",message="spec.forProvider.dataSource is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.outputPath) || (has(self.initProvider) && has(self.initProvider.outputPath))",message="spec.forProvider.outputPath is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   DatasetSpec   `json:"spec"`
	Status DatasetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasetList contains a list of Datasets
type DatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dataset `json:"items"`
}

// Repository type metadata.
var (
	Dataset_Kind             = "Dataset"
	Dataset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Dataset_Kind}.String()
	Dataset_KindAPIVersion   = Dataset_Kind + "." + CRDGroupVersion.String()
	Dataset_GroupVersionKind = CRDGroupVersion.WithKind(Dataset_Kind)
)

func init() {
	SchemeBuilder.Register(&Dataset{}, &DatasetList{})
}