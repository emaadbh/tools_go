package tools

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/go-ping/ping"
	"strings"
)

func RangePing() {
	red := color.New(color.FgRed, color.Bold)

	var baseIP string

	red.Println("Enter the base IP (range: e.g. 192.168.1 ,single e.g 192.168.1.1): ")
	_, _ = fmt.Scan(&baseIP)

	sections := strings.Split(baseIP, ".")
	if len(sections) == 4 {
		pingIP(baseIP)
	} else if len(sections) == 2 {
		pingIP(baseIP)

	} else if len(sections) >= 4 {
		red.Println("WTF .. :/")
	} else {
		rangIpPing(baseIP)
	}
}

func rangIpPing(baseIP string) {

	var startRange, endRange int

	fmt.Print("Enter the start range (e.g., 1): ")
	_, _ = fmt.Scan(&startRange)

	fmt.Print("Enter the end range (e.g., 254): ")
	_, _ = fmt.Scan(&endRange)

	// پینگ کردن IPها در رنج مشخص شده
	for i := startRange; i <= endRange; i++ {
		ip := fmt.Sprintf("%s.%d", baseIP, i)
		fmt.Printf("----------------%v----------------\n", ip)
		pingIP(ip)
	}
}

func pingIP(ip string) {
	green := color.New(color.FgGreen, color.Bold)

	pinger, err := ping.NewPinger(ip)
	pinger.SetPrivileged(true)

	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}

	state := pinger.Statistics()

	green.Printf("ave : %v - lost: %v - MinRTT: %v - MaxRtt: %v \n", state.AvgRtt, state.PacketLoss, state.MinRtt, state.MaxRtt)
	green.Printf("IP : %v - PacketsSent: %v - PacketsRecv: %v - PacketsRecvDuplicates: %v \n", state.IPAddr, state.PacketsSent, state.PacketsRecv, state.PacketsRecvDuplicates)

}
