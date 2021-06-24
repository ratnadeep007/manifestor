package utils

import "github.com/manifoldco/promptui"

func GetPromptTemplate() *promptui.PromptTemplates {
	return &promptui.PromptTemplates{
		Prompt:  "{{ . }}: ",
		Valid:   "\U00002705 {{ . | green }}",
		Invalid: "\U00002717 {{ . | red}}",
		Success: "{{ . | bold | green }}",
	}

}
