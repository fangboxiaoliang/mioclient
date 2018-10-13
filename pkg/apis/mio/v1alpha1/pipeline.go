package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/hidevopsio/hiboot/pkg/system"
)

const (
	CODEPATH   string = "codepath"
	CLONE      string = "clone"
	COMPILE    string = "compile"
	BUILDIMAGE string = "buildImage"
	PUSHIMAGE  string = "pushImage"
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
	Status PipelineStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
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
	EventTypes        []string          `json:"event_type"`
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
	ForceUpdate    bool         `json:"force_update"`
	Replicas       int32        `json:"replicas"`
	Env            []system.Env `json:"env"`
	Labels         Labels       `json:"labels"`
	Project        string       `json:"project"`
}

type Labels struct {
	App     string `json:"app"`
	Version string `json:"version"`
	Cluster string `json:"cluster"`
}

type BuildConfigs struct {
	TagFrom     string       `json:"tag_from"`
	ImageStream string       `json:"image_stream"`
	Env         []system.Env `json:"env"`
	Rebuild     bool         `json:"rebuild"`
	Project     string       `json:"project"`
	Namespace   string       `json:"namespace"`
	Branch      string       `json:"branch"`
}

type GatewayConfigs struct {
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

type PipelineStatus struct {
	King           string           `json:"king" protobuf:"bytes,1,opt,name=kind"`
	Name           string           `json:"name" protobuf:"bytes,2,opt,name=name"`
	Namespace      string           `json:"namespace" protobuf:"bytes,3,opt,name=namespace"`
	Phase          string           `json:"phase"  protobuf:"bytes,4,opt,name=phase"`
	Stages         []PipelineStages `json:"stages" protobuf:"bytes,5,opt,name=stages"`
	StartTimestamp metav1.Time      `json:"startTimestamp" protobuf:"bytes,6,opt,name=startTimestamp"`
}

type PipelineStages struct {
	Name                 string `json:"name" protobuf:"bytes,1,opt,name=name"`
	StartTime            int64  `json:"startTime" protobuf:"bytes,2,opt,name=startTime"`
	DurationMilliseconds int64  `json:"durationMilliseconds" protobuf:"bytes,3,opt,name=durationMilliseconds"`
}
