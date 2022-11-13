input {{.ModelName}}Sort {
{{range .Fields}}    {{.LowerCamelFieldName}}: {{.SortScalar}}  {{ printf "\n" }}{{end}}
}
