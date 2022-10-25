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
	clientset "github.com/Tencent/bk-bcs/bcs-services/bcs-clb-controller/pkg/client/internalclientset"
	clbv1 "github.com/Tencent/bk-bcs/bcs-services/bcs-clb-controller/pkg/client/internalclientset/typed/clb/v1"
	fakeclbv1 "github.com/Tencent/bk-bcs/bcs-services/bcs-clb-controller/pkg/client/internalclientset/typed/clb/v1/fake"
	meshv1 "github.com/Tencent/bk-bcs/bcs-services/bcs-clb-controller/pkg/client/internalclientset/typed/mesh/v1"
	fakemeshv1 "github.com/Tencent/bk-bcs/bcs-services/bcs-clb-controller/pkg/client/internalclientset/typed/mesh/v1/fake"
	networkv1 "github.com/Tencent/bk-bcs/bcs-services/bcs-clb-controller/pkg/client/internalclientset/typed/network/v1"
	fakenetworkv1 "github.com/Tencent/bk-bcs/bcs-services/bcs-clb-controller/pkg/client/internalclientset/typed/network/v1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// ClbV1 retrieves the ClbV1Client
func (c *Clientset) ClbV1() clbv1.ClbV1Interface {
	return &fakeclbv1.FakeClbV1{Fake: &c.Fake}
}

// Clb retrieves the ClbV1Client
func (c *Clientset) Clb() clbv1.ClbV1Interface {
	return &fakeclbv1.FakeClbV1{Fake: &c.Fake}
}

// MeshV1 retrieves the MeshV1Client
func (c *Clientset) MeshV1() meshv1.MeshV1Interface {
	return &fakemeshv1.FakeMeshV1{Fake: &c.Fake}
}

// Mesh retrieves the MeshV1Client
func (c *Clientset) Mesh() meshv1.MeshV1Interface {
	return &fakemeshv1.FakeMeshV1{Fake: &c.Fake}
}

// NetworkV1 retrieves the NetworkV1Client
func (c *Clientset) NetworkV1() networkv1.NetworkV1Interface {
	return &fakenetworkv1.FakeNetworkV1{Fake: &c.Fake}
}

// Network retrieves the NetworkV1Client
func (c *Clientset) Network() networkv1.NetworkV1Interface {
	return &fakenetworkv1.FakeNetworkV1{Fake: &c.Fake}
}
