package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/hidevopsio/hiboot/pkg/system"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo is a specification for a Foo resource
type Pipeline struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of the Deployment.
	// +optional
	Spec PipelineSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Most recently observed status of the Deployment.
	// +optional
	Status metav1.Status `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type PipelineSpec struct {
	Name              string            `json:"name" validate:"required"`
	App               string            `json:"app" validate:"required"`
	Profile           string            `json:"profile"`
	Project           string            `json:"project" validate:"required"`
	Cluster           string            `json:"cluster"`
	Namespace         string            `json:"namespace"`
	Scm               Scm               `json:"scm"`
	Version           string            `json:"version"`
	DockerRegistry    string            `json:"docker_registry"`
	NodeSelector      string            `json:"node_selector"`
	Identifiers       []string          `json:"identifiers"`
	ConfigFiles       []string          `json:"config_files"`
	Ports             []Ports           `json:"ports"`
	BuildConfigs      BuildConfigs      `json:"build_configs"`
	DeploymentConfigs DeploymentConfigs `json:"deployment_configs"`
	GatewayConfigs    GatewayConfigs    `json:"gateway_configs"`
}

type Ports struct {
	Name          string `json:"name"`
	Port          int32  `json:"port" `
	ContainerPort int32  `json:"container_port"`
	Protocol      string `json:"protocol,omitempty"`
}

type Scm struct {
	Type     string `json:"type"`
	Url      string `json:"url"`
	Ref      string `json:"ref"`
	UserName string `json:"user_name"`
}

type DeploymentConfigs struct {
	HealthEndPoint string       `json:"health_end_point"`
	Enable         bool         `json:"enable"`
	ForceUpdate    bool         `json:"force_update"`
	Replicas       int32        `json:"replicas"`
	Env            []system.Env `json:"env"`
	Labels         Labels       `json:"labels"`
	Project        string       `json:"project"`
	RemoteEnable   bool         `json:"remote_enable"`
}

type Labels struct {
	App     string `json:"app"`
	Version string `json:"version"`
	Cluster string `json:"cluster"`
}

type BuildConfigs struct {
	TagEnable   bool         `json:"tag_enable"`
	Enable      bool         `json:"enable"`
	TagFrom     string       `json:"tag_from"`
	ImageStream string       `json:"image_stream"`
	Env         []system.Env `json:"env"`
	Rebuild     bool         `json:"rebuild"`
	Project     string       `json:"project"`
	Namespace   string       `json:"namespace"`
	Branch      string       `json:"branch"`
}

type GatewayConfigs struct {
	Enable      bool   `json:"enable"`
	Uri         string `json:"uri"`
	UpstreamUrl string `json:"upstream_url"`
}


// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo resources
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Pipeline `json:"items"`
}
