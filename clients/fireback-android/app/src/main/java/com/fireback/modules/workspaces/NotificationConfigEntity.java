package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class NotificationConfigEntity extends JsonSerializable {
    public Boolean cascadeToSubWorkspaces;
    public Boolean forcedCascadeEmailProvider;
    public EmailProviderEntity generalEmailProvider;
    public GsmProviderEntity generalGsmProvider;
    public String inviteToWorkspaceContent;
    public String inviteToWorkspaceContentExcerpt;
    public String inviteToWorkspaceContentDefault;
    public String inviteToWorkspaceContentDefaultExcerpt;
    public String inviteToWorkspaceTitle;
    public String inviteToWorkspaceTitleDefault;
    public EmailSenderEntity inviteToWorkspaceSender;
    public EmailSenderEntity accountCenterEmailSender;
    public String forgetPasswordContent;
    public String forgetPasswordContentExcerpt;
    public String forgetPasswordContentDefault;
    public String forgetPasswordContentDefaultExcerpt;
    public String forgetPasswordTitle;
    public String forgetPasswordTitleDefault;
    public EmailSenderEntity forgetPasswordSender;
    public String acceptLanguage;
    public EmailSenderEntity confirmEmailSender;
    public String confirmEmailContent;
    public String confirmEmailContentExcerpt;
    public String confirmEmailContentDefault;
    public String confirmEmailContentDefaultExcerpt;
    public String confirmEmailTitle;
    public String confirmEmailTitleDefault;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: CascadeToSubWorkspaces cascadeToSubWorkspaces
    private MutableLiveData< Boolean > cascadeToSubWorkspaces = new MutableLiveData<>();
    public MutableLiveData< Boolean > getCascadeToSubWorkspaces() {
        return cascadeToSubWorkspaces;
    }
    public void setCascadeToSubWorkspaces( Boolean  v) {
        cascadeToSubWorkspaces.setValue(v);
    }
    // upper: ForcedCascadeEmailProvider forcedCascadeEmailProvider
    private MutableLiveData< Boolean > forcedCascadeEmailProvider = new MutableLiveData<>();
    public MutableLiveData< Boolean > getForcedCascadeEmailProvider() {
        return forcedCascadeEmailProvider;
    }
    public void setForcedCascadeEmailProvider( Boolean  v) {
        forcedCascadeEmailProvider.setValue(v);
    }
    // upper: GeneralEmailProvider generalEmailProvider
    private MutableLiveData< EmailProviderEntity > generalEmailProvider = new MutableLiveData<>();
    public MutableLiveData< EmailProviderEntity > getGeneralEmailProvider() {
        return generalEmailProvider;
    }
    public void setGeneralEmailProvider( EmailProviderEntity  v) {
        generalEmailProvider.setValue(v);
    }
    // upper: GeneralGsmProvider generalGsmProvider
    private MutableLiveData< GsmProviderEntity > generalGsmProvider = new MutableLiveData<>();
    public MutableLiveData< GsmProviderEntity > getGeneralGsmProvider() {
        return generalGsmProvider;
    }
    public void setGeneralGsmProvider( GsmProviderEntity  v) {
        generalGsmProvider.setValue(v);
    }
    // upper: InviteToWorkspaceContent inviteToWorkspaceContent
    private MutableLiveData< String > inviteToWorkspaceContent = new MutableLiveData<>();
    public MutableLiveData< String > getInviteToWorkspaceContent() {
        return inviteToWorkspaceContent;
    }
    public void setInviteToWorkspaceContent( String  v) {
        inviteToWorkspaceContent.setValue(v);
    }
    // upper: InviteToWorkspaceContentExcerpt inviteToWorkspaceContentExcerpt
    private MutableLiveData< String > inviteToWorkspaceContentExcerpt = new MutableLiveData<>();
    public MutableLiveData< String > getInviteToWorkspaceContentExcerpt() {
        return inviteToWorkspaceContentExcerpt;
    }
    public void setInviteToWorkspaceContentExcerpt( String  v) {
        inviteToWorkspaceContentExcerpt.setValue(v);
    }
    // upper: InviteToWorkspaceContentDefault inviteToWorkspaceContentDefault
    private MutableLiveData< String > inviteToWorkspaceContentDefault = new MutableLiveData<>();
    public MutableLiveData< String > getInviteToWorkspaceContentDefault() {
        return inviteToWorkspaceContentDefault;
    }
    public void setInviteToWorkspaceContentDefault( String  v) {
        inviteToWorkspaceContentDefault.setValue(v);
    }
    // upper: InviteToWorkspaceContentDefaultExcerpt inviteToWorkspaceContentDefaultExcerpt
    private MutableLiveData< String > inviteToWorkspaceContentDefaultExcerpt = new MutableLiveData<>();
    public MutableLiveData< String > getInviteToWorkspaceContentDefaultExcerpt() {
        return inviteToWorkspaceContentDefaultExcerpt;
    }
    public void setInviteToWorkspaceContentDefaultExcerpt( String  v) {
        inviteToWorkspaceContentDefaultExcerpt.setValue(v);
    }
    // upper: InviteToWorkspaceTitle inviteToWorkspaceTitle
    private MutableLiveData< String > inviteToWorkspaceTitle = new MutableLiveData<>();
    public MutableLiveData< String > getInviteToWorkspaceTitle() {
        return inviteToWorkspaceTitle;
    }
    public void setInviteToWorkspaceTitle( String  v) {
        inviteToWorkspaceTitle.setValue(v);
    }
    // upper: InviteToWorkspaceTitleDefault inviteToWorkspaceTitleDefault
    private MutableLiveData< String > inviteToWorkspaceTitleDefault = new MutableLiveData<>();
    public MutableLiveData< String > getInviteToWorkspaceTitleDefault() {
        return inviteToWorkspaceTitleDefault;
    }
    public void setInviteToWorkspaceTitleDefault( String  v) {
        inviteToWorkspaceTitleDefault.setValue(v);
    }
    // upper: InviteToWorkspaceSender inviteToWorkspaceSender
    private MutableLiveData< EmailSenderEntity > inviteToWorkspaceSender = new MutableLiveData<>();
    public MutableLiveData< EmailSenderEntity > getInviteToWorkspaceSender() {
        return inviteToWorkspaceSender;
    }
    public void setInviteToWorkspaceSender( EmailSenderEntity  v) {
        inviteToWorkspaceSender.setValue(v);
    }
    // upper: AccountCenterEmailSender accountCenterEmailSender
    private MutableLiveData< EmailSenderEntity > accountCenterEmailSender = new MutableLiveData<>();
    public MutableLiveData< EmailSenderEntity > getAccountCenterEmailSender() {
        return accountCenterEmailSender;
    }
    public void setAccountCenterEmailSender( EmailSenderEntity  v) {
        accountCenterEmailSender.setValue(v);
    }
    // upper: ForgetPasswordContent forgetPasswordContent
    private MutableLiveData< String > forgetPasswordContent = new MutableLiveData<>();
    public MutableLiveData< String > getForgetPasswordContent() {
        return forgetPasswordContent;
    }
    public void setForgetPasswordContent( String  v) {
        forgetPasswordContent.setValue(v);
    }
    // upper: ForgetPasswordContentExcerpt forgetPasswordContentExcerpt
    private MutableLiveData< String > forgetPasswordContentExcerpt = new MutableLiveData<>();
    public MutableLiveData< String > getForgetPasswordContentExcerpt() {
        return forgetPasswordContentExcerpt;
    }
    public void setForgetPasswordContentExcerpt( String  v) {
        forgetPasswordContentExcerpt.setValue(v);
    }
    // upper: ForgetPasswordContentDefault forgetPasswordContentDefault
    private MutableLiveData< String > forgetPasswordContentDefault = new MutableLiveData<>();
    public MutableLiveData< String > getForgetPasswordContentDefault() {
        return forgetPasswordContentDefault;
    }
    public void setForgetPasswordContentDefault( String  v) {
        forgetPasswordContentDefault.setValue(v);
    }
    // upper: ForgetPasswordContentDefaultExcerpt forgetPasswordContentDefaultExcerpt
    private MutableLiveData< String > forgetPasswordContentDefaultExcerpt = new MutableLiveData<>();
    public MutableLiveData< String > getForgetPasswordContentDefaultExcerpt() {
        return forgetPasswordContentDefaultExcerpt;
    }
    public void setForgetPasswordContentDefaultExcerpt( String  v) {
        forgetPasswordContentDefaultExcerpt.setValue(v);
    }
    // upper: ForgetPasswordTitle forgetPasswordTitle
    private MutableLiveData< String > forgetPasswordTitle = new MutableLiveData<>();
    public MutableLiveData< String > getForgetPasswordTitle() {
        return forgetPasswordTitle;
    }
    public void setForgetPasswordTitle( String  v) {
        forgetPasswordTitle.setValue(v);
    }
    // upper: ForgetPasswordTitleDefault forgetPasswordTitleDefault
    private MutableLiveData< String > forgetPasswordTitleDefault = new MutableLiveData<>();
    public MutableLiveData< String > getForgetPasswordTitleDefault() {
        return forgetPasswordTitleDefault;
    }
    public void setForgetPasswordTitleDefault( String  v) {
        forgetPasswordTitleDefault.setValue(v);
    }
    // upper: ForgetPasswordSender forgetPasswordSender
    private MutableLiveData< EmailSenderEntity > forgetPasswordSender = new MutableLiveData<>();
    public MutableLiveData< EmailSenderEntity > getForgetPasswordSender() {
        return forgetPasswordSender;
    }
    public void setForgetPasswordSender( EmailSenderEntity  v) {
        forgetPasswordSender.setValue(v);
    }
    // upper: AcceptLanguage acceptLanguage
    private MutableLiveData< String > acceptLanguage = new MutableLiveData<>();
    public MutableLiveData< String > getAcceptLanguage() {
        return acceptLanguage;
    }
    public void setAcceptLanguage( String  v) {
        acceptLanguage.setValue(v);
    }
    // upper: ConfirmEmailSender confirmEmailSender
    private MutableLiveData< EmailSenderEntity > confirmEmailSender = new MutableLiveData<>();
    public MutableLiveData< EmailSenderEntity > getConfirmEmailSender() {
        return confirmEmailSender;
    }
    public void setConfirmEmailSender( EmailSenderEntity  v) {
        confirmEmailSender.setValue(v);
    }
    // upper: ConfirmEmailContent confirmEmailContent
    private MutableLiveData< String > confirmEmailContent = new MutableLiveData<>();
    public MutableLiveData< String > getConfirmEmailContent() {
        return confirmEmailContent;
    }
    public void setConfirmEmailContent( String  v) {
        confirmEmailContent.setValue(v);
    }
    // upper: ConfirmEmailContentExcerpt confirmEmailContentExcerpt
    private MutableLiveData< String > confirmEmailContentExcerpt = new MutableLiveData<>();
    public MutableLiveData< String > getConfirmEmailContentExcerpt() {
        return confirmEmailContentExcerpt;
    }
    public void setConfirmEmailContentExcerpt( String  v) {
        confirmEmailContentExcerpt.setValue(v);
    }
    // upper: ConfirmEmailContentDefault confirmEmailContentDefault
    private MutableLiveData< String > confirmEmailContentDefault = new MutableLiveData<>();
    public MutableLiveData< String > getConfirmEmailContentDefault() {
        return confirmEmailContentDefault;
    }
    public void setConfirmEmailContentDefault( String  v) {
        confirmEmailContentDefault.setValue(v);
    }
    // upper: ConfirmEmailContentDefaultExcerpt confirmEmailContentDefaultExcerpt
    private MutableLiveData< String > confirmEmailContentDefaultExcerpt = new MutableLiveData<>();
    public MutableLiveData< String > getConfirmEmailContentDefaultExcerpt() {
        return confirmEmailContentDefaultExcerpt;
    }
    public void setConfirmEmailContentDefaultExcerpt( String  v) {
        confirmEmailContentDefaultExcerpt.setValue(v);
    }
    // upper: ConfirmEmailTitle confirmEmailTitle
    private MutableLiveData< String > confirmEmailTitle = new MutableLiveData<>();
    public MutableLiveData< String > getConfirmEmailTitle() {
        return confirmEmailTitle;
    }
    public void setConfirmEmailTitle( String  v) {
        confirmEmailTitle.setValue(v);
    }
    // upper: ConfirmEmailTitleDefault confirmEmailTitleDefault
    private MutableLiveData< String > confirmEmailTitleDefault = new MutableLiveData<>();
    public MutableLiveData< String > getConfirmEmailTitleDefault() {
        return confirmEmailTitleDefault;
    }
    public void setConfirmEmailTitleDefault( String  v) {
        confirmEmailTitleDefault.setValue(v);
    }
    // Handling error message for each field
    // upper: CascadeToSubWorkspaces cascadeToSubWorkspaces
    private MutableLiveData<String> cascadeToSubWorkspacesMsg = new MutableLiveData<>();
    public MutableLiveData<String> getCascadeToSubWorkspacesMsg() {
        return cascadeToSubWorkspacesMsg;
    }
    public void setCascadeToSubWorkspacesMsg(String v) {
        cascadeToSubWorkspacesMsg.setValue(v);
    }
    // upper: ForcedCascadeEmailProvider forcedCascadeEmailProvider
    private MutableLiveData<String> forcedCascadeEmailProviderMsg = new MutableLiveData<>();
    public MutableLiveData<String> getForcedCascadeEmailProviderMsg() {
        return forcedCascadeEmailProviderMsg;
    }
    public void setForcedCascadeEmailProviderMsg(String v) {
        forcedCascadeEmailProviderMsg.setValue(v);
    }
    // upper: GeneralEmailProvider generalEmailProvider
    private MutableLiveData<String> generalEmailProviderMsg = new MutableLiveData<>();
    public MutableLiveData<String> getGeneralEmailProviderMsg() {
        return generalEmailProviderMsg;
    }
    public void setGeneralEmailProviderMsg(String v) {
        generalEmailProviderMsg.setValue(v);
    }
    // upper: GeneralGsmProvider generalGsmProvider
    private MutableLiveData<String> generalGsmProviderMsg = new MutableLiveData<>();
    public MutableLiveData<String> getGeneralGsmProviderMsg() {
        return generalGsmProviderMsg;
    }
    public void setGeneralGsmProviderMsg(String v) {
        generalGsmProviderMsg.setValue(v);
    }
    // upper: InviteToWorkspaceContent inviteToWorkspaceContent
    private MutableLiveData<String> inviteToWorkspaceContentMsg = new MutableLiveData<>();
    public MutableLiveData<String> getInviteToWorkspaceContentMsg() {
        return inviteToWorkspaceContentMsg;
    }
    public void setInviteToWorkspaceContentMsg(String v) {
        inviteToWorkspaceContentMsg.setValue(v);
    }
    // upper: InviteToWorkspaceContentExcerpt inviteToWorkspaceContentExcerpt
    private MutableLiveData<String> inviteToWorkspaceContentExcerptMsg = new MutableLiveData<>();
    public MutableLiveData<String> getInviteToWorkspaceContentExcerptMsg() {
        return inviteToWorkspaceContentExcerptMsg;
    }
    public void setInviteToWorkspaceContentExcerptMsg(String v) {
        inviteToWorkspaceContentExcerptMsg.setValue(v);
    }
    // upper: InviteToWorkspaceContentDefault inviteToWorkspaceContentDefault
    private MutableLiveData<String> inviteToWorkspaceContentDefaultMsg = new MutableLiveData<>();
    public MutableLiveData<String> getInviteToWorkspaceContentDefaultMsg() {
        return inviteToWorkspaceContentDefaultMsg;
    }
    public void setInviteToWorkspaceContentDefaultMsg(String v) {
        inviteToWorkspaceContentDefaultMsg.setValue(v);
    }
    // upper: InviteToWorkspaceContentDefaultExcerpt inviteToWorkspaceContentDefaultExcerpt
    private MutableLiveData<String> inviteToWorkspaceContentDefaultExcerptMsg = new MutableLiveData<>();
    public MutableLiveData<String> getInviteToWorkspaceContentDefaultExcerptMsg() {
        return inviteToWorkspaceContentDefaultExcerptMsg;
    }
    public void setInviteToWorkspaceContentDefaultExcerptMsg(String v) {
        inviteToWorkspaceContentDefaultExcerptMsg.setValue(v);
    }
    // upper: InviteToWorkspaceTitle inviteToWorkspaceTitle
    private MutableLiveData<String> inviteToWorkspaceTitleMsg = new MutableLiveData<>();
    public MutableLiveData<String> getInviteToWorkspaceTitleMsg() {
        return inviteToWorkspaceTitleMsg;
    }
    public void setInviteToWorkspaceTitleMsg(String v) {
        inviteToWorkspaceTitleMsg.setValue(v);
    }
    // upper: InviteToWorkspaceTitleDefault inviteToWorkspaceTitleDefault
    private MutableLiveData<String> inviteToWorkspaceTitleDefaultMsg = new MutableLiveData<>();
    public MutableLiveData<String> getInviteToWorkspaceTitleDefaultMsg() {
        return inviteToWorkspaceTitleDefaultMsg;
    }
    public void setInviteToWorkspaceTitleDefaultMsg(String v) {
        inviteToWorkspaceTitleDefaultMsg.setValue(v);
    }
    // upper: InviteToWorkspaceSender inviteToWorkspaceSender
    private MutableLiveData<String> inviteToWorkspaceSenderMsg = new MutableLiveData<>();
    public MutableLiveData<String> getInviteToWorkspaceSenderMsg() {
        return inviteToWorkspaceSenderMsg;
    }
    public void setInviteToWorkspaceSenderMsg(String v) {
        inviteToWorkspaceSenderMsg.setValue(v);
    }
    // upper: AccountCenterEmailSender accountCenterEmailSender
    private MutableLiveData<String> accountCenterEmailSenderMsg = new MutableLiveData<>();
    public MutableLiveData<String> getAccountCenterEmailSenderMsg() {
        return accountCenterEmailSenderMsg;
    }
    public void setAccountCenterEmailSenderMsg(String v) {
        accountCenterEmailSenderMsg.setValue(v);
    }
    // upper: ForgetPasswordContent forgetPasswordContent
    private MutableLiveData<String> forgetPasswordContentMsg = new MutableLiveData<>();
    public MutableLiveData<String> getForgetPasswordContentMsg() {
        return forgetPasswordContentMsg;
    }
    public void setForgetPasswordContentMsg(String v) {
        forgetPasswordContentMsg.setValue(v);
    }
    // upper: ForgetPasswordContentExcerpt forgetPasswordContentExcerpt
    private MutableLiveData<String> forgetPasswordContentExcerptMsg = new MutableLiveData<>();
    public MutableLiveData<String> getForgetPasswordContentExcerptMsg() {
        return forgetPasswordContentExcerptMsg;
    }
    public void setForgetPasswordContentExcerptMsg(String v) {
        forgetPasswordContentExcerptMsg.setValue(v);
    }
    // upper: ForgetPasswordContentDefault forgetPasswordContentDefault
    private MutableLiveData<String> forgetPasswordContentDefaultMsg = new MutableLiveData<>();
    public MutableLiveData<String> getForgetPasswordContentDefaultMsg() {
        return forgetPasswordContentDefaultMsg;
    }
    public void setForgetPasswordContentDefaultMsg(String v) {
        forgetPasswordContentDefaultMsg.setValue(v);
    }
    // upper: ForgetPasswordContentDefaultExcerpt forgetPasswordContentDefaultExcerpt
    private MutableLiveData<String> forgetPasswordContentDefaultExcerptMsg = new MutableLiveData<>();
    public MutableLiveData<String> getForgetPasswordContentDefaultExcerptMsg() {
        return forgetPasswordContentDefaultExcerptMsg;
    }
    public void setForgetPasswordContentDefaultExcerptMsg(String v) {
        forgetPasswordContentDefaultExcerptMsg.setValue(v);
    }
    // upper: ForgetPasswordTitle forgetPasswordTitle
    private MutableLiveData<String> forgetPasswordTitleMsg = new MutableLiveData<>();
    public MutableLiveData<String> getForgetPasswordTitleMsg() {
        return forgetPasswordTitleMsg;
    }
    public void setForgetPasswordTitleMsg(String v) {
        forgetPasswordTitleMsg.setValue(v);
    }
    // upper: ForgetPasswordTitleDefault forgetPasswordTitleDefault
    private MutableLiveData<String> forgetPasswordTitleDefaultMsg = new MutableLiveData<>();
    public MutableLiveData<String> getForgetPasswordTitleDefaultMsg() {
        return forgetPasswordTitleDefaultMsg;
    }
    public void setForgetPasswordTitleDefaultMsg(String v) {
        forgetPasswordTitleDefaultMsg.setValue(v);
    }
    // upper: ForgetPasswordSender forgetPasswordSender
    private MutableLiveData<String> forgetPasswordSenderMsg = new MutableLiveData<>();
    public MutableLiveData<String> getForgetPasswordSenderMsg() {
        return forgetPasswordSenderMsg;
    }
    public void setForgetPasswordSenderMsg(String v) {
        forgetPasswordSenderMsg.setValue(v);
    }
    // upper: AcceptLanguage acceptLanguage
    private MutableLiveData<String> acceptLanguageMsg = new MutableLiveData<>();
    public MutableLiveData<String> getAcceptLanguageMsg() {
        return acceptLanguageMsg;
    }
    public void setAcceptLanguageMsg(String v) {
        acceptLanguageMsg.setValue(v);
    }
    // upper: ConfirmEmailSender confirmEmailSender
    private MutableLiveData<String> confirmEmailSenderMsg = new MutableLiveData<>();
    public MutableLiveData<String> getConfirmEmailSenderMsg() {
        return confirmEmailSenderMsg;
    }
    public void setConfirmEmailSenderMsg(String v) {
        confirmEmailSenderMsg.setValue(v);
    }
    // upper: ConfirmEmailContent confirmEmailContent
    private MutableLiveData<String> confirmEmailContentMsg = new MutableLiveData<>();
    public MutableLiveData<String> getConfirmEmailContentMsg() {
        return confirmEmailContentMsg;
    }
    public void setConfirmEmailContentMsg(String v) {
        confirmEmailContentMsg.setValue(v);
    }
    // upper: ConfirmEmailContentExcerpt confirmEmailContentExcerpt
    private MutableLiveData<String> confirmEmailContentExcerptMsg = new MutableLiveData<>();
    public MutableLiveData<String> getConfirmEmailContentExcerptMsg() {
        return confirmEmailContentExcerptMsg;
    }
    public void setConfirmEmailContentExcerptMsg(String v) {
        confirmEmailContentExcerptMsg.setValue(v);
    }
    // upper: ConfirmEmailContentDefault confirmEmailContentDefault
    private MutableLiveData<String> confirmEmailContentDefaultMsg = new MutableLiveData<>();
    public MutableLiveData<String> getConfirmEmailContentDefaultMsg() {
        return confirmEmailContentDefaultMsg;
    }
    public void setConfirmEmailContentDefaultMsg(String v) {
        confirmEmailContentDefaultMsg.setValue(v);
    }
    // upper: ConfirmEmailContentDefaultExcerpt confirmEmailContentDefaultExcerpt
    private MutableLiveData<String> confirmEmailContentDefaultExcerptMsg = new MutableLiveData<>();
    public MutableLiveData<String> getConfirmEmailContentDefaultExcerptMsg() {
        return confirmEmailContentDefaultExcerptMsg;
    }
    public void setConfirmEmailContentDefaultExcerptMsg(String v) {
        confirmEmailContentDefaultExcerptMsg.setValue(v);
    }
    // upper: ConfirmEmailTitle confirmEmailTitle
    private MutableLiveData<String> confirmEmailTitleMsg = new MutableLiveData<>();
    public MutableLiveData<String> getConfirmEmailTitleMsg() {
        return confirmEmailTitleMsg;
    }
    public void setConfirmEmailTitleMsg(String v) {
        confirmEmailTitleMsg.setValue(v);
    }
    // upper: ConfirmEmailTitleDefault confirmEmailTitleDefault
    private MutableLiveData<String> confirmEmailTitleDefaultMsg = new MutableLiveData<>();
    public MutableLiveData<String> getConfirmEmailTitleDefaultMsg() {
        return confirmEmailTitleDefaultMsg;
    }
    public void setConfirmEmailTitleDefaultMsg(String v) {
        confirmEmailTitleDefaultMsg.setValue(v);
    }
  }
}