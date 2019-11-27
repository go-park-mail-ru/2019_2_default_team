// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package comments_microservice_client

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

func easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient(in *jlexer.Lexer, out *commentsManagerClient) {
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
func easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient(out *jwriter.Writer, in commentsManagerClient) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v commentsManagerClient) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v commentsManagerClient) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *commentsManagerClient) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *commentsManagerClient) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient(l, v)
}
func easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient1(in *jlexer.Lexer, out *UnimplementedCommentsManagerServer) {
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
func easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient1(out *jwriter.Writer, in UnimplementedCommentsManagerServer) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UnimplementedCommentsManagerServer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UnimplementedCommentsManagerServer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UnimplementedCommentsManagerServer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UnimplementedCommentsManagerServer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient1(l, v)
}
func easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient2(in *jlexer.Lexer, out *Nothing) {
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
func easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient2(out *jwriter.Writer, in Nothing) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Nothing) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Nothing) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Nothing) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Nothing) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient2(l, v)
}
func easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient3(in *jlexer.Lexer, out *CommentsResponse) {
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
		case "comments":
			if in.IsNull() {
				in.Skip()
				out.Comments = nil
			} else {
				in.Delim('[')
				if out.Comments == nil {
					if !in.IsDelim(']') {
						out.Comments = make([]*Comment, 0, 8)
					} else {
						out.Comments = []*Comment{}
					}
				} else {
					out.Comments = (out.Comments)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Comment
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Comment)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Comments = append(out.Comments, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient3(out *jwriter.Writer, in CommentsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Comments) != 0 {
		const prefix string = ",\"comments\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v2, v3 := range in.Comments {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CommentsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CommentsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CommentsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CommentsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient3(l, v)
}
func easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient4(in *jlexer.Lexer, out *CommentID) {
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
		case "Film":
			out.Film = string(in.String())
		case "User":
			out.User = string(in.String())
		case "CID":
			out.CID = uint64(in.Uint64())
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
func easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient4(out *jwriter.Writer, in CommentID) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Film != "" {
		const prefix string = ",\"Film\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Film))
	}
	if in.User != "" {
		const prefix string = ",\"User\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.User))
	}
	if in.CID != 0 {
		const prefix string = ",\"CID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.CID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CommentID) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CommentID) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CommentID) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CommentID) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient4(l, v)
}
func easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient5(in *jlexer.Lexer, out *Comment) {
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
		case "Username":
			out.Username = string(in.String())
		case "FilmTitle":
			out.FilmTitle = string(in.String())
		case "Text":
			out.Text = string(in.String())
		case "ID":
			out.ID = uint64(in.Uint64())
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
func easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient5(out *jwriter.Writer, in Comment) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Username != "" {
		const prefix string = ",\"Username\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Username))
	}
	if in.FilmTitle != "" {
		const prefix string = ",\"FilmTitle\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FilmTitle))
	}
	if in.Text != "" {
		const prefix string = ",\"Text\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Text))
	}
	if in.ID != 0 {
		const prefix string = ",\"ID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Comment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Comment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson873bff8EncodeKinoBackendCommentsMicroserviceClient5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Comment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Comment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson873bff8DecodeKinoBackendCommentsMicroserviceClient5(l, v)
}
