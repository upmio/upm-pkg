package v1alpha1

import (
	"fmt"
	utilErrors "k8s.io/apimachinery/pkg/util/errors"
)

func (u *Unitset) Valid() error {
	errs := []error{}

	return utilErrors.NewAggregate(errs)
}

// GetUnitsetPodTemplateName
// e.g.: mysql-8.0.34.1-amd64-pod-tmpl
func GetUnitsetPodTemplateName(unitset Unitset) string {
	return unitset.Spec.Image.TypeWithVersionAndArchByHyphen() + "-pod-tmpl"
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