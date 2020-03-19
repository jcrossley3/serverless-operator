package common

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/types"
	servingv1alpha1 "knative.dev/serving-operator/pkg/apis/serving/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var Log = logf.Log.WithName("knative").WithName("openshift")

// Configure is a  helper to set a value for a key, potentially overriding existing contents.
func Configure(ks *servingv1alpha1.KnativeServing, cm, key, value string) bool {
	if ks.Spec.Config == nil {
		ks.Spec.Config = map[string]map[string]string{}
	}

	old, found := ks.Spec.Config[cm][key]
	if found && value == old {
		return false
	}

	if ks.Spec.Config[cm] == nil {
		ks.Spec.Config[cm] = map[string]string{}
	}

	ks.Spec.Config[cm][key] = value
	Log.Info("Configured", "map", cm, key, value, "old value", old)
	return true
}

// IngressNamespace returns namespace where ingress is deployed.
func IngressNamespace(servingNamespace string) string {
	return servingNamespace + "-ingress"
}

// NamespaceFromIngressNamespace returns namespace after trim trailing "-ingress".
func NamespaceFromIngressNamespace(ingressNamespace string) string {
	return strings.TrimSuffix(ingressNamespace, "-ingress")
}

// KnativeServing gets KnativeServing in the specified namespace
func KnativeServing(api client.Client, ksNamespace string) (types.NamespacedName, error) {
	ksList := &servingv1alpha1.KnativeServingList{}
	if err := api.List(context.TODO(), &client.ListOptions{Namespace: ksNamespace}, ksList); err != nil {
		return types.NamespacedName{}, fmt.Errorf("failed to List KnativeServing: %w", err)
	}
	if len(ksList.Items) == 0 {
		return types.NamespacedName{}, fmt.Errorf("KnativeServing is not found in namespace %q", ksNamespace)
	}
	return types.NamespacedName{
		Namespace: ksList.Items[0].Namespace,
		Name:      ksList.Items[0].Name,
	}, nil
}
