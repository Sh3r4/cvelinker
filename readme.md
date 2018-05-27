# CVElinker

``` txt
                               $$\ $$\           $$\
                               $$ |\__|          $$ |
 $$$$$$$\ $$\    $$\  $$$$$$\  $$ |$$\ $$$$$$$\  $$ |  $$\  $$$$$$\   $$$$$$\
$$  _____|\$$\  $$  |$$  __$$\ $$ |$$ |$$  __$$\ $$ | $$  |$$  __$$\ $$  __$$\
$$ /       \$$\$$  / $$$$$$$$ |$$ |$$ |$$ |  $$ |$$$$$$  / $$$$$$$$ |$$ |  \__|
$$ |        \$$$  /  $$   ____|$$ |$$ |$$ |  $$ |$$  _$$<  $$   ____|$$ |
\$$$$$$$\    \$  /   \$$$$$$$\ $$ |$$ |$$ |  $$ |$$ | \$$\ \$$$$$$$\ $$ |
 \_______|    \_/     \_______|\__|\__|\__|  \__|\__|  \__| \_______|\__|

Author:  Morgaine "sectorsect" Timms
License: MIT
```

I threw this together to make it easier to dig through advisory bulletins that name CVEs but don't link to any further information on them.

It will look in the provided input for matches against the `CVE-yyyy-ID` format ignoring case.

Hopefully it will help someone save at least a little bit of time.

## Installation

`go get github.com/sectorsect/cvelinker`

## Usage

CVElinker will read the clipboard and extract CVE IDs.

If network mode is enabled with `-a`, data from the CIRCL CVE-Search api will be included in the output.

```
Things It Does:
  -a, --api-enabled
    	Gathers Extra information via API
  -i, --input string
    	Pull CVEs from specified input file
  -o, --output string
    	Specify file to output formatted markdown report to
  -v, --verbose
    	Set verbose mode on
```

## Output Formats

CVELinker will output to stdout by default, or a markdown file optionally.

[Example Markdown Report Link](exampleReport.md)

example CLI output with api-calls enabled:

```

========================= CVE-2018-3640 ==========================
Published: 2018-05-22T08:29:00.327000
Modified:  2018-05-22T08:29:00.343000
Summary: Systems with microprocessors utilizing speculative
execution and that perform speculative reads of system registers
may allow unauthorized disclosure of system parameters to an
attacker with local user access via a side-channel analysis, aka
Rogue System Register Read (RSRE), Variant 3a.

[NVD__] https://nvd.nist.gov/vuln/detail/CVE-2018-3640
[MITRE] https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-3640
[LVD__] https://lwn.net/Search/DoSearch?words=CVE-2018-3640
[MTASP] https://www.rapid7.com/db/search?q=CVE-2018-3640
[DEB__] https://security-tracker.debian.org/tracker/CVE-2018-3640
[RHEL_] https://bugzilla.redhat.com/show_bug.cgi?id=CVE-2018-3640
[GOOGL] https://www.google.com/search?q=CVE-2018-3640
[GITHB] https://github.com/search?q=%22CVE-2018-3640%22&type=Issues
        *** Supplemental ***
    [-] http://support.lenovo.com/us/en/solutions/LEN-22133
    [-] http://www.securityfocus.com/bid/104228
    [-] http://www.securitytracker.com/id/1040949
    [-] https://developer.arm.com/support/arm-security-updates/speculative-processor-vulnerability
    [-] https://portal.msrc.microsoft.com/en-us/security-guidance/advisory/ADV180013
    [-] https://security.netapp.com/advisory/ntap-20180521-0001/
    [-] https://tools.cisco.com/security/center/content/CiscoSecurityAdvisory/cisco-sa-20180521-cpusidechannel
    [-] https://www.intel.com/content/www/us/en/security-center/advisory/intel-sa-00115.html
    [-] https://www.kb.cert.org/vuls/id/180049
    [-] https://www.synology.com/support/security/Synology_SA_18_23
    [-] https://www.us-cert.gov/ncas/alerts/TA18-141A



========================= CVE-2018-3639 ==========================
Published: 2018-05-22T08:29:00.250000
Modified:  2018-05-22T08:29:00.280000
Summary: Systems with microprocessors utilizing speculative
execution and speculative execution of memory reads before the
addresses of all prior memory writes are known may allow
unauthorized disclosure of information to an attacker with local
user access via a side-channel analysis, aka Speculative Store
Bypass (SSB), Variant 4.

[NVD__] https://nvd.nist.gov/vuln/detail/CVE-2018-3639
[MITRE] https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-3639
[LVD__] https://lwn.net/Search/DoSearch?words=CVE-2018-3639
[MTASP] https://www.rapid7.com/db/search?q=CVE-2018-3639
[DEB__] https://security-tracker.debian.org/tracker/CVE-2018-3639
[RHEL_] https://bugzilla.redhat.com/show_bug.cgi?id=CVE-2018-3639
[GOOGL] https://www.google.com/search?q=CVE-2018-3639
[GITHB] https://github.com/search?q=%22CVE-2018-3639%22&type=Issues
        *** Supplemental ***
    [-] http://support.lenovo.com/us/en/solutions/LEN-22133
    [-] http://www.securityfocus.com/bid/104232
    [-] http://www.securitytracker.com/id/1040949
    [-] http://xenbits.xen.org/xsa/advisory-263.html
    [-] https://access.redhat.com/errata/RHSA-2018:1629
    [-] https://access.redhat.com/errata/RHSA-2018:1630
    [-] https://access.redhat.com/errata/RHSA-2018:1632
    [-] https://access.redhat.com/errata/RHSA-2018:1633
    [-] https://access.redhat.com/errata/RHSA-2018:1635
    [-] https://access.redhat.com/errata/RHSA-2018:1636
    [-] https://access.redhat.com/errata/RHSA-2018:1642
    [-] https://access.redhat.com/errata/RHSA-2018:1643
    [-] https://access.redhat.com/errata/RHSA-2018:1644
    [-] https://access.redhat.com/errata/RHSA-2018:1645
    [-] https://access.redhat.com/errata/RHSA-2018:1646
    [-] https://access.redhat.com/errata/RHSA-2018:1647
    [-] https://access.redhat.com/errata/RHSA-2018:1648
    [-] https://access.redhat.com/errata/RHSA-2018:1649
    [-] https://access.redhat.com/errata/RHSA-2018:1650
    [-] https://access.redhat.com/errata/RHSA-2018:1651
    [-] https://access.redhat.com/errata/RHSA-2018:1652
    [-] https://access.redhat.com/errata/RHSA-2018:1653
    [-] https://access.redhat.com/errata/RHSA-2018:1654
    [-] https://access.redhat.com/errata/RHSA-2018:1655
    [-] https://access.redhat.com/errata/RHSA-2018:1656
    [-] https://access.redhat.com/errata/RHSA-2018:1657
    [-] https://access.redhat.com/errata/RHSA-2018:1658
    [-] https://access.redhat.com/errata/RHSA-2018:1659
    [-] https://access.redhat.com/errata/RHSA-2018:1660
    [-] https://access.redhat.com/errata/RHSA-2018:1661
    [-] https://access.redhat.com/errata/RHSA-2018:1662
    [-] https://access.redhat.com/errata/RHSA-2018:1663
    [-] https://access.redhat.com/errata/RHSA-2018:1664
    [-] https://access.redhat.com/errata/RHSA-2018:1665
    [-] https://access.redhat.com/errata/RHSA-2018:1666
    [-] https://access.redhat.com/errata/RHSA-2018:1667
    [-] https://access.redhat.com/errata/RHSA-2018:1668
    [-] https://access.redhat.com/errata/RHSA-2018:1669
    [-] https://access.redhat.com/errata/RHSA-2018:1674
    [-] https://access.redhat.com/errata/RHSA-2018:1675
    [-] https://access.redhat.com/errata/RHSA-2018:1676
    [-] https://access.redhat.com/errata/RHSA-2018:1686
    [-] https://access.redhat.com/errata/RHSA-2018:1688
    [-] https://access.redhat.com/errata/RHSA-2018:1689
    [-] https://access.redhat.com/errata/RHSA-2018:1690
    [-] https://access.redhat.com/errata/RHSA-2018:1696
    [-] https://access.redhat.com/errata/RHSA-2018:1710
    [-] https://access.redhat.com/errata/RHSA-2018:1711
    [-] https://bugs.chromium.org/p/project-zero/issues/detail?id=1528
    [-] https://developer.arm.com/support/arm-security-updates/speculative-processor-vulnerability
    [-] https://portal.msrc.microsoft.com/en-US/security-guidance/advisory/ADV180012
    [-] https://security.netapp.com/advisory/ntap-20180521-0001/
    [-] https://support.citrix.com/article/CTX235225
    [-] https://tools.cisco.com/security/center/content/CiscoSecurityAdvisory/cisco-sa-20180521-cpusidechannel
    [-] https://usn.ubuntu.com/3651-1/
    [-] https://usn.ubuntu.com/3652-1/
    [-] https://usn.ubuntu.com/3653-1/
    [-] https://usn.ubuntu.com/3653-2/
    [-] https://usn.ubuntu.com/3654-1/
    [-] https://usn.ubuntu.com/3654-2/
    [-] https://usn.ubuntu.com/3655-2/
    [-] https://www.exploit-db.com/exploits/44695/
    [-] https://www.intel.com/content/www/us/en/security-center/advisory/intel-sa-00115.html
    [-] https://www.kb.cert.org/vuls/id/180049
    [-] https://www.synology.com/support/security/Synology_SA_18_23
    [-] https://www.us-cert.gov/ncas/alerts/TA18-141A

```