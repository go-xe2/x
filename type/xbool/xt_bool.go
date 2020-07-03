package xbool

type Bool bool

func New(b ...bool) Bool {
	if len(b) > 0 {
		return Bool(b[0])
	}
	return Bool(false)
}

func (b Bool) String() string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func (b Bool) Not() bool {
	return !bool(b)
}

func (b Bool) Or(other bool) bool {
	return bool(b) || other
}

func (b Bool) And(other bool) bool {
	return bool(b) && other
}

func (b Bool) Xor(other bool) bool {
	var n1, n2 int8
	if b {
		n1 = 1
	}
	if other {
		n2 = 1
	}
	n3 := n1 ^ n2
	return n3 == 1
}
