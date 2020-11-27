package alexa

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const baseApiUrl = "http://data.alexa.com/data?cli=10&url="

// GlobalRank get Global Rank of website in Alexa
// It returns global rank in string type and any rank error encountered.
func GlobalRank(url string) (string, error) {
	element := get(url, "POPULARITY")
	if len(element.Attr) >= 2 {
		return element.Attr[1].Value, nil
	} else {
		return "No rank", nil
	}
}

// GlobalRank get Global Rank of website in Alexa
// It returns country rank in string, country name in string, country code in string and any rank error encountered.
func CountryRank(url string) (string, string, string, error) {
	element := get(url, "COUNTRY")
	if len(element.Attr) >= 3 {
		return element.Attr[2].Value, element.Attr[1].Value, element.Attr[0].Value, nil
	} else {
		return "", "", "", errors.New("no country rank")
	}
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
		log.Println(err)
	}
}
