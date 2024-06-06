//go:generate mockgen -source=./email_grpc.pb.go -destination=../mock/email_grpc_mock.go -package=mock proto EmailServiceClient

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: email.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	EmailService_GetAllIncoming_FullMethodName      = "/proto.EmailService/GetAllIncoming"
	EmailService_GetAllSent_FullMethodName          = "/proto.EmailService/GetAllSent"
	EmailService_GetDraftEmails_FullMethodName      = "/proto.EmailService/GetDraftEmails"
	EmailService_GetSpamEmails_FullMethodName       = "/proto.EmailService/GetSpamEmails"
	EmailService_GetEmailByID_FullMethodName        = "/proto.EmailService/GetEmailByID"
	EmailService_CreateEmail_FullMethodName         = "/proto.EmailService/CreateEmail"
	EmailService_CreateProfileEmail_FullMethodName  = "/proto.EmailService/CreateProfileEmail"
	EmailService_CheckRecipientEmail_FullMethodName = "/proto.EmailService/CheckRecipientEmail"
	EmailService_UpdateEmail_FullMethodName         = "/proto.EmailService/UpdateEmail"
	EmailService_DeleteEmail_FullMethodName         = "/proto.EmailService/DeleteEmail"
	EmailService_AddEmailDraft_FullMethodName       = "/proto.EmailService/AddEmailDraft"
	EmailService_AddAttachment_FullMethodName       = "/proto.EmailService/AddAttachment"
	EmailService_GetFileByID_FullMethodName         = "/proto.EmailService/GetFileByID"
	EmailService_GetFilesByEmailID_FullMethodName   = "/proto.EmailService/GetFilesByEmailID"
	EmailService_DeleteFileByID_FullMethodName      = "/proto.EmailService/DeleteFileByID"
	EmailService_UpdateFileByID_FullMethodName      = "/proto.EmailService/UpdateFileByID"
	EmailService_AddFile_FullMethodName             = "/proto.EmailService/AddFile"
	EmailService_AddFileToEmail_FullMethodName      = "/proto.EmailService/AddFileToEmail"
)

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailServiceClient interface {
	GetAllIncoming(ctx context.Context, in *LoginOffsetLimit, opts ...grpc.CallOption) (*Emails, error)
	GetAllSent(ctx context.Context, in *LoginOffsetLimit, opts ...grpc.CallOption) (*Emails, error)
	GetDraftEmails(ctx context.Context, in *LoginOffsetLimit, opts ...grpc.CallOption) (*Emails, error)
	GetSpamEmails(ctx context.Context, in *LoginOffsetLimit, opts ...grpc.CallOption) (*Emails, error)
	GetEmailByID(ctx context.Context, in *EmailIdAndLogin, opts ...grpc.CallOption) (*Email, error)
	CreateEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*EmailWithID, error)
	CreateProfileEmail(ctx context.Context, in *IdSenderRecipient, opts ...grpc.CallOption) (*EmptyEmail, error)
	CheckRecipientEmail(ctx context.Context, in *Recipient, opts ...grpc.CallOption) (*EmptyEmail, error)
	UpdateEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*StatusEmail, error)
	DeleteEmail(ctx context.Context, in *LoginWithID, opts ...grpc.CallOption) (*StatusEmail, error)
	AddEmailDraft(ctx context.Context, in *Email, opts ...grpc.CallOption) (*EmailWithID, error)
	AddAttachment(ctx context.Context, in *AddAttachmentRequest, opts ...grpc.CallOption) (*AddAttachmentReply, error)
	GetFileByID(ctx context.Context, in *GetFileByIDRequest, opts ...grpc.CallOption) (*GetFileByIDReply, error)
	GetFilesByEmailID(ctx context.Context, in *GetFilesByEmailIDRequest, opts ...grpc.CallOption) (*GetFilesByEmailIDReply, error)
	DeleteFileByID(ctx context.Context, in *DeleteFileByIDRequest, opts ...grpc.CallOption) (*DeleteFileByIDReply, error)
	UpdateFileByID(ctx context.Context, in *UpdateFileByIDRequest, opts ...grpc.CallOption) (*UpdateFileByIDReply, error)
	AddFile(ctx context.Context, in *AddFileRequest, opts ...grpc.CallOption) (*AddFileReply, error)
	AddFileToEmail(ctx context.Context, in *AddFileToEmailRequest, opts ...grpc.CallOption) (*AddFileToEmailReply, error)
}

type emailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailServiceClient(cc grpc.ClientConnInterface) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) GetAllIncoming(ctx context.Context, in *LoginOffsetLimit, opts ...grpc.CallOption) (*Emails, error) {
	out := new(Emails)
	err := c.cc.Invoke(ctx, EmailService_GetAllIncoming_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) GetAllSent(ctx context.Context, in *LoginOffsetLimit, opts ...grpc.CallOption) (*Emails, error) {
	out := new(Emails)
	err := c.cc.Invoke(ctx, EmailService_GetAllSent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) GetDraftEmails(ctx context.Context, in *LoginOffsetLimit, opts ...grpc.CallOption) (*Emails, error) {
	out := new(Emails)
	err := c.cc.Invoke(ctx, EmailService_GetDraftEmails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) GetSpamEmails(ctx context.Context, in *LoginOffsetLimit, opts ...grpc.CallOption) (*Emails, error) {
	out := new(Emails)
	err := c.cc.Invoke(ctx, EmailService_GetSpamEmails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) GetEmailByID(ctx context.Context, in *EmailIdAndLogin, opts ...grpc.CallOption) (*Email, error) {
	out := new(Email)
	err := c.cc.Invoke(ctx, EmailService_GetEmailByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) CreateEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*EmailWithID, error) {
	out := new(EmailWithID)
	err := c.cc.Invoke(ctx, EmailService_CreateEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) CreateProfileEmail(ctx context.Context, in *IdSenderRecipient, opts ...grpc.CallOption) (*EmptyEmail, error) {
	out := new(EmptyEmail)
	err := c.cc.Invoke(ctx, EmailService_CreateProfileEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) CheckRecipientEmail(ctx context.Context, in *Recipient, opts ...grpc.CallOption) (*EmptyEmail, error) {
	out := new(EmptyEmail)
	err := c.cc.Invoke(ctx, EmailService_CheckRecipientEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) UpdateEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*StatusEmail, error) {
	out := new(StatusEmail)
	err := c.cc.Invoke(ctx, EmailService_UpdateEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) DeleteEmail(ctx context.Context, in *LoginWithID, opts ...grpc.CallOption) (*StatusEmail, error) {
	out := new(StatusEmail)
	err := c.cc.Invoke(ctx, EmailService_DeleteEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) AddEmailDraft(ctx context.Context, in *Email, opts ...grpc.CallOption) (*EmailWithID, error) {
	out := new(EmailWithID)
	err := c.cc.Invoke(ctx, EmailService_AddEmailDraft_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) AddAttachment(ctx context.Context, in *AddAttachmentRequest, opts ...grpc.CallOption) (*AddAttachmentReply, error) {
	out := new(AddAttachmentReply)
	err := c.cc.Invoke(ctx, EmailService_AddAttachment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) GetFileByID(ctx context.Context, in *GetFileByIDRequest, opts ...grpc.CallOption) (*GetFileByIDReply, error) {
	out := new(GetFileByIDReply)
	err := c.cc.Invoke(ctx, EmailService_GetFileByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) GetFilesByEmailID(ctx context.Context, in *GetFilesByEmailIDRequest, opts ...grpc.CallOption) (*GetFilesByEmailIDReply, error) {
	out := new(GetFilesByEmailIDReply)
	err := c.cc.Invoke(ctx, EmailService_GetFilesByEmailID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) DeleteFileByID(ctx context.Context, in *DeleteFileByIDRequest, opts ...grpc.CallOption) (*DeleteFileByIDReply, error) {
	out := new(DeleteFileByIDReply)
	err := c.cc.Invoke(ctx, EmailService_DeleteFileByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) UpdateFileByID(ctx context.Context, in *UpdateFileByIDRequest, opts ...grpc.CallOption) (*UpdateFileByIDReply, error) {
	out := new(UpdateFileByIDReply)
	err := c.cc.Invoke(ctx, EmailService_UpdateFileByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) AddFile(ctx context.Context, in *AddFileRequest, opts ...grpc.CallOption) (*AddFileReply, error) {
	out := new(AddFileReply)
	err := c.cc.Invoke(ctx, EmailService_AddFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) AddFileToEmail(ctx context.Context, in *AddFileToEmailRequest, opts ...grpc.CallOption) (*AddFileToEmailReply, error) {
	out := new(AddFileToEmailReply)
	err := c.cc.Invoke(ctx, EmailService_AddFileToEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceServer is the server API for EmailService service.
// All implementations must embed UnimplementedEmailServiceServer
// for forward compatibility
type EmailServiceServer interface {
	GetAllIncoming(context.Context, *LoginOffsetLimit) (*Emails, error)
	GetAllSent(context.Context, *LoginOffsetLimit) (*Emails, error)
	GetDraftEmails(context.Context, *LoginOffsetLimit) (*Emails, error)
	GetSpamEmails(context.Context, *LoginOffsetLimit) (*Emails, error)
	GetEmailByID(context.Context, *EmailIdAndLogin) (*Email, error)
	CreateEmail(context.Context, *Email) (*EmailWithID, error)
	CreateProfileEmail(context.Context, *IdSenderRecipient) (*EmptyEmail, error)
	CheckRecipientEmail(context.Context, *Recipient) (*EmptyEmail, error)
	UpdateEmail(context.Context, *Email) (*StatusEmail, error)
	DeleteEmail(context.Context, *LoginWithID) (*StatusEmail, error)
	AddEmailDraft(context.Context, *Email) (*EmailWithID, error)
	AddAttachment(context.Context, *AddAttachmentRequest) (*AddAttachmentReply, error)
	GetFileByID(context.Context, *GetFileByIDRequest) (*GetFileByIDReply, error)
	GetFilesByEmailID(context.Context, *GetFilesByEmailIDRequest) (*GetFilesByEmailIDReply, error)
	DeleteFileByID(context.Context, *DeleteFileByIDRequest) (*DeleteFileByIDReply, error)
	UpdateFileByID(context.Context, *UpdateFileByIDRequest) (*UpdateFileByIDReply, error)
	AddFile(context.Context, *AddFileRequest) (*AddFileReply, error)
	AddFileToEmail(context.Context, *AddFileToEmailRequest) (*AddFileToEmailReply, error)
	mustEmbedUnimplementedEmailServiceServer()
}

// UnimplementedEmailServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmailServiceServer struct {
}

func (UnimplementedEmailServiceServer) GetAllIncoming(context.Context, *LoginOffsetLimit) (*Emails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllIncoming not implemented")
}
func (UnimplementedEmailServiceServer) GetAllSent(context.Context, *LoginOffsetLimit) (*Emails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSent not implemented")
}
func (UnimplementedEmailServiceServer) GetDraftEmails(context.Context, *LoginOffsetLimit) (*Emails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDraftEmails not implemented")
}
func (UnimplementedEmailServiceServer) GetSpamEmails(context.Context, *LoginOffsetLimit) (*Emails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpamEmails not implemented")
}
func (UnimplementedEmailServiceServer) GetEmailByID(context.Context, *EmailIdAndLogin) (*Email, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmailByID not implemented")
}
func (UnimplementedEmailServiceServer) CreateEmail(context.Context, *Email) (*EmailWithID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmail not implemented")
}
func (UnimplementedEmailServiceServer) CreateProfileEmail(context.Context, *IdSenderRecipient) (*EmptyEmail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfileEmail not implemented")
}
func (UnimplementedEmailServiceServer) CheckRecipientEmail(context.Context, *Recipient) (*EmptyEmail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRecipientEmail not implemented")
}
func (UnimplementedEmailServiceServer) UpdateEmail(context.Context, *Email) (*StatusEmail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmail not implemented")
}
func (UnimplementedEmailServiceServer) DeleteEmail(context.Context, *LoginWithID) (*StatusEmail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmail not implemented")
}
func (UnimplementedEmailServiceServer) AddEmailDraft(context.Context, *Email) (*EmailWithID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEmailDraft not implemented")
}
func (UnimplementedEmailServiceServer) AddAttachment(context.Context, *AddAttachmentRequest) (*AddAttachmentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAttachment not implemented")
}
func (UnimplementedEmailServiceServer) GetFileByID(context.Context, *GetFileByIDRequest) (*GetFileByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileByID not implemented")
}
func (UnimplementedEmailServiceServer) GetFilesByEmailID(context.Context, *GetFilesByEmailIDRequest) (*GetFilesByEmailIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilesByEmailID not implemented")
}
func (UnimplementedEmailServiceServer) DeleteFileByID(context.Context, *DeleteFileByIDRequest) (*DeleteFileByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFileByID not implemented")
}
func (UnimplementedEmailServiceServer) UpdateFileByID(context.Context, *UpdateFileByIDRequest) (*UpdateFileByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFileByID not implemented")
}
func (UnimplementedEmailServiceServer) AddFile(context.Context, *AddFileRequest) (*AddFileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFile not implemented")
}
func (UnimplementedEmailServiceServer) AddFileToEmail(context.Context, *AddFileToEmailRequest) (*AddFileToEmailReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFileToEmail not implemented")
}
func (UnimplementedEmailServiceServer) mustEmbedUnimplementedEmailServiceServer() {}

// UnsafeEmailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailServiceServer will
// result in compilation errors.
type UnsafeEmailServiceServer interface {
	mustEmbedUnimplementedEmailServiceServer()
}

func RegisterEmailServiceServer(s grpc.ServiceRegistrar, srv EmailServiceServer) {
	s.RegisterService(&EmailService_ServiceDesc, srv)
}

func _EmailService_GetAllIncoming_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginOffsetLimit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).GetAllIncoming(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_GetAllIncoming_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).GetAllIncoming(ctx, req.(*LoginOffsetLimit))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_GetAllSent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginOffsetLimit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).GetAllSent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_GetAllSent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).GetAllSent(ctx, req.(*LoginOffsetLimit))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_GetDraftEmails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginOffsetLimit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).GetDraftEmails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_GetDraftEmails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).GetDraftEmails(ctx, req.(*LoginOffsetLimit))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_GetSpamEmails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginOffsetLimit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).GetSpamEmails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_GetSpamEmails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).GetSpamEmails(ctx, req.(*LoginOffsetLimit))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_GetEmailByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailIdAndLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).GetEmailByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_GetEmailByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).GetEmailByID(ctx, req.(*EmailIdAndLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_CreateEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).CreateEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_CreateEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).CreateEmail(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_CreateProfileEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdSenderRecipient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).CreateProfileEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_CreateProfileEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).CreateProfileEmail(ctx, req.(*IdSenderRecipient))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_CheckRecipientEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Recipient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).CheckRecipientEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_CheckRecipientEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).CheckRecipientEmail(ctx, req.(*Recipient))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_UpdateEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).UpdateEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_UpdateEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).UpdateEmail(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_DeleteEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginWithID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).DeleteEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_DeleteEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).DeleteEmail(ctx, req.(*LoginWithID))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_AddEmailDraft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).AddEmailDraft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_AddEmailDraft_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).AddEmailDraft(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_AddAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAttachmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).AddAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_AddAttachment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).AddAttachment(ctx, req.(*AddAttachmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_GetFileByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).GetFileByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_GetFileByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).GetFileByID(ctx, req.(*GetFileByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_GetFilesByEmailID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilesByEmailIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).GetFilesByEmailID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_GetFilesByEmailID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).GetFilesByEmailID(ctx, req.(*GetFilesByEmailIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_DeleteFileByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).DeleteFileByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_DeleteFileByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).DeleteFileByID(ctx, req.(*DeleteFileByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_UpdateFileByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFileByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).UpdateFileByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_UpdateFileByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).UpdateFileByID(ctx, req.(*UpdateFileByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_AddFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).AddFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_AddFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).AddFile(ctx, req.(*AddFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_AddFileToEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFileToEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).AddFileToEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_AddFileToEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).AddFileToEmail(ctx, req.(*AddFileToEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailService_ServiceDesc is the grpc.ServiceDesc for EmailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllIncoming",
			Handler:    _EmailService_GetAllIncoming_Handler,
		},
		{
			MethodName: "GetAllSent",
			Handler:    _EmailService_GetAllSent_Handler,
		},
		{
			MethodName: "GetDraftEmails",
			Handler:    _EmailService_GetDraftEmails_Handler,
		},
		{
			MethodName: "GetSpamEmails",
			Handler:    _EmailService_GetSpamEmails_Handler,
		},
		{
			MethodName: "GetEmailByID",
			Handler:    _EmailService_GetEmailByID_Handler,
		},
		{
			MethodName: "CreateEmail",
			Handler:    _EmailService_CreateEmail_Handler,
		},
		{
			MethodName: "CreateProfileEmail",
			Handler:    _EmailService_CreateProfileEmail_Handler,
		},
		{
			MethodName: "CheckRecipientEmail",
			Handler:    _EmailService_CheckRecipientEmail_Handler,
		},
		{
			MethodName: "UpdateEmail",
			Handler:    _EmailService_UpdateEmail_Handler,
		},
		{
			MethodName: "DeleteEmail",
			Handler:    _EmailService_DeleteEmail_Handler,
		},
		{
			MethodName: "AddEmailDraft",
			Handler:    _EmailService_AddEmailDraft_Handler,
		},
		{
			MethodName: "AddAttachment",
			Handler:    _EmailService_AddAttachment_Handler,
		},
		{
			MethodName: "GetFileByID",
			Handler:    _EmailService_GetFileByID_Handler,
		},
		{
			MethodName: "GetFilesByEmailID",
			Handler:    _EmailService_GetFilesByEmailID_Handler,
		},
		{
			MethodName: "DeleteFileByID",
			Handler:    _EmailService_DeleteFileByID_Handler,
		},
		{
			MethodName: "UpdateFileByID",
			Handler:    _EmailService_UpdateFileByID_Handler,
		},
		{
			MethodName: "AddFile",
			Handler:    _EmailService_AddFile_Handler,
		},
		{
			MethodName: "AddFileToEmail",
			Handler:    _EmailService_AddFileToEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email.proto",
}