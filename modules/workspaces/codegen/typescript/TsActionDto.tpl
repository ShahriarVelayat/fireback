
{{ define  "tsactionimport" }}

{{ range $key, $value := . }}
{{ if $value.Items}}
import {
  {{ range $value.Items }}
    {{ .}},
  {{ end }}

} from "{{ $key}}"
{{ end }}
{{ end }}


{{ end }}

{{ template "tsactionimport" .tsactionimports }}


{{ range .m.Actions }}

    {{ if .In.Fields }}

export class {{ .Upper }}ActionReqDto {
    {{ template "definitionrow" .In.Fields }}
    {{ template "actionDtoFields" .In.Fields }}
}

    {{ end }}

    {{ if .Out.Fields }}

export class {{ .Upper }}ActionResDto {
    {{ template "definitionrow" .Out.Fields }}
    {{ template "actionDtoFields" .Out.Fields }}
}

    {{ end }}

{{ end }}