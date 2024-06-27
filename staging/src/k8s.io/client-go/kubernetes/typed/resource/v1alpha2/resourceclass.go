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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"

	v1alpha2 "k8s.io/api/resource/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	resourcev1alpha2 "k8s.io/client-go/applyconfigurations/resource/v1alpha2"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// ResourceClassesGetter has a method to return a ResourceClassInterface.
// A group's client should implement this interface.
type ResourceClassesGetter interface {
	ResourceClasses() ResourceClassInterface
}

// ResourceClassInterface has methods to work with ResourceClass resources.
type ResourceClassInterface interface {
	Create(ctx context.Context, resourceClass *v1alpha2.ResourceClass, opts v1.CreateOptions) (*v1alpha2.ResourceClass, error)
	Update(ctx context.Context, resourceClass *v1alpha2.ResourceClass, opts v1.UpdateOptions) (*v1alpha2.ResourceClass, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.ResourceClass, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.ResourceClassList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ResourceClass, err error)
	Apply(ctx context.Context, resourceClass *resourcev1alpha2.ResourceClassApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha2.ResourceClass, err error)
	ResourceClassExpansion
}

// resourceClasses implements ResourceClassInterface
type resourceClasses struct {
	*gentype.ClientWithListAndApply[*v1alpha2.ResourceClass, *v1alpha2.ResourceClassList, *resourcev1alpha2.ResourceClassApplyConfiguration]
}

// newResourceClasses returns a ResourceClasses
func newResourceClasses(c *ResourceV1alpha2Client) *resourceClasses {
	return &resourceClasses{
		gentype.NewClientWithListAndApply[*v1alpha2.ResourceClass, *v1alpha2.ResourceClassList, *resourcev1alpha2.ResourceClassApplyConfiguration](
			"resourceclasses",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha2.ResourceClass { return &v1alpha2.ResourceClass{} },
			func() *v1alpha2.ResourceClassList { return &v1alpha2.ResourceClassList{} }),
	}
}
