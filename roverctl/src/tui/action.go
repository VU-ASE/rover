package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

// This is an action that can be performed by a page
type Action[T interface{}] struct {
	Name     string
	Result   bool
	Error    error
	Started  bool
	Finished bool
	Attempt  uint // allows you to ignore results from previous attempts
	Data     *T
}

type ActionInit[T interface{}] struct {
	Name string
}

// This is a message returned after an action is performed
// it describes the action and the data that was returned
type ActionResult[T interface{}] struct {
	Name       string
	Result     bool
	Error      error
	ForAttempt uint // allows you to ignore results from previous attempts
	Data       *T
}

// A collection of Actions that can be used by a model
type Actions []*Action[any]

// Check if a given result is for a specific action and attempt
func (a ActionResult[T]) IsFor(action *Action[T]) bool {
	return a.Name == action.Name && a.ForAttempt == action.Attempt
}

// ProcessResult takes an ActionResult and updates the Actions where the name matches
func (a Actions) ProcessResults(res ActionResult[any]) {
	for _, action := range a {
		if res.IsFor(action) {
			action.Result = res.Result
			action.Error = res.Error
			action.Finished = true
			action.Data = res.Data
		}
	}
}

// Check if an init matches an action and if so, start the action
func (action *Action[T]) ProcessInit(a ActionInit[T]) {
	if action.Name == a.Name {
		action.Reset()
		action.Start()
	}
}

// Checks if action matches result, and then updates the action with the result
func (action *Action[t]) ProcessResult(res ActionResult[t]) {
	if res.IsFor(action) {
		action.Result = res.Result
		action.Error = res.Error
		action.Finished = true
		action.Data = res.Data
	}
}

func (a Action[T]) IsLoading() bool {
	return a.Started && !a.Finished
}

func (a Action[T]) IsSuccess() bool {
	return a.Started && a.Finished && a.Result
}

// Can be used for optimistic updates, where you want to use the previous data while the new data is loading
func (a Action[T]) HasData() bool {
	return a.Data != nil
}

func (a Action[T]) IsError() bool {
	return a.Started && a.Finished && !a.Result
}

func (a Action[T]) IsDone() bool {
	return a.Started && a.Finished
}

func (a *Action[T]) Reset() {
	a.Started = false
	a.Finished = false
	a.Result = false
	a.Error = nil
	// a.Data = nil
}

func (a *Action[T]) Restart() {
	a.Reset()
	a.Start()
}

func (a *Action[T]) Start() {
	a.Attempt++
	a.Started = true
}

func (a *Action[T]) StartTea() tea.Cmd {
	return func() tea.Msg {
		a.Start()
		return nil
	}
}

func (a *Action[T]) ResetTea() tea.Cmd {
	return func() tea.Msg {
		a.Reset()
		return nil
	}
}

// Generate a new ActionResult from an Action
func NewResult[T interface{}](a Action[T], success bool, err error, data *T, attempt uint) ActionResult[T] {
	return ActionResult[T]{
		Name:       a.Name,
		Result:     success,
		Error:      err,
		ForAttempt: attempt,
		Data:       data,
	}
}

func NewAction[T interface{}](name string) Action[T] {
	return Action[T]{
		Name:     name,
		Result:   false,
		Error:    nil,
		Started:  false,
		Finished: false,
		Attempt:  0,
		Data:     nil,
	}
}

type ActionFunction[T interface{}] func() (*T, error)

// Wrapper that makes your life easier when performing an action
// You can also use the oldschool method of creating a function that returns a tea.Cmd, and use tui.NewResult() with that
func PerformAction[T interface{}](action *Action[T], f ActionFunction[T]) tea.Cmd {
	attempt := action.Attempt + 1
	init := func() tea.Cmd {
		return func() tea.Msg {
			return ActionInit[T]{Name: action.Name}
		}
	}
	run := func() tea.Cmd {
		return func() tea.Msg {
			data, err := f()
			return NewResult(*action, err == nil, err, data, attempt)
		}
	}
	return tea.Sequence(init(), run())
}

//
// ActionV2 definitions (which should make using actions a bit more ergonomic)
//

type ActionV2[Req interface{}, Res interface{}] struct {
	id      string // unique identifier to match updates with actions
	request ActionRequestV2[Req]
	result  ActionResultV2[Res]
}

type ActionRequestV2[Req interface{}] struct {
	attempt   uint  // 0 indicates not started, >0 indicates the attempt identifier
	createdAt int64 // time in milliseconds since epoch
	data      Req
}

type ActionResultV2[Res interface{}] struct {
	data      Res
	attempt   uint    // 0 indicates not started, >0 indicates the attempt identifier
	createdAt int64   // time in milliseconds since epoch
	errors    []error // any errors that might have occurred
}

func NewNamedActionV2[Req interface{}, Res interface{}](id string) ActionV2[Req, Res] {
	return ActionV2[Req, Res]{
		id: id,
		request: ActionRequestV2[Req]{
			attempt:   0,
			createdAt: time.Now().UnixMilli(),
		},
		result: ActionResultV2[Res]{
			attempt:   0,
			createdAt: time.Now().UnixMilli(),
		},
	}
}

func NewActionV2[Req interface{}, Res interface{}]() ActionV2[Req, Res] {
	// Auto-generate ID
	id := "action-" + uuid.New().String()
	return NewNamedActionV2[Req, Res](id)
}

type ActionFunctionV2[Req interface{}, Res interface{}] func() (*Res, []error)

// Public API as used by the pages
func PerformActionV2[Req interface{}, Res interface{}](action *ActionV2[Req, Res], req *Req, f ActionFunctionV2[Req, Res]) tea.Cmd {
	attempt := action.request.attempt + 1
	init := func() tea.Cmd {
		return func() tea.Msg {
			return ActionV2Init[Req, Res]{id: action.id, req: req}
		}
	}
	run := func() tea.Cmd {
		return func() tea.Msg {
			data, err := f()
			return ActionV2Result[Req, Res]{id: action.id, forAttempt: attempt, data: data, errors: err}
		}
	}
	return tea.Sequence(init(), run())
}

func (a *ActionV2[Req, Res]) start(req *Req) {
	a.request.attempt++
	a.request.createdAt = time.Now().UnixMilli()

	if req != nil {
		a.request.data = *req
	}
}

func (a *ActionV2[Req, Res]) finish(attempt uint, data *Res, errors []error) {
	if (attempt == 0) || (attempt != a.request.attempt) {
		return
	}
	if data != nil {
		a.result.data = *data
	}
	a.result.attempt = attempt
	a.result.createdAt = time.Now().UnixMilli()
	a.result.errors = errors
}

// Interface to constrain the type of action updates
type ActionUpdate[Req interface{}, Res interface{}] interface {
	isForAction() string // returns the ID of the action that this update is for
}
type ActionV2Init[Req interface{}, Res interface{}] struct {
	id  string
	req *Req
}

func (a ActionV2Init[Req, Res]) isForAction() string {
	return a.id
}

type ActionV2Result[Req interface{}, Res interface{}] struct {
	id         string
	forAttempt uint // allows you to ignore results from previous attempts
	data       *Res
	errors     []error
}

func (a ActionV2Result[Req, Res]) isForAction() string {
	return a.id
}

func (action *ActionV2[Req, Res]) ProcessUpdate(update ActionUpdate[Req, Res]) {
	if action.id == update.isForAction() {
		switch u := update.(type) {
		case ActionV2Init[Req, Res]:
			action.start(u.req)
		case ActionV2Result[Req, Res]:
			action.finish(u.forAttempt, u.data, u.errors)
		}
	}
}

func (update ActionV2Init[Req, Res]) IsFor(action ActionUpdate[Req, Res]) bool {
	return update.isForAction() == action.isForAction()
}

func (action ActionV2[Req, Res]) IsLoading() bool {
	return action.request.attempt > 0 && action.result.attempt != action.request.attempt
}

func (action ActionV2[Req, Res]) IsDone() bool {
	return action.request.attempt > 0 && action.result.attempt == action.request.attempt
}

func (action ActionV2[Req, Res]) IsSuccess() bool {
	return action.IsDone() && len(action.result.errors) == 0
}

func (action ActionV2[Req, Res]) IsError() bool {
	return action.IsDone() && len(action.result.errors) > 0
}

func (action ActionV2[Req, Res]) HasRequest() bool {
	return action.request.attempt > 0
}

func (action ActionV2[Req, Res]) Request() Req {
	return action.request.data
}

func (action ActionV2[Req, Res]) Started() bool {
	return action.request.attempt > 0
}

func (action ActionV2[Req, Res]) HasResult() bool {
	return action.result.attempt > 0 && action.result.errors == nil
}

func (action ActionV2[Req, Res]) Result() Res {
	return action.result.data
}

func (action ActionV2[Req, Res]) Errors() []error {
	return action.result.errors
}

func AllSuccess(actions ...interface{ IsSuccess() bool }) bool {
	for _, action := range actions {
		if !action.IsSuccess() {
			return false
		}
	}
	return true
}

type NoRequest struct{}
