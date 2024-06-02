package workspaces

/*
*	Generated by fireback 1.1.16
*	Written by Ali Torabi.
*	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"strings"
)

func CastOkayResponseFromCli(c *cli.Context) *OkayResponseDto {
	template := &OkayResponseDto{}
	return template
}

var OkayResponseDtoCommonCliFlagsOptional = []cli.Flag{
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
}

type OkayResponseDto struct {
}

func (x *OkayResponseDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	// Intentional trim (so strings lib is always imported)
	return strings.TrimSpace("")
}
func (x *OkayResponseDto) JsonPrint() {
	fmt.Println(x.Json())
}

// This is an experimental way to create new dtos, with exluding the pointers as helper.
func NewOkayResponseDto() OkayResponseDto {
	return OkayResponseDto{}
}
