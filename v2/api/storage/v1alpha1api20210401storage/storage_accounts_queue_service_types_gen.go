// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=storage.azure.com,resources=storageaccountsqueueservices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=storage.azure.com,resources={storageaccountsqueueservices/status,storageaccountsqueueservices/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210401.StorageAccountsQueueService
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/resourceDefinitions/storageAccounts_queueServices
type StorageAccountsQueueService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountsQueueServices_Spec `json:"spec,omitempty"`
	Status            QueueServiceProperties_Status     `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsQueueService{}

// GetConditions returns the conditions of the resource
func (storageAccountsQueueService *StorageAccountsQueueService) GetConditions() conditions.Conditions {
	return storageAccountsQueueService.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (storageAccountsQueueService *StorageAccountsQueueService) SetConditions(conditions conditions.Conditions) {
	storageAccountsQueueService.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &StorageAccountsQueueService{}

// AzureName returns the Azure name of the resource (always "default")
func (storageAccountsQueueService *StorageAccountsQueueService) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (storageAccountsQueueService StorageAccountsQueueService) GetAPIVersion() string {
	return "2021-04-01"
}

// GetResourceKind returns the kind of the resource
func (storageAccountsQueueService *StorageAccountsQueueService) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (storageAccountsQueueService *StorageAccountsQueueService) GetSpec() genruntime.ConvertibleSpec {
	return &storageAccountsQueueService.Spec
}

// GetStatus returns the status of this resource
func (storageAccountsQueueService *StorageAccountsQueueService) GetStatus() genruntime.ConvertibleStatus {
	return &storageAccountsQueueService.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/queueServices"
func (storageAccountsQueueService *StorageAccountsQueueService) GetType() string {
	return "Microsoft.Storage/storageAccounts/queueServices"
}

// NewEmptyStatus returns a new empty (blank) status
func (storageAccountsQueueService *StorageAccountsQueueService) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &QueueServiceProperties_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (storageAccountsQueueService *StorageAccountsQueueService) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(storageAccountsQueueService.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  storageAccountsQueueService.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (storageAccountsQueueService *StorageAccountsQueueService) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*QueueServiceProperties_Status); ok {
		storageAccountsQueueService.Status = *st
		return nil
	}

	// Convert status to required version
	var st QueueServiceProperties_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	storageAccountsQueueService.Status = st
	return nil
}

// Hub marks that this StorageAccountsQueueService is the hub type for conversion
func (storageAccountsQueueService *StorageAccountsQueueService) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (storageAccountsQueueService *StorageAccountsQueueService) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: storageAccountsQueueService.Spec.OriginalVersion,
		Kind:    "StorageAccountsQueueService",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210401.StorageAccountsQueueService
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/resourceDefinitions/storageAccounts_queueServices
type StorageAccountsQueueServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsQueueService `json:"items"`
}

//Storage version of v1alpha1api20210401.QueueServiceProperties_Status
type QueueServiceProperties_Status struct {
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Cors        *CorsRules_Status      `json:"cors,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &QueueServiceProperties_Status{}

// ConvertStatusFrom populates our QueueServiceProperties_Status from the provided source
func (queueServicePropertiesStatus *QueueServiceProperties_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == queueServicePropertiesStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(queueServicePropertiesStatus)
}

// ConvertStatusTo populates the provided destination from our QueueServiceProperties_Status
func (queueServicePropertiesStatus *QueueServiceProperties_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == queueServicePropertiesStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(queueServicePropertiesStatus)
}

//Storage version of v1alpha1api20210401.StorageAccountsQueueServices_Spec
type StorageAccountsQueueServices_Spec struct {
	Cors            *CorsRules `json:"cors,omitempty"`
	Location        *string    `json:"location,omitempty"`
	OriginalVersion string     `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner" kind:"StorageAccount"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccountsQueueServices_Spec{}

// ConvertSpecFrom populates our StorageAccountsQueueServices_Spec from the provided source
func (storageAccountsQueueServicesSpec *StorageAccountsQueueServices_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == storageAccountsQueueServicesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(storageAccountsQueueServicesSpec)
}

// ConvertSpecTo populates the provided destination from our StorageAccountsQueueServices_Spec
func (storageAccountsQueueServicesSpec *StorageAccountsQueueServices_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == storageAccountsQueueServicesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(storageAccountsQueueServicesSpec)
}

func init() {
	SchemeBuilder.Register(&StorageAccountsQueueService{}, &StorageAccountsQueueServiceList{})
}