package logics

import (
	"bytes"
	"encoding/json"
	"os"
	"strconv"
)

type Entity struct {
	Id         int     `json:"id"`
	DataDriven bool    `json:"data_driven"`
	Modifier   float32 `json:"modifier"`
}

func (e *Entity) Serialize() string {
	repr, _ := json.Marshal(e)
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, repr, "", "	")
	return prettyJSON.String()
}

func (e *Entity) Save() {
	data := []byte(e.Serialize())
	err := os.WriteFile(e.makeFileName(), data, 0644)
	if err != nil {
	}
}

func (e *Entity) makeFileName() string {
	return strconv.Itoa(e.Id) + ".txt"
}
