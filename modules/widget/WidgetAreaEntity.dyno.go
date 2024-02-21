package widget
import (
    "github.com/gin-gonic/gin"
	"pixelplux.com/fireback/modules/workspaces"
	"log"
	"os"
	"fmt"
	"encoding/json"
	"strings"
	"github.com/schollz/progressbar/v3"
	"github.com/gookit/event"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	jsoniter "github.com/json-iterator/go"
	"embed"
	reflect "reflect"
	"github.com/urfave/cli"
	mocks "pixelplux.com/fireback/modules/widget/mocks/WidgetArea"
)
type WidgetAreaWidgets struct {
    Visibility       *string                         `json:"visibility,omitempty" yaml:"visibility"`
    WorkspaceId      *string                         `json:"workspaceId,omitempty" yaml:"workspaceId"`
    LinkerId         *string                         `json:"linkerId,omitempty" yaml:"linkerId"`
    ParentId         *string                         `json:"parentId,omitempty" yaml:"parentId"`
    UniqueId         string                          `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
    UserId           *string                         `json:"userId,omitempty" yaml:"userId"`
    Rank             int64                           `json:"rank,omitempty" gorm:"type:int;name:rank"`
    Updated          int64                           `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
    Created          int64                           `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
    CreatedFormatted string                          `json:"createdFormatted,omitempty" sql:"-"`
    UpdatedFormatted string                          `json:"updatedFormatted,omitempty" sql:"-"`
    Title   *string `json:"title" yaml:"title"        translate:"true" `
    // Datenano also has a text representation
    Widget   *  WidgetEntity `json:"widget" yaml:"widget"    gorm:"foreignKey:WidgetId;references:UniqueId"     `
    // Datenano also has a text representation
        WidgetId *string `json:"widgetId" yaml:"widgetId"`
    X   *int64 `json:"x" yaml:"x"       `
    // Datenano also has a text representation
    Y   *int64 `json:"y" yaml:"y"       `
    // Datenano also has a text representation
    W   *int64 `json:"w" yaml:"w"       `
    // Datenano also has a text representation
    H   *int64 `json:"h" yaml:"h"       `
    // Datenano also has a text representation
    Data   *string `json:"data" yaml:"data"       `
    // Datenano also has a text representation
	LinkedTo *WidgetAreaEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
func ( x * WidgetAreaWidgets) RootObjectName() string {
	return "WidgetAreaEntity"
}
type WidgetAreaEntity struct {
    Visibility       *string                         `json:"visibility,omitempty" yaml:"visibility"`
    WorkspaceId      *string                         `json:"workspaceId,omitempty" yaml:"workspaceId"`
    LinkerId         *string                         `json:"linkerId,omitempty" yaml:"linkerId"`
    ParentId         *string                         `json:"parentId,omitempty" yaml:"parentId"`
    UniqueId         string                          `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
    UserId           *string                         `json:"userId,omitempty" yaml:"userId"`
    Rank             int64                           `json:"rank,omitempty" gorm:"type:int;name:rank"`
    Updated          int64                           `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
    Created          int64                           `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
    CreatedFormatted string                          `json:"createdFormatted,omitempty" sql:"-"`
    UpdatedFormatted string                          `json:"updatedFormatted,omitempty" sql:"-"`
    Name   *string `json:"name" yaml:"name"        translate:"true" `
    // Datenano also has a text representation
    Layouts   *string `json:"layouts" yaml:"layouts"       `
    // Datenano also has a text representation
    Widgets   []*  WidgetAreaWidgets `json:"widgets" yaml:"widgets"    gorm:"foreignKey:LinkerId;references:UniqueId"     `
    // Datenano also has a text representation
    Translations     []*WidgetAreaEntityPolyglot `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    Children []*WidgetAreaEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *WidgetAreaEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var WidgetAreaPreloadRelations []string = []string{}
var WIDGETAREA_EVENT_CREATED = "widgetArea.created"
var WIDGETAREA_EVENT_UPDATED = "widgetArea.updated"
var WIDGETAREA_EVENT_DELETED = "widgetArea.deleted"
var WIDGETAREA_EVENTS = []string{
	WIDGETAREA_EVENT_CREATED,
	WIDGETAREA_EVENT_UPDATED,
	WIDGETAREA_EVENT_DELETED,
}
type WidgetAreaFieldMap struct {
		Name workspaces.TranslatedString `yaml:"name"`
		Layouts workspaces.TranslatedString `yaml:"layouts"`
		Widgets workspaces.TranslatedString `yaml:"widgets"`
}
var WidgetAreaEntityMetaConfig map[string]int64 = map[string]int64{
}
var WidgetAreaEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&WidgetAreaEntity{}))
  type WidgetAreaEntityPolyglot struct {
    LinkerId string `gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId"`
    LanguageId string `gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId"`
        Name string `yaml:"name" json:"name"`
  }
func WidgetAreaWidgetsActionCreate(
  dto *WidgetAreaWidgets ,
  query workspaces.QueryDSL,
) (*WidgetAreaWidgets , *workspaces.IError) {
    dto.LinkerId = &query.LinkerId
    var dbref *gorm.DB = nil
    if query.Tx == nil {
        dbref = workspaces.GetDbRef()
    } else {
        dbref = query.Tx
    }
    query.Tx = dbref
    if dto.UniqueId == "" {
        dto.UniqueId = workspaces.UUID()
    }
    err := dbref.Create(&dto).Error
    if err != nil {
        err := workspaces.GormErrorToIError(err)
        return dto, err
    }
    return dto, nil
}
func WidgetAreaWidgetsActionUpdate(
    query workspaces.QueryDSL,
    dto *WidgetAreaWidgets,
) (*WidgetAreaWidgets, *workspaces.IError) {
    dto.LinkerId = &query.LinkerId
    var dbref *gorm.DB = nil
    if query.Tx == nil {
        dbref = workspaces.GetDbRef()
    } else {
        dbref = query.Tx
    }
    query.Tx = dbref
    err := dbref.UpdateColumns(&dto).Error
    if err != nil {
        err := workspaces.GormErrorToIError(err)
        return dto, err
    }
    return dto, nil
}
func WidgetAreaWidgetsActionGetOne(
    query workspaces.QueryDSL,
) (*WidgetAreaWidgets , *workspaces.IError) {
    refl := reflect.ValueOf(&WidgetAreaWidgets {})
    item, err := workspaces.GetOneEntity[WidgetAreaWidgets ](query, refl)
    return item, err
}
func entityWidgetAreaFormatter(dto *WidgetAreaEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
	if dto.Created > 0 {
		dto.CreatedFormatted = workspaces.FormatDateBasedOnQuery(dto.Created, query)
	}
	if dto.Updated > 0 {
		dto.CreatedFormatted = workspaces.FormatDateBasedOnQuery(dto.Updated, query)
	}
}
func WidgetAreaMockEntity() *WidgetAreaEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &WidgetAreaEntity{
      Name : &stringHolder,
      Layouts : &stringHolder,
	}
	return entity
}
func WidgetAreaActionSeeder(query workspaces.QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := WidgetAreaMockEntity()
		_, err := WidgetAreaActionCreate(entity, query)
		if err == nil {
			successInsert++
		} else {
			fmt.Println(err)
			failureInsert++
		}
		bar.Add(1)
	}
	fmt.Println("Success", successInsert, "Failure", failureInsert)
}
    func (x*WidgetAreaEntity) GetNameTranslated(language string) string{
      if x.Translations != nil && len(x.Translations) > 0{
        for _, item := range x.Translations {
          if item.LanguageId == language {
              return item.Name
          }
        }
      }
      return ""
    }
  func WidgetAreaActionSeederInit(query workspaces.QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*WidgetAreaEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &WidgetAreaEntity{
          Name: &tildaRef,
          Layouts: &tildaRef,
          Widgets: []*WidgetAreaWidgets{{}},
    }
    data = append(data, entity)
    if format == "yml" || format == "yaml" {
      body, err = yaml.Marshal(data)
      if err != nil {
        log.Fatal(err)
      }
    }
    if format == "json" {
      body, err = json.MarshalIndent(data, "", "  ")
      if err != nil {
        log.Fatal(err)
      }
      file = strings.Replace(file, ".yml", ".json", -1)
    }
    os.WriteFile(file, body, 0644)
  }
  func WidgetAreaAssociationCreate(dto *WidgetAreaEntity, query workspaces.QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func WidgetAreaRelationContentCreate(dto *WidgetAreaEntity, query workspaces.QueryDSL) error {
return nil
}
func WidgetAreaRelationContentUpdate(dto *WidgetAreaEntity, query workspaces.QueryDSL) error {
	return nil
}
func WidgetAreaPolyglotCreateHandler(dto *WidgetAreaEntity, query workspaces.QueryDSL) {
	if dto == nil {
		return
	}
    workspaces.PolyglotCreateHandler(dto, &WidgetAreaEntityPolyglot{}, query)
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func WidgetAreaValidator(dto *WidgetAreaEntity, isPatch bool) *workspaces.IError {
    err := workspaces.CommonStructValidatorPointer(dto, isPatch)
        if dto != nil && dto.Widgets != nil {
          workspaces.AppendSliceErrors(dto.Widgets, isPatch, "widgets", err)
        }
    return err
  }
func WidgetAreaEntityPreSanitize(dto *WidgetAreaEntity, query workspaces.QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func WidgetAreaEntityBeforeCreateAppend(dto *WidgetAreaEntity, query workspaces.QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = workspaces.UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    WidgetAreaRecursiveAddUniqueId(dto, query)
  }
  func WidgetAreaRecursiveAddUniqueId(dto *WidgetAreaEntity, query workspaces.QueryDSL) {
      if dto.Widgets != nil && len(dto.Widgets) > 0 {
        for index0 := range dto.Widgets {
          if (dto.Widgets[index0].UniqueId == "") {
            dto.Widgets[index0].UniqueId = workspaces.UUID()
          }
        }
    }
  }
func WidgetAreaActionBatchCreateFn(dtos []*WidgetAreaEntity, query workspaces.QueryDSL) ([]*WidgetAreaEntity, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*WidgetAreaEntity{}
		for _, item := range dtos {
			s, err := WidgetAreaActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func WidgetAreaActionCreateFn(dto *WidgetAreaEntity, query workspaces.QueryDSL) (*WidgetAreaEntity, *workspaces.IError) {
	// 1. Validate always
	if iError := WidgetAreaValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	WidgetAreaEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	WidgetAreaEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	WidgetAreaPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	WidgetAreaRelationContentCreate(dto, query)
	// 4. Create the entity
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = workspaces.GetDbRef()
	} else {
		dbref = query.Tx
	}
	query.Tx = dbref;
	err := dbref.Create(&dto).Error
	if err != nil {
		err := workspaces.GormErrorToIError(err)
		return dto, err
	}
	// 5. Create sub entities, objects or arrays, association to other entities
	WidgetAreaAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(WIDGETAREA_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": workspaces.GetTypeString(&WidgetAreaEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func WidgetAreaActionGetOne(query workspaces.QueryDSL) (*WidgetAreaEntity, *workspaces.IError) {
    refl := reflect.ValueOf(&WidgetAreaEntity{})
    item, err := workspaces.GetOneEntity[WidgetAreaEntity](query, refl)
    entityWidgetAreaFormatter(item, query)
    return item, err
  }
  func WidgetAreaActionQuery(query workspaces.QueryDSL) ([]*WidgetAreaEntity, *workspaces.QueryResultMeta, error) {
    refl := reflect.ValueOf(&WidgetAreaEntity{})
    items, meta, err := workspaces.QueryEntitiesPointer[WidgetAreaEntity](query, refl)
    for _, item := range items {
      entityWidgetAreaFormatter(item, query)
    }
    return items, meta, err
  }
  func WidgetAreaUpdateExec(dbref *gorm.DB, query workspaces.QueryDSL, fields *WidgetAreaEntity) (*WidgetAreaEntity, *workspaces.IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = WIDGETAREA_EVENT_UPDATED
    WidgetAreaEntityPreSanitize(fields, query)
    var item WidgetAreaEntity
    q := dbref.
      Where(&WidgetAreaEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, workspaces.GormErrorToIError(err)
    }
    query.Tx = dbref
    WidgetAreaRelationContentUpdate(fields, query)
    WidgetAreaPolyglotCreateHandler(fields, query)
    // @meta(update has many)
      if fields.Widgets != nil {
       linkerId := uniqueId;
        dbref.Debug().
          Where(&WidgetAreaWidgets {LinkerId: &linkerId}).
          Delete(&WidgetAreaWidgets {})
        for _, newItem := range fields.Widgets {
          newItem.UniqueId = workspaces.UUID()
          newItem.LinkerId = &linkerId
          dbref.Create(&newItem)
        }
      }
    err = dbref.
      Preload(clause.Associations).
      Where(&WidgetAreaEntity{UniqueId: uniqueId}).
      First(&item).Error
    event.MustFire(query.TriggerEventName, event.M{
      "entity":   &item,
      "target":   "workspace",
      "unqiueId": query.WorkspaceId,
    })
    if err != nil {
      return &item, workspaces.GormErrorToIError(err)
    }
    return &item, nil
  }
  func WidgetAreaActionUpdateFn(query workspaces.QueryDSL, fields *WidgetAreaEntity) (*WidgetAreaEntity, *workspaces.IError) {
    if fields == nil {
      return nil, workspaces.CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := WidgetAreaValidator(fields, true); iError != nil {
      return nil, iError
    }
    WidgetAreaRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = workspaces.GetDbRef()
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        _, err := WidgetAreaUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return nil, workspaces.CastToIError(vf)
    } else {
      dbref = query.Tx
      return WidgetAreaUpdateExec(dbref, query, fields)
    }
  }
var WidgetAreaWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire widgetareas ",
	Action: func(c *cli.Context) error {
		query := workspaces.CommonCliQueryDSLBuilder(c)
		count, _ := WidgetAreaActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func WidgetAreaActionRemove(query workspaces.QueryDSL) (int64, *workspaces.IError) {
	refl := reflect.ValueOf(&WidgetAreaEntity{})
	query.ActionRequires = []string{PERM_ROOT_WIDGETAREA_DELETE}
	return workspaces.RemoveEntity[WidgetAreaEntity](query, refl)
}
func WidgetAreaActionWipeClean(query workspaces.QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
			{
				subCount, subErr := workspaces.WipeCleanEntity[WidgetAreaWidgets]()
				if (subErr != nil) {
					fmt.Println("Error while wiping 'WidgetAreaWidgets'", subErr)
					return count, subErr
				} else {
					count += subCount
				}
			}
	{
		subCount, subErr := workspaces.WipeCleanEntity[WidgetAreaEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'WidgetAreaEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func WidgetAreaActionBulkUpdate(
    query workspaces.QueryDSL, dto *workspaces.BulkRecordRequest[WidgetAreaEntity]) (
    *workspaces.BulkRecordRequest[WidgetAreaEntity], *workspaces.IError,
  ) {
    result := []*WidgetAreaEntity{}
    err := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := WidgetAreaActionUpdate(query, record)
        if err != nil {
          return err
        } else {
          result = append(result, item)
        }
      }
      return nil
    })
    if err == nil {
      return dto, nil
    }
    return nil, err.(*workspaces.IError)
  }
func (x *WidgetAreaEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var WidgetAreaEntityMeta = workspaces.TableMetaData{
	EntityName:    "WidgetArea",
	ExportKey:    "widget-areas",
	TableNameInDb: "fb_widgetarea_entities",
	EntityObject:  &WidgetAreaEntity{},
	ExportStream: WidgetAreaActionExportT,
	ImportQuery: WidgetAreaActionImport,
}
func WidgetAreaActionExport(
	query workspaces.QueryDSL,
) (chan []byte, *workspaces.IError) {
	return workspaces.YamlExporterChannel[WidgetAreaEntity](query, WidgetAreaActionQuery, WidgetAreaPreloadRelations)
}
func WidgetAreaActionExportT(
	query workspaces.QueryDSL,
) (chan []interface{}, *workspaces.IError) {
	return workspaces.YamlExporterChannelT[WidgetAreaEntity](query, WidgetAreaActionQuery, WidgetAreaPreloadRelations)
}
func WidgetAreaActionImport(
	dto interface{}, query workspaces.QueryDSL,
) *workspaces.IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content WidgetAreaEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return workspaces.CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := WidgetAreaActionCreate(&content, query)
	return err
}
var WidgetAreaCommonCliFlags = []cli.Flag{
  &cli.StringFlag{
    Name:     "wid",
    Required: false,
    Usage:    "Provide workspace id, if you want to change the data workspace",
  },
  &cli.StringFlag{
    Name:     "uid",
    Required: false,
    Usage:    "uniqueId (primary key)",
  },
  &cli.StringFlag{
    Name:     "pid",
    Required: false,
    Usage:    " Parent record id of the same type",
  },
    &cli.StringFlag{
      Name:     "name",
      Required: false,
      Usage:    "name",
    },
    &cli.StringFlag{
      Name:     "layouts",
      Required: false,
      Usage:    "layouts",
    },
    &cli.StringSliceFlag{
      Name:     "widgets",
      Required: false,
      Usage:    "widgets",
    },
}
var WidgetAreaCommonInteractiveCliFlags = []workspaces.CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: false,
		Usage:    "name",
		Type: "string",
	},
	{
		Name:     "layouts",
		StructField:     "Layouts",
		Required: false,
		Usage:    "layouts",
		Type: "string",
	},
}
var WidgetAreaCommonCliFlagsOptional = []cli.Flag{
  &cli.StringFlag{
    Name:     "wid",
    Required: false,
    Usage:    "Provide workspace id, if you want to change the data workspace",
  },
  &cli.StringFlag{
    Name:     "uid",
    Required: false,
    Usage:    "uniqueId (primary key)",
  },
  &cli.StringFlag{
    Name:     "pid",
    Required: false,
    Usage:    " Parent record id of the same type",
  },
    &cli.StringFlag{
      Name:     "name",
      Required: false,
      Usage:    "name",
    },
    &cli.StringFlag{
      Name:     "layouts",
      Required: false,
      Usage:    "layouts",
    },
    &cli.StringSliceFlag{
      Name:     "widgets",
      Required: false,
      Usage:    "widgets",
    },
}
  var WidgetAreaCreateCmd cli.Command = cli.Command{
    Name:    "create",
    Aliases: []string{"c"},
    Flags: WidgetAreaCommonCliFlags,
    Usage: "Create a new template",
    Action: func(c *cli.Context) {
      query := workspaces.CommonCliQueryDSLBuilder(c)
      entity := CastWidgetAreaFromCli(c)
      if entity, err := WidgetAreaActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var WidgetAreaCreateInteractiveCmd cli.Command = cli.Command{
    Name:  "ic",
    Usage: "Creates a new template, using requied fields in an interactive name",
    Flags: []cli.Flag{
      &cli.BoolFlag{
        Name:  "all",
        Usage: "Interactively asks for all inputs, not only required ones",
      },
    },
    Action: func(c *cli.Context) {
      query := workspaces.CommonCliQueryDSLBuilder(c)
      entity := &WidgetAreaEntity{}
      for _, item := range WidgetAreaCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := workspaces.AskForInput(item.Name, "")
        workspaces.SetFieldString(entity, item.StructField, result)
      }
      if entity, err := WidgetAreaActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var WidgetAreaUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: WidgetAreaCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := workspaces.CommonCliQueryDSLBuilder(c)
      entity := CastWidgetAreaFromCli(c)
      if entity, err := WidgetAreaActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func CastWidgetAreaFromCli (c *cli.Context) *WidgetAreaEntity {
	template := &WidgetAreaEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("name") {
        value := c.String("name")
        template.Name = &value
      }
      if c.IsSet("layouts") {
        value := c.String("layouts")
        template.Layouts = &value
      }
	return template
}
  func WidgetAreaSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      WidgetAreaActionCreate,
      reflect.ValueOf(&WidgetAreaEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func WidgetAreaImportMocks() {
    workspaces.SeederFromFSImport(
      workspaces.QueryDSL{},
      WidgetAreaActionCreate,
      reflect.ValueOf(&WidgetAreaEntity{}).Elem(),
      &mocks.ViewsFs,
      []string{},
      false,
    )
  }
  func WidgetAreaWriteQueryMock(ctx workspaces.MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := WidgetAreaActionQuery(f)
      result := workspaces.QueryEntitySuccessResult(f, items, count)
      workspaces.WriteMockDataToFile(lang, "", "WidgetArea", result)
    }
  }
var WidgetAreaImportExportCommands = []cli.Command{
	{
		Name:  "mock",
		Usage: "Generates mock records based on the entity definition",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "count",
				Usage: "how many activation key do you need to be generated and stored in database",
				Value: 10,
			},
		},
		Action: func(c *cli.Context) error {
			query := workspaces.CommonCliQueryDSLBuilder(c)
			WidgetAreaActionSeeder(query, c.Int("count"))
			return nil
		},
	},
	{
		Name:    "init",
		Aliases: []string{"i"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Usage: "The address of file you want the csv be exported to",
				Value: "widget-area-seeder.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Creates a basic seeder file for you, based on the definition module we have. You can populate this file as an example",
		Action: func(c *cli.Context) error {
			f := workspaces.CommonCliQueryDSLBuilder(c)
			WidgetAreaActionSeederInit(f, c.String("file"), c.String("format"))
			return nil
		},
	},
	{
		Name:    "validate",
		Aliases: []string{"v"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Usage: "Validates an import file, such as yaml, json, csv, and gives some insights how the after import it would look like",
				Value: "widget-area-seeder-widget-area.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of widget-areas, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]WidgetAreaEntity{}
			workspaces.ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
		cli.Command{
			Name:  "mocks",
			Usage: "Prints the list of mocks",
			Action: func(c *cli.Context) error {
				if entity, err := workspaces.GetSeederFilenames(&mocks.ViewsFs, ""); err != nil {
					fmt.Println(err.Error())
				} else {
					f, _ := json.MarshalIndent(entity, "", "  ")
					fmt.Println(string(f))
				}
				return nil
			},
		},
		cli.Command{
			Name:  "msync",
			Usage: "Tries to sync mocks into the system",
			Action: func(c *cli.Context) error {
				workspaces.CommonCliImportEmbedCmd(c,
					WidgetAreaActionCreate,
					reflect.ValueOf(&WidgetAreaEntity{}).Elem(),
					&mocks.ViewsFs,
				)
				return nil
			},
		},
	cli.Command{
		Name:    "import",
		Flags: append(workspaces.CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv be imported from",
				Required: true,
			}),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			workspaces.CommonCliImportCmd(c,
				WidgetAreaActionCreate,
				reflect.ValueOf(&WidgetAreaEntity{}).Elem(),
				c.String("file"),
			)
			return nil
		},
	},
}
    var WidgetAreaCliCommands []cli.Command = []cli.Command{
      workspaces.GetCommonQuery(WidgetAreaActionQuery),
      workspaces.GetCommonTableQuery(reflect.ValueOf(&WidgetAreaEntity{}).Elem(), WidgetAreaActionQuery),
          WidgetAreaCreateCmd,
          WidgetAreaUpdateCmd,
          WidgetAreaCreateInteractiveCmd,
          WidgetAreaWipeCmd,
          workspaces.GetCommonRemoveQuery(reflect.ValueOf(&WidgetAreaEntity{}).Elem(), WidgetAreaActionRemove),
  }
  func WidgetAreaCliFn() cli.Command {
    WidgetAreaCliCommands = append(WidgetAreaCliCommands, WidgetAreaImportExportCommands...)
    return cli.Command{
      Name:        "widgetArea",
      Description: "WidgetAreas module actions (sample module to handle complex entities)",
      Usage:       "Widget areas are groups of widgets, which can be placed on a special place such as dashboard",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: WidgetAreaCliCommands,
    }
  }
  /**
  *	Override this function on WidgetAreaEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendWidgetAreaRouter = func(r *[]workspaces.Module2Action) {}
  func GetWidgetAreaModule2Actions() []workspaces.Module2Action {
    routes := []workspaces.Module2Action{
       {
        Method: "GET",
        Url:    "/widget-areas",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_WIDGETAREA_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpQueryEntity(c, WidgetAreaActionQuery)
          },
        },
        Format: "QUERY",
        Action: WidgetAreaActionQuery,
        ResponseEntity: &[]WidgetAreaEntity{},
      },
      {
        Method: "GET",
        Url:    "/widget-areas/export",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_WIDGETAREA_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpStreamFileChannel(c, WidgetAreaActionExport)
          },
        },
        Format: "QUERY",
        Action: WidgetAreaActionExport,
        ResponseEntity: &[]WidgetAreaEntity{},
      },
      {
        Method: "GET",
        Url:    "/widget-area/:uniqueId",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_WIDGETAREA_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpGetEntity(c, WidgetAreaActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: WidgetAreaActionGetOne,
        ResponseEntity: &WidgetAreaEntity{},
      },
      {
        Method: "POST",
        Url:    "/widget-area",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_WIDGETAREA_CREATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpPostEntity(c, WidgetAreaActionCreate)
          },
        },
        Action: WidgetAreaActionCreate,
        Format: "POST_ONE",
        RequestEntity: &WidgetAreaEntity{},
        ResponseEntity: &WidgetAreaEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/widget-area",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_WIDGETAREA_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpUpdateEntity(c, WidgetAreaActionUpdate)
          },
        },
        Action: WidgetAreaActionUpdate,
        RequestEntity: &WidgetAreaEntity{},
        Format: "PATCH_ONE",
        ResponseEntity: &WidgetAreaEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/widget-areas",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_WIDGETAREA_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpUpdateEntities(c, WidgetAreaActionBulkUpdate)
          },
        },
        Action: WidgetAreaActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &workspaces.BulkRecordRequest[WidgetAreaEntity]{},
        ResponseEntity: &workspaces.BulkRecordRequest[WidgetAreaEntity]{},
      },
      {
        Method: "DELETE",
        Url:    "/widget-area",
        Format: "DELETE_DSL",
        SecurityModel: workspaces.SecurityModel{
          ActionRequires: []string{PERM_ROOT_WIDGETAREA_DELETE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            workspaces.HttpRemoveEntity(c, WidgetAreaActionRemove)
          },
        },
        Action: WidgetAreaActionRemove,
        RequestEntity: &workspaces.DeleteRequest{},
        ResponseEntity: &workspaces.DeleteResponse{},
        TargetEntity: &WidgetAreaEntity{},
      },
          {
            Method: "PATCH",
            Url:    "/widget-area/:linkerId/widgets/:uniqueId",
            SecurityModel: workspaces.SecurityModel{
              ActionRequires: []string{PERM_ROOT_WIDGETAREA_UPDATE},
            },
            Handlers: []gin.HandlerFunc{
              func (
                c *gin.Context,
              ) {
                workspaces.HttpUpdateEntity(c, WidgetAreaWidgetsActionUpdate)
              },
            },
            Action: WidgetAreaWidgetsActionUpdate,
            Format: "PATCH_ONE",
            RequestEntity: &WidgetAreaWidgets{},
            ResponseEntity: &WidgetAreaWidgets{},
          },
          {
            Method: "GET",
            Url:    "/widget-area/widgets/:linkerId/:uniqueId",
            SecurityModel: workspaces.SecurityModel{
              ActionRequires: []string{PERM_ROOT_WIDGETAREA_QUERY},
            },
            Handlers: []gin.HandlerFunc{
              func (
                c *gin.Context,
              ) {
                workspaces.HttpGetEntity(c, WidgetAreaWidgetsActionGetOne)
              },
            },
            Action: WidgetAreaWidgetsActionGetOne,
            Format: "GET_ONE",
            ResponseEntity: &WidgetAreaWidgets{},
          },
          {
            Method: "POST",
            Url:    "/widget-area/:linkerId/widgets",
            SecurityModel: workspaces.SecurityModel{
              ActionRequires: []string{PERM_ROOT_WIDGETAREA_CREATE},
            },
            Handlers: []gin.HandlerFunc{
              func (
                c *gin.Context,
              ) {
                workspaces.HttpPostEntity(c, WidgetAreaWidgetsActionCreate)
              },
            },
            Action: WidgetAreaWidgetsActionCreate,
            Format: "POST_ONE",
            RequestEntity: &WidgetAreaWidgets{},
            ResponseEntity: &WidgetAreaWidgets{},
          },
    }
    // Append user defined functions
    AppendWidgetAreaRouter(&routes)
    return routes
  }
  func CreateWidgetAreaRouter(r *gin.Engine) []workspaces.Module2Action {
    httpRoutes := GetWidgetAreaModule2Actions()
    workspaces.CastRoutes(httpRoutes, r)
    workspaces.WriteHttpInformationToFile(&httpRoutes, WidgetAreaEntityJsonSchema, "widget-area-http", "widget")
    workspaces.WriteEntitySchema("WidgetAreaEntity", WidgetAreaEntityJsonSchema, "widget")
    return httpRoutes
  }
var PERM_ROOT_WIDGETAREA_DELETE = "root/widgetarea/delete"
var PERM_ROOT_WIDGETAREA_CREATE = "root/widgetarea/create"
var PERM_ROOT_WIDGETAREA_UPDATE = "root/widgetarea/update"
var PERM_ROOT_WIDGETAREA_QUERY = "root/widgetarea/query"
var PERM_ROOT_WIDGETAREA = "root/widgetarea"
var ALL_WIDGETAREA_PERMISSIONS = []string{
	PERM_ROOT_WIDGETAREA_DELETE,
	PERM_ROOT_WIDGETAREA_CREATE,
	PERM_ROOT_WIDGETAREA_UPDATE,
	PERM_ROOT_WIDGETAREA_QUERY,
	PERM_ROOT_WIDGETAREA,
}