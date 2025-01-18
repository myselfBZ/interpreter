package evaluator

import (
	"github.com/myselfBZ/interpreter/internal/ast"
	"github.com/myselfBZ/interpreter/internal/object"
)

var(
    NULL = &object.Null{}
    TRUE = &object.Boolean{Value: true}
    FALSE = &object.Boolean{Value: false}
)


func boolToBoolOBJ(b bool) *object.Boolean{
    if b {
        return TRUE
    }
    return FALSE
}

func Eval(node ast.Node) object.Object{
    switch node := node.(type){
    case *ast.Program:
        return evalStatements(node.Statements)
    case *ast.IntLiteral:
        return &object.Integer{Value: int(node.Value)}
    case *ast.ExpressionStatement:
        return Eval(node.Expression)
    case *ast.Boolean:
        if node.Value {
            return TRUE
        }
        return FALSE
    case *ast.PrefixExpression:
        return evalPrefix(node, node.Operator)
    case *ast.InfixExperssion:
        right := Eval(node.Right)
        left := Eval(node.Left)
        return evalInfix(right, left, node.Operator)
    default:
        return NULL
    }
}


func evalPrefix(node *ast.PrefixExpression, op string) object.Object{
    v := Eval(node.Right)
    switch op{
    case "!":
        return evalBang(v)
    case "-":
        return evalMinus(v)
    default:
        return NULL
    }
}

func evalBang(o object.Object) object.Object {
    switch o {
    case TRUE:
        return FALSE
    case FALSE:
        return TRUE
    case NULL:
        return TRUE
    default:
        return FALSE
    } 
}


func evalMinus(o object.Object) object.Object {
    if o.Type() != object.INTEGER_OBJ{
        return NULL
    }
    val := o.(*object.Integer).Value
    return &object.Integer{Value:-val}
}

func evalStatements(s []ast.Statement) object.Object{
   var result object.Object
   for _, v := range s{
       result = Eval(v)
   }
   return result
}

func evalIntInfix(right object.Object, left object.Object, oprtr string) object.Object{
    leftValue := left.(*object.Integer).Value
    rightValue := right.(*object.Integer).Value
    switch oprtr{
    case "+":
        return &object.Integer{Value: leftValue + rightValue}
    case "-":
        return &object.Integer{Value: leftValue - rightValue}
    case "*":
        return &object.Integer{Value: leftValue * rightValue}
    case "/":
        return &object.Integer{Value: leftValue / rightValue}
    case "==":
        return boolToBoolOBJ(leftValue == rightValue)
    case "!=":
        return boolToBoolOBJ(leftValue != rightValue)
    case ">=":
        return boolToBoolOBJ(leftValue >= rightValue)
    case "<=":
        return boolToBoolOBJ(leftValue <= rightValue)
    case ">":
        return boolToBoolOBJ(leftValue > rightValue)
    case "<":
        return boolToBoolOBJ(leftValue < rightValue)
    default:
        return NULL
    }
}

func evalInfix(right object.Object, left object.Object, oprtr string) object.Object{
    if right.Type() == object.INTEGER_OBJ && left.Type() == object.INTEGER_OBJ{
        return evalIntInfix(right, left, oprtr)
    }
    return evalBoolInfix(right, left, oprtr)
}


func compareBool(right object.Object, left object.Object, oprtr string) object.Object{
    leftValue := left.(*object.Boolean).Value
    rightValue := right.(*object.Boolean).Value
    switch oprtr{
    case "==":
        return &object.Boolean{Value: leftValue == rightValue}
    case "!=":
        return &object.Boolean{Value: leftValue != rightValue}
    default:
        return NULL
    }
}


func evalBoolInfix(right object.Object, left object.Object, oprtr string) object.Object{
    if right.Type() != left.Type() {
        return NULL
    } 
    if right.Type() == object.BOOLEAN_OBJ{
        return compareBool(right, left, oprtr)
    }
    return NULL
}

