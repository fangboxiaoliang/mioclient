package mio

import (
	"github.com/hidevopsio/hiboot/pkg/app"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/hioak/starter/kube"
	"github.com/hidevopsio/mioclient/pkg/client/clientset/versioned/typed/mio/v1alpha1"
)

type configuration struct {
	app.Configuration `depends:"kube"`
}

func init() {
	app.AutoConfiguration(newConfiguration)
}

func newConfiguration() *configuration {
	return &configuration{}
}

func (c *configuration) BuildConfig(restConfig *kube.RestConfig) *BuildConfig {
	clientSet, err := v1alpha1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("v1alpha1.NewForConfig %v", err)
		return nil
	}
	return newBuildConfig(clientSet)
}

func (c *configuration) Build(restConfig *kube.RestConfig) *Build {
	clientSet, err := v1alpha1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("v1alpha1.NewForConfig %v", err)
		return nil
	}
	return newBuild(clientSet)
}

func (c *configuration) Pipeline(restConfig *kube.RestConfig) *Pipeline {
	clientSet, err := v1alpha1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("v1alpha1.NewForConfig %v", err)
		return nil
	}
	return newPipeline(clientSet)
}

func (c *configuration) PipelineConfig(restConfig *kube.RestConfig) *PipelineConfig {
	clientSet, err := v1alpha1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("v1alpha1.NewForConfig %v", err)
		return nil
	}
	return newPipelineConfig(clientSet)
}

func (c *configuration) SourceConfig(restConfig *kube.RestConfig) *SourceConfig {
	clientSet, err := v1alpha1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("v1alpha1.NewForConfig %v", err)
		return nil
	}
	return newSourceConfig(clientSet)
}

func (c *configuration) DeploymentConfig(restConfig *kube.RestConfig) *DeploymentConfig {
	clientSet, err := v1alpha1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("v1alpha1.NewForConfig %v", err)
		return nil
	}
	return newDeploymentConfig(clientSet)
}

func (c *configuration) GatewayConfig(restConfig *kube.RestConfig) *GatewayConfig {
	clientSet, err := v1alpha1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("v1alpha1.NewForConfig %v", err)
		return nil
	}
	return newGatewayConfig(clientSet)
}

func (c *configuration) ServiceConfig(restConfig *kube.RestConfig) *ServiceConfig {
	clientSet, err := v1alpha1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("v1alpha1.NewForConfig %v", err)
		return nil
	}
	return newServiceConfig(clientSet)
}

func (c *configuration) Deployment(restConfig *kube.RestConfig) *Deployment {
	clientSet, err := v1alpha1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("v1alpha1.NewForConfig %v", err)
		return nil
	}
	return newDeployment(clientSet)
}
