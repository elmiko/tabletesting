package tabletesting

import (
    "testing"
)

func TestStringTableTesterRunAgainst(t *testing.T) {
    tester := StringTableTester{ "Field1", "test value"}
    testee := struct{Field1 string}{"test value"}
    if tester.RunAgainst(testee, t) == false {
        t.Error("Failed RunAgainst\n")
    }
}
