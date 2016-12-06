package settings

import "github.com/spf13/viper"

// Solver contains all information regarding a FlatZinc solver and its output
type Solver struct {
	Name    string // Solver name
	Binary  string // Binary location
	Globals string // Globals directory
	Flags   string // FZN solver flags

	Extractors *ExtractionCluster
}

// SolversFromViper extracts all solver information from Viper
func SolversFromViper() []Solver {
	var solvers []Solver

	for key := range viper.GetStringMap("solvers") {
		options := viper.GetStringMapString("solvers." + key)

		solver := Solver{
			Name:    key,
			Flags:   options["flags"],
			Globals: options["globals"],
		}
		if bin, exists := options["binary"]; exists {
			solver.Binary = bin
		} else {
			solver.Binary = solver.Name
		}

		solver.Extractors = ExtractorsFromViper("solvers." + key)

		solvers = append(solvers, solver)
	}

	return solvers
}
