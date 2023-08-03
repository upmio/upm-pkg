package v1alpha1

import (
	"fmt"
	"golang.org/x/xerrors"
	utilErrors "k8s.io/apimachinery/pkg/util/errors"
)

func (u *Unitset) Valid() error {
	errs := []error{}

	if u.Spec.InitOnly == nil {
		errs = append(errs, xerrors.New("spec.init_only is required"))
	}

	if u.Spec.Arch.Mode == "" {
		errs = append(errs, xerrors.New("arch mode is required"))
	}

	if u.Spec.Arch.Nodes <= 0 {
		errs = append(errs, xerrors.New("spec.arch.nodes is required"))
	}

	if (u.Spec.ZoneAffinity.Preferred == nil && len(u.Spec.ZoneAffinity.Preferred) == 0) &&
		(u.Spec.ZoneAffinity.Required == nil && len(u.Spec.ZoneAffinity.Required) == 0) {
		errs = append(errs, xerrors.New("ZoneAffinity.Preferred or ZoneAffinity.Required is required"))
	}

	if u.Spec.Image.Type == "" {
		errs = append(errs, xerrors.New("image is required"))
	}

	if u.Spec.ResourceRequests.Storage != nil {
		if u.Spec.ResourceRequests.Storage.StorageClassID != "" {
			if len(u.Spec.ResourceRequests.Storage.Volumes) == 0 {
				errs = append(errs, xerrors.New("volumes is required"))
			} else {
				for i := range u.Spec.ResourceRequests.Storage.Volumes {
					if u.Spec.ResourceRequests.Storage.Volumes[i].Capacity <= int64(100) {
						errs = append(errs, xerrors.New("storage volumes capacity do not allow less than or equal to 100"))
					}
				}
			}
		}

		if len(u.Spec.ResourceRequests.Storage.Volumes) != 0 {
			if u.Spec.ResourceRequests.Storage.StorageClassID == "" {
				errs = append(errs, xerrors.New("has data or log volume, storageclass id is required"))
			}
		}
	}

	if u.Spec.ResourceRequests.Storage == nil && u.Spec.ResourceRequests.Cache == nil {
		errs = append(errs, xerrors.New("at least one cache volume or data、log volume"))
	}

	if u.Spec.ResourceRequests.MiliCPU <= 0 {
		errs = append(errs, xerrors.New("unit resource milicpu is required"))
	}

	if u.Spec.ResourceRequests.Memory <= 0 {
		errs = append(errs, xerrors.New("unit resource memory is required"))
	}

	if u.Spec.Service.Type == "" {
		errs = append(errs, xerrors.New("k8s service type is required, support:['none','ClusterIP','NodePort','LoadBalancer']"))
	} else {
		if u.Spec.Service.Type != "none" && u.Spec.Service.Type != "ClusterIP" && u.Spec.Service.Type != "NodePort" && u.Spec.Service.Type != "LoadBalancer" {
			errs = append(errs, xerrors.Errorf("not support the ingress methods:[%s] for a Kubernetes service,"+
				"only support:['none','ClusterIP','NodePort','LoadBalancer']", u.Spec.Service.Type))
		}
	}

	return utilErrors.NewAggregate(errs)
}

// GetUnitsetPodTemplateName
// e.g.: mysql-8.0.34.1-amd64-pod-tmpl
func GetUnitsetPodTemplateName(unitset Unitset) string {
	return unitset.Spec.Image.TypeWithVersionAndArchByHyphen() + "-pod-tmpl"
}

// GetUnitImageWithImageRepositoryAddr
// e.g.: repositoryAddr/mysql:8.0.34.1-amd64
func GetUnitImageWithImageRepositoryAddr(unitset Unitset) string {
	return fmt.Sprintf("%s/%s", unitset.Spec.ImageRepositoryAddr, unitset.Spec.Image.TypeWithVersionAndArchByColon())
}

// Version
// e.g.: 5.7.0.1
func (iv *ImageVersion) Version() string {
	return fmt.Sprintf("%d.%d.%d.%d", iv.Major, iv.Minor, iv.Patch, iv.Dev)
}

// VersionWithArch
// e.g.: 5.7.0.1-amd64
func (iv *ImageVersion) VersionWithArch() string {
	return fmt.Sprintf("%d.%d.%d.%d-%s", iv.Major, iv.Minor, iv.Patch, iv.Dev, iv.Arch)
}

// MainVersion
// e.g.: 5.7
func (iv *ImageVersion) MainVersion() string {
	return fmt.Sprintf("%d.%d", iv.Major, iv.Minor)
}

// TypeWithVersionByColon
// e.g.: mysql:5.7.0.1
func (iv *ImageVersion) TypeWithVersionByColon() string {
	return fmt.Sprintf("%s:%d.%d.%d.%d", iv.Type, iv.Major, iv.Minor, iv.Patch, iv.Dev)
}

// TypeWithVersionByHyphen
// e.g.: mysql-5.7.0.1
func (iv *ImageVersion) TypeWithVersionByHyphen() string {
	return fmt.Sprintf("%s-%d.%d.%d.%d", iv.Type, iv.Major, iv.Minor, iv.Patch, iv.Dev)
}

// TypeWithVersionAndArchByColon
// e.g.: mysql:5.7.0.1-amd64
func (iv *ImageVersion) TypeWithVersionAndArchByColon() string {
	return fmt.Sprintf("%s:%d.%d.%d.%d-%s", iv.Type, iv.Major, iv.Minor, iv.Patch, iv.Dev, iv.Arch)
}

// TypeWithVersionAndArchByHyphen
// e.g.: mysql-5.7.0.1-amd64
func (iv *ImageVersion) TypeWithVersionAndArchByHyphen() string {
	return fmt.Sprintf("%s-%d.%d.%d.%d-%s", iv.Type, iv.Major, iv.Minor, iv.Patch, iv.Dev, iv.Arch)
}