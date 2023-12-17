package websocket

type MulticastHub struct {
	broadcast  chan string
	register   chan *Client
	unregister chan *Client
	multicast  []*Multicast
}

type Multicast struct {
	id      int
	clients []*Client
}

func newMulticastHub() *MulticastHub {
	return &MulticastHub{
		broadcast:  make(chan string),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		multicast:  make([]*Multicast, 100),
	}
}

func (s *MulticastHub) run() {
}
