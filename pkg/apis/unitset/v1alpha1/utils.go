package v1alpha1

import (
	"golang.org/x/xerrors"
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

	image := u.Spec.Image.ID

	if image == "" {
		errs = append(errs, xerrors.New("image is required"))
	}

	if u.Spec.Unit.Resources.Requests.Storage != nil {
		if u.Spec.Unit.Resources.Requests.Storage.StorageClass.ID != "" {
			if len(u.Spec.Unit.Resources.Requests.Storage.Volumes) == 0 {
				errs = append(errs, xerrors.New("volumes is required"))
			} else {
				for i := range u.Spec.Unit.Resources.Requests.Storage.Volumes {
					if u.Spec.Unit.Resources.Requests.Storage.Volumes[i].Capacity <= int64(100) {
						errs = append(errs, xerrors.New("storage volumes capacity do not allow less than or equal to 100"))
					}
				}
			}
		}

		if len(u.Spec.Unit.Resources.Requests.Storage.Volumes) != 0 {
			if u.Spec.Unit.Resources.Requests.Storage.StorageClass.ID == "" {
				errs = append(errs, xerrors.New("has data or log volume, storageclass id is required"))
			}
		}
	}

	if u.Spec.Unit.Resources.Requests.Storage == nil && u.Spec.Unit.Resources.Requests.Cache == nil {
		errs = append(errs, xerrors.New("at least one cache volume or data、log volume"))
	}

	if u.Spec.Unit.Resources.Requests.MiliCPU <= 0 {
		errs = append(errs, xerrors.New("unit resource milicpu is required"))
	}

	if u.Spec.Unit.Resources.Requests.Memory <= 0 {
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

	return nil
}