package agent

import (
	"errors"
	"net"

	"github.com/go-distributed/gog/message"
	"github.com/go-distributed/gog/node"

	"code.google.com/p/gogoprotobuf/proto"
)

var (
	ErrInvalidMessageType = errors.New("Invalid message type")
)

// disconnect() sends a Disconnect message to the node and close the connection.
// TODO(yifan): cache the connection.
func (ag *agent) disconnect(node *node.Node) {
	msg := &message.Disconnect{Id: proto.String(ag.id)}
	ag.codec.WriteMsg(msg, node.Conn)
	node.Conn.Close()
	node.Conn = nil
}

// forwardJoin() sends a ForwardJoin message to the node. The message
// will include the Id and Addr of the source node, as the receiver might
// use these information to establish a connection.
func (ag *agent) forwardJoin(node, newNode *node.Node, ttl uint32) error {
	msg := &message.ForwardJoin{
		Id:         proto.String(ag.id),
		SourceId:   proto.String(newNode.Id),
		SourceAddr: proto.String(newNode.Addr),
		Ttl:        proto.Uint32(ttl),
	}
	return ag.codec.WriteMsg(msg, node.Conn)
}

// rejectNeighbor() sends a the NeighborReply with accept = false.
func (ag *agent) rejectNeighbor(node *node.Node) error {
	msg := &message.NeighborReply{
		Id:     proto.String(ag.id),
		Accept: proto.Bool(false),
	}
	return ag.codec.WriteMsg(msg, node.Conn)
}

// acceptNeighbor() sends a the NeighborReply with accept = true.
func (ag *agent) acceptNeighbor(node *node.Node) error {
	msg := &message.NeighborReply{
		Id:     proto.String(ag.id),
		Accept: proto.Bool(true),
	}
	return ag.codec.WriteMsg(msg, node.Conn)
}

func (ag *agent) join(node *node.Node) error {
	msg := &message.Join{
		Id:   proto.String(ag.id),
		Addr: proto.String(ag.cfg.AddrStr),
	}
	return ag.codec.WriteMsg(msg, node.Conn)
}

// neighbor() sends a Neighbor message, and wait for the reply.
// If the other side accepts the request, we add the node to the active view.
func (ag *agent) neighbor(node *node.Node, priority message.Neighbor_Priority) (error, bool) {
	if node.Conn == nil {
		addr, err := net.ResolveTCPAddr(ag.cfg.Net, node.Addr)
		if err != nil {
			// TODO(yifan) log.
			return err, false
		}
		conn, err := net.DialTCP(ag.cfg.Net, nil, addr)
		if err != nil {
			// TODO(yifan) log.
			return err, false
		}
		node.Conn = conn
	}
	msg := &message.Neighbor{
		Id:       proto.String(ag.id),
		Addr:     proto.String(ag.cfg.AddrStr),
		Priority: priority.Enum(),
	}
	if err := ag.codec.WriteMsg(msg, node.Conn); err != nil {
		// TODO(yifan) log.
		return err, false
	}
	recvMsg, err := ag.codec.ReadMsg(node.Conn)
	if err != nil {
		// TODO(yifan) log.
		node.Conn.Close()
		return err, false
	}
	reply, ok := recvMsg.(*message.NeighborReply)
	if !ok {
		node.Conn.Close()
		return ErrInvalidMessageType, false
	}
	if reply.GetAccept() {
		ag.addNodeActiveView(node)
		go ag.serveConn(node.Conn)
		return nil, true
	}
	ag.addNodePassiveView(node)
	return nil, false
}

// userMessage() sends a user message to the node.
func (ag *agent) userMessage(node *node.Node, msg proto.Message) error {
	return ag.codec.WriteMsg(msg, node.Conn)
}

func (ag *agent) forwardShuffle(node *node.Node, msg *message.Shuffle) error {
	msg.Id = proto.String(ag.id)
	return ag.codec.WriteMsg(msg, node.Conn)
}

func (ag *agent) shuffleReply(msg *message.Shuffle) error {
	return errors.New("implement this")
}

func (ag *agent) shuffle(node *node.Node) error {
	msg := &message.Shuffle{
		Id:         proto.String(ag.id),
		SourceId:   proto.String(ag.id),
		Addr:       proto.String(ag.cfg.AddrStr),
		Candidates: nil,
		Ttl:        proto.Uint32(uint32(ag.cfg.SRWL)),
	}
	return ag.codec.WriteMsg(msg, node.Conn)
}