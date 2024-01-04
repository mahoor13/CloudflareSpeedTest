package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/XIU2/CloudflareSpeedTest/task"
	"github.com/XIU2/CloudflareSpeedTest/utils"
)

var (
	version, versionNew string
)

func init() {
	var printVersion bool
	var help = `
CloudflareSpeedTest ` + version + `
Test the latency and speed of all IP addresses associated with Cloudflare CDN, and obtain the fastest IP (IPv4+IPv6)!
https://github.com/XIU2/CloudflareSpeedTest

Parameters:
    -n 200
        Number of latency testing threads; more threads lead to faster latency testing, avoid high values for low-performance devices (e.g., routers); (default 200, maximum 1000)
    -t 4
        Number of latency testing times per IP; (default 4 times)
    -dn 10
        Number of download speed tests; after latency testing and sorting, the number of IPs to perform download speed tests on; (default 10)
    -dt 10
        Download speed test time per IP; should not be too short; (default 10 seconds)
    -tp 443
        Specify the testing port; the port used for latency testing/download speed testing; (default 443)
    -url https://cf.xiu2.xyz/url
        Specify the testing address; the address used for latency testing (HTTPing)/download speed testing; the default address is not guaranteed to be available, it is recommended to use a custom one;

    -httping
        Switch testing mode; change latency testing mode to HTTP protocol, using the address specified by [-url] parameter; (default TCPing)
    -httping-code 200
        Effective status codes; HTTPing latency testing's valid HTTP status codes, only one; (default 200 301 302)
    -cfcolo HKG,KHH,NRT,LAX,SEA,SJC,FRA,MAD
        Match specified regions; region names are local airport codes, separated by commas, only available in HTTPing mode; (default all regions)

    -tl 200
        Average latency upper limit; only output IPs with average latency lower than the specified limit, various upper and lower limit conditions can be combined; (default 9999 ms)
    -tll 40
        Average latency lower limit; only output IPs with average latency higher than the specified limit; (default 0 ms)
    -tlr 0.2
        Packet loss rate upper limit; only output IPs with packet loss rate lower than or equal to the specified rate, range 0.00~1.00, 0 filters out any packet loss IPs; (default 1.00)
    -sl 5
        Download speed lower limit; only output IPs with download speed higher than the specified limit, testing stops when reaching the specified quantity [-dn]; (default 0.00 MB/s)

    -p 10
        Display number of results; directly display the specified number of results after testing, when 0, results are not displayed and the program exits directly; (default 10)
    -f ip.txt
        IP range data file; if the path contains spaces, please use quotes; supports other CDN IP ranges; (default ip.txt)
    -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
        Specify IP range data; directly specify the IP range data to be tested through parameters, separated by commas; (default empty)
    -o result.csv
        Write results to file; if the path contains spaces, please use quotes; when empty, do not write to file [-o ""]; (default result.csv)

    -dd
        Disable download speed testing; when disabled, the results are sorted by latency (default sorting by download speed); (default enabled)
    -allip
        Test all IPs; test each IP in the IP range (IPv4 only); (default randomly test one IP per /24 range)

    -v
        Print program version + check for updates
    -h
        Print help instructions
`
	var minDelay, maxDelay, downloadTime int
	var maxLossRate float64
	flag.IntVar(&task.Routines, "n", 200, "Number of latency testing threads")
	flag.IntVar(&task.PingTimes, "t", 4, "Number of latency testing times")
	flag.IntVar(&task.TestCount, "dn", 10, "Number of download speed tests")
	flag.IntVar(&downloadTime, "dt", 10, "Download speed test time")
	flag.IntVar(&task.TCPPort, "tp", 443, "Specify testing port")
	flag.StringVar(&task.URL, "url", "https://cf.xiu2.xyz/url", "Specify testing address")

	flag.BoolVar(&task.Httping, "httping", false, "Switch testing mode")
	flag.IntVar(&task.HttpingStatusCode, "httping-code", 0, "Effective status codes")
	flag.StringVar(&task.HttpingCFColo, "cfcolo", "", "Match specified regions")

	flag.IntVar(&maxDelay, "tl", 9999, "Average latency upper limit")
	flag.IntVar(&minDelay, "tll", 0, "Average latency lower limit")
	flag.Float64Var(&maxLossRate, "tlr", 1, "Packet loss rate upper limit")
	flag.Float64Var(&task.MinSpeed, "sl", 0, "Download speed lower limit")

	flag.IntVar(&utils.PrintNum, "p", 10, "Display number of results")
	flag.StringVar(&task.IPFile, "f", "ip.txt", "IP range data file")
	flag.StringVar(&task.IPText, "ip", "", "Specify IP range data")
	flag.StringVar(&utils.Output, "o", "result.csv", "Output results file")

	flag.BoolVar(&task.Disable, "dd", false, "Disable download speed testing")
	flag.BoolVar(&task.TestAll, "allip", false, "Test all IPs")

	flag.BoolVar(&printVersion, "v", false, "Print program version")
	flag.Usage = func() { fmt.Print(help) }
	flag.Parse()

	if task.MinSpeed > 0 && time.Duration(maxDelay)*time.Millisecond == utils.InputMaxDelay {
		fmt.Println("[Tip] When using the [-sl] parameter, it is recommended to use the [-tl] parameter to avoid continuous testing due to insufficient quantity [-dn]...")
	}
	utils.InputMaxDelay = time.Duration(maxDelay) * time.Millisecond
	utils.InputMinDelay = time.Duration(minDelay) * time.Millisecond
	utils.InputMaxLossRate = float32(maxLossRate)
	task.Timeout = time.Duration(downloadTime) * time.Second
	task.HttpingCFColomap = task.MapColoMap()

	if printVersion {
		println(version)
		fmt.Println("Checking for updates...")
		checkUpdate()
		if versionNew != "" {
			fmt.Printf("*** New version [%s] found! Please visit [https://github.com/XIU2/CloudflareSpeedTest] to update! ***", versionNew)
		} else {
			fmt.Println("Current version is the latest [" + version + "]!")
		}
		os.Exit(0)
	}
}

func main() {
	task.InitRandSeed() // Set random seed

	fmt.Printf("# XIU2/CloudflareSpeedTest %s \n\n", version)

	// Start latency testing + filter latency/packet loss
	pingData := task.NewPing().Run().FilterDelay().FilterLossRate()
	// Start download speed testing
	speedData := task.TestDownloadSpeed(pingData)
	utils.ExportCsv(speedData) // Output file
	speedData.Print()          // Print results

	if versionNew != "" {
		fmt.Printf("\n*** New version [%s] found! Please visit [https://github.com/XIU2/CloudflareSpeedTest] to update! ***\n", versionNew)
	}
	endPrint()
}

func endPrint() {
	if utils.NoPrintResult() {
		return
	}
	if runtime.GOOS == "windows" { // If it is a Windows system, press Enter or Ctrl+C to exit (to avoid closing directly after testing when run by double-clicking)
		fmt.Printf("Press Enter or Ctrl+C to exit.")
		fmt.Scanln()
	}
}

// Check for updates
func checkUpdate() {
	timeout := 10 * time.Second
	client := http.Client{Timeout: timeout}
	res, err := client.Get("https://api.xiu2.xyz/ver/cloudflarespeedtest.txt")
	if err != nil {
		return
	}
	// Read resource data body: []byte
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	// Close resource stream
	defer res.Body.Close()
	if string(body) != version {
		versionNew = string(body)
	}
}
