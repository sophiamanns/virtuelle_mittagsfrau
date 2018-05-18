//usr/bin/env go run $0 $@; exit $?

// Run `./fotothek < fotothek.xml` or `go run fotothek.go < fotothek.xml` to parse out interesting records from
// deutschefotothek.xml (OAI-DC). A OAI dump can be created with harvesters like https://github.com/miku/metha.
//
// A non-thumbnail version might be found under URLs like:
// https://fotothek.slub-dresden.de/fotos/df/hauptkatalog/0412000/df_hauptkatalog_0412193.jpg
// https://fotothek.slub-dresden.de/fotos/df/dat/0014000/df_dat_0014640_001.jpg
// https://fotothek.slub-dresden.de/fotos/wob/hei/0009000/wob_hei_0009013_001_001.jpg

// <Record><header status=""><identifier>oai::a8450::obj|70065238|df_hauptkatalog_0415609</identifier><datestamp>2016-04-15T18:49:17Z</datestamp><setSpec>fotos</setSpec></header><metadata>
// <oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/"
//            xmlns:europeana="http://europeana.eu/terms"
//            xmlns:edpm="http://www.stegmannsystems.com/edp-meta/2009"
//            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
//            xmlns:edp="http://www.stegmannsystems.com/edp/2009"
//            xmlns:dc="http://purl.org/dc/elements/1.1/"
//            xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
//    <dc:coverage>Dresden-Friedrichstadt</dc:coverage>
//    <dc:creator>Rous, André (Fotograf)</dc:creator>
//    <dc:title>Dresden-Friedrichstadt, Schäferstraße 58. Wohnhaus (um 1865)</dc:title>
//    <dc:subject>Architektur</dc:subject>
//    <dc:subject>Wohnhaus</dc:subject>
//    <dc:subject>Fotos</dc:subject>
//    <dc:date>1865</dc:date>
//    <dc:description>Dresden-Friedrichstadt, Schäferstraße 58. Wohnhaus (um 1865) Dresden-Friedrichstadt</dc:description>
//    <dc:format>image/jpeg</dc:format>
//    <dc:source>SLUB/Deutsche Fotothek</dc:source>
//    <dc:type>image</dc:type>
//    <europeana:provider>Deutsche Fotothek</europeana:provider>
//    <europeana:type>IMAGE</europeana:type>
//    <dc:identifier>http://www.deutschefotothek.de/obj70065238.html</dc:identifier>
//    <europeana:isShownAt>http://www.deutschefotothek.de/obj70065238.html</europeana:isShownAt>
//    <europeana:isShownBy>http://fotothek.slub-dresden.de/thumbs/df_hauptkatalog_0415609.jpg</europeana:isShownBy>
//    <europeana:object>http://fotothek.slub-dresden.de/thumbs/df_hauptkatalog_0415609.jpg</europeana:object>
//    <dc:language>de-DE</dc:language>
//    <dc:rights>SLUB / Deutsche Fotothek</dc:rights>
//    <europeana:country>DE</europeana:country>
//    <europeana:language>de</europeana:language>
// </oai_dc:dc></metadata><about></about></Record>

package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/miku/xmlstream"
)

// areasOfInterest are keywords to look for in various record fields.
var areasOfInterest = []string{
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

func (r *Record) Snippet() string {
	m := map[string]string{
		"obj":  strings.TrimSpace(r.Metadata.Dc.Object.Text),
		"id":   strings.TrimSpace(r.Metadata.Dc.Identifier.Text),
		"desc": strings.TrimSpace(r.Metadata.Dc.Identifier.Text),
	}
	vs := []string{}

	for _, v := range r.Metadata.Dc.Title {
		vs = append(vs, strings.TrimSpace(v.Text))
	}
	m["title"] = strings.Join(vs, ", ")
	vs = vs[:0]

	for _, v := range r.Metadata.Dc.Coverage {
		vs = append(vs, strings.TrimSpace(v.Text))
	}
	m["coverage"] = strings.Join(vs, ", ")
	vs = vs[:0]

	for _, v := range r.Metadata.Dc.Date {
		vs = append(vs, strings.TrimSpace(v.Text))
	}
	m["date"] = strings.Join(vs, ", ")
	vs = vs[:0]

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(m); err != nil {
		log.Fatal(err)
	}

	return buf.String()
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
	scanner := xmlstream.NewScanner(os.Stdin, new(Record))
	for scanner.Scan() {
		tag := scanner.Element()
		switch el := tag.(type) {
		case *Record:
			record := *el
		L:
			for _, v := range record.Metadata.Dc.Coverage {
				if matchAnyString(v.Text, areasOfInterest) {
					fmt.Printf(record.Snippet())
					break L
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scan failed: %v\n", err)
	}
}
