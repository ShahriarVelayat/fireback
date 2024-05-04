package com.fireback.modules.licenses;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
class LicensePermissions extends JsonSerializable {
    public com.fireback.modules.workspaces.CapabilityEntity capability;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Capability capability
    private MutableLiveData<com.fireback.modules. workspaces.CapabilityEntity > capability = new MutableLiveData<>();
    public MutableLiveData<com.fireback.modules. workspaces.CapabilityEntity > getCapability() {
        return capability;
    }
    public void setCapability(com.fireback.modules. workspaces.CapabilityEntity  v) {
        capability.setValue(v);
    }
    // Handling error message for each field
    // upper: Capability capability
    private MutableLiveData<String> capabilityMsg = new MutableLiveData<>();
    public MutableLiveData<String> getCapabilityMsg() {
        return capabilityMsg;
    }
    public void setCapabilityMsg(String v) {
        capabilityMsg.setValue(v);
    }
  }
}
public class LicenseEntity extends JsonSerializable {
    public String name;
    public String signedLicense;
    public java.util.Date validityStartDate;
    public java.util.Date validityEndDate;
    public LicensePermissions[] permissions;
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
    // upper: SignedLicense signedLicense
    private MutableLiveData< String > signedLicense = new MutableLiveData<>();
    public MutableLiveData< String > getSignedLicense() {
        return signedLicense;
    }
    public void setSignedLicense( String  v) {
        signedLicense.setValue(v);
    }
    // upper: ValidityStartDate validityStartDate
    private MutableLiveData< java.util.Date > validityStartDate = new MutableLiveData<>();
    public MutableLiveData< java.util.Date > getValidityStartDate() {
        return validityStartDate;
    }
    public void setValidityStartDate( java.util.Date  v) {
        validityStartDate.setValue(v);
    }
    // upper: ValidityEndDate validityEndDate
    private MutableLiveData< java.util.Date > validityEndDate = new MutableLiveData<>();
    public MutableLiveData< java.util.Date > getValidityEndDate() {
        return validityEndDate;
    }
    public void setValidityEndDate( java.util.Date  v) {
        validityEndDate.setValue(v);
    }
    // upper: Permissions permissions
    private MutableLiveData< LicensePermissions[] > permissions = new MutableLiveData<>();
    public MutableLiveData< LicensePermissions[] > getPermissions() {
        return permissions;
    }
    public void setPermissions( LicensePermissions[]  v) {
        permissions.setValue(v);
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
    // upper: SignedLicense signedLicense
    private MutableLiveData<String> signedLicenseMsg = new MutableLiveData<>();
    public MutableLiveData<String> getSignedLicenseMsg() {
        return signedLicenseMsg;
    }
    public void setSignedLicenseMsg(String v) {
        signedLicenseMsg.setValue(v);
    }
    // upper: ValidityStartDate validityStartDate
    private MutableLiveData<String> validityStartDateMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValidityStartDateMsg() {
        return validityStartDateMsg;
    }
    public void setValidityStartDateMsg(String v) {
        validityStartDateMsg.setValue(v);
    }
    // upper: ValidityEndDate validityEndDate
    private MutableLiveData<String> validityEndDateMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValidityEndDateMsg() {
        return validityEndDateMsg;
    }
    public void setValidityEndDateMsg(String v) {
        validityEndDateMsg.setValue(v);
    }
    // upper: Permissions permissions
    private MutableLiveData<String> permissionsMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPermissionsMsg() {
        return permissionsMsg;
    }
    public void setPermissionsMsg(String v) {
        permissionsMsg.setValue(v);
    }
  }
}