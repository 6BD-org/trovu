/*


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

// PayloadKeyValPair organizes extra service information as key-value pairs
type PayloadKeyValPair struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

// Payload carries extra information of a service
type Payload struct {
	KeyValPairs []PayloadKeyValPair `json:"keyValPairs"`
}

// ServiceEntry is one single entry for a service, which may contain multiple hosts
type ServiceEntry struct {
	ServiceName string  `json:"serviceName"`
	ServiceHost string  `json:"serviceHosts"`
	Payload     Payload `json:"payload,omitempty"`
}

// PathFinderSpec defines the desired state of PathFinder
type PathFinderSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of PathFinder. Edit PathFinder_types.go to remove/update
	ClusterDomain string `json:"clusterDomain,omitempty"`
	Region        string `json:"region"`
}

// PathFinderStatus defines the observed state of PathFinder
type PathFinderStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ServiceEntries []ServiceEntry `json:"serviceEntries,omitempty"`
}

// +kubebuilder:object:root=true

// PathFinder is the Schema for the pathfinders API
type PathFinder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PathFinderSpec   `json:"spec,omitempty"`
	Status PathFinderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PathFinderList contains a list of PathFinder
type PathFinderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PathFinder `json:"items"`
}

// FindServiceEntry find service entry by name, if not found, second returned val
func (status PathFinderStatus) FindServiceEntry(name string) (ServiceEntry, bool) {
	for _, entry := range status.ServiceEntries {
		if entry.ServiceName == name {
			return entry, true
		}
	}
	return ServiceEntry{}, false
}

func init() {
	SchemeBuilder.Register(&PathFinder{}, &PathFinderList{})
}
