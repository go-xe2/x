package xbinder

import (
	"github.com/go-xe2/x/sync/xsafeMap"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

var dbQueryBinders = xsafeMap.NewStrAnyMap()

func RegQueryBinder(binderName TBinderName, binder DbQueryBinder) {
	BootEntry.Register(binderName.Name(), binder, binderName.Options())
}

func GetQueryBinder(binderName TBinderName) DbQueryBinder {
	return BootEntry.GetBinder(binderName)
}

func registerInternalBinder(entry *xBinderBootEntry) {
	entry.Register(MapBinder.Name(), NewQryMapBinder(nil), MapBinder.Options())
	entry.Register(JsonBinder.Name(), NewQryJsonBinder(nil), JsonBinder.Options())
	entry.Register(XmlBinder.Name(), NewQryXmlBinder(nil), XmlBinder.Options())
	entry.Register(DatasetBinder.Name(), NewQryDatasetBinder(nil), DatasetBinder.Options())
	entry.Register(SliceBinder.Name(), NewQrySliceBinder(nil), SliceBinder.Options())
}
