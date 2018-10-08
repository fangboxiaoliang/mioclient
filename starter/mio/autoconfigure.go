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

func (c *configuration) KubeBuildConfig(restConfig *kube.RestConfig) *BuildConfig {
	clientSet, err := v1alpha1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("oauthv1.NewForConfig %v", err)
		return nil
	}
	return newBuildConfig(clientSet)
}
