package main

import "fmt"

func sqlQuote(x interface{}) string {
	if x == nil {
		return ""
	} else if _, ok := x.(int); ok {
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(uint); ok {
		return fmt.Sprintf("%d", x)
	} else if b, ok := x.(bool); ok {
		if !b {
			return "FALSE"
		}
		return "TRUE"
	} else if s, ok := x.(string); ok {
		return sqlQuoteStrings(s)
	} else {
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func sqlQuote2(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if !x {
			return "FALSE"
		}
		return "TRUE"
	case string:
		return sqlQuoteString(x)
	default:
		panic(fmt.Sprintf("unexpected type: %T: %v", x, x))
	}
}
