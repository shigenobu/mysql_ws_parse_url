package test

import (
	"../func"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFull(t *testing.T) {
	input := "https://user:pass@www.sample.com:8080/path/to/site.html?k1=v1&k2=v2&u[]=1&u[]=2#top"

	output := _func.ParseUrl(input)
	fmt.Println(output)

	var result _func.Result
	json.Unmarshal([]byte(output), &result)
	assert.Equal(t, "https", result.Scheme)
	assert.Equal(t, "user", result.User)
	assert.Equal(t, "pass", result.Pass)
	assert.Equal(t, "www.sample.com", result.Host)
	assert.Equal(t, "8080", result.Port)
	assert.Equal(t, "/path/to/site.html", result.Path)
	assert.Equal(t, "k1=v1&k2=v2&u[]=1&u[]=2", result.Query)

	assert.Equal(t, "v1", result.Params["k1"][0])
	assert.Equal(t, "v2", result.Params["k2"][0])
	assert.Equal(t, "1", result.Params["u"][0])
	assert.Equal(t, "2", result.Params["u"][1])

	assert.Equal(t, "top", result.Fragment)
}

func TestParseMini(t *testing.T) {
	input := "https://www.sample.com"

	output := _func.ParseUrl(input)
	fmt.Println(output)
}

func TestOnlyPathAndQuery(t *testing.T) {
	input := "/path/to/site.html?k1=v1&k2=v2&u[]=1&u[]=2"

	output := _func.ParseUrl(input)
	fmt.Println(output)
}

func TestOnlyPath(t *testing.T) {
	input := "/path/to/site.html"

	output := _func.ParseUrl(input)
	fmt.Println(output)
}

func TestOnlyQuery(t *testing.T) {
	input := "?k1=v1&k2=v2&u[]=1&u[]=2"

	output := _func.ParseUrl(input)
	fmt.Println(output)
}

func TestOnlyQueryEncoded(t *testing.T) {
	input := "?t=&list%5B%5D=1&list%5B%5D=2&list%5B%5D=3&id=187"

	output := _func.ParseUrl(input)
	fmt.Println(output)

	var result _func.Result
	json.Unmarshal([]byte(output), &result)
	assert.Equal(t, "1", result.Params["list"][0])
	assert.Equal(t, "2", result.Params["list"][1])
	assert.Equal(t, "3", result.Params["list"][2])
}

func TestNumber(t *testing.T) {
	input := "0"

	output := _func.ParseUrl(input)
	if output == "" {
		fmt.Println("this is empty")
	}
	fmt.Println(output)
}

func TestEmpty(t *testing.T) {
	input := ""

	output := _func.ParseUrl(input)
	if output == "" {
		fmt.Println("this is empty")
	}
	fmt.Println(output)
}
