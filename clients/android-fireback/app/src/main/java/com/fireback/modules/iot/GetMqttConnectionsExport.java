package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.fireback.FirebackConfig;
import java.util.concurrent.TimeUnit;
import io.reactivex.rxjava3.core.Observable;
import io.reactivex.rxjava3.schedulers.Schedulers;
import android.content.Context;
import com.fireback.ArrayResponse;
import com.fireback.SessionManager;
import com.fireback.SingleResponse;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import java.io.IOException;
import java.util.concurrent.TimeUnit;
/*
*/
public class GetMqttConnectionsExport {
    private Context context;
    public GetMqttConnectionsExport(Context ctx ) {
        context = ctx;
    }
    public static String Url  = FirebackConfig.getInstance().BuildUrl("/mqtt-connections/export");
    private Response makeHttpRequest() throws IOException {
        OkHttpClient client = new OkHttpClient();
        Request request = new Request.Builder()
                .header("authorization", SessionManager.getInstance(context).getUserSession().token)
                .url(Url)
                .build();
        return client.newCall(request).execute();
    }
    public Observable<ArrayResponse<MqttConnectionEntity>> query() {
        return Observable.just("")
            .observeOn(Schedulers.io())
            .map(tick -> makeHttpRequest())
            .map(response -> {
                if (response.isSuccessful()) {
                    ArrayResponse<MqttConnectionEntity> res = ArrayResponse.fromJson(response.body().string(), MqttConnectionEntity.class);
                    response.close();
                    return res;
                } else {
                    throw new IOException("Request failed with code: " + response.code());
                }
            });
    }
}
