package StitchingSQLGo

import (
	"testing"
)

func TestLimit_Limit_0(t *testing.T) {
	s, l := SqlBuilder{}, Limit(0)

	except := ""
	if err := l.limit(&s); err != nil {
		t.Fatal(err)
	}
	if s.String() != except {
		t.Fatalf("except:%s\nnow:%s", except, s.String())
	}
}

func TestLimit_Limit_1(t *testing.T) {
	s, l := SqlBuilder{}, Limit(1)

	except := " limit 1"
	if err := l.limit(&s); err != nil {
		t.Fatal(err)
	}
	if s.String() != except {
		t.Fatalf("except:%s\nnow:%s", except, s.String())
	}
}
