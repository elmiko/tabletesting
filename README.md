tabletesting
============

A Go library to aid in running table based tests.

Basic Usage
-----------

A TableTester will attempt to check each field name in the test data supplied
to TableTester.RunAgainst with the reference value specified. This example
shows a string based test.

    func RunOneTest(t *testing.t) {
        test := StringTableTester{"SomeStringField", "test data"}
        actualdata := struct {SomeStringField string}{"test data"}
        if err := test.RunAgainst(actualdata); err != nil {
            t.Errorf("Test failed with error \"%s\"\n", err)
        }
    }

This example shows how to run multiple tests using the RunSeries function.

    func RunMultipleTests(t *testing.t) {
        testtable := []TableTester{
            StringTableTester{"SomeStringField", "test data"},
            StringTableTester{"SomeOtherStringField", "more test data"},
        }
        actualdata := struct {
            SomeStringField string
            SomeOtherStringField string
        }{ "test data", "more test data"}

        if err := RunSeries(testtable, actualdata); err != nil {
            t.Errorf("A test has failed with error \"%s\"\n", err)
        }
    }
