package stub

import (
	"context"

	"github.com/operator-framework/operator-marketplace/pkg/apis/marketplace/v1alpha1"
	"github.com/operator-framework/operator-marketplace/pkg/operatorsource"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
)

func NewHandler() sdk.Handler {
	osHandler, _, err := operatorsource.NewHandler()
	if err != nil {
		logrus.Fatalf("failed to initialize - %v", err)
	}

	return &Handler{
		operatorSourceHanlder: osHandler,
	}
}

type Handler struct {
	operatorSourceHanlder operatorsource.Handler
}

// Handle function for handling CatalogSourceConfig and OperatorSource CR events
func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.CatalogSourceConfig:
		// Ignore the delete event as the garbage collector will clean up the created resources as per the OwnerReference
		if event.Deleted {
			logrus.Infof("Deleted %s CatalogSourceConfig in %s namespace", o.Name, o.Spec.TargetNamespace)
			return nil
		}
		return createCatalogSource(o)

	case *v1alpha1.OperatorSource:
		if err := h.operatorSourceHanlder.Handle(ctx, event); err != nil {
			logrus.Errorf("reconciliation error: %v", err)
			return err
		}
	}

	return nil
}
