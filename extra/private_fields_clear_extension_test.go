package extra

import (
"fmt"
"github.com/qluvio/json-iterator"
"testing"
)

func TestClearExtensions(t *testing.T) {

	type testStruct struct {
		PublicField  string
		privateField string
	}


	myStruct1 := testStruct{PublicField: "public", privateField: "private"}
	expected := `{"PublicField":"public"}`
	actual, err := jsoniter.Marshal(myStruct1)
	if err != nil {
		t.Error(err)
	}
	if expected != string(actual) {
		t.Error(fmt.Sprintf("public-only JSON does not match - expected: '%s' actual:'%s'", expected, string(actual)))
	}

	SupportPrivateFields()

	type testStructAlias1 testStruct
	myStruct2 := testStructAlias1(myStruct1)

	expected = `{"PublicField":"public","privateField":"private"}`
	actual, err = jsoniter.Marshal(myStruct2)
	if err != nil {
		t.Error(err)
	}
	if expected != string(actual) {
		t.Error(fmt.Sprintf("public+private JSON does not match - expected: '%s' actual:'%s'", expected, string(actual)))
	}

	jsoniter.ClearExtensions()

	type testStructAlias2 testStruct
	myStruct3 := testStructAlias2(myStruct1)

	expected = `{"PublicField":"public"}`
	actual, err = jsoniter.Marshal(myStruct3)
	if err != nil {
		t.Error(err)
	}
	if expected != string(actual) {
		t.Error(fmt.Sprintf("public-only JSON after ClearExtensions() does not match - expected: '%s' actual:'%s'", expected, string(actual)))
	}

}

