package v1alpha1

import (
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Unitset is a specification for a Unitset resource
type Unitset struct {
	metaV1.TypeMeta   `json:",inline"`
	metaV1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec UnitsetSpec `json:"spec"`

	// +optional
	Status UnitsetStatus `json:"status"`
}

// UnitsetSpec is the spec for a Unitset resource
type UnitsetSpec struct {
	// +optional
	Image ImageVersion `json:"image"`

	// +optional
	Architecture Architecture `json:"architecture"`
	// shared config configmap name
	// 如果非空，先检查是否存在该cm，如果没有则报错
	// 则使用该configmap作为shared config
	// +optional
	SharedConfigName string `json:"sharedConfigName"`

	// +optional
	VolumeClaimTemplates []coreV1.PersistentVolumeClaim `json:"volumeClaimTemplates"`

	// +optional
	Action Action `json:"action,omitempty"`

	// +optional
	Template UnitTemplate `json:"template"`

	// +optional
	ExternalSecret ExternalSecretInfo `json:"externalSecret,omitempty"`
}

type UnitTemplate struct {

	// +optional
	Metadata UnitMetadata `json:"metadata"`
	// 是否与node绑定，默认false=绑定

	// +optional
	UnbindNode bool `json:"unbindNode,omitempty"`

	// +optional
	Env []coreV1.EnvVar `json:"env"`
	// 主容器资源配置

	// +optional
	Resource coreV1.ResourceRequirements `json:"resource,omitempty"`

	// +optional
	Volumes []coreV1.Volume `json:"volumes"`

	// +optional
	VolumeMounts []coreV1.VolumeMount `json:"volumeMounts"`

	// +optional
	Affinity *coreV1.Affinity `json:"affinity"`

	// +optional
	Ports []ContainerPort `json:"ports"`

	// default: true

	// +optional
	ShareProcessNamespace *bool `json:"shareProcessNamespace,omitempty"`

	// +optional
	ServiceAccount string `json:"serviceAccount"`
}

type UnitMetadata struct {

	// +optional
	Name string `json:"name"`

	// +optional
	Labels map[string]string `json:"labels"`

	// +optional
	Annotations map[string]string `json:"annotations"`
}

type Action struct {
}

type ExternalSecretInfo struct {

	// +optional
	Organization string `json:"organization"`

	// +optional
	RootSecret string `json:"root_secret"`
}

type ContainerPort struct {

	// +optional
	Port int32 `json:"port"`

	// +optional
	Name string `json:"name"`
}

// ImageVersion 镜像版本
type ImageVersion struct {
	// 镜像类型
	// required: true
	// example: infini-gateway
	Type string `json:"type"`
	// 主版本号
	// required: true
	// minimum: 0
	Major int `json:"major"`
	// 小版本号
	// required: true
	// minimum: 0
	Minor int `json:"minor"`
	// 小更新版本号
	// required: true
	// minimum: 0
	Patch int `json:"patch"`
	// 编译版本号
	// required: true
	// minimum: 0
	Dev int `json:"build"`
	// 架构
	// required: true
	// enum: arm64,amd64
	Arch string `json:"arch"`
}

type Architecture struct {

	// +optional
	Units int `json:"units"`

	// enum: single,clone,replication_async,replication_semi_sync
	// +optional
	Mode string `json:"mode"`
}

// UnitsetStatus is the status for a Unitset resource
type UnitsetStatus struct {

	// +optional
	Units int `json:"units"`

	// +optional
	ReadyUnits int `json:"readyUnits"`
}

type ConditionStatus string
type ConditionType string
type Condition struct {
	// +optional
	Type ConditionType `json:"type"`

	// +optional
	Status ConditionStatus `json:"status"`
}

type ErrMsg struct {

	// +optional
	Time metaV1.Time `json:"time"`

	// +optional
	Err string `json:"err"`

	// +optional
	Mode string `json:"mode"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UnitsetList is a list of Unitset resources
type UnitsetList struct {
	metaV1.TypeMeta `json:",inline"`
	metaV1.ListMeta `json:"metadata"`

	Items []Unitset `json:"items"`
}