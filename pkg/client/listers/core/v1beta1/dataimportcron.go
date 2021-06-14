/*
Copyright 2018 The CDI Authors.

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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "kubevirt.io/containerized-data-importer/pkg/apis/core/v1beta1"
)

// DataImportCronLister helps list DataImportCrons.
// All objects returned here must be treated as read-only.
type DataImportCronLister interface {
	// List lists all DataImportCrons in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.DataImportCron, err error)
	// DataImportCrons returns an object that can list and get DataImportCrons.
	DataImportCrons(namespace string) DataImportCronNamespaceLister
	DataImportCronListerExpansion
}

// dataImportCronLister implements the DataImportCronLister interface.
type dataImportCronLister struct {
	indexer cache.Indexer
}

// NewDataImportCronLister returns a new DataImportCronLister.
func NewDataImportCronLister(indexer cache.Indexer) DataImportCronLister {
	return &dataImportCronLister{indexer: indexer}
}

// List lists all DataImportCrons in the indexer.
func (s *dataImportCronLister) List(selector labels.Selector) (ret []*v1beta1.DataImportCron, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.DataImportCron))
	})
	return ret, err
}

// DataImportCrons returns an object that can list and get DataImportCrons.
func (s *dataImportCronLister) DataImportCrons(namespace string) DataImportCronNamespaceLister {
	return dataImportCronNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DataImportCronNamespaceLister helps list and get DataImportCrons.
// All objects returned here must be treated as read-only.
type DataImportCronNamespaceLister interface {
	// List lists all DataImportCrons in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.DataImportCron, err error)
	// Get retrieves the DataImportCron from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.DataImportCron, error)
	DataImportCronNamespaceListerExpansion
}

// dataImportCronNamespaceLister implements the DataImportCronNamespaceLister
// interface.
type dataImportCronNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DataImportCrons in the indexer for a given namespace.
func (s dataImportCronNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.DataImportCron, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.DataImportCron))
	})
	return ret, err
}

// Get retrieves the DataImportCron from the indexer for a given namespace and name.
func (s dataImportCronNamespaceLister) Get(name string) (*v1beta1.DataImportCron, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("dataimportcron"), name)
	}
	return obj.(*v1beta1.DataImportCron), nil
}
