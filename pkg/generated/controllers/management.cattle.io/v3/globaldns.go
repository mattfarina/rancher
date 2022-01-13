/*
Copyright 2022 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type GlobalDnsHandler func(string, *v3.GlobalDns) (*v3.GlobalDns, error)

type GlobalDnsController interface {
	generic.ControllerMeta
	GlobalDnsClient

	OnChange(ctx context.Context, name string, sync GlobalDnsHandler)
	OnRemove(ctx context.Context, name string, sync GlobalDnsHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() GlobalDnsCache
}

type GlobalDnsClient interface {
	Create(*v3.GlobalDns) (*v3.GlobalDns, error)
	Update(*v3.GlobalDns) (*v3.GlobalDns, error)
	UpdateStatus(*v3.GlobalDns) (*v3.GlobalDns, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v3.GlobalDns, error)
	List(namespace string, opts metav1.ListOptions) (*v3.GlobalDnsList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.GlobalDns, err error)
}

type GlobalDnsCache interface {
	Get(namespace, name string) (*v3.GlobalDns, error)
	List(namespace string, selector labels.Selector) ([]*v3.GlobalDns, error)

	AddIndexer(indexName string, indexer GlobalDnsIndexer)
	GetByIndex(indexName, key string) ([]*v3.GlobalDns, error)
}

type GlobalDnsIndexer func(obj *v3.GlobalDns) ([]string, error)

type globalDnsController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewGlobalDnsController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) GlobalDnsController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &globalDnsController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromGlobalDnsHandlerToHandler(sync GlobalDnsHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.GlobalDns
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.GlobalDns))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *globalDnsController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.GlobalDns))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateGlobalDnsDeepCopyOnChange(client GlobalDnsClient, obj *v3.GlobalDns, handler func(obj *v3.GlobalDns) (*v3.GlobalDns, error)) (*v3.GlobalDns, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *globalDnsController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *globalDnsController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *globalDnsController) OnChange(ctx context.Context, name string, sync GlobalDnsHandler) {
	c.AddGenericHandler(ctx, name, FromGlobalDnsHandlerToHandler(sync))
}

func (c *globalDnsController) OnRemove(ctx context.Context, name string, sync GlobalDnsHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromGlobalDnsHandlerToHandler(sync)))
}

func (c *globalDnsController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *globalDnsController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *globalDnsController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *globalDnsController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *globalDnsController) Cache() GlobalDnsCache {
	return &globalDnsCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *globalDnsController) Create(obj *v3.GlobalDns) (*v3.GlobalDns, error) {
	result := &v3.GlobalDns{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *globalDnsController) Update(obj *v3.GlobalDns) (*v3.GlobalDns, error) {
	result := &v3.GlobalDns{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *globalDnsController) UpdateStatus(obj *v3.GlobalDns) (*v3.GlobalDns, error) {
	result := &v3.GlobalDns{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *globalDnsController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *globalDnsController) Get(namespace, name string, options metav1.GetOptions) (*v3.GlobalDns, error) {
	result := &v3.GlobalDns{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *globalDnsController) List(namespace string, opts metav1.ListOptions) (*v3.GlobalDnsList, error) {
	result := &v3.GlobalDnsList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *globalDnsController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *globalDnsController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v3.GlobalDns, error) {
	result := &v3.GlobalDns{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type globalDnsCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *globalDnsCache) Get(namespace, name string) (*v3.GlobalDns, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.GlobalDns), nil
}

func (c *globalDnsCache) List(namespace string, selector labels.Selector) (ret []*v3.GlobalDns, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.GlobalDns))
	})

	return ret, err
}

func (c *globalDnsCache) AddIndexer(indexName string, indexer GlobalDnsIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.GlobalDns))
		},
	}))
}

func (c *globalDnsCache) GetByIndex(indexName, key string) (result []*v3.GlobalDns, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.GlobalDns, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.GlobalDns))
	}
	return result, nil
}

type GlobalDnsStatusHandler func(obj *v3.GlobalDns, status v3.GlobalDNSStatus) (v3.GlobalDNSStatus, error)

type GlobalDnsGeneratingHandler func(obj *v3.GlobalDns, status v3.GlobalDNSStatus) ([]runtime.Object, v3.GlobalDNSStatus, error)

func RegisterGlobalDnsStatusHandler(ctx context.Context, controller GlobalDnsController, condition condition.Cond, name string, handler GlobalDnsStatusHandler) {
	statusHandler := &globalDnsStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromGlobalDnsHandlerToHandler(statusHandler.sync))
}

func RegisterGlobalDnsGeneratingHandler(ctx context.Context, controller GlobalDnsController, apply apply.Apply,
	condition condition.Cond, name string, handler GlobalDnsGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &globalDnsGeneratingHandler{
		GlobalDnsGeneratingHandler: handler,
		apply:                      apply,
		name:                       name,
		gvk:                        controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterGlobalDnsStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type globalDnsStatusHandler struct {
	client    GlobalDnsClient
	condition condition.Cond
	handler   GlobalDnsStatusHandler
}

func (a *globalDnsStatusHandler) sync(key string, obj *v3.GlobalDns) (*v3.GlobalDns, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type globalDnsGeneratingHandler struct {
	GlobalDnsGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *globalDnsGeneratingHandler) Remove(key string, obj *v3.GlobalDns) (*v3.GlobalDns, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v3.GlobalDns{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *globalDnsGeneratingHandler) Handle(obj *v3.GlobalDns, status v3.GlobalDNSStatus) (v3.GlobalDNSStatus, error) {
	objs, newStatus, err := a.GlobalDnsGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
