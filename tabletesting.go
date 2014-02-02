// Tabletesting package provides helper structures and interfaces for dealing
// with table based tests.
package tabletesting

import (
    "fmt"
    "reflect"
)

type TableTester interface {
    RunAgainst(interface{}) error
}

// Run a series of testers.
// Logs errors and returns false on the first failure.
func RunSeries(testers []TableTester, data interface{}) error {
    for _, tester := range testers {
        if err := tester.RunAgainst(data); err != nil {
            return err
        }
    }
    return nil
}

// A TableTester for string based data.
type StringTableTester struct {
    Field string
    Value string
}

// Run the TableTester against the supplied data.
// Looks for a field in the supplied data name Field and compares it's value
// to Value, returns error or nil.
func (stt StringTableTester) RunAgainst(ti interface{}) error {
    field := reflect.ValueOf(ti).FieldByName(stt.Field)
    if field.IsValid() == false {
        err := fmt.Errorf("Failed to find field \"%s\"\n", stt.Field)
        return err
    }
    actual := field.Interface()
    expected := stt.Value
    if actual != expected {
        err := fmt.Errorf("Fail on %s, expected \"%s\", actual \"%s\"\n", stt.Field, expected, actual)
        return err
    }
    return nil
}
