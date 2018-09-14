package operatorsource

import (
	"context"

	"github.com/operator-framework/operator-marketplace/pkg/apis/marketplace/v1alpha1"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	// Handle hanldes a new event generated by the operator sdk
	Handle(ctx context.Context, event sdk.Event) error
}

type handler struct {
	reconciler Reconciler
}

func (h *handler) Handle(ctx context.Context, event sdk.Event) error {
	os := event.Object.(*v1alpha1.OperatorSource)

	if event.Deleted {
		logrus.Infof("No action taken, object has been deleted [type=%s object=%s/%s]",
			os.TypeMeta.Kind, os.ObjectMeta.Namespace, os.ObjectMeta.Name)

		return nil
	}

	reconciled, err := h.reconciler.IsAlreadyReconciled(os)
	if err != nil {
		return err
	}

	if reconciled {
		logrus.Infof("Already reconciled, no action taken [type=%s object=%s/%s]",
			os.TypeMeta.Kind, os.ObjectMeta.Namespace, os.ObjectMeta.Name)

		return nil
	}

	if err := h.reconciler.Reconcile(os); err != nil {
		return err
	}

	logrus.Infof("Reconciliation completed successfully [type=%s object=%s/%s]",
		os.TypeMeta.Kind, os.ObjectMeta.Namespace, os.ObjectMeta.Name)
	return nil
}