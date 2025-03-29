package services

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"

	uuid "github.com/google/uuid"
)

type Rooms interface {
	Init()
	GetParticipants(roomID string) []Participant
	CreateChannel() string
	DeleteChannel(chID string) error
	Connect(chID string, participant Participant) error
}

type Participant interface {
	Write(msg map[string]interface{}) error
	ID() string
	Close() error
}

type participant struct {
	isHost bool
	id     string
	ip     string
	name   string
	conn   *websocket.Conn
}

type rooms struct {
	sync.RWMutex
	channels map[string][]Participant
}

func NewRoom() Rooms {
	return &rooms{}
}

func NewParticipant(id, ip, name string, conn *websocket.Conn) Participant {
	return &participant{
		id:   id,
		ip:   ip,
		name: name,
		conn: conn,
	}
}

func (r *rooms) Init() {
	r.channels = make(map[string][]Participant)
}

func (r *rooms) GetParticipants(chID string) []Participant {
	r.RLock()
	defer r.RUnlock()

	return r.channels[chID]
}

func (r *rooms) CreateChannel() string {
	r.Lock()
	defer r.Unlock()

	id := uuid.New()

	r.channels[id.String()] = make([]Participant, 0)

	return id.String()
}

func (r *rooms) Connect(chID string, participant Participant) error {
	r.Lock()
	defer r.Unlock()

	_, ok := r.channels[chID]
	if !ok {
		return fmt.Errorf("channel not exists")
	}

	r.channels[chID] = append(r.channels[chID], participant)

	return nil
}

func (r *rooms) DeleteChannel(chID string) error {
	r.Lock()
	defer r.Unlock()

	_, ok := r.channels[chID]
	if !ok {
		return fmt.Errorf("channel not exists")
	}

	delete(r.channels, chID)

	return nil
}

func (r *participant) Write(msg map[string]interface{}) error {
	return r.conn.WriteJSON(msg)
}

func (r *participant) ID() string {
	return r.id
}

func (r *participant) Close() error {
	return r.conn.Close()
}
