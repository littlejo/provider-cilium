package clustermesh

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cilium_clustermesh", func(r *config.Resource) {
		r.ShortGroup = "cmesh"
		r.Kind = "CiliumClusterMeshEnabler"
	})

	p.AddResourceConfigurator("cilium_clustermesh_connection", func(r *config.Resource) {
		r.ShortGroup = "cmesh"
		r.Kind = "CiliumClusterMeshConnection"
	})
}
