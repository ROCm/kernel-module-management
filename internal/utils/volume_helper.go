/*
Copyright (c) Advanced Micro Devices, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the \"License\");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an \"AS IS\" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	v1 "k8s.io/api/core/v1"
)

func MakeSecretVolume(secretRef *v1.LocalObjectReference, key string, path string) v1.Volume {
	if secretRef == nil {
		return v1.Volume{}
	}

	vol := v1.Volume{
		Name: volumeNameFromSecretRef(*secretRef),
		VolumeSource: v1.VolumeSource{
			Secret: &v1.SecretVolumeSource{
				SecretName: secretRef.Name,
			},
		},
	}

	if key != "" {
		vol.VolumeSource.Secret.Items = []v1.KeyToPath{
			{
				Key:  key,
				Path: path,
			},
		}
	}

	return vol
}

func MakeSecretVolumeMount(secretRef *v1.LocalObjectReference, mountPath string, readOnly bool) v1.VolumeMount {
	if secretRef == nil {
		return v1.VolumeMount{}
	}

	return v1.VolumeMount{
		Name:      volumeNameFromSecretRef(*secretRef),
		ReadOnly:  readOnly,
		MountPath: mountPath,
	}
}

func volumeNameFromSecretRef(ref v1.LocalObjectReference) string {
	return "secret-" + ref.Name
}
