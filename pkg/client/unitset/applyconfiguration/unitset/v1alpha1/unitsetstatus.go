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
	unitsetv1alpha1 "xxx/pkg/apis/unitset/v1alpha1"
)

// UnitsetStatusApplyConfiguration represents an declarative configuration of the UnitsetStatus type for use
// with apply.
type UnitsetStatusApplyConfiguration struct {
	ErrMessages []ErrMsgApplyConfiguration  `json:"err_messages,omitempty"`
	Conditions  []unitsetv1alpha1.Condition `json:"conditions,omitempty"`
}

// UnitsetStatusApplyConfiguration constructs an declarative configuration of the UnitsetStatus type for use with
// apply.
func UnitsetStatus() *UnitsetStatusApplyConfiguration {
	return &UnitsetStatusApplyConfiguration{}
}

// WithErrMessages adds the given value to the ErrMessages field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ErrMessages field.
func (b *UnitsetStatusApplyConfiguration) WithErrMessages(values ...*ErrMsgApplyConfiguration) *UnitsetStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithErrMessages")
		}
		b.ErrMessages = append(b.ErrMessages, *values[i])
	}
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *UnitsetStatusApplyConfiguration) WithConditions(values ...unitsetv1alpha1.Condition) *UnitsetStatusApplyConfiguration {
	for i := range values {
		b.Conditions = append(b.Conditions, values[i])
	}
	return b
}
