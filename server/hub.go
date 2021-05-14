package server

import (
	"encoding/json"
)

var H = hub{
	c: make(map[*connection]bool),
	u: make(chan *connection),
	b: make(chan []byte),
	r: make(chan *connection),
}

type hub struct {
	c map[*connection]bool
	b chan []byte
	r chan *connection
	u chan *connection
}

func (H *hub) Run() {
	for {
		select {
		case c := <-H.r:
			H.c[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			c.sc <- data_b
		case c := <-H.u:
			if _, ok := H.c[c]; ok {
				delete(H.c, c)
				close(c.sc)
			}
		case data := <-H.b:
			for c := range H.c {
				select {
				case c.sc <- data:
				default:
					delete(H.c, c)
					close(c.sc)
				}
			}
		}
	}
}