/*
 *	Generated by fireback 1.1.16
 *	Written by Ali Torabi.
 *	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
package com.fireback.modules.workspaces;

import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.JsonSerializable;

public class PermissionInfoDto extends JsonSerializable {
  public String name;
  public String description;
  public String completeKey;

  public static class VM extends ViewModel {
    // upper: Name name
    private MutableLiveData<String> name = new MutableLiveData<>();

    public MutableLiveData<String> getName() {
      return name;
    }

    public void setName(String v) {
      name.setValue(v);
    }

    // upper: Description description
    private MutableLiveData<String> description = new MutableLiveData<>();

    public MutableLiveData<String> getDescription() {
      return description;
    }

    public void setDescription(String v) {
      description.setValue(v);
    }

    // upper: CompleteKey completeKey
    private MutableLiveData<String> completeKey = new MutableLiveData<>();

    public MutableLiveData<String> getCompleteKey() {
      return completeKey;
    }

    public void setCompleteKey(String v) {
      completeKey.setValue(v);
    }
  }
}
