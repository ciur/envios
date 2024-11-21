package parser

import (
	"testing"

	"github.com/ciur/enward/profiles"
)

func TestParseProfileLineBasicProfileName(t *testing.T) {
	actual_tokens, err := ParseProfileLine("[coco]\n\n")
	var tokens = []Token{
		{name: PROFILE_NAME, value: "coco"},
	}

	if err != "" {
		t.Errorf("Error while parsing profile line %v", err)
	}

	if len(actual_tokens) != 1 {
		t.Errorf("Expected: len(actual_tokens) == 1, got %d", len(actual_tokens))
	}

	for i, token := range tokens {
		if token != actual_tokens[i] {
			t.Errorf("%s != %s", token, actual_tokens[i])
		}
	}
}

func TestParseProfileLineProfileNameWithWhiteSpacesInFront(t *testing.T) {
	actual_tokens, err := ParseProfileLine("    	 [coco]  \n\n")
	var tokens = []Token{
		{name: PROFILE_NAME, value: "coco"},
	}

	if err != "" {
		t.Errorf("Error while parsing profile line %v", err)
	}

	if len(actual_tokens) != 1 {
		t.Errorf("Expected: len(actual_tokens) == 1, got %d", len(actual_tokens))
	}

	for i, token := range tokens {
		if token != actual_tokens[i] {
			t.Errorf("%s != %s", token, actual_tokens[i])
		}
	}
}

func TestParseProfileLineProfileNameWithWhiteSpacesAroundName(t *testing.T) {
	actual_tokens, err := ParseProfileLine("    	 [     some_name    ]  \n\n")
	var tokens = []Token{
		{name: PROFILE_NAME, value: "some_name"},
	}

	if err != "" {
		t.Errorf("Error while parsing profile line %v", err)
	}

	if len(actual_tokens) != 1 {
		t.Errorf("Expected: len(actual_tokens) == 1, got %d", len(actual_tokens))
	}

	for i, token := range tokens {
		if token != actual_tokens[i] {
			t.Errorf("%s != %s", token, actual_tokens[i])
		}
	}
}

func TestParseProfileLine1(t *testing.T) {
	actual_tokens, err := ParseProfileLine("    	 [     some_name:common    ]  \n\n")
	var tokens = []Token{
		{name: PROFILE_NAME, value: "some_name"},
		{name: INHERITED_PROFILE_NAME, value: "common"},
	}

	if err != "" {
		t.Errorf("Error while parsing profile line %v", err)
	}

	if len(actual_tokens) != 2 {
		t.Errorf("Expected: len(actual_tokens) == 2, got %d", len(actual_tokens))
	}

	for i, token := range tokens {
		if token != actual_tokens[i] {
			t.Errorf("%s != %s", token, actual_tokens[i])
		}
	}
}

func TestParseProfileLine2(t *testing.T) {
	actual_tokens, err := ParseProfileLine("    	 [     some_name   :   common    ]  \n\n")
	var tokens = []Token{
		{name: PROFILE_NAME, value: "some_name"},
		{name: INHERITED_PROFILE_NAME, value: "common"},
	}

	if err != "" {
		t.Errorf("Error while parsing profile line %v", err)
	}

	if len(actual_tokens) != 2 {
		t.Errorf("Expected: len(actual_tokens) == 2, got %d", len(actual_tokens))
	}

	for i, token := range tokens {
		if token != actual_tokens[i] {
			t.Errorf("%s != %s", token, actual_tokens[i])
		}
	}
}

func TestParseProfileLine3(t *testing.T) {
	actual_tokens, err := ParseProfileLine("    	 [     some_name   :   common:default   ]  \n\n")
	var tokens = []Token{
		{name: PROFILE_NAME, value: "some_name"},
		{name: INHERITED_PROFILE_NAME, value: "common"},
		{name: DEFAULT_SWITCH, value: "default"},
	}

	if err != "" {
		t.Errorf("Error while parsing profile line %v", err)
	}

	if len(actual_tokens) != 3 {
		t.Errorf("Expected: len(actual_tokens) == 3, got %d", len(actual_tokens))
	}

	for i, token := range tokens {
		if token != actual_tokens[i] {
			t.Errorf("%s != %s", token, actual_tokens[i])
		}
	}
}

func TestParseProfileLine4(t *testing.T) {
	_, err := ParseProfileLine("    	 [     some_name   :   :default   ]  \n\n")

	if err != "inherited profile name cannot be empty" {
		t.Error("Parsing error missed")
	}
}

func TestParseProfileLine5(t *testing.T) {
	_, err := ParseProfileLine("[   :   :default   ]")

	if err != "profile name cannot be empty" {
		t.Error("Parsing error missed")
	}
}

func TestParseProfilNameContent1(t *testing.T) {
	actual_tokens, err := ParseProfileContent(" test:common")
	var tokens = []Token{
		{name: PROFILE_NAME, value: "test"},
		{name: INHERITED_PROFILE_NAME, value: string("common")},
	}

	if err != "" {
		t.Errorf("Unexpected error %v", err)
	}

	if len(actual_tokens) != 2 {
		t.Errorf("Expected: len(actual_tokens) == 2, got %d", len(actual_tokens))
	}

	for i, token := range tokens {
		if token != actual_tokens[i] {
			t.Errorf("%s != %s", token, actual_tokens[i])
		}
	}
}

func TestParseProfilNameContent2(t *testing.T) {
	actual_tokens, err := ParseProfileContent(" test:common:default")

	var tokens = []Token{
		{name: PROFILE_NAME, value: "test"},
		{name: INHERITED_PROFILE_NAME, value: string("common")},
		{name: DEFAULT_SWITCH, value: string("default")},
	}

	if len(actual_tokens) != 3 {
		t.Errorf("Expected: len(actual_tokens) == 3, got %d", len(actual_tokens))
	}

	if err != "" {
		t.Errorf("Unexpected error %v", err)
	}

	for i, token := range tokens {
		if token != actual_tokens[i] {
			t.Errorf("%s != %s", token, actual_tokens[i])
		}
	}
}

func TestParseProfilNameContent3(t *testing.T) {
	_, err := ParseProfileContent(" :")

	if err != "profile name cannot be empty" {
		t.Error("Empty profile problem was not reported")
	}
}

func TestParseProfilNameContent4(t *testing.T) {
	_, err := ParseProfileContent(" :comm")

	if err != "profile name cannot be empty" {
		t.Error("Empty profile problem was not reported")
	}
}

func TestParseProfilNameContent5(t *testing.T) {
	_, err := ParseProfileContent(" some::def")

	if err != "inherited profile name cannot be empty" {
		t.Error("Empty profile problem was not reported")
	}
}

func TestParseProfilNameContent6(t *testing.T) {
	_, err := ParseProfileContent(" some::")

	if err != "inherited profile name cannot be empty" {
		t.Error("Empty profile problem was not reported")
	}
}

func TestParseVariableLine1(t *testing.T) {
	actual_tokens := ParseVariableLine("  x = y  \n\n")
	var tokens = []Token{
		{name: VAR_NAME, value: string("x")},
		{name: VAR_VALUE, value: string("y")},
	}

	if len(actual_tokens) != 2 {
		t.Errorf("Expected: len(actual_tokens) == 2, got %d", len(actual_tokens))
	}

	for i, token := range tokens {
		if token != actual_tokens[i] {
			t.Errorf("%s != %s", token, actual_tokens[i])
		}
	}
}

func TestParseVariableLine2(t *testing.T) {
	actual_tokens := ParseVariableLine("  PAPER_URL = http://paper  \n\n")
	var tokens = []Token{
		{name: VAR_NAME, value: string("PAPER_URL")},
		{name: VAR_VALUE, value: string("http://paper")},
	}

	if len(actual_tokens) != 2 {
		t.Errorf("Expected: len(actual_tokens) == 2, got %d", len(actual_tokens))
	}

	for i, token := range tokens {
		if token != actual_tokens[i] {
			t.Errorf("%s != %s", token, actual_tokens[i])
		}
	}
}

func TestBuildProfiles1(t *testing.T) {
	var tokens = []Token{
		{name: PROFILE_NAME, value: "test"},
		{name: VAR_NAME, value: string("x")},
		{name: VAR_VALUE, value: string("y")},
	}

	actual_profiles := BuildProfiles(tokens)

	if len(actual_profiles) != 1 {
		t.Errorf("Expected: len(actual_profiles) == 1, got %d", len(actual_profiles))
	}

	if actual_profiles[0].Name != "test" {
		t.Errorf("expected profile name == 'test', got %s", actual_profiles[0].Name)
	}
	if actual_profiles[0].Variables[0].Name != "x" {
		t.Errorf("expected variable == 'x', got %s", actual_profiles[0].Variables[0].Name)
	}
	if actual_profiles[0].Variables[0].Value != "y" {
		t.Errorf("expected value == 'y', got %s", actual_profiles[0].Variables[0].Value)
	}
}

func TestBuildProfiles2(t *testing.T) {
	var tokens = []Token{
		{name: PROFILE_NAME, value: "test"},
		{name: VAR_NAME, value: string("x")},
		{name: VAR_VALUE, value: string("y")},
		{name: VAR_NAME, value: string("a")},
		{name: VAR_VALUE, value: string("b")},
	}

	actual_profiles := BuildProfiles(tokens)

	if len(actual_profiles) != 1 {
		t.Errorf("Expected: len(actual_profiles) == 1, got %d", len(actual_profiles))
	}

	if actual_profiles[0].Name != "test" {
		t.Errorf("expected profile name == 'test', got %s", actual_profiles[0].Name)
	}
	if actual_profiles[0].Variables[0].Name != "x" {
		t.Errorf("expected variable == 'x', got %s", actual_profiles[0].Variables[0].Name)
	}
	if actual_profiles[0].Variables[0].Value != "y" {
		t.Errorf("expected value == 'y', got %s", actual_profiles[0].Variables[0].Value)
	}
	if actual_profiles[0].Variables[1].Name != "a" {
		t.Errorf("expected variable == 'a', got %s", actual_profiles[0].Variables[1].Name)
	}
	if actual_profiles[0].Variables[1].Value != "b" {
		t.Errorf("expected value == 'b', got %s", actual_profiles[0].Variables[1].Value)
	}
}

func TestBuildProfiles3(t *testing.T) {
	/*
	* Read two empty profiles (i.e. profiles without any variables)
	 */
	var tokens = []Token{
		{name: PROFILE_NAME, value: "test1"},
		{name: PROFILE_NAME, value: "test2"},
	}

	actual_profiles := BuildProfiles(tokens)

	if len(actual_profiles) != 2 {
		t.Errorf("Expected: len(actual_profiles) == 2, got %d", len(actual_profiles))
	}

	if actual_profiles[0].Name != "test1" {
		t.Errorf("expected profile name == 'test1', got %s", actual_profiles[0].Name)
	}

	if actual_profiles[1].Name != "test2" {
		t.Errorf("expected profile name == 'test2', got %s", actual_profiles[0].Name)
	}

}

func TestConfigFile2(t *testing.T) {
	testFile := "test_data/config2.ini"
	actual_profiles, error := LoadConfig("test_data/config2.ini")
	expected_profile1 := profiles.Profile{
		Name: "test1",
		Variables: []profiles.ProfileVariable{
			{Name: "x", Value: "1"},
			{Name: "y", Value: "2"},
		},
	}
	expected_profile2 := profiles.Profile{
		Name: "test2",
	}

	if error != "" {
		t.Errorf("Error while opening %s", testFile)
	}

	if len(actual_profiles) != 2 {
		t.Errorf("Expecting 2 profiles, got %d\n", len(actual_profiles))
	}

	if actual_profiles[0].NotEqual(expected_profile1) {
		t.Errorf("Profile 1 expected to be %v, got %v", expected_profile1, actual_profiles[0])
	}

	if actual_profiles[1].NotEqual(expected_profile2) {
		t.Errorf("Profile 2 expected to be %v, got %v", expected_profile2, actual_profiles[1])
	}
}
