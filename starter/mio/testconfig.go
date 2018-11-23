package mio

import (
	"fmt"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/mioclient/pkg/apis/mio/v1alpha1"
	miov1alpha1 "hidevops.io/mioclient/pkg/client/clientset/versioned/typed/mio/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type testConfig struct {
	clientSet miov1alpha1.MioV1alpha1Interface
}

func newTestConfig(clientSet miov1alpha1.MioV1alpha1Interface) *testConfig {
	return &testConfig{
		clientSet: clientSet,
	}
}

func (d *testConfig) Create(testConfig *v1alpha1.TestConfig) (result *v1alpha1.TestConfig, err error) {
	log.Debugf("deployConfig create : %v", testConfig.Name)
	result, err = d.clientSet.TestConfigs(testConfig.Namespace).Create(testConfig)
	if err != nil {
		return nil, err
	}
	return
}

func (d *testConfig) Get(name, namespace string) (config *v1alpha1.TestConfig, err error) {
	log.Info(fmt.Sprintf("get testConfig %s in namespace %s", name, namespace))
	result, err := d.clientSet.TestConfigs(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *testConfig) Delete(name, namespace string) error {
	log.Info(fmt.Sprintf("delete testConfig %s in namespace %s", name, namespace))
	err := d.clientSet.TestConfigs(namespace).Delete(name, &v1.DeleteOptions{})
	return err
}

func (d *testConfig) Update(name, namespace string, testConfig *v1alpha1.TestConfig) (*v1alpha1.TestConfig, error) {
	log.Info(fmt.Sprintf("update testConfig %s in namespace %s", name, namespace))
	result, err := d.clientSet.TestConfigs(namespace).Update(testConfig)
	return result, err
}

func (d *testConfig) List(namespace string, option v1.ListOptions) (*v1alpha1.TestConfigList, error) {
	log.Info(fmt.Sprintf("list deployConfig in namespace %s", namespace))
	result, err := d.clientSet.TestConfigs(namespace).List(option)
	return result, err
}

func (d *testConfig) Watch(listOptions v1.ListOptions, namespace string) (watch.Interface, error) {
	log.Info(fmt.Sprintf("watch label for %s testconfig，in the namespace %s", listOptions.LabelSelector, namespace))
	w, err := d.clientSet.TestConfigs(namespace).Watch(listOptions)
	if err != nil {
		return nil, err
	}
	return w, nil
}
