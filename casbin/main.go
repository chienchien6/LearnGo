package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	a, _ := gormadapter.NewAdapter("mysql", "root:2Password@tcp(127.0.0.1:3306)/casbin", true) // Your driver and data source.
	e, _ := casbin.NewEnforcer("casbin/model.conf", a)
	e.AddFunction("my_func", KeyMatchFunc)

	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.
	added, err := e.AddGroupingPolicy("alice", "data2_admin")
	fmt.Println(added)
	fmt.Println(err)
	//filteredPolicy, _ := e.GetFilteredPolicy(0, "data1")

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		fmt.Printf("%s", err)
		// handle err
	}

	if ok {
		// permit alice to read data1
		fmt.Println("通過")
	} else {
		// deny the request, show an error
		fmt.Println("未通過")
	}

	// You could use BatchEnforce() to enforce some requests in batches.
	// This method returns a bool slice, and this slice's index corresponds to the row index of the two-dimensional array.
	// e.g. results[0] is the result of {"alice", "data1", "read"}
	//results, err := e.BatchEnforce([][]interface{}{{"alice", "data1", "read"}, {"bob", "data2", "write"}, {"jack", "data3", "read"}})
}

func KeyMatch(key1 string, key2 string) bool {

	return key1 == key2
}

func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(KeyMatch(name1, name2)), nil
}
