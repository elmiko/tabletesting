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
    field := reflect.ValueOf(ti).FieldByName(stt.Field)
    if field.IsValid() == false {
        t.Logf("Failed to find field %s\n", stt.Field)
        return false
    }
    actual := field.Interface()
    expected := stt.Value
    if actual != expected {
        t.Logf("Fail on %s, expected %s, actual %s\n", stt.Field, expected, actual)
        return false
    }
    return true
}
