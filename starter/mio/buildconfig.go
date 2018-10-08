package mio

import (
	"fmt"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/mioclient/pkg/apis/mio/v1alpha1"
	miov1 "github.com/hidevopsio/mioclient/pkg/client/clientset/versioned/typed/mio/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type BuildConfig struct {
	clientSet miov1.MioV1alpha1Interface
}

func newBuildConfig(clientSet miov1.MioV1alpha1Interface) *BuildConfig {
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
	config, err = b.clientSet.BuildConfigs(namespace).Create(config)
	if err != nil {
		return nil, err
	}
	return
}

func (b *BuildConfig) Get(name, namespace string) (config *v1alpha1.BuildConfig, err error) {
	log.Info("get config map :", name)
	result, err := b.clientSet.BuildConfigs(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *BuildConfig) Delete(name, namespace string) error {
	log.Info("get config map :", name)
	err := b.clientSet.BuildConfigs(namespace).Delete(name, &v1.DeleteOptions{})
	return err
}

func (b *BuildConfig) Update(name, namespace string, config *v1alpha1.BuildConfig) (*v1alpha1.BuildConfig, error) {
	log.Info("get build config :", name)
	result, err := b.clientSet.BuildConfigs(namespace).Update(config)
	return result, err
}

func (b *BuildConfig) Watch(name, namespace string) error {
	log.Info("get build config :", name)
	config := v1.ListOptions{
		LabelSelector: "app=" + name,
		Watch:         true,
	}
	w, err := b.clientSet.BuildConfigs(namespace).Watch(config)
	if err != nil {
		return err
	}
	for {
		select {
		case event, ok := <-w.ResultChan():
			if !ok {
				log.Errorf("failed on RC watching %v", ok)
				return fmt.Errorf("failed on RC watching %v", ok)
			}
			switch event.Type {
			case watch.Added:
				rc := event.Object.(*v1alpha1.BuildConfig)
				log.Debug(rc.Name)
			case watch.Modified:
				rc := event.Object.(*v1alpha1.BuildConfig)
				log.Debugf("RC: %s", rc.Name)
			case watch.Deleted:
				log.Info("Deleted: ", event.Object)
			default:
				log.Error("Failed")
			}
		}
	}
	return nil
}
