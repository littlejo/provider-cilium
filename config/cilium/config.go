package cilium

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cilium_deploy", func(r *config.Resource) {
		r.ShortGroup = "core"
		r.Kind = "CiliumDeployment"
	})
}
