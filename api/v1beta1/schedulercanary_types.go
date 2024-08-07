/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1 "k8s.io/api/core/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SchedulerCanarySpec defines the desired state of SchedulerCanary
type SchedulerCanarySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// MaxPodCompletionTimeout is the maximum amount of time to wait for a pod to complete.
	// After the timeout expires, the pod will be deleted, the canary will be marked as failed, and a new canary will be created.
	// The default is 15 minutes.
	MaxPodCompletionTimeout metav1.Duration `json:"maxPodCompletionTimeout,omitempty"`

	// Interval is the interval at which a canary pod will be created.
	// The interval is not very precise.
	// The default is 1 minute.
	Interval metav1.Duration `json:"interval,omitempty"`

	// ForbidParallelRuns will prevent the creation of a new canary pod if there is already a canary pod running.
	// The default is false.
	ForbidParallelRuns bool `json:"forbidParallelRuns,omitempty"`

	// PodTemplate is the pod template to use for the canary pods.
	PodTemplate corev1.PodTemplateSpec `json:"podTemplate,omitempty"`
}

func (s SchedulerCanarySpec) MaxPodCompletionTimeoutWithDefault() time.Duration {
	d := s.MaxPodCompletionTimeout.Duration
	if d == 0 {
		return 15 * time.Minute
	}
	return d
}

func (s SchedulerCanarySpec) IntervalWithDefault() time.Duration {
	d := s.Interval.Duration
	if d == 0 {
		return time.Minute
	}
	return d
}

// SchedulerCanaryStatus defines the observed state of SchedulerCanary
type SchedulerCanaryStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// LastCanaryCreatedAt is the timestamp of the last canary creation.
	LastCanaryCreatedAt metav1.Time `json:"lastCanaryCreatedAt,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="LASTEXEC",type="string",JSONPath=`.status.lastCanaryCreatedAt`

// SchedulerCanary is the Schema for the schedulercanaries API
type SchedulerCanary struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SchedulerCanarySpec   `json:"spec,omitempty"`
	Status SchedulerCanaryStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SchedulerCanaryList contains a list of SchedulerCanary
type SchedulerCanaryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SchedulerCanary `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SchedulerCanary{}, &SchedulerCanaryList{})
}
