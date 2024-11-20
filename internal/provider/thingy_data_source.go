package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ThingyDataSource{}

func NewThingyDataSource() datasource.DataSource {
	return &ThingyDataSource{}
}

// ThingyDataSource defines the data source implementation.
type ThingyDataSource struct {
	client *http.Client
}

// ThingyDataSourceModel describes the data source data model.
type ThingyDataSourceModel struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
}

func (d *ThingyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_thingy"
}

func (d *ThingyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Thingy data source",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Thingy name",
				Required:            true,
			},
			"id": schema.ListAttribute{
				MarkdownDescription: "List of Thingy Ids matching the provided name",
				Computed:            true,
			},
		},
	}
}

func (d *ThingyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
}
