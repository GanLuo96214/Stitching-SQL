package StitchingSQLGo

import "testing"

func TestField1_Field(t *testing.T) {
	s, fs := SqlBuilder{}, Field1{}

	if err := fs.Field(&s, false); err != nil {
		t.Fatal(err)
	}

	except := " 1"
	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}

}
