package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type cve struct {
	id      string
	details cveDeets
	links   []cveLinkData
}

type cveDeets struct {
	Published  string   `json:"Published"`
	Modified   string   `json:"Modified"`
	Cvss       string   `json:"cvss"`
	References []string `json:"references"`
	Summary    string   `json:"summary"`
}

func (c *cve) Init(cve string, linksData []cveLinkData) {
	if testTokenForCveNess(strings.ToUpper(cve)) {
		c.id = strings.ToUpper(cve)
		c.links = linksData
	}
}

func (c *cve) PopulateCveDetails(apiURL string) {
	// test the id
	if len(c.id) < 9 {
		o.Warn.Println("Setting %[s] details to nil. Malformed cve?", c.id)
		return

	}

	// make the request to the api
	resp, err := http.Get(apiURL + strings.ToUpper(c.id))
	if err != nil {
		o.Warn.Println("Failed to query API at: " + apiURL + strings.ToUpper(c.id))
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		o.Warn.Println("Failed to query API at: " + apiURL + strings.ToUpper(c.id))
		return
	}

	// unmarshal the data
	err = json.Unmarshal(body, &c.details)
	if err != nil {
		o.Warn.Println("Failed to interpret JSON Response for " + strings.ToUpper(c.id))
		c.details = cveDeets{}
		return
	}

	// instigate a fallback plan -- do nothing
	if string(body) == "null" || c.details.Summary == "" {
		c.details.Published = "n/a"
		c.details.Modified = "n/a"
		c.details.Cvss = "n/a"
		c.details.Summary = "This CVE ID either does not exist or has not been publicly disclosed yet. Check the links for more details"
	}

}

func (c *cve) FormatOutputLines() string {
	var lines string
	// format ID line
	idline := "=" + o.green.Sprintf(" %s ", c.id) + "="
	idline = rightPad2Len(idline, "=", 40+len(c.id))
	idline = leftPad2Len(idline, "=", 90-len(c.id))
	lines += idline + "\n"

	// deal with details
	if len(c.details.Published) > 0 {
		lines += "Published: " + c.details.Published + "\n"
	}
	if len(c.details.Modified) > 0 {
		lines += "Modified:  " + c.details.Modified + "\n"
	}
	if len(c.details.Cvss) > 0 {
		lines += "CVSS: " + c.details.Cvss + "\n"
	}
	if len(c.details.Summary) > 0 {
		lines += wordWrap("Summary: "+c.details.Summary, 65) + "\n"
	}

	// output all the links with formatted names
	lines += "\n"
	for _, ref := range c.links {
		// build a sweet square bracket id for the link
		lineID := "[" + ref.name
		for i := len(ref.name); i < 5; i++ {
			lineID += "_"
		}
		lineID += "] "

		lines += o.green.Sprintf(lineID) + ref.url + strings.ToUpper(c.id) + "\n"
	}

	// deal with special case links which are dumb
	lines+= o.green.Sprintf("[GITHB] ") + "https://github.com/search?q=%22" + c.id + "%22&type=Issues\n"


	// output extra references
	if len(c.details.References) > 0 {
		lines += o.green.Sprint("        *** Supplemental ***\n")
		for _, link := range c.details.References {
			lines += "    [-] " + link + "\n"
		}
	}

	lines += "\n\n"
	return lines
}

func (c *cve) FormatOutputMarkdown() []byte {

	lines := "## " + c.id + "\n\n"

	if len(c.details.Cvss) > 0 {
		lines += "**CVSS Score:** " + c.details.Cvss + "\n\n"
	} else {
		lines += "**CVSS Score:** n/a\n\n"
	}

	if len(c.details.Published) > 0 {
		lines += "**Published:** " + c.details.Published + "\n\n"
	} else {
		lines += "**Published:** n/a\n\n"
	}
	if len(c.details.Modified) > 0 {
		lines += "**Modified :** " + c.details.Modified + "\n\n"
	} else {
		lines += "**Modified :** n/a\n\n"
	}






	lines += "### Summary\n\n"

	lines += c.details.Summary + "\n\n"

	lines += "### Generated Links\n\n"

	for _, l := range c.links {
		// build a sweet square bracket id for the link
		lineID := l.name
		for i := len(l.name); i < 5; i++ {
			lineID += "-"
		}
		lines += "* [" + lineID + "> " + l.url + strings.ToUpper(c.id) + "](" + l.url + strings.ToUpper(c.id) + ")\n"

	}
	lines += "\n"

	if len(c.details.References) > 0 {
		lines += "### Additional References\n\n"
		for _, l := range c.details.References {
			lines += "* [" + l + "](" + l + ")\n"
		}
		lines += "\n"
	}

	return []byte(lines)
}

// make an object for stuff so more info can be added later
type cveLinkData struct {
	name string
	url  string
}
