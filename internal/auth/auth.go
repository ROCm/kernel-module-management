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

package auth

import (
	"context"
	"fmt"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/authn/kubernetes"
	"github.com/kubernetes-sigs/kernel-module-management/internal/api"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//go:generate mockgen -source=auth.go -package=auth -destination=mock_auth.go

type RegistryAuthGetter interface {
	GetKeyChain(ctx context.Context) (authn.Keychain, error)
}

type registrySecretAuthGetter struct {
	client         client.Client
	namespacedName types.NamespacedName
}

func NewRegistryAuthGetter(client client.Client, namespacedName types.NamespacedName) RegistryAuthGetter {
	return &registrySecretAuthGetter{
		client:         client,
		namespacedName: namespacedName,
	}
}

func (rsag *registrySecretAuthGetter) GetKeyChain(ctx context.Context) (authn.Keychain, error) {

	secret := v1.Secret{}
	if err := rsag.client.Get(ctx, rsag.namespacedName, &secret); err != nil {
		return nil, fmt.Errorf("cannot find secret %s: %w", rsag.namespacedName, err)
	}

	keychain, err := kubernetes.NewFromPullSecrets(ctx, []v1.Secret{secret})
	if err != nil {
		return nil, fmt.Errorf("could not create a keycahin from secret %v: %w", secret, err)
	}

	return keychain, nil
}

func NewRegistryAuthGetterFrom(client client.Client, mld *api.ModuleLoaderData) RegistryAuthGetter {
	if mld.ImageRepoSecret != nil {
		namespacedName := types.NamespacedName{
			Name:      mld.ImageRepoSecret.Name,
			Namespace: mld.Namespace,
		}
		return NewRegistryAuthGetter(client, namespacedName)
	}
	return nil
}
