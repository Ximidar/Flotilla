package Helm

type ExecRule struct {
	Command   string
	Args      []string
	WD        string
	Container bool
	Name      string
	Info      interface{} // TODO change to whatever the Command Type is
}

// NewExecRule will set up a new rule
func NewExecRule(Command string, Args []string, WD string, Container bool, Name string) *ExecRule {
	rule := new(ExecRule)
	rule.Command = Command
	rule.Args = Args
	rule.WD = WD
	rule.Container = Container
	rule.Name = Name
	return rule
}

// Start will begin execution of the rule
func (rule *ExecRule) Start() error {
	return nil
}

// Stop will stop the rule
func (rule *ExecRule) Stop(force bool) error {
	return nil
}
