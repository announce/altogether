package domain

type TypedConfig interface {
	Path(Type) string
}

type ConfigPath map[Type]string

func (c ConfigPath) Path(p Type) string {
	return c[p]
}
