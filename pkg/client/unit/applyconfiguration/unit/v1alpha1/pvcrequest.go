/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "tesseract/pkg/apis/unit/v1alpha1"

	v1 "k8s.io/api/core/v1"
)

// PVCRequestApplyConfiguration represents an declarative configuration of the PVCRequest type for use
// with apply.
type PVCRequestApplyConfiguration struct {
	Name             *string                         `json:"name,omitempty"`
	StorageMode      *v1alpha1.StorageModeType       `json:"storageMode,omitempty"`
	StorageClassName *string                         `json:"storageClassName,omitempty"`
	Storage          *StorageApplyConfiguration      `json:"storage,omitempty"`
	AccessModes      []v1.PersistentVolumeAccessMode `json:"accessModes,omitempty"`
}

// PVCRequestApplyConfiguration constructs an declarative configuration of the PVCRequest type for use with
// apply.
func PVCRequest() *PVCRequestApplyConfiguration {
	return &PVCRequestApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PVCRequestApplyConfiguration) WithName(value string) *PVCRequestApplyConfiguration {
	b.Name = &value
	return b
}

// WithStorageMode sets the StorageMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageMode field is set to the value of the last call.
func (b *PVCRequestApplyConfiguration) WithStorageMode(value v1alpha1.StorageModeType) *PVCRequestApplyConfiguration {
	b.StorageMode = &value
	return b
}

// WithStorageClassName sets the StorageClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageClassName field is set to the value of the last call.
func (b *PVCRequestApplyConfiguration) WithStorageClassName(value string) *PVCRequestApplyConfiguration {
	b.StorageClassName = &value
	return b
}

// WithStorage sets the Storage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Storage field is set to the value of the last call.
func (b *PVCRequestApplyConfiguration) WithStorage(value *StorageApplyConfiguration) *PVCRequestApplyConfiguration {
	b.Storage = value
	return b
}

// WithAccessModes adds the given value to the AccessModes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AccessModes field.
func (b *PVCRequestApplyConfiguration) WithAccessModes(values ...v1.PersistentVolumeAccessMode) *PVCRequestApplyConfiguration {
	for i := range values {
		b.AccessModes = append(b.AccessModes, values[i])
	}
	return b
}
