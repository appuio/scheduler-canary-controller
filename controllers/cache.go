package controllers

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewCacheOptions creates cache options that restricts the pod cache to pods with the `instanceLabel` set.
func NewCacheOptions() (cache.Options, error) {
	lblReq, err := labels.NewRequirement(instanceLabel, selection.Exists, nil)
	if err != nil {
		return cache.Options{}, fmt.Errorf("failed to create label requirement for pod cache: %w", err)
	}
	return cache.Options{
		ByObject: map[client.Object]cache.ByObject{
			&corev1.Pod{}: {
				Label: labels.NewSelector().Add(*lblReq),
			},
		},
	}, nil
}
