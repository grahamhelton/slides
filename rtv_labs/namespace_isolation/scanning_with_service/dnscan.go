package main

// https://raw.githubusercontent.com/LowOrbitSecurity/dnscan/refs/heads/main/main.go
// Forked from https://gist.github.com/nirohfeld/c596898673ead369cb8992d97a1c764e
// All credit goes to github.com/nirohfeld

import (
    "flag"
    "fmt"
    "net"
    "os"
    "strings"
    "sync"
    "time"
    "github.com/fatih/color"
)

func performLookup(ip net.IP, results chan<- string, retryCount int) {
    var err error
    var names []string
    for i := 0; i < retryCount; i++ {
        names, err = net.LookupAddr(ip.String())
        if err == nil {
            break
        }
        time.Sleep(10 * time.Millisecond) // Shorter delay between retries
    }

    if err == nil && len(names) > 0 {
        fmt.Println(color.GreenString(ip.String()) + "," + names[0])
    }
}

func main() {
    fmt.Println(color.BlueString("Starting Scan"))
    fmt.Println(color.BlueString("------------"))
    var inputFlag string
    flag.StringVar(&inputFlag, "subnet", "", "Input to scan, CIDR notation (e.g., 10.5.0.0/24) or wildcard (e.g., 10.5.0.*)")
    flag.Parse()

    if inputFlag == "" {
        fmt.Println("Usage: ./dnscan -subnet <CIDR or Wildcard>")
        os.Exit(1)
    }

    ips, err := parseInput(inputFlag)
    if err != nil {
        fmt.Printf("Error processing subnet: %v\n", err)
        os.Exit(1)
    }

    results := make(chan string)
    retryCount := 2

    var wg sync.WaitGroup
    rateLimit := time.Tick(1 * time.Millisecond)

    for _, ip := range ips {
        wg.Add(1)
        <-rateLimit
        go func(ip net.IP) {
            defer wg.Done()
            performLookup(ip, results, retryCount)
        }(ip)
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    for r := range results {
        fmt.Println(r)
    }
}

func parseInput(input string) ([]net.IP, error) {
    if strings.Contains(input, "*") {
        // Convert wildcard to CIDR
        cidrs, err := wildcardToCIDRs(input)
        if err != nil {
            return nil, err
        }
        var ips []net.IP
        for _, cidr := range cidrs {
            _, ipnet, _ := net.ParseCIDR(cidr)
            ips = append(ips, generateIPsFromCIDR(ipnet)...)
        }
        return ips, nil
    } else {
        _, ipnet, err := net.ParseCIDR(input)
        if err != nil {
            return nil, fmt.Errorf("invalid input: %v", err)
        }
        return generateIPsFromCIDR(ipnet), nil
    }
}

func wildcardToCIDRs(wildcard string) ([]string, error) {
    parts := strings.Split(wildcard, ".")
    if len(parts) != 4 {
        return nil, fmt.Errorf("invalid wildcard format")
    }

    cidr := ""
    for i, part := range parts {
        if part == "*" {
            parts[i] = "0"
        }
        cidr += parts[i]
        if i < 3 {
            cidr += "."
        }
    }

    wildcardCount := strings.Count(wildcard, "*")
    if wildcardCount < 1 || wildcardCount > 4 {
        return nil, fmt.Errorf("invalid wildcard usage")
    }

    // Calculate CIDR block size based on wildcard count
    cidrBlockSize := 32 - (wildcardCount * 8)
    cidr += fmt.Sprintf("/%d", cidrBlockSize)

    return []string{cidr}, nil
}

func generateIPsFromCIDR(ipnet *net.IPNet) []net.IP {
    var ips []net.IP
    for ip := ipnet.IP.Mask(ipnet.Mask); ipnet.Contains(ip); incrementIP(ip) {
        ips = append(ips, net.ParseIP(ip.String()))
    }
    return ips
}

func incrementIP(ip net.IP) {
    for j := len(ip) - 1; j >= 0; j-- {
        ip[j]++
        if ip[j] > 0 {
            break
        }
    }
}

