package com.fireback.modules.geo;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class GeoProvinceEntity extends JsonSerializable {
    public String name;
    public GeoCountryEntity country;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Country country
    private MutableLiveData< GeoCountryEntity > country = new MutableLiveData<>();
    public MutableLiveData< GeoCountryEntity > getCountry() {
        return country;
    }
    public void setCountry( GeoCountryEntity  v) {
        country.setValue(v);
    }
    // Handling error message for each field
    // upper: Name name
    private MutableLiveData<String> nameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getNameMsg() {
        return nameMsg;
    }
    public void setNameMsg(String v) {
        nameMsg.setValue(v);
    }
    // upper: Country country
    private MutableLiveData<String> countryMsg = new MutableLiveData<>();
    public MutableLiveData<String> getCountryMsg() {
        return countryMsg;
    }
    public void setCountryMsg(String v) {
        countryMsg.setValue(v);
    }
  }
}