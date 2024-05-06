// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConfigInitParameters struct {

	// (Boolean) Restart Cilium pods (Default: true).
	// Restart Cilium pods (Default: `true`).
	Restart *bool `json:"restart,omitempty" tf:"restart,omitempty"`

	// (String) Value of the key
	// Value of the key
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConfigObservation struct {

	// (String) Cilium config identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Key of the config
	// Key of the config
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (Boolean) Restart Cilium pods (Default: true).
	// Restart Cilium pods (Default: `true`).
	Restart *bool `json:"restart,omitempty" tf:"restart,omitempty"`

	// (String) Value of the key
	// Value of the key
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConfigParameters struct {

	// (String) Key of the config
	// Key of the config
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// (Boolean) Restart Cilium pods (Default: true).
	// Restart Cilium pods (Default: `true`).
	// +kubebuilder:validation:Optional
	Restart *bool `json:"restart,omitempty" tf:"restart,omitempty"`

	// (String) Value of the key
	// Value of the key
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// ConfigSpec defines the desired state of Config
type ConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ConfigInitParameters `json:"initProvider,omitempty"`
}

// ConfigStatus defines the observed state of Config.
type ConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Config is the Schema for the Configs API. Config resource for Cilium. This is equivalent to cilium cli: cilium config: It manages the cilium Kubernetes ConfigMap resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cilium}
type Config struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || (has(self.initProvider) && has(self.initProvider.value))",message="spec.forProvider.value is a required parameter"
	Spec   ConfigSpec   `json:"spec"`
	Status ConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigList contains a list of Configs
type ConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Config `json:"items"`
}

// Repository type metadata.
var (
	Config_Kind             = "Config"
	Config_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Config_Kind}.String()
	Config_KindAPIVersion   = Config_Kind + "." + CRDGroupVersion.String()
	Config_GroupVersionKind = CRDGroupVersion.WithKind(Config_Kind)
)

func init() {
	SchemeBuilder.Register(&Config{}, &ConfigList{})
}
