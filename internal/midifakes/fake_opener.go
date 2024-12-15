// Code generated by counterfeiter. DO NOT EDIT.
package midifakes

import (
	"sync"

	"github.com/cholick/midi-cli/pkg/midi"
)

type FakeOpener struct {
	NewDefaultOutStub        func() (midi.Out, error)
	newDefaultOutMutex       sync.RWMutex
	newDefaultOutArgsForCall []struct {
	}
	newDefaultOutReturns struct {
		result1 midi.Out
		result2 error
	}
	newDefaultOutReturnsOnCall map[int]struct {
		result1 midi.Out
		result2 error
	}
	NewOutForPortStub        func(string) (midi.Out, error)
	newOutForPortMutex       sync.RWMutex
	newOutForPortArgsForCall []struct {
		arg1 string
	}
	newOutForPortReturns struct {
		result1 midi.Out
		result2 error
	}
	newOutForPortReturnsOnCall map[int]struct {
		result1 midi.Out
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOpener) NewDefaultOut() (midi.Out, error) {
	fake.newDefaultOutMutex.Lock()
	ret, specificReturn := fake.newDefaultOutReturnsOnCall[len(fake.newDefaultOutArgsForCall)]
	fake.newDefaultOutArgsForCall = append(fake.newDefaultOutArgsForCall, struct {
	}{})
	stub := fake.NewDefaultOutStub
	fakeReturns := fake.newDefaultOutReturns
	fake.recordInvocation("NewDefaultOut", []interface{}{})
	fake.newDefaultOutMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOpener) NewDefaultOutCallCount() int {
	fake.newDefaultOutMutex.RLock()
	defer fake.newDefaultOutMutex.RUnlock()
	return len(fake.newDefaultOutArgsForCall)
}

func (fake *FakeOpener) NewDefaultOutCalls(stub func() (midi.Out, error)) {
	fake.newDefaultOutMutex.Lock()
	defer fake.newDefaultOutMutex.Unlock()
	fake.NewDefaultOutStub = stub
}

func (fake *FakeOpener) NewDefaultOutReturns(result1 midi.Out, result2 error) {
	fake.newDefaultOutMutex.Lock()
	defer fake.newDefaultOutMutex.Unlock()
	fake.NewDefaultOutStub = nil
	fake.newDefaultOutReturns = struct {
		result1 midi.Out
		result2 error
	}{result1, result2}
}

func (fake *FakeOpener) NewDefaultOutReturnsOnCall(i int, result1 midi.Out, result2 error) {
	fake.newDefaultOutMutex.Lock()
	defer fake.newDefaultOutMutex.Unlock()
	fake.NewDefaultOutStub = nil
	if fake.newDefaultOutReturnsOnCall == nil {
		fake.newDefaultOutReturnsOnCall = make(map[int]struct {
			result1 midi.Out
			result2 error
		})
	}
	fake.newDefaultOutReturnsOnCall[i] = struct {
		result1 midi.Out
		result2 error
	}{result1, result2}
}

func (fake *FakeOpener) NewOutForPort(arg1 string) (midi.Out, error) {
	fake.newOutForPortMutex.Lock()
	ret, specificReturn := fake.newOutForPortReturnsOnCall[len(fake.newOutForPortArgsForCall)]
	fake.newOutForPortArgsForCall = append(fake.newOutForPortArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.NewOutForPortStub
	fakeReturns := fake.newOutForPortReturns
	fake.recordInvocation("NewOutForPort", []interface{}{arg1})
	fake.newOutForPortMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOpener) NewOutForPortCallCount() int {
	fake.newOutForPortMutex.RLock()
	defer fake.newOutForPortMutex.RUnlock()
	return len(fake.newOutForPortArgsForCall)
}

func (fake *FakeOpener) NewOutForPortCalls(stub func(string) (midi.Out, error)) {
	fake.newOutForPortMutex.Lock()
	defer fake.newOutForPortMutex.Unlock()
	fake.NewOutForPortStub = stub
}

func (fake *FakeOpener) NewOutForPortArgsForCall(i int) string {
	fake.newOutForPortMutex.RLock()
	defer fake.newOutForPortMutex.RUnlock()
	argsForCall := fake.newOutForPortArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOpener) NewOutForPortReturns(result1 midi.Out, result2 error) {
	fake.newOutForPortMutex.Lock()
	defer fake.newOutForPortMutex.Unlock()
	fake.NewOutForPortStub = nil
	fake.newOutForPortReturns = struct {
		result1 midi.Out
		result2 error
	}{result1, result2}
}

func (fake *FakeOpener) NewOutForPortReturnsOnCall(i int, result1 midi.Out, result2 error) {
	fake.newOutForPortMutex.Lock()
	defer fake.newOutForPortMutex.Unlock()
	fake.NewOutForPortStub = nil
	if fake.newOutForPortReturnsOnCall == nil {
		fake.newOutForPortReturnsOnCall = make(map[int]struct {
			result1 midi.Out
			result2 error
		})
	}
	fake.newOutForPortReturnsOnCall[i] = struct {
		result1 midi.Out
		result2 error
	}{result1, result2}
}

func (fake *FakeOpener) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newDefaultOutMutex.RLock()
	defer fake.newDefaultOutMutex.RUnlock()
	fake.newOutForPortMutex.RLock()
	defer fake.newOutForPortMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOpener) recordInvocation(key string, args []interface{}) {
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

var _ midi.Opener = new(FakeOpener)