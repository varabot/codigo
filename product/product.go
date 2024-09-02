// The product package provides data structures to configure a VSCode product.
package product

// This is lightly hand-massaged code generated by using [CodeConvert](https://www.codeconvert.ai/app)
// against the VSCode source file: ./src/vs/base/common/product.ts@1.92.1

type BuiltInExtension struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Repo     string `json:"repo"`
	Metadata any    `json:"metadata"`
}

type ProductWalkthrough struct {
	ID    string                   `json:"id"`
	Steps []ProductWalkthroughStep `json:"steps"`
}

type ProductWalkthroughStep struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	When        string `json:"when"`
	Description string `json:"description"`
	Media       any    `json:"media"`
}

type FeaturedExtension struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImagePath   string `json:"imagePath"`
}

type ConfigurationSyncStore struct {
	URL                     string `json:"url"`
	InsidersURL             string `json:"insidersUrl"`
	StableURL               string `json:"stableUrl"`
	CanSwitch               *bool  `json:"canSwitch,omitempty"`
	AuthenticationProviders map[string]struct {
		Scopes []string `json:"scopes"`
	} `json:"authenticationProviders"`
}

type ExtensionUntrustedWorkspaceSupport struct {
	Default  *bool `json:"default,omitempty"`
	Override *bool `json:"override,omitempty"`
}

type ExtensionVirtualWorkspaceSupport struct {
	Default  *bool `json:"default,omitempty"`
	Override *bool `json:"override,omitempty"`
}

type Configuration struct {
	Version                               string               `json:"version,omitempty"`
	Date                                  *string              `json:"date,omitempty"`
	Quality                               *string              `json:"quality,omitempty"`
	Commit                                *string              `json:"commit,omitempty"`
	NameShort                             string               `json:"nameShort,omitempty"`
	NameLong                              string               `json:"nameLong,omitempty"`
	Win32AppUserModelID                   *string              `json:"win32AppUserModelId,omitempty"`
	Win32MutexName                        *string              `json:"win32MutexName,omitempty"`
	Win32RegValueName                     *string              `json:"win32RegValueName,omitempty"`
	ApplicationName                       string               `json:"applicationName,omitempty"`
	EmbedderIdentifier                    *string              `json:"embedderIdentifier,omitempty"`
	URLProtocol                           string               `json:"urlProtocol,omitempty"`
	DataFolderName                        string               `json:"dataFolderName,omitempty"`
	BuiltInExtensions                     []BuiltInExtension   `json:"builtInExtensions,omitempty"`
	WalkthroughMetadata                   []ProductWalkthrough `json:"walkthroughMetadata,omitempty"`
	FeaturedExtensions                    []FeaturedExtension  `json:"featuredExtensions,omitempty"`
	DownloadURL                           *string              `json:"downloadUrl,omitempty"`
	UpdateURL                             *string              `json:"updateUrl,omitempty"`
	WebURL                                *string              `json:"webUrl,omitempty"`
	WebEndpointURLTemplate                *string              `json:"webEndpointUrlTemplate,omitempty"`
	WebviewContentExternalBaseURLTemplate *string              `json:"webviewContentExternalBaseUrlTemplate,omitempty"`
	Target                                *string              `json:"target,omitempty"`
	NlsCoreBaseURL                        *string              `json:"nlsCoreBaseUrl,omitempty"`
	SettingsSearchBuildID                 *int                 `json:"settingsSearchBuildId,omitempty"`
	SettingsSearchURL                     *string              `json:"settingsSearchUrl,omitempty"`
	TasConfig                             *struct {
		Endpoint                               string `json:"endpoint"`
		TelemetryEventName                     string `json:"telemetryEventName"`
		AssignmentContextTelemetryPropertyName string `json:"assignmentContextTelemetryPropertyName"`
	} `json:"tasConfig,omitempty"`
	ExtensionsGallery *struct {
		ServiceURL          string  `json:"serviceUrl"`
		ServicePPEURL       *string `json:"servicePPEUrl,omitempty"`
		SearchURL           *string `json:"searchUrl,omitempty"`
		ItemURL             string  `json:"itemUrl"`
		PublisherURL        string  `json:"publisherUrl"`
		ResourceURLTemplate string  `json:"resourceUrlTemplate"`
		ControlURL          string  `json:"controlUrl"`
		NlsBaseURL          string  `json:"nlsBaseUrl"`
	} `json:"extensionsGallery,omitempty"`
	ExtensionRecommendations          map[string]ExtensionRecommendations     `json:"extensionRecommendations,omitempty"`
	ConfigBasedExtensionTips          map[string]ConfigBasedExtensionTip      `json:"configBasedExtensionTips,omitempty"`
	ExeBasedExtensionTips             map[string]ExeBasedExtensionTip         `json:"exeBasedExtensionTips,omitempty"`
	RemoteExtensionTips               map[string]RemoteExtensionTip           `json:"remoteExtensionTips,omitempty"`
	VirtualWorkspaceExtensionTips     map[string]VirtualWorkspaceExtensionTip `json:"virtualWorkspaceExtensionTips,omitempty"`
	ExtensionKeywords                 map[string][]string                     `json:"extensionKeywords,omitempty"`
	KeymapExtensionTips               []string                                `json:"keymapExtensionTips,omitempty"`
	WebExtensionTips                  []string                                `json:"webExtensionTips,omitempty"`
	LanguageExtensionTips             []string                                `json:"languageExtensionTips,omitempty"`
	TrustedExtensionURLPublicKeys     map[string][]string                     `json:"trustedExtensionUrlPublicKeys,omitempty"`
	TrustedExtensionAuthAccess        any                                     `json:"trustedExtensionAuthAccess,omitempty"`
	TrustedExtensionProtocolHandlers  []string                                `json:"trustedExtensionProtocolHandlers,omitempty"`
	CommandPaletteSuggestedCommandIDs []string                                `json:"commandPaletteSuggestedCommandIds,omitempty"`
	CrashReporter                     *struct {
		CompanyName string `json:"companyName"`
		ProductName string `json:"productName"`
	} `json:"crashReporter,omitempty"`
	RemoveTelemetryMachineID *bool `json:"removeTelemetryMachineId,omitempty"`
	EnabledTelemetryLevels   struct {
		Error bool `json:"error"`
		Usage bool `json:"usage"`
	} `json:"enabledTelemetryLevels"`
	EnableTelemetry       *bool `json:"enableTelemetry,omitempty"`
	OpenToWelcomeMainPage *bool `json:"openToWelcomeMainPage,omitempty"`
	AiConfig              *struct {
		AriaKey string `json:"ariaKey"`
	} `json:"aiConfig,omitempty"`
	DocumentationURL                        *string                                       `json:"documentationUrl,omitempty"`
	ServerDocumentationURL                  *string                                       `json:"serverDocumentationUrl,omitempty"`
	ReleaseNotesURL                         *string                                       `json:"releaseNotesUrl,omitempty"`
	KeyboardShortcutsURLMac                 *string                                       `json:"keyboardShortcutsUrlMac,omitempty"`
	KeyboardShortcutsURLLinux               *string                                       `json:"keyboardShortcutsUrlLinux,omitempty"`
	KeyboardShortcutsURLWin                 *string                                       `json:"keyboardShortcutsUrlWin,omitempty"`
	IntroductoryVideosURL                   *string                                       `json:"introductoryVideosUrl,omitempty"`
	TipsAndTricksURL                        *string                                       `json:"tipsAndTricksUrl,omitempty"`
	NewsletterSignupURL                     *string                                       `json:"newsletterSignupUrl,omitempty"`
	YouTubeURL                              *string                                       `json:"youTubeUrl,omitempty"`
	RequestFeatureURL                       *string                                       `json:"requestFeatureUrl,omitempty"`
	ReportIssueURL                          *string                                       `json:"reportIssueUrl,omitempty"`
	ReportMarketplaceIssueURL               *string                                       `json:"reportMarketplaceIssueUrl,omitempty"`
	LicenseURL                              *string                                       `json:"licenseUrl,omitempty"`
	ServerLicenseURL                        *string                                       `json:"serverLicenseUrl,omitempty"`
	PrivacyStatementURL                     *string                                       `json:"privacyStatementUrl,omitempty"`
	ShowTelemetryOptOut                     *bool                                         `json:"showTelemetryOptOut,omitempty"`
	ServerGreeting                          []string                                      `json:"serverGreeting,omitempty"`
	ServerLicense                           []string                                      `json:"serverLicense,omitempty"`
	ServerLicensePrompt                     *string                                       `json:"serverLicensePrompt,omitempty"`
	ServerApplicationName                   string                                        `json:"serverApplicationName"`
	ServerDataFolderName                    *string                                       `json:"serverDataFolderName,omitempty"`
	TunnelApplicationName                   *string                                       `json:"tunnelApplicationName,omitempty"`
	TunnelApplicationConfig                 *TunnelApplicationConfig                      `json:"tunnelApplicationConfig,omitempty"`
	NpsSurveyURL                            *string                                       `json:"npsSurveyUrl,omitempty"`
	CesSurveyURL                            *string                                       `json:"cesSurveyUrl,omitempty"`
	Surveys                                 []SurveyData                                  `json:"surveys,omitempty"`
	Checksums                               map[string]string                             `json:"checksums,omitempty"`
	ChecksumFailMoreInfoURL                 *string                                       `json:"checksumFailMoreInfoUrl,omitempty"`
	AppCenter                               *AppCenterConfiguration                       `json:"appCenter,omitempty"`
	Portable                                *string                                       `json:"portable,omitempty"`
	ExtensionKind                           map[string][]string                           `json:"extensionKind,omitempty"`
	ExtensionPointExtensionKind             map[string][]string                           `json:"extensionPointExtensionKind,omitempty"`
	ExtensionSyncedKeys                     map[string][]string                           `json:"extensionSyncedKeys,omitempty"`
	ExtensionsEnabledWithApiProposalVersion []string                                      `json:"extensionsEnabledWithApiProposalVersion,omitempty"`
	ExtensionEnabledApiProposals            map[string][]string                           `json:"extensionEnabledApiProposals,omitempty"`
	ExtensionUntrustedWorkspaceSupport      map[string]ExtensionUntrustedWorkspaceSupport `json:"extensionUntrustedWorkspaceSupport,omitempty"`
	ExtensionVirtualWorkspacesSupport       map[string]ExtensionVirtualWorkspaceSupport   `json:"extensionVirtualWorkspacesSupport,omitempty"`
	MsftInternalDomains                     []string                                      `json:"msftInternalDomains,omitempty"`
	LinkProtectionTrustedDomains            []string                                      `json:"linkProtectionTrustedDomains,omitempty"`
	ConfigurationSyncStore                  *ConfigurationSyncStore                       `json:"configurationSync.store,omitempty"`
	EditSessionsStore                       *ConfigurationSyncStore                       `json:"editSessions.store,omitempty"`
	DarwinUniversalAssetID                  *string                                       `json:"darwinUniversalAssetId,omitempty"`
	ProfileTemplatesURL                     *string                                       `json:"profileTemplatesUrl,omitempty"`
	CommonlyUsedSettings                    []string                                      `json:"commonlyUsedSettings,omitempty"`
	AiGeneratedWorkspaceTrust               *AIGeneratedWorkspaceTrust                    `json:"aiGeneratedWorkspaceTrust,omitempty"`
	GitHubEntitlement                       *GitHubEntitlement                            `json:"gitHubEntitlement,omitempty"`
	ChatParticipantRegistry                 *string                                       `json:"chatParticipantRegistry,omitempty"`
}

type TunnelApplicationConfig struct {
	AuthenticationProviders map[string]struct {
		Scopes []string `json:"scopes"`
	} `json:"authenticationProviders"`
	EditorWebURL string             `json:"editorWebUrl"`
	Extension    RemoteExtensionTip `json:"extension"`
}

type ExtensionRecommendations struct {
	OnFileOpen           []FileOpenCondition          `json:"onFileOpen"`
	OnSettingsEditorOpen *SettingsEditorOpenCondition `json:"onSettingsEditorOpen,omitempty"`
}

type SettingsEditorOpenCondition struct {
	Prerelease          *any    `json:"prerelease,omitempty"`
	DescriptionOverride *string `json:"descriptionOverride,omitempty"`
}

type ExtensionRecommendationCondition struct {
	Important        *bool    `json:"important,omitempty"`
	WhenInstalled    []string `json:"whenInstalled,omitempty"`
	WhenNotInstalled []string `json:"whenNotInstalled,omitempty"`
}

type FileOpenCondition any

type FileLanguageCondition struct {
	Languages []string `json:"languages"`
}

type FilePathCondition struct {
	PathGlob string `json:"pathGlob"`
}

type FileContentCondition struct {
	ExtensionRecommendationCondition
	FileLanguageCondition
	FilePathCondition
	ContentPattern string `json:"contentPattern"`
}

type AppCenterConfiguration struct {
	Win32X64        string `json:"win32-x64"`
	Win32Arm64      string `json:"win32-arm64"`
	LinuxX64        string `json:"linux-x64"`
	Darwin          string `json:"darwin"`
	DarwinUniversal string `json:"darwin-universal"`
	DarwinArm64     string `json:"darwin-arm64"`
}

type ConfigBasedExtensionTip struct {
	ConfigPath      string  `json:"configPath"`
	ConfigName      string  `json:"configName"`
	ConfigScheme    *string `json:"configScheme,omitempty"`
	Recommendations map[string]struct {
		Name             string   `json:"name"`
		ContentPattern   *string  `json:"contentPattern,omitempty"`
		Important        *bool    `json:"important,omitempty"`
		IsExtensionPack  *bool    `json:"isExtensionPack,omitempty"`
		WhenNotInstalled []string `json:"whenNotInstalled,omitempty"`
	} `json:"recommendations"`
}

type ExeBasedExtensionTip struct {
	FriendlyName    string  `json:"friendlyName"`
	WindowsPath     *string `json:"windowsPath,omitempty"`
	Important       *bool   `json:"important,omitempty"`
	Recommendations map[string]struct {
		Name             string   `json:"name"`
		Important        *bool    `json:"important,omitempty"`
		IsExtensionPack  *bool    `json:"isExtensionPack,omitempty"`
		WhenNotInstalled []string `json:"whenNotInstalled,omitempty"`
	} `json:"recommendations"`
}

type RemoteExtensionTip struct {
	FriendlyName       string   `json:"friendlyName"`
	ExtensionID        string   `json:"extensionId"`
	SupportedPlatforms []string `json:"supportedPlatforms,omitempty"`
	StartEntry         struct {
		HelpLink          string `json:"helpLink"`
		StartConnectLabel string `json:"startConnectLabel"`
		StartCommand      string `json:"startCommand"`
		Priority          int    `json:"priority"`
	} `json:"startEntry"`
}

type VirtualWorkspaceExtensionTip struct {
	FriendlyName       string   `json:"friendlyName"`
	ExtensionID        string   `json:"extensionId"`
	SupportedPlatforms []string `json:"supportedPlatforms,omitempty"`
	StartEntry         struct {
		HelpLink          string `json:"helpLink"`
		StartConnectLabel string `json:"startConnectLabel"`
		StartCommand      string `json:"startCommand"`
		Priority          int    `json:"priority"`
	} `json:"startEntry"`
}

type SurveyData struct {
	SurveyID        string  `json:"surveyId"`
	SurveyURL       string  `json:"surveyUrl"`
	LanguageID      string  `json:"languageId"`
	EditCount       int     `json:"editCount"`
	UserProbability float64 `json:"userProbability"`
}

type AIGeneratedWorkspaceTrust struct {
	Title                        string `json:"title"`
	CheckboxText                 string `json:"checkboxText"`
	TrustOption                  string `json:"trustOption"`
	DontTrustOption              string `json:"dontTrustOption"`
	StartupTrustRequestLearnMore string `json:"startupTrustRequestLearnMore"`
}

type GitHubEntitlement struct {
	ProviderID string `json:"providerId"`
	Command    struct {
		Title                   string `json:"title"`
		TitleWithoutPlaceHolder string `json:"titleWithoutPlaceHolder"`
		Action                  string `json:"action"`
		When                    string `json:"when"`
	} `json:"command"`
	EntitlementURL      string `json:"entitlementUrl"`
	ExtensionID         string `json:"extensionId"`
	EnablementKey       string `json:"enablementKey"`
	ConfirmationMessage string `json:"confirmationMessage"`
	ConfirmationAction  string `json:"confirmationAction"`
}
