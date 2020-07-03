package exception

type ICause interface {
	Cause() error
}
