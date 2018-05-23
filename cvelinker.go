package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
	flag "github.com/ogier/pflag"
)

func main() {
	// declarations
	var cveInput string
	var cveOutput string
	var links []string
	var builderData []cveLinkData = []cveLinkData{
		cveLinkData{"NVD", "https://nvd.nist.gov/vuln/detail/"},
		cveLinkData{"MITRE", "https://cve.mitre.org/cgi-bin/cvename.cgi?name="},
		cveLinkData{"LVD", "https://lwn.net/Search/DoSearch?words="},
		cveLinkData{"MTASP", "https://www.rapid7.com/db/search?q="},
		cveLinkData{"DEB", "https://security-tracker.debian.org/tracker/"},
		cveLinkData{"RHEL", "https://bugzilla.redhat.com/show_bug.cgi?id="},
		cveLinkData{"GOOGL", "https://www.google.com/search?q="},
	}

	// declare flags and parse
	flag.StringVarP(&cveInput, "input", "i", "", "Pull CVEs from specified input file")
	flag.StringVarP(&cveOutput, "output", "o", "", "Specify file to output formatted links to")
	flag.Parse()

	// deal with an input file flag
	if len(cveInput) > 0 {
		//attempt to open and read into buf
		content, err := ioutil.ReadFile(cveInput)
		if err != nil {
			fmt.Println("[!] Error opening file for read: ")
			fmt.Println(err)
			os.Exit(1)
		}

		// coerce buf into string
		strContent := string(content)

		// ingest the tokens and add links for valid CVE IDs to the list
		var cveTokens []string = ingestCveTokens(strContent)
		for _, argument := range cveTokens {
			links = appendSlice(links, collectAllCveLinksAndFormat(argument, builderData))
		}

	}

	// if there were no args, look for clipboard contents
	if len(flag.Args()) < 1 {
		// take from current clipboard
		clipboardcontents, err := clipboard.ReadAll()
		if err != nil {
			fmt.Println("[!] Error reading from clipboard")
			fmt.Println(err)
			os.Exit(1)
		}

		// ingest the tokens and add links for valid CVE IDs to the list
		var cveTokens []string = ingestCveTokens(clipboardcontents)
		for _, argument := range cveTokens {
			links = appendSlice(links, collectAllCveLinksAndFormat(argument, builderData))
		}
	} else { // Since there were excess arguments, treat each one as a potential cve ID
		for _, argument := range flag.Args() {

			// test uppercased arg for cve and add if valid, discard if not
			upperargument := strings.ToUpper(argument)
			if testTokenForCveNess(upperargument) {
				links = appendSlice(links, collectAllCveLinksAndFormat(upperargument, builderData))
			} else {
				fmt.Println("[*] Discarded invalid argument: " + upperargument)
			}
		}
	}

	// if output to file is set, open the default output path and overwrite everything
	if len(cveOutput) > 0 {
		outputstring := "\n"
		for _, l := range links {
			outputstring += l + "\n"
		}

		// coerce into bytes
		var outputbytes = []byte(outputstring + "\n")

		// open output file
		fo, err := os.Create(cveOutput)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// close fo on exit and check for its returned error
		defer func() {
			if err := fo.Close(); err != nil {
				fmt.Println("[!] Error closing output file")
				fmt.Println(err)
				os.Exit(1)
			}
		}()

		// write all lines
		if _, err := fo.Write(outputbytes); err != nil {
			panic(err)
		}
	} else { // if not write to file, just output the lines to stdout
		for _, l := range links {
			fmt.Println(l)
		}
	}
}

func collectAllCveLinksAndFormat(cve string, builderData []cveLinkData) []string {
	// start off with a banner and the 'standard' links
	lines := []string{"========== " + cve + " Links =========="}
	lines = appendSlice(lines, buildLinksFromData(builderData, cve))

	// deal with special cases
	lines = appendSlice(lines, buildLinksForSpecialCases(cve))
	lines = append(lines, "\n")

	return lines
}

// buildLinksFromData builds formatted output for a single cve.
// It returns an array of string lines ready for output to file or console
func buildLinksFromData(builderData []cveLinkData, cve string) []string {
	lines := []string{}

	// iterate the data array
	for _, data := range builderData {
		// build a sweet square bracket id for the link
		line := "[" + data.name
		for i := len(data.name); i < 5; i++ {
			line += "_"
		}
		line += "] "

		// create the link and append to the slice
		line += data.url + cve
		lines = append(lines, line)
	}

	return lines
}

// buildLinksForSpecialCases handles building URLs which are not in the more standard prepend format
func buildLinksForSpecialCases(cve string) []string {
	var links = []string{}

	// special case for github issue search
	githubLink := "[GITHB] https://github.com/search?q=%22" + cve + "%22&type=Issues"
	links = append(links, githubLink)

	return links
}

// ingestCveTokens extracts all CVE IDs from a given input string.
// It returns a string array of the IDs in 'CVE-yyyy-idnum' format.
func ingestCveTokens(input string) []string {
	re := regexp.MustCompile(`([cC][vV][eE]\-\d{4}\-\d{4,})`)
	tokens := re.FindAllString(input, -1)
	cves := []string{}

	// iterate the tokens
	for _, token := range tokens {
		// convert to Upper cause all good vulns are uppercase
		uppertoken := strings.ToUpper(token)

		// append if it is a cveid
		if testTokenForCveNess(uppertoken) {
			cves = append(cves, uppertoken)
		}
	}
	return cves
}

// appendSlice appends the contents of a string array to another string array.
// It returns the combined string array
func appendSlice(frontSlice []string, rearSlice []string) []string {
	slice := frontSlice

	for _, stringVal := range rearSlice {
		slice = append(slice, stringVal)
	}

	return slice
}

// testTokenForCveNess checks a given string against a rudimentary regex for the CVE ID format.
// It returns true if the input matches the CVEID pattern.
func testTokenForCveNess(input string) bool {
	var validCve = regexp.MustCompile(`CVE\-\d{4}\-\d{4,}`)
	return validCve.MatchString(input)
}

// make an object for stuff so more info can be added later
type cveLinkData struct {
	name string
	url  string
}
