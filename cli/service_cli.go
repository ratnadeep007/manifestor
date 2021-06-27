package cli

import (
	"errors"
	"fmt"
	"log"
	"manifest_creator/manifests"
	"manifest_creator/utils"
	"strconv"

	"github.com/manifoldco/promptui"
)

func ServiceQuestions() {
	serviceName := NamePromptService()
	selector := SelectorPromptService()
	serviceType := ServiceTypePrompt()
	portsCount := PortsCountPrompt()
	ports := PortsListPrompt(portsCount, serviceType)
	service := manifests.Service{
		Name:       serviceName,
		Selector:   selector,
		Type:       serviceType,
		PortsCount: portsCount,
		Ports:      ports,
	}
	serviceOutput := manifests.ServiceOutput{}
	dataByte := serviceOutput.ServiceGetYAML(service)
	utils.CreateFile(serviceName+"_service", dataByte)
}

func NamePromptService() string {
	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("length must be greater than 3")
		}
		return nil
	}

	promptName := promptui.Prompt{
		Label:     "Service Name: ",
		Validate:  validate,
		Templates: utils.GetPromptTemplate(),
	}

	name, err := promptName.Run()
	utils.GraceFullExit(err)

	return name
}

func SelectorPromptService() string {
	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("length must be greater than 3")
		}
		return nil
	}

	promptSelector := promptui.Prompt{
		Label:     "Selector Name: ",
		Validate:  validate,
		Templates: utils.GetPromptTemplate(),
	}

	name, err := promptSelector.Run()
	utils.GraceFullExit(err)

	return name
}

func ServiceTypePrompt() string {
	promptServiceType := promptui.Select{
		Label:     "Select Type of service",
		Items:     []string{"NodePort", "ClusterIP", "LoadBalancer"},
		Templates: utils.GetSelectTemplate(),
	}
	_, resultServiceType, err := promptServiceType.Run()
	utils.GraceFullExit(err)

	return resultServiceType
}

func PortsCountPrompt() int {
	validate := func(input string) error {
		number, err := strconv.Atoi(input)

		if err != nil {
			return errors.New("this must be number")
		}
		if number < 1 {
			return errors.New("number must be greter than 0")
		}
		return nil
	}

	promptPortCount := promptui.Prompt{
		Validate:  validate,
		Label:     "Number of ports: ",
		Templates: utils.GetPromptTemplate(),
	}
	portCountResult, err := promptPortCount.Run()
	utils.GraceFullExit(err)

	count, err := strconv.Atoi(portCountResult)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func PortsListPrompt(portsCount int, serviceType string) []manifests.Port {
	ports := []manifests.Port{}
	for i := 0; i < portsCount; i++ {
		port := PortDetailPrompt(i, serviceType)
		ports = append(ports, port)
	}
	return ports
}

func PortDetailPrompt(portsCount int, serviceType string) manifests.Port {
	fmt.Println(portsCount)
	validateName := func(input string) error {
		if portsCount == 0 {
			return nil
		} else if len(input) < 3 {
			return errors.New("length must be greater than 3")
		}
		return nil
	}

	namePrompt := promptui.Prompt{
		Label:     "Name of port " + strconv.Itoa(portsCount) + ":",
		Validate:  validateName,
		Templates: utils.GetPromptTemplate(),
	}
	name, err := namePrompt.Run()
	utils.GraceFullExit(err)

	promptPortProtocol := promptui.Select{
		Label:     "Select Type of service",
		Items:     []string{"TCP", "UDP"},
		Templates: utils.GetSelectTemplate(),
	}
	_, resultPortProtocol, err := promptPortProtocol.Run()
	utils.GraceFullExit(err)

	validatePort := func(input string) error {
		_, err := strconv.Atoi(input)

		if err != nil {
			return errors.New("this must be number")
		}
		return nil
	}

	promptPort := promptui.Prompt{
		Validate:  validatePort,
		Label:     "Port for container: ",
		Templates: utils.GetPromptTemplate(),
	}
	portString, err := promptPort.Run()
	utils.GraceFullExit(err)
	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatal(err)
	}

	promptTargetPort := promptui.Prompt{
		Validate:  validatePort,
		Label:     "Target Port: ",
		Templates: utils.GetPromptTemplate(),
	}
	targetPortString, err := promptTargetPort.Run()
	utils.GraceFullExit(err)
	targetPort, err := strconv.Atoi(targetPortString)
	if err != nil {
		log.Fatal(err)
	}
	var nodePort int
	if serviceType == "NodePort" {
		validateNodePort := func(input string) error {
			port, err := strconv.Atoi(input)

			if err != nil {
				return errors.New("this must be number")
			}
			if port > 32767 || port < 30000 {
				return errors.New("node port must be in range 30000 - 32767")
			}
			return nil
		}
		nodePortPrompt := promptui.Prompt{
			Validate:  validateNodePort,
			Label:     "Target Port: ",
			Templates: utils.GetPromptTemplate(),
		}
		nodePortString, err := nodePortPrompt.Run()
		utils.GraceFullExit(err)
		nodePort, err = strconv.Atoi(nodePortString)
		if err != nil {
			log.Fatal(err)
		}
	}
	return manifests.Port{
		Name:       name,
		TargetPort: targetPort,
		Protocol:   resultPortProtocol,
		Port:       port,
		NodePort:   nodePort,
	}
}
