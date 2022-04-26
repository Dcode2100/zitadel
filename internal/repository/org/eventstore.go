package org

import (
	"github.com/zitadel/zitadel/internal/eventstore"
)

func RegisterEventMappers(es *eventstore.Eventstore) {
	es.RegisterFilterEventMapper(OrgAddedEventType, OrgAddedEventMapper).
		RegisterFilterEventMapper(OrgChangedEventType, OrgChangedEventMapper).
		RegisterFilterEventMapper(OrgDeactivatedEventType, OrgDeactivatedEventMapper).
		RegisterFilterEventMapper(OrgReactivatedEventType, OrgReactivatedEventMapper).
		RegisterFilterEventMapper(OrgDomainAddedEventType, DomainAddedEventMapper).
		RegisterFilterEventMapper(OrgDomainVerificationAddedEventType, DomainVerificationAddedEventMapper).
		RegisterFilterEventMapper(OrgDomainVerificationFailedEventType, DomainVerificationFailedEventMapper).
		RegisterFilterEventMapper(OrgDomainVerifiedEventType, DomainVerifiedEventMapper).
		RegisterFilterEventMapper(OrgDomainPrimarySetEventType, DomainPrimarySetEventMapper).
		RegisterFilterEventMapper(OrgDomainRemovedEventType, DomainRemovedEventMapper).
		RegisterFilterEventMapper(MemberAddedEventType, MemberAddedEventMapper).
		RegisterFilterEventMapper(MemberChangedEventType, MemberChangedEventMapper).
		RegisterFilterEventMapper(MemberRemovedEventType, MemberRemovedEventMapper).
		RegisterFilterEventMapper(MemberCascadeRemovedEventType, MemberCascadeRemovedEventMapper).
		RegisterFilterEventMapper(LabelPolicyAddedEventType, LabelPolicyAddedEventMapper).
		RegisterFilterEventMapper(LabelPolicyChangedEventType, LabelPolicyChangedEventMapper).
		RegisterFilterEventMapper(LabelPolicyActivatedEventType, LabelPolicyActivatedEventMapper).
		RegisterFilterEventMapper(LabelPolicyRemovedEventType, LabelPolicyRemovedEventMapper).
		RegisterFilterEventMapper(LabelPolicyLogoAddedEventType, LabelPolicyLogoAddedEventMapper).
		RegisterFilterEventMapper(LabelPolicyLogoRemovedEventType, LabelPolicyLogoRemovedEventMapper).
		RegisterFilterEventMapper(LabelPolicyIconAddedEventType, LabelPolicyIconAddedEventMapper).
		RegisterFilterEventMapper(LabelPolicyIconRemovedEventType, LabelPolicyIconRemovedEventMapper).
		RegisterFilterEventMapper(LabelPolicyLogoDarkAddedEventType, LabelPolicyLogoDarkAddedEventMapper).
		RegisterFilterEventMapper(LabelPolicyLogoDarkRemovedEventType, LabelPolicyLogoDarkRemovedEventMapper).
		RegisterFilterEventMapper(LabelPolicyIconDarkAddedEventType, LabelPolicyIconDarkAddedEventMapper).
		RegisterFilterEventMapper(LabelPolicyIconDarkRemovedEventType, LabelPolicyIconDarkRemovedEventMapper).
		RegisterFilterEventMapper(LabelPolicyFontAddedEventType, LabelPolicyFontAddedEventMapper).
		RegisterFilterEventMapper(LabelPolicyFontRemovedEventType, LabelPolicyFontRemovedEventMapper).
		RegisterFilterEventMapper(LabelPolicyAssetsRemovedEventType, LabelPolicyAssetsRemovedEventMapper).
		RegisterFilterEventMapper(LoginPolicyAddedEventType, LoginPolicyAddedEventMapper).
		RegisterFilterEventMapper(LoginPolicyChangedEventType, LoginPolicyChangedEventMapper).
		RegisterFilterEventMapper(LoginPolicyRemovedEventType, LoginPolicyRemovedEventMapper).
		RegisterFilterEventMapper(LoginPolicySecondFactorAddedEventType, SecondFactorAddedEventEventMapper).
		RegisterFilterEventMapper(LoginPolicySecondFactorRemovedEventType, SecondFactorRemovedEventEventMapper).
		RegisterFilterEventMapper(LoginPolicyMultiFactorAddedEventType, MultiFactorAddedEventEventMapper).
		RegisterFilterEventMapper(LoginPolicyMultiFactorRemovedEventType, MultiFactorRemovedEventEventMapper).
		RegisterFilterEventMapper(LoginPolicyIDPProviderAddedEventType, IdentityProviderAddedEventMapper).
		RegisterFilterEventMapper(LoginPolicyIDPProviderRemovedEventType, IdentityProviderRemovedEventMapper).
		RegisterFilterEventMapper(LoginPolicyIDPProviderCascadeRemovedEventType, IdentityProviderCascadeRemovedEventMapper).
		RegisterFilterEventMapper(DomainPolicyAddedEventType, DomainPolicyAddedEventMapper).
		RegisterFilterEventMapper(DomainPolicyChangedEventType, DomainPolicyChangedEventMapper).
		RegisterFilterEventMapper(DomainPolicyRemovedEventType, DomainPolicyRemovedEventMapper).
		RegisterFilterEventMapper(PasswordAgePolicyAddedEventType, PasswordAgePolicyAddedEventMapper).
		RegisterFilterEventMapper(PasswordAgePolicyChangedEventType, PasswordAgePolicyChangedEventMapper).
		RegisterFilterEventMapper(PasswordAgePolicyRemovedEventType, PasswordAgePolicyRemovedEventMapper).
		RegisterFilterEventMapper(PasswordComplexityPolicyAddedEventType, PasswordComplexityPolicyAddedEventMapper).
		RegisterFilterEventMapper(PasswordComplexityPolicyChangedEventType, PasswordComplexityPolicyChangedEventMapper).
		RegisterFilterEventMapper(PasswordComplexityPolicyRemovedEventType, PasswordComplexityPolicyRemovedEventMapper).
		RegisterFilterEventMapper(LockoutPolicyAddedEventType, LockoutPolicyAddedEventMapper).
		RegisterFilterEventMapper(LockoutPolicyChangedEventType, LockoutPolicyChangedEventMapper).
		RegisterFilterEventMapper(LockoutPolicyRemovedEventType, LockoutPolicyRemovedEventMapper).
		RegisterFilterEventMapper(PrivacyPolicyAddedEventType, PrivacyPolicyAddedEventMapper).
		RegisterFilterEventMapper(PrivacyPolicyChangedEventType, PrivacyPolicyChangedEventMapper).
		RegisterFilterEventMapper(PrivacyPolicyRemovedEventType, PrivacyPolicyRemovedEventMapper).
		RegisterFilterEventMapper(MailTemplateAddedEventType, MailTemplateAddedEventMapper).
		RegisterFilterEventMapper(MailTemplateChangedEventType, MailTemplateChangedEventMapper).
		RegisterFilterEventMapper(MailTemplateRemovedEventType, MailTemplateRemovedEventMapper).
		RegisterFilterEventMapper(MailTextAddedEventType, MailTextAddedEventMapper).
		RegisterFilterEventMapper(MailTextChangedEventType, MailTextChangedEventMapper).
		RegisterFilterEventMapper(MailTextRemovedEventType, MailTextRemovedEventMapper).
		RegisterFilterEventMapper(CustomTextSetEventType, CustomTextSetEventMapper).
		RegisterFilterEventMapper(CustomTextRemovedEventType, CustomTextRemovedEventMapper).
		RegisterFilterEventMapper(CustomTextTemplateRemovedEventType, CustomTextTemplateRemovedEventMapper).
		RegisterFilterEventMapper(IDPConfigAddedEventType, IDPConfigAddedEventMapper).
		RegisterFilterEventMapper(IDPConfigChangedEventType, IDPConfigChangedEventMapper).
		RegisterFilterEventMapper(IDPConfigRemovedEventType, IDPConfigRemovedEventMapper).
		RegisterFilterEventMapper(IDPConfigDeactivatedEventType, IDPConfigDeactivatedEventMapper).
		RegisterFilterEventMapper(IDPConfigReactivatedEventType, IDPConfigReactivatedEventMapper).
		RegisterFilterEventMapper(IDPOIDCConfigAddedEventType, IDPOIDCConfigAddedEventMapper).
		RegisterFilterEventMapper(IDPOIDCConfigChangedEventType, IDPOIDCConfigChangedEventMapper).
		RegisterFilterEventMapper(IDPJWTConfigAddedEventType, IDPJWTConfigAddedEventMapper).
		RegisterFilterEventMapper(IDPJWTConfigChangedEventType, IDPJWTConfigChangedEventMapper).
		RegisterFilterEventMapper(FeaturesSetEventType, FeaturesSetEventMapper).
		RegisterFilterEventMapper(FeaturesRemovedEventType, FeaturesRemovedEventMapper).
		RegisterFilterEventMapper(TriggerActionsSetEventType, TriggerActionsSetEventMapper).
		RegisterFilterEventMapper(TriggerActionsCascadeRemovedEventType, TriggerActionsCascadeRemovedEventMapper).
		RegisterFilterEventMapper(FlowClearedEventType, FlowClearedEventMapper)
}
