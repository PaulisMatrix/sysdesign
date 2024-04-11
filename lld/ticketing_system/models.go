package ticketing_system

type TicketType string
type TicketStatus string

const (
	WalletBalance TicketType = "check-wallet-balance"
	ChangeLang    TicketType = "change-language"
	Others        TicketType = "others"
)

const (
	OPEN     TicketStatus = "OPEN"
	ASSIGNED TicketStatus = "ASSIGNED"
	CLOSED   TicketStatus = "CLOSED"
)

type Ticket struct {
	TicketNo       int
	Type           TicketType
	Description    string
	IsManualAction bool
	Resolution     string
	TicketStatus   TicketStatus
	ResolvedBy     Employee
	VerifiedBy     Supervisor
}

// maintain a list of all tickets
var Tickets []Ticket

// each ticket type has an action associated with it
type AutoActions map[TicketType]string

var autoActions = AutoActions{
	WalletBalance: "sent automatic SMS to customer",
	ChangeLang:    "automatic IVR call made to the customer",
	Others:        "add to the queue",
}

// ticket to be assigned by an action assign-ticket, no ticket is automatically assigned
// Employee to handle the ticket and Supervisor to verify the ticket
// For employee, ticket is assigned from the open-tickets bucket
// For supervisor, ticket is assigned from verify-resolved-ticket bucket

type Employee struct {
	Name string
}

type Supervisor struct {
	Name string
}

// map to check if the person handling ticket is supervisor or an employee
var emps = []string{"Sam", "Tom"}
var supervisors = []string{"Harry"}

// all commands
// 1. create-ticket <ticket_type> <ticket_desc>
// 2. status <ticket_no>
// 3. assign-ticket <emp_name>
// 4. resolve-ticket <emp_name> <resolution_desc>
// 5. verify-ticket-resolution <supervisor_name>
// 6. status (if not ticket no is specified, display all tickets)

type CommandType string

const (
	CREATETICKET  CommandType = "create-ticket"
	ASSIGNTICKET  CommandType = "assign-ticket"
	RESOLVETICKET CommandType = "resolve-ticket"
	VERIFYTICKET  CommandType = "verify-ticket-resolution"
	STATUS        CommandType = "status"
)

// comment to be shown to the end user after each input execution
type OperationResult struct {
	Comment string
}
