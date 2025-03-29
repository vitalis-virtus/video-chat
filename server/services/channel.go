package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/vitalis-virtus/video-chat/models"
)

type broadcastMsg struct {
	Message   map[string]interface{}
	ChannelID string
	SenderID  string
	Client    *websocket.Conn
}

func (s *service) broadcaster() {
	for {
		msg := <-s.broadcast
		for _, client := range s.rooms.GetParticipants(msg.ChannelID) {

			if client.ID() != msg.SenderID {
				err := client.Write(msg.Message)

				if err != nil {
					log.Println(err)

					log.Println(client.Close())
				}
			}
		}
	}
}

func (s *service) CreateChannel() string {
	id := s.rooms.CreateChannel()

	return id
}

func (s *service) JoinChannel(conn *websocket.Conn, joinParams *models.JoinChannelQuery) {
	participantID := uuid.New()

	participant := NewParticipant(participantID.String(), joinParams.IP, joinParams.Name, conn)

	err := s.rooms.Connect(joinParams.ChannelID, participant)

	if err != nil {
		log.Println(err)
		conn.Close()
	}

	go s.broadcaster()

	for {
		var msg broadcastMsg

		err := conn.ReadJSON(&msg.Message)
		if err != nil {
			log.Fatal("Read Error:", err)
		}

		msg.Client = conn
		msg.ChannelID = joinParams.ChannelID
		msg.SenderID = participantID.String()

		// log.Printf("send msg {%v} from %v", msg.Message, participantID.String())

		s.broadcast <- msg
	}
}
