package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input", err)
	}

}

func checkDomain(domain string) {
	var hasDMARC, hasMX, hasSPF bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Println("error ", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}
	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Println("error ", err)
	}

	for _, r := range txtRecords {
		if strings.HasPrefix(r, "v=spf1") {
			hasSPF = true
			spfRecord = r
			break
		}

	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Println("error ", err)

	}
	for _, r := range dmarcRecords {
		if strings.HasPrefix(r, "v=dmarc1") {
			hasDMARC = true
			dmarcRecord = r
			break
		}
	}
	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
	fmt.Println()
}
