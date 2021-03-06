// Package memory provides the Chrome Debugging Protocol
// commands, types, and events for the Chrome Memory domain.
//
// Generated by the chromedp-gen command.
package memory

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
)

// GetDOMCountersParams [no description].
type GetDOMCountersParams struct{}

// GetDOMCounters [no description].
func GetDOMCounters() *GetDOMCountersParams {
	return &GetDOMCountersParams{}
}

// GetDOMCountersReturns return values.
type GetDOMCountersReturns struct {
	Documents        int64 `json:"documents,omitempty"`
	Nodes            int64 `json:"nodes,omitempty"`
	JsEventListeners int64 `json:"jsEventListeners,omitempty"`
}

// Do executes Memory.getDOMCounters.
//
// returns:
//   documents
//   nodes
//   jsEventListeners
func (p *GetDOMCountersParams) Do(ctxt context.Context, h cdp.FrameHandler) (documents int64, nodes int64, jsEventListeners int64, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandMemoryGetDOMCounters, cdp.Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return 0, 0, 0, cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r GetDOMCountersReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return 0, 0, 0, cdp.ErrInvalidResult
			}

			return r.Documents, r.Nodes, r.JsEventListeners, nil

		case error:
			return 0, 0, 0, v
		}

	case <-ctxt.Done():
		return 0, 0, 0, cdp.ErrContextDone
	}

	return 0, 0, 0, cdp.ErrUnknownResult
}

// SetPressureNotificationsSuppressedParams enable/disable suppressing memory
// pressure notifications in all processes.
type SetPressureNotificationsSuppressedParams struct {
	Suppressed bool `json:"suppressed"` // If true, memory pressure notifications will be suppressed.
}

// SetPressureNotificationsSuppressed enable/disable suppressing memory
// pressure notifications in all processes.
//
// parameters:
//   suppressed - If true, memory pressure notifications will be suppressed.
func SetPressureNotificationsSuppressed(suppressed bool) *SetPressureNotificationsSuppressedParams {
	return &SetPressureNotificationsSuppressedParams{
		Suppressed: suppressed,
	}
}

// Do executes Memory.setPressureNotificationsSuppressed.
func (p *SetPressureNotificationsSuppressedParams) Do(ctxt context.Context, h cdp.FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandMemorySetPressureNotificationsSuppressed, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return cdp.ErrContextDone
	}

	return cdp.ErrUnknownResult
}

// SimulatePressureNotificationParams simulate a memory pressure notification
// in all processes.
type SimulatePressureNotificationParams struct {
	Level PressureLevel `json:"level"` // Memory pressure level of the notification.
}

// SimulatePressureNotification simulate a memory pressure notification in
// all processes.
//
// parameters:
//   level - Memory pressure level of the notification.
func SimulatePressureNotification(level PressureLevel) *SimulatePressureNotificationParams {
	return &SimulatePressureNotificationParams{
		Level: level,
	}
}

// Do executes Memory.simulatePressureNotification.
func (p *SimulatePressureNotificationParams) Do(ctxt context.Context, h cdp.FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandMemorySimulatePressureNotification, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return cdp.ErrContextDone
	}

	return cdp.ErrUnknownResult
}
