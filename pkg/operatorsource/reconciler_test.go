package operatorsource

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/operator-framework/operator-marketplace/pkg/appregistry"
)

func TestReconcile(t *testing.T) {
	os.Setenv("KUBERNETES_CONFIG", getKubeConfigFilePath())
	fmt.Printf("KUBERNETES_CONFIG=%s\n", os.Getenv("KUBERNETES_CONFIG"))

	r := &reconciler{
		factory:   appregistry.NewClientFactory(),
		datastore: newDatastore(),
	}
	os := newOperatorSourceType("marketplace", "localhost")

	err := r.Reconcile(os)

	assert.NoError(t, err)
}

func getKubeConfigFilePath() string {
	home := os.Getenv("HOME")
	return filepath.Join(home, ".kube", "config")
}
