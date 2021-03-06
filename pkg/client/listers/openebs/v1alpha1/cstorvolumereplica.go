/*
Copyright 2017 The OpenEBS Authors

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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CStorVolumeReplicaLister helps list CStorVolumeReplicas.
type CStorVolumeReplicaLister interface {
	// List lists all CStorVolumeReplicas in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CStorVolumeReplica, err error)
	// Get retrieves the CStorVolumeReplica from the index for a given name.
	Get(name string) (*v1alpha1.CStorVolumeReplica, error)
	CStorVolumeReplicaListerExpansion
}

// cStorVolumeReplicaLister implements the CStorVolumeReplicaLister interface.
type cStorVolumeReplicaLister struct {
	indexer cache.Indexer
}

// NewCStorVolumeReplicaLister returns a new CStorVolumeReplicaLister.
func NewCStorVolumeReplicaLister(indexer cache.Indexer) CStorVolumeReplicaLister {
	return &cStorVolumeReplicaLister{indexer: indexer}
}

// List lists all CStorVolumeReplicas in the indexer.
func (s *cStorVolumeReplicaLister) List(selector labels.Selector) (ret []*v1alpha1.CStorVolumeReplica, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CStorVolumeReplica))
	})
	return ret, err
}

// Get retrieves the CStorVolumeReplica from the index for a given name.
func (s *cStorVolumeReplicaLister) Get(name string) (*v1alpha1.CStorVolumeReplica, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cstorvolumereplica"), name)
	}
	return obj.(*v1alpha1.CStorVolumeReplica), nil
}
