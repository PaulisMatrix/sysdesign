package ticketing_system

type Verifier struct {
	commandType CommandType
}

func (v *Verifier) VerifyCreateTicket(userInput string) string {
	return ""
}

func (v *Verifier) VerifyAssignTicket(userInput string) string {
	return ""
}

func (v *Verifier) VerifyResolveTicket(userInput string) string {
	return ""
}

func (v *Verifier) VerifyResolutionTicket(userInput string) string {
	return ""
}

func (v *Verifier) VerifyStatus(userInput string) string {
	return ""
}
