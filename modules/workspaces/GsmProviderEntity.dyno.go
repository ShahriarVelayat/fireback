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
	mocks "pixelplux.com/fireback/modules/workspaces/mocks/GsmProvider"
)
type GsmProviderEntity struct {
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
    ApiKey   *string `json:"apiKey" yaml:"apiKey"       `
    // Datenano also has a text representation
    MainSenderNumber   *string `json:"mainSenderNumber" yaml:"mainSenderNumber"  validate:"required"       `
    // Datenano also has a text representation
    Type   *string `json:"type" yaml:"type"  validate:"required"       `
    // Datenano also has a text representation
    InvokeUrl   *string `json:"invokeUrl" yaml:"invokeUrl"       `
    // Datenano also has a text representation
    InvokeBody   *string `json:"invokeBody" yaml:"invokeBody"       `
    // Datenano also has a text representation
    Children []*GsmProviderEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *GsmProviderEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var GsmProviderPreloadRelations []string = []string{}
var GSMPROVIDER_EVENT_CREATED = "gsmProvider.created"
var GSMPROVIDER_EVENT_UPDATED = "gsmProvider.updated"
var GSMPROVIDER_EVENT_DELETED = "gsmProvider.deleted"
var GSMPROVIDER_EVENTS = []string{
	GSMPROVIDER_EVENT_CREATED,
	GSMPROVIDER_EVENT_UPDATED,
	GSMPROVIDER_EVENT_DELETED,
}
type GsmProviderFieldMap struct {
		ApiKey TranslatedString `yaml:"apiKey"`
		MainSenderNumber TranslatedString `yaml:"mainSenderNumber"`
		Type TranslatedString `yaml:"type"`
		InvokeUrl TranslatedString `yaml:"invokeUrl"`
		InvokeBody TranslatedString `yaml:"invokeBody"`
}
var GsmProviderEntityMetaConfig map[string]int64 = map[string]int64{
}
var GsmProviderEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&GsmProviderEntity{}))
func entityGsmProviderFormatter(dto *GsmProviderEntity, query QueryDSL) {
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
func GsmProviderMockEntity() *GsmProviderEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &GsmProviderEntity{
      ApiKey : &stringHolder,
      MainSenderNumber : &stringHolder,
      Type : &stringHolder,
      InvokeUrl : &stringHolder,
      InvokeBody : &stringHolder,
	}
	return entity
}
func GsmProviderActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := GsmProviderMockEntity()
		_, err := GsmProviderActionCreate(entity, query)
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
  func GsmProviderActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*GsmProviderEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &GsmProviderEntity{
          ApiKey: &tildaRef,
          MainSenderNumber: &tildaRef,
          Type: &tildaRef,
          InvokeUrl: &tildaRef,
          InvokeBody: &tildaRef,
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
  func GsmProviderAssociationCreate(dto *GsmProviderEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func GsmProviderRelationContentCreate(dto *GsmProviderEntity, query QueryDSL) error {
return nil
}
func GsmProviderRelationContentUpdate(dto *GsmProviderEntity, query QueryDSL) error {
	return nil
}
func GsmProviderPolyglotCreateHandler(dto *GsmProviderEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func GsmProviderValidator(dto *GsmProviderEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func GsmProviderEntityPreSanitize(dto *GsmProviderEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func GsmProviderEntityBeforeCreateAppend(dto *GsmProviderEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    GsmProviderRecursiveAddUniqueId(dto, query)
  }
  func GsmProviderRecursiveAddUniqueId(dto *GsmProviderEntity, query QueryDSL) {
  }
func GsmProviderActionBatchCreateFn(dtos []*GsmProviderEntity, query QueryDSL) ([]*GsmProviderEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*GsmProviderEntity{}
		for _, item := range dtos {
			s, err := GsmProviderActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func GsmProviderActionCreateFn(dto *GsmProviderEntity, query QueryDSL) (*GsmProviderEntity, *IError) {
	// 1. Validate always
	if iError := GsmProviderValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	GsmProviderEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	GsmProviderEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	GsmProviderPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	GsmProviderRelationContentCreate(dto, query)
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
	GsmProviderAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(GSMPROVIDER_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&GsmProviderEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func GsmProviderActionGetOne(query QueryDSL) (*GsmProviderEntity, *IError) {
    refl := reflect.ValueOf(&GsmProviderEntity{})
    item, err := GetOneEntity[GsmProviderEntity](query, refl)
    entityGsmProviderFormatter(item, query)
    return item, err
  }
  func GsmProviderActionQuery(query QueryDSL) ([]*GsmProviderEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&GsmProviderEntity{})
    items, meta, err := QueryEntitiesPointer[GsmProviderEntity](query, refl)
    for _, item := range items {
      entityGsmProviderFormatter(item, query)
    }
    return items, meta, err
  }
  func GsmProviderUpdateExec(dbref *gorm.DB, query QueryDSL, fields *GsmProviderEntity) (*GsmProviderEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = GSMPROVIDER_EVENT_UPDATED
    GsmProviderEntityPreSanitize(fields, query)
    var item GsmProviderEntity
    q := dbref.
      Where(&GsmProviderEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    GsmProviderRelationContentUpdate(fields, query)
    GsmProviderPolyglotCreateHandler(fields, query)
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&GsmProviderEntity{UniqueId: uniqueId}).
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
  func GsmProviderActionUpdateFn(query QueryDSL, fields *GsmProviderEntity) (*GsmProviderEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := GsmProviderValidator(fields, true); iError != nil {
      return nil, iError
    }
    GsmProviderRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        _, err := GsmProviderUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return nil, CastToIError(vf)
    } else {
      dbref = query.Tx
      return GsmProviderUpdateExec(dbref, query, fields)
    }
  }
var GsmProviderWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire gsmproviders ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilder(c)
		count, _ := GsmProviderActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func GsmProviderActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&GsmProviderEntity{})
	query.ActionRequires = []string{PERM_ROOT_GSMPROVIDER_DELETE}
	return RemoveEntity[GsmProviderEntity](query, refl)
}
func GsmProviderActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[GsmProviderEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'GsmProviderEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func GsmProviderActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[GsmProviderEntity]) (
    *BulkRecordRequest[GsmProviderEntity], *IError,
  ) {
    result := []*GsmProviderEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := GsmProviderActionUpdate(query, record)
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
func (x *GsmProviderEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var GsmProviderEntityMeta = TableMetaData{
	EntityName:    "GsmProvider",
	ExportKey:    "gsm-providers",
	TableNameInDb: "fb_gsmprovider_entities",
	EntityObject:  &GsmProviderEntity{},
	ExportStream: GsmProviderActionExportT,
	ImportQuery: GsmProviderActionImport,
}
func GsmProviderActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[GsmProviderEntity](query, GsmProviderActionQuery, GsmProviderPreloadRelations)
}
func GsmProviderActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[GsmProviderEntity](query, GsmProviderActionQuery, GsmProviderPreloadRelations)
}
func GsmProviderActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content GsmProviderEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := GsmProviderActionCreate(&content, query)
	return err
}
var GsmProviderCommonCliFlags = []cli.Flag{
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
      Name:     "api-key",
      Required: false,
      Usage:    "apiKey",
    },
    &cli.StringFlag{
      Name:     "main-sender-number",
      Required: true,
      Usage:    "mainSenderNumber",
    },
    &cli.StringFlag{
      Name:     "type",
      Required: true,
      Usage:    "One of: 'url', 'terminal', 'mediana'",
    },
    &cli.StringFlag{
      Name:     "invoke-url",
      Required: false,
      Usage:    "invokeUrl",
    },
    &cli.StringFlag{
      Name:     "invoke-body",
      Required: false,
      Usage:    "invokeBody",
    },
}
var GsmProviderCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "apiKey",
		StructField:     "ApiKey",
		Required: false,
		Usage:    "apiKey",
		Type: "string",
	},
	{
		Name:     "mainSenderNumber",
		StructField:     "MainSenderNumber",
		Required: true,
		Usage:    "mainSenderNumber",
		Type: "string",
	},
	{
		Name:     "type",
		StructField:     "Type",
		Required: true,
		Usage:    "One of: 'url', 'terminal', 'mediana'",
		Type: "string",
	},
	{
		Name:     "invokeUrl",
		StructField:     "InvokeUrl",
		Required: false,
		Usage:    "invokeUrl",
		Type: "string",
	},
	{
		Name:     "invokeBody",
		StructField:     "InvokeBody",
		Required: false,
		Usage:    "invokeBody",
		Type: "string",
	},
}
var GsmProviderCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "api-key",
      Required: false,
      Usage:    "apiKey",
    },
    &cli.StringFlag{
      Name:     "main-sender-number",
      Required: true,
      Usage:    "mainSenderNumber",
    },
    &cli.StringFlag{
      Name:     "type",
      Required: true,
      Usage:    "One of: 'url', 'terminal', 'mediana'",
    },
    &cli.StringFlag{
      Name:     "invoke-url",
      Required: false,
      Usage:    "invokeUrl",
    },
    &cli.StringFlag{
      Name:     "invoke-body",
      Required: false,
      Usage:    "invokeBody",
    },
}
  var GsmProviderCreateCmd cli.Command = cli.Command{
    Name:    "create",
    Aliases: []string{"c"},
    Flags: GsmProviderCommonCliFlags,
    Usage: "Create a new template",
    Action: func(c *cli.Context) {
      query := CommonCliQueryDSLBuilder(c)
      entity := CastGsmProviderFromCli(c)
      if entity, err := GsmProviderActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var GsmProviderCreateInteractiveCmd cli.Command = cli.Command{
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
      entity := &GsmProviderEntity{}
      for _, item := range GsmProviderCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := GsmProviderActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var GsmProviderUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: GsmProviderCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilder(c)
      entity := CastGsmProviderFromCli(c)
      if entity, err := GsmProviderActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func CastGsmProviderFromCli (c *cli.Context) *GsmProviderEntity {
	template := &GsmProviderEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("api-key") {
        value := c.String("api-key")
        template.ApiKey = &value
      }
      if c.IsSet("main-sender-number") {
        value := c.String("main-sender-number")
        template.MainSenderNumber = &value
      }
      if c.IsSet("type") {
        value := c.String("type")
        template.Type = &value
      }
      if c.IsSet("invoke-url") {
        value := c.String("invoke-url")
        template.InvokeUrl = &value
      }
      if c.IsSet("invoke-body") {
        value := c.String("invoke-body")
        template.InvokeBody = &value
      }
	return template
}
  func GsmProviderSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      GsmProviderActionCreate,
      reflect.ValueOf(&GsmProviderEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func GsmProviderImportMocks() {
    SeederFromFSImport(
      QueryDSL{},
      GsmProviderActionCreate,
      reflect.ValueOf(&GsmProviderEntity{}).Elem(),
      &mocks.ViewsFs,
      []string{},
      false,
    )
  }
  func GsmProviderWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := GsmProviderActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "GsmProvider", result)
    }
  }
var GsmProviderImportExportCommands = []cli.Command{
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
			GsmProviderActionSeeder(query, c.Int("count"))
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
				Value: "gsm-provider-seeder.yml",
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
			GsmProviderActionSeederInit(f, c.String("file"), c.String("format"))
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
				Value: "gsm-provider-seeder-gsm-provider.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of gsm-providers, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]GsmProviderEntity{}
			ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
		cli.Command{
			Name:  "mocks",
			Usage: "Prints the list of mocks",
			Action: func(c *cli.Context) error {
				if entity, err := GetSeederFilenames(&mocks.ViewsFs, ""); err != nil {
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
				CommonCliImportEmbedCmd(c,
					GsmProviderActionCreate,
					reflect.ValueOf(&GsmProviderEntity{}).Elem(),
					&mocks.ViewsFs,
				)
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
				GsmProviderActionCreate,
				reflect.ValueOf(&GsmProviderEntity{}).Elem(),
				c.String("file"),
			)
			return nil
		},
	},
}
    var GsmProviderCliCommands []cli.Command = []cli.Command{
      GetCommonQuery(GsmProviderActionQuery),
      GetCommonTableQuery(reflect.ValueOf(&GsmProviderEntity{}).Elem(), GsmProviderActionQuery),
          GsmProviderCreateCmd,
          GsmProviderUpdateCmd,
          GsmProviderCreateInteractiveCmd,
          GsmProviderWipeCmd,
          GetCommonRemoveQuery(reflect.ValueOf(&GsmProviderEntity{}).Elem(), GsmProviderActionRemove),
  }
  func GsmProviderCliFn() cli.Command {
    GsmProviderCliCommands = append(GsmProviderCliCommands, GsmProviderImportExportCommands...)
    return cli.Command{
      Name:        "gsmProvider",
      Description: "GsmProviders module actions (sample module to handle complex entities)",
      Usage:       "",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: GsmProviderCliCommands,
    }
  }
  /**
  *	Override this function on GsmProviderEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendGsmProviderRouter = func(r *[]Module2Action) {}
  func GetGsmProviderModule2Actions() []Module2Action {
    routes := []Module2Action{
       {
        Method: "GET",
        Url:    "/gsm-providers",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_GSMPROVIDER_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpQueryEntity(c, GsmProviderActionQuery)
          },
        },
        Format: "QUERY",
        Action: GsmProviderActionQuery,
        ResponseEntity: &[]GsmProviderEntity{},
      },
      {
        Method: "GET",
        Url:    "/gsm-providers/export",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_GSMPROVIDER_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpStreamFileChannel(c, GsmProviderActionExport)
          },
        },
        Format: "QUERY",
        Action: GsmProviderActionExport,
        ResponseEntity: &[]GsmProviderEntity{},
      },
      {
        Method: "GET",
        Url:    "/gsm-provider/:uniqueId",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_GSMPROVIDER_QUERY},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpGetEntity(c, GsmProviderActionGetOne)
          },
        },
        Format: "GET_ONE",
        Action: GsmProviderActionGetOne,
        ResponseEntity: &GsmProviderEntity{},
      },
      {
        Method: "POST",
        Url:    "/gsm-provider",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_GSMPROVIDER_CREATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpPostEntity(c, GsmProviderActionCreate)
          },
        },
        Action: GsmProviderActionCreate,
        Format: "POST_ONE",
        RequestEntity: &GsmProviderEntity{},
        ResponseEntity: &GsmProviderEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/gsm-provider",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_GSMPROVIDER_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpUpdateEntity(c, GsmProviderActionUpdate)
          },
        },
        Action: GsmProviderActionUpdate,
        RequestEntity: &GsmProviderEntity{},
        Format: "PATCH_ONE",
        ResponseEntity: &GsmProviderEntity{},
      },
      {
        Method: "PATCH",
        Url:    "/gsm-providers",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_GSMPROVIDER_UPDATE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpUpdateEntities(c, GsmProviderActionBulkUpdate)
          },
        },
        Action: GsmProviderActionBulkUpdate,
        Format: "PATCH_BULK",
        RequestEntity:  &BulkRecordRequest[GsmProviderEntity]{},
        ResponseEntity: &BulkRecordRequest[GsmProviderEntity]{},
      },
      {
        Method: "DELETE",
        Url:    "/gsm-provider",
        Format: "DELETE_DSL",
        SecurityModel: SecurityModel{
          ActionRequires: []string{PERM_ROOT_GSMPROVIDER_DELETE},
        },
        Handlers: []gin.HandlerFunc{
          func (c *gin.Context) {
            HttpRemoveEntity(c, GsmProviderActionRemove)
          },
        },
        Action: GsmProviderActionRemove,
        RequestEntity: &DeleteRequest{},
        ResponseEntity: &DeleteResponse{},
        TargetEntity: &GsmProviderEntity{},
      },
    }
    // Append user defined functions
    AppendGsmProviderRouter(&routes)
    return routes
  }
  func CreateGsmProviderRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetGsmProviderModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, GsmProviderEntityJsonSchema, "gsm-provider-http", "workspaces")
    WriteEntitySchema("GsmProviderEntity", GsmProviderEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_GSMPROVIDER_DELETE = "root/gsmprovider/delete"
var PERM_ROOT_GSMPROVIDER_CREATE = "root/gsmprovider/create"
var PERM_ROOT_GSMPROVIDER_UPDATE = "root/gsmprovider/update"
var PERM_ROOT_GSMPROVIDER_QUERY = "root/gsmprovider/query"
var PERM_ROOT_GSMPROVIDER = "root/gsmprovider"
var ALL_GSMPROVIDER_PERMISSIONS = []string{
	PERM_ROOT_GSMPROVIDER_DELETE,
	PERM_ROOT_GSMPROVIDER_CREATE,
	PERM_ROOT_GSMPROVIDER_UPDATE,
	PERM_ROOT_GSMPROVIDER_QUERY,
	PERM_ROOT_GSMPROVIDER,
}
var GsmProviderType = newGsmProviderType()
func newGsmProviderType() *xGsmProviderType {
	return &xGsmProviderType{
      Url: "url",
      Terminal: "terminal",
      Mediana: "mediana",
	}
}
type xGsmProviderType struct {
    Url string
    Terminal string
    Mediana string
}