module github.com/sighupio/fip-commons

go 1.16

require (
	k8s.io/api v0.18.14
	k8s.io/apimachinery v0.18.14
	k8s.io/client-go v0.18.4
)

replace k8s.io/client-go => k8s.io/client-go v0.18.14
