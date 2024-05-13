// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	ciliumclustermeshconnection "github.com/littlejo/provider-cilium/internal/controller/cmesh/ciliumclustermeshconnection"
	ciliumclustermeshenabler "github.com/littlejo/provider-cilium/internal/controller/cmesh/ciliumclustermeshenabler"
	ciliumdeployment "github.com/littlejo/provider-cilium/internal/controller/core/ciliumdeployment"
	ciliumconfig "github.com/littlejo/provider-cilium/internal/controller/opt/ciliumconfig"
	hubble "github.com/littlejo/provider-cilium/internal/controller/opt/hubble"
	kubeproxyfree "github.com/littlejo/provider-cilium/internal/controller/opt/kubeproxyfree"
	providerconfig "github.com/littlejo/provider-cilium/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		ciliumclustermeshconnection.Setup,
		ciliumclustermeshenabler.Setup,
		ciliumdeployment.Setup,
		ciliumconfig.Setup,
		hubble.Setup,
		kubeproxyfree.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
