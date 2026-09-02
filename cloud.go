package fyne

type CloudProvider interface {
	ProviderDescription() string

	ProviderIcon() Resource

	ProviderName() string

	Cleanup(App)

	Setup(App) error
}

type CloudProviderPreferences interface {
	CloudPreferences(App) Preferences
}

type CloudProviderStorage interface {
	CloudStorage(App) Storage
}
