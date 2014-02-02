tabletesting
============

A Go library to aid in running table based tests.

Basic Usage
-----------

A TableTester will attempt to check each field name in the test data supplied
to TableTester.RunAgainst with the reference value specified. This example
shows a string based test.

    func TestSomeTest(t *testing.t) {
        // a table with 2 tests
        testtable = []StringTableTester{
            {"SomeStringField", "test data"},
            {"SomeOtherStringField", "more test data"},
        }
        actualtestdata := struct {
            SomeStringField string
            SomeOtherStringField string
        }{ "test data", "more test data"}

        // this will return true if all tests succeed
        if testtable.RunAgainst(actualtestdata, t) == false {
            t.Error("A test has failed!\n")
        }
    }
