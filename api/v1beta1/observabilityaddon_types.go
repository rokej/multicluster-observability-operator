/*
Copyright 2021.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StatusCondition contains condition information for an observability addon
type StatusCondition struct {
	Type               string                 `json:"type"`
	Status             metav1.ConditionStatus `json:"status"`
	LastTransitionTime metav1.Time            `json:"lastTransitionTime"`
	Reason             string                 `json:"reason"`
	Message            string                 `json:"message"`
}

// ObservabilityAddonStatus defines the observed state of ObservabilityAddon
type ObservabilityAddonStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Conditions []StatusCondition `json:"conditions"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ObservabilityAddon is the Schema for the observabilityaddon API
// +kubebuilder:resource:path=observabilityaddons,scope=Namespaced,shortName=oba
type ObservabilityAddon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ObservabilityAddonSpec   `json:"spec,omitempty"`
	Status ObservabilityAddonStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObservabilityAddonList contains a list of ObservabilityAddon
type ObservabilityAddonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObservabilityAddon `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ObservabilityAddon{}, &ObservabilityAddonList{})
}