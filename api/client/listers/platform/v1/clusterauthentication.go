/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "tkestack.io/tke/api/platform/v1"
)

// ClusterAuthenticationLister helps list ClusterAuthentications.
type ClusterAuthenticationLister interface {
	// List lists all ClusterAuthentications in the indexer.
	List(selector labels.Selector) (ret []*v1.ClusterAuthentication, err error)
	// ClusterAuthentications returns an object that can list and get ClusterAuthentications.
	ClusterAuthentications(namespace string) ClusterAuthenticationNamespaceLister
	ClusterAuthenticationListerExpansion
}

// clusterAuthenticationLister implements the ClusterAuthenticationLister interface.
type clusterAuthenticationLister struct {
	indexer cache.Indexer
}

// NewClusterAuthenticationLister returns a new ClusterAuthenticationLister.
func NewClusterAuthenticationLister(indexer cache.Indexer) ClusterAuthenticationLister {
	return &clusterAuthenticationLister{indexer: indexer}
}

// List lists all ClusterAuthentications in the indexer.
func (s *clusterAuthenticationLister) List(selector labels.Selector) (ret []*v1.ClusterAuthentication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterAuthentication))
	})
	return ret, err
}

// ClusterAuthentications returns an object that can list and get ClusterAuthentications.
func (s *clusterAuthenticationLister) ClusterAuthentications(namespace string) ClusterAuthenticationNamespaceLister {
	return clusterAuthenticationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterAuthenticationNamespaceLister helps list and get ClusterAuthentications.
type ClusterAuthenticationNamespaceLister interface {
	// List lists all ClusterAuthentications in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ClusterAuthentication, err error)
	// Get retrieves the ClusterAuthentication from the indexer for a given namespace and name.
	Get(name string) (*v1.ClusterAuthentication, error)
	ClusterAuthenticationNamespaceListerExpansion
}

// clusterAuthenticationNamespaceLister implements the ClusterAuthenticationNamespaceLister
// interface.
type clusterAuthenticationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterAuthentications in the indexer for a given namespace.
func (s clusterAuthenticationNamespaceLister) List(selector labels.Selector) (ret []*v1.ClusterAuthentication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterAuthentication))
	})
	return ret, err
}

// Get retrieves the ClusterAuthentication from the indexer for a given namespace and name.
func (s clusterAuthenticationNamespaceLister) Get(name string) (*v1.ClusterAuthentication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clusterauthentication"), name)
	}
	return obj.(*v1.ClusterAuthentication), nil
}
