package xclass

type ObjectHello interface {
	Hello() string
}

func (o *TObject) SayHello() string {
	if v, ok := o.this.(ObjectHello); ok {
		return v.Hello()
	}
	return ""
}
