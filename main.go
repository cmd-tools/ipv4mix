package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// ConvertToHex converts an integer to its hexadecimal representation.
func toHex(ipPart int) string {
	return fmt.Sprintf("0x%X", ipPart)
}

// ConvertToOctal converts an integer to its octal representation.
func toOct(ipPart int) string {
	return fmt.Sprintf("0%o", ipPart)
}

// ValidateIP validates the input IP address.
func ValidateIP(ip string) (net.IP, error) {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil || parsedIP.To4() == nil {
		return nil, fmt.Errorf("invalid IPv4 address: %s", ip)
	}
	return parsedIP.To4(), nil
}

// parseIPPart parses a part of an IP address, correctly handling octal and decimal values.
func parseIPPart(part string) (int, error) {
	// Check if the part is octal (starts with 0 and has more than one digit)
	if len(part) > 1 && part[0] == '0' {
		val, err := strconv.ParseInt(part, 8, 0) // Parse as octal
		if err != nil {
			return 0, err
		}
		return int(val), nil // Convert to int
	}
	// Otherwise, parse it as a decimal
	val, err := strconv.Atoi(part)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ipv4mix <IPv4 address>")
		return
	}

	inputIP := os.Args[1]

	_, err := ValidateIP(inputIP)
	if err != nil {
		fmt.Printf("Invalid IP address: %s\n", inputIP)
	}

	parts := strings.Split(inputIP, ".")

	// Validate and parse IP parts
	octets := make([]int, 4)
	switch len(parts) {
	case 1:
		// Single part, interpret as a full 32-bit integer
		fullIP, err := strconv.ParseUint(parts[0], 10, 32)
		if err != nil {
			fmt.Println("Invalid IP address:", inputIP)
			return
		}
		octets[0] = int(fullIP >> 24 & 0xFF)
		octets[1] = int(fullIP >> 16 & 0xFF)
		octets[2] = int(fullIP >> 8 & 0xFF)
		octets[3] = int(fullIP & 0xFF)
	case 2:
		// Two parts
		p1, err := parseIPPart(parts[0])
		if err != nil {
			fmt.Println("Invalid IP address part:", parts[0])
			return
		}
		p2, err := strconv.ParseUint(parts[1], 10, 32)
		if err != nil {
			fmt.Println("Invalid IP address part:", parts[1])
			return
		}
		octets[0] = p1
		octets[1] = int(p2 >> 16 & 0xFF)
		octets[2] = int(p2 >> 8 & 0xFF)
		octets[3] = int(p2 & 0xFF)
	case 3:
		// Three parts
		p1, err := parseIPPart(parts[0])
		if err != nil {
			fmt.Println("Invalid IP address part:", parts[0])
			return
		}
		p2, err := parseIPPart(parts[1])
		if err != nil {
			fmt.Println("Invalid IP address part:", parts[1])
			return
		}
		p3, err := strconv.ParseUint(parts[2], 10, 32)
		if err != nil {
			fmt.Println("Invalid IP address part:", parts[2])
			return
		}
		octets[0] = p1
		octets[1] = p2
		octets[2] = int(p3 >> 8 & 0xFF)
		octets[3] = int(p3 & 0xFF)
	case 4:
		// Four parts
		for i := 0; i < 4; i++ {
			val, err := parseIPPart(parts[i])
			if err != nil {
				fmt.Println("Invalid IP address part:", parts[i])
				return
			}
			octets[i] = val
		}
	default:
		fmt.Println("Invalid IP address format")
		return
	}

	// 1. Dotted Decimal
	fmt.Printf("%d.%d.%d.%d\n", octets[0], octets[1], octets[2], octets[3])

	// 2. Dotted Hexadecimal
	fmt.Printf("%s.%s.%s.%s\n", toHex(octets[0]), toHex(octets[1]), toHex(octets[2]), toHex(octets[3]))

	// 3. Dotted Octal
	fmt.Printf("%s.%s.%s.%s\n", toOct(octets[0]), toOct(octets[1]), toOct(octets[2]), toOct(octets[3]))

	// 4. Mixed Decimal, Hex, Octal Combinations
	fmt.Printf("%d.%s.%d.%d\n", octets[0], toHex(octets[1]), octets[2], octets[3])
	fmt.Printf("%d.%d.%s.%d\n", octets[0], octets[1], toHex(octets[2]), octets[3])
	fmt.Printf("%d.%d.%d.%s\n", octets[0], octets[1], octets[2], toHex(octets[3]))

	fmt.Printf("%d.%s.%d.%d\n", octets[0], toOct(octets[1]), octets[2], octets[3])
	fmt.Printf("%d.%d.%s.%d\n", octets[0], octets[1], toOct(octets[2]), octets[3])
	fmt.Printf("%d.%d.%d.%s\n", octets[0], octets[1], octets[2], toOct(octets[3]))

	fmt.Printf("%s.%d.%d.%d\n", toHex(octets[0]), octets[1], octets[2], octets[3])
	fmt.Printf("%s.%s.%d.%d\n", toHex(octets[0]), toHex(octets[1]), octets[2], octets[3])
	fmt.Printf("%s.%d.%s.%d\n", toHex(octets[0]), octets[1], toHex(octets[2]), octets[3])
	fmt.Printf("%s.%d.%d.%s\n", toHex(octets[0]), octets[1], octets[2], toHex(octets[3]))

	fmt.Printf("%s.%s.%d.%d\n", toOct(octets[0]), toOct(octets[1]), octets[2], octets[3])
	fmt.Printf("%d.%s.%d.%s\n", octets[0], toHex(octets[1]), octets[2], toHex(octets[3]))
	fmt.Printf("%d.%s.%s.%d\n", octets[0], toHex(octets[1]), toHex(octets[2]), octets[3])

	// 5. Single Decimal
	decimalIP := (octets[0] << 24) | (octets[1] << 16) | (octets[2] << 8) | octets[3]
	fmt.Printf("%d\n", decimalIP)

	// 6. Single Hexadecimal
	fmt.Printf("0x%X\n", decimalIP)

	// 7. Single Octal
	fmt.Printf("0%o\n", decimalIP)

	// 8. Last two octets combined into a single decimal
	combinedLastTwo := (octets[2] << 8) | octets[3]
	fmt.Printf("%d.%d.%d\n", octets[0], octets[1], combinedLastTwo)
	fmt.Printf("%d.%s.%d\n", octets[0], toHex(octets[1]), combinedLastTwo)
	fmt.Printf("%d.%s.%d\n", octets[0], toOct(octets[1]), combinedLastTwo)

	// 9. Last three octets combined into a single decimal
	combinedLastThree := (octets[1] << 16) | (octets[2] << 8) | octets[3]
	fmt.Printf("%d.%d\n", octets[0], combinedLastThree)
	fmt.Printf("%s.%d\n", toHex(octets[0]), combinedLastThree)
	fmt.Printf("%s.%d\n", toOct(octets[0]), combinedLastThree)

	// 10. First two octets in octal, last two combined into a single decimal
	combinedLastTwoDecimal := (octets[2] << 8) | octets[3]
	fmt.Printf("%s.%s.%d\n", toOct(octets[0]), toOct(octets[1]), combinedLastTwoDecimal)
	fmt.Printf("%s.%s.%s\n", toOct(octets[0]), toOct(octets[1]), toHex(combinedLastTwoDecimal))
	fmt.Printf("%s.%s.%s\n", toOct(octets[0]), toOct(octets[1]), toOct(combinedLastTwoDecimal))
}
