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

type CiliumDeploymentInitParameters struct {

	// Datapath mode to use { tunnel | native | aws-eni | gke | azure | aks-byocni } (Default: `autodetected`).
	DataPath *string `json:"dataPath,omitempty" tf:"data_path,omitempty"`

	// Helm chart repository to download Cilium charts from (Default: `https://helm.cilium.io`).
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// When upgrading, reset the helm values to the ones built into the chart (Default: `false`).
	Reset *bool `json:"reset,omitempty" tf:"reset,omitempty"`

	// When upgrading, reuse the helm values from the latest release unless any overrides from are set from other flags. This option takes precedence over HelmResetValues (Default: `true`).
	Reuse *bool `json:"reuse,omitempty" tf:"reuse,omitempty"`

	// Set helm values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2 (Default: `[]`).
	Set []*string `json:"set,omitempty" tf:"set,omitempty"`

	// values in raw yaml to pass to helm. (Default: `empty`).
	Values *string `json:"values,omitempty" tf:"values,omitempty"`

	// Version of Cilium (Default: `v1.15.4`).
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Wait for Cilium status is ok (Default: `true`).
	Wait *bool `json:"wait,omitempty" tf:"wait,omitempty"`
}

type CiliumDeploymentObservation struct {

	// Datapath mode to use { tunnel | native | aws-eni | gke | azure | aks-byocni } (Default: `autodetected`).
	DataPath *string `json:"dataPath,omitempty" tf:"data_path,omitempty"`

	// Helm values (`helm get values -n kube-system cilium`)
	HelmValues *string `json:"helmValues,omitempty" tf:"helm_values,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Helm chart repository to download Cilium charts from (Default: `https://helm.cilium.io`).
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// When upgrading, reset the helm values to the ones built into the chart (Default: `false`).
	Reset *bool `json:"reset,omitempty" tf:"reset,omitempty"`

	// When upgrading, reuse the helm values from the latest release unless any overrides from are set from other flags. This option takes precedence over HelmResetValues (Default: `true`).
	Reuse *bool `json:"reuse,omitempty" tf:"reuse,omitempty"`

	// Set helm values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2 (Default: `[]`).
	Set []*string `json:"set,omitempty" tf:"set,omitempty"`

	// values in raw yaml to pass to helm. (Default: `empty`).
	Values *string `json:"values,omitempty" tf:"values,omitempty"`

	// Version of Cilium (Default: `v1.15.4`).
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Wait for Cilium status is ok (Default: `true`).
	Wait *bool `json:"wait,omitempty" tf:"wait,omitempty"`
}

type CiliumDeploymentParameters struct {

	// Datapath mode to use { tunnel | native | aws-eni | gke | azure | aks-byocni } (Default: `autodetected`).
	// +kubebuilder:validation:Optional
	DataPath *string `json:"dataPath,omitempty" tf:"data_path,omitempty"`

	// Helm chart repository to download Cilium charts from (Default: `https://helm.cilium.io`).
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// When upgrading, reset the helm values to the ones built into the chart (Default: `false`).
	// +kubebuilder:validation:Optional
	Reset *bool `json:"reset,omitempty" tf:"reset,omitempty"`

	// When upgrading, reuse the helm values from the latest release unless any overrides from are set from other flags. This option takes precedence over HelmResetValues (Default: `true`).
	// +kubebuilder:validation:Optional
	Reuse *bool `json:"reuse,omitempty" tf:"reuse,omitempty"`

	// Set helm values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2 (Default: `[]`).
	// +kubebuilder:validation:Optional
	Set []*string `json:"set,omitempty" tf:"set,omitempty"`

	// values in raw yaml to pass to helm. (Default: `empty`).
	// +kubebuilder:validation:Optional
	Values *string `json:"values,omitempty" tf:"values,omitempty"`

	// Version of Cilium (Default: `v1.15.4`).
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Wait for Cilium status is ok (Default: `true`).
	// +kubebuilder:validation:Optional
	Wait *bool `json:"wait,omitempty" tf:"wait,omitempty"`
}

// CiliumDeploymentSpec defines the desired state of CiliumDeployment
type CiliumDeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CiliumDeploymentParameters `json:"forProvider"`
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
	InitProvider CiliumDeploymentInitParameters `json:"initProvider,omitempty"`
}

// CiliumDeploymentStatus defines the observed state of CiliumDeployment.
type CiliumDeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CiliumDeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CiliumDeployment is the Schema for the CiliumDeployments API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cilium}
type CiliumDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CiliumDeploymentSpec   `json:"spec"`
	Status            CiliumDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CiliumDeploymentList contains a list of CiliumDeployments
type CiliumDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CiliumDeployment `json:"items"`
}

// Repository type metadata.
var (
	CiliumDeployment_Kind             = "CiliumDeployment"
	CiliumDeployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CiliumDeployment_Kind}.String()
	CiliumDeployment_KindAPIVersion   = CiliumDeployment_Kind + "." + CRDGroupVersion.String()
	CiliumDeployment_GroupVersionKind = CRDGroupVersion.WithKind(CiliumDeployment_Kind)
)

func init() {
	SchemeBuilder.Register(&CiliumDeployment{}, &CiliumDeploymentList{})
}