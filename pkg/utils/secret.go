package utils

import (
	"encoding/json"
	"k8s.io/utils/mount"
	"strings"
)

func maskSecretKeys(inp interface{}) {
	if arr, ok := inp.([]interface{}); ok {
		for _, f := range arr {
			maskSecretKeys(f)
		}
		return
	}
	if form, ok := inp.(map[string]interface{}); ok {
	loop0:
		for k, v := range form {
			for _, m := range []string{"password", "secret", "jointoken", "unlockkey", "signingcakey"} {
				if strings.EqualFold(m, k) {
					form[k] = "*****"
					continue loop0
				}
			}
			maskSecretKeys(v)
		}
	}
}

func MaskJsonSecret(in []byte) ([]byte, error) {
	var form map[string]interface{}

	if err := json.Unmarshal(in, &form); err != nil {
		return nil, err
	}

	maskSecretKeys(form)

	return json.Marshal(form)
}

func IsMountPath(path string) bool {

	var (
		err       error
		mountList []mount.MountPoint
	)

	mounter := mount.New("")
	// Get list of mounted paths present with the node
	if mountList, err = mounter.List(); err != nil {
		return false
	}
	for _, mntInfo := range mountList {
		if mntInfo.Path == path {
			return true
		}
	}
	return false
}
