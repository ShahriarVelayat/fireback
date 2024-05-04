package com.fireback.modules.{{ .m.Path }};
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.ResponseErrorException;

public class {{ .action.Upper }}Action {
    {{ if .action.In.Fields }}

    public static class Req extends JsonSerializable {
        {{ template "definitionrow" .action.In.Fields }}
        {{ template "viewModelMessageRow" .action.In.Fields }}
    }
    
    
    public static class ReqViewModel extends ViewModel {
        {{ template "viewmodelrow" .action.In.Fields }}
        {{ template "viewModelMessageRow" .action.In.Fields }}
        {{ template "applyExceptionOnViewModel" .action.In.Fields }}
    }

    {{ end }}

    {{ if .action.Out.Fields }}

    public static class Res extends JsonSerializable {
        {{ template "definitionrow" .action.Out.Fields }}
        {{ template "viewModelMessageRow" .action.Out.Fields }}
    }

    {{ end }}
}