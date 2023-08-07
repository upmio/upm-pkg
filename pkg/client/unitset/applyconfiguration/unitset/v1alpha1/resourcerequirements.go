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
	resource "k8s.io/apimachinery/pkg/api/resource"
)

// ResourceRequirementsApplyConfiguration represents an declarative configuration of the ResourceRequirements type for use
// with apply.
type ResourceRequirementsApplyConfiguration struct {
	MiliCPU *resource.Quantity             `json:"milicpu,omitempty"`
	Memory  *resource.Quantity             `json:"memory,omitempty"`
	Cache   *CacheInfoApplyConfiguration   `json:"cache,omitempty"`
	Storage *StorageInfoApplyConfiguration `json:"storage,omitempty"`
}

// ResourceRequirementsApplyConfiguration constructs an declarative configuration of the ResourceRequirements type for use with
// apply.
func ResourceRequirements() *ResourceRequirementsApplyConfiguration {
	return &ResourceRequirementsApplyConfiguration{}
}

// WithMiliCPU sets the MiliCPU field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MiliCPU field is set to the value of the last call.
func (b *ResourceRequirementsApplyConfiguration) WithMiliCPU(value resource.Quantity) *ResourceRequirementsApplyConfiguration {
	b.MiliCPU = &value
	return b
}

// WithMemory sets the Memory field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Memory field is set to the value of the last call.
func (b *ResourceRequirementsApplyConfiguration) WithMemory(value resource.Quantity) *ResourceRequirementsApplyConfiguration {
	b.Memory = &value
	return b
}

// WithCache sets the Cache field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cache field is set to the value of the last call.
func (b *ResourceRequirementsApplyConfiguration) WithCache(value *CacheInfoApplyConfiguration) *ResourceRequirementsApplyConfiguration {
	b.Cache = value
	return b
}

// WithStorage sets the Storage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Storage field is set to the value of the last call.
func (b *ResourceRequirementsApplyConfiguration) WithStorage(value *StorageInfoApplyConfiguration) *ResourceRequirementsApplyConfiguration {
	b.Storage = value
	return b
}
