package main

import "fmt"

//类型分支
func main() {
	fmt.Println(sqlQuota(12))
}

func sqlQuota(v interface{}) string {
	switch v := v.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", v)
	default:
		return ""
	}
}
