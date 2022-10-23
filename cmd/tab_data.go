package cmd

type TabData struct {
	CommandInfo CommandInfo
}
type CommandInfo struct {
	Title       string
	Description string
	List        []string
}
