package presentation

type Presenter[T any] interface {
	Parse(T) map[string]any
}
