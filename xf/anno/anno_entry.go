package anno

import "github.com/go-xe2/x/xf/xfboot"

// 元注解启动初始化入口
type AnnotationEntry interface {
	xfboot.BootEntry
	Register(ann Annotation) bool
}

type xfAnnEntry struct {
	items map[string]Annotation
}

var _ xfboot.BootEntry = (*xfAnnEntry)(nil)
var _ AnnotationEntry = (*xfAnnEntry)(nil)

const AnnotationEntryName = "annotation"

var annotationEntry = xfboot.GetEntryOrRegister(AnnotationEntryName, func() xfboot.BootEntry {
	return &xfAnnEntry{
		items: make(map[string]Annotation),
	}
}).(*xfAnnEntry)

func (ae *xfAnnEntry) EntryName() string {
	return AnnotationEntryName
}

func (ae *xfAnnEntry) Entry() interface{} {
	return ae
}

func (ae *xfAnnEntry) Register(ann Annotation) bool {
	if _, ok := ae.items[ann.AnnotationName()]; ok {
		return false
	}
	ae.items[ann.AnnotationName()] = ann
	return true
}

func (ae *xfAnnEntry) Items() map[string]Annotation {
	return ae.items
}

func (ae *xfAnnEntry) HasAnnotation(annName string) bool {
	if _, ok := ae.items[annName]; ok {
		return true
	}
	return false
}

func (ae *xfAnnEntry) GetAnnotation(annName string) Annotation {
	if v, ok := ae.items[annName]; ok {
		return v
	}
	return nil
}

func (ae *xfAnnEntry) IsInit() bool {
	return xfboot.IsEntryInit(ae.EntryName())
}

func (ae *xfAnnEntry) Init() {
	if ae.IsInit() {
		return
	}
	xfboot.InitEntry(ae.EntryName())
}
