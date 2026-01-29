// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net/http"

	"github.com/umccr/terraform-provider-remscontent/internal/provider/data_sources"
	"github.com/umccr/terraform-provider-remscontent/internal/provider/functions"
	"github.com/umccr/terraform-provider-remscontent/internal/provider/resources"
	remsclient "github.com/umccr/terraform-provider-remscontent/internal/remsclient"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure RemsContentProvider satisfies various provider interfaces.
var _ provider.Provider = &RemsContentProvider{}
var _ provider.ProviderWithFunctions = &RemsContentProvider{}
var _ provider.ProviderWithEphemeralResources = &RemsContentProvider{}

// RemsContentProvider defines the provider implementation.
type RemsContentProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// RemsContentProviderModel describes the provider data model.
type RemsContentProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	ApiUser  types.String `tfsdk:"api_user"`
	ApiKey   types.String `tfsdk:"api_key"`
}

func (p *RemsContentProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "remscontent"
	resp.Version = p.version
}

func (p *RemsContentProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "REMS instance endpoint (DNS name only, not URI)",
				Required:            true,
			},
			"api_user": schema.StringAttribute{
				MarkdownDescription: "REMS API user",
				Required:            true,
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: "REMS API key",
				Required:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *RemsContentProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data RemsContentProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// configure a client to hit the authenticated endpoint
	cfg := remsclient.NewConfiguration()
	cfg.Host = data.Endpoint.ValueString()
	cfg.Scheme = "https"
	cfg.DefaultHeader = map[string]string{
		"x-rems-user-id": data.ApiUser.ValueString(),
		"x-rems-api-key": data.ApiKey.ValueString(),
		"Content-Type":   "application/json",
	}

	//transport := &BasePathRoundTripper{
	//	BasePath: "/api/",
	//	Base:     http.DefaultTransport,
	//}

	//transport := &BasePathRoundTripper{
	//	BasePath: "/api/",
	//	Base:     &DebugRoundTripper{Base: http.DefaultTransport, Ctx: ctx},
	//}

	cfg.HTTPClient = &http.Client{
		//	Transport: transport,
	}

	client := remsclient.NewAPIClient(cfg)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *RemsContentProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewCatalogueItemResource,
		resources.NewCategoryResource,
		resources.NewFormResource,
		resources.NewLicenseResource,
		resources.NewResourceResource,
		resources.NewWorkflowResource,
	}
}

func (p *RemsContentProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{}
}

func (p *RemsContentProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		data_sources.NewOrganizationDataSource,
	}
}

// :description :email :date :phone-number :table :header :texta :option :label :multiselect :ip-address :attachment :text

func (p *RemsContentProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		functions.NewFormFieldHeaderFunction,
		functions.NewFormFieldLabelFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &RemsContentProvider{
			version: version,
		}
	}
}
