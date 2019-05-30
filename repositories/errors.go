package repositories

type NotFoundError struct {
	message string
}

func (n *NotFoundError) Error() string {
	return n.message
}
