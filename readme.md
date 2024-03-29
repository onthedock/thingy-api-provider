# Terraform Provider Thingy

> Uso como referencia el repositorio [terraform-provider-scaffolding-framework](https://github.com/hashicorp/terraform-provider-scaffolding-framework/tree/main)

## Pre-requisitos

Terraform versión 1.0 o superior:

```console
$ terraform version
Terraform v1.7.5
on linux_amd64
```

Go versión 1.21 o superior:

```console
$ go version
go version go1.22.0 linux/amd64
```

## Inicializar el módulo del *provider*

```console
go mod init github.com/onthedock/terraform-provider-thingy
```

Obtenemos el paquete: `github.com/hashicorp/terraform-plugin-framework`

```console
$ go get github.com/hashicorp/terraform-plugin-framework
go: downloading github.com/hashicorp/terraform-plugin-framework v1.7.0
go: added github.com/hashicorp/terraform-plugin-framework v1.7.0
```

Creamos el fichero `main.go`.

Definimos la variable para definir la versión del *provider*:

```go
var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	version string = "dev"

	// goreleaser can pass other information to the main package, such as the specific commit
	// https://goreleaser.com/cookbooks/using-main.version/
)
```

Y a continuación, la función `main`:

```go
func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		// TODO: Update this string with the published name of your provider.
		// Also update the tfplugindocs generate command to either remove the
		// -provider-name flag or set its value to the updated provider name.
		Address: "registry.terraform.io/xaviaznar/thingy",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
```

La función `main` crea una nueva instancia del *provider* mediante `provider.New`.
