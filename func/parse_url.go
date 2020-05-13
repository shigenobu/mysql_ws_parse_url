/*
 parse url
 */
package _func

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strconv"
)

type Result struct {
	Scheme string `json:"scheme,omitempty"`
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
	User string `json:"user,omitempty"`
	Pass string `json:"pass,omitempty"`
	Path string `json:"path,omitempty"`
	Query string `json:"query,omitempty"`
	Params map[string][]string `json:"params,omitempty"`
	Fragment string `json:"fragment,omitempty"`
}

func ParseUrl(uri string) string {
	if uri == "" {
		return "{}"
	}
	if _, err := strconv.Atoi(uri); err == nil {
		return "{}"
	}
	
	u, err := url.Parse(uri)
	if err != nil {
		return "{}"
	}

	// pass
	var pass string
	p, ret := u.User.Password()
	if ret {
		pass = p
	}

	result := Result{
		Scheme:   u.Scheme,
		Host:     u.Hostname(),
		Port:     u.Port(),
		User:     u.User.Username(),
		Pass:     pass,
		Path:     u.Path,
		Query:    u.RawQuery,
		Params:   nil,
		Fragment: u.Fragment,
	}
	if len(u.Query()) > 0 {
		result.Params = make(map[string][]string, len(u.Query()))
		for key, values := range u.Query() {
			var params []string
			for _, v := range values {
				params = append(params, v)
			}
			result.Params[key] = params
		}
	}

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err = encoder.Encode(result)
	if err != nil {
		return "{}"
	}
	return string(buffer.Bytes())
}