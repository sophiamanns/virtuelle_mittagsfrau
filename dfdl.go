// dfdl will try to harvest image material from DF. Will download all images into cache/df.
//
// $ open $(go run dfdl.go < data/fotothek.jsonl 2> /dev/null | shuf -n 1)
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/dchest/safefile"
	log "github.com/sirupsen/logrus"
)

var (
	userAgent        = "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)"
	imageDownloadDir = "./cache/df"
)

// Doc is a subset of record information.
type Doc struct {
	Coverage string `json:"coverage"`
	Date     string `json:"date"`
	Desc     string `json:"desc"`
	Id       string `json:"id"`
	Obj      string `json:"obj"`
	Title    string `json:"title"`
}

// findImageURL turns a thumbnail URL Like
// http://fotothek.slub-dresden.de/thumbs/df_bika029_0000150_motiv.jpg into a
// non-thumb version like
// https://fotothek.slub-dresden.de/fotos/bz/si/0054000/bz_si_0054194.jpg.
func findImageURL(u *url.URL) (string, error) {
	filename := filepath.Base(u.Path)
	parts := strings.Split(filename, "_")
	if len(parts) < 3 {
		return "", fmt.Errorf("unexpected name: %s", filename)
	}
	source, collection, id := parts[0], parts[1], parts[2]
	nameext := strings.Split(id, ".")
	if len(nameext) != 2 {
		return "", fmt.Errorf("unexpected name: %s", id)
	}
	nameonly := nameext[0]
	if len(nameonly) != 7 {
		return "", fmt.Errorf("some other length: %s", nameonly)
	}
	shard := nameonly[:4] + "000"
	return fmt.Sprintf("https://fotothek.slub-dresden.de/fotos/%s/%s/%s/%s", source, collection, shard, filename), nil
}

func downloadFile(link, filepath string) error {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		out, err := safefile.Create(filepath, 0644)
		if err != nil {
			return err
		}
		defer out.Close()

		req, err := http.NewRequest("GET", link, nil)
		if err != nil {
			return err
		}
		req.Header.Set("User-Agent", userAgent)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("bad status at %s: %s", link, resp.Status)
		}
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return err
		}
		out.Commit()
	}
	return nil
}

func main() {
	if err := os.MkdirAll(imageDownloadDir, 0755); err != nil {
		log.Fatal(err)
	}
	br := bufio.NewReader(os.Stdin)
	for {
		b, err := br.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var doc Doc
		if err := json.Unmarshal(b, &doc); err != nil {
			log.Fatal(err)
		}
		u, err := url.Parse(doc.Obj)
		if err != nil {
			log.Println(err)
			continue
		}
		imageURL, err := findImageURL(u)
		if err != nil {
			log.Println(err)
			continue
		}
		dst := filepath.Join(imageDownloadDir, filepath.Base(imageURL))
		if err := downloadFile(imageURL, dst); err != nil {
			log.Fatal(err)
		}
		log.Printf("ok: %v", imageURL)
	}
}
