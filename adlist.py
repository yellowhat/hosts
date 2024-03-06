#!/usr/bin/env python
"""Check adlist"""

import re
from urllib.request import urlopen

URLS = [
    # From firebog.net
    "https://adaway.org/hosts.txt",
    "https://bitbucket.org/ethanr/dns-blacklists/raw/8575c9f96e5b4a1308f2f12394abd86d0927a4a0/bad_lists/Mandiant_APT1_Report_Appendix_D.txt",
    "https://gitlab.com/quidsup/notrack-blocklists/raw/master/notrack-blocklist.txt",
    "https://gitlab.com/quidsup/notrack-blocklists/raw/master/notrack-malware.txt",
    "https://hostfiles.frogeye.fr/firstparty-trackers-hosts.txt",
    "https://malware-filter.gitlab.io/malware-filter/phishing-filter-hosts.txt",
    "https://osint.digitalside.it/Threat-Intel/lists/latestdomains.txt",
    "https://paulgb.github.io/BarbBlock/blacklists/hosts-file.txt",
    "https://pgl.yoyo.org/adservers/serverlist.php?hostformat=hosts&showintro=0&mimetype=plaintext",
    "https://phishing.army/download/phishing_army_blocklist_extended.txt",
    "https://raw.githubusercontent.com/AssoEchap/stalkerware-indicators/master/generated/hosts",
    "https://raw.githubusercontent.com/DandelionSprout/adfilt/master/Alternate%20versions%20Anti-Malware%20List/AntiMalwareHosts.txt",
    "https://raw.githubusercontent.com/FadeMind/hosts.extras/master/UncheckyAds/hosts",
    "https://raw.githubusercontent.com/FadeMind/hosts.extras/master/add.2o7Net/hosts",
    "https://raw.githubusercontent.com/FadeMind/hosts.extras/master/add.Risk/hosts",
    "https://raw.githubusercontent.com/FadeMind/hosts.extras/master/add.Spam/hosts",
    "https://raw.githubusercontent.com/Perflyst/PiHoleBlocklist/master/AmazonFireTV.txt",
    "https://raw.githubusercontent.com/Perflyst/PiHoleBlocklist/master/SmartTV.txt",
    "https://raw.githubusercontent.com/Perflyst/PiHoleBlocklist/master/android-tracking.txt",
    "https://raw.githubusercontent.com/PolishFiltersTeam/KADhosts/master/KADhosts.txt",
    "https://raw.githubusercontent.com/RooneyMcNibNug/pihole-stuff/master/SNAFU.txt",
    "https://raw.githubusercontent.com/Spam404/lists/master/main-blacklist.txt",
    "https://raw.githubusercontent.com/VeleSila/yhosts/master/hosts",
    "https://raw.githubusercontent.com/anudeepND/blacklist/master/adservers.txt",
    "https://raw.githubusercontent.com/anudeepND/blacklist/master/facebook.txt",
    "https://raw.githubusercontent.com/bigdargon/hostsVN/master/hosts",
    "https://raw.githubusercontent.com/crazy-max/WindowsSpyBlocker/master/data/hosts/spy.txt",
    "https://raw.githubusercontent.com/jdlingyu/ad-wars/master/hosts",
    "https://raw.githubusercontent.com/matomo-org/referrer-spam-blacklist/master/spammers.txt",
    "https://s3.amazonaws.com/lists.disconnect.me/simple_ad.txt",
    "https://s3.amazonaws.com/lists.disconnect.me/simple_malvertising.txt",
    "https://someonewhocares.org/hosts/zero/hosts",
    "https://urlhaus.abuse.ch/downloads/hostfile/",
    "https://v.firebog.net/hosts/AdguardDNS.txt",
    "https://v.firebog.net/hosts/Admiral.txt",
    "https://v.firebog.net/hosts/Easylist.txt",
    "https://v.firebog.net/hosts/Easyprivacy.txt",
    "https://v.firebog.net/hosts/Prigent-Ads.txt",
    # "https://v.firebog.net/hosts/Prigent-Adult.txt",  # Huge 4.5M
    "https://v.firebog.net/hosts/Prigent-Crypto.txt",
    "https://v.firebog.net/hosts/Prigent-Malware.txt",
    "https://v.firebog.net/hosts/RPiList-Malware.txt",
    "https://v.firebog.net/hosts/RPiList-Phishing.txt",
    "https://v.firebog.net/hosts/neohostsbasic.txt",
    "https://v.firebog.net/hosts/static/w3kbl.txt",
    "https://winhelp2002.mvps.org/hosts.txt",
    "https://www.github.developerdan.com/hosts/lists/ads-and-tracking-extended.txt",
    "https://zerodot1.gitlab.io/CoinBlockerLists/hosts_browser",
    # Others
    "https://badmojr.gitlab.io/1hosts/Pro/domains.txt",
    "https://gist.githubusercontent.com/eterps/9ddb13a118a21a7d9c12c6165e0bbff5/raw/0ba4b04802a4b478d7777fb7abe76c8eac0c5bfc/Samsung%2520Smart-TV%2520Blocklist%2520Adlist%2520(for%2520PiHole)",
    "https://gist.githubusercontent.com/wassname/b594c63222f9e4c83ea23c818440901b/raw/1b0afd2aecf3a099f1681b1cf18fc0e6e2fa116a/Samsung%2520Smart-TV%2520Blocklist%2520Adlist%2520(for%2520PiHole)",
    "https://gitlab.com/Kurobeats/phishing_hosts/raw/master/hosts",
    "https://perflyst.github.io/PiHoleBlocklist/SmartTV.txt",
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
    "https://raw.githubusercontent.com/kboghdady/youTube_ads_4_pi-hole/master/crowed_list.txt",
    "https://raw.githubusercontent.com/kboghdady/youTube_ads_4_pi-hole/master/youtubelist.txt",
    "https://raw.githubusercontent.com/pirat28/IHateTracker/master/iHateTracker.txt",
    "https://raw.githubusercontent.com/r-a-y/mobile-hosts/master/AdguardMobileSpyware.txt",
    "https://reddestdream.github.io/Projects/MinimalHosts/etc/MinimalHostsBlocker/minimalhosts",
    "https://raw.githubusercontent.com/esc0rtd3w/firestick-loader/master/misc/hosts.adfree",
]


def read_txt(filename: str) -> list[str]:
    """Read a file and remove empty and comment line"""

    with open(filename, encoding="UTF8") as obj:
        lines = obj.read().splitlines()

    return [line for line in lines if line and line[0] != "#"]


def get_url(_url: str) -> set[str]:
    """Fetch domains from url"""

    # nosemgrep: python.lang.security.audit.dynamic-urllib-use-detected.dynamic-urllib-use-detected
    with urlopen(_url) as obj:  # nosec B310
        items = set(obj.read().decode("utf8").splitlines())
    print(f" - {len(items):9,} - {_url}")

    return items


if __name__ == "__main__":

    allowlist = read_txt("allowlist.txt")
    blocklist = read_txt("blocklist.txt")

    # Gather from urls
    domains = set()
    for idx, url in enumerate(URLS):
        print(f"{idx: 2d}/{len(URLS): 2d}", end="")
        domains.update(get_url(url))
    domains.remove("")
    print()
    print(f"[INFO] Read: {len(domains):,}")
    print()

    # Filter only domain
    domains_filtered = set()
    for domain in domains:
        if "#" in domain:
            continue
        # match = re.sub("#.*", "", domain).strip().split()[-1].strip()
        # matches = re.findall(r"[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+", domain)
        # https://regexr.com/3au3g
        matches = re.findall(
            r"(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]",
            domain,
        )
        for match in matches:
            domains_filtered.add(match)

    # Add blocklist
    for domain in blocklist:
        domains_filtered.add(domain)

    # Remove from allowlist
    for domain in allowlist:
        domains_filtered.discard(domain)

    print(f"[INFO] Final: {len(domains_filtered):,}")
    print()

    with open("adlist.txt", "w", encoding="utf8") as writer:
        for domain in sorted(domains_filtered):
            writer.write(f"{domain}\n")
