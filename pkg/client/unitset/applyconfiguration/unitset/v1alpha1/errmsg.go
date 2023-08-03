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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ErrMsgApplyConfiguration represents an declarative configuration of the ErrMsg type for use
// with apply.
type ErrMsgApplyConfiguration struct {
	Time *v1.Time `json:"time,omitempty"`
	Err  *string  `json:"err,omitempty"`
	Mode *string  `json:"mode,omitempty"`
}

// ErrMsgApplyConfiguration constructs an declarative configuration of the ErrMsg type for use with
// apply.
func ErrMsg() *ErrMsgApplyConfiguration {
	return &ErrMsgApplyConfiguration{}
}

// WithTime sets the Time field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Time field is set to the value of the last call.
func (b *ErrMsgApplyConfiguration) WithTime(value v1.Time) *ErrMsgApplyConfiguration {
	b.Time = &value
	return b
}

// WithErr sets the Err field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Err field is set to the value of the last call.
func (b *ErrMsgApplyConfiguration) WithErr(value string) *ErrMsgApplyConfiguration {
	b.Err = &value
	return b
}

// WithMode sets the Mode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Mode field is set to the value of the last call.
func (b *ErrMsgApplyConfiguration) WithMode(value string) *ErrMsgApplyConfiguration {
	b.Mode = &value
	return b
}