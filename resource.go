package fyne

type Resource interface {
	Name() string
	Content() []byte
}

type ThemedResource interface {
	Resource
	ThemeColorName() ThemeColorName
}

type StaticResource struct {
	StaticName    string
	StaticContent []byte
}

func (r *StaticResource) Name() string { _ = "STUB: not implemented"; return "" }

func (r *StaticResource) Content() []byte { _ = "STUB: not implemented"; return nil }

func NewStaticResource(name string, content []byte) *StaticResource {
	_ = "STUB: not implemented"
	return nil
}

func LoadResourceFromPath(path string) (Resource, error) {
	_ = "STUB: not implemented"
	return *new(Resource), nil
}

func LoadResourceFromURLString(urlStr string) (Resource, error) {
	_ = "STUB: not implemented"
	return *new(Resource), nil
}

//gosec:disable G107 -- applying security measures to the URL is the caller’s responsibility

func CacheResourceFromURLString(urlStr string) (Resource, error) {
	_ = "STUB: not implemented"
	return *new(Resource), nil
}
