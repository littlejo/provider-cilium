/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"cilium_":                       config.TemplatedStringAsIdentifier("", "{{ .setup.configuration.helm_release }}"),
	"cilium_hubble":                 config.TemplatedStringAsIdentifier("", "cilium-hubble"),
	"cilium_config":                 config.TemplatedStringAsIdentifier("", "cilium-config-{{ .parameters.key }}"),
	"cilium_clustermesh":            config.TemplatedStringAsIdentifier("", "ciliumclustermeshenable"),
	"cilium_clustermesh_connection": config.TemplatedStringAsIdentifier("", "ciliumclustermeshconnect-{{ .parameters.destination_context }}"),
	"cilium_kubeproxy_free":         config.TemplatedStringAsIdentifier("", "cilium-kubeproxy-less"),
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
