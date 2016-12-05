package settings

import (
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Instance contains the information to run one particular instance of a model
type Instance struct {
	Model string
	Data  string
}

// InstancesFromViper extracts all instance information from Viper
func InstancesFromViper() []Instance {
	var instances []Instance

	data := viper.GetStringSlice("data")
	for _, model := range viper.GetStringSlice("models") {
		if len(data) > 0 {
			for _, dat := range data {
				instances = append(instances, Instance{Model: model, Data: dat})
			}
		} else {
			instances = append(instances, Instance{Model: model})
		}
	}

	return instances
}

// OutPath returns the intended output location of the instance given the solver
func (i *Instance) OutPath(solver string) string {
	path := viper.GetString("raw_dir")
	path = filepath.Join(path, solver)
	cleanModel := strings.TrimSuffix(i.Model, filepath.Ext(i.Model))
	path = filepath.Join(path, cleanModel)
	if i.Data != "" {
		cleanData := strings.TrimSuffix(i.Data, filepath.Ext(i.Data))
		path = filepath.Join(path, cleanData)
	}
	return path + ".dat"
}
