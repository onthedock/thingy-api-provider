package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
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
