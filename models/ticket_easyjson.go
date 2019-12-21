// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonB39c1f48DecodeKinoBackendModels(in *jlexer.Lexer, out *TicketProfile) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ticket_id":
			out.TicketID = uint(in.Uint())
		case "user_id":
			out.UserID = uint(in.Uint())
		case "ms_id":
			out.MSID = uint(in.Uint())
		case "seat_id":
			out.SeatID = uint(in.Uint())
		case "price":
			out.Price = uint(in.Uint())
		case "start_datetime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Date).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB39c1f48EncodeKinoBackendModels(out *jwriter.Writer, in TicketProfile) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ticket_id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.TicketID))
	}
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix)
		out.Uint(uint(in.UserID))
	}
	{
		const prefix string = ",\"ms_id\":"
		out.RawString(prefix)
		out.Uint(uint(in.MSID))
	}
	{
		const prefix string = ",\"seat_id\":"
		out.RawString(prefix)
		out.Uint(uint(in.SeatID))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.Uint(uint(in.Price))
	}
	{
		const prefix string = ",\"start_datetime\":"
		out.RawString(prefix)
		out.Raw((in.Date).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TicketProfile) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB39c1f48EncodeKinoBackendModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TicketProfile) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB39c1f48EncodeKinoBackendModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TicketProfile) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB39c1f48DecodeKinoBackendModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TicketProfile) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB39c1f48DecodeKinoBackendModels(l, v)
}
func easyjsonB39c1f48DecodeKinoBackendModels1(in *jlexer.Lexer, out *Ticket) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ticket_id":
			out.TicketID = uint(in.Uint())
		case "user_id":
			out.UserID = uint(in.Uint())
		case "ms_id":
			out.MSID = uint(in.Uint())
		case "seat_id":
			out.SeatID = uint(in.Uint())
		case "price":
			out.Price = uint(in.Uint())
		case "start_datetime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Date).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB39c1f48EncodeKinoBackendModels1(out *jwriter.Writer, in Ticket) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ticket_id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.TicketID))
	}
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix)
		out.Uint(uint(in.UserID))
	}
	{
		const prefix string = ",\"ms_id\":"
		out.RawString(prefix)
		out.Uint(uint(in.MSID))
	}
	{
		const prefix string = ",\"seat_id\":"
		out.RawString(prefix)
		out.Uint(uint(in.SeatID))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.Uint(uint(in.Price))
	}
	{
		const prefix string = ",\"start_datetime\":"
		out.RawString(prefix)
		out.Raw((in.Date).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Ticket) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB39c1f48EncodeKinoBackendModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Ticket) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB39c1f48EncodeKinoBackendModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Ticket) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB39c1f48DecodeKinoBackendModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Ticket) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB39c1f48DecodeKinoBackendModels1(l, v)
}
func easyjsonB39c1f48DecodeKinoBackendModels2(in *jlexer.Lexer, out *Seat) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "seat_id":
			out.SeatID = uint(in.Uint())
		case "hall_name":
			out.HallName = string(in.String())
		case "movie_session_id":
			out.MovieSessionID = int(in.Int())
		case "is_taken":
			out.IsTaken = bool(in.Bool())
		case "row":
			out.Row = int(in.Int())
		case "seat_number":
			out.SeatNumber = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB39c1f48EncodeKinoBackendModels2(out *jwriter.Writer, in Seat) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"seat_id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.SeatID))
	}
	{
		const prefix string = ",\"hall_name\":"
		out.RawString(prefix)
		out.String(string(in.HallName))
	}
	{
		const prefix string = ",\"movie_session_id\":"
		out.RawString(prefix)
		out.Int(int(in.MovieSessionID))
	}
	{
		const prefix string = ",\"is_taken\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsTaken))
	}
	{
		const prefix string = ",\"row\":"
		out.RawString(prefix)
		out.Int(int(in.Row))
	}
	{
		const prefix string = ",\"seat_number\":"
		out.RawString(prefix)
		out.Int(int(in.SeatNumber))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Seat) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB39c1f48EncodeKinoBackendModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Seat) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB39c1f48EncodeKinoBackendModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Seat) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB39c1f48DecodeKinoBackendModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Seat) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB39c1f48DecodeKinoBackendModels2(l, v)
}
func easyjsonB39c1f48DecodeKinoBackendModels3(in *jlexer.Lexer, out *RequestTicket) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ticket_id":
			out.TicketID = uint(in.Uint())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB39c1f48EncodeKinoBackendModels3(out *jwriter.Writer, in RequestTicket) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ticket_id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.TicketID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestTicket) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB39c1f48EncodeKinoBackendModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestTicket) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB39c1f48EncodeKinoBackendModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestTicket) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB39c1f48DecodeKinoBackendModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestTicket) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB39c1f48DecodeKinoBackendModels3(l, v)
}
func easyjsonB39c1f48DecodeKinoBackendModels4(in *jlexer.Lexer, out *RegisterTicket) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "user_id":
			out.UserID = uint(in.Uint())
		case "ms_id":
			out.MSID = uint(in.Uint())
		case "seat_id":
			out.SeatID = uint(in.Uint())
		case "price":
			out.Price = uint(in.Uint())
		case "start_datetime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Date).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB39c1f48EncodeKinoBackendModels4(out *jwriter.Writer, in RegisterTicket) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.UserID))
	}
	{
		const prefix string = ",\"ms_id\":"
		out.RawString(prefix)
		out.Uint(uint(in.MSID))
	}
	{
		const prefix string = ",\"seat_id\":"
		out.RawString(prefix)
		out.Uint(uint(in.SeatID))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.Uint(uint(in.Price))
	}
	{
		const prefix string = ",\"start_datetime\":"
		out.RawString(prefix)
		out.Raw((in.Date).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RegisterTicket) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB39c1f48EncodeKinoBackendModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RegisterTicket) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB39c1f48EncodeKinoBackendModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RegisterTicket) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB39c1f48DecodeKinoBackendModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RegisterTicket) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB39c1f48DecodeKinoBackendModels4(l, v)
}
