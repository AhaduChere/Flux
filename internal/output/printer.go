package output

import (
	"bytes"
	"encoding/json"
	"flux/internal/request"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"net/http"
)

// NOTE:Prints response
func Print(resp request.Response) {
	color := "2"
	symbol := "✓"
	if resp.StatusCode >= 400 {
		color = "1"
		symbol = "✗"
	}

	style := lipgloss.NewStyle().Foreground(lipgloss.Color(color))

	//NOTE: Print status code and duration
	fmt.Println(style.Render(fmt.Sprintf("%s %d %s · %dms", symbol, resp.StatusCode, http.StatusText(resp.StatusCode), resp.Duration.Milliseconds())))

	//NOTE: Print Body
	var indented bytes.Buffer
	json.Indent(&indented, []byte(resp.Body), "", "  ")
	fmt.Println(indented.String())
}
