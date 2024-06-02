/*
 *	Generated by fireback 1.1.16
 *	Written by Ali Torabi.
 *	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
package com.fireback.modules.geo;

import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.JsonSerializable;
import com.fireback.modules.workspaces.*;

public class GeoStateEntity extends JsonSerializable {
  public String name;

  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Name name
    private MutableLiveData<String> name = new MutableLiveData<>();

    public MutableLiveData<String> getName() {
      return name;
    }

    public void setName(String v) {
      name.setValue(v);
    }

    // Handling error message for each field
    // upper: Name name
    private MutableLiveData<String> nameMsg = new MutableLiveData<>();

    public MutableLiveData<String> getNameMsg() {
      return nameMsg;
    }

    public void setNameMsg(String v) {
      nameMsg.setValue(v);
    }
  }
}
