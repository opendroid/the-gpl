package structs

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetAThakur(t *testing.T) {
	aThakur := GetAThakur()
	aThakur.Describe()

	bThakur := &Thakur{Age: 20, Name: "Aiden"}
	bThakur.Describe()
	ChangeThakur(bThakur)
	bThakur.Describe()
	bThakur.ValueDescribe() // Operate on copy of bThakur
	bThakur.GotMarried()
	bThakur.Describe()

	cThakur := NewThakur("Angela", 20)
	cThakur.ThoughtIGotMarried()
	cThakur.Describe()
	cThakur.GotMarried()
	cThakur.Describe()
}

func TestGetAThakur_Json(t *testing.T) {
	testString := `{"age": 40, "name": "Temp Name"}`

	var tmpThakur Thakur
	err := json.Unmarshal([]byte(testString), &tmpThakur)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tmpThakur)
}
