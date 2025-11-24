package models

import (
    "database/sql/driver"
    "errors"
    "strings"
    "time"
)

type GrowwTime struct {
    time.Time
}

func (gt *GrowwTime) UnmarshalJSON(b []byte) error {
    s := strings.Trim(string(b), `"`)
    if s == "" || s == "null" {
        return nil
    }

    // Try without timezone
    t, err := time.Parse("2006-01-02T15:04:05", s)
    if err == nil {
        gt.Time = t
        return nil
    }

    // Try with timezone
    t, err = time.Parse(time.RFC3339, s+"Z")
    if err == nil {
        gt.Time = t
        return nil
    }

    return err
}

//
// REQUIRED BY GORM
//

// Scan implements the sql.Scanner interface
func (gt *GrowwTime) Scan(value interface{}) error {
    if value == nil {
        return nil
    }

    switch v := value.(type) {
    case time.Time:
        gt.Time = v
        return nil
    default:
        return errors.New("invalid time format for GrowwTime")
    }
}

// Value implements the driver.Valuer interface
func (gt GrowwTime) Value() (driver.Value, error) {
    return gt.Time, nil
}
