package broadpeakio

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

const baseUrl = "https://api.broadpeak.io/v1/"

func makeFirstCharLower(s string) string {
	if len(s) == 0 {
		return s // Return empty string if input is empty
	}

	firstChar := strings.ToLower(string(s[0]))
	return firstChar + s[1:]
}

func structToMap(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	value := reflect.ValueOf(s)
	typeOfS := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldName := typeOfS.Field(i).Name
		m[fieldName] = field.Interface()
	}

	return m
}

func getReqUrlWithQuery(apiUrl string, options map[string]interface{}) (string, error) {
	requestURL, err := url.Parse(apiUrl)
	if err != nil {
		return "", err
	}

	queryParams := requestURL.Query()

	// Set optional parameters as query parameters
	for key, val := range options {
		valueStr, okStr := val.(string)
		if okStr && valueStr != "" {
			queryParams.Set(makeFirstCharLower(key), valueStr)
		}

		valueInt, okInt := val.(uint)
		if okInt && valueInt != 0 {
			queryParams.Set(makeFirstCharLower(key), fmt.Sprint(valueInt))
		}
	}

	requestURL.RawQuery = queryParams.Encode()

	reqUrl := requestURL.String()

	for key, val := range options {
		valueUintList, okUintList := val.([]uint)
		if okUintList {
			questionMark := '?'
			for _, valUint := range valueUintList {
				if strings.ContainsRune(reqUrl, questionMark) {
					reqUrl = reqUrl + "&" + makeFirstCharLower(key) + "=" + fmt.Sprint(valUint)
				} else {
					reqUrl = reqUrl + "?" + makeFirstCharLower(key) + "=" + fmt.Sprint(valUint)
				}
			}
		}
	}

	return reqUrl, nil
}

func checkRequiredFieldsNonEmpty(obj interface{}, requiredFields []string) (bool, string) {
	v := reflect.ValueOf(obj)
	typeOfV := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := typeOfV.Field(i).Name
		if isRequiredField(fieldName, requiredFields) && isEmptyValue(field) {
			return false, field.Interface().(string)
		}
	}

	return true, ""
}

func isRequiredField(field interface{}, requiredFileds []string) bool {
	for _, reqField := range requiredFileds {
		fieldStr, isOk := field.(string)
		if isOk && fieldStr == reqField {
			return true
		}
	}
	return false
}

func isEmptyValue(field reflect.Value) bool {
	switch field.Kind() {
	case reflect.String:
		return field.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return field.Int() == 0
	default:
		// Rachit add a check for other types like Identifiable
		/* if field.Interface().(string) == "{0}" {
			return false
		} */
		return false
	}
}

func addOffsetUrl(url string, offset uint, limit uint) string {
	questionMark := '?'

	if offset != 0 {
		if strings.ContainsRune(url, questionMark) {
			url = url + "&offset=" + fmt.Sprint(offset)
		} else {
			url = url + "?offset=" + fmt.Sprint(offset)
		}
	}

	if limit != 0 {
		if strings.ContainsRune(url, questionMark) {
			url = url + "&limit=" + fmt.Sprint(limit)
		} else {
			url = url + "?limit=" + fmt.Sprint(limit)
		}
	}

	return url
}

// Deserialising the API response
func deserialiseApiResponse(resp string, target interface{}) error {
	err := json.Unmarshal([]byte(resp), &target)
	if err != nil {
		return err
	}
	return nil
}
