// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Redis_PatchSchedule_Spec. Use v1beta20201201.Redis_PatchSchedule_Spec instead
type Redis_PatchSchedule_SpecARM struct {
	Location   *string             `json:"location,omitempty"`
	Name       string              `json:"name,omitempty"`
	Properties *ScheduleEntriesARM `json:"properties,omitempty"`
	Tags       map[string]string   `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Redis_PatchSchedule_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (schedule Redis_PatchSchedule_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (schedule *Redis_PatchSchedule_SpecARM) GetName() string {
	return schedule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/patchSchedules"
func (schedule *Redis_PatchSchedule_SpecARM) GetType() string {
	return "Microsoft.Cache/redis/patchSchedules"
}

// Deprecated version of ScheduleEntries. Use v1beta20201201.ScheduleEntries instead
type ScheduleEntriesARM struct {
	ScheduleEntries []ScheduleEntryARM `json:"scheduleEntries,omitempty"`
}

// Deprecated version of ScheduleEntry. Use v1beta20201201.ScheduleEntry instead
type ScheduleEntryARM struct {
	DayOfWeek         *ScheduleEntry_DayOfWeek `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                  `json:"maintenanceWindow,omitempty"`
	StartHourUtc      *int                     `json:"startHourUtc,omitempty"`
}