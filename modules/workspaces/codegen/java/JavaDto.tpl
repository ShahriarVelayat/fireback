package com.fireback.modules.{{ .m.Path }};
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;


{{ range .children }}
class {{ .FullName }} extends JsonSerializable {
  {{ template "definitionrow" .CompleteFields }}

}
{{ end }}


public class {{ .e.DtoName }} extends JsonSerializable {
    {{ template "definitionrow" .e.CompleteFields }}

    public static class VM extends ViewModel {
        {{ template "viewmodelrow" .e.CompleteFields }}
    }
}