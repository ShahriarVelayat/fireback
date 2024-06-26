/*
 *	Generated by fireback 1.1.16
 *	Written by Ali Torabi.
 *	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
package com.fireback.modules.workspaces;

import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.JsonSerializable;

public class FileEntity extends JsonSerializable {
  public String name;
  public String diskPath;
  public int size;
  public String virtualPath;
  public String type;

  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Name name
    private MutableLiveData<String> name = new MutableLiveData<>();

    public MutableLiveData<String> getName() {
      return name;
    }

    public void setName(String v) {
      name.setValue(v);
    }

    // upper: DiskPath diskPath
    private MutableLiveData<String> diskPath = new MutableLiveData<>();

    public MutableLiveData<String> getDiskPath() {
      return diskPath;
    }

    public void setDiskPath(String v) {
      diskPath.setValue(v);
    }

    // upper: Size size
    private MutableLiveData<Integer> size = new MutableLiveData<>();

    public MutableLiveData<Integer> getSize() {
      return size;
    }

    public void setSize(Integer v) {
      size.setValue(v);
    }

    // upper: VirtualPath virtualPath
    private MutableLiveData<String> virtualPath = new MutableLiveData<>();

    public MutableLiveData<String> getVirtualPath() {
      return virtualPath;
    }

    public void setVirtualPath(String v) {
      virtualPath.setValue(v);
    }

    // upper: Type type
    private MutableLiveData<String> type = new MutableLiveData<>();

    public MutableLiveData<String> getType() {
      return type;
    }

    public void setType(String v) {
      type.setValue(v);
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

    // upper: DiskPath diskPath
    private MutableLiveData<String> diskPathMsg = new MutableLiveData<>();

    public MutableLiveData<String> getDiskPathMsg() {
      return diskPathMsg;
    }

    public void setDiskPathMsg(String v) {
      diskPathMsg.setValue(v);
    }

    // upper: Size size
    private MutableLiveData<String> sizeMsg = new MutableLiveData<>();

    public MutableLiveData<String> getSizeMsg() {
      return sizeMsg;
    }

    public void setSizeMsg(String v) {
      sizeMsg.setValue(v);
    }

    // upper: VirtualPath virtualPath
    private MutableLiveData<String> virtualPathMsg = new MutableLiveData<>();

    public MutableLiveData<String> getVirtualPathMsg() {
      return virtualPathMsg;
    }

    public void setVirtualPathMsg(String v) {
      virtualPathMsg.setValue(v);
    }

    // upper: Type type
    private MutableLiveData<String> typeMsg = new MutableLiveData<>();

    public MutableLiveData<String> getTypeMsg() {
      return typeMsg;
    }

    public void setTypeMsg(String v) {
      typeMsg.setValue(v);
    }
  }
}
