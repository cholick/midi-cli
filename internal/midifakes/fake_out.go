// Code generated by counterfeiter. DO NOT EDIT.
package midifakes

import (
	"sync"

	"github.com/cholick/midi-cli/pkg/midi"
)

type FakeOut struct {
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	ControlChangeStub        func(int, int, int) error
	controlChangeMutex       sync.RWMutex
	controlChangeArgsForCall []struct {
		arg1 int
		arg2 int
		arg3 int
	}
	controlChangeReturns struct {
		result1 error
	}
	controlChangeReturnsOnCall map[int]struct {
		result1 error
	}
	ListPortsStub        func() ([]string, error)
	listPortsMutex       sync.RWMutex
	listPortsArgsForCall []struct {
	}
	listPortsReturns struct {
		result1 []string
		result2 error
	}
	listPortsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	NoteOffStub        func(string, int, int) error
	noteOffMutex       sync.RWMutex
	noteOffArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 int
	}
	noteOffReturns struct {
		result1 error
	}
	noteOffReturnsOnCall map[int]struct {
		result1 error
	}
	NoteOnStub        func(string, int, int) error
	noteOnMutex       sync.RWMutex
	noteOnArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 int
	}
	noteOnReturns struct {
		result1 error
	}
	noteOnReturnsOnCall map[int]struct {
		result1 error
	}
	OpenPortStub        func(string) error
	openPortMutex       sync.RWMutex
	openPortArgsForCall []struct {
		arg1 string
	}
	openPortReturns struct {
		result1 error
	}
	openPortReturnsOnCall map[int]struct {
		result1 error
	}
	PanicStub        func(int) error
	panicMutex       sync.RWMutex
	panicArgsForCall []struct {
		arg1 int
	}
	panicReturns struct {
		result1 error
	}
	panicReturnsOnCall map[int]struct {
		result1 error
	}
	PanicAllStub        func() error
	panicAllMutex       sync.RWMutex
	panicAllArgsForCall []struct {
	}
	panicAllReturns struct {
		result1 error
	}
	panicAllReturnsOnCall map[int]struct {
		result1 error
	}
	ProgramChangeStub        func(int, int) error
	programChangeMutex       sync.RWMutex
	programChangeArgsForCall []struct {
		arg1 int
		arg2 int
	}
	programChangeReturns struct {
		result1 error
	}
	programChangeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOut) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeOut) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeOut) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeOut) ControlChange(arg1 int, arg2 int, arg3 int) error {
	fake.controlChangeMutex.Lock()
	ret, specificReturn := fake.controlChangeReturnsOnCall[len(fake.controlChangeArgsForCall)]
	fake.controlChangeArgsForCall = append(fake.controlChangeArgsForCall, struct {
		arg1 int
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	stub := fake.ControlChangeStub
	fakeReturns := fake.controlChangeReturns
	fake.recordInvocation("ControlChange", []interface{}{arg1, arg2, arg3})
	fake.controlChangeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOut) ControlChangeCallCount() int {
	fake.controlChangeMutex.RLock()
	defer fake.controlChangeMutex.RUnlock()
	return len(fake.controlChangeArgsForCall)
}

func (fake *FakeOut) ControlChangeCalls(stub func(int, int, int) error) {
	fake.controlChangeMutex.Lock()
	defer fake.controlChangeMutex.Unlock()
	fake.ControlChangeStub = stub
}

func (fake *FakeOut) ControlChangeArgsForCall(i int) (int, int, int) {
	fake.controlChangeMutex.RLock()
	defer fake.controlChangeMutex.RUnlock()
	argsForCall := fake.controlChangeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeOut) ControlChangeReturns(result1 error) {
	fake.controlChangeMutex.Lock()
	defer fake.controlChangeMutex.Unlock()
	fake.ControlChangeStub = nil
	fake.controlChangeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOut) ControlChangeReturnsOnCall(i int, result1 error) {
	fake.controlChangeMutex.Lock()
	defer fake.controlChangeMutex.Unlock()
	fake.ControlChangeStub = nil
	if fake.controlChangeReturnsOnCall == nil {
		fake.controlChangeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.controlChangeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOut) ListPorts() ([]string, error) {
	fake.listPortsMutex.Lock()
	ret, specificReturn := fake.listPortsReturnsOnCall[len(fake.listPortsArgsForCall)]
	fake.listPortsArgsForCall = append(fake.listPortsArgsForCall, struct {
	}{})
	stub := fake.ListPortsStub
	fakeReturns := fake.listPortsReturns
	fake.recordInvocation("ListPorts", []interface{}{})
	fake.listPortsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOut) ListPortsCallCount() int {
	fake.listPortsMutex.RLock()
	defer fake.listPortsMutex.RUnlock()
	return len(fake.listPortsArgsForCall)
}

func (fake *FakeOut) ListPortsCalls(stub func() ([]string, error)) {
	fake.listPortsMutex.Lock()
	defer fake.listPortsMutex.Unlock()
	fake.ListPortsStub = stub
}

func (fake *FakeOut) ListPortsReturns(result1 []string, result2 error) {
	fake.listPortsMutex.Lock()
	defer fake.listPortsMutex.Unlock()
	fake.ListPortsStub = nil
	fake.listPortsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeOut) ListPortsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listPortsMutex.Lock()
	defer fake.listPortsMutex.Unlock()
	fake.ListPortsStub = nil
	if fake.listPortsReturnsOnCall == nil {
		fake.listPortsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listPortsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeOut) NoteOff(arg1 string, arg2 int, arg3 int) error {
	fake.noteOffMutex.Lock()
	ret, specificReturn := fake.noteOffReturnsOnCall[len(fake.noteOffArgsForCall)]
	fake.noteOffArgsForCall = append(fake.noteOffArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	stub := fake.NoteOffStub
	fakeReturns := fake.noteOffReturns
	fake.recordInvocation("NoteOff", []interface{}{arg1, arg2, arg3})
	fake.noteOffMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOut) NoteOffCallCount() int {
	fake.noteOffMutex.RLock()
	defer fake.noteOffMutex.RUnlock()
	return len(fake.noteOffArgsForCall)
}

func (fake *FakeOut) NoteOffCalls(stub func(string, int, int) error) {
	fake.noteOffMutex.Lock()
	defer fake.noteOffMutex.Unlock()
	fake.NoteOffStub = stub
}

func (fake *FakeOut) NoteOffArgsForCall(i int) (string, int, int) {
	fake.noteOffMutex.RLock()
	defer fake.noteOffMutex.RUnlock()
	argsForCall := fake.noteOffArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeOut) NoteOffReturns(result1 error) {
	fake.noteOffMutex.Lock()
	defer fake.noteOffMutex.Unlock()
	fake.NoteOffStub = nil
	fake.noteOffReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOut) NoteOffReturnsOnCall(i int, result1 error) {
	fake.noteOffMutex.Lock()
	defer fake.noteOffMutex.Unlock()
	fake.NoteOffStub = nil
	if fake.noteOffReturnsOnCall == nil {
		fake.noteOffReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.noteOffReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOut) NoteOn(arg1 string, arg2 int, arg3 int) error {
	fake.noteOnMutex.Lock()
	ret, specificReturn := fake.noteOnReturnsOnCall[len(fake.noteOnArgsForCall)]
	fake.noteOnArgsForCall = append(fake.noteOnArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	stub := fake.NoteOnStub
	fakeReturns := fake.noteOnReturns
	fake.recordInvocation("NoteOn", []interface{}{arg1, arg2, arg3})
	fake.noteOnMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOut) NoteOnCallCount() int {
	fake.noteOnMutex.RLock()
	defer fake.noteOnMutex.RUnlock()
	return len(fake.noteOnArgsForCall)
}

func (fake *FakeOut) NoteOnCalls(stub func(string, int, int) error) {
	fake.noteOnMutex.Lock()
	defer fake.noteOnMutex.Unlock()
	fake.NoteOnStub = stub
}

func (fake *FakeOut) NoteOnArgsForCall(i int) (string, int, int) {
	fake.noteOnMutex.RLock()
	defer fake.noteOnMutex.RUnlock()
	argsForCall := fake.noteOnArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeOut) NoteOnReturns(result1 error) {
	fake.noteOnMutex.Lock()
	defer fake.noteOnMutex.Unlock()
	fake.NoteOnStub = nil
	fake.noteOnReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOut) NoteOnReturnsOnCall(i int, result1 error) {
	fake.noteOnMutex.Lock()
	defer fake.noteOnMutex.Unlock()
	fake.NoteOnStub = nil
	if fake.noteOnReturnsOnCall == nil {
		fake.noteOnReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.noteOnReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOut) OpenPort(arg1 string) error {
	fake.openPortMutex.Lock()
	ret, specificReturn := fake.openPortReturnsOnCall[len(fake.openPortArgsForCall)]
	fake.openPortArgsForCall = append(fake.openPortArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.OpenPortStub
	fakeReturns := fake.openPortReturns
	fake.recordInvocation("OpenPort", []interface{}{arg1})
	fake.openPortMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOut) OpenPortCallCount() int {
	fake.openPortMutex.RLock()
	defer fake.openPortMutex.RUnlock()
	return len(fake.openPortArgsForCall)
}

func (fake *FakeOut) OpenPortCalls(stub func(string) error) {
	fake.openPortMutex.Lock()
	defer fake.openPortMutex.Unlock()
	fake.OpenPortStub = stub
}

func (fake *FakeOut) OpenPortArgsForCall(i int) string {
	fake.openPortMutex.RLock()
	defer fake.openPortMutex.RUnlock()
	argsForCall := fake.openPortArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOut) OpenPortReturns(result1 error) {
	fake.openPortMutex.Lock()
	defer fake.openPortMutex.Unlock()
	fake.OpenPortStub = nil
	fake.openPortReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOut) OpenPortReturnsOnCall(i int, result1 error) {
	fake.openPortMutex.Lock()
	defer fake.openPortMutex.Unlock()
	fake.OpenPortStub = nil
	if fake.openPortReturnsOnCall == nil {
		fake.openPortReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.openPortReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOut) Panic(arg1 int) error {
	fake.panicMutex.Lock()
	ret, specificReturn := fake.panicReturnsOnCall[len(fake.panicArgsForCall)]
	fake.panicArgsForCall = append(fake.panicArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.PanicStub
	fakeReturns := fake.panicReturns
	fake.recordInvocation("Panic", []interface{}{arg1})
	fake.panicMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOut) PanicCallCount() int {
	fake.panicMutex.RLock()
	defer fake.panicMutex.RUnlock()
	return len(fake.panicArgsForCall)
}

func (fake *FakeOut) PanicCalls(stub func(int) error) {
	fake.panicMutex.Lock()
	defer fake.panicMutex.Unlock()
	fake.PanicStub = stub
}

func (fake *FakeOut) PanicArgsForCall(i int) int {
	fake.panicMutex.RLock()
	defer fake.panicMutex.RUnlock()
	argsForCall := fake.panicArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOut) PanicReturns(result1 error) {
	fake.panicMutex.Lock()
	defer fake.panicMutex.Unlock()
	fake.PanicStub = nil
	fake.panicReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOut) PanicReturnsOnCall(i int, result1 error) {
	fake.panicMutex.Lock()
	defer fake.panicMutex.Unlock()
	fake.PanicStub = nil
	if fake.panicReturnsOnCall == nil {
		fake.panicReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.panicReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOut) PanicAll() error {
	fake.panicAllMutex.Lock()
	ret, specificReturn := fake.panicAllReturnsOnCall[len(fake.panicAllArgsForCall)]
	fake.panicAllArgsForCall = append(fake.panicAllArgsForCall, struct {
	}{})
	stub := fake.PanicAllStub
	fakeReturns := fake.panicAllReturns
	fake.recordInvocation("PanicAll", []interface{}{})
	fake.panicAllMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOut) PanicAllCallCount() int {
	fake.panicAllMutex.RLock()
	defer fake.panicAllMutex.RUnlock()
	return len(fake.panicAllArgsForCall)
}

func (fake *FakeOut) PanicAllCalls(stub func() error) {
	fake.panicAllMutex.Lock()
	defer fake.panicAllMutex.Unlock()
	fake.PanicAllStub = stub
}

func (fake *FakeOut) PanicAllReturns(result1 error) {
	fake.panicAllMutex.Lock()
	defer fake.panicAllMutex.Unlock()
	fake.PanicAllStub = nil
	fake.panicAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOut) PanicAllReturnsOnCall(i int, result1 error) {
	fake.panicAllMutex.Lock()
	defer fake.panicAllMutex.Unlock()
	fake.PanicAllStub = nil
	if fake.panicAllReturnsOnCall == nil {
		fake.panicAllReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.panicAllReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOut) ProgramChange(arg1 int, arg2 int) error {
	fake.programChangeMutex.Lock()
	ret, specificReturn := fake.programChangeReturnsOnCall[len(fake.programChangeArgsForCall)]
	fake.programChangeArgsForCall = append(fake.programChangeArgsForCall, struct {
		arg1 int
		arg2 int
	}{arg1, arg2})
	stub := fake.ProgramChangeStub
	fakeReturns := fake.programChangeReturns
	fake.recordInvocation("ProgramChange", []interface{}{arg1, arg2})
	fake.programChangeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOut) ProgramChangeCallCount() int {
	fake.programChangeMutex.RLock()
	defer fake.programChangeMutex.RUnlock()
	return len(fake.programChangeArgsForCall)
}

func (fake *FakeOut) ProgramChangeCalls(stub func(int, int) error) {
	fake.programChangeMutex.Lock()
	defer fake.programChangeMutex.Unlock()
	fake.ProgramChangeStub = stub
}

func (fake *FakeOut) ProgramChangeArgsForCall(i int) (int, int) {
	fake.programChangeMutex.RLock()
	defer fake.programChangeMutex.RUnlock()
	argsForCall := fake.programChangeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOut) ProgramChangeReturns(result1 error) {
	fake.programChangeMutex.Lock()
	defer fake.programChangeMutex.Unlock()
	fake.ProgramChangeStub = nil
	fake.programChangeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOut) ProgramChangeReturnsOnCall(i int, result1 error) {
	fake.programChangeMutex.Lock()
	defer fake.programChangeMutex.Unlock()
	fake.ProgramChangeStub = nil
	if fake.programChangeReturnsOnCall == nil {
		fake.programChangeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.programChangeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOut) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.controlChangeMutex.RLock()
	defer fake.controlChangeMutex.RUnlock()
	fake.listPortsMutex.RLock()
	defer fake.listPortsMutex.RUnlock()
	fake.noteOffMutex.RLock()
	defer fake.noteOffMutex.RUnlock()
	fake.noteOnMutex.RLock()
	defer fake.noteOnMutex.RUnlock()
	fake.openPortMutex.RLock()
	defer fake.openPortMutex.RUnlock()
	fake.panicMutex.RLock()
	defer fake.panicMutex.RUnlock()
	fake.panicAllMutex.RLock()
	defer fake.panicAllMutex.RUnlock()
	fake.programChangeMutex.RLock()
	defer fake.programChangeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOut) recordInvocation(key string, args []interface{}) {
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

var _ midi.Out = new(FakeOut)
