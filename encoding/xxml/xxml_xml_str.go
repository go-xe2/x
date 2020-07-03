package xxml

type XmlStr string

func (xs XmlStr) String() string {
	return string(xs)
}
