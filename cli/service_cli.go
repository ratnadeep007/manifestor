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
	ports := PortsListPrompt(portsCount)
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
	if err != nil {
		fmt.Printf("Failed %v\n", err)
	}

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
	if err != nil {
		fmt.Printf("Failed %v\n", err)
	}

	return name
}

func ServiceTypePrompt() string {
	promptServiceType := promptui.Select{
		Label:     "Select Type of service",
		Items:     []string{"NodePort", "ClusterIP", "LoadBalancer"},
		Templates: utils.GetSelectTemplate(),
	}
	_, resultServiceType, err := promptServiceType.Run()
	if err != nil {
		fmt.Printf("Failed %v\n", err)
	}
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
	if err != nil {
		fmt.Printf("Failed %v\n", err)
	}
	count, err := strconv.Atoi(portCountResult)
	if err != nil {
		fmt.Printf("Failed %v\n", err)
	}
	return count
}

func PortsListPrompt(portsCount int) []manifests.Port {
	ports := []manifests.Port{}
	for i := 0; i < portsCount; i++ {
		port := PortDetailPrompt(i)
		ports = append(ports, port)
	}
	return ports
}

func PortDetailPrompt(portsCount int) manifests.Port {
	validateName := func(input string) error {
		if len(input) < 3 {
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
	if err != nil {
		fmt.Println("Failed %v\n", err)
	}

	promptPortProtocol := promptui.Select{
		Label:     "Select Type of service",
		Items:     []string{"TCP", "UDP"},
		Templates: utils.GetSelectTemplate(),
	}
	_, resultPortProtocol, err := promptPortProtocol.Run()
	if err != nil {
		fmt.Printf("Failed %v\n", err)
	}

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
	if err != nil {
		fmt.Printf("Failed %v\n", err)
	}
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
	if err != nil {
		fmt.Printf("Failed %v\n", err)
	}
	targetPort, err := strconv.Atoi(targetPortString)
	if err != nil {
		log.Fatal(err)
	}
	return manifests.Port{
		Name:       name,
		TargetPort: targetPort,
		Protocol:   resultPortProtocol,
		Port:       port,
	}
}
