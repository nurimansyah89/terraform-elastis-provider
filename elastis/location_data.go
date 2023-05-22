package elastis

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nurimansyah89/terraform-elastis-provider/elastis/client"
	"github.com/nurimansyah89/terraform-elastis-provider/elastis/client/handlers"
)

type locationDataSource struct {
	client *client.ElastisClient
}

var (
	_ datasource.DataSource              = &locationDataSource{}
	_ datasource.DataSourceWithConfigure = &locationDataSource{}
)

func NewLocationDataSource() datasource.DataSource {
	return &locationDataSource{}
}

func (d *locationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_location"
}

func (d *locationDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"locations": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"display_name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *locationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.ElastisClient)
}
func (d *locationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Get current state
	var state LocationModel

	// Get locations
	locationHandler := handlers.NewLocation(d.client)
	locations, err := locationHandler.GetLocations()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Locations",
			"There was error when reading locations: "+err.Error(),
		)
		return
	}

	state.Locations = []Locations{}
	for _, location := range locations {
		state.Locations = append(state.Locations, Locations{
			DisplayName: types.StringValue(location.DisplayName),
		})
	}

	// Set refreshed state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
