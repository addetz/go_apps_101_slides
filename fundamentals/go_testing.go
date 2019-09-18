package repo_test

//START OMIT
import "testing"

func TestFirst(t *testing.T) {
    if true {
        t.Error("expected false, got true")
    }
}
//END OMIT
