package vars

const (
	CalicoNetworkMode  = "calico"
	MacvlanNetworkMode = "macvlan"
	SriovNetworkMode   = "sriov"
	RavenNetworkMode   = "raven"
	KubeovnNetworkMode = "kube-ovn"

	EndpointPodNameMode = "pod_name"
	EndpointPodIPMode   = "pod_ip"

	VolumePathStorageMode = "volumepath"
	PVCStorageMode        = "pvc"

	LoadBalancerNoneMode = "none"
	LoadBalancerAVIMode  = "avi-loadbalancer"

	ModeArchModeMasterDataCluster = "master_data_cluster"
	ModeArchModeData              = "data"
)
