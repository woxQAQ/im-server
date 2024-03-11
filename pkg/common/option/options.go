package option

type Option interface {
	Default()

	Load()
}
