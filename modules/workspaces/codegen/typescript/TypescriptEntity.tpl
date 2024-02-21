{{ template "tsimport" . }}

// In this section we have sub entities related to this object
{{ range .children }}
export class {{ .FullName }} extends BaseEntity {
  {{ template "definitionrow" .CompleteFields }}
}
{{ end }}

// Class body

export type {{ .e.EntityName }}Keys =
  keyof typeof {{ .e.EntityName }}.Fields;

export class {{ .e.EntityName }} extends BaseEntity {

  {{ template "definitionrow" .e.CompleteFields }}

  {{ if eq .ctx.Ts.IncludeStaticNavigation true }}
    {{ template "staticnavigation" . }}
  {{ end }}
  

  {{ if eq .ctx.Ts.IncludeFirebackDef true }}
    {{ template "fbdefinition" . }}
  {{ end }}

  {{ if eq .ctx.Ts.IncludeStaticField true }}
    {{ template "staticfield" . }}
  {{ end }}
  
  
}
