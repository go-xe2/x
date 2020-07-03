package anno

type tAnnotationContainer struct {
	annotation Annotation
	annParams  map[string]interface{}
}

var _ AnnotationContainer = (*tAnnotationContainer)(nil)

func newAnnotationContainer(ann Annotation, annParams map[string]interface{}) AnnotationContainer {
	return &tAnnotationContainer{
		annotation: ann,
		annParams:  annParams,
	}
}

func (ac *tAnnotationContainer) AnnName() string {
	if ac.annotation == nil {
		return ""
	}
	return ac.annotation.AnnotationName()
}

func (ac *tAnnotationContainer) Ann() Annotation {
	return ac.annotation
}

func (ac *tAnnotationContainer) AnnParams() map[string]interface{} {
	return ac.annParams
}

func (ac *tAnnotationContainer) Create(caller interface{}, params ...interface{}) interface{} {
	if ac.annotation == nil {
		return nil
	}
	return ac.annotation.AnnCreate(caller, ac.annParams, params...)
}
