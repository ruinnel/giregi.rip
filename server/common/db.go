package common

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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
	Like Op = "LIKE"
)

type Order string

const (
	Asc  Order = "ASC"
	Desc Order = "DESC"
)

type Condition struct {
	Field string
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
		return strings.Compare(conditions[i].Field, conditions[j].Field) > 0
	})
	for _, condition := range conditions {
		list = append(list, condition.String())
	}
	return strings.Join(list, ",")
}

func ConditionsToQueries(conditions []Condition) []qm.QueryMod {
	var queries []qm.QueryMod
	for idx, condition := range conditions {
		col := condition.Field
		val := condition.Val
		op := condition.Op
		if op == Like {
			val = fmt.Sprintf("%%%v%%", val)
		}
		if idx == 0 {
			queries = append(queries, qm.Where(fmt.Sprintf("%s %s ?", col, op), val))
		} else {
			queries = append(queries, qm.And(fmt.Sprintf("%s %s ?", col, op), val))
		}
	}
	return queries
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
