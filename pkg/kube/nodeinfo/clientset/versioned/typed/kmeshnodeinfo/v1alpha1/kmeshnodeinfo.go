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
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	kmeshnodeinfov1alpha1 "kmesh.net/kmesh/pkg/kube/apis/kmeshnodeinfo/v1alpha1"
	scheme "kmesh.net/kmesh/pkg/kube/nodeinfo/clientset/versioned/scheme"
)

// KmeshNodeInfosGetter has a method to return a KmeshNodeInfoInterface.
// A group's client should implement this interface.
type KmeshNodeInfosGetter interface {
	KmeshNodeInfos(namespace string) KmeshNodeInfoInterface
}

// KmeshNodeInfoInterface has methods to work with KmeshNodeInfo resources.
type KmeshNodeInfoInterface interface {
	Create(ctx context.Context, kmeshNodeInfo *kmeshnodeinfov1alpha1.KmeshNodeInfo, opts v1.CreateOptions) (*kmeshnodeinfov1alpha1.KmeshNodeInfo, error)
	Update(ctx context.Context, kmeshNodeInfo *kmeshnodeinfov1alpha1.KmeshNodeInfo, opts v1.UpdateOptions) (*kmeshnodeinfov1alpha1.KmeshNodeInfo, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, kmeshNodeInfo *kmeshnodeinfov1alpha1.KmeshNodeInfo, opts v1.UpdateOptions) (*kmeshnodeinfov1alpha1.KmeshNodeInfo, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*kmeshnodeinfov1alpha1.KmeshNodeInfo, error)
	List(ctx context.Context, opts v1.ListOptions) (*kmeshnodeinfov1alpha1.KmeshNodeInfoList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kmeshnodeinfov1alpha1.KmeshNodeInfo, err error)
	KmeshNodeInfoExpansion
}

// kmeshNodeInfos implements KmeshNodeInfoInterface
type kmeshNodeInfos struct {
	*gentype.ClientWithList[*kmeshnodeinfov1alpha1.KmeshNodeInfo, *kmeshnodeinfov1alpha1.KmeshNodeInfoList]
}

// newKmeshNodeInfos returns a KmeshNodeInfos
func newKmeshNodeInfos(c *KmeshV1alpha1Client, namespace string) *kmeshNodeInfos {
	return &kmeshNodeInfos{
		gentype.NewClientWithList[*kmeshnodeinfov1alpha1.KmeshNodeInfo, *kmeshnodeinfov1alpha1.KmeshNodeInfoList](
			"kmeshnodeinfos",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *kmeshnodeinfov1alpha1.KmeshNodeInfo { return &kmeshnodeinfov1alpha1.KmeshNodeInfo{} },
			func() *kmeshnodeinfov1alpha1.KmeshNodeInfoList { return &kmeshnodeinfov1alpha1.KmeshNodeInfoList{} },
		),
	}
}
