package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class PassportEntity extends JsonSerializable {
    public String type;
    public UserEntity user;
    public String value;
    public String password;
    public Boolean confirmed;
    public String accessToken;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Type type
    private MutableLiveData< String > type = new MutableLiveData<>();
    public MutableLiveData< String > getType() {
        return type;
    }
    public void setType( String  v) {
        type.setValue(v);
    }
    // upper: User user
    private MutableLiveData< UserEntity > user = new MutableLiveData<>();
    public MutableLiveData< UserEntity > getUser() {
        return user;
    }
    public void setUser( UserEntity  v) {
        user.setValue(v);
    }
    // upper: Value value
    private MutableLiveData< String > value = new MutableLiveData<>();
    public MutableLiveData< String > getValue() {
        return value;
    }
    public void setValue( String  v) {
        value.setValue(v);
    }
    // upper: Password password
    private MutableLiveData< String > password = new MutableLiveData<>();
    public MutableLiveData< String > getPassword() {
        return password;
    }
    public void setPassword( String  v) {
        password.setValue(v);
    }
    // upper: Confirmed confirmed
    private MutableLiveData< Boolean > confirmed = new MutableLiveData<>();
    public MutableLiveData< Boolean > getConfirmed() {
        return confirmed;
    }
    public void setConfirmed( Boolean  v) {
        confirmed.setValue(v);
    }
    // upper: AccessToken accessToken
    private MutableLiveData< String > accessToken = new MutableLiveData<>();
    public MutableLiveData< String > getAccessToken() {
        return accessToken;
    }
    public void setAccessToken( String  v) {
        accessToken.setValue(v);
    }
    // Handling error message for each field
    // upper: Type type
    private MutableLiveData<String> typeMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTypeMsg() {
        return typeMsg;
    }
    public void setTypeMsg(String v) {
        typeMsg.setValue(v);
    }
    // upper: User user
    private MutableLiveData<String> userMsg = new MutableLiveData<>();
    public MutableLiveData<String> getUserMsg() {
        return userMsg;
    }
    public void setUserMsg(String v) {
        userMsg.setValue(v);
    }
    // upper: Value value
    private MutableLiveData<String> valueMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValueMsg() {
        return valueMsg;
    }
    public void setValueMsg(String v) {
        valueMsg.setValue(v);
    }
    // upper: Password password
    private MutableLiveData<String> passwordMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPasswordMsg() {
        return passwordMsg;
    }
    public void setPasswordMsg(String v) {
        passwordMsg.setValue(v);
    }
    // upper: Confirmed confirmed
    private MutableLiveData<String> confirmedMsg = new MutableLiveData<>();
    public MutableLiveData<String> getConfirmedMsg() {
        return confirmedMsg;
    }
    public void setConfirmedMsg(String v) {
        confirmedMsg.setValue(v);
    }
    // upper: AccessToken accessToken
    private MutableLiveData<String> accessTokenMsg = new MutableLiveData<>();
    public MutableLiveData<String> getAccessTokenMsg() {
        return accessTokenMsg;
    }
    public void setAccessTokenMsg(String v) {
        accessTokenMsg.setValue(v);
    }
  }
}