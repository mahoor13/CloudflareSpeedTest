package task

import (
	"bufio"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const defaultInputFile = "ip.txt"

var (
	// TestAll test all IP addresses
	TestAll = false
	// IPFile is the filename of IP Ranges
	IPFile = defaultInputFile
	IPText string
)

func InitRandSeed() {
	rand.Seed(time.Now().UnixNano())
}

func isIPv4(ip string) bool {
	return strings.Contains(ip, ".")
}

func randIPEndWith(num byte) byte {
	if num == 0 { // For single IP addresses like /32
		return byte(0)
	}
	return byte(rand.Intn(int(num)))
}

type IPRanges struct {
	ips     []*net.IPAddr
	mask    string
	firstIP net.IP
	ipNet   *net.IPNet
}

func newIPRanges() *IPRanges {
	return &IPRanges{
		ips: make([]*net.IPAddr, 0),
	}
}

// If it is a single IP, add the subnet mask; otherwise, get the subnet mask (r.mask)
func (r *IPRanges) fixIP(ip string) string {
	// If it does not contain '/', it is not an IP range but a single IP, so add /32 or /128 subnet mask accordingly
	if i := strings.IndexByte(ip, '/'); i < 0 {
		if isIPv4(ip) {
			r.mask = "/32"
		} else {
			r.mask = "/128"
		}
		ip += r.mask
	} else {
		r.mask = ip[i:]
	}
	return ip
}

// Parse IP range, obtain IP, IP range, and subnet mask
func (r *IPRanges) parseCIDR(ip string) {
	var err error
	if r.firstIP, r.ipNet, err = net.ParseCIDR(r.fixIP(ip)); err != nil {
		log.Fatalln("ParseCIDR err", err)
	}
}

func (r *IPRanges) appendIPv4(d byte) {
	r.appendIP(net.IPv4(r.firstIP[12], r.firstIP[13], r.firstIP[14], d))
}

func (r *IPRanges) appendIP(ip net.IP) {
	r.ips = append(r.ips, &net.IPAddr{IP: ip})
}

// Return the minimum value and available number of the fourth segment of the IP address
func (r *IPRanges) getIPRange() (minIP, hosts byte) {
	minIP = r.firstIP[15] & r.ipNet.Mask[3] // Minimum value of the fourth segment of the IP address

	// Get the number of hosts based on the subnet mask
	m := net.IPv4Mask(255, 255, 255, 255)
	for i, v := range r.ipNet.Mask {
		m[i] ^= v
	}
	total, _ := strconv.ParseInt(m.String(), 16, 32) // Total available IP addresses
	if total > 255 {                                 // Correct the available number of the fourth segment
		hosts = 255
		return
	}
	hosts = byte(total)
	return
}

func (r *IPRanges) chooseIPv4() {
	if r.mask == "/32" { // For single IP addresses, no need to randomize, directly add itself
		r.appendIP(r.firstIP)
	} else {
		minIP, hosts := r.getIPRange()    // Return the minimum value and available number of the fourth segment of the IP address
		for r.ipNet.Contains(r.firstIP) { // Continue looping as long as the IP address does not exceed the IP range
			if TestAll { // If testing all IP addresses
				for i := 0; i <= int(hosts); i++ { // Traverse from the minimum value to the maximum value of the fourth segment of the IP address
					r.appendIPv4(byte(i) + minIP)
				}
			} else { // Randomize the last segment of the IP address 0.0.0.X
				r.appendIPv4(minIP + randIPEndWith(hosts))
			}
			r.firstIP[14]++ // 0.0.(X+1).X
			if r.firstIP[14] == 0 {
				r.firstIP[13]++ // 0.(X+1).X.X
				if r.firstIP[13] == 0 {
					r.firstIP[12]++ // (X+1).X.X.X
				}
			}
		}
	}
}

func (r *IPRanges) chooseIPv6() {
	if r.mask == "/128" { // For single IP addresses, no need to randomize, directly add itself
		r.appendIP(r.firstIP)
	} else {
		var tempIP uint8                  // Temporary variable to record the value of the previous position
		for r.ipNet.Contains(r.firstIP) { // Continue looping as long as the IP address does not exceed the IP range
			r.firstIP[15] = randIPEndWith(255) // Randomize the last segment of the IP address
			r.firstIP[14] = randIPEndWith(255) // Randomize the last segment of the IP address

			targetIP := make([]byte, len(r.firstIP))
			copy(targetIP, r.firstIP)
			r.appendIP(targetIP) // Add to the IP address pool

			for i := 13; i >= 0; i-- { // Randomize from the third-to-last segment
				tempIP = r.firstIP[i]              // Save the value of the previous segment
				r.firstIP[i] += randIPEndWith(255) // Randomize 0-255 and add to the current segment
				if r.firstIP[i] >= tempIP {        // If the current segment value is greater than or equal to the previous segment value, the randomization is successful, and the loop can be exited
					break
				}
			}
		}
	}
}

func loadIPRanges() []*net.IPAddr {
	ranges := newIPRanges()
	if IPText != "" { // Get IP range data from parameters
		IPs := strings.Split(IPText, ",") // Split into an array using commas and loop through
		for _, IP := range IPs {
			IP = strings.TrimSpace(IP) // Remove leading and trailing whitespace characters (spaces, tabs, line breaks, etc.)
			if IP == "" {              // Skip empty ones (e.g., at the beginning, end, or consecutive ,,)
				continue
			}
			ranges.parseCIDR(IP) // Parse IP range, obtain IP, IP range, and subnet mask
			if isIPv4(IP) {      // Generate all IPv4 / IPv6 addresses to be tested (single/random/all)
				ranges.chooseIPv4()
			} else {
				ranges.chooseIPv6()
			}
		}
	} else { // Get IP range data from a file
		if IPFile == "" {
			IPFile = defaultInputFile
		}
		file, err := os.Open(IPFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() { // Loop through each line of the file
			line := strings.TrimSpace(scanner.Text()) // Remove leading and trailing whitespace characters (spaces, tabs, line breaks, etc.)
			if line == "" {                           // Skip empty lines
				continue
			}
			ranges.parseCIDR(line) // Parse IP range, obtain IP, IP range, and subnet mask
			if isIPv4(line) {      // Generate all IPv4 / IPv6 addresses to be tested (single/random/all)
				ranges.chooseIPv4()
			} else {
				ranges.chooseIPv6()
			}
		}
	}
	return ranges.ips
}
