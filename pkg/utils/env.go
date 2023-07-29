package utils

import (
	"os"
	"path/filepath"
)

func GetBaseKubeScriptDir() string {
	dir := os.Getenv("KUBE_SCRIPT_BASE_DIR")
	if dir == "" {
		//return "/opt/kube/scripts/"
		return "/opt/cluster_engine/script"
	}
	return dir
}

func GetNodeInitDir() string {
	bdir := GetBaseKubeScriptDir()
	return filepath.Join(bdir, "host-init")
}

// func GetNodeInitDir() string {
// 	dir := os.Getenv("NODE_INIT_DIR")
// 	if dir == "" {
// 		return "/opt/kube/scripts/host-init"
// 	}
// 	return dir
// }

func GetBaseClusterAgentScriptDir() string {
	dir := os.Getenv("KUBE_SCRIPT_BASE_DIR")
	if dir == "" {
		//return "/opt/kube/scripts/"
		return "/opt/cluster_agent/script"
	}
	return dir
}

func GetNodeDeployDir() string {
	bdir := "/opt/cluster-engine/script"
	return filepath.Join(bdir, "hostMGR")
}
