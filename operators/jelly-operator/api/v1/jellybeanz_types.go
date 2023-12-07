/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// JellybeanzSpec defines the desired state of Jellybeanz
type JellybeanzSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Jellybeanz. Edit jellybeanz_types.go to remove/update
	Foo      string `json:"foo,omitempty"`
	StartJob bool   `json:"startJob,omitempty"`
}

// JellybeanzStatus defines the observed state of Jellybeanz
type JellybeanzStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	JobStatus bool `json:"jobStatus,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Jellybeanz is the Schema for the jellybeanzs API
type Jellybeanz struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JellybeanzSpec   `json:"spec,omitempty"`
	Status JellybeanzStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// JellybeanzList contains a list of Jellybeanz
type JellybeanzList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Jellybeanz `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Jellybeanz{}, &JellybeanzList{})
}
