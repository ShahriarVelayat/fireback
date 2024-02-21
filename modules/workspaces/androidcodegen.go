package workspaces

import (
	javatpl "pixelplux.com/fireback/modules/workspaces/codegen/java"
	javainclude "pixelplux.com/fireback/modules/workspaces/codegen/java/include"
)

func JavaPrimitve(primitive string) string {
	switch primitive {
	case "string", "text":
		return "String"
	case "int64", "int32", "int":
		return "Integer"
	case "float64", "float32", "float":
		return "Float"
	case "bool":
		return "Boolean"
	case "double":
		return "Double"
	default:
		return "unknown"
	}
}

func JavaComputedField(field *Module2Field) string {
	switch field.Type {
	case "string", "text":
		return "String"
	case "one":
		return field.Target + ""
	case "int64", "int32", "int":
		return "int"
	case "float64", "float32", "float":
		return "float"
	case "html":
		// Think about this, make an object called Html maybe?
		return "String"
	case "json":
		return "String"
		// return TsCalcJsonField(field)
	case "array":
		return field.PublicName() + "[]"
	case "many2many":
		if field.Module != "" {
			return field.Module + "." + field.Target + "[]"
		}
		return field.Target + "[]"
	case "object":
		return field.PublicName()
	case "arrayP":
		return JavaPrimitve(field.Primitive) + "[]"
	case "bool":
		return "Boolean"
	case "enum":
		return "String"
	case "date":
		return "java.util.Date"
	case "Timestamp", "datenano":
		return "String"
	case "double":
		return "double"
	case "any":
		return "String"

	default:
		return field.Type
	}
}

func JavaEntityDiskName(x *Module2Entity) string {
	return ToUpper(x.Name) + "Entity.java"
}

func JavaDtoDiskName(x *Module2DtoBase) string {
	return ToUpper(x.Name) + "Dto.java"
}

func JavaFormDiskName(x *Module2Entity) string {
	return ToUpper(x.Name) + "Form.java"
}

func JavaRpcCommonDiskName(x *Module2Action) string {
	return ToUpper(x.GetFuncName()) + ".java"
}

func JavaActionDiskName(action *Module2Action, moduleName string) string {
	return ToUpper(action.Name) + "Action.java"
}

var JavaGenCatalog CodeGenCatalog = CodeGenCatalog{
	LanguageName:            "java",
	ComputeField:            JavaComputedField,
	RpcPostDiskName:         JavaRpcCommonDiskName,
	RpcPost:                 "JavaRpcPost.tpl",
	Templates:               javatpl.JavaTpls,
	IncludeDirectory:        &javainclude.JavaInclude,
	RpcQueryDiskName:        JavaRpcCommonDiskName,
	RpcQuery:                "JavaRpcQuery.tpl",
	EntityGeneratorTemplate: "JavaEntity.tpl",
	DtoGeneratorTemplate:    "JavaDto.tpl",
	ActionGeneratorTemplate: "JavaActionDto.tpl",
	SingleActionDiskName:    JavaActionDiskName,

	EntityDiskName: JavaEntityDiskName,
	DtoDiskName:    JavaDtoDiskName,
}
