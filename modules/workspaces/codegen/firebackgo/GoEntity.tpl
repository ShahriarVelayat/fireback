package {{ .m.Path }}

/*
*	Generated by fireback {{ .fv }}
*	Written by Ali Torabi.
*	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
*/

import (
    "github.com/gin-gonic/gin"

    {{ if ne .m.Path "workspaces" }}
	"github.com/torabian/fireback/modules/workspaces"
	{{ end }}
 
	"fmt"
	"encoding/json"
	"strings"
	"github.com/schollz/progressbar/v3"
	"github.com/gookit/event"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	jsoniter "github.com/json-iterator/go"

    {{ if or (.e.Cte) (.e.Queries) }}
    queries "{{ .gofModule }}/modules/{{ .m.Path }}/queries"
    {{ end }}

	"embed"
	reflect "reflect"

	"github.com/urfave/cli"
	{{ if .hasSeeders }}
	seeders "{{ .gofModule }}/modules/{{ .m.Path }}/seeders/{{ .e.Upper }}"
	{{ end }}
	{{ if .hasMocks }}
	mocks "{{ .gofModule }}/modules/{{ .m.Path }}/mocks/{{ .e.Upper }}"
	{{ end }}
	{{ if .hasMetas }}
	metas "{{ .gofModule }}/modules/{{ .m.Path }}/metas"
	{{ end }}
   
)


{{ if .e.PrependScript }}
	{{ .e.PrependScript }}
{{ end }}

{{ template "goimport" . }}


{{ if .hasSeeders }}
var {{ .e.Name }}SeedersFs = &seeders.ViewsFs
{{ else }}
var {{ .e.Name }}SeedersFs *embed.FS = nil
{{ end }}

func Reset{{ .e.Upper }}Seeders(fs *embed.FS) {
	{{ .e.Name }}SeedersFs = fs
}

{{ range .children }}
type {{ .FullName }} struct {
    {{ template "defaultgofields" . }}
    {{ template "definitionrow" (arr .CompleteFields $.wsprefix) }}

	{{ if .LinkedTo }}
	LinkedTo *{{ .LinkedTo }} `yaml:"-" gorm:"-" json:"-" sql:"-"`
	{{ end }}
}

func ( x * {{ .FullName }}) RootObjectName() string {
	return "{{ $.e.EntityName }}"
}

{{ end }}

type {{ .e.EntityName }} struct {
    {{ template "defaultgofields" .e }}
    {{ template "definitionrow" (arr .e.CompleteFields $.wsprefix) }}

    {{ if .e.HasTranslations }}
    Translations     []*{{ .e.PolyglotName}} `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId;constraint:OnDelete:CASCADE"`
    {{ end }}

    Children []*{{ .e.EntityName }} `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`

    LinkedTo *{{ .e.EntityName }} `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var {{ .e.Upper }}PreloadRelations []string = []string{}


{{ template "eventsAndMeta" . }}

{{ template "polyglottable" . }}

{{ template "entitychildactions" . }}

{{ template "entityformatting" . }}

{{ template "mockingentity" . }}

{{ template "getEntityTranslateFields" . }}

{{ template "entitySeederInit" . }}

{{ template "entityAssociationCreate" . }}

{{ template "entityRelationContentCreation" . }}

{{ template "relationContentUpdate" . }}

{{ template "polyglot" . }}

{{ template "entityValidator" . }}

{{ template "entitySanitize" . }}

{{ template "entityBeforeCreateActions" . }}

{{ template "batchActionCreate" . }}

{{ template "entityDeleteEntireChildren" . }}

{{ template "entityActionCreate" . }}

{{ template "entityActionGetAndQuery" . }}

{{ template "queriesAndPivot" . }}

{{ template "entityUpdateExec" . }}

{{ template "entityUpdateAction" . }}

{{ template "entityRemoveAndCleanActions" . }}

{{ template "entityBulkUpdate" . }}

{{ template "entityExtensions" . }}

{{ template "entityMeta" . }}

{{ template "entityImportExport" . }}

{{ template "cliFlags" . }}

{{ template "entityCliCommands" . }}

{{ template "entityCastFromCli" . }}

{{ template "entityMockAndSeeders" . }}

{{ template "entityCliImportExportCmd" . }}

{{ template "entityCliActionsCmd" . }}

{{ template "entityHttp" . }}

{{ template "entityPermissions" . }}

{{ template "recursiveGetEnums" (arr .e.CompleteFields .e.Upper)}}

{{ template "entityDistinctOperations" . }}

{{ if .e.Messages }}
{{ template "messageCode" (arr .e.Name .e.Messages $.wsprefix)}}
{{ end }}

var {{ .e.EntityName }}Bundle = {{ $.wsprefix }}EntityBundle{
	Permissions: ALL_{{ .e.AllUpper }}_PERMISSIONS,
	CliCommands: []cli.Command{
		{{ .e.Upper }}CliFn(),
	},
	Actions: Get{{ .e.Upper }}Module2Actions(),
	AutoMigrationEntities: []interface{}{
		&{{ .e.EntityName }}{},
		{{ range .children }}
		&{{ .FullName }}{},
		{{ end }}
  	},
}