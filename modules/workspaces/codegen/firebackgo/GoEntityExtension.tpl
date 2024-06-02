package {{ .m.Path }}

/*
*	Generated by fireback {{ .fv }}
*	Written by Ali Torabi.
*	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
*/

import "github.com/torabian/fireback/modules/workspaces"

func {{ .e.Upper }}ActionCreate(
	dto *{{ .e.EntityName }}, query {{ .wsprefix }}QueryDSL,
) (*{{ .e.EntityName }}, *{{ .wsprefix }}IError) {
	return {{ .e.Upper }}ActionCreateFn(dto, query)
}

func {{ .e.Upper }}ActionUpdate(
	query {{ .wsprefix }}QueryDSL,
	fields *{{ .e.EntityName }},
) (*{{ .e.EntityName }}, *{{ .wsprefix }}IError) {
	return {{ .e.Upper }}ActionUpdateFn(query, fields)
}