package elastis

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nurimansyah89/terraform-elastis-provider/elastis/client"
)

type elastisProvider struct{}

// Main provider configuration
type elastisProviderModel struct {
	Token types.String `tfsdk:"token"`
}

var (
	_ provider.Provider = &elastisProvider{}
)

func New() provider.Provider {
	return &elastisProvider{}
}

func (p *elastisProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "elastis"
}

func (p *elastisProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"token": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (p *elastisProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config elastisProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check unknown
	if config.Token.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Unknown Elastis.id Token",
			"Please provide the token entity",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// Assign from config
	var token string
	if !config.Token.IsNull() {
		token = config.Token.ValueString()
	}

	// Check empty
	if token == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Missing Elastis.id Token",
			"Please provide the token API! You can get the token inside your console account to generate a new one.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// Create elastis client
	client := client.New(token)

	// Resources
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *elastisProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewLocationDataSource,
	}
}

func (p *elastisProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewVMResource,
	}
}
