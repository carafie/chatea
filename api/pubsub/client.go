package pubsub

import "sync"

type Client struct {
	topics map[string]*subscription
	mu     sync.RWMutex
}

func NewClient() *Client {
	return &Client{
		topics: make(map[string]*subscription),
	}
}

func (c *Client) Connect(topic string) *Connection {
	c.mu.Lock()
	defer c.mu.Unlock()

	sub, ok := c.topics[topic]
	if !ok {
		sub = &subscription{
			topic: topic,
			conns: make(map[*Connection]struct{}),
		}
		c.topics[topic] = sub
	}

	conn := &Connection{
		client: c,
		sub:    sub,
		ch:     make(chan []byte, 10),
	}
	sub.mu.Lock()
	sub.conns[conn] = struct{}{}
	sub.mu.Unlock()

	return conn
}

type Connection struct {
	client *Client
	sub    *subscription
	ch     chan []byte
	close  sync.Once
}

func (c *Connection) Publish(message []byte) {
	c.sub.mu.RLock()
	defer c.sub.mu.RUnlock()

	for conn := range c.sub.conns {
		select {
		case conn.ch <- message:
		default:
		}
	}
}

func (c *Connection) Subscribe() <-chan []byte {
	return c.ch
}

func (c *Connection) Close() {
	c.close.Do(func() {
		c.client.mu.Lock()
		c.sub.mu.Lock()

		delete(c.sub.conns, c)
		close(c.ch)
		if len(c.sub.conns) == 0 {
			delete(c.client.topics, c.sub.topic)
		}

		c.sub.mu.Unlock()
		c.client.mu.Unlock()
	})
}

type subscription struct {
	topic string
	conns map[*Connection]struct{}
	mu    sync.RWMutex
}
