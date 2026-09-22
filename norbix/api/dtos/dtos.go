// Code generated from the Norbix API contract. DO NOT EDIT.
//
// Regenerate with the type generator described in README.md ("Regenerate the
// types"). Editing this file by hand is lost on the next run.

// Package dtos holds the data types of the Norbix API surface. A field
// that holds another generated type is a pointer; every field carries a JSON
// tag with omitempty, because the gateway leaves absent values out.
package dtos

// TimeUnit enum.
type TimeUnit string

const (
	TimeUnitTicks        TimeUnit = "Ticks"
	TimeUnitMilliseconds TimeUnit = "Milliseconds"
	TimeUnitSeconds      TimeUnit = "Seconds"
	TimeUnitMinutes      TimeUnit = "Minutes"
	TimeUnitHours        TimeUnit = "Hours"
)

// ResourceRefKind enum.
type ResourceRefKind string

const (
	ResourceRefKindContact         ResourceRefKind = "Contact"
	ResourceRefKindDocument        ResourceRefKind = "Document"
	ResourceRefKindFile            ResourceRefKind = "File"
	ResourceRefKindPaymentCustomer ResourceRefKind = "PaymentCustomer"
	ResourceRefKindOrder           ResourceRefKind = "Order"
	ResourceRefKindPayment         ResourceRefKind = "Payment"
	ResourceRefKindProduct         ResourceRefKind = "Product"
	ResourceRefKindIntegration     ResourceRefKind = "Integration"
)

// ResourceSource enum.
type ResourceSource string

const (
	ResourceSourceNorbix       ResourceSource = "Norbix"
	ResourceSourceStripe       ResourceSource = "Stripe"
	ResourceSourceShopify      ResourceSource = "Shopify"
	ResourceSourcePayPal       ResourceSource = "PayPal"
	ResourceSourceAdyen        ResourceSource = "Adyen"
	ResourceSourceMollie       ResourceSource = "Mollie"
	ResourceSourcePaddle       ResourceSource = "Paddle"
	ResourceSourceLemonSqueezy ResourceSource = "LemonSqueezy"
	ResourceSourceAppleInApp   ResourceSource = "AppleInApp"
	ResourceSourceGoogleInApp  ResourceSource = "GoogleInApp"
	ResourceSourceAuthorizeNet ResourceSource = "AuthorizeNet"
	ResourceSourceBraintree    ResourceSource = "Braintree"
	ResourceSourceCheckOutCom  ResourceSource = "CheckOutCom"
	ResourceSourceWooCommerce  ResourceSource = "WooCommerce"
	ResourceSourceMagento      ResourceSource = "Magento"
	ResourceSourceWorldpay     ResourceSource = "Worldpay"
)

// CommunicationChannel enum.
type CommunicationChannel string

const (
	CommunicationChannelTransactional CommunicationChannel = "Transactional"
	CommunicationChannelMarketing     CommunicationChannel = "Marketing"
	CommunicationChannelSystem        CommunicationChannel = "System"
)

// DeliveryChannel enum.
type DeliveryChannel string

const (
	DeliveryChannelEmail        DeliveryChannel = "Email"
	DeliveryChannelPush         DeliveryChannel = "Push"
	DeliveryChannelSms          DeliveryChannel = "Sms"
	DeliveryChannelWebPush      DeliveryChannel = "WebPush"
	DeliveryChannelInApp        DeliveryChannel = "InApp"
	DeliveryChannelChatBot      DeliveryChannel = "ChatBot"
	DeliveryChannelChatPlatform DeliveryChannel = "ChatPlatform"
)

// Continent enum.
type Continent string

const (
	ContinentAfrica       Continent = "Africa"
	ContinentAntarctica   Continent = "Antarctica"
	ContinentAsia         Continent = "Asia"
	ContinentEurope       Continent = "Europe"
	ContinentNorthAmerica Continent = "NorthAmerica"
	ContinentOceania      Continent = "Oceania"
	ContinentSouthAmerica Continent = "SouthAmerica"
)

// FileProvider enum.
type FileProvider string

const (
	FileProviderLocal              FileProvider = "Local"
	FileProviderAwsS3              FileProvider = "AwsS3"
	FileProviderAzureBlobStorage   FileProvider = "AzureBlobStorage"
	FileProviderGoogleCloudStorage FileProvider = "GoogleCloudStorage"
	FileProviderFtp                FileProvider = "Ftp"
	FileProviderAppleICloud        FileProvider = "AppleICloud"
	FileProviderDropBox            FileProvider = "DropBox"
	FileProviderGoogleDrive        FileProvider = "GoogleDrive"
)

// DeviceType enum.
type DeviceType string

const (
	DeviceTypeUnknown DeviceType = "Unknown"
	DeviceTypePhone   DeviceType = "Phone"
	DeviceTypeTablet  DeviceType = "Tablet"
	DeviceTypeDesktop DeviceType = "Desktop"
	DeviceTypeTv      DeviceType = "Tv"
)

// PushDeviceDeliveryFamily enum.
type PushDeviceDeliveryFamily string

const (
	PushDeviceDeliveryFamilyIos     PushDeviceDeliveryFamily = "Ios"
	PushDeviceDeliveryFamilyAndroid PushDeviceDeliveryFamily = "Android"
	PushDeviceDeliveryFamilyChrome  PushDeviceDeliveryFamily = "Chrome"
	PushDeviceDeliveryFamilySafari  PushDeviceDeliveryFamily = "Safari"
	PushDeviceDeliveryFamilyExpo    PushDeviceDeliveryFamily = "Expo"
)

// Gender enum.
type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"
	GenderOther  Gender = "Other"
)

// MarketingBlockReason enum.
type MarketingBlockReason string

const (
	MarketingBlockReasonUnspecified  MarketingBlockReason = "Unspecified"
	MarketingBlockReasonUnsubscribed MarketingBlockReason = "Unsubscribed"
	MarketingBlockReasonComplaint    MarketingBlockReason = "Complaint"
	MarketingBlockReasonHardBounce   MarketingBlockReason = "HardBounce"
	MarketingBlockReasonInvalidEmail MarketingBlockReason = "InvalidEmail"
	MarketingBlockReasonAdminBlock   MarketingBlockReason = "AdminBlock"
)

// PushProvider enum.
type PushProvider string

const (
	PushProviderAppleApns            PushProvider = "AppleApns"
	PushProviderSafariWeb            PushProvider = "SafariWeb"
	PushProviderSafariPush           PushProvider = "SafariPush"
	PushProviderAndroidFirebase      PushProvider = "AndroidFirebase"
	PushProviderChromeWeb            PushProvider = "ChromeWeb"
	PushProviderFirefoxWeb           PushProvider = "FirefoxWeb"
	PushProviderEdgeWeb              PushProvider = "EdgeWeb"
	PushProviderChromePush           PushProvider = "ChromePush"
	PushProviderCodeMashIosApp       PushProvider = "CodeMashIosApp"
	PushProviderCodeMashAndroidApp   PushProvider = "CodeMashAndroidApp"
	PushProviderCodeMashSafariPlugin PushProvider = "CodeMashSafariPlugin"
	PushProviderCodeMashSafariWeb    PushProvider = "CodeMashSafariWeb"
	PushProviderCodeMashChromePlugin PushProvider = "CodeMashChromePlugin"
	PushProviderCodeMashChromeWeb    PushProvider = "CodeMashChromeWeb"
	PushProviderExpo                 PushProvider = "Expo"
	PushProviderFake                 PushProvider = "Fake"
)

// CodeMashRelease enum.
type CodeMashRelease string

const (
	CodeMashReleaseNotSet         CodeMashRelease = "NotSet"
	CodeMashReleaseCommunity      CodeMashRelease = "Community"
	CodeMashReleaseManagedService CodeMashRelease = "ManagedService"
	CodeMashReleaseEnterprise     CodeMashRelease = "Enterprise"
)

// CodeMashRuntime enum.
type CodeMashRuntime string

const (
	CodeMashRuntimeDevelopment CodeMashRuntime = "Development"
	CodeMashRuntimeCI          CodeMashRuntime = "CI"
	CodeMashRuntimeStaging     CodeMashRuntime = "Staging"
	CodeMashRuntimeProduction  CodeMashRuntime = "Production"
)

// AuthType enum.
type AuthType string

const (
	AuthTypeService  AuthType = "Service"
	AuthTypeEmail    AuthType = "Email"
	AuthTypeUserName AuthType = "UserName"
	AuthTypePhone    AuthType = "Phone"
	AuthTypeGuest    AuthType = "Guest"
	AuthTypeSocial   AuthType = "Social"
)

// AuthStatus enum.
type AuthStatus string

const (
	AuthStatusRegistered        AuthStatus = "Registered"
	AuthStatusPendingValidation AuthStatus = "PendingValidation"
	AuthStatusActive            AuthStatus = "Active"
	AuthStatusUnregistered      AuthStatus = "Unregistered"
	AuthStatusSuspended         AuthStatus = "Suspended"
	AuthStatusInActive          AuthStatus = "InActive"
	AuthStatusBlocked           AuthStatus = "Blocked"
)

// TriggerType enum.
type TriggerType string

const (
	TriggerTypeMembership TriggerType = "Membership"
	TriggerTypeSchema     TriggerType = "Schema"
	TriggerTypeFiles      TriggerType = "Files"
	TriggerTypePayments   TriggerType = "Payments"
)

// TriggerActionType enum.
type TriggerActionType string

const (
	TriggerActionTypeCode        TriggerActionType = "Code"
	TriggerActionTypePush        TriggerActionType = "Push"
	TriggerActionTypeSms         TriggerActionType = "Sms"
	TriggerActionTypeEmail       TriggerActionType = "Email"
	TriggerActionTypeWebhookCall TriggerActionType = "WebhookCall"
	TriggerActionTypeSseCall     TriggerActionType = "SseCall"
	TriggerActionTypeMarketplace TriggerActionType = "Marketplace"
)

// IReturn DTO.
type IReturn[T any] struct {
}

// IReturnVoid DTO.
type IReturnVoid struct {
}

// IHasSessionId DTO.
type IHasSessionId struct {
}

// IHasBearerToken DTO.
type IHasBearerToken struct {
}

// IPost DTO.
type IPost struct {
}

// RequestBase DTO.
type RequestBase struct {
	CultureCode   string `json:"cultureCode,omitempty"`
	TimeZoneId    string `json:"timeZoneId,omitempty"`
	Version       string `json:"version,omitempty"`
	CorrelationId string `json:"correlationId,omitempty"`
}

// ICultureBasedRequest DTO.
type ICultureBasedRequest struct {
}

// IVersionBasedRequest DTO.
type IVersionBasedRequest struct {
}

// IHasCorrelationIdRequest DTO.
type IHasCorrelationIdRequest struct {
}

// EmailAddress DTO.
type EmailAddress struct {
	Address string `json:"address,omitempty"`
}

// DisplayName DTO.
type DisplayName struct {
	Value string `json:"value,omitempty"`
}

// AggregateId DTO.
type AggregateId struct {
	Value string `json:"value,omitempty"`
}

// AccountId DTO.
type AccountId struct {
	AggregateId
}

// UtcDateTime DTO.
type UtcDateTime struct {
}

// ExpirationToken DTO.
type ExpirationToken struct {
	Items float64  `json:"items,omitempty"`
	Unit  TimeUnit `json:"unit,omitempty"`
	Value float64  `json:"value,omitempty"`
}

// CodeMashSubscriptionId DTO.
type CodeMashSubscriptionId struct {
	AggregateId
}

// ProjectId DTO.
type ProjectId struct {
	AggregateId
}

// IntegrationId DTO.
type IntegrationId struct {
	AggregateId
}

// ResourceRef DTO.
type ResourceRef struct {
	ProjectId     *ProjectId      `json:"projectId,omitempty"`
	IntegrationId *IntegrationId  `json:"integrationId,omitempty"`
	Kind          ResourceRefKind `json:"kind,omitempty"`
}

// PaymentCustomerRef DTO.
type PaymentCustomerRef struct {
	ResourceRef
	Kind       ResourceRefKind `json:"kind,omitempty"`
	Source     ResourceSource  `json:"source,omitempty"`
	ExternalId string          `json:"externalId,omitempty"`
}

// Quantity DTO.
type Quantity struct {
	Value float64 `json:"value,omitempty"`
}

// CodeMashManagedServiceSubscription DTO.
type CodeMashManagedServiceSubscription struct {
	SubscriptionId     *CodeMashSubscriptionId `json:"subscriptionId,omitempty"`
	PaymentCustomerRef *PaymentCustomerRef     `json:"paymentCustomerRef,omitempty"`
	RefSubscriptionId  string                  `json:"refSubscriptionId,omitempty"`
	IssuedOn           *UtcDateTime            `json:"issuedOn,omitempty"`
	WillExpireOn       *UtcDateTime            `json:"willExpireOn,omitempty"`
	ProjectCap         *Quantity               `json:"projectCap,omitempty"`
	IsTrial            bool                    `json:"isTrial,omitempty"`
}

// DomainUrl DTO.
type DomainUrl struct {
	Value string `json:"value,omitempty"`
}

// CodeMashLicense DTO.
type CodeMashLicense struct {
	CodeMashManagedServiceSubscription
	Domain       *DomainUrl `json:"domain,omitempty"`
	AccountId    *AccountId `json:"accountId,omitempty"`
	IsEnterprise bool       `json:"isEnterprise,omitempty"`
}

// Tag DTO.
type Tag struct {
}

// TagDescription DTO.
type TagDescription struct {
	DisplayName *DisplayName `json:"displayName,omitempty"`
	Description string       `json:"description,omitempty"`
}

// MessageTranslation DTO.
type MessageTranslation[TContent any] struct {
}

// TagTranslation DTO.
type TagTranslation struct {
	MessageTranslation[*TagDescription]
}

// BaseTagDefinition DTO.
type BaseTagDefinition struct {
	Tag          *Tag              `json:"tag,omitempty"`
	Translations []*TagTranslation `json:"translations,omitempty"`
}

// GroupDefinition DTO.
type GroupDefinition struct {
	BaseTagDefinition
}

// TagDefinition DTO.
type TagDefinition struct {
	BaseTagDefinition
	DefaultDelivery map[string]bool `json:"defaultDelivery,omitempty"`
}

// ProjectName DTO.
type ProjectName struct {
	Name       string `json:"name,omitempty"`
	UniqueName string `json:"uniqueName,omitempty"`
}

// NorbixRegion DTO.
type NorbixRegion struct {
	Code string `json:"code,omitempty"`
}

// ProjectRegion DTO.
type ProjectRegion struct {
	Region    *NorbixRegion `json:"region,omitempty"`
	Name      string        `json:"name,omitempty"`
	Continent Continent     `json:"continent,omitempty"`
}

// Language DTO.
type Language struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// FileResourceId DTO.
type FileResourceId struct {
	Value string `json:"value,omitempty"`
}

// FileChecksum DTO.
type FileChecksum struct {
	Algorithm string `json:"algorithm,omitempty"`
	Hash      string `json:"hash,omitempty"`
}

// FileResource DTO.
type FileResource struct {
	Id               *FileResourceId `json:"id,omitempty"`
	OriginalFileName string          `json:"originalFileName,omitempty"`
	Extension        string          `json:"extension,omitempty"`
	SizeBytes        float64         `json:"sizeBytes,omitempty"`
	Checksum         *FileChecksum   `json:"checksum,omitempty"`
	StoredFileName   string          `json:"storedFileName,omitempty"`
}

// FileResourceRef DTO.
type FileResourceRef struct {
	Resource      *FileResource  `json:"resource,omitempty"`
	IntegrationId *IntegrationId `json:"integrationId,omitempty"`
	Provider      FileProvider   `json:"provider,omitempty"`
	Path          string         `json:"path,omitempty"`
}

// ProjectLogo DTO.
type ProjectLogo struct {
	FileResource *FileResourceRef `json:"fileResource,omitempty"`
	PublicUrl    string           `json:"publicUrl,omitempty"`
}

// ProjectIcon DTO.
type ProjectIcon struct {
	FileResource *FileResourceRef `json:"fileResource,omitempty"`
	PublicUrl    string           `json:"publicUrl,omitempty"`
}

// BrandColor DTO.
type BrandColor struct {
	Value string `json:"value,omitempty"`
}

// TimeZone DTO.
type TimeZone struct {
	ZoneId string `json:"zoneId,omitempty"`
}

// GroupTags DTO.
type GroupTags struct {
	Group *Tag   `json:"group,omitempty"`
	Tags  []*Tag `json:"tags,omitempty"`
}

// ProjectCommunicationChannel DTO.
type ProjectCommunicationChannel struct {
	Channel CommunicationChannel `json:"channel,omitempty"`
	Groups  []*GroupTags         `json:"groups,omitempty"`
}

// ProjectCommunication DTO.
type ProjectCommunication struct {
	Channels []*ProjectCommunicationChannel `json:"channels,omitempty"`
	Groups   []*GroupDefinition             `json:"groups,omitempty"`
	Tags     []*TagDefinition               `json:"tags,omitempty"`
}

// AuthId DTO.
type AuthId struct {
	Value string `json:"value,omitempty"`
}

// DeviceId DTO.
type DeviceId struct {
	Id string `json:"id,omitempty"`
}

// PushDeviceToken DTO.
type PushDeviceToken struct {
	Token string `json:"token,omitempty"`
}

// PushDeviceDeliveryToken DTO.
type PushDeviceDeliveryToken struct {
	PushDeviceToken *PushDeviceToken         `json:"pushDeviceToken,omitempty"`
	DeliveryFamily  PushDeviceDeliveryFamily `json:"deliveryFamily,omitempty"`
}

// PushDevice DTO.
type PushDevice struct {
	Id               *DeviceId                `json:"id,omitempty"`
	Brand            string                   `json:"brand,omitempty"`
	Manufacturer     string                   `json:"manufacturer,omitempty"`
	ModelName        string                   `json:"modelName,omitempty"`
	DeviceName       string                   `json:"deviceName,omitempty"`
	DeviceType       DeviceType               `json:"deviceType,omitempty"`
	OsName           string                   `json:"osName,omitempty"`
	OsVersion        string                   `json:"osVersion,omitempty"`
	PlatformApiLevel float64                  `json:"platformApiLevel,omitempty"`
	Token            *PushDeviceDeliveryToken `json:"token,omitempty"`
}

// CodeMashRequestBase DTO.
type CodeMashRequestBase struct {
	RequestBase
	ProjectId string `json:"projectId,omitempty"`
	Env       string `json:"env,omitempty"`
}

// IHasProjectId DTO.
type IHasProjectId struct {
}

// IHasEnv DTO.
type IHasEnv struct {
}

// UserGeneralInfoDto DTO.
type UserGeneralInfoDto struct {
	Phone                     string                 `json:"phone,omitempty"`
	PrimaryEmail              string                 `json:"primaryEmail,omitempty"`
	DisplayName               string                 `json:"displayName,omitempty"`
	FirstName                 string                 `json:"firstName,omitempty"`
	LastName                  string                 `json:"lastName,omitempty"`
	FullName                  string                 `json:"fullName,omitempty"`
	AddressLine1              string                 `json:"addressLine1,omitempty"`
	AddressLine2              string                 `json:"addressLine2,omitempty"`
	Country                   string                 `json:"country,omitempty"`
	City                      string                 `json:"city,omitempty"`
	State                     string                 `json:"state,omitempty"`
	PostalCode                string                 `json:"postalCode,omitempty"`
	Company                   string                 `json:"company,omitempty"`
	Gender                    Gender                 `json:"gender,omitempty"`
	BirthDate                 float64                `json:"birthDate,omitempty"`
	TimeZone                  string                 `json:"timeZone,omitempty"`
	Language                  string                 `json:"language,omitempty"`
	BlockAllMarketingMessages bool                   `json:"blockAllMarketingMessages,omitempty"`
	BlockedTags               map[string][]string    `json:"blockedTags,omitempty"`
	BlockReasons              []MarketingBlockReason `json:"blockReasons,omitempty"`
	ExtraMetadata             string                 `json:"extraMetadata,omitempty"`
	Notes                     string                 `json:"notes,omitempty"`
}

// SaveUser DTO.
type SaveUser struct {
	CodeMashRequestBase
	DatabaseIntegrationId     string              `json:"databaseIntegrationId,omitempty"`
	UserGeneralInfo           *UserGeneralInfoDto `json:"userGeneralInfo,omitempty"`
	UserId                    string              `json:"userId,omitempty"`
	IgnoreUserRegistersAsRole bool                `json:"ignoreUserRegistersAsRole,omitempty"`
}

// SaveUserWithRolesBase DTO.
type SaveUserWithRolesBase struct {
	SaveUser
	Roles []string `json:"roles,omitempty"`
}

// Env DTO.
type Env struct {
	Value  string `json:"value,omitempty"`
	IsProd bool   `json:"isProd,omitempty"`
}

// CursorArgs DTO.
type CursorArgs struct {
	Field string  `json:"field,omitempty"`
	Order float64 `json:"order,omitempty"`
}

// PagingArgs DTO.
type PagingArgs struct {
	CursorArgs    *CursorArgs `json:"cursorArgs,omitempty"`
	PageSize      float64     `json:"pageSize,omitempty"`
	StartingAfter string      `json:"startingAfter,omitempty"`
	EndingBefore  string      `json:"endingBefore,omitempty"`
}

// CodeMashListPaginationRequestBase DTO.
type CodeMashListPaginationRequestBase struct {
	RequestBase
	ProjectId     string      `json:"projectId,omitempty"`
	Env           string      `json:"env,omitempty"`
	ResolvedEnv   *Env        `json:"resolvedEnv,omitempty"`
	StartingAfter string      `json:"startingAfter,omitempty"`
	EndingBefore  string      `json:"endingBefore,omitempty"`
	PageSize      float64     `json:"pageSize,omitempty"`
	Paging        *PagingArgs `json:"paging,omitempty"`
}

// IPasskeyCeremonyRequest DTO.
type IPasskeyCeremonyRequest struct {
}

// Integration DTO.
type Integration struct {
	IntegrationId                    *IntegrationId `json:"integrationId,omitempty"`
	Env                              *Env           `json:"env,omitempty"`
	Capability                       string         `json:"capability,omitempty"`
	IsSystemOwned                    bool           `json:"isSystemOwned,omitempty"`
	IntegrationName                  *DisplayName   `json:"integrationName,omitempty"`
	IsEnabled                        bool           `json:"isEnabled,omitempty"`
	IsConfigured                     bool           `json:"isConfigured,omitempty"`
	LastIntegrationTestAtUtc         string         `json:"lastIntegrationTestAtUtc,omitempty"`
	LastIntegrationTestSucceeded     bool           `json:"lastIntegrationTestSucceeded,omitempty"`
	LastIntegrationTestErrorMessages []string       `json:"lastIntegrationTestErrorMessages,omitempty"`
	HumanDeliveryConfirmedAtUtc      string         `json:"humanDeliveryConfirmedAtUtc,omitempty"`
	IsApprovedThatItWorks            bool           `json:"isApprovedThatItWorks,omitempty"`
}

// PushIntegration DTO.
type PushIntegration struct {
	Integration
	Provider PushProvider `json:"provider,omitempty"`
}

// TemplateId DTO.
type TemplateId struct {
	Value string `json:"value,omitempty"`
}

// Template DTO.
type Template[TMessageContent any] struct {
	TemplateId           *TemplateId                           `json:"templateId,omitempty"`
	TemplateName         *DisplayName                          `json:"templateName,omitempty"`
	Translations         []MessageTranslation[TMessageContent] `json:"translations,omitempty"`
	CommunicationChannel CommunicationChannel                  `json:"communicationChannel,omitempty"`
	IsActive             bool                                  `json:"isActive,omitempty"`
	Description          string                                `json:"description,omitempty"`
	Tags                 []*Tag                                `json:"tags,omitempty"`
	FileIntegrationId    *IntegrationId                        `json:"fileIntegrationId,omitempty"`
	Env                  *Env                                  `json:"env,omitempty"`
}

// TemplateCode DTO.
type TemplateCode struct {
}

// PushTitle DTO.
type PushTitle struct {
	Value *TemplateCode `json:"value,omitempty"`
}

// PushBody DTO.
type PushBody struct {
	Value *TemplateCode `json:"value,omitempty"`
}

// PushMessageContent DTO.
type PushMessageContent struct {
	Title    *PushTitle `json:"title,omitempty"`
	SubTitle *PushTitle `json:"subTitle,omitempty"`
	Body     *PushBody  `json:"body,omitempty"`
}

// PushTemplate DTO.
type PushTemplate struct {
	Template[*PushMessageContent]
}

// EchoLicenseDto DTO.
type EchoLicenseDto struct {
	Domain    string  `json:"domain,omitempty"`
	AccountId string  `json:"accountId,omitempty"`
	Email     string  `json:"email,omitempty"`
	Release   string  `json:"release,omitempty"`
	Expire    float64 `json:"expire,omitempty"`
	IsTrial   bool    `json:"isTrial,omitempty"`
	Cap       float64 `json:"cap,omitempty"`
}

// EchoRegionDto DTO.
type EchoRegionDto struct {
	Code        string `json:"code,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	ApiUrl      string `json:"apiUrl,omitempty"`
	HubUrl      string `json:"hubUrl,omitempty"`
}

// PublicBrandDto DTO.
type PublicBrandDto struct {
	DisplayName string `json:"displayName,omitempty"`
	MainColor   string `json:"mainColor,omitempty"`
	AccentColor string `json:"accentColor,omitempty"`
	LogoUrl     string `json:"logoUrl,omitempty"`
	IconUrl     string `json:"iconUrl,omitempty"`
}

// PublicPasswordPolicyDto DTO.
type PublicPasswordPolicyDto struct {
	MinLength      float64 `json:"minLength,omitempty"`
	MaxLength      float64 `json:"maxLength,omitempty"`
	MinNumbers     float64 `json:"minNumbers,omitempty"`
	MinUpper       float64 `json:"minUpper,omitempty"`
	MinLower       float64 `json:"minLower,omitempty"`
	MinSpecial     float64 `json:"minSpecial,omitempty"`
	AllowedSpecial string  `json:"allowedSpecial,omitempty"`
}

// PublicAuthDto DTO.
type PublicAuthDto struct {
	SocialProviders []string                 `json:"socialProviders,omitempty"`
	Passkey         bool                     `json:"passkey,omitempty"`
	Methods         []string                 `json:"methods,omitempty"`
	PasswordPolicy  *PublicPasswordPolicyDto `json:"passwordPolicy,omitempty"`
}

// ErrorDto DTO.
type ErrorDto struct {
	Message    string            `json:"message,omitempty"`
	ErrorCode  string            `json:"errorCode,omitempty"`
	Context    map[string]string `json:"context,omitempty"`
	StackTrace []*ErrorDto       `json:"stackTrace,omitempty"`
}

// CodeMashResponseStatus DTO.
type CodeMashResponseStatus struct {
	IsSuccess bool        `json:"isSuccess,omitempty"`
	Errors    []*ErrorDto `json:"errors,omitempty"`
}

// ResponseBase DTO.
type ResponseBase struct {
	ResponseStatus *CodeMashResponseStatus `json:"responseStatus,omitempty"`
}

// AccessInformationDto DTO.
type AccessInformationDto struct {
	Ip       string `json:"ip,omitempty"`
	Date     string `json:"date,omitempty"`
	TimeZone string `json:"timeZone,omitempty"`
}

// RegistrationDto DTO.
type RegistrationDto struct {
	RegistrationInformation *AccessInformationDto `json:"registrationInformation,omitempty"`
}

// LoginDto DTO.
type LoginDto struct {
	NeedChangePasswordOnNextLogin bool                  `json:"needChangePasswordOnNextLogin,omitempty"`
	LastAccessInformation         *AccessInformationDto `json:"lastAccessInformation,omitempty"`
}

// AuthDto DTO.
type AuthDto struct {
	Id           string              `json:"id,omitempty"`
	Type         AuthType            `json:"type,omitempty"`
	Email        string              `json:"email,omitempty"`
	UserName     string              `json:"userName,omitempty"`
	Registration *RegistrationDto    `json:"registration,omitempty"`
	Login        *LoginDto           `json:"login,omitempty"`
	GeneralInfo  *UserGeneralInfoDto `json:"generalInfo,omitempty"`
	Roles        []string            `json:"roles,omitempty"`
	PushDevices  []string            `json:"pushDevices,omitempty"`
	Tags         []string            `json:"tags,omitempty"`
	Status       AuthStatus          `json:"status,omitempty"`
	CreatedOn    string              `json:"createdOn,omitempty"`
	ModifiedOn   string              `json:"modifiedOn,omitempty"`
}

// PaginatedResponse DTO.
type PaginatedResponse[TViewModelProjection any] struct {
	Items         []TViewModelProjection `json:"items,omitempty"`
	HasMore       bool                   `json:"hasMore,omitempty"`
	HasPrevious   bool                   `json:"hasPrevious,omitempty"`
	StartingAfter string                 `json:"startingAfter,omitempty"`
	EndingBefore  string                 `json:"endingBefore,omitempty"`
}

// UserMarketingPreferencesDto DTO.
type UserMarketingPreferencesDto struct {
	BlockAllMarketingMessages bool                   `json:"blockAllMarketingMessages,omitempty"`
	BlockedTags               map[string][]string    `json:"blockedTags,omitempty"`
	BlockReasons              []MarketingBlockReason `json:"blockReasons,omitempty"`
}

// PasskeyListItemDto DTO.
type PasskeyListItemDto struct {
	CredentialId    string `json:"credentialId,omitempty"`
	FriendlyName    string `json:"friendlyName,omitempty"`
	RegisteredOnUtc string `json:"registeredOnUtc,omitempty"`
	LastUsedOnUtc   string `json:"lastUsedOnUtc,omitempty"`
	IsRevoked       bool   `json:"isRevoked,omitempty"`
}

// TermMultiParentDto DTO.
type TermMultiParentDto struct {
	TaxonomyId string            `json:"taxonomyId,omitempty"`
	ParentId   string            `json:"parentId,omitempty"`
	Name       string            `json:"name,omitempty"`
	Names      map[string]string `json:"names,omitempty"`
}

// TermTreeDto DTO.
type TermTreeDto struct {
	Id           string                `json:"id,omitempty"`
	TaxonomyId   string                `json:"taxonomyId,omitempty"`
	TaxonomyName string                `json:"taxonomyName,omitempty"`
	ParentId     string                `json:"parentId,omitempty"`
	Order        float64               `json:"order,omitempty"`
	Name         string                `json:"name,omitempty"`
	Names        map[string]string     `json:"names,omitempty"`
	Description  string                `json:"description,omitempty"`
	Descriptions map[string]string     `json:"descriptions,omitempty"`
	MultiParents []*TermMultiParentDto `json:"multiParents,omitempty"`
	Meta         map[string]any        `json:"meta,omitempty"`
	Children     []*TermTreeDto        `json:"children,omitempty"`
}

// TaxonomyTreeDto DTO.
type TaxonomyTreeDto struct {
	ViewId       string             `json:"viewId,omitempty"`
	TaxonomyName string             `json:"taxonomyName,omitempty"`
	TaxonomySlug string             `json:"taxonomySlug,omitempty"`
	ParentId     string             `json:"parentId,omitempty"`
	Children     []*TaxonomyTreeDto `json:"children,omitempty"`
	Terms        []*TermTreeDto     `json:"terms,omitempty"`
}

// TermDto DTO.
type TermDto struct {
	Id           string                `json:"id,omitempty"`
	TaxonomyId   string                `json:"taxonomyId,omitempty"`
	TaxonomyName string                `json:"taxonomyName,omitempty"`
	ParentId     string                `json:"parentId,omitempty"`
	Order        float64               `json:"order,omitempty"`
	Name         string                `json:"name,omitempty"`
	Names        map[string]string     `json:"names,omitempty"`
	Description  string                `json:"description,omitempty"`
	Descriptions map[string]string     `json:"descriptions,omitempty"`
	MultiParents []*TermMultiParentDto `json:"multiParents,omitempty"`
	Meta         map[string]any        `json:"meta,omitempty"`
}

// JsonSchemaFieldDto DTO.
type JsonSchemaFieldDto struct {
	FieldName string `json:"fieldName,omitempty"`
}

// DataSchemaDto DTO.
type DataSchemaDto struct {
	Json   string                `json:"json,omitempty"`
	Fields []*JsonSchemaFieldDto `json:"fields,omitempty"`
}

// VisualSchemaDto DTO.
type VisualSchemaDto struct {
	Json string `json:"json,omitempty"`
}

// SchemaSettingsDto DTO.
type SchemaSettingsDto struct {
	SoftDelete     bool   `json:"softDelete,omitempty"`
	HasRecordOwner bool   `json:"hasRecordOwner,omitempty"`
	Description    string `json:"description,omitempty"`
}

// TriggerActionDto DTO.
type TriggerActionDto struct {
	Type          TriggerActionType `json:"type,omitempty"`
	IntegrationId string            `json:"integrationId,omitempty"`
}

// TriggerDto DTO.
type TriggerDto struct {
	Type           TriggerType       `json:"type,omitempty"`
	ViewId         string            `json:"viewId,omitempty"`
	Name           string            `json:"name,omitempty"`
	ThenAction     *TriggerActionDto `json:"thenAction,omitempty"`
	Description    string            `json:"description,omitempty"`
	IsEnabled      bool              `json:"isEnabled,omitempty"`
	ActivationCode string            `json:"activationCode,omitempty"`
	SavedByAuthId  string            `json:"savedByAuthId,omitempty"`
}

// SchemaDto DTO.
type SchemaDto struct {
	ViewId            string             `json:"viewId,omitempty"`
	SchemaName        string             `json:"schemaName,omitempty"`
	SchemaSlug        string             `json:"schemaSlug,omitempty"`
	Version           float64            `json:"version,omitempty"`
	MetaSchemaVersion float64            `json:"metaSchemaVersion,omitempty"`
	DataSchema        *DataSchemaDto     `json:"dataSchema,omitempty"`
	VisualSchema      *VisualSchemaDto   `json:"visualSchema,omitempty"`
	PublishedAt       string             `json:"publishedAt,omitempty"`
	Settings          *SchemaSettingsDto `json:"settings,omitempty"`
	Triggers          []*TriggerDto      `json:"triggers,omitempty"`
}

// SchemaListProjection DTO.
type SchemaListProjection struct {
	ViewId            string  `json:"viewId,omitempty"`
	SchemaName        string  `json:"schemaName,omitempty"`
	SchemaTitle       string  `json:"schemaTitle,omitempty"`
	LatestVersion     float64 `json:"latestVersion,omitempty"`
	HasDraft          bool    `json:"hasDraft,omitempty"`
	MetaSchemaVersion float64 `json:"metaSchemaVersion,omitempty"`
	Description       string  `json:"description,omitempty"`
}

// FileChecksumDto DTO.
type FileChecksumDto struct {
	Algorithm string `json:"algorithm,omitempty"`
	Hash      string `json:"hash,omitempty"`
}

// FileResourceDto DTO.
type FileResourceDto struct {
	Id               string           `json:"id,omitempty"`
	OriginalFileName string           `json:"originalFileName,omitempty"`
	Extension        string           `json:"extension,omitempty"`
	StoredFileName   string           `json:"storedFileName,omitempty"`
	SizeBytes        float64          `json:"sizeBytes,omitempty"`
	Checksum         *FileChecksumDto `json:"checksum,omitempty"`
}

// FileResourceRefDto DTO.
type FileResourceRefDto struct {
	Resource      *FileResourceDto `json:"resource,omitempty"`
	IntegrationId string           `json:"integrationId,omitempty"`
	Provider      FileProvider     `json:"provider,omitempty"`
	Path          string           `json:"path,omitempty"`
	PublicUrl     string           `json:"publicUrl,omitempty"`
	IsPublic      bool             `json:"isPublic,omitempty"`
}

// PublicFolderDto DTO.
type PublicFolderDto struct {
	Path      string `json:"path,omitempty"`
	PublicId  string `json:"publicId,omitempty"`
	PublicUrl string `json:"publicUrl,omitempty"`
	Inherited bool   `json:"inherited,omitempty"`
}

// IntegrationTestResultItemDto DTO.
type IntegrationTestResultItemDto struct {
	Operation string   `json:"operation,omitempty"`
	Result    string   `json:"result,omitempty"`
	Errors    []string `json:"errors,omitempty"`
}

// ResponseError DTO.
type ResponseError struct {
	ErrorCode string            `json:"errorCode,omitempty"`
	FieldName string            `json:"fieldName,omitempty"`
	Message   string            `json:"message,omitempty"`
	Meta      map[string]string `json:"meta,omitempty"`
}

// ResponseStatus DTO.
type ResponseStatus struct {
	ErrorCode  string            `json:"errorCode,omitempty"`
	Message    string            `json:"message,omitempty"`
	StackTrace string            `json:"stackTrace,omitempty"`
	Errors     []*ResponseError  `json:"errors,omitempty"`
	Meta       map[string]string `json:"meta,omitempty"`
}

// IHasDomainEntityId DTO.
type IHasDomainEntityId struct {
}

// IIntegrationIdentification DTO.
type IIntegrationIdentification struct {
}

// IBindableContract DTO.
type IBindableContract struct {
}

// IHasViewId DTO.
type IHasViewId struct {
}

// ICursorArgs DTO.
type ICursorArgs struct {
}

// StringFieldDto DTO.
type StringFieldDto struct {
	JsonSchemaFieldDto
	Format           string            `json:"format,omitempty"`
	Pattern          string            `json:"pattern,omitempty"`
	MinLength        float64           `json:"minLength,omitempty"`
	MaxLength        float64           `json:"maxLength,omitempty"`
	TranslateOptions map[string]string `json:"translateOptions,omitempty"`
}

// DecimalFieldDto DTO.
type DecimalFieldDto struct {
	JsonSchemaFieldDto
	Minimum    float64 `json:"minimum,omitempty"`
	Maximum    float64 `json:"maximum,omitempty"`
	MultipleOf float64 `json:"multipleOf,omitempty"`
}

// CurrencyFieldDto DTO.
type CurrencyFieldDto struct {
	JsonSchemaFieldDto
	AllowedCurrencies []string `json:"allowedCurrencies,omitempty"`
}

// BooleanFieldDto DTO.
type BooleanFieldDto struct {
	JsonSchemaFieldDto
}

// DateFieldDto DTO.
type DateFieldDto struct {
	JsonSchemaFieldDto
	Minimum float64 `json:"minimum,omitempty"`
	Maximum float64 `json:"maximum,omitempty"`
}

// IntegerFieldDto DTO.
type IntegerFieldDto struct {
	JsonSchemaFieldDto
	Minimum float64 `json:"minimum,omitempty"`
	Maximum float64 `json:"maximum,omitempty"`
}

// GeolocationFieldDto DTO.
type GeolocationFieldDto struct {
	JsonSchemaFieldDto
	AllowedTypes []string `json:"allowedTypes,omitempty"`
}

// TagsFieldDto DTO.
type TagsFieldDto struct {
	JsonSchemaFieldDto
}

// FileFieldDto DTO.
type FileFieldDto struct {
	JsonSchemaFieldDto
	Storages []string `json:"storages,omitempty"`
}

// TaxonomySelectionFieldDto DTO.
type TaxonomySelectionFieldDto struct {
	JsonSchemaFieldDto
	TaxonomyId string `json:"taxonomyId,omitempty"`
	Multiple   bool   `json:"multiple,omitempty"`
}

// CollectionSelectionFieldDto DTO.
type CollectionSelectionFieldDto struct {
	JsonSchemaFieldDto
	CollectionId string `json:"collectionId,omitempty"`
	DisplayField string `json:"displayField,omitempty"`
	Multiple     bool   `json:"multiple,omitempty"`
}

// UserSelectionFieldDto DTO.
type UserSelectionFieldDto struct {
	JsonSchemaFieldDto
	Multiple bool `json:"multiple,omitempty"`
}

// RoleSelectionFieldDto DTO.
type RoleSelectionFieldDto struct {
	JsonSchemaFieldDto
	Multiple bool `json:"multiple,omitempty"`
}

// EnumSelectionFieldDto DTO.
type EnumSelectionFieldDto struct {
	JsonSchemaFieldDto
	Values   []string `json:"values,omitempty"`
	Multiple bool     `json:"multiple,omitempty"`
}

// EchoResponse DTO.
type EchoResponse struct {
	ContainerName                string           `json:"containerName,omitempty"`
	Ip                           string           `json:"ip,omitempty"`
	Release                      CodeMashRelease  `json:"release,omitempty"`
	Runtime                      CodeMashRuntime  `json:"runtime,omitempty"`
	ManagedServiceHubUrl         string           `json:"managedServiceHubUrl,omitempty"`
	ManagedServiceApiUrl         string           `json:"managedServiceApiUrl,omitempty"`
	HubUrl                       string           `json:"hubUrl,omitempty"`
	ApiUrl                       string           `json:"apiUrl,omitempty"`
	ApiVersion                   string           `json:"apiVersion,omitempty"`
	HubVersion                   string           `json:"hubVersion,omitempty"`
	MjmlUrl                      string           `json:"mjmlUrl,omitempty"`
	AdminUrlTemplate             string           `json:"adminUrlTemplate,omitempty"`
	License                      *EchoLicenseDto  `json:"license,omitempty"`
	AskForEnterpriseLicenseEmail string           `json:"askForEnterpriseLicenseEmail,omitempty"`
	EmailServiceConfigured       bool             `json:"emailServiceConfigured,omitempty"`
	RootBootstrapPasswordSource  string           `json:"rootBootstrapPasswordSource,omitempty"`
	Regions                      []*EchoRegionDto `json:"regions,omitempty"`
	IsProductionInstallation     bool             `json:"isProductionInstallation,omitempty"`
	LicensingMode                string           `json:"licensingMode,omitempty"`
	GraceDaysLeft                float64          `json:"graceDaysLeft,omitempty"`
	InstallationDomain           string           `json:"installationDomain,omitempty"`
	LicensingDocsUrl             string           `json:"licensingDocsUrl,omitempty"`
}

// PublicProjectConfigDto DTO.
type PublicProjectConfigDto struct {
	DisplayName        string          `json:"displayName,omitempty"`
	AdminPortalEnabled bool            `json:"adminPortalEnabled,omitempty"`
	Branding           *PublicBrandDto `json:"branding,omitempty"`
	Auth               *PublicAuthDto  `json:"auth,omitempty"`
}

// PublicLegalDocumentDto DTO.
type PublicLegalDocumentDto struct {
	Kind      string `json:"kind,omitempty"`
	Title     string `json:"title,omitempty"`
	Body      string `json:"body,omitempty"`
	Available bool   `json:"available,omitempty"`
}

// AskChatResponse DTO.
type AskChatResponse struct {
	ResponseBase
	Result string `json:"result,omitempty"`
}

// EmptyResponse DTO.
type EmptyResponse struct {
	ResponseBase
}

// IdResponse DTO.
type IdResponse struct {
	ResponseBase
	Id     string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
}

// GetUserResponse DTO.
type GetUserResponse struct {
	ResponseBase
	User *AuthDto `json:"user,omitempty"`
}

// GetUsersResponse DTO.
type GetUsersResponse struct {
	ResponseBase
	List PaginatedResponse[*AuthDto] `json:"list,omitempty"`
}

// GetUserPreferencesResponse DTO.
type GetUserPreferencesResponse struct {
	ResponseBase
	Preferences *UserMarketingPreferencesDto `json:"preferences,omitempty"`
}

// PasskeyOkResponse DTO.
type PasskeyOkResponse struct {
	ResponseBase
}

// PasskeyCeremonyOptionsResponse DTO.
type PasskeyCeremonyOptionsResponse struct {
	ResponseBase
	CeremonyId  string `json:"ceremonyId,omitempty"`
	OptionsJson string `json:"optionsJson,omitempty"`
}

// PasskeyAuthTokensResponse DTO.
type PasskeyAuthTokensResponse struct {
	ResponseBase
	AccessToken      string   `json:"accessToken,omitempty"`
	RefreshToken     string   `json:"refreshToken,omitempty"`
	ExpiresInSeconds float64  `json:"expiresInSeconds,omitempty"`
	RecoveryCodes    []string `json:"recoveryCodes,omitempty"`
}

// PasskeyListResponse DTO.
type PasskeyListResponse struct {
	ResponseBase
	Passkeys []*PasskeyListItemDto `json:"passkeys,omitempty"`
}

// PasskeyRecoveryResponse DTO.
type PasskeyRecoveryResponse struct {
	ResponseBase
	AccessToken      string  `json:"accessToken,omitempty"`
	RefreshToken     string  `json:"refreshToken,omitempty"`
	ExpiresInSeconds float64 `json:"expiresInSeconds,omitempty"`
	RemainingCodes   float64 `json:"remainingCodes,omitempty"`
}

// PasskeyVerificationTokenResponse DTO.
type PasskeyVerificationTokenResponse struct {
	ResponseBase
	VerificationToken string `json:"verificationToken,omitempty"`
}

// FindMergedTermTreeResponse DTO.
type FindMergedTermTreeResponse struct {
	ResponseBase
	Tree []*TermTreeDto `json:"tree,omitempty"`
}

// FindTaxonomyTreeResponse DTO.
type FindTaxonomyTreeResponse struct {
	ResponseBase
	Tree []*TaxonomyTreeDto `json:"tree,omitempty"`
}

// FindTermsResponse DTO.
type FindTermsResponse struct {
	ResponseBase
	List PaginatedResponse[*TermDto] `json:"list,omitempty"`
}

// FindTermsChildrenResponse DTO.
type FindTermsChildrenResponse struct {
	ResponseBase
	List PaginatedResponse[*TermDto] `json:"list,omitempty"`
}

// FindTermTreeResponse DTO.
type FindTermTreeResponse struct {
	ResponseBase
	Tree []*TermTreeDto `json:"tree,omitempty"`
}

// GetDatabaseSchemaResponse DTO.
type GetDatabaseSchemaResponse struct {
	ResponseBase
	Item *SchemaDto `json:"item,omitempty"`
}

// GetDatabaseSchemasResponse DTO.
type GetDatabaseSchemasResponse struct {
	ResponseBase
	List PaginatedResponse[*SchemaListProjection] `json:"list,omitempty"`
}

// AggregateResponse DTO.
type AggregateResponse struct {
	ResponseBase
	Result []map[string]any `json:"result,omitempty"`
}

// CountResponse DTO.
type CountResponse struct {
	ResponseBase
	Count float64 `json:"count,omitempty"`
}

// DistinctResponse DTO.
type DistinctResponse struct {
	ResponseBase
	Values []map[string]any `json:"values,omitempty"`
}

// ExecuteAggregateResponse DTO.
type ExecuteAggregateResponse struct {
	ResponseBase
	Result []map[string]any `json:"result,omitempty"`
}

// FindResponse DTO.
type FindResponse struct {
	ResponseBase
	List PaginatedResponse[map[string]any] `json:"list,omitempty"`
}

// FindOneResponse DTO.
type FindOneResponse struct {
	ResponseBase
	Result map[string]any `json:"result,omitempty"`
}

// GetFileInfoResponse DTO.
type GetFileInfoResponse struct {
	ResponseBase
	File      *FileResourceRefDto `json:"file,omitempty"`
	IsPublic  bool                `json:"isPublic,omitempty"`
	PublicUrl string              `json:"publicUrl,omitempty"`
}

// GetSignedUrlResponse DTO.
type GetSignedUrlResponse struct {
	ResponseBase
	Url string `json:"url,omitempty"`
}

// ListFilesResponse DTO.
type ListFilesResponse struct {
	ResponseBase
	List          PaginatedResponse[*FileResourceRefDto] `json:"list,omitempty"`
	Folders       []string                               `json:"folders,omitempty"`
	PublicFolders []*PublicFolderDto                     `json:"publicFolders,omitempty"`
}

// RequestUploadUrlResponse DTO.
type RequestUploadUrlResponse struct {
	ResponseBase
	Url string `json:"url,omitempty"`
}

// TestFilesIntegrationResponse DTO.
type TestFilesIntegrationResponse struct {
	ResponseBase
	Items []*IntegrationTestResultItemDto `json:"items,omitempty"`
}

// AuthenticateResponse DTO.
type AuthenticateResponse struct {
	UserId             string            `json:"userId,omitempty"`
	SessionId          string            `json:"sessionId,omitempty"`
	UserName           string            `json:"userName,omitempty"`
	DisplayName        string            `json:"displayName,omitempty"`
	ReferrerUrl        string            `json:"referrerUrl,omitempty"`
	BearerToken        string            `json:"bearerToken,omitempty"`
	RefreshToken       string            `json:"refreshToken,omitempty"`
	RefreshTokenExpiry string            `json:"refreshTokenExpiry,omitempty"`
	ProfileUrl         string            `json:"profileUrl,omitempty"`
	Roles              []string          `json:"roles,omitempty"`
	Permissions        []string          `json:"permissions,omitempty"`
	AuthProvider       string            `json:"authProvider,omitempty"`
	ResponseStatus     *ResponseStatus   `json:"responseStatus,omitempty"`
	Meta               map[string]string `json:"meta,omitempty"`
}

// GetAccessTokenResponse DTO.
type GetAccessTokenResponse struct {
	AccessToken    string            `json:"accessToken,omitempty"`
	Meta           map[string]string `json:"meta,omitempty"`
	ResponseStatus *ResponseStatus   `json:"responseStatus,omitempty"`
}

// Echo DTO.
type Echo struct {
	RequestBase
}

// GetPublicProjectConfig DTO.
type GetPublicProjectConfig struct {
	RequestBase
	ProjectId string `json:"projectId,omitempty"`
}

// GetPublicProjectLegal DTO.
type GetPublicProjectLegal struct {
	RequestBase
	ProjectId string `json:"projectId,omitempty"`
	Kind      string `json:"kind,omitempty"`
}

// AccountCreated DTO.
type AccountCreated struct {
	Email       *EmailAddress `json:"email,omitempty"`
	DisplayName *DisplayName  `json:"displayName,omitempty"`
	AccountId   *AccountId    `json:"accountId,omitempty"`
	CreatedOn   *UtcDateTime  `json:"createdOn,omitempty"`
}

// AccountVerified DTO.
type AccountVerified struct {
}

// AccountSetAsActive DTO.
type AccountSetAsActive struct {
}

// AccountValidationTokenIssued DTO.
type AccountValidationTokenIssued struct {
	Expiration *ExpirationToken `json:"expiration,omitempty"`
}

// AccountBlocked DTO.
type AccountBlocked struct {
}

// AccountProfileUpdated DTO.
type AccountProfileUpdated struct {
	DisplayName     *DisplayName  `json:"displayName,omitempty"`
	BillingEmail    *EmailAddress `json:"billingEmail,omitempty"`
	OperationsEmail *EmailAddress `json:"operationsEmail,omitempty"`
	SecurityEmail   *EmailAddress `json:"securityEmail,omitempty"`
}

// AccountSetAsInactive DTO.
type AccountSetAsInactive struct {
}

// AccountUnregistered DTO.
type AccountUnregistered struct {
}

// LicenseCreated DTO.
type LicenseCreated struct {
	License *CodeMashLicense `json:"license,omitempty"`
}

// CustomerCreated DTO.
type CustomerCreated struct {
	PaymentCustomerRef *PaymentCustomerRef `json:"paymentCustomerRef,omitempty"`
}

// SubscriptionChanged DTO.
type SubscriptionChanged struct {
	Subscription *CodeMashManagedServiceSubscription `json:"subscription,omitempty"`
}

// SubscriptionCanceled DTO.
type SubscriptionCanceled struct {
	PaymentCustomerRef *PaymentCustomerRef `json:"paymentCustomerRef,omitempty"`
	SubscriptionId     string              `json:"subscriptionId,omitempty"`
}

// ProjectCommunicationGroupSaved DTO.
type ProjectCommunicationGroupSaved struct {
	Group         *GroupDefinition     `json:"group,omitempty"`
	Channel       CommunicationChannel `json:"channel,omitempty"`
	OriginChannel CommunicationChannel `json:"originChannel,omitempty"`
}

// ProjectCommunicationTagFromGroupDeleted DTO.
type ProjectCommunicationTagFromGroupDeleted struct {
	GroupTag   *Tag `json:"groupTag,omitempty"`
	RemovedTag *Tag `json:"removedTag,omitempty"`
}

// ProjectCommunicationGroupDeleted DTO.
type ProjectCommunicationGroupDeleted struct {
	GroupTag *Tag `json:"groupTag,omitempty"`
}

// ProjectCommunicationTagSaved DTO.
type ProjectCommunicationTagSaved struct {
	Tag      *TagDefinition       `json:"tag,omitempty"`
	GroupTag *Tag                 `json:"groupTag,omitempty"`
	Channel  CommunicationChannel `json:"channel,omitempty"`
}

// ProjectCommunicationTagDeleted DTO.
type ProjectCommunicationTagDeleted struct {
	Tag *Tag `json:"tag,omitempty"`
}

// ProjectCreated DTO.
type ProjectCreated struct {
	Id                    *ProjectId       `json:"id,omitempty"`
	Name                  *ProjectName     `json:"name,omitempty"`
	DatabaseIntegrationId *IntegrationId   `json:"databaseIntegrationId,omitempty"`
	PrimaryRegion         *ProjectRegion   `json:"primaryRegion,omitempty"`
	AdditionalRegions     []*ProjectRegion `json:"additionalRegions,omitempty"`
	Description           string           `json:"description,omitempty"`
	IsProvisioning        bool             `json:"isProvisioning,omitempty"`
}

// ProjectDeleted DTO.
type ProjectDeleted struct {
}

// ProjectActivated DTO.
type ProjectActivated struct {
}

// ProjectDisabled DTO.
type ProjectDisabled struct {
}

// ProjectNameChanged DTO.
type ProjectNameChanged struct {
	ProjectName *ProjectName `json:"projectName,omitempty"`
}

// ProjectDescriptionChanged DTO.
type ProjectDescriptionChanged struct {
	Description string `json:"description,omitempty"`
}

// ProjectMarketingUrlChanged DTO.
type ProjectMarketingUrlChanged struct {
	Url *DomainUrl `json:"url,omitempty"`
}

// ProjectAllowedOriginsChanged DTO.
type ProjectAllowedOriginsChanged struct {
	Origins []*DomainUrl `json:"origins,omitempty"`
}

// ProjectDefaultLanguageChanged DTO.
type ProjectDefaultLanguageChanged struct {
	Language *Language `json:"language,omitempty"`
}

// ProjectLanguagesChanged DTO.
type ProjectLanguagesChanged struct {
	Languages []*Language `json:"languages,omitempty"`
}

// ProjectLogoChanged DTO.
type ProjectLogoChanged struct {
	Logo *ProjectLogo `json:"logo,omitempty"`
}

// ProjectIconChanged DTO.
type ProjectIconChanged struct {
	Icon *ProjectIcon `json:"icon,omitempty"`
}

// ProjectMainColorChanged DTO.
type ProjectMainColorChanged struct {
	Color *BrandColor `json:"color,omitempty"`
}

// ProjectAccentColorChanged DTO.
type ProjectAccentColorChanged struct {
	Color *BrandColor `json:"color,omitempty"`
}

// ProjectRegionsChanged DTO.
type ProjectRegionsChanged struct {
	PrimaryRegion     *ProjectRegion   `json:"primaryRegion,omitempty"`
	AdditionalRegions []*ProjectRegion `json:"additionalRegions,omitempty"`
}

// ProjectTimeZoneChanged DTO.
type ProjectTimeZoneChanged struct {
	TimeZone *TimeZone `json:"timeZone,omitempty"`
}

// ProjectPaymentZonesChanged DTO.
type ProjectPaymentZonesChanged struct {
	PaymentZones []*TimeZone `json:"paymentZones,omitempty"`
}

// ProjectCommunicationSet DTO.
type ProjectCommunicationSet struct {
	ProjectCommunication *ProjectCommunication `json:"projectCommunication,omitempty"`
}

// AccountUserPushDeviceCreated DTO.
type AccountUserPushDeviceCreated struct {
	AuthId     *AuthId     `json:"authId,omitempty"`
	PushDevice *PushDevice `json:"pushDevice,omitempty"`
}

// AskChatRequest DTO.
type AskChatRequest struct {
	CodeMashRequestBase
	Prompt string `json:"prompt,omitempty"`
}

// BlockUserRequest DTO.
type BlockUserRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// SaveSystemUserWithPermissions DTO.
type SaveSystemUserWithPermissions struct {
	SaveUserWithRolesBase
}

// SaveGuestUser DTO.
type SaveGuestUser struct {
	SaveUser
}

// SaveUserNameUser DTO.
type SaveUserNameUser struct {
	SaveUser
	Password string `json:"password,omitempty"`
	UserName string `json:"userName,omitempty"`
}

// SaveEmailUser DTO.
type SaveEmailUser struct {
	SaveUser
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

// SavePhoneUser DTO.
type SavePhoneUser struct {
	SaveUser
	Phone string `json:"phone,omitempty"`
}

// SavePhoneUserNameWithPermissions DTO.
type SavePhoneUserNameWithPermissions struct {
	SaveUserWithRolesBase
	Phone string `json:"phone,omitempty"`
}

// SaveEmailUserNameWithPermissions DTO.
type SaveEmailUserNameWithPermissions struct {
	SaveUserWithRolesBase
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

// SaveUserNameWithPermissions DTO.
type SaveUserNameWithPermissions struct {
	SaveUserWithRolesBase
	Password string `json:"password,omitempty"`
	UserName string `json:"userName,omitempty"`
}

// DeleteUserRequest DTO.
type DeleteUserRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetUserRequest DTO.
type GetUserRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetUsersRequest DTO.
type GetUsersRequest struct {
	CodeMashListPaginationRequestBase
	DatabaseIntegrationId    string   `json:"databaseIntegrationId,omitempty"`
	IncludePermissions       bool     `json:"includePermissions,omitempty"`
	UserShouldHavePushDevice bool     `json:"userShouldHavePushDevice,omitempty"`
	UserShouldHaveEmail      bool     `json:"userShouldHaveEmail,omitempty"`
	IncludeMeta              bool     `json:"includeMeta,omitempty"`
	RoleNames                []string `json:"roleNames,omitempty"`
	UserIds                  []string `json:"userIds,omitempty"`
}

// GetUserPreferencesRequest DTO.
type GetUserPreferencesRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GrantContactConsentRequest DTO.
type GrantContactConsentRequest struct {
	CodeMashRequestBase
	ContactId   string `json:"contactId,omitempty"`
	Channel     string `json:"channel,omitempty"`
	LawfulBasis string `json:"lawfulBasis,omitempty"`
	Source      string `json:"source,omitempty"`
	EvidenceRef string `json:"evidenceRef,omitempty"`
}

// InviteUserRequest DTO.
type InviteUserRequest struct {
	CodeMashRequestBase
	Email                 string `json:"email,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// LinkIdentityRequest DTO.
type LinkIdentityRequest struct {
	CodeMashRequestBase
	UserId                string `json:"userId,omitempty"`
	Provider              string `json:"provider,omitempty"`
	ProviderToken         string `json:"providerToken,omitempty"`
	EmailToVerify         string `json:"emailToVerify,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// MapAuthToUserRequest DTO.
type MapAuthToUserRequest struct {
	CodeMashRequestBase
	UserId                string `json:"userId,omitempty"`
	AuthId                string `json:"authId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// AssignRolePermissionsRequest DTO.
type AssignRolePermissionsRequest struct {
	CodeMashRequestBase
	Id                    string   `json:"id,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
	Roles                 []string `json:"roles,omitempty"`
}

// SetContactRolesRequest DTO.
type SetContactRolesRequest struct {
	CodeMashRequestBase
	UserId                string   `json:"userId,omitempty"`
	Roles                 []string `json:"roles,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
}

// SetContactTagSubscriptionRequest DTO.
type SetContactTagSubscriptionRequest struct {
	CodeMashRequestBase
	ContactId   string `json:"contactId,omitempty"`
	CommChannel string `json:"commChannel,omitempty"`
	Channel     string `json:"channel,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Subscribed  bool   `json:"subscribed,omitempty"`
}

// UnblockUserRequest DTO.
type UnblockUserRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// UnsubscribeContactRequest DTO.
type UnsubscribeContactRequest struct {
	CodeMashRequestBase
	ContactId string `json:"contactId,omitempty"`
	Channel   string `json:"channel,omitempty"`
	Reason    string `json:"reason,omitempty"`
}

// UpdateUserRequest DTO.
type UpdateUserRequest struct {
	SaveUser
	Id string `json:"id,omitempty"`
}

// UpdateUserPreferencesRequest DTO.
type UpdateUserPreferencesRequest struct {
	CodeMashRequestBase
	Id                        string              `json:"id,omitempty"`
	BlockAllMarketingMessages bool                `json:"blockAllMarketingMessages,omitempty"`
	BlockedTags               map[string][]string `json:"blockedTags,omitempty"`
	DatabaseIntegrationId     string              `json:"databaseIntegrationId,omitempty"`
}

// ChangePasswordRequest DTO.
type ChangePasswordRequest struct {
	CodeMashRequestBase
	CurrentPassword       string `json:"currentPassword,omitempty"`
	NewPassword           string `json:"newPassword,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// RequestPasswordResetRequest DTO.
type RequestPasswordResetRequest struct {
	CodeMashRequestBase
	Email string `json:"email,omitempty"`
}

// ConfirmPasswordResetRequest DTO.
type ConfirmPasswordResetRequest struct {
	CodeMashRequestBase
	Token                 string `json:"token,omitempty"`
	NewPassword           string `json:"newPassword,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// PasskeyAuthenticationOptionsRequest DTO.
type PasskeyAuthenticationOptionsRequest struct {
	CodeMashRequestBase
	Email string `json:"email,omitempty"`
}

// VerifyPasskeyAuthenticationRequest DTO.
type VerifyPasskeyAuthenticationRequest struct {
	CodeMashRequestBase
	CeremonyId        string `json:"ceremonyId,omitempty"`
	AssertionResponse string `json:"assertionResponse,omitempty"`
}

// ListPasskeysRequest DTO.
type ListPasskeysRequest struct {
	CodeMashRequestBase
}

// RenamePasskeyRequest DTO.
type RenamePasskeyRequest struct {
	CodeMashRequestBase
	CredentialId string `json:"credentialId,omitempty"`
	FriendlyName string `json:"friendlyName,omitempty"`
}

// RevokePasskeyRequest DTO.
type RevokePasskeyRequest struct {
	CodeMashRequestBase
	CredentialId string `json:"credentialId,omitempty"`
}

// UseRecoveryCodeRequest DTO.
type UseRecoveryCodeRequest struct {
	CodeMashRequestBase
	Email        string `json:"email,omitempty"`
	RecoveryCode string `json:"recoveryCode,omitempty"`
}

// RequestMagicLinkRequest DTO.
type RequestMagicLinkRequest struct {
	CodeMashRequestBase
	Email string `json:"email,omitempty"`
}

// ConsumeMagicLinkRequest DTO.
type ConsumeMagicLinkRequest struct {
	CodeMashRequestBase
	Token string `json:"token,omitempty"`
}

// HasPasskeyRequest DTO.
type HasPasskeyRequest struct {
	CodeMashRequestBase
	Email string `json:"email,omitempty"`
}

// StartEmailVerificationRequest DTO.
type StartEmailVerificationRequest struct {
	CodeMashRequestBase
	Email string `json:"email,omitempty"`
}

// ConfirmEmailVerificationRequest DTO.
type ConfirmEmailVerificationRequest struct {
	CodeMashRequestBase
	Email string `json:"email,omitempty"`
	Code  string `json:"code,omitempty"`
}

// PasskeyRegistrationOptionsRequest DTO.
type PasskeyRegistrationOptionsRequest struct {
	CodeMashRequestBase
	VerificationToken string `json:"verificationToken,omitempty"`
}

// VerifyPasskeyRegistrationRequest DTO.
type VerifyPasskeyRegistrationRequest struct {
	CodeMashRequestBase
	VerificationToken   string `json:"verificationToken,omitempty"`
	CeremonyId          string `json:"ceremonyId,omitempty"`
	AttestationResponse string `json:"attestationResponse,omitempty"`
	FriendlyName        string `json:"friendlyName,omitempty"`
}

// RefreshPasskeyTokenRequest DTO.
type RefreshPasskeyTokenRequest struct {
	CodeMashRequestBase
	RefreshToken string `json:"refreshToken,omitempty"`
}

// PasskeyLogoutRequest DTO.
type PasskeyLogoutRequest struct {
	CodeMashRequestBase
	RefreshToken string `json:"refreshToken,omitempty"`
}

// FindMergedTermTreeRequest DTO.
type FindMergedTermTreeRequest struct {
	CodeMashRequestBase
	TaxonomyName          string `json:"taxonomyName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// FindTaxonomyTreeRequest DTO.
type FindTaxonomyTreeRequest struct {
	CodeMashRequestBase
	IncludeTerms          bool   `json:"includeTerms,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// FindTermsRequest DTO.
type FindTermsRequest struct {
	CodeMashListPaginationRequestBase
	TaxonomyName          string      `json:"taxonomyName,omitempty"`
	DatabaseIntegrationId string      `json:"databaseIntegrationId,omitempty"`
	Filter                string      `json:"filter,omitempty"`
	SortDescending        bool        `json:"sortDescending,omitempty"`
	PagingArgs            *PagingArgs `json:"pagingArgs,omitempty"`
}

// FindTermsChildrenRequest DTO.
type FindTermsChildrenRequest struct {
	CodeMashListPaginationRequestBase
	TaxonomyName          string      `json:"taxonomyName,omitempty"`
	ParentId              string      `json:"parentId,omitempty"`
	DatabaseIntegrationId string      `json:"databaseIntegrationId,omitempty"`
	Filter                string      `json:"filter,omitempty"`
	PagingArgs            *PagingArgs `json:"pagingArgs,omitempty"`
}

// FindTermTreeRequest DTO.
type FindTermTreeRequest struct {
	CodeMashRequestBase
	TaxonomyName          string  `json:"taxonomyName,omitempty"`
	RootTermId            string  `json:"rootTermId,omitempty"`
	Depth                 float64 `json:"depth,omitempty"`
	DatabaseIntegrationId string  `json:"databaseIntegrationId,omitempty"`
}

// GetDatabaseSchemaRequest DTO.
type GetDatabaseSchemaRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetDatabaseSchemasRequest DTO.
type GetDatabaseSchemasRequest struct {
	CodeMashListPaginationRequestBase
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

// AggregateRequest DTO.
type AggregateRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Pipeline              string `json:"pipeline,omitempty"`
}

// ChangeResponsibilityRequest DTO.
type ChangeResponsibilityRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	NewResponsibleUserId  string `json:"newResponsibleUserId,omitempty"`
}

// CountRequest DTO.
type CountRequest struct {
	CodeMashRequestBase
	CollectionName        string  `json:"collectionName,omitempty"`
	DatabaseIntegrationId string  `json:"databaseIntegrationId,omitempty"`
	Filter                string  `json:"filter,omitempty"`
	SchemaVersion         float64 `json:"schemaVersion,omitempty"`
}

// DeleteManyRequest DTO.
type DeleteManyRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Filter                string `json:"filter,omitempty"`
}

// DeleteOneRequest DTO.
type DeleteOneRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// DistinctRequest DTO.
type DistinctRequest struct {
	CodeMashRequestBase
	CollectionName        string  `json:"collectionName,omitempty"`
	DatabaseIntegrationId string  `json:"databaseIntegrationId,omitempty"`
	Field                 string  `json:"field,omitempty"`
	Filter                string  `json:"filter,omitempty"`
	SchemaVersion         float64 `json:"schemaVersion,omitempty"`
}

// ExecuteAggregateRequest DTO.
type ExecuteAggregateRequest struct {
	CodeMashRequestBase
	CollectionName        string            `json:"collectionName,omitempty"`
	AggregateId           string            `json:"aggregateId,omitempty"`
	DatabaseIntegrationId string            `json:"databaseIntegrationId,omitempty"`
	Tokens                map[string]string `json:"tokens,omitempty"`
}

// FindRequest DTO.
type FindRequest struct {
	CodeMashListPaginationRequestBase
	CollectionName        string      `json:"collectionName,omitempty"`
	DatabaseIntegrationId string      `json:"databaseIntegrationId,omitempty"`
	Filter                string      `json:"filter,omitempty"`
	SchemaVersion         float64     `json:"schemaVersion,omitempty"`
	PagingArgs            *PagingArgs `json:"pagingArgs,omitempty"`
	SortBy                string      `json:"sortBy,omitempty"`
	SortOrder             float64     `json:"sortOrder,omitempty"`
}

// FindOneRequest DTO.
type FindOneRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// FindOwnRequest DTO.
type FindOwnRequest struct {
	CodeMashListPaginationRequestBase
	CollectionName        string      `json:"collectionName,omitempty"`
	DatabaseIntegrationId string      `json:"databaseIntegrationId,omitempty"`
	Filter                string      `json:"filter,omitempty"`
	SchemaVersion         float64     `json:"schemaVersion,omitempty"`
	PagingArgs            *PagingArgs `json:"pagingArgs,omitempty"`
}

// InsertManyRequest DTO.
type InsertManyRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Documents             string `json:"documents,omitempty"`
}

// InsertOneRequest DTO.
type InsertOneRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Document              string `json:"document,omitempty"`
}

// ReplaceOneRequest DTO.
type ReplaceOneRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Replacement           string `json:"replacement,omitempty"`
}

// UpdateManyRequest DTO.
type UpdateManyRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Filter                string `json:"filter,omitempty"`
	Update                string `json:"update,omitempty"`
}

// UpdateOneRequest DTO.
type UpdateOneRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Update                string `json:"update,omitempty"`
}

// CommitUploadRequest DTO.
type CommitUploadRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string  `json:"filesIntegrationId,omitempty"`
	Path               string  `json:"path,omitempty"`
	ContentType        string  `json:"contentType,omitempty"`
	SizeBytes          float64 `json:"sizeBytes,omitempty"`
	FileName           string  `json:"fileName,omitempty"`
}

// GetFileContentRequest DTO.
type GetFileContentRequest struct {
	RequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
	Token              string `json:"token,omitempty"`
}

// PutFileContentRequest DTO.
type PutFileContentRequest struct {
	RequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
	Token              string `json:"token,omitempty"`
}

// DeleteFileApiRequest DTO.
type DeleteFileApiRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
}

// DeleteManyFilesApiRequest DTO.
type DeleteManyFilesApiRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string   `json:"filesIntegrationId,omitempty"`
	Paths__            []string `json:"paths__,omitempty"`
}

// DownloadFileApiRequest DTO.
type DownloadFileApiRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
}

// GetFileInfoRequest DTO.
type GetFileInfoRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
}

// GetSignedUrlRequest DTO.
type GetSignedUrlRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string  `json:"filesIntegrationId,omitempty"`
	Path               string  `json:"path,omitempty"`
	ExpirationSeconds  float64 `json:"expirationSeconds,omitempty"`
}

// ListFilesRequest DTO.
type ListFilesRequest struct {
	CodeMashListPaginationRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
}

// GetPublicFileRequest DTO.
type GetPublicFileRequest struct {
	RequestBase
	PublicId string `json:"publicId,omitempty"`
	Name     string `json:"name,omitempty"`
}

// RequestUploadUrlRequest DTO.
type RequestUploadUrlRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string  `json:"filesIntegrationId,omitempty"`
	Path               string  `json:"path,omitempty"`
	ContentType        string  `json:"contentType,omitempty"`
	ExpirationSeconds  float64 `json:"expirationSeconds,omitempty"`
}

// TestFilesIntegrationRequest DTO.
type TestFilesIntegrationRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
}

// PushIntegrationSaved DTO.
type PushIntegrationSaved struct {
	Integration *PushIntegration `json:"integration,omitempty"`
}

// PushIntegrationRenamed DTO.
type PushIntegrationRenamed struct {
	Id   *IntegrationId `json:"id,omitempty"`
	Name *DisplayName   `json:"name,omitempty"`
	Env  *Env           `json:"env,omitempty"`
}

// PushIntegrationSetAsDefault DTO.
type PushIntegrationSetAsDefault struct {
	Env *Env           `json:"env,omitempty"`
	Id  *IntegrationId `json:"id,omitempty"`
}

// PushIntegrationDeleted DTO.
type PushIntegrationDeleted struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// PushIntegrationEnabled DTO.
type PushIntegrationEnabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// PushIntegrationDisabled DTO.
type PushIntegrationDisabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// PushServiceEstablished DTO.
type PushServiceEstablished struct {
	DefaultTemplates []*PushTemplate `json:"defaultTemplates,omitempty"`
}

// PushServiceEnabled DTO.
type PushServiceEnabled struct {
}

// PushServiceDisabled DTO.
type PushServiceDisabled struct {
}

// PushTemplateCreated DTO.
type PushTemplateCreated struct {
	TemplateId   *TemplateId                               `json:"templateId,omitempty"`
	DisplayName  *DisplayName                              `json:"displayName,omitempty"`
	Translations []MessageTranslation[*PushMessageContent] `json:"translations,omitempty"`
	Channel      CommunicationChannel                      `json:"channel,omitempty"`
	Description  string                                    `json:"description,omitempty"`
	Tags         []*Tag                                    `json:"tags,omitempty"`
	Env          *Env                                      `json:"env,omitempty"`
}

// PushTemplateUpdated DTO.
type PushTemplateUpdated struct {
	TemplateId   *TemplateId                               `json:"templateId,omitempty"`
	DisplayName  *DisplayName                              `json:"displayName,omitempty"`
	Translations []MessageTranslation[*PushMessageContent] `json:"translations,omitempty"`
	Channel      CommunicationChannel                      `json:"channel,omitempty"`
	Description  string                                    `json:"description,omitempty"`
	Tags         []*Tag                                    `json:"tags,omitempty"`
	Env          *Env                                      `json:"env,omitempty"`
}

// PushTemplateDeleted DTO.
type PushTemplateDeleted struct {
	TemplateId *TemplateId `json:"templateId,omitempty"`
	Env        *Env        `json:"env,omitempty"`
}

// PushTemplateArchived DTO.
type PushTemplateArchived struct {
	TemplateId *TemplateId `json:"templateId,omitempty"`
	Env        *Env        `json:"env,omitempty"`
}

// PushTemplateUnArchived DTO.
type PushTemplateUnArchived struct {
	TemplateId *TemplateId `json:"templateId,omitempty"`
	Env        *Env        `json:"env,omitempty"`
}

// PushTemplateMirrored DTO.
type PushTemplateMirrored struct {
	Template *PushTemplate `json:"template,omitempty"`
}

// Authenticate DTO.
type Authenticate struct {
	Provider          string            `json:"provider,omitempty"`
	UserName          string            `json:"userName,omitempty"`
	Password          string            `json:"password,omitempty"`
	RememberMe        bool              `json:"rememberMe,omitempty"`
	AccessToken       string            `json:"accessToken,omitempty"`
	AccessTokenSecret string            `json:"accessTokenSecret,omitempty"`
	ReturnUrl         string            `json:"returnUrl,omitempty"`
	ErrorView         string            `json:"errorView,omitempty"`
	Meta              map[string]string `json:"meta,omitempty"`
}

// GetAccessToken DTO.
type GetAccessToken struct {
	RefreshToken string            `json:"refreshToken,omitempty"`
	Meta         map[string]string `json:"meta,omitempty"`
}
