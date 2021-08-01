package messages

const (
	AttributeNotFound = "Attribute not found!!"
)

var (
	Message MessageInterface = &Messages{}
)

type MessageInterface interface {
	//
}

type Messages struct {
	//
}
