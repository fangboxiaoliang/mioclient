package mio

import (
	"fmt"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/mioclient/pkg/apis/mio/v1alpha1"
	miov1alpha1 "hidevops.io/mioclient/pkg/client/clientset/versioned/typed/mio/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type SourceConfig struct {
	clientSet miov1alpha1.MioV1alpha1Interface
}

func newSourceConfig(clientSet miov1alpha1.MioV1alpha1Interface) *SourceConfig {
	return &SourceConfig{
		clientSet: clientSet,
	}
}

func (b *SourceConfig) Create(build *v1alpha1.SourceConfig) (config *v1alpha1.SourceConfig, err error) {
	log.Debugf("source config create : %v", build.Name)
	config, err = b.clientSet.SourceConfigs(build.Namespace).Create(build)
	if err != nil {
		return nil, err
	}
	return
}

func (b *SourceConfig) Get(name, namespace string) (config *v1alpha1.SourceConfig, err error) {
	log.Info(fmt.Sprintf("get source config app %s in namespace %s:", name, namespace))
	result, err := b.clientSet.SourceConfigs(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *SourceConfig) Delete(name, namespace string) error {
	log.Info(fmt.Sprintf("delete source config app %s in namespace %s:", name, namespace))
	err := b.clientSet.SourceConfigs(namespace).Delete(name, &v1.DeleteOptions{})
	return err
}

func (b *SourceConfig) Update(name, namespace string, config *v1alpha1.SourceConfig) (*v1alpha1.SourceConfig, error) {
	log.Info(fmt.Sprintf("update source config app %s in namespace %s:", name, namespace))
	result, err := b.clientSet.SourceConfigs(namespace).Update(config)
	return result, err
}

func (b *SourceConfig) List(namespace string, option v1.ListOptions) (*v1alpha1.SourceConfigList, error) {
	log.Info(fmt.Sprintf("list source config in namespace %s:", namespace))
	result, err := b.clientSet.SourceConfigs(namespace).List(option)
	return result, err
}

func (b *SourceConfig) Watch(listOptions v1.ListOptions, namespace, name string) (watch.Interface, error) {
	log.Info(fmt.Sprintf("watch source config app %s in namespace %s:", name, namespace))
	listOptions.LabelSelector = fmt.Sprintf("app=%s", name)
	listOptions.Watch = true

	w, err := b.clientSet.SourceConfigs(namespace).Watch(listOptions)
	if err != nil {
		return nil, err
	}
	return w, nil
}
