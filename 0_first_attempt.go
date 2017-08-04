package main

// #cgo LDFLAGS: -lfbembed
//
// #include "ibase.h"
import "C"

import (
    "fmt"
)

func main() {
    var status C.ISC_STATUS_ARRAY
    var newdb C.isc_db_handle
    var trans C.isc_tr_handle
    var dbname string = "CREATE DATABASE 'foo.fdb'"
    
    fmt.Println(dbname, status, newdb, trans)

    ret := C.isc_dsql_execute_immediate(&status[0], &newdb, &trans, 0, (*C.ISC_SCHAR)(C.CString(dbname)), 1, nil)
    
    fmt.Println("ret", ret)

    if ret == 0 {
        sqlcode := C.isc_sqlcode(&status[0])
        fmt.Println("sqlcode", sqlcode)
    } else {
        C.isc_print_status(&status[0])
    }
}