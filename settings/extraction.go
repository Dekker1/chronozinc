package settings

import (
	"log"
	"regexp"

	"github.com/spf13/viper"
)

type ExtractionCluster struct {
	Extractors     map[string]*regexp.Regexp
	LastExtractors map[string]*regexp.Regexp
	Matchers       map[string]map[string]*regexp.Regexp
}

func ExtractorsFromViper(viperSpace string) *ExtractionCluster {
	var cluster ExtractionCluster
	if viperSpace != "" {
		viperSpace += "."
	}

	cluster.Extractors = make(map[string]*regexp.Regexp)
	for i, str := range viper.GetStringMapString(viperSpace + "extractors") {
		reg, err := regexp.Compile(str)
		if err != nil {
			log.Panicf("Error compiling extractor `%s:%s`: %s",
				viperSpace, i, err)
		} else {
			cluster.Extractors[i] = reg
		}
	}

	cluster.LastExtractors = make(map[string]*regexp.Regexp)
	for i, str := range viper.GetStringMapString(viperSpace + "last_extractors") {
		reg, err := regexp.Compile(str)
		if err != nil {
			log.Panicf("Error compiling extractor `%s:%s`: %s",
				viperSpace, i, err)
		} else {
			cluster.LastExtractors[i] = reg
		}
	}

	cluster.Matchers = make(map[string]map[string]*regexp.Regexp)
	for i := range viper.GetStringMap(viperSpace + "matchers") {
		matcher := make(map[string]*regexp.Regexp)
		for j, str := range viper.GetStringMapString(viperSpace + "matchers." + i) {
			reg, err := regexp.Compile(str)
			if err != nil {
				log.Panicf("Error compiling match case `%s:%s:`: %s",
					viperSpace, i, err)
			} else {
				matcher[j] = reg
			}
		}
		cluster.Matchers[i] = matcher
	}

	return &cluster
}
