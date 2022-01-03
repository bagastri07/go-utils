package xgqlgen

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/kumparan/go-utils"
	"gorm.io/gorm"
)

// MarshalInt64ID marshal int64 to string ID
func MarshalInt64ID(i int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(fmt.Sprintf(`"%d"`, i)))
	})
}

// UnmarshalInt64ID unmarshal ID into int64
func UnmarshalInt64ID(v interface{}) (int64, error) {
	switch v := v.(type) {
	case string:
		return utils.StringToInt64(v), nil
	case json.Number:
		return utils.StringToInt64(string(v)), nil
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case int32:
		return int64(v), nil
	default:
		return 0, fmt.Errorf("%T is not a int64", v)
	}
}

// MarshalGormDeletedAt marshal gorm.DeletedAt to string
func MarshalGormDeletedAt(gd gorm.DeletedAt) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		ts := utils.FormatTimeRFC3339(&gd.Time)
		w.Write([]byte(fmt.Sprintf(`"%s"`, ts)))
	})
}

// UnmarshalGormDeletedAt unmarshal v into gorm.DeletedAt
func UnmarshalGormDeletedAt(v interface{}) (gorm.DeletedAt, error) {
	switch v := v.(type) {
	case string:
		tt, err := time.Parse(time.RFC3339Nano, v)
		if err != nil {
			return gorm.DeletedAt{}, err
		}
		gd := gorm.DeletedAt{}
		err = gd.Scan(tt)
		return gd, err
	default:
		return gorm.DeletedAt{}, errors.New("v is not a valid string time")
	}
}

// MarshalTimeRFC3339Nano marshal time.Time to string RFC3339Nano
func MarshalTimeRFC3339Nano(tt time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		ts := utils.FormatTimeRFC3339(&tt)
		w.Write([]byte(fmt.Sprintf(`"%s"`, ts)))
	})
}

// UnmarshalTimeRFC3339Nano unmarshal v into time.Time
func UnmarshalTimeRFC3339Nano(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		tt, err := time.Parse(time.RFC3339Nano, v)
		return tt, err
	default:
		return time.Time{}, fmt.Errorf("%T is not a valid RFC3339Nano time", v)
	}
}