package parser

import (
	"testing"
)

func TestParseProfileLineBasicProfileName(t *testing.T) {
	actual_tokens, err := ParseProfileLine("[coco]\n\n")
	var tokens = []Token{
		{name: LEFT_BRA, value: string("[")},
		{name: PROFILE_NAME, value: "coco"},
		{name: RIGHT_BRA, value: string("]")},
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

func TestParseProfileLineProfileNameWithWhiteSpacesInFront(t *testing.T) {
	actual_tokens, err := ParseProfileLine("    	 [coco]  \n\n")
	var tokens = []Token{
		{name: LEFT_BRA, value: string("[")},
		{name: PROFILE_NAME, value: "coco"},
		{name: RIGHT_BRA, value: string("]")},
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

func TestParseProfileLineProfileNameWithWhiteSpacesAroundName(t *testing.T) {
	actual_tokens, err := ParseProfileLine("    	 [     some_name    ]  \n\n")
	var tokens = []Token{
		{name: LEFT_BRA, value: string("[")},
		{name: PROFILE_NAME, value: "some_name"},
		{name: RIGHT_BRA, value: string("]")},
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

func TestParseProfileLine1(t *testing.T) {
	actual_tokens, err := ParseProfileLine("    	 [     some_name:common    ]  \n\n")
	var tokens = []Token{
		{name: LEFT_BRA, value: string("[")},
		{name: PROFILE_NAME, value: "some_name"},
		{name: INHERITED_PROFILE_NAME, value: "common"},
		{name: RIGHT_BRA, value: string("]")},
	}

	if err != "" {
		t.Errorf("Error while parsing profile line %v", err)
	}

	if len(actual_tokens) != 4 {
		t.Errorf("Expected: len(actual_tokens) == 4, got %d", len(actual_tokens))
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
		{name: LEFT_BRA, value: string("[")},
		{name: PROFILE_NAME, value: "some_name"},
		{name: INHERITED_PROFILE_NAME, value: "common"},
		{name: RIGHT_BRA, value: string("]")},
	}

	if err != "" {
		t.Errorf("Error while parsing profile line %v", err)
	}

	if len(actual_tokens) != 4 {
		t.Errorf("Expected: len(actual_tokens) == 4, got %d", len(actual_tokens))
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
		{name: LEFT_BRA, value: string("[")},
		{name: PROFILE_NAME, value: "some_name"},
		{name: INHERITED_PROFILE_NAME, value: "common"},
		{name: DEFAULT_SWITCH, value: "default"},
		{name: RIGHT_BRA, value: string("]")},
	}

	if err != "" {
		t.Errorf("Error while parsing profile line %v", err)
	}

	if len(actual_tokens) != 5 {
		t.Errorf("Expected: len(actual_tokens) == 5, got %d", len(actual_tokens))
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
		{name: EQUAL, value: "="},
		{name: VAR_VALUE, value: string("y")},
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

func TestParseVariableLine2(t *testing.T) {
	actual_tokens := ParseVariableLine("  PAPER_URL = http://paper  \n\n")
	var tokens = []Token{
		{name: VAR_NAME, value: string("PAPER_URL")},
		{name: EQUAL, value: "="},
		{name: VAR_VALUE, value: string("http://paper")},
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

func TestBuildProfiles1(t *testing.T) {
	var tokens = []Token{
		{name: LEFT_BRA, value: string("[")},
		{name: PROFILE_NAME, value: "test"},
		{name: RIGHT_BRA, value: string("]")},
		{name: VAR_NAME, value: string("x")},
		{name: EQUAL, value: string("=")},
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
		{name: LEFT_BRA, value: string("[")},
		{name: PROFILE_NAME, value: "test"},
		{name: RIGHT_BRA, value: string("]")},
		{name: VAR_NAME, value: string("x")},
		{name: EQUAL, value: string("=")},
		{name: VAR_VALUE, value: string("y")},
		{name: VAR_NAME, value: string("a")},
		{name: EQUAL, value: string("=")},
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
