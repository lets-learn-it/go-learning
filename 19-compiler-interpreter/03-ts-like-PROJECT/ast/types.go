package ast

type NodeType string

const (
	Program             NodeType = "Program"
	NumericLiteral      NodeType = "NumericLiteral"
	Identifier          NodeType = "Identifier"
	BinaryExpr          NodeType = "BinaryExpr"
	CallExpr            NodeType = "CallExpr"
	UnaryExpr           NodeType = "UnaryExpr"
	FunctionDeclaration NodeType = "FunctionDeclaration"
)

type Stmt struct {
	Kind NodeType
}

type ProgramStmt struct {
	Kind NodeType
	Body []Stmt
}

type Expr struct {
	Kind NodeType
}

type BinaryExprExpr struct {
	Kind  NodeType
	Left  Expr
	Right Expr
	Op    string
}

type IdentifierExpr struct {
	Kind   NodeType
	Symbol string
}
