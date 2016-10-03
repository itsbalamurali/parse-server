package models

type App struct {
	AppName string `json:"appName"`
	ApplicationID string `json:"applicationId"`
	ClientClassCreationEnabled bool `json:"clientClassCreationEnabled"`
	ClientPushEnabled bool `json:"clientPushEnabled"`
	DashboardURL string `json:"dashboardURL"`
	JavascriptKey string `json:"javascriptKey"`
	MasterKey string `json:"masterKey"`
	RequireRevocableSessions bool `json:"requireRevocableSessions"`
	RestKey string `json:"restKey"`
	RevokeSessionOnPasswordChange bool `json:"revokeSessionOnPasswordChange"`
	WebhookKey string `json:"webhookKey"`
	WindowsKey string `json:"windowsKey"`
}

