package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure ThingyProvider satisfies various provider interfaces.
var _ provider.Provider = &ThingyProvider{}

// ThingyProvider defines the provider implementation.
type ThingyProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// ThingyProviderModel describes the provider data model.
type ThingyProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *ThingyProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "thingy"
	resp.Version = p.version
}

func (p *ThingyProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Endpoint for the Thingy provider to connect to",
				Optional:            true,
			},
		},
	}
}

func (p *ThingyProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data ThingyProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	client := http.DefaultClient
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *ThingyProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
	// return []func() datasource.DataSource{
	// 	NewExampleDataSource,
	// }
}

func (p *ThingyProvider) Resources(ctx context.Context) []func() resource.Resource {
	return nil
	// return []func() resource.Resource{
	// 	NewExampleResource,
	// }
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ThingyProvider{
			version: version,
		}
	}
}
