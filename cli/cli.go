package cli

import (
	"fmt"
	"manifest_creator/utils"

	"github.com/manifoldco/promptui"
)

func Run() {

	promptKind := promptui.Select{
		Label:     "Select Kind",
		Items:     []string{"Deployment", "Service"},
		Templates: utils.GetSelectTemplate(),
	}

	_, result, err := promptKind.Run()
	if err != nil {
		fmt.Printf("Failed %v\n", err)
		return
	}

	switch result {
	case "Deployment":
		DeploymentQuestions()
	case "Service":
		ServiceQuestions()
	}

}
