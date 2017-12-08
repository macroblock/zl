package events

type tActionMap struct {
	byName     map[string]IAction
	byEventKey map[string]IAction
	mode       string
}

var _mAM = "actionmap "

// ActionMap -
var ActionMap tActionMap

func initActionMap() {
	ActionMap.byName = map[string]IAction{}
	ActionMap.byEventKey = map[string]IAction{}
}

// Add -
func (o *tActionMap) Add(action IAction) {
	if ActionMap.byName == nil {
		initActionMap()
	}
	_, ok := o.byName[action.Name()]
	log.Warning(ok, _mAM, "action Add: action is already set (overwriten)")
	o.byName[action.Name()] = action
}

// Delete -
func (o *tActionMap) Delete(action IAction) {
	_, ok := o.byName[action.Name()]
	log.Warning(!ok, _mAM, "Delete: action isn't set (skiped)")
	delete(o.byName, action.Name())
}

// Apply -
func (o *tActionMap) Apply() {
	o.byEventKey = map[string]IAction{}
	for _, act := range o.byName {
		_, ok := o.byEventKey[act.EventKey()]
		log.Warning(ok, _mAM, "Apply: action is already set (overwriten)")
		o.byEventKey[act.EventKey()] = act
	}
}

// SetMode -
func (o *tActionMap) SetMode(mode string) {
	o.mode = mode
	if len(o.mode) > 0 {
		o.mode += "/"
	}
}

// HandleEvent -
func HandleEvent(ev IEvent) {
	log.Debug("HandleEvent")
	if act, ok := ActionMap.byName[ev.Type()]; ok {
		log.Debug(ok)
		act.Do(ev)
	}
	if act, ok := ActionMap.byEventKey[ev.EventKey()]; ok {
		log.Debug(ok)
		act.Do(ev)
	}
}
