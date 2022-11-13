package s2g

import (
	"reflect"
	"strings"
	"text/template"
)

type generator struct {
	Ctx
}

func RemoveStringQuote(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

func NewGenerator(model interface{}) generator {
	g := generator{
		Ctx{},
	}
	rt := reflect.TypeOf(model)
	for rt.Kind() == reflect.Ptr ||
		rt.Kind() == reflect.Interface {
		rt = rt.Elem()
	}
	g.Ctx.ModelName = rt.Name()

	for i := 0; i < rt.NumField(); i++ {
		f := ModelField{}
		field := rt.Field(i)
		f.Name = field.Name

		// TODO: derives primitive type by reflect
		if value, ok := field.Tag.Lookup("filterScalar"); !ok {
			continue
		} else {
			f.Scalar = RemoveStringQuote(value)
		}

		if value, ok := field.Tag.Lookup("filter"); ok {
			tags := strings.Split(RemoveStringQuote(value), " ")
			for _, tag := range tags {
				if filters, ok := FilterMap[tag]; ok {
					f.FilterOperators = append(f.FilterOperators, filters...)
				}
			}
		}
		g.Ctx.ModelField = append(g.Ctx.ModelField, f)
	}
	return g
}

func (g generator) Generate() (string, error) {
	sb := strings.Builder{}
	c := g.ToTplCtx()

	tpl, err := template.ParseFiles("./field_filter.tpl")
	if err != nil {
		return "", err
	}

	for _, f := range c.Fields {
		err = tpl.Execute(&sb, f)
		if err != nil {
			return "", err
		}
	}

	tpl, err = template.ParseFiles("./model_filter.tpl")
	if err != nil {
		return "", err
	}

	err = tpl.Execute(&sb, c)
	if err != nil {
		return "", err
	}

	return sb.String(), nil
}
