# XIU2/CloudflareSpeedTest

[![Go Version](https://img.shields.io/github/go-mod/go-version/XIU2/CloudflareSpeedTest.svg?style=flat-square&label=Go&color=00ADD8&logo=go)](https://github.com/XIU2/CloudflareSpeedTest/)
[![Release Version](https://img.shields.io/github/v/release/XIU2/CloudflareSpeedTest.svg?style=flat-square&label=Release&color=00ADD8&logo=github)](https://github.com/XIU2/CloudflareSpeedTest/releases/latest)
[![GitHub license](https://img.shields.io/github/license/XIU2/CloudflareSpeedTest.svg?style=flat-square&label=License&color=00ADD8&logo=github)](https://github.com/XIU2/CloudflareSpeedTest/)
[![GitHub Star](https://img.shields.io/github/stars/XIU2/CloudflareSpeedTest.svg?style=flat-square&label=Star&color=00ADD8&logo=github)](https://github.com/XIU2/CloudflareSpeedTest/)
[![GitHub Fork](https://img.shields.io/github/forks/XIU2/CloudflareSpeedTest.svg?style=flat-square&label=Fork&color=00ADD8&logo=github)](https://github.com/XIU2/CloudflareSpeedTest/)

Many foreign websites use Cloudflare CDN, but the IP addresses assigned to visitors in mainland China are not friendly (high latency, high packet loss, slow speed).  
Although Cloudflare has publicly disclosed all [IP ranges](https://www.cloudflare.com/zh-cn/ips/), finding the suitable one among so many IPs can be challenging. Thus, this software was created.

**"Choose Your Preferred IP" to test Cloudflare CDN latency and speed, and get the fastest IP (IPv4+IPv6)! If you find it useful, give it a ⭐!**

> _Check out my other open-source projects: [**TrackersList.com** - A collection of popular BT Tracker lists! Effectively improves BT download speed~](https://github.com/XIU2/TrackersListCollection) <img src="https://img.shields.io/github/stars/XIU2/TrackersListCollection.svg?style=flat-square&label=Star&color=4285dd&logo=github" height="16px" />_  
> _[**UserScript** - 🐵 Github high-speed downloads, enhanced Zhihu, automatic seamless page flipping, eye protection mode, and more than a dozen **Monkey scripts**~](https://github.com/XIU2/UserScript) <img src="https://img.shields.io/github/stars/XIU2/UserScript.svg?style=flat-square&label=Star&color=4285dd&logo=github" height="16px" />_  
> _[**SNIProxy** - 🧷 Simple SNI Proxy for personal use (supports all platforms, systems, front-end proxy, simple configuration, etc.)~](https://github.com/XIU2/SNIProxy) <img src="https://img.shields.io/github/stars/XIU2/SNIProxy.svg?style=flat-square&label=Star&color=4285dd&logo=github" height="16px" />_  

> This project also supports latency testing for **other CDN / website IPs** (such as: [CloudFront](https://github.com/XIU2/CloudflareSpeedTest/discussions/304), [Gcore](https://github.com/XIU2/CloudflareSpeedTest/discussions/303) CDN), but you need to find the addresses yourself for downloading speed tests.

> For users with **proxies through Cloudflare CDN**, be aware that this should be considered as an **alternative solution**, not the **only solution**. Avoid excessive reliance on [#382](https://github.com/XIU2/CloudflareSpeedTest/discussions/382) [#383](https://github.com/XIU2/CloudflareSpeedTest/discussions/383)

****
## \# Quick Usage

### Download and Run

1. Download the compiled executable file ([Github Releases](https://github.com/XIU2/CloudflareSpeedTest/releases) / [Lanzou Cloud](https://pan.lanzouj.com/b0742hkxe)) and unzip it.  
2. Double-click the `CloudflareST.exe` file (for Windows systems) and wait for the speed test to complete...

<details>
<summary><code><strong>「 Click to view the usage example on Linux systems 」</strong></code></summary>

****

The following commands are just examples. Please check [**Releases**](https://github.com/XIU2/CloudflareSpeedTest/releases) for version numbers and file names.

``` yaml
# If it is the first time, it is recommended to create a new folder (skip this step for subsequent updates)
mkdir CloudflareST

# Enter the folder (for subsequent updates, just repeat the following download and unzip commands from here)
cd CloudflareST

# Download the CloudflareST archive (replace [Version] and [Filename] in the URL according to your needs)
wget -N https://github.com/XIU2/CloudflareSpeedTest/releases/download/v2.2.4/CloudflareST_linux_amd64.tar.gz
# If you are downloading on a server in China, use one of the following mirrors for acceleration:
# wget -N https://download.fgit.ml/XIU2/CloudflareSpeedTest/releases/download/v2.2.4/CloudflareST_linux_amd64.tar.gz
# wget -N https://download.fgit.gq/XIU2/CloudflareSpeedTest/releases/download/v2.2.4/CloudflareST_linux_amd64.tar.gz
# wget -N https://ghproxy.com/https://github.com/XIU2/CloudflareSpeedTest/releases/download/v2.2.4/CloudflareST_linux_amd64.tar.gz
# If the download fails, try removing the -N parameter (if updating, remember to delete the old archive rm CloudflareST_linux_amd64.tar.gz)

# Unzip (no need to delete old files, they will be overwritten directly; replace [Filename] according to your needs)
tar -zxf CloudflareST_linux_amd64.tar.gz

# Grant execute permissions
chmod +x CloudflareST

# Run (without parameters)
./CloudflareST

# Run (with example parameters)
./CloudflareST -dd -tll 90
```
> If the average latency is very low (e.g., 0.xx), it indicates that CloudflareST is using a proxy during the speed test. Please disable the proxy software before testing the speed.
> If running on a router, it is recommended to disable any proxies in the router (or exclude them), as speed test results may be inaccurate/unusable.

</details>

# XIU2/CloudflareSpeedTest vX.X.X

**Quick Guide to Independently Running CloudflareST Speed Test on Mobile Phones: [Android](https://github.com/XIU2/CloudflareSpeedTest/discussions/61), [Android APP](https://github.com/xianshenglu/cloudflare-ip-tester-app), [IOS](https://github.com/XIU2/CloudflareSpeedTest/discussions/321)**

**Note! This software is only applicable to websites and does not support selecting Cloudflare WARP preferred IP. See details: [#392](https://github.com/XIU2/CloudflareSpeedTest/discussions/392)**

### Example Results

After the speed test, it will display the **top 10 fastest IPs** by default, for example:

```bash
IP Address        Sent    Received    Loss Rate    Average Latency    Download Speed (MB/s)
104.27.200.69     4       4           0.00         146.23             28.64
172.67.60.78      4       4           0.00         139.82             15.02
104.25.140.153    4       4           0.00         146.49             14.90
104.27.192.65     4       4           0.00         140.28             14.07
172.67.62.214     4       4           0.00         139.29             12.71
104.27.207.5      4       4           0.00         145.92             11.95
172.67.54.193     4       4           0.00         146.71             11.55
104.22.66.8       4       4           0.00         147.42             11.11
104.27.197.63     4       4           0.00         131.29             10.26
172.67.58.91      4       4           0.00         140.19             9.14
...
```

The first line of the speed test results is the **fastest IP with both the highest download speed and lowest average latency**!

The complete results are saved in the `result.csv` file in the current directory. Open it with **Notepad/Spreadsheet software**, and the format is as follows:

```
IP Address, Sent, Received, Loss Rate, Average Latency, Download Speed (MB/s)
104.27.200.69,4,4,0.00,146.23,28.64
```

> Customize the complete results according to your needs, or explore advanced usage with **specified filtering conditions**!

### Advanced Usage

To achieve more comprehensive and customized speed test results, you can use custom parameters.

```css
C:\>CloudflareST.exe -h

CloudflareSpeedTest vX.X.X
Test the latency and speed of all IPs of Cloudflare CDN, get the fastest IP (IPv4+IPv6)!
https://github.com/XIU2/CloudflareSpeedTest

Parameters:
    -n 200
        Latency test threads; the more threads, the faster the latency test. Devices with weak performance (such as routers) should not set it too high; (default 200, maximum 1000)
    -t 4
        Number of latency test times; the number of times a single IP's latency is tested; (default 4 times)
    -dn 10
        Number of download speed tests; after latency test and sorting, the number of IPs to perform download speed test, starting from the lowest latency; (default 10)
    -dt 10
        Download speed test time; maximum time for a single IP's download speed test; it shouldn't be too short; (default 10 seconds)
    -tp 443
        Specify the test port; the port used during latency test/download speed test; (default port 443)
    -url https://cf.xiu2.xyz/url
        Specify the test address; the address used during latency test (HTTPing)/download speed test. The default address is not guaranteed to be available; it is recommended to use a self-built address;

    -httping
        Switch test mode; change the latency test mode to HTTP protocol, and use the test address specified by [-url] parameter; (default TCPing)
        Note: HTTPing is essentially a kind of network scanning behavior. If you run it on a server, you need to reduce concurrency (-n), otherwise, it may be temporarily suspended by some strict service providers. If you encounter a situation where the number of available IPs is normal during the first HTTPing test, but it decreases or even directly becomes 0 in subsequent tests, but it recovers after a while, it may also be recognized as a temporary restriction mechanism triggered by the carrier or Cloudflare CDN due to network scanning. Therefore, it will recover after a while, and it is recommended to reduce the concurrency (-n) to reduce the occurrence of this situation.
    -httping-code 200
        Effective status code; the effective HTTP status code returned when HTTPing tests the latency. Only one code is allowed; (default 200 301 302)
    -cfcolo HKG,KHH,NRT,LAX,SEA,SJC,FRA,MAD
        Match specified regions; the region name is the local airport code, separated by English commas, supports lowercase, supports Cloudflare, AWS CloudFront, only available in HTTPing mode; (default all regions)

    -tl 200
        Upper limit of average latency; only output IPs with average latency lower than the specified limit, various upper and lower limit conditions can be used together; (default 9999 ms)
    -tll 40
        Lower limit of average latency; only output IPs with average latency higher than the specified limit; (default 0 ms)
    -tlr 0.2
        Upper limit of packet loss rate; only output IPs with packet loss rate lower than or equal to the specified rate, range 0.00~1.00, 0 filters out any IP with packet loss; (default 1.00)
    -sl 5
        Lower limit of download speed; only output IPs with download speed higher than the specified speed, stop the speed test only when the specified number [-dn] is reached; (default 0.00 MB/s)

    -p 10
        Display result quantity; directly display the specified quantity of results after the speed test, set to 0 to exit without displaying results; (default 10)
    -f ip.txt
        IP range data file; if the path contains spaces, please add quotes; supports other CDN IP ranges; (default ip.txt)
    -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
        Specify IP range data; directly specify the IP range data to be tested through parameters, separated by English commas; (default empty)
    -o result.csv
        Write result file; if

 the path contains spaces, please add quotes; when the value is empty, do not write to the file [-o ""]; (default result.csv)

    -dd
        Disable download speed test; after disabling, the speed test results will be sorted by latency (default sorted by download speed); (default enabled)
    -allip
        Test all IPs; test each IP in the IP range (IPv4 only); (default randomly test one IP in each /24 range)

    -v
        Print program version + check for version updates
    -h
        Print help instructions
```

### Interface Explanation

To avoid misunderstanding the **output content during the speed test process (usable, queue numbers, interruption during half of the download speed test?)**, I specifically explain it.

<details>
<summary><code><strong>「 Click to expand to see the content 」</strong></code></summary>

****

> This example has added common parameters, namely: `-ttl 40 -tl 150 -sl 1 -dn 5`, and the final output results are as follows:

```bash
# XIU2/CloudflareSpeedTest vX.X.X

Start latency test (mode: TCP, port: 443, range: 40 ~ 150 ms, loss: 1.00)
321 / 321 [----------------------------------------------------------------------------------] Available: 30
Start download speed test (lower limit: 1.00 MB/s, quantity: 5, queue: 10)
3 / 5 [---------------------------------------------------------↗---------------------------]
IP Address        Sent    Received    Loss Rate    Average Latency    Download Speed (MB/s)
XXX.XXX.XXX.XXX   4       4           0.00         83.32              3.66
XXX.XXX.XXX.XXX   4       4           0.00         107.81             2.49
XXX.XXX.XXX.XXX   4       3           0.25         149.59             1.04

Complete speed test results have been written to the result.csv file, you can use Notepad/Spreadsheet software to view it.
Press Enter or Ctrl+C to exit.
```
# For those who have just started using CloudflareST, you might be confused: **There were 30 usable IPs in the latency test, but why are there only 3 left at the end?**
What does the queue in the download speed test mean? Do I have to wait in line for the download speed test?

CloudflareST will perform latency tests first, and during this process, the progress bar on the right will real-time display the number of available IPs (`Available: 30`). However, note that this available quantity refers to the **number of IPs that passed the test without timing out**, unrelated to latency upper and lower limits, or packet loss conditions. After the latency test is completed, due to specifying latency upper and lower limits and packet loss conditions, only `10` IPs remain (waiting for download speed tests, i.e., `Queue: 10`).

So, in the example above, after the latency test of `321` IPs is completed, only `30` IPs passed the test without timing out. Then, based on the latency upper and lower limits (`40 ~ 150 ms`) and packet loss upper limit conditions, only `10` IPs that meet the requirements remain. If you have disabled download speed test (`-dd`), then these `10` IPs will be directly output. However, in this example, the download speed test is not disabled, so the software will continue to perform download speed tests on these `10` IPs (`Queue: 10`).

> Because download speed tests are done one by one in a single-threaded manner, the number of IPs waiting for download speed determines the `queue`.

****

> You may have noticed: **You specified to find 5 IPs that meet the download speed conditions, but why did it "interrupt" at 3?**

In the download speed test progress bar, `3 / 5`, the former indicates that `3` IPs meeting the download speed lower limit conditions (i.e., download speed above `1 MB/s`) have been found. The latter `5` indicates that you requested to find `5` IPs meeting the download speed lower limit conditions (`-dn 5`).

> Additionally, it is worth noting that if you set `-dn` greater than the download speed test queue, for example, after the latency test, only `4` IPs remain, then the numbers in the download speed test progress bar will both be `4`, not the `5` specified by `-dn`.

After the software finishes testing these `10` IPs, it found only `3` IPs with download speeds above `1 MB/s`, and the remaining `7` IPs are "not qualified."

So, it's not "it interrupts before reaching 5 every time," but rather all IPs have completed the download speed test, but only `3` IPs meeting the conditions were found.

****

If you don't want to encounter a situation where there are not many IPs meeting the conditions after all the tests, you can **lower the download speed upper limit parameter `-sl`**, or remove it.

Because as long as the `-sl` parameter is specified, the speed test will continue until the specified number of `-dn` (default 10) is reached or all IPs have completed the test. Removing `-sl` and adding `-dn 20` will only test the latency of the top 20 IPs with the lowest latency and stop after testing, saving time.

****

In addition, if all IPs in the queue have completed the test but none of them meet the download speed conditions, then it will **directly output the download speed results of all IPs in the queue**. This way, you can see the download speeds of these IPs, and then **try lowering `-sl` appropriately**.

Similarly, for latency tests, the values of `Available: 30` and `Queue: 10` can tell you whether the latency conditions you set are too strict for you. If there are plenty of available IPs, but after filtering conditions, only 2 or 3 remain, then it's clear that you need to **lower the expected latency/packet loss conditions**.

These two mechanisms, one tells you about **latency and packet loss conditions**, and the other tells you about **download speed conditions**.

</details>

****

### Usage Examples

On Windows, specifying parameters requires running in CMD or adding parameters to the shortcut target.

> **Note**: All parameters have **default values**, and parameters using default values can be omitted (**choose as needed**).  
> **Tip**: In Windows **PowerShell**, just change `CloudflareST.exe` in the command to `.\CloudflareST.exe`.  
> **Tip**: For Linux systems, just change `CloudflareST.exe` in the command to `./CloudflareST`.

****

#### \# Run CloudflareST in CMD with Parameters

For those not familiar with the command-line program, you might not know how to run it with parameters. Let me explain briefly.

<details>
<summary><code><strong>「 Click to expand to see the content 」</strong></code></summary>

****

Many people opening CMD to run CloudflareST with an **absolute path** will encounter errors. This is because the default `-f ip.txt` parameter is a relative

 path and requires specifying the absolute path of ip.txt. However, this is too troublesome, so it is recommended to enter the CloudflareST program directory and run it with a **relative path**:

**Method 1**:
1. Open the directory where CloudflareST is located.
2. Right-click on a blank space, press <kbd>Shift + Right-click</kbd> to display the context menu.
3. Choose **\[Open command window here\]** to open the CMD window, which is now located in the current directory.
4. Enter the command with parameters, such as `CloudflareST.exe -tll 50 -tl 200` to run.

**Method 2**:
1. Open the directory where CloudflareST is located.
2. Directly select and enter `cmd` in the folder address bar, then press Enter to open the CMD window, which is now located in the current directory.
3. Enter the command with parameters, such as `CloudflareST.exe -tll 50 -tl 200` to run.

> Of course, you can also randomly open a CMD window and then enter something like `cd /d "D:\Program Files\CloudflareST"` to enter the program directory.

> **Tip**: If you are using **PowerShell**, just change `CloudflareST.exe` in the command to `.\CloudflareST.exe`.

</details>

****

#### \# Run CloudflareST with Parameters using Windows Shortcut

If you don't often modify the run parameters (such as usually directly double-clicking to run), it is recommended to use a shortcut, which is more convenient.

<details>
<summary><code><strong>「 Click to expand to see the content 」</strong></code></summary>

Right-click the `CloudflareST.exe` file - **\[Create Shortcut\]**, then right-click on the shortcut - **\[Properties\]**, and modify its **Target**:

```bash
# If you don't want to output result files, add -o " " (quotes contain a space, without a space, this parameter will be omitted).
D:\ABC\CloudflareST\CloudflareST.exe -n 500 -t 4 -dn 20 -dt 5 -o " "

# If the file path contains quotes, place the startup parameters outside the quotes, and remember there is a space between quotes and -.
"D:\Program Files\CloudflareST\CloudflareST.exe" -n 500 -t 4 -dn 20 -dt 5 -o " "

# Note! The shortcut - Start in location cannot be empty; otherwise, it will not find the ip.txt file due to the absolute path.
```

</details>

****

#### \# IPv4/IPv6

<details>
<summary><code><strong>「 Click to expand to see the content 」</strong></code></summary>

****
```bash
# Specify the built-in IPv4 data file to test these IPv4 addresses (the -f default value is ip.txt, so this parameter can be omitted).
CloudflareST.exe -f ip.txt

# Specify the built-in IPv6 data file to test these IPv6 addresses
# Additionally, starting from version 2.1.0, IPv4+IPv6 mixed speed testing is supported, and the -ipv6 parameter has been removed. Therefore, a file can contain both IPv4+IPv6 addresses.
CloudflareST.exe -f ipv6.txt

# You can also specify the IP to be tested directly through parameters.
CloudflareST.exe -ip 1.1.1.1,2606:4700::/32

> When testing IPv6, you may notice that the number of tests varies each time. Understand the reason: [#120](https://github.com/XIU2/CloudflareSpeedTest/issues/120)
> Because there are too many IPv6 addresses (in the order of hundreds of millions), and the vast majority of IP ranges are not enabled, I only scanned a part of the available IPv6 ranges and wrote them to the `ipv6.txt` file. If you are interested, you can scan and add/delete them yourself. ASN data source is from: [bgp.he.net](https://bgp.he.net/AS13335#_prefixes6)
```

</details>

****

#### \# HTTPing

<details>
<summary><code><strong>「 Click to expand to see the content 」</strong></code></summary>

****

There are currently two latency test modes: **TCP protocol** and **HTTP protocol**.
The TCP protocol takes less time and consumes fewer resources, with a timeout of 1 second; this is the default mode.
The HTTP protocol is suitable for quickly testing if a certain domain points to a certain IP, with a timeout of 2 seconds.
For the same IP, the latency obtained by each protocol generally follows: **ICMP < TCP < HTTP**, with the rightmost being more sensitive to network fluctuations such as packet loss.

> Note: HTTPing is essentially a kind of **network scanning** behavior. If you run it on a server, you need to **reduce concurrency** (`-n`), or you may be temporarily suspended by some strict providers. If you encounter a situation where the first HTTPing test has a normal number of available IPs, but the subsequent tests have fewer and even become 0, but it recovers after a while, it may be due to triggering a temporary restriction mechanism by the ISP or Cloudflare CDN for **network scanning**. It will recover after a while. It is recommended to **reduce concurrency** (`-n`) to reduce the occurrence of this situation.

> Also, this software's HTTPing only obtains **response headers**, and does not retrieve the content of the body (i.e., the URL file size does not affect HTTPing tests. However, if you still want to perform download speed tests, you will need a large file).

```bash
# Just add the -httping parameter to switch to HTTP protocol latency test mode.
CloudflareST.exe -httping

# The software will determine the availability based on the effective HTTP status codes returned when accessing the webpage (of course, timeout is also considered). By default, responses with 200, 301, and 302 HTTP status codes are considered valid. You can manually specify the HTTP status code considered valid, but only one can be specified (you need to determine in advance which status code the test address will return under normal circumstances).
CloudflareST.exe -httping -httping-code 200

# Use the -url parameter to specify the HTTPing test address (it can be any webpage URL, not limited to a specific file address).
CloudflareST.exe -httping -url https://cf.xiu2.xyz/url
# If you want to HTTPing test other websites/CDNs, then specify an address that uses that website/CDN (because the default address is Cloudflare's and can only be used to test Cloudflare's IP).
```

</details>

****

#### \# Match specified region (colo airport three-letter code)

<details>
<summary><code><strong>「 Click to expand to see the content 」</strong></code></summary>

****

```bash
# This feature supports Cloudflare CDN and AWS CloudFront CDN, and the three-letter codes for these two CDNs are universal.
# Note: If you want to use it to filter AWS CloudFront CDN regions, you need to specify an address using that CDN through the -url parameter (because the default address is Cloudflare's).

# After specifying the region name, the results obtained after the latency test will all be IPs of the specified region (you can continue to perform download speed tests).
# The node region name is the local airport three-letter code, and multiple names can be specified, separated by commas. Starting from version 2.2.3, lowercase is supported.

CloudflareST.exe -cfcolo HKG,KHH,NRT,LAX,SEA,SJC,FRA,MAD

# Note that this parameter is only available in HTTPing latency test mode (because it needs to access the webpage to obtain it).
```

> The two CDN airport three-letter codes are universal, so you can see the names of each region here: https://www.cloudflarestatus.com/
</details>

****

#### \# Relative/Absolute File Paths

<details>
<summary><code><strong>「 Click to expand to see the content 」</strong></code></summary>

****

```bash
# Specify the IPv4 data file, do not display the result directly exit, and output the result to the file (-p value is 0).
CloudflareST.exe -f 1.txt -p 0 -dd

# Specify the IPv4 data file, do not output the result to the file, and directly display the result (-p value is 10, -o value is empty but quotes cannot be less).
CloudflareST.exe -f 2.txt -o "" -p 10 -dd

# Specify the IPv4 data file and output the result

 to the file (relative path, i.e., in the current directory, if there are spaces, please add quotes).
CloudflareST.exe -f 3.txt -o result.txt -dd

# Specify the IPv4 data file and output the result to the file (relative path, i.e., in the current directory in the abc folder, if there are spaces, please add quotes)
# Linux (Inside the abc folder in the CloudflareST program directory)
./CloudflareST -f abc/3.txt -o abc/result.txt -dd

# Windows (Note the backslash)
CloudflareST.exe -f abc\3.txt -o abc\result.txt -dd

# Specify the IPv4 data file and output the result to the file (absolute path, i.e., under the C:\abc\ directory, if there are spaces, please add quotes)
# Linux (Under the /abc/ directory)
./CloudflareST -f /abc/4.txt -o /abc/result.csv -dd

# Windows (Note the backslash)
CloudflareST.exe -f C:\abc\4.txt -o C:\abc\result.csv -dd

# If you want to run CloudflareST with an absolute path, then the file name in the -f / -o parameter must also be an absolute path, otherwise an error will be reported file not found!
# Linux (Under the /abc/ directory)
/abc/CloudflareST -f /abc/4.txt -o /abc/result.csv -dd

# Windows (Note the backslash)
C:\abc\CloudflareST.exe -f C:\abc\4.txt -o C:\abc\result.csv -dd
```
</details>

#### \# Test Other Ports

<details>
<summary><code><strong>「 Click to expand to see the content 」</strong></code></summary>

****

```bash
# If you want to test ports other than the default 443, you need to specify them using the -tp parameter (this parameter will affect the port used for latency testing/download speed testing).

# If you want to latency test port 80 + download speed test (if -dd disables download speed testing, it is not needed), you also need to specify an http:// protocol download speed test address (and the address will not be forcibly redirected to HTTPS because that will be port 443).
CloudflareST.exe -tp 80 -url http://cdn.cloudflare.steamstatic.com/steam/apps/5952/movie_max.webm

# If it is a non-80 443 port, you need to make sure that the download speed test address you are using supports access through that non-standard port.
```

</details>

****

#### \# Custom Test Address

<details>
<summary><code><strong>「 Click to expand to see the content 」</strong></code></summary>

****

```bash
# This parameter is applicable to download speed testing and latency testing using the HTTP protocol. For the latter, the address can be any webpage URL (not limited to a specific file address).

# Address requirements: Direct download, file size over 200MB, and using Cloudflare CDN.
CloudflareST.exe -url https://cf.xiu2.xyz/url

# Note: If the test address is HTTP protocol (the address cannot be forcibly redirected to HTTPS), remember to add -tp 80 (this parameter will affect the port used for latency testing/download speed testing). If it is a non-80 443 port, you need to make sure that the download speed test address supports access through that port.
CloudflareST.exe -tp 80 -url http://cdn.cloudflare.steamstatic.com/steam/apps/5952/movie_max.webm
```

</details>

****

#### \# Custom Test Conditions (Specify the Target Range for Latency/Packet Loss/Download Speed)

<details>
<summary><code><strong>「 Click to expand to see the content 」</strong></code></summary>

****

> Note: The "available quantity" on the right of the latency test progress bar refers only to the number of IPs that did not time out during the latency test and is unrelated to the latency upper and lower limit conditions.

- Specify only the **[Average Latency Upper Limit]** condition

```bash
# Average latency upper limit: 200 ms, download speed lower limit: 0 MB/s
# That is, find IPs with an average latency below 200 ms, and then perform 10 download speed tests in ascending order of latency.
CloudflareST.exe -tl 200
```

> If **no IP meets the latency** condition, nothing will be output.

****

- Specify only the **[Average Latency Upper Limit]** condition and **only latency testing, no download speed testing**

```bash
# Average latency upper limit: 200 ms, download speed lower limit: 0 MB/s, quantity: unknown
# That is, only output IPs with latency below 200 ms and no longer perform download speed tests (because no longer performing download speed tests, the -dn parameter is invalid).
CloudflareST.exe -tl 200 -dd
```

- Specify only the **[Packet Loss Upper Limit]** condition

```bash
# Packet loss upper limit: 0.25
# That is, find IPs with a packet loss rate less than or equal to 0.25, with a range of 0.00 to 1.00. If -tlr 0 is specified, it means to filter out any IPs with packet loss.
CloudflareST.exe -tlr 0.25
```

****

- Specify only the **[Download Speed Lower Limit]** condition

```bash
# Average latency upper limit: 9999 ms, download speed lower limit: 5 MB/s, quantity: 10 (optional)
# That is, 10 IPs with an average latency below 9999 ms and a download speed above 5 MB/s must be found to stop the speed test.
CloudflareST.exe -sl 5 -dn 10
```

> If **no IP meets the speed** condition, it will **ignore the condition and output all IP speed test results** (convenient for adjusting conditions for the next test).

> If you **do not specify the average latency upper limit**, and the number of IPs that meet the condition is not reached, it will continue to test the speed indefinitely.
> So, it is recommended to **simultaneously specify [Download Speed Lower Limit] + [Average Latency Upper Limit]**. This way, if the test is not enough to meet the latency upper limit, it will stop the speed test.

****

- Specify both the **[Average Latency Upper Limit] + [Download Speed Lower Limit]** conditions

```bash
# Both the average latency upper limit and download speed lower limit support decimals (such as -sl 0.5).
# Average latency upper limit: 200 ms, download speed lower limit: 5.6 MB/s, quantity: 10 (optional)
# That is, 10 IPs with an average latency below 200 ms and a download speed above 5.6 MB/s must be found to stop the speed test.
CloudflareST.exe -tl 200 -sl 5.6 -dn 10
```

> If **no IP meets the latency** condition, nothing will be output.
> If **no IP meets the speed** condition, it will ignore the condition and output all IP speed test results (convenient for adjusting conditions for the next test).
> So, it is recommended to run the test once without specifying conditions to see the approximate range of average latency and download speed to avoid specifying conditions that are too low or too high!

> Because the IP ranges published by Cloudflare are **origin IPs + anycast IPs**, and **origin IPs** cannot be used, so the download speed is 0.00.
> You can add -sl 0.01 (download speed lower limit) at runtime to filter out **origin IPs** (results with download speeds below 0.01MB/s).
</details>

****

#### \# Test a Single or Multiple IPs Separately

<details>
<summary><code><strong>「 Click to expand to see the content 」</strong></code></summary>

****

**Method One**:
Specify the IP range data to be tested directly through parameters.
```bash
# Enter the directory where CloudflareST is located, and then run:
# Windows system (run in CMD)
CloudflareST.exe -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32

# Linux system
./CloudflareST -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
```

****

**Method Two**:
Alternatively, write these IPs in any text file in the following format, for example, `1.txt`

```
1.1.1.

1
1.1.1.200
1.0.0.1/24
2606:4700::/32
```

> For a single IP, you can omit the `/32` subnet mask (i.e., `1.1.1.1` is equivalent to `1.1.1.1/32`).
> The subnet mask `/24` refers to the last segment of this IP, that is, `1.0.0.1~1.0.0.255`.

Then run CloudflareST with the startup parameter `-f 1.txt` to specify the IP range data file.

```bash
# Enter the directory where CloudflareST is located, and then run:
# Windows system (run in CMD)
CloudflareST.exe -f 1.txt

# Linux system
./CloudflareST -f 1.txt

# For IP ranges like 1.0.0.1/24, only the last segment will be random (1.0.0.1~255). If you want to test all IPs in this range, please add the -allip parameter.
```

</details>

****

#### \# One-stop Acceleration for All Websites Using Cloudflare CDN (No Need to Add Domains to Hosts One by One)

I've mentioned before that the purpose of developing this software project is to accelerate access to websites using Cloudflare CDN through **modifying the Hosts file**.

However, as mentioned in [**#8**](https://github.com/XIU2/CloudflareSpeedTest/issues/8), adding domains to Hosts one by one is **too troublesome**, so I found a **one-stop solution**! You can check this [**Still adding Hosts one by one? Perfect local acceleration method for all websites using Cloudflare CDN is here!**](https://github.com/XIU2/CloudflareSpeedTest/discussions/71) and another [tutorial on modifying domain resolution IP to a custom IP using local DNS service](https://github.com/XIU2/CloudflareSpeedTest/discussions/317).

****

#### \# Automatically Update Hosts

Considering that many people need to replace the IP in the Hosts file after obtaining the fastest Cloudflare CDN IP.

You can check this [**Issues**](https://github.com/XIU2/CloudflareSpeedTest/discussions/312) to get **Windows/Linux automatic update Hosts script**!

****

## Issue Feedback

If you encounter any issues, you can first check [**Issues**](https://github.com/XIU2/CloudflareSpeedTest/issues) and [Discussions](https://github.com/XIU2/CloudflareSpeedTest/discussions) to see if others have asked similar questions (remember to check [**Closed**](https://github.com/XIU2/CloudflareSpeedTest/issues?q=is%3Aissue+is%3Aclosed)).  
If you don't find a similar issue, please open a new [**Issues**](https://github.com/XIU2/CloudflareSpeedTest/issues/new) to let me know!

> **Note**! For things unrelated to `feedback issues, feature suggestions`, please go to the project's internal forum for discussion (the `💬 Discussions` above).

****

## Support and Appreciation

![WeChat Appreciation](https://github.com/XIU2/XIU2/blob/master/img/zs-01.png)![Alipay Appreciation](https://github.com/XIU2/XIU2/blob/master/img/zs-02.png)

****

## Derivative Projects

- _https://github.com/xianshenglu/cloudflare-ip-tester-app_  
_**CloudflareST Android APP [#202](https://github.com/XIU2/CloudflareSpeedTest/discussions/320)**_

- _https://github.com/mingxiaoyu/luci-app-cloudflarespeedtest_  
_**CloudflareST OpenWrt Router Plugin Version [#174](https://github.com/XIU2/CloudflareSpeedTest/discussions/319)**_

- _https://github.com/immortalwrt-collections/openwrt-cdnspeedtest_  
_**CloudflareST OpenWrt Native Compilation Version [#64](https://github.com/XIU2/CloudflareSpeedTest/discussions/64)**_

- _https://github.com/hoseinnikkhah/CloudflareSpeedTest-English_  
_**English language version of CloudflareST (Text language differences only) [#64](https://github.com/XIU2/CloudflareSpeedTest/issues/68)**_

> _Here only collects some CloudflareST related derivative projects that have been promoted in this project. If there are omissions, please let me know~_

****

## Acknowledgments

- _https://github.com/Spedoske/CloudflareScanner_

> _Because this project has not been updated for a long time, and I have many functional requirements, so I temporarily learned Go language and started (novice)..._  
> _This software is made based on that project, but has **added a lot of features and fixed bugs**, and actively added and optimized features based on user feedback (leisure)..._

****

## Manual Compilation

<details>
<summary><code><strong>「 Click to expand to see the content 」</strong></code></summary>

****

For convenience, I write the version number into the code's `version` variable during compilation. Therefore, when you manually compile, you need to use the `-ldflags` parameter after the `go build` command to specify the version number, like this:

```bash
go build -ldflags "-s -w -X main.version=v2.3.3"
# Run this command in the CloudflareSpeedTest directory via the command line (e.g., CMD, Batch script), and you can compile a binary program that can run in an environment with the same system, bit, and architecture as the current device (Go will automatically detect your system bit and architecture) and the version number is v2.3.3.
```

If you want to compile for **other systems, architectures, and bits** on a Windows 64-bit system, you need to specify the **GOOS** and **GOARCH** variables.

For example, to compile a binary program suitable for **Linux system amd architecture 64-bit** on a Windows system:

```bat
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-s -w -X main.version=v2.3.3"
```

For example, to compile a binary program suitable for **Windows system amd architecture 32-bit** on a Linux system:

```bash
GOOS=windows
GOARCH=386
go build -ldflags "-s -w -X main.version=v2.3.3"
```

> You can run `go tool dist list` to see which combinations are supported by the current Go version.

****

Of course, for convenience of batch compilation, I will specifically set a variable for the version number, and subsequent compilations can directly call this version number variable.  
At the same time, for batch compilation, you need to separate them into different folders (or use different file names), and you need to add the `-o` parameter to specify.

```bat
:: For Windows system:
SET version=v2.3.3
SET GOOS=linux
SET GOARCH=amd64
go build -o Releases\CloudflareST_linux_amd64\CloudflareST -ldflags "-s -w -X main.version=%version%"
```

```bash
# For Linux system:
version=v2.3.3
GOOS=windows
GOARCH=386
go build -o Releases/CloudflareST_windows_386/CloudflareST.exe -ldflags "-s -w -X main.version=${version}"
```

</details>

****

## License

The GPL-3.0 License.