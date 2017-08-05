package main

/* 
#cgo LDFLAGS: -lfbembed

#include "ibase.h"
#include <stdio.h>

int isc_start_transaction2(ISC_STATUS *status_vector, isc_tr_handle *trans_handle, short db_handle_count, isc_db_handle *db_handle, unsigned short tpb_length, char *tpb_address)
{
    return isc_start_transaction(status_vector, trans_handle, db_handle_count, db_handle, tpb_length, tpb_address);
}

*/
import "C"

import (
    "errors"
    "fmt"
)

var status C.ISC_STATUS_ARRAY
var newdb C.isc_db_handle
var trans C.isc_tr_handle
var dbname string = "CREATE DATABASE 'foo.fdb'"

func transaction(query string) error {
    C.isc_start_transaction2(
        &status[0], &trans, 1, &newdb, 0, nil,
    )
    ret := C.isc_dsql_execute_immediate(&status[0], &newdb, &trans, 0, (*C.ISC_SCHAR)(C.CString(query)), 1, nil)
    
    if ret != 0 {
        //C.isc_rollback_transaction(&status[0], &trans)
        C.isc_print_status(&status[0])
        return errors.New("command failed")
    } else {
        C.isc_commit_transaction(&status[0], &trans)
    }
    return nil
}

func main() {
    fmt.Println(dbname, status, newdb, trans)

    ret := C.isc_dsql_execute_immediate(&status[0], &newdb, &trans, 0, (*C.ISC_SCHAR)(C.CString(dbname)), 1, nil)
    
    fmt.Println("ret", ret)

    if ret == 0 {
        sqlcode := C.isc_sqlcode(&status[0])
        fmt.Println("sqlcode", sqlcode)
        C.isc_commit_transaction(&status[0], &trans)
        C.isc_detach_database(&status[0], &newdb)
        dbname = "foo.fdb"
        C.isc_attach_database(&status[0], 0, (*C.ISC_SCHAR)(C.CString(dbname)), &newdb, 0, nil) 
        defer C.isc_detach_database(&status[0], &newdb)

        fmt.Println("after attach")
        C.isc_print_status(&status[0])
        err := transaction("CREATE TABLE test (when_created DATE);")
        if err != nil { fmt.Println(err) }
        transaction("INSERT INTO test VALUES ('NOW');")
        

    } else {
        C.isc_print_status(&status[0])
    }
    
}