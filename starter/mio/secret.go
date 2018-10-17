package mio

import (
	"github.com/hidevopsio/mioclient/pkg/apis/mio/v1alpha1"
	"github.com/hidevopsio/hiboot/pkg/log"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"
	"k8s.io/apimachinery/pkg/watch"
	miov1alpha1 "github.com/hidevopsio/mioclient/pkg/client/clientset/versioned/typed/mio/v1alpha1"
)

type Secret struct {
	clientSet miov1alpha1.MioV1alpha1Interface
}

func newSecret(clientSet miov1alpha1.MioV1alpha1Interface) *Secret {
	return &Secret{
		clientSet: clientSet,
	}
}

func (b *Secret) Create(secret *v1alpha1.Secret) (result *v1alpha1.Secret, err error) {
	log.Debugf("Secret create : %v", secret.Name)
	result, err = b.clientSet.Secrets(secret.Namespace).Create(secret)
	if err != nil {
		return nil, err
	}
	return
}

func (b *Secret) Get(name, namespace string) (config *v1alpha1.Secret, err error) {
	log.Info("get Secret :", name)
	result, err := b.clientSet.Secrets(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *Secret) Delete(name, namespace string) error {
	log.Info("delete Secret :", name)
	err := b.clientSet.Secrets(namespace).Delete(name, &v1.DeleteOptions{})
	return err
}

func (b *Secret) Update(name, namespace string, secret *v1alpha1.Secret) (*v1alpha1.Secret, error) {
	log.Info("update Secret :", name)
	result, err := b.clientSet.Secrets(namespace).Update(secret)
	return result, err
}

func (b *Secret) List(namespace string) (*v1alpha1.SecretList, error) {
	log.Info(fmt.Sprintf("list in namespace %s:", namespace))
	option := v1.ListOptions{
	}
	result, err := b.clientSet.Secrets(namespace).List(option)
	return result, err
}

func (b *Secret) Watch(listOptions v1.ListOptions, namespace, name string) (watch.Interface, error) {
	log.Info(fmt.Sprintf("watch app %s in namespace %s:", name, namespace))

	listOptions.LabelSelector = fmt.Sprintf("app=%s", name)
	listOptions.Watch = true

	w, err := b.clientSet.Secrets(namespace).Watch(listOptions)
	if err != nil {
		return nil, err
	}
	return w, nil
}
