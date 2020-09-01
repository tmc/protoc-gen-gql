package gengraphql

import (
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
)

func tmplFuncs() template.FuncMap {
	var additionalFns = template.FuncMap{
		"fmtUnions": func(types []string) string {
			return strings.Join(types, " | ")
		},
		"fmtDoc": func(description string, prepends ...string) string {
			trimmed := strings.TrimSpace(description)
			if trimmed == "" {
				return ""
			}
			pre := ""
			for _, p := range prepends {
				pre = pre + p
			}
			lines := strings.Split(trimmed, "\n")
			for i, l := range lines {
				lines[i] = pre + strings.TrimSpace(l)
				if i == len(lines)-1 && len(l) > 0 {
					if l[len(l)-1] != '.' {
						lines[i] = lines[i] + "."
					}
				}
			}
			var docStr = pre + `"""`
			toprepend := append([]string{docStr}, lines...)
			toprepend = append(toprepend, docStr)
			return strings.Join(toprepend, "\n")
		},
	}
	fm := sprig.TxtFuncMap()
	for k, fn := range additionalFns {
		fm[k] = fn
	}
	return fm
}

const schemaTemplate = `
{{ if (gt (len .Service.Methods) 0) }}

type Query { {{ range .Service.Methods }}
    {{- fmtDoc .Doc "    " }}
    {{ .Name }}{{ .Request }}: {{ .Response }}!{{ end }}
}

{{ end }}

{{ if (gt (len .Service.Mutations) 0) }}

type Mutation { {{ range .Service.Mutations }}
    {{- fmtDoc .Doc "    " }}
    {{ .Name }}{{ .Request }}: {{ .Response }}!{{ end }}
}

{{ end }}

{{ range .Types }}
{{ fmtDoc .Doc }}
type {{ .Name }} {
{{ range .Fields }}
    {{- fmtDoc .Doc "    " }}
    {{ .Name }}: {{ .Type }}!
{{ end }}
    {{- if (eq (len .Fields) 0) }}
    responseMessage: String!
    {{- end }}
}
{{ end }}
{{ range .Inputs }}
{{ fmtDoc .Doc }}
input {{ .Name }} {
{{ range .Fields }}
    {{- fmtDoc .Doc "    " }}
    {{ .Name }}: {{ .Type }}
{{ end }}
{{- if (eq (len .Fields) 0) }}
    # This is being place because the number of fields would be zero otherwise. This is a bug in the generator.
    invalidInput: String!
{{- end }}
}
{{ end }}
{{ range .Enums }}
{{ fmtDoc .Doc }}
enum {{ .Name }} { {{ range .Fields }}
    {{- fmtDoc .Doc "    " }}
    {{ .Name }}{{ end }}
}{{ end }}
{{ range .Scalars }}
scalar {{ . }}
{{ end }}
{{ range .Unions }}
union {{ .Name }} = {{ fmtUnions .Types }}
{{ end }}
`
