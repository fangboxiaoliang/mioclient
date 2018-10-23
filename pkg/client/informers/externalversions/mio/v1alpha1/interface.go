/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/hidevopsio/mioclient/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Builds returns a BuildInformer.
	Builds() BuildInformer
	// BuildConfigs returns a BuildConfigInformer.
	BuildConfigs() BuildConfigInformer
	// DeploymentConfigs returns a DeploymentConfigInformer.
	DeploymentConfigs() DeploymentConfigInformer
	// Pipelines returns a PipelineInformer.
	Pipelines() PipelineInformer
	// PipelineConfigs returns a PipelineConfigInformer.
	PipelineConfigs() PipelineConfigInformer
	// SourceConfigs returns a SourceConfigInformer.
	SourceConfigs() SourceConfigInformer
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

// Builds returns a BuildInformer.
func (v *version) Builds() BuildInformer {
	return &buildInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BuildConfigs returns a BuildConfigInformer.
func (v *version) BuildConfigs() BuildConfigInformer {
	return &buildConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DeploymentConfigs returns a DeploymentConfigInformer.
func (v *version) DeploymentConfigs() DeploymentConfigInformer {
	return &deploymentConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Pipelines returns a PipelineInformer.
func (v *version) Pipelines() PipelineInformer {
	return &pipelineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PipelineConfigs returns a PipelineConfigInformer.
func (v *version) PipelineConfigs() PipelineConfigInformer {
	return &pipelineConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SourceConfigs returns a SourceConfigInformer.
func (v *version) SourceConfigs() SourceConfigInformer {
	return &sourceConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
