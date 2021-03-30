// +groupName=stable.example.com
// +kubebuilder:object:generator=true
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubeapis,singular=kubeapi,shortName=kapi,categories={}
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name=Age,type=date
// +kubebuilder:printcolumn:JSONPath=".status.deploymentName",name=Deployment,type=string
// +kubebuilder:printcolumn:JSONPath=".status.serviceName",name=Service,type=string
// +kubebuilder:printcolumn:JSONPath=".status.replicas",name=Replicas,type=integer
// +kubebuilder:printcolumn:JSONPath=".status.phase",name=Status,type=string
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Disappointment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DisappointmentSpec `json:"spec"`
	// +optional
	Status DisappointmentStatus `json:"status"`
}

type DisappointmentSpec struct {
	// +optional
	Version string `json:"version,omitempty"`

	// +kubebuilder:default=1
	// +optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=20
	Replicas *int32 `json:"replicas"`

	// +optional
	DeploymentName string `json:"deploymentName"`

	// +optional
	ServiceName string `json:"serviceName"`

	// +kubebuilder:default=ClusterIP
	ServiceType ServiceType `json:"serviceType"`

	Container ContainerSpec `json:"container"`
}

// +kubebuilder:validation:Enum=ClusterIP;NodePort
type ServiceType string

type ContainerSpec struct {
	// Container image of the Api Server
	Image string `json:"image"`

	// Container port of Api Server
	ContainerPort int32 `json:"containerPort"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DisappointmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Disappointment `json:"items"`
}

type DisappointmentStatus struct {
	Phase    string `json:"phase"`
	Replicas int32  `json:"replicas"`
}
