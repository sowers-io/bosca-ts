package common

type WorkflowTransition struct {
	MetadataId            string
	WorkflowConfiguration map[string]string
	StateConfiguration    map[string]string
}
