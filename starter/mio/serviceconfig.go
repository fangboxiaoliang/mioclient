package mio

import (
	"fmt"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/mioclient/pkg/apis/mio/v1alpha1"
	miov1 "github.com/hidevopsio/mioclient/pkg/client/clientset/versioned/typed/mio/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type ServiceConfig struct {
	clientSet miov1.MioV1alpha1Interface
}

func newServiceConfig(clientSet miov1.MioV1alpha1Interface) *ServiceConfig {
	return &ServiceConfig{
		clientSet: clientSet,
	}
}

func (s *ServiceConfig) Create(serviceConfig *v1alpha1.ServiceConfig) (config *v1alpha1.ServiceConfig, err error) {
	log.Debugf("serviceConfig create : %v", serviceConfig.Name)
	config, err = s.clientSet.ServiceConfigs(serviceConfig.Namespace).Create(serviceConfig)
	if err != nil {
		return nil, err
	}
	return
}

func (s *ServiceConfig) Get(name, namespace string) (config *v1alpha1.ServiceConfig, err error) {
	log.Info("get ServiceConfigs :", name)
	result, err := s.clientSet.ServiceConfigs(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ServiceConfig) Delete(name, namespace string) error {
	log.Info("delete ServiceConfigs :", name)
	err := s.clientSet.ServiceConfigs(namespace).Delete(name, &v1.DeleteOptions{})
	return err
}

func (s *ServiceConfig) Update(name, namespace string, config *v1alpha1.ServiceConfig) (*v1alpha1.ServiceConfig, error) {
	log.Info("update ServiceConfigs :", name)
	result, err := s.clientSet.ServiceConfigs(namespace).Update(config)
	return result, err
}

func (s *ServiceConfig) List(namespace string, option v1.ListOptions) (*v1alpha1.ServiceConfigList, error) {
	log.Info(fmt.Sprintf("list ServiceConfigs in namespace %s:", namespace))
	result, err := s.clientSet.ServiceConfigs(namespace).List(option)
	return result, err
}

func (s *ServiceConfig) Watch(listOptions v1.ListOptions, namespace, name string) (watch.Interface, error) {
	log.Info(fmt.Sprintf("watch ServiceConfigs app %s in namespace %s:", name, namespace))

	listOptions.LabelSelector = fmt.Sprintf("app=%s", name)
	listOptions.Watch = true

	w, err := s.clientSet.ServiceConfigs(namespace).Watch(listOptions)
	if err != nil {
		return nil, err
	}
	return w, nil
}
