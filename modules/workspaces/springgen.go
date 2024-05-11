package workspaces

import (
	firebackspring "github.com/torabian/fireback/modules/workspaces/codegen/spring"
	firebackinclude "github.com/torabian/fireback/modules/workspaces/codegen/spring/include"
)

func SpringComputedField(field *Module2Field, isWorkspace bool) string {
	return JavaComputedField(field, isWorkspace)
}

func SpringEntityDiskName(x *Module2Entity) string {
	return ToUpper(x.Name) + "Entity.java"
}

var SpringGenCatalog CodeGenCatalog = CodeGenCatalog{
	LanguageName:            "FirebackSpring",
	ComputeField:            SpringComputedField,
	IncludeDirectory:        &firebackinclude.SpringInclude,
	Templates:               firebackspring.FbSpringTpl,
	EntityGeneratorTemplate: "SpringEntity.tpl",
	EntityDiskName:          SpringEntityDiskName,
}