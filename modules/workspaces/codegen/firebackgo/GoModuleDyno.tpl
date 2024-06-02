package {{ .m.Path }}

/*
*	Generated by fireback {{ .fv }}
*	Written by Ali Torabi.
*	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
*/

{{ if ne .wsprefix "" }}
import "github.com/torabian/fireback/modules/workspaces"
{{ end }}
// Module dynamic things comes here. Don't touch it :D

var PERM_ROOT_{{ .m.AllUpper }}_EVERYTHING = {{ .wsprefix }}PermissionInfo{
  CompleteKey: "root/{{.m.AllLower}}/*",
}
var ALL_PERM_{{ .m.AllUpper }}_MODULE = []{{ .wsprefix }}PermissionInfo{
  PERM_ROOT_{{ .m.AllUpper }}_EVERYTHING,
}
