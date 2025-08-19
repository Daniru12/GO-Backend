package error_handler


type AppErrorInterface interface {
	Error() string      
	Status() int         
	Code() int           
	Details() interface{} 
	Type() string       
}

type DomainError struct {
	Message    string
	CodeValue  int
	DetailsVal interface{}
}

func NewDomainError(message string, code int, details interface{}) *DomainError {
	return &DomainError{
		Message:    message,
		CodeValue:  code,
		DetailsVal: details,
	}
}

func (e *DomainError) Error() string {
	return e.Message
}

func (e *DomainError) Status() int {
	return 400 
}

func (e *DomainError) Code() int {
	return e.CodeValue
}

func (e *DomainError) Details() interface{} {
	return e.DetailsVal
}

func (e *DomainError) Type() string {
	return "DomainError"
}

type ApplicationError struct {
	Message    string
	CodeValue  int
	StatusCode int
	DetailsVal interface{}
}

func NewApplicationError(message string, statusCode int, code int, details interface{}) *ApplicationError {
	return &ApplicationError{
		Message:    message,
		StatusCode: statusCode,
		CodeValue:  code,
		DetailsVal: details,
	}
}

func (e *ApplicationError) Error() string {
	return e.Message
}

func (e *ApplicationError) Status() int {
	return e.StatusCode
}

func (e *ApplicationError) Code() int {
	return e.CodeValue
}

func (e *ApplicationError) Details() interface{} {
	return e.DetailsVal
}

func (e *ApplicationError) Type() string {
	return "ApplicationError"
}

func IsDomain(err error) bool {
	_, ok := err.(*DomainError)
	return ok
}

func IsApplication(err error) bool {
	_, ok := err.(*ApplicationError)
	return ok
}
