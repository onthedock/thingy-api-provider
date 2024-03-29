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

## Creamos la carpeta para el *provider*

```console
mkdir -p internal/provider
touch internal/provider/provider.go
```

Usamos la técnica de asignar el *struct* del *provider* al *blank identifier* para validar que el nuevo tipo valida todos los *interfaces* del *provider*:

```go
// Ensure ThingyProvider satisfies various provider interfaces.
var _ provider.Provider = &ThingyProvider{}
```

A continuación, definimos el *struct*, que únicamente tiene un campo:

```go
// ThingyProvider defines the provider implementation.
type ThingyProvider struct {
    // version is set to the provider version on release, "dev" when the
    // provider is built and ran locally, and "test" when running acceptance
    // testing.
    version string
}
```

Tamnbién definimos el *model* para el *provider*; este *struct* se usa para *serializar* los datos proporcionados por el *consumer* (el usuario) a la hora de configurar el *provider*.

En el recurso de ejemplo, únicamente es necesario proporcionar una propiedad, el *endpoint* al que se contectará el *provider*.

## Métodos del *provider*

El *provider* tiene varios métodos que deben satisfacerse:

- `Metadata`
- `Schema`
- `Configure`
- `DataSources`
- `Resources`

Así que tenemos que definirlos para que nuestro *provider* sea un *provider* válido.

### `Metadata`

```go
func (p *ThingyProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
    resp.TypeName = "thingy"
    resp.Version = p.version
}
```

### `Schema`

En `Schema` definimos los atributos del *provider*; en nuestro caso, sólo tenemos un atributo opcional, el *endpoint*:

```go
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
```

### `Configure`

El método `Configure` es donde realizamos la lectura del fichero de configuración del *provider*, así como de las variables de entorno relevantes requeridas para configurar el cliente usado por el *provider* para conectar con la API.

De momento, lo obviamos.

```go
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
```

En el bloque comentado, como hemos definido el *endpoint* como opcional, deberíamos establecer algún valor por defecto, por ejemplo, conectar con un *endpoint* local o algo por el estilo.

Después de leer la configuración, si no está especificado el *endpoint*, podríamos intentar obtener la configuración desde una variable de entorno.

De nuevo, dejamos esta configuración del *provider* para más tarde.

### `DataSources` y `Resources`

De momento, sólo estamos añadiendo los métodos requeridos para que el nuevo *provier* valide la definición del `provider.Provider`.
Por ello, dejamos vacíos tanto los *recursos* como los *datasources* gestionados por el *provider*.

### `New`

La función `New()` instancia el *provider*.

```go
func New(version string) func() provider.Provider {
    return func() provider.Provider {
        return &ThingyProvider{
            version: version,
        }
    }
}
```

Con esta configuración básica deberíamos ser capaces de compilar el *provider*.

```go
$ go install .
$ ls /go/bin/ | grep -i terraform
terraform-provider-thingy
```

Y si intentamos ejecutarlo:

```console
 $ /go/bin/terraform-provider-thingy 
This binary is a plugin. These are not meant to be executed directly.
Please execute the program that consumes these plugins, which will
load any plugins automatically
```
