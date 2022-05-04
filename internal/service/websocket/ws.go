package websocket

import (
	"fmt"
	"github.com/wujunyi792/where-it-go/internal/logger"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var manager *socketManager

const (
	writeWait = 10 * time.Second
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type socketManager struct {
	clients    map[*socketClient]bool
	register   chan *socketClient
	unregister chan *socketClient
	receive    chan map[*socketClient][]byte
	broadcast  chan []byte
}

/**
 * @description: 创建Socket管理器
 */
func newManager() *socketManager {
	return &socketManager{
		clients:    make(map[*socketClient]bool),
		register:   make(chan *socketClient),
		unregister: make(chan *socketClient),
		receive:    make(chan map[*socketClient][]byte),
		broadcast:  make(chan []byte), // 广播
	}
}

/**
 * @description: 接收Socket连接
 */
func (m *socketManager) run() {
	for {
		select {
		// 连接加入
		case client := <-m.register:
			// 设置clinet为true
			m.clients[client] = true
		// 连接关闭
		case client := <-m.unregister:
			if _, ok := m.clients[client]; ok {
				delete(m.clients, client)
				close(client.send)
			}
		// 发送message
		case message := <-m.broadcast:
			for client := range m.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(m.clients, client)
				}
			}
		}
	}
}

type socketClient struct {
	name    string
	manager *socketManager
	conn    *websocket.Conn
	send    chan []byte
}

/**
 * @description: 接收Socket信息
 */
func (c *socketClient) readPump() {
	defer func() {
		c.manager.unregister <- c
		err := c.conn.Close()
		if err != nil {
			logger.Error.Println(err)
		}
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			logger.Error.Println(fmt.Sprintf("error: %v", err))
			break
		}
		c.manager.receive <- map[*socketClient][]byte{
			c: message,
		}
	}
}

/**
 * @description: 发送Socket信息
 */
func (c *socketClient) writePump() {
	defer func() {
		err := c.conn.Close()
		if err != nil {
			logger.Error.Println(err)
		}
	}()
	for {
		message, ok := <-c.send
		c.conn.SetWriteDeadline(time.Now().Add(writeWait))
		if !ok {
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		w, err := c.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			return
		}
		w.Write(message)

		// Add queued chat messages to the current websocket message.
		n := len(c.send)
		for i := 0; i < n; i++ {
			w.Write(<-c.send)
		}

		if err := w.Close(); err != nil {
			return
		}
	}
}

/**
 * @description: 发送Socket信息
 * @param {string} name
 * @param {string} message
 */
func SendClientSocket(name string, message string) {
	for k := range manager.clients {
		if k.name == name {
			k.send <- []byte(message)
		}
	}
}

func SendAllSocket(message string) {
	manager.broadcast <- []byte(message)
}

func SocketServer(w http.ResponseWriter, r *http.Request, n string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &socketClient{name: n, manager: manager, conn: conn, send: make(chan []byte, 256)}
	client.manager.register <- client
	go client.writePump()
	go client.readPump()
}

func SocketInit() {
	manager = newManager()
	go manager.run()
}
