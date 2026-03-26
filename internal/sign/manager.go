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

package sign

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubernetes-sigs/kernel-module-management/internal/api"
	"github.com/kubernetes-sigs/kernel-module-management/internal/utils"
)

//go:generate mockgen -source=manager.go -package=sign -destination=mock_manager.go

type SignManager interface {
	GarbageCollect(ctx context.Context, modName, namespace string, owner metav1.Object) ([]string, error)

	ShouldSync(ctx context.Context, mld *api.ModuleLoaderData) (bool, error)

	Sync(
		ctx context.Context,
		mld *api.ModuleLoaderData,
		imageToSign string,
		pushImage bool,
		owner metav1.Object) (utils.Status, error)
}
