package cni

import (
	"testing"

	"github.com/maistra/istio-operator/pkg/controller/common"
	"github.com/maistra/istio-operator/pkg/controller/common/test/assert"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func TestInitConfig_disablingCNI(t *testing.T) {
	operatorNamespace := "istio-operator"
	InitializeGlobals(operatorNamespace)()
	var m manager.Manager
	config, err := InitConfig(m)
	assert.Equals(err, nil, "", t)
	assert.Equals(config.Enabled, false, "", t)

}

func InitializeGlobals(operatorNamespace string) func() {
	return func() {
		// make sure globals are initialized for testing
		common.Config.OLM.CNIDisabled = true
	}
}
