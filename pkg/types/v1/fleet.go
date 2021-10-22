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
	Paths   []Path `json:"paths" bson:"paths"`
}

type Path struct {
	Name     string `json:"name" bson:"name"`
	Endpoint string `json:"endpoint" bson:"endpoint"`
}

type MongoTarget struct {
	TargetName string `json:"name" bson:"name"`
	Collection string `json:"collection" bson:"collection"`
}
