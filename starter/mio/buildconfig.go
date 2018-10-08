package mio

import (
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/mioclient/pkg/client/clientset/versioned"
	"github.com/hidevopsio/mioclient/pkg/apis/mio/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BuildConfig struct {
	clientSet versioned.Interface
}

func newBuildConfig(clientSet versioned.Interface) *BuildConfig {
	return &BuildConfig{
		clientSet: clientSet,
	}
}

func (b *BuildConfig) Create(name, namespace string) (config *v1alpha1.BuildConfig, err error) {
	log.Debugf("config map create : %v", name)
	cm, err := b.Get(name, namespace)
	config = &v1alpha1.BuildConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	log.Debug("config map get :", cm)
	if err == nil {
		config, err = b.Update(name, namespace, config)
		return
	}
	config, err = b.clientSet.MioV1alpha1().BuildConfigs(namespace).Create(config)
	if err != nil {
		return nil, err
	}
	return
}


func (b *BuildConfig) Get(name, namespace string) (config *v1alpha1.BuildConfig, err error) {
	log.Info("get config map :", name)
	result, err := b.clientSet.MioV1alpha1().BuildConfigs(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *BuildConfig) Delete(name, namespace string) error {
	log.Info("get config map :", name)
	err := b.clientSet.MioV1alpha1().BuildConfigs(namespace).Delete(name, &v1.DeleteOptions{})
	return err
}

func (b *BuildConfig) Update(name, namespace string, config *v1alpha1.BuildConfig) (*v1alpha1.BuildConfig, error) {
	log.Info("get build config :", name)
	result, err := b.clientSet.MioV1alpha1().BuildConfigs(namespace).Update(config)
	return result, err
}
