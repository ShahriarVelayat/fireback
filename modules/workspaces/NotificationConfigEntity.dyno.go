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
var notificationConfigSeedersFs *embed.FS = nil
func ResetNotificationConfigSeeders(fs *embed.FS) {
	notificationConfigSeedersFs = fs
}
type NotificationConfigEntity struct {
    Visibility       *string                         `json:"visibility,omitempty" yaml:"visibility"`
    WorkspaceId      *string                         `json:"workspaceId,omitempty" yaml:"workspaceId" gorm:"unique;not null;" `
    LinkerId         *string                         `json:"linkerId,omitempty" yaml:"linkerId"`
    ParentId         *string                         `json:"parentId,omitempty" yaml:"parentId"`
    IsDeletable         *bool                         `json:"isDeletable,omitempty" yaml:"isDeletable" gorm:"default:true"`
    IsUpdatable         *bool                         `json:"isUpdatable,omitempty" yaml:"isUpdatable" gorm:"default:true"`
    UniqueId         string                          `json:"uniqueId,omitempty" gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId"`
    UserId           *string                         `json:"userId,omitempty" yaml:"userId"`
    Rank             int64                           `json:"rank,omitempty" gorm:"type:int;name:rank"`
    Updated          int64                           `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
    Created          int64                           `json:"created,omitempty" gorm:"autoUpdateTime:nano"`
    CreatedFormatted string                          `json:"createdFormatted,omitempty" sql:"-" gorm:"-"`
    UpdatedFormatted string                          `json:"updatedFormatted,omitempty" sql:"-" gorm:"-"`
    CascadeToSubWorkspaces   *bool `json:"cascadeToSubWorkspaces" yaml:"cascadeToSubWorkspaces"       `
    // Datenano also has a text representation
    ForcedCascadeEmailProvider   *bool `json:"forcedCascadeEmailProvider" yaml:"forcedCascadeEmailProvider"       `
    // Datenano also has a text representation
    GeneralEmailProvider   *  EmailProviderEntity `json:"generalEmailProvider" yaml:"generalEmailProvider"    gorm:"foreignKey:GeneralEmailProviderId;references:UniqueId"     `
    // Datenano also has a text representation
        GeneralEmailProviderId *string `json:"generalEmailProviderId" yaml:"generalEmailProviderId"`
    GeneralGsmProvider   *  GsmProviderEntity `json:"generalGsmProvider" yaml:"generalGsmProvider"    gorm:"foreignKey:GeneralGsmProviderId;references:UniqueId"     `
    // Datenano also has a text representation
        GeneralGsmProviderId *string `json:"generalGsmProviderId" yaml:"generalGsmProviderId"`
    InviteToWorkspaceContent   *string `json:"inviteToWorkspaceContent" yaml:"inviteToWorkspaceContent"    gorm:"text"     `
    // Datenano also has a text representation
    InviteToWorkspaceContentExcerpt   *string `json:"inviteToWorkspaceContentExcerpt" yaml:"inviteToWorkspaceContentExcerpt"    gorm:"text"     `
    // Datenano also has a text representation
    InviteToWorkspaceContentDefault   *string `json:"inviteToWorkspaceContentDefault" yaml:"inviteToWorkspaceContentDefault"    gorm:"text"     sql:"false"  `
    // Datenano also has a text representation
    InviteToWorkspaceContentDefaultExcerpt   *string `json:"inviteToWorkspaceContentDefaultExcerpt" yaml:"inviteToWorkspaceContentDefaultExcerpt"    gorm:"text"     sql:"false"  `
    // Datenano also has a text representation
    InviteToWorkspaceTitle   *string `json:"inviteToWorkspaceTitle" yaml:"inviteToWorkspaceTitle"       `
    // Datenano also has a text representation
    InviteToWorkspaceTitleDefault   *string `json:"inviteToWorkspaceTitleDefault" yaml:"inviteToWorkspaceTitleDefault"       sql:"false"  `
    // Datenano also has a text representation
    InviteToWorkspaceSender   *  EmailSenderEntity `json:"inviteToWorkspaceSender" yaml:"inviteToWorkspaceSender"    gorm:"foreignKey:InviteToWorkspaceSenderId;references:UniqueId"     `
    // Datenano also has a text representation
        InviteToWorkspaceSenderId *string `json:"inviteToWorkspaceSenderId" yaml:"inviteToWorkspaceSenderId"`
    AccountCenterEmailSender   *  EmailSenderEntity `json:"accountCenterEmailSender" yaml:"accountCenterEmailSender"    gorm:"foreignKey:AccountCenterEmailSenderId;references:UniqueId"     `
    // Datenano also has a text representation
        AccountCenterEmailSenderId *string `json:"accountCenterEmailSenderId" yaml:"accountCenterEmailSenderId"`
    ForgetPasswordContent   *string `json:"forgetPasswordContent" yaml:"forgetPasswordContent"    gorm:"text"     `
    // Datenano also has a text representation
    ForgetPasswordContentExcerpt   *string `json:"forgetPasswordContentExcerpt" yaml:"forgetPasswordContentExcerpt"    gorm:"text"     `
    // Datenano also has a text representation
    ForgetPasswordContentDefault   *string `json:"forgetPasswordContentDefault" yaml:"forgetPasswordContentDefault"    gorm:"text"     sql:"false"  `
    // Datenano also has a text representation
    ForgetPasswordContentDefaultExcerpt   *string `json:"forgetPasswordContentDefaultExcerpt" yaml:"forgetPasswordContentDefaultExcerpt"    gorm:"text"     sql:"false"  `
    // Datenano also has a text representation
    ForgetPasswordTitle   *string `json:"forgetPasswordTitle" yaml:"forgetPasswordTitle"    gorm:"text"     `
    // Datenano also has a text representation
    ForgetPasswordTitleDefault   *string `json:"forgetPasswordTitleDefault" yaml:"forgetPasswordTitleDefault"    gorm:"text"     sql:"false"  `
    // Datenano also has a text representation
    ForgetPasswordSender   *  EmailSenderEntity `json:"forgetPasswordSender" yaml:"forgetPasswordSender"    gorm:"foreignKey:ForgetPasswordSenderId;references:UniqueId"     `
    // Datenano also has a text representation
        ForgetPasswordSenderId *string `json:"forgetPasswordSenderId" yaml:"forgetPasswordSenderId"`
    AcceptLanguage   *string `json:"acceptLanguage" yaml:"acceptLanguage"    gorm:"text"     `
    // Datenano also has a text representation
    AcceptLanguageExcerpt *string `json:"acceptLanguageExcerpt" yaml:"acceptLanguageExcerpt"`
    ConfirmEmailSender   *  EmailSenderEntity `json:"confirmEmailSender" yaml:"confirmEmailSender"    gorm:"foreignKey:ConfirmEmailSenderId;references:UniqueId"     `
    // Datenano also has a text representation
        ConfirmEmailSenderId *string `json:"confirmEmailSenderId" yaml:"confirmEmailSenderId"`
    ConfirmEmailContent   *string `json:"confirmEmailContent" yaml:"confirmEmailContent"    gorm:"text"     `
    // Datenano also has a text representation
    ConfirmEmailContentExcerpt   *string `json:"confirmEmailContentExcerpt" yaml:"confirmEmailContentExcerpt"    gorm:"text"     `
    // Datenano also has a text representation
    ConfirmEmailContentDefault   *string `json:"confirmEmailContentDefault" yaml:"confirmEmailContentDefault"    gorm:"text"     sql:"false"  `
    // Datenano also has a text representation
    ConfirmEmailContentDefaultExcerpt   *string `json:"confirmEmailContentDefaultExcerpt" yaml:"confirmEmailContentDefaultExcerpt"    gorm:"text"     sql:"false"  `
    // Datenano also has a text representation
    ConfirmEmailTitle   *string `json:"confirmEmailTitle" yaml:"confirmEmailTitle"       `
    // Datenano also has a text representation
    ConfirmEmailTitleDefault   *string `json:"confirmEmailTitleDefault" yaml:"confirmEmailTitleDefault"       sql:"false"  `
    // Datenano also has a text representation
    Children []*NotificationConfigEntity `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`
    LinkedTo *NotificationConfigEntity `yaml:"-" gorm:"-" json:"-" sql:"-"`
}
var NotificationConfigPreloadRelations []string = []string{}
var NOTIFICATION_CONFIG_EVENT_CREATED = "notificationConfig.created"
var NOTIFICATION_CONFIG_EVENT_UPDATED = "notificationConfig.updated"
var NOTIFICATION_CONFIG_EVENT_DELETED = "notificationConfig.deleted"
var NOTIFICATION_CONFIG_EVENTS = []string{
	NOTIFICATION_CONFIG_EVENT_CREATED,
	NOTIFICATION_CONFIG_EVENT_UPDATED,
	NOTIFICATION_CONFIG_EVENT_DELETED,
}
type NotificationConfigFieldMap struct {
		CascadeToSubWorkspaces TranslatedString `yaml:"cascadeToSubWorkspaces"`
		ForcedCascadeEmailProvider TranslatedString `yaml:"forcedCascadeEmailProvider"`
		GeneralEmailProvider TranslatedString `yaml:"generalEmailProvider"`
		GeneralGsmProvider TranslatedString `yaml:"generalGsmProvider"`
		InviteToWorkspaceContent TranslatedString `yaml:"inviteToWorkspaceContent"`
		InviteToWorkspaceContentExcerpt TranslatedString `yaml:"inviteToWorkspaceContentExcerpt"`
		InviteToWorkspaceContentDefault TranslatedString `yaml:"inviteToWorkspaceContentDefault"`
		InviteToWorkspaceContentDefaultExcerpt TranslatedString `yaml:"inviteToWorkspaceContentDefaultExcerpt"`
		InviteToWorkspaceTitle TranslatedString `yaml:"inviteToWorkspaceTitle"`
		InviteToWorkspaceTitleDefault TranslatedString `yaml:"inviteToWorkspaceTitleDefault"`
		InviteToWorkspaceSender TranslatedString `yaml:"inviteToWorkspaceSender"`
		AccountCenterEmailSender TranslatedString `yaml:"accountCenterEmailSender"`
		ForgetPasswordContent TranslatedString `yaml:"forgetPasswordContent"`
		ForgetPasswordContentExcerpt TranslatedString `yaml:"forgetPasswordContentExcerpt"`
		ForgetPasswordContentDefault TranslatedString `yaml:"forgetPasswordContentDefault"`
		ForgetPasswordContentDefaultExcerpt TranslatedString `yaml:"forgetPasswordContentDefaultExcerpt"`
		ForgetPasswordTitle TranslatedString `yaml:"forgetPasswordTitle"`
		ForgetPasswordTitleDefault TranslatedString `yaml:"forgetPasswordTitleDefault"`
		ForgetPasswordSender TranslatedString `yaml:"forgetPasswordSender"`
		AcceptLanguage TranslatedString `yaml:"acceptLanguage"`
		ConfirmEmailSender TranslatedString `yaml:"confirmEmailSender"`
		ConfirmEmailContent TranslatedString `yaml:"confirmEmailContent"`
		ConfirmEmailContentExcerpt TranslatedString `yaml:"confirmEmailContentExcerpt"`
		ConfirmEmailContentDefault TranslatedString `yaml:"confirmEmailContentDefault"`
		ConfirmEmailContentDefaultExcerpt TranslatedString `yaml:"confirmEmailContentDefaultExcerpt"`
		ConfirmEmailTitle TranslatedString `yaml:"confirmEmailTitle"`
		ConfirmEmailTitleDefault TranslatedString `yaml:"confirmEmailTitleDefault"`
}
var NotificationConfigEntityMetaConfig map[string]int64 = map[string]int64{
            "AcceptLanguageExcerptSize": 100,
}
var NotificationConfigEntityJsonSchema = ExtractEntityFields(reflect.ValueOf(&NotificationConfigEntity{}))
func entityNotificationConfigFormatter(dto *NotificationConfigEntity, query QueryDSL) {
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
func NotificationConfigMockEntity() *NotificationConfigEntity {
	stringHolder := "~"
	int64Holder := int64(10)
	float64Holder := float64(10)
	_ = stringHolder
	_ = int64Holder
	_ = float64Holder
	entity := &NotificationConfigEntity{
      InviteToWorkspaceContent : &stringHolder,
      InviteToWorkspaceContentExcerpt : &stringHolder,
      InviteToWorkspaceContentDefault : &stringHolder,
      InviteToWorkspaceContentDefaultExcerpt : &stringHolder,
      InviteToWorkspaceTitle : &stringHolder,
      InviteToWorkspaceTitleDefault : &stringHolder,
      ForgetPasswordContent : &stringHolder,
      ForgetPasswordContentExcerpt : &stringHolder,
      ForgetPasswordContentDefault : &stringHolder,
      ForgetPasswordContentDefaultExcerpt : &stringHolder,
      ForgetPasswordTitle : &stringHolder,
      ForgetPasswordTitleDefault : &stringHolder,
      ConfirmEmailContent : &stringHolder,
      ConfirmEmailContentExcerpt : &stringHolder,
      ConfirmEmailContentDefault : &stringHolder,
      ConfirmEmailContentDefaultExcerpt : &stringHolder,
      ConfirmEmailTitle : &stringHolder,
      ConfirmEmailTitleDefault : &stringHolder,
	}
	return entity
}
func NotificationConfigActionSeeder(query QueryDSL, count int) {
	successInsert := 0
	failureInsert := 0
	bar := progressbar.Default(int64(count))
	for i := 1; i <= count; i++ {
		entity := NotificationConfigMockEntity()
		_, err := NotificationConfigActionCreate(entity, query)
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
  func NotificationConfigActionSeederInit(query QueryDSL, file string, format string) {
    body := []byte{}
    var err error
    data := []*NotificationConfigEntity{}
    tildaRef := "~"
    _ = tildaRef
    entity := &NotificationConfigEntity{
          InviteToWorkspaceContent: &tildaRef,
          InviteToWorkspaceContentExcerpt: &tildaRef,
          InviteToWorkspaceContentDefault: &tildaRef,
          InviteToWorkspaceContentDefaultExcerpt: &tildaRef,
          InviteToWorkspaceTitle: &tildaRef,
          InviteToWorkspaceTitleDefault: &tildaRef,
          ForgetPasswordContent: &tildaRef,
          ForgetPasswordContentExcerpt: &tildaRef,
          ForgetPasswordContentDefault: &tildaRef,
          ForgetPasswordContentDefaultExcerpt: &tildaRef,
          ForgetPasswordTitle: &tildaRef,
          ForgetPasswordTitleDefault: &tildaRef,
          ConfirmEmailContent: &tildaRef,
          ConfirmEmailContentExcerpt: &tildaRef,
          ConfirmEmailContentDefault: &tildaRef,
          ConfirmEmailContentDefaultExcerpt: &tildaRef,
          ConfirmEmailTitle: &tildaRef,
          ConfirmEmailTitleDefault: &tildaRef,
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
  func NotificationConfigAssociationCreate(dto *NotificationConfigEntity, query QueryDSL) error {
    return nil
  }
/**
* These kind of content are coming from another entity, which is indepndent module
* If we want to create them, we need to do it before. This is not association.
**/
func NotificationConfigRelationContentCreate(dto *NotificationConfigEntity, query QueryDSL) error {
return nil
}
func NotificationConfigRelationContentUpdate(dto *NotificationConfigEntity, query QueryDSL) error {
	return nil
}
func NotificationConfigPolyglotCreateHandler(dto *NotificationConfigEntity, query QueryDSL) {
	if dto == nil {
		return
	}
}
  /**
  * This will be validating your entity fully. Important note is that, you add validate:* tag
  * in your entity, it will automatically work here. For slices inside entity, make sure you add
  * extra line of AppendSliceErrors, otherwise they won't be detected
  */
  func NotificationConfigValidator(dto *NotificationConfigEntity, isPatch bool) *IError {
    err := CommonStructValidatorPointer(dto, isPatch)
    return err
  }
func NotificationConfigEntityPreSanitize(dto *NotificationConfigEntity, query QueryDSL) {
	var stripPolicy = bluemonday.StripTagsPolicy()
	var ugcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
	_ = stripPolicy
	_ = ugcPolicy
}
  func NotificationConfigEntityBeforeCreateAppend(dto *NotificationConfigEntity, query QueryDSL) {
    if (dto.UniqueId == "") {
      dto.UniqueId = UUID()
    }
    dto.WorkspaceId = &query.WorkspaceId
    dto.UserId = &query.UserId
    NotificationConfigRecursiveAddUniqueId(dto, query)
  }
  func NotificationConfigRecursiveAddUniqueId(dto *NotificationConfigEntity, query QueryDSL) {
  }
func NotificationConfigActionBatchCreateFn(dtos []*NotificationConfigEntity, query QueryDSL) ([]*NotificationConfigEntity, *IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*NotificationConfigEntity{}
		for _, item := range dtos {
			s, err := NotificationConfigActionCreateFn(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil;
}
func NotificationConfigDeleteEntireChildren(query QueryDSL, dto *NotificationConfigEntity) (*IError) {
  // intentionally removed this. It's hard to implement it, and probably wrong without
  // proper on delete cascade
  return nil
}
func NotificationConfigActionCreateFn(dto *NotificationConfigEntity, query QueryDSL) (*NotificationConfigEntity, *IError) {
	// 1. Validate always
	if iError := NotificationConfigValidator(dto, false); iError != nil {
		return nil, iError
	}
	// 1.5 Sanitize the content coming of the front-end
	NotificationConfigEntityPreSanitize(dto, query)
	// 2. Append the necessary information about user, workspace
	NotificationConfigEntityBeforeCreateAppend(dto, query)
	// 3. Append the necessary translations, even if english
	NotificationConfigPolyglotCreateHandler(dto, query)
	// 3.5. Create other entities if we want select from them
	NotificationConfigRelationContentCreate(dto, query)
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
	NotificationConfigAssociationCreate(dto, query)
	// 6. Fire the event into system
	event.MustFire(NOTIFICATION_CONFIG_EVENT_CREATED, event.M{
		"entity":   dto,
		"entityKey": GetTypeString(&NotificationConfigEntity{}),
		"target":   "workspace",
		"unqiueId": query.WorkspaceId,
	})
	return dto, nil
}
  func NotificationConfigActionGetOne(query QueryDSL) (*NotificationConfigEntity, *IError) {
    refl := reflect.ValueOf(&NotificationConfigEntity{})
    item, err := GetOneEntity[NotificationConfigEntity](query, refl)
    entityNotificationConfigFormatter(item, query)
    return item, err
  }
  func NotificationConfigActionQuery(query QueryDSL) ([]*NotificationConfigEntity, *QueryResultMeta, error) {
    refl := reflect.ValueOf(&NotificationConfigEntity{})
    items, meta, err := QueryEntitiesPointer[NotificationConfigEntity](query, refl)
    for _, item := range items {
      entityNotificationConfigFormatter(item, query)
    }
    return items, meta, err
  }
  func NotificationConfigUpdateExec(dbref *gorm.DB, query QueryDSL, fields *NotificationConfigEntity) (*NotificationConfigEntity, *IError) {
    uniqueId := fields.UniqueId
    query.TriggerEventName = NOTIFICATION_CONFIG_EVENT_UPDATED
    NotificationConfigEntityPreSanitize(fields, query)
    var item NotificationConfigEntity
    q := dbref.
      Where(&NotificationConfigEntity{UniqueId: uniqueId}).
      FirstOrCreate(&item)
    err := q.UpdateColumns(fields).Error
    if err != nil {
      return nil, GormErrorToIError(err)
    }
    query.Tx = dbref
    NotificationConfigRelationContentUpdate(fields, query)
    NotificationConfigPolyglotCreateHandler(fields, query)
    if ero := NotificationConfigDeleteEntireChildren(query, fields); ero != nil {
      return nil, ero
    }
    // @meta(update has many)
    err = dbref.
      Preload(clause.Associations).
      Where(&NotificationConfigEntity{UniqueId: uniqueId}).
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
  func NotificationConfigActionUpdateFn(query QueryDSL, fields *NotificationConfigEntity) (*NotificationConfigEntity, *IError) {
    if fields == nil {
      return nil, CreateIErrorString("ENTITY_IS_NEEDED", []string{}, 403)
    }
    // 1. Validate always
    if iError := NotificationConfigValidator(fields, true); iError != nil {
      return nil, iError
    }
    // Let's not add this. I am not sure of the consequences
    // NotificationConfigRecursiveAddUniqueId(fields, query)
    var dbref *gorm.DB = nil
    if query.Tx == nil {
      dbref = GetDbRef()
      var item *NotificationConfigEntity
      vf := dbref.Transaction(func(tx *gorm.DB) error {
        dbref = tx
        var err *IError
        item, err = NotificationConfigUpdateExec(dbref, query, fields)
        if err == nil {
          return nil
        } else {
          return err
        }
      })
      return item, CastToIError(vf)
    } else {
      dbref = query.Tx
      return NotificationConfigUpdateExec(dbref, query, fields)
    }
  }
var NotificationConfigWipeCmd cli.Command = cli.Command{
	Name:  "wipe",
	Usage: "Wipes entire notificationconfigs ",
	Action: func(c *cli.Context) error {
		query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
      ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_DELETE},
    })
		count, _ := NotificationConfigActionWipeClean(query)
		fmt.Println("Removed", count, "of entities")
		return nil
	},
}
func NotificationConfigActionRemove(query QueryDSL) (int64, *IError) {
	refl := reflect.ValueOf(&NotificationConfigEntity{})
	query.ActionRequires = []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_DELETE}
	return RemoveEntity[NotificationConfigEntity](query, refl)
}
func NotificationConfigActionWipeClean(query QueryDSL) (int64, error) {
	var err error;
	var count int64 = 0;
	{
		subCount, subErr := WipeCleanEntity[NotificationConfigEntity]()	
		if (subErr != nil) {
			fmt.Println("Error while wiping 'NotificationConfigEntity'", subErr)
			return count, subErr
		} else {
			count += subCount
		}
	}
	return count, err
}
  func NotificationConfigActionBulkUpdate(
    query QueryDSL, dto *BulkRecordRequest[NotificationConfigEntity]) (
    *BulkRecordRequest[NotificationConfigEntity], *IError,
  ) {
    result := []*NotificationConfigEntity{}
    err := GetDbRef().Transaction(func(tx *gorm.DB) error {
      query.Tx = tx
      for _, record := range dto.Records {
        item, err := NotificationConfigActionUpdate(query, record)
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
func (x *NotificationConfigEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}
var NotificationConfigEntityMeta = TableMetaData{
	EntityName:    "NotificationConfig",
	ExportKey:    "notification-configs",
	TableNameInDb: "fb_notification-config_entities",
	EntityObject:  &NotificationConfigEntity{},
	ExportStream: NotificationConfigActionExportT,
	ImportQuery: NotificationConfigActionImport,
}
func NotificationConfigActionExport(
	query QueryDSL,
) (chan []byte, *IError) {
	return YamlExporterChannel[NotificationConfigEntity](query, NotificationConfigActionQuery, NotificationConfigPreloadRelations)
}
func NotificationConfigActionExportT(
	query QueryDSL,
) (chan []interface{}, *IError) {
	return YamlExporterChannelT[NotificationConfigEntity](query, NotificationConfigActionQuery, NotificationConfigPreloadRelations)
}
func NotificationConfigActionImport(
	dto interface{}, query QueryDSL,
) *IError {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content NotificationConfigEntity
	cx, err2 := json.Marshal(dto)
	if err2 != nil {
		return CreateIErrorString("INVALID_CONTENT", []string{}, 501)
	}
	json.Unmarshal(cx, &content)
	_, err := NotificationConfigActionCreate(&content, query)
	return err
}
var NotificationConfigCommonCliFlags = []cli.Flag{
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
    &cli.BoolFlag{
      Name:     "cascade-to-sub-workspaces",
      Required: false,
      Usage:    "cascadeToSubWorkspaces",
    },
    &cli.BoolFlag{
      Name:     "forced-cascade-email-provider",
      Required: false,
      Usage:    "forcedCascadeEmailProvider",
    },
    &cli.StringFlag{
      Name:     "general-email-provider-id",
      Required: false,
      Usage:    "generalEmailProvider",
    },
    &cli.StringFlag{
      Name:     "general-gsm-provider-id",
      Required: false,
      Usage:    "generalGsmProvider",
    },
    &cli.StringFlag{
      Name:     "invite-to-workspace-content",
      Required: false,
      Usage:    "inviteToWorkspaceContent",
    },
    &cli.StringFlag{
      Name:     "invite-to-workspace-content-excerpt",
      Required: false,
      Usage:    "inviteToWorkspaceContentExcerpt",
    },
    &cli.StringFlag{
      Name:     "invite-to-workspace-content-default",
      Required: false,
      Usage:    "inviteToWorkspaceContentDefault",
    },
    &cli.StringFlag{
      Name:     "invite-to-workspace-content-default-excerpt",
      Required: false,
      Usage:    "inviteToWorkspaceContentDefaultExcerpt",
    },
    &cli.StringFlag{
      Name:     "invite-to-workspace-title",
      Required: false,
      Usage:    "inviteToWorkspaceTitle",
    },
    &cli.StringFlag{
      Name:     "invite-to-workspace-title-default",
      Required: false,
      Usage:    "inviteToWorkspaceTitleDefault",
    },
    &cli.StringFlag{
      Name:     "invite-to-workspace-sender-id",
      Required: false,
      Usage:    "inviteToWorkspaceSender",
    },
    &cli.StringFlag{
      Name:     "account-center-email-sender-id",
      Required: false,
      Usage:    "accountCenterEmailSender",
    },
    &cli.StringFlag{
      Name:     "forget-password-content",
      Required: false,
      Usage:    "forgetPasswordContent",
    },
    &cli.StringFlag{
      Name:     "forget-password-content-excerpt",
      Required: false,
      Usage:    "forgetPasswordContentExcerpt",
    },
    &cli.StringFlag{
      Name:     "forget-password-content-default",
      Required: false,
      Usage:    "forgetPasswordContentDefault",
    },
    &cli.StringFlag{
      Name:     "forget-password-content-default-excerpt",
      Required: false,
      Usage:    "forgetPasswordContentDefaultExcerpt",
    },
    &cli.StringFlag{
      Name:     "forget-password-title",
      Required: false,
      Usage:    "forgetPasswordTitle",
    },
    &cli.StringFlag{
      Name:     "forget-password-title-default",
      Required: false,
      Usage:    "forgetPasswordTitleDefault",
    },
    &cli.StringFlag{
      Name:     "forget-password-sender-id",
      Required: false,
      Usage:    "forgetPasswordSender",
    },
    &cli.StringFlag{
      Name:     "accept-language",
      Required: false,
      Usage:    "acceptLanguage",
    },
    &cli.StringFlag{
      Name:     "confirm-email-sender-id",
      Required: false,
      Usage:    "confirmEmailSender",
    },
    &cli.StringFlag{
      Name:     "confirm-email-content",
      Required: false,
      Usage:    "confirmEmailContent",
    },
    &cli.StringFlag{
      Name:     "confirm-email-content-excerpt",
      Required: false,
      Usage:    "confirmEmailContentExcerpt",
    },
    &cli.StringFlag{
      Name:     "confirm-email-content-default",
      Required: false,
      Usage:    "confirmEmailContentDefault",
    },
    &cli.StringFlag{
      Name:     "confirm-email-content-default-excerpt",
      Required: false,
      Usage:    "confirmEmailContentDefaultExcerpt",
    },
    &cli.StringFlag{
      Name:     "confirm-email-title",
      Required: false,
      Usage:    "confirmEmailTitle",
    },
    &cli.StringFlag{
      Name:     "confirm-email-title-default",
      Required: false,
      Usage:    "confirmEmailTitleDefault",
    },
}
var NotificationConfigCommonInteractiveCliFlags = []CliInteractiveFlag{
	{
		Name:     "cascadeToSubWorkspaces",
		StructField:     "CascadeToSubWorkspaces",
		Required: false,
		Usage:    "cascadeToSubWorkspaces",
		Type: "bool",
	},
	{
		Name:     "forcedCascadeEmailProvider",
		StructField:     "ForcedCascadeEmailProvider",
		Required: false,
		Usage:    "forcedCascadeEmailProvider",
		Type: "bool",
	},
	{
		Name:     "inviteToWorkspaceContent",
		StructField:     "InviteToWorkspaceContent",
		Required: false,
		Usage:    "inviteToWorkspaceContent",
		Type: "string",
	},
	{
		Name:     "inviteToWorkspaceContentExcerpt",
		StructField:     "InviteToWorkspaceContentExcerpt",
		Required: false,
		Usage:    "inviteToWorkspaceContentExcerpt",
		Type: "string",
	},
	{
		Name:     "inviteToWorkspaceContentDefault",
		StructField:     "InviteToWorkspaceContentDefault",
		Required: false,
		Usage:    "inviteToWorkspaceContentDefault",
		Type: "string",
	},
	{
		Name:     "inviteToWorkspaceContentDefaultExcerpt",
		StructField:     "InviteToWorkspaceContentDefaultExcerpt",
		Required: false,
		Usage:    "inviteToWorkspaceContentDefaultExcerpt",
		Type: "string",
	},
	{
		Name:     "inviteToWorkspaceTitle",
		StructField:     "InviteToWorkspaceTitle",
		Required: false,
		Usage:    "inviteToWorkspaceTitle",
		Type: "string",
	},
	{
		Name:     "inviteToWorkspaceTitleDefault",
		StructField:     "InviteToWorkspaceTitleDefault",
		Required: false,
		Usage:    "inviteToWorkspaceTitleDefault",
		Type: "string",
	},
	{
		Name:     "forgetPasswordContent",
		StructField:     "ForgetPasswordContent",
		Required: false,
		Usage:    "forgetPasswordContent",
		Type: "string",
	},
	{
		Name:     "forgetPasswordContentExcerpt",
		StructField:     "ForgetPasswordContentExcerpt",
		Required: false,
		Usage:    "forgetPasswordContentExcerpt",
		Type: "string",
	},
	{
		Name:     "forgetPasswordContentDefault",
		StructField:     "ForgetPasswordContentDefault",
		Required: false,
		Usage:    "forgetPasswordContentDefault",
		Type: "string",
	},
	{
		Name:     "forgetPasswordContentDefaultExcerpt",
		StructField:     "ForgetPasswordContentDefaultExcerpt",
		Required: false,
		Usage:    "forgetPasswordContentDefaultExcerpt",
		Type: "string",
	},
	{
		Name:     "forgetPasswordTitle",
		StructField:     "ForgetPasswordTitle",
		Required: false,
		Usage:    "forgetPasswordTitle",
		Type: "string",
	},
	{
		Name:     "forgetPasswordTitleDefault",
		StructField:     "ForgetPasswordTitleDefault",
		Required: false,
		Usage:    "forgetPasswordTitleDefault",
		Type: "string",
	},
	{
		Name:     "confirmEmailContent",
		StructField:     "ConfirmEmailContent",
		Required: false,
		Usage:    "confirmEmailContent",
		Type: "string",
	},
	{
		Name:     "confirmEmailContentExcerpt",
		StructField:     "ConfirmEmailContentExcerpt",
		Required: false,
		Usage:    "confirmEmailContentExcerpt",
		Type: "string",
	},
	{
		Name:     "confirmEmailContentDefault",
		StructField:     "ConfirmEmailContentDefault",
		Required: false,
		Usage:    "confirmEmailContentDefault",
		Type: "string",
	},
	{
		Name:     "confirmEmailContentDefaultExcerpt",
		StructField:     "ConfirmEmailContentDefaultExcerpt",
		Required: false,
		Usage:    "confirmEmailContentDefaultExcerpt",
		Type: "string",
	},
	{
		Name:     "confirmEmailTitle",
		StructField:     "ConfirmEmailTitle",
		Required: false,
		Usage:    "confirmEmailTitle",
		Type: "string",
	},
	{
		Name:     "confirmEmailTitleDefault",
		StructField:     "ConfirmEmailTitleDefault",
		Required: false,
		Usage:    "confirmEmailTitleDefault",
		Type: "string",
	},
}
var NotificationConfigCommonCliFlagsOptional = []cli.Flag{
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
    &cli.BoolFlag{
      Name:     "cascade-to-sub-workspaces",
      Required: false,
      Usage:    "cascadeToSubWorkspaces",
    },
    &cli.BoolFlag{
      Name:     "forced-cascade-email-provider",
      Required: false,
      Usage:    "forcedCascadeEmailProvider",
    },
    &cli.StringFlag{
      Name:     "general-email-provider-id",
      Required: false,
      Usage:    "generalEmailProvider",
    },
    &cli.StringFlag{
      Name:     "general-gsm-provider-id",
      Required: false,
      Usage:    "generalGsmProvider",
    },
    &cli.StringFlag{
      Name:     "invite-to-workspace-content",
      Required: false,
      Usage:    "inviteToWorkspaceContent",
    },
    &cli.StringFlag{
      Name:     "invite-to-workspace-content-excerpt",
      Required: false,
      Usage:    "inviteToWorkspaceContentExcerpt",
    },
    &cli.StringFlag{
      Name:     "invite-to-workspace-content-default",
      Required: false,
      Usage:    "inviteToWorkspaceContentDefault",
    },
    &cli.StringFlag{
      Name:     "invite-to-workspace-content-default-excerpt",
      Required: false,
      Usage:    "inviteToWorkspaceContentDefaultExcerpt",
    },
    &cli.StringFlag{
      Name:     "invite-to-workspace-title",
      Required: false,
      Usage:    "inviteToWorkspaceTitle",
    },
    &cli.StringFlag{
      Name:     "invite-to-workspace-title-default",
      Required: false,
      Usage:    "inviteToWorkspaceTitleDefault",
    },
    &cli.StringFlag{
      Name:     "invite-to-workspace-sender-id",
      Required: false,
      Usage:    "inviteToWorkspaceSender",
    },
    &cli.StringFlag{
      Name:     "account-center-email-sender-id",
      Required: false,
      Usage:    "accountCenterEmailSender",
    },
    &cli.StringFlag{
      Name:     "forget-password-content",
      Required: false,
      Usage:    "forgetPasswordContent",
    },
    &cli.StringFlag{
      Name:     "forget-password-content-excerpt",
      Required: false,
      Usage:    "forgetPasswordContentExcerpt",
    },
    &cli.StringFlag{
      Name:     "forget-password-content-default",
      Required: false,
      Usage:    "forgetPasswordContentDefault",
    },
    &cli.StringFlag{
      Name:     "forget-password-content-default-excerpt",
      Required: false,
      Usage:    "forgetPasswordContentDefaultExcerpt",
    },
    &cli.StringFlag{
      Name:     "forget-password-title",
      Required: false,
      Usage:    "forgetPasswordTitle",
    },
    &cli.StringFlag{
      Name:     "forget-password-title-default",
      Required: false,
      Usage:    "forgetPasswordTitleDefault",
    },
    &cli.StringFlag{
      Name:     "forget-password-sender-id",
      Required: false,
      Usage:    "forgetPasswordSender",
    },
    &cli.StringFlag{
      Name:     "accept-language",
      Required: false,
      Usage:    "acceptLanguage",
    },
    &cli.StringFlag{
      Name:     "confirm-email-sender-id",
      Required: false,
      Usage:    "confirmEmailSender",
    },
    &cli.StringFlag{
      Name:     "confirm-email-content",
      Required: false,
      Usage:    "confirmEmailContent",
    },
    &cli.StringFlag{
      Name:     "confirm-email-content-excerpt",
      Required: false,
      Usage:    "confirmEmailContentExcerpt",
    },
    &cli.StringFlag{
      Name:     "confirm-email-content-default",
      Required: false,
      Usage:    "confirmEmailContentDefault",
    },
    &cli.StringFlag{
      Name:     "confirm-email-content-default-excerpt",
      Required: false,
      Usage:    "confirmEmailContentDefaultExcerpt",
    },
    &cli.StringFlag{
      Name:     "confirm-email-title",
      Required: false,
      Usage:    "confirmEmailTitle",
    },
    &cli.StringFlag{
      Name:     "confirm-email-title-default",
      Required: false,
      Usage:    "confirmEmailTitleDefault",
    },
}
  var NotificationConfigCreateCmd cli.Command = NOTIFICATION_CONFIG_ACTION_POST_ONE.ToCli()
  var NotificationConfigCreateInteractiveCmd cli.Command = cli.Command{
    Name:  "ic",
    Usage: "Creates a new template, using requied fields in an interactive name",
    Flags: []cli.Flag{
      &cli.BoolFlag{
        Name:  "all",
        Usage: "Interactively asks for all inputs, not only required ones",
      },
    },
    Action: func(c *cli.Context) {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_CREATE},
      })
      entity := &NotificationConfigEntity{}
      for _, item := range NotificationConfigCommonInteractiveCliFlags {
        if !item.Required && c.Bool("all") == false {
          continue
        }
        result := AskForInput(item.Name, "")
        SetFieldString(entity, item.StructField, result)
      }
      if entity, err := NotificationConfigActionCreate(entity, query); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
    },
  }
  var NotificationConfigUpdateCmd cli.Command = cli.Command{
    Name:    "update",
    Aliases: []string{"u"},
    Flags: NotificationConfigCommonCliFlagsOptional,
    Usage:   "Updates a template by passing the parameters",
    Action: func(c *cli.Context) error {
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_UPDATE},
      })
      entity := CastNotificationConfigFromCli(c)
      if entity, err := NotificationConfigActionUpdate(query, entity); err != nil {
        fmt.Println(err.Error())
      } else {
        f, _ := json.MarshalIndent(entity, "", "  ")
        fmt.Println(string(f))
      }
      return nil
    },
  }
func (x* NotificationConfigEntity) FromCli(c *cli.Context) *NotificationConfigEntity {
	return CastNotificationConfigFromCli(c)
}
func CastNotificationConfigFromCli (c *cli.Context) *NotificationConfigEntity {
	template := &NotificationConfigEntity{}
	if c.IsSet("uid") {
		template.UniqueId = c.String("uid")
	}
	if c.IsSet("pid") {
		x := c.String("pid")
		template.ParentId = &x
	}
      if c.IsSet("general-email-provider-id") {
        value := c.String("general-email-provider-id")
        template.GeneralEmailProviderId = &value
      }
      if c.IsSet("general-gsm-provider-id") {
        value := c.String("general-gsm-provider-id")
        template.GeneralGsmProviderId = &value
      }
      if c.IsSet("invite-to-workspace-content") {
        value := c.String("invite-to-workspace-content")
        template.InviteToWorkspaceContent = &value
      }
      if c.IsSet("invite-to-workspace-content-excerpt") {
        value := c.String("invite-to-workspace-content-excerpt")
        template.InviteToWorkspaceContentExcerpt = &value
      }
      if c.IsSet("invite-to-workspace-content-default") {
        value := c.String("invite-to-workspace-content-default")
        template.InviteToWorkspaceContentDefault = &value
      }
      if c.IsSet("invite-to-workspace-content-default-excerpt") {
        value := c.String("invite-to-workspace-content-default-excerpt")
        template.InviteToWorkspaceContentDefaultExcerpt = &value
      }
      if c.IsSet("invite-to-workspace-title") {
        value := c.String("invite-to-workspace-title")
        template.InviteToWorkspaceTitle = &value
      }
      if c.IsSet("invite-to-workspace-title-default") {
        value := c.String("invite-to-workspace-title-default")
        template.InviteToWorkspaceTitleDefault = &value
      }
      if c.IsSet("invite-to-workspace-sender-id") {
        value := c.String("invite-to-workspace-sender-id")
        template.InviteToWorkspaceSenderId = &value
      }
      if c.IsSet("account-center-email-sender-id") {
        value := c.String("account-center-email-sender-id")
        template.AccountCenterEmailSenderId = &value
      }
      if c.IsSet("forget-password-content") {
        value := c.String("forget-password-content")
        template.ForgetPasswordContent = &value
      }
      if c.IsSet("forget-password-content-excerpt") {
        value := c.String("forget-password-content-excerpt")
        template.ForgetPasswordContentExcerpt = &value
      }
      if c.IsSet("forget-password-content-default") {
        value := c.String("forget-password-content-default")
        template.ForgetPasswordContentDefault = &value
      }
      if c.IsSet("forget-password-content-default-excerpt") {
        value := c.String("forget-password-content-default-excerpt")
        template.ForgetPasswordContentDefaultExcerpt = &value
      }
      if c.IsSet("forget-password-title") {
        value := c.String("forget-password-title")
        template.ForgetPasswordTitle = &value
      }
      if c.IsSet("forget-password-title-default") {
        value := c.String("forget-password-title-default")
        template.ForgetPasswordTitleDefault = &value
      }
      if c.IsSet("forget-password-sender-id") {
        value := c.String("forget-password-sender-id")
        template.ForgetPasswordSenderId = &value
      }
      if c.IsSet("accept-language") {
        value := c.String("accept-language")
        template.AcceptLanguage = &value
      }
      if c.IsSet("confirm-email-sender-id") {
        value := c.String("confirm-email-sender-id")
        template.ConfirmEmailSenderId = &value
      }
      if c.IsSet("confirm-email-content") {
        value := c.String("confirm-email-content")
        template.ConfirmEmailContent = &value
      }
      if c.IsSet("confirm-email-content-excerpt") {
        value := c.String("confirm-email-content-excerpt")
        template.ConfirmEmailContentExcerpt = &value
      }
      if c.IsSet("confirm-email-content-default") {
        value := c.String("confirm-email-content-default")
        template.ConfirmEmailContentDefault = &value
      }
      if c.IsSet("confirm-email-content-default-excerpt") {
        value := c.String("confirm-email-content-default-excerpt")
        template.ConfirmEmailContentDefaultExcerpt = &value
      }
      if c.IsSet("confirm-email-title") {
        value := c.String("confirm-email-title")
        template.ConfirmEmailTitle = &value
      }
      if c.IsSet("confirm-email-title-default") {
        value := c.String("confirm-email-title-default")
        template.ConfirmEmailTitleDefault = &value
      }
	return template
}
  func NotificationConfigSyncSeederFromFs(fsRef *embed.FS, fileNames []string) {
    SeederFromFSImport(
      QueryDSL{},
      NotificationConfigActionCreate,
      reflect.ValueOf(&NotificationConfigEntity{}).Elem(),
      fsRef,
      fileNames,
      true,
    )
  }
  func NotificationConfigWriteQueryMock(ctx MockQueryContext) {
    for _, lang := range ctx.Languages  {
      itemsPerPage := 9999
      if (ctx.ItemsPerPage > 0) {
        itemsPerPage = ctx.ItemsPerPage
      }
      f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
      items, count, _ := NotificationConfigActionQuery(f)
      result := QueryEntitySuccessResult(f, items, count)
      WriteMockDataToFile(lang, "", "NotificationConfig", result)
    }
  }
var NotificationConfigImportExportCommands = []cli.Command{
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
			query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_CREATE},
      })
			NotificationConfigActionSeeder(query, c.Int("count"))
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
				Value: "notification-config-seeder.yml",
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
      query := CommonCliQueryDSLBuilderAuthorize(c, &SecurityModel{
        ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_CREATE},
      })
			NotificationConfigActionSeederInit(query, c.String("file"), c.String("format"))
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
				Value: "notification-config-seeder-notification-config.yml",
				// Uncomment before publish, they need to specify
				// Required: true,
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Format of the export or import file. Can be 'yaml', 'yml', 'json', 'sql', 'csv'",
				Value: "yaml",
			},
		},
		Usage: "Reads a yaml file containing an array of notification-configs, you can run this to validate if your import file is correct, and how it would look like after import",
		Action: func(c *cli.Context) error {
			data := &[]NotificationConfigEntity{}
			ReadYamlFile(c.String("file"), data)
			fmt.Println(data)
			return nil
		},
	},
	cli.Command{
		Name:    "import",
    Flags: append(
			append(
				CommonQueryFlags,
				&cli.StringFlag{
					Name:     "file",
					Usage:    "The address of file you want the csv be imported from",
					Required: true,
				}),
			NotificationConfigCommonCliFlagsOptional...,
		),
		Usage: "imports csv/yaml/json file and place it and its children into database",
		Action: func(c *cli.Context) error {
			CommonCliImportCmdAuthorized(c,
				NotificationConfigActionCreate,
				reflect.ValueOf(&NotificationConfigEntity{}).Elem(),
				c.String("file"),
        &SecurityModel{
					ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_CREATE},
				},
        func() NotificationConfigEntity {
					v := CastNotificationConfigFromCli(c)
					return *v
				},
			)
			return nil
		},
	},
}
    var NotificationConfigCliCommands []cli.Command = []cli.Command{
      NOTIFICATION_CONFIG_ACTION_QUERY.ToCli(),
      NOTIFICATION_CONFIG_ACTION_TABLE.ToCli(),
      NotificationConfigCreateCmd,
      NotificationConfigUpdateCmd,
      NotificationConfigCreateInteractiveCmd,
      NotificationConfigWipeCmd,
      GetCommonRemoveQuery(reflect.ValueOf(&NotificationConfigEntity{}).Elem(), NotificationConfigActionRemove),
  }
  func NotificationConfigCliFn() cli.Command {
    NotificationConfigCliCommands = append(NotificationConfigCliCommands, NotificationConfigImportExportCommands...)
    return cli.Command{
      Name:        "notificationConfig",
      ShortName:   "config",
      Description: "NotificationConfigs module actions (sample module to handle complex entities)",
      Usage:       "Configuration for the notifications used in the app, such as default gsm number, email senders, and many more",
      Flags: []cli.Flag{
        &cli.StringFlag{
          Name:  "language",
          Value: "en",
        },
      },
      Subcommands: NotificationConfigCliCommands,
    }
  }
var NOTIFICATION_CONFIG_ACTION_TABLE = Module2Action{
  Name:    "table",
  ActionName: "table",
  ActionAliases: []string{"t"},
  Flags:  CommonQueryFlags,
  Description:   "Table formatted queries all of the entities in database based on the standard query format",
  Action: NotificationConfigActionQuery,
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    CommonCliTableCmd2(c,
      NotificationConfigActionQuery,
      security,
      reflect.ValueOf(&NotificationConfigEntity{}).Elem(),
    )
    return nil
  },
}
var NOTIFICATION_CONFIG_ACTION_QUERY = Module2Action{
  Method: "GET",
  Url:    "/notification-configs",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_QUERY},
  },
  Group: "notificationConfig",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpQueryEntity(c, NotificationConfigActionQuery)
    },
  },
  Format: "QUERY",
  Action: NotificationConfigActionQuery,
  ResponseEntity: &[]NotificationConfigEntity{},
  Out: Module2ActionBody{
		Entity: "NotificationConfigEntity",
	},
  CliAction: func(c *cli.Context, security *SecurityModel) error {
		CommonCliQueryCmd2(
			c,
			NotificationConfigActionQuery,
			security,
		)
		return nil
	},
	CliName:       "query",
	ActionName:    "query",
	ActionAliases: []string{"q"},
	Flags:         CommonQueryFlags,
	Description:   "Queries all of the entities in database based on the standard query format (s+)",
}
var NOTIFICATION_CONFIG_ACTION_EXPORT = Module2Action{
  Method: "GET",
  Url:    "/notification-configs/export",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_QUERY},
  },
  Group: "notificationConfig",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpStreamFileChannel(c, NotificationConfigActionExport)
    },
  },
  Format: "QUERY",
  Action: NotificationConfigActionExport,
  ResponseEntity: &[]NotificationConfigEntity{},
  Out: Module2ActionBody{
		Entity: "NotificationConfigEntity",
	},
}
var NOTIFICATION_CONFIG_ACTION_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/notification-config/:uniqueId",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_QUERY},
  },
  Group: "notificationConfig",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, NotificationConfigActionGetOne)
    },
  },
  Format: "GET_ONE",
  Action: NotificationConfigActionGetOne,
  ResponseEntity: &NotificationConfigEntity{},
  Out: Module2ActionBody{
		Entity: "NotificationConfigEntity",
	},
}
var NOTIFICATION_CONFIG_ACTION_POST_ONE = Module2Action{
  ActionName:    "create",
  ActionAliases: []string{"c"},
  Description: "Create new notificationConfig",
  Flags: NotificationConfigCommonCliFlags,
  Method: "POST",
  Url:    "/notification-config",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_CREATE},
  },
  Group: "notificationConfig",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpPostEntity(c, NotificationConfigActionCreate)
    },
  },
  CliAction: func(c *cli.Context, security *SecurityModel) error {
    result, err := CliPostEntity(c, NotificationConfigActionCreate, security)
    HandleActionInCli(c, result, err, map[string]map[string]string{})
    return err
  },
  Action: NotificationConfigActionCreate,
  Format: "POST_ONE",
  RequestEntity: &NotificationConfigEntity{},
  ResponseEntity: &NotificationConfigEntity{},
  Out: Module2ActionBody{
		Entity: "NotificationConfigEntity",
	},
  In: Module2ActionBody{
		Entity: "NotificationConfigEntity",
	},
}
var NOTIFICATION_CONFIG_ACTION_PATCH = Module2Action{
  ActionName:    "update",
  ActionAliases: []string{"u"},
  Flags: NotificationConfigCommonCliFlagsOptional,
  Method: "PATCH",
  Url:    "/notification-config",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_UPDATE},
  },
  Group: "notificationConfig",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, NotificationConfigActionUpdate)
    },
  },
  Action: NotificationConfigActionUpdate,
  RequestEntity: &NotificationConfigEntity{},
  ResponseEntity: &NotificationConfigEntity{},
  Format: "PATCH_ONE",
  Out: Module2ActionBody{
		Entity: "NotificationConfigEntity",
	},
  In: Module2ActionBody{
		Entity: "NotificationConfigEntity",
	},
}
var NOTIFICATION_CONFIG_ACTION_PATCH_BULK = Module2Action{
  Method: "PATCH",
  Url:    "/notification-configs",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_UPDATE},
  },
  Group: "notificationConfig",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntities(c, NotificationConfigActionBulkUpdate)
    },
  },
  Action: NotificationConfigActionBulkUpdate,
  Format: "PATCH_BULK",
  RequestEntity:  &BulkRecordRequest[NotificationConfigEntity]{},
  ResponseEntity: &BulkRecordRequest[NotificationConfigEntity]{},
  Out: Module2ActionBody{
		Entity: "NotificationConfigEntity",
	},
  In: Module2ActionBody{
		Entity: "NotificationConfigEntity",
	},
}
var NOTIFICATION_CONFIG_ACTION_DELETE = Module2Action{
  Method: "DELETE",
  Url:    "/notification-config",
  Format: "DELETE_DSL",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_DELETE},
  },
  Group: "notificationConfig",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpRemoveEntity(c, NotificationConfigActionRemove)
    },
  },
  Action: NotificationConfigActionRemove,
  RequestEntity: &DeleteRequest{},
  ResponseEntity: &DeleteResponse{},
  TargetEntity: &NotificationConfigEntity{},
}
var NOTIFICATION_CONFIG_ACTION_DISTINCT_PATCH_ONE = Module2Action{
  Method: "PATCH",
  Url:    "/notification-config/distinct",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_UPDATE_DISTINCT_WORKSPACE},
  },
  Group: "notificationConfig",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpUpdateEntity(c, NotificationConfigDistinctActionUpdate)
    },
  },
  Action: NotificationConfigDistinctActionUpdate,
  Format: "PATCH_ONE",
  RequestEntity: &NotificationConfigEntity{},
  ResponseEntity: &NotificationConfigEntity{},
  Out: Module2ActionBody{
		Entity: "NotificationConfigEntity",
	},
  In: Module2ActionBody{
		Entity: "NotificationConfigEntity",
	},
}
var NOTIFICATION_CONFIG_ACTION_DISTINCT_GET_ONE = Module2Action{
  Method: "GET",
  Url:    "/notification-config/distinct",
  SecurityModel: &SecurityModel{
    ActionRequires: []PermissionInfo{PERM_ROOT_NOTIFICATION_CONFIG_GET_DISTINCT_WORKSPACE},
  },
  Group: "notificationConfig",
  Handlers: []gin.HandlerFunc{
    func (c *gin.Context) {
      HttpGetEntity(c, NotificationConfigDistinctActionGetOne)
    },
  },
  Action: NotificationConfigDistinctActionGetOne,
  Format: "GET_ONE",
  ResponseEntity: &NotificationConfigEntity{},
  Out: Module2ActionBody{
		Entity: "NotificationConfigEntity",
	},
}
  /**
  *	Override this function on NotificationConfigEntityHttp.go,
  *	In order to add your own http
  **/
  var AppendNotificationConfigRouter = func(r *[]Module2Action) {}
  func GetNotificationConfigModule2Actions() []Module2Action {
    routes := []Module2Action{
      NOTIFICATION_CONFIG_ACTION_QUERY,
      NOTIFICATION_CONFIG_ACTION_EXPORT,
      NOTIFICATION_CONFIG_ACTION_GET_ONE,
      NOTIFICATION_CONFIG_ACTION_POST_ONE,
      NOTIFICATION_CONFIG_ACTION_PATCH,
      NOTIFICATION_CONFIG_ACTION_PATCH_BULK,
      NOTIFICATION_CONFIG_ACTION_DELETE,
      NOTIFICATION_CONFIG_ACTION_DISTINCT_PATCH_ONE,
      NOTIFICATION_CONFIG_ACTION_DISTINCT_GET_ONE,
    }
    // Append user defined functions
    AppendNotificationConfigRouter(&routes)
    return routes
  }
  func CreateNotificationConfigRouter(r *gin.Engine) []Module2Action {
    httpRoutes := GetNotificationConfigModule2Actions()
    CastRoutes(httpRoutes, r)
    WriteHttpInformationToFile(&httpRoutes, NotificationConfigEntityJsonSchema, "notification-config-http", "workspaces")
    WriteEntitySchema("NotificationConfigEntity", NotificationConfigEntityJsonSchema, "workspaces")
    return httpRoutes
  }
var PERM_ROOT_NOTIFICATION_CONFIG_DELETE = PermissionInfo{
  CompleteKey: "root/workspaces/notification-config/delete",
  Name: "Delete notification config",
}
var PERM_ROOT_NOTIFICATION_CONFIG_CREATE = PermissionInfo{
  CompleteKey: "root/workspaces/notification-config/create",
  Name: "Create notification config",
}
var PERM_ROOT_NOTIFICATION_CONFIG_UPDATE = PermissionInfo{
  CompleteKey: "root/workspaces/notification-config/update",
  Name: "Update notification config",
}
var PERM_ROOT_NOTIFICATION_CONFIG_QUERY = PermissionInfo{
  CompleteKey: "root/workspaces/notification-config/query",
  Name: "Query notification config",
}
  var PERM_ROOT_NOTIFICATION_CONFIG_GET_DISTINCT_WORKSPACE = PermissionInfo{
    CompleteKey: "root/workspaces/notification-config/get-distinct-workspace",
    Name: "Get notification config Distinct",
  }
  var PERM_ROOT_NOTIFICATION_CONFIG_UPDATE_DISTINCT_WORKSPACE = PermissionInfo{
    CompleteKey: "root/workspaces/notification-config/update-distinct-workspace",
    Name: "Update notification config Distinct",
  }
var PERM_ROOT_NOTIFICATION_CONFIG = PermissionInfo{
  CompleteKey: "root/workspaces/notification-config/*",
  Name: "Entire notification config actions (*)",
}
var ALL_NOTIFICATION_CONFIG_PERMISSIONS = []PermissionInfo{
	PERM_ROOT_NOTIFICATION_CONFIG_DELETE,
	PERM_ROOT_NOTIFICATION_CONFIG_CREATE,
	PERM_ROOT_NOTIFICATION_CONFIG_UPDATE,
    PERM_ROOT_NOTIFICATION_CONFIG_GET_DISTINCT_WORKSPACE,
    PERM_ROOT_NOTIFICATION_CONFIG_UPDATE_DISTINCT_WORKSPACE,
	PERM_ROOT_NOTIFICATION_CONFIG_QUERY,
	PERM_ROOT_NOTIFICATION_CONFIG,
}
  func NotificationConfigDistinctActionUpdate(
    query QueryDSL,
    fields *NotificationConfigEntity,
  ) (*NotificationConfigEntity, *IError) {
    query.UniqueId = query.UserId
    entity, err := NotificationConfigActionGetOne(query)
    if err != nil || entity.UniqueId == "" {
      fields.UniqueId = query.UserId
      return NotificationConfigActionCreateFn(fields, query)
    } else {
      fields.UniqueId = query.UniqueId
      return NotificationConfigActionUpdateFn(query, fields)
    }
  }
  func NotificationConfigDistinctActionGetOne(
    query QueryDSL,
  ) (*NotificationConfigEntity, *IError) {
    query.UniqueId = query.UserId
    entity, err := NotificationConfigActionGetOne(query)
    if err != nil && err.HttpCode == 404 {
      return &NotificationConfigEntity{}, nil
    }
    return entity, err
  }