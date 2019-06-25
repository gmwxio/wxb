package adlproc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type readast struct {
	File string `type:"arg" help:"adl ast file" predict:"files"`
}

func (et *readast) Run() error {
	by, err := ioutil.ReadFile(et.File)
	if err != nil {
		return err
	}
	var m Module
	err = json.Unmarshal(by, &m)
	if err != nil {
		return err
	}
	// fmt.Printf("%v\n", m)
	b2, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", string(b2))
	return nil
}

type jsonObj struct {
}

func (et *jsonObj) Run() error {
	x := JsonObj{
		JsonObjs: map[string]Json{"a": "b"},
	}
	// x := map[string]interface{}{
	// 	"a": "b",
	// }
	b2, err := json.MarshalIndent(x, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", string(b2))

	js, err := UnmarshalJSON([]byte(`{"a" : ["b", 1, 2, 3, truez] }`))
	fmt.Printf("%v\n", js)
	// err = json.Unmarshal(by, &m)
	if err != nil {
		return err
	}
	b2, err = json.MarshalIndent(js, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("js - %v\n", string(b2))

	return nil

}
