package core

import (
	"errors"

	"example.com/go-messages-grpc/messagesToo"
)

const (
	clientId = "clientID"
	deviceId = "deviceId"
	score = "score"
)



type service struct {
	clients map[messagesToo.Client.ClientId]messagesToo.Client.Score
}

func DataBase(
	return &service{
		clients: map[messagesToo.Client.ClientId]messagesToo.Client.Score{
			"one": "good",
			"two": "bad",
			"three": "very bad",
		}
	}
)

var sub *messagesToo.StreamingRawBytes
sub = &messagesToo.StreamingRawBytes{
	Subscribe: &messagesToo.Subscribe{
		ClientId: clientId,
	},
}

func (s *service) SubscribeService(sub  *StreamingRawBytes) (subEvent *StreamingRawBytes, err error) {
	if val, ok := s.clients[sub.Subscribe.ClientId]; ok{
		return subEvent, nil //TODO na ftiakso to subEvent
	}

	return subEvent, messagesToo.ErrNotFound
}

var ts *messagesToo.StreamingRawBytes
ts = &messagesToo.StreamingRawBytes{
	CreateTrustScore: &messagesToo.CreateTrustScore{
		DeviceID: deviceId,
		Score: score,
	},
}

func (s *service) CreateTrustScoreService(ts *StreamingRawBytes) (tsEvent *StreamingRawBytes, err error) {
	s.clients[ts.CreateTrustScore.DeviceID] = ts.CreateTrustScore.Score

	retunr tsEvent, nil // TODO na ftiakso to tsEvent
}

rts = *messagesToo.StreamingRawBytes
rts = &messagesToo.StreamingRawBytes{
	ReadTrustScore: &messagesToo.ReadTrustScore{
		DeviceID: deviceId,
	},
}

func (s *service) ReadTrustScoreService(rts *StreamingRawBytes) (rtsEvent *ReadTrustScoreEvent, err error) {
	
	if val, ok := s.clients[rts.ReadTrustScore.DeviceID]; ok
	
	rtsEvent = *messagesToo.StreamingRawBytes
	rtsEvent = &messagesToo.StreamingRawBytes{
		ReadTrustScoreEvent: &messagesToo.ReadTrustScoreevent{
			Score val//TODO na fiakso san diktionary to service
		}
	}

	if ok {
		return rtsEvent, nil
	}
	return rtsEvent, messagesToo.ErrNotFound
}



