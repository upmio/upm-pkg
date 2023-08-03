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

// AffinityApplyConfiguration represents an declarative configuration of the Affinity type for use
// with apply.
type AffinityApplyConfiguration struct {
	Required  []string `json:"required,omitempty"`
	Preferred []string `json:"preferred,omitempty"`
}

// AffinityApplyConfiguration constructs an declarative configuration of the Affinity type for use with
// apply.
func Affinity() *AffinityApplyConfiguration {
	return &AffinityApplyConfiguration{}
}

// WithRequired adds the given value to the Required field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Required field.
func (b *AffinityApplyConfiguration) WithRequired(values ...string) *AffinityApplyConfiguration {
	for i := range values {
		b.Required = append(b.Required, values[i])
	}
	return b
}

// WithPreferred adds the given value to the Preferred field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Preferred field.
func (b *AffinityApplyConfiguration) WithPreferred(values ...string) *AffinityApplyConfiguration {
	for i := range values {
		b.Preferred = append(b.Preferred, values[i])
	}
	return b
}