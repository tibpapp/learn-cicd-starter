
// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    // Step 1: Call the Add function with some input
    result := Add(2, 3)

    // Step 2: Check the output and behavior
    expected := 5
    if result != expected {
        t.Errorf("Expected %d, but got %d", expected, result)
    }

    // Step 3: If something is wrong, use t.Error() or t.Errorf() to report it
    // t.Error("Some issue occurred")
}

