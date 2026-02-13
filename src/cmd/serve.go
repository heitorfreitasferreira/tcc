package cmd

import (
	"fmt"
	"net/http"
	"strings"
	"tcc/web"
	"time"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a web server with a Hello World page",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		addr, err := cmd.Flags().GetString("addr")
		if err != nil {
			return fmt.Errorf("get --addr: %w", err)
		}

		handler, err := web.NewHandler()
		if err != nil {
			return err
		}

		server := &http.Server{
			Addr:              addr,
			Handler:           handler,
			ReadHeaderTimeout: 5 * time.Second,
		}

		cmd.Printf("web server running at %s\n", formatServerURL(addr))

		err = server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("start web server: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().String("addr", ":8080", "Address where the web server listens")
}

func formatServerURL(addr string) string {
	if strings.HasPrefix(addr, ":") {
		return "http://localhost" + addr
	}

	if strings.HasPrefix(addr, "0.0.0.0:") {
		port := strings.TrimPrefix(addr, "0.0.0.0:")
		return "http://localhost:" + port
	}

	return "http://" + addr
}
