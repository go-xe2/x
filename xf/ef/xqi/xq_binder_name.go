package xqi

type TBinderName struct {
	name    string
	options map[string]interface{}
}

var (
	MapBinder     = BinderName("map")
	JsonBinder    = BinderName("json")
	XmlBinder     = BinderName("xml")
	DatasetBinder = BinderName("dataset")
	SliceBinder   = BinderName("slice")
)

func BinderName(name string, options ...map[string]interface{}) TBinderName {
	opts := map[string]interface{}{}
	if len(options) > 0 {
		opts = options[0]
	}
	return TBinderName{
		name:    name,
		options: opts,
	}
}

func (bn TBinderName) String() string {
	return bn.name
}

func (bn TBinderName) Name() string {
	return bn.name
}

func (bn TBinderName) Options() map[string]interface{} {
	return bn.options
}
