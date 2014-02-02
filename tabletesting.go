// Tabletesting package provides helper structures and interfaces for dealing
// with table based tests.
package tabletesting

import (
    "reflect"
    "testing"
)

type TableTester interface {
    RunAgainst(interface{}, *testing.T) bool
}

type StringTableTester struct {
    Field string
    Value string
}

func (stt StringTableTester) RunAgainst(ti interface{}, t *testing.T) bool {
    retval := true
    actual := reflect.ValueOf(ti).FieldByName(stt.Field).Interface()
    expected := stt.Value
    if actual != expected {
        t.Logf("Fail on %s, expected %s, actual %s\n", stt.Field, expected, actual)
        retval = false
    }
    return retval
}
