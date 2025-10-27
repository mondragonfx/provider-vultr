// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	kubernetes "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/kubernetes/kubernetes"
	nodepools "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/kubernetes/nodepools"
	providerconfig "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/providerconfig"
	instance "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/vultr/instance"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		kubernetes.Setup,
		nodepools.Setup,
		providerconfig.Setup,
		instance.Setup,
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
		kubernetes.SetupGated,
		nodepools.SetupGated,
		providerconfig.SetupGated,
		instance.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
