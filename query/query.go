package query

import (
	"fmt"
	"strings"
)

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

func (r ExecutionResult) getDebugMessage() string {
	return r.status
}

func syntaxAnalysis(query string) (err error, n int) {
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
	return nil, len(query)
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
	if err, _ := syntaxAnalysis(query); err == nil {
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
