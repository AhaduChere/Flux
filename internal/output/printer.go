package output

import (
	"flux/internal/request"
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func Print(resp request.Response) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("green"))
	fmt.Println(style.Render(fmt.Sprintf("%d", resp.StatusCode)))
	fmt.Println(style.Render(fmt.Sprintf("%v", resp.Duration)))
	fmt.Println(style.Render(fmt.Sprintf("%s", resp.Body)))
}
