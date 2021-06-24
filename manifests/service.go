package manifests

type ServiceType int

const (
	ClusterIP ServiceType = iota
	NodePort
	LoadBalancer
	ExternalName
)

type Service struct {
	Name       string
	Kind       string
	Selector   string
	PortsCount int
	Type       ServiceType
}
