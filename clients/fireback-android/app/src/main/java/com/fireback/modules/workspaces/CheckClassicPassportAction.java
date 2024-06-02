/*
 *	Generated by fireback 1.1.16
 *	Written by Ali Torabi.
 *	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
package com.fireback.modules.workspaces;

import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.JsonSerializable;
import com.fireback.ResponseErrorException;

public class CheckClassicPassportAction {
  public static class Req extends JsonSerializable {
    public String value;
    // upper: Value value
    private MutableLiveData<String> valueMsg = new MutableLiveData<>();

    public MutableLiveData<String> getValueMsg() {
      return valueMsg;
    }

    public void setValueMsg(String v) {
      valueMsg.setValue(v);
    }
  }

  public static class ReqViewModel extends ViewModel {
    // upper: Value value
    private MutableLiveData<String> value = new MutableLiveData<>();

    public MutableLiveData<String> getValue() {
      return value;
    }

    public void setValue(String v) {
      value.setValue(v);
    }

    // upper: Value value
    private MutableLiveData<String> valueMsg = new MutableLiveData<>();

    public MutableLiveData<String> getValueMsg() {
      return valueMsg;
    }

    public void setValueMsg(String v) {
      valueMsg.setValue(v);
    }

    public void applyException(Throwable e) {
      if (!(e instanceof ResponseErrorException)) {
        return;
      }
      ResponseErrorException responseError = (ResponseErrorException) e;
      // @todo on fireback: This needs to be recursive.
      responseError.error.errors.forEach(
          item -> {
            if (item.location != null && item.location.equals("value")) {
              this.setValueMsg(item.messageTranslated);
            }
          });
    }
  }

  public static class Res extends JsonSerializable {
    public Boolean exists;
    // upper: Exists exists
    private MutableLiveData<String> existsMsg = new MutableLiveData<>();

    public MutableLiveData<String> getExistsMsg() {
      return existsMsg;
    }

    public void setExistsMsg(String v) {
      existsMsg.setValue(v);
    }
  }
}
