package mio

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/hidevopsio/mioclient/pkg/client/clientset/versioned/fake"
	"github.com/hidevopsio/mioclient/pkg/apis/mio/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestBuildconfigCurd(t *testing.T) {
	name := "test"
	namespace := "demo-dev"
	clientSet := fake.NewSimpleClientset()
	config := newBuildConfig(clientSet)
	result, err := config.Create(name, namespace)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	result, err = config.Get(name, namespace)
	assert.Equal(t, nil, err)
	//assert.Equal(t, name, result.Name)

	con := &v1alpha1.BuildConfig{
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
