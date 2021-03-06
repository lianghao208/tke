/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
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

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
	notify "tkestack.io/tke/api/notify"
)

// MessageRequestsGetter has a method to return a MessageRequestInterface.
// A group's client should implement this interface.
type MessageRequestsGetter interface {
	MessageRequests(namespace string) MessageRequestInterface
}

// MessageRequestInterface has methods to work with MessageRequest resources.
type MessageRequestInterface interface {
	Create(*notify.MessageRequest) (*notify.MessageRequest, error)
	Update(*notify.MessageRequest) (*notify.MessageRequest, error)
	UpdateStatus(*notify.MessageRequest) (*notify.MessageRequest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*notify.MessageRequest, error)
	List(opts v1.ListOptions) (*notify.MessageRequestList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *notify.MessageRequest, err error)
	MessageRequestExpansion
}

// messageRequests implements MessageRequestInterface
type messageRequests struct {
	client rest.Interface
	ns     string
}

// newMessageRequests returns a MessageRequests
func newMessageRequests(c *NotifyClient, namespace string) *messageRequests {
	return &messageRequests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the messageRequest, and returns the corresponding messageRequest object, and an error if there is any.
func (c *messageRequests) Get(name string, options v1.GetOptions) (result *notify.MessageRequest, err error) {
	result = &notify.MessageRequest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("messagerequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MessageRequests that match those selectors.
func (c *messageRequests) List(opts v1.ListOptions) (result *notify.MessageRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &notify.MessageRequestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("messagerequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested messageRequests.
func (c *messageRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("messagerequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a messageRequest and creates it.  Returns the server's representation of the messageRequest, and an error, if there is any.
func (c *messageRequests) Create(messageRequest *notify.MessageRequest) (result *notify.MessageRequest, err error) {
	result = &notify.MessageRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("messagerequests").
		Body(messageRequest).
		Do().
		Into(result)
	return
}

// Update takes the representation of a messageRequest and updates it. Returns the server's representation of the messageRequest, and an error, if there is any.
func (c *messageRequests) Update(messageRequest *notify.MessageRequest) (result *notify.MessageRequest, err error) {
	result = &notify.MessageRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("messagerequests").
		Name(messageRequest.Name).
		Body(messageRequest).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *messageRequests) UpdateStatus(messageRequest *notify.MessageRequest) (result *notify.MessageRequest, err error) {
	result = &notify.MessageRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("messagerequests").
		Name(messageRequest.Name).
		SubResource("status").
		Body(messageRequest).
		Do().
		Into(result)
	return
}

// Delete takes name of the messageRequest and deletes it. Returns an error if one occurs.
func (c *messageRequests) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("messagerequests").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *messageRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("messagerequests").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched messageRequest.
func (c *messageRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *notify.MessageRequest, err error) {
	result = &notify.MessageRequest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("messagerequests").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
