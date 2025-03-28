/*
 * Copyright The Kmesh Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
	kmeshnodeinfov1alpha1 "kmesh.net/kmesh/pkg/kube/apis/kmeshnodeinfo/v1alpha1"
)

// KmeshNodeInfoLister helps list KmeshNodeInfos.
// All objects returned here must be treated as read-only.
type KmeshNodeInfoLister interface {
	// List lists all KmeshNodeInfos in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kmeshnodeinfov1alpha1.KmeshNodeInfo, err error)
	// KmeshNodeInfos returns an object that can list and get KmeshNodeInfos.
	KmeshNodeInfos(namespace string) KmeshNodeInfoNamespaceLister
	KmeshNodeInfoListerExpansion
}

// kmeshNodeInfoLister implements the KmeshNodeInfoLister interface.
type kmeshNodeInfoLister struct {
	listers.ResourceIndexer[*kmeshnodeinfov1alpha1.KmeshNodeInfo]
}

// NewKmeshNodeInfoLister returns a new KmeshNodeInfoLister.
func NewKmeshNodeInfoLister(indexer cache.Indexer) KmeshNodeInfoLister {
	return &kmeshNodeInfoLister{listers.New[*kmeshnodeinfov1alpha1.KmeshNodeInfo](indexer, kmeshnodeinfov1alpha1.Resource("kmeshnodeinfo"))}
}

// KmeshNodeInfos returns an object that can list and get KmeshNodeInfos.
func (s *kmeshNodeInfoLister) KmeshNodeInfos(namespace string) KmeshNodeInfoNamespaceLister {
	return kmeshNodeInfoNamespaceLister{listers.NewNamespaced[*kmeshnodeinfov1alpha1.KmeshNodeInfo](s.ResourceIndexer, namespace)}
}

// KmeshNodeInfoNamespaceLister helps list and get KmeshNodeInfos.
// All objects returned here must be treated as read-only.
type KmeshNodeInfoNamespaceLister interface {
	// List lists all KmeshNodeInfos in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kmeshnodeinfov1alpha1.KmeshNodeInfo, err error)
	// Get retrieves the KmeshNodeInfo from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*kmeshnodeinfov1alpha1.KmeshNodeInfo, error)
	KmeshNodeInfoNamespaceListerExpansion
}

// kmeshNodeInfoNamespaceLister implements the KmeshNodeInfoNamespaceLister
// interface.
type kmeshNodeInfoNamespaceLister struct {
	listers.ResourceIndexer[*kmeshnodeinfov1alpha1.KmeshNodeInfo]
}
