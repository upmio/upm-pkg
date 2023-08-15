package v1alpha1

import (
	unitv1alpha1 "github.com/upmio/upm-pkg/pkg/apis/unit/v1alpha1"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Unitset is a specification for a Unitset resource
type Unitset struct {
	metaV1.TypeMeta   `json:",inline"`
	metaV1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UnitsetSpec   `json:"spec"`
	Status UnitsetStatus `json:"status"`
}

// UnitsetSpec is the spec for a Unitset resource
type UnitsetSpec struct {
	Action Action `json:"action"`
	// required: false
	// shared config configmap name
	// 如果非空，先检查是否存在该cm，如果没有则报错
	// 则使用该configmap作为shared config
	SharedConfigName string `json:"sharedConfigName"`
	// 是否与node绑定，默认false=绑定
	// required: false
	UnbindNode bool `json:"unbindNode,omitempty"`
	// 原InitOnly: true: 服务启动；false: 不用启动
	Startup  *bool            `json:"startup"`
	Affinity *coreV1.Affinity `json:"affinity"`
	// required: true
	Architecture Architecture `json:"architecture"`
	// required: true
	Image ImageVersion `json:"image"`
	// required: false
	Ports          []ContainerPort    `json:"ports"`
	Env            []coreV1.EnvVar    `json:"env"`
	ExternalSecret ExternalSecretInfo `json:"externalSecret"`
	Options        map[string]string  `json:"options"`
	// required: true
	Service K8sService `json:"service"`
	// required: false
	// default: true
	ShareProcessNamespace *bool `json:"share_process_namespace"`
	// 主容器资源配置
	Resource             coreV1.ResourceRequirements    `json:"resource"`
	Volumes              []coreV1.Volume                `json:"volumes"`
	VolumeClaimTemplates []coreV1.PersistentVolumeClaim `json:"volumeClaimTemplates"`
	VolumeMounts         []coreV1.VolumeMount           `json:"volumeMounts"`
}

type Action struct {
	Delete *unitv1alpha1.DeleteAction `json:"delete,omitempty"`
}

type DeleteAction struct {
	Force   bool               `json:"force,omitempty"`
	PreStop *coreV1.ExecAction `json:"exec"`
}

type ExternalSecretInfo struct {
	Organization string `json:"organization"`
	RootSecret   string `json:"root_secret"`
}

type ResourceRequirements struct {
	// required: true
	// minimum: 1
	MiliCPU resource.Quantity `json:"milicpu"`
	// required: true
	// minimum: 1
	Memory  resource.Quantity `json:"memory"`
	Cache   *CacheInfo        `json:"cache,omitempty"`
	Storage *StorageInfo      `json:"storage,omitempty"`
}

type CacheInfo struct {
	CacheType string `json:"type"`
}

type StorageInfo struct {
	Volumes        []VolumeRequirement `json:"volumes"`
	StorageClassID string              `json:"storageclass_id"`
}

type VolumeRequirement struct {
	Capacity  resource.Quantity `json:"capacity"`
	Type      string            `json:"type"`
	MountPath string            `json:"mount_path"`
}

type ContainerPort struct {
	Port int32  `json:"port"`
	Name string `json:"name"`
}

type K8sService struct {
	// enum: ClusterIP, NodePort, LoadBalancer
	Type                     string `json:"type"`
	PublishNotReadyAddresses *bool  `json:"publish_not_ready_addresses"`
}

// ConfigSet 服务配置信息(创建时用户指定的）
type ConfigSet struct {
	Section string `json:"section"`
	Key     string `json:"key"`
	Value   string `json:"value"`
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
	// required: true
	Nodes int `json:"nodes"`
	// required: true
	// enum: single,clone,replication_async,replication_semi_sync
	Mode string `json:"mode"`
}

// Affinity ...
type Affinity struct {
	Required  []string `json:"required"`
	Preferred []string `json:"preferred"`
}

// AffinityNew ...
// [key] = [value1, value2, ...]
type AffinityNew struct {
	Required  map[string][]string `json:"required"`
	Preferred map[string][]string `json:"preferred"`
}

// UnitsetStatus is the status for a Unitset resource
type UnitsetStatus struct {
	ErrMessages []ErrMsg    `json:"err_messages"`
	Conditions  []Condition `json:"conditions"`
}

type ConditionStatus string
type ConditionType string
type Condition struct {
	Type   ConditionType
	Status ConditionStatus
}

type ErrMsg struct {
	Time metaV1.Time `json:"time"`
	Err  string      `json:"err"`
	Mode string      `json:"mode"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UnitsetList is a list of Unitset resources
type UnitsetList struct {
	metaV1.TypeMeta `json:",inline"`
	metaV1.ListMeta `json:"metadata"`

	Items []Unitset `json:"items"`
}