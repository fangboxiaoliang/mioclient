package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo is a specification for a Foo resource
type BuildConfig struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of the Deployment.
	// +optional
	Spec BuildConfigSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Most recently observed status of the Deployment.
	// +optional
	Status metav1.Status `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type BuildConfigSpec struct {
	CloneConfig CloneConfig `json:"cloneConfig"  protobuf:"bytes,1,opt,name=cloneConfig"`
	App         string      `json:"app"  protobuf:"bytes,1,opt,name=app"`
	//代码类型
	CodeType    string      `json:"codeType"  protobuf:"bytes,1,opt,name=codeType"`
	CompileCmd  []CMD       `json:"compileCmd"  protobuf:"bytes,1,opt,name=compileCmd"`
	//获取类型
	CloneType   string      `json:"cloneType"  protobuf:"bytes,1,opt,name=cloneType"`
	//基础镜像包
	S2iImage    string      `json:"s2iImage"  protobuf:"bytes,1,opt,name=s2iImage"`
	//版本
	Tags        []string    `json:"tags"  protobuf:"bytes,1,opt,name=tags"`
	DockerFile  []string    `json:"dockerFile"  protobuf:"bytes,1,opt,name=dockerFile"`
}

type CloneConfig struct {
	Url      string `json:"url"  protobuf:"bytes,1,opt,name=url"`
	Branch   string `json:"branch"  protobuf:"bytes,1,opt,name=branch"`
	//clone 路径
	DstDir   string `json:"dstDir"  protobuf:"bytes,1,opt,name=dstDir"`
	Username string `json:"username"  protobuf:"bytes,1,opt,name=username"`
	Password string `json:"password"  protobuf:"bytes,1,opt,name=password"`
}

type CMD struct {
	CommandName string   `json:"commandName"  protobuf:"bytes,1,opt,name=commandName"`
	Params      []string `json:"params"  protobuf:"bytes,1,opt,name=params"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo resources
type BuildConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []BuildConfig `json:"items"`
}