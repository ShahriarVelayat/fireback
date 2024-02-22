package com.fireback.modules.workspaces;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class ClassicSignupAction {
    public static class Req extends JsonSerializable {
    public String value;
    public String type;
    public String password;
    public String firstName;
    public String lastName;
    public String inviteId;
    public String publicJoinKeyId;
    public String workspaceTypeId;
    }
    public static class ReqViewModel extends ViewModel {
    // upper: Value value
    private MutableLiveData< String > value = new MutableLiveData<>();
    public MutableLiveData< String > getValue() {
        return value;
    }
    public void setValue( String  v) {
        value.setValue(v);
    }
    // upper: Type type
    private MutableLiveData< String > type = new MutableLiveData<>();
    public MutableLiveData< String > getType() {
        return type;
    }
    public void setType( String  v) {
        type.setValue(v);
    }
    // upper: Password password
    private MutableLiveData< String > password = new MutableLiveData<>();
    public MutableLiveData< String > getPassword() {
        return password;
    }
    public void setPassword( String  v) {
        password.setValue(v);
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
    // upper: InviteId inviteId
    private MutableLiveData< String > inviteId = new MutableLiveData<>();
    public MutableLiveData< String > getInviteId() {
        return inviteId;
    }
    public void setInviteId( String  v) {
        inviteId.setValue(v);
    }
    // upper: PublicJoinKeyId publicJoinKeyId
    private MutableLiveData< String > publicJoinKeyId = new MutableLiveData<>();
    public MutableLiveData< String > getPublicJoinKeyId() {
        return publicJoinKeyId;
    }
    public void setPublicJoinKeyId( String  v) {
        publicJoinKeyId.setValue(v);
    }
    // upper: WorkspaceTypeId workspaceTypeId
    private MutableLiveData< String > workspaceTypeId = new MutableLiveData<>();
    public MutableLiveData< String > getWorkspaceTypeId() {
        return workspaceTypeId;
    }
    public void setWorkspaceTypeId( String  v) {
        workspaceTypeId.setValue(v);
    }
    }
}