package bot

import (
	"fmt"
	"github.com/bivas/rivi/util"
	"github.com/spf13/viper"
)

type Action interface {
	Apply(config Configuration, meta EventData)
}

type Rule interface {
	Name() string
	Accept(meta EventData) bool
	Actions() []Action
}

type rule struct {
	name      string
	condition Condition
	actions   []Action
}

func (r *rule) Name() string {
	return r.name
}

func (r *rule) String() string {
	return fmt.Sprintf("%#v", r)
}

func (r *rule) Accept(meta EventData) bool {
	accept := r.condition.Match(meta)
	if !accept {
		util.Logger.Debug("Skipping rule '%s'", r.name)
	}
	return accept
}

func (r *rule) Actions() []Action {
	return r.actions
}

type ActionFactory interface {
	BuildAction(config map[string]interface{}) Action
}

var actions map[string]ActionFactory = make(map[string]ActionFactory)
var supportedActions []string = make([]string, 0)

func RegisterAction(kind string, action ActionFactory) {
	actions[kind] = action
	supportedActions = append(supportedActions, kind)
	util.Logger.Debug("running with support for %s", kind)
}

func buildActionsFromConfiguration(config *viper.Viper) []Action {
	result := make([]Action, 0)
	for setting := range config.AllSettings() {
		if setting == "condition" {
			continue
		}
		for _, support := range supportedActions {
			if setting == support {
				factory := actions[setting]
				result = append(result, factory.BuildAction(config.GetStringMap(setting)))
			}
		}
	}
	return result
}
