package agt

import (
	"fmt"
	"reflect"
	"td3/comsoc"
)

type Agent struct {
	ID    AgentID
	Name  string
	Prefs []comsoc.Alternative
}

type AgentID string

type AgentI interface {
	Equal(ag AgentI) bool
	DeepEqual(ag AgentI) bool
	Clone() AgentI
	String() string
	Prefers(a comsoc.Alternative, b comsoc.Alternative) bool
	Start()
}

func (a *Agent) Equal(ag AgentI) bool {
	// assertion de type, on obtient le résultat de la conversion et un booléen indiquant le succès de l'opération
	agentConverti, success := ag.(*Agent)
	if !success {
		return false
	}
	return a.ID == agentConverti.ID
}

func (a *Agent) DeepEqual(ag AgentI) bool {
	agentConverti, success := ag.(*Agent)
	if !success {
		return false
	}
	//reflect.DeepEqual fait une comparaison profonde
	return (a.ID == agentConverti.ID) &&
		(a.Name == agentConverti.Name) &&
		(reflect.DeepEqual(a.Prefs, agentConverti.Prefs))
}

func (a *Agent) Clone() AgentI {
	prefsCopy := make([]comsoc.Alternative, len(a.Prefs))
	copy(prefsCopy, a.Prefs)
	return &Agent{ID: a.ID, Name: a.Name, Prefs: prefsCopy}
}

func (a *Agent) String() string {
	return fmt.Sprintf("Agent{ID: %d, Name: %s, Prefs: %v}", a.ID, a.Name, a.Prefs)
}

func (a *Agent) Prefers(a1, a2 comsoc.Alternative) bool {
	return comsoc.IsPref(a1, a1, a.Prefs)
}
func (a *Agent) Start() {
	//start
}
