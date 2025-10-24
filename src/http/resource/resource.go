package resource

type IResource interface {
	Collection(pagination Pagination, data interface{}) interface{}
	Single(data interface{}) interface{}
}
