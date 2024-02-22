package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class EmailConfirmationEntity extends JsonSerializable {
    public UserEntity user;
    public String status;
    public String email;
    public String key;
    public String expiresAt;
    public static class VM extends ViewModel {
    // upper: User user
    private MutableLiveData< UserEntity > user = new MutableLiveData<>();
    public MutableLiveData< UserEntity > getUser() {
        return user;
    }
    public void setUser( UserEntity  v) {
        user.setValue(v);
    }
    // upper: Status status
    private MutableLiveData< String > status = new MutableLiveData<>();
    public MutableLiveData< String > getStatus() {
        return status;
    }
    public void setStatus( String  v) {
        status.setValue(v);
    }
    // upper: Email email
    private MutableLiveData< String > email = new MutableLiveData<>();
    public MutableLiveData< String > getEmail() {
        return email;
    }
    public void setEmail( String  v) {
        email.setValue(v);
    }
    // upper: Key key
    private MutableLiveData< String > key = new MutableLiveData<>();
    public MutableLiveData< String > getKey() {
        return key;
    }
    public void setKey( String  v) {
        key.setValue(v);
    }
    // upper: ExpiresAt expiresAt
    private MutableLiveData< String > expiresAt = new MutableLiveData<>();
    public MutableLiveData< String > getExpiresAt() {
        return expiresAt;
    }
    public void setExpiresAt( String  v) {
        expiresAt.setValue(v);
    }
    }
}