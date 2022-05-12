// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/cilium/tetragon/pkg/k8s/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// TracingPolicies returns a TracingPolicyInformer.
	TracingPolicies() TracingPolicyInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// TracingPolicies returns a TracingPolicyInformer.
func (v *version) TracingPolicies() TracingPolicyInformer {
	return &tracingPolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
