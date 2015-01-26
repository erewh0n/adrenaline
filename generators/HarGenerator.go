package hargenerator

import (
	"bytes"
	"har"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HarGenerator struct {
	har   *har.HARLog
	host  string
	index int
}

func Create(har *har.HARLog, host string) *HarGenerator {
	return &HarGenerator{
		har:   har,
		host:  host,
		index: 0,
	}
}

func (gen HarGenerator) Generate(res *http.Response) (*http.Request, error) {
	if gen.index > (len(gen.har.Log.Entries) - 1) {
		gen.index = 0
	}
	var harRequest = gen.har.Log.Entries[gen.index].Request
	gen.index++

	url, err := url.Parse(harRequest.Url)
	if err != nil {
		return nil, err
	}
	url.Host = gen.host

	request := &http.Request{}
	request.Method = harRequest.Method
	request.URL = url
	request.Body = ioutil.NopCloser(bytes.NewBufferString(harRequest.PostData.Text))
	request.Header = http.Header{}

	for i := 0; i < len(harRequest.Headers); i++ {
		headerName := harRequest.Headers[i].Name
		if headerName == "authorization" {
			request.Header.Add("authorization", harRequest.Headers[i].Value)
		}
	}

	return request, nil
}
