package config

import "testing"

func TestParse(t *testing.T) {
	file, _ := parse(`
	[[group]]
	path = "systems"
	aliases = [ "s" ]

	[[task]]
	name = "deploy"
	aliases = [ "d" ]
	groups = [ "systems" ]
	commands = [
		"up -d $"
	]
	`)

	if len(file.Groups) != 1 {
		t.Errorf("Groups parsed incorrectly, found %d", len(file.Groups))
	}
	
	if group := file.Groups[0]; group.Path != "systems" || len(group.Aliases) != 1 || group.Aliases[0] != "s" {
		t.Errorf("Failed to parse group properly, found %s", group)
	}

	if len(file.Tasks) != 1 {
		t.Errorf("Tasks parsed incorrectly, found %d", len(file.Tasks))
	}

	if task := file.Tasks[0]; task.Name != "deploy" || len(task.Aliases) != 1 || task.Aliases[0] != "d" || len(task.Groups) != 1 || task.Groups[0] != "systems" || len(task.Commands) != 1 || task.Commands[0] != "up -d $" {
		t.Errorf("Failed to parse task properly, found %s", task)
	}
}