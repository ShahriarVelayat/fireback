package workspaces
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
func CastPermissionInfoFromCli (c *cli.Context) *PermissionInfoDto {
	template := &PermissionInfoDto{}
      if c.IsSet("name") {
        value := c.String("name")
        template.Name = &value
      }
      if c.IsSet("description") {
        value := c.String("description")
        template.Description = &value
      }
      if c.IsSet("complete-key") {
        value := c.String("complete-key")
        template.CompleteKey = &value
      }
	return template
}
var PermissionInfoDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "description",
      Required: false,
      Usage:    "description",
    },
    &cli.StringFlag{
      Name:     "complete-key",
      Required: false,
      Usage:    "completeKey",
    },
}
type PermissionInfoDto struct {
    Name   *string `json:"name" yaml:"name"       `
    // Datenano also has a text representation
    Description   *string `json:"description" yaml:"description"       `
    // Datenano also has a text representation
    CompleteKey   *string `json:"completeKey" yaml:"completeKey"       `
    // Datenano also has a text representation
}
func (x* PermissionInfoDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* PermissionInfoDto) JsonPrint()  {
    fmt.Println(x.Json())
}
// This is an experimental way to create new dtos, with exluding the pointers as helper.
func NewPermissionInfoDto(
	Name string,
	Description string,
	CompleteKey string,
) PermissionInfoDto {
    return PermissionInfoDto{
	Name: &Name,
	Description: &Description,
	CompleteKey: &CompleteKey,
    }
}