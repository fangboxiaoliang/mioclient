package mio

import (
	"github.com/hidevopsio/mioclient/pkg/apis/mio/v1alpha1"
	"github.com/hidevopsio/hiboot/pkg/log"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"
	miov1 "github.com/hidevopsio/mioclient/pkg/client/clientset/versioned/typed/mio/v1alpha1"
	"k8s.io/apimachinery/pkg/watch"
)

type PipelineConfig struct {
	clientSet miov1.MioV1alpha1Interface
}

func newPipelineConfig(clientSet miov1.MioV1alpha1Interface) *PipelineConfig {
	return &PipelineConfig{
		clientSet: clientSet,
	}
}

func (b *PipelineConfig) Create(pipeline *v1alpha1.PipelineConfig) (config *v1alpha1.PipelineConfig, err error) {
	log.Debugf("pipelineConfig create : %v", pipeline.Name)
	config, err = b.clientSet.PipelineConfigs(pipeline.Namespace).Create(pipeline)
	if err != nil {
		return nil, err
	}
	return
}

func (b *PipelineConfig) Get(name, namespace string) (config *v1alpha1.PipelineConfig, err error) {
	log.Info("get pipelineConfig :", name)
	result, err := b.clientSet.PipelineConfigs(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *PipelineConfig) Delete(name, namespace string) error {
	log.Info("delete pipelineConfig :", name)
	err := b.clientSet.PipelineConfigs(namespace).Delete(name, &v1.DeleteOptions{})
	return err
}

func (b *PipelineConfig) Update(name, namespace string, config *v1alpha1.PipelineConfig) (*v1alpha1.PipelineConfig, error) {
	log.Info("update pipelineConfig :", name)
	result, err := b.clientSet.PipelineConfigs(namespace).Update(config)
	return result, err
}

func (b *PipelineConfig) List(namespace string) (*v1alpha1.PipelineConfigList, error) {
	log.Info(fmt.Sprintf("list pipelineConfig in namespace %s:", namespace))
	option := v1.ListOptions{
	}
	result, err := b.clientSet.PipelineConfigs(namespace).List(option)
	return result, err
}

func (b *PipelineConfig) Watch(listOptions v1.ListOptions,namespace,name string) (watch.Interface, error) {
	log.Info(fmt.Sprintf("watch pipelineConfig app %s in namespace %s:", name,namespace))

	listOptions.LabelSelector = fmt.Sprintf("app=%s",name)
	listOptions.Watch = true

	w, err := b.clientSet.PipelineConfigs(namespace).Watch(listOptions)
	if err != nil {
		return nil,err
	}
	return w,nil
}
