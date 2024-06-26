/*
 *	Generated by fireback 1.1.16
 *	Written by Ali Torabi.
 *	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
package com.fireback.modules.workspaces;

import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.JsonSerializable;

public class PublicJoinKeyEntity extends JsonSerializable {
  public RoleEntity role;
  public WorkspaceEntity workspace;

  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Role role
    private MutableLiveData<RoleEntity> role = new MutableLiveData<>();

    public MutableLiveData<RoleEntity> getRole() {
      return role;
    }

    public void setRole(RoleEntity v) {
      role.setValue(v);
    }

    // upper: Workspace workspace
    private MutableLiveData<WorkspaceEntity> workspace = new MutableLiveData<>();

    public MutableLiveData<WorkspaceEntity> getWorkspace() {
      return workspace;
    }

    public void setWorkspace(WorkspaceEntity v) {
      workspace.setValue(v);
    }

    // Handling error message for each field
    // upper: Role role
    private MutableLiveData<String> roleMsg = new MutableLiveData<>();

    public MutableLiveData<String> getRoleMsg() {
      return roleMsg;
    }

    public void setRoleMsg(String v) {
      roleMsg.setValue(v);
    }

    // upper: Workspace workspace
    private MutableLiveData<String> workspaceMsg = new MutableLiveData<>();

    public MutableLiveData<String> getWorkspaceMsg() {
      return workspaceMsg;
    }

    public void setWorkspaceMsg(String v) {
      workspaceMsg.setValue(v);
    }
  }
}
