package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo is a specification for a Foo resource
type TestConfig struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of the Test.
	// +optional
	Spec TestConfigSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Most recently observed status of the Test.
	// +optional
	Status TestConfigStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}


type TestCmd struct {
	ExecType      string   `protobuf:"bytes,1,opt,name=execType,proto3" json:"execType,omitempty"`
	Script        string   `protobuf:"bytes,2,opt,name=Script,proto3" json:"Script,omitempty"`
	CommandName   string   `protobuf:"bytes,3,opt,name=commandName,proto3" json:"commandName,omitempty"`
	CommandParams []string `protobuf:"bytes,4,rep,name=commandParams,proto3" json:"commandParams,omitempty"`
}


type TestConfigSpec struct {
	CodeType string		  `json:"codeType" protobuf:"bytes,1,opt,name=codeType"`
	TestCmd  TestCmd      `json:"testCmd" protobuf:"bytes,2,opt,name=testCmd"`
}

type TestConfigStatus struct {
	LastVersion int `json:"lastVersion,omitempty" protobuf:"bytes,1,opt,name=lastVersion"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo resources
type TestConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []TestConfig `json:"items"`
}
