package parser

import (
	"envios/profiles"
	"strings"
)

const (
	LEFT_BRA               = "LEFT_BRA"
	RIGHT_BRA              = "RIGHT_BRA"
	PROFILE_NAME           = "PROFILE_NAME"
	INHERITED_PROFILE_NAME = "INHERITED_PROFILE_NAME"
	DEFAULT_SWITCH         = "DEFAULT_SWITCH"
	VAR_NAME               = "VAR_NAME"
	EQUAL                  = "EQUAL"
	VAR_VALUE              = "VAR_VALUE"
)

type Token struct {
	name  string
	value string
}

func ParseProfileLine(line string) ([]Token, string) {
	var tokens []Token
	var profile_name []string
	var gather_profile_name = false
	var error string

	for i := 0; i < len(line); i++ {
		if line[i] == '[' {
			gather_profile_name = true
			token := Token{name: LEFT_BRA, value: string('[')}
			tokens = append(tokens, token)
			continue
		}

		if line[i] == ']' {
			if len(profile_name) > 0 {
				profile_content := strings.TrimSpace(strings.Join(profile_name, ""))
				profile_tokens, err := ParseProfileContent(profile_content)

				if err != "" {
					return tokens, err
				}

				tokens = append(tokens, profile_tokens...)

			}
			gather_profile_name = false
			token := Token{name: RIGHT_BRA, value: string(']')}
			tokens = append(tokens, token)
			break
		}

		if gather_profile_name {
			profile_name = append(profile_name, string(line[i]))
		}
	}

	return tokens, error
}

func ParseProfileContent(content_line string) ([]Token, string) {
	var tokens []Token
	var error string

	parts := strings.Split(content_line, ":")

	if len(parts) == 1 {
		value := strings.TrimSpace(parts[0])
		token := Token{name: PROFILE_NAME, value: value}
		tokens = append(tokens, token)
	}

	if len(parts) == 2 {
		value := strings.TrimSpace(parts[0])

		if len(value) == 0 {
			error = "profile name cannot be empty"
			return tokens, error
		}

		token := Token{name: PROFILE_NAME, value: value}
		tokens = append(tokens, token)

		value = strings.TrimSpace(parts[1])

		if len(value) == 0 {
			error = "inherited profile name cannot be empty"
			return tokens, error
		}

		token = Token{name: INHERITED_PROFILE_NAME, value: value}
		tokens = append(tokens, token)
	}

	if len(parts) == 3 {
		value := strings.TrimSpace(parts[0])

		if len(value) == 0 {
			error = "profile name cannot be empty"
			return tokens, error
		}

		token := Token{name: PROFILE_NAME, value: value}
		tokens = append(tokens, token)

		value = strings.TrimSpace(parts[1])

		if len(value) == 0 {
			error = "inherited profile name cannot be empty"
			return tokens, error
		}

		token = Token{name: INHERITED_PROFILE_NAME, value: value}
		tokens = append(tokens, token)

		value = strings.TrimSpace(parts[2])

		if len(value) == 0 {
			error = "default switch cannot be empty"
			return tokens, error
		}

		token = Token{name: DEFAULT_SWITCH, value: value}
		tokens = append(tokens, token)
	}

	return tokens, error
}

func ParseVariableLine(line string) []Token {
	var tokens []Token

	parts := strings.Split(line, "=")

	if len(parts) == 2 {
		value := strings.TrimSpace(parts[0])
		token := Token{name: VAR_NAME, value: value}
		tokens = append(tokens, token)

		token = Token{name: EQUAL, value: "="}
		tokens = append(tokens, token)

		value = strings.TrimSpace(parts[1])
		token = Token{name: VAR_VALUE, value: value}
		tokens = append(tokens, token)
	}

	return tokens
}

func BuildProfiles(tokens []Token) []profiles.Profile {
	var items []profiles.Profile
	var index int

	for i, token := range tokens {
		if index >= len(tokens) {
			break
		}
		if token.name == PROFILE_NAME {
			prof := profiles.Profile{Name: token.value, DefaultSwitch: false}
			if tokens[i+1].name == INHERITED_PROFILE_NAME {
				prof.InheritFrom = tokens[i+1].value
			}
			if tokens[i+2].name == DEFAULT_SWITCH {
				prof.DefaultSwitch = true
			}
			for j := i + 1; j < len(tokens); j++ {
				if tokens[j].name == PROFILE_NAME {
					break
				}
				variable := profiles.ProfileVariable{}
				if tokens[j].name == VAR_NAME {
					variable.Name = tokens[j].value
					if tokens[j+2].name == VAR_VALUE {
						variable.Value = tokens[j+2].value
						prof.Variables = append(prof.Variables, variable)
					}
				}
				index = j + 1
			}
			items = append(items, prof)
		}
	}

	return items
}
