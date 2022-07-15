// Copyright (c) 2022 SIGHUP s.r.l All rights reserved.
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

type ClusterDataUsefulLinks struct {
	Label string `bson:"label" json:"label"`
	Links []Link `bson:"url" json:"url"`
}

type Link struct {
	Name string   `bson:"name" json:"name"`
	URLs []string `bson:"urls" json:"urls"`
}

type ClusterDataContacts struct {
	Fullname string `bson:"fullname" json:"fullname"`
	Email    string `bson:"email" json:"email"`
	Phone    string `bson:"phone" json:"phone"`
	OnCall   bool   `bson:"onCall" json:"onCall"`
}

type ClusterDataFuryModules struct {
	Name    string `bson:"name" json:"name"`
	Version string `bson:"version" json:"version"`
}

type ClusterDataFury struct {
	Version string                   `bson:"version" json:"version"`
	Modules []ClusterDataFuryModules `bson:"modules" json:"modules"`
}

type ClusterDataHardwareInfo struct {
	Quantity int    `bson:"quantity" json:"quantity"`
	Unit     string `bson:"unit" json:"unit"`
}

type ClusterData struct {
	Id                string                  `bson:"_id,omitempty" json:"id"`
	Uuid              string                  `bson:"uuid" json:"uuid"`
	Name              string                  `bson:"name" json:"name"`
	Slug              string                  `bson:"slug" json:"slug"`
	Provider          string                  `bson:"provider" json:"provider"`
	KubernetesVersion string                  `bson:"kubernetesVersion" json:"kubernetesVersion"`
	Os                string                  `bson:"os" json:"os"`
	ContainerRuntime  string                  `bson:"containerRuntime" json:"containerRuntime"`
	Cpu               ClusterDataHardwareInfo `bson:"cpu" json:"cpu"`
	Ram               ClusterDataHardwareInfo `bson:"ram" json:"ram"`
	PkiCert           CertificateSpec         `bson:"pkiCert" json:"pkiCert"`
	EtcdCert          CertificateSpec         `bson:"etcdCert" json:"etcdCert"`
	ApplicationCert   CertificateSpec         `bson:"applicationCert,omitempty" json:"applicationCert"`
	WorkerNodes       int                     `bson:"workerNodes" json:"workerNodes"`
	Fury              ClusterDataFury         `bson:"fury" json:"fury"`
	OnCall            bool                    `bson:"onCall" json:"onCall"`
	UsefulLinks       ClusterDataUsefulLinks  `bson:"usefulLinks" json:"usefulLinks"`
	Environment       string                  `bson:"environment" json:"environment"`

	Status       ClusterStatus `bson:"status" json:"status"`
	HealthChecks []HealthCheck `bson:"healthChecks" json:"healthChecks"`
	CreatedAt    string        `bson:"createdAt" json:"createdAt"`

	GitRepository       string                `bson:"gitRepository" json:"gitRepository"`
	KubeConfigPath      string                `bson:"kubeconfigPath" json:"kubeconfigPath"`
	GitCryptName        string                `bson:"gitCryptName" json:"gitCryptName"`
	ProviderCredentials []ProviderCredentials `bson:"providerCredentials" json:"providerCredentials"`

	Tags  []string `bson:"tags" json:"tags"`
	Notes string   `bson:"notes" json:"notes"`
}

type ClusterStatus struct {
	Name          string `bson:"name" json:"name"`
	LastUpdatedAt string `bson:"lastUpdatedAt" json:"lastUpdatedAt"`
}

type ProviderCredentials struct {
	ProviderName    string `bson:"providerName" json:"providerName"`
	CredentialsName string `bson:"credentialsName" json:"credentialsName"`
}

type HealthCheck struct {
	Category   string       `json:"category"`
	Severities SeveritySpec `json:"severities"`
	Alerts     []Alert      `json:"alerts"`
}

type Alert struct {
	Severity string `json:"severity"`
	Name     string `json:"name"`
	Message  string `json:"message"`
	URL      string `json:"url"`
}

type SeveritySpec struct {
	Critical SeverityCount `json:"CRITICAL,omitempty"`
	Warning  SeverityCount `json:"WARNING,omitempty"`
}

type SeverityCount struct {
	Count int `json:"count"`
}

type CertificateSpec struct {
	Name     string `bson:"name" json:"name"`
	NotAfter string `bson:"notAfter" json:"notAfter"`
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
}

type ClusterGroup struct {
	Id       string                `bson:"_id,omitempty" json:"id"`
	Name     string                `bson:"name" json:"name"`
	Clusters []ClusterGroupCluster `bson:"clusters" json:"clusters"`
}
