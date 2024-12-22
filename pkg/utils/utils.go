package utils

import (
	"database/sql"
	"encoding/json"
	"strings"
)

func IsEmptyString(s string) bool {
	return strings.TrimSpace(s) == ""
}

func ParseDates(dates string) ([]string, error) {
	var dateArray []string
	err := json.Unmarshal([]byte(dates), &dateArray)
	if err != nil {
		return nil, err
	}

	return dateArray, nil
}

func ConvertToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}

func ConvertToJSONString(slice []string) (string, error) {
	var result []byte
	result, err := json.Marshal(slice)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
