package cli

import (
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
	utils.GraceFullExit(err)

	switch result {
	case "Deployment":
		DeploymentQuestions()
	case "Service":
		ServiceQuestions()
	}

}
