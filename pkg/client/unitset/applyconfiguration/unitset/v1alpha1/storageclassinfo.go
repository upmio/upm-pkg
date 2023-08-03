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

// StorageclassInfoApplyConfiguration represents an declarative configuration of the StorageclassInfo type for use
// with apply.
type StorageclassInfoApplyConfiguration struct {
	ID          *string            `json:"id,omitempty"`
	DisplayName *string            `json:"display_name,omitempty"`
	Parameters  *map[string]string `json:"parameters,omitempty"`
}

// StorageclassInfoApplyConfiguration constructs an declarative configuration of the StorageclassInfo type for use with
// apply.
func StorageclassInfo() *StorageclassInfoApplyConfiguration {
	return &StorageclassInfoApplyConfiguration{}
}

// WithID sets the ID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ID field is set to the value of the last call.
func (b *StorageclassInfoApplyConfiguration) WithID(value string) *StorageclassInfoApplyConfiguration {
	b.ID = &value
	return b
}

// WithDisplayName sets the DisplayName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisplayName field is set to the value of the last call.
func (b *StorageclassInfoApplyConfiguration) WithDisplayName(value string) *StorageclassInfoApplyConfiguration {
	b.DisplayName = &value
	return b
}

// WithParameters sets the Parameters field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Parameters field is set to the value of the last call.
func (b *StorageclassInfoApplyConfiguration) WithParameters(value map[string]string) *StorageclassInfoApplyConfiguration {
	b.Parameters = &value
	return b
}