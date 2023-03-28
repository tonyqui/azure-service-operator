// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Servers_Databases_TransparentDataEncryption_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *TransparentDataEncryptionProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Servers_Databases_TransparentDataEncryption_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (encryption Servers_Databases_TransparentDataEncryption_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (encryption *Servers_Databases_TransparentDataEncryption_Spec_ARM) GetName() string {
	return encryption.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases/transparentDataEncryption"
func (encryption *Servers_Databases_TransparentDataEncryption_Spec_ARM) GetType() string {
	return "Microsoft.Sql/servers/databases/transparentDataEncryption"
}

// Properties of a transparent data encryption.
type TransparentDataEncryptionProperties_ARM struct {
	// State: Specifies the state of the transparent data encryption.
	State *TransparentDataEncryptionProperties_State `json:"state,omitempty"`
}