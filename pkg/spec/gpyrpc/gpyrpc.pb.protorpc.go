// Code generated by protoc-gen-protorpc. DO NOT EDIT.
//
// plugin: https://github.com/chai2010/protorpc/tree/master/protoc-gen-plugin
// plugin: https://github.com/chai2010/protorpc/tree/master/protoc-gen-protorpc
//
// source: gpyrpc.proto

package gpyrpc

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
	"time"

	"github.com/chai2010/protorpc"
	"github.com/golang/protobuf/proto"
)

var (
	_ = fmt.Sprint
	_ = io.Reader(nil)
	_ = log.Print
	_ = net.Addr(nil)
	_ = rpc.Call{}
	_ = time.Second

	_ = proto.String
	_ = protorpc.Dial
)

type PROTORPC_BuiltinService interface {
	Ping(in *Ping_Args, out *Ping_Result) error
	ListMethod(in *ListMethod_Args, out *ListMethod_Result) error
}

// PROTORPC_AcceptBuiltinServiceClient accepts connections on the listener and serves requests
// for each incoming connection.  Accept blocks; the caller typically
// invokes it in a go statement.
func PROTORPC_AcceptBuiltinServiceClient(lis net.Listener, x PROTORPC_BuiltinService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("BuiltinService", x); err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeCodec(protorpc.NewServerCodec(conn))
	}
}

// PROTORPC_RegisterBuiltinService publish the given PROTORPC_BuiltinService implementation on the server.
func PROTORPC_RegisterBuiltinService(srv *rpc.Server, x PROTORPC_BuiltinService) error {
	if err := srv.RegisterName("BuiltinService", x); err != nil {
		return err
	}
	return nil
}

// PROTORPC_NewBuiltinServiceServer returns a new PROTORPC_BuiltinService Server.
func PROTORPC_NewBuiltinServiceServer(x PROTORPC_BuiltinService) *rpc.Server {
	srv := rpc.NewServer()
	if err := srv.RegisterName("BuiltinService", x); err != nil {
		log.Fatal(err)
	}
	return srv
}

// PROTORPC_ListenAndServeBuiltinService listen announces on the local network address laddr
// and serves the given BuiltinService implementation.
func PROTORPC_ListenAndServeBuiltinService(network, addr string, x PROTORPC_BuiltinService) error {
	lis, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	defer lis.Close()

	srv := rpc.NewServer()
	if err := srv.RegisterName("BuiltinService", x); err != nil {
		return err
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeCodec(protorpc.NewServerCodec(conn))
	}
}

// PROTORPC_ServeBuiltinService serves the given PROTORPC_BuiltinService implementation.
func PROTORPC_ServeBuiltinService(conn io.ReadWriteCloser, x PROTORPC_BuiltinService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("BuiltinService", x); err != nil {
		log.Fatal(err)
	}
	srv.ServeCodec(protorpc.NewServerCodec(conn))
}

type PROTORPC_BuiltinServiceClient struct {
	*rpc.Client
}

// PROTORPC_NewBuiltinServiceClient returns a PROTORPC_BuiltinService stub to handle
// requests to the set of PROTORPC_BuiltinService at the other end of the connection.
func PROTORPC_NewBuiltinServiceClient(conn io.ReadWriteCloser) *PROTORPC_BuiltinServiceClient {
	c := rpc.NewClientWithCodec(protorpc.NewClientCodec(conn))
	return &PROTORPC_BuiltinServiceClient{c}
}

func (c *PROTORPC_BuiltinServiceClient) Ping(in *Ping_Args) (out *Ping_Result, err error) {
	if in == nil {
		in = new(Ping_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Ping_Result)
	if err = c.Call("BuiltinService.Ping", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_BuiltinServiceClient) AsyncPing(in *Ping_Args, out *Ping_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(Ping_Args)
	}
	return c.Go(
		"BuiltinService.Ping",
		in, out,
		done,
	)
}

func (c *PROTORPC_BuiltinServiceClient) ListMethod(in *ListMethod_Args) (out *ListMethod_Result, err error) {
	if in == nil {
		in = new(ListMethod_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ListMethod_Result)
	if err = c.Call("BuiltinService.ListMethod", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_BuiltinServiceClient) AsyncListMethod(in *ListMethod_Args, out *ListMethod_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ListMethod_Args)
	}
	return c.Go(
		"BuiltinService.ListMethod",
		in, out,
		done,
	)
}

// PROTORPC_DialBuiltinService connects to an PROTORPC_BuiltinService at the specified network address.
func PROTORPC_DialBuiltinService(network, addr string) (*PROTORPC_BuiltinServiceClient, error) {
	c, err := protorpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &PROTORPC_BuiltinServiceClient{c}, nil
}

// PROTORPC_DialBuiltinServiceTimeout connects to an PROTORPC_BuiltinService at the specified network address.
func PROTORPC_DialBuiltinServiceTimeout(network, addr string, timeout time.Duration) (*PROTORPC_BuiltinServiceClient, error) {
	c, err := protorpc.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, err
	}
	return &PROTORPC_BuiltinServiceClient{c}, nil
}

type PROTORPC_KclvmService interface {
	Ping(in *Ping_Args, out *Ping_Result) error
	ParseFile_LarkTree(in *ParseFile_LarkTree_Args, out *ParseFile_LarkTree_Result) error
	ParseFile_AST(in *ParseFile_AST_Args, out *ParseFile_AST_Result) error
	ParseProgram_AST(in *ParseProgram_AST_Args, out *ParseProgram_AST_Result) error
	ExecProgram(in *ExecProgram_Args, out *ExecProgram_Result) error
	ResetPlugin(in *ResetPlugin_Args, out *ResetPlugin_Result) error
	FormatCode(in *FormatCode_Args, out *FormatCode_Result) error
	FormatPath(in *FormatPath_Args, out *FormatPath_Result) error
	LintPath(in *LintPath_Args, out *LintPath_Result) error
	OverrideFile(in *OverrideFile_Args, out *OverrideFile_Result) error
	EvalCode(in *EvalCode_Args, out *EvalCode_Result) error
	ResolveCode(in *ResolveCode_Args, out *ResolveCode_Result) error
	GetSchemaType(in *GetSchemaType_Args, out *GetSchemaType_Result) error
	ValidateCode(in *ValidateCode_Args, out *ValidateCode_Result) error
	SpliceCode(in *SpliceCode_Args, out *SpliceCode_Result) error
	Complete(in *Complete_Args, out *Complete_Result) error
	GoToDef(in *GoToDef_Args, out *GoToDef_Result) error
	DocumentSymbol(in *DocumentSymbol_Args, out *DocumentSymbol_Result) error
	Hover(in *Hover_Args, out *Hover_Result) error
	ListDepFiles(in *ListDepFiles_Args, out *ListDepFiles_Result) error
	ListUpStreamFiles(in *ListUpStreamFiles_Args, out *ListUpStreamFiles_Result) error
	ListDownStreamFiles(in *ListDownStreamFiles_Args, out *ListDownStreamFiles_Result) error
	LoadSettingsFiles(in *LoadSettingsFiles_Args, out *LoadSettingsFiles_Result) error
}

// PROTORPC_AcceptKclvmServiceClient accepts connections on the listener and serves requests
// for each incoming connection.  Accept blocks; the caller typically
// invokes it in a go statement.
func PROTORPC_AcceptKclvmServiceClient(lis net.Listener, x PROTORPC_KclvmService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("KclvmService", x); err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeCodec(protorpc.NewServerCodec(conn))
	}
}

// PROTORPC_RegisterKclvmService publish the given PROTORPC_KclvmService implementation on the server.
func PROTORPC_RegisterKclvmService(srv *rpc.Server, x PROTORPC_KclvmService) error {
	if err := srv.RegisterName("KclvmService", x); err != nil {
		return err
	}
	return nil
}

// PROTORPC_NewKclvmServiceServer returns a new PROTORPC_KclvmService Server.
func PROTORPC_NewKclvmServiceServer(x PROTORPC_KclvmService) *rpc.Server {
	srv := rpc.NewServer()
	if err := srv.RegisterName("KclvmService", x); err != nil {
		log.Fatal(err)
	}
	return srv
}

// PROTORPC_ListenAndServeKclvmService listen announces on the local network address laddr
// and serves the given KclvmService implementation.
func PROTORPC_ListenAndServeKclvmService(network, addr string, x PROTORPC_KclvmService) error {
	lis, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	defer lis.Close()

	srv := rpc.NewServer()
	if err := srv.RegisterName("KclvmService", x); err != nil {
		return err
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeCodec(protorpc.NewServerCodec(conn))
	}
}

// PROTORPC_ServeKclvmService serves the given PROTORPC_KclvmService implementation.
func PROTORPC_ServeKclvmService(conn io.ReadWriteCloser, x PROTORPC_KclvmService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("KclvmService", x); err != nil {
		log.Fatal(err)
	}
	srv.ServeCodec(protorpc.NewServerCodec(conn))
}

type PROTORPC_KclvmServiceClient struct {
	*rpc.Client
}

// PROTORPC_NewKclvmServiceClient returns a PROTORPC_KclvmService stub to handle
// requests to the set of PROTORPC_KclvmService at the other end of the connection.
func PROTORPC_NewKclvmServiceClient(conn io.ReadWriteCloser) *PROTORPC_KclvmServiceClient {
	c := rpc.NewClientWithCodec(protorpc.NewClientCodec(conn))
	return &PROTORPC_KclvmServiceClient{c}
}

func (c *PROTORPC_KclvmServiceClient) Ping(in *Ping_Args) (out *Ping_Result, err error) {
	if in == nil {
		in = new(Ping_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Ping_Result)
	if err = c.Call("KclvmService.Ping", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncPing(in *Ping_Args, out *Ping_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(Ping_Args)
	}
	return c.Go(
		"KclvmService.Ping",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ParseFile_LarkTree(in *ParseFile_LarkTree_Args) (out *ParseFile_LarkTree_Result, err error) {
	if in == nil {
		in = new(ParseFile_LarkTree_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ParseFile_LarkTree_Result)
	if err = c.Call("KclvmService.ParseFile_LarkTree", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncParseFile_LarkTree(in *ParseFile_LarkTree_Args, out *ParseFile_LarkTree_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ParseFile_LarkTree_Args)
	}
	return c.Go(
		"KclvmService.ParseFile_LarkTree",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ParseFile_AST(in *ParseFile_AST_Args) (out *ParseFile_AST_Result, err error) {
	if in == nil {
		in = new(ParseFile_AST_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ParseFile_AST_Result)
	if err = c.Call("KclvmService.ParseFile_AST", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncParseFile_AST(in *ParseFile_AST_Args, out *ParseFile_AST_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ParseFile_AST_Args)
	}
	return c.Go(
		"KclvmService.ParseFile_AST",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ParseProgram_AST(in *ParseProgram_AST_Args) (out *ParseProgram_AST_Result, err error) {
	if in == nil {
		in = new(ParseProgram_AST_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ParseProgram_AST_Result)
	if err = c.Call("KclvmService.ParseProgram_AST", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncParseProgram_AST(in *ParseProgram_AST_Args, out *ParseProgram_AST_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ParseProgram_AST_Args)
	}
	return c.Go(
		"KclvmService.ParseProgram_AST",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ExecProgram(in *ExecProgram_Args) (out *ExecProgram_Result, err error) {
	if in == nil {
		in = new(ExecProgram_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ExecProgram_Result)
	if err = c.Call("KclvmService.ExecProgram", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncExecProgram(in *ExecProgram_Args, out *ExecProgram_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ExecProgram_Args)
	}
	return c.Go(
		"KclvmService.ExecProgram",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ResetPlugin(in *ResetPlugin_Args) (out *ResetPlugin_Result, err error) {
	if in == nil {
		in = new(ResetPlugin_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ResetPlugin_Result)
	if err = c.Call("KclvmService.ResetPlugin", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncResetPlugin(in *ResetPlugin_Args, out *ResetPlugin_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ResetPlugin_Args)
	}
	return c.Go(
		"KclvmService.ResetPlugin",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) FormatCode(in *FormatCode_Args) (out *FormatCode_Result, err error) {
	if in == nil {
		in = new(FormatCode_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(FormatCode_Result)
	if err = c.Call("KclvmService.FormatCode", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncFormatCode(in *FormatCode_Args, out *FormatCode_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(FormatCode_Args)
	}
	return c.Go(
		"KclvmService.FormatCode",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) FormatPath(in *FormatPath_Args) (out *FormatPath_Result, err error) {
	if in == nil {
		in = new(FormatPath_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(FormatPath_Result)
	if err = c.Call("KclvmService.FormatPath", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncFormatPath(in *FormatPath_Args, out *FormatPath_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(FormatPath_Args)
	}
	return c.Go(
		"KclvmService.FormatPath",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) LintPath(in *LintPath_Args) (out *LintPath_Result, err error) {
	if in == nil {
		in = new(LintPath_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(LintPath_Result)
	if err = c.Call("KclvmService.LintPath", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncLintPath(in *LintPath_Args, out *LintPath_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(LintPath_Args)
	}
	return c.Go(
		"KclvmService.LintPath",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) OverrideFile(in *OverrideFile_Args) (out *OverrideFile_Result, err error) {
	if in == nil {
		in = new(OverrideFile_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(OverrideFile_Result)
	if err = c.Call("KclvmService.OverrideFile", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncOverrideFile(in *OverrideFile_Args, out *OverrideFile_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(OverrideFile_Args)
	}
	return c.Go(
		"KclvmService.OverrideFile",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) EvalCode(in *EvalCode_Args) (out *EvalCode_Result, err error) {
	if in == nil {
		in = new(EvalCode_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(EvalCode_Result)
	if err = c.Call("KclvmService.EvalCode", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncEvalCode(in *EvalCode_Args, out *EvalCode_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(EvalCode_Args)
	}
	return c.Go(
		"KclvmService.EvalCode",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ResolveCode(in *ResolveCode_Args) (out *ResolveCode_Result, err error) {
	if in == nil {
		in = new(ResolveCode_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ResolveCode_Result)
	if err = c.Call("KclvmService.ResolveCode", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncResolveCode(in *ResolveCode_Args, out *ResolveCode_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ResolveCode_Args)
	}
	return c.Go(
		"KclvmService.ResolveCode",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) GetSchemaType(in *GetSchemaType_Args) (out *GetSchemaType_Result, err error) {
	if in == nil {
		in = new(GetSchemaType_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(GetSchemaType_Result)
	if err = c.Call("KclvmService.GetSchemaType", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncGetSchemaType(in *GetSchemaType_Args, out *GetSchemaType_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(GetSchemaType_Args)
	}
	return c.Go(
		"KclvmService.GetSchemaType",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ValidateCode(in *ValidateCode_Args) (out *ValidateCode_Result, err error) {
	if in == nil {
		in = new(ValidateCode_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ValidateCode_Result)
	if err = c.Call("KclvmService.ValidateCode", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncValidateCode(in *ValidateCode_Args, out *ValidateCode_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ValidateCode_Args)
	}
	return c.Go(
		"KclvmService.ValidateCode",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) SpliceCode(in *SpliceCode_Args) (out *SpliceCode_Result, err error) {
	if in == nil {
		in = new(SpliceCode_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(SpliceCode_Result)
	if err = c.Call("KclvmService.SpliceCode", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncSpliceCode(in *SpliceCode_Args, out *SpliceCode_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(SpliceCode_Args)
	}
	return c.Go(
		"KclvmService.SpliceCode",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) Complete(in *Complete_Args) (out *Complete_Result, err error) {
	if in == nil {
		in = new(Complete_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Complete_Result)
	if err = c.Call("KclvmService.Complete", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncComplete(in *Complete_Args, out *Complete_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(Complete_Args)
	}
	return c.Go(
		"KclvmService.Complete",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) GoToDef(in *GoToDef_Args) (out *GoToDef_Result, err error) {
	if in == nil {
		in = new(GoToDef_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(GoToDef_Result)
	if err = c.Call("KclvmService.GoToDef", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncGoToDef(in *GoToDef_Args, out *GoToDef_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(GoToDef_Args)
	}
	return c.Go(
		"KclvmService.GoToDef",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) DocumentSymbol(in *DocumentSymbol_Args) (out *DocumentSymbol_Result, err error) {
	if in == nil {
		in = new(DocumentSymbol_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(DocumentSymbol_Result)
	if err = c.Call("KclvmService.DocumentSymbol", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncDocumentSymbol(in *DocumentSymbol_Args, out *DocumentSymbol_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(DocumentSymbol_Args)
	}
	return c.Go(
		"KclvmService.DocumentSymbol",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) Hover(in *Hover_Args) (out *Hover_Result, err error) {
	if in == nil {
		in = new(Hover_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Hover_Result)
	if err = c.Call("KclvmService.Hover", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncHover(in *Hover_Args, out *Hover_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(Hover_Args)
	}
	return c.Go(
		"KclvmService.Hover",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ListDepFiles(in *ListDepFiles_Args) (out *ListDepFiles_Result, err error) {
	if in == nil {
		in = new(ListDepFiles_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ListDepFiles_Result)
	if err = c.Call("KclvmService.ListDepFiles", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncListDepFiles(in *ListDepFiles_Args, out *ListDepFiles_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ListDepFiles_Args)
	}
	return c.Go(
		"KclvmService.ListDepFiles",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ListUpStreamFiles(in *ListUpStreamFiles_Args) (out *ListUpStreamFiles_Result, err error) {
	if in == nil {
		in = new(ListUpStreamFiles_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ListUpStreamFiles_Result)
	if err = c.Call("KclvmService.ListUpStreamFiles", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncListUpStreamFiles(in *ListUpStreamFiles_Args, out *ListUpStreamFiles_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ListUpStreamFiles_Args)
	}
	return c.Go(
		"KclvmService.ListUpStreamFiles",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ListDownStreamFiles(in *ListDownStreamFiles_Args) (out *ListDownStreamFiles_Result, err error) {
	if in == nil {
		in = new(ListDownStreamFiles_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ListDownStreamFiles_Result)
	if err = c.Call("KclvmService.ListDownStreamFiles", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncListDownStreamFiles(in *ListDownStreamFiles_Args, out *ListDownStreamFiles_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ListDownStreamFiles_Args)
	}
	return c.Go(
		"KclvmService.ListDownStreamFiles",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) LoadSettingsFiles(in *LoadSettingsFiles_Args) (out *LoadSettingsFiles_Result, err error) {
	if in == nil {
		in = new(LoadSettingsFiles_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(LoadSettingsFiles_Result)
	if err = c.Call("KclvmService.LoadSettingsFiles", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncLoadSettingsFiles(in *LoadSettingsFiles_Args, out *LoadSettingsFiles_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(LoadSettingsFiles_Args)
	}
	return c.Go(
		"KclvmService.LoadSettingsFiles",
		in, out,
		done,
	)
}

// PROTORPC_DialKclvmService connects to an PROTORPC_KclvmService at the specified network address.
func PROTORPC_DialKclvmService(network, addr string) (*PROTORPC_KclvmServiceClient, error) {
	c, err := protorpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &PROTORPC_KclvmServiceClient{c}, nil
}

// PROTORPC_DialKclvmServiceTimeout connects to an PROTORPC_KclvmService at the specified network address.
func PROTORPC_DialKclvmServiceTimeout(network, addr string, timeout time.Duration) (*PROTORPC_KclvmServiceClient, error) {
	c, err := protorpc.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, err
	}
	return &PROTORPC_KclvmServiceClient{c}, nil
}
