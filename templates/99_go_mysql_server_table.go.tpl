{{- $alias := .Aliases.Table .Table.Name}}
{{- $orig_tbl_name := .Table.Name}}

{{- $toSimSQL := dict "string" "simsql.Text" "null.String" "simsql.Text" "int" "simsql.Int64" "null.Int" "simsql.Int64" "bool" "simsql.Boolean" "null.Bool" "simsql.Boolean"}}

func CreateDummy{{$alias.UpSingular}}Table(db *memory.Database) *memory.Table {
    tableName := "{{$alias.DownSingular}}"
    table := memory.NewTable(tableName, simsql.NewPrimaryKeySchema(simsql.Schema{
		{{- range $column := .Table.Columns}}
		{Name: "{{$column.Name}}", Type: {{$column.Type | get $toSimSQL}}, Nullable: {{$column.Nullable}}, Source: tableName, PrimaryKey: {{$.Table.PKey.Columns | has $column.Name}}},
		{{- end}}
	}), db.GetForeignKeyCollection())

    return table
}

// HACK: only for consuming unused imports
var _ = strconv.Itoa(42)
var _ null.String
