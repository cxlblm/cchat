package websocket

import (
	"bytes"
	"github.com/gorilla/websocket"
	"io"
	"log/slog"
	"time"
	"unsafe"
)

type Client struct {
	conn *websocket.Conn
	send chan string
	hub  *Hub
	log  *slog.Logger
	name string
}

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (s *Client) read() {

	defer func() {
		s.hub.unregister <- s
		err := s.conn.Close()
		if err != nil {
			s.log.Error("conn close error: " + err.Error())
		}
	}()

	s.conn.SetReadLimit(maxMessageSize)
	s.conn.SetReadDeadline(time.Now().Add(pongWait))
	s.conn.SetPongHandler(func(string) error {
		err := s.conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			return err
		}
		return nil
	})

	for {
		_, message, err := s.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				s.log.Error("websocket read error: " + err.Error())
			}
			s.log.Error("websocket read error: " + err.Error())
			break
		}

		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		messageStr := s.name + ": " + string(message)
		s.hub.broadcast <- messageStr
	}
}

func writeToConn(conn io.WriteCloser, m string) (int, error) {
	s := unsafe.Slice(unsafe.StringData(m), len(m))
	return conn.Write(s)
}

func (s *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		err := s.conn.Close()
		if err != nil {
			s.log.Error("websocket close error: " + err.Error())
		}
	}()

	for {
		select {
		case message, ok := <-s.send:
			s.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				s.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := s.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				s.log.Error("next writer error: " + err.Error())
				return
			}
			w.Write([]byte(message))

			n := len(s.send)

			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write([]byte(<-s.send))
			}

			if err := w.Close(); err != nil {
				s.log.Error("websocket close error: " + err.Error())
				return
			}
		case <-ticker.C:
			s.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := s.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}

}
