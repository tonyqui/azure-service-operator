// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
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

var _ conversion.Convertible = &StorageAccountsQueueService{}

// ConvertFrom populates our StorageAccountsQueueService from the provided hub StorageAccountsQueueService
func (storageAccountsQueueService *StorageAccountsQueueService) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1alpha1api20210401storage.StorageAccountsQueueService)
	if !ok {
		return fmt.Errorf("expected storage:storage/v1alpha1api20210401storage/StorageAccountsQueueService but received %T instead", hub)
	}

	return storageAccountsQueueService.AssignPropertiesFromStorageAccountsQueueService(source)
}

// ConvertTo populates the provided hub StorageAccountsQueueService from our StorageAccountsQueueService
func (storageAccountsQueueService *StorageAccountsQueueService) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20210401storage.StorageAccountsQueueService)
	if !ok {
		return fmt.Errorf("expected storage:storage/v1alpha1api20210401storage/StorageAccountsQueueService but received %T instead", hub)
	}

	return storageAccountsQueueService.AssignPropertiesToStorageAccountsQueueService(destination)
}

// +kubebuilder:webhook:path=/mutate-storage-azure-com-v1alpha1api20210401-storageaccountsqueueservice,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountsqueueservices,verbs=create;update,versions=v1alpha1api20210401,name=default.v1alpha1api20210401.storageaccountsqueueservices.storage.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &StorageAccountsQueueService{}

// Default applies defaults to the StorageAccountsQueueService resource
func (storageAccountsQueueService *StorageAccountsQueueService) Default() {
	storageAccountsQueueService.defaultImpl()
	var temp interface{} = storageAccountsQueueService
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the StorageAccountsQueueService resource
func (storageAccountsQueueService *StorageAccountsQueueService) defaultImpl() {}

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

// +kubebuilder:webhook:path=/validate-storage-azure-com-v1alpha1api20210401-storageaccountsqueueservice,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountsqueueservices,verbs=create;update,versions=v1alpha1api20210401,name=validate.v1alpha1api20210401.storageaccountsqueueservices.storage.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &StorageAccountsQueueService{}

// ValidateCreate validates the creation of the resource
func (storageAccountsQueueService *StorageAccountsQueueService) ValidateCreate() error {
	validations := storageAccountsQueueService.createValidations()
	var temp interface{} = storageAccountsQueueService
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (storageAccountsQueueService *StorageAccountsQueueService) ValidateDelete() error {
	validations := storageAccountsQueueService.deleteValidations()
	var temp interface{} = storageAccountsQueueService
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (storageAccountsQueueService *StorageAccountsQueueService) ValidateUpdate(old runtime.Object) error {
	validations := storageAccountsQueueService.updateValidations()
	var temp interface{} = storageAccountsQueueService
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (storageAccountsQueueService *StorageAccountsQueueService) createValidations() []func() error {
	return []func() error{storageAccountsQueueService.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (storageAccountsQueueService *StorageAccountsQueueService) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (storageAccountsQueueService *StorageAccountsQueueService) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return storageAccountsQueueService.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (storageAccountsQueueService *StorageAccountsQueueService) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&storageAccountsQueueService.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromStorageAccountsQueueService populates our StorageAccountsQueueService from the provided source StorageAccountsQueueService
func (storageAccountsQueueService *StorageAccountsQueueService) AssignPropertiesFromStorageAccountsQueueService(source *v1alpha1api20210401storage.StorageAccountsQueueService) error {

	// ObjectMeta
	storageAccountsQueueService.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec StorageAccountsQueueServices_Spec
	err := spec.AssignPropertiesFromStorageAccountsQueueServicesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromStorageAccountsQueueServicesSpec() to populate field Spec")
	}
	storageAccountsQueueService.Spec = spec

	// Status
	var status QueueServiceProperties_Status
	err = status.AssignPropertiesFromQueueServicePropertiesStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromQueueServicePropertiesStatus() to populate field Status")
	}
	storageAccountsQueueService.Status = status

	// No error
	return nil
}

// AssignPropertiesToStorageAccountsQueueService populates the provided destination StorageAccountsQueueService from our StorageAccountsQueueService
func (storageAccountsQueueService *StorageAccountsQueueService) AssignPropertiesToStorageAccountsQueueService(destination *v1alpha1api20210401storage.StorageAccountsQueueService) error {

	// ObjectMeta
	destination.ObjectMeta = *storageAccountsQueueService.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20210401storage.StorageAccountsQueueServices_Spec
	err := storageAccountsQueueService.Spec.AssignPropertiesToStorageAccountsQueueServicesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToStorageAccountsQueueServicesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20210401storage.QueueServiceProperties_Status
	err = storageAccountsQueueService.Status.AssignPropertiesToQueueServicePropertiesStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToQueueServicePropertiesStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (storageAccountsQueueService *StorageAccountsQueueService) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: storageAccountsQueueService.Spec.OriginalVersion(),
		Kind:    "StorageAccountsQueueService",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/resourceDefinitions/storageAccounts_queueServices
type StorageAccountsQueueServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsQueueService `json:"items"`
}

type QueueServiceProperties_Status struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//Cors: Specifies CORS rules for the Queue service. You can include up to five
	//CorsRule elements in the request. If no CorsRule elements are included in the
	//request body, all CORS rules will be deleted, and CORS will be disabled for the
	//Queue service.
	Cors *CorsRules_Status `json:"cors,omitempty"`

	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
	//"Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &QueueServiceProperties_Status{}

// ConvertStatusFrom populates our QueueServiceProperties_Status from the provided source
func (queueServicePropertiesStatus *QueueServiceProperties_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20210401storage.QueueServiceProperties_Status)
	if ok {
		// Populate our instance from source
		return queueServicePropertiesStatus.AssignPropertiesFromQueueServicePropertiesStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210401storage.QueueServiceProperties_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = queueServicePropertiesStatus.AssignPropertiesFromQueueServicePropertiesStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our QueueServiceProperties_Status
func (queueServicePropertiesStatus *QueueServiceProperties_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20210401storage.QueueServiceProperties_Status)
	if ok {
		// Populate destination from our instance
		return queueServicePropertiesStatus.AssignPropertiesToQueueServicePropertiesStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210401storage.QueueServiceProperties_Status{}
	err := queueServicePropertiesStatus.AssignPropertiesToQueueServicePropertiesStatus(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &QueueServiceProperties_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (queueServicePropertiesStatus *QueueServiceProperties_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &QueueServiceProperties_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (queueServicePropertiesStatus *QueueServiceProperties_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(QueueServiceProperties_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected QueueServiceProperties_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Cors’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Cors != nil {
			var cors1 CorsRules_Status
			err := cors1.PopulateFromARM(owner, *typedInput.Properties.Cors)
			if err != nil {
				return err
			}
			cors := cors1
			queueServicePropertiesStatus.Cors = &cors
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		queueServicePropertiesStatus.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		queueServicePropertiesStatus.Name = &name
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		queueServicePropertiesStatus.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromQueueServicePropertiesStatus populates our QueueServiceProperties_Status from the provided source QueueServiceProperties_Status
func (queueServicePropertiesStatus *QueueServiceProperties_Status) AssignPropertiesFromQueueServicePropertiesStatus(source *v1alpha1api20210401storage.QueueServiceProperties_Status) error {

	// Conditions
	queueServicePropertiesStatus.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Cors
	if source.Cors != nil {
		var cor CorsRules_Status
		err := cor.AssignPropertiesFromCorsRulesStatus(source.Cors)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCorsRulesStatus() to populate field Cors")
		}
		queueServicePropertiesStatus.Cors = &cor
	} else {
		queueServicePropertiesStatus.Cors = nil
	}

	// Id
	queueServicePropertiesStatus.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	queueServicePropertiesStatus.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	queueServicePropertiesStatus.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToQueueServicePropertiesStatus populates the provided destination QueueServiceProperties_Status from our QueueServiceProperties_Status
func (queueServicePropertiesStatus *QueueServiceProperties_Status) AssignPropertiesToQueueServicePropertiesStatus(destination *v1alpha1api20210401storage.QueueServiceProperties_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(queueServicePropertiesStatus.Conditions)

	// Cors
	if queueServicePropertiesStatus.Cors != nil {
		var cor v1alpha1api20210401storage.CorsRules_Status
		err := queueServicePropertiesStatus.Cors.AssignPropertiesToCorsRulesStatus(&cor)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCorsRulesStatus() to populate field Cors")
		}
		destination.Cors = &cor
	} else {
		destination.Cors = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(queueServicePropertiesStatus.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(queueServicePropertiesStatus.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(queueServicePropertiesStatus.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"2021-04-01"}
type StorageAccountsQueueServicesSpecAPIVersion string

const StorageAccountsQueueServicesSpecAPIVersion20210401 = StorageAccountsQueueServicesSpecAPIVersion("2021-04-01")

type StorageAccountsQueueServices_Spec struct {
	//Cors: Sets the CORS rules. You can include up to five CorsRule elements in the
	//request.
	Cors *CorsRules `json:"cors,omitempty"`

	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner" kind:"StorageAccount"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &StorageAccountsQueueServices_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (storageAccountsQueueServicesSpec *StorageAccountsQueueServices_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if storageAccountsQueueServicesSpec == nil {
		return nil, nil
	}
	var result StorageAccountsQueueServices_SpecARM

	// Set property ‘Location’:
	if storageAccountsQueueServicesSpec.Location != nil {
		location := *storageAccountsQueueServicesSpec.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if storageAccountsQueueServicesSpec.Cors != nil {
		result.Properties = &QueueServicePropertiesPropertiesARM{}
	}
	if storageAccountsQueueServicesSpec.Cors != nil {
		corsARM, err := (*storageAccountsQueueServicesSpec.Cors).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		cors := corsARM.(CorsRulesARM)
		result.Properties.Cors = &cors
	}

	// Set property ‘Tags’:
	if storageAccountsQueueServicesSpec.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range storageAccountsQueueServicesSpec.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (storageAccountsQueueServicesSpec *StorageAccountsQueueServices_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &StorageAccountsQueueServices_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (storageAccountsQueueServicesSpec *StorageAccountsQueueServices_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(StorageAccountsQueueServices_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected StorageAccountsQueueServices_SpecARM, got %T", armInput)
	}

	// Set property ‘Cors’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Cors != nil {
			var cors1 CorsRules
			err := cors1.PopulateFromARM(owner, *typedInput.Properties.Cors)
			if err != nil {
				return err
			}
			cors := cors1
			storageAccountsQueueServicesSpec.Cors = &cors
		}
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		storageAccountsQueueServicesSpec.Location = &location
	}

	// Set property ‘Owner’:
	storageAccountsQueueServicesSpec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		storageAccountsQueueServicesSpec.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			storageAccountsQueueServicesSpec.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &StorageAccountsQueueServices_Spec{}

// ConvertSpecFrom populates our StorageAccountsQueueServices_Spec from the provided source
func (storageAccountsQueueServicesSpec *StorageAccountsQueueServices_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20210401storage.StorageAccountsQueueServices_Spec)
	if ok {
		// Populate our instance from source
		return storageAccountsQueueServicesSpec.AssignPropertiesFromStorageAccountsQueueServicesSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210401storage.StorageAccountsQueueServices_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = storageAccountsQueueServicesSpec.AssignPropertiesFromStorageAccountsQueueServicesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our StorageAccountsQueueServices_Spec
func (storageAccountsQueueServicesSpec *StorageAccountsQueueServices_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20210401storage.StorageAccountsQueueServices_Spec)
	if ok {
		// Populate destination from our instance
		return storageAccountsQueueServicesSpec.AssignPropertiesToStorageAccountsQueueServicesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210401storage.StorageAccountsQueueServices_Spec{}
	err := storageAccountsQueueServicesSpec.AssignPropertiesToStorageAccountsQueueServicesSpec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromStorageAccountsQueueServicesSpec populates our StorageAccountsQueueServices_Spec from the provided source StorageAccountsQueueServices_Spec
func (storageAccountsQueueServicesSpec *StorageAccountsQueueServices_Spec) AssignPropertiesFromStorageAccountsQueueServicesSpec(source *v1alpha1api20210401storage.StorageAccountsQueueServices_Spec) error {

	// Cors
	if source.Cors != nil {
		var cor CorsRules
		err := cor.AssignPropertiesFromCorsRules(source.Cors)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCorsRules() to populate field Cors")
		}
		storageAccountsQueueServicesSpec.Cors = &cor
	} else {
		storageAccountsQueueServicesSpec.Cors = nil
	}

	// Location
	storageAccountsQueueServicesSpec.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	storageAccountsQueueServicesSpec.Owner = source.Owner.Copy()

	// Tags
	storageAccountsQueueServicesSpec.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToStorageAccountsQueueServicesSpec populates the provided destination StorageAccountsQueueServices_Spec from our StorageAccountsQueueServices_Spec
func (storageAccountsQueueServicesSpec *StorageAccountsQueueServices_Spec) AssignPropertiesToStorageAccountsQueueServicesSpec(destination *v1alpha1api20210401storage.StorageAccountsQueueServices_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Cors
	if storageAccountsQueueServicesSpec.Cors != nil {
		var cor v1alpha1api20210401storage.CorsRules
		err := storageAccountsQueueServicesSpec.Cors.AssignPropertiesToCorsRules(&cor)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCorsRules() to populate field Cors")
		}
		destination.Cors = &cor
	} else {
		destination.Cors = nil
	}

	// Location
	destination.Location = genruntime.ClonePointerToString(storageAccountsQueueServicesSpec.Location)

	// OriginalVersion
	destination.OriginalVersion = storageAccountsQueueServicesSpec.OriginalVersion()

	// Owner
	destination.Owner = storageAccountsQueueServicesSpec.Owner.Copy()

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(storageAccountsQueueServicesSpec.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (storageAccountsQueueServicesSpec *StorageAccountsQueueServices_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

func init() {
	SchemeBuilder.Register(&StorageAccountsQueueService{}, &StorageAccountsQueueServiceList{})
}