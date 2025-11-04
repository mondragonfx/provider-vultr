// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	metalserver "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/baremetal/metalserver"
	storage "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/block/storage"
	connectionpool "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/database/connectionpool"
	database "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/database/database"
	db "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/database/db"
	replica "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/database/replica"
	user "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/database/user"
	nodepools "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/kubernetes/nodepools"
	balancer "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/loadbalancer/balancer"
	storageobject "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/object/storage"
	providerconfig "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/providerconfig"
	instance "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/vultr/instance"
	kubernetes "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/vultr/kubernetes"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		metalserver.Setup,
		storage.Setup,
		connectionpool.Setup,
		database.Setup,
		db.Setup,
		replica.Setup,
		user.Setup,
		nodepools.Setup,
		balancer.Setup,
		storageobject.Setup,
		providerconfig.Setup,
		instance.Setup,
		kubernetes.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		metalserver.SetupGated,
		storage.SetupGated,
		connectionpool.SetupGated,
		database.SetupGated,
		db.SetupGated,
		replica.SetupGated,
		user.SetupGated,
		nodepools.SetupGated,
		balancer.SetupGated,
		storageobject.SetupGated,
		providerconfig.SetupGated,
		instance.SetupGated,
		kubernetes.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
