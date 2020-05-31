package StitchingSQLGo

import (
	"testing"
)

func TestWhereItemLinker_WhereItem_And(t *testing.T) {
	s := SqlBuilder{}

	if err := AND.WhereItem(&s); err != nil {
		t.Fatal(err)
	}

	except := " and"
	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}

}
func TestWhereItemLinker_WhereItem_Or(t *testing.T) {
	s := SqlBuilder{}

	if err := OR.WhereItem(&s); err != nil {
		t.Fatal(err)
	}

	except := " or"
	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}

}
