package gomydb_test

import (
	"fmt"
	. "gomydb"
	"testing"
)

func TestGetString(t *testing.T) {
	config := NewConfig("conf/db.conf")
	fmt.Println(config.Get("test_db", "host"))
	fmt.Println(config.Get("test_db", "port"))
}
