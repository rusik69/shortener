package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rusik69/shortener/pkg/server"
)

func main() {
	var (
		port     = flag.String("port", "8080", "listen port")
		host     = flag.String("host", "localhost", "listen host")
		baseURL  = flag.String("base-url", "http://localhost:8080", "Base URL for shortened links")
		dbPath   = flag.String("db", "urls.db", "Path to database file")
		showHelp = flag.Bool("help", false, "Show help message")
	)

	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	log.Printf("Starting URL shortener service")
	log.Printf("Port: %s", *port)
	log.Printf("Host: %s", *host)
	log.Printf("Base URL: %s", *baseURL)
	log.Printf("Database: %s", *dbPath)

	if err := server.Run(*port, *host, *baseURL, *dbPath); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
