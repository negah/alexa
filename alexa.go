package alexa

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const baseApiUrl = "http://data.alexa.com/data?cli=10&url="

func GlobalRank(url string) string {
	element := get(url, "POPULARITY")
	return element.Attr[1].Value
}

func CountryRank(url string) (string, string, string) {
	element := get(url, "COUNTRY")
	return element.Attr[2].Value, element.Attr[1].Value, element.Attr[0].Value
}

func get(url string, tag string) xml.StartElement {
	// Get Data from alexa api
	data := getData(baseApiUrl + url)

	// Get element
	element := getElement(data, tag)

	return element
}

func getElement(data *xml.Decoder, elementName string) xml.StartElement {
	// loop till we get REACH token
	for {
		token, _ := data.Token()

		if token == nil {
			break
		}

		switch element := token.(type) {
		case xml.StartElement:
			if element.Name.Local == elementName {
				return element
			}
		}
	}
	return xml.StartElement{}
}

func getData(url string) *xml.Decoder {
	// Get Data from alexa api
	data, err := http.Get(url)

	checkError(err)

	defer data.Body.Close()

	// Parse fetched data
	parsed, err := ioutil.ReadAll(data.Body)

	checkError(err)

	// convert []byte to type io.Reader with strings.NewReader()
	decoder := xml.NewDecoder(strings.NewReader(string(parsed)))

	return decoder
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
