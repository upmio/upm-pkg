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

	Spec   UnitsetSpec   `json:"spec"`
	Status UnitsetStatus `json:"status"`
}

// UnitsetSpec is the spec for a Unitset resource
type UnitsetSpec struct {
	// 软件的管理员用户、密码
	// 最终会挂在到容器中，挂载目录：/etc/secret-volume
	Secret string `json:"secret"`
	// required: false
	// shared config configmap name
	// 如果非空，先检查是否存在该cm，如果没有则报错
	// 则使用该configmap作为shared config
	SharedConfigName string `json:"shared_config_name"`
	// 是否与node绑定，默认false=绑定
	// required: false
	UnbindNode bool `json:"unbind_node,omitempty"`
	// 是否只做初始化,false:初始化并且启动; true:仅初始化,不启动
	// required: true
	InitOnly *bool `json:"init_only"`
	// 容器亲缘性，值为有亲缘性需求的service id
	// required: false
	Affinity Affinity `json:"affinity"`
	// required: false
	// enum: preferred, required
	// default: required
	PodAntiAffinity string `json:"pod_anti_affinity"`
	// required: true
	ZoneAffinity Affinity `json:"zone_affinity"`
	// required: false
	SourceAffinity AffinityNew `json:"source_affinity"`
	// required: true
	Arch Arch `json:"arch"`
	// required: true
	Image ImageVersion `json:"image"`
	// ImageRepositoryAddr is the address of image repository
	ImageRepositoryAddr string `json:"image_repository_addr"`
	// required: false
	ConfigSets []ConfigSet `json:"config_sets,omitempty"`
	// required: false
	// enum: service, pod_ip
	EndpointMode string `json:"endpoint_mode"`
	// required: false
	Ports            []ContainerPort      `json:"ports"`
	Options          map[string]string    `json:"options"`
	ResourceRequests ResourceRequirements `json:"resource_requests"`
	Env              []coreV1.EnvVar      `json:"env"`

	// required: false
	AuthSecret     string             `json:"auth_secret"`
	ExternalSecret ExternalSecretInfo `json:"external_secret"`
	// required: true
	Service K8sService `json:"service"`
	// required: false
	// default: true
	ShareProcessNamespace *bool `json:"share_process_namespace"`
}

type ExternalSecretInfo struct {
	Organization string `json:"organization"`
	RootSecret   string `json:"root_secret"`
}

type ResourceRequirements struct {
	// required: true
	// minimum: 1
	MiliCPU int64 `json:"milicpu"`
	// required: true
	// minimum: 1
	Memory  int64        `json:"memory"`
	Cache   *CacheInfo   `json:"cache,omitempty"`
	Storage *StorageInfo `json:"storage,omitempty"`
}

type CacheInfo struct {
	CacheType string `json:"type"`
}

type StorageInfo struct {
	Volumes        []VolumeRequirement `json:"volumes"`
	StorageClassID string              `json:"storageclass_id"`
}

type VolumeRequirement struct {
	Capacity  int64  `json:"capacity"`
	Type      string `json:"type"`
	MountPath string `json:"mount_path"`
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

type Arch struct {
	// required: true
	Nodes int `json:"nodes"`
	// 服务分片数量
	Shards int `json:"shards"`
	// required: true
	// enum: single,clone,replication_async,replication_semi_sync
	Mode string `json:"mode"`
	Desc string `json:"desc"`
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