package main

import (
	files "backend/src/core/files/usecase"
	"backend/src/internal/cloud/objectstorage/container"
	"backend/src/internal/middleware"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct{}

func (mdl model) Init() tea.Cmd                           { return nil }
func (mdl model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return mdl, nil }
func (mdl model) View() string                            { return "Backend running on :8081\nPress Ctrl+C to exit.\n" }

func main() {
	port := ":8081"
	mux := http.NewServeMux()
	corsHandler := middleware.EnableCORS(mux)

	upload := files.UploadFile{}
	upload.RegisterRoutes(mux)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	err := objectstorage.CreateNamespaceAndMount("/tmp/bucket-test", "/tmp/ns-test")
	if err != nil {
		fmt.Errorf("error creating namespace and mount: %w", err)
	}

	link, err := os.Readlink("/proc/self/ns/mnt")
	if err != nil {
		fmt.Errorf("failed to read mount namespace link: %w", err)
	}

	fmt.Println("Current mount namespace:", link)

	go func() {
		fmt.Printf("starting backend on port %s\n", port)
		log.Fatal(http.ListenAndServe(":8081", corsHandler))

		err := http.ListenAndServe(port, mux)
		if err != nil {
			fmt.Errorf("error starting backend", err)
			os.Exit(1)
		}
	}()

	//go func() {
	//	p := tea.NewProgram(initialModel())
	//	if _, err := p.Run(); err != nil {
	//		log.Fatalf("TUI error: %v", err)
	//	}
	//}()

	<-ctx.Done()
	fmt.Println("\nShutting down server...")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("server shutdown error %e", err)
	}

	fmt.Println("Exiting...")

}

func initialModel() tea.Model {
	return model{}
}
