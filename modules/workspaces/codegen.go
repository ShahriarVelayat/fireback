/**
* Code generation for fireback projects.
* This is |NOT| golang code generation, rather generates code after compile time for different
* Platforms
 */
package workspaces

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	reflect "reflect"
	"regexp"
	"strings"
	"text/template"

	"github.com/gertd/go-pluralize"
	"github.com/gin-gonic/gin"
	"github.com/swaggest/openapi-go/openapi3"
	firebackgo "github.com/torabian/fireback/modules/workspaces/codegen/firebackgo"
	"gopkg.in/yaml.v2"
)

var FIELD_TYPE_ARRAY string = "array"
var FIELD_TYPE_ARRAYP string = "arrayP"
var FIELD_TYPE_JSON string = "json"
var FIELD_TYPE_ONE string = "one"
var FIELD_TYPE_DATE string = "date"
var FIELD_TYPE_MANY2MANY string = "many2many"
var FIELD_TYPE_OBJECT string = "object"
var FIELD_TYPE_ENUM string = "enum"
var FIELD_TYPE_COMPUTED string = "computed"
var FIELD_TYPE_TEXT string = "text"
var FIELD_TYPE_STRING string = "string"
var FIELD_TYPE_ANY string = "any"
var ROUTE_FORMAT_DELETE string = "DELETE_DSL"
var ROUTE_FORMAT_QUERY string = "QUERY"
var ROUTE_FORMAT_POST string = "POST_ONE"
var ROUTE_FORMAT_GET_ONE string = "GET_ONE"
var ROUTE_FORMAT_REACTIVE string = "REACTIVE"
var ROUTE_FORMAT_PATCH string = "PATCH_ONE"
var ROUTE_FORMAT_PATCH_BULK string = "PATCH_BULK"

// We extend things in the front-end from a base entity, this is not necessary anymore
// But let it be for reference, maybe later we want add something to it (?)
var DefaultFields []*Module2Field = []*Module2Field{
	// {Name: "visibility", Type: "string"},
	// {Name: "parentId", Type: "string"},
	// {Name: "linkerId", Type: "string"},
	// {Name: "workspaceId", Type: "string"},
	// {Name: "linkedId", Type: "string"},
	// {Name: "uniqueId", Type: "string"},
	// {Name: "userId", Type: "string"},
	// // {Name: "rank", Type: "float64"},
	// {Name: "updated", Type: "float64"},
	// {Name: "created", Type: "float64"},
	// {Name: "createdFormatted", Type: "string"},
	// {Name: "updatedFormatted", Type: "string"},
}

func (x *Module2FieldMatch) PublicName() string {
	if x.Dto == nil {
		return ""
	}

	return ToUpper(*x.Dto) + "Dto"
}

func (x *Module2ActionBody) EntityPure() string {
	if x.Entity != "" {
		return strings.ReplaceAll(x.Entity, "Entity", "")
	}

	return ""
}

func (x *Module2Action) DashedName() string {
	return ToSnakeCase(x.Name)
}

func (x *Module2) Upper() string {
	if x.Name == "" {
		return ToUpper(x.Path)
	}
	return ToUpper(x.Name)
}

func (x *Module2Field) ComputedCliName() string {
	return strings.ReplaceAll(ToSnakeCase((x.Name)), "_", "-")
}

func (x *Module2Field) DistinctBy() string {
	return ""
}

func (x *Module2Action) ComputedCliName() string {
	if x.CliName != "" {
		return x.CliName
	}
	return strings.ReplaceAll(ToSnakeCase((x.Name)), "_", "-")
}

func (x *Module2ActionBody) Template() string {
	return "-"
}

func (x *Module2Action) Template() string {
	return "-"
	// return x.DashedName()
}

func (x *Module2Action) ComputeRequestEntity() string {
	if x.In.Entity != "" {
		return "&" + x.In.Entity + "{}"
	}
	if x.In.Dto != "" {
		return "&" + x.In.Dto + "{}"
	}

	if len(x.In.Fields) > 0 {
		return "&" + x.Upper() + "ActionReqDto{}"
	}

	return ""
}

func (x *Module2Action) ComputeRequestEntityS() string {
	if x.In.Entity != "" {
		return x.In.Entity
	}
	if x.In.Dto != "" {
		return x.In.Dto
	}

	if len(x.In.Fields) > 0 {
		return x.Upper() + "ActionReqDto"
	}

	return ""
}

func (x *Module2Action) FormatComputed() string {
	if x.Method == "REACTIVE" || x.Method == "reactive" {
		return "REACTIVE"
	}
	if x.Format != "" {
		return strings.ToUpper(x.Format)
	}

	if x.Method == "get" {
		return "GET_ONE"
	}

	return "POST_ONE"
}

func (x *Module2Action) DashedPluralName() string {

	pluralize2 := pluralize.NewClient()
	return strings.ReplaceAll(ToSnakeCase(pluralize2.Plural(x.Name)), "_", "-")
}
func (x *Module2Action) ComputedUrl() string {

	if x.Url != "" {
		return x.Url
	}

	return "/" + x.DashedPluralName()
}

func (x *Module2Action) MethodAllUpper() string {

	return strings.ToUpper(x.Method)
}

func (x *Module2FieldOf) KeyUpper() string {
	return ToUpper(x.Key)
}

func (x *Module2Action) ComputeResponseEntity() string {
	if x.Out.Entity != "" {
		return "&" + x.Out.Entity + "{}"
	}
	if x.Out.Dto != "" {
		return "&" + x.Out.Dto + "{}"
	}

	if len(x.Out.Fields) > 0 {
		return "&" + x.Upper() + "ActionResDto{}"
	}

	return "&OkayResponseDto{}"
}
func (x *Module2Action) ComputeResponseEntityS() string {
	if x.Out.Entity != "" {
		return x.Out.Entity
	}
	if x.Out.Dto != "" {
		return x.Out.Dto
	}

	if len(x.Out.Fields) > 0 {
		return x.Upper() + "ActionResDto"
	}

	return "OkayResponseDto"
}

type TypeScriptGenContext struct {
	IncludeStaticField      bool
	IncludeFirebackDef      bool
	IncludeStaticNavigation bool
}

type CodeGenContext struct {
	// Where the content will be exported to
	Path string

	EntityPath string

	// Type of the generation, (swift, etc)
	Type string

	// Generation
	OpenApiFile string

	// Generation
	GofModuleName string

	// Only build specific modules
	Modules []string

	// Only build specific modules
	ModulesOnDisk []string

	// Set of functions, and meta data to generate for that specific target
	Catalog CodeGenCatalog

	NoCache bool

	Ts TypeScriptGenContext
}

// Used in go structs to configurate the gorm tag
func (x *Module2Field) ComputedExcerptSize() int {
	if x.ExcerptSize == 0 {
		return 100
	}

	return x.ExcerptSize
}

func (x *Module2Field) ComputedGormTag() string {
	if x.Gorm != "" {
		return x.Gorm
	}

	if x.Type == FIELD_TYPE_TEXT {
		return "text"
	}

	if x.Type == FIELD_TYPE_MANY2MANY {
		return "many2many:" + x.BelongingEntityName + "_" + x.PrivateName() + ";foreignKey:UniqueId;references:UniqueId"
	}

	if x.Type == FIELD_TYPE_ARRAY || x.Type == FIELD_TYPE_OBJECT {
		return "foreignKey:LinkerId;references:UniqueId;constraint:OnDelete:CASCADE"
	}

	if x.Type == FIELD_TYPE_ONE {
		return "foreignKey:" + x.PublicName() + "Id;references:UniqueId"
	}

	if x.Type == FIELD_TYPE_ANY {
		return "-"
	}

	return ""
}

func (x *Module2Field) ComputedSqlTag() string {
	if x.Sql != "" {
		return x.Sql
	}

	if x.Type == FIELD_TYPE_COMPUTED || x.Type == FIELD_TYPE_ANY {
		return "-"
	}

	return ""
}
func (x *Module2Field) PrivateNameUnderscore() string {
	return "_" + x.Name
}
func (x *Module2Field) UpperPlural() string {
	pluralize2 := pluralize.NewClient()
	return ToUpper(pluralize2.Plural(x.Name))
}

func CalcAllPolyglotEntities(m []*Module2Field) []string {
	items := []string{}
	for _, item := range m {
		if item.Translate {
			items = append(items, item.Name)
		}

		if item.Type == FIELD_TYPE_OBJECT || item.Type == FIELD_TYPE_ARRAY {
			items = append(items, CalcAllPolyglotEntities(item.Fields)...)
		}
	}

	return items
}

func (x *Module2Entity) CompletePolyglotFields() []string {
	return CalcAllPolyglotEntities(x.Fields)
}

func (x *Module2Entity) DistinctByAllUpper() string {
	return strings.ToUpper(x.DistinctBy)
}

func (x *Module2Entity) DistinctByAllLower() string {
	return strings.ToLower(x.DistinctBy)
}

func (x *Module2Entity) ComputedCliName() string {
	if x.CliName != "" {
		return x.CliName
	}
	return x.Name
}

func (x *Module2Entity) ComputedCliDescription() string {
	if x.CliDescription != "" {
		return x.CliDescription
	}
	return ""

}

func (x *Module2Field) ComputedCliDescription() string {

	if x.Type == FIELD_TYPE_ENUM {
		items := []string{}
		for _, item := range x.OfType {
			items = append(items, "'"+item.Key+"'")
		}
		return "One of: " + strings.Join(items, ", ")
	}

	if x.Description != "" {
		return x.Description
	}

	return x.Name
}

func (x *Module2Field) CompleteFields() []*Module2Field {
	return x.Fields
}

func (x *Module2Field) IsRequired() bool {
	return strings.Contains(x.Validate, "required")
}

func (x *Module2Field) PrivateName() string {
	return x.Name
}

func (x *Module2) PublicName() string {
	return ToUpper(x.Path)
}
func (x *Module2Field) PublicName() string {
	return ToUpper(x.Name)
}
func (x *Module2Field) AllUpper() string {
	return strings.ToUpper(CamelCaseToWordsUnderlined(x.Name))
}
func (x *Module2Field) UnderscoreName() string {
	return strings.ToLower(CamelCaseToWordsUnderlined(x.Name))
}
func (x *Module2Entity) UnderscoreName() string {
	return strings.ToLower(CamelCaseToWordsUnderlined(x.Name))
}
func (x *Module2Field) TargetWithModule() string {
	if x.Module != "" {
		return x.Module + "." + ToUpper(x.Target)
	}
	return ToUpper(x.Target)
}
func (x *Module2Field) TargetWithoutEntity() string {
	return strings.ReplaceAll(x.Target, "Entity", "")
}
func (x *Module2Field) TargetWithoutEntityPlural() string {
	pluralize2 := pluralize.NewClient()
	return ToUpper(pluralize2.Plural(x.TargetWithoutEntity()))
}
func (x *Module2Field) TargetWithModuleWithoutEntity() string {
	return strings.ReplaceAll(x.TargetWithModule(), "Entity", "")
}
func (x *Module2Action) Upper() string {
	return ToUpper(x.Name)
}

func (x *Module2Action) ActionReqDto() string {

	if x.In.Entity != "" {
		return "*" + x.In.Entity
	}
	if x.In.Dto != "" {
		return "*" + x.In.Dto
	}

	if len(x.In.Fields) > 0 {
		return "*" + x.Upper() + "ActionReqDto"
	}

	return "nil"
}

func (x *Module2Action) ActionResDto() string {
	prefix := ""
	if x.Format == "QUERY" {
		prefix = "[]"
	}
	if x.Out.Entity != "" {
		return prefix + "*" + x.Out.Entity
	}
	if x.Out.Dto != "" {
		return prefix + "*" + x.Out.Dto
	}

	if len(x.Out.Fields) > 0 {
		return prefix + "*" + x.Upper() + "ActionResDto"
	}

	return "nil"
}

func (x *Module2Field) DashedName() string {
	return ToSnakeCase(x.Name)
}
func (x *Module2Field) DefaultEmptySymbol() string {
	switch x.Type {
	case "string", "text":
		return `""`
	case "one":
		return x.Target + "?"
	case "array", "many2many":
		return "nil"
	case "int64", "int32", "int":
		return "0"
	case "float64", "float32", "float":
		return "0.0"
	case "bool":
		return "FALSE"
	case "Timestamp", "datenano":
		return `""`
	case "double":
		return "0.0"
	default:
		return "nil"
	}

}
func (x *Module2Entity) CompleteFields() []*Module2Field {
	var all []*Module2Field = []*Module2Field{}
	all = append(all,
		x.Fields...,
	)
	all = append(all,
		DefaultFields...,
	)

	return all
}

func (x *Module2DtoBase) CompleteFields() []*Module2Field {
	var all []*Module2Field = []*Module2Field{}
	all = append(all,
		x.Fields...,
	)
	all = append(all,
		DefaultFields...,
	)

	return all
}

// Represents things which are needed to be imported
type ImportDependencyStrategy struct {
	Items []string
	Path  string
}

type ImportMapRow struct {
	Items []string
}
type ImportMap map[string]*ImportMapRow

func mergeImportMaps(maps ...ImportMap) ImportMap {
	result := make(ImportMap)

	for _, m := range maps {
		for key, row := range m {
			if _, ok := result[key]; !ok {
				result[key] = &ImportMapRow{}
			}
			// Merge items and make them unique
			itemSet := make(map[string]struct{})
			for _, item := range append(result[key].Items, row.Items...) {
				itemSet[item] = struct{}{}
			}
			result[key].Items = make([]string, 0, len(itemSet))
			for item := range itemSet {
				result[key].Items = append(result[key].Items, item)
			}
		}
	}

	return result
}

var generatorHash map[string]string = map[string]string{}

func GetMD5Hash(text []byte) string {
	hash := md5.Sum(text)
	return hex.EncodeToString(hash[:])
}

func WriteFileGen(ctx *CodeGenContext, name string, data []byte, perm os.FileMode) error {

	gen, okay := generatorHash[name]

	newGen := GetMD5Hash(data)
	if okay && gen == newGen {
		return nil
	}

	generatorHash[name] = newGen

	err := os.WriteFile(name, data, perm)

	if err != nil {
		return err
	}

	if ctx.Catalog.Prettier != nil {
		ctx.Catalog.Prettier(name)
	}

	return nil
}

type ReconfigDto struct {
	ProjectSource  string
	NewProjectName string
	BinaryName     string
	Description    string
	Languages      []string
}

/**
* Renames the project into another project, useful for fast bootstrap
**/
func Reconfig(scheme ReconfigDto) error {

	// Change the make file
	{
		data, _ := ioutil.ReadFile("Makefile")
		d := string(data)
		d = strings.ReplaceAll(d, scheme.ProjectSource, scheme.NewProjectName)
		os.WriteFile("Makefile", []byte(d), 0644)
	}
	{
		data, _ := ioutil.ReadFile(filepath.Join(".vscode", "tasks.json"))
		d := string(data)
		d = strings.ReplaceAll(d, ` f `, scheme.BinaryName+` `)
		os.WriteFile(filepath.Join(".vscode", "tasks.json"), []byte(d), 0644)
	}
	{
		data, _ := ioutil.ReadFile(filepath.Join("cmd", "fireback", "main.go"))
		d := string(data)
		d = strings.ReplaceAll(d, `var PRODUCT_NAMESPACENAME = "fireback"`, `var PRODUCT_NAMESPACENAME = "`+scheme.NewProjectName+`"`)
		d = strings.ReplaceAll(d, `var PRODUCT_DESCRIPTION = "Fireback core microservice"`, `var PRODUCT_DESCRIPTION = "`+scheme.Description+`"`)
		os.WriteFile(filepath.Join("cmd", "fireback", "main.go"), []byte(d), 0644)
	}

	{
		data, _ := ioutil.ReadFile(filepath.Join("cmd", "fireback", "Makefile"))
		d := string(data)
		d = strings.ReplaceAll(d, "project = fireback", "project = "+scheme.NewProjectName)
		d = strings.ReplaceAll(d, "projectBinary = f", "projectBinary = "+scheme.BinaryName)
		os.WriteFile(filepath.Join("cmd", "fireback", "Makefile"), []byte(d), 0644)
	}

	{
		err := os.Rename(
			filepath.Join("cmd", "fireback"),
			filepath.Join("cmd", scheme.NewProjectName+"-server"),
		)
		if err != nil {
			return err
		}
	}

	return nil

}

func GenerateRpcCodeOnDisk(ctx *CodeGenContext, route Module2Action, exportDir string) {

	content, exportPath := GenerateRpcCodeString(ctx, route, exportDir)

	if exportPath != "" {

		err3 := WriteFileGen(ctx, exportPath, content, 0644)
		if err3 != nil {
			fmt.Println("Error on writing content on RPC:", exportPath, err3)
		}
	}
}

func GenerateRpcCodeString(ctx *CodeGenContext, route Module2Action, exportDir string) ([]byte, string) {
	method := strings.ToUpper(route.Method)
	if (route.Format == ROUTE_FORMAT_POST || method == "POST") && ctx.Catalog.RpcPost != "" {
		data, err := route.RenderTemplate(ctx, ctx.Catalog.Templates, ctx.Catalog.RpcPost)
		if err != nil {
			log.Fatalln("Generating post call error", err)
			return []byte(""), ""
		}
		exportPath := filepath.Join(exportDir, ctx.Catalog.RpcPostDiskName(&route))
		return EscapeLines(data), exportPath
	}

	if route.Format == ROUTE_FORMAT_QUERY && ctx.Catalog.RpcQuery != "" {
		data, err := route.RenderTemplate(ctx, ctx.Catalog.Templates, ctx.Catalog.RpcQuery)
		if err != nil {
			log.Fatalln("Generating rpc query call error", err)
			return []byte(""), ""
		}
		exportPath := filepath.Join(exportDir, ctx.Catalog.RpcQueryDiskName(&route))
		return EscapeLines(data), exportPath
	}
	if (route.Format == ROUTE_FORMAT_DELETE || method == "DELETE") && ctx.Catalog.RpcDelete != "" {
		data, err := route.RenderTemplate(ctx, ctx.Catalog.Templates, ctx.Catalog.RpcDelete)
		if err != nil {
			log.Fatalln("Generating delete rpc call error", err)
			return []byte(""), ""
		}
		exportPath := filepath.Join(exportDir, ctx.Catalog.RpcDeleteDiskName(&route))
		return EscapeLines(data), exportPath
	}
	if (route.Format == ROUTE_FORMAT_PATCH || method == "PATCH") && ctx.Catalog.RpcPatch != "" {
		data, err := route.RenderTemplate(ctx, ctx.Catalog.Templates, ctx.Catalog.RpcPatch)
		if err != nil {
			log.Fatalln("Generating rpc patch call error", err)
			return []byte(""), ""
		}
		exportPath := filepath.Join(exportDir, ctx.Catalog.RpcPatchDiskName(&route))
		return EscapeLines(data), exportPath
	}
	if route.Format == ROUTE_FORMAT_PATCH_BULK && ctx.Catalog.RpcPatchBulk != "" {
		data, err := route.RenderTemplate(ctx, ctx.Catalog.Templates, ctx.Catalog.RpcPatchBulk)
		if err != nil {
			log.Fatalln("Generating rpc patch call error", err)
			return []byte(""), ""
		}
		exportPath := filepath.Join(exportDir, ctx.Catalog.RpcPatchBulkDiskName(&route))
		return EscapeLines(data), exportPath
	}
	if (route.Format == ROUTE_FORMAT_REACTIVE || method == ROUTE_FORMAT_REACTIVE) && ctx.Catalog.RpcReactive != "" {
		data, err := route.RenderTemplate(ctx, ctx.Catalog.Templates, ctx.Catalog.RpcReactive)
		if err != nil {
			log.Fatalln("Generating rpc reactive call error", err)
			return []byte(""), ""
		}
		exportPath := filepath.Join(exportDir, ctx.Catalog.RpcReactiveDiskName(&route))
		return EscapeLines(data), exportPath
	}
	if route.Format == ROUTE_FORMAT_GET_ONE && ctx.Catalog.RpcGetOne != "" {
		data, err := route.RenderTemplate(ctx, ctx.Catalog.Templates, ctx.Catalog.RpcGetOne)
		if err != nil {
			log.Fatalln("Generating rpc get one call error", err)
			return []byte(""), ""
		}
		exportPath := filepath.Join(exportDir, ctx.Catalog.RpcGetOneDiskName(&route))
		return EscapeLines(data), exportPath
	}
	return []byte(""), ""
}

func GenMoveIncludeDir(ctx *CodeGenContext) {

	// Move the include directory
	if ctx.Catalog.IncludeDirectory != nil {
		files, err := GetAllFilenames(ctx.Catalog.IncludeDirectory, "")
		if err != nil {
			log.Fatalln(err)
		}

		for _, file := range files {
			exportFile := filepath.Join(ctx.Path, file)
			content, err4 := ReadEmbedFileContent(ctx.Catalog.IncludeDirectory, file)
			if err4 != nil {
				log.Fatalln(err)
			}
			dir := filepath.Dir(exportFile)
			os.MkdirAll(dir, os.ModePerm)
			err3 := WriteFileGen(ctx, exportFile, []byte(content), 0644)
			if err3 != nil {
				log.Fatalln(err)
			}
		}
	}

}

func GenGetModules(xapp *XWebServer, ctx *CodeGenContext) []*Module2 {

	j := []*Module2{}
	if len(ctx.ModulesOnDisk) > 0 && ctx.ModulesOnDisk[0] != "" {
		j = append(j, ListModule2FilesFromDisk(ctx.ModulesOnDisk)...)
	} else {
		j = append(j, ListModule2Files(xapp)...)
	}

	return j
}

func writeGenCache(ctx *CodeGenContext) {
	cacheMapFile := filepath.Join(ctx.Path, "cachemap.json")

	cache, _ := json.MarshalIndent(generatorHash, "", "  ")
	os.WriteFile(cacheMapFile, cache, 0644)
}

func ReadGenCache(ctx *CodeGenContext) {
	cacheMapFile := filepath.Join(ctx.Path, "cachemap.json")
	if !ctx.NoCache {
		ReadJsonFile(cacheMapFile, &generatorHash)

	}
}

func GenRpcCode(ctx *CodeGenContext, modules []*ModuleProvider, mode string) {

	for _, item := range modules {
		m := item.ToModule2()
		exportDir := filepath.Join(ctx.Path, "modules", item.Name)

		perr := os.MkdirAll(exportDir, os.ModePerm)
		if perr != nil {
			log.Fatalln(perr)
		}

		actions := item.Actions

		for _, bundle := range item.EntityBundles {
			actions = append(actions, bundle.Actions)
		}

		for _, actions := range actions {
			if mode == "disk" {
				for _, action := range actions {
					action.RootModule = &m
					GenerateRpcCodeOnDisk(ctx, action, exportDir)
				}
			}

			if mode == "class" {
				if len(actions) > 0 {

					importMap := ImportMap{}
					fileToWrite := filepath.Join(exportDir, ctx.Catalog.EntityClassDiskName(&actions[0]))
					content := []byte("")

					for _, action := range actions {

						action.RootModule = &m
						maps := actions[0].ImportDependecies()
						importMap = mergeImportMaps(importMap, maps)
						partial, _ := GenerateRpcCodeString(ctx, action, exportDir)

						content = append(content, []byte("\r\n")...)
						content = append(content, EscapeLines(partial)...)
						content = append(content, []byte("\r\n")...)
					}

					render, err2 := RenderRpcGroupClassBody(
						ctx,
						ctx.Catalog.Templates,
						ctx.Catalog.EntityClassTemplate,
						content,
						actions[0].Group,
						importMap,
					)
					if err2 != nil {
						fmt.Println("Error on rendering the content", err2)
					}

					err3 := WriteFileGen(ctx, fileToWrite, (render), 0644)
					if err3 != nil {
						fmt.Println("Error on writing content for Actions class file:", fileToWrite, err3)
					}
				}

			}
		}
	}
}

func GenRpcCodeExternal(ctx *CodeGenContext, modules []*Module2, mode string) {

	for _, item := range modules {

		exportDir := filepath.Join(ctx.Path, "modules", item.Name)
		perr := os.MkdirAll(exportDir, os.ModePerm)
		if perr != nil {
			log.Fatalln(perr)
		}

		if mode == "disk" {
			for _, action := range item.Actions {
				action.RootModule = item
				GenerateRpcCodeOnDisk(ctx, action, exportDir)
			}
		}

		if mode == "class" {
			if len(item.Actions) > 0 {

				importMap := ImportMap{}
				fileToWrite := filepath.Join(exportDir, ctx.Catalog.EntityClassDiskName(&item.Actions[0]))
				content := []byte("")

				for _, action := range item.Actions {
					action.RootModule = item
					maps := action.ImportDependecies()
					importMap = mergeImportMaps(importMap, maps)
					partial, _ := GenerateRpcCodeString(ctx, action, exportDir)

					content = append(content, []byte("\r\n")...)
					content = append(content, EscapeLines(partial)...)
					content = append(content, []byte("\r\n")...)
				}

				render, err2 := RenderRpcGroupClassBody(
					ctx,
					ctx.Catalog.Templates,
					ctx.Catalog.EntityClassTemplate,
					content,
					item.Actions[0].Group,
					importMap,
				)
				if err2 != nil {
					fmt.Println("Error on rendering the content", err2)
				}

				err3 := WriteFileGen(ctx, fileToWrite, (render), 0644)
				if err3 != nil {
					fmt.Println("Error on writing content for Actions class file:", fileToWrite, err3)
				}
			}

		}
	}
}

// For openapi3, we create xwebserver not from internal, rather an external json
func GetOpenAPiXServer(ctx *CodeGenContext) (*XWebServer, []*Module2) {
	data, _ := ioutil.ReadFile(ctx.OpenApiFile)
	s := openapi3.Spec{}

	if err := s.UnmarshalJSON(data); err != nil {
		log.Fatal("Converting json content:", err)
	}

	virtualModule := OpenApiToFireback(s)
	modules := []*Module2{
		virtualModule,
	}
	app := &XWebServer{
		Modules: []*ModuleProvider{
			{
				Actions: [][]Module2Action{
					virtualModule.Actions,
				},
			},
		},
	}
	return app, modules
}

func extractModuleName(goModFilePath string) (string, error) {
	file, err := os.Open(goModFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				return parts[1], nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("module name not found in %s", goModFilePath)
}

/*
* Creates a yml, and also a golang module file in modules directory
 */
func NewGoNativeModule(name string, dist string, autoImport string) error {

	folderName := strings.ToLower(dist)
	args := gin.H{
		"path": folderName,
		"Name": ToUpper(folderName),
		"name": folderName,
	}
	goModule, err := CompileString(&firebackgo.FbGoTpl, "GoModule.tpl", args)
	if err != nil {
		return err
	}

	goModuleDef, err := CompileString(&firebackgo.FbGoTpl, "GoModuleDef.tpl", args)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join("modules", folderName), os.ModePerm); err != nil {
		return err
	}

	moduleName := filepath.Join("modules", folderName, ToUpper(name)+"Module.go")

	if err := os.WriteFile(moduleName, []byte(goModule), 0644); err != nil {
		return err
	}

	yamlName := filepath.Join("modules", folderName, ToUpper(name)+"Module3.yml")
	if err := os.WriteFile(yamlName, []byte(goModuleDef), 0644); err != nil {
		return err
	}

	MAGIC_LINE := "// do not remove this comment line - it's used by fireback to append new modules"

	if autoImport != "" {

		if !Exists("go.mod") {
			return nil
		}

		moduleName, err0 := extractModuleName("go.mod")
		if err0 != nil {
			return nil
		}

		fmt.Println(autoImport)
		if data, err := os.ReadFile(autoImport); err != nil {
			return err
		} else {
			j := string(data)
			m := strings.ReplaceAll(
				j,
				MAGIC_LINE,
				MAGIC_LINE+"\r\n\t\t"+name+"."+ToUpper(name)+"ModuleSetup(),",
			)

			m = strings.ReplaceAll(
				m,
				"import (",
				"import ("+"\r\n\t\""+moduleName+"/modules/"+name+"\"\r\n",
			)
			os.WriteFile(autoImport, []byte(m), 0644)
		}
	}

	return nil
}

func CompileString(fs *embed.FS, fname string, params gin.H) (string, error) {
	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer

	err = t.ExecuteTemplate(&tpl, fname, params)

	return tpl.String(), err
}

func RunCodeGen(xapp *XWebServer, ctx *CodeGenContext) error {
	// For now, I have separate work flow for binary and external definitions
	// Which might need to be fixed.

	if len(ctx.ModulesOnDisk) > 0 || ctx.OpenApiFile != "" {
		return RunCodeGenExternal(ctx)
	}

	os.MkdirAll(ctx.Path, os.ModePerm)
	ReadGenCache(ctx)
	GenMoveIncludeDir(ctx)

	app := xapp
	modules := GenGetModules(xapp, ctx)

	// Generate the classes, definitions, structs
	for _, item := range modules {
		item.Generate(ctx)
	}

	mode := "disk"
	if ctx.Catalog.EntityClassTemplate != "" {
		mode = "class"
	}
	GenRpcCode(ctx, app.Modules, mode)

	writeGenCache(ctx)

	return nil
}

// This is used for pure items from definition, not internal binary
func RunCodeGenExternal(ctx *CodeGenContext) error {

	os.MkdirAll(ctx.Path, os.ModePerm)
	ReadGenCache(ctx)
	GenMoveIncludeDir(ctx)

	modules := []*Module2{}

	for _, def := range ctx.ModulesOnDisk {
		var m Module2
		ReadYamlFile[Module2](def, &m)
		modules = append(modules, &m)
	}

	if ctx.OpenApiFile != "" {
		fmt.Println("Catched open api :D ")
		_, modules = GetOpenAPiXServer(ctx)
	}

	// Generate the classes, definitions, structs
	for _, item := range modules {
		// if len(ctx.Modules) > 0 && !Contains(ctx.Modules, item.Name) {
		// 	continue
		// }
		item.Generate(ctx)
	}

	mode := "disk"
	if ctx.Catalog.EntityClassTemplate != "" {
		mode = "class"
	}
	GenRpcCodeExternal(ctx, modules, mode)

	writeGenCache(ctx)

	return nil
}

func FindModuleByPath(xapp *XWebServer, modulePath string) *Module2 {
	for _, item := range xapp.Modules {
		if item.Definitions == nil {
			continue
		}

		if defFile, err := GetSeederFilenames(item.Definitions, ""); err != nil {
			fmt.Println(err.Error())
		} else {

			for _, path := range defFile {
				var mod2 Module2
				ReadYamlFileEmbed(item.Definitions, path, &mod2)
				ComputeMacros(&mod2)
				if mod2.Path == modulePath {
					return &mod2
				}
			}
		}

	}

	return nil
}

func ListModule2WithEntities(xapp *XWebServer) []string {
	items := []string{}
	for _, item := range xapp.Modules {
		if item.Definitions == nil {
			continue
		}

		if defFile, err := GetSeederFilenames(item.Definitions, ""); err != nil {
			fmt.Println(err.Error())
		} else {

			for _, path := range defFile {
				var mod2 Module2
				ReadYamlFileEmbed(item.Definitions, path, &mod2)
				ComputeMacros(&mod2)
				for _, entity := range mod2.Entities {
					items = append(items, mod2.Path+"."+entity.Name)
				}
			}
		}
	}

	return items
}

func FindModule2Entity(xapp *XWebServer, address string) *Module2Entity {
	for _, item := range xapp.Modules {
		if item.Definitions == nil {
			continue
		}

		if defFile, err := GetSeederFilenames(item.Definitions, ""); err != nil {
			fmt.Println(err.Error())
		} else {

			for _, path := range defFile {
				var mod2 Module2
				ReadYamlFileEmbed(item.Definitions, path, &mod2)
				ComputeMacros(&mod2)
				for _, entity := range mod2.Entities {
					if mod2.Path+"."+entity.Name == address {
						return &entity
					}
				}
			}
		}
	}

	return nil
}

func ListModule2Files(xapp *XWebServer) []*Module2 {
	items := []*Module2{}
	for _, item := range xapp.Modules {
		if item.Definitions == nil {
			continue
		}

		if defFile, err := GetSeederFilenames(item.Definitions, ""); err != nil {
			fmt.Println(err.Error())
		} else {

			for _, path := range defFile {
				var mod2 Module2
				ReadYamlFileEmbed(item.Definitions, path, &mod2)
				ComputeMacros(&mod2)
				items = append(items, &mod2)
			}
		}

	}

	return items
}

func ListModule2Entities(xapp *XWebServer, modulePath string) []Module2Entity {
	module := FindModuleByPath(xapp, modulePath)
	if module == nil {
		return []Module2Entity{}
	}
	return module.Entities
}

func ListModule2FilesFromDisk(files []string) []*Module2 {
	items := []*Module2{}

	for _, item := range files {
		var mod2 Module2
		ReadYamlFile(item, &mod2)
		ComputeMacros(&mod2)
		items = append(items, &mod2)
	}

	return items
}

func ComputeComplexGormField(entity *Module2Entity, fields []*Module2Field) {
	if len(fields) == 0 {
		return
	}

	for _, field := range fields {
		field.BelongingEntityName = entity.Name

		if field.Type == FIELD_TYPE_OBJECT || field.Type == FIELD_TYPE_ARRAY {
			ComputeComplexGormField(entity, field.Fields)
		}
	}
}

func ComputeFieldTypes(fields []*Module2Field, isWorkspace bool, fn func(field *Module2Field, isWorkspace bool) string,
) {
	if len(fields) == 0 {
		return
	}

	for _, field := range fields {
		field.ComputedType = fn(field, isWorkspace)

		if field.Type == FIELD_TYPE_OBJECT || field.Type == FIELD_TYPE_ARRAY {
			ComputeFieldTypes(field.Fields, isWorkspace, fn)
		}
	}
}

type CodeGenCatalog struct {
	LanguageName            string
	ComputeField            func(field *Module2Field, isWorkspace bool) string
	EntityDiskName          func(x *Module2Entity) string
	EntityExtensionDiskName func(x *Module2Entity) string
	EntityClassDiskName     func(x *Module2Action) string

	// A function that does a post formatting on the file, when it's saved
	Prettier func(string)

	// Maybe only useful for C/C++
	EntityHeaderDiskName          func(x *Module2Entity) string
	EntityHeaderExtensionDiskName func(x *Module2Entity) string

	ActionDiskName func(modulename string) string

	// When you want each action to be written in separate file
	SingleActionDiskName             func(action *Module2Action, modulename string) string
	DtoDiskName                      func(x *Module2DtoBase) string
	FormDiskName                     func(x *Module2Entity) string
	RpcQueryDiskName                 func(x *Module2Action) string
	RpcDeleteDiskName                func(x *Module2Action) string
	RpcGetOneDiskName                func(x *Module2Action) string
	RpcPatchBulkDiskName             func(x *Module2Action) string
	RpcReactiveDiskName              func(x *Module2Action) string
	RpcPatchDiskName                 func(x *Module2Action) string
	RpcPostDiskName                  func(x *Module2Action) string
	Templates                        embed.FS
	IncludeDirectory                 *embed.FS
	Partials                         *embed.FS
	EntityClassTemplate              string
	EntityGeneratorTemplate          string
	EntityExtensionGeneratorTemplate string
	DtoGeneratorTemplate             string
	CteSqlTemplate                   string
	ActionGeneratorTemplate          string
	FormGeneratorTemplate            string
	RpcQuery                         string
	RpcPatch                         string
	RpcPatchBulk                     string
	RpcGetOne                        string
	RpcReactive                      string
	RpcPost                          string
	RpcDelete                        string
}

func EscapeLines(data []byte) []byte {
	d := string(data)

	re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`)

	return []byte(re.ReplaceAllString(d, ""))

}

func ComputeMacros(x *Module2) {
	for _, item := range x.Macros {
		if item.Using == "eav" {
			EavMacro(item, x)
		}
	}
}

/**
*	Common code generator
**/
func (x *Module2) Generate(ctx *CodeGenContext) {

	isWorkspace := false
	if x.Path == "workspaces" {
		isWorkspace = true
	}
	os.MkdirAll(ctx.Path, os.ModePerm)
	exportDir := filepath.Join(ctx.Path, "modules", x.Path)

	perr := os.MkdirAll(exportDir, os.ModePerm)
	if perr != nil {
		log.Fatalln(perr)
	}

	ComputeMacros(x)

	if ctx.Catalog.LanguageName == "FirebackGo" {
		exportPath := filepath.Join(exportDir, ToUpper(x.Name)+"Module.dyno.go")

		data, err := x.RenderTemplate(ctx, ctx.Catalog.Templates, "GoModuleDyno.tpl")
		if err != nil {
			fmt.Println("Error on module dyno file:", exportPath, err)
		}

		err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
		if err3 != nil {
			fmt.Println("Error on writing content:", exportPath, err3)
		}

	}

	for _, dto := range x.Dto {

		// Computing field types is important for target writter.
		ComputeFieldTypes(dto.CompleteFields(), isWorkspace, ctx.Catalog.ComputeField)

		// Step 0: Generate the Entity
		if ctx.Catalog.DtoGeneratorTemplate != "" {
			exportPath := filepath.Join(exportDir, ctx.Catalog.DtoDiskName(&dto))

			data, err := dto.RenderTemplate(
				ctx,
				ctx.Catalog.Templates,
				ctx.Catalog.DtoGeneratorTemplate,
				x,
			)
			if err != nil {
				fmt.Println("Error on dto generation:", err)
			} else {
				err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
				if err3 != nil {
					fmt.Println("Error on writing content:", exportPath, err3)
				}
			}
		}

	}

	// Render actions specific dtos if they have their own
	if ctx.Catalog.ActionDiskName != nil {
		exportPath := filepath.Join(exportDir, ctx.Catalog.ActionDiskName(x.Path))

		if len(x.Actions) > 0 {

			data, err := x.RenderActions(
				ctx,
				ctx.Catalog.Templates,
				ctx.Catalog.ActionGeneratorTemplate,
			)

			if err != nil {
				fmt.Println("Error on action generation:", err)
			} else {
				err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
				if err3 != nil {
					fmt.Println("Error on writing content:", exportPath, err3)
				}
			}
		}
	}

	if ctx.Catalog.SingleActionDiskName != nil {
		if len(x.Actions) > 0 {

			for _, action := range x.Actions {
				exportPath := filepath.Join(exportDir, ctx.Catalog.SingleActionDiskName(&action, x.Path))

				data, err := action.Render(
					x,
					ctx,
					ctx.Catalog.Templates,
					ctx.Catalog.ActionGeneratorTemplate,
				)
				if err != nil {
					fmt.Println("Error on action generation:", err)
				} else {
					err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
					if err3 != nil {
						fmt.Println("Error on writing content:", exportPath, err3)
					}
				}
			}
		}
	}

	for _, entity := range x.Entities {

		// Computing field types is important for target writter.
		ComputeFieldTypes(entity.CompleteFields(), isWorkspace, ctx.Catalog.ComputeField)
		ComputeComplexGormField(&entity, entity.CompleteFields())

		entityAddress := filepath.Join(exportDir, ctx.Catalog.EntityDiskName(&entity))

		if ctx.Catalog.EntityHeaderDiskName != nil {
			exportPath := filepath.Join(exportDir, ctx.Catalog.EntityHeaderDiskName(&entity))

			params := gin.H{
				"implementation": "skip",
			}

			data, err := entity.RenderTemplate(
				ctx,
				ctx.Catalog.Templates,
				ctx.Catalog.EntityGeneratorTemplate,
				x,
				params,
			)

			if err != nil {
				fmt.Println("Error on entity extension generation:", err)
			} else {
				err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
				if err3 != nil {
					fmt.Println("Error on writing content:", exportPath, err3)
				}
			}
		}

		if ctx.Catalog.EntityExtensionGeneratorTemplate != "" {
			exportPath := filepath.Join(exportDir, ctx.Catalog.EntityExtensionDiskName(&entity))

			// We only render the extension, if this entity is first time being created
			if !Exists(exportPath) && !Exists(entityAddress) {
				data, err := entity.RenderTemplate(
					ctx,
					ctx.Catalog.Templates,
					ctx.Catalog.EntityExtensionGeneratorTemplate,
					x,
					nil,
				)
				if err != nil {
					fmt.Println("Error on entity extension generation:", err)
				} else {
					err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
					if err3 != nil {
						fmt.Println("Error on writing content:", exportPath, err3)
					}
				}

			}

		}

		// Step 0: Generate the Entity
		if ctx.Catalog.EntityGeneratorTemplate != "" {

			data, err := entity.RenderTemplate(
				ctx,
				ctx.Catalog.Templates,
				ctx.Catalog.EntityGeneratorTemplate,
				x,
				nil,
			)
			if err != nil {
				fmt.Println("Error on entity generation:", err)
			} else {
				err3 := WriteFileGen(ctx, entityAddress, EscapeLines(data), 0644)
				if err3 != nil {
					fmt.Println("Error on writing content:", entityAddress, err3)
				}
			}
		}

		// Step 0: Cte SQL Render
		if entity.Cte && ctx.Catalog.CteSqlTemplate != "" {

			{

				os.MkdirAll(filepath.Join(exportDir, "queries"), os.ModePerm)
				exportPath := filepath.Join(exportDir, "queries", entity.Upper()+"Cte.vsql")
				data, err := entity.RenderCteSqlTemplate(
					ctx,
					ctx.Catalog.Templates,
					ctx.Catalog.CteSqlTemplate,
					x,
					"sql",
				)
				if err != nil {
					fmt.Println("Error on cte sql generation:", err)
				} else {
					err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
					if err3 != nil {
						fmt.Println("Error on writing content:", exportPath, err3)
					}
				}
			}
		}

		// Step 1: Generate the form, (fields, sections, inputs, etc...)
		if ctx.Catalog.FormGeneratorTemplate != "" {
			exportPath := filepath.Join(exportDir, ctx.Catalog.FormDiskName(&entity))

			data, err := entity.RenderTemplate(
				ctx,
				ctx.Catalog.Templates,
				ctx.Catalog.FormGeneratorTemplate,
				x,
				nil,
			)
			if err != nil {
				fmt.Println("Error on UI generation:", err)
			} else {
				err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
				if err3 != nil {
					fmt.Println("Error on writing content:", exportPath, err3)
				}
			}
		}

	}
}

func ToUpper(t string) string {
	if t == "" {
		return ""
	}
	return strings.ToUpper(t[0:1]) + t[1:]
}
func ToLower(t string) string {
	if t == "" {
		return ""
	}
	return strings.ToLower(t[0:1]) + t[1:]
}

func (x *Module2Entity) EntityName() string {
	return ToUpper(x.Name) + "Entity"
}

func (x *Module2Entity) ObjectName() string {
	return x.EntityName()
}

func (x *Module2DtoBase) ObjectName() string {
	return x.DtoName()
}

func (x *Module2Entity) HasExtendedQuer() bool {
	return len(x.Queries) > 0 && Contains(x.Queries, "extended")
}

func (x *Module2Entity) EventCreated() string {
	return x.AllUpper() + "_EVENT_CREATED"
}

func (x *Module2Entity) EventUpdated() string {
	return x.AllUpper() + "_EVENT_UPDATED"
}

func (x *Module2Entity) AllUpper() string {
	return strings.ToUpper(CamelCaseToWordsUnderlined(x.Name))
}

func (x *Module2Entity) AllLower() string {
	return strings.ToLower(CamelCaseToWordsDashed(x.Name))
}

func (x *Module2Entity) HumanReadable() string {
	return strings.ToLower(CamelCaseToWords(x.Name))
}

func (x *Module2) AllUpper() string {
	return strings.ToUpper(CamelCaseToWordsUnderlined(x.Name))
}

func (x *Module2) AllLower() string {
	return strings.ToLower(CamelCaseToWordsDashed(x.Name))
}

func (x *Module2Permission) AllUpper() string {
	return strings.ToUpper(CamelCaseToWordsUnderlined(x.Key))
}

func (x *Module2Permission) AllLower() string {
	return strings.ToLower(CamelCaseToWordsDashed(x.Key))
}

func (x *Module2Entity) PolyglotName() string {
	return x.EntityName() + "Polyglot"
}
func (x *Module2Entity) HasTranslations() bool {
	for _, field := range x.Fields {
		if field.Translate {
			return true
		}
	}
	return false
}
func (x *Module2Entity) DefinitionJson() string {
	data, _ := json.MarshalIndent(x, "", "  ")
	return string(data)
}
func (x *Module2DtoBase) DtoName() string {
	return ToUpper(x.Name) + "Dto"
}
func (x *Module2DtoBase) DefinitionJson() string {
	data, _ := json.MarshalIndent(x, "", "  ")
	return string(data)
}
func (x *Module2Entity) Upper() string {
	return ToUpper(x.Name)
}

func (x *Module2DtoBase) Upper() string {
	return ToUpper(x.Name)
}

/**
*	In module2 definitions we do have array and object fields,
*	which need to be stored in database in their own table
*	so we need to create those classes, etc...
**/
func GetArrayOrObjectFieldsFlatten(depth int, parentType string, depthName string, fields []*Module2Field, ctx *CodeGenContext, isWorkspace bool) []*Module2Field {
	items := []*Module2Field{}
	if len(fields) == 0 {
		return items
	}

	for _, item := range fields {
		if item.Type != FIELD_TYPE_OBJECT && item.Type != FIELD_TYPE_ARRAY {
			item.ComputedType = ctx.Catalog.ComputeField(item, isWorkspace)
			continue
		} else {
			item.LinkedTo = depthName
			if depth == 0 {
				item.LinkedTo += parentType
			}
			item.ComputedType = depthName + ctx.Catalog.ComputeField(item, isWorkspace)

		}

		item.FullName = depthName + item.PublicName()
		items = append(items, item)
		items = append(items, GetArrayOrObjectFieldsFlatten(
			depth+1,
			parentType,
			item.FullName,
			item.Fields, ctx, isWorkspace)...,
		)
	}

	return items
}

func ChildItems(x *Module2Entity, ctx *CodeGenContext, isWorkspace bool) []*Module2Field {

	return GetArrayOrObjectFieldsFlatten(0, "Entity", x.Upper(), x.Fields, ctx, isWorkspace)

}

func ChildItemsX(x *Module2DtoBase, ctx *CodeGenContext, isWorkspace bool) []*Module2Field {

	return GetArrayOrObjectFieldsFlatten(0, "Dto", x.Upper(), x.Fields, ctx, isWorkspace)

}

func (x *Module2Field) PluralName() string {

	pluralize2 := pluralize.NewClient()
	return pluralize2.Plural(x.Name)
}

func (x *Module2Entity) PluralNameUpper() string {

	pluralize2 := pluralize.NewClient()
	return ToUpper(pluralize2.Plural(x.Name))
}

func (x *Module2Entity) PluralName() string {

	pluralize2 := pluralize.NewClient()
	return pluralize2.Plural(x.Name)
}

func (x *Module2DtoBase) Template() string {
	return x.DashedName()
}

func (x *Module2DtoBase) Templates() string {
	pluralize2 := pluralize.NewClient()
	return strings.ToLower(pluralize2.Plural(x.Name))
}

func (x *Module2DtoBase) TemplatesLower() string {
	return x.PluralNameLower()
}

func (x *Module2Entity) Template() string {
	return x.DashedName()
}

func (x *Module2Entity) Templates() string {
	pluralize2 := pluralize.NewClient()
	return strings.ToLower(pluralize2.Plural(x.Name))
}

func (x *Module2Entity) TemplatesLower() string {
	return x.PluralNameLower()
}

func (x *Module2Entity) PluralNameLower() string {

	pluralize2 := pluralize.NewClient()
	return strings.ToLower(pluralize2.Plural(x.Name))
}

func (x *Module2DtoBase) PluralNameLower() string {

	pluralize2 := pluralize.NewClient()
	return strings.ToLower(pluralize2.Plural(x.Name))
}

func (x *Module2Entity) DashedPluralName() string {

	pluralize2 := pluralize.NewClient()
	return strings.ReplaceAll(ToSnakeCase(pluralize2.Plural(x.Name)), "_", "-")
}

func (x *Module2Entity) TableName() string {

	return ToSnakeCase((x.Name))
}

func (x *Module2DtoBase) DashedName() string {
	return strings.ReplaceAll(ToSnakeCase(x.Name), "_", "-")
}
func (x *Module2Entity) DashedName() string {
	return strings.ReplaceAll(ToSnakeCase(x.Name), "_", "-")
}

func (x *Module2Entity) FormName() string {
	return ToUpper(x.Name) + "Form"
}

func ImportDependecies(fields []*Module2Field) []ImportDependencyStrategy {
	items := []ImportDependencyStrategy{
		{
			Path:  "../../core/definitions",
			Items: []string{"BaseDto", "BaseEntity"},
		},
	}

	for _, field := range fields {

		if field.Type == FIELD_TYPE_JSON {

			if len(field.Matches) > 0 {
				for _, item := range field.Matches {
					if item.Dto != nil {
						items = append(items, ImportDependencyStrategy{
							Items: []string{item.PublicName()},

							// This loads only the same directory items
							Path: "./" + item.PublicName(),
						})
					}
				}
			}
		}

		if field.Type == FIELD_TYPE_ARRAY || field.Type == FIELD_TYPE_OBJECT {
			items = append(items, ImportDependecies(field.Fields)...)
		}

		if field.Type != FIELD_TYPE_ONE && field.Type != FIELD_TYPE_MANY2MANY {
			continue
		}

		// Computed path is based on the idea that every thing is only one folder
		// inside. This might required revise if we want to build for C++, etc

		computedPath := ""
		if field.Module != "" {
			computedPath = "../" + field.Module + "/" + field.Target
		} else {
			if field.RootClass != "" {
				computedPath = "./" + field.RootClass

			} else {
				computedPath = "./" + field.Target

			}
		}

		items = append(items, ImportDependencyStrategy{
			Items: []string{field.Target},
			Path:  computedPath,
		})
	}

	return items
}

func (x *Module2) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

func (x *Module2) Yaml() string {
	if x != nil {
		str, _ := yaml.Marshal(x)
		return (string(str))
	}
	return ""
}

/*
Give me golang function, which gets and string, and convert it into camelCase,
and output should only consist of numbers and english alphabet lowercase and uppercase (only).
*/

func ToCamelCaseClean(input string) string {
	splitBySpecial := regexp.MustCompile("[^A-Za-z0-9]+")
	words := splitBySpecial.Split(input, -1)

	var result string
	for _, word := range words {
		// Convert each word to camel case
		word = strings.ToLower(word)
		if word == "" {
			continue
		}
		word = strings.ToUpper(word[0:1]) + word[1:]
		result += word
	}

	// Remove non-alphanumeric characters
	nonAlphaNumeric := regexp.MustCompile("[^A-Za-z0-9]")
	result = nonAlphaNumeric.ReplaceAllString(result, "")

	return ToLower(result)
}

func ImportGoDependencies(fields []*Module2Field, importGroupPrefix string) []ImportDependencyStrategy {
	items := []ImportDependencyStrategy{}

	for _, field := range fields {
		actualPrefix := importGroupPrefix

		if field.Provider != "" {
			actualPrefix = field.Provider

			// this is magical keyword resolves to the fireback on github
			if field.Provider == "fireback" {
				actualPrefix = "github.com/torabian/fireback/modules/"
			}

			if !strings.HasSuffix(actualPrefix, "/") {
				actualPrefix = actualPrefix + "/"
			}
		}
		// if field.Type == FIELD_TYPE_JSON {
		// 	items = append(items, ImportDependencyStrategy{
		// 		Items: []string{field.Target},
		// 	})
		// }

		if field.Type == FIELD_TYPE_ARRAY || field.Type == FIELD_TYPE_OBJECT {

			items = append(items, ImportGoDependencies(field.Fields, actualPrefix)...)
		}

		if field.Type != FIELD_TYPE_ONE && field.Type != FIELD_TYPE_MANY2MANY {
			continue
		}

		if field.Type == FIELD_TYPE_DATE {
			items = append(items, ImportDependencyStrategy{
				Path: "time",
			})
		}

		if field.Module != "" && field.Module != "workspaces" {
			items = append(items, ImportDependencyStrategy{
				Items: []string{field.Target},
				Path:  actualPrefix + field.Module,
			})

		}
	}

	return items
}

func (x Module2Action) RequestRootObjectName() string {
	reqValue := reflect.ValueOf(x.RequestEntity)
	if reqValue.MethodByName("RootObjectName").IsValid() {
		res := reqValue.MethodByName("RootObjectName").Call(nil)

		if len(res) > 0 {
			return res[0].String()
		}
	}
	return ""
}

func (x Module2Action) ResponseRootObjectName() string {
	reqValue := reflect.ValueOf(x.ResponseEntity)
	if reqValue.MethodByName("RootObjectName").IsValid() {
		res := reqValue.MethodByName("RootObjectName").Call(nil)

		if len(res) > 0 {
			return res[0].String()
		}
	}
	return ""
}

func (x Module2Action) RequestExample() string {
	if x.RequestEntity == nil {
		return ""
	}

	reqValue := reflect.ValueOf(x.RequestEntity)
	if reqValue.MethodByName("Seeder").IsValid() {
		res := reqValue.MethodByName("Seeder").Call(nil)

		if len(res) > 0 {
			return res[0].String()
		}
	}
	return ""
}

// In Typescript, we put all req/res dtos related to custom actions
// into a single file, this is why we need to get the correct name
func TsObjectName(objectName string, rootObjectName string) string {

	if strings.Contains(objectName, "ActionReq") || strings.Contains(objectName, "ActionRes") {
		return ToUpper(rootObjectName) + "ActionsDto"
	}

	return rootObjectName
}

func (x ImportMap) AppendImportMapRow(key string, row *ImportMapRow) {
	if x[key] == nil {
		x[key] = row
	} else {
		// Complete this logic over time, if necessary
		for _, item := range row.Items {
			if !Contains(x[key].Items, item) {
				x[key].Items = append(x[key].Items, item)
			}
		}
	}
}

func (x Module2Action) ImportDependecies() ImportMap {
	m := ImportMap{}

	if x.RequestEntity != nil {
		meta := x.RequestEntityMeta()
		fileName := meta.ClassName
		if name := x.RequestRootObjectName(); name != "" {
			fileName = TsObjectName(meta.ClassName, name)
		}

		m.AppendImportMapRow("../"+meta.Module+"/"+fileName, &ImportMapRow{
			Items: []string{meta.ClassName},
		})
	} else if x.RequestEntityComputed() != "" {
		// This is more for external defined files
		m.AppendImportMapRow("./"+x.RequestEntityComputed(), &ImportMapRow{
			Items: []string{x.RequestEntityComputed()},
		})
	}

	if x.ResponseEntity != nil {
		meta := x.ResponseEntityMeta()
		fileName := meta.ClassName
		if name := x.ResponseRootObjectName(); name != "" {
			fileName = TsObjectName(meta.ClassName, name)
		}

		m.AppendImportMapRow("../"+meta.Module+"/"+fileName, &ImportMapRow{
			Items: []string{meta.ClassName},
		})

	} else if x.ResponseEntityComputed() != "" {
		// This is more for external defined files
		m.AppendImportMapRow("./"+x.ResponseEntityComputed(), &ImportMapRow{
			Items: []string{x.ResponseEntityComputed()},
		})
	}

	return m
}

// Converts import strategy into unique map to be ported into the template.
// ImportDependencies might generate duplicate elements, here we make them unique
// or any other last moment changes
func (x *Module2Entity) ImportGroupResolver(prefix string) ImportMap {

	deps := ImportGoDependencies(x.Fields, prefix)

	m := ImportMap{}
	for _, dep := range deps {
		if m[dep.Path] == nil {
			m[dep.Path] = &ImportMapRow{}
		}

		for _, klass := range dep.Items {
			if !Contains(m[dep.Path].Items, klass) && klass != x.EntityName() {
				m[dep.Path].Items = append(m[dep.Path].Items, klass)
			}
		}

	}

	return m

}
func (x *Module2) TsActionsImport() ImportMap {
	m := ImportMap{}

	for _, action := range x.Actions {

		{

			deps := ImportDependecies(action.In.Fields)

			for _, dep := range deps {
				if m[dep.Path] == nil {
					m[dep.Path] = &ImportMapRow{}
				}

				for _, klass := range dep.Items {
					if !Contains(m[dep.Path].Items, klass) {
						m[dep.Path].Items = append(m[dep.Path].Items, klass)
					}
				}

			}
		}

		{
			deps := ImportDependecies(action.Out.Fields)

			for _, dep := range deps {
				if m[dep.Path] == nil {
					m[dep.Path] = &ImportMapRow{}
				}

				for _, klass := range dep.Items {
					if !Contains(m[dep.Path].Items, klass) {
						m[dep.Path].Items = append(m[dep.Path].Items, klass)
					}
				}

			}
		}
	}
	return m

}
func (x *Module2Entity) ImportDependecies() ImportMap {

	deps := ImportDependecies(x.Fields)

	m := ImportMap{}
	for _, dep := range deps {
		if m[dep.Path] == nil {
			m[dep.Path] = &ImportMapRow{}
		}

		for _, klass := range dep.Items {
			if !Contains(m[dep.Path].Items, klass) && klass != x.EntityName() {
				m[dep.Path].Items = append(m[dep.Path].Items, klass)
			}
		}

	}

	return m

}
func (x *Module2DtoBase) ImportDependecies() ImportMap {

	deps := ImportDependecies(x.Fields)

	m := ImportMap{}
	for _, dep := range deps {
		if m[dep.Path] == nil {
			m[dep.Path] = &ImportMapRow{}
		}

		for _, klass := range dep.Items {
			if !Contains(m[dep.Path].Items, klass) && klass != x.DtoName() {
				m[dep.Path].Items = append(m[dep.Path].Items, klass)
			}
		}

	}

	return m

}

func HasSeeders(module *Module2, entity *Module2Entity) bool {
	checkee := filepath.Join("modules", module.Path, "seeders", entity.Upper())

	if _, err := os.Stat(checkee); !os.IsNotExist(err) {

		return true
	}

	return false
}

func HasMetas(module *Module2, entity *Module2Entity) bool {
	checkee := filepath.Join("modules", module.Path, "seeders", entity.Upper())

	if _, err := os.Stat(checkee); !os.IsNotExist(err) {

		return true
	}

	return false
}

func HasMocks(module *Module2, entity *Module2Entity) bool {
	mocks := filepath.Join("modules", module.Path, "mocks", entity.Upper())

	if _, err := os.Stat(mocks); !os.IsNotExist(err) {
		return true
	}
	return false
}

func generateRange(start, end int) []int {
	result := make([]int, end-start+1)
	for i := range result {
		result[i] = i + start
	}
	return result
}

var CommonMap = template.FuncMap{
	"until": generateRange,
	"join":  strings.Join,
	"upper": ToUpper,
	"arr":   func(els ...any) []any { return els },
	"inc": func(i int) int {
		return i + 1
	},
	"fx": func(fieldName string, depth int) string {
		return fieldName + "[index" + fmt.Sprintf("%v", depth) + "]."
	},
}

func mergeMaps(map1, map2 map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})

	// Copy values from map1 to merged
	for key, value := range map1 {
		merged[key] = value
	}

	// Copy values from map2 to merged, overwriting existing keys
	for key, value := range map2 {
		merged[key] = value
	}

	return merged
}
func (x *Module2Entity) RenderTemplate(
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
	module *Module2,
	map2 map[string]interface{},
) ([]byte, error) {
	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}
	var tpl bytes.Buffer

	wsPrefix := "workspaces."
	isWorkspace := false
	if module.Path == "workspaces" {
		wsPrefix = ""
		isWorkspace = true
	}

	params := gin.H{
		"e":           x,
		"m":           module,
		"fv":          FIREBACK_VERSION,
		"gofModule":   ctx.GofModuleName,
		"ctx":         ctx,
		"children":    ChildItems(x, ctx, isWorkspace),
		"imports":     x.ImportDependecies(),
		"goimports":   x.ImportGroupResolver(ctx.GofModuleName + "/"),
		"javaimports": x.ImportGroupResolver("com.fireback.modules."),
		"wsprefix":    wsPrefix,
		"hasSeeders":  HasSeeders(module, x),
		"hasMetas":    HasMetas(module, x),
		"hasMocks":    HasMocks(module, x),
	}

	err = t.ExecuteTemplate(&tpl, fname, mergeMaps(params, map2))

	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}
func (x *Module2Entity) RenderCteSqlTemplate(
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
	module *Module2,
	tp string,
) ([]byte, error) {
	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}
	var tpl bytes.Buffer

	wsPrefix := "workspaces."
	if module.Path == "workspaces" {
		wsPrefix = ""
	}

	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"m":        module,
		"wsprefix": wsPrefix,
		"type":     tp,
		"e":        x,
		"fv":       FIREBACK_VERSION,
		"ctx":      ctx,
	})
	if err != nil {
		return []byte{}, err
	}

	y := string(tpl.Bytes())
	y = strings.ReplaceAll(y, "template", x.TableName())
	return []byte(y), nil
}

func (action *Module2Action) Render(
	x *Module2,
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
) ([]byte, error) {
	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}
	var tpl bytes.Buffer

	isWorkspace := false

	wsPrefix := "workspaces."
	if x.Path == "workspaces" {
		wsPrefix = ""
		isWorkspace = true
	}

	if len(action.In.Fields) > 0 {
		ComputeFieldTypes(action.In.Fields, isWorkspace, ctx.Catalog.ComputeField)
	}
	if len(action.Out.Fields) > 0 {

		ComputeFieldTypes(action.Out.Fields, isWorkspace, ctx.Catalog.ComputeField)
	}

	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"m":               x,
		"ctx":             ctx,
		"fv":              FIREBACK_VERSION,
		"wsprefix":        wsPrefix,
		"action":          action,
		"tsactionimports": x.TsActionsImport(),
	})
	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}

func (x *Module2) RenderActions(
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
) ([]byte, error) {
	isWorkspace := false

	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}
	var tpl bytes.Buffer

	wsPrefix := "workspaces."
	if x.Path == "workspaces" {
		wsPrefix = ""
		isWorkspace = true

	}

	for _, action := range x.Actions {

		if len(action.In.Fields) > 0 {
			ComputeFieldTypes(action.In.Fields, isWorkspace, ctx.Catalog.ComputeField)
		}
		if len(action.Out.Fields) > 0 {
			ComputeFieldTypes(action.Out.Fields, isWorkspace, ctx.Catalog.ComputeField)
		}

	}

	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"m":               x,
		"fv":              FIREBACK_VERSION,
		"wsprefix":        wsPrefix,
		"tsactionimports": x.TsActionsImport(),
		"ctx":             ctx,
	})
	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}

func (x *Module2Entity) GetSqlFields() []string {
	items := []string{
		"fb_template_entities.parent_id",
		"fb_template_entities.visibility",
		"fb_template_entities.updated",
		"fb_template_entities.created",
	}
	for _, field := range x.Fields {
		if field.Type == "object" {
			continue
		}
		if field.Type == "one" {
			items = append(items, "fb_template_entities."+ToSnakeCase(field.Name)+"_id")
		} else {
			items = append(items, "fb_template_entities."+ToSnakeCase(field.Name))
		}

	}

	return items
}

func (x *Module2Entity) GetSqlFieldNames() []string {
	items := []string{"parent_id", "visibility", "updated", "created"}
	for _, field := range x.Fields {
		if field.Type == "object" {
			continue
		}

		if field.Type == "one" {
			items = append(items, ToSnakeCase(field.Name)+"_id")
		} else {
			items = append(items, ToSnakeCase(field.Name))
		}

	}

	return items
}

func (x *Module2Entity) GetSqlFieldNamesAfter() []string {
	items := []string{
		"fb_template_entities_cte.parent_id",
		"fb_template_entities_cte.visibility",
		"fb_template_entities_cte.updated",
		"fb_template_entities_cte.created",
	}
	for _, field := range x.Fields {
		if field.Type == "object" {
			continue
		}

		if field.Type == "one" {
			items = append(items, "fb_template_entities_cte."+ToSnakeCase(field.Name)+"_id\n")
		} else {

			if field.Translate {
				items = append(items, "fb_template_entity_polyglots."+ToSnakeCase(field.Name)+"\n")
			} else {
				items = append(items, "fb_template_entities_cte."+ToSnakeCase(field.Name)+"\n")
			}
		}

	}

	return items
}

func (x *Module2DtoBase) RenderTemplate(
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
	module *Module2,
) ([]byte, error) {

	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}
	var tpl bytes.Buffer

	wsPrefix := "workspaces."
	isWorkspace := false
	if module.Path == "workspaces" {
		wsPrefix = ""
		isWorkspace = true
	}

	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"e":        x,
		"children": ChildItemsX(x, ctx, isWorkspace),
		"imports":  x.ImportDependecies(),
		"m":        module,
		"ctx":      ctx,
		"fv":       FIREBACK_VERSION,
		"wsprefix": wsPrefix,
	})

	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}

func (x *Module2) RenderTemplate(
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
) ([]byte, error) {

	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}
	var tpl bytes.Buffer

	wsPrefix := "workspaces."
	if x.Path == "workspaces" {
		wsPrefix = ""
	}

	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"m":        x,
		"fv":       FIREBACK_VERSION,
		"wsprefix": wsPrefix,
	})

	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}

func (x Module2Action) RenderTemplate(ctx *CodeGenContext, fs embed.FS, fname string) ([]byte, error) {
	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}

	wsPrefix := "workspaces."

	var tpl bytes.Buffer
	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"r":        x,
		"m":        x.RootModule,
		"ctx":      ctx,
		"fv":       FIREBACK_VERSION,
		"imports":  x.ImportDependecies(),
		"wsprefix": wsPrefix,
	})
	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}

func RenderRpcGroupClassBody(
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
	content []byte,
	group string,
	importMap ImportMap,
) ([]byte, error) {
	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}

	wsPrefix := "workspaces."

	var tpl bytes.Buffer
	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"content":  content,
		"ctx":      ctx,
		"group":    group,
		"Group":    ToUpper(group),
		"imports":  importMap,
		"fv":       FIREBACK_VERSION,
		"wsprefix": wsPrefix,
	})
	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}

// Detect undefined and null in golang
// I think there is a huge problem to use this
// How to pass it to gorm?
// type SampleEntity or Dto struct {
//     Field1 Optional[string] `json:"field1"`
//     Field2 Optional[bool]   `json:"field2"`
//     Field3 Optional[int32]  `json:"field3"`
// }

type Optional[T any] struct {
	Defined bool
	Value   *T
}

// UnmarshalJSON is implemented by deferring to the wrapped type (T).
// It will be called only if the value is defined in the JSON payload.
func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	o.Defined = true
	return json.Unmarshal(data, &o.Value)
}
