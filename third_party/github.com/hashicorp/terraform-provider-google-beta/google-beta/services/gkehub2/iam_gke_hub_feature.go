// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package gkehub2

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var GKEHub2FeatureIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"location": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"name": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
	},
}

type GKEHub2FeatureIamUpdater struct {
	project  string
	location string
	name     string
	d        tpgresource.TerraformResourceData
	Config   *transport_tpg.Config
}

func GKEHub2FeatureIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	location, _ := tpgresource.GetLocation(d, config)
	if location != "" {
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	values["location"] = location
	if v, ok := d.GetOk("name"); ok {
		values["name"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/features/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)", "(?P<location>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config, d.Get("name").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &GKEHub2FeatureIamUpdater{
		project:  values["project"],
		location: values["location"],
		name:     values["name"],
		d:        d,
		Config:   config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("location", u.location); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("name", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}

	return u, nil
}

func GKEHub2FeatureIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		values["project"] = project
	}

	location, _ := tpgresource.GetLocation(d, config)
	if location != "" {
		values["location"] = location
	}

	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/features/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)", "(?P<location>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &GKEHub2FeatureIamUpdater{
		project:  values["project"],
		location: values["location"],
		name:     values["name"],
		d:        d,
		Config:   config,
	}
	if err := d.Set("name", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *GKEHub2FeatureIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyFeatureUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "GET",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
	})
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = tpgresource.Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *GKEHub2FeatureIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyFeatureUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   u.d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *GKEHub2FeatureIamUpdater) qualifyFeatureUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{GKEHub2BasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/features/%s", u.project, u.location, u.name), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *GKEHub2FeatureIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/features/%s", u.project, u.location, u.name)
}

func (u *GKEHub2FeatureIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-gkehub2-feature-%s", u.GetResourceId())
}

func (u *GKEHub2FeatureIamUpdater) DescribeResource() string {
	return fmt.Sprintf("gkehub2 feature %q", u.GetResourceId())
}