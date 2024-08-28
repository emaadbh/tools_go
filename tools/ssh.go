package tools

import (
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh"
	"log"
	"strings"
)

func SshTools() {
	red := color.New(color.FgRed, color.Bold)
	var host string
	var user string
	var pwd string
	red.Println("connect to server (ssh) and change ip in netplan (ubuntu) \n\n** Deletes all old IPs \n\n\n")

	fmt.Print("ip server main (e.g 192.168.1.1:22): ")
	_, _ = fmt.Scan(&host)

	fmt.Print("user (e.g root): ")
	_, _ = fmt.Scan(&user)

	fmt.Print("password: ")
	_, _ = fmt.Scan(&pwd)
	connectSSH(host, user, pwd)
}

func connectSSH(host string, user string, pwd string) {
	var ipNew string
	var gtwNew string

	conf := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(pwd),
		},
	}

	client, errSSH := ssh.Dial("tcp", host, conf)
	if errSSH != nil {
		fmt.Println(errSSH.Error())
	}
	defer client.Close()

	fmt.Println("ENTER NEW IP :")
	fmt.Scan(&ipNew)
	fmt.Println("ENTER NEW Gateway :")
	fmt.Scan(&gtwNew)

	networkName, _ := runCommand(client, "ls /sys/class/net | head -n 1")

	commands := []string{
		"netplan -h",
		" rm -rf /etc/netplan/*",
		netplanTEXT(networkName, ipNew, gtwNew),
		"netplan apply",
	}

	for _, cmd := range commands {
		_, err := runCommand(client, cmd)
		if err != nil {
			log.Fatalf("Failed to run command '%s': %s", cmd, err)
		}
		fmt.Println("DONE")
	}

}

func runCommand(client *ssh.Client, command string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func netplanTEXT(networkName string, ipNew string, gtwNew string) string {
	value := `cat <<EOF > /etc/netplan/config.yaml
network:
  version: 2
  renderer: networkd
  ethernets:
    @name:
      addresses:
        - @address/24
      routes:
        - to: 0.0.0.0/0
          via: @gateway
      dhcp4: no
      accept-ra: false
      nameservers:
        addresses:
          - 4.2.2.4
          - 8.8.4.4
EOF
`
	value = strings.Replace(value, "@name", networkName, -1)

	value = strings.Replace(value, "@address", ipNew, -1)
	value = strings.Replace(value, "@gateway", gtwNew, -1)

	return value
}
