// +build !test

package mio

import (
	"github.com/hidevopsio/hiboot/pkg/app"
	"github.com/hidevopsio/hiboot/pkg/app/web"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMioConfig(t *testing.T) {
	testApp := web.NewTestApplication(t).(app.ApplicationContext)

	t.Run("should get instance", func(t *testing.T) {
		mioBuildConfig := testApp.GetInstance("mioBuildConfig")
		assert.NotEqual(t, nil, mioBuildConfig)
	})
}
