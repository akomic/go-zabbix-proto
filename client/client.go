package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"
)

// ProxyClient
type Client struct {
	Host string
	Port int
}

// Client constructor.
func NewClient(host string, port int) (c *Client) {
	c = &Client{Host: host, Port: port}
	return
}

// Client method, return Zabbix header.
func (c *Client) getHeader() []byte {
	return []byte("ZBXD\x01")
}

// Client method, resolve uri by name:port.
func (c *Client) getTCPAddr() (iaddr *net.TCPAddr, err error) {
	// format: hostname:port
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)

	// Resolve hostname:port to ip:port
	iaddr, err = net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		err = fmt.Errorf("Connection failed: %s", err.Error())
		return
	}

	return
}

// Client method, make connection to uri.
func (c *Client) connect() (conn *net.TCPConn, err error) {

	type DialResp struct {
		Conn  *net.TCPConn
		Error error
	}

	// Open connection to Zabbix host
	iaddr, err := c.getTCPAddr()
	if err != nil {
		return
	}

	// dial tcp and handle timeouts
	ch := make(chan DialResp)

	go func() {
		conn, err = net.DialTCP("tcp", nil, iaddr)
		ch <- DialResp{Conn: conn, Error: err}
	}()

	select {
	case <-time.After(5 * time.Second):
		err = fmt.Errorf("Connection Timeout")
	case resp := <-ch:
		if resp.Error != nil {
			err = resp.Error
			break
		}

		conn = resp.Conn
	}

	return
}

// Client method, read data from connection.
func (c *Client) read(conn *net.TCPConn) (res []byte, err error) {
	res = make([]byte, 1024)
	res, err = ioutil.ReadAll(conn)
	if err != nil {
		err = fmt.Errorf("Error while receiving the data: %s", err.Error())
		return
	}

	return
}

// Client method, sends packet to Zabbix.
func (c *Client) Send(packet *Packet) (res []byte, err error) {
	log.Println("Sending packet with request:", packet.Request)
	conn, err := c.connect()
	if err != nil {
		return
	}
	defer conn.Close()

	// Fill buffer
	buffer := append(c.getHeader(), packet.DataLen()...)
	buffer = append(buffer, packet.Data...)

	// Send packet to Zabbix
	_, err = conn.Write(buffer)
	if err != nil {
		err = fmt.Errorf("Error while sending the data: %s", err.Error())
		return
	}

	res, err = c.read(conn)

	return
}
