/*
 *	Generated by fireback 1.1.16
 *	Written by Ali Torabi.
 *	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
package com.fireback.modules.currency;

import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.JsonSerializable;
import com.fireback.modules.workspaces.*;

class PriceTagVariations extends JsonSerializable {
  public CurrencyEntity currency;
  public float amount;

  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Currency currency
    private MutableLiveData<CurrencyEntity> currency = new MutableLiveData<>();

    public MutableLiveData<CurrencyEntity> getCurrency() {
      return currency;
    }

    public void setCurrency(CurrencyEntity v) {
      currency.setValue(v);
    }

    // upper: Amount amount
    private MutableLiveData<Float> amount = new MutableLiveData<>();

    public MutableLiveData<Float> getAmount() {
      return amount;
    }

    public void setAmount(Float v) {
      amount.setValue(v);
    }

    // Handling error message for each field
    // upper: Currency currency
    private MutableLiveData<String> currencyMsg = new MutableLiveData<>();

    public MutableLiveData<String> getCurrencyMsg() {
      return currencyMsg;
    }

    public void setCurrencyMsg(String v) {
      currencyMsg.setValue(v);
    }

    // upper: Amount amount
    private MutableLiveData<String> amountMsg = new MutableLiveData<>();

    public MutableLiveData<String> getAmountMsg() {
      return amountMsg;
    }

    public void setAmountMsg(String v) {
      amountMsg.setValue(v);
    }
  }
}

public class PriceTagEntity extends JsonSerializable {
  public PriceTagVariations[] variations;

  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Variations variations
    private MutableLiveData<PriceTagVariations[]> variations = new MutableLiveData<>();

    public MutableLiveData<PriceTagVariations[]> getVariations() {
      return variations;
    }

    public void setVariations(PriceTagVariations[] v) {
      variations.setValue(v);
    }

    // Handling error message for each field
    // upper: Variations variations
    private MutableLiveData<String> variationsMsg = new MutableLiveData<>();

    public MutableLiveData<String> getVariationsMsg() {
      return variationsMsg;
    }

    public void setVariationsMsg(String v) {
      variationsMsg.setValue(v);
    }
  }
}
