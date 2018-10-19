package mio

import (
	"github.com/hidevopsio/mioclient/pkg/apis/mio/v1alpha1"
	"github.com/hidevopsio/mioclient/pkg/client/clientset/versioned/fake"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/hidevopsio/hiboot/pkg/log"
)

func TestSourceConfigCurd(t *testing.T) {
	name := "test"
	name1 := "test1"
	namespace := "demo-dev"
	version := "v1"
	clientSet := fake.NewSimpleClientset().MioV1alpha1()
	config := newSourceConfig(clientSet)
	config1 := &v1alpha1.SourceConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels: map[string]string{
				"app":name,
			},
		},
	}
	result, err := config.Create(config1)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	config2 := &v1alpha1.SourceConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name1,
			Namespace: namespace,
			Annotations: map[string]string{
				"mio.io/build.number":"3",
			},
			Labels: map[string]string{
				"app":name,
				"version": version,
			},
		},
	}
	result, err = config.Create(config2)
	assert.Equal(t, nil, err)
	assert.Equal(t, name1, result.Name)
	option := v1.ListOptions{

		LabelSelector: "app=" + name,
	}
	list, err := config.List(namespace, option)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(list.Items))

	result, err = config.Get(name, namespace)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	con := &v1alpha1.SourceConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	result, err = config.Update(name, namespace, con)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	err = config.Delete(name, namespace)
	assert.Equal(t, nil, err)

	listOptions := metav1.ListOptions{}
	i, err := config.Watch(listOptions, name, namespace)
	log.Infof("i", i)
	assert.Equal(t, nil, err)
}
