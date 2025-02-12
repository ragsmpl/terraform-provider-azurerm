package resourceproviders

// Required returns all of the Resource Providers used by the AzureRM Provider
// whilst all may not be used by every user - the intention is that we determine which should be
// registered such that we can avoid obscure errors where Resource Providers aren't registered.
// new Resource Providers should be added to this list as they're used in the Provider
// (this is the approach used by Microsoft in their tooling)
func Required() map[string]struct{} {
	// NOTE: Resource Providers in this list are case sensitive
	return map[string]struct{}{
		"Microsoft.ApiManagement":           {},
		"Microsoft.AppPlatform":             {},
		"Microsoft.Authorization":           {},
		"Microsoft.Automation":              {},
		"Microsoft.AVS":                     {},
		"Microsoft.Blueprint":               {},
		"Microsoft.BotService":              {},
		"Microsoft.Cache":                   {},
		"Microsoft.Cdn":                     {},
		"Microsoft.CognitiveServices":       {},
		"Microsoft.Compute":                 {},
		"Microsoft.ContainerInstance":       {},
		"Microsoft.ContainerRegistry":       {},
		"Microsoft.ContainerService":        {},
		"Microsoft.CostManagement":          {},
		"Microsoft.CustomProviders":         {},
		"Microsoft.Databricks":              {},
		"Microsoft.DataLakeAnalytics":       {},
		"Microsoft.DataLakeStore":           {},
		"Microsoft.DataMigration":           {},
		"Microsoft.DataProtection":          {},
		"Microsoft.DBforMariaDB":            {},
		"Microsoft.DBforMySQL":              {},
		"Microsoft.DBforPostgreSQL":         {},
		"Microsoft.DesktopVirtualization":   {},
		"Microsoft.Devices":                 {},
		"Microsoft.DevTestLab":              {},
		"Microsoft.DocumentDB":              {},
		"Microsoft.EventGrid":               {},
		"Microsoft.EventHub":                {},
		"Microsoft.HDInsight":               {},
		"Microsoft.HealthcareApis":          {},
		"Microsoft.GuestConfiguration":      {},
		"Microsoft.KeyVault":                {},
		"Microsoft.Kusto":                   {},
		"microsoft.insights":                {},
		"Microsoft.Logic":                   {},
		"Microsoft.MachineLearningServices": {},
		"Microsoft.Maintenance":             {},
		"Microsoft.ManagedIdentity":         {},
		"Microsoft.ManagedServices":         {},
		"Microsoft.Management":              {},
		"Microsoft.Maps":                    {},
		"Microsoft.MarketplaceOrdering":     {},
		"Microsoft.Media":                   {},
		"Microsoft.MixedReality":            {},
		"Microsoft.Network":                 {},
		"Microsoft.NotificationHubs":        {},
		"Microsoft.OperationalInsights":     {},
		"Microsoft.OperationsManagement":    {},
		"Microsoft.PolicyInsights":          {},
		"Microsoft.PowerBIDedicated":        {},
		"Microsoft.Relay":                   {},
		"Microsoft.RecoveryServices":        {},
		"Microsoft.Resources":               {},
		"Microsoft.Search":                  {},
		"Microsoft.Security":                {},
		"Microsoft.SecurityInsights":        {},
		"Microsoft.ServiceBus":              {},
		"Microsoft.ServiceFabric":           {},
		"Microsoft.Sql":                     {},
		"Microsoft.Storage":                 {},
		"Microsoft.StreamAnalytics":         {},
		"Microsoft.TimeSeriesInsights":      {},
		"Microsoft.Web":                     {},
	}
}
