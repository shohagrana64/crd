/*
Copyright Shohag Rana(shohagrana@appscode.com).

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

	v1alpha1 "github.com/shohagrana64/crd/pkg/apis/stable.example.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDisappointments implements DisappointmentInterface
type FakeDisappointments struct {
	Fake *FakeStableV1alpha1
	ns   string
}

var disappointmentsResource = schema.GroupVersionResource{Group: "stable.example.com", Version: "v1alpha1", Resource: "disappointments"}

var disappointmentsKind = schema.GroupVersionKind{Group: "stable.example.com", Version: "v1alpha1", Kind: "Disappointment"}

// Get takes name of the disappointment, and returns the corresponding disappointment object, and an error if there is any.
func (c *FakeDisappointments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Disappointment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(disappointmentsResource, c.ns, name), &v1alpha1.Disappointment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Disappointment), err
}

// List takes label and field selectors, and returns the list of Disappointments that match those selectors.
func (c *FakeDisappointments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DisappointmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(disappointmentsResource, disappointmentsKind, c.ns, opts), &v1alpha1.DisappointmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DisappointmentList{ListMeta: obj.(*v1alpha1.DisappointmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.DisappointmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested disappointments.
func (c *FakeDisappointments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(disappointmentsResource, c.ns, opts))

}

// Create takes the representation of a disappointment and creates it.  Returns the server's representation of the disappointment, and an error, if there is any.
func (c *FakeDisappointments) Create(ctx context.Context, disappointment *v1alpha1.Disappointment, opts v1.CreateOptions) (result *v1alpha1.Disappointment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(disappointmentsResource, c.ns, disappointment), &v1alpha1.Disappointment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Disappointment), err
}

// Update takes the representation of a disappointment and updates it. Returns the server's representation of the disappointment, and an error, if there is any.
func (c *FakeDisappointments) Update(ctx context.Context, disappointment *v1alpha1.Disappointment, opts v1.UpdateOptions) (result *v1alpha1.Disappointment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(disappointmentsResource, c.ns, disappointment), &v1alpha1.Disappointment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Disappointment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDisappointments) UpdateStatus(ctx context.Context, disappointment *v1alpha1.Disappointment, opts v1.UpdateOptions) (*v1alpha1.Disappointment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(disappointmentsResource, "status", c.ns, disappointment), &v1alpha1.Disappointment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Disappointment), err
}

// Delete takes name of the disappointment and deletes it. Returns an error if one occurs.
func (c *FakeDisappointments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(disappointmentsResource, c.ns, name), &v1alpha1.Disappointment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDisappointments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(disappointmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DisappointmentList{})
	return err
}

// Patch applies the patch and returns the patched disappointment.
func (c *FakeDisappointments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Disappointment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(disappointmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Disappointment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Disappointment), err
}
