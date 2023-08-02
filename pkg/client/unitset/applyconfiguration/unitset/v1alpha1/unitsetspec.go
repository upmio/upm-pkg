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

// UnitsetSpecApplyConfiguration represents an declarative configuration of the UnitsetSpec type for use
// with apply.
type UnitsetSpecApplyConfiguration struct {
	SharedConfigName      *string                               `json:"shared_config_name,omitempty"`
	UnbindNode            *bool                                 `json:"unbind_node,omitempty"`
	InitOnly              *bool                                 `json:"init_only,omitempty"`
	Affinity              *AffinityApplyConfiguration           `json:"affinity,omitempty"`
	PodAntiAffinity       *string                               `json:"pod_anti_affinity,omitempty"`
	ZoneAffinity          *AffinityApplyConfiguration           `json:"zone_affinity,omitempty"`
	SourceAffinity        *AffinityNewApplyConfiguration        `json:"source_affinity,omitempty"`
	Arch                  *ArchApplyConfiguration               `json:"arch,omitempty"`
	Image                 *ImageVersionApplyConfiguration       `json:"image,omitempty"`
	ImageRepositoryAddr   *string                               `json:"image_repository_addr,omitempty"`
	ConfigSets            []ConfigSetApplyConfiguration         `json:"config_sets,omitempty"`
	EndpointMode          *string                               `json:"endpoint_mode,omitempty"`
	Ports                 []ContainerPortApplyConfiguration     `json:"ports,omitempty"`
	Options               map[string]string                     `json:"options,omitempty"`
	Unit                  *UnitSpecApplyConfiguration           `json:"unit,omitempty"`
	AuthSecret            *string                               `json:"auth_secret,omitempty"`
	CASecret              *string                               `json:"ca_secret,omitempty"`
	ExternalSecret        *ExternalSecretInfoApplyConfiguration `json:"external_secret,omitempty"`
	Service               *K8sServiceApplyConfiguration         `json:"service,omitempty"`
	ShareProcessNamespace *bool                                 `json:"share_process_namespace,omitempty"`
}

// UnitsetSpecApplyConfiguration constructs an declarative configuration of the UnitsetSpec type for use with
// apply.
func UnitsetSpec() *UnitsetSpecApplyConfiguration {
	return &UnitsetSpecApplyConfiguration{}
}

// WithSharedConfigName sets the SharedConfigName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SharedConfigName field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithSharedConfigName(value string) *UnitsetSpecApplyConfiguration {
	b.SharedConfigName = &value
	return b
}

// WithUnbindNode sets the UnbindNode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UnbindNode field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithUnbindNode(value bool) *UnitsetSpecApplyConfiguration {
	b.UnbindNode = &value
	return b
}

// WithInitOnly sets the InitOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InitOnly field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithInitOnly(value bool) *UnitsetSpecApplyConfiguration {
	b.InitOnly = &value
	return b
}

// WithAffinity sets the Affinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affinity field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithAffinity(value *AffinityApplyConfiguration) *UnitsetSpecApplyConfiguration {
	b.Affinity = value
	return b
}

// WithPodAntiAffinity sets the PodAntiAffinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodAntiAffinity field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithPodAntiAffinity(value string) *UnitsetSpecApplyConfiguration {
	b.PodAntiAffinity = &value
	return b
}

// WithZoneAffinity sets the ZoneAffinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ZoneAffinity field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithZoneAffinity(value *AffinityApplyConfiguration) *UnitsetSpecApplyConfiguration {
	b.ZoneAffinity = value
	return b
}

// WithSourceAffinity sets the SourceAffinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SourceAffinity field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithSourceAffinity(value *AffinityNewApplyConfiguration) *UnitsetSpecApplyConfiguration {
	b.SourceAffinity = value
	return b
}

// WithArch sets the Arch field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Arch field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithArch(value *ArchApplyConfiguration) *UnitsetSpecApplyConfiguration {
	b.Arch = value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithImage(value *ImageVersionApplyConfiguration) *UnitsetSpecApplyConfiguration {
	b.Image = value
	return b
}

// WithImageRepositoryAddr sets the ImageRepositoryAddr field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImageRepositoryAddr field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithImageRepositoryAddr(value string) *UnitsetSpecApplyConfiguration {
	b.ImageRepositoryAddr = &value
	return b
}

// WithConfigSets adds the given value to the ConfigSets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ConfigSets field.
func (b *UnitsetSpecApplyConfiguration) WithConfigSets(values ...*ConfigSetApplyConfiguration) *UnitsetSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConfigSets")
		}
		b.ConfigSets = append(b.ConfigSets, *values[i])
	}
	return b
}

// WithEndpointMode sets the EndpointMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndpointMode field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithEndpointMode(value string) *UnitsetSpecApplyConfiguration {
	b.EndpointMode = &value
	return b
}

// WithPorts adds the given value to the Ports field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ports field.
func (b *UnitsetSpecApplyConfiguration) WithPorts(values ...*ContainerPortApplyConfiguration) *UnitsetSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPorts")
		}
		b.Ports = append(b.Ports, *values[i])
	}
	return b
}

// WithOptions puts the entries into the Options field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Options field,
// overwriting an existing map entries in Options field with the same key.
func (b *UnitsetSpecApplyConfiguration) WithOptions(entries map[string]string) *UnitsetSpecApplyConfiguration {
	if b.Options == nil && len(entries) > 0 {
		b.Options = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Options[k] = v
	}
	return b
}

// WithUnit sets the Unit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Unit field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithUnit(value *UnitSpecApplyConfiguration) *UnitsetSpecApplyConfiguration {
	b.Unit = value
	return b
}

// WithAuthSecret sets the AuthSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AuthSecret field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithAuthSecret(value string) *UnitsetSpecApplyConfiguration {
	b.AuthSecret = &value
	return b
}

// WithCASecret sets the CASecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CASecret field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithCASecret(value string) *UnitsetSpecApplyConfiguration {
	b.CASecret = &value
	return b
}

// WithExternalSecret sets the ExternalSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExternalSecret field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithExternalSecret(value *ExternalSecretInfoApplyConfiguration) *UnitsetSpecApplyConfiguration {
	b.ExternalSecret = value
	return b
}

// WithService sets the Service field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Service field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithService(value *K8sServiceApplyConfiguration) *UnitsetSpecApplyConfiguration {
	b.Service = value
	return b
}

// WithShareProcessNamespace sets the ShareProcessNamespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ShareProcessNamespace field is set to the value of the last call.
func (b *UnitsetSpecApplyConfiguration) WithShareProcessNamespace(value bool) *UnitsetSpecApplyConfiguration {
	b.ShareProcessNamespace = &value
	return b
}
