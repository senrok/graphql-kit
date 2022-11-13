input {{.ModelName}}Filter {
{{range .Fields}}    {{.LowerCamelFieldName}}: {{.ModelName}}{{.FieldName}}Filter {{ printf "\n" }}{{end}}
}
