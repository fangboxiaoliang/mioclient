package mio

import (
	"github.com/hidevopsio/mioclient/pkg/apis/mio/v1alpha1"
	"github.com/hidevopsio/hiboot/pkg/log"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"
	miov1 "github.com/hidevopsio/mioclient/pkg/client/clientset/versioned/typed/mio/v1alpha1"
	)

type Pipeline struct {
	clientSet miov1.MioV1alpha1Interface
}

func newPipeline(clientSet miov1.MioV1alpha1Interface) *Pipeline {
	return &Pipeline{
		clientSet: clientSet,
	}
}

func (b *Pipeline) Create(pipeline *v1alpha1.Pipeline) (config *v1alpha1.Pipeline, err error) {
	log.Debugf("config map create : %v", pipeline.Name)
	cm, err := b.Get(pipeline.Name, pipeline.Namespace)
	log.Debug("config map get :", cm)
	if err == nil {
		config, err = b.Update(pipeline.Name, pipeline.Namespace, config)
		return
	}
	config, err = b.clientSet.Pipelines(pipeline.Namespace).Create(pipeline)
	if err != nil {
		return nil, err
	}
	return
}

func (b *Pipeline) Get(name, namespace string) (config *v1alpha1.Pipeline, err error) {
	log.Info("get config map :", name)
	result, err := b.clientSet.Pipelines(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *Pipeline) Delete(name, namespace string) error {
	log.Info("get config map :", name)
	err := b.clientSet.Pipelines(namespace).Delete(name, &v1.DeleteOptions{})
	return err
}

func (b *Pipeline) Update(name, namespace string, config *v1alpha1.Pipeline) (*v1alpha1.Pipeline, error) {
	log.Info("get build config :", name)
	result, err := b.clientSet.Pipelines(namespace).Update(config)
	return result, err
}

func (b *Pipeline) List(namespace string) (*v1alpha1.PipelineList, error) {
	log.Info(fmt.Sprintf("list in namespace %s:", namespace))
	option := v1.ListOptions{
	}
	result, err := b.clientSet.Pipelines(namespace).List(option)
	return result, err
}

func (b *Pipeline) Watch(name, namespace string) (*v1alpha1.Pipeline, error) {
	log.Info("get build config :", name)
	config := v1.ListOptions{
		LabelSelector: "app=" + name,
		Watch:         true,
	}
	w, err := b.clientSet.Pipelines(namespace).Watch(config)
	if err != nil {
		return nil, err
	}
	for {
		select {
		case event, ok := <-w.ResultChan():
			if !ok {
				log.Errorf("failed on RC watching %v", ok)
				return nil, fmt.Errorf("failed on RC watching %v", ok)
			}
			rc := event.Object.(*v1alpha1.Pipeline)
			return rc, nil
		}
	}
}

