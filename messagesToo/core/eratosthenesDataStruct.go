package core

import (

	messagesToo "github.com/EvanFr/messagesToo/messagesToo"
)

const (
	clientId = "clientID"
	deviceId = "deviceId"
	score = "score"
)



type service struct {
	clients map[string]string
}

func NewService() messagesToo.Service {
	return &service{
		clients: map[string]string{
			"one": "good",
			"two": "bad",
			"three": "very bad",
		},
	}
}

/*
var sub *messagesToo.StreamingRawBytes
sub = &messagesToo.StreamingRawBytes{
	Subscribe: &messagesToo.Subscribe{
		ClientId: clientId,
	},
}
*/

func (s *service) SubscribeService(sub  messagesToo.StreamingRawBytes) (subEvent messagesToo.StreamingRawBytes, err error) {
	if _, ok := s.clients[sub.Subscribe.ClientId]; ok{
		return subEvent, nil //TODO na ftiakso to subEvent
	}

	return subEvent, nil 
	//return subEvent, messagesToo.ErrNotFound
}

/*
var ts *messagesToo.StreamingRawBytes
ts = &messagesToo.StreamingRawBytes{
	CreateTrustScore: &messagesToo.CreateTrustScore{
		DeviceID: deviceId,
		Score: score,
	},
}
*/

/*
func (s *service) CreateTrustScoreService(ts *messagesToo.StreamingRawBytes) (tsEvent*messagesToo.StreamingRawBytes, err error) {
	s.clients[ts.CreateTrustScore.DeviceID] = ts.CreateTrustScore.Score

	return tsEvent, nil // TODO na ftiakso to tsEvent
}
*/

/*
rts = *messagesToo.StreamingRawBytes
rts = &messagesToo.StreamingRawBytes{
	ReadTrustScore: &messagesToo.ReadTrustScore{
		DeviceID: deviceId,
	},
}
*/

/*
func (s *service) ReadTrustScoreService(rts *messagesToo.StreamingRawBytes) (rtsEvent *messagesToo.StreamingRawBytes, err error) {
	
	var value string
	var flag bool

	value = "notfound"
	flag = false

	if val, ok := s.clients[rts.ReadTrustScore.DeviceID]; ok{
		value = val
		flag = ok
	}
	
	rtsEvent = &messagesToo.StreamingRawBytes{
		ReadTrustScoreEvent: &messagesToo.ReadTrustScoreEvent{
			Score: value,
		},
	}

	if flag {
		return rtsEvent, nil
	}
	return rtsEvent, nil

	//return rtsEvent, messagesToo.ErrNotFound //TODO
}
*/


