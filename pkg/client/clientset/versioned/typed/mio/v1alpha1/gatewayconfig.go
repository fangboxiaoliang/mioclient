/*
Copyright 2018 The Kubernetes Authors.

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
	v1alpha1 "hidevops.io/mioclient/pkg/apis/mio/v1alpha1"
	scheme "hidevops.io/mioclient/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GatewayConfigsGetter has a method to return a GatewayConfigInterface.
// A group's client should implement this interface.
type GatewayConfigsGetter interface {
	GatewayConfigs(namespace string) GatewayConfigInterface
}

// GatewayConfigInterface has methods to work with GatewayConfig resources.
type GatewayConfigInterface interface {
	Create(*v1alpha1.GatewayConfig) (*v1alpha1.GatewayConfig, error)
	Update(*v1alpha1.GatewayConfig) (*v1alpha1.GatewayConfig, error)
	UpdateStatus(*v1alpha1.GatewayConfig) (*v1alpha1.GatewayConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GatewayConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.GatewayConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GatewayConfig, err error)
	GatewayConfigExpansion
}

// gatewayConfigs implements GatewayConfigInterface
type gatewayConfigs struct {
	client rest.Interface
	ns     string
}

// newGatewayConfigs returns a GatewayConfigs
func newGatewayConfigs(c *MioV1alpha1Client, namespace string) *gatewayConfigs {
	return &gatewayConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gatewayConfig, and returns the corresponding gatewayConfig object, and an error if there is any.
func (c *gatewayConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.GatewayConfig, err error) {
	result = &v1alpha1.GatewayConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gatewayconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GatewayConfigs that match those selectors.
func (c *gatewayConfigs) List(opts v1.ListOptions) (result *v1alpha1.GatewayConfigList, err error) {
	result = &v1alpha1.GatewayConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gatewayconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gatewayConfigs.
func (c *gatewayConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gatewayconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a gatewayConfig and creates it.  Returns the server's representation of the gatewayConfig, and an error, if there is any.
func (c *gatewayConfigs) Create(gatewayConfig *v1alpha1.GatewayConfig) (result *v1alpha1.GatewayConfig, err error) {
	result = &v1alpha1.GatewayConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gatewayconfigs").
		Body(gatewayConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gatewayConfig and updates it. Returns the server's representation of the gatewayConfig, and an error, if there is any.
func (c *gatewayConfigs) Update(gatewayConfig *v1alpha1.GatewayConfig) (result *v1alpha1.GatewayConfig, err error) {
	result = &v1alpha1.GatewayConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gatewayconfigs").
		Name(gatewayConfig.Name).
		Body(gatewayConfig).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *gatewayConfigs) UpdateStatus(gatewayConfig *v1alpha1.GatewayConfig) (result *v1alpha1.GatewayConfig, err error) {
	result = &v1alpha1.GatewayConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gatewayconfigs").
		Name(gatewayConfig.Name).
		SubResource("status").
		Body(gatewayConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the gatewayConfig and deletes it. Returns an error if one occurs.
func (c *gatewayConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gatewayconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gatewayConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gatewayconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gatewayConfig.
func (c *gatewayConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GatewayConfig, err error) {
	result = &v1alpha1.GatewayConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gatewayconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
