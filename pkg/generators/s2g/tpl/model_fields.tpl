enum {{.ModelName}}Fields {
{{range .Fields}}    {{.OriginBsonName}} {{ printf "\n" }}{{end}}}
