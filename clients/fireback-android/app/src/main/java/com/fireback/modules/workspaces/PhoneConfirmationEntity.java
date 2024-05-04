package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class PhoneConfirmationEntity extends JsonSerializable {
    public UserEntity user;
    public String status;
    public String phoneNumber;
    public String key;
    public String expiresAt;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
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
    // upper: PhoneNumber phoneNumber
    private MutableLiveData< String > phoneNumber = new MutableLiveData<>();
    public MutableLiveData< String > getPhoneNumber() {
        return phoneNumber;
    }
    public void setPhoneNumber( String  v) {
        phoneNumber.setValue(v);
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
    // Handling error message for each field
    // upper: User user
    private MutableLiveData<String> userMsg = new MutableLiveData<>();
    public MutableLiveData<String> getUserMsg() {
        return userMsg;
    }
    public void setUserMsg(String v) {
        userMsg.setValue(v);
    }
    // upper: Status status
    private MutableLiveData<String> statusMsg = new MutableLiveData<>();
    public MutableLiveData<String> getStatusMsg() {
        return statusMsg;
    }
    public void setStatusMsg(String v) {
        statusMsg.setValue(v);
    }
    // upper: PhoneNumber phoneNumber
    private MutableLiveData<String> phoneNumberMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPhoneNumberMsg() {
        return phoneNumberMsg;
    }
    public void setPhoneNumberMsg(String v) {
        phoneNumberMsg.setValue(v);
    }
    // upper: Key key
    private MutableLiveData<String> keyMsg = new MutableLiveData<>();
    public MutableLiveData<String> getKeyMsg() {
        return keyMsg;
    }
    public void setKeyMsg(String v) {
        keyMsg.setValue(v);
    }
    // upper: ExpiresAt expiresAt
    private MutableLiveData<String> expiresAtMsg = new MutableLiveData<>();
    public MutableLiveData<String> getExpiresAtMsg() {
        return expiresAtMsg;
    }
    public void setExpiresAtMsg(String v) {
        expiresAtMsg.setValue(v);
    }
  }
}