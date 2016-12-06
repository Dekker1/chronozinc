package parsing

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"

	"github.com/jjdekker/chronozinc/settings"
	"github.com/spf13/viper"
)

// ParseAll parses all viper given parameters for all instances and saves them
// to a CSV file
func ParseAll(solvers []settings.Solver, instances []settings.Instance) {
	params := viper.GetStringSlice("parameters")
	if len(params) > 0 {
		f, err := os.Create(viper.GetString("output"))
		if err != nil {
			log.Panicf("Unable to create file %s", viper.GetString("output"))
		}
		defer f.Close()
		w := csv.NewWriter(f)
		defer w.Flush()

		headers := append(persistantHeaders(), params...)
		w.Write(headers)
		for i := range solvers {
			for j := range instances {
				record := make([]string, 0, len(headers))
				record = append(record,
					[]string{solvers[i].Name, instances[j].Model}...)
				if instances[j].Data != "" {
					record = append(record, instances[j].Data)
				}

				for k := range params {
					record = append(record,
						ParseParameter(&solvers[i], &instances[j], params[k]))
				}

				w.Write(record)
			}
		}
	}
}

func persistantHeaders() []string {
	headers := []string{"solver", "model"}
	if len(viper.GetStringSlice("data")) > 0 {
		headers = append(headers, "data")
	}
	return headers
}

// ParseParameter returns the parsed result of an Extraction for a given
// instance if found.
func ParseParameter(solver *settings.Solver, instance *settings.Instance,
	parameter string) string {
	if f, err := ioutil.ReadFile(instance.OutPath(solver.Name)); err != nil {
		log.Printf("Unable to open file %s", instance.OutPath(solver.Name))
	} else {
		clusters := []*settings.ExtractionCluster{solver.Extractors, settings.GlobalExtractors()}
		for _, c := range clusters {
			switch {
			case (c.Extractors[parameter] != nil):
				return Extract(f, c.Extractors[parameter])
			case (c.LastExtractors[parameter] != nil):
				return ExtractLast(f, c.LastExtractors[parameter])
			case (c.Matchers[parameter] != nil):
				return Match(f, c.Matchers[parameter])
			}
		}
	}
	return ""
}
