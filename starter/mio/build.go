package mio

import (
	"fmt"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/mioclient/pkg/apis/mio/v1alpha1"
	miov1alpha1 "github.com/hidevopsio/mioclient/pkg/client/clientset/versioned/typed/mio/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type Build struct {
	clientSet miov1alpha1.MioV1alpha1Interface
}

func newBuild(clientSet miov1alpha1.MioV1alpha1Interface) *Build {
	return &Build{
		clientSet: clientSet,
	}
}

func (b *Build) Create(build *v1alpha1.Build) (config *v1alpha1.Build, err error) {
	log.Debugf("config map create : %v", build.Name)
	config, err = b.clientSet.Builds(build.Namespace).Create(build)
	if err != nil {
		return nil, err
	}
	return
}

func (b *Build) Get(name, namespace string) (config *v1alpha1.Build, err error) {
	log.Info(fmt.Sprintf("get app %s in namespace %s:", name,namespace))
	result, err := b.clientSet.Builds(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *Build) Delete(name, namespace string) error {
	log.Info(fmt.Sprintf("delete app %s in namespace %s:", name,namespace))
	err := b.clientSet.Builds(namespace).Delete(name, &v1.DeleteOptions{})
	return err
}

func (b *Build) Update(name, namespace string, config *v1alpha1.Build) (*v1alpha1.Build, error) {
	log.Info(fmt.Sprintf("update app %s in namespace %s:", name,namespace))
	result, err := b.clientSet.Builds(namespace).Update(config)
	return result, err
}

func (b *Build) List(namespace string) (*v1alpha1.BuildList, error) {
	log.Info(fmt.Sprintf("list in namespace %s:", namespace))
	option := v1.ListOptions{
	}
	result, err := b.clientSet.Builds(namespace).List(option)
	return result, err
}

func (b *Build) Watch(listOptions v1.ListOptions,namespace,name string) (watch.Interface, error) {
	log.Info(fmt.Sprintf("watch app %s in namespace %s:", name,namespace))

	listOptions.LabelSelector = fmt.Sprintf("app=%s",name)
	listOptions.Watch = true

	w, err := b.clientSet.Builds(namespace).Watch(listOptions)
	if err != nil {
		return nil,err
	}
	return w,nil
}
