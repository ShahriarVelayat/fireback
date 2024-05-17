package {{ .m.Path }}

import (
	"github.com/gin-gonic/gin"
    "github.com/urfave/cli"
    {{ if ne .wsprefix "" }}
	"github.com/torabian/fireback/modules/workspaces"
    {{ end }}
)


{{ range .m.Actions }}

{{ if .SecurityModel }}
var {{ .Upper }}SecurityModel = &{{ $.wsprefix }}SecurityModel{
   ActionRequires: []{{ $.wsprefix }}PermissionInfo{ 
        {{ range .SecurityModel.ActionRequires }}
            {
				CompleteKey: "{{ .CompleteKey }}",
			},
        {{ end }}
    },
}
{{ else }}
var {{ .Upper }}SecurityModel *{{ $.wsprefix }}SecurityModel = nil
{{ end }}

    {{ if .In.Fields }}


type {{ .Upper }}ActionReqDto struct {
    {{ template "definitionrow" (arr .In.Fields $.wsprefix) }}
}

func ( x * {{ .Upper }}ActionReqDto) RootObjectName() string {
	return "{{ $.m.Path }}"
}

var {{ .Upper }}CommonCliFlagsOptional = []cli.Flag{
  {{ template "dtoCliFlag" (arr .In.Fields "") }}
}


func {{ .Upper }}ActionReqValidator(dto *{{ .Upper }}ActionReqDto) *{{ $.wsprefix }}IError {
    err := {{ $.wsprefix }}CommonStructValidatorPointer(dto, false)

    {{ range .In.Fields }}
      {{ if  eq .Type "array"  }}
        if dto != nil && dto.{{ .UpperPlural }} != nil {
          {{ $.wsprefix }}AppendSliceErrors(dto.{{ .UpperPlural }}, false, "{{ .Name}}", err)
        }
      {{ end }}
    {{ end }}
    return err
  }


func Cast{{ .Upper }}FromCli (c *cli.Context) *{{ .Upper }}ActionReqDto {
	template := &{{ .Upper }}ActionReqDto{}

	{{ template "entityCliCastRecursive" (arr .In.Fields "")}}

	return template
}

    {{ end }}

    {{ if .Out.Fields }}

type {{ .Upper }}ActionResDto struct {
    {{ template "definitionrow" (arr .Out.Fields $.wsprefix) }}
}

func ( x * {{ .Upper }}ActionResDto) RootObjectName() string {
	return "{{ $.m.Path }}"
}


    {{ end }}

{{ if or (eq .Method "reactive")}}


var {{ .Upper }}ActionImp = {{ $.wsprefix }}DefaultEmptyReactiveAction

{{ else }}

type {{ .Name }}ActionImpSig func(
    {{ if ne .ActionReqDto "nil" }}req {{ .ActionReqDto }}, {{ end}}
    q {{ $.wsprefix }}QueryDSL) ({{ .ActionResDto }},
    {{ if (eq .FormatComputed "QUERY") }} *workspaces.QueryResultMeta, {{ end }}
    *{{ $.wsprefix }}IError,
)
var {{ .Upper }}ActionImp {{ .Name }}ActionImpSig
{{ end }}


{{ if or (eq .Method "reactive")}}
// Reactive action does not have that
{{ else }}
func {{ .Upper }}ActionFn(
    {{ if ne .ActionReqDto "nil" }}req {{ .ActionReqDto }}, {{ end}}
    q {{ $.wsprefix }}QueryDSL,
) (
    {{ .ActionResDto }},
    {{ if (eq .FormatComputed "QUERY") }} *workspaces.QueryResultMeta, {{ end }}
    *{{ $.wsprefix }}IError,
) {

    if {{ .Upper }}ActionImp == nil {
        return nil, {{ if (eq .FormatComputed "QUERY") }} nil, {{ end }} nil
    }

    return {{ .Upper }}ActionImp({{ if ne .ActionReqDto "nil" }}req, {{ end}} q)
}
{{ end }}

var {{ .Upper }}ActionCmd cli.Command = cli.Command{
	Name:  "{{ .ComputedCliName }}",
	Usage: "{{ .Description }}",
    {{ if (eq .FormatComputed "QUERY") }}
    Flags: workspaces.CommonQueryFlags,
    {{ end }}
    {{ if .In.Fields }}
	Flags: {{ .Upper }}CommonCliFlagsOptional,
    {{ else if .In.Entity }}
	Flags: {{ .In.EntityPure }}CommonCliFlagsOptional,
    {{ end }}
	Action: func(c *cli.Context) {

		query := {{ $.wsprefix }}CommonCliQueryDSLBuilderAuthorize(c, {{ .Upper }}SecurityModel)

        {{ if or (ne .Method "reactive")}}
        {{ if .In.Fields }}
		dto := Cast{{ .Upper }}FromCli(c)
        {{ else if .In.Entity }}
		dto := Cast{{ .In.EntityPure }}FromCli(c)
        {{ end }}

        {{ if or (eq .FormatComputed "QUERY")}}
		result, _, err := {{ .Upper }}ActionFn(query)
        {{ else }}
		result, err := {{ .Upper }}ActionFn(dto, query)
        {{ end }}

		{{ $.wsprefix }}HandleActionInCli(c, result, err, map[string]map[string]string{})
        {{ else }}
        {{ $.wsprefix }}CliReactivePipeHandler(query, {{ .Upper }}ActionImp)
        {{end}}
	},
}

{{ end }}


func {{ .m.PublicName }}CustomActions() []{{ $.wsprefix }}Module2Action {
	routes := []{{ $.wsprefix }}Module2Action{
        {{ range .m.Actions }}
		{
			Method: "{{ .MethodAllUpper }}",
			Url:    "{{ .ComputedUrl }}",
            SecurityModel: {{ .Upper }}SecurityModel,
            Group: "{{ $.m.PublicName }}Custom",
			Handlers: []gin.HandlerFunc{
                {{ if or (eq .Method "reactive")}}
                {{ $.wsprefix }}ReactiveSocketHandler({{ .Upper }}ActionImp),
                {{ else }}
				func(c *gin.Context) {
                    // {{ .FormatComputed }} - {{ .Method }}
                    
                    {{ if or (eq .FormatComputed "POST") (eq .Method "POST") (eq .Method "post") }}
                        {{ $.wsprefix }}HttpPostEntity(c, {{ .Upper }}ActionFn)
                    {{ end }}
                    {{ if or (eq .FormatComputed "QUERY")}}
                        {{ $.wsprefix }}HttpQueryEntity2(c, {{ .Upper }}ActionFn)
                    {{ end }}
                },
                {{ end }}
			},
			Format:         "{{ .FormatComputed }}",
            {{ if or (ne .Method "reactive")}}
			Action:         {{ .Upper }}ActionFn,
            {{end}}
            {{ if .ComputeResponseEntity }}
			ResponseEntity: {{.ComputeResponseEntity}},
            Out: {{ $.wsprefix }}Module2ActionBody{
                Entity: "{{ .ComputeResponseEntityS }}",
            },
            {{ end }}
            {{ if .ComputeRequestEntity}}
			RequestEntity: {{.ComputeRequestEntity}},
            In: {{ $.wsprefix }}Module2ActionBody{
                Entity: "{{ .ComputeRequestEntityS }}",
            },
            {{ end }}
		},
        {{ end }}
	}
    
	return routes
}

var {{ .m.Upper }}CustomActionsCli = []cli.Command {
{{ range .m.Actions }}
    {{ .Upper }}ActionCmd,
{{ end }}
}

func Get{{ .m.Upper }}ActionsCli() []cli.Command {
	return {{ .m.Upper }}CustomActionsCli
}
