package command

import (
	"fmt"

	"github.com/mrtc0/wazuh"
)

func GetAllAgents(client *wazuh.Client) {
	agents, _ := client.GetAllAgents()
	fmt.Printf("%-20v%-20v%-20v%-20v%-20v\n", "ID", "STATUS", "NAME", "IP", "OS")
	for _, agent := range *agents {
		fmt.Printf("%-20v%-20v%-20v%-20v%-20v\n", agent.ID, agent.Status, agent.Name, agent.IP, agent.Os.Name)
	}
}
