// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package delivery_models

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

func easyjson408d1214DecodeMailInternalModelsDeliveryModels(in *jlexer.Lexer, out *FolderEmail) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "folderId":
			out.FolderID = uint32(in.Uint32())
		case "emailId":
			out.EmailID = uint32(in.Uint32())
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
func easyjson408d1214EncodeMailInternalModelsDeliveryModels(out *jwriter.Writer, in FolderEmail) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"folderId\":"
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.FolderID))
	}
	{
		const prefix string = ",\"emailId\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.EmailID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FolderEmail) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson408d1214EncodeMailInternalModelsDeliveryModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FolderEmail) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson408d1214EncodeMailInternalModelsDeliveryModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FolderEmail) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson408d1214DecodeMailInternalModelsDeliveryModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FolderEmail) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson408d1214DecodeMailInternalModelsDeliveryModels(l, v)
}
func easyjson408d1214DecodeMailInternalModelsDeliveryModels1(in *jlexer.Lexer, out *Folder) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint32(in.Uint32())
		case "profileId":
			out.ProfileId = uint32(in.Uint32())
		case "name":
			out.Name = string(in.String())
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
func easyjson408d1214EncodeMailInternalModelsDeliveryModels1(out *jwriter.Writer, in Folder) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.ID))
	}
	if in.ProfileId != 0 {
		const prefix string = ",\"profileId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(in.ProfileId))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Folder) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson408d1214EncodeMailInternalModelsDeliveryModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Folder) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson408d1214EncodeMailInternalModelsDeliveryModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Folder) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson408d1214DecodeMailInternalModelsDeliveryModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Folder) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson408d1214DecodeMailInternalModelsDeliveryModels1(l, v)
}