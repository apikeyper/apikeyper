package main

import (
	"apikeyper/internal/server"
	"fmt"
	"log/slog"
)

func main() {

	server := server.NewServer()

	slog.Info(fmt.Sprintf("Server listening on %s", server.Addr))
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
