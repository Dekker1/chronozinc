package settings

import (
	"fmt"
	"log"
	"regexp"

	"github.com/spf13/viper"
)

// Solver contains all information regarding a FlatZinc solver and its output
type Solver struct {
	Name    string // Solver name
	Binary  string // Binary location
	Globals string // Globals directory
	Flags   string // FZN solver flags

	Extractors     map[string]*regexp.Regexp
	LastExtractors map[string]*regexp.Regexp
	Matchers       map[string]map[string]*regexp.Regexp
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

		solver.Extractors = make(map[string]*regexp.Regexp)
		for i, str := range viper.GetStringMapString("solvers." + key + ".extractors") {
			reg, err := regexp.Compile(str)
			if err != nil {
				log.Panicf("Error compiling extractor `%s:%s`: %s",
					key, i, err)
			} else {
				solver.Extractors[i] = reg
			}
		}

		solver.LastExtractors = make(map[string]*regexp.Regexp)
		for i, str := range viper.GetStringMapString("solvers." + key + ".last_extractors") {
			reg, err := regexp.Compile(str)
			if err != nil {
				log.Panicf("Error compiling extractor `%s:%s`: %s",
					key, i, err)
			} else {
				solver.LastExtractors[i] = reg
			}
		}

		solver.Matchers = make(map[string]map[string]*regexp.Regexp)
		for i := range viper.GetStringMap("solvers." + key + ".matchers") {
			matcher := make(map[string]*regexp.Regexp)
			for j, str := range viper.GetStringMapString("solvers." + key + ".matchers." + i) {
				reg, err := regexp.Compile(str)
				if err != nil {
					log.Panicf("Error compiling match case `%s:%s:`: %s",
						key, i, err)
				} else {
					matcher[j] = reg
				}
			}
			fmt.Println(matcher)
			solver.Matchers[i] = matcher
		}

		solvers = append(solvers, solver)
	}

	return solvers
}
