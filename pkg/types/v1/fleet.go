// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

type ClusterInfo struct {
	Name      string        `json:"name" bson:"name"`
	ID        string        `json:"id" bson:"_id,omitempty"`
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

type ClusterDataUsefulLink struct {
	Name string `bson:"name" json:"name"`
	URL  string `bson:"url" json:"url"`
}

type ClusterDataContact struct {
	Fullname string `bson:"fullname" json:"fullname"`
	Email    string `bson:"email" json:"email"`
	Phone    string `bson:"phone" json:"phone"`
	OnCall   bool   `bson:"onCall" json:"onCall"`
}

type ClusterDataHardwareInfo struct {
	Quantity int    `bson:"quantity" json:"quantity"`
	Unit     string `bson:"unit" json:"unit"`
}

type ClusterData struct {
	ID                     string                  `bson:"_id,omitempty" json:"id"`
	Name                   string                  `bson:"name" json:"name"`
	Slug                   string                  `bson:"slug" json:"slug"`
	Provider               string                  `bson:"provider" json:"provider"`
	PkiCertExpirationDate  string                  `bson:"pkiCertExpirationDate" json:"pkiCertExpirationDate"`
	EtcdCertExpirationDate string                  `bson:"etcdCertExpirationDate" json:"etcdCertExpirationDate"`
	KubernetesVersion      string                  `bson:"kubernetesVersion" json:"kubernetesVersion"`
	OS                     string                  `bson:"os" json:"os"`
	ContainerRuntime       string                  `bson:"containerRuntime" json:"containerRuntime"`
	CPU                    ClusterDataHardwareInfo `bson:"cpu" json:"cpu"`
	RAM                    ClusterDataHardwareInfo `bson:"ram" json:"ram"`
	WorkerNodes            int                     `bson:"workerNodes" json:"workerNodes"`
	OnCall                 bool                    `bson:"onCall" json:"onCall"`
	UsefulLinks            []ClusterDataUsefulLink `bson:"usefulLinks" json:"usefulLinks"`
	Contacts               []ClusterDataContact    `bson:"contacts" json:"contacts"`
	Fury                   KFDReleaseDef           `bson:"fury" json:"fury"`
}

type ClusterGroupClusterStatus struct {
	Name          string `bson:"name" json:"name"`
	LastUpdatedAt string `bson:"lastUpdatedAt" json:"lastUpdatedAt"`
}

type ClusterGroupCluster struct {
	Name        string                    `bson:"name" json:"name"`
	Slug        string                    `bson:"slug" json:"slug"`
	Provider    string                    `bson:"provider" json:"provider"`
	Environment string                    `bson:"environment" json:"environment"`
	Status      ClusterGroupClusterStatus `bson:"status" json:"status"`
	Contacts    []ClusterDataContact      `bson:"contacts" json:"contacts"`
}

type ClusterGroup struct {
	Id       string                `bson:"_id,omitempty" json:"id"`
	Name     string                `bson:"name" json:"name"`
	Clusters []ClusterGroupCluster `bson:"clusters" json:"clusters"`
}
