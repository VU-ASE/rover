package components

// A list of actions that can be performed with this utility
type ActionItem struct {
	Name, Desc string
}

func (i ActionItem) Title() string       { return i.Name }
func (i ActionItem) Description() string { return i.Desc }
func (i ActionItem) FilterValue() string { return i.Name }
