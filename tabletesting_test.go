package tabletesting

import (
    "testing"
)

func TestRunSeries(t *testing.T) {
    testers := []TableTester{
        StringTableTester{"Field1", "test value"},
        StringTableTester{"Field2", "another test value"},
    }
    testee := struct{
            Field1 string
            Field2 string
        }{"test value", "another test value"}
    if err := RunSeries([]TableTester(testers), testee); err != nil {
        t.Errorf("RunSeries failed with error \"%s\"\n", err)
    }
}

func TestStringTableTesterRunAgainst(t *testing.T) {
    tester := StringTableTester{ "Field1", "test value"}
    testee := struct{Field1 string}{"test value"}
    if err := tester.RunAgainst(testee); err != nil {
        t.Errorf("RunAgainst failed with error \"%s\"\n", err)
    }
}

func TestStringTableTesterRunAgainstFieldFailure(t *testing.T) {
    tester := StringTableTester{ "Field1", "test value"}
    testee := struct{Field2 string}{"test value"}
    if err := tester.RunAgainst(testee); err == nil {
        t.Error("RunAgainst produced no error when given a mismatched field entry\n")
    }
}

func TestStringTableTesterRunAgainstValueFailure(t *testing.T) {
    tester := StringTableTester{ "Field1", "test value"}
    testee := struct{Field1 string}{"bad test value"}
    if err := tester.RunAgainst(testee); err == nil {
        t.Error("RunAgainst produced no error when given a mismatched value entry\n")
    }
}
