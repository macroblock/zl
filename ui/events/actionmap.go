package events

type tActionMap struct {
	byName     map[string]IAction
	byEventKey map[string]IAction
	mode       string
}

// ActionMap -
var ActionMap tActionMap

func initActionMap() {
	ActionMap = tActionMap{}
	ActionMap.byName = map[string]IAction{}
	ActionMap.Apply()
}

// Add -
func (am *tActionMap) Add(action IAction) {
	_, ok := am.byName[action.Name()]
	log.Warning(ok, "Add(): action is already set (overwriten)")
	am.byName[action.Name()] = action
}

// Delete -
func (am *tActionMap) Delete(action IAction) {
	_, ok := am.byName[action.Name()]
	log.Warning(!ok, "Delete(): action isn't set (skiped)")
	delete(am.byName, action.Name())
}

// Apply -
func (am *tActionMap) Apply() {
	am.byEventKey = map[string]IAction{}
	for _, act := range am.byName {
		_, ok := am.byEventKey[act.EventKey()]
		log.Warning(!ok, "Apply(): action is already set (overwriten)")
		am.byEventKey[act.EventKey()] = act
	}
}

// HandleEvent -
func HandleEvent(ev IEvent) {
	if act, ok := ActionMap.byEventKey[ActionMap.mode+ev.EventKey()]; ok {
		act.Do(ev)
	}
}
