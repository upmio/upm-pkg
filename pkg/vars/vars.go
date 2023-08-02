package vars

import (
	"k8s.io/klog/v2"
	"os"
)

var (
	// GITCOMMIT will be overwritten automatically by the build system
	GITCOMMIT = "HEAD"

	// BUILDTIME will be overwritten automatically by the build system
	BUILDTIME = "<unknown>"

	SeCretAESKey       = "!QAw2f#dfwef234df345dfsddd3e1"
	CBC256SeCretAESKey = "U2FsdGVkX18wJo2/q5Utz7T30qXCEbIiYVApKmaOTsU="
	AescbcwithivKey    = "7e20c20ea7564231a76dd83ac1cf7013"

	ManagerNamespace = "unknown"
	// EngineImage      = "unknown"
	ProjectName = "upm"
	IpFamily    = "IPv4"
)

func init() {
	managerNamespace := os.Getenv("NAMESPACE")
	if managerNamespace == "" {
		klog.Fatalf("not found env: [NAMESPACE], can't start service...")
	} else {
		ManagerNamespace = managerNamespace
	}

	ipFamily := os.Getenv("IP_FAMILY")
	if ipFamily == "" {
		klog.Infof("not found env: [IP_FAMILY], only support [SingleStack:IPv4]...")
	} else {
		klog.Infof("found env: [IP_FAMILY], only support [%s]...", ipFamily)
		IpFamily = ipFamily
	}
}

var (
	ServiceAccountSuffix     = ProjectName + "-sa"
	ClusterRoleSuffix        = ProjectName + "-clusterrole"
	ClusterRoleBindingSuffix = ProjectName + "-clusterrolebinding"

	ClusterManagerName = ProjectName + "-manager"
	ClusterEngineName  = ProjectName + "-engine"

	UnitAgentName     = "unit-agent"
	UnitAgentImage    string
	UnitAgentHostType = "domain"
)

const (
	InternalSchedulerName = "volumepath-scheduler"

	Certmanager                  = "cert-manager"
	CertmanagerIssuerSuffix      = "certmanager-issuer"
	CertmanagerCertificateSuffix = "certmanager-ca"
	CertmanagerSecretNameSuffix  = "secret"

	PodAffinityPreferredType = "preferred"
	PodAffinityRequiredType  = "required"
	LicenseSecretNamePrefix  = "license"
	CSIVomepathDriverName    = "infini.volumepath.csi"

	SharedConfigNameSuffix                    = "-shared-config"
	SharedConfigDataPluginsKey                = "plugins"
	SharedConfigDataPluginDictListKeySuffix   = "_dict_list"
	SharedConfigDataPluginDictConfigKeySuffix = "_dict_config"

	LogstashPipelineConfigNameSuffix = "-logstash-pipeline"
	LogstashPipelineConfigKey        = "pipelines.yml"
	LogstashPipelineDefaultID        = "default"
	LogstashPipelineBaseConfigPath   = "/DAT_DIR/pipelines/conf"

	// SharedConfigDiscoverySeedHostsKey
	// value格式: "["c5ffc20b-h2v-esmaster-0","c5ffc20b-h2v-esmaster-1"]" es master的域名+port
	SharedConfigDiscoverySeedHostsKey = "discovery_seed_hosts"

	// SharedConfigHttpsMasterHosts
	// value格式: "["https://c5ffc20b-h2v-esmaster-0:3306","https://c5ffc20b-h2v-esmaster-1:3306"]" es master的域名+port
	// 用于apm 的elasticsearch.hosts
	// 用于kibana 的defaults.elasticsearch_hosts
	SharedConfigHttpsMasterHosts = "https_master_hosts"

	// SharedConfigMasterHosts
	// value: "["c5ffc20b-h2v-esmaster-0:3306","c5ffc20b-h2v-esmaster-1:3306"]" es master的域名+port
	// 用于gateway 的elasticsearch.hosts
	SharedConfigMasterHosts = "master_hosts"

	// SharedConfigHttpKibanaHost
	// value格式: "http://c5ffc20b-h2v-esmaster-0:3306" kibana中一个+port
	// 用于apm 的apm_server.kibana_host
	SharedConfigHttpKibanaHost = "kibana_host"

	ServiceMonitorCrdName           = "servicemonitors.monitoring.coreos.com"
	MonitorSidecarName              = "monitor"
	MonitorServiceNameSuffix        = "-exporter-svc"
	MonitorServiceMonitorNameSuffix = "-exporter-svcmon"

	DefaultImagePullSecret = "regcred"
	OptionArch             = "arch"

	DefaultCacheVolumeMountName = "cache-volume"
	DefaultCacheMount           = "/CACHE"
	DefaultDataMount            = "/DAT_DIR"
	DefaultLogMount             = "/LOG_DIR"

	ClusterAgent  = "cluster-agent"
	ClusterEngine = "cluster-engine"

	HttpSchemeMode  = "http"
	HttpsSchemeMode = "https"

	TanzuSiteMode     = "tanzu"
	KubeSpraySiteMode = "kubespray"

	FileDateSource      = "file"
	ConfigMapDateSource = "configmap"
	ConfigSiteJsonTab   = "site.json"

	NodeRoleSpare = "spare"
	NodeRoleNode  = "node"
	JobName       = "job-name"

	PortNameHttp           = "http"
	PortNameAPMServer      = "apm-server"
	PortNameAPIPort        = "api"
	PortNameEntryPort      = "entry"
	PortNameKafka          = "kafka"
	PortNameTransportPort  = "transport"
	PortNameLeadershipPort = "leadership"
	PortNameJmx            = "jmx"
	PortNameClient         = "client"
	PortNameApiHttp        = "http"

	EnvNameUnitName = "UNIT_NAME"
)

const (
	EnvPipLineConfDir = "PIPELINE_CONF_DIR"
)