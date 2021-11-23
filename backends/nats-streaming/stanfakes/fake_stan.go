// Code generated by counterfeiter. DO NOT EDIT.
package stanfakes

import (
	"sync"

	nats "github.com/nats-io/nats.go"
	stan "github.com/nats-io/stan.go"
)

type FakeConn struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	NatsConnStub        func() *nats.Conn
	natsConnMutex       sync.RWMutex
	natsConnArgsForCall []struct {
	}
	natsConnReturns struct {
		result1 *nats.Conn
	}
	natsConnReturnsOnCall map[int]struct {
		result1 *nats.Conn
	}
	PublishStub        func(string, []byte) error
	publishMutex       sync.RWMutex
	publishArgsForCall []struct {
		arg1 string
		arg2 []byte
	}
	publishReturns struct {
		result1 error
	}
	publishReturnsOnCall map[int]struct {
		result1 error
	}
	PublishAsyncStub        func(string, []byte, stan.AckHandler) (string, error)
	publishAsyncMutex       sync.RWMutex
	publishAsyncArgsForCall []struct {
		arg1 string
		arg2 []byte
		arg3 stan.AckHandler
	}
	publishAsyncReturns struct {
		result1 string
		result2 error
	}
	publishAsyncReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	QueueSubscribeStub        func(string, string, stan.MsgHandler, ...stan.SubscriptionOption) (stan.Subscription, error)
	queueSubscribeMutex       sync.RWMutex
	queueSubscribeArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 stan.MsgHandler
		arg4 []stan.SubscriptionOption
	}
	queueSubscribeReturns struct {
		result1 stan.Subscription
		result2 error
	}
	queueSubscribeReturnsOnCall map[int]struct {
		result1 stan.Subscription
		result2 error
	}
	SubscribeStub        func(string, stan.MsgHandler, ...stan.SubscriptionOption) (stan.Subscription, error)
	subscribeMutex       sync.RWMutex
	subscribeArgsForCall []struct {
		arg1 string
		arg2 stan.MsgHandler
		arg3 []stan.SubscriptionOption
	}
	subscribeReturns struct {
		result1 stan.Subscription
		result2 error
	}
	subscribeReturnsOnCall map[int]struct {
		result1 stan.Subscription
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConn) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConn) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeConn) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeConn) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConn) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConn) NatsConn() *nats.Conn {
	fake.natsConnMutex.Lock()
	ret, specificReturn := fake.natsConnReturnsOnCall[len(fake.natsConnArgsForCall)]
	fake.natsConnArgsForCall = append(fake.natsConnArgsForCall, struct {
	}{})
	stub := fake.NatsConnStub
	fakeReturns := fake.natsConnReturns
	fake.recordInvocation("NatsConn", []interface{}{})
	fake.natsConnMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConn) NatsConnCallCount() int {
	fake.natsConnMutex.RLock()
	defer fake.natsConnMutex.RUnlock()
	return len(fake.natsConnArgsForCall)
}

func (fake *FakeConn) NatsConnCalls(stub func() *nats.Conn) {
	fake.natsConnMutex.Lock()
	defer fake.natsConnMutex.Unlock()
	fake.NatsConnStub = stub
}

func (fake *FakeConn) NatsConnReturns(result1 *nats.Conn) {
	fake.natsConnMutex.Lock()
	defer fake.natsConnMutex.Unlock()
	fake.NatsConnStub = nil
	fake.natsConnReturns = struct {
		result1 *nats.Conn
	}{result1}
}

func (fake *FakeConn) NatsConnReturnsOnCall(i int, result1 *nats.Conn) {
	fake.natsConnMutex.Lock()
	defer fake.natsConnMutex.Unlock()
	fake.NatsConnStub = nil
	if fake.natsConnReturnsOnCall == nil {
		fake.natsConnReturnsOnCall = make(map[int]struct {
			result1 *nats.Conn
		})
	}
	fake.natsConnReturnsOnCall[i] = struct {
		result1 *nats.Conn
	}{result1}
}

func (fake *FakeConn) Publish(arg1 string, arg2 []byte) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.publishMutex.Lock()
	ret, specificReturn := fake.publishReturnsOnCall[len(fake.publishArgsForCall)]
	fake.publishArgsForCall = append(fake.publishArgsForCall, struct {
		arg1 string
		arg2 []byte
	}{arg1, arg2Copy})
	stub := fake.PublishStub
	fakeReturns := fake.publishReturns
	fake.recordInvocation("Publish", []interface{}{arg1, arg2Copy})
	fake.publishMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConn) PublishCallCount() int {
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	return len(fake.publishArgsForCall)
}

func (fake *FakeConn) PublishCalls(stub func(string, []byte) error) {
	fake.publishMutex.Lock()
	defer fake.publishMutex.Unlock()
	fake.PublishStub = stub
}

func (fake *FakeConn) PublishArgsForCall(i int) (string, []byte) {
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	argsForCall := fake.publishArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeConn) PublishReturns(result1 error) {
	fake.publishMutex.Lock()
	defer fake.publishMutex.Unlock()
	fake.PublishStub = nil
	fake.publishReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConn) PublishReturnsOnCall(i int, result1 error) {
	fake.publishMutex.Lock()
	defer fake.publishMutex.Unlock()
	fake.PublishStub = nil
	if fake.publishReturnsOnCall == nil {
		fake.publishReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConn) PublishAsync(arg1 string, arg2 []byte, arg3 stan.AckHandler) (string, error) {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.publishAsyncMutex.Lock()
	ret, specificReturn := fake.publishAsyncReturnsOnCall[len(fake.publishAsyncArgsForCall)]
	fake.publishAsyncArgsForCall = append(fake.publishAsyncArgsForCall, struct {
		arg1 string
		arg2 []byte
		arg3 stan.AckHandler
	}{arg1, arg2Copy, arg3})
	stub := fake.PublishAsyncStub
	fakeReturns := fake.publishAsyncReturns
	fake.recordInvocation("PublishAsync", []interface{}{arg1, arg2Copy, arg3})
	fake.publishAsyncMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConn) PublishAsyncCallCount() int {
	fake.publishAsyncMutex.RLock()
	defer fake.publishAsyncMutex.RUnlock()
	return len(fake.publishAsyncArgsForCall)
}

func (fake *FakeConn) PublishAsyncCalls(stub func(string, []byte, stan.AckHandler) (string, error)) {
	fake.publishAsyncMutex.Lock()
	defer fake.publishAsyncMutex.Unlock()
	fake.PublishAsyncStub = stub
}

func (fake *FakeConn) PublishAsyncArgsForCall(i int) (string, []byte, stan.AckHandler) {
	fake.publishAsyncMutex.RLock()
	defer fake.publishAsyncMutex.RUnlock()
	argsForCall := fake.publishAsyncArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeConn) PublishAsyncReturns(result1 string, result2 error) {
	fake.publishAsyncMutex.Lock()
	defer fake.publishAsyncMutex.Unlock()
	fake.PublishAsyncStub = nil
	fake.publishAsyncReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) PublishAsyncReturnsOnCall(i int, result1 string, result2 error) {
	fake.publishAsyncMutex.Lock()
	defer fake.publishAsyncMutex.Unlock()
	fake.PublishAsyncStub = nil
	if fake.publishAsyncReturnsOnCall == nil {
		fake.publishAsyncReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.publishAsyncReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) QueueSubscribe(arg1 string, arg2 string, arg3 stan.MsgHandler, arg4 ...stan.SubscriptionOption) (stan.Subscription, error) {
	fake.queueSubscribeMutex.Lock()
	ret, specificReturn := fake.queueSubscribeReturnsOnCall[len(fake.queueSubscribeArgsForCall)]
	fake.queueSubscribeArgsForCall = append(fake.queueSubscribeArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 stan.MsgHandler
		arg4 []stan.SubscriptionOption
	}{arg1, arg2, arg3, arg4})
	stub := fake.QueueSubscribeStub
	fakeReturns := fake.queueSubscribeReturns
	fake.recordInvocation("QueueSubscribe", []interface{}{arg1, arg2, arg3, arg4})
	fake.queueSubscribeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConn) QueueSubscribeCallCount() int {
	fake.queueSubscribeMutex.RLock()
	defer fake.queueSubscribeMutex.RUnlock()
	return len(fake.queueSubscribeArgsForCall)
}

func (fake *FakeConn) QueueSubscribeCalls(stub func(string, string, stan.MsgHandler, ...stan.SubscriptionOption) (stan.Subscription, error)) {
	fake.queueSubscribeMutex.Lock()
	defer fake.queueSubscribeMutex.Unlock()
	fake.QueueSubscribeStub = stub
}

func (fake *FakeConn) QueueSubscribeArgsForCall(i int) (string, string, stan.MsgHandler, []stan.SubscriptionOption) {
	fake.queueSubscribeMutex.RLock()
	defer fake.queueSubscribeMutex.RUnlock()
	argsForCall := fake.queueSubscribeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeConn) QueueSubscribeReturns(result1 stan.Subscription, result2 error) {
	fake.queueSubscribeMutex.Lock()
	defer fake.queueSubscribeMutex.Unlock()
	fake.QueueSubscribeStub = nil
	fake.queueSubscribeReturns = struct {
		result1 stan.Subscription
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) QueueSubscribeReturnsOnCall(i int, result1 stan.Subscription, result2 error) {
	fake.queueSubscribeMutex.Lock()
	defer fake.queueSubscribeMutex.Unlock()
	fake.QueueSubscribeStub = nil
	if fake.queueSubscribeReturnsOnCall == nil {
		fake.queueSubscribeReturnsOnCall = make(map[int]struct {
			result1 stan.Subscription
			result2 error
		})
	}
	fake.queueSubscribeReturnsOnCall[i] = struct {
		result1 stan.Subscription
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) Subscribe(arg1 string, arg2 stan.MsgHandler, arg3 ...stan.SubscriptionOption) (stan.Subscription, error) {
	fake.subscribeMutex.Lock()
	ret, specificReturn := fake.subscribeReturnsOnCall[len(fake.subscribeArgsForCall)]
	fake.subscribeArgsForCall = append(fake.subscribeArgsForCall, struct {
		arg1 string
		arg2 stan.MsgHandler
		arg3 []stan.SubscriptionOption
	}{arg1, arg2, arg3})
	stub := fake.SubscribeStub
	fakeReturns := fake.subscribeReturns
	fake.recordInvocation("Subscribe", []interface{}{arg1, arg2, arg3})
	fake.subscribeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConn) SubscribeCallCount() int {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return len(fake.subscribeArgsForCall)
}

func (fake *FakeConn) SubscribeCalls(stub func(string, stan.MsgHandler, ...stan.SubscriptionOption) (stan.Subscription, error)) {
	fake.subscribeMutex.Lock()
	defer fake.subscribeMutex.Unlock()
	fake.SubscribeStub = stub
}

func (fake *FakeConn) SubscribeArgsForCall(i int) (string, stan.MsgHandler, []stan.SubscriptionOption) {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	argsForCall := fake.subscribeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeConn) SubscribeReturns(result1 stan.Subscription, result2 error) {
	fake.subscribeMutex.Lock()
	defer fake.subscribeMutex.Unlock()
	fake.SubscribeStub = nil
	fake.subscribeReturns = struct {
		result1 stan.Subscription
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) SubscribeReturnsOnCall(i int, result1 stan.Subscription, result2 error) {
	fake.subscribeMutex.Lock()
	defer fake.subscribeMutex.Unlock()
	fake.SubscribeStub = nil
	if fake.subscribeReturnsOnCall == nil {
		fake.subscribeReturnsOnCall = make(map[int]struct {
			result1 stan.Subscription
			result2 error
		})
	}
	fake.subscribeReturnsOnCall[i] = struct {
		result1 stan.Subscription
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.natsConnMutex.RLock()
	defer fake.natsConnMutex.RUnlock()
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	fake.publishAsyncMutex.RLock()
	defer fake.publishAsyncMutex.RUnlock()
	fake.queueSubscribeMutex.RLock()
	defer fake.queueSubscribeMutex.RUnlock()
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConn) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ stan.Conn = new(FakeConn)
