/**
Current file is set of definitions, to create Module2 yaml files.
Module2 is a declarative way of creating backend entities, crud actions on them,
and many complex operation. Fireback would generate those codes for many languages
both for backend and front-end purposes.

Backend code can be generated in: C and Golang
Front-end code can be generated in: Angular, React, Pure TypeScript, Android Java, Swift
*/

package workspaces

import (
	"github.com/gin-gonic/gin"
)

// Module2 struct represents the entire file tree
type Module2 struct {
	Path        string           `yaml:"path,omitempty" json:"path,omitempty"`
	Description string           `yaml:"description,omitempty" json:"description,omitempty"`
	Version     string           `yaml:"version,omitempty" json:"version,omitempty"`
	Name        string           `yaml:"name,omitempty" json:"name,omitempty"`
	Entities    []Module2Entity  `yaml:"entities,omitempty" json:"entities,omitempty"`
	Dto         []Module2DtoBase `yaml:"dto,omitempty" json:"dto,omitempty"`
	Actions     []Module2Action  `yaml:"actions,omitempty" json:"actions,omitempty"`
}

type Module2FieldOf struct {
	Key string `yaml:"k,omitempty" json:"k,omitempty"`
}

type Module2Field struct {
	LinkedTo            string               `yaml:"linkedTo,omitempty" json:"linkedTo,omitempty"`
	Description         string               `yaml:"description,omitempty" json:"description,omitempty"`
	Name                string               `yaml:"name,omitempty" json:"name,omitempty"`
	Type                string               `yaml:"type,omitempty" json:"type,omitempty"`
	Primitive           string               `yaml:"primitive,omitempty" json:"primitive,omitempty"`
	Target              string               `yaml:"target,omitempty" json:"target,omitempty"`
	Validate            string               `yaml:"validate,omitempty" json:"validate,omitempty"`
	ExcerptSize         int                  `yaml:"excerptSize,omitempty" json:"excerptSize,omitempty"`
	Default             interface{}          `yaml:"default,omitempty" json:"default,omitempty"`
	Translate           bool                 `yaml:"translate,omitempty" json:"translate,omitempty"`
	Unsafe              bool                 `yaml:"unsafe,omitempty" json:"unsafe,omitempty"`
	AllowCreate         bool                 `yaml:"allowCreate,omitempty" json:"allowCreate,omitempty"`
	Module              string               `yaml:"module,omitempty" json:"module,omitempty"`
	Json                string               `yaml:"json,omitempty" json:"json,omitempty"`
	OfType              []*Module2FieldOf    `yaml:"of,omitempty" json:"of,omitempty"`
	Yaml                string               `yaml:"yaml,omitempty" json:"yaml,omitempty"`
	IdFieldGorm         string               `yaml:"idFieldGorm,omitempty" json:"idFieldGorm,omitempty"`
	ComputedType        string               `yaml:"computedType,omitempty" json:"computedType,omitempty"`
	ComputedTypeClass   string               `yaml:"computedTypeClass,omitempty" json:"computedTypeClass,omitempty"`
	BelongingEntityName string               `yaml:"-" json:"-"`
	Matches             []*Module2FieldMatch `yaml:"matches,omitempty" json:"matches,omitempty"`
	Gorm                string               `yaml:"gorm,omitempty" json:"json,omitempty"`
	GormMap             GormOverrideMap      `yaml:"gormMap,omitempty" json:"gormMap,omitempty"`

	Sql string `yaml:"sql,omitempty" json:"sql,omitempty"`
	// This is the name of field considering how deep it is
	FullName string          `yaml:"fullName,omitempty" json:"fullName,omitempty"`
	Fields   []*Module2Field `yaml:"fields,omitempty" json:"fields,omitempty"`
}

type Module2FieldMatch struct {
	Dto *string `yaml:"dto,omitempty" json:"dto,omitempty"`
}

type GormOverrideMap struct {
	WorkspaceId string `yaml:"workspaceId,omitempty" json:"workspaceId,omitempty"`
	UserId      string `yaml:"userId,omitempty" json:"userId,omitempty"`
}

type Security struct {
	Model string `yaml:"model,omitempty" json:"model,omitempty"`
}

type Module2Http struct {
	Query bool `yaml:"query,omitempty" json:"query,omitempty"`
}

type Module2Entity struct {
	Name           string          `yaml:"name,omitempty" json:"name,omitempty"`
	DistinctBy     string          `yaml:"distinctBy,omitempty" json:"distinctBy,omitempty"`
	NoQuery        bool            `yaml:"noQuery,omitempty" json:"noQuery,omitempty"`
	Access         string          `yaml:"access,omitempty" json:"access,omitempty"`
	QueryScope     string          `yaml:"queryScope,omitempty" json:"queryScope,omitempty"`
	Http           Module2Http     `yaml:"http,omitempty" json:"http,omitempty"`
	Patch          bool            `yaml:"patch,omitempty" json:"patch,omitempty"`
	Queries        []string        `yaml:"queries,omitempty" json:"queries,omitempty"`
	Get            bool            `yaml:"get,omitempty" json:"get,omitempty"`
	GormMap        GormOverrideMap `yaml:"gormMap,omitempty" json:"gormMap,omitempty"`
	Query          bool            `yaml:"query,omitempty" json:"query,omitempty"`
	Post           bool            `yaml:"post,omitempty" json:"post,omitempty"`
	ImportList     []string        `yaml:"importList,omitempty" json:"importList,omitempty"`
	Fields         []*Module2Field `yaml:"fields,omitempty" json:"fields,omitempty"`
	C              bool            `yaml:"c,omitempty" json:"c,omitempty"`
	CliName        string          `yaml:"cliName,omitempty" json:"cliName,omitempty"`
	CliShort       string          `yaml:"cliShort,omitempty" json:"cliShort,omitempty"`
	CliDescription string          `yaml:"cliDescription,omitempty" json:"cliDescription,omitempty"`
	Cte            bool            `yaml:"cte,omitempty" json:"cte,omitempty"`
	PostFormatter  string          `yaml:"postFormatter,omitempty" json:"postFormatter,omitempty"`
}

// This is the new dto version
type Module2DtoBase struct {
	Name       string          `yaml:"name,omitempty" json:"name,omitempty"`
	ImportList []string        `yaml:"importList,omitempty" json:"importList,omitempty"`
	Fields     []*Module2Field `yaml:"fields,omitempty" json:"fields,omitempty"`
}

type Module2ActionBody struct {
	Fields []*Module2Field `yaml:"fields,omitempty" json:"fields,omitempty"`
	Dto    string          `yaml:"dto,omitempty" json:"dto,omitempty"`
	Entity string          `yaml:"entity,omitempty" json:"entity,omitempty"`
}

type Module2Action struct {
	CliName string `yaml:"cliName,omitempty" json:"cliName,omitempty"`
	Name    string `yaml:"name,omitempty" json:"name,omitempty"`
	Url     string `yaml:"url,omitempty" json:"url,omitempty"`
	Method  string `yaml:"method,omitempty" json:"method,omitempty"`

	Fn          string `yaml:"fn,omitempty" json:"fn,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	Format        string            `yaml:"format,omitempty" json:"format,omitempty"`
	In            Module2ActionBody `yaml:"in,omitempty" json:"in,omitempty"`
	Out           Module2ActionBody `yaml:"out,omitempty" json:"out,omitempty"`
	SecurityModel SecurityModel     `yaml:"security,omitempty" json:"security,omitempty"`

	Virtual        bool
	Handlers       []gin.HandlerFunc `yaml:"-" json:"-"`
	ExternFuncName string            `yaml:"-" json:"-"`
	RequestEntity  any               `yaml:"-" json:"-"`
	ResponseEntity any               `yaml:"-" json:"-"`
	Action         any               `yaml:"-" json:"-"`
	TargetEntity   any               `yaml:"-" json:"-"`
}