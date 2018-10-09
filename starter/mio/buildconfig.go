package mio

import (
	"fmt"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/mioclient/pkg/apis/mio/v1alpha1"
	miov1alpha1 "github.com/hidevopsio/mioclient/pkg/client/clientset/versioned/typed/mio/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type BuildConfig struct {
	clientSet miov1alpha1.MioV1alpha1Interface
}

func newBuildConfig(clientSet miov1alpha1.MioV1alpha1Interface) *BuildConfig {
	return &BuildConfig{
		clientSet: clientSet,
	}
}

func (b *BuildConfig) Create(build *v1alpha1.BuildConfig) (config *v1alpha1.BuildConfig, err error) {
	log.Debugf("config map create : %v", build.Name)
	cm, err := b.Get(build.Name, build.Namespace)
	log.Debug("config map get :", cm)
	if err == nil {
		config, err = b.Update(build.Name, build.Namespace, config)
		return
	}
	config, err = b.clientSet.BuildConfigs(build.Namespace).Create(build)
	if err != nil {
		return nil, err
	}
	return
}

func (b *BuildConfig) Get(name, namespace string) (config *v1alpha1.BuildConfig, err error) {
	log.Info(fmt.Sprintf("get app %s in namespace %s:", name,namespace))
	result, err := b.clientSet.BuildConfigs(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *BuildConfig) Delete(name, namespace string) error {
	log.Info(fmt.Sprintf("delete app %s in namespace %s:", name,namespace))
	err := b.clientSet.BuildConfigs(namespace).Delete(name, &v1.DeleteOptions{})
	return err
}

func (b *BuildConfig) Update(name, namespace string, config *v1alpha1.BuildConfig) (*v1alpha1.BuildConfig, error) {
	log.Info(fmt.Sprintf("update app %s in namespace %s:", name,namespace))
	result, err := b.clientSet.BuildConfigs(namespace).Update(config)
	return result, err
}

func (b *BuildConfig) Watch(listOptions v1.ListOptions,namespace,name string) (watch.Interface, error) {
	log.Info(fmt.Sprintf("watch app %s in namespace %s:", name,namespace))

	listOptions.LabelSelector = fmt.Sprintf("app=%s",name)
	listOptions.Watch = true

	w, err := b.clientSet.BuildConfigs(namespace).Watch(listOptions)
	if err != nil {
		return nil,err
	}
	return w,nil
}
