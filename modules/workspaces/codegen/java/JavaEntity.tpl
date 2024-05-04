package com.fireback.modules.{{ .m.Path }};
{{ template "javaimport" . }}
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;

{{ range .children }}
class {{ .FullName }} extends JsonSerializable {
  {{ template "javaClassContent" . }}
}
{{ end }}


public class {{ .e.EntityName }} extends JsonSerializable {
  {{ template "javaClassContent" .e }}
}