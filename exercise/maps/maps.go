//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printServerStatus(servers map[string]int) {
	fmt.Println()
	fmt.Println("    --Server Status--")
	for key, value := range servers {
		fmt.Println(key, "is in", value, "mode")
	}
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	mapServers := make(map[string]int)
	for _, server := range servers {
		mapServers[server] = Online
	}

	printServerStatus(mapServers)

	mapServers["darkstar"] = Retired
	mapServers["aiur"] = Offline
	printServerStatus(mapServers)

	for key := range mapServers {
		mapServers[key] = Maintenance
	}
	printServerStatus(mapServers)

}
