package main

import "fmt"

type StructuredQuery struct {
}

type ExecutionResult struct {
	status string
}

func (r ExecutionResult) getDebugMessage() string {
	return r.status
}

func syntaxAnalysis(query string) (err error, n int) {
	return nil, len(query)
}

func parseQuery(query string) StructuredQuery {
	fmt.Println("DEBUG: Parsed query ----- ", query)
	return StructuredQuery{}
}

func executeQuery(query StructuredQuery) (error, ExecutionResult) {
	fmt.Println("DEBUG: Execution in progress ...", query)
	return nil, ExecutionResult{status: "success"}
}

func handleQuery(query string) {
	fmt.Println("processing query ...", query)
	if err, queryLength := syntaxAnalysis(query); err == nil {
		fmt.Println("Length of query =", queryLength)
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
}
