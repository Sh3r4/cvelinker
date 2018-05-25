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

var o levelOutput
var online bool
var circlAPI = "https://cve.circl.lu/api/cve/"

func main() {
	// declarations
	var verbose bool
	var cveInput string
	var cveOutput string
	var builderData = []cveLinkData{
		cveLinkData{"NVD", "https://nvd.nist.gov/vuln/detail/"},
		cveLinkData{"MITRE", "https://cve.mitre.org/cgi-bin/cvename.cgi?name="},
		cveLinkData{"LVD", "https://lwn.net/Search/DoSearch?words="},
		cveLinkData{"MTASP", "https://www.rapid7.com/db/search?q="},
		cveLinkData{"DEB", "https://security-tracker.debian.org/tracker/"},
		cveLinkData{"RHEL", "https://bugzilla.redhat.com/show_bug.cgi?id="},
		cveLinkData{"GOOGL", "https://www.google.com/search?q="},
	}

	// declare flags and parse
	flag.BoolVarP(&online, "api-enabled", "a", false, "Gathers Extra information via API")
	flag.StringVarP(&cveInput, "input", "i", "", "Pull CVEs from specified input file")
	flag.StringVarP(&cveOutput, "output", "o", "", "Specify file to output formatted markdown report to")
	flag.BoolVarP(&verbose, "verbose", "v", false, "Set verbose mode on")
	flag.Parse()

	if verbose {
		o.Init(4, false, false)
	} else {
		o.Init(3, false, false)
	}

	cves := make(map[string]cve)

	// deal with an input file flag
	if len(cveInput) > 0 {
		//attempt to open and read into buf
		content, err := ioutil.ReadFile(cveInput)
		if err != nil {
			o.Fatality.Fatalf("Error opening file for read: %s", err)
			os.Exit(1)
		}

		// ingest the cve tokens
		strContent := string(content)
		var cveTokens = ingestCveTokens(strContent)
		for _, argument := range cveTokens {
			cves = addUniqueCVE(cves, argument, builderData)
		}

	}

	// if there were no args, look for clipboard contents
	if len(flag.Args()) < 1 {
		// take from current clipboard
		clipboardcontents, err := clipboard.ReadAll()
		if err != nil {
			o.Fatality.Fatalf("Error reading clipboard: %s", err)
		}

		// ingest the tokens and add links for valid CVE IDs to the list
		var cveTokens = ingestCveTokens(clipboardcontents)
		for _, argument := range cveTokens {
			cves = addUniqueCVE(cves, argument, builderData)
		}
	} else { // Since there were excess arguments, treat each one as a potential cve ID
		for _, argument := range flag.Args() {

			// test uppercased arg for cve and add if valid, discard if not
			upperArgument := strings.ToUpper(argument)
			if testTokenForCveNess(upperArgument) {
				cves = addUniqueCVE(cves, upperArgument, builderData)
			} else {
				o.OverShare.Println("Discarded invalid argument: " + upperArgument)
			}
		}
	}

	// if output to file is set, open the default output path and overwrite everything
	if len(cveOutput) > 0 {
		outputbytes := orchestrateMarkdownBuild(cves)

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
		for _, c := range cves {
			o.Print.Println(c.FormatOutputLines())
		}
	}
}

func addUniqueCVE(cveMap map[string]cve, cveid string, builderData []cveLinkData) map[string]cve {
	id := strings.ToUpper(cveid)

	// skip if it already exists in the map
	if _,ok := cveMap[id]; !ok {
		cveMap[id] = collectAndBuildCveData(id, builderData)
	}

	return cveMap
}

func orchestrateMarkdownBuild(cves map[string]cve) []byte {
	lines := []byte("# CVE Evaluation Report\n\n")

	for _, c := range cves {
		lines = appendSlice(lines, c.FormatOutputMarkdown())
	}

	return lines
}

func collectAndBuildCveData(cvestr string, builderData []cveLinkData) cve {

	builder := cve{}
	builder.Init(cvestr, builderData)
	if online {
		builder.PopulateCveDetails(circlAPI)
	}

	return builder
}

// ingestCveTokens extracts all CVE IDs from a given input string.
// It returns a string array of the IDs in 'CVE-yyyy-idnum' format.
func ingestCveTokens(input string) []string {
	re := regexp.MustCompile(`([cC][vV][eE]\-\d{4}\-\d{4,})`)
	tokens := re.FindAllString(input, -1)
	var cveIds []string

	// iterate the tokens
	for _, token := range tokens {
		// convert to Upper cause all good vulns are uppercase
		upperToken := strings.ToUpper(token)

		// append if it is a cveid
		if testTokenForCveNess(upperToken) {
			cveIds = append(cveIds, upperToken)
		}
	}
	return cveIds
}

// appendSlice appends the contents of a string array to another string array.
// It returns the combined string array
func appendSlice(frontSlice []byte, rearSlice []byte) []byte {
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
