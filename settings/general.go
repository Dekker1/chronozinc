package settings

var extractors *ExtractionCluster

// GlobalExtractors returns the extractors defined on a global level
func GlobalExtractors() *ExtractionCluster {
	if extractors == nil {
		extractors = ExtractorsFromViper("")
	}
	return extractors
}
