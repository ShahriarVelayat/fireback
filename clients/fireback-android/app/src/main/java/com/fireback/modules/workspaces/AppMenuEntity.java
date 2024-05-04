package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class AppMenuEntity extends JsonSerializable {
    public String href;
    public String icon;
    public String label;
    public String activeMatcher;
    public String applyType;
    public CapabilityEntity capability;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Href href
    private MutableLiveData< String > href = new MutableLiveData<>();
    public MutableLiveData< String > getHref() {
        return href;
    }
    public void setHref( String  v) {
        href.setValue(v);
    }
    // upper: Icon icon
    private MutableLiveData< String > icon = new MutableLiveData<>();
    public MutableLiveData< String > getIcon() {
        return icon;
    }
    public void setIcon( String  v) {
        icon.setValue(v);
    }
    // upper: Label label
    private MutableLiveData< String > label = new MutableLiveData<>();
    public MutableLiveData< String > getLabel() {
        return label;
    }
    public void setLabel( String  v) {
        label.setValue(v);
    }
    // upper: ActiveMatcher activeMatcher
    private MutableLiveData< String > activeMatcher = new MutableLiveData<>();
    public MutableLiveData< String > getActiveMatcher() {
        return activeMatcher;
    }
    public void setActiveMatcher( String  v) {
        activeMatcher.setValue(v);
    }
    // upper: ApplyType applyType
    private MutableLiveData< String > applyType = new MutableLiveData<>();
    public MutableLiveData< String > getApplyType() {
        return applyType;
    }
    public void setApplyType( String  v) {
        applyType.setValue(v);
    }
    // upper: Capability capability
    private MutableLiveData< CapabilityEntity > capability = new MutableLiveData<>();
    public MutableLiveData< CapabilityEntity > getCapability() {
        return capability;
    }
    public void setCapability( CapabilityEntity  v) {
        capability.setValue(v);
    }
    // Handling error message for each field
    // upper: Href href
    private MutableLiveData<String> hrefMsg = new MutableLiveData<>();
    public MutableLiveData<String> getHrefMsg() {
        return hrefMsg;
    }
    public void setHrefMsg(String v) {
        hrefMsg.setValue(v);
    }
    // upper: Icon icon
    private MutableLiveData<String> iconMsg = new MutableLiveData<>();
    public MutableLiveData<String> getIconMsg() {
        return iconMsg;
    }
    public void setIconMsg(String v) {
        iconMsg.setValue(v);
    }
    // upper: Label label
    private MutableLiveData<String> labelMsg = new MutableLiveData<>();
    public MutableLiveData<String> getLabelMsg() {
        return labelMsg;
    }
    public void setLabelMsg(String v) {
        labelMsg.setValue(v);
    }
    // upper: ActiveMatcher activeMatcher
    private MutableLiveData<String> activeMatcherMsg = new MutableLiveData<>();
    public MutableLiveData<String> getActiveMatcherMsg() {
        return activeMatcherMsg;
    }
    public void setActiveMatcherMsg(String v) {
        activeMatcherMsg.setValue(v);
    }
    // upper: ApplyType applyType
    private MutableLiveData<String> applyTypeMsg = new MutableLiveData<>();
    public MutableLiveData<String> getApplyTypeMsg() {
        return applyTypeMsg;
    }
    public void setApplyTypeMsg(String v) {
        applyTypeMsg.setValue(v);
    }
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