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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "harmonycloud.cn/middleware-operator-manager/pkg/apis/redis/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RedisClusterLister helps list RedisClusters.
type RedisClusterLister interface {
	// List lists all RedisClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RedisCluster, err error)
	// RedisClusters returns an object that can list and get RedisClusters.
	RedisClusters(namespace string) RedisClusterNamespaceLister
	RedisClusterListerExpansion
}

// redisClusterLister implements the RedisClusterLister interface.
type redisClusterLister struct {
	indexer cache.Indexer
}

// NewRedisClusterLister returns a new RedisClusterLister.
func NewRedisClusterLister(indexer cache.Indexer) RedisClusterLister {
	return &redisClusterLister{indexer: indexer}
}

// List lists all RedisClusters in the indexer.
func (s *redisClusterLister) List(selector labels.Selector) (ret []*v1alpha1.RedisCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RedisCluster))
	})
	return ret, err
}

// RedisClusters returns an object that can list and get RedisClusters.
func (s *redisClusterLister) RedisClusters(namespace string) RedisClusterNamespaceLister {
	return redisClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RedisClusterNamespaceLister helps list and get RedisClusters.
type RedisClusterNamespaceLister interface {
	// List lists all RedisClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RedisCluster, err error)
	// Get retrieves the RedisCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RedisCluster, error)
	RedisClusterNamespaceListerExpansion
}

// redisClusterNamespaceLister implements the RedisClusterNamespaceLister
// interface.
type redisClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RedisClusters in the indexer for a given namespace.
func (s redisClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RedisCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RedisCluster))
	})
	return ret, err
}

// Get retrieves the RedisCluster from the indexer for a given namespace and name.
func (s redisClusterNamespaceLister) Get(name string) (*v1alpha1.RedisCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("rediscluster"), name)
	}
	return obj.(*v1alpha1.RedisCluster), nil
}
