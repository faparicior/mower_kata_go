package movemowers

type MoveMowersCommand struct {
	fileName string
}

func BuildMoveMowersCommand(filename string) MoveMowersCommand {
	return MoveMowersCommand{fileName: filename}
}

func (command MoveMowersCommand) filename() string {
	return command.fileName
}
