package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	port := ":8081"

	go func() {
		fmt.Printf("starting backend on port %s\n", port)

		err := http.ListenAndServe(port, mux)
		if err != nil {
			fmt.Errorf("error starting backend", err)
			os.Exit(1)
		}
	}()

	//p := tea.NewProgram(initialModel())
	//if _, err := p.Run(); err != nil {
	//	fmt.Errorf("tui error", err)
	//}
}
