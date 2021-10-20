package spanner

import "cloud.google.com/go/spanner"

// string型をspanner.NullStringに変換します。
func ConvSpannerNullString(param string) spanner.NullString {
	if len(param) == 0 {
		return spanner.NullString{}
	} else {
		return spanner.NullString{StringVal: param, Valid: true}
	}
}
