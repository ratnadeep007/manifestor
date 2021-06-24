package cli

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func Run() {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\U00002192 {{ . | red }}",
		Inactive: "  {{ . | cyan }}",
		Selected: "  {{ . | green | bold }}",
	}

	promptKind := promptui.Select{
		Label:     "Select Kind",
		Items:     []string{"Deployment", "Service"},
		Templates: templates,
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
