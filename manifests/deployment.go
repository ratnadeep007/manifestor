package manifests

import (
	"log"

	"gopkg.in/yaml.v2"
)

type Deployment struct {
	Name                string
	Kind                string
	ContainerCount      int
	Containers          []Container
	InitContainerNeeded bool
	InitContainers      []InitContainer
}

type Labels struct {
	App string
}

type Metadata struct {
	Name      string
	Namespace string
	Labels
}

type MatchLabels struct {
	App string
}

type Selector struct {
	MatchLabels
}

type TemplateMetadata struct {
	Labels
}

type Template struct {
	Metadata TemplateMetadata
}

type Container struct {
	Name  string
	Image string
	Ports []map[string]int
}

type InitContainer struct {
	Name    string
	Image   string
	Command []string
}

type TemplateSpec struct {
	Containers     []Container
	InitContainers []InitContainer `yaml:"initContainers,omitempty"`
}

type SpecTemplate struct {
	Metadata TemplateMetadata
	Spec     TemplateSpec
}

type Spec struct {
	Replicas int
	Selector
	Template SpecTemplate
}

type DeploymentOutput struct {
	ApiVersion string
	Kind       string
	Metadata
	Spec
}

// Function to prepare struct that can be converted to yaml
func (do DeploymentOutput) PrepareDeployment(deployement Deployment) DeploymentOutput {
	do.ApiVersion = "apps/v1"
	do.Kind = "Deployment"
	deploymentMetadata := Metadata{
		Name:      deployement.Name,
		Namespace: "default",
		Labels: Labels{
			App: deployement.Name,
		},
	}
	do.Metadata = deploymentMetadata
	deploymentSpec := Spec{
		Replicas: 1,
		Selector: Selector{
			MatchLabels: MatchLabels{
				App: deployement.Name,
			},
		},
		Template: SpecTemplate{
			Metadata: TemplateMetadata{
				Labels: Labels{
					App: deployement.Name,
				},
			},
			Spec: TemplateSpec{},
		},
	}
	deploymentSpec.Template.Spec.Containers = do.AddContainers(deployement)
	if deployement.InitContainerNeeded {
		deploymentSpec.Template.Spec.InitContainers = do.AddInitContainers(deployement)
	}
	do.Spec = deploymentSpec
	return do
}

// Function to add containers
func (do DeploymentOutput) AddContainers(deploy Deployment) []Container {
	containers := []Container{}
	for _, v := range deploy.Containers {
		container := Container{
			Name:  v.Name,
			Image: v.Image,
			Ports: []map[string]int{{"containerPort": 8000}},
		}
		containers = append(containers, container)
	}
	return containers
}

func (do DeploymentOutput) AddInitContainers(deploy Deployment) []InitContainer {
	containers := []InitContainer{}
	for _, v := range deploy.InitContainers {
		container := InitContainer{
			Name:    v.Name,
			Image:   v.Image,
			Command: []string{},
		}
		containers = append(containers, container)
	}
	return containers
}

// Convert struct to YAML
func (do DeploymentOutput) GetYAML(deploy Deployment) []byte {
	deployment := do.PrepareDeployment(deploy)
	d, err := yaml.Marshal(&deployment)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
