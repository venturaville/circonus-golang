package circonus

import (
  "encoding/json"
  "os"
  "fmt"
//  "strings"
//  "time"
  "io/ioutil"
  "crypto/tls"
  "net/http"
)

const defaultAppName string = "curl"
const defaultApiServer string = "api.circonus.com"

type Circonus struct {
  AppName string
  ApiServer string
  ApiToken string
}

var defaultCirconus = Circonus{
  AppName: "",
  ApiServer: "",
  ApiToken: "",
}

// get_url takes a collection string, and returns a URL to query
func (c *Circonus) get_url(collection string) string {
  url := fmt.Sprintf("https://%s/v2/%s",c.ApiServer,collection)
//  fmt.Println("URL:",url)
  return url
}

// NewCirconus initializes our Circonus struct
func NewCirconus(c *Circonus) *Circonus {
  if c == nil {
    c = &defaultCirconus
  }

  appname := os.Getenv("CIRCONUS_APPNAME")
  apitoken := os.Getenv("CIRCONUS_APITOKEN")
  apiserver := os.Getenv("CIRCONUS_APISERVER")
  if (c.AppName == "") && (appname != "") { c.AppName = appname }
  if (c.ApiServer == "") && (apiserver != "") { c.ApiServer = apiserver }
  if (c.ApiToken == "") && (apitoken != "") { c.ApiToken = apitoken }
  if c.AppName == "" { c.AppName = defaultAppName }
  if c.ApiServer == "" { c.ApiServer = defaultApiServer }
//  fmt.Println("apitoken:", c.ApiToken)
//  fmt.Println("apiserver:", c.ApiServer)
//  fmt.Println("appname:", c.AppName)
  return c
}

// do_request takes:
//   A method string ("GET", "POST","PUT", "DELETE")
//   The action to take ("list","add","get","delete")
//   The collection ("broker","check","graph",etc.)
// It returns the prepared response from the request
func (c *Circonus) do_request(method string, action string, collection string) []interface{} {
  url := c.get_url(collection)
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Set("X-Circonus-Auth-Token", c.ApiToken)
  req.Header.Set("X-Circonus-App-Name", c.AppName)
  req.Header.Set("Accept", "application/json")
  req.Header.Set("Content-Type", "application/json")

  config := &tls.Config{InsecureSkipVerify: true}
  tr := &http.Transport{ TLSClientConfig: config }
  client := &http.Client{Transport: tr}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }

  defer resp.Body.Close()
//  fmt.Println("response Status:", resp.Status)
//  fmt.Println("response Headers:", resp.Header)
  body, _ := ioutil.ReadAll(resp.Body)
//  fmt.Println("response Body:", string(body))

  //var data map[string] interface{}
  var data [] interface{}
  json.Unmarshal([]byte(body), &data)
//  fmt.Println("%v", data)
//  entry := data[0]
//  fmt.Println("%v", entry)
//  fmt.Println("%v", entry.(map[string]interface{})["_name"])
  return data
}

func (c *Circonus) ListAccount() []interface{}{ return c.do_request("GET","list","ccount") }
func (c *Circonus) ListAnnotation() []interface{}{ return c.do_request("GET","list","annotation") }
func (c *Circonus) ListBroker() []interface{}{ return c.do_request("GET","list","broker") }
func (c *Circonus) ListCheck() []interface{}{ return c.do_request("GET","list","check") }
func (c *Circonus) ListCheckBundle() []interface{}{ return c.do_request("GET","list","check_bundle") }
func (c *Circonus) ListContactGroup() []interface{}{ return c.do_request("GET","list","contact_group") }
func (c *Circonus) ListGraph() []interface{}{ return c.do_request("GET","list","graph") }
func (c *Circonus) ListMetricCluster() []interface{}{ return c.do_request("GET","list","metric_clusterj") }
func (c *Circonus) ListRuleSet() []interface{}{ return c.do_request("GET","list","rule_set") }
func (c *Circonus) ListTag() []interface{}{ return c.do_request("GET","list","tag") }
func (c *Circonus) ListTemplate() []interface{}{ return c.do_request("GET","list","template") }
func (c *Circonus) ListUser() []interface{}{ return c.do_request("GET","list","user") }
func (c *Circonus) ListWorksheet() []interface{}{ return c.do_request("GET","list","worksheet") }

