package snyk

import "time"

type ListAllIssuesResp struct {
	Ok              bool   `json:"ok"`
	Issues          Issues `json:"issues"`
	DependencyCount int    `json:"dependencyCount"`
	PackageManager  string `json:"packageManager"`
}
type Semver struct {
	Unaffected string `json:"unaffected"`
	Vulnerable string `json:"vulnerable"`
}
type Identifiers struct {
	CVE         []string `json:"CVE"`
	CWE         []string `json:"CWE"`
	ALTERNATIVE []string `json:"ALTERNATIVE"`
}
type Patches struct {
	ID               string    `json:"id"`
	Urls             []string  `json:"urls"`
	Version          string    `json:"version"`
	Comments         []string  `json:"comments"`
	ModificationTime time.Time `json:"modificationTime"`
}
type Ignored struct {
	Reason  string    `json:"reason"`
	Expires time.Time `json:"expires"`
	Source  string    `json:"source"`
}
type Patched struct {
	Patched time.Time `json:"patched"`
}
type Vulnerabilities struct {
	ID              string      `json:"id"`
	URL             string      `json:"url"`
	Title           string      `json:"title"`
	Type            string      `json:"type"`
	Description     string      `json:"description"`
	From            []string    `json:"from"`
	Package         string      `json:"package"`
	Version         string      `json:"version"`
	Severity        string      `json:"severity"`
	Language        string      `json:"language"`
	PackageManager  string      `json:"packageManager"`
	Semver          Semver      `json:"semver"`
	PublicationTime time.Time   `json:"publicationTime"`
	DisclosureTime  time.Time   `json:"disclosureTime"`
	IsUpgradable    bool        `json:"isUpgradable"`
	IsPatchable     bool        `json:"isPatchable"`
	Identifiers     Identifiers `json:"identifiers"`
	Credit          []string    `json:"credit"`
	CVSSv3          string      `json:"CVSSv3"`
	CvssScore       float64     `json:"cvssScore"`
	Patches         []Patches   `json:"patches"`
	IsIgnored       bool        `json:"isIgnored"`
	IsPatched       bool        `json:"isPatched"`
	UpgradePath     []string    `json:"upgradePath"`
	Ignored         []Ignored   `json:"ignored,omitempty"`
	Patched         []Patched   `json:"patched,omitempty"`
}
type Issues struct {
	Vulnerabilities []Vulnerabilities `json:"vulnerabilities"`
	Licenses        []string          `json:"licenses"`
}
