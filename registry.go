package featuremgmt

// FeatureFlag describes a feature toggle in the Grafana feature management system.
type FeatureFlag struct {
	Name            string
	Description     string
	Stage           string
	Owner           string
	FrontendOnly    bool
	RequiresRestart bool
	Expression      string
}

const (
	FeatureStageExperimental       = "experimental"
	FeatureStagePublicPreview      = "preview"
	FeatureStageGeneralAvailability = "ga"
)

const (
	grafanaObservabilityLogsSquad       = "grafana-observability-logs"
	grafanaFrontendPlatformSquad        = "grafana-frontend-platform"
	grafanaDatavizSquad                 = "grafana-dataviz"
	grafanaAlertingSquad                = "grafana-alerting"
	grafanaDashboardsSquad              = "grafana-dashboards"
	grafanaSearchAndStorageSquad        = "grafana-search-and-storage"
	grafanaPluginsPlatformSquad         = "grafana-plugins-platform"
	grafanaBackendServicesSquad         = "grafana-backend-services"
)

var (
	standardFeatureFlags = []FeatureFlag{
		{
			Name:        "panelTitleSearch",
			Description: "Search for dashboards using panel title",
			Stage:       FeatureStagePublicPreview,
			Owner:       grafanaSearchAndStorageSquad,
			Expression:  "true",
		},
		{
			Name:        "publicDashboards",
			Description: "Enables public access to dashboards",
			Stage:       FeatureStageGeneralAvailability,
			Owner:       grafanaDashboardsSquad,
			Expression:  "true",
		},
		{
			Name:        "liveAPIServer",
			Description: "Registers a live apiserver",
			Stage:       FeatureStageExperimental,
			Owner:       grafanaBackendServicesSquad,
			Expression:  "false",
		},
		{
			Name:        "trimDefaults",
			Description: "Use cue schema to remove default values before storing in database",
			Stage:       FeatureStageGeneralAvailability,
			Owner:       grafanaDashboardsSquad,
			Expression:  "true",
		},
		{
			Name:        "prometheusBufferedClient",
			Description: "Enable buffered client for Prometheus",
			Stage:       FeatureStagePublicPreview,
			Owner:       grafanaObservabilityLogsSquad,
			Expression:  "false",
		},
		{
			Name:        "lokiLogsDataplane",
			Description: "Changes Loki responses to be compliant with the dataplane specification",
			Stage:       FeatureStageGeneralAvailability,
			Owner:       grafanaObservabilityLogsSquad,
			Expression:  "true",
		},
		{
			Name:        "alertingNotificationsPoliciesMatchingInstances",
			Description: "Enables the preview of matching instances for notification policies",
			Stage:       FeatureStagePublicPreview,
			Owner:       grafanaAlertingSquad,
			Expression:  "false",
		},
		{
			Name:        "pluginsDynamicAngularDetectionPatterns",
			Description: "Enables dynamic detection of Angular plugins using patterns",
			Stage:       FeatureStageExperimental,
			Owner:       grafanaPluginsPlatformSquad,
			Expression:  "false",
		},
		{
			Name:        "transformationsVariableSupport",
			Description: "Allows using variables in transformations",
			Stage:       FeatureStageGeneralAvailability,
			Owner:       grafanaDashboardsSquad,
			Expression:  "true",
		},
		{
			Name:            "colorblindThemes",
			Description:     "Enables the new colorblind-friendly themes",
			Stage:           FeatureStageGeneralAvailability,
			Owner:           grafanaFrontendPlatformSquad,
			RequiresRestart: true,
			Expression:      "false",
		},
		{
			Name:        "adhocFilterLabelsFromPanels",
			Description: "Get ad-hoc filter label values from visible panels in the dashboard",
			Stage:       FeatureStageExperimental,
			Owner:       grafanaDashboardsSquad,
			Expression:  "false",
		},
	}
)

// GetFlags returns all registered feature flags.
func GetFlags() []FeatureFlag {
	return standardFeatureFlags
}
