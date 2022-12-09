input {{.ModelName}}{{.FieldName}}Filter {
{{range .Operators}}    {{.Name}}: {{.ScalarName}} {{.CustomTags}}{{ printf "\n" }}{{end}}}
