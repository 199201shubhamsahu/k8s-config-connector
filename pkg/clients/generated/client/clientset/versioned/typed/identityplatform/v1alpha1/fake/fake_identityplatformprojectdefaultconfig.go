// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/identityplatform/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIdentityPlatformProjectDefaultConfigs implements IdentityPlatformProjectDefaultConfigInterface
type FakeIdentityPlatformProjectDefaultConfigs struct {
	Fake *FakeIdentityplatformV1alpha1
	ns   string
}

var identityplatformprojectdefaultconfigsResource = schema.GroupVersionResource{Group: "identityplatform.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "identityplatformprojectdefaultconfigs"}

var identityplatformprojectdefaultconfigsKind = schema.GroupVersionKind{Group: "identityplatform.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "IdentityPlatformProjectDefaultConfig"}

// Get takes name of the identityPlatformProjectDefaultConfig, and returns the corresponding identityPlatformProjectDefaultConfig object, and an error if there is any.
func (c *FakeIdentityPlatformProjectDefaultConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IdentityPlatformProjectDefaultConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(identityplatformprojectdefaultconfigsResource, c.ns, name), &v1alpha1.IdentityPlatformProjectDefaultConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityPlatformProjectDefaultConfig), err
}

// List takes label and field selectors, and returns the list of IdentityPlatformProjectDefaultConfigs that match those selectors.
func (c *FakeIdentityPlatformProjectDefaultConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IdentityPlatformProjectDefaultConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(identityplatformprojectdefaultconfigsResource, identityplatformprojectdefaultconfigsKind, c.ns, opts), &v1alpha1.IdentityPlatformProjectDefaultConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IdentityPlatformProjectDefaultConfigList{ListMeta: obj.(*v1alpha1.IdentityPlatformProjectDefaultConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.IdentityPlatformProjectDefaultConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested identityPlatformProjectDefaultConfigs.
func (c *FakeIdentityPlatformProjectDefaultConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(identityplatformprojectdefaultconfigsResource, c.ns, opts))

}

// Create takes the representation of a identityPlatformProjectDefaultConfig and creates it.  Returns the server's representation of the identityPlatformProjectDefaultConfig, and an error, if there is any.
func (c *FakeIdentityPlatformProjectDefaultConfigs) Create(ctx context.Context, identityPlatformProjectDefaultConfig *v1alpha1.IdentityPlatformProjectDefaultConfig, opts v1.CreateOptions) (result *v1alpha1.IdentityPlatformProjectDefaultConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(identityplatformprojectdefaultconfigsResource, c.ns, identityPlatformProjectDefaultConfig), &v1alpha1.IdentityPlatformProjectDefaultConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityPlatformProjectDefaultConfig), err
}

// Update takes the representation of a identityPlatformProjectDefaultConfig and updates it. Returns the server's representation of the identityPlatformProjectDefaultConfig, and an error, if there is any.
func (c *FakeIdentityPlatformProjectDefaultConfigs) Update(ctx context.Context, identityPlatformProjectDefaultConfig *v1alpha1.IdentityPlatformProjectDefaultConfig, opts v1.UpdateOptions) (result *v1alpha1.IdentityPlatformProjectDefaultConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(identityplatformprojectdefaultconfigsResource, c.ns, identityPlatformProjectDefaultConfig), &v1alpha1.IdentityPlatformProjectDefaultConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityPlatformProjectDefaultConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIdentityPlatformProjectDefaultConfigs) UpdateStatus(ctx context.Context, identityPlatformProjectDefaultConfig *v1alpha1.IdentityPlatformProjectDefaultConfig, opts v1.UpdateOptions) (*v1alpha1.IdentityPlatformProjectDefaultConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(identityplatformprojectdefaultconfigsResource, "status", c.ns, identityPlatformProjectDefaultConfig), &v1alpha1.IdentityPlatformProjectDefaultConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityPlatformProjectDefaultConfig), err
}

// Delete takes name of the identityPlatformProjectDefaultConfig and deletes it. Returns an error if one occurs.
func (c *FakeIdentityPlatformProjectDefaultConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(identityplatformprojectdefaultconfigsResource, c.ns, name, opts), &v1alpha1.IdentityPlatformProjectDefaultConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIdentityPlatformProjectDefaultConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(identityplatformprojectdefaultconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IdentityPlatformProjectDefaultConfigList{})
	return err
}

// Patch applies the patch and returns the patched identityPlatformProjectDefaultConfig.
func (c *FakeIdentityPlatformProjectDefaultConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IdentityPlatformProjectDefaultConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(identityplatformprojectdefaultconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IdentityPlatformProjectDefaultConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityPlatformProjectDefaultConfig), err
}