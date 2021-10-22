package v1

type ClusterInfo struct {
	Name      string        `json:"name" bson:"name"`
	ID        string        `json:"id" bson:"_id"`
	Slug      string        `json:"slug" bson:"slug"`
	Endpoints Endpoints     `json:"apis" bson:"apis"`
	Targets   []MongoTarget `json:"targets" bson:"targets"`
}

type Endpoints struct {
	BaseURL string `json:"baseUrl" bson:"baseUrl"`
	Paths   Paths  `json:"paths" bson:"paths"`
}

type Paths struct {
	ClusterMetadata     string `json:"clusterMetadata" bson:"clusterMetadata"`
	FuryMetadata        string `json:"furyMetadata" bson:"furyMetadata"`
	ApplicationMetadata string `json:"applicationMetadata" bson:"applicationMetadata"`
}

type MongoTarget struct {
	TargetName string `json:"targetName" bson:"targetName"`
	Collection string `json:"collection" bson:"collection"`
}
