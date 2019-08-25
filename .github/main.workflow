workflow "Build & Test" {
	on = "push"
	resolves = ["Build"]
}

action "Build" {
	uses = "./workflows/go"
}
