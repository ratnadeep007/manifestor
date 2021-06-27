package cli

import (
	"errors"
	"log"
	"manifest_creator/manifests"
	"manifest_creator/utils"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

func DeploymentQuestions() {
	deploymentName := NamePrompt()
	containerCount := ContainerCountPrompt()
	containers := ContainerDetailListPrompt(containerCount)
	initContainerNeeded, initContainer := InitContainerPrompt()
	deployment := manifests.Deployment{
		Name:                deploymentName,
		Kind:                "Deployment",
		ContainerCount:      containerCount,
		Containers:          containers,
		InitContainers:      initContainer,
		InitContainerNeeded: initContainerNeeded,
	}
	deploymentOutput := manifests.DeploymentOutput{}
	dataByte := deploymentOutput.GetYAML(deployment)
	utils.CreateFile(deploymentName+"_deployment", dataByte)
}

func NamePrompt() string {
	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("length must be greater than 3")
		}
		return nil
	}

	promptName := promptui.Prompt{
		Label:     "Deployment Name: ",
		Validate:  validate,
		Templates: utils.GetPromptTemplate(),
	}

	name, err := promptName.Run()
	utils.GraceFullExit(err)

	return name
}

func ContainerCountPrompt() int {
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

	promptContainerCount := promptui.Prompt{
		Validate:  validate,
		Label:     "Number of containers: ",
		Templates: utils.GetPromptTemplate(),
	}
	containerCountResult, err := promptContainerCount.Run()
	utils.GraceFullExit(err)
	count, err := strconv.Atoi(containerCountResult)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func ContainerDetailListPrompt(count int) []manifests.Container {
	containers := []manifests.Container{}
	for i := 0; i < count; i++ {
		container := ContainerDetailPrompt(i)
		containers = append(containers, container)
	}
	return containers
}

func ContainerDetailPrompt(containerNumber int) manifests.Container {
	validateName := func(input string) error {
		if len(input) < 3 {
			return errors.New("length must be greater than 3")
		}
		return nil
	}

	namePrompt := promptui.Prompt{
		Label:     "Name of container " + strconv.Itoa(containerNumber+1) + ": ",
		Validate:  validateName,
		Templates: utils.GetPromptTemplate(),
	}
	name, err := namePrompt.Run()
	utils.GraceFullExit(err)

	validateImage := func(input string) error {
		if len(input) == 0 {
			return nil
		}
		if len(input) > 0 && len(input) < 5 {
			return errors.New("length must be greater than 5") // can be left blank
		}
		if !strings.Contains(input, ":") {
			return errors.New("please add a tag")
		}
		return nil
	}

	imagePrompt := promptui.Prompt{
		Label:     "Image url with tag: ",
		Templates: utils.GetPromptTemplate(),
		Validate:  validateImage,
	}
	image, err := imagePrompt.Run()

	if image == "" {
		image = "<insert image url here>"
	}

	utils.GraceFullExit(err)
	return manifests.Container{
		Name:  name,
		Image: image,
	}
}

func InitContainerPrompt() (bool, []manifests.InitContainer) {
	promptInitContainer := promptui.Prompt{
		Label:     "Add init container?",
		IsConfirm: true,
	}
	initContainerResult, err := promptInitContainer.Run()
	utils.GraceFullExit(err)
	if initContainerResult == "y" || initContainerResult == "Y" {
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
		promptContainerCount := promptui.Prompt{
			Validate:  validate,
			Label:     "Number of containers: ",
			Templates: utils.GetPromptTemplate(),
		}
		containerCountResult, err := promptContainerCount.Run()
		utils.GraceFullExit(err)
		count, err := strconv.Atoi(containerCountResult)
		if err != nil {
			log.Fatal(err)
		}
		return true, InitContainerDetailList(count)
	}
	return false, []manifests.InitContainer{}
}

func InitContainerDetailList(count int) []manifests.InitContainer {
	containers := []manifests.InitContainer{}
	for i := 0; i < count; i++ {
		container := InitContainerDetailPrompt(i)
		containers = append(containers, container)
	}
	return containers
}

func InitContainerDetailPrompt(containerNumber int) manifests.InitContainer {
	validateName := func(input string) error {
		if len(input) < 3 {
			return errors.New("length must be greater than 3")
		}
		return nil
	}

	namePrompt := promptui.Prompt{
		Label:     "Name of container " + strconv.Itoa(containerNumber+1) + ": ",
		Validate:  validateName,
		Templates: utils.GetPromptTemplate(),
	}
	name, err := namePrompt.Run()
	utils.GraceFullExit(err)

	validateImage := func(input string) error {
		if len(input) == 0 {
			return nil
		}
		if len(input) > 0 && len(input) < 5 {
			return errors.New("length must be greater than 5") // can be left blank
		}
		if !strings.Contains(input, ":") {
			return errors.New("please add a tag")
		}
		return nil
	}

	imagePrompt := promptui.Prompt{
		Label:     "Image url with tag: ",
		Templates: utils.GetPromptTemplate(),
		Validate:  validateImage,
	}
	image, err := imagePrompt.Run()

	if image == "" {
		image = "<insert image url here>"
	}

	utils.GraceFullExit(err)

	return manifests.InitContainer{
		Name:  name,
		Image: image,
	}
}
