package query

import (
	"fmt"
	"strings"
)

type Type int
type Operator int

type Query struct {
	Type       Type
	TableName  string
	Conditions []Condition
	Updates    map[string]string
	Inserts    [][]string
	Fields     []string
	Aliases    map[string]string
}

type Condition struct {
	OperandLeft         string
	OperandLeftIsField  bool
	Operator            Operator
	OperandRight        string
	OperandRightIsField bool
}

func (Statement Statement) AddStatementToExecutionPlan() int {
	return 0
}

type StructuredQuery struct{}

type OperationType struct {
	statement string
	flag      string
}

type Statement struct {
	operation OperationType
}

type ExecutionResult struct {
	status string
	result string
}

type ExecutionTreeDump struct{}

func (query StructuredQuery) dumpExecutionTree() ExecutionTreeDump {
	return ExecutionTreeDump{}
}

func (r ExecutionResult) getDebugMessage() string {
	return r.status
}

func syntaxAnalysis(query string) (n int, err error) {
	var tempStatement Statement
	statements := strings.Split(query, " ")
	for _, statement := range statements {
		switch statement {
		case "SELECT":
			tempStatement = Statement{operation: OperationType{flag: "SELECT", statement: statement}}
		case "INSERT":
			tempStatement = Statement{operation: OperationType{flag: "INSERT", statement: statement}}
		case "UPDATE":
			tempStatement = Statement{operation: OperationType{flag: "UPDATE", statement: statement}}
		case "DELETE":
			tempStatement = Statement{operation: OperationType{flag: "DELETE", statement: statement}}
		case "CREATE":
			tempStatement = Statement{operation: OperationType{flag: "CREATE", statement: statement}}
		}
		tempStatement.AddStatementToExecutionPlan()
	}
	return len(query), nil
}

func parseQuery(query string) StructuredQuery {
	fmt.Println("DEBUG: Parsed query ----- ", query)
	return StructuredQuery{}
}

func executeQuery(query StructuredQuery) (error, ExecutionResult) {
	// TODO set up logging system.
	fmt.Println("DEBUG: Execution in progress ...", query)
	return nil, ExecutionResult{status: "success"}
}

type ResponseQuery struct {
	status uint8
	result string
}

func (response ResponseQuery) GetResponseResult() string {
	return response.result
}

func defaultResponseQuery() ResponseQuery {
	return ResponseQuery{}
}

func HandleQuery(query string) ResponseQuery {
	// TODO set up logging system.
	if _, err := syntaxAnalysis(query); err == nil {
		structuredQuery := parseQuery(query)
		possibleExecutionError, executionResult := executeQuery(structuredQuery)
		if possibleExecutionError != nil {
			fmt.Println(possibleExecutionError)
		} else {
			fmt.Println(executionResult.getDebugMessage())
		}
	} else {
		fmt.Println(err)
	}
	return defaultResponseQuery()
}
