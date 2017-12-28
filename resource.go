package bamboo

// ResourceMetadata holds metadata about the API service response
// - Expand: Element of the expand parameter used for the service
// - Link: See ServiceLink
type ResourceMetadata struct {
	Expand string `json:"expand"`
	Link   *Link  `json:"link"`
}

// Link holds link information for the service
// - HREF: Relationship between link and element (defaults to "self")
// - Rel:  URL for the project
type Link struct {
	HREF string `json:"href"`
	Rel  string `json:"rel"`
}

// CollectionMetadata holds metadata about a collection of Bamboo resources
// - Size:       Number of resources
// - Expand:     Element of the expand parameter used for the collection
// - StartIndex: Index from which to the request started gathering resources
// - MaxResult:  The maximum number of returned resources for the request
type CollectionMetadata struct {
	Size       int    `json:"size"`
	Expand     string `json:"expand"`
	StartIndex int    `json:"start-index"`
	MaxResult  int    `json:"max-result"`
}
