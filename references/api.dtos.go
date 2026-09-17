//go:build ignore

/* Options:
Date: 2026-09-04 14:56:08
Version: 10.08
Tip: To override a DTO option, remove "//" prefix before updating
BaseUrl: http://localhost:5002

//GlobalNamespace:
//MakePropertiesOptional: False
//AddServiceStackTypes: True
//AddResponseStatus: False
//AddImplicitVersion:
//AddDescriptionAsComments: True
//IncludeTypes:
//ExcludeTypes:
//DefaultImports:
*/

package dtos

import (
	"time"
)

type IReturn struct {
}

type IReturnVoid struct {
}

type IHasSessionId struct {
	SessionId *string `json:"sessionId,omitempty"`
}

type IHasBearerToken struct {
	BearerToken *string `json:"bearerToken,omitempty"`
}

type IPost struct {
}

// @DataContract(Namespace="http://codemash.io/types/")
type RequestBase struct {
	/** @description Specify culture code when your response from the API should be localised. E.g.: en */
	// @DataMember
	// @ApiMember(DataType="string", Description="Specify culture code when your response from the API should be localised. E.g.: en", Name="CultureCode", ParameterType="header")
	CultureCode *string `json:"cultureCode,omitempty"`
	/** @description TimeZone */
	// @DataMember
	// @ApiMember(DataType="string", Description="TimeZone", Name="TimeZoneId", ParameterType="header")
	TimeZoneId *string `json:"timeZoneId,omitempty"`
	/** @description The CodeMash API version used to fetch data from the API. If not specified, the last version will be used.  E.g.: v3 */
	// @DataMember
	// @ApiMember(DataType="string", Description="The CodeMash API version used to fetch data from the API. If not specified, the last version will be used.  E.g.: v3", IsRequired=true, Name="version", ParameterType="path")
	Version string `json:"version"`
	/** @description CorrelationId for each request */
	// @DataMember
	// @ApiMember(DataType="string", Description="CorrelationId for each request", Name="CorrelationId", ParameterType="header")
	CorrelationId *string `json:"correlationId,omitempty"`
}

type ICultureBasedRequest struct {
	CultureCode *string `json:"cultureCode,omitempty"`
}

type IVersionBasedRequest struct {
	Version string `json:"version"`
}

type IHasCorrelationIdRequest struct {
	CorrelationId *string `json:"correlationId,omitempty"`
}

type EmailAddress struct {
	Address string `json:"address"`
}

type DisplayName struct {
	Value string `json:"value"`
}

type AggregateId struct {
	Value string `json:"value,omitempty"`
}

type AccountId struct {
	AggregateId
}

type UtcDateTime struct {
}

type TimeUnit string

const (
	TimeUnitTicks        TimeUnit = "Ticks"
	TimeUnitMilliseconds          = "Milliseconds"
	TimeUnitSeconds               = "Seconds"
	TimeUnitMinutes               = "Minutes"
	TimeUnitHours                 = "Hours"
)

type ExpirationToken struct {
	Items int64    `json:"items,omitempty"`
	Unit  TimeUnit `json:"unit,omitempty"`
	Value int64    `json:"value,omitempty"`
}

type CodeMashSubscriptionId struct {
	AggregateId
}

type ProjectId struct {
	AggregateId
}

type IntegrationId struct {
	AggregateId
}

type ResourceRefKind string

const (
	ResourceRefKindContact         ResourceRefKind = "Contact"
	ResourceRefKindDocument                        = "Document"
	ResourceRefKindFile                            = "File"
	ResourceRefKindPaymentCustomer                 = "PaymentCustomer"
	ResourceRefKindOrder                           = "Order"
	ResourceRefKindPayment                         = "Payment"
	ResourceRefKindProduct                         = "Product"
	ResourceRefKindIntegration                     = "Integration"
)

type ResourceRef struct {
	ProjectId     ProjectId       `json:"projectId"`
	IntegrationId *IntegrationId  `json:"integrationId,omitempty"`
	Kind          ResourceRefKind `json:"kind,omitempty"`
}

type ResourceSource string

const (
	ResourceSourceNorbix       ResourceSource = "Norbix"
	ResourceSourceStripe                      = "Stripe"
	ResourceSourceShopify                     = "Shopify"
	ResourceSourcePayPal                      = "PayPal"
	ResourceSourceAdyen                       = "Adyen"
	ResourceSourceMollie                      = "Mollie"
	ResourceSourcePaddle                      = "Paddle"
	ResourceSourceLemonSqueezy                = "LemonSqueezy"
	ResourceSourceAppleInApp                  = "AppleInApp"
	ResourceSourceGoogleInApp                 = "GoogleInApp"
	ResourceSourceAuthorizeNet                = "AuthorizeNet"
	ResourceSourceBraintree                   = "Braintree"
	ResourceSourceCheckOutCom                 = "CheckOutCom"
	ResourceSourceWooCommerce                 = "WooCommerce"
	ResourceSourceMagento                     = "Magento"
	ResourceSourceWorldpay                    = "Worldpay"
)

type PaymentCustomerRef struct {
	ResourceRef
	Kind       ResourceRefKind `json:"kind,omitempty"`
	Source     ResourceSource  `json:"source,omitempty"`
	ExternalId string          `json:"externalId"`
}

type Quantity struct {
	Value int `json:"value,omitempty"`
}

type CodeMashManagedServiceSubscription struct {
	SubscriptionId     CodeMashSubscriptionId `json:"subscriptionId"`
	PaymentCustomerRef PaymentCustomerRef     `json:"paymentCustomerRef"`
	RefSubscriptionId  string                 `json:"refSubscriptionId"`
	IssuedOn           UtcDateTime            `json:"issuedOn"`
	WillExpireOn       UtcDateTime            `json:"willExpireOn"`
	ProjectCap         Quantity               `json:"projectCap"`
	IsTrial            bool                   `json:"isTrial,omitempty"`
}

type DomainUrl struct {
	Value string `json:"value"`
}

type CodeMashLicense struct {
	CodeMashManagedServiceSubscription
	Domain       DomainUrl `json:"domain"`
	AccountId    AccountId `json:"accountId"`
	IsEnterprise bool      `json:"isEnterprise,omitempty"`
}

type Tag struct {
}

type TagDescription struct {
	DisplayName DisplayName `json:"displayName"`
	Description *string     `json:"description,omitempty"`
}

// @DataContract
type MessageTranslation struct {
}

type TagTranslation struct {
	MessageTranslation
}

type BaseTagDefinition struct {
	Tag          Tag              `json:"tag"`
	Translations []TagTranslation `json:"translations"`
}

type GroupDefinition struct {
	BaseTagDefinition
}

type CommunicationChannel string

const (
	CommunicationChannelTransactional CommunicationChannel = "Transactional"
	CommunicationChannelMarketing                          = "Marketing"
	CommunicationChannelSystem                             = "System"
)

type DeliveryChannel string

const (
	DeliveryChannelEmail        DeliveryChannel = "Email"
	DeliveryChannelPush                         = "Push"
	DeliveryChannelSms                          = "Sms"
	DeliveryChannelWebPush                      = "WebPush"
	DeliveryChannelInApp                        = "InApp"
	DeliveryChannelChatBot                      = "ChatBot"
	DeliveryChannelChatPlatform                 = "ChatPlatform"
)

type TagDefinition struct {
	BaseTagDefinition
	DefaultDelivery map[DeliveryChannel]bool `json:"defaultDelivery"`
}

// @DataContract
type ProjectName struct {
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	UniqueName string `json:"uniqueName"`
}

type NorbixRegion struct {
	Code string `json:"code"`
}

type Continent string

const (
	ContinentAfrica       Continent = "Africa"
	ContinentAntarctica             = "Antarctica"
	ContinentAsia                   = "Asia"
	ContinentEurope                 = "Europe"
	ContinentNorthAmerica           = "NorthAmerica"
	ContinentOceania                = "Oceania"
	ContinentSouthAmerica           = "SouthAmerica"
)

// @DataContract
type ProjectRegion struct {
	// @DataMember
	Region NorbixRegion `json:"region"`
	// @DataMember
	Name *string `json:"name,omitempty"`
	// @DataMember
	Continent *Continent `json:"continent,omitempty"`
}

type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type FileResourceId struct {
	Value string `json:"value,omitempty"`
}

type FileChecksum struct {
	Algorithm string `json:"algorithm"`
	Hash      string `json:"hash"`
}

// @DataContract
type FileResource struct {
	// @DataMember
	Id FileResourceId `json:"id"`
	// @DataMember
	OriginalFileName string `json:"originalFileName"`
	// @DataMember
	Extension string `json:"extension"`
	// @DataMember
	SizeBytes *int64 `json:"sizeBytes,omitempty"`
	// @DataMember
	Checksum *FileChecksum `json:"checksum,omitempty"`
	// @DataMember
	StoredFileName string `json:"storedFileName"`
}

type FileProvider string

const (
	FileProviderLocal              FileProvider = "Local"
	FileProviderAwsS3                           = "AwsS3"
	FileProviderAzureBlobStorage                = "AzureBlobStorage"
	FileProviderGoogleCloudStorage              = "GoogleCloudStorage"
	FileProviderFtp                             = "Ftp"
	FileProviderAppleICloud                     = "AppleICloud"
	FileProviderDropBox                         = "DropBox"
	FileProviderGoogleDrive                     = "GoogleDrive"
)

// @DataContract
type FileResourceRef struct {
	// @DataMember(Order=1)
	Resource FileResource `json:"resource"`
	// @DataMember(Order=2)
	IntegrationId IntegrationId `json:"integrationId"`
	// @DataMember(Order=3)
	Provider FileProvider `json:"provider,omitempty"`
	// @DataMember(Order=4)
	Path string `json:"path"`
}

type ProjectLogo struct {
	FileResource FileResourceRef `json:"fileResource"`
	PublicUrl    string          `json:"publicUrl"`
}

type ProjectIcon struct {
	FileResource FileResourceRef `json:"fileResource"`
	PublicUrl    string          `json:"publicUrl"`
}

// @DataContract
type BrandColor struct {
	// @DataMember
	Value string `json:"value"`
}

// @DataContract
type TimeZone struct {
	// @DataMember
	ZoneId string `json:"zoneId"`
}

// @DataContract
type GroupTags struct {
	// @DataMember
	Group Tag `json:"group"`
	// @DataMember
	Tags []Tag `json:"tags"`
}

// @DataContract
type ProjectCommunicationChannel struct {
	// @DataMember
	Channel CommunicationChannel `json:"channel,omitempty"`
	// @DataMember
	Groups []GroupTags `json:"groups"`
}

// @DataContract
type ProjectCommunication struct {
	// @DataMember
	Channels []ProjectCommunicationChannel `json:"channels"`
	// @DataMember
	Groups []GroupDefinition `json:"groups"`
	// @DataMember
	Tags []TagDefinition `json:"tags"`
}

type AuthId struct {
	Value string `json:"value,omitempty"`
}

type DeviceId struct {
	Id string `json:"id,omitempty"`
}

type DeviceType string

const (
	DeviceTypeUnknown DeviceType = "Unknown"
	DeviceTypePhone              = "Phone"
	DeviceTypeTablet             = "Tablet"
	DeviceTypeDesktop            = "Desktop"
	DeviceTypeTv                 = "Tv"
)

type PushDeviceToken struct {
	Token string `json:"token"`
}

type PushDeviceDeliveryFamily string

const (
	PushDeviceDeliveryFamilyIos     PushDeviceDeliveryFamily = "Ios"
	PushDeviceDeliveryFamilyAndroid                          = "Android"
	PushDeviceDeliveryFamilyChrome                           = "Chrome"
	PushDeviceDeliveryFamilySafari                           = "Safari"
	PushDeviceDeliveryFamilyExpo                             = "Expo"
)

// @DataContract
type PushDeviceDeliveryToken struct {
	// @DataMember
	PushDeviceToken PushDeviceToken `json:"pushDeviceToken"`
	// @DataMember
	DeliveryFamily PushDeviceDeliveryFamily `json:"deliveryFamily,omitempty"`
}

// @DataContract
type PushDevice struct {
	// @DataMember
	Id DeviceId `json:"id"`
	// @DataMember
	Brand *string `json:"brand,omitempty"`
	// @DataMember
	Manufacturer *string `json:"manufacturer,omitempty"`
	// @DataMember
	ModelName *string `json:"modelName,omitempty"`
	// @DataMember
	DeviceName *string `json:"deviceName,omitempty"`
	// @DataMember
	DeviceType *DeviceType `json:"deviceType,omitempty"`
	// @DataMember
	OsName *string `json:"osName,omitempty"`
	// @DataMember
	OsVersion *string `json:"osVersion,omitempty"`
	// @DataMember
	PlatformApiLevel *int `json:"platformApiLevel,omitempty"`
	// @DataMember
	Token PushDeviceDeliveryToken `json:"token"`
}

// @DataContract(Namespace="http://codemash.io/types/")
type CodeMashRequestBase struct {
	RequestBase
	/** @description ID of your project. Can be passed in a header as norbix-project-id. */
	// @DataMember
	// @ApiMember(DataType="string", Description="ID of your project. Can be passed in a header as norbix-project-id.", IsRequired=true, Name="norbix-project-id", ParameterType="header")
	ProjectId string `json:"projectId"`
	/** @description Target environment for this request (e.g. TEST, STAGING). Optional — when omitted the request runs against PROD. Can be passed in a header as norbix-env. */
	// @DataMember
	// @ApiMember(DataType="string", Description="Target environment for this request (e.g. TEST, STAGING). Optional — when omitted the request runs against PROD. Can be passed in a header as norbix-env.", Name="norbix-env", ParameterType="header")
	Env *string `json:"env,omitempty"`
}

type IHasProjectId struct {
	ProjectId string `json:"projectId"`
}

type IHasEnv struct {
	Env *string `json:"env,omitempty"`
}

type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale        = "Female"
	GenderOther         = "Other"
)

type MarketingBlockReason string

const (
	MarketingBlockReasonUnspecified  MarketingBlockReason = "Unspecified"
	MarketingBlockReasonUnsubscribed                      = "Unsubscribed"
	MarketingBlockReasonComplaint                         = "Complaint"
	MarketingBlockReasonHardBounce                        = "HardBounce"
	MarketingBlockReasonInvalidEmail                      = "InvalidEmail"
	MarketingBlockReasonAdminBlock                        = "AdminBlock"
)

type UserGeneralInfoDto struct {
	Phone                     *string                `json:"phone,omitempty"`
	PrimaryEmail              *string                `json:"primaryEmail,omitempty"`
	DisplayName               *string                `json:"displayName,omitempty"`
	FirstName                 *string                `json:"firstName,omitempty"`
	LastName                  *string                `json:"lastName,omitempty"`
	FullName                  *string                `json:"fullName,omitempty"`
	AddressLine1              *string                `json:"addressLine1,omitempty"`
	AddressLine2              *string                `json:"addressLine2,omitempty"`
	Country                   *string                `json:"country,omitempty"`
	City                      *string                `json:"city,omitempty"`
	State                     *string                `json:"state,omitempty"`
	PostalCode                *string                `json:"postalCode,omitempty"`
	Company                   *string                `json:"company,omitempty"`
	Gender                    *Gender                `json:"gender,omitempty"`
	BirthDate                 *int64                 `json:"birthDate,omitempty"`
	TimeZone                  *string                `json:"timeZone,omitempty"`
	Language                  *string                `json:"language,omitempty"`
	BlockAllMarketingMessages bool                   `json:"blockAllMarketingMessages,omitempty"`
	BlockedTags               map[string][]string    `json:"blockedTags,omitempty"`
	BlockReasons              []MarketingBlockReason `json:"blockReasons,omitempty"`
	ExtraMetadata             *string                `json:"extraMetadata,omitempty"`
	Notes                     *string                `json:"notes,omitempty"`
}

// @DataContract
type SaveUser struct {
	CodeMashRequestBase
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description User Info */
	// @DataMember
	// @ApiMember(DataType="object", Description="User Info", Name="UserGeneralInfo", ParameterType="body")
	UserGeneralInfo *UserGeneralInfoDto `json:"userGeneralInfo,omitempty"`
	/** @description Attach this login to an existing user id. Optional. */
	// @DataMember
	// @ApiMember(Description="Attach this login to an existing user id. Optional.")
	UserId *string `json:"userId,omitempty"`
	/** @description Ignore UserRegistersAsRole from Membership Settings */
	// @DataMember
	// @ApiMember(DataType="boolean", Description="Ignore UserRegistersAsRole from Membership Settings", Name="IgnoreUserRegistersAsRole", ParameterType="body")
	IgnoreUserRegistersAsRole bool `json:"ignoreUserRegistersAsRole,omitempty"`
}

// @DataContract
type SaveUserWithRolesBase struct {
	SaveUser
	// @DataMember
	Roles []string `json:"roles"`
}

type Env struct {
	Value  string `json:"value"`
	IsProd bool   `json:"isProd,omitempty"`
}

type CursorArgs struct {
	Field string `json:"field"`
	Order int    `json:"order,omitempty"`
}

type PagingArgs struct {
	CursorArgs    *CursorArgs `json:"cursorArgs,omitempty"`
	PageSize      *int        `json:"pageSize,omitempty"`
	StartingAfter *string     `json:"startingAfter,omitempty"`
	EndingBefore  *string     `json:"endingBefore,omitempty"`
}

type CodeMashListPaginationRequestBase struct {
	RequestBase
	/** @description ID of your project. Can be passed in a header as norbix-project-id. */
	// @DataMember
	// @ApiMember(DataType="string", Description="ID of your project. Can be passed in a header as norbix-project-id.", IsRequired=true, Name="norbix-project-id", ParameterType="header")
	ProjectId string `json:"projectId"`
	/** @description Target environment for this request (e.g. TEST, STAGING). Optional — when omitted the request runs against PROD. Can be passed in a header as norbix-env. */
	// @DataMember
	// @ApiMember(DataType="string", Description="Target environment for this request (e.g. TEST, STAGING). Optional — when omitted the request runs against PROD. Can be passed in a header as norbix-env.", Name="norbix-env", ParameterType="header")
	Env         *string `json:"env,omitempty"`
	ResolvedEnv Env     `json:"resolvedEnv"`
	/** @description Cursor token — fetch the page AFTER this item. */
	// @DataMember
	// @ApiMember(DataType="string", Description="Cursor token — fetch the page AFTER this item.", Name="startingAfter", ParameterType="query")
	StartingAfter *string `json:"startingAfter,omitempty"`
	/** @description Cursor token — fetch the page BEFORE this item. */
	// @DataMember
	// @ApiMember(DataType="string", Description="Cursor token — fetch the page BEFORE this item.", Name="endingBefore", ParameterType="query")
	EndingBefore *string `json:"endingBefore,omitempty"`
	/** @description Amount of records to return. */
	// @DataMember
	// @ApiMember(DataType="integer", Description="Amount of records to return.", Format="int32", Name="pageSize", ParameterType="query")
	PageSize *int `json:"pageSize,omitempty"`
	/** @description Paging */
	// @ApiMember(DataType="object", Description="Paging", Name="paging", ParameterType="body")
	Paging *PagingArgs `json:"paging,omitempty"`
}

type IPasskeyCeremonyRequest struct {
}

type Integration struct {
	IntegrationId                    IntegrationId `json:"integrationId"`
	Env                              Env           `json:"env"`
	Capability                       string        `json:"capability,omitempty"`
	IsSystemOwned                    bool          `json:"isSystemOwned,omitempty"`
	IntegrationName                  DisplayName   `json:"integrationName"`
	IsEnabled                        bool          `json:"isEnabled,omitempty"`
	IsConfigured                     bool          `json:"isConfigured,omitempty"`
	LastIntegrationTestAtUtc         *time.Time    `json:"lastIntegrationTestAtUtc,omitempty"`
	LastIntegrationTestSucceeded     *bool         `json:"lastIntegrationTestSucceeded,omitempty"`
	LastIntegrationTestErrorMessages IReadOnlyList `json:"lastIntegrationTestErrorMessages"`
	HumanDeliveryConfirmedAtUtc      *time.Time    `json:"humanDeliveryConfirmedAtUtc,omitempty"`
	IsApprovedThatItWorks            bool          `json:"isApprovedThatItWorks,omitempty"`
}

// @DataContract
type PushProvider string

const (
	PushProviderAppleApns            PushProvider = "AppleApns"
	PushProviderSafariWeb                         = "SafariWeb"
	PushProviderSafariPush                        = "SafariPush"
	PushProviderAndroidFirebase                   = "AndroidFirebase"
	PushProviderChromeWeb                         = "ChromeWeb"
	PushProviderFirefoxWeb                        = "FirefoxWeb"
	PushProviderEdgeWeb                           = "EdgeWeb"
	PushProviderChromePush                        = "ChromePush"
	PushProviderCodeMashIosApp                    = "CodeMashIosApp"
	PushProviderCodeMashAndroidApp                = "CodeMashAndroidApp"
	PushProviderCodeMashSafariPlugin              = "CodeMashSafariPlugin"
	PushProviderCodeMashSafariWeb                 = "CodeMashSafariWeb"
	PushProviderCodeMashChromePlugin              = "CodeMashChromePlugin"
	PushProviderCodeMashChromeWeb                 = "CodeMashChromeWeb"
	PushProviderExpo                              = "Expo"
	PushProviderFake                              = "Fake"
)

type PushIntegration struct {
	Integration
	Provider PushProvider `json:"provider,omitempty"`
}

type TemplateId struct {
	Value string `json:"value,omitempty"`
}

// @DataContract
type Template struct {
	// @DataMember
	TemplateId TemplateId `json:"templateId"`
	// @DataMember
	TemplateName DisplayName `json:"templateName"`
	// @DataMember
	Translations []MessageTranslation `json:"translations"`
	// @DataMember
	CommunicationChannel CommunicationChannel `json:"communicationChannel,omitempty"`
	// @DataMember
	IsActive bool `json:"isActive,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	Tags []Tag `json:"tags,omitempty"`
	// @DataMember
	FileIntegrationId *IntegrationId `json:"fileIntegrationId,omitempty"`
	// @DataMember
	Env Env `json:"env"`
}

// @DataContract
type TemplateCode struct {
}

// @DataContract
type PushTitle struct {
	// @DataMember
	Value TemplateCode `json:"value"`
}

type PushBody struct {
	Value TemplateCode `json:"value"`
}

// @DataContract
type PushMessageContent struct {
	// @DataMember(Order=1)
	Title PushTitle `json:"title"`
	// @DataMember(Order=1)
	SubTitle *PushTitle `json:"subTitle,omitempty"`
	// @DataMember(Order=2)
	Body PushBody `json:"body"`
}

// @DataContract
type PushTemplate struct {
	Template
}

// @DataContract
type CodeMashRelease string

const (
	CodeMashReleaseNotSet         CodeMashRelease = "NotSet"
	CodeMashReleaseCommunity                      = "Community"
	CodeMashReleaseManagedService                 = "ManagedService"
	CodeMashReleaseEnterprise                     = "Enterprise"
)

type CodeMashRuntime string

const (
	CodeMashRuntimeDevelopment CodeMashRuntime = "Development"
	CodeMashRuntimeCI                          = "CI"
	CodeMashRuntimeStaging                     = "Staging"
	CodeMashRuntimeProduction                  = "Production"
)

// @DataContract
type EchoLicenseDto struct {
	// @DataMember(Name="domain")
	Domain *string `json:"domain,omitempty"`
	// @DataMember(Name="accountId")
	AccountId *string `json:"accountId,omitempty"`
	// @DataMember(Name="email")
	Email *string `json:"email,omitempty"`
	// @DataMember(Name="release")
	Release *string `json:"release,omitempty"`
	// @DataMember(Name="expire")
	Expire int64 `json:"expire,omitempty"`
	// @DataMember(Name="isTrial")
	IsTrial bool `json:"isTrial,omitempty"`
	// @DataMember(Name="cap")
	Cap int `json:"cap,omitempty"`
}

type EchoRegionDto struct {
	Code        string `json:"code"`
	DisplayName string `json:"displayName"`
	ApiUrl      string `json:"apiUrl"`
	HubUrl      string `json:"hubUrl"`
}

type PublicBrandDto struct {
	DisplayName string  `json:"displayName"`
	MainColor   *string `json:"mainColor,omitempty"`
	AccentColor *string `json:"accentColor,omitempty"`
	LogoUrl     *string `json:"logoUrl,omitempty"`
	IconUrl     *string `json:"iconUrl,omitempty"`
}

type PublicPasswordPolicyDto struct {
	MinLength      int     `json:"minLength,omitempty"`
	MaxLength      *int    `json:"maxLength,omitempty"`
	MinNumbers     *int    `json:"minNumbers,omitempty"`
	MinUpper       *int    `json:"minUpper,omitempty"`
	MinLower       *int    `json:"minLower,omitempty"`
	MinSpecial     *int    `json:"minSpecial,omitempty"`
	AllowedSpecial *string `json:"allowedSpecial,omitempty"`
}

type PublicAuthDto struct {
	SocialProviders []string                 `json:"socialProviders"`
	Passkey         bool                     `json:"passkey,omitempty"`
	Methods         []string                 `json:"methods,omitempty"`
	PasswordPolicy  *PublicPasswordPolicyDto `json:"passwordPolicy,omitempty"`
}

type ErrorDto struct {
	Message    string            `json:"message"`
	ErrorCode  *string           `json:"errorCode,omitempty"`
	Context    map[string]string `json:"context,omitempty"`
	StackTrace []ErrorDto        `json:"stackTrace,omitempty"`
}

type CodeMashResponseStatus struct {
	IsSuccess bool       `json:"isSuccess,omitempty"`
	Errors    []ErrorDto `json:"errors,omitempty"`
}

// @DataContract
type ResponseBase struct {
	// @DataMember
	ResponseStatus CodeMashResponseStatus `json:"responseStatus"`
}

type AuthType string

const (
	AuthTypeService  AuthType = "Service"
	AuthTypeEmail             = "Email"
	AuthTypeUserName          = "UserName"
	AuthTypePhone             = "Phone"
	AuthTypeGuest             = "Guest"
	AuthTypeSocial            = "Social"
)

type AccessInformationDto struct {
	Ip       *string    `json:"ip,omitempty"`
	Date     *time.Time `json:"date,omitempty"`
	TimeZone *string    `json:"timeZone,omitempty"`
}

type RegistrationDto struct {
	RegistrationInformation AccessInformationDto `json:"registrationInformation"`
}

type LoginDto struct {
	NeedChangePasswordOnNextLogin bool                  `json:"needChangePasswordOnNextLogin,omitempty"`
	LastAccessInformation         *AccessInformationDto `json:"lastAccessInformation,omitempty"`
}

type AuthStatus string

const (
	AuthStatusRegistered        AuthStatus = "Registered"
	AuthStatusPendingValidation            = "PendingValidation"
	AuthStatusActive                       = "Active"
	AuthStatusUnregistered                 = "Unregistered"
	AuthStatusSuspended                    = "Suspended"
	AuthStatusInActive                     = "InActive"
	AuthStatusBlocked                      = "Blocked"
)

type AuthDto struct {
	Id           string              `json:"id"`
	Type         AuthType            `json:"type,omitempty"`
	Email        *string             `json:"email,omitempty"`
	UserName     *string             `json:"userName,omitempty"`
	Registration *RegistrationDto    `json:"registration,omitempty"`
	Login        *LoginDto           `json:"login,omitempty"`
	GeneralInfo  *UserGeneralInfoDto `json:"generalInfo,omitempty"`
	Roles        []string            `json:"roles,omitempty"`
	PushDevices  []string            `json:"pushDevices,omitempty"`
	Tags         []string            `json:"tags,omitempty"`
	Status       AuthStatus          `json:"status,omitempty"`
	CreatedOn    time.Time           `json:"createdOn,omitempty"`
	ModifiedOn   time.Time           `json:"modifiedOn,omitempty"`
}

type PaginatedResponse struct {
	Items         IList   `json:"items"`
	HasMore       bool    `json:"hasMore,omitempty"`
	HasPrevious   bool    `json:"hasPrevious,omitempty"`
	StartingAfter *string `json:"startingAfter,omitempty"`
	EndingBefore  *string `json:"endingBefore,omitempty"`
}

type UserMarketingPreferencesDto struct {
	BlockAllMarketingMessages bool                   `json:"blockAllMarketingMessages,omitempty"`
	BlockedTags               map[string][]string    `json:"blockedTags,omitempty"`
	BlockReasons              []MarketingBlockReason `json:"blockReasons,omitempty"`
}

type PasskeyListItemDto struct {
	CredentialId    string    `json:"credentialId"`
	FriendlyName    string    `json:"friendlyName"`
	RegisteredOnUtc time.Time `json:"registeredOnUtc,omitempty"`
	LastUsedOnUtc   time.Time `json:"lastUsedOnUtc,omitempty"`
	IsRevoked       bool      `json:"isRevoked,omitempty"`
}

type TermMultiParentDto struct {
	// @DataMember
	TaxonomyId string `json:"taxonomyId"`
	// @DataMember
	ParentId string `json:"parentId"`
	// @DataMember
	Name *string `json:"name,omitempty"`
	// @DataMember
	Names map[string]string `json:"names,omitempty"`
}

type TermTreeDto struct {
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	TaxonomyId *string `json:"taxonomyId,omitempty"`
	// @DataMember
	TaxonomyName *string `json:"taxonomyName,omitempty"`
	// @DataMember
	ParentId *string `json:"parentId,omitempty"`
	// @DataMember
	Order *int `json:"order,omitempty"`
	// @DataMember
	Name *string `json:"name,omitempty"`
	// @DataMember
	Names map[string]string `json:"names,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	Descriptions map[string]string `json:"descriptions,omitempty"`
	// @DataMember
	MultiParents []TermMultiParentDto `json:"multiParents,omitempty"`
	// @DataMember
	Meta *Object `json:"meta,omitempty"`
	// @DataMember
	Children []TermTreeDto `json:"children,omitempty"`
}

type TaxonomyTreeDto struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	TaxonomyName string `json:"taxonomyName"`
	// @DataMember
	TaxonomySlug string `json:"taxonomySlug"`
	// @DataMember
	ParentId *string `json:"parentId,omitempty"`
	// @DataMember
	Children []TaxonomyTreeDto `json:"children,omitempty"`
	// @DataMember
	Terms []TermTreeDto `json:"terms,omitempty"`
}

type TermDto struct {
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	TaxonomyId *string `json:"taxonomyId,omitempty"`
	// @DataMember
	TaxonomyName *string `json:"taxonomyName,omitempty"`
	// @DataMember
	ParentId *string `json:"parentId,omitempty"`
	// @DataMember
	Order *int `json:"order,omitempty"`
	// @DataMember
	Name *string `json:"name,omitempty"`
	// @DataMember
	Names map[string]string `json:"names,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	Descriptions map[string]string `json:"descriptions,omitempty"`
	// @DataMember
	MultiParents []TermMultiParentDto `json:"multiParents,omitempty"`
	// @DataMember
	Meta *Object `json:"meta,omitempty"`
}

type JsonSchemaFieldDto struct {
	// @DataMember
	FieldName string `json:"fieldName"`
}

type DataSchemaDto struct {
	// @DataMember
	Json string `json:"json"`
	// @DataMember
	Fields []JsonSchemaFieldDto `json:"fields"`
}

type VisualSchemaDto struct {
	// @DataMember
	Json string `json:"json"`
}

type SchemaSettingsDto struct {
	// @DataMember
	SoftDelete bool `json:"softDelete,omitempty"`
	// @DataMember
	HasRecordOwner bool `json:"hasRecordOwner,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
}

type TriggerType string

const (
	TriggerTypeMembership TriggerType = "Membership"
	TriggerTypeSchema                 = "Schema"
	TriggerTypeFiles                  = "Files"
	TriggerTypePayments               = "Payments"
)

type TriggerActionType string

const (
	TriggerActionTypeCode        TriggerActionType = "Code"
	TriggerActionTypePush                          = "Push"
	TriggerActionTypeSms                           = "Sms"
	TriggerActionTypeEmail                         = "Email"
	TriggerActionTypeWebhookCall                   = "WebhookCall"
	TriggerActionTypeSseCall                       = "SseCall"
	TriggerActionTypeMarketplace                   = "Marketplace"
)

// @DataContract
type TriggerActionDto struct {
	// @DataMember
	Type TriggerActionType `json:"type,omitempty"`
	// @DataMember
	IntegrationId *string `json:"integrationId,omitempty"`
}

// @DataContract
type TriggerDto struct {
	// @DataMember
	Type TriggerType `json:"type,omitempty"`
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	ThenAction TriggerActionDto `json:"thenAction"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	ActivationCode *string `json:"activationCode,omitempty"`
}

type SchemaDto struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	SchemaName string `json:"schemaName"`
	// @DataMember
	SchemaSlug *string `json:"schemaSlug,omitempty"`
	// @DataMember
	Version int `json:"version,omitempty"`
	// @DataMember
	MetaSchemaVersion int `json:"metaSchemaVersion,omitempty"`
	// @DataMember
	DataSchema DataSchemaDto `json:"dataSchema"`
	// @DataMember
	VisualSchema VisualSchemaDto `json:"visualSchema"`
	// @DataMember
	PublishedAt time.Time `json:"publishedAt,omitempty"`
	// @DataMember
	Settings *SchemaSettingsDto `json:"settings,omitempty"`
	// @DataMember
	Triggers []TriggerDto `json:"triggers,omitempty"`
}

type SchemaListProjection struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	SchemaName string `json:"schemaName"`
	// @DataMember
	SchemaTitle string `json:"schemaTitle"`
	// @DataMember
	LatestVersion *int `json:"latestVersion,omitempty"`
	// @DataMember
	HasDraft bool `json:"hasDraft,omitempty"`
	// @DataMember
	MetaSchemaVersion int `json:"metaSchemaVersion,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
}

// @DataContract
type FileChecksumDto struct {
	// @DataMember(Order=1)
	Algorithm string `json:"algorithm"`
	// @DataMember(Order=2)
	Hash string `json:"hash"`
}

// @DataContract
type FileResourceDto struct {
	// @DataMember(Order=1)
	Id string `json:"id"`
	// @DataMember(Order=2)
	OriginalFileName string `json:"originalFileName"`
	// @DataMember(Order=3)
	Extension string `json:"extension"`
	// @DataMember(Order=4)
	StoredFileName string `json:"storedFileName"`
	// @DataMember(Order=5)
	SizeBytes *int64 `json:"sizeBytes,omitempty"`
	// @DataMember(Order=6)
	Checksum *FileChecksumDto `json:"checksum,omitempty"`
}

// @DataContract
type FileResourceRefDto struct {
	// @DataMember(Order=1)
	Resource FileResourceDto `json:"resource"`
	// @DataMember(Order=2)
	IntegrationId string `json:"integrationId"`
	// @DataMember(Order=3)
	Provider FileProvider `json:"provider,omitempty"`
	// @DataMember(Order=4)
	Path string `json:"path"`
}

// @DataContract
type ResponseError struct {
	// @DataMember(Order=1)
	ErrorCode string `json:"errorCode"`
	// @DataMember(Order=2)
	FieldName string `json:"fieldName"`
	// @DataMember(Order=3)
	Message string `json:"message"`
	// @DataMember(Order=4)
	Meta map[string]string `json:"meta,omitempty"`
}

// @DataContract
type ResponseStatus struct {
	// @DataMember(Order=1)
	ErrorCode string `json:"errorCode"`
	// @DataMember(Order=2)
	Message *string `json:"message,omitempty"`
	// @DataMember(Order=3)
	StackTrace *string `json:"stackTrace,omitempty"`
	// @DataMember(Order=4)
	Errors []ResponseError `json:"errors,omitempty"`
	// @DataMember(Order=5)
	Meta map[string]string `json:"meta,omitempty"`
}

type IHasDomainEntityId struct {
	ViewId string `json:"viewId"`
}

type IIntegrationIdentification struct {
	IntegrationId IntegrationId `json:"integrationId"`
	Capability    string        `json:"capability,omitempty"`
	IsSystemOwned bool          `json:"isSystemOwned,omitempty"`
}

type IBindableContract struct {
}

type IHasViewId struct {
	ViewId string `json:"viewId"`
}

type ICursorArgs struct {
	Field string `json:"field"`
	Order int    `json:"order,omitempty"`
}

type StringFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Format *string `json:"format,omitempty"`
	// @DataMember
	Pattern *string `json:"pattern,omitempty"`
	// @DataMember
	MinLength *int `json:"minLength,omitempty"`
	// @DataMember
	MaxLength *int `json:"maxLength,omitempty"`
	// @DataMember
	TranslateOptions *IReadOnlyDictionary `json:"translateOptions,omitempty"`
}

type DecimalFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Minimum *float64 `json:"minimum,omitempty"`
	// @DataMember
	Maximum *float64 `json:"maximum,omitempty"`
	// @DataMember
	MultipleOf *float64 `json:"multipleOf,omitempty"`
}

type CurrencyFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	AllowedCurrencies *IReadOnlyList `json:"allowedCurrencies,omitempty"`
}

type BooleanFieldDto struct {
	JsonSchemaFieldDto
}

type DateFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Minimum *int64 `json:"minimum,omitempty"`
	// @DataMember
	Maximum *int64 `json:"maximum,omitempty"`
}

type IntegerFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Minimum *int64 `json:"minimum,omitempty"`
	// @DataMember
	Maximum *int64 `json:"maximum,omitempty"`
}

type GeolocationFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	AllowedTypes *IReadOnlyList `json:"allowedTypes,omitempty"`
}

type TagsFieldDto struct {
	JsonSchemaFieldDto
}

type FileFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Storages *IReadOnlyList `json:"storages,omitempty"`
}

type TaxonomySelectionFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	TaxonomyId *string `json:"taxonomyId,omitempty"`
	// @DataMember
	Multiple bool `json:"multiple,omitempty"`
}

type CollectionSelectionFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	CollectionId *string `json:"collectionId,omitempty"`
	// @DataMember
	DisplayField *string `json:"displayField,omitempty"`
	// @DataMember
	Multiple bool `json:"multiple,omitempty"`
}

type UserSelectionFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Multiple bool `json:"multiple,omitempty"`
}

type RoleSelectionFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Multiple bool `json:"multiple,omitempty"`
}

type EnumSelectionFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Values *IReadOnlyList `json:"values,omitempty"`
	// @DataMember
	Multiple bool `json:"multiple,omitempty"`
}

type EchoResponse struct {
	ContainerName                *string         `json:"containerName,omitempty"`
	Ip                           string          `json:"ip"`
	Release                      CodeMashRelease `json:"release,omitempty"`
	Runtime                      CodeMashRuntime `json:"runtime,omitempty"`
	ManagedServiceHubUrl         string          `json:"managedServiceHubUrl"`
	ManagedServiceApiUrl         string          `json:"managedServiceApiUrl"`
	HubUrl                       string          `json:"hubUrl"`
	ApiUrl                       string          `json:"apiUrl"`
	ApiVersion                   string          `json:"apiVersion"`
	HubVersion                   string          `json:"hubVersion"`
	MjmlUrl                      string          `json:"mjmlUrl"`
	AdminUrlTemplate             *string         `json:"adminUrlTemplate,omitempty"`
	License                      *EchoLicenseDto `json:"license,omitempty"`
	AskForEnterpriseLicenseEmail *string         `json:"askForEnterpriseLicenseEmail,omitempty"`
	EmailServiceConfigured       bool            `json:"emailServiceConfigured,omitempty"`
	RootBootstrapPasswordSource  *string         `json:"rootBootstrapPasswordSource,omitempty"`
	Regions                      []EchoRegionDto `json:"regions,omitempty"`
	IsProductionInstallation     bool            `json:"isProductionInstallation,omitempty"`
	LicensingMode                string          `json:"licensingMode"`
	GraceDaysLeft                *int            `json:"graceDaysLeft,omitempty"`
	InstallationDomain           *string         `json:"installationDomain,omitempty"`
	LicensingDocsUrl             *string         `json:"licensingDocsUrl,omitempty"`
}

type PublicProjectConfigDto struct {
	DisplayName        string          `json:"displayName"`
	AdminPortalEnabled bool            `json:"adminPortalEnabled,omitempty"`
	Branding           *PublicBrandDto `json:"branding,omitempty"`
	Auth               PublicAuthDto   `json:"auth"`
}

type PublicLegalDocumentDto struct {
	Kind      string  `json:"kind"`
	Title     *string `json:"title,omitempty"`
	Body      string  `json:"body"`
	Available bool    `json:"available,omitempty"`
}

type AskChatResponse struct {
	ResponseBase
	Result *string `json:"result,omitempty"`
}

type EmptyResponse struct {
	ResponseBase
}

// @DataContract
type IdResponse struct {
	ResponseBase
	// @DataMember
	Id *string `json:"id,omitempty"`
	// @DataMember
	Status *string `json:"status,omitempty"`
}

type GetUserResponse struct {
	ResponseBase
	User *AuthDto `json:"user,omitempty"`
}

type GetUsersResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetUserPreferencesResponse struct {
	ResponseBase
	Preferences *UserMarketingPreferencesDto `json:"preferences,omitempty"`
}

type PasskeyCeremonyOptionsResponse struct {
	ResponseBase
	CeremonyId  string `json:"ceremonyId"`
	OptionsJson string `json:"optionsJson"`
}

type PasskeyAuthTokensResponse struct {
	ResponseBase
	AccessToken      string   `json:"accessToken"`
	RefreshToken     string   `json:"refreshToken"`
	ExpiresInSeconds int      `json:"expiresInSeconds,omitempty"`
	RecoveryCodes    []string `json:"recoveryCodes,omitempty"`
}

type PasskeyListResponse struct {
	ResponseBase
	Passkeys []PasskeyListItemDto `json:"passkeys"`
}

type PasskeyOkResponse struct {
	ResponseBase
}

type PasskeyRecoveryResponse struct {
	ResponseBase
	AccessToken      string `json:"accessToken"`
	RefreshToken     string `json:"refreshToken"`
	ExpiresInSeconds int    `json:"expiresInSeconds,omitempty"`
	RemainingCodes   int    `json:"remainingCodes,omitempty"`
}

type PasskeyVerificationTokenResponse struct {
	ResponseBase
	VerificationToken string `json:"verificationToken"`
}

type FindMergedTermTreeResponse struct {
	ResponseBase
	Tree []TermTreeDto `json:"tree,omitempty"`
}

type FindTaxonomyTreeResponse struct {
	ResponseBase
	Tree []TaxonomyTreeDto `json:"tree,omitempty"`
}

type FindTermsResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type FindTermsChildrenResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type FindTermTreeResponse struct {
	ResponseBase
	Tree []TermTreeDto `json:"tree,omitempty"`
}

type GetDatabaseSchemaResponse struct {
	ResponseBase
	Item *SchemaDto `json:"item,omitempty"`
}

type GetDatabaseSchemasResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type AggregateResponse struct {
	ResponseBase
	Result []Object `json:"result,omitempty"`
}

type CountResponse struct {
	ResponseBase
	Count int64 `json:"count,omitempty"`
}

type DistinctResponse struct {
	ResponseBase
	Values []Object `json:"values,omitempty"`
}

type ExecuteAggregateResponse struct {
	ResponseBase
	Result []Object `json:"result,omitempty"`
}

type FindResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type FindOneResponse struct {
	ResponseBase
	Result *Object `json:"result,omitempty"`
}

type GetFileInfoResponse struct {
	ResponseBase
	File      *FileResourceRefDto `json:"file,omitempty"`
	IsPublic  *bool               `json:"isPublic,omitempty"`
	PublicUrl *string             `json:"publicUrl,omitempty"`
}

type GetSignedUrlResponse struct {
	ResponseBase
	Url *string `json:"url,omitempty"`
}

type ListFilesResponse struct {
	ResponseBase
	List    *PaginatedResponse `json:"list,omitempty"`
	Folders *IList             `json:"folders,omitempty"`
}

type RequestUploadUrlResponse struct {
	ResponseBase
	Url *string `json:"url,omitempty"`
}

// @DataContract
type AuthenticateResponse struct {
	// @DataMember(Order=1)
	UserId *string `json:"userId,omitempty"`
	// @DataMember(Order=2)
	SessionId *string `json:"sessionId,omitempty"`
	// @DataMember(Order=3)
	UserName *string `json:"userName,omitempty"`
	// @DataMember(Order=4)
	DisplayName *string `json:"displayName,omitempty"`
	// @DataMember(Order=5)
	ReferrerUrl *string `json:"referrerUrl,omitempty"`
	// @DataMember(Order=6)
	BearerToken *string `json:"bearerToken,omitempty"`
	// @DataMember(Order=7)
	RefreshToken *string `json:"refreshToken,omitempty"`
	// @DataMember(Order=8)
	RefreshTokenExpiry *time.Time `json:"refreshTokenExpiry,omitempty"`
	// @DataMember(Order=9)
	ProfileUrl *string `json:"profileUrl,omitempty"`
	// @DataMember(Order=10)
	Roles []string `json:"roles,omitempty"`
	// @DataMember(Order=11)
	Permissions []string `json:"permissions,omitempty"`
	// @DataMember(Order=12)
	AuthProvider *string `json:"authProvider,omitempty"`
	// @DataMember(Order=13)
	ResponseStatus *ResponseStatus `json:"responseStatus,omitempty"`
	// @DataMember(Order=14)
	Meta map[string]string `json:"meta,omitempty"`
}

// @DataContract
type GetAccessTokenResponse struct {
	// @DataMember(Order=1)
	AccessToken *string `json:"accessToken,omitempty"`
	// @DataMember(Order=2)
	Meta map[string]string `json:"meta,omitempty"`
	// @DataMember(Order=3)
	ResponseStatus *ResponseStatus `json:"responseStatus,omitempty"`
}

// @Route("/{version}/echo", "GET")
type Echo struct {
	RequestBase
}

// @Route("/{version}/public/projects/{ProjectId}/config", "GET")
type GetPublicProjectConfig struct {
	RequestBase
	ProjectId *string `json:"projectId,omitempty"`
}

// @Route("/{version}/public/projects/{ProjectId}/legal/{Kind}", "GET")
type GetPublicProjectLegal struct {
	RequestBase
	ProjectId *string `json:"projectId,omitempty"`
	Kind      *string `json:"kind,omitempty"`
}

type AccountCreated struct {
	Email       EmailAddress `json:"email"`
	DisplayName DisplayName  `json:"displayName"`
	AccountId   AccountId    `json:"accountId"`
	CreatedOn   UtcDateTime  `json:"createdOn"`
}

type AccountVerified struct {
}

type AccountSetAsActive struct {
}

type AccountValidationTokenIssued struct {
	Expiration ExpirationToken `json:"expiration"`
}

type AccountBlocked struct {
}

type AccountProfileUpdated struct {
	DisplayName     DisplayName   `json:"displayName"`
	BillingEmail    *EmailAddress `json:"billingEmail,omitempty"`
	OperationsEmail *EmailAddress `json:"operationsEmail,omitempty"`
	SecurityEmail   *EmailAddress `json:"securityEmail,omitempty"`
}

type AccountSetAsInactive struct {
}

type AccountUnregistered struct {
}

type LicenseCreated struct {
	License CodeMashLicense `json:"license"`
}

type CustomerCreated struct {
	PaymentCustomerRef PaymentCustomerRef `json:"paymentCustomerRef"`
}

type SubscriptionChanged struct {
	Subscription CodeMashManagedServiceSubscription `json:"subscription"`
}

type SubscriptionCanceled struct {
	PaymentCustomerRef PaymentCustomerRef `json:"paymentCustomerRef"`
	SubscriptionId     string             `json:"subscriptionId"`
}

type ProjectCommunicationGroupSaved struct {
	Group         GroupDefinition       `json:"group"`
	Channel       CommunicationChannel  `json:"channel,omitempty"`
	OriginChannel *CommunicationChannel `json:"originChannel,omitempty"`
}

type ProjectCommunicationTagFromGroupDeleted struct {
	GroupTag   Tag `json:"groupTag"`
	RemovedTag Tag `json:"removedTag"`
}

type ProjectCommunicationGroupDeleted struct {
	GroupTag Tag `json:"groupTag"`
}

type ProjectCommunicationTagSaved struct {
	Tag      TagDefinition         `json:"tag"`
	GroupTag *Tag                  `json:"groupTag,omitempty"`
	Channel  *CommunicationChannel `json:"channel,omitempty"`
}

type ProjectCommunicationTagDeleted struct {
	Tag Tag `json:"tag"`
}

type ProjectCreated struct {
	Id                    ProjectId       `json:"id"`
	Name                  ProjectName     `json:"name"`
	DatabaseIntegrationId IntegrationId   `json:"databaseIntegrationId"`
	PrimaryRegion         *ProjectRegion  `json:"primaryRegion,omitempty"`
	AdditionalRegions     []ProjectRegion `json:"additionalRegions,omitempty"`
	Description           *string         `json:"description,omitempty"`
	IsProvisioning        bool            `json:"isProvisioning,omitempty"`
}

type ProjectDeleted struct {
}

type ProjectActivated struct {
}

type ProjectDisabled struct {
}

type ProjectNameChanged struct {
	ProjectName ProjectName `json:"projectName"`
}

type ProjectDescriptionChanged struct {
	Description *string `json:"description,omitempty"`
}

type ProjectMarketingUrlChanged struct {
	Url *DomainUrl `json:"url,omitempty"`
}

type ProjectAllowedOriginsChanged struct {
	Origins []DomainUrl `json:"origins,omitempty"`
}

type ProjectDefaultLanguageChanged struct {
	Language Language `json:"language"`
}

type ProjectLanguagesChanged struct {
	Languages []Language `json:"languages"`
}

type ProjectLogoChanged struct {
	Logo *ProjectLogo `json:"logo,omitempty"`
}

type ProjectIconChanged struct {
	Icon *ProjectIcon `json:"icon,omitempty"`
}

type ProjectMainColorChanged struct {
	Color BrandColor `json:"color"`
}

type ProjectAccentColorChanged struct {
	Color BrandColor `json:"color"`
}

type ProjectRegionsChanged struct {
	PrimaryRegion     *ProjectRegion  `json:"primaryRegion,omitempty"`
	AdditionalRegions []ProjectRegion `json:"additionalRegions,omitempty"`
}

type ProjectTimeZoneChanged struct {
	TimeZone *TimeZone `json:"timeZone,omitempty"`
}

type ProjectPaymentZonesChanged struct {
	PaymentZones []TimeZone `json:"paymentZones,omitempty"`
}

type ProjectCommunicationSet struct {
	ProjectCommunication ProjectCommunication `json:"projectCommunication"`
}

type AccountUserPushDeviceCreated struct {
	AuthId     AuthId     `json:"authId"`
	PushDevice PushDevice `json:"pushDevice"`
}

/** @description AI */
// @Route("/{version}/chat/complete", "POST")
// @Api(Description="AI")
// @DataContract
type AskChatRequest struct {
	CodeMashRequestBase
	// @DataMember
	Prompt string `json:"prompt"`
}

/** @description Membership */
// @Route("/{version}/membership/auth/block", "PATCH")
// @Api(Description="Membership")
// @DataContract
type BlockUserRequest struct {
	CodeMashRequestBase
	/** @description Id of the user to block, from get_users. */
	// @DataMember
	// @ApiMember(Description="Id of the user to block, from get_users.", IsRequired=true)
	Id string `json:"id"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Membership */
// @Route("/{version}/membership/auth/register/service", "POST")
// @Api(Description="Membership")
// @DataContract
type SaveSystemUserWithPermissions struct {
	SaveUserWithRolesBase
}

/** @description Membership */
// @Route("/{version}/membership/auth/register/guest", "POST")
// @Api(Description="Membership")
// @DataContract
type SaveGuestUser struct {
	SaveUser
}

/** @description Membership */
// @Route("/{version}/membership/auth/register/user-name", "POST")
// @Api(Description="Membership")
// @DataContract
type SaveUserNameUser struct {
	SaveUser
	// @DataMember
	Password string `json:"password"`
	// @DataMember
	UserName string `json:"userName"`
}

/** @description Membership */
// @Route("/{version}/membership/auth/register/email", "POST")
// @Api(Description="Membership")
// @DataContract
type SaveEmailUser struct {
	SaveUser
	// @DataMember
	Password string `json:"password"`
	// @DataMember
	Email string `json:"email"`
}

/** @description Membership */
// @Route("/{version}/membership/auth/register/phone", "POST")
// @Api(Description="Membership")
// @DataContract
type SavePhoneUser struct {
	SaveUser
	/** @description Phone number for the new user, in E.164 format. */
	// @DataMember
	// @ApiMember(Description="Phone number for the new user, in E.164 format.", IsRequired=true)
	Phone string `json:"phone"`
}

/** @description Membership */
// @Route("/{version}/membership/auth/register/phone-with-permissions", "POST")
// @Api(Description="Membership")
// @DataContract
type SavePhoneUserNameWithPermissions struct {
	SaveUserWithRolesBase
	// @DataMember
	Phone string `json:"phone"`
}

/** @description Membership */
// @Route("/{version}/membership/auth/register/email-with-permissions", "POST")
// @Api(Description="Membership")
// @DataContract
type SaveEmailUserNameWithPermissions struct {
	SaveUserWithRolesBase
	// @DataMember
	Password string `json:"password"`
	// @DataMember
	Email string `json:"email"`
}

/** @description Membership */
// @Route("/{version}/membership/auth/register/user-name-with-permissions", "POST")
// @Api(Description="Membership")
// @DataContract
type SaveUserNameWithPermissions struct {
	SaveUserWithRolesBase
	// @DataMember
	Password string `json:"password"`
	// @DataMember
	UserName string `json:"userName"`
}

/** @description Membership */
// @Route("/{version}/membership/auth", "DELETE")
// @Api(Description="Membership")
// @DataContract
type DeleteUserRequest struct {
	CodeMashRequestBase
	/** @description Id of the user to delete, from get_users. */
	// @DataMember
	// @ApiMember(Description="Id of the user to delete, from get_users.", IsRequired=true)
	Id string `json:"id"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Membership */
// @Route("/{version}/membership/auth/{id}", "GET")
// @Api(Description="Membership")
// @DataContract
type GetUserRequest struct {
	CodeMashRequestBase
	/** @description Id of the user to fetch, from get_users. */
	// @DataMember
	// @ApiMember(Description="Id of the user to fetch, from get_users.", IsRequired=true)
	Id string `json:"id"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Membership */
// @Route("/{version}/membership/auth", "GET")
// @Api(Description="Membership")
// @DataContract
type GetUsersRequest struct {
	CodeMashListPaginationRequestBase
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Include each user's effective permissions in the result. */
	// @DataMember
	// @ApiMember(Description="Include each user's effective permissions in the result.")
	IncludePermissions bool `json:"includePermissions,omitempty"`
	/** @description Only return users that have a registered push device. */
	// @DataMember
	// @ApiMember(Description="Only return users that have a registered push device.")
	UserShouldHavePushDevice bool `json:"userShouldHavePushDevice,omitempty"`
	/** @description Only return users that have an email address. */
	// @DataMember
	// @ApiMember(Description="Only return users that have an email address.")
	UserShouldHaveEmail bool `json:"userShouldHaveEmail,omitempty"`
	/** @description Include each user's metadata in the result. */
	// @DataMember
	// @ApiMember(Description="Include each user's metadata in the result.")
	IncludeMeta bool `json:"includeMeta,omitempty"`
	/** @description Filter to users that have any of these role names. */
	// @DataMember
	// @ApiMember(Description="Filter to users that have any of these role names.")
	RoleNames []string `json:"roleNames,omitempty"`
	/** @description Filter to these specific user ids. */
	// @DataMember
	// @ApiMember(Description="Filter to these specific user ids.")
	UserIds []string `json:"userIds,omitempty"`
}

/** @description Membership */
// @Route("/{version}/membership/auth/{id}/preferences", "GET")
// @Api(Description="Membership")
// @DataContract
type GetUserPreferencesRequest struct {
	CodeMashRequestBase
	/** @description Id of the user whose preferences to fetch, from get_users. */
	// @DataMember
	// @ApiMember(Description="Id of the user whose preferences to fetch, from get_users.", IsRequired=true)
	Id string `json:"id"`
	/** @description Database integration id. Optional — defaults to the project's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the project's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

// @Route("/{version}/membership/users/{contactId}/marketing-state/{channel}/consent", "POST")
type GrantContactConsentRequest struct {
	CodeMashRequestBase
	/** @description Id of the user (contact) to grant consent for. */
	// @ApiMember(Description="Id of the user (contact) to grant consent for.", IsRequired=true)
	ContactId string `json:"contactId"`
	/** @description Delivery channel to grant consent on: Email, Sms, or Push. */
	// @ApiMember(Description="Delivery channel to grant consent on: Email, Sms, or Push.", IsRequired=true)
	Channel string `json:"channel"`
	/** @description Lawful basis for the consent, e.g. Consent. Defaults to Consent. */
	// @ApiMember(Description="Lawful basis for the consent, e.g. Consent. Defaults to Consent.")
	LawfulBasis string `json:"lawfulBasis"`
	/** @description Source of the consent, e.g. UserOptIn. Defaults to UserOptIn. */
	// @ApiMember(Description="Source of the consent, e.g. UserOptIn. Defaults to UserOptIn.")
	Source string `json:"source"`
	/** @description Optional free-text reference to evidence of consent (e.g. a form submission id). */
	// @ApiMember(Description="Optional free-text reference to evidence of consent (e.g. a form submission id).")
	EvidenceRef *string `json:"evidenceRef,omitempty"`
}

/** @description Membership */
// @Route("/{version}/membership/auth/invite", "POST")
// @Api(Description="Membership")
// @DataContract
type InviteUserRequest struct {
	CodeMashRequestBase
	/** @description Email address the invitation is sent to. */
	// @DataMember
	// @ApiMember(Description="Email address the invitation is sent to.", IsRequired=true)
	Email string `json:"email"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Membership */
// @Route("/{version}/membership/auth/{userId}/link-identity", "POST")
// @Api(Description="Membership")
// @DataContract
type LinkIdentityRequest struct {
	CodeMashRequestBase
	// @DataMember
	UserId string `json:"userId"`
	// @DataMember
	Provider string `json:"provider"`
	// @DataMember
	ProviderToken string `json:"providerToken"`
	// @DataMember
	EmailToVerify *string `json:"emailToVerify,omitempty"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Membership */
// @Route("/{version}/membership/users/{userId}/map-auth", "POST")
// @Api(Description="Membership")
type MapAuthToUserRequest struct {
	CodeMashRequestBase
	UserId                string  `json:"userId"`
	AuthId                string  `json:"authId"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Membership */
// @Route("/{version}/membership/auth/assign-roles", "PUT")
// @Api(Description="Membership")
// @DataContract
type AssignRolePermissionsRequest struct {
	CodeMashRequestBase
	/** @description Id of the user login to assign roles to, from get_users. */
	// @DataMember
	// @ApiMember(Description="Id of the user login to assign roles to, from get_users.", IsRequired=true)
	Id string `json:"id"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description The complete new list of role names (full replacement), from get_roles. */
	// @DataMember
	// @ApiMember(Description="The complete new list of role names (full replacement), from get_roles.")
	Roles []string `json:"roles,omitempty"`
}

/** @description Membership */
// @Route("/{version}/membership/users/{userId}/roles", "PUT")
// @Api(Description="Membership")
// @DataContract
type SetContactRolesRequest struct {
	CodeMashRequestBase
	/** @description Id of the human user to assign roles to. */
	// @DataMember
	// @ApiMember(Description="Id of the human user to assign roles to.", IsRequired=true)
	UserId string `json:"userId"`
	/** @description The complete new list of role ids (full replacement), from get_roles. Empty/omitted clears all roles. */
	// @DataMember
	// @ApiMember(Description="The complete new list of role ids (full replacement), from get_roles. Empty/omitted clears all roles.")
	Roles []string `json:"roles,omitempty"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

// @Route("/{version}/membership/users/{contactId}/marketing-state/{commChannel}/{channel}/tags/{tag}", "PUT")
type SetContactTagSubscriptionRequest struct {
	CodeMashRequestBase
	/** @description Id of the user (contact) to update. */
	// @ApiMember(Description="Id of the user (contact) to update.", IsRequired=true)
	ContactId string `json:"contactId"`
	/** @description Communication channel type: Marketing or Transactional. */
	// @ApiMember(Description="Communication channel type: Marketing or Transactional.", IsRequired=true)
	CommChannel string `json:"commChannel"`
	/** @description Delivery channel: Email, Sms, or Push. */
	// @ApiMember(Description="Delivery channel: Email, Sms, or Push.", IsRequired=true)
	Channel string `json:"channel"`
	/** @description The tag name; must already exist for the communication channel. */
	// @ApiMember(Description="The tag name; must already exist for the communication channel.", IsRequired=true)
	Tag string `json:"tag"`
	/** @description True to subscribe (unblock) the tag, false to block it. */
	// @ApiMember(Description="True to subscribe (unblock) the tag, false to block it.")
	Subscribed bool `json:"subscribed,omitempty"`
}

/** @description Membership */
// @Route("/{version}/membership/auth/unblock", "PATCH")
// @Api(Description="Membership")
// @DataContract
type UnblockUserRequest struct {
	CodeMashRequestBase
	/** @description Id of the user to unblock, from get_users. */
	// @DataMember
	// @ApiMember(Description="Id of the user to unblock, from get_users.", IsRequired=true)
	Id string `json:"id"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

// @Route("/{version}/membership/users/{contactId}/marketing-state/{channel}/unsubscribe", "POST")
type UnsubscribeContactRequest struct {
	CodeMashRequestBase
	/** @description Id of the user (contact) to unsubscribe. */
	// @ApiMember(Description="Id of the user (contact) to unsubscribe.", IsRequired=true)
	ContactId string `json:"contactId"`
	/** @description Delivery channel to unsubscribe from: Email, Sms, or Push. */
	// @ApiMember(Description="Delivery channel to unsubscribe from: Email, Sms, or Push.", IsRequired=true)
	Channel string `json:"channel"`
	/** @description Optional suppression reason name explaining why consent was revoked. */
	// @ApiMember(Description="Optional suppression reason name explaining why consent was revoked.")
	Reason *string `json:"reason,omitempty"`
}

/** @description Membership */
// @Route("/{version}/membership/auth", "PUT")
// @Api(Description="Membership")
// @DataContract
type UpdateUserRequest struct {
	SaveUser
	/** @description Id of the user to update, from get_users. */
	// @DataMember
	// @ApiMember(Description="Id of the user to update, from get_users.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Membership */
// @Route("/{version}/membership/auth/{id}/preferences", "PUT")
// @Api(Description="Membership")
// @DataContract
type UpdateUserPreferencesRequest struct {
	CodeMashRequestBase
	/** @description Id of the user to update, from get_users. */
	// @DataMember
	// @ApiMember(Description="Id of the user to update, from get_users.", IsRequired=true)
	Id string `json:"id"`
	/** @description When true, blocks all marketing messages to this user. */
	// @DataMember
	// @ApiMember(Description="When true, blocks all marketing messages to this user.")
	BlockAllMarketingMessages bool `json:"blockAllMarketingMessages,omitempty"`
	/** @description Per communication channel, the set of tags blocked for this user. Full replacement. */
	// @DataMember
	// @ApiMember(Description="Per communication channel, the set of tags blocked for this user. Full replacement.")
	BlockedTags map[string][]string `json:"blockedTags,omitempty"`
	/** @description Database integration id. Optional — defaults to the project's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the project's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/passkey/authentication-options", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type PasskeyAuthenticationOptionsRequest struct {
	CodeMashRequestBase
	// @DataMember
	Email string `json:"email"`
}

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/passkey/verify-authentication", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type VerifyPasskeyAuthenticationRequest struct {
	CodeMashRequestBase
	// @DataMember
	CeremonyId string `json:"ceremonyId"`
	// @DataMember
	AssertionResponse string `json:"assertionResponse"`
}

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/passkeys", "GET")
// @Api(Description="Membership · Passkey")
// @DataContract
type ListPasskeysRequest struct {
	CodeMashRequestBase
}

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/passkeys/{CredentialId}/rename", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type RenamePasskeyRequest struct {
	CodeMashRequestBase
	/** @description Base64 credential id of the passkey to rename, from list_passkeys. */
	// @DataMember
	// @ApiMember(Description="Base64 credential id of the passkey to rename, from list_passkeys.", IsRequired=true)
	CredentialId string `json:"credentialId"`
	/** @description The new friendly name for the passkey. */
	// @DataMember
	// @ApiMember(Description="The new friendly name for the passkey.", IsRequired=true)
	FriendlyName string `json:"friendlyName"`
}

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/passkeys/{CredentialId}/revoke", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type RevokePasskeyRequest struct {
	CodeMashRequestBase
	/** @description Base64 credential id of the passkey to revoke, from list_passkeys. */
	// @DataMember
	// @ApiMember(Description="Base64 credential id of the passkey to revoke, from list_passkeys.", IsRequired=true)
	CredentialId string `json:"credentialId"`
}

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/recovery/use-code", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type UseRecoveryCodeRequest struct {
	CodeMashRequestBase
	// @DataMember
	Email string `json:"email"`
	// @DataMember
	RecoveryCode string `json:"recoveryCode"`
}

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/recovery/magic-link/request", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type RequestMagicLinkRequest struct {
	CodeMashRequestBase
	// @DataMember
	Email string `json:"email"`
}

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/recovery/magic-link/consume", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type ConsumeMagicLinkRequest struct {
	CodeMashRequestBase
	// @DataMember
	Token string `json:"token"`
}

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/has-passkey", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type HasPasskeyRequest struct {
	CodeMashRequestBase
	// @DataMember
	Email string `json:"email"`
}

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/email/start-verification", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type StartEmailVerificationRequest struct {
	CodeMashRequestBase
	// @DataMember
	Email string `json:"email"`
}

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/email/confirm-verification", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type ConfirmEmailVerificationRequest struct {
	CodeMashRequestBase
	// @DataMember
	Email string `json:"email"`
	// @DataMember
	Code string `json:"code"`
}

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/passkey/registration-options", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type PasskeyRegistrationOptionsRequest struct {
	CodeMashRequestBase
	// @DataMember
	VerificationToken string `json:"verificationToken"`
}

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/passkey/verify-registration", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type VerifyPasskeyRegistrationRequest struct {
	CodeMashRequestBase
	// @DataMember
	VerificationToken string `json:"verificationToken"`
	// @DataMember
	CeremonyId string `json:"ceremonyId"`
	// @DataMember
	AttestationResponse string `json:"attestationResponse"`
	// @DataMember
	FriendlyName *string `json:"friendlyName,omitempty"`
}

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/token/refresh", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type RefreshPasskeyTokenRequest struct {
	CodeMashRequestBase
	// @DataMember
	RefreshToken *string `json:"refreshToken,omitempty"`
}

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/logout", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type PasskeyLogoutRequest struct {
	CodeMashRequestBase
	// @DataMember
	RefreshToken *string `json:"refreshToken,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/taxonomies/{taxonomyName}/merged-tree", "GET")
// @Api(Description="Database")
// @DataContract
type FindMergedTermTreeRequest struct {
	CodeMashRequestBase
	// @DataMember
	TaxonomyName string `json:"taxonomyName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/taxonomies/tree", "GET")
// @Api(Description="Database")
// @DataContract
type FindTaxonomyTreeRequest struct {
	CodeMashRequestBase
	// @DataMember
	IncludeTerms bool `json:"includeTerms,omitempty"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/taxonomies/{taxonomyName}/terms", "GET")
// @Api(Description="Database")
// @DataContract
type FindTermsRequest struct {
	CodeMashListPaginationRequestBase
	// @DataMember
	TaxonomyName string `json:"taxonomyName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Filter *string `json:"filter,omitempty"`
	// @DataMember
	SortDescending bool `json:"sortDescending,omitempty"`
	// @DataMember
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/taxonomies/{taxonomyName}/terms/{parentId}/children", "GET")
// @Api(Description="Database")
// @DataContract
type FindTermsChildrenRequest struct {
	CodeMashListPaginationRequestBase
	// @DataMember
	TaxonomyName string `json:"taxonomyName"`
	// @DataMember
	ParentId string `json:"parentId"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Filter *string `json:"filter,omitempty"`
	// @DataMember
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/taxonomies/{taxonomyName}/terms/tree", "GET")
// @Api(Description="Database")
// @DataContract
type FindTermTreeRequest struct {
	CodeMashRequestBase
	// @DataMember
	TaxonomyName string `json:"taxonomyName"`
	// @DataMember
	RootTermId *string `json:"rootTermId,omitempty"`
	// @DataMember
	Depth *int `json:"depth,omitempty"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/schemas/{id}", "GET")
// @Api(Description="Database")
// @DataContract
type GetDatabaseSchemaRequest struct {
	CodeMashRequestBase
	// @DataMember
	Id string `json:"id"`
}

/** @description Database */
// @Route("/{version}/database/schemas", "GET")
// @Api(Description="Database")
// @DataContract
type GetDatabaseSchemasRequest struct {
	CodeMashListPaginationRequestBase
	// @DataMember
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/aggregate", "POST")
// @Api(Description="Database")
// @DataContract
type AggregateRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Pipeline string `json:"pipeline"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/{id}/responsibility", "PUT")
// @Api(Description="Database")
// @DataContract
type ChangeResponsibilityRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	NewResponsibleUserId string `json:"newResponsibleUserId"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/count", "GET")
// @Api(Description="Database")
// @DataContract
type CountRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Filter *string `json:"filter,omitempty"`
	// @DataMember
	SchemaVersion *int `json:"schemaVersion,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/many", "DELETE")
// @Api(Description="Database")
// @DataContract
type DeleteManyRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Filter string `json:"filter"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/{id}", "DELETE")
// @Api(Description="Database")
// @DataContract
type DeleteOneRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/distinct", "GET")
// @Api(Description="Database")
// @DataContract
type DistinctRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Field string `json:"field"`
	// @DataMember
	Filter *string `json:"filter,omitempty"`
	// @DataMember
	SchemaVersion *int `json:"schemaVersion,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/aggregates/{aggregateId}/execute", "POST")
// @Api(Description="Database")
// @DataContract
type ExecuteAggregateRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	AggregateId string `json:"aggregateId"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Tokens map[string]string `json:"tokens,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}", "GET")
// @Api(Description="Database")
// @DataContract
type FindRequest struct {
	CodeMashListPaginationRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Filter *string `json:"filter,omitempty"`
	// @DataMember
	SchemaVersion *int `json:"schemaVersion,omitempty"`
	// @DataMember
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
	// @DataMember
	SortBy *string `json:"sortBy,omitempty"`
	// @DataMember
	SortOrder *int `json:"sortOrder,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/{id}", "GET")
// @Api(Description="Database")
// @DataContract
type FindOneRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/own", "GET")
// @Api(Description="Database")
// @DataContract
type FindOwnRequest struct {
	CodeMashListPaginationRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Filter *string `json:"filter,omitempty"`
	// @DataMember
	SchemaVersion *int `json:"schemaVersion,omitempty"`
	// @DataMember
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/many", "POST")
// @Api(Description="Database")
// @DataContract
type InsertManyRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Documents string `json:"documents"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}", "POST")
// @Api(Description="Database")
// @DataContract
type InsertOneRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Document string `json:"document"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/{id}/replace", "PUT")
// @Api(Description="Database")
// @DataContract
type ReplaceOneRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Replacement string `json:"replacement"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/many", "PUT")
// @Api(Description="Database")
// @DataContract
type UpdateManyRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Filter string `json:"filter"`
	// @DataMember
	Update string `json:"update"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/{id}", "PUT")
// @Api(Description="Database")
// @DataContract
type UpdateOneRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Update string `json:"update"`
}

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/commit", "POST")
// @Api(Description="Files")
// @DataContract
type CommitUploadRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path string `json:"path"`
	// @DataMember
	ContentType *string `json:"contentType,omitempty"`
	// @DataMember
	SizeBytes *int64 `json:"sizeBytes,omitempty"`
	// @DataMember
	FileName *string `json:"fileName,omitempty"`
}

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}", "DELETE")
// @Api(Description="Files")
// @DataContract
type DeleteFileApiRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path string `json:"path"`
}

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/bulk", "DELETE")
// @Api(Description="Files")
// @DataContract
type DeleteManyFilesApiRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember(Name="paths[]")
	Paths []string `json:"paths__"`
}

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/download", "GET")
// @Api(Description="Files")
// @DataContract
type DownloadFileApiRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path string `json:"path"`
}

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/info", "GET")
// @Api(Description="Files")
// @DataContract
type GetFileInfoRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path string `json:"path"`
}

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/sign", "GET")
// @Api(Description="Files")
// @DataContract
type GetSignedUrlRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path string `json:"path"`
	// @DataMember
	ExpirationSeconds *int `json:"expirationSeconds,omitempty"`
}

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}", "GET")
// @Api(Description="Files")
// @DataContract
type ListFilesRequest struct {
	CodeMashListPaginationRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path *string `json:"path,omitempty"`
}

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/upload-url", "POST")
// @Api(Description="Files")
// @DataContract
type RequestUploadUrlRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path string `json:"path"`
	// @DataMember
	ContentType string `json:"contentType"`
	// @DataMember
	ExpirationSeconds *int `json:"expirationSeconds,omitempty"`
}

type PushIntegrationSaved struct {
	Integration PushIntegration `json:"integration"`
}

type PushIntegrationRenamed struct {
	Id   IntegrationId `json:"id"`
	Name DisplayName   `json:"name"`
	Env  *Env          `json:"env,omitempty"`
}

type PushIntegrationSetAsDefault struct {
	Env Env           `json:"env"`
	Id  IntegrationId `json:"id"`
}

type PushIntegrationDeleted struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type PushIntegrationEnabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type PushIntegrationDisabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type PushServiceEstablished struct {
	DefaultTemplates []PushTemplate `json:"defaultTemplates,omitempty"`
}

type PushServiceEnabled struct {
}

type PushServiceDisabled struct {
}

type PushTemplateCreated struct {
	TemplateId   TemplateId           `json:"templateId"`
	DisplayName  DisplayName          `json:"displayName"`
	Translations []MessageTranslation `json:"translations"`
	Channel      CommunicationChannel `json:"channel,omitempty"`
	Description  *string              `json:"description,omitempty"`
	Tags         []Tag                `json:"tags,omitempty"`
	Env          *Env                 `json:"env,omitempty"`
}

type PushTemplateUpdated struct {
	TemplateId   TemplateId           `json:"templateId"`
	DisplayName  DisplayName          `json:"displayName"`
	Translations []MessageTranslation `json:"translations"`
	Channel      CommunicationChannel `json:"channel,omitempty"`
	Description  *string              `json:"description,omitempty"`
	Tags         []Tag                `json:"tags,omitempty"`
	Env          *Env                 `json:"env,omitempty"`
}

type PushTemplateDeleted struct {
	TemplateId TemplateId `json:"templateId"`
	Env        *Env       `json:"env,omitempty"`
}

type PushTemplateArchived struct {
	TemplateId TemplateId `json:"templateId"`
	Env        *Env       `json:"env,omitempty"`
}

type PushTemplateUnArchived struct {
	TemplateId TemplateId `json:"templateId"`
	Env        *Env       `json:"env,omitempty"`
}

type PushTemplateMirrored struct {
	Template PushTemplate `json:"template"`
}

/** @description Sign In */
// @Route("/auth", "GET,POST")
// @Route("/auth/{provider}", "GET,POST")
// @Route("/v3/auth", "POST,GET,OPTIONS")
// @Route("/v3/auth/{provider}", "POST,GET,OPTIONS")
// @Api(Description="Sign In")
// @DataContract
type Authenticate struct {
	/** @description AuthProvider, e.g. credentials */
	// @DataMember(Order=1)
	Provider *string `json:"provider,omitempty"`
	// @DataMember(Order=2)
	UserName *string `json:"userName,omitempty"`
	// @DataMember(Order=3)
	Password *string `json:"password,omitempty"`
	// @DataMember(Order=4)
	RememberMe *bool `json:"rememberMe,omitempty"`
	// @DataMember(Order=5)
	AccessToken *string `json:"accessToken,omitempty"`
	// @DataMember(Order=6)
	AccessTokenSecret *string `json:"accessTokenSecret,omitempty"`
	// @DataMember(Order=7)
	ReturnUrl *string `json:"returnUrl,omitempty"`
	// @DataMember(Order=8)
	ErrorView *string `json:"errorView,omitempty"`
	// @DataMember(Order=9)
	Meta map[string]string `json:"meta,omitempty"`
}

// @Route("/access-token")
// @DataContract
type GetAccessToken struct {
	// @DataMember(Order=1)
	RefreshToken *string `json:"refreshToken,omitempty"`
	// @DataMember(Order=2)
	Meta map[string]string `json:"meta,omitempty"`
}
