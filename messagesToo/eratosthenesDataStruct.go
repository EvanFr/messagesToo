package request

import (
	status "google.golang.org/genproto/googleapis/rpc/status" 	// GRPC status
	did "github.com/hyperledger/aries-framework-go/pkg/doc/did"	// Aries Framework GO
)

type StreamingRawBytes struct {
	Subscribe             *Subscribe             `json:"subscribe,omitempty"`
	SubscribeEvent        *SubscribeEvent        `json:"subscribeEvent,omitempty"`
	CreateDoc             *CreateDoc	         `json:"createDoc,omitempty"`
	CreateDocEvent        *CreateDocEvent        `json:"createDocEvent,omitempty"`
	ResolveDID            *ResolveDID	         `json:"resolveDID,omitempty"`
	ResolveDIDEvent       *ResolveDIDEvent       `json:"resolveDIDEvent,omitempty"`
	DisseminateDID        *DisseminateDID        `json:"disseminateDID,omitempty"`
	CreateTrustScore      *CreateTrustScore      `json:"createTrustScore,omitempty"`
	CreateTrustScoreEvent *CreateTrustScoreEvent `json:"createTrustScoreEvent,omitempty"`
	ReadTrustScore	      *ReadTrustScore	     `json:"readTrustScore,omitempty"`
	ReadTrustScoreEvent   *ReadTrustScoreEvent   `json:"readTrustScoreEvent,omitempty"`
	Error                 *status.Status         `json:"error,omitempty"`
}

/* Subscribe Request */
type Subscribe struct {
	ClientId string
}

type SubscribeEvent struct{
	Event string
}

/* CreateDoc Request */
type CreateDoc struct {
	Doc *did.Doc
}

type CreateDocEvent struct{
	DocResolution *did.DocResolution
}

/* ResolveDID Request */
type ResolveDID struct {
	DocId string
}

type ResolveDIDEvent struct {
	DocResolution *did.DocResolution
}

/* Disseminate to TMB */
type DisseminateDID struct {
	DID string
}

/* Create Trust Score */
type CreateTrustScore struct {
	DeviceID string
	Score    string
}

type CreateTrustScoreEvent struct {
	Event string
}

/* Create Trust Score */
type ReadTrustScore struct {
	DeviceID string
}

type ReadTrustScoreEvent struct {
	Score string
}

type Service interface {
	SubscribeService(StreamingRawBytes) (StreamingRawBytes, error)
	//CreateTrustScoreService(StreamingRawBytes) (StreamingRawBytes, error)
	//ReadTrustScoreService(StreamingRawBytes) (StreamingRawBytes, error)
}

type Client struct {
	ClientId string
	Score string
}

