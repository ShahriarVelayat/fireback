package {{ .m.Path }}

import (
	"encoding/json"
	"fmt"

	"strings"
	"github.com/urfave/cli"
)

{{ template "goimport" . }}

{{ range .children }}
type {{ .FullName }} struct {
    {{ template "definitionrow" (arr .CompleteFields $.wsprefix) }}
}
func ( x * {{ .FullName }}) RootObjectName() string {
	return "{{ $.e.DtoName }}"
}
{{ end }}

{{ template "dtoCastFromCli" . }}


var {{ .e.DtoName }}CommonCliFlagsOptional = []cli.Flag{
  {{ template "entityCommonCliFlag" (arr .e.CompleteFields "") }}
}

type {{ .e.DtoName }} struct {
    {{ template "definitionrow" (arr .e.CompleteFields $.wsprefix) }}
}

func (x* {{ .e.DtoName }}) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}

func (x* {{ .e.DtoName }}) JsonPrint()  {
    fmt.Println(x.Json())
}

// This is an experimental way to create new dtos, with exluding the pointers as helper.
func New{{ .e.DtoName }}(
{{ range .e.CompleteFields }}
	{{ if eq .Type "string" }}
	{{ .PublicName }} string,
	{{ end }}
{{ end }}	
) {{ .e.DtoName }} {
    return {{ .e.DtoName }}{
	{{ range .e.CompleteFields }}

	{{ if eq .Type "string" }}
	{{ .PublicName }}: &{{ .PublicName }},
	{{ end }}

	{{ end }}
    }
}