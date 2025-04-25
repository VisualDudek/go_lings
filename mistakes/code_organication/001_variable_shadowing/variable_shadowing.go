//go:build ignore

package main

import (
	"log"
	"net/http"
)

// są dwa rozwiązania na shadowing 1) temp var w nested block
// 2)

func main() {

	var client *http.Client // this will be always nil
	tracing := true

	if tracing {
		client, err := createClientWithTracing()
		if err != nil {
			return err
		}
		log.Print(client)
	} else {
		client, err := createDefaultClient()
		if err != nil {
			return err
		}
		log.Print(client)
	}

	log.Print(client)

	// --- temp var fix ---
	if tracing {
		c, err := createClientWithTracing()
		if err != nil {
			return err
		}
		client = c // FIX
		log.Print(client)

	} else {
		c, err := createDefaultClient()
		if err != nil {
			return err
		}
		client = c // FIX
		log.Print(client)
	}

	// --- declare err fix ---

	var err error

	if tracing {
		client, err = createClientWithTracing()
	} else {
		client, err = createDefaultClient()
	}

	if err != nil {
		return err
	}

}
