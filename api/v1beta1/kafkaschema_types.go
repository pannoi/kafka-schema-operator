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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type KafkaSchemaData struct {
	ConfigRef     string `json:"configRef"`     // Reference to configmap with values
	Format        string `json:"format"`        // avro/json/protobuf <-> schemeType:
	Compatibility string `json:"compatibility"` // BACKWARD | BACKWARD_TRANSITIVE | FORWARD | FORWARD_TRANSITIVE | FULL | FULL_TRANSITIVE | NONE
}

// KafkaSchemaSpec defines the desired state of KafkaSchema
type KafkaSchemaSpec struct {
	Name                  string          `json:"name"`
	SchemaSerializer      string          `json:"schemaSerializer"`
	AutoReconciliation    bool            `json:"autoReconciliation,omitempty"`
	DeletionPolicy        string          `json:"deletionPolicy,omitempty"`
	TerminationProtection bool            `json:"terminationProtection,omitempty"`
	Data                  KafkaSchemaData `json:"data"`
}

// KafkaSchemaStatus defines the observed state of KafkaSchema
type KafkaSchemaStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// KafkaSchema is the Schema for the kafkaschemas API
type KafkaSchema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KafkaSchemaSpec   `json:"spec,omitempty"`
	Status KafkaSchemaStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KafkaSchemaList contains a list of KafkaSchema
type KafkaSchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaSchema `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KafkaSchema{}, &KafkaSchemaList{})
}
