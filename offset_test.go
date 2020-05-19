package StitchingSQLGo

import (
	"testing"
)

func TestOffset_Offset0(t *testing.T) {
	s, o := sql{}, Offset(0)

	except := ""
	if err := o.offset(&s); err != nil {
		t.Fatal(err)
	} else if s.String() != except {
		t.Fatalf("except:%s\nnow:%s", except, s.String())
	}
}

func TestOffset_Offset1(t *testing.T) {
	s, o := sql{}, Offset(1)
	except := " offset 1"
	if err := o.offset(&s); err != nil {
		t.Fatal(err)
	} else if s.String() != except {
		t.Fatalf("except:%s\nnow:%s", except, s.String())
	}
}
