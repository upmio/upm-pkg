package v1alpha1

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (unit *Unit) PodName() string {
	return unit.GetName()
}

func (unit *Unit) ServiceName() string {
	return unit.Name
}

func (unit *Unit) ScriptName() string {
	return unit.Name + "-script"
}

func (u *Unit) Valid() error {
	return nil
}

func volumeName(name string) string {
	return strings.Replace(name, ".", "-", -1)
}

func GetVolumePathName(unit *Unit, volume string) string {
	return GetPersistentVolumeName(unit, volume)
}

func GetLunGroupName(unit *Unit, volume string) string {
	return GetPersistentVolumeName(unit, volume)
}

func GetPersistentVolumeClaimName(unit *Unit, volume string) string {
	return GetPersistentVolumeName(unit, volume)
}

func GetPodName(unit *Unit) string {
	return unit.GetName()
}

func GetRetainVolumeName(unit *Unit, volume string) string {
	return fmt.Sprintf("%s-%d-%s", unit.GetName(), unit.Status.RebuildStatus.RetainVolumeSuffix, volume)
}

func GetPersistentVolumeName(unit *Unit, volume string) string {
	if unit.Status.RebuildStatus != nil && unit.Status.RebuildStatus.CurVolumeSuffix != 0 {
		return fmt.Sprintf("%s-%d-%s", unit.GetName(), unit.Status.RebuildStatus.CurVolumeSuffix, volume)
	}

	return fmt.Sprintf("%s-%s", unit.GetName(), volume)
}

func GetScriptTemplateName(unit *Unit) string {
	return unit.Spec.MainContainerName + "-" + unit.Spec.MainImageVersion + "-script"
}

func GetTemplateConfigName(unit *Unit) string {
	return unit.Spec.MainContainerName + "-" + unit.Spec.MainImageVersion + "-config-template"
}

func GetConfdToml(unit *Unit) string {
	return unit.Spec.MainContainerName + "-" + unit.Spec.MainImageVersion + "-confd-toml"
}

func GetConfdTmpl(unit *Unit) string {
	return unit.Spec.MainContainerName + "-" + unit.Spec.MainImageVersion + "-confd-tmpl"
}

func GetUnitConfdToml(unit *Unit) string {
	return unit.GetName() + "-confd-toml"
}

func GetUnitConfdTmpl(unit *Unit) string {
	return unit.GetName() + "-confd-tmpl"
}

func GetUnitScriptConfigName(unit *Unit) string {
	return unit.GetName() + "-script"
}

// func GetUnitConfigTemplateName(unit *Unit) string {
// 	return unit.GetName() + "-conf-tmpl-main"
// }
//
// func GetUnitConfigName(unit *Unit) string {
// 	return unit.GetName() + "-service-config"
// }

func GetUnitConfigTemplateName(unit *Unit) string {
	return unit.Spec.MainContainerName + "-" + unit.Spec.MainImageVersion + "-conf-tmpl-main"
}

func GetUnitConfigName(unit *Unit) string {
	return unit.GetName() + "-conf-tmpl-main"
}

func GetNetworkClaimName(unit *Unit) string {
	return unit.GetName()
}

func GetServiceName(unit *Unit) string {
	return unit.ServiceName()
}

// {"default/pod1":1,"default/pod2":2}
func GetUnitGroups(unit *Unit) (map[string]int, error) {
	group := map[string]int{}
	value, ok := unit.Annotations[PodGroupAnnotation]
	if !ok {
		return group, fmt.Errorf("not find [%s] annotation on unit", PodGroupAnnotation)
	}
	err := json.Unmarshal([]byte(value), &group)
	if err != nil {
		return group, fmt.Errorf("[%s:%s] Unmarshal fail:[%s] ", PodGroupAnnotation, value, err.Error())
	}
	return group, nil
}

// func GetUnitPort(unit *Unit, portName string) (int, bool) {
//	find := false
//	port := 0
//	for _, container := range unit.Spec.Template.Spec.Containers {
//		if container.Name == unit.Spec.MainContainerName {
//			for _, p := range container.Ports {
//				if p.Name == portName {
//					if p.ContainerPort != 0 {
//						port = int(p.ContainerPort)
//						find = true
//						break
//					}
//				}
//			}
//		}
//	}
//
//	if !find {
//		return 0, false
//	}
//
//	return port, true
// }

func GetUnitNameInGroups(podKey string) string {
	slices := strings.Split(podKey, "/")
	if len(slices) > 1 {
		return slices[1]
	}
	return slices[0]
}

// func GetVolumeMountPath(unit *Unit, volumeName string) string {
//	mountPath := ""
//	for _, one := range unit.Spec.VolumeClaims {
//		if one.Name == volumeName {
//			mountPath = one.Storage.Mounter
//			break
//		}
//	}
//
//	if mountPath == "" {
//		switch volumeName {
//		case "log":
//			mountPath = vars.DefaultLogMount
//		case "data":
//			mountPath = vars.DefaultDataMount
//		case "cache":
//			mountPath = vars.DefaultCacheMount
//		}
//	}
//
//	return mountPath
// }