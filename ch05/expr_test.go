package main

import "testing"

type Var string

type Env map[Var]float64

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}

	var prevExpr string
	for _, tt := range tests {
		if tt.expr != prevExpr {
			fmt.Printf("\n%s\n", tt.expr)
			prevExpr = tt.expr
		}

		expr, err := Parse(tt.expr)
		if err != nil {
			t.Error(err)
			continue
		}

		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}

		// t.Run(tt.name, func(t *testing.T) {
		// 	if got := tt.l.Eval(tt.args.in0); got != tt.want {
		// 		t.Errorf("literial.Eval() = %v, want %v", got, tt.want)
		// 	}
		// })
	}
}
