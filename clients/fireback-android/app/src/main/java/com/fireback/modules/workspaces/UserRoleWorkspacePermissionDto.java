package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class UserRoleWorkspacePermissionDto extends JsonSerializable {
    public String workspaceId;
    public String userId;
    public String roleId;
    public String capabilityId;
    public String type;
    public static class VM extends ViewModel {
    // upper: WorkspaceId workspaceId
    private MutableLiveData< String > workspaceId = new MutableLiveData<>();
    public MutableLiveData< String > getWorkspaceId() {
        return workspaceId;
    }
    public void setWorkspaceId( String  v) {
        workspaceId.setValue(v);
    }
    // upper: UserId userId
    private MutableLiveData< String > userId = new MutableLiveData<>();
    public MutableLiveData< String > getUserId() {
        return userId;
    }
    public void setUserId( String  v) {
        userId.setValue(v);
    }
    // upper: RoleId roleId
    private MutableLiveData< String > roleId = new MutableLiveData<>();
    public MutableLiveData< String > getRoleId() {
        return roleId;
    }
    public void setRoleId( String  v) {
        roleId.setValue(v);
    }
    // upper: CapabilityId capabilityId
    private MutableLiveData< String > capabilityId = new MutableLiveData<>();
    public MutableLiveData< String > getCapabilityId() {
        return capabilityId;
    }
    public void setCapabilityId( String  v) {
        capabilityId.setValue(v);
    }
    // upper: Type type
    private MutableLiveData< String > type = new MutableLiveData<>();
    public MutableLiveData< String > getType() {
        return type;
    }
    public void setType( String  v) {
        type.setValue(v);
    }
    }
}