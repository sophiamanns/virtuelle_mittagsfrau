// Run `go run fotothek.go` to parse out interesting records from
// deutschefotothek.xml (OAI-DC).
package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
	"strings"

	"github.com/miku/xmlstream"
)

// Record was generated 2018-05-18 16:00:31 by tir on hayiti.
type Record struct {
	XMLName xml.Name `xml:"Record"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Identifier struct {
			Text string `xml:",chardata"` // oai::a8450::obj|01e4efa5-...
		} `xml:"identifier"`
		Datestamp struct {
			Text string `xml:",chardata"` // 2015-10-01T15:21:32Z, 201...
		} `xml:"datestamp"`
		SetSpec []struct {
			Text string `xml:",chardata"` // fotos, fotos, fotos, foto...
		} `xml:"setSpec"`
	} `xml:"header"`
	Metadata struct {
		Text string `xml:",chardata"`
		Dc   struct {
			Text           string `xml:",chardata"`
			Dc             string `xml:"dc,attr"`
			Edpm           string `xml:"edpm,attr"`
			Xsi            string `xml:"xsi,attr"`
			OaiDc          string `xml:"oai_dc,attr"`
			Europeana      string `xml:"europeana,attr"`
			Edp            string `xml:"edp,attr"`
			SchemaLocation string `xml:"schemaLocation,attr"`
			Coverage       []struct {
				Text string `xml:",chardata"` // Leipzig, Leipzig, Leipzig...
			} `xml:"coverage"`
			Creator []struct {
				Text string `xml:",chardata"` // Weigt, Ernst (Fotograf), ...
			} `xml:"creator"`
			Title []struct {
				Text string `xml:",chardata"` // Bilderschließung in Vorb...
			} `xml:"title"`
			Subject []struct {
				Text string `xml:",chardata"` // Fotografie, Foto, XXX, Fo...
			} `xml:"subject"`
			Date []struct {
				Text string `xml:",chardata"` // 1935, 1935, 1935, 1935, 1...
			} `xml:"date"`
			Description struct {
				Text string `xml:",chardata"` // XXX, Thekla, Bilderschlie...
			} `xml:"description"`
			Format struct {
				Text string `xml:",chardata"` // image/jpeg, image/jpeg, i...
			} `xml:"format"`
			Source struct {
				Text string `xml:",chardata"` // SLUB/Deutsche Fotothek, S...
			} `xml:"source"`
			Type struct {
				Text string `xml:",chardata"` // image, IMAGE, image, IMAG...
			} `xml:"type"`
			Provider struct {
				Text string `xml:",chardata"` // Deutsche Fotothek, Deutsc...
			} `xml:"provider"`
			Identifier struct {
				Text string `xml:",chardata"` // http://www.deutschefototh...
			} `xml:"identifier"`
			IsShownAt struct {
				Text string `xml:",chardata"` // http://www.deutschefototh...
			} `xml:"isShownAt"`
			IsShownBy struct {
				Text string `xml:",chardata"` // http://fotothek.slub-dres...
			} `xml:"isShownBy"`
			Object struct {
				Text string `xml:",chardata"` // http://fotothek.slub-dres...
			} `xml:"object"`
			Language struct {
				Text string `xml:",chardata"` // de-DE, de, de-DE, de, de-...
			} `xml:"language"`
			Rights struct {
				Text string `xml:",chardata"` // SLUB / Deutsche Fotothek,...
			} `xml:"rights"`
			Country struct {
				Text string `xml:",chardata"` // DE, DE, DE, DE, DE, DE, D...
			} `xml:"country"`
			Contributor []struct {
				Text string `xml:",chardata"` // Oppenheim, Martin Wilhelm...
			} `xml:"contributor"`
			HasObject struct {
				Text string `xml:",chardata"` // false, false, false, fals...
			} `xml:"hasObject"`
		} `xml:"dc"`
	} `xml:"metadata"`
	About struct {
		Text string `xml:",chardata"`
	} `xml:"about"`
}

func matchAnyString(haystack string, needles []string) bool {
	for _, n := range needles {
		if strings.Contains(haystack, n) {
			return true
		}
	}
	return false
}

func main() {
	keywords := []string{
		"Sabrodt",
		"Seidewinkel",
		"Hoyerswerda",
		"Parcow",
		"Sprjowje",
		"Kreckwitz",
		"Rowno",
		"Niederlausitz",
		"Bluno",
		"Lauske bei Hochkirch",
		"Geierswalde",
		"Oberlausitz",
	}

	encoder := json.NewEncoder(os.Stdout)
	scanner := xmlstream.NewScanner(os.Stdin, new(Record))

	for scanner.Scan() {
		tag := scanner.Element()
		switch el := tag.(type) {
		case *Record:
			record := *el
		L:
			for _, v := range record.Metadata.Dc.Coverage {
				if matchAnyString(v.Text, keywords) {
					if err := encoder.Encode(record); err != nil {
						log.Fatal(err)
					}
					break L
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scan failed: %v\n", err)
	}
}
