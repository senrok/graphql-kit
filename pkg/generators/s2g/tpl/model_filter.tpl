input {{.ModelName}}Filter {
{{range .Fields}}    {{.LowerCamelFieldName}}: {{.ModelName}}{{.FieldName}}Filter {{.CustomTags}}  {{ printf "\n" }}{{end}}}
