package mio

import (
	"github.com/hidevopsio/hiboot/pkg/app"
	"github.com/hidevopsio/hioak/starter/kube"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/mioclient/pkg/client/clientset/versioned/typed/mio/v1alpha1"
	)

type configuration struct {
	app.Configuration `depends:"kube"`

}

func init() {
	app.AutoConfiguration(newConfiguration)
}

func newConfiguration() *configuration {
	return &configuration{
	}
}

func (c *configuration) MioBuildConfig(restConfig *kube.RestConfig) *BuildConfig {
	clientSet, err := v1alpha1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("v1alpha1.NewForConfig %v", err)
		return nil
	}
	return newBuildConfig(clientSet)
}

func (c *configuration) MioPipeline(restConfig *kube.RestConfig) *Pipeline {
	clientSet, err := v1alpha1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("v1alpha1.NewForConfig %v", err)
		return nil
	}
	return newPipeline(clientSet)
}