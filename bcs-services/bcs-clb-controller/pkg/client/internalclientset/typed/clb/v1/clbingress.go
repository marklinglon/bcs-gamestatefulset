/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/Tencent/bk-bcs/bcs-services/bcs-clb-controller/pkg/apis/clb/v1"
	scheme "github.com/Tencent/bk-bcs/bcs-services/bcs-clb-controller/pkg/client/internalclientset/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClbIngressesGetter has a method to return a ClbIngressInterface.
// A group's client should implement this interface.
type ClbIngressesGetter interface {
	ClbIngresses(namespace string) ClbIngressInterface
}

// ClbIngressInterface has methods to work with ClbIngress resources.
type ClbIngressInterface interface {
	Create(*v1.ClbIngress) (*v1.ClbIngress, error)
	Update(*v1.ClbIngress) (*v1.ClbIngress, error)
	UpdateStatus(*v1.ClbIngress) (*v1.ClbIngress, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.ClbIngress, error)
	List(opts meta_v1.ListOptions) (*v1.ClbIngressList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClbIngress, err error)
	ClbIngressExpansion
}

// clbIngresses implements ClbIngressInterface
type clbIngresses struct {
	client rest.Interface
	ns     string
}

// newClbIngresses returns a ClbIngresses
func newClbIngresses(c *ClbV1Client, namespace string) *clbIngresses {
	return &clbIngresses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clbIngress, and returns the corresponding clbIngress object, and an error if there is any.
func (c *clbIngresses) Get(name string, options meta_v1.GetOptions) (result *v1.ClbIngress, err error) {
	result = &v1.ClbIngress{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clbingresses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClbIngresses that match those selectors.
func (c *clbIngresses) List(opts meta_v1.ListOptions) (result *v1.ClbIngressList, err error) {
	result = &v1.ClbIngressList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clbingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clbIngresses.
func (c *clbIngresses) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clbingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a clbIngress and creates it.  Returns the server's representation of the clbIngress, and an error, if there is any.
func (c *clbIngresses) Create(clbIngress *v1.ClbIngress) (result *v1.ClbIngress, err error) {
	result = &v1.ClbIngress{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clbingresses").
		Body(clbIngress).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clbIngress and updates it. Returns the server's representation of the clbIngress, and an error, if there is any.
func (c *clbIngresses) Update(clbIngress *v1.ClbIngress) (result *v1.ClbIngress, err error) {
	result = &v1.ClbIngress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clbingresses").
		Name(clbIngress.Name).
		Body(clbIngress).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clbIngresses) UpdateStatus(clbIngress *v1.ClbIngress) (result *v1.ClbIngress, err error) {
	result = &v1.ClbIngress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clbingresses").
		Name(clbIngress.Name).
		SubResource("status").
		Body(clbIngress).
		Do().
		Into(result)
	return
}

// Delete takes name of the clbIngress and deletes it. Returns an error if one occurs.
func (c *clbIngresses) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clbingresses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clbIngresses) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clbingresses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clbIngress.
func (c *clbIngresses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClbIngress, err error) {
	result = &v1.ClbIngress{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clbingresses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
