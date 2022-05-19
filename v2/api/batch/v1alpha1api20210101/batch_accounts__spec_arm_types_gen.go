// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of BatchAccounts_Spec. Use v1beta20210101.BatchAccounts_Spec instead
type BatchAccounts_SpecARM struct {
	Identity   *BatchAccountIdentityARM         `json:"identity,omitempty"`
	Location   *string                          `json:"location,omitempty"`
	Name       string                           `json:"name,omitempty"`
	Properties *BatchAccountCreatePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &BatchAccounts_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01"
func (accounts BatchAccounts_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (accounts BatchAccounts_SpecARM) GetName() string {
	return accounts.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Batch/batchAccounts"
func (accounts BatchAccounts_SpecARM) GetType() string {
	return "Microsoft.Batch/batchAccounts"
}

// Deprecated version of BatchAccountCreateProperties. Use v1beta20210101.BatchAccountCreateProperties instead
type BatchAccountCreatePropertiesARM struct {
	AutoStorage         *AutoStorageBasePropertiesARM                    `json:"autoStorage,omitempty"`
	Encryption          *EncryptionPropertiesARM                         `json:"encryption,omitempty"`
	KeyVaultReference   *KeyVaultReferenceARM                            `json:"keyVaultReference,omitempty"`
	PoolAllocationMode  *BatchAccountCreatePropertiesPoolAllocationMode  `json:"poolAllocationMode,omitempty"`
	PublicNetworkAccess *BatchAccountCreatePropertiesPublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
}

// Deprecated version of BatchAccountIdentity. Use v1beta20210101.BatchAccountIdentity instead
type BatchAccountIdentityARM struct {
	Type *BatchAccountIdentityType `json:"type,omitempty"`
}

// Deprecated version of AutoStorageBaseProperties. Use v1beta20210101.AutoStorageBaseProperties instead
type AutoStorageBasePropertiesARM struct {
	StorageAccountId *string `json:"storageAccountId,omitempty"`
}

// Deprecated version of BatchAccountIdentityType. Use v1beta20210101.BatchAccountIdentityType instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","UserAssigned"}
type BatchAccountIdentityType string

const (
	BatchAccountIdentityTypeNone           = BatchAccountIdentityType("None")
	BatchAccountIdentityTypeSystemAssigned = BatchAccountIdentityType("SystemAssigned")
	BatchAccountIdentityTypeUserAssigned   = BatchAccountIdentityType("UserAssigned")
)

// Deprecated version of EncryptionProperties. Use v1beta20210101.EncryptionProperties instead
type EncryptionPropertiesARM struct {
	KeySource          *EncryptionPropertiesKeySource `json:"keySource,omitempty"`
	KeyVaultProperties *KeyVaultPropertiesARM         `json:"keyVaultProperties,omitempty"`
}

// Deprecated version of KeyVaultReference. Use v1beta20210101.KeyVaultReference instead
type KeyVaultReferenceARM struct {
	Id  *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
}

// Deprecated version of KeyVaultProperties. Use v1beta20210101.KeyVaultProperties instead
type KeyVaultPropertiesARM struct {
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
}
