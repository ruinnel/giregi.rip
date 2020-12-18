package common

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/asdine/storm/v3/q"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Op string

const (
	// Op
	Lt   Op = "<"
	Lte  Op = "<="
	Eq   Op = "="
	Gt   Op = ">"
	Gte  Op = ">="
	In   Op = "IN"
	Like Op = "LIKE"
)

type Order string

const (
	Asc  Order = "ASC"
	Desc Order = "DESC"
)

type Condition struct {
	Field reflect.StructField
	Op    Op
	Val   interface{}
}

type FetchResult struct {
	Total      int64       `json:"total"`
	Data       interface{} `json:"data"`
	NextCursor string      `json:"nextCursor"`
}

const (
	timeFormat = "2006-01-02T15:04:05.999Z"
)

func (c Condition) String() string {
	return fmt.Sprintf("%s%s%v", c.Field, string(c.Op), c.Val)
}

func hash(src string) string {
	h := sha256.New()
	h.Write([]byte(src))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func conditionToString(conditions []Condition) string {
	var list []string

	sort.Slice(conditions, func(i, j int) bool {
		return strings.Compare(conditions[i].Field.Name, conditions[j].Field.Name) > 0
	})
	for _, condition := range conditions {
		list = append(list, condition.String())
	}
	return strings.Join(list, ",")
}

func toStringList(list []interface{}) []string {
	var result []string
	for _, v := range list {
		result = append(result, fmt.Sprintf("%v", v))
	}
	return result
}

func makeQuery(condition Condition, first bool) qm.QueryMod {
	field := condition.Field
	val := condition.Val
	op := condition.Op
	col := field.Tag.Get("mysql")
	switch op {
	case In:
		switch val.(type) {
		case []interface{}:
			list := val.([]interface{})
			if len(list) > 0 {
				val = strings.Join(toStringList(list), ",")
			}
		}
	case Like:
		val = fmt.Sprintf("%%%v%%", val)
	default:
		break
	}
	if first {
		return qm.Where(fmt.Sprintf("%s %s ?", col, op), val)
	} else {
		return qm.And(fmt.Sprintf("%s %s ?", col, op), val)
	}
}

func ConditionsToQueries(conditions []Condition) []qm.QueryMod {
	var queries []qm.QueryMod
	for idx, condition := range conditions {
		queries = append(queries, makeQuery(condition, idx == 0))
	}
	return queries
}

func makeMatcher(condition Condition) q.Matcher {
	col := condition.Field.Name
	val := condition.Val
	op := condition.Op
	switch op {
	case Lt:
		return q.Lt(col, val)
	case Lte:
		return q.Lte(col, val)
	case Eq:
		return q.Eq(col, val)
	case Gt:
		return q.Gt(col, val)
	case Gte:
		return q.Gte(col, val)
	case Like:
		return q.Re(col, fmt.Sprintf(`^.*%s.*$`, val))
	case In:
		return q.In(col, val)
	default:
		return nil
	}
}

func ConditionsToMatchers(conditions []Condition) []q.Matcher {
	var matchers []q.Matcher

	for _, condition := range conditions {
		matchers = append(matchers, makeMatcher(condition))
	}
	return matchers
}

func EncodeCursor(conditions []Condition, t time.Time) string {
	return fmt.Sprintf("%d_%s", t.Unix(), hash(conditionToString(conditions)))
}

func DecodeCursor(conditions []Condition, cursor string) (*time.Time, error) {
	if len(cursor) == 0 {
		return nil, nil
	}
	slice := strings.Split(cursor, "_")
	if len(slice) != 2 {
		return nil, errors.New("invalid cursor(split)")
	}

	t := slice[0]
	checksum := slice[1]

	if checksum != hash(conditionToString(conditions)) {
		return nil, errors.New("invalid cursor(checksum)")
	}

	unix, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return nil, errors.New("invalid cursor(parse)")
	}

	c := time.Unix(unix, int64(0))
	return &c, nil
}
