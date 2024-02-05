package etch

import (
	"fmt"
	"html"
)

func Escape(a any) string {
	return html.EscapeString(anyToString(a))
}

func anyToString(a any) string {
	if v, ok := a.(string); ok {
		return v
	} else if v, ok := a.(int); ok {
		return fmt.Sprintf("%d", v)
	} else if v, ok := a.(float64); ok {
		return fmt.Sprintf("%f", v)
	} else if v, ok := a.(*string); ok {
		if v == nil {
			return ""
		} else {
			return *v
		}
	} else {
		return fmt.Sprintf("%s", a)
	}
}
