package operatorsource

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/operator-framework/operator-marketplace/pkg/appregistry"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
)

func TestHandle(t *testing.T) {
	os.Setenv("KUBERNETES_CONFIG", "/home/akashem/.kube/config")
	fmt.Printf("KUBERNETES_CONFIG=%s\n", os.Getenv("KUBERNETES_CONFIG"))

	r := &reconciler{
		factory:   appregistry.NewClientFactory(),
		datastore: newDatastore(),
	}

	os := newOperatorSourceType("marketplace", "localhost")

	event := sdk.Event{
		Deleted: false,
		Object:  os,
	}

	handler := &handler{reconciler: r}
	handler.Handle(context.Background(), event)
}
