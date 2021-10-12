package v1

type KFDRelease struct {
	Version     string   `json:"version"`
	ReleaseDate string   `json:"release_date"`
	Modules     []Module `json:"modules"`
}

type Module struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ModuleRelease struct {
	Version     string `json:"version"`
	ReleaseDate string `json:"release_date"`
}

type KFDReleaseDef struct {
	Version              string                `json:"version"`
	ReleaseDate          string                `json:"release_date"`
	Description          string                `json:"description"`
	Repo                 string                `json:"repo_url"`
	Doc                  string                `json:"doc_url"`
	ReleaseNotes         map[string][]string   `json:"release_notes"`
	SupportedK8sRuntimes []float32             `json:"supported_k8s_runtimes"`
	Modules              []KFDModuleReleaseDef `json:"modules"`
}

type KFDModuleReleaseDef struct {
	Module                 string                `json:"module"`
	Description            string                `json:"description"`
	Version                string                `json:"version"`
	ReleaseDate            string                `json:"release_date"`
	Repo                   string                `json:"repo_url"`
	ReleaseNotes           map[string][]string   `json:"release_notes"`
	SupportedK8sRuntimes   []float32             `json:"supported_k8s_runtimes"`
	UpgradeWarnings        []string              `json:"upgrade_warnings"`
	KnownIssues            []string              `json:"known_issues"`
	K8sCompatibilityMatrix []CompatibilityMatrix `json:"k8s_compatibility_matrix"`
	Components             []Component           `json:"components"`
}

type CompatibilityMatrix struct {
	K8sVersion float32 `json:"k8s_version"`
	State      string  `json:"state"`
}

type Component struct {
	Name      string     `json:"name"`
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Kind       string      `json:"kind"`
	Name       string      `json:"name"`
	Namespace  string      `json:"namespace"`
	APIVersion string      `json:"apiVersion"`
	Containers []Container `json:"containers"`
}

type Container struct {
	Image           string `json:"image"`
	Version         string `json:"version"`
	UpstreamImage   string `json:"upstreamImage"`
	UpstreamVersion string `json:"upstreamVersion"`
}
