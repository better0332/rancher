package common

import (
	"testing"

	"github.com/rancher/rancher/pkg/settings"
	"github.com/rancher/types/apis/project.cattle.io/v3"
	"github.com/stretchr/testify/assert"
)

func Test_injectDefaultRegistry(t *testing.T) {
	testRegistry := "test.registry.com"
	err := settings.SystemDefaultRegistry.Set(testRegistry)
	assert.Nil(t, err, "failed to set system default registry settings")

	testCases := []struct {
		app  *v3.App
		want bool
	}{
		{
			app: &v3.App{
				Spec: v3.AppSpec{
					ExternalID: "catalog://?catalog=library&template=wordpress&version=2.1.11",
				},
			},
			want: false,
		},
		{
			app: &v3.App{
				Spec: v3.AppSpec{
					ExternalID: settings.SystemMonitoringCatalogID.Get(),
				},
			},
			want: true,
		},
	}

	for _, testCase := range testCases {
		testApp := testCase.app
		injectMap := injectDefaultRegistry(testApp)
		if !testCase.want {
			assert.Nilf(t, injectMap, "catalog id %s should not get default registry parameters", testApp.Spec.ExternalID)
		} else {
			v, _ := injectMap["systemDefaultRegistry"]
			assert.Equalf(t, testRegistry, v, "catalog id %s should not get default registry parameters", testApp.Spec.ExternalID)
		}
	}
}
