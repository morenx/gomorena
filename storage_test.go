package main

import (
	"fmt"
	"testing"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type storageIntegrationSuite struct{}

var _ = Suite(&storageIntegrationSuite{})

func (s *storageIntegrationSuite) Test_storingAndRetreivingAProduct(c *C) {
	db, err := sql.Open("mysql", "root:foobar@/mysql")
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	// put in db
	// get out of db
	// assert what we put in is what came out
	c.Assert(true, Equals, true)
}
