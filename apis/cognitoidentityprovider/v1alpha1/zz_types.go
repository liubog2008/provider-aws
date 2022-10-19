/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

// +kubebuilder:skipversion
type AccountRecoverySettingType struct {
	RecoveryMechanisms []*RecoveryOptionType `json:"recoveryMechanisms,omitempty"`
}

// +kubebuilder:skipversion
type AdminCreateUserConfigType struct {
	AllowAdminCreateUserOnly *bool `json:"allowAdminCreateUserOnly,omitempty"`
	// The message template structure.
	InviteMessageTemplate *MessageTemplateType `json:"inviteMessageTemplate,omitempty"`
}

// +kubebuilder:skipversion
type AnalyticsConfigurationType struct {
	ApplicationARN *string `json:"applicationARN,omitempty"`

	ApplicationID *string `json:"applicationID,omitempty"`

	ExternalID *string `json:"externalID,omitempty"`

	RoleARN *string `json:"roleARN,omitempty"`

	UserDataShared *bool `json:"userDataShared,omitempty"`
}

// +kubebuilder:skipversion
type AnalyticsMetadataType struct {
	AnalyticsEndpointID *string `json:"analyticsEndpointID,omitempty"`
}

// +kubebuilder:skipversion
type AuthEventType struct {
	CreationDate *metav1.Time `json:"creationDate,omitempty"`

	EventID *string `json:"eventID,omitempty"`
}

// +kubebuilder:skipversion
type AuthenticationResultType struct {
	ExpiresIn *int64 `json:"expiresIn,omitempty"`

	TokenType *string `json:"tokenType,omitempty"`
}

// +kubebuilder:skipversion
type CodeDeliveryDetailsType struct {
	Destination *string `json:"destination,omitempty"`
}

// +kubebuilder:skipversion
type ContextDataType struct {
	EncodedData *string `json:"encodedData,omitempty"`

	IPAddress *string `json:"ipAddress,omitempty"`

	ServerName *string `json:"serverName,omitempty"`

	ServerPath *string `json:"serverPath,omitempty"`
}

// +kubebuilder:skipversion
type CustomDomainConfigType struct {
	CertificateARN *string `json:"certificateARN,omitempty"`
}

// +kubebuilder:skipversion
type CustomEmailLambdaVersionConfigType struct {
	LambdaARN *string `json:"lambdaARN,omitempty"`

	LambdaVersion *string `json:"lambdaVersion,omitempty"`
}

// +kubebuilder:skipversion
type CustomSMSLambdaVersionConfigType struct {
	LambdaARN *string `json:"lambdaARN,omitempty"`

	LambdaVersion *string `json:"lambdaVersion,omitempty"`
}

// +kubebuilder:skipversion
type DeviceConfigurationType struct {
	ChallengeRequiredOnNewDevice *bool `json:"challengeRequiredOnNewDevice,omitempty"`

	DeviceOnlyRememberedOnUserPrompt *bool `json:"deviceOnlyRememberedOnUserPrompt,omitempty"`
}

// +kubebuilder:skipversion
type DeviceSecretVerifierConfigType struct {
	PasswordVerifier *string `json:"passwordVerifier,omitempty"`

	Salt *string `json:"salt,omitempty"`
}

// +kubebuilder:skipversion
type DeviceType struct {
	DeviceCreateDate *metav1.Time `json:"deviceCreateDate,omitempty"`

	DeviceLastAuthenticatedDate *metav1.Time `json:"deviceLastAuthenticatedDate,omitempty"`

	DeviceLastModifiedDate *metav1.Time `json:"deviceLastModifiedDate,omitempty"`
}

// +kubebuilder:skipversion
type DomainDescriptionType struct {
	AWSAccountID *string `json:"awsAccountID,omitempty"`

	CloudFrontDistribution *string `json:"cloudFrontDistribution,omitempty"`
	// The configuration for a custom domain that hosts the sign-up and sign-in
	// webpages for your application.
	CustomDomainConfig *CustomDomainConfigType `json:"customDomainConfig,omitempty"`

	Domain *string `json:"domain,omitempty"`

	S3Bucket *string `json:"s3Bucket,omitempty"`

	Status *string `json:"status,omitempty"`

	UserPoolID *string `json:"userPoolID,omitempty"`

	Version *string `json:"version,omitempty"`
}

// +kubebuilder:skipversion
type EmailConfigurationType struct {
	ConfigurationSet *string `json:"configurationSet,omitempty"`

	EmailSendingAccount *string `json:"emailSendingAccount,omitempty"`

	From *string `json:"from,omitempty"`

	ReplyToEmailAddress *string `json:"replyToEmailAddress,omitempty"`

	SourceARN *string `json:"sourceARN,omitempty"`
}

// +kubebuilder:skipversion
type EventContextDataType struct {
	City *string `json:"city,omitempty"`

	Country *string `json:"country,omitempty"`

	DeviceName *string `json:"deviceName,omitempty"`

	IPAddress *string `json:"ipAddress,omitempty"`

	Timezone *string `json:"timezone,omitempty"`
}

// +kubebuilder:skipversion
type EventFeedbackType struct {
	FeedbackDate *metav1.Time `json:"feedbackDate,omitempty"`

	Provider *string `json:"provider,omitempty"`
}

// +kubebuilder:skipversion
type EventRiskType struct {
	CompromisedCredentialsDetected *bool `json:"compromisedCredentialsDetected,omitempty"`
}

// +kubebuilder:skipversion
type GroupType struct {
	CreationDate *metav1.Time `json:"creationDate,omitempty"`

	Description *string `json:"description,omitempty"`

	GroupName *string `json:"groupName,omitempty"`

	LastModifiedDate *metav1.Time `json:"lastModifiedDate,omitempty"`

	Precedence *int64 `json:"precedence,omitempty"`

	RoleARN *string `json:"roleARN,omitempty"`

	UserPoolID *string `json:"userPoolID,omitempty"`
}

// +kubebuilder:skipversion
type HTTPHeader struct {
	HeaderName *string `json:"headerName,omitempty"`

	HeaderValue *string `json:"headerValue,omitempty"`
}

// +kubebuilder:skipversion
type IdentityProviderType struct {
	AttributeMapping map[string]*string `json:"attributeMapping,omitempty"`

	CreationDate *metav1.Time `json:"creationDate,omitempty"`

	IDpIdentifiers []*string `json:"idpIdentifiers,omitempty"`

	LastModifiedDate *metav1.Time `json:"lastModifiedDate,omitempty"`

	ProviderName *string `json:"providerName,omitempty"`

	ProviderType *string `json:"providerType,omitempty"`

	UserPoolID *string `json:"userPoolID,omitempty"`
}

// +kubebuilder:skipversion
type LambdaConfigType struct {
	CreateAuthChallenge *string `json:"createAuthChallenge,omitempty"`
	// A custom email sender Lambda configuration type.
	CustomEmailSender *CustomEmailLambdaVersionConfigType `json:"customEmailSender,omitempty"`

	CustomMessage *string `json:"customMessage,omitempty"`
	// A custom SMS sender Lambda configuration type.
	CustomSMSSender *CustomSMSLambdaVersionConfigType `json:"customSMSSender,omitempty"`

	DefineAuthChallenge *string `json:"defineAuthChallenge,omitempty"`

	KMSKeyID *string `json:"kmsKeyID,omitempty"`

	PostAuthentication *string `json:"postAuthentication,omitempty"`

	PostConfirmation *string `json:"postConfirmation,omitempty"`

	PreAuthentication *string `json:"preAuthentication,omitempty"`

	PreSignUp *string `json:"preSignUp,omitempty"`

	PreTokenGeneration *string `json:"preTokenGeneration,omitempty"`

	UserMigration *string `json:"userMigration,omitempty"`

	VerifyAuthChallengeResponse *string `json:"verifyAuthChallengeResponse,omitempty"`
}

// +kubebuilder:skipversion
type MessageTemplateType struct {
	EmailMessage *string `json:"emailMessage,omitempty"`

	EmailSubject *string `json:"emailSubject,omitempty"`

	SMSMessage *string `json:"sMSMessage,omitempty"`
}

// +kubebuilder:skipversion
type NewDeviceMetadataType struct {
	DeviceGroupKey *string `json:"deviceGroupKey,omitempty"`
}

// +kubebuilder:skipversion
type NotifyConfigurationType struct {
	From *string `json:"from,omitempty"`

	ReplyTo *string `json:"replyTo,omitempty"`

	SourceARN *string `json:"sourceARN,omitempty"`
}

// +kubebuilder:skipversion
type NumberAttributeConstraintsType struct {
	MaxValue *string `json:"maxValue,omitempty"`

	MinValue *string `json:"minValue,omitempty"`
}

// +kubebuilder:skipversion
type PasswordPolicyType struct {
	MinimumLength *int64 `json:"minimumLength,omitempty"`

	RequireLowercase *bool `json:"requireLowercase,omitempty"`

	RequireNumbers *bool `json:"requireNumbers,omitempty"`

	RequireSymbols *bool `json:"requireSymbols,omitempty"`

	RequireUppercase *bool `json:"requireUppercase,omitempty"`

	TemporaryPasswordValidityDays *int64 `json:"temporaryPasswordValidityDays,omitempty"`
}

// +kubebuilder:skipversion
type ProviderDescription struct {
	CreationDate *metav1.Time `json:"creationDate,omitempty"`

	LastModifiedDate *metav1.Time `json:"lastModifiedDate,omitempty"`

	ProviderName *string `json:"providerName,omitempty"`

	ProviderType *string `json:"providerType,omitempty"`
}

// +kubebuilder:skipversion
type ProviderUserIdentifierType struct {
	ProviderAttributeName *string `json:"providerAttributeName,omitempty"`

	ProviderAttributeValue *string `json:"providerAttributeValue,omitempty"`

	ProviderName *string `json:"providerName,omitempty"`
}

// +kubebuilder:skipversion
type RecoveryOptionType struct {
	Name *string `json:"name,omitempty"`

	Priority *int64 `json:"priority,omitempty"`
}

// +kubebuilder:skipversion
type ResourceServerScopeType struct {
	ScopeDescription *string `json:"scopeDescription,omitempty"`

	ScopeName *string `json:"scopeName,omitempty"`
}

// +kubebuilder:skipversion
type ResourceServerType struct {
	Identifier *string `json:"identifier,omitempty"`

	Name *string `json:"name,omitempty"`

	Scopes []*ResourceServerScopeType `json:"scopes,omitempty"`

	UserPoolID *string `json:"userPoolID,omitempty"`
}

// +kubebuilder:skipversion
type RiskConfigurationType struct {
	ClientID *string `json:"clientID,omitempty"`

	LastModifiedDate *metav1.Time `json:"lastModifiedDate,omitempty"`

	UserPoolID *string `json:"userPoolID,omitempty"`
}

// +kubebuilder:skipversion
type SMSMFASettingsType struct {
	Enabled *bool `json:"enabled,omitempty"`

	PreferredMFA *bool `json:"preferredMFA,omitempty"`
}

// +kubebuilder:skipversion
type SchemaAttributeType struct {
	AttributeDataType *string `json:"attributeDataType,omitempty"`

	DeveloperOnlyAttribute *bool `json:"developerOnlyAttribute,omitempty"`

	Mutable *bool `json:"mutable,omitempty"`

	Name *string `json:"name,omitempty"`
	// The minimum and maximum values of an attribute that is of the number data
	// type.
	NumberAttributeConstraints *NumberAttributeConstraintsType `json:"numberAttributeConstraints,omitempty"`

	Required *bool `json:"required,omitempty"`
	// The constraints associated with a string attribute.
	StringAttributeConstraints *StringAttributeConstraintsType `json:"stringAttributeConstraints,omitempty"`
}

// +kubebuilder:skipversion
type SmsConfigurationType struct {
	ExternalID *string `json:"externalID,omitempty"`

	SNSCallerARN *string `json:"snsCallerARN,omitempty"`

	SNSRegion *string `json:"snsRegion,omitempty"`
}

// +kubebuilder:skipversion
type SmsMFAConfigType struct {
	SmsAuthenticationMessage *string `json:"smsAuthenticationMessage,omitempty"`
	// The SMS configuration type is the settings that your Amazon Cognito user
	// pool must use to send an SMS message from your Amazon Web Services account
	// through Amazon Simple Notification Service. To send SMS messages with Amazon
	// SNS in the Amazon Web Services Region that you want, the Amazon Cognito user
	// pool uses an Identity and Access Management (IAM) role in your Amazon Web
	// Services account.
	SmsConfiguration *SmsConfigurationType `json:"smsConfiguration,omitempty"`
}

// +kubebuilder:skipversion
type SoftwareTokenMFAConfigType struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// +kubebuilder:skipversion
type SoftwareTokenMFASettingsType struct {
	Enabled *bool `json:"enabled,omitempty"`

	PreferredMFA *bool `json:"preferredMFA,omitempty"`
}

// +kubebuilder:skipversion
type StringAttributeConstraintsType struct {
	MaxLength *string `json:"maxLength,omitempty"`

	MinLength *string `json:"minLength,omitempty"`
}

// +kubebuilder:skipversion
type TokenValidityUnitsType struct {
	AccessToken *string `json:"accessToken,omitempty"`

	IDToken *string `json:"idToken,omitempty"`

	RefreshToken *string `json:"refreshToken,omitempty"`
}

// +kubebuilder:skipversion
type UICustomizationType struct {
	ClientID *string `json:"clientID,omitempty"`

	CreationDate *metav1.Time `json:"creationDate,omitempty"`

	LastModifiedDate *metav1.Time `json:"lastModifiedDate,omitempty"`

	UserPoolID *string `json:"userPoolID,omitempty"`
}

// +kubebuilder:skipversion
type UserContextDataType struct {
	EncodedData *string `json:"encodedData,omitempty"`
}

// +kubebuilder:skipversion
type UserImportJobType struct {
	CloudWatchLogsRoleARN *string `json:"cloudWatchLogsRoleARN,omitempty"`

	CompletionDate *metav1.Time `json:"completionDate,omitempty"`

	CreationDate *metav1.Time `json:"creationDate,omitempty"`

	StartDate *metav1.Time `json:"startDate,omitempty"`

	UserPoolID *string `json:"userPoolID,omitempty"`
}

// +kubebuilder:skipversion
type UserPoolAddOnsType struct {
	AdvancedSecurityMode *string `json:"advancedSecurityMode,omitempty"`
}

// +kubebuilder:skipversion
type UserPoolClientDescription struct {
	ClientID *string `json:"clientID,omitempty"`

	ClientName *string `json:"clientName,omitempty"`

	UserPoolID *string `json:"userPoolID,omitempty"`
}

// +kubebuilder:skipversion
type UserPoolClientType struct {
	AccessTokenValidity *int64 `json:"accessTokenValidity,omitempty"`

	AllowedOAuthFlows []*string `json:"allowedOAuthFlows,omitempty"`

	AllowedOAuthFlowsUserPoolClient *bool `json:"allowedOAuthFlowsUserPoolClient,omitempty"`

	AllowedOAuthScopes []*string `json:"allowedOAuthScopes,omitempty"`
	// The Amazon Pinpoint analytics configuration for collecting metrics for a
	// user pool.
	//
	// In Regions where Amazon Pinpointisn't available, user pools only support
	// sending events to Amazon Pinpoint projects in us-east-1. In Regions where
	// Amazon Pinpoint is available, user pools support sending events to Amazon
	// Pinpoint projects within that same Region.
	AnalyticsConfiguration *AnalyticsConfigurationType `json:"analyticsConfiguration,omitempty"`

	CallbackURLs []*string `json:"callbackURLs,omitempty"`

	ClientID *string `json:"clientID,omitempty"`

	ClientName *string `json:"clientName,omitempty"`

	ClientSecret *string `json:"clientSecret,omitempty"`

	CreationDate *metav1.Time `json:"creationDate,omitempty"`

	DefaultRedirectURI *string `json:"defaultRedirectURI,omitempty"`

	EnableTokenRevocation *bool `json:"enableTokenRevocation,omitempty"`

	ExplicitAuthFlows []*string `json:"explicitAuthFlows,omitempty"`

	IDTokenValidity *int64 `json:"idTokenValidity,omitempty"`

	LastModifiedDate *metav1.Time `json:"lastModifiedDate,omitempty"`

	LogoutURLs []*string `json:"logoutURLs,omitempty"`

	PreventUserExistenceErrors *string `json:"preventUserExistenceErrors,omitempty"`

	ReadAttributes []*string `json:"readAttributes,omitempty"`

	RefreshTokenValidity *int64 `json:"refreshTokenValidity,omitempty"`

	SupportedIdentityProviders []*string `json:"supportedIdentityProviders,omitempty"`
	// The data type for TokenValidityUnits that specifics the time measurements
	// for token validity.
	TokenValidityUnits *TokenValidityUnitsType `json:"tokenValidityUnits,omitempty"`

	UserPoolID *string `json:"userPoolID,omitempty"`

	WriteAttributes []*string `json:"writeAttributes,omitempty"`
}

// +kubebuilder:skipversion
type UserPoolDescriptionType struct {
	CreationDate *metav1.Time `json:"creationDate,omitempty"`

	ID *string `json:"id,omitempty"`
	// Specifies the configuration for Lambda triggers.
	LambdaConfig *LambdaConfigType `json:"lambdaConfig,omitempty"`

	LastModifiedDate *metav1.Time `json:"lastModifiedDate,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`
}

// +kubebuilder:skipversion
type UserPoolPolicyType struct {
	// The password policy type.
	PasswordPolicy *PasswordPolicyType `json:"passwordPolicy,omitempty"`
}

// +kubebuilder:skipversion
type UserPoolType struct {
	// The data type for AccountRecoverySetting.
	AccountRecoverySetting *AccountRecoverySettingType `json:"accountRecoverySetting,omitempty"`
	// The configuration for creating a new user profile.
	AdminCreateUserConfig *AdminCreateUserConfigType `json:"adminCreateUserConfig,omitempty"`

	AliasAttributes []*string `json:"aliasAttributes,omitempty"`

	ARN *string `json:"arn,omitempty"`

	AutoVerifiedAttributes []*string `json:"autoVerifiedAttributes,omitempty"`

	CreationDate *metav1.Time `json:"creationDate,omitempty"`

	CustomDomain *string `json:"customDomain,omitempty"`
	// The device tracking configuration for a user pool. A user pool with device
	// tracking deactivated returns a null value.
	//
	// When you provide values for any DeviceConfiguration field, you activate device
	// tracking.
	DeviceConfiguration *DeviceConfigurationType `json:"deviceConfiguration,omitempty"`

	Domain *string `json:"domain,omitempty"`
	// The email configuration of your user pool. The email configuration type sets
	// your preferred sending method, Amazon Web Services Region, and sender for
	// messages from your user pool.
	//
	// Amazon Cognito can send email messages with Amazon Simple Email Service resources
	// in the Amazon Web Services Region where you created your user pool, and in
	// alternate Regions in some cases. For more information on the supported Regions,
	// see Email settings for Amazon Cognito user pools (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-email.html).
	EmailConfiguration *EmailConfigurationType `json:"emailConfiguration,omitempty"`

	EmailConfigurationFailure *string `json:"emailConfigurationFailure,omitempty"`

	EmailVerificationMessage *string `json:"emailVerificationMessage,omitempty"`

	EmailVerificationSubject *string `json:"emailVerificationSubject,omitempty"`

	EstimatedNumberOfUsers *int64 `json:"estimatedNumberOfUsers,omitempty"`

	ID *string `json:"id,omitempty"`
	// Specifies the configuration for Lambda triggers.
	LambdaConfig *LambdaConfigType `json:"lambdaConfig,omitempty"`

	LastModifiedDate *metav1.Time `json:"lastModifiedDate,omitempty"`

	MFAConfiguration *string `json:"mfaConfiguration,omitempty"`

	Name *string `json:"name,omitempty"`
	// The policy associated with a user pool.
	Policies *UserPoolPolicyType `json:"policies,omitempty"`

	SchemaAttributes []*SchemaAttributeType `json:"schemaAttributes,omitempty"`

	SmsAuthenticationMessage *string `json:"smsAuthenticationMessage,omitempty"`
	// The SMS configuration type is the settings that your Amazon Cognito user
	// pool must use to send an SMS message from your Amazon Web Services account
	// through Amazon Simple Notification Service. To send SMS messages with Amazon
	// SNS in the Amazon Web Services Region that you want, the Amazon Cognito user
	// pool uses an Identity and Access Management (IAM) role in your Amazon Web
	// Services account.
	SmsConfiguration *SmsConfigurationType `json:"smsConfiguration,omitempty"`

	SmsConfigurationFailure *string `json:"smsConfigurationFailure,omitempty"`

	SmsVerificationMessage *string `json:"smsVerificationMessage,omitempty"`

	Status *string `json:"status,omitempty"`
	// The user pool add-ons type.
	UserPoolAddOns *UserPoolAddOnsType `json:"userPoolAddOns,omitempty"`

	UserPoolTags map[string]*string `json:"userPoolTags,omitempty"`

	UsernameAttributes []*string `json:"usernameAttributes,omitempty"`
	// The username configuration type.
	UsernameConfiguration *UsernameConfigurationType `json:"usernameConfiguration,omitempty"`
	// The template for verification messages.
	VerificationMessageTemplate *VerificationMessageTemplateType `json:"verificationMessageTemplate,omitempty"`
}

// +kubebuilder:skipversion
type UserType struct {
	Enabled *bool `json:"enabled,omitempty"`

	UserCreateDate *metav1.Time `json:"userCreateDate,omitempty"`

	UserLastModifiedDate *metav1.Time `json:"userLastModifiedDate,omitempty"`
}

// +kubebuilder:skipversion
type UsernameConfigurationType struct {
	CaseSensitive *bool `json:"caseSensitive,omitempty"`
}

// +kubebuilder:skipversion
type VerificationMessageTemplateType struct {
	DefaultEmailOption *string `json:"defaultEmailOption,omitempty"`

	EmailMessage *string `json:"emailMessage,omitempty"`

	EmailMessageByLink *string `json:"emailMessageByLink,omitempty"`

	EmailSubject *string `json:"emailSubject,omitempty"`

	EmailSubjectByLink *string `json:"emailSubjectByLink,omitempty"`

	SmsMessage *string `json:"smsMessage,omitempty"`
}
