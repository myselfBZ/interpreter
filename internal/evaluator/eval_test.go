package evaluator
// testing comment
import (
	"testing"

	"github.com/myselfBZ/interpreter/internal/lexer"
	"github.com/myselfBZ/interpreter/internal/object"
	"github.com/myselfBZ/interpreter/internal/parser"
)

func TestEvaluatorInt(t *testing.T){
    input := struct{
        input string
        expected int
    }{
        "6",
        6,
    }
    l := lexer.New(input.input)
    p := parser.New(l)
    program := p.ParseProgram()
    obj := Eval(program)
    i, ok := obj.(*object.Integer)
    if !ok{
        t.Fatalf("expected int got %T\n", obj.(*object.Integer))
    }
    if i.Value != input.expected{
        t.Fatalf("expected %d got %d", input.expected, i.Value)
    }
}

func TestEvalBoolean(t *testing.T) {
    input := struct{
        input string
        expct bool
    }{
        "false;",
        false,
    }
    l := lexer.New(input.input)
    p := parser.New(l)
    program := p.ParseProgram()
    obj := Eval(program)
    b, ok := obj.(*object.Boolean)
    if !ok{
        t.Fatalf("expected boolean got %T", obj)
    }
    if b.Value != input.expct{
        t.Fatalf("expected %v got %v", input.expct, b.Value)
    }
}

