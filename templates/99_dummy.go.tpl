{{- $alias := .Aliases.Table .Table.Name}}
{{- $orig_tbl_name := .Table.Name}}

func {{$alias.UpSingular}}Hello() {
    fmt.Println("Hello {{$alias.UpSingular}}")
}
