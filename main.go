package main

import (
	"fmt"
	"github.com/fatih/color"
	"golang/tools"
	"golang/utils"
)

func main() {
	blue := color.New(color.FgBlue, color.Bold)
	yellow := color.New(color.FgYellow, color.Bold)

	_, _ = blue.Println(utils.EmadGenerator())

	for {
		var index int
		_, _ = blue.Println("1 - DNS CHECKER")
		_, _ = blue.Println("2 - SSL CHECKER")
		_, _ = blue.Println("3 - Ping (Rang / Single)")
		_, _ = blue.Println("4 - connect to SSH and set new IP (netplan)")

		_, _ = yellow.Println("Choose the tool you want :")

		fmt.Scan(&index)
		switchTools(index)

		if utils.Exit() {
			continue
		} else {
			break
		}

	}

}

func switchTools(index int) {
	switch index {
	case 1:
		tools.DnsShow()
	case 2:
		tools.SslChecker()
	case 3:
		tools.RangePing()
	case 4:
		tools.SshTools()
	}
}
