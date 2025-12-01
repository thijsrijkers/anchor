package parser

import(
    "anchor/lexer"
)

type OperationType string

const(
    ASSIGNMENT OperationType = "ASSIGMENT"
    ADDITION OperationType = "ADDITION"
    SUBSTRACTION OperationType = "SUBSTRACTION"
)

type Node Interface{}

type Expression interface{
    Node
}

type TreeNode struct {
    operation OperationType
    left *Expression
    right *Expression
    rightOperation: bool
}

func ParseTokens(tokens []Token) []TreeNode {
    var treeNodes []TreeNode
    var currentTreeNodes []TreeNode

    position := 0
    tokensLength := len(tokens)

    for position < tokensLength {
        if token.Type == "INDETIFER" {
	    hasRightOperation := tokesn[i+2] != TokenType.CLOSING

	    if !hasRightOperation {
	        rightOperation := tokens[i+1].Value
	        position += 2
	    } else {
	        rightOperation := nil
	    }

	    currentTreeNodes = appends(&TreeNode{
	        operation: ASSIGNMENT
		left: token.Value
	        right: rightOperation
		rightOperation: hasRightOperation
	    })
	    
	    if !hasRightOperation {
	        treeNodes := append(treeNodes, currentTreeNodes...)
                currentTreeNodes = []TreeNode
	    }
	}
	
	position++
    }
    return treeNodes
}
