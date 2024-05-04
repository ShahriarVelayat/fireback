package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class WorkspaceInviteEntity extends JsonSerializable {
    public String coverLetter;
    public String targetUserLocale;
    public String value;
    public WorkspaceEntity workspace;
    public String firstName;
    public String lastName;
    public Boolean used;
    public RoleEntity role;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: CoverLetter coverLetter
    private MutableLiveData< String > coverLetter = new MutableLiveData<>();
    public MutableLiveData< String > getCoverLetter() {
        return coverLetter;
    }
    public void setCoverLetter( String  v) {
        coverLetter.setValue(v);
    }
    // upper: TargetUserLocale targetUserLocale
    private MutableLiveData< String > targetUserLocale = new MutableLiveData<>();
    public MutableLiveData< String > getTargetUserLocale() {
        return targetUserLocale;
    }
    public void setTargetUserLocale( String  v) {
        targetUserLocale.setValue(v);
    }
    // upper: Value value
    private MutableLiveData< String > value = new MutableLiveData<>();
    public MutableLiveData< String > getValue() {
        return value;
    }
    public void setValue( String  v) {
        value.setValue(v);
    }
    // upper: Workspace workspace
    private MutableLiveData< WorkspaceEntity > workspace = new MutableLiveData<>();
    public MutableLiveData< WorkspaceEntity > getWorkspace() {
        return workspace;
    }
    public void setWorkspace( WorkspaceEntity  v) {
        workspace.setValue(v);
    }
    // upper: FirstName firstName
    private MutableLiveData< String > firstName = new MutableLiveData<>();
    public MutableLiveData< String > getFirstName() {
        return firstName;
    }
    public void setFirstName( String  v) {
        firstName.setValue(v);
    }
    // upper: LastName lastName
    private MutableLiveData< String > lastName = new MutableLiveData<>();
    public MutableLiveData< String > getLastName() {
        return lastName;
    }
    public void setLastName( String  v) {
        lastName.setValue(v);
    }
    // upper: Used used
    private MutableLiveData< Boolean > used = new MutableLiveData<>();
    public MutableLiveData< Boolean > getUsed() {
        return used;
    }
    public void setUsed( Boolean  v) {
        used.setValue(v);
    }
    // upper: Role role
    private MutableLiveData< RoleEntity > role = new MutableLiveData<>();
    public MutableLiveData< RoleEntity > getRole() {
        return role;
    }
    public void setRole( RoleEntity  v) {
        role.setValue(v);
    }
    // Handling error message for each field
    // upper: CoverLetter coverLetter
    private MutableLiveData<String> coverLetterMsg = new MutableLiveData<>();
    public MutableLiveData<String> getCoverLetterMsg() {
        return coverLetterMsg;
    }
    public void setCoverLetterMsg(String v) {
        coverLetterMsg.setValue(v);
    }
    // upper: TargetUserLocale targetUserLocale
    private MutableLiveData<String> targetUserLocaleMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTargetUserLocaleMsg() {
        return targetUserLocaleMsg;
    }
    public void setTargetUserLocaleMsg(String v) {
        targetUserLocaleMsg.setValue(v);
    }
    // upper: Value value
    private MutableLiveData<String> valueMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValueMsg() {
        return valueMsg;
    }
    public void setValueMsg(String v) {
        valueMsg.setValue(v);
    }
    // upper: Workspace workspace
    private MutableLiveData<String> workspaceMsg = new MutableLiveData<>();
    public MutableLiveData<String> getWorkspaceMsg() {
        return workspaceMsg;
    }
    public void setWorkspaceMsg(String v) {
        workspaceMsg.setValue(v);
    }
    // upper: FirstName firstName
    private MutableLiveData<String> firstNameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getFirstNameMsg() {
        return firstNameMsg;
    }
    public void setFirstNameMsg(String v) {
        firstNameMsg.setValue(v);
    }
    // upper: LastName lastName
    private MutableLiveData<String> lastNameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getLastNameMsg() {
        return lastNameMsg;
    }
    public void setLastNameMsg(String v) {
        lastNameMsg.setValue(v);
    }
    // upper: Used used
    private MutableLiveData<String> usedMsg = new MutableLiveData<>();
    public MutableLiveData<String> getUsedMsg() {
        return usedMsg;
    }
    public void setUsedMsg(String v) {
        usedMsg.setValue(v);
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