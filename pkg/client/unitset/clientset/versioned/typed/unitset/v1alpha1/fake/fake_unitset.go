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

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/upmio/upm-pkg/pkg/apis/unitset/v1alpha1"
	unitsetv1alpha1 "github.com/upmio/upm-pkg/pkg/client/unitset/applyconfiguration/unitset/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUnitsets implements UnitsetInterface
type FakeUnitsets struct {
	Fake *FakeCrdV1alpha1
	ns   string
}

var unitsetsResource = v1alpha1.SchemeGroupVersion.WithResource("unitsets")

var unitsetsKind = v1alpha1.SchemeGroupVersion.WithKind("Unitset")

// Get takes name of the unitset, and returns the corresponding unitset object, and an error if there is any.
func (c *FakeUnitsets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Unitset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(unitsetsResource, c.ns, name), &v1alpha1.Unitset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Unitset), err
}

// List takes label and field selectors, and returns the list of Unitsets that match those selectors.
func (c *FakeUnitsets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UnitsetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(unitsetsResource, unitsetsKind, c.ns, opts), &v1alpha1.UnitsetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UnitsetList{ListMeta: obj.(*v1alpha1.UnitsetList).ListMeta}
	for _, item := range obj.(*v1alpha1.UnitsetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested unitsets.
func (c *FakeUnitsets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(unitsetsResource, c.ns, opts))

}

// Create takes the representation of a unitset and creates it.  Returns the server's representation of the unitset, and an error, if there is any.
func (c *FakeUnitsets) Create(ctx context.Context, unitset *v1alpha1.Unitset, opts v1.CreateOptions) (result *v1alpha1.Unitset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(unitsetsResource, c.ns, unitset), &v1alpha1.Unitset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Unitset), err
}

// Update takes the representation of a unitset and updates it. Returns the server's representation of the unitset, and an error, if there is any.
func (c *FakeUnitsets) Update(ctx context.Context, unitset *v1alpha1.Unitset, opts v1.UpdateOptions) (result *v1alpha1.Unitset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(unitsetsResource, c.ns, unitset), &v1alpha1.Unitset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Unitset), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUnitsets) UpdateStatus(ctx context.Context, unitset *v1alpha1.Unitset, opts v1.UpdateOptions) (*v1alpha1.Unitset, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(unitsetsResource, "status", c.ns, unitset), &v1alpha1.Unitset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Unitset), err
}

// Delete takes name of the unitset and deletes it. Returns an error if one occurs.
func (c *FakeUnitsets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(unitsetsResource, c.ns, name, opts), &v1alpha1.Unitset{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUnitsets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(unitsetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.UnitsetList{})
	return err
}

// Patch applies the patch and returns the patched unitset.
func (c *FakeUnitsets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Unitset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(unitsetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Unitset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Unitset), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied unitset.
func (c *FakeUnitsets) Apply(ctx context.Context, unitset *unitsetv1alpha1.UnitsetApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Unitset, err error) {
	if unitset == nil {
		return nil, fmt.Errorf("unitset provided to Apply must not be nil")
	}
	data, err := json.Marshal(unitset)
	if err != nil {
		return nil, err
	}
	name := unitset.Name
	if name == nil {
		return nil, fmt.Errorf("unitset.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(unitsetsResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.Unitset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Unitset), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeUnitsets) ApplyStatus(ctx context.Context, unitset *unitsetv1alpha1.UnitsetApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Unitset, err error) {
	if unitset == nil {
		return nil, fmt.Errorf("unitset provided to Apply must not be nil")
	}
	data, err := json.Marshal(unitset)
	if err != nil {
		return nil, err
	}
	name := unitset.Name
	if name == nil {
		return nil, fmt.Errorf("unitset.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(unitsetsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha1.Unitset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Unitset), err
}
