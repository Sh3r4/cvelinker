# CVElinker

I threw this together to make it easier to dig through advisory bulletins that name CVEs but don't link to any further information on them.

It will look in the provided input for matches against the `CVE-yyyy-ID` format ignoring case.

Hopefully it will help someone save at least a little bit of time.

## Example Output

This example takes input from the clipboard.

``` txt
$ ./cvelinker
========== CVE-2018-3640 Links ==========
[NVD__] https://nvd.nist.gov/vuln/detail/CVE-2018-3640
[MITRE] https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-3640
[LVD__] https://lwn.net/Search/DoSearch?words=CVE-2018-3640
[MTASP] https://www.rapid7.com/db/search?q=CVE-2018-3640
[DEB__] https://security-tracker.debian.org/tracker/CVE-2018-3640
[RHEL_] https://bugzilla.redhat.com/show_bug.cgi?id=CVE-2018-3640
[GOOGL] https://www.google.com/search?q=CVE-2018-3640
[GITHB] https://github.com/search?q=%22CVE-2018-3640%22&type=Issues


========== CVE-2018-3639 Links ==========
[NVD__] https://nvd.nist.gov/vuln/detail/CVE-2018-3639
[MITRE] https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-3639
[LVD__] https://lwn.net/Search/DoSearch?words=CVE-2018-3639
[MTASP] https://www.rapid7.com/db/search?q=CVE-2018-3639
[DEB__] https://security-tracker.debian.org/tracker/CVE-2018-3639
[RHEL_] https://bugzilla.redhat.com/show_bug.cgi?id=CVE-2018-3639
[GOOGL] https://www.google.com/search?q=CVE-2018-3639
[GITHB] https://github.com/search?q=%22CVE-2018-3639%22&type=Issues
```

## Usage

``` txt
Usage of ./cvelinker:
  -i, --input string
    	Pull CVEs from specified input file
  -o, --output string
    	Specify file to output formatted links to
```

### Clipboard

This is the main intended usage path.

1. Copy text which includes CVEs in `CVE-yyyy-ID` format to clipboard
1. Run CVElinker
