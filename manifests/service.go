package manifests

import (
	"log"

	"gopkg.in/yaml.v2"
)

type Service struct {
	Name       string
	Kind       string
	Selector   string
	PortsCount int
	Ports      []Port
	Namespace  string
	Label      string
	Type       string
}

type ServiceSelector struct {
	App string
}

type ServiceSpec struct {
	Selector ServiceSelector
	Type     string
	Ports    []Port
}

type Port struct {
	Name       string `yaml:",omitempty"`
	Protocol   string
	Port       int
	TargetPort int `yaml:"targetPort"`
	NodePort   int `yaml:",omitempty"`
}

type ServiceOutput struct {
	ApiVersion string
	Kind       string
	Metadata   Metadata
	Spec       ServiceSpec
}

func (so ServiceOutput) PrepareService(service Service) ServiceOutput {
	so.ApiVersion = "v1"
	so.Kind = "Service"
	serviceMetadata := Metadata{
		Name:      service.Name,
		Namespace: "default",
		Labels: Labels{
			App: service.Name,
		},
	}
	so.Metadata = serviceMetadata
	so.Spec = ServiceSpec{
		Selector: ServiceSelector{
			App: service.Name,
		},
		Type:  service.Type,
		Ports: so.AddPorts(service),
	}
	return so
}

func (so ServiceOutput) AddPorts(service Service) []Port {
	var ports []Port
	for _, v := range service.Ports {
		port := Port{
			Protocol:   "TCP",
			Name:       v.Name,
			Port:       v.Port,
			TargetPort: v.TargetPort,
		}
		ports = append(ports, port)
	}
	return ports
}

func (so ServiceOutput) ServiceGetYAML(svc Service) []byte {
	service := so.PrepareService(svc)
	d, err := yaml.Marshal(&service)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
