package structs

func SliceStruct(structPtr interface{}) ([]IField, error) {
	obj, err := New(structPtr)
	if err != nil {
		return nil, err
	}
	return obj.Fields(), nil
}
