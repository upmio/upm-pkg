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
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/upmio/upm-pkg/pkg/apis/unitset/v1alpha1"
	unitsetv1alpha1 "github.com/upmio/upm-pkg/pkg/client/unitset/applyconfiguration/unitset/v1alpha1"
	scheme "github.com/upmio/upm-pkg/pkg/client/unitset/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UnitsetsGetter has a method to return a UnitsetInterface.
// A group's client should implement this interface.
type UnitsetsGetter interface {
	Unitsets(namespace string) UnitsetInterface
}

// UnitsetInterface has methods to work with Unitset resources.
type UnitsetInterface interface {
	Create(ctx context.Context, unitset *v1alpha1.Unitset, opts v1.CreateOptions) (*v1alpha1.Unitset, error)
	Update(ctx context.Context, unitset *v1alpha1.Unitset, opts v1.UpdateOptions) (*v1alpha1.Unitset, error)
	UpdateStatus(ctx context.Context, unitset *v1alpha1.Unitset, opts v1.UpdateOptions) (*v1alpha1.Unitset, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Unitset, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.UnitsetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Unitset, err error)
	Apply(ctx context.Context, unitset *unitsetv1alpha1.UnitsetApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Unitset, err error)
	ApplyStatus(ctx context.Context, unitset *unitsetv1alpha1.UnitsetApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Unitset, err error)
	UnitsetExpansion
}

// unitsets implements UnitsetInterface
type unitsets struct {
	client rest.Interface
	ns     string
}

// newUnitsets returns a Unitsets
func newUnitsets(c *CrdV1alpha1Client, namespace string) *unitsets {
	return &unitsets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the unitset, and returns the corresponding unitset object, and an error if there is any.
func (c *unitsets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Unitset, err error) {
	result = &v1alpha1.Unitset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("unitsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Unitsets that match those selectors.
func (c *unitsets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UnitsetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.UnitsetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("unitsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested unitsets.
func (c *unitsets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("unitsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a unitset and creates it.  Returns the server's representation of the unitset, and an error, if there is any.
func (c *unitsets) Create(ctx context.Context, unitset *v1alpha1.Unitset, opts v1.CreateOptions) (result *v1alpha1.Unitset, err error) {
	result = &v1alpha1.Unitset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("unitsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(unitset).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a unitset and updates it. Returns the server's representation of the unitset, and an error, if there is any.
func (c *unitsets) Update(ctx context.Context, unitset *v1alpha1.Unitset, opts v1.UpdateOptions) (result *v1alpha1.Unitset, err error) {
	result = &v1alpha1.Unitset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("unitsets").
		Name(unitset.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(unitset).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *unitsets) UpdateStatus(ctx context.Context, unitset *v1alpha1.Unitset, opts v1.UpdateOptions) (result *v1alpha1.Unitset, err error) {
	result = &v1alpha1.Unitset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("unitsets").
		Name(unitset.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(unitset).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the unitset and deletes it. Returns an error if one occurs.
func (c *unitsets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("unitsets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *unitsets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("unitsets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched unitset.
func (c *unitsets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Unitset, err error) {
	result = &v1alpha1.Unitset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("unitsets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied unitset.
func (c *unitsets) Apply(ctx context.Context, unitset *unitsetv1alpha1.UnitsetApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Unitset, err error) {
	if unitset == nil {
		return nil, fmt.Errorf("unitset provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(unitset)
	if err != nil {
		return nil, err
	}
	name := unitset.Name
	if name == nil {
		return nil, fmt.Errorf("unitset.Name must be provided to Apply")
	}
	result = &v1alpha1.Unitset{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("unitsets").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *unitsets) ApplyStatus(ctx context.Context, unitset *unitsetv1alpha1.UnitsetApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Unitset, err error) {
	if unitset == nil {
		return nil, fmt.Errorf("unitset provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(unitset)
	if err != nil {
		return nil, err
	}

	name := unitset.Name
	if name == nil {
		return nil, fmt.Errorf("unitset.Name must be provided to Apply")
	}

	result = &v1alpha1.Unitset{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("unitsets").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
