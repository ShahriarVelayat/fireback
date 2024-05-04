package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class WorkspaceConfigEntity extends JsonSerializable {
    public int disablePublicWorkspaceCreation;
    public WorkspaceEntity workspace;
    public String zoomClientId;
    public String zoomClientSecret;
    public Boolean allowPublicToJoinTheWorkspace;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: DisablePublicWorkspaceCreation disablePublicWorkspaceCreation
    private MutableLiveData< Integer > disablePublicWorkspaceCreation = new MutableLiveData<>();
    public MutableLiveData< Integer > getDisablePublicWorkspaceCreation() {
        return disablePublicWorkspaceCreation;
    }
    public void setDisablePublicWorkspaceCreation( Integer  v) {
        disablePublicWorkspaceCreation.setValue(v);
    }
    // upper: Workspace workspace
    private MutableLiveData< WorkspaceEntity > workspace = new MutableLiveData<>();
    public MutableLiveData< WorkspaceEntity > getWorkspace() {
        return workspace;
    }
    public void setWorkspace( WorkspaceEntity  v) {
        workspace.setValue(v);
    }
    // upper: ZoomClientId zoomClientId
    private MutableLiveData< String > zoomClientId = new MutableLiveData<>();
    public MutableLiveData< String > getZoomClientId() {
        return zoomClientId;
    }
    public void setZoomClientId( String  v) {
        zoomClientId.setValue(v);
    }
    // upper: ZoomClientSecret zoomClientSecret
    private MutableLiveData< String > zoomClientSecret = new MutableLiveData<>();
    public MutableLiveData< String > getZoomClientSecret() {
        return zoomClientSecret;
    }
    public void setZoomClientSecret( String  v) {
        zoomClientSecret.setValue(v);
    }
    // upper: AllowPublicToJoinTheWorkspace allowPublicToJoinTheWorkspace
    private MutableLiveData< Boolean > allowPublicToJoinTheWorkspace = new MutableLiveData<>();
    public MutableLiveData< Boolean > getAllowPublicToJoinTheWorkspace() {
        return allowPublicToJoinTheWorkspace;
    }
    public void setAllowPublicToJoinTheWorkspace( Boolean  v) {
        allowPublicToJoinTheWorkspace.setValue(v);
    }
    // Handling error message for each field
    // upper: DisablePublicWorkspaceCreation disablePublicWorkspaceCreation
    private MutableLiveData<String> disablePublicWorkspaceCreationMsg = new MutableLiveData<>();
    public MutableLiveData<String> getDisablePublicWorkspaceCreationMsg() {
        return disablePublicWorkspaceCreationMsg;
    }
    public void setDisablePublicWorkspaceCreationMsg(String v) {
        disablePublicWorkspaceCreationMsg.setValue(v);
    }
    // upper: Workspace workspace
    private MutableLiveData<String> workspaceMsg = new MutableLiveData<>();
    public MutableLiveData<String> getWorkspaceMsg() {
        return workspaceMsg;
    }
    public void setWorkspaceMsg(String v) {
        workspaceMsg.setValue(v);
    }
    // upper: ZoomClientId zoomClientId
    private MutableLiveData<String> zoomClientIdMsg = new MutableLiveData<>();
    public MutableLiveData<String> getZoomClientIdMsg() {
        return zoomClientIdMsg;
    }
    public void setZoomClientIdMsg(String v) {
        zoomClientIdMsg.setValue(v);
    }
    // upper: ZoomClientSecret zoomClientSecret
    private MutableLiveData<String> zoomClientSecretMsg = new MutableLiveData<>();
    public MutableLiveData<String> getZoomClientSecretMsg() {
        return zoomClientSecretMsg;
    }
    public void setZoomClientSecretMsg(String v) {
        zoomClientSecretMsg.setValue(v);
    }
    // upper: AllowPublicToJoinTheWorkspace allowPublicToJoinTheWorkspace
    private MutableLiveData<String> allowPublicToJoinTheWorkspaceMsg = new MutableLiveData<>();
    public MutableLiveData<String> getAllowPublicToJoinTheWorkspaceMsg() {
        return allowPublicToJoinTheWorkspaceMsg;
    }
    public void setAllowPublicToJoinTheWorkspaceMsg(String v) {
        allowPublicToJoinTheWorkspaceMsg.setValue(v);
    }
  }
}