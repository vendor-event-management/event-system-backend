package utils

import (
	"database/sql"
	"encoding/json"
	"strings"
	"time"
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

func ConvertStringToSQLNullString(s string) sql.NullString {
	if IsEmptyString(s) {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func ConvertTimeToSQLNullTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{
			Time: t, Valid: false,
		}
	}
	return sql.NullTime{Time: t, Valid: true}
}

func ConvertStringToTime(dateStr string) (time.Time, error) {
	format := "02-01-2006"
	parsedTime, err := time.Parse(format, dateStr)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, err
}

func ConvertToJSONString(slice []string) (string, error) {
	var result []byte
	result, err := json.Marshal(slice)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
