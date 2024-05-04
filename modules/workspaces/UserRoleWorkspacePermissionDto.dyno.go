package workspaces
import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/urfave/cli"
)
func CastUserRoleWorkspacePermissionFromCli (c *cli.Context) *UserRoleWorkspacePermissionDto {
	template := &UserRoleWorkspacePermissionDto{}
      if c.IsSet("workspace-id") {
        value := c.String("workspace-id")
        template.WorkspaceId = &value
      }
      if c.IsSet("user-id") {
        value := c.String("user-id")
        template.UserId = &value
      }
      if c.IsSet("role-id") {
        value := c.String("role-id")
        template.RoleId = &value
      }
      if c.IsSet("capability-id") {
        value := c.String("capability-id")
        template.CapabilityId = &value
      }
      if c.IsSet("type") {
        value := c.String("type")
        template.Type = &value
      }
	return template
}
var UserRoleWorkspacePermissionDtoCommonCliFlagsOptional = []cli.Flag{
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
      Name:     "workspace-id",
      Required: false,
      Usage:    "workspaceId",
    },
    &cli.StringFlag{
      Name:     "user-id",
      Required: false,
      Usage:    "userId",
    },
    &cli.StringFlag{
      Name:     "role-id",
      Required: false,
      Usage:    "roleId",
    },
    &cli.StringFlag{
      Name:     "capability-id",
      Required: false,
      Usage:    "capabilityId",
    },
    &cli.StringFlag{
      Name:     "type",
      Required: false,
      Usage:    "type",
    },
}
type UserRoleWorkspacePermissionDto struct {
    WorkspaceId   *string `json:"workspaceId" yaml:"workspaceId"       `
    // Datenano also has a text representation
    UserId   *string `json:"userId" yaml:"userId"       `
    // Datenano also has a text representation
    RoleId   *string `json:"roleId" yaml:"roleId"       `
    // Datenano also has a text representation
    CapabilityId   *string `json:"capabilityId" yaml:"capabilityId"       `
    // Datenano also has a text representation
    Type   *string `json:"type" yaml:"type"       `
    // Datenano also has a text representation
}
func (x* UserRoleWorkspacePermissionDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x* UserRoleWorkspacePermissionDto) JsonPrint()  {
    fmt.Println(x.Json())
}
// This is an experimental way to create new dtos, with exluding the pointers as helper.
func NewUserRoleWorkspacePermissionDto(
	WorkspaceId string,
	UserId string,
	RoleId string,
	CapabilityId string,
	Type string,
) UserRoleWorkspacePermissionDto {
    return UserRoleWorkspacePermissionDto{
	WorkspaceId: &WorkspaceId,
	UserId: &UserId,
	RoleId: &RoleId,
	CapabilityId: &CapabilityId,
	Type: &Type,
    }
}