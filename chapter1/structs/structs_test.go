package structs

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetAThakur(t *testing.T) {
	aThakur := GetAThakur()
	aThakur.Describe()

	bThakur := &ThakurCopy{Age: 20, Name: "Aiden"}
	bThakur.Describe()
	bThakur.ChangeToHeadOfHousehold() // This is a value receiver
	bThakur.Describe()                // notice, did not change
	bThakur.ValueDescribe()           // Operate on copy of bThakur
	bThakur.GotMarried()
	bThakur.Describe() // notice, did not change value of bThakur

	cThakur := NewThakur("Angela", 20)
	cThakur.ThoughtIGotMarried() // Operate on cThakur
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
