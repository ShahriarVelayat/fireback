package worldtimezone

import (
	"embed"

	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func LicensesModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "worldtimezone",
		Definitions: &Module2Definitions,
	}

	module.ProvidePermissionHandler(ALL_TIMEZONEGROUP_PERMISSIONS)

	module.Actions = [][]workspaces.Module2Action{
		GetTimezoneGroupModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) {
		dbref.AutoMigrate(&TimezoneGroupEntity{})
		dbref.AutoMigrate(&TimezoneGroupEntityPolyglot{})
		dbref.AutoMigrate(&TimezoneGroupUtcItems{})
	})

	module.ProvideCliHandlers([]cli.Command{
		TimezoneGroupCliFn(),
	})

	return module
}
