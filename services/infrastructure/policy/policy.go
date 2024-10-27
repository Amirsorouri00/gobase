package policy

import (
	"net"
	"net/rpc"
)

// Handler is an interface that defines the policy handler.
type Handler interface {

	// Handle returns whether the given action is allowed.
	Handle(actor, action string) (bool, error)
}

// HandlerFunc is a function that determines whether an action is allowed.
type HandlerFunc func(actor, action string) (bool, error)

// Handle returns whether the given action is allowed.
func (h HandlerFunc) Handle(actor, action string) (bool, error) {
	return h(actor, action)
}

// Server is a policy server that can be used to enforce policies.
type Server struct {

	// Addr is the address to listen on.
	Addr string

	// Handler is the policy handler.
	Handler Handler
}

type Request struct {
	Actor  string
	Action string
}

type Response struct {
	Allowed bool
}

type rpcReceiver struct {
	Handler Handler
}

// Can returns whether the given action is allowed.
func (h *rpcReceiver) Can(req *Request, resp *Response) error {
	allowed, err := h.Handler.Handle(req.Actor, req.Action)
	if err != nil {
		return err
	}
	resp.Allowed = allowed
	return nil
}

// Run starts the policy server.
func (s *Server) Run() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	defer listener.Close()
	server := rpc.NewServer()
	server.RegisterName("Policy", &rpcReceiver{Handler: s.Handler})
	server.Accept(listener)
	return nil
}

// Client is a policy client that can be used to enforce policies.
type Client struct {

	// Addr is the address to connect to.
	Addr string
}

// Can returns whether the given action is allowed.
func (c *Client) Can(actor, action string) (bool, error) {
	client, err := rpc.Dial("tcp", c.Addr)
	if err != nil {
		return false, err
	}
	defer client.Close()
	req := &Request{Actor: actor, Action: action}
	resp := &Response{}
	if err := client.Call("Policy.Can", req, resp); err != nil {
		return false, err
	}
	return resp.Allowed, nil
}
