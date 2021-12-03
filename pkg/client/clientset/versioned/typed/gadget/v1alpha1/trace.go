// Copyright 2019-2021 The Inspektor Gadget authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kinvolk/inspektor-gadget/pkg/apis/gadget/v1alpha1"
	scheme "github.com/kinvolk/inspektor-gadget/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TracesGetter has a method to return a TraceInterface.
// A group's client should implement this interface.
type TracesGetter interface {
	Traces(namespace string) TraceInterface
}

// TraceInterface has methods to work with Trace resources.
type TraceInterface interface {
	Create(ctx context.Context, trace *v1alpha1.Trace, opts v1.CreateOptions) (*v1alpha1.Trace, error)
	Update(ctx context.Context, trace *v1alpha1.Trace, opts v1.UpdateOptions) (*v1alpha1.Trace, error)
	UpdateStatus(ctx context.Context, trace *v1alpha1.Trace, opts v1.UpdateOptions) (*v1alpha1.Trace, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Trace, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.TraceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Trace, err error)
	TraceExpansion
}

// traces implements TraceInterface
type traces struct {
	client rest.Interface
	ns     string
}

// newTraces returns a Traces
func newTraces(c *GadgetV1alpha1Client, namespace string) *traces {
	return &traces{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the trace, and returns the corresponding trace object, and an error if there is any.
func (c *traces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Trace, err error) {
	result = &v1alpha1.Trace{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("traces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Traces that match those selectors.
func (c *traces) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TraceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TraceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("traces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested traces.
func (c *traces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("traces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a trace and creates it.  Returns the server's representation of the trace, and an error, if there is any.
func (c *traces) Create(ctx context.Context, trace *v1alpha1.Trace, opts v1.CreateOptions) (result *v1alpha1.Trace, err error) {
	result = &v1alpha1.Trace{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("traces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(trace).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a trace and updates it. Returns the server's representation of the trace, and an error, if there is any.
func (c *traces) Update(ctx context.Context, trace *v1alpha1.Trace, opts v1.UpdateOptions) (result *v1alpha1.Trace, err error) {
	result = &v1alpha1.Trace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("traces").
		Name(trace.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(trace).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *traces) UpdateStatus(ctx context.Context, trace *v1alpha1.Trace, opts v1.UpdateOptions) (result *v1alpha1.Trace, err error) {
	result = &v1alpha1.Trace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("traces").
		Name(trace.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(trace).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the trace and deletes it. Returns an error if one occurs.
func (c *traces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("traces").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *traces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("traces").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched trace.
func (c *traces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Trace, err error) {
	result = &v1alpha1.Trace{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("traces").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
