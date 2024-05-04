package com.fireback.modules.geo;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class GeoCountryEntity extends JsonSerializable {
    public String status;
    public String flag;
    public String commonName;
    public String officialName;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Status status
    private MutableLiveData< String > status = new MutableLiveData<>();
    public MutableLiveData< String > getStatus() {
        return status;
    }
    public void setStatus( String  v) {
        status.setValue(v);
    }
    // upper: Flag flag
    private MutableLiveData< String > flag = new MutableLiveData<>();
    public MutableLiveData< String > getFlag() {
        return flag;
    }
    public void setFlag( String  v) {
        flag.setValue(v);
    }
    // upper: CommonName commonName
    private MutableLiveData< String > commonName = new MutableLiveData<>();
    public MutableLiveData< String > getCommonName() {
        return commonName;
    }
    public void setCommonName( String  v) {
        commonName.setValue(v);
    }
    // upper: OfficialName officialName
    private MutableLiveData< String > officialName = new MutableLiveData<>();
    public MutableLiveData< String > getOfficialName() {
        return officialName;
    }
    public void setOfficialName( String  v) {
        officialName.setValue(v);
    }
    // Handling error message for each field
    // upper: Status status
    private MutableLiveData<String> statusMsg = new MutableLiveData<>();
    public MutableLiveData<String> getStatusMsg() {
        return statusMsg;
    }
    public void setStatusMsg(String v) {
        statusMsg.setValue(v);
    }
    // upper: Flag flag
    private MutableLiveData<String> flagMsg = new MutableLiveData<>();
    public MutableLiveData<String> getFlagMsg() {
        return flagMsg;
    }
    public void setFlagMsg(String v) {
        flagMsg.setValue(v);
    }
    // upper: CommonName commonName
    private MutableLiveData<String> commonNameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getCommonNameMsg() {
        return commonNameMsg;
    }
    public void setCommonNameMsg(String v) {
        commonNameMsg.setValue(v);
    }
    // upper: OfficialName officialName
    private MutableLiveData<String> officialNameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getOfficialNameMsg() {
        return officialNameMsg;
    }
    public void setOfficialNameMsg(String v) {
        officialNameMsg.setValue(v);
    }
  }
}