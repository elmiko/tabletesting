package tabletesting

import (
    "testing"
)

func TestStringTableTesterRunAgainst(t *testing.T) {
    tester := StringTableTester{ "Field1", "test value"}
    testee := struct{Field1 string}{"test value"}
    if tester.RunAgainst(testee, t) == false {
        t.Error("RunAgainst returned false, expected true\n")
    }
}

func TestStringTableTesterRunAgainstFieldFailure(t *testing.T) {
    tester := StringTableTester{ "Field1", "test value"}
    testee := struct{Field2 string}{"test value"}
    if tester.RunAgainst(testee, t) == true {
        t.Error("RunAgainst returned true, expected false\n")
    }
}

func TestStringTableTesterRunAgainstValueFailure(t *testing.T) {
    tester := StringTableTester{ "Field1", "test value"}
    testee := struct{Field1 string}{"bad test value"}
    if tester.RunAgainst(testee, t) == true {
        t.Error("RunAgainst returned true, expected false\n")
    }
}
