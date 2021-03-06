package json

import (
  "bytes"
  "encoding/json"
  "net/http"
)
/*
"github.com/leonard7e/hookgit/error"
*/
type JsonMap map[string]interface{}

func ReadJson(r *http.Request) (*JsonMap) {
  buf := new(bytes.Buffer)
  buf.ReadFrom(r.Body)
  s := buf.Bytes()

  jm := new (JsonMap)
  json.Unmarshal(s, jm)
  // err := json.Unmarshal(s, jm)
  // error.AssertNoErr(err)

  return jm
}
