package main

import (
	"fmt"
	"log"

	"github.com/go-routeros/routeros"
)

func main() {

	operationRouter()
}
func operationRouter() {
	// Connection details to MikroTik
	routerAddress := "192.168.1.1"
	routerUsername := "admin"
	routerPassword := "password"

	// Connect to MikroTik
	client, err := routeros.Dial(routerAddress, routerUsername, routerPassword)
	if err != nil {
		log.Fatalf("Unable to connect to the router: %v", err)
	}
	defer client.Close()

	// Block HTTP and HTTPS requests
	blockHTTP(client)
	blockHTTPS(client)

	fmt.Println("Blocking HTTP and HTTPS requests completed.")
}

// blockHTTP Firewall filter settings to block HTTP requests
func blockHTTP(client *routeros.Client) {
	cmd := "/ip/firewall/filter/add"
	params := map[string]string{
		"chain":    "forward",
		"action":   "drop",
		"protocol": "tcp",
		"dst-port": "80",
		"comment":  "Block HTTP",
		"disabled": "no",
	}

	_, err := client.Call(cmd, params)
	if err != nil {
		log.Printf("Failed to block HTTP: %v", err)
	} else {
		fmt.Println("HTTP blocking rule added.")
	}
}

// blockHTTPS Firewall filter settings to block HTTPS requests
func blockHTTPS(client *routeros.Client) {
	cmd := "/ip/firewall/filter/add"
	params := map[string]string{
		"chain":    "forward",
		"action":   "drop",
		"protocol": "tcp",
		"dst-port": "443",
		"comment":  "Block HTTPS",
		"disabled": "no",
	}

	_, err := client.Call(cmd, params)
	if err != nil {
		log.Printf("Failed to block HTTPS: %v", err)
	} else {
		fmt.Println("HTTPS blocking rule added.")
	}
}
