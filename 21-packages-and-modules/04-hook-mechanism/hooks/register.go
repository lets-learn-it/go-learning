package hooks

type Hook interface {
	Init()
	Perform()
	Destroy()
}

var hooks []Hook

func Register(h Hook) {
	hooks = append(hooks, h)
}

func GetHooks() []Hook {
	return hooks
}
