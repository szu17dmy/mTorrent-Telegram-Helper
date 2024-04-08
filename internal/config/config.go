package config

type Configs struct {
	Jobs []*Job
}

type Job struct {
	Name string
	Spec string
	Func func()
}
