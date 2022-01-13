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

type TemplateHandler func(string, *v3.Template) (*v3.Template, error)

type TemplateController interface {
	generic.ControllerMeta
	TemplateClient

	OnChange(ctx context.Context, name string, sync TemplateHandler)
	OnRemove(ctx context.Context, name string, sync TemplateHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() TemplateCache
}

type TemplateClient interface {
	Create(*v3.Template) (*v3.Template, error)
	Update(*v3.Template) (*v3.Template, error)
	UpdateStatus(*v3.Template) (*v3.Template, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v3.Template, error)
	List(opts metav1.ListOptions) (*v3.TemplateList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.Template, err error)
}

type TemplateCache interface {
	Get(name string) (*v3.Template, error)
	List(selector labels.Selector) ([]*v3.Template, error)

	AddIndexer(indexName string, indexer TemplateIndexer)
	GetByIndex(indexName, key string) ([]*v3.Template, error)
}

type TemplateIndexer func(obj *v3.Template) ([]string, error)

type templateController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewTemplateController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) TemplateController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &templateController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromTemplateHandlerToHandler(sync TemplateHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.Template
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.Template))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *templateController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.Template))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateTemplateDeepCopyOnChange(client TemplateClient, obj *v3.Template, handler func(obj *v3.Template) (*v3.Template, error)) (*v3.Template, error) {
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

func (c *templateController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *templateController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *templateController) OnChange(ctx context.Context, name string, sync TemplateHandler) {
	c.AddGenericHandler(ctx, name, FromTemplateHandlerToHandler(sync))
}

func (c *templateController) OnRemove(ctx context.Context, name string, sync TemplateHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromTemplateHandlerToHandler(sync)))
}

func (c *templateController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *templateController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *templateController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *templateController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *templateController) Cache() TemplateCache {
	return &templateCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *templateController) Create(obj *v3.Template) (*v3.Template, error) {
	result := &v3.Template{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *templateController) Update(obj *v3.Template) (*v3.Template, error) {
	result := &v3.Template{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *templateController) UpdateStatus(obj *v3.Template) (*v3.Template, error) {
	result := &v3.Template{}
	return result, c.client.UpdateStatus(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *templateController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *templateController) Get(name string, options metav1.GetOptions) (*v3.Template, error) {
	result := &v3.Template{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *templateController) List(opts metav1.ListOptions) (*v3.TemplateList, error) {
	result := &v3.TemplateList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *templateController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *templateController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v3.Template, error) {
	result := &v3.Template{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type templateCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *templateCache) Get(name string) (*v3.Template, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.Template), nil
}

func (c *templateCache) List(selector labels.Selector) (ret []*v3.Template, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.Template))
	})

	return ret, err
}

func (c *templateCache) AddIndexer(indexName string, indexer TemplateIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.Template))
		},
	}))
}

func (c *templateCache) GetByIndex(indexName, key string) (result []*v3.Template, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.Template, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.Template))
	}
	return result, nil
}

type TemplateStatusHandler func(obj *v3.Template, status v3.TemplateStatus) (v3.TemplateStatus, error)

type TemplateGeneratingHandler func(obj *v3.Template, status v3.TemplateStatus) ([]runtime.Object, v3.TemplateStatus, error)

func RegisterTemplateStatusHandler(ctx context.Context, controller TemplateController, condition condition.Cond, name string, handler TemplateStatusHandler) {
	statusHandler := &templateStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromTemplateHandlerToHandler(statusHandler.sync))
}

func RegisterTemplateGeneratingHandler(ctx context.Context, controller TemplateController, apply apply.Apply,
	condition condition.Cond, name string, handler TemplateGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &templateGeneratingHandler{
		TemplateGeneratingHandler: handler,
		apply:                     apply,
		name:                      name,
		gvk:                       controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterTemplateStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type templateStatusHandler struct {
	client    TemplateClient
	condition condition.Cond
	handler   TemplateStatusHandler
}

func (a *templateStatusHandler) sync(key string, obj *v3.Template) (*v3.Template, error) {
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

type templateGeneratingHandler struct {
	TemplateGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *templateGeneratingHandler) Remove(key string, obj *v3.Template) (*v3.Template, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v3.Template{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *templateGeneratingHandler) Handle(obj *v3.Template, status v3.TemplateStatus) (v3.TemplateStatus, error) {
	objs, newStatus, err := a.TemplateGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
