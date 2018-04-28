// This file was generated by github.com/vektah/gqlgen, DO NOT EDIT

package models

import (
	fmt "fmt"
	io "io"
	strconv "strconv"
)

type DateFilter struct {
	Value    string
	Timezone *string
	Op       *DateFilterOp
}

type DateFilterOp string

const (
	DateFilterOpEq  DateFilterOp = "EQ"
	DateFilterOpNeq DateFilterOp = "NEQ"
	DateFilterOpGt  DateFilterOp = "GT"
	DateFilterOpGte DateFilterOp = "GTE"
	DateFilterOpLt  DateFilterOp = "LT"
	DateFilterOpLte DateFilterOp = "LTE"
)

func (e DateFilterOp) IsValid() bool {
	switch e {
	case DateFilterOpEq, DateFilterOpNeq, DateFilterOpGt, DateFilterOpGte, DateFilterOpLt, DateFilterOpLte:
		return true
	}
	return false
}

func (e DateFilterOp) String() string {
	return string(e)
}

func (e *DateFilterOp) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DateFilterOp(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DATE_FILTER_OP", str)
	}
	return nil
}

func (e DateFilterOp) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
