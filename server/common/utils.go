package common

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"github.com/speps/go-hashids"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Result struct {
	Result       bool        `json:"result,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	ErrorCode    *ErrorCode  `json:"errorCode,omitempty"`
	ErrorMessage *string     `json:"errorMessage,omitempty"`
}

var hashid, _ = hashids.NewWithData(&hashids.HashIDData{
	Alphabet:  hashids.DefaultAlphabet,
	MinLength: 20,
	Salt:      "giregi_rip",
})

func EncodeHashId(val int64) string {
	encoded, _ := hashid.EncodeInt64([]int64{val})
	return encoded
}

func DecodeHashId(val string) (int64, error) {
	decoded, err := hashid.DecodeInt64WithError(val)
	if err != nil {
		return -1, err
	} else {
		return decoded[0], nil
	}

}

func ParseJsonBody(request *http.Request, param interface{}) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, param)
	if err != nil {
		return err
	}
	return nil
}

func GenerateAccessToken(userId, userAgent string) string {
	hash := sha512.New()
	hash.Write([]byte(fmt.Sprintf("%s,%s,%d", userId, userAgent, time.Now().Unix())))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func WriteJson(writer http.ResponseWriter, data interface{}) {
	result := Result{
		Result:       true,
		Data:         data,
		ErrorMessage: nil,
	}
	body, err := json.Marshal(result)
	writer.Header().Add("Content-Type", "application/json")
	if err == nil {
		writer.Write(body)
	} else {
		writer.Write([]byte("{\"result\": false,\"errorMessage\": \"to json fail\"}"))
	}
}

func WriteJsonWithCount(writer http.ResponseWriter, count int64, data interface{}) {
	result := Result{
		Result: true,
		Data: map[string]interface{}{
			"total": count,
			"data":  data,
		},
		ErrorMessage: nil,
	}
	body, err := json.Marshal(result)
	writer.Header().Add("Content-Type", "application/json")
	if err == nil {
		writer.Write(body)
	} else {
		writer.Write([]byte("{\"result\": false,\"errorMessage\": \"to json fail\"}"))
	}
}

func WriteError(writer http.ResponseWriter, err error, data interface{}) {
	writeError(writer, err, data)
}

func writeError(writer http.ResponseWriter, err error, data interface{}) {
	result := Result{Result: false, Data: data}
	statusErr, ok := err.(*StatusError)
	if ok {
		result.ErrorCode = &statusErr.Code
		result.ErrorMessage = &statusErr.Message
		writer.WriteHeader(statusErr.Status)
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Header().Add("Content-Type", "application/json")

	body, e := json.Marshal(result)
	if e == nil {
		writer.Write(body)
	} else {
		writer.Write([]byte("{\"result\": false,\"errorMessage\": \"to json fail\"}"))
	}
}

func getJsonKeyAndValue(fieldType reflect.StructField, value interface{}) (*string, interface{}) {
	tag := fieldType.Tag.Get("json")
	tags := strings.Split(tag, ",")
	name := fieldType.Name
	ignore := false
	omitempty := false
	if len(tags) >= 1 {
		if tags[0] == "-" {
			ignore = true
		} else {
			name = tags[0]
		}
	}
	if len(tags) >= 2 {
		omitempty = tags[1] == "omitempty"
	}

	if ignore || (omitempty && value == nil) {
		return nil, nil
	}

	switch value.(type) {
	case *time.Time:
		if value.(*time.Time) != nil {
			return &name, value.(*time.Time).UnixNano() / int64(time.Millisecond)
		}
	case time.Time:
		return &name, value.(time.Time).UnixNano() / int64(time.Millisecond)
	default:
		return &name, value
	}
	return nil, nil
}

// time.Time를 json 변환시 timestamp(millisecond)로 처리.
// type 선언 후 MarshalJSON()을 구현하는 방식은 firestore sdk에서 model에 데이터를 맵핑 할때 문제가 생김.
// BaseModel을 flat 하게 json으로 변환
func MarshalJSON(src interface{}) ([]byte, error) {
	fields := reflect.ValueOf(src)
	m := map[string]interface{}{}
	for i := 0; i < fields.NumField(); i++ {
		fieldType := fields.Type().Field(i)
		value := fields.Field(i).Interface()
		if fieldType.Name == "BaseModel" {
			baseModelFields := reflect.ValueOf(value)
			for j := 0; j < baseModelFields.NumField(); j++ {
				baseModelFieldType := baseModelFields.Type().Field(j)
				baseModelValue := baseModelFields.Field(j).Interface()
				key, value := getJsonKeyAndValue(baseModelFieldType, baseModelValue)
				if key != nil {
					m[*key] = value
				}
			}
		} else {
			key, value := getJsonKeyAndValue(fieldType, value)
			if key != nil {
				m[*key] = value
			}
		}
	}
	return json.Marshal(m)
}

func GetIdFromPath(request *http.Request) (int64, error) {
	path := request.URL.Path
	if strings.HasSuffix(path, "/") {
		path = path[0 : len(path)-1]
	}
	splits := strings.Split(path, "/")
	idParam := splits[len(splits)-1]

	return strconv.ParseInt(idParam, 10, 64)
}

func ParsePath(regex *regexp.Regexp, path string) map[string]string {
	match := regex.FindStringSubmatch(path)
	paramsMap := make(map[string]string)
	for i, name := range regex.SubexpNames() {
		if i > 0 && i <= len(match) {
			val := match[i]
			if len(val) > 0 {
				paramsMap[name] = match[i]
			}
		}
	}
	return paramsMap
}
