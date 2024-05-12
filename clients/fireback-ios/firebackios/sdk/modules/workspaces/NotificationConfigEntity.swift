import Foundation
class NotificationConfigEntity : Codable, Identifiable {
    var cascadeToSubWorkspaces: Bool? = nil
    var forcedCascadeEmailProvider: Bool? = nil
    var generalEmailProvider: EmailProviderEntity? = nil
    // var generalEmailProviderId: String? = nil
    var generalGsmProvider: GsmProviderEntity? = nil
    // var generalGsmProviderId: String? = nil
    var inviteToWorkspaceContent: String? = nil
    var inviteToWorkspaceContentExcerpt: String? = nil
    var inviteToWorkspaceContentDefault: String? = nil
    var inviteToWorkspaceContentDefaultExcerpt: String? = nil
    var inviteToWorkspaceTitle: String? = nil
    var inviteToWorkspaceTitleDefault: String? = nil
    var inviteToWorkspaceSender: EmailSenderEntity? = nil
    // var inviteToWorkspaceSenderId: String? = nil
    var accountCenterEmailSender: EmailSenderEntity? = nil
    // var accountCenterEmailSenderId: String? = nil
    var forgetPasswordContent: String? = nil
    var forgetPasswordContentExcerpt: String? = nil
    var forgetPasswordContentDefault: String? = nil
    var forgetPasswordContentDefaultExcerpt: String? = nil
    var forgetPasswordTitle: String? = nil
    var forgetPasswordTitleDefault: String? = nil
    var forgetPasswordSender: EmailSenderEntity? = nil
    // var forgetPasswordSenderId: String? = nil
    var acceptLanguage: String? = nil
    var confirmEmailSender: EmailSenderEntity? = nil
    // var confirmEmailSenderId: String? = nil
    var confirmEmailContent: String? = nil
    var confirmEmailContentExcerpt: String? = nil
    var confirmEmailContentDefault: String? = nil
    var confirmEmailContentDefaultExcerpt: String? = nil
    var confirmEmailTitle: String? = nil
    var confirmEmailTitleDefault: String? = nil
}
class NotificationConfigEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var cascadeToSubWorkspaces: Bool? = nil
  @Published var cascadeToSubWorkspacesErrorMessage: Bool? = nil
  @Published var forcedCascadeEmailProvider: Bool? = nil
  @Published var forcedCascadeEmailProviderErrorMessage: Bool? = nil
  @Published var inviteToWorkspaceContent: String? = nil
  @Published var inviteToWorkspaceContentErrorMessage: String? = nil
  @Published var inviteToWorkspaceContentExcerpt: String? = nil
  @Published var inviteToWorkspaceContentExcerptErrorMessage: String? = nil
  @Published var inviteToWorkspaceContentDefault: String? = nil
  @Published var inviteToWorkspaceContentDefaultErrorMessage: String? = nil
  @Published var inviteToWorkspaceContentDefaultExcerpt: String? = nil
  @Published var inviteToWorkspaceContentDefaultExcerptErrorMessage: String? = nil
  @Published var inviteToWorkspaceTitle: String? = nil
  @Published var inviteToWorkspaceTitleErrorMessage: String? = nil
  @Published var inviteToWorkspaceTitleDefault: String? = nil
  @Published var inviteToWorkspaceTitleDefaultErrorMessage: String? = nil
  @Published var forgetPasswordContent: String? = nil
  @Published var forgetPasswordContentErrorMessage: String? = nil
  @Published var forgetPasswordContentExcerpt: String? = nil
  @Published var forgetPasswordContentExcerptErrorMessage: String? = nil
  @Published var forgetPasswordContentDefault: String? = nil
  @Published var forgetPasswordContentDefaultErrorMessage: String? = nil
  @Published var forgetPasswordContentDefaultExcerpt: String? = nil
  @Published var forgetPasswordContentDefaultExcerptErrorMessage: String? = nil
  @Published var forgetPasswordTitle: String? = nil
  @Published var forgetPasswordTitleErrorMessage: String? = nil
  @Published var forgetPasswordTitleDefault: String? = nil
  @Published var forgetPasswordTitleDefaultErrorMessage: String? = nil
  @Published var confirmEmailContent: String? = nil
  @Published var confirmEmailContentErrorMessage: String? = nil
  @Published var confirmEmailContentExcerpt: String? = nil
  @Published var confirmEmailContentExcerptErrorMessage: String? = nil
  @Published var confirmEmailContentDefault: String? = nil
  @Published var confirmEmailContentDefaultErrorMessage: String? = nil
  @Published var confirmEmailContentDefaultExcerpt: String? = nil
  @Published var confirmEmailContentDefaultExcerptErrorMessage: String? = nil
  @Published var confirmEmailTitle: String? = nil
  @Published var confirmEmailTitleErrorMessage: String? = nil
  @Published var confirmEmailTitleDefault: String? = nil
  @Published var confirmEmailTitleDefaultErrorMessage: String? = nil
  func getDto() -> NotificationConfigEntity {
      var dto = NotificationConfigEntity()
    dto.cascadeToSubWorkspaces = self.cascadeToSubWorkspaces
    dto.forcedCascadeEmailProvider = self.forcedCascadeEmailProvider
    dto.inviteToWorkspaceContent = self.inviteToWorkspaceContent
    dto.inviteToWorkspaceContentExcerpt = self.inviteToWorkspaceContentExcerpt
    dto.inviteToWorkspaceContentDefault = self.inviteToWorkspaceContentDefault
    dto.inviteToWorkspaceContentDefaultExcerpt = self.inviteToWorkspaceContentDefaultExcerpt
    dto.inviteToWorkspaceTitle = self.inviteToWorkspaceTitle
    dto.inviteToWorkspaceTitleDefault = self.inviteToWorkspaceTitleDefault
    dto.forgetPasswordContent = self.forgetPasswordContent
    dto.forgetPasswordContentExcerpt = self.forgetPasswordContentExcerpt
    dto.forgetPasswordContentDefault = self.forgetPasswordContentDefault
    dto.forgetPasswordContentDefaultExcerpt = self.forgetPasswordContentDefaultExcerpt
    dto.forgetPasswordTitle = self.forgetPasswordTitle
    dto.forgetPasswordTitleDefault = self.forgetPasswordTitleDefault
    dto.confirmEmailContent = self.confirmEmailContent
    dto.confirmEmailContentExcerpt = self.confirmEmailContentExcerpt
    dto.confirmEmailContentDefault = self.confirmEmailContentDefault
    dto.confirmEmailContentDefaultExcerpt = self.confirmEmailContentDefaultExcerpt
    dto.confirmEmailTitle = self.confirmEmailTitle
    dto.confirmEmailTitleDefault = self.confirmEmailTitleDefault
      return dto
  }
}