/*
 *	Generated by fireback 1.1.16
 *	Written by Ali Torabi.
 *	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
package com.fireback.modules.workspaces;

import android.content.Context;
import com.fireback.ArrayResponse;
import com.fireback.FirebackConfig;
import com.fireback.SessionManager;
import io.reactivex.rxjava3.core.Observable;
import io.reactivex.rxjava3.schedulers.Schedulers;
import java.io.IOException;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;

/*
 */
public class GetRolesExport {
  private Context context;

  public GetRolesExport(Context ctx) {
    context = ctx;
  }

  private String getUrl() {
    return FirebackConfig.getInstance().BuildUrl("/roles/export");
  }

  private Response makeHttpRequest() throws IOException {
    OkHttpClient client = new OkHttpClient();
    Request request =
        new Request.Builder()
            .header("authorization", SessionManager.getInstance(context).getUserSession().token)
            .url(getUrl())
            .build();
    return client.newCall(request).execute();
  }

  public Observable<ArrayResponse<RoleEntity>> query() {
    return Observable.just("")
        .observeOn(Schedulers.io())
        .map(tick -> makeHttpRequest())
        .map(
            response -> {
              if (response.isSuccessful()) {
                ArrayResponse<RoleEntity> res =
                    ArrayResponse.fromJson(response.body().string(), RoleEntity.class);
                response.close();
                return res;
              } else {
                throw new IOException("Request failed with code: " + response.code());
              }
            });
  }
}
