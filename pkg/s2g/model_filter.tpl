input {{.ModelName}}Filter {
{{range .Fields}}    {{.LowerCamelFieldName}}: {{.ModelName}}{{.FieldName}}Filter {{end}}
}
