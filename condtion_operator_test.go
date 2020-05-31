package StitchingSQLGo

import (
	"testing"
)

func TestConditionOperator_ConditionOperator_EQ(t *testing.T) {
	s := SqlBuilder{}

	if err := EQ.conditionOperator(&s); err != nil {
		t.Fatal(err)
	}

	except := " ="
	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}
}
func TestConditionOperator_ConditionOperator_NEQ(t *testing.T) {
	s := SqlBuilder{}

	if err := NEQ.conditionOperator(&s); err != nil {
		t.Fatal(err)
	}

	except := " !="
	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}
}
func TestConditionOperator_ConditionOperator_GT(t *testing.T) {
	s := SqlBuilder{}

	if err := GT.conditionOperator(&s); err != nil {
		t.Fatal(err)
	}

	except := " >"
	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}
}
func TestConditionOperator_ConditionOperator_LT(t *testing.T) {
	s := SqlBuilder{}

	if err := LT.conditionOperator(&s); err != nil {
		t.Fatal(err)
	}

	except := " <"
	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}
}
func TestConditionOperator_ConditionOperator_GTEQ(t *testing.T) {
	s := SqlBuilder{}

	if err := GTEQ.conditionOperator(&s); err != nil {
		t.Fatal(err)
	}

	except := " >="
	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}
}
func TestConditionOperator_ConditionOperator_LTEQ(t *testing.T) {
	s := SqlBuilder{}

	if err := LTEQ.conditionOperator(&s); err != nil {
		t.Fatal(err)
	}

	except := " <="
	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}
}
func TestConditionOperator_ConditionOperator_NGT(t *testing.T) {
	s := SqlBuilder{}

	if err := NGT.conditionOperator(&s); err != nil {
		t.Fatal(err)
	}

	except := " !>"
	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}
}
func TestConditionOperator_ConditionOperator_NLT(t *testing.T) {
	s := SqlBuilder{}

	if err := NLT.conditionOperator(&s); err != nil {
		t.Fatal(err)
	}

	except := " !<"
	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}
}

func TestConditionOperator_ConditionOperator_IS(t *testing.T) {
	s := SqlBuilder{}

	if err := IS.conditionOperator(&s); err != nil {
		t.Fatal(err)
	}

	except := " is"
	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}
}

func TestConditionOperator_ConditionOperator_ISNOT(t *testing.T) {
	s := SqlBuilder{}

	if err := ISNot.conditionOperator(&s); err != nil {
		t.Fatal(err)
	}

	except := " is not"
	if s.String() != except {
		t.Fatalf("except\n%s\nnow\n%s", except, s.String())
	}
}
