package v1alpha1

import (
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
	Secret            string `json:"secret"`
	MainContainerName string `json:"mainContainerName"`
	MainImageVersion  string `json:"mainImageVersion"`
	UnService         bool   `json:"unService"`
	UnBindNode        bool   `json:"unBindNode,omitempty"`
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