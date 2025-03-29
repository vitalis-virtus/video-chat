package server

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"

	uuid "github.com/google/uuid"
)

type Rooms interface {
	Init()
	GetParticipants(roomID string) []Participant
	CreateRoom() string
	DeleteChannel(chID string) error
}

type Participant interface {
}

type participant struct {
	isHost bool
	id     string
	conn   *websocket.Conn
}

type rooms struct {
	sync.RWMutex
	channels map[string][]Participant
}

func NewRoom() Rooms {
	return &rooms{}
}

func NewParticipant(id string) Participant {
	return &participant{id: id}
}

func (r *rooms) Init() {
	r.channels = make(map[string][]Participant)
}

func (r *rooms) GetParticipants(chID string) []Participant {
	r.RLock()
	defer r.RUnlock()

	return r.channels[chID]
}

func (r *rooms) CreateRoom() string {
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
