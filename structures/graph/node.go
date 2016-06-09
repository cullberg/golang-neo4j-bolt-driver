package messages

const (
	// NodeSignature is the signature byte for a Node object
	NodeSignature = 0x4E
)

// Node Represents a Node structure
type Node struct {
	NodeIdentity int
	Labels       []string
	Properties   map[string]interface{}
}

// Signature gets the signature byte for the struct
func (n Node) Signature() int {
	return NodeSignature
}
