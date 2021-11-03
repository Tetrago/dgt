package config

type Group struct {
	Path string
	Aliases []string
}

type Target struct {
	Group *Group
	Name string
}

type Task struct {
	Name string
	Aliases []string
	Groups []*Group
	Targets *[]Target
}

func (task *Task) IsTargetBased() bool {
	// `task.Targets` will be null when the task is group based (Ex. `docker compose up`)
	return task.Targets != nil
}

type Config struct {
	Groups []Group
	Tasks []Task
}