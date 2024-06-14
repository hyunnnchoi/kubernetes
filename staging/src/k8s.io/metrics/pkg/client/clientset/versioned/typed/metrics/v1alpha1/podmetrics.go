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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	consistencydetector "k8s.io/client-go/util/consistencydetector"
	watchlist "k8s.io/client-go/util/watchlist"
	"k8s.io/klog/v2"
	v1alpha1 "k8s.io/metrics/pkg/apis/metrics/v1alpha1"
	scheme "k8s.io/metrics/pkg/client/clientset/versioned/scheme"
)

// PodMetricsesGetter has a method to return a PodMetricsInterface.
// A group's client should implement this interface.
type PodMetricsesGetter interface {
	PodMetricses(namespace string) PodMetricsInterface
}

// PodMetricsInterface has methods to work with PodMetrics resources.
type PodMetricsInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PodMetrics, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PodMetricsList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	PodMetricsExpansion
}

// podMetricses implements PodMetricsInterface
type podMetricses struct {
	client rest.Interface
	ns     string
}

// newPodMetricses returns a PodMetricses
func newPodMetricses(c *MetricsV1alpha1Client, namespace string) *podMetricses {
	return &podMetricses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the podMetrics, and returns the corresponding podMetrics object, and an error if there is any.
func (c *podMetricses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PodMetrics, err error) {
	result = &v1alpha1.PodMetrics{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pods").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PodMetricses that match those selectors.
func (c *podMetricses) List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PodMetricsList, error) {
	if watchListOptions, hasWatchListOptionsPrepared, watchListOptionsErr := watchlist.PrepareWatchListOptionsFromListOptions(opts); watchListOptionsErr != nil {
		klog.Warningf("Failed preparing watchlist options for pods, falling back to the standard LIST semantics, err = %v", watchListOptionsErr)
	} else if hasWatchListOptionsPrepared {
		result, err := c.watchList(ctx, watchListOptions)
		if err == nil {
			consistencydetector.CheckWatchListFromCacheDataConsistencyIfRequested(ctx, "watchlist request for pods", c.list, opts, result)
			return result, nil
		}
		klog.Warningf("The watchlist request for pods ended with an error, falling back to the standard LIST semantics, err = %v", err)
	}
	result, err := c.list(ctx, opts)
	if err == nil {
		consistencydetector.CheckListFromCacheDataConsistencyIfRequested(ctx, "list request for pods", c.list, opts, result)
	}
	return result, err
}

// list takes label and field selectors, and returns the list of PodMetricses that match those selectors.
func (c *podMetricses) list(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PodMetricsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PodMetricsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pods").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// watchList establishes a watch stream with the server and returns the list of PodMetricses
func (c *podMetricses) watchList(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PodMetricsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PodMetricsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pods").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		WatchList(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested podMetricses.
func (c *podMetricses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pods").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}
