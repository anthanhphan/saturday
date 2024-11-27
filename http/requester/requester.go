package requester

// CtxRequester defines an interface for accessing request context information.
type CtxRequester interface {
	GetUserId() any
}

var _ CtxRequester = (*ctxRequester)(nil)

type ctxRequester struct {
	UserId any
}

// NewCtxRequester creates a new CtxRequester instance with the specified user ID.
//
// Parameters:
//   - userId: The user's unique identifier of any type
//
// Returns:
//   - CtxRequester: A new instance implementing the CtxRequester interface
//
// Example:
//
//	requester := NewCtxRequester(123)
//	userId := requester.GetUserId() // returns 123
func NewCtxRequester(userId any) CtxRequester {
	return &ctxRequester{
		UserId: userId,
	}
}

// GetUserId returns the ID of the authenticated user.
//
// Returns:
//   - any: The user's unique identifier of any type
//
// Example:
//
//	requester := NewCtxRequester("user-123")
//	userId := requester.GetUserId() // returns "user-123"
func (r *ctxRequester) GetUserId() any {
	return r.UserId
}
