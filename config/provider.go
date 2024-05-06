/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/littlejo/provider-cilium/config/cilium"
	"github.com/littlejo/provider-cilium/config/clustermesh"
	"github.com/littlejo/provider-cilium/config/config"
	"github.com/littlejo/provider-cilium/config/hubble"
	"github.com/littlejo/provider-cilium/config/kubeproxy"
)

const (
	resourcePrefix = "cilium"
	modulePath     = "github.com/littlejo/provider-cilium"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("cilium.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		hubble.Configure,
		cilium.Configure,
		config.Configure,
		clustermesh.Configure,
		kubeproxy.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
