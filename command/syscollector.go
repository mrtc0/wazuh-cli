package command

import (
	"fmt"
	"html/template"
	"os"
	"text/tabwriter"

	"github.com/mrtc0/wazuh"
)

func GetHardwareInfo(client *wazuh.Client, agent_id string) {
	hw, _ := client.GetHardwareInfo(agent_id)
	text := `RAM:
  Usage: {{.RAM.Usage}}
  Total: {{.RAM.Total}}
  Free: {{.RAM.Free}}
CPU:
  Cores: {{.CPU.Cores}}
  Mhz: {{.CPU.Mhz}}
  Name: {{.CPU.Name}}
Scan:
  ID: {{.Scan.ID}}
  Time: {{.Scan.Time}}
`

	tmpl, err := template.New("hw").Parse(text)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, hw)
	if err != nil {
		panic(err)
	}
}

func GetPackagesInfo(client *wazuh.Client, agent_id string) {
	packages, _ := client.GetPackagesInfo(agent_id)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.DiscardEmptyColumns)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", "Name", "Version", "Source", "Description")
	for _, p := range *packages {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", p.Name, p.Version, p.Source, p.Description)
	}
	w.Flush()
}

func GetNetworkAddrInfo(client *wazuh.Client, agent_id string) {
	nwaddr, _ := client.GetNetworkAddrInfo(agent_id)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.DiscardEmptyColumns)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", "Address", "Netmask", "Broadcast", "Proto")
	for _, n := range *nwaddr {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", n.Address, n.Netmask, n.Broadcast, n.Proto)
	}
	w.Flush()
}

func GetOpenPorts(client *wazuh.Client, agent_id string) {
	ports, _ := client.GetPortsInfo(agent_id)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', tabwriter.DiscardEmptyColumns)
	fmt.Fprintf(w, "%s\t%s\t%s\n", "Proto", "Local", "State")
	for _, p := range *ports {
		fmt.Fprintf(w, "%s\t%s:%d\t%s\n", p.Protocol, p.Local.IP, p.Local.Port, p.State)
	}
	w.Flush()
}

func GetProcesses(client *wazuh.Client, agent_id string) {
	processes, _ := client.GetProcessesInfo(agent_id)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', tabwriter.DiscardEmptyColumns)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", "User", "PID", "STAT", "TIME", "Name", "CMD")
	for _, p := range *processes {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", p.Euser, p.Pid, p.State, p.Name, p.Cmd)
	}
	w.Flush()
}
