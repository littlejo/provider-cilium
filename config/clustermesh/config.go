package clustermesh

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cilium_clustermesh", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "cilium"
		r.ShortGroup = "cmesh"
		r.Kind = "Enabler"
	})

	p.AddResourceConfigurator("cilium_clustermesh_connection", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "cilium"
		r.ShortGroup = "cmesh"
		r.Kind = "Connection"
	})
}
