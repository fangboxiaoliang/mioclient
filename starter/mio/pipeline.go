package mio

import (
	"github.com/hidevopsio/mioclient/pkg/apis/mio/v1alpha1"
	"github.com/hidevopsio/hiboot/pkg/log"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"
	miov1 "github.com/hidevopsio/mioclient/pkg/client/clientset/versioned/typed/mio/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Pipeline struct {
	clientSet miov1.MioV1alpha1Interface
}

func newPipeline(clientSet miov1.MioV1alpha1Interface) *Pipeline {
	return &Pipeline{
		clientSet: clientSet,
	}
}

func (b *Pipeline) Create(name, namespace string) (config *v1alpha1.Pipeline, err error) {
	log.Debugf("config map create : %v", name)
	cm, err := b.Get(name, namespace)
	config = &v1alpha1.Pipeline{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	log.Debug("config map get :", cm)
	if err == nil {
		config, err = b.Update(name, namespace, config)
		return
	}
	config, err = b.clientSet.Pipelines(namespace).Create(config)
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

