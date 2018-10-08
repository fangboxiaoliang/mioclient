package mio_test

import (
	"testing"
	"github.com/hidevopsio/hiboot/pkg/app/web"
	"github.com/hidevopsio/hiboot/pkg/app"
	"github.com/stretchr/testify/assert"
)

func TestMioConfig(t *testing.T) {
	testApp := web.NewTestApplication(t).(app.ApplicationContext)

	t.Run("should get instance", func(t *testing.T) {
		kubeBuildConfig := testApp.GetInstance("mioBuildConfig")
		assert.NotEqual(t, nil, kubeBuildConfig)
	})
}
