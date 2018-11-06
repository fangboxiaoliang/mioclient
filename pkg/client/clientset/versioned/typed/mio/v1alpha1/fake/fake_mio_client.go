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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "hidevops.io/mioclient/pkg/client/clientset/versioned/typed/mio/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMioV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMioV1alpha1) Builds(namespace string) v1alpha1.BuildInterface {
	return &FakeBuilds{c, namespace}
}

func (c *FakeMioV1alpha1) BuildConfigs(namespace string) v1alpha1.BuildConfigInterface {
	return &FakeBuildConfigs{c, namespace}
}

func (c *FakeMioV1alpha1) Deployments(namespace string) v1alpha1.DeploymentInterface {
	return &FakeDeployments{c, namespace}
}

func (c *FakeMioV1alpha1) DeploymentConfigs(namespace string) v1alpha1.DeploymentConfigInterface {
	return &FakeDeploymentConfigs{c, namespace}
}

func (c *FakeMioV1alpha1) GatewayConfigs(namespace string) v1alpha1.GatewayConfigInterface {
	return &FakeGatewayConfigs{c, namespace}
}

func (c *FakeMioV1alpha1) Pipelines(namespace string) v1alpha1.PipelineInterface {
	return &FakePipelines{c, namespace}
}

func (c *FakeMioV1alpha1) PipelineConfigs(namespace string) v1alpha1.PipelineConfigInterface {
	return &FakePipelineConfigs{c, namespace}
}

func (c *FakeMioV1alpha1) ServiceConfigs(namespace string) v1alpha1.ServiceConfigInterface {
	return &FakeServiceConfigs{c, namespace}
}

func (c *FakeMioV1alpha1) SourceConfigs(namespace string) v1alpha1.SourceConfigInterface {
	return &FakeSourceConfigs{c, namespace}
}

func (c *FakeMioV1alpha1) TestConfigs(namespace string) v1alpha1.TestConfigInterface {
	return &FakeTestConfigs{c, namespace}
}

func (c *FakeMioV1alpha1) Testses(namespace string) v1alpha1.TestsInterface {
	return &FakeTestses{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMioV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
