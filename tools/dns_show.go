package tools

import (
	"fmt"
	"github.com/fatih/color"
	"golang/utils"
	"net"
)

func DnsShow() {
	red := color.New(color.FgRed, color.Bold)

	var domain string

	for {
		fmt.Print("Enter domain: ")
		_, _ = fmt.Scan(&domain)

		_, _ = red.Printf(utils.StarGenerator(25)+" Domain: %v "+utils.StarGenerator(25)+"\n", domain)

		aRecords, err := net.LookupIP(domain)
		printRecord(err, aRecords)

		mxRecords, errMX := net.LookupMX(domain)
		printRecord(errMX, mxRecords)

		nsRecords, errNs := net.LookupNS(domain)
		printRecord(errNs, nsRecords)

		if utils.Exit() {
			continue
		} else {
			break
		}
	}

}

func printRecord(err error, record interface{}) {
	green := color.New(color.FgGreen, color.Bold)

	if err != nil {
		fmt.Println(err)
	} else {

		switch v := record.(type) {

		case []*net.MX:
			fmt.Printf(utils.ConstValue(), "MX")

			for _, mx := range v {
				_, _ = green.Println(" $ MX: ", mx)
			}
		case []net.IP:
			fmt.Printf(utils.ConstValue(), "A")
			for _, ip := range v {
				_, _ = green.Println(" $ A: ", ip)
			}
		case []*net.NS:
			fmt.Printf(utils.ConstValue(), "NS")
			for _, ip := range v {
				_, _ = green.Println(" $ NS: ", ip)
			}
		}

	}
}
