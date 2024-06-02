/*
 *	Generated by fireback 1.1.16
 *	Written by Ali Torabi.
 *	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
package com.fireback.modules.workspaces;

import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.JsonSerializable;

public class WorkspaceRoleEntity extends JsonSerializable {
  public UserWorkspaceEntity userWorkspace;
  public RoleEntity role;

  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: UserWorkspace userWorkspace
    private MutableLiveData<UserWorkspaceEntity> userWorkspace = new MutableLiveData<>();

    public MutableLiveData<UserWorkspaceEntity> getUserWorkspace() {
      return userWorkspace;
    }

    public void setUserWorkspace(UserWorkspaceEntity v) {
      userWorkspace.setValue(v);
    }

    // upper: Role role
    private MutableLiveData<RoleEntity> role = new MutableLiveData<>();

    public MutableLiveData<RoleEntity> getRole() {
      return role;
    }

    public void setRole(RoleEntity v) {
      role.setValue(v);
    }

    // Handling error message for each field
    // upper: UserWorkspace userWorkspace
    private MutableLiveData<String> userWorkspaceMsg = new MutableLiveData<>();

    public MutableLiveData<String> getUserWorkspaceMsg() {
      return userWorkspaceMsg;
    }

    public void setUserWorkspaceMsg(String v) {
      userWorkspaceMsg.setValue(v);
    }

    // upper: Role role
    private MutableLiveData<String> roleMsg = new MutableLiveData<>();

    public MutableLiveData<String> getRoleMsg() {
      return roleMsg;
    }

    public void setRoleMsg(String v) {
      roleMsg.setValue(v);
    }
  }
}
