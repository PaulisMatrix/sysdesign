package ticketing_system

import "strings"

type Parser struct {
}

func (p *Parser) ParseCommand(userInput string) OperationResult {
	// split the user input by space
	// command_type, first_field, second_field
	// first_field, second_field depends on the command_type
	command := strings.Split(userInput, " ")[0]
	cmd := CommandType(command)
	v := &Verifier{
		commandType: cmd,
	}

	switch cmd {
	case CREATETICKET:
		//create-ticket
		result := v.VerifyCreateTicket(userInput)
		return OperationResult{Comment: result}
	case ASSIGNTICKET:
		//assign-ticket
		result := v.VerifyAssignTicket(userInput)
		return OperationResult{Comment: result}
	case RESOLVETICKET:
		//resolve-ticket
		result := v.VerifyResolveTicket(userInput)
		return OperationResult{Comment: result}
	case VERIFYTICKET:
		//verify-ticket-resolution
		result := v.VerifyResolutionTicket(userInput)
		return OperationResult{Comment: result}
	case STATUS:
		//status
		result := v.VerifyStatus(userInput)
		return OperationResult{Comment: result}
	default:
		return OperationResult{Comment: "unkown command. pls check for any typos."}
	}
}
