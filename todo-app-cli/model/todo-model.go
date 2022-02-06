package model

import (
   "fmt"
)

// Cryptoresponse is exported, it models the data we receive.
type Todo []struct {
   oid              string    `json:"_id"`
   name             string    `json:"name"`
   date             string    `json:"date"`
   completed        string    `json:"completed"`
}

//TextOutput is exported,it formats the data to plain text.
func (c Todo) TextOutput() string {
p := fmt.Sprintf(
  "Name: %s\nCompleted : %s\n_id: %s\ndate: %s\n",
  c[0].name, c[0].completed, c[0].oid, c[0].date)
  return p
}