package ui

//	"gitlab.com/gomidi/midi"

/*
func findInPort(name string) midi.In {
		for _, port := range ins {
			if port.String() == name {
				return port
			}
		}
	return nil
}
*/

/*
func findOutPort(name string) midi.Out {
		for _, port := range outs {
			if port.String() == name {
				return port
			}
		}
	return nil
}
*/

/*
func getOutPorts() (ports []string) {
		for _, port := range outs {
			ports = append(ports, port.String())
		}
	return
}
*/

/*
func getInPorts() (ports []string) {
		for _, port := range ins {
			ports = append(ports, port.String())
		}
	return
}
*/

/*
func getLines() (res []string) {
		for _, st := range data.Lines {
			res = append(res, st.Name)
		}
	return res
}
*/

/*
func getActions() (res []string) {
		for _, st := range data.Actions {
			res = append(res, st.Name)
		}
	return res
}
*/

/*
func getConditions() (res []string) {
		for _, st := range data.Conditions {
			res = append(res, st.Name)
		}
	return res
}
*/

/*
func getConditionMakers() (res [][2]string) {
		names := lineconfig.RegisteredConditionMaker()
		infos := lineconfig.RegisteredConditionMakerInfos()
		sort.Strings(names)

		for _, name := range names {
			res = append(res, [2]string{name, infos[name]})
		}
	return
}
*/

/*
func getActionMakers() (res [][2]string) {
		names := lineconfig.RegisteredActionMaker()
		infos := lineconfig.RegisteredActionMakerInfos()
		sort.Strings(names)

		for _, name := range names {
			res = append(res, [2]string{name, infos[name]})
		}
	return
}
*/

/*
func addLine(m lineconfig.Line) {
	data.Lines = append(data.Lines, m)
}

func addAction(m lineconfig.Action) {
	data.Actions = append(data.Actions, m)
}

func addCondition(m lineconfig.Condition) {
	data.Conditions = append(data.Conditions, m)
}

func replaceLine(oldname string, n lineconfig.Line) {
	var nm = []lineconfig.Line{n}

	for _, m := range data.Lines {
		if m.Name != oldname {
			nm = append(nm, m)
		}
	}

	data.Lines = nm
}

func replaceAction(oldname string, n lineconfig.Action) {
	var nm = []lineconfig.Action{n}

	for _, m := range data.Actions {
		if m.Name != oldname {
			nm = append(nm, m)
		}
	}

	data.Actions = nm
}

func removeCondition(name string) {
	var nm = []lineconfig.Condition{}

	for _, m := range data.Conditions {
		if m.Name != name {
			nm = append(nm, m)
		}
	}

	data.Conditions = nm
}

func removeLine(name string) {
	var nm = []lineconfig.Line{}

	for _, m := range data.Lines {
		if m.Name != name {
			nm = append(nm, m)
		}
	}

	data.Lines = nm
}

func removeAction(name string) {
	var nm = []lineconfig.Action{}

	for _, m := range data.Actions {
		if m.Name != name {
			nm = append(nm, m)
		}
	}

	data.Actions = nm
}

func replaceCondition(oldname string, n lineconfig.Condition) {
	var nm = []lineconfig.Condition{n}

	for _, m := range data.Conditions {
		if m.Name != oldname {
			nm = append(nm, m)
		}
	}

	data.Conditions = nm
}

//var runScreen *runnerScreen

func saveConfig() error {
		bt, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}

		return ioutil.WriteFile(CONFIG_FILE, bt, 0644)
	return nil
}

*/
