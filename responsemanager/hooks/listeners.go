package hooks

import (
	"github.com/hannahhoward/go-pubsub"
	peer "github.com/libp2p/go-libp2p-core/peer"

	"github.com/ipfs/go-graphsync"
)

// CompletedResponseListeners is a set of listeners for completed responses
type CompletedResponseListeners struct {
	pubSub *pubsub.PubSub
}

type internalCompletedResponseEvent struct {
	p       peer.ID
	request graphsync.RequestData
	status  graphsync.ResponseStatusCode
}

func completedResponseDispatcher(event pubsub.Event, subscriberFn pubsub.SubscriberFn) error {
	ie := event.(internalCompletedResponseEvent)
	listener := subscriberFn.(graphsync.OnResponseCompletedListener)
	listener(ie.p, ie.request, ie.status)
	return nil
}

// NewCompletedResponseListeners returns a new list of completed response listeners
func NewCompletedResponseListeners() *CompletedResponseListeners {
	return &CompletedResponseListeners{pubSub: pubsub.New(completedResponseDispatcher)}
}

// Register registers an listener for completed responses
func (crl *CompletedResponseListeners) Register(listener graphsync.OnResponseCompletedListener) graphsync.UnregisterHookFunc {
	return graphsync.UnregisterHookFunc(crl.pubSub.Subscribe(listener))
}

// NotifyCompletedListeners runs notifies all completed listeners that a response has completed
func (crl *CompletedResponseListeners) NotifyCompletedListeners(p peer.ID, request graphsync.RequestData, status graphsync.ResponseStatusCode) {
	_ = crl.pubSub.Publish(internalCompletedResponseEvent{p, request, status})
}

// RequestorCancelledListeners is a set of listeners for when requestors cancel
type RequestorCancelledListeners struct {
	pubSub *pubsub.PubSub
}

type internalRequestorCancelledEvent struct {
	p       peer.ID
	request graphsync.RequestData
}

func requestorCancelledDispatcher(event pubsub.Event, subscriberFn pubsub.SubscriberFn) error {
	ie := event.(internalRequestorCancelledEvent)
	listener := subscriberFn.(graphsync.OnRequestorCancelledListener)
	listener(ie.p, ie.request)
	return nil
}

// NewRequestorCancelledListeners returns a new list of listeners for when requestors cancel
func NewRequestorCancelledListeners() *RequestorCancelledListeners {
	return &RequestorCancelledListeners{pubSub: pubsub.New(requestorCancelledDispatcher)}
}

// Register registers an listener for completed responses
func (rcl *RequestorCancelledListeners) Register(listener graphsync.OnRequestorCancelledListener) graphsync.UnregisterHookFunc {
	return graphsync.UnregisterHookFunc(rcl.pubSub.Subscribe(listener))
}

// NotifyCancelledListeners notifies all listeners that a requestor cancelled a response
func (rcl *RequestorCancelledListeners) NotifyCancelledListeners(p peer.ID, request graphsync.RequestData) {
	_ = rcl.pubSub.Publish(internalRequestorCancelledEvent{p, request})
}

// BlockSentListeners is a set of listeners for when requestors cancel
type BlockSentListeners struct {
	pubSub *pubsub.PubSub
}

type internalBlockSentEvent struct {
	p       peer.ID
	request graphsync.RequestData
	block   graphsync.BlockData
}

func blockSentDispatcher(event pubsub.Event, subscriberFn pubsub.SubscriberFn) error {
	ie := event.(internalBlockSentEvent)
	listener := subscriberFn.(graphsync.OnBlockSentListener)
	listener(ie.p, ie.request, ie.block)
	return nil
}

// NewBlockSentListeners returns a new list of listeners for when requestors cancel
func NewBlockSentListeners() *BlockSentListeners {
	return &BlockSentListeners{pubSub: pubsub.New(blockSentDispatcher)}
}

// Register registers an listener for completed responses
func (bsl *BlockSentListeners) Register(listener graphsync.OnBlockSentListener) graphsync.UnregisterHookFunc {
	return graphsync.UnregisterHookFunc(bsl.pubSub.Subscribe(listener))
}

// NotifyBlockSentListeners notifies all listeners that a requestor cancelled a response
func (bsl *BlockSentListeners) NotifyBlockSentListeners(p peer.ID, request graphsync.RequestData, block graphsync.BlockData) {
	_ = bsl.pubSub.Publish(internalBlockSentEvent{p, request, block})
}

// NetworkErrorListeners is a set of listeners for when requestors cancel
type NetworkErrorListeners struct {
	pubSub *pubsub.PubSub
}

type internalNetworkErrorEvent struct {
	p       peer.ID
	request graphsync.RequestData
	err     error
}

func networkErrorDispatcher(event pubsub.Event, subscriberFn pubsub.SubscriberFn) error {
	ie := event.(internalNetworkErrorEvent)
	listener := subscriberFn.(graphsync.OnNetworkErrorListener)
	listener(ie.p, ie.request, ie.err)
	return nil
}

// NewNetworkErrorListeners returns a new list of listeners for when requestors cancel
func NewNetworkErrorListeners() *NetworkErrorListeners {
	return &NetworkErrorListeners{pubSub: pubsub.New(networkErrorDispatcher)}
}

// Register registers an listener for completed responses
func (nel *NetworkErrorListeners) Register(listener graphsync.OnNetworkErrorListener) graphsync.UnregisterHookFunc {
	return graphsync.UnregisterHookFunc(nel.pubSub.Subscribe(listener))
}

// NotifyNetworkErrorListeners notifies all listeners that a requestor cancelled a response
func (nel *NetworkErrorListeners) NotifyNetworkErrorListeners(p peer.ID, request graphsync.RequestData, err error) {
	_ = nel.pubSub.Publish(internalNetworkErrorEvent{p, request, err})
}
