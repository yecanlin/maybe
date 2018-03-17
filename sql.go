package main

import (
	"fmt"
	"errors"
)

var ErrConnDone = errors.New("database/sql: connection is already closed")

var ErrNoRows = errors.New("sql: no rows in result set")

var ErrTxDone = errors.New("sql: Transaction has already been committed or rolled back")
