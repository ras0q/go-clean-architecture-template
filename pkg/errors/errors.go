package errors

type Code interface {
	~uint
	Messsage() string
}

type CodeIsImplementedBy[T Code] struct{}

type CodeError[T Code] struct {
	Code     T
	internal error
}

func (e *CodeError[T]) Error() string {
	if e.internal == nil {
		return e.Code.Messsage()
	}

	return e.Code.Messsage() + ": " + e.internal.Error()
}

func New[T Code](code T) error {
	return &CodeError[T]{
		Code: code,
	}
}

func Wrap[T Code](code T, internal error) error {
	return &CodeError[T]{
		Code:     code,
		internal: internal,
	}
}
