package s2g

import (
	"context"
	"github.com/iancoleman/strcase"
	"github.com/senrok/yadal"
	"github.com/senrok/yadal/providers/fs"
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

func generateFields(rt reflect.Type) []ModelField {
	var mf []ModelField
	for i := 0; i < rt.NumField(); i++ {
		f := ModelField{}
		field := rt.Field(i)
		f.Name = field.Name
		f.StructField = field

		if field.Anonymous {
			mf = append(mf, generateFields(field.Type)...)
			continue
		}

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
		} else {
			continue
		}
		mf = append(mf, f)
	}
	return mf
}

func newGenerator(rt reflect.Type) *generator {
	g := generator{}

	for rt.Kind() == reflect.Ptr ||
		rt.Kind() == reflect.Interface {
		rt = rt.Elem()
	}
	g.Ctx.ModelName = rt.Name()
	g.Ctx.ModelField = generateFields(rt)
	return &g
}

func NewGenerator(model interface{}) *generator {
	rt := reflect.TypeOf(model)
	return newGenerator(rt)
}

var fns = template.FuncMap{
	"last": func(x int, a interface{}) bool {
		return x == reflect.ValueOf(a).Len()-1
	},
}

// go:embed ./tpl/field_filter.tpl
var fieldFilterTemplate string

// go:embed ./tpl/model_filter.tpl
var modelFilterTemplate string

// go:embed ./tpl/model_fields.tpl
var modelFieldsTemplate string

// go:embed ./tpl/model_sort.tpl
var modelSortTemplate string

func (g generator) Generate() (string, error) {
	sb := strings.Builder{}
	c := g.ToTplCtx()

	tpl, err := template.New("fieldFilterTemplate").Parse(fieldFilterTemplate)
	if err != nil {
		return "", err
	}

	for _, f := range c.Fields {
		err = tpl.Execute(&sb, f)
		if err != nil {
			return "", err
		}
	}

	tpl, err = template.New("modelFilterTemplate").Parse(modelFilterTemplate)
	if err != nil {
		return "", err
	}

	err = tpl.Execute(&sb, c)
	if err != nil {
		return "", err
	}

	tpl, err = template.New("modelFieldsTemplate").Parse(modelFieldsTemplate)
	if err != nil {
		return "", err
	}

	err = tpl.Execute(&sb, c)
	if err != nil {
		return "", err
	}

	tpl, err = template.New("modelSortTemplate").Parse(modelSortTemplate)
	if err != nil {
		return "", err
	}

	err = tpl.Execute(&sb, c)
	if err != nil {
		return "", err
	}

	return sb.String(), nil
}

type Config struct {
	Model          interface{}
	OutputFileName string
}

func NewBatchGenerator(dir string, configs ...Config) error {
	acc := fs.NewDriver(fs.Options{Root: dir})
	op := yadal.NewOperatorFromAccessor(acc)

	for _, config := range configs {
		rv := reflect.ValueOf(config.Model)
		rv = reflect.Indirect(rv)
		rt := rv.Type()

		if config.OutputFileName == "" {
			config.OutputFileName = strcase.ToKebab(rt.Name()) + ".generated.graphql"
		}
		output, err := newGenerator(rt).Generate()
		if err != nil {
			return err
		}
		obj := op.Object(config.OutputFileName)
		err = obj.Write(context.TODO(), []byte(output))
		if err != nil {
			return err
		}
	}
	return nil
}
