package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo is a specification for a Foo resource
type Notify struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of the Deployment.
	// +optional
	Spec NotifySpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Most recently observed status of the Deployment.
	// +optional
	Status BuildStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type NotifySpec struct {
	RedirectUrl string `json:"redirectUrl" protobuf:"bytes,1,opt,name=hosts"`
	Roles       []Role `json:"roles" protobuf:"bytes,2,opt,name=roles"`
	Reason      string `json:"reason" protobuf:"bytes,3,opt,name=reason"`
	Version     string `json:"version" protobuf:"bytes,4,opt,name=version"`
	Profile     string `json:"profile" protobuf:"bytes,5,opt,name=version"`
}

type Role struct {
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	Pass bool   `json:"pass" protobuf:"bytes,2,opt,name=pass"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo resources
type NotifyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Notify `json:"items"`
}
