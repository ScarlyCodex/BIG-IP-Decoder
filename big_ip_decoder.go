package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)


func printBanner() {
	banner := `
 ____  __    ___        __   ____        ____  ____   ___   __   ____  ____  ____ 
(  _ \(  )  / __) ___  (  ) (  _ \      (    \(  __) / __) /  \ (    \(  __)(  _ \
 ) _ ( )(  ( (_ \(___)  )(   ) __/       ) D ( ) _) ( (__ (  O ) ) D ( ) _)  )   /
(____/(__)  \___/      (__) (__)        (____/(____) \___) \__/ (____/(____)(__\_)
          
 `
	color.Cyan(banner)
}

// Function to decode the BIG-IP cookie value and extract the internal IP address and port
func decodeBigIP(value string) (string, int) {
	// Split the cookie value into IP and port components
	parts := strings.Split(value, ".")
	if len(parts) < 2 {
		return "Invalid BIG-IP cookie format.", 0
	}

	// Convert the IP part to an integer
	ipDec, err := strconv.Atoi(parts[0])
	if err != nil {
		return "Error converting the IP part: " + err.Error(), 0
	}

	// Convert the decimal value to hexadecimal
	hexIP := fmt.Sprintf("%08x", ipDec)

	// Split the address into octets and reverse the order (little-endian)
	var ipParts []string
	for i := 0; i < len(hexIP); i += 2 {
		ipParts = append(ipParts, hexIP[i:i+2])
	}
	// Reverse the octet order
	for i, j := 0, len(ipParts)-1; i < j; i, j = i+1, j-1 {
		ipParts[i], ipParts[j] = ipParts[j], ipParts[i]
	}

	// Convert each part to decimal
	var decodedIP []string
	for _, part := range ipParts {
		ip, _ := strconv.ParseInt(part, 16, 64)
		decodedIP = append(decodedIP, strconv.Itoa(int(ip)))
	}

	// Convert the port part
	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return strings.Join(decodedIP, "."), 0
	}

	// Force the port to 443 if required
	port = 443

	// Return the decoded IP and the forced port
	return strings.Join(decodedIP, "."), port
}

// Function to extract the pool name from the cookie
func extractPoolName(cookieName string) string {
	// Regex-like operation to extract the pool name
	poolParts := strings.Split(cookieName, "BIGipServer")
	if len(poolParts) > 1 {
		// Extract the first part of the pool name (before the underscore)
		poolNameParts := strings.Split(poolParts[1], "_")
		if len(poolNameParts) > 0 {
			return poolNameParts[0]
		}
	}
	return "Unknown"
}

// Main function
func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <URL>\n", os.Args[0])
		return
	}

	url := os.Args[1]
	printBanner()

	color.Yellow("[*] URL to request: %s\n", url)

	// Create custom HTTP client with SSL certificate verification disabled
	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: customTransport}

	// Make the HTTP request
	resp, err := client.Get(url)
	if err != nil {
		color.Red("Error in the HTTP request: %s", err)
		return
	}
	defer resp.Body.Close()

	// Check the cookies to find the BIG-IP cookie
	for _, cookie := range resp.Cookies() {
		if strings.HasPrefix(cookie.Name, "BIGipServer") {
			color.Green("[*] Cookie to decode: %s=%s\n", cookie.Name, cookie.Value)
			
			// Extract the pool name
			poolName := extractPoolName(cookie.Name)

			// Decode the IP and port from the cookie value
			decodedIP, port := decodeBigIP(cookie.Value)

			// Print the formatted output with colors
			color.Cyan("[*] Pool name: %s", poolName)
			color.Cyan("[*] Decoded IP and Port: %s:%d\n", decodedIP, port)
			return
		}
	}

	color.Red("No BIG-IP cookie found in the response.")
}
