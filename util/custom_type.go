package util

import (
	"database/sql"
	"encoding/json"
)

type NullInt32 struct {
	sql.NullInt32
}

func (ni *NullInt32) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int32)
	}
	return json.Marshal(nil)
}

func (ni *NullInt32) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ni.Valid = false
		return nil
	}

	var val int32
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	ni.Int32 = val
	ni.Valid = true
	return nil
}
