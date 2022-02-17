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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MongoserviceSpec defines the desired state of Mongoservice
type MongoserviceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Mongoservice. Edit mongoservice_types.go to remove/update
	Size int32 `json:"size"`
}

// MongoserviceStatus defines the observed state of Mongoservice
type MongoserviceStatus struct {
	Nodes []string `json:"nodes"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Mongoservice is the Schema for the mongoservices API
type Mongoservice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MongoserviceSpec   `json:"spec,omitempty"`
	Status MongoserviceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MongoserviceList contains a list of Mongoservice
type MongoserviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mongoservice `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Mongoservice{}, &MongoserviceList{})
}
