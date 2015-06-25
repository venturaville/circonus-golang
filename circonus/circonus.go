package circonus

import (
//        "encoding/json"
  "fmt"
//  "strings"
//  "time"
  "io/ioutil"
  "crypto/tls"
  "net/http"
)

type Circonus struct {
  AppName string "curl"
  ApiServer string "api.circonus.com"
  ApiToken string ""
}

// get_url takes a collection string, and returns a URL to query
func (c *Circonus) get_url(collection string) string {
  url := fmt.Sprintf("https://%s/v2/%s",c.ApiServer,collection)
  fmt.Println("URL:",url)
  return url
}

// do_request takes:
//   A method string ("GET", "POST","PUT", "DELETE")
//   The action to take ("list","add","get","delete")
//   The collection ("broker","check","graph",etc.)
// It returns the prepared response from the request
func (c *Circonus) do_request(method string, action string, collection string) []byte {
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
  fmt.Println("response Status:", resp.Status)
  fmt.Println("response Headers:", resp.Header)
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println("response Body:", string(body))
  return body
}

func (c *Circonus) ListAccount() []byte { return c.do_request("GET","list","ccount") }
func (c *Circonus) ListAnnotation() []byte { return c.do_request("GET","list","annotation") }
func (c *Circonus) ListBroker() []byte { return c.do_request("GET","list","broker") }
func (c *Circonus) ListCheck() []byte { return c.do_request("GET","list","check") }
func (c *Circonus) ListCheckBundle() []byte { return c.do_request("GET","list","check_bundle") }
func (c *Circonus) ListContactGroup() []byte { return c.do_request("GET","list","contact_group") }
func (c *Circonus) ListGraph() []byte { return c.do_request("GET","list","graph") }
func (c *Circonus) ListMetricCluster() []byte { return c.do_request("GET","list","metric_clusterj") }
func (c *Circonus) ListRuleSet() []byte { return c.do_request("GET","list","rule_set") }
func (c *Circonus) ListTag() []byte { return c.do_request("GET","list","tag") }
func (c *Circonus) ListTemplate() []byte { return c.do_request("GET","list","template") }
func (c *Circonus) ListUser() []byte { return c.do_request("GET","list","user") }
func (c *Circonus) ListWorksheet() []byte { return c.do_request("GET","list","worksheet") }
