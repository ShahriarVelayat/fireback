package shop

import (
	"embed"
	"fmt"

	// "github.com/urfave/cli"
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func ShopModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "shop",
		Definitions: &Module2Definitions,
	}

	module.ProvideMockImportHandler(func() {
		// FormImportMocks()
	})

	module.ProvideSeederImportHandler(func() {

	})

	module.ProvideMockWriterHandler(func(languages []string) {

	})

	module.ProvidePermissionHandler(
		ALL_PRODUCT_PERMISSIONS,
		ALL_PRODUCTSUBMISSION_PERMISSIONS,
		ALL_BRAND_PERMISSIONS,
		ALL_CATEGORY_PERMISSIONS,
		ALL_TAG_PERMISSIONS,
	)

	module.Actions = [][]workspaces.Module2Action{
		GetProductSubmissionModule2Actions(),
		GetProductModule2Actions(),
		GetTagModule2Actions(),
		GetCategoryModule2Actions(),
		GetBrandModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) {
		if err := dbref.AutoMigrate(
			&ProductEntity{},
			&ProductFields{},
			&ProductSubmissionEntity{},
			&ProductSubmissionValues{},
			&ProductSubmissionPrice{},
			&ProductSubmissionPriceVariations{},
			&CategoryEntity{},
			&CategoryEntityPolyglot{},
			&TagEntity{},
			&TagEntityPolyglot{},
			&BrandEntity{},
			&BrandEntityPolyglot{},
		); err != nil {
			fmt.Println(err.Error())
		}

	})

	module.ProvideCliHandlers([]cli.Command{
		ProductCliFn(),
		ProductSubmissionCliFn(),
		CategoryCliFn(),
		TagCliFn(),
		BrandCliFn(),
	})

	return module
}

func QueryProductSubmissionsReact(query workspaces.QueryDSL, chanStream chan *workspaces.ReactiveSearchResultDto) {
	actionFnNavigate := "navigate"

	query.Query = "name %" + query.SearchPhrase + "%"
	items, _, _ := ProductSubmissionActionQuery(query)

	product := "$product"
	products := "products"
	for _, item := range items {
		loc := "/product-submission/" + item.UniqueId

		uid := workspaces.UUID()

		chanStream <- &workspaces.ReactiveSearchResultDto{
			Phrase:      item.Name,
			Icon:        &product,
			Description: item.Name,
			Group:       &products,
			ActionFn:    &actionFnNavigate,
			UiLocation:  &loc,
			UniqueId:    &uid,
		}
	}

}
