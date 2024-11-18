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

package v1

import (
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
	example2v1 "k8s.io/code-generator/examples/apiserver/apis/example2/v1"
)

// TestTypeLister helps list TestTypes.
// All objects returned here must be treated as read-only.
type TestTypeLister interface {
	// List lists all TestTypes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*example2v1.TestType, err error)
	// TestTypes returns an object that can list and get TestTypes.
	TestTypes(namespace string) TestTypeNamespaceLister
	TestTypeListerExpansion
}

// testTypeLister implements the TestTypeLister interface.
type testTypeLister struct {
	listers.ResourceIndexer[*example2v1.TestType]
}

// NewTestTypeLister returns a new TestTypeLister.
func NewTestTypeLister(indexer cache.Indexer) TestTypeLister {
	return &testTypeLister{listers.New[*example2v1.TestType](indexer, example2v1.Resource("testtype"))}
}

// TestTypes returns an object that can list and get TestTypes.
func (s *testTypeLister) TestTypes(namespace string) TestTypeNamespaceLister {
	return testTypeNamespaceLister{listers.NewNamespaced[*example2v1.TestType](s.ResourceIndexer, namespace)}
}

// TestTypeNamespaceLister helps list and get TestTypes.
// All objects returned here must be treated as read-only.
type TestTypeNamespaceLister interface {
	// List lists all TestTypes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*example2v1.TestType, err error)
	// Get retrieves the TestType from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*example2v1.TestType, error)
	TestTypeNamespaceListerExpansion
}

// testTypeNamespaceLister implements the TestTypeNamespaceLister
// interface.
type testTypeNamespaceLister struct {
	listers.ResourceIndexer[*example2v1.TestType]
}
