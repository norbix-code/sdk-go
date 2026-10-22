// Code generated from the Norbix API contract. DO NOT EDIT.
//
// Regenerate with the type generator described in README.md ("Regenerate the
// types"). Editing this file by hand is lost on the next run.

// Package dtos holds the data types of the Norbix API surface. A field
// that holds another generated type is a pointer. A nullable true/false or
// number field is a pointer too (*bool, *float64): nil is left out of the
// request, false and 0 are sent. A required true/false field is always sent.
// Every other field carries omitempty, because the gateway leaves absent
// values out.
package dtos

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
	AuthStatusInvited           AuthStatus = "Invited"
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
	TriggerTypeAi         TriggerType = "Ai"
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

// JsonSchemaFieldDto DTO.
type JsonSchemaFieldDto struct {
	FieldName string `json:"fieldName,omitempty"`
}

// StringFieldDto DTO.
type StringFieldDto struct {
	JsonSchemaFieldDto
	Format           string            `json:"format,omitempty"`
	Pattern          string            `json:"pattern,omitempty"`
	MinLength        *float64          `json:"minLength,omitempty"`
	MaxLength        *float64          `json:"maxLength,omitempty"`
	TranslateOptions map[string]string `json:"translateOptions,omitempty"`
	Default          string            `json:"default,omitempty"`
	Unique           *bool             `json:"unique,omitempty"`
}

// DecimalFieldDto DTO.
type DecimalFieldDto struct {
	JsonSchemaFieldDto
	Minimum    *float64 `json:"minimum,omitempty"`
	Maximum    *float64 `json:"maximum,omitempty"`
	MultipleOf *float64 `json:"multipleOf,omitempty"`
	Default    *float64 `json:"default,omitempty"`
	Unique     *bool    `json:"unique,omitempty"`
}

// CurrencyDefaultDto DTO.
type CurrencyDefaultDto struct {
	Value    float64 `json:"value,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

// CurrencyFieldDto DTO.
type CurrencyFieldDto struct {
	JsonSchemaFieldDto
	AllowedCurrencies []string            `json:"allowedCurrencies,omitempty"`
	MultipleOf        *float64            `json:"multipleOf,omitempty"`
	Minimum           *float64            `json:"minimum,omitempty"`
	Maximum           *float64            `json:"maximum,omitempty"`
	Default           *CurrencyDefaultDto `json:"default,omitempty"`
}

// BooleanFieldDto DTO.
type BooleanFieldDto struct {
	JsonSchemaFieldDto
	Default *bool `json:"default,omitempty"`
}

// DateFieldDto DTO.
type DateFieldDto struct {
	JsonSchemaFieldDto
	Minimum *float64 `json:"minimum,omitempty"`
	Maximum *float64 `json:"maximum,omitempty"`
	Default *float64 `json:"default,omitempty"`
}

// IntegerFieldDto DTO.
type IntegerFieldDto struct {
	JsonSchemaFieldDto
	Minimum *float64 `json:"minimum,omitempty"`
	Maximum *float64 `json:"maximum,omitempty"`
	Default *float64 `json:"default,omitempty"`
	Unique  *bool    `json:"unique,omitempty"`
}

// GeolocationFieldDto DTO.
type GeolocationFieldDto struct {
	JsonSchemaFieldDto
	AllowedTypes []string `json:"allowedTypes,omitempty"`
}

// TagsFieldDto DTO.
type TagsFieldDto struct {
	JsonSchemaFieldDto
	MinItems *float64 `json:"minItems,omitempty"`
	MaxItems *float64 `json:"maxItems,omitempty"`
	Default  []string `json:"default,omitempty"`
}

// FileFieldDto DTO.
type FileFieldDto struct {
	JsonSchemaFieldDto
	Storages        []string `json:"storages,omitempty"`
	MinItems        *float64 `json:"minItems,omitempty"`
	MaxItems        *float64 `json:"maxItems,omitempty"`
	AllowedFileType string   `json:"allowedFileType,omitempty"`
	MaxSizeMb       *float64 `json:"maxSizeMb,omitempty"`
}

// TaxonomySelectionFieldDto DTO.
type TaxonomySelectionFieldDto struct {
	JsonSchemaFieldDto
	TaxonomyId   string `json:"taxonomyId,omitempty"`
	Multiple     bool   `json:"multiple"`
	DisplayField string `json:"displayField,omitempty"`
}

// CollectionSelectionFieldDto DTO.
type CollectionSelectionFieldDto struct {
	JsonSchemaFieldDto
	CollectionId string `json:"collectionId,omitempty"`
	DisplayField string `json:"displayField,omitempty"`
	Multiple     bool   `json:"multiple"`
}

// UserSelectionFieldDto DTO.
type UserSelectionFieldDto struct {
	JsonSchemaFieldDto
	Multiple     bool   `json:"multiple"`
	DisplayField string `json:"displayField,omitempty"`
}

// RoleSelectionFieldDto DTO.
type RoleSelectionFieldDto struct {
	JsonSchemaFieldDto
	Multiple     bool   `json:"multiple"`
	DisplayField string `json:"displayField,omitempty"`
}

// EnumSelectionFieldDto DTO.
type EnumSelectionFieldDto struct {
	JsonSchemaFieldDto
	Values   []string `json:"values,omitempty"`
	Multiple bool     `json:"multiple"`
	Default  []string `json:"default,omitempty"`
}

// ObjectFieldDto DTO.
type ObjectFieldDto struct {
	JsonSchemaFieldDto
	Properties []*JsonSchemaFieldDto `json:"properties,omitempty"`
	Required   []string              `json:"required,omitempty"`
}

// ArrayFieldDto DTO.
type ArrayFieldDto struct {
	JsonSchemaFieldDto
	Items       *JsonSchemaFieldDto `json:"items,omitempty"`
	MinItems    *float64            `json:"minItems,omitempty"`
	MaxItems    *float64            `json:"maxItems,omitempty"`
	UniqueItems *bool               `json:"uniqueItems,omitempty"`
}

// JsonFieldDto DTO.
type JsonFieldDto struct {
	JsonSchemaFieldDto
	MaxBytes *float64 `json:"maxBytes,omitempty"`
}

// RequestBase DTO.
type RequestBase struct {
	TimeZoneId    string `json:"timeZoneId,omitempty"`
	Version       string `json:"version,omitempty"`
	CorrelationId string `json:"correlationId,omitempty"`
}

// IVersionBasedRequest DTO.
type IVersionBasedRequest struct {
}

// IHasCorrelationIdRequest DTO.
type IHasCorrelationIdRequest struct {
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
	BirthDate                 *float64               `json:"birthDate,omitempty"`
	TimeZone                  string                 `json:"timeZone,omitempty"`
	Language                  string                 `json:"language,omitempty"`
	BlockAllMarketingMessages bool                   `json:"blockAllMarketingMessages"`
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
	IgnoreUserRegistersAsRole bool                `json:"ignoreUserRegistersAsRole"`
}

// SaveUserWithRolesBase DTO.
type SaveUserWithRolesBase struct {
	SaveUser
	Roles []string `json:"roles,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

// IPublicRegistrationRequest DTO.
type IPublicRegistrationRequest struct {
}

// CodeMashListPaginationRequestBase DTO.
type CodeMashListPaginationRequestBase struct {
	RequestBase
	ProjectId     string   `json:"projectId,omitempty"`
	Env           string   `json:"env,omitempty"`
	StartingAfter string   `json:"startingAfter,omitempty"`
	EndingBefore  string   `json:"endingBefore,omitempty"`
	PageSize      *float64 `json:"pageSize,omitempty"`
}

// IPasskeyCeremonyRequest DTO.
type IPasskeyCeremonyRequest struct {
}

// CursorArgs DTO.
type CursorArgs struct {
	Field string  `json:"field,omitempty"`
	Order float64 `json:"order,omitempty"`
}

// PagingArgs DTO.
type PagingArgs struct {
	CursorArgs    *CursorArgs `json:"cursorArgs,omitempty"`
	PageSize      *float64    `json:"pageSize,omitempty"`
	StartingAfter string      `json:"startingAfter,omitempty"`
	EndingBefore  string      `json:"endingBefore,omitempty"`
}

// RegisterFileItem DTO.
type RegisterFileItem struct {
	Path        string   `json:"path,omitempty"`
	OwnerUserId string   `json:"ownerUserId,omitempty"`
	ContentType string   `json:"contentType,omitempty"`
	SizeBytes   *float64 `json:"sizeBytes,omitempty"`
	FileName    string   `json:"fileName,omitempty"`
}

// WriteLogEntry DTO.
type WriteLogEntry struct {
	Level         string            `json:"level,omitempty"`
	Title         string            `json:"title,omitempty"`
	Message       string            `json:"message,omitempty"`
	EventCode     string            `json:"eventCode,omitempty"`
	CorrelationId string            `json:"correlationId,omitempty"`
	Timestamp     string            `json:"timestamp,omitempty"`
	Meta          map[string]string `json:"meta,omitempty"`
}

// PaymentOrderLineInput DTO.
type PaymentOrderLineInput struct {
	Sku             string  `json:"sku,omitempty"`
	Title           string  `json:"title,omitempty"`
	Quantity        float64 `json:"quantity,omitempty"`
	UnitAmountMinor float64 `json:"unitAmountMinor,omitempty"`
}

// EchoRegionDto DTO.
type EchoRegionDto struct {
	Code        string `json:"code,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	ApiUrl      string `json:"apiUrl,omitempty"`
	HubUrl      string `json:"hubUrl,omitempty"`
}

// EchoAgentSnippetDto DTO.
type EchoAgentSnippetDto struct {
	Client      string `json:"client,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Recommended bool   `json:"recommended"`
	Auth        string `json:"auth,omitempty"`
	Config      string `json:"config,omitempty"`
}

// EchoAgentDto DTO.
type EchoAgentDto struct {
	McpUrl                 string                 `json:"mcpUrl,omitempty"`
	OAuthMetadataUrl       string                 `json:"oAuthMetadataUrl,omitempty"`
	InstallationType       string                 `json:"installationType,omitempty"`
	OnboardingDocsUrl      string                 `json:"onboardingDocsUrl,omitempty"`
	ToolsUrl               string                 `json:"toolsUrl,omitempty"`
	CliPackage             string                 `json:"cliPackage,omitempty"`
	CliInstallCommand      string                 `json:"cliInstallCommand,omitempty"`
	DeviceAuthorizationUrl string                 `json:"deviceAuthorizationUrl,omitempty"`
	DeviceTokenUrl         string                 `json:"deviceTokenUrl,omitempty"`
	Snippets               []*EchoAgentSnippetDto `json:"snippets,omitempty"`
}

// EchoHostingPlatformDto DTO.
type EchoHostingPlatformDto struct {
	Provider string `json:"provider,omitempty"`
	IsAws    bool   `json:"isAws"`
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
	MinLength      float64  `json:"minLength,omitempty"`
	MaxLength      *float64 `json:"maxLength,omitempty"`
	MinNumbers     *float64 `json:"minNumbers,omitempty"`
	MinUpper       *float64 `json:"minUpper,omitempty"`
	MinLower       *float64 `json:"minLower,omitempty"`
	MinSpecial     *float64 `json:"minSpecial,omitempty"`
	AllowedSpecial string   `json:"allowedSpecial,omitempty"`
}

// PublicAuthDto DTO.
type PublicAuthDto struct {
	SocialProviders []string                 `json:"socialProviders,omitempty"`
	Passkey         bool                     `json:"passkey"`
	Methods         []string                 `json:"methods,omitempty"`
	PasswordPolicy  *PublicPasswordPolicyDto `json:"passwordPolicy,omitempty"`
}

// PublicAiAssistantDto DTO.
type PublicAiAssistantDto struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Welcome string `json:"welcome,omitempty"`
}

// PublicAiChatDto DTO.
type PublicAiChatDto struct {
	Enabled    bool                    `json:"enabled"`
	Assistants []*PublicAiAssistantDto `json:"assistants,omitempty"`
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
	IsSuccess bool        `json:"isSuccess"`
	Errors    []*ErrorDto `json:"errors,omitempty"`
}

// ResponseBase DTO.
type ResponseBase struct {
	ResponseStatus *CodeMashResponseStatus `json:"responseStatus,omitempty"`
}

// EndUserChatAttachment DTO.
type EndUserChatAttachment struct {
	Id           string  `json:"id,omitempty"`
	SessionId    string  `json:"sessionId,omitempty"`
	FileName     string  `json:"fileName,omitempty"`
	ContentType  string  `json:"contentType,omitempty"`
	Kind         string  `json:"kind,omitempty"`
	Size         float64 `json:"size,omitempty"`
	Summary      string  `json:"summary,omitempty"`
	CreatedAtUtc string  `json:"createdAtUtc,omitempty"`
}

// EndUserChatMemoryNote DTO.
type EndUserChatMemoryNote struct {
	Id           string `json:"id,omitempty"`
	SessionId    string `json:"sessionId,omitempty"`
	Kind         string `json:"kind,omitempty"`
	Text         string `json:"text,omitempty"`
	CreatedAtUtc string `json:"createdAtUtc,omitempty"`
}

// EndUserChatAssistant DTO.
type EndUserChatAssistant struct {
	Id             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	WelcomeMessage string `json:"welcomeMessage,omitempty"`
	IsDefault      bool   `json:"isDefault"`
	MemoryEnabled  bool   `json:"memoryEnabled"`
}

// EndUserChatPlan DTO.
type EndUserChatPlan struct {
	Id           string  `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	QuotaUnit    string  `json:"quotaUnit,omitempty"`
	MonthlyQuota float64 `json:"monthlyQuota,omitempty"`
	Used         float64 `json:"used,omitempty"`
	Remaining    float64 `json:"remaining,omitempty"`
	Attachments  bool    `json:"attachments"`
	Rag          bool    `json:"rag"`
	Memory       bool    `json:"memory"`
}

// EndUserChatSession DTO.
type EndUserChatSession struct {
	Id           string  `json:"id,omitempty"`
	AssistantId  string  `json:"assistantId,omitempty"`
	Title        string  `json:"title,omitempty"`
	IsPinned     bool    `json:"isPinned"`
	IsArchived   bool    `json:"isArchived"`
	LastSeq      float64 `json:"lastSeq,omitempty"`
	CreatedAtUtc string  `json:"createdAtUtc,omitempty"`
	UpdatedAtUtc string  `json:"updatedAtUtc,omitempty"`
}

// AiChatEntryWireDto DTO.
type AiChatEntryWireDto struct {
	Kind                 string  `json:"kind,omitempty"`
	Id                   string  `json:"id,omitempty"`
	Seq                  float64 `json:"seq,omitempty"`
	AtUtc                string  `json:"atUtc,omitempty"`
	RefEntryId           string  `json:"refEntryId,omitempty"`
	WorkItemId           string  `json:"workItemId,omitempty"`
	Feedback             string  `json:"feedback,omitempty"`
	FeedbackAtUtc        string  `json:"feedbackAtUtc,omitempty"`
	FeedbackByUserAuthId string  `json:"feedbackByUserAuthId,omitempty"`
}

// EndUserAiToolParameter DTO.
type EndUserAiToolParameter struct {
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Required    bool   `json:"required"`
	Description string `json:"description,omitempty"`
}

// EndUserAiTool DTO.
type EndUserAiTool struct {
	Name                 string                    `json:"name,omitempty"`
	Description          string                    `json:"description,omitempty"`
	Toolsets             []string                  `json:"toolsets,omitempty"`
	RequiresConfirmation bool                      `json:"requiresConfirmation"`
	Parameters           []*EndUserAiToolParameter `json:"parameters,omitempty"`
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
	NeedChangePasswordOnNextLogin bool                  `json:"needChangePasswordOnNextLogin"`
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
	HasMore       bool                   `json:"hasMore"`
	HasPrevious   bool                   `json:"hasPrevious"`
	StartingAfter string                 `json:"startingAfter,omitempty"`
	EndingBefore  string                 `json:"endingBefore,omitempty"`
}

// UserMarketingPreferencesDto DTO.
type UserMarketingPreferencesDto struct {
	BlockAllMarketingMessages bool                   `json:"blockAllMarketingMessages"`
	BlockedTags               map[string][]string    `json:"blockedTags,omitempty"`
	BlockReasons              []MarketingBlockReason `json:"blockReasons,omitempty"`
}

// PasskeyListItemDto DTO.
type PasskeyListItemDto struct {
	CredentialId    string `json:"credentialId,omitempty"`
	FriendlyName    string `json:"friendlyName,omitempty"`
	RegisteredOnUtc string `json:"registeredOnUtc,omitempty"`
	LastUsedOnUtc   string `json:"lastUsedOnUtc,omitempty"`
	IsRevoked       bool   `json:"isRevoked"`
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
	Order        *float64              `json:"order,omitempty"`
	Name         string                `json:"name,omitempty"`
	Names        map[string]string     `json:"names,omitempty"`
	Slug         string                `json:"slug,omitempty"`
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
	Order        *float64              `json:"order,omitempty"`
	Name         string                `json:"name,omitempty"`
	Names        map[string]string     `json:"names,omitempty"`
	Slug         string                `json:"slug,omitempty"`
	Description  string                `json:"description,omitempty"`
	Descriptions map[string]string     `json:"descriptions,omitempty"`
	MultiParents []*TermMultiParentDto `json:"multiParents,omitempty"`
	Meta         map[string]any        `json:"meta,omitempty"`
	Origin       string                `json:"origin,omitempty"`
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

// SchemaUniqueIndexDto DTO.
type SchemaUniqueIndexDto struct {
	Fields []string `json:"fields,omitempty"`
}

// SchemaSettingsDto DTO.
type SchemaSettingsDto struct {
	SoftDelete        bool                    `json:"softDelete"`
	HasRecordOwner    bool                    `json:"hasRecordOwner"`
	Description       string                  `json:"description,omitempty"`
	UniqueIndexes     []*SchemaUniqueIndexDto `json:"uniqueIndexes,omitempty"`
	EndUserOwnRecords *bool                   `json:"endUserOwnRecords,omitempty"`
}

// SchemaEmbedSettingsDto DTO.
type SchemaEmbedSettingsDto struct {
	Enabled                bool     `json:"enabled"`
	Fields                 []string `json:"fields,omitempty"`
	EmbeddingIntegrationId string   `json:"embeddingIntegrationId,omitempty"`
	PerUser                bool     `json:"perUser"`
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
	IsEnabled      bool              `json:"isEnabled"`
	ActivationCode string            `json:"activationCode,omitempty"`
	Order          *float64          `json:"order,omitempty"`
	BreakOnError   bool              `json:"breakOnError"`
	SavedByAuthId  string            `json:"savedByAuthId,omitempty"`
}

// SchemaDto DTO.
type SchemaDto struct {
	ViewId            string                  `json:"viewId,omitempty"`
	SchemaName        string                  `json:"schemaName,omitempty"`
	SchemaSlug        string                  `json:"schemaSlug,omitempty"`
	Version           float64                 `json:"version,omitempty"`
	MetaSchemaVersion float64                 `json:"metaSchemaVersion,omitempty"`
	DataSchema        *DataSchemaDto          `json:"dataSchema,omitempty"`
	VisualSchema      *VisualSchemaDto        `json:"visualSchema,omitempty"`
	PublishedAt       string                  `json:"publishedAt,omitempty"`
	Settings          *SchemaSettingsDto      `json:"settings,omitempty"`
	Embed             *SchemaEmbedSettingsDto `json:"embed,omitempty"`
	Triggers          []*TriggerDto           `json:"triggers,omitempty"`
}

// SchemaListProjection DTO.
type SchemaListProjection struct {
	ViewId            string   `json:"viewId,omitempty"`
	SchemaName        string   `json:"schemaName,omitempty"`
	SchemaTitle       string   `json:"schemaTitle,omitempty"`
	LatestVersion     *float64 `json:"latestVersion,omitempty"`
	HasDraft          bool     `json:"hasDraft"`
	MetaSchemaVersion float64  `json:"metaSchemaVersion,omitempty"`
	Description       string   `json:"description,omitempty"`
	Env               string   `json:"env,omitempty"`
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
	SizeBytes        *float64         `json:"sizeBytes,omitempty"`
	Checksum         *FileChecksumDto `json:"checksum,omitempty"`
}

// FileResourceRefDto DTO.
type FileResourceRefDto struct {
	Resource      *FileResourceDto `json:"resource,omitempty"`
	IntegrationId string           `json:"integrationId,omitempty"`
	Provider      FileProvider     `json:"provider,omitempty"`
	Path          string           `json:"path,omitempty"`
	PublicUrl     string           `json:"publicUrl,omitempty"`
	IsPublic      bool             `json:"isPublic"`
}

// PublicFolderDto DTO.
type PublicFolderDto struct {
	Path      string `json:"path,omitempty"`
	PublicId  string `json:"publicId,omitempty"`
	PublicUrl string `json:"publicUrl,omitempty"`
	Inherited bool   `json:"inherited"`
}

// RegisterFileResult DTO.
type RegisterFileResult struct {
	Index       float64  `json:"index,omitempty"`
	Path        string   `json:"path,omitempty"`
	OwnerUserId string   `json:"ownerUserId,omitempty"`
	Status      string   `json:"status,omitempty"`
	FileId      string   `json:"fileId,omitempty"`
	Verified    *bool    `json:"verified,omitempty"`
	Warning     string   `json:"warning,omitempty"`
	ErrorCode   string   `json:"errorCode,omitempty"`
	ErrorStatus *float64 `json:"errorStatus,omitempty"`
	Error       string   `json:"error,omitempty"`
}

// IntegrationTestResultItemDto DTO.
type IntegrationTestResultItemDto struct {
	Operation string   `json:"operation,omitempty"`
	Result    string   `json:"result,omitempty"`
	Errors    []string `json:"errors,omitempty"`
}

// OwnPushDeviceDto DTO.
type OwnPushDeviceDto struct {
	Id             string `json:"id,omitempty"`
	Provider       string `json:"provider,omitempty"`
	Token          string `json:"token,omitempty"`
	Platform       string `json:"platform,omitempty"`
	OsName         string `json:"osName,omitempty"`
	OsVersion      string `json:"osVersion,omitempty"`
	AppVersion     string `json:"appVersion,omitempty"`
	DeviceName     string `json:"deviceName,omitempty"`
	ModelName      string `json:"modelName,omitempty"`
	Brand          string `json:"brand,omitempty"`
	InstallationId string `json:"installationId,omitempty"`
	Locale         string `json:"locale,omitempty"`
	TimeZone       string `json:"timeZone,omitempty"`
	UpdatedAtUtc   string `json:"updatedAtUtc,omitempty"`
}

// OwnPushMessageDto DTO.
type OwnPushMessageDto struct {
	Id         string `json:"id,omitempty"`
	CampaignId string `json:"campaignId,omitempty"`
	Title      string `json:"title,omitempty"`
	Body       string `json:"body,omitempty"`
	SentAtUtc  string `json:"sentAtUtc,omitempty"`
	IsRead     bool   `json:"isRead"`
	ReadAtUtc  string `json:"readAtUtc,omitempty"`
}

// PaymentOrderLineDto DTO.
type PaymentOrderLineDto struct {
	Sku             string  `json:"sku,omitempty"`
	Title           string  `json:"title,omitempty"`
	Quantity        float64 `json:"quantity,omitempty"`
	UnitAmountMinor float64 `json:"unitAmountMinor,omitempty"`
	AmountMinor     float64 `json:"amountMinor,omitempty"`
}

// PaymentOrderCheckoutDto DTO.
type PaymentOrderCheckoutDto struct {
	Attempt           float64 `json:"attempt,omitempty"`
	IntegrationId     string  `json:"integrationId,omitempty"`
	Provider          string  `json:"provider,omitempty"`
	ProviderOrderId   string  `json:"providerOrderId,omitempty"`
	ProviderReference string  `json:"providerReference,omitempty"`
	PaymentUrl        string  `json:"paymentUrl,omitempty"`
	Method            string  `json:"method,omitempty"`
	StartedOnUtc      string  `json:"startedOnUtc,omitempty"`
	ExpiresAtUtc      string  `json:"expiresAtUtc,omitempty"`
}

// PaymentOrderDto DTO.
type PaymentOrderDto struct {
	Id                string                     `json:"id,omitempty"`
	Env               string                     `json:"env,omitempty"`
	BuyerUserId       string                     `json:"buyerUserId,omitempty"`
	MerchantReference string                     `json:"merchantReference,omitempty"`
	Currency          string                     `json:"currency,omitempty"`
	PriceVersion      string                     `json:"priceVersion,omitempty"`
	Lines             []*PaymentOrderLineDto     `json:"lines,omitempty"`
	TotalMinor        float64                    `json:"totalMinor,omitempty"`
	Metadata          map[string]string          `json:"metadata,omitempty"`
	Status            string                     `json:"status,omitempty"`
	CreatedOnUtc      string                     `json:"createdOnUtc,omitempty"`
	UpdatedOnUtc      string                     `json:"updatedOnUtc,omitempty"`
	Checkouts         []*PaymentOrderCheckoutDto `json:"checkouts,omitempty"`
	SettledMinor      *float64                   `json:"settledMinor,omitempty"`
	RefundedMinor     float64                    `json:"refundedMinor,omitempty"`
	PaidOnUtc         string                     `json:"paidOnUtc,omitempty"`
	ProviderPaymentId string                     `json:"providerPaymentId,omitempty"`
	Mismatches        []string                   `json:"mismatches,omitempty"`
}

// InAppNotificationDto DTO.
type InAppNotificationDto struct {
	Id           string `json:"id,omitempty"`
	ProjectId    string `json:"projectId,omitempty"`
	UserId       string `json:"userId,omitempty"`
	Channel      string `json:"channel,omitempty"`
	EventName    string `json:"eventName,omitempty"`
	PayloadJson  string `json:"payloadJson,omitempty"`
	IsRead       bool   `json:"isRead"`
	SeenAtUtc    string `json:"seenAtUtc,omitempty"`
	ReadAtUtc    string `json:"readAtUtc,omitempty"`
	CreatedAtUtc string `json:"createdAtUtc,omitempty"`
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

// IBindableContract DTO.
type IBindableContract struct {
}

// IHasViewId DTO.
type IHasViewId struct {
}

// ICursorArgs DTO.
type ICursorArgs struct {
}

// IQueuedTrigger DTO.
type IQueuedTrigger struct {
}

// EchoResponse DTO.
type EchoResponse struct {
	Release                     CodeMashRelease         `json:"release,omitempty"`
	Runtime                     CodeMashRuntime         `json:"runtime,omitempty"`
	ManagedServiceHubUrl        string                  `json:"managedServiceHubUrl,omitempty"`
	ManagedServiceApiUrl        string                  `json:"managedServiceApiUrl,omitempty"`
	HubUrl                      string                  `json:"hubUrl,omitempty"`
	ApiUrl                      string                  `json:"apiUrl,omitempty"`
	ApiVersion                  string                  `json:"apiVersion,omitempty"`
	HubVersion                  string                  `json:"hubVersion,omitempty"`
	AdminUrlTemplate            string                  `json:"adminUrlTemplate,omitempty"`
	EmailServiceConfigured      bool                    `json:"emailServiceConfigured"`
	RootBootstrapPasswordSource string                  `json:"rootBootstrapPasswordSource,omitempty"`
	Regions                     []*EchoRegionDto        `json:"regions,omitempty"`
	Agent                       *EchoAgentDto           `json:"agent,omitempty"`
	HostingPlatform             *EchoHostingPlatformDto `json:"hostingPlatform,omitempty"`
}

// PublicProjectConfigDto DTO.
type PublicProjectConfigDto struct {
	DisplayName        string           `json:"displayName,omitempty"`
	AdminPortalEnabled bool             `json:"adminPortalEnabled"`
	Branding           *PublicBrandDto  `json:"branding,omitempty"`
	Auth               *PublicAuthDto   `json:"auth,omitempty"`
	AiChat             *PublicAiChatDto `json:"aiChat,omitempty"`
}

// PublicLegalDocumentDto DTO.
type PublicLegalDocumentDto struct {
	Kind      string `json:"kind,omitempty"`
	Title     string `json:"title,omitempty"`
	Body      string `json:"body,omitempty"`
	Available bool   `json:"available"`
}

// IdResponse DTO.
type IdResponse struct {
	ResponseBase
	Id       string   `json:"id,omitempty"`
	Status   string   `json:"status,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

// ListEndUserChatAttachmentsResponse DTO.
type ListEndUserChatAttachmentsResponse struct {
	ResponseBase
	Attachments []*EndUserChatAttachment `json:"attachments,omitempty"`
}

// EmptyResponse DTO.
type EmptyResponse struct {
	ResponseBase
}

// ListEndUserChatMemoryResponse DTO.
type ListEndUserChatMemoryResponse struct {
	ResponseBase
	Notes []*EndUserChatMemoryNote `json:"notes,omitempty"`
}

// GetEndUserChatAvailabilityResponse DTO.
type GetEndUserChatAvailabilityResponse struct {
	ResponseBase
	Enabled            bool                    `json:"enabled"`
	Available          bool                    `json:"available"`
	Reason             string                  `json:"reason,omitempty"`
	DefaultAssistantId string                  `json:"defaultAssistantId,omitempty"`
	Assistants         []*EndUserChatAssistant `json:"assistants,omitempty"`
	Plan               *EndUserChatPlan        `json:"plan,omitempty"`
}

// ListEndUserChatSessionsResponse DTO.
type ListEndUserChatSessionsResponse struct {
	ResponseBase
	Sessions []*EndUserChatSession `json:"sessions,omitempty"`
}

// GetEndUserChatSessionResponse DTO.
type GetEndUserChatSessionResponse struct {
	ResponseBase
	Session *EndUserChatSession `json:"session,omitempty"`
}

// GetEndUserChatEntriesResponse DTO.
type GetEndUserChatEntriesResponse struct {
	ResponseBase
	SessionId string                `json:"sessionId,omitempty"`
	Entries   []*AiChatEntryWireDto `json:"entries,omitempty"`
	LastSeq   float64               `json:"lastSeq,omitempty"`
	HasMore   bool                  `json:"hasMore"`
}

// StartEndUserChatTurnResponse DTO.
type StartEndUserChatTurnResponse struct {
	ResponseBase
	TurnId    string `json:"turnId,omitempty"`
	SessionId string `json:"sessionId,omitempty"`
	Channel   string `json:"channel,omitempty"`
}

// GetEndUserAiToolsResponse DTO.
type GetEndUserAiToolsResponse struct {
	ResponseBase
	Tools []*EndUserAiTool `json:"tools,omitempty"`
}

// InvokeEndUserAiToolResponse DTO.
type InvokeEndUserAiToolResponse struct {
	ResponseBase
	Result string `json:"result,omitempty"`
}

// SaveGuestUserResponse DTO.
type SaveGuestUserResponse struct {
	IdResponse
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
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

// SendTestPasswordResetResponse DTO.
type SendTestPasswordResetResponse struct {
	ResponseBase
	Queued             bool     `json:"queued"`
	Reason             string   `json:"reason,omitempty"`
	Env                string   `json:"env,omitempty"`
	CorrelationId      string   `json:"correlationId,omitempty"`
	ExpiresAtUtc       string   `json:"expiresAtUtc,omitempty"`
	IntegrationId      string   `json:"integrationId,omitempty"`
	IntegrationName    string   `json:"integrationName,omitempty"`
	Provider           string   `json:"provider,omitempty"`
	IntegrationProblem string   `json:"integrationProblem,omitempty"`
	TemplateId         string   `json:"templateId,omitempty"`
	TemplateName       string   `json:"templateName,omitempty"`
	TemplateSource     string   `json:"templateSource,omitempty"`
	Language           string   `json:"language,omitempty"`
	LanguageRule       string   `json:"languageRule,omitempty"`
	TemplateLanguages  []string `json:"templateLanguages,omitempty"`
	Fallback           string   `json:"fallback,omitempty"`
	MjmlRenderer       string   `json:"mjmlRenderer,omitempty"`
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

// GetFileByIdResponse DTO.
type GetFileByIdResponse struct {
	ResponseBase
	File      *FileResourceRefDto `json:"file,omitempty"`
	IsPublic  *bool               `json:"isPublic,omitempty"`
	PublicUrl string              `json:"publicUrl,omitempty"`
}

// GetFileInfoResponse DTO.
type GetFileInfoResponse struct {
	ResponseBase
	File      *FileResourceRefDto `json:"file,omitempty"`
	IsPublic  *bool               `json:"isPublic,omitempty"`
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

// RegisterFilesResponse DTO.
type RegisterFilesResponse struct {
	ResponseBase
	DryRun     bool                  `json:"dryRun"`
	Registered float64               `json:"registered,omitempty"`
	Unchanged  float64               `json:"unchanged,omitempty"`
	Failed     float64               `json:"failed,omitempty"`
	Items      []*RegisterFileResult `json:"items,omitempty"`
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

// GetOwnPushDevicesResponse DTO.
type GetOwnPushDevicesResponse struct {
	ResponseBase
	Items []*OwnPushDeviceDto `json:"items,omitempty"`
}

// GetOwnPushMessagesResponse DTO.
type GetOwnPushMessagesResponse struct {
	ResponseBase
	List PaginatedResponse[*OwnPushMessageDto] `json:"list,omitempty"`
}

// WriteLogsResponse DTO.
type WriteLogsResponse struct {
	ResponseBase
	Accepted float64 `json:"accepted,omitempty"`
}

// GetPaymentOrderResponse DTO.
type GetPaymentOrderResponse struct {
	ResponseBase
	Item *PaymentOrderDto `json:"item,omitempty"`
}

// GetPaymentOrdersResponse DTO.
type GetPaymentOrdersResponse struct {
	ResponseBase
	List PaginatedResponse[*PaymentOrderDto] `json:"list,omitempty"`
}

// StartPaymentCheckoutResponse DTO.
type StartPaymentCheckoutResponse struct {
	ResponseBase
	Checkout *PaymentOrderCheckoutDto `json:"checkout,omitempty"`
	Reused   bool                     `json:"reused"`
	Order    *PaymentOrderDto         `json:"order,omitempty"`
}

// GetNotificationInboxResponse DTO.
type GetNotificationInboxResponse struct {
	ResponseBase
	List PaginatedResponse[*InAppNotificationDto] `json:"list,omitempty"`
}

// GetUnreadCountResponse DTO.
type GetUnreadCountResponse struct {
	ResponseBase
	Count float64 `json:"count,omitempty"`
}

// MarkNotificationReadResponse DTO.
type MarkNotificationReadResponse struct {
	ResponseBase
}

// MarkAllNotificationsReadResponse DTO.
type MarkAllNotificationsReadResponse struct {
	ResponseBase
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

// SharedTypeGen DTO.
type SharedTypeGen struct {
	Typegen_StringFieldDto              *StringFieldDto              `json:"typegen_StringFieldDto,omitempty"`
	Typegen_DecimalFieldDto             *DecimalFieldDto             `json:"typegen_DecimalFieldDto,omitempty"`
	Typegen_CurrencyFieldDto            *CurrencyFieldDto            `json:"typegen_CurrencyFieldDto,omitempty"`
	Typegen_BooleanFieldDto             *BooleanFieldDto             `json:"typegen_BooleanFieldDto,omitempty"`
	Typegen_DateFieldDto                *DateFieldDto                `json:"typegen_DateFieldDto,omitempty"`
	Typegen_IntegerFieldDto             *IntegerFieldDto             `json:"typegen_IntegerFieldDto,omitempty"`
	Typegen_GeolocationFieldDto         *GeolocationFieldDto         `json:"typegen_GeolocationFieldDto,omitempty"`
	Typegen_TagsFieldDto                *TagsFieldDto                `json:"typegen_TagsFieldDto,omitempty"`
	Typegen_FileFieldDto                *FileFieldDto                `json:"typegen_FileFieldDto,omitempty"`
	Typegen_TaxonomySelectionFieldDto   *TaxonomySelectionFieldDto   `json:"typegen_TaxonomySelectionFieldDto,omitempty"`
	Typegen_CollectionSelectionFieldDto *CollectionSelectionFieldDto `json:"typegen_CollectionSelectionFieldDto,omitempty"`
	Typegen_UserSelectionFieldDto       *UserSelectionFieldDto       `json:"typegen_UserSelectionFieldDto,omitempty"`
	Typegen_RoleSelectionFieldDto       *RoleSelectionFieldDto       `json:"typegen_RoleSelectionFieldDto,omitempty"`
	Typegen_EnumSelectionFieldDto       *EnumSelectionFieldDto       `json:"typegen_EnumSelectionFieldDto,omitempty"`
	Typegen_ObjectFieldDto              *ObjectFieldDto              `json:"typegen_ObjectFieldDto,omitempty"`
	Typegen_ArrayFieldDto               *ArrayFieldDto               `json:"typegen_ArrayFieldDto,omitempty"`
	Typegen_JsonFieldDto                *JsonFieldDto                `json:"typegen_JsonFieldDto,omitempty"`
}

// Echo DTO.
type Echo struct {
	RequestBase
}

// GetPublicProjectBrandAsset DTO.
type GetPublicProjectBrandAsset struct {
	RequestBase
	ProjectId string `json:"projectId,omitempty"`
	Kind      string `json:"kind,omitempty"`
	V         string `json:"v,omitempty"`
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

// UploadEndUserChatAttachmentRequest DTO.
type UploadEndUserChatAttachmentRequest struct {
	CodeMashRequestBase
	SessionId     string `json:"sessionId,omitempty"`
	FileName      string `json:"fileName,omitempty"`
	ContentType   string `json:"contentType,omitempty"`
	Base64Content string `json:"base64Content,omitempty"`
}

// ListEndUserChatAttachmentsRequest DTO.
type ListEndUserChatAttachmentsRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId,omitempty"`
}

// DeleteEndUserChatAttachmentRequest DTO.
type DeleteEndUserChatAttachmentRequest struct {
	CodeMashRequestBase
	AttachmentId string `json:"attachmentId,omitempty"`
}

// SetEndUserChatEntryFeedbackRequest DTO.
type SetEndUserChatEntryFeedbackRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId,omitempty"`
	EntryId   string `json:"entryId,omitempty"`
	Feedback  string `json:"feedback,omitempty"`
}

// ListEndUserChatMemoryRequest DTO.
type ListEndUserChatMemoryRequest struct {
	CodeMashRequestBase
	Take *float64 `json:"take,omitempty"`
}

// ForgetEndUserChatMemoryRequest DTO.
type ForgetEndUserChatMemoryRequest struct {
	CodeMashRequestBase
	NoteId string `json:"noteId,omitempty"`
}

// GetEndUserChatAvailabilityRequest DTO.
type GetEndUserChatAvailabilityRequest struct {
	CodeMashRequestBase
}

// ListEndUserChatSessionsRequest DTO.
type ListEndUserChatSessionsRequest struct {
	CodeMashRequestBase
	Take            *float64 `json:"take,omitempty"`
	IncludeArchived bool     `json:"includeArchived"`
}

// CreateEndUserChatSessionRequest DTO.
type CreateEndUserChatSessionRequest struct {
	CodeMashRequestBase
	AssistantId string `json:"assistantId,omitempty"`
	Title       string `json:"title,omitempty"`
}

// GetEndUserChatSessionRequest DTO.
type GetEndUserChatSessionRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId,omitempty"`
}

// RenameEndUserChatSessionRequest DTO.
type RenameEndUserChatSessionRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId,omitempty"`
	Title     string `json:"title,omitempty"`
}

// PinEndUserChatSessionRequest DTO.
type PinEndUserChatSessionRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId,omitempty"`
	Pinned    bool   `json:"pinned"`
}

// ArchiveEndUserChatSessionRequest DTO.
type ArchiveEndUserChatSessionRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId,omitempty"`
	Archived  bool   `json:"archived"`
}

// DeleteEndUserChatSessionRequest DTO.
type DeleteEndUserChatSessionRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId,omitempty"`
}

// GetEndUserChatEntriesRequest DTO.
type GetEndUserChatEntriesRequest struct {
	CodeMashRequestBase
	SessionId string   `json:"sessionId,omitempty"`
	AfterSeq  *float64 `json:"afterSeq,omitempty"`
	Take      *float64 `json:"take,omitempty"`
}

// StartEndUserChatTurnRequest DTO.
type StartEndUserChatTurnRequest struct {
	CodeMashRequestBase
	SessionId   string `json:"sessionId,omitempty"`
	AssistantId string `json:"assistantId,omitempty"`
	Message     string `json:"message,omitempty"`
}

// GetEndUserAiToolsRequest DTO.
type GetEndUserAiToolsRequest struct {
	RequestBase
}

// InvokeEndUserAiToolRequest DTO.
type InvokeEndUserAiToolRequest struct {
	RequestBase
	ToolName      string `json:"toolName,omitempty"`
	ArgumentsJson string `json:"argumentsJson,omitempty"`
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

// SavePhoneUser DTO.
type SavePhoneUser struct {
	SaveUser
	Phone string `json:"phone,omitempty"`
}

// SaveUserNameUser DTO.
type SaveUserNameUser struct {
	SaveUser
	Password string   `json:"password,omitempty"`
	UserName string   `json:"userName,omitempty"`
	Roles    []string `json:"roles,omitempty"`
}

// SaveEmailUser DTO.
type SaveEmailUser struct {
	SaveUser
	Password string   `json:"password,omitempty"`
	Email    string   `json:"email,omitempty"`
	Roles    []string `json:"roles,omitempty"`
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
	IncludePermissions       bool     `json:"includePermissions"`
	UserShouldHavePushDevice bool     `json:"userShouldHavePushDevice"`
	UserShouldHaveEmail      bool     `json:"userShouldHaveEmail"`
	IncludeMeta              bool     `json:"includeMeta"`
	RoleNames                []string `json:"roleNames,omitempty"`
	UserIds                  []string `json:"userIds,omitempty"`
	Search                   string   `json:"search,omitempty"`
	Email                    string   `json:"email,omitempty"`
	UserName                 string   `json:"userName,omitempty"`
	Statuses                 []string `json:"statuses,omitempty"`
	Types                    []string `json:"types,omitempty"`
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
	FullName              string `json:"fullName,omitempty"`
	Salutation            string `json:"salutation,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// FinishInvitationRequest DTO.
type FinishInvitationRequest struct {
	CodeMashRequestBase
	Token       string `json:"token,omitempty"`
	Password    string `json:"password,omitempty"`
	FullName    string `json:"fullName,omitempty"`
	Salutation  string `json:"salutation,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
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
	Subscribed  bool   `json:"subscribed"`
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
	BlockAllMarketingMessages bool                `json:"blockAllMarketingMessages"`
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

// SendTestPasswordResetRequest DTO.
type SendTestPasswordResetRequest struct {
	CodeMashRequestBase
	Email string `json:"email,omitempty"`
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
	AllLanguages          bool   `json:"allLanguages"`
	TaxonomyName          string `json:"taxonomyName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// FindTaxonomyTreeRequest DTO.
type FindTaxonomyTreeRequest struct {
	CodeMashRequestBase
	AllLanguages          bool   `json:"allLanguages"`
	IncludeTerms          bool   `json:"includeTerms"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// FindTermsRequest DTO.
type FindTermsRequest struct {
	CodeMashListPaginationRequestBase
	AllLanguages          bool        `json:"allLanguages"`
	TaxonomyName          string      `json:"taxonomyName,omitempty"`
	DatabaseIntegrationId string      `json:"databaseIntegrationId,omitempty"`
	Filter                string      `json:"filter,omitempty"`
	SortDescending        bool        `json:"sortDescending"`
	PagingArgs            *PagingArgs `json:"pagingArgs,omitempty"`
}

// FindTermsChildrenRequest DTO.
type FindTermsChildrenRequest struct {
	CodeMashListPaginationRequestBase
	AllLanguages          bool        `json:"allLanguages"`
	TaxonomyName          string      `json:"taxonomyName,omitempty"`
	ParentId              string      `json:"parentId,omitempty"`
	DatabaseIntegrationId string      `json:"databaseIntegrationId,omitempty"`
	Filter                string      `json:"filter,omitempty"`
	PagingArgs            *PagingArgs `json:"pagingArgs,omitempty"`
}

// FindTermTreeRequest DTO.
type FindTermTreeRequest struct {
	CodeMashRequestBase
	AllLanguages          bool     `json:"allLanguages"`
	TaxonomyName          string   `json:"taxonomyName,omitempty"`
	RootTermId            string   `json:"rootTermId,omitempty"`
	Depth                 *float64 `json:"depth,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
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
	CollectionName        string   `json:"collectionName,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
	Filter                string   `json:"filter,omitempty"`
	SchemaVersion         *float64 `json:"schemaVersion,omitempty"`
}

// DeleteManyRequest DTO.
type DeleteManyRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Filter                string `json:"filter,omitempty"`
	AllRecords            *bool  `json:"allRecords,omitempty"`
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
	CollectionName        string   `json:"collectionName,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
	Field                 string   `json:"field,omitempty"`
	Filter                string   `json:"filter,omitempty"`
	SchemaVersion         *float64 `json:"schemaVersion,omitempty"`
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
	SchemaVersion         *float64    `json:"schemaVersion,omitempty"`
	PagingArgs            *PagingArgs `json:"pagingArgs,omitempty"`
	SortBy                string      `json:"sortBy,omitempty"`
	SortOrder             *float64    `json:"sortOrder,omitempty"`
	ExpandReferences      bool        `json:"expandReferences"`
	AllLanguages          bool        `json:"allLanguages"`
}

// FindOneRequest DTO.
type FindOneRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	ExpandReferences      bool   `json:"expandReferences"`
	AllLanguages          bool   `json:"allLanguages"`
}

// FindOwnRequest DTO.
type FindOwnRequest struct {
	CodeMashListPaginationRequestBase
	CollectionName        string      `json:"collectionName,omitempty"`
	DatabaseIntegrationId string      `json:"databaseIntegrationId,omitempty"`
	Filter                string      `json:"filter,omitempty"`
	SchemaVersion         *float64    `json:"schemaVersion,omitempty"`
	PagingArgs            *PagingArgs `json:"pagingArgs,omitempty"`
	ExpandReferences      bool        `json:"expandReferences"`
	AllLanguages          bool        `json:"allLanguages"`
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
	AllRecords            *bool  `json:"allRecords,omitempty"`
	Update                string `json:"update,omitempty"`
	ArrayFilters          string `json:"arrayFilters,omitempty"`
}

// UpdateOneRequest DTO.
type UpdateOneRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Update                string `json:"update,omitempty"`
	ArrayFilters          string `json:"arrayFilters,omitempty"`
}

// CommitUploadRequest DTO.
type CommitUploadRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string   `json:"filesIntegrationId,omitempty"`
	Path               string   `json:"path,omitempty"`
	ContentType        string   `json:"contentType,omitempty"`
	SizeBytes          *float64 `json:"sizeBytes,omitempty"`
	FileName           string   `json:"fileName,omitempty"`
	OwnerUserId        string   `json:"ownerUserId,omitempty"`
	ReassignOwner      *bool    `json:"reassignOwner,omitempty"`
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

// GetFileByIdRequest DTO.
type GetFileByIdRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Id                 string `json:"id,omitempty"`
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
	FilesIntegrationId string   `json:"filesIntegrationId,omitempty"`
	Path               string   `json:"path,omitempty"`
	ExpirationSeconds  *float64 `json:"expirationSeconds,omitempty"`
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

// RegisterFilesRequest DTO.
type RegisterFilesRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string              `json:"filesIntegrationId,omitempty"`
	Items              []*RegisterFileItem `json:"items,omitempty"`
	ReassignOwner      *bool               `json:"reassignOwner,omitempty"`
	DryRun             *bool               `json:"dryRun,omitempty"`
}

// RequestUploadUrlRequest DTO.
type RequestUploadUrlRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string   `json:"filesIntegrationId,omitempty"`
	Path               string   `json:"path,omitempty"`
	ContentType        string   `json:"contentType,omitempty"`
	ExpirationSeconds  *float64 `json:"expirationSeconds,omitempty"`
	OwnerUserId        string   `json:"ownerUserId,omitempty"`
	ReassignOwner      *bool    `json:"reassignOwner,omitempty"`
}

// TestFilesIntegrationRequest DTO.
type TestFilesIntegrationRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
}

// UseFileRequest DTO.
type UseFileRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
	OwnerUserId        string `json:"ownerUserId,omitempty"`
	ReassignOwner      *bool  `json:"reassignOwner,omitempty"`
}

// RegisterOwnPushDeviceRequest DTO.
type RegisterOwnPushDeviceRequest struct {
	CodeMashRequestBase
	Provider       string `json:"provider,omitempty"`
	Token          string `json:"token,omitempty"`
	Platform       string `json:"platform,omitempty"`
	OsName         string `json:"osName,omitempty"`
	OsVersion      string `json:"osVersion,omitempty"`
	AppVersion     string `json:"appVersion,omitempty"`
	DeviceName     string `json:"deviceName,omitempty"`
	ModelName      string `json:"modelName,omitempty"`
	Brand          string `json:"brand,omitempty"`
	Manufacturer   string `json:"manufacturer,omitempty"`
	InstallationId string `json:"installationId,omitempty"`
	Locale         string `json:"locale,omitempty"`
	TimeZone       string `json:"timeZone,omitempty"`
}

// GetOwnPushDevicesRequest DTO.
type GetOwnPushDevicesRequest struct {
	CodeMashRequestBase
}

// DeleteOwnPushDeviceRequest DTO.
type DeleteOwnPushDeviceRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DeleteOwnPushDeviceByTokenRequest DTO.
type DeleteOwnPushDeviceByTokenRequest struct {
	CodeMashRequestBase
	Token string `json:"token,omitempty"`
}

// GetOwnPushMessagesRequest DTO.
type GetOwnPushMessagesRequest struct {
	CodeMashListPaginationRequestBase
	UnreadOnly bool `json:"unreadOnly"`
}

// MarkOwnPushMessageReadRequest DTO.
type MarkOwnPushMessageReadRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// MarkAllOwnPushMessagesReadRequest DTO.
type MarkAllOwnPushMessagesReadRequest struct {
	CodeMashRequestBase
}

// WriteLogs DTO.
type WriteLogs struct {
	CodeMashRequestBase
	Entries []*WriteLogEntry `json:"entries,omitempty"`
}

// CreatePaymentOrderRequest DTO.
type CreatePaymentOrderRequest struct {
	CodeMashRequestBase
	BuyerUserId       string                   `json:"buyerUserId,omitempty"`
	MerchantReference string                   `json:"merchantReference,omitempty"`
	IdempotencyKey    string                   `json:"idempotencyKey,omitempty"`
	Currency          string                   `json:"currency,omitempty"`
	PriceVersion      string                   `json:"priceVersion,omitempty"`
	Lines             []*PaymentOrderLineInput `json:"lines,omitempty"`
	Metadata          map[string]string        `json:"metadata,omitempty"`
}

// GetPaymentOrderRequest DTO.
type GetPaymentOrderRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetPaymentOrdersRequest DTO.
type GetPaymentOrdersRequest struct {
	CodeMashListPaginationRequestBase
	BuyerUserId       string `json:"buyerUserId,omitempty"`
	MerchantReference string `json:"merchantReference,omitempty"`
	Status            string `json:"status,omitempty"`
}

// StartPaymentCheckoutRequest DTO.
type StartPaymentCheckoutRequest struct {
	CodeMashRequestBase
	Id                string `json:"id,omitempty"`
	IntegrationId     string `json:"integrationId,omitempty"`
	Method            string `json:"method,omitempty"`
	ReturnUrl         string `json:"returnUrl,omitempty"`
	PreferredCountry  string `json:"preferredCountry,omitempty"`
	PreferredProvider string `json:"preferredProvider,omitempty"`
}

// GetNotificationInboxRequest DTO.
type GetNotificationInboxRequest struct {
	CodeMashListPaginationRequestBase
	UnreadOnly bool   `json:"unreadOnly"`
	Channel    string `json:"channel,omitempty"`
}

// GetUnreadCountRequest DTO.
type GetUnreadCountRequest struct {
	CodeMashRequestBase
}

// MarkNotificationReadRequest DTO.
type MarkNotificationReadRequest struct {
	CodeMashRequestBase
	NotificationId string `json:"notificationId,omitempty"`
}

// MarkAllNotificationsReadRequest DTO.
type MarkAllNotificationsReadRequest struct {
	CodeMashRequestBase
}

// Authenticate DTO.
type Authenticate struct {
	Provider          string            `json:"provider,omitempty"`
	UserName          string            `json:"userName,omitempty"`
	Password          string            `json:"password,omitempty"`
	RememberMe        *bool             `json:"rememberMe,omitempty"`
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
