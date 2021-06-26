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

func GetSelectTemplate() *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\U00002192 {{ . | red }}",
		Inactive: "  {{ . | cyan }}",
		Selected: "  {{ . | green | bold }}",
	}
}
