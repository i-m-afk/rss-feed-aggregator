package scraper

import "encoding/xml"

// A common RSS structure
type Rss struct {
	Extra       map[string]interface{} `xml:",any"`
	Title       string                 `xml:"title"`
	Link        string                 `xml:"link"`
	Description string                 `xml:"description"`
	PublishDate string                 `xml:"pubDate"`
	XMLName     xml.Name               `xml:"rss"`
	Text        string                 `xml:",chardata"`
	Version     string                 `xml:"version,attr"`
	Atom        string                 `xml:"atom,attr"`
	Channel     struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Description   string `xml:"description"`
		Generator     string `xml:"generator"`
		Language      string `xml:"language"`
		LastBuildDate string `xml:"lastBuildDate"`
		Item          []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			PubDate     string `xml:"pubDate"`
			Guid        string `xml:"guid"`
			Description string `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
}
