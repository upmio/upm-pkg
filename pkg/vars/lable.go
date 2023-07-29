package vars

var (
	LabelKey                      = ProjectName
	LabelValue                    = ProjectName
	LabelIo                       = ProjectName + ".io"
	LabelCluster                  = ProjectName + ".cluster"
	LabelZone                     = ProjectName + ".zone"
	LabelUnitName                 = ProjectName + ".unit.name"
	LabelSite                     = ProjectName + ".site"
	LabelResourceLimit            = ProjectName + ".soft-resource.limit"
	LabelServiceID                = ProjectName + ".service.id"
	LabelServiceName              = ProjectName + ".service.name"
	LabelServiceImageType         = ProjectName + ".service.image.name"
	LabelGroupName                = ProjectName + ".service-group"
	LabelGroupType                = ProjectName + ".service-group.type"
	LabelHostID                   = ProjectName + ".host.id"
	LabelStorageDriveType         = ProjectName + ".storage.drive.type"
	LabelCPUModelName             = ProjectName + ".cpu.model.name"
	LabelHostRoom                 = ProjectName + ".host.room"
	LabelHostSeat                 = ProjectName + ".host.seat"
	LabelArchMode                 = ProjectName + ".arch.mode"
	LabelNodeLocalVolumeTmpl      = ProjectName + ".localvolume.%s"
	LabelNodeLocalVolumeLimitTmpl = ProjectName + ".localvolume.resource.limit.%s"
	LabelIoRole                   = ProjectName + ".io/role"
	LabelSecretType               = ProjectName + ".secret.type"

	UnitMetricsAnnotation       = ProjectName + ".unit.metrics"
	UnitMetricsStatusAnnotation = ProjectName + ".unit.metrics.status"
)

const (
	LabelCASecret             = "ca"
	LabelUnschedulable        = "unschedulable"
	LabelArch                 = "kubernetes.io/arch"
	LabelOS                   = "kubernetes.io/os"
	LabelImageFullName        = "image"      // infini-gateway-1.7.0.720-amd64
	LabelImageVersionWithArch = "image.tag"  // 1.7.0.720-amd64
	LabelImageType            = "image.type" // infini-gateway
	PreparedNodeRoleLabel     = "prepared"

	AnnotationHostUsageLimit  = "node.usage.limit"
	AnnotationHostStorageInfo = "host.storage.info"

	LastUnitBelongNodeAnnotation = "last.unit.belong.node"
	CalicoPodIPAnnotation        = "cni.projectcalico.org/podIP"
	CalicoPodIPsAnnotation       = "cni.projectcalico.org/podIPs"
	CalicoIpAddrsAnnotation      = "cni.projectcalico.org/ipAddrs"
	KubeOvnPodIPAnnotation       = "ovn.kubernetes.io/ip_address"
)