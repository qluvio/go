package extra

import (
	"github.com/qluvio/json-iterator"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_lower_case_with_underscores(t *testing.T) {
	should := require.New(t)
	should.Equal("hello_world", LowerCaseWithUnderscores("helloWorld"))
	should.Equal("hello_world", LowerCaseWithUnderscores("HelloWorld"))
	SetNamingStrategy(LowerCaseWithUnderscores)
	output, err := jsoniter.Marshal(struct {
		UserName      string
		FirstLanguage string
	}{
		UserName:      "taowen",
		FirstLanguage: "Chinese",
	})
	should.Nil(err)
	should.Equal(`{"user_name":"taowen","first_language":"Chinese"}`, string(output))
}

func Test_set_naming_strategy_with_overrides(t *testing.T) {
	should := require.New(t)
	SetNamingStrategy(LowerCaseWithUnderscores)
	output, err := jsoniter.Marshal(struct {
		UserName      string `json:"UserName"`
		FirstLanguage string
	}{
		UserName:      "taowen",
		FirstLanguage: "Chinese",
	})
	should.Nil(err)
	should.Equal(`{"UserName":"taowen","first_language":"Chinese"}`, string(output))
}

func Test_set_naming_strategy_with_omitempty(t *testing.T) {
	should := require.New(t)
	SetNamingStrategy(LowerCaseWithUnderscores)
	output, err := jsoniter.Marshal(struct {
		UserName      string
		FirstLanguage string `json:",omitempty"`
	}{
		UserName: "taowen",
	})
	should.Nil(err)
	should.Equal(`{"user_name":"taowen"}`, string(output))
}

func Test_set_naming_strategy_public_private(t *testing.T) {
	should := require.New(t)
	jsoniter.ResetAll()
	jsoniter.ClearExtensions()
	SetNamingStrategy(LowerCaseWithUnderscores)
	SupportPrivateFields()
	output, err := jsoniter.Marshal(struct {
		PublicField  string
		privateField string
	}{
		PublicField:  "public",
		privateField: "private",
	})
	should.Nil(err)
	should.Equal(`{"public_field":"public","private_field":"private"}`, string(output))
}


func Test_set_naming_strategy_public_only(t *testing.T) {
	should := require.New(t)
	jsoniter.ResetAll()
	SetNamingStrategy(LowerCaseWithUnderscores)
	output, err := jsoniter.Marshal(struct {
		PublicField  string
		privateField string
	}{
		PublicField:  "public",
		privateField: "private",
	})
	should.Nil(err)
	should.Equal(`{"public_field":"public"}`, string(output))
}

func Test_single_proc(t *testing.T) {
	should := require.New(t)
	should.Equal("hello_world", LowerCaseWithUnderscores("helloWorld"))
	should.Equal("hello_world", LowerCaseWithUnderscores("HelloWorld"))
	SetNamingStrategy(LowerCaseWithUnderscores)
	output, err := jsoniter.Marshal(struct {
		UserName      string
		FirstLanguage string
	}{
		UserName:      "taowen",
		FirstLanguage: "Chinese",
	})
	should.Nil(err)
	should.Equal(`{"user_name":"taowen","first_language":"Chinese"}`, string(output))

	output, err = jsoniter.Marshal(struct {
		UserName      string `json:"UserName"`
		FirstLanguage string
	}{
		UserName:      "taowen",
		FirstLanguage: "Chinese",
	})
	should.Nil(err)
	should.Equal(`{"UserName":"taowen","first_language":"Chinese"}`, string(output))

	jsoniter.ResetAll()
	SetNamingStrategy(LowerCaseWithUnderscores)
	output, err = jsoniter.Marshal(struct {
		UserName      string
		FirstLanguage string `json:",omitempty"`
	}{
		UserName: "taowen",
	})
	should.Nil(err)
	should.Equal(`{"user_name":"taowen"}`, string(output))



	output, err = jsoniter.Marshal(struct {
		PublicField  string
		privateField string
	}{
		PublicField:  "public",
		privateField: "private",
	})
	should.Nil(err)
	should.Equal(`{"public_field":"public"}`, string(output))

	jsoniter.ResetAll()
	SupportPrivateFields()
	SetNamingStrategy(LowerCaseWithUnderscores)
	output, err = jsoniter.Marshal(struct {
		PublicField  string
		privateField string
	}{
		PublicField:  "public",
		privateField: "private",
	})
	should.Nil(err)
	should.Equal(`{"public_field":"public","private_field":"private"}`, string(output))

}
