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
type WorkspaceConfigEntity struct {
    Visibility       *string                         `json:"visibility,omitempty" yaml:"visibility"`
    WorkspaceId      *string                         `json:"workspaceId,omitempty" yaml:"workspaceId" gorm:"unique;not null;" `
    LinkerId         *string                         `json:"linkerId,omitempty" yaml:"linkerId"`
    ParentId         *string                         `json:"parentId,omitempty" yaml:"parentId"`
    UniqueId         string                          `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
    UserId           *string                         `json:"userId,omitempty" yaml:"userId"`
    Rank             int64                           `json:"rank,omitempty" gorm:"type:int;name:rank"`
    Updated          int64                           `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
    Created          int64                           `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
    CreatedFormatted string                          `json:"createdFormatted,omitempty" sql:"-"`
    UpdatedFormatted string                          `json:"updatedFormatted,omitempty" sql:"-"`
    DisablePublicWorkspaceCreation   *int64 `json:"disablePublicWorkspaceCreation" yaml:"disablePublicWorkspaceCreation"       `
    // Datenano also has a text representation
    Workspace   *  WorkspaceEntity `json:"workspace" yaml:"workspace"    gorm:"foreignKey:WorkspaceId;references:UniqueId"     `
    // Datenano also has a text representation
    ZoomClientId   *string `json:"zoomClientId" yaml:"zoomClientId"       `
    // Datenano also has a text representation
    ZoomClientSecret   *string `json:"zoomClientSecret" yaml:"zoomClientSecret"       `
    // Datenano also has a text representation
    AllowPublicToJoinTheWorkspace   *bool `json:"allowPublicToJoinTheWorkspace" yaml:"allowPublicToJoinTheWorkspace"       `
    // Datenano also has a text representation
    Children []*WorkspaceConfigEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *WorkspaceConfigEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var WorkspaceConfigPreloadRelations []string = []string{}
var WORKSPACECONFIG_EVENT_CREATED = "workspaceConfig.created"
var WORKSPACECONFIG_EVENT_UPDATED = "workspaceConfig.updated"
var WORKSPACECONFIG_EVENT_DELETED = "workspaceConfig.deleted"
var WORKSPACECONFIG_EVENTS = []string{
	WORKSPACECONFIG_EVENT_CREATED,
	WORKSPACECONFIG_EVENT_UPDATED,
	WORKSPACECONFIG_EVENT_DELETED,
}
type WorkspaceConfigFieldMap struct {
		DisablePublicWorkspaceCreation TranslatedString `yaml:"disablePublicWorkspaceCreation"`
		Workspace TranslatedString `yaml:"workspace"`
		ZoomClientId TranslatedString `yaml:"zoomClientId"`
		ZoomClientSecret TranslatedString `yaml:"zoomClientSecret"`
		AllowPublicToJoinTheWorkspace TranslatedString `yaml:"allowPublicToJoinTheWorkspace"`
}
var WorkspaceConfigEntityMetaConfig map[string]int64 = map[string]int64{
}
var WorkspaceConfigEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&WorkspaceConfigEntity{}))
func entityWorkspaceConfigFormatter(dto *WorkspaceConfigEntity, query QueryDSL) {
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
func WorkspaceConfigMockEntity() *WorkspaceConfigEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &WorkspaceConfigEntity{
      DisablePublicWorkspaceCreation : &int64Holder,
      ZoomClientId : &stringHolder,
      ZoomClientSecret : &stringHolder,
	}
	return entity
}
func WorkspaceConfigActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := WorkspaceConfigMockEntity()
		_, err := WorkspaceConfigActionCreate(entity, query)
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
  func WorkspaceConfigActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*WorkspaceConfigEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &WorkspaceConfigEntity{
          ZoomClientId: &tildaRef,
          ZoomClientSecret: &tildaRef,
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
  func WorkspaceConfigAssociationCreate(dto *WorkspaceConfigEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func WorkspaceConfigRelationContentCreate(dto *WorkspaceConfigEntity, query QueryDSL) error {
return nil
}
func WorkspaceConfigRelationContentUpdate(dto *WorkspaceConfigEntity, query QueryDSL) error {
	return nil
}
func WorkspaceConfigPolyglotCreateHandler(dto *WorkspaceConfigEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func WorkspaceConfigValidator(dto *WorkspaceConfigEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func WorkspaceConfigEntityPreSanitize(dto *WorkspaceConfigEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func WorkspaceConfigEntityBeforeCreateAppend(dto *WorkspaceConfigEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    WorkspaceConfigRecursiveAddUniqueId(dto, query)
  }
  func WorkspaceConfigRecursiveAddUniqueId(dto *WorkspaceConfigEntity, query QueryDSL) {
  }
func WorkspaceConfigActionBatchCreateFn(dtos []*WorkspaceConfigEntity, query QueryDSL) ([]*WorkspaceConfigEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*WorkspaceConfigEntity{}
		for _, item := range dtos {
			s, err := WorkspaceConfigActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func WorkspaceConfigActionCreateFn(dto *WorkspaceConfigEntity, query QueryDSL) (*WorkspaceConfigEntity, *IError) {
	// 1. Validate always
	if iError := WorkspaceConfigValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	WorkspaceConfigEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	WorkspaceConfigEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	WorkspaceConfigPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	WorkspaceConfigRelationContentCreate(dto, query)
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
	WorkspaceConfigAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(WORKSPACECONFIG_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&WorkspaceConfigEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func WorkspaceConfigActionGetOne(query QueryDSL) (*WorkspaceConfigEntity, *IError) {
    refl := reflect.ValueOf(&WorkspaceConfigEntity{})
    item, err := GetOneEntity[WorkspaceConfigEntity](query, refl)
    entityWorkspaceConfigFormatter(item, query)
    return item, err
  }
  func WorkspaceConfigActionQuery(query QueryDSL) ([]*WorkspaceConfigEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&WorkspaceConfigEntity{})
    items, meta, err := QueryEntitiesPointer[WorkspaceConfigEntity](query, refl)
    for _, item := range items {
      entityWorkspaceConfigFormatter(item, query)
    }
    return items, meta, err
  }
  func WorkspaceConfigUpdateExec(dbref *gorm.DB, query QueryDSL, fields *WorkspaceConfigEntity) (*WorkspaceConfigEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = WORKSPACECONFIG_EVENT_UPDATED
    WorkspaceConfigEntityPreSanitize(fields, query)
    var item WorkspaceConfigEntity
    q := dbref.
      Where(&WorkspaceConfigEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    WorkspaceConfigRelationContentUpdate(fields, query)
    WorkspaceConfigPolyglotCreateHandler(fields, query)
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&WorkspaceConfigEntity{UniqueId: uniqueId}).
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
  func WorkspaceConfigActionUpdateFn(query QueryDSL, fields *WorkspaceConfigEntity) (*WorkspaceConfigEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := WorkspaceConfigValidator(fields, true); iError != nil {
      return nil, iError
    }
    WorkspaceConfigRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        _, err := WorkspaceConfigUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return nil, CastToIError(vf)
    } else {
      dbref = query.Tx
      return WorkspaceConfigUpdateExec(dbref, query, fields)
    }
  }
var WorkspaceConfigWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire workspaceconfigs ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilder(c)
		count, _ := WorkspaceConfigActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func WorkspaceConfigActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&WorkspaceConfigEntity{})
	query.ActionRequires = []string{PERM_ROOT_WORKSPACECONFIG_DELETE}
	return RemoveEntity[WorkspaceConfigEntity](query, refl)
}
func WorkspaceConfigActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[WorkspaceConfigEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'WorkspaceConfigEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func WorkspaceConfigActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[WorkspaceConfigEntity]) (
    *BulkRecordRequest[WorkspaceConfigEntity], *IError,
  ) {
    result := []*WorkspaceConfigEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := WorkspaceConfigActionUpdate(query, record)
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
func (x *WorkspaceConfigEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var WorkspaceConfigEntityMeta = TableMetaData{
	EntityName:    "WorkspaceConfig",
	ExportKey:    "workspace-configs",
	TableNameInDb: "fb_workspaceconfig_entities",
	EntityObject:  &WorkspaceConfigEntity{},
	ExportStream: WorkspaceConfigActionExportT,
	ImportQuery: WorkspaceConfigActionImport,
}
func WorkspaceConfigActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[WorkspaceConfigEntity](query, WorkspaceConfigActionQuery, WorkspaceConfigPreloadRelations)
}
func WorkspaceConfigActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[WorkspaceConfigEntity](query, WorkspaceConfigActionQuery, WorkspaceConfigPreloadRelations)
}
func WorkspaceConfigActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content WorkspaceConfigEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := WorkspaceConfigActionCreate(&content, query)
	return err
}
var WorkspaceConfigCommonCliFlags = []cli.Flag{
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
    &cli.Int64Flag{
      Name:     "disable-public-workspace-creation",
      Required: false,
      Usage:    "disablePublicWorkspaceCreation",
      Value: 1,
    },
    &cli.StringFlag{
      Name:     "workspace-id",
      Required: false,
      Usage:    "workspace",
    },
    &cli.StringFlag{
      Name:     "zoom-client-id",
      Required: false,
      Usage:    "zoomClientId",
    },
    &cli.StringFlag{
      Name:     "zoom-client-secret",
      Required: false,
      Usage:    "zoomClientSecret",
    },
    &cli.BoolFlag{
      Name:     "allow-public-to-join-the-workspace",
      Required: false,
      Usage:    "allowPublicToJoinTheWorkspace",
    },
}
var WorkspaceConfigCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "disablePublicWorkspaceCreation",
		StructField:     "DisablePublicWorkspaceCreation",
		Required: false,
		Usage:    "disablePublicWorkspaceCreation",
		Type: "int64",
	},
	{
		Name:     "zoomClientId",
		StructField:     "ZoomClientId",
		Required: false,
		Usage:    "zoomClientId",
		Type: "string",
	},
	{
		Name:     "zoomClientSecret",
		StructField:     "ZoomClientSecret",
		Required: false,
		Usage:    "zoomClientSecret",
		Type: "string",
	},
	{
		Name:     "allowPublicToJoinTheWorkspace",
		StructField:     "AllowPublicToJoinTheWorkspace",
		Required: false,
		Usage:    "allowPublicToJoinTheWorkspace",
		Type: "bool",
	},
}
var WorkspaceConfigCommonCliFlagsOptional = []cli.Flag{
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
    &cli.Int64Flag{
      Name:     "disable-public-workspace-creation",
      Required: false,
      Usage:    "disablePublicWorkspaceCreation",
      Value: 1,
    },
    &cli.StringFlag{
      Name:     "workspace-id",
      Required: false,
      Usage:    "workspace",
    },
    &cli.StringFlag{
      Name:     "zoom-client-id",
      Required: false,
      Usage:    "zoomClientId",
    },
    &cli.StringFlag{
      Name:     "zoom-client-secret",
      Required: false,
      Usage:    "zoomClientSecret",
    },
    &cli.BoolFlag{
      Name:     "allow-public-to-join-the-workspace",
      Required: false,
      Usage:    "allowPublicToJoinTheWorkspace",
    },
}
  var WorkspaceConfigCreateCmd cli.Command = cli.Command{
    Name:    "create",
    Aliases: []string{"c"},
    Flags: WorkspaceConfigCommonCliFlags,
    Usage: "Create a new template",
    Action: func(c *cli.Context) {
      query := CommonCliQueryDSLBuilder(c)
      entity := CastWorkspaceConfigFromCli(c)
      if entity, err := WorkspaceConfigActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var WorkspaceConfigCreateInteractiveCmd cli.Command = cli.Command{
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
      entity := &WorkspaceConfigEntity{}
      for _, item := range WorkspaceConfigCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := WorkspaceConfigActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var WorkspaceConfigUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: WorkspaceConfigCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilder(c)
      entity := CastWorkspaceConfigFromCli(c)
      if entity, err := WorkspaceConfigActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func CastWorkspaceConfigFromCli (c *cli.Context) *WorkspaceConfigEntity {
	template := &WorkspaceConfigEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("workspace-id") {
        value := c.String("workspace-id")
        template.WorkspaceId = &value
      }
      if c.IsSet("zoom-client-id") {
        value := c.String("zoom-client-id")
        template.ZoomClientId = &value
      }
      if c.IsSet("zoom-client-secret") {
        value := c.String("zoom-client-secret")
        template.ZoomClientSecret = &value
      }
	return template
}
  func WorkspaceConfigSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      WorkspaceConfigActionCreate,
      reflect.ValueOf(&WorkspaceConfigEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func WorkspaceConfigWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := WorkspaceConfigActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "WorkspaceConfig", result)
    }
  }
var WorkspaceConfigImportExportCommands = []cli.Command{
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
			WorkspaceConfigActionSeeder(query, c.Int("count"))
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
				Value: "workspace-config-seeder.yml",
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
			WorkspaceConfigActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "workspace-config-seeder-workspace-config.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of workspace-configs, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]WorkspaceConfigEntity{}
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
				WorkspaceConfigActionCreate,
				reflect.ValueOf(&WorkspaceConfigEntity{}).Elem(),
				c.String("file"),
			)
			return nil
		},
	},
}
    var WorkspaceConfigCliCommands []cli.Command = []cli.Command{
      GetCommonQuery(WorkspaceConfigActionQuery),
      GetCommonTableQuery(reflect.ValueOf(&WorkspaceConfigEntity{}).Elem(), WorkspaceConfigActionQuery),
          WorkspaceConfigCreateCmd,
          WorkspaceConfigUpdateCmd,
          WorkspaceConfigCreateInteractiveCmd,
          WorkspaceConfigWipeCmd,
          GetCommonRemoveQuery(reflect.ValueOf(&WorkspaceConfigEntity{}).Elem(), WorkspaceConfigActionRemove),
  }
  func WorkspaceConfigCliFn() cli.Command {
    WorkspaceConfigCliCommands = append(WorkspaceConfigCliCommands, WorkspaceConfigImportExportCommands...)
    return cli.Command{
      Name:        "config",
      Description: "WorkspaceConfigs module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: WorkspaceConfigCliCommands,
    }
  }
  /**
  *	Override this function on WorkspaceConfigEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendWorkspaceConfigRouter = func(r *[]Module2Action) {}
  func GetWorkspaceConfigModule2Actions() []Module2Action {
    routes := []Module2Action{
       {
        Method: "GET",
        Url:    "/workspace-configs",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_WORKSPACECONFIG_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpQueryEntity(c, WorkspaceConfigActionQuery)
          },
        },
        Format: "QUERY",
        Action: WorkspaceConfigActionQuery,
        ResponseEntity: &[]WorkspaceConfigEntity{},
      },
      {
        Method: "GET",
        Url:    "/workspace-configs/export",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_WORKSPACECONFIG_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpStreamFileChannel(c, WorkspaceConfigActionExport)
          },
        },
        Format: "QUERY",
        Action: WorkspaceConfigActionExport,
        ResponseEntity: &[]WorkspaceConfigEntity{},
      },
      {
        Method: "GET",
        Url:    "/workspace-config/:uniqueId",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_WORKSPACECONFIG_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpGetEntity(c, WorkspaceConfigActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: WorkspaceConfigActionGetOne,
        ResponseEntity: &WorkspaceConfigEntity{},
      },
      {
        Method: "POST",
        Url:    "/workspace-config",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_WORKSPACECONFIG_CREATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpPostEntity(c, WorkspaceConfigActionCreate)
          },
        },
        Action: WorkspaceConfigActionCreate,
        Format: "POST_ONE",
        RequestEntity: &WorkspaceConfigEntity{},
        ResponseEntity: &WorkspaceConfigEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/workspace-config",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_WORKSPACECONFIG_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpUpdateEntity(c, WorkspaceConfigActionUpdate)
          },
        },
        Action: WorkspaceConfigActionUpdate,
        RequestEntity: &WorkspaceConfigEntity{},
        Format: "PATCH_ONE",
        ResponseEntity: &WorkspaceConfigEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/workspace-configs",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_WORKSPACECONFIG_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpUpdateEntities(c, WorkspaceConfigActionBulkUpdate)
          },
        },
        Action: WorkspaceConfigActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &BulkRecordRequest[WorkspaceConfigEntity]{},
        ResponseEntity: &BulkRecordRequest[WorkspaceConfigEntity]{},
      },
      {
        Method: "DELETE",
        Url:    "/workspace-config",
        Format: "DELETE_DSL",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_WORKSPACECONFIG_DELETE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpRemoveEntity(c, WorkspaceConfigActionRemove)
          },
        },
        Action: WorkspaceConfigActionRemove,
        RequestEntity: &DeleteRequest{},
        ResponseEntity: &DeleteResponse{},
        TargetEntity: &WorkspaceConfigEntity{},
      },
          {
            Method: "PATCH",
            Url:    "/workspace-config/distinct",
            SecurityModel: SecurityModel{
              ActionRequires: []string{PERM_ROOT_WORKSPACECONFIG_UPDATE_DISTINCT_WORKSPACE},
            },
            Handlers: []gin.HandlerFunc{
              func (c *gin.Context) {
                HttpUpdateEntity(c, WorkspaceConfigDistinctActionUpdate)
              },
            },
            Action: WorkspaceConfigDistinctActionUpdate,
            Format: "PATCH_ONE",
            RequestEntity: &WorkspaceConfigEntity{},
            ResponseEntity: &WorkspaceConfigEntity{},
          },
          {
            Method: "GET",
            Url:    "/workspace-config/distinct",
            SecurityModel: SecurityModel{
              ActionRequires: []string{PERM_ROOT_WORKSPACECONFIG_GET_DISTINCT_WORKSPACE},
            },
            Handlers: []gin.HandlerFunc{
              func (c *gin.Context) {
                HttpGetEntity(c, WorkspaceConfigDistinctActionGetOne)
              },
            },
            Action: WorkspaceConfigDistinctActionGetOne,
            Format: "GET_ONE",
            ResponseEntity: &WorkspaceConfigEntity{},
          },
    }
    // Append user defined functions
    AppendWorkspaceConfigRouter(&routes)
    return routes
  }
  func CreateWorkspaceConfigRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetWorkspaceConfigModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, WorkspaceConfigEntityJsonSchema, "workspace-config-http", "workspaces")
    WriteEntitySchema("WorkspaceConfigEntity", WorkspaceConfigEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_WORKSPACECONFIG_DELETE = "root/workspaceconfig/delete"
var PERM_ROOT_WORKSPACECONFIG_CREATE = "root/workspaceconfig/create"
var PERM_ROOT_WORKSPACECONFIG_UPDATE = "root/workspaceconfig/update"
var PERM_ROOT_WORKSPACECONFIG_QUERY = "root/workspaceconfig/query"
  var PERM_ROOT_WORKSPACECONFIG_GET_DISTINCT_WORKSPACE = "root/workspaceconfig/get-distinct-workspace"
  var PERM_ROOT_WORKSPACECONFIG_UPDATE_DISTINCT_WORKSPACE = "root/workspaceconfig/update-distinct-workspace"
var PERM_ROOT_WORKSPACECONFIG = "root/workspaceconfig"
var ALL_WORKSPACECONFIG_PERMISSIONS = []string{
	PERM_ROOT_WORKSPACECONFIG_DELETE,
	PERM_ROOT_WORKSPACECONFIG_CREATE,
	PERM_ROOT_WORKSPACECONFIG_UPDATE,
    PERM_ROOT_WORKSPACECONFIG_GET_DISTINCT_WORKSPACE,
    PERM_ROOT_WORKSPACECONFIG_UPDATE_DISTINCT_WORKSPACE,
	PERM_ROOT_WORKSPACECONFIG_QUERY,
	PERM_ROOT_WORKSPACECONFIG,
}
  func WorkspaceConfigDistinctActionUpdate(
    query QueryDSL,
    fields *WorkspaceConfigEntity,
  ) (*WorkspaceConfigEntity, *IError) {
    query.UniqueId = query.UserId
    entity, err := WorkspaceConfigActionGetOne(query)
    if err != nil || entity.UniqueId == "" {
      fields.UniqueId = query.UserId
      return WorkspaceConfigActionCreateFn(fields, query)
    } else {
      fields.UniqueId = query.UniqueId
      return WorkspaceConfigActionUpdateFn(query, fields)
    }
  }
  func WorkspaceConfigDistinctActionGetOne(
    query QueryDSL,
  ) (*WorkspaceConfigEntity, *IError) {
    query.UniqueId = query.UserId
    entity, err := WorkspaceConfigActionGetOne(query)
    if err != nil && err.HttpCode == 404 {
      return &WorkspaceConfigEntity{}, nil
    }
    return entity, err
  }