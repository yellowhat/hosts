package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var urls = []string{
	// From firebog.net
	// - Suspicious Lists
	"https://raw.githubusercontent.com/PolishFiltersTeam/KADhosts/master/KADhosts.txt",
	"https://raw.githubusercontent.com/FadeMind/hosts.extras/master/add.Spam/hosts",
	"https://v.firebog.net/hosts/static/w3kbl.txt",
	"https://raw.githubusercontent.com/matomo-org/referrer-spam-blacklist/master/spammers.txt",
	"https://someonewhocares.org/hosts/zero/hosts",
	"https://raw.githubusercontent.com/RooneyMcNibNug/pihole-stuff/master/SNAFU.txt",
	// - Advertising Lists
	"https://adaway.org/hosts.txt",
	"https://v.firebog.net/hosts/AdguardDNS.txt",
	"https://v.firebog.net/hosts/Admiral.txt",
	"https://raw.githubusercontent.com/anudeepND/blacklist/master/adservers.txt",
	"https://v.firebog.net/hosts/Easylist.txt",
	"https://pgl.yoyo.org/adservers/serverlist.php?hostformat=hosts&showintro=0&mimetype=plaintext",
	"https://raw.githubusercontent.com/FadeMind/hosts.extras/master/UncheckyAds/hosts",
	"https://raw.githubusercontent.com/bigdargon/hostsVN/master/hosts",
	// - Tracking & Telemetry Lists
	"https://v.firebog.net/hosts/Easyprivacy.txt",
	"https://v.firebog.net/hosts/Prigent-Ads.txt",
	"https://raw.githubusercontent.com/FadeMind/hosts.extras/master/add.2o7Net/hosts",
	"https://raw.githubusercontent.com/crazy-max/WindowsSpyBlocker/master/data/hosts/spy.txt",
	"https://hostfiles.frogeye.fr/firstparty-trackers-hosts.txt",
	// - Malicious Lists
	"https://raw.githubusercontent.com/DandelionSprout/adfilt/master/Alternate%20versions%20Anti-Malware%20List/AntiMalwareHosts.txt",
	"https://v.firebog.net/hosts/Prigent-Crypto.txt",
	"https://raw.githubusercontent.com/FadeMind/hosts.extras/master/add.Risk/hosts",
	"https://phishing.army/download/phishing_army_blocklist_extended.txt",
	"https://gitlab.com/quidsup/notrack-blocklists/raw/master/notrack-malware.txt",
	"https://v.firebog.net/hosts/RPiList-Malware.txt",
	"https://raw.githubusercontent.com/Spam404/lists/master/main-blacklist.txt",
	"https://raw.githubusercontent.com/AssoEchap/stalkerware-indicators/master/generated/hosts",
	"https://urlhaus.abuse.ch/downloads/hostfile",
	"https://lists.cyberhost.uk/malware.txt",
	"https://malware-filter.gitlab.io/malware-filter/phishing-filter-hosts.txt",
	"https://v.firebog.net/hosts/Prigent-Malware.txt",
	"https://raw.githubusercontent.com/jarelllama/Scam-Blocklist/main/lists/wildcard_domains/scams.txt",
	"https://v.firebog.net/hosts/RPiList-Phishing.txt",
	// - Other Lists
	"https://raw.githubusercontent.com/chadmayfield/my-pihole-blocklists/master/lists/pi_blocklist_porn_top1m.list",
	// "https://v.firebog.net/hosts/Prigent-Adult.txt", # Huge 4.5M
	"https://raw.githubusercontent.com/anudeepND/blacklist/master/facebook.txt",
	// Others
	"https://badmojr.gitlab.io/1hosts/Pro/domains.txt",
	"https://gist.githubusercontent.com/eterps/9ddb13a118a21a7d9c12c6165e0bbff5/raw/0ba4b04802a4b478d7777fb7abe76c8eac0c5bfc/Samsung%2520Smart-TV%2520Blocklist%2520Adlist%2520(for%2520PiHole)",
	"https://gist.githubusercontent.com/wassname/b594c63222f9e4c83ea23c818440901b/raw/1b0afd2aecf3a099f1681b1cf18fc0e6e2fa116a/Samsung%2520Smart-TV%2520Blocklist%2520Adlist%2520(for%2520PiHole)",
	"https://gitlab.com/Kurobeats/phishing_hosts/raw/master/hosts",
	"https://raw.githubusercontent.com/MetaMask/eth-phishing-detect/master/src/hosts.txt",
	"https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/fakenews-gambling-porn/hosts",
	"https://raw.githubusercontent.com/blocklistproject/Lists/master/ads.txt",
	"https://raw.githubusercontent.com/blocklistproject/Lists/master/basic.txt",
	"https://raw.githubusercontent.com/blocklistproject/Lists/master/malware.txt",
	"https://raw.githubusercontent.com/blocklistproject/Lists/master/phishing.txt",
	"https://raw.githubusercontent.com/blocklistproject/Lists/master/ransomware.txt",
	"https://raw.githubusercontent.com/deathbybandaid/piholeparser/master/Subscribable-Lists/ParsedBlacklists/CryptoWall-Ransomware-C2-Domain-blocklist.txt",
	"https://raw.githubusercontent.com/deathbybandaid/piholeparser/master/Subscribable-Lists/ParsedBlacklists/Locky-Ransomware-C2-Domain-Blocklist.txt",
	"https://raw.githubusercontent.com/deathbybandaid/piholeparser/master/Subscribable-Lists/ParsedBlacklists/Ransomware-Domain-Blocklist.txt",
	"https://raw.githubusercontent.com/deathbybandaid/piholeparser/master/Subscribable-Lists/ParsedBlacklists/TeslaCrypt-Ransomware-C2-Domain-Blocklist.txt",
	"https://raw.githubusercontent.com/deathbybandaid/piholeparser/master/Subscribable-Lists/ParsedBlacklists/TeslaCrypt-Ransomware-Payment-Sites-Domain-Blocklist.txt",
	"https://raw.githubusercontent.com/deathbybandaid/piholeparser/master/Subscribable-Lists/ParsedBlacklists/TorrentLocker-Ransomware-C2-Domain-Blocklist.txt",
	"https://raw.githubusercontent.com/deathbybandaid/piholeparser/master/Subscribable-Lists/ParsedBlacklists/TorrentLocker-Ransomware-Payment-Sites-Domain-Blocklist.txt",
	"https://raw.githubusercontent.com/deathbybandaid/piholeparser/master/Subscribable-Lists/ParsedBlacklists/WindowsSpyBlocker81.txt",
	"https://raw.githubusercontent.com/esc0rtd3w/firestick-loader/master/misc/hosts.adfree",
	"https://raw.githubusercontent.com/kboghdady/youTube_ads_4_pi-hole/master/crowed_list.txt",
	"https://raw.githubusercontent.com/kboghdady/youTube_ads_4_pi-hole/master/youtubelist.txt",
	"https://raw.githubusercontent.com/pirat28/IHateTracker/master/iHateTracker.txt",
	"https://raw.githubusercontent.com/r-a-y/mobile-hosts/master/AdguardMobileSpyware.txt",
	"https://reddestdream.github.io/Projects/MinimalHosts/etc/MinimalHostsBlocker/minimalhosts",
}

func cleanLines(body []byte) []string {
	var lines []string
	for _, line := range strings.Split(string(body), "\n") {
		if !strings.HasPrefix(line, "#") {
			lines = append(lines, line)
		}
	}

	return lines
}

func readURL(url string) []string {
	response, err := http.Get(url)
	if err != nil {
		log.Panic("Error making GET request:", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Panic("Error reading response body:", err)
	}

	err = os.MkdirAll("raw", 0755)
	if err != nil {
		log.Panic("Error creating raw directory:", err)
	}

	filename := strings.TrimPrefix(url, "https://")
	filename = strings.ReplaceAll(filename, "/", "_")

	filePath := filepath.Join("raw", filename)
	err = os.WriteFile(filePath, body, 0644)
	if err != nil {
		log.Panic("Error writing file:", err)
	}

	return cleanLines(body)
}

func readFile(path string) []string {
	body, err := os.ReadFile(path)
	if err != nil {
		log.Panic("Error opening file:", err)
	}

	return cleanLines(body)
}

func writeToFile(urls []string) {
	file, err := os.Create("adlist.txt")
	if err != nil {
		log.Panic("Error creating file:", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, str := range urls {
		_, err := writer.WriteString(str + "\n") // Write string with a newline
		if err != nil {
			log.Panic("Error writing to file:", err)
		}
	}

	err = writer.Flush()
	if err != nil {
		log.Panic("Error flushing writer:", err)
	}
}

func main() {
	result := make(map[string]struct{})
	for _, url := range urls {
		lines := readURL(url)
		log.Printf("Read %6d - %s\n", len(lines), url)

		for _, line := range lines {
			// Split each line by spaces, due to:
			// 0.0.0.0 ai-apps-empire.com
			parts := strings.Fields(line)
			if len(parts) > 0 {
				for _, part := range parts {
					// Not sure why but https://v.firebog.net/hosts/Prigent-Crypto.txt, is like:
					// 0.0.0.0adminer.com
					row := strings.ReplaceAll(part, "0.0.0.0", "")

					result[row] = struct{}{}
				}
			}
		}
	}
	log.Printf("Read %d from urls\n", len(result))

	linesBlock := readFile("blocklist.txt")
	log.Printf("Read %d from blocklist.txt\n", len(linesBlock))
	for _, line := range linesBlock {
		result[line] = struct{}{}
	}

	linesAllow := readFile("allowlist.txt")
	log.Printf("Read %d from allowlist.txt\n", len(linesAllow))
	for _, line := range linesAllow {
		_, exists := result[line]
		if exists {
			delete(result, line)
		}
	}

	hosts := []string{}
	for item := range result {
		// Do not understand why but domain with "/" pass
		if strings.Contains(item, "/") {
			continue
		}

		host, err := normalizeHostsListEntry(item)
		if err == nil {
			hosts = append(hosts, host)
		}
	}
	sort.Strings(hosts)

	log.Printf("Read %d after normalizeHostsListEntry\n", len(hosts))

	writeToFile(hosts)
}
