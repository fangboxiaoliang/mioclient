package mio

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/hidevopsio/mioclient/pkg/apis/mio/v1alpha1"
	"github.com/hidevopsio/mioclient/pkg/client/clientset/versioned/fake"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPipelineCurd(t *testing.T) {
	name := "test"
	namespace := "demo-dev"
	clientSet := fake.NewSimpleClientset().MioV1alpha1()
	config := newPipeline(clientSet)
	config1 := &v1alpha1.Pipeline{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Namespace:namespace,
		},
	}
	result, err := config.Create(config1)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	list, err := config.List(namespace)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(list.Items))

	result, err = config.Get(name, namespace)
	assert.Equal(t, nil, err)
	//assert.Equal(t, name, result.Name)

	con := &v1alpha1.Pipeline{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	result, err = config.Update(name, namespace, con)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	err = config.Delete(name, namespace)
	assert.Equal(t, nil, err)

}