package resource

var _ IResource = (*AppResource)(nil)

type Configs struct {
	Configs map[string]string `json:"configs"`
} //@name ConfigsResource

type Config struct {
	Config string `json:"config"`
} //@name ConfigResource

type AppResource struct {
}

func (r AppResource) Collection(pagination Pagination, data interface{}) interface{} {
	if items, ok := data.(map[string]string); ok {
		return Configs{Configs: items}
	}
	return data
}

func (r AppResource) Single(data interface{}) interface{} {
	return nil
}
