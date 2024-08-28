package tools

import (
	"crypto/tls"
	"fmt"
	"github.com/fatih/color"
	"golang/utils"
	"net"
	"time"
)

func SslChecker() {
	green := color.New(color.FgGreen, color.Bold)
	red := color.New(color.FgRed, color.Bold)

	var domain string

	for {

		fmt.Print("Enter domain: ")
		_, _ = fmt.Scan(&domain)

		//dialer := &net.Dialer{Timeout: 10 * time.Second}
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
		}

		conn, _ := net.DialTimeout("tcp", domain+":443", 5*time.Second)
		defer conn.Close()

		conect := tls.Client(conn, tlsConfig)

		err := conect.Handshake()
		if err != nil {
			red.Printf("TLS handshake failed: %v", err)
		}

		// Get the certificate chain
		certs := conect.ConnectionState().PeerCertificates

		// Loop over the certificates and print details
		for _, cert := range certs {
			green.Printf("Issuer: %s\n", cert.Issuer.Organization)
			green.Printf("Subject: %s\n", cert.Subject.CommonName)
			green.Printf("\n SSL TIME : %v -- DateEX: %v \n", CalculateRemainingDays(cert.NotAfter), cert.NotAfter)
			break
		}

		//Expire Date

		red.Println("\n ------------------END-----------------")

		if utils.Exit() {
			continue
		} else {
			break
		}
	}

}

func CalculateRemainingDays(expiryDate time.Time) int {
	return int(expiryDate.Sub(time.Now()).Hours() / 24)
}
