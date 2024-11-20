# DataSource

Un *DataSource* es un elemento de sólo lectura, por lo que es más fácil implementarlo de entrada.

Empezamos creando el fichero `interal/provider/thingy_data_source.go`.

Lo asignamos a `datasource.DataSource` para asegurarnos que `ThingyDataSource` valida todos los interfaces del *framework*:

```go
// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ThingyDataSource{}
```

A continuación, definimos la función auxiliar que crea una instancia del *datasource*:

```go
func NewThingyDataSource() datasource.DataSource {
    return &ThingyDataSource{}
}
```

Creamos el *struct* para validar la definición del usuario del *DataSource*:

```go
// ThingyDataSourceModel describes the data source data model.
type ThingyDataSourceModel struct {
    Name types.String `tfsdk:"name"`
    Id   types.String `tfsdk:"id"`
}
```

Definimos el `DataSource`:

```go
// ThingyDataSource defines the data source implementation.
type ThingyDataSource struct {
    client *http.Client
}
```

Para *serializar* la configuración proporcionada por el usuario en un *struct* que podamos gestionar en Go, definimos un *model*:

```go
// ThingyDataSourceModel describes the data source data model.
type ThingyDataSourceModel struct {
    Name types.String `tfsdk:"name"`
    Id   types.String `tfsdk:"id"`
}
```

Como para el *provider*, el *type* `DataSource` también requiere unos métodos, aunque al tratarse de un elemento de sólo lectura, además del `Metadata` y el `Configure`, sólo usa el método `Read`:

- `Schema`
- `Metadata`
- `Configure`
- `Read`

## `Schema`

```go
func (d *ThingyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = schema.Schema{
        // This description is used by the documentation generator and the language server.
        MarkdownDescription: "Thingy data source",

        Attributes: map[string]schema.Attribute{
            "name": schema.StringAttribute{
                MarkdownDescription: "Thingy name",
                Required:            true,
            },
            "id": schema.StringAttribute{
                MarkdownDescription: "Id of any Thingy matching the provided name",
                Computed:            true,
            },
        },
    }
}
```

## `Metadata`

En el repositorio de *scaffolding*, el *datasource* de ejemplo añade el sufijo `_example` al nombre del *proveedor*, pero no tengo claro qué efecto tiene ésto en los recursos devueltos por el *datasource*.
