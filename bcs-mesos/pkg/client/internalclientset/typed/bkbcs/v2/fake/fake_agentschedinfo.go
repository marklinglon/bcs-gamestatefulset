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

package fake

import (
	v2 "github.com/Tencent/bk-bcs/bcs-mesos/pkg/apis/bkbcs/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAgentSchedInfos implements AgentSchedInfoInterface
type FakeAgentSchedInfos struct {
	Fake *FakeBkbcsV2
	ns   string
}

var agentschedinfosResource = schema.GroupVersionResource{Group: "bkbcs.tencent.com", Version: "v2", Resource: "agentschedinfos"}

var agentschedinfosKind = schema.GroupVersionKind{Group: "bkbcs.tencent.com", Version: "v2", Kind: "AgentSchedInfo"}

// Get takes name of the agentSchedInfo, and returns the corresponding agentSchedInfo object, and an error if there is any.
func (c *FakeAgentSchedInfos) Get(name string, options v1.GetOptions) (result *v2.AgentSchedInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(agentschedinfosResource, c.ns, name), &v2.AgentSchedInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.AgentSchedInfo), err
}

// List takes label and field selectors, and returns the list of AgentSchedInfos that match those selectors.
func (c *FakeAgentSchedInfos) List(opts v1.ListOptions) (result *v2.AgentSchedInfoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(agentschedinfosResource, agentschedinfosKind, c.ns, opts), &v2.AgentSchedInfoList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.AgentSchedInfoList{ListMeta: obj.(*v2.AgentSchedInfoList).ListMeta}
	for _, item := range obj.(*v2.AgentSchedInfoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested agentSchedInfos.
func (c *FakeAgentSchedInfos) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(agentschedinfosResource, c.ns, opts))

}

// Create takes the representation of a agentSchedInfo and creates it.  Returns the server's representation of the agentSchedInfo, and an error, if there is any.
func (c *FakeAgentSchedInfos) Create(agentSchedInfo *v2.AgentSchedInfo) (result *v2.AgentSchedInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(agentschedinfosResource, c.ns, agentSchedInfo), &v2.AgentSchedInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.AgentSchedInfo), err
}

// Update takes the representation of a agentSchedInfo and updates it. Returns the server's representation of the agentSchedInfo, and an error, if there is any.
func (c *FakeAgentSchedInfos) Update(agentSchedInfo *v2.AgentSchedInfo) (result *v2.AgentSchedInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(agentschedinfosResource, c.ns, agentSchedInfo), &v2.AgentSchedInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.AgentSchedInfo), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAgentSchedInfos) UpdateStatus(agentSchedInfo *v2.AgentSchedInfo) (*v2.AgentSchedInfo, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(agentschedinfosResource, "status", c.ns, agentSchedInfo), &v2.AgentSchedInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.AgentSchedInfo), err
}

// Delete takes name of the agentSchedInfo and deletes it. Returns an error if one occurs.
func (c *FakeAgentSchedInfos) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(agentschedinfosResource, c.ns, name), &v2.AgentSchedInfo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAgentSchedInfos) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(agentschedinfosResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v2.AgentSchedInfoList{})
	return err
}

// Patch applies the patch and returns the patched agentSchedInfo.
func (c *FakeAgentSchedInfos) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.AgentSchedInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(agentschedinfosResource, c.ns, name, data, subresources...), &v2.AgentSchedInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.AgentSchedInfo), err
}