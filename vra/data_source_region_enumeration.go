package vra

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/vra-sdk-go/pkg/client/cloud_account"
	"github.com/vmware/vra-sdk-go/pkg/models"
)

func dataSourceRegionEnumeration() *schema.Resource {
	return &schema.Resource{
		Read:               dataSourceRegionEnumerationRead,
		DeprecationMessage: "'region_enumeration' is deprecated. Use 'region_enumeration_vsphere' instead.",

		Schema: map[string]*schema.Schema{
			"accept_self_signed_cert": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dcid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"regions": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceRegionEnumerationRead(d *schema.ResourceData, meta interface{}) error {
	apiClient := meta.(*Client).apiClient

	getResp, err := apiClient.CloudAccount.EnumerateVSphereRegions(cloud_account.NewEnumerateVSphereRegionsParams().WithBody(&models.CloudAccountVsphereSpecification{
		AcceptSelfSignedCertificate: d.Get("accept_self_signed_cert").(bool),
		Dcid:                        d.Get("dcid").(string),
		HostName:                    withString(d.Get("hostname").(string)),
		Password:                    withString(d.Get("password").(string)),
		Username:                    withString(d.Get("username").(string)),
	}))

	if err != nil {
		return err
	}

	d.Set("regions", getResp.Payload.ExternalRegionIds)
	d.SetId(d.Get("hostname").(string))

	return nil
}
