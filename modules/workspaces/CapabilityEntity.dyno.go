package workspaces
import (
    "github.com/gin-gonic/gin"
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
)
type CapabilityEntity struct {
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
    Name   *string `json:"name" yaml:"name"       `
    // Datenano also has a text representation
    Children []*CapabilityEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *CapabilityEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var CapabilityPreloadRelations []string = []string{}
var CAPABILITY_EVENT_CREATED = "capability.created"
var CAPABILITY_EVENT_UPDATED = "capability.updated"
var CAPABILITY_EVENT_DELETED = "capability.deleted"
var CAPABILITY_EVENTS = []string{
	CAPABILITY_EVENT_CREATED,
	CAPABILITY_EVENT_UPDATED,
	CAPABILITY_EVENT_DELETED,
}
type CapabilityFieldMap struct {
		Name TranslatedString `yaml:"name"`
}
var CapabilityEntityMetaConfig map[string]int64 = map[string]int64{
}
var CapabilityEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&CapabilityEntity{}))
func entityCapabilityFormatter(dto *CapabilityEntity, query QueryDSL) {
	if dto == nil {
		return
	}
	if dto.Created > 0 {
		dto.CreatedFormatted = FormatDateBasedOnQuery(dto.Created, query)
	}
	if dto.Updated > 0 {
		dto.CreatedFormatted = FormatDateBasedOnQuery(dto.Updated, query)
	}
}
func CapabilityMockEntity() *CapabilityEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &CapabilityEntity{
      Name : &stringHolder,
	}
	return entity
}
func CapabilityActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := CapabilityMockEntity()
		_, err := CapabilityActionCreate(entity, query)
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
  func CapabilityActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*CapabilityEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &CapabilityEntity{
          Name: &tildaRef,
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
  func CapabilityAssociationCreate(dto *CapabilityEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func CapabilityRelationContentCreate(dto *CapabilityEntity, query QueryDSL) error {
return nil
}
func CapabilityRelationContentUpdate(dto *CapabilityEntity, query QueryDSL) error {
	return nil
}
func CapabilityPolyglotCreateHandler(dto *CapabilityEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func CapabilityValidator(dto *CapabilityEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func CapabilityEntityPreSanitize(dto *CapabilityEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func CapabilityEntityBeforeCreateAppend(dto *CapabilityEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    CapabilityRecursiveAddUniqueId(dto, query)
  }
  func CapabilityRecursiveAddUniqueId(dto *CapabilityEntity, query QueryDSL) {
  }
func CapabilityActionBatchCreateFn(dtos []*CapabilityEntity, query QueryDSL) ([]*CapabilityEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*CapabilityEntity{}
		for _, item := range dtos {
			s, err := CapabilityActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func CapabilityActionCreateFn(dto *CapabilityEntity, query QueryDSL) (*CapabilityEntity, *IError) {
	// 1. Validate always
	if iError := CapabilityValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	CapabilityEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	CapabilityEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	CapabilityPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	CapabilityRelationContentCreate(dto, query)
	// 4. Create the entity
	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
	} else {
		dbref = query.Tx
	}
	query.Tx = dbref;
	err := dbref.Create(&dto).Error
	if err != nil {
		err := GormErrorToIError(err)
		return dto, err
	}
	// 5. Create sub entities, objects or arrays, association to other entities
	CapabilityAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(CAPABILITY_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&CapabilityEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func CapabilityActionGetOne(query QueryDSL) (*CapabilityEntity, *IError) {
    refl := reflect.ValueOf(&CapabilityEntity{})
    item, err := GetOneEntity[CapabilityEntity](query, refl)
    entityCapabilityFormatter(item, query)
    return item, err
  }
  func CapabilityActionQuery(query QueryDSL) ([]*CapabilityEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&CapabilityEntity{})
    items, meta, err := QueryEntitiesPointer[CapabilityEntity](query, refl)
    for _, item := range items {
      entityCapabilityFormatter(item, query)
    }
    return items, meta, err
  }
  func CapabilityUpdateExec(dbref *gorm.DB, query QueryDSL, fields *CapabilityEntity) (*CapabilityEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = CAPABILITY_EVENT_UPDATED
    CapabilityEntityPreSanitize(fields, query)
    var item CapabilityEntity
    q := dbref.
      Where(&CapabilityEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    CapabilityRelationContentUpdate(fields, query)
    CapabilityPolyglotCreateHandler(fields, query)
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&CapabilityEntity{UniqueId: uniqueId}).
      First(&item).Error
    event.MustFire(query.TriggerEventName, event.M{
      "entity":   &item,
      "target":   "workspace",
      "unqiueId": query.WorkspaceId,
    })
    if err != nil {
      return &item, GormErrorToIError(err)
    }
    return &item, nil
  }
  func CapabilityActionUpdateFn(query QueryDSL, fields *CapabilityEntity) (*CapabilityEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := CapabilityValidator(fields, true); iError != nil {
      return nil, iError
    }
    CapabilityRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        _, err := CapabilityUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return nil, CastToIError(vf)
    } else {
      dbref = query.Tx
      return CapabilityUpdateExec(dbref, query, fields)
    }
  }
var CapabilityWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire capabilities ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilder(c)
		count, _ := CapabilityActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func CapabilityActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&CapabilityEntity{})
	query.ActionRequires = []string{PERM_ROOT_CAPABILITY_DELETE}
	return RemoveEntity[CapabilityEntity](query, refl)
}
func CapabilityActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[CapabilityEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'CapabilityEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func CapabilityActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[CapabilityEntity]) (
    *BulkRecordRequest[CapabilityEntity], *IError,
  ) {
    result := []*CapabilityEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := CapabilityActionUpdate(query, record)
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
    return nil, err.(*IError)
  }
func (x *CapabilityEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var CapabilityEntityMeta = TableMetaData{
	EntityName:    "Capability",
	ExportKey:    "capabilities",
	TableNameInDb: "fb_capability_entities",
	EntityObject:  &CapabilityEntity{},
	ExportStream: CapabilityActionExportT,
	ImportQuery: CapabilityActionImport,
}
func CapabilityActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[CapabilityEntity](query, CapabilityActionQuery, CapabilityPreloadRelations)
}
func CapabilityActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[CapabilityEntity](query, CapabilityActionQuery, CapabilityPreloadRelations)
}
func CapabilityActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content CapabilityEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := CapabilityActionCreate(&content, query)
	return err
}
var CapabilityCommonCliFlags = []cli.Flag{
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
}
var CapabilityCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "name",
		StructField:     "Name",
		Required: false,
		Usage:    "name",
		Type: "string",
	},
}
var CapabilityCommonCliFlagsOptional = []cli.Flag{
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
}
  var CapabilityCreateCmd cli.Command = cli.Command{
    Name:    "create",
    Aliases: []string{"c"},
    Flags: CapabilityCommonCliFlags,
    Usage: "Create a new template",
    Action: func(c *cli.Context) {
      query := CommonCliQueryDSLBuilder(c)
      entity := CastCapabilityFromCli(c)
      if entity, err := CapabilityActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var CapabilityCreateInteractiveCmd cli.Command = cli.Command{
    Name:  "ic",
    Usage: "Creates a new template, using requied fields in an interactive name",
    Flags: []cli.Flag{
      &cli.BoolFlag{
        Name:  "all",
        Usage: "Interactively asks for all inputs, not only required ones",
      },
    },
    Action: func(c *cli.Context) {
      query := CommonCliQueryDSLBuilder(c)
      entity := &CapabilityEntity{}
      for _, item := range CapabilityCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := CapabilityActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var CapabilityUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: CapabilityCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilder(c)
      entity := CastCapabilityFromCli(c)
      if entity, err := CapabilityActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func CastCapabilityFromCli (c *cli.Context) *CapabilityEntity {
	template := &CapabilityEntity{}
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
	return template
}
  func CapabilitySyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      CapabilityActionCreate,
      reflect.ValueOf(&CapabilityEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func CapabilityWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := CapabilityActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "Capability", result)
    }
  }
var CapabilityImportExportCommands = []cli.Command{
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
			query := CommonCliQueryDSLBuilder(c)
			CapabilityActionSeeder(query, c.Int("count"))
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
				Value: "capability-seeder.yml",
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
			f := CommonCliQueryDSLBuilder(c)
			CapabilityActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "capability-seeder-capability.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of capabilities, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]CapabilityEntity{}
			ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:    "import",
		Flags: append(CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv be imported from",
				Required: true,
			}),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmd(c,
				CapabilityActionCreate,
				reflect.ValueOf(&CapabilityEntity{}).Elem(),
				c.String("file"),
			)
			return nil
		},
	},
}
    var CapabilityCliCommands []cli.Command = []cli.Command{
      GetCommonQuery(CapabilityActionQuery),
      GetCommonTableQuery(reflect.ValueOf(&CapabilityEntity{}).Elem(), CapabilityActionQuery),
          CapabilityCreateCmd,
          CapabilityUpdateCmd,
          CapabilityCreateInteractiveCmd,
          CapabilityWipeCmd,
          GetCommonRemoveQuery(reflect.ValueOf(&CapabilityEntity{}).Elem(), CapabilityActionRemove),
  }
  func CapabilityCliFn() cli.Command {
    CapabilityCliCommands = append(CapabilityCliCommands, CapabilityImportExportCommands...)
    return cli.Command{
      Name:        "capability",
      ShortName:   "cap",
      Description: "Capabilitys module actions (sample module to handle complex entities)",
      Usage:       "Manage the capabilities inside the application, both builtin to core and custom defined ones",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: CapabilityCliCommands,
    }
  }
  /**
  *	Override this function on CapabilityEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendCapabilityRouter = func(r *[]Module2Action) {}
  func GetCapabilityModule2Actions() []Module2Action {
    routes := []Module2Action{
       {
        Method: "GET",
        Url:    "/capabilities",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_CAPABILITY_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpQueryEntity(c, CapabilityActionQuery)
          },
        },
        Format: "QUERY",
        Action: CapabilityActionQuery,
        ResponseEntity: &[]CapabilityEntity{},
      },
      {
        Method: "GET",
        Url:    "/capabilities/export",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_CAPABILITY_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpStreamFileChannel(c, CapabilityActionExport)
          },
        },
        Format: "QUERY",
        Action: CapabilityActionExport,
        ResponseEntity: &[]CapabilityEntity{},
      },
      {
        Method: "GET",
        Url:    "/capability/:uniqueId",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_CAPABILITY_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpGetEntity(c, CapabilityActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: CapabilityActionGetOne,
        ResponseEntity: &CapabilityEntity{},
      },
      {
        Method: "POST",
        Url:    "/capability",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_CAPABILITY_CREATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpPostEntity(c, CapabilityActionCreate)
          },
        },
        Action: CapabilityActionCreate,
        Format: "POST_ONE",
        RequestEntity: &CapabilityEntity{},
        ResponseEntity: &CapabilityEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/capability",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_CAPABILITY_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpUpdateEntity(c, CapabilityActionUpdate)
          },
        },
        Action: CapabilityActionUpdate,
        RequestEntity: &CapabilityEntity{},
        Format: "PATCH_ONE",
        ResponseEntity: &CapabilityEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/capabilities",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_CAPABILITY_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpUpdateEntities(c, CapabilityActionBulkUpdate)
          },
        },
        Action: CapabilityActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &BulkRecordRequest[CapabilityEntity]{},
        ResponseEntity: &BulkRecordRequest[CapabilityEntity]{},
      },
      {
        Method: "DELETE",
        Url:    "/capability",
        Format: "DELETE_DSL",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_CAPABILITY_DELETE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpRemoveEntity(c, CapabilityActionRemove)
          },
        },
        Action: CapabilityActionRemove,
        RequestEntity: &DeleteRequest{},
        ResponseEntity: &DeleteResponse{},
        TargetEntity: &CapabilityEntity{},
      },
    }
    // Append user defined functions
    AppendCapabilityRouter(&routes)
    return routes
  }
  func CreateCapabilityRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetCapabilityModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, CapabilityEntityJsonSchema, "capability-http", "workspaces")
    WriteEntitySchema("CapabilityEntity", CapabilityEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_CAPABILITY_DELETE = "root/capability/delete"
var PERM_ROOT_CAPABILITY_CREATE = "root/capability/create"
var PERM_ROOT_CAPABILITY_UPDATE = "root/capability/update"
var PERM_ROOT_CAPABILITY_QUERY = "root/capability/query"
var PERM_ROOT_CAPABILITY = "root/capability"
var ALL_CAPABILITY_PERMISSIONS = []string{
	PERM_ROOT_CAPABILITY_DELETE,
	PERM_ROOT_CAPABILITY_CREATE,
	PERM_ROOT_CAPABILITY_UPDATE,
	PERM_ROOT_CAPABILITY_QUERY,
	PERM_ROOT_CAPABILITY,
}