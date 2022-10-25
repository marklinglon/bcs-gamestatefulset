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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedServiceAccountLister helps list FederatedServiceAccounts.
type FederatedServiceAccountLister interface {
	// List lists all FederatedServiceAccounts in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.FederatedServiceAccount, err error)
	// FederatedServiceAccounts returns an object that can list and get FederatedServiceAccounts.
	FederatedServiceAccounts(namespace string) FederatedServiceAccountNamespaceLister
	FederatedServiceAccountListerExpansion
}

// federatedServiceAccountLister implements the FederatedServiceAccountLister interface.
type federatedServiceAccountLister struct {
	indexer cache.Indexer
}

// NewFederatedServiceAccountLister returns a new FederatedServiceAccountLister.
func NewFederatedServiceAccountLister(indexer cache.Indexer) FederatedServiceAccountLister {
	return &federatedServiceAccountLister{indexer: indexer}
}

// List lists all FederatedServiceAccounts in the indexer.
func (s *federatedServiceAccountLister) List(selector labels.Selector) (ret []*v1beta1.FederatedServiceAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedServiceAccount))
	})
	return ret, err
}

// FederatedServiceAccounts returns an object that can list and get FederatedServiceAccounts.
func (s *federatedServiceAccountLister) FederatedServiceAccounts(namespace string) FederatedServiceAccountNamespaceLister {
	return federatedServiceAccountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedServiceAccountNamespaceLister helps list and get FederatedServiceAccounts.
type FederatedServiceAccountNamespaceLister interface {
	// List lists all FederatedServiceAccounts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.FederatedServiceAccount, err error)
	// Get retrieves the FederatedServiceAccount from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.FederatedServiceAccount, error)
	FederatedServiceAccountNamespaceListerExpansion
}

// federatedServiceAccountNamespaceLister implements the FederatedServiceAccountNamespaceLister
// interface.
type federatedServiceAccountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedServiceAccounts in the indexer for a given namespace.
func (s federatedServiceAccountNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.FederatedServiceAccount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedServiceAccount))
	})
	return ret, err
}

// Get retrieves the FederatedServiceAccount from the indexer for a given namespace and name.
func (s federatedServiceAccountNamespaceLister) Get(name string) (*v1beta1.FederatedServiceAccount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("federatedserviceaccount"), name)
	}
	return obj.(*v1beta1.FederatedServiceAccount), nil
}
