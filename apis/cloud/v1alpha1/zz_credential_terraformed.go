/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Credential
func (mg *Credential) GetTerraformResourceType() string {
	return "rancher2_cloud_credential"
}

// GetConnectionDetailsMapping for this Credential
func (tr *Credential) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"amazonec2_credential_config[*].access_key": "spec.forProvider.amazonec2CredentialConfig[*].accessKeySecretRef", "amazonec2_credential_config[*].secret_key": "spec.forProvider.amazonec2CredentialConfig[*].secretKeySecretRef", "azure_credential_config[*].client_id": "spec.forProvider.azureCredentialConfig[*].clientIDSecretRef", "azure_credential_config[*].client_secret": "spec.forProvider.azureCredentialConfig[*].clientSecretSecretRef", "azure_credential_config[*].subscription_id": "spec.forProvider.azureCredentialConfig[*].subscriptionIDSecretRef", "digitalocean_credential_config[*].access_token": "spec.forProvider.digitaloceanCredentialConfig[*].accessTokenSecretRef", "google_credential_config[*].auth_encoded_json": "spec.forProvider.googleCredentialConfig[*].authEncodedJSONSecretRef", "linode_credential_config[*].token": "spec.forProvider.linodeCredentialConfig[*].tokenSecretRef", "openstack_credential_config[*].password": "spec.forProvider.openstackCredentialConfig[*].passwordSecretRef", "s3_credential_config[*].access_key": "spec.forProvider.s3CredentialConfig[*].accessKeySecretRef", "s3_credential_config[*].default_endpoint_ca": "spec.forProvider.s3CredentialConfig[*].defaultEndpointCASecretRef", "s3_credential_config[*].secret_key": "spec.forProvider.s3CredentialConfig[*].secretKeySecretRef", "vsphere_credential_config[*].password": "spec.forProvider.vsphereCredentialConfig[*].passwordSecretRef"}
}

// GetObservation of this Credential
func (tr *Credential) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Credential
func (tr *Credential) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Credential
func (tr *Credential) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Credential
func (tr *Credential) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Credential
func (tr *Credential) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Credential using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Credential) LateInitialize(attrs []byte) (bool, error) {
	params := &CredentialParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Credential) GetTerraformSchemaVersion() int {
	return 0
}