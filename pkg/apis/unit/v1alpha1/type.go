package v1alpha1

import (
	"time"

	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ScriptSyncTab                string = "sync-from"
	ScriptSyncResourceVersionTab string = "sync-resource-version"
	ScriptDataTab                string = "unitMGR"

	VGRequestAnnotation      string = "vg.localvolume.request"
	PodOptionsAnnotation     string = "options"
	PodGroupAnnotation       string = "groups"
	PodDependAnnotation      string = "depend"
	EndpointModeAnnotation   string = "endpoint.mode"
	ConfigSetsAnnotation     string = "config.sets"
	ExternalSecretAnnotation string = "external.secret"

	ConfigDataTab     string = "content"
	ConfigFilePathTab string = "filepath"

	ConfdTomlDataTab     string = "confd.toml"
	ConfdTemplateDataTab string = "confd.tmpl"

	ConfigSourceTypeConfigmap string = "configmap"
	ConfigSourceTypeTemplate  string = "template"

	StorageLocalType string = "bsgchina.com.storage.local"
)

type Phase string
type NetworkType string
type NetworkModeType string
type StorageModeType string

const (
	VolumePathAnnotation             = "volumepath"
	NetworkInternal      NetworkType = "bsgchina.networkClaim.internal"
	NetworkExternal      NetworkType = "bsgchina.networkClaim.external"

	RavenNetworkMode    NetworkModeType = "raven"
	CalicoNetworkMode   NetworkModeType = "calico"
	MacVlanNetworkMode  NetworkModeType = "macvlan"
	SriovNetworkMode    NetworkModeType = "sriov"
	ExternalNetworkMode NetworkModeType = "external"
	KubeOvnNetworkMode  NetworkModeType = "kube-ovn"

	VolumePathStorageMode StorageModeType = "volumepath"
	PVCStorageMode        StorageModeType = "pvc"
)

type ConditionStatus string
type ConditionType string
type Condition struct {
	Type   ConditionType
	Status ConditionStatus
}

const (
	ConditionUnset ConditionStatus = ""
	ConditionTrue  ConditionStatus = "True"
	ConditionFalse ConditionStatus = "False"

	InitStartCondition ConditionType = "initStart"
)

type UnitMetricsCfg struct {
	Metrics    []string `json:"metrics"`
	ExportPort int      `json:"export_port"`
}

type UnitMetricsStatus struct {
	MetricStatus []UnitMetricStatus `json:"metric_status"`
	TimeStamp    int64              `json:"timestamp"`
}

type UnitMetricStatus struct {
	Key     string            `json:"key"`
	Value   string            `json:"value"`
	Comment map[string]string `json:"comment"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Unit is a specification for a Unit resource
type Unit struct {
	metaV1.TypeMeta   `json:",inline"`
	metaV1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UnitSpec   `json:"spec"`
	Status UnitStatus `json:"status"`
}

// UnitSpec is the spec for a Unit resource
type UnitSpec struct {
	Secret            string `json:"secret"`
	MainContainerName string `json:"mainContainerName"`
	MainImageVersion  string `json:"mainImageVersion"`
	UnService         bool   `json:"unService"`
	UnBindNode        bool   `json:"unBindNode,omitempty"`

	Template     coreV1.PodTemplateSpec `json:"template"`
	ConfigSource []ConfigSource         `json:"config_source,omitempty"`
	Networking   NetworkingRequest      `json:"networking,omitempty"`
	VolumeClaims []PVCRequest           `json:"claims,omitempty"`
	Action       Action                 `json:"action"`
}

//	{
//	 "config_source": [
//	   {
//	     "type": "configmaps/templates",
//	     "namespaces": "",
//	     "configmap_name": "",
//	     "volumemount_mountpath": "",
//	     "volume_name": ""
//	   },
//	   {
//	     "type": "configmaps/templates",
//	     "namespaces": "",
//	     "configmap_name": "",
//	     "volumemount_mountpath": "",
//	     "volume_name": ""
//	   }
//	 ]
//	}
//
// ConfigSource
// 如果type = configmap,
// 那么在unit所在namespace下生成一个configmap,
// configmap_name="${UNIT_NAME}-service-config"
type ConfigSource struct {
	Type                 string `json:"type"`
	Namespace            string `json:"namespace"`
	ConfigmapName        string `json:"configmap_name"`
	DefaultMode          string `json:"default_mode"`
	VolumemountMountpath string `json:"volumemount_mountpath"`
	VolumeName           string `json:"volume_name"`
}

type MigrateAction struct {
	NodeName string `json:"nodeName,omitempty"`
	Force    bool   `json:"force,omitempty"`
}

type RebuildAction struct {
	OnlyPod  bool    `json:"only_pod,omitempty"`
	Start    bool    `json:"start,omitempty"`
	Force    bool    `json:"force,omitempty"`
	NodeName *string `json:"nodeName,omitempty"`

	RetainVolume *bool `json:"retain_volume,omitempty"`
}

type DeleteAction struct {
	Force   bool               `json:"force,omitempty"`
	PreStop *coreV1.ExecAction `json:"exec"`
}

type Action struct {
	Delete            *DeleteAction            `json:"delete,omitempty"`
	Rebuild           *RebuildAction           `json:"rebuild,omitempty"`
	Migrate           *MigrateAction           `json:"migrate,omitempty"`
	ReuseRetainVolume *ReuseRetainVolumeAction `json:"reuse_retain_volume,omitempty"`
}

type ReuseRetainVolumeAction struct {
	Force bool `json:"force,omitempty"`
}

type NetworkingRequest struct {
	Mode      NetworkModeType `json:"mode"`
	Bandwidth int32           `json:"bandwidth,omitempty"`
	Network   string          `json:"network,omitempty"`
	Type      NetworkType     `json:"type,omitempty"` // default NetworkInternal
}

type PVCRequest struct {
	Name             string          `json:"name"`
	StorageMode      StorageModeType `json:"storageMode"`
	StorageClassName string          `json:"storageClassName,omitempty"` // volumepath
	Storage          Storage         `json:"storage"`

	AccessModes []coreV1.PersistentVolumeAccessMode `json:"accessModes"`
}

type Storage struct {
	Type      string `json:"type,omitempty"`      // "infinilabs.com.storage.local" / "infinilabs.com.storage.remote"
	Level     string `json:"level,omitempty"`     // high-performance / normal-performance / low-performance
	AllocType string `json:"allocType,omitempty"` // "thick / thin
	FsType    string `json:"fsType,omitempty"`
	Mounter   string `json:"mounter,omitempty"`

	Request    resource.Quantity            `json:"request"`
	VolumeMode *coreV1.PersistentVolumeMode `json:"volumeMode,omitempty"`
}

func AddErrMsg(unit *Unit, msg ErrMsg) {
	if len(msg.Err) > 300 {
		msg.Err = msg.Err[:300]
	}

	if len(unit.Status.ErrMessages) == 0 {
		unit.Status.ErrMessages = []ErrMsg{msg}
		return
	}

	msgs := []ErrMsg{msg}
	for i, curMsg := range unit.Status.ErrMessages {
		if i > 3 {
			break
		}
		if curMsg.Time.Second()+30*60 > time.Now().Second() {
			msgs = append(msgs, curMsg)
		}
	}
	unit.Status.ErrMessages = msgs
}

func SetConditionStatus(unit *Unit, conditon Condition) {

	for i := range unit.Status.Conditions {
		if unit.Status.Conditions[i].Type == conditon.Type {
			unit.Status.Conditions[i].Status = conditon.Status
			return
		}
	}

	unit.Status.Conditions = append(unit.Status.Conditions, conditon)

}

func GetConditionStatus(unit *Unit, ty ConditionType) ConditionStatus {
	for _, condition := range unit.Status.Conditions {
		if condition.Type == ty {
			return condition.Status
		}
	}
	return ConditionUnset
}

type ErrMsg struct {
	Time metaV1.Time `json:"time"`
	Err  string      `json:"err"`
	Mode string      `json:"mode"`
}

type RebuildVolumeStatus struct {
	RetainVolumeSuffix int `json:"retain_volume_suffix"`
	CurVolumeSuffix    int `json:"cur_volume_suffix"`
}

// UnitStatus is the status for a Unit resource
type UnitStatus struct {
	RebuildStatus *RebuildVolumeStatus `json:"volume_suffix,omitempty"`
	Conditions    []Condition          `json:"conditions"`
	// Phase      string      `json:"phase"`
	ErrMessages []ErrMsg `json:"err_messages"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UnitList is a list of Unit resources
type UnitList struct {
	metaV1.TypeMeta `json:",inline"`
	metaV1.ListMeta `json:"metadata"`

	Items []Unit `json:"items"`
}