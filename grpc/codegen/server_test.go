package codegen

import (
	"testing"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/expr"
	"goa.design/goa/v3/grpc/codegen/testdata"
)

func TestServerGRPCInterface(t *testing.T) {
	cases := []struct {
		Name string
		DSL  func()
		Code string
	}{
		{"unary-rpcs", testdata.UnaryRPCsDSL, testdata.UnaryRPCsServerInterfaceCode},
		{"unary-rpc-no-payload", testdata.UnaryRPCNoPayloadDSL, testdata.UnaryRPCNoPayloadServerInterfaceCode},
		{"unary-rpc-no-result", testdata.UnaryRPCNoResultDSL, testdata.UnaryRPCNoResultServerInterfaceCode},
		{"unary-rpc-with-errors", testdata.UnaryRPCWithErrorsDSL, testdata.UnaryRPCWithErrorsServerInterfaceCode},
		{"unary-rpc-with-overriding-errors", testdata.UnaryRPCWithOverridingErrorsDSL, testdata.UnaryRPCWithOverridingErrorsServerInterfaceCode},
		{"server-streaming-rpc", testdata.ServerStreamingRPCDSL, testdata.ServerStreamingRPCServerInterfaceCode},
		{"client-streaming-rpc", testdata.ClientStreamingRPCDSL, testdata.ClientStreamingRPCServerInterfaceCode},
		{"client-streaming-rpc-with-payload", testdata.ClientStreamingRPCWithPayloadDSL, testdata.ClientStreamingRPCWithPayloadServerInterfaceCode},
		{"bidirectional-streaming-rpc", testdata.BidirectionalStreamingRPCDSL, testdata.BidirectionalStreamingRPCServerInterfaceCode},
		{"bidirectional-streaming-rpc-with-payload", testdata.BidirectionalStreamingRPCWithPayloadDSL, testdata.BidirectionalStreamingRPCWithPayloadServerInterfaceCode},
		{"bidirectional-streaming-rpc-with-errors", testdata.BidirectionalStreamingRPCWithErrorsDSL, testdata.BidirectionalStreamingRPCWithErrorsServerInterfaceCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			RunGRPCDSL(t, c.DSL)
			fs := ServerFiles("", expr.Root)
			if len(fs) != 2 {
				t.Fatalf("got %d files, expected two", len(fs))
			}
			sections := fs[0].Section("server-grpc-interface")
			if len(sections) == 0 {
				t.Fatalf("got zero sections, expected at least one")
			}
			code := codegen.SectionsCode(t, sections)
			if code != c.Code {
				t.Errorf("%s: got\n%s\ngot vs. expected:\n%s", c.Name, code, codegen.Diff(t, code, c.Code))
			}
		})
	}
}

func TestServerHandlerInit(t *testing.T) {
	cases := []struct {
		Name string
		DSL  func()
		Code string
	}{
		{"unary-rpcs", testdata.UnaryRPCsDSL, testdata.UnaryRPCsServerHandlerInitCode},
		{"unary-rpc-no-payload", testdata.UnaryRPCNoPayloadDSL, testdata.UnaryRPCNoPayloadServerHandlerInitCode},
		{"unary-rpc-no-result", testdata.UnaryRPCNoResultDSL, testdata.UnaryRPCNoResultServerHandlerInitCode},
		{"server-streaming-rpc", testdata.ServerStreamingRPCDSL, testdata.ServerStreamingRPCServerHandlerInitCode},
		{"client-streaming-rpc", testdata.ClientStreamingRPCDSL, testdata.ClientStreamingRPCServerHandlerInitCode},
		{"client-streaming-rpc-with-payload", testdata.ClientStreamingRPCWithPayloadDSL, testdata.ClientStreamingRPCWithPayloadServerHandlerInitCode},
		{"bidirectional-streaming-rpc", testdata.BidirectionalStreamingRPCDSL, testdata.BidirectionalStreamingRPCServerHandlerInitCode},
		{"bidirectional-streaming-rpc-with-payload", testdata.BidirectionalStreamingRPCWithPayloadDSL, testdata.BidirectionalStreamingRPCWithPayloadServerHandlerInitCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			RunGRPCDSL(t, c.DSL)
			fs := ServerFiles("", expr.Root)
			if len(fs) != 2 {
				t.Fatalf("got %d files, expected two", len(fs))
			}
			sections := fs[0].Section("grpc-handler-init")
			if len(sections) == 0 {
				t.Fatalf("got zero sections, expected at least one")
			}
			code := codegen.SectionsCode(t, sections)
			if code != c.Code {
				t.Errorf("%s: got\n%s\ngot vs. expected:\n%s", c.Name, code, codegen.Diff(t, code, c.Code))
			}
		})
	}
}

func TestRequestDecoder(t *testing.T) {
	cases := []struct {
		Name string
		DSL  func()
		Code string
	}{
		{"request-decoder-payload-user-type", testdata.MessageUserTypeWithNestedUserTypesDSL, testdata.PayloadUserTypeRequestDecoderCode},
		{"request-decoder-payload-array", testdata.UnaryRPCNoResultDSL, testdata.PayloadArrayRequestDecoderCode},
		{"request-decoder-payload-map", testdata.MessageMapDSL, testdata.PayloadMapRequestDecoderCode},
		{"request-decoder-payload-primitive", testdata.ServerStreamingRPCDSL, testdata.PayloadPrimitiveRequestDecoderCode},
		{"request-decoder-payload-primitive-with-streaming-payload", testdata.ClientStreamingRPCWithPayloadDSL, testdata.PayloadPrimitiveWithStreamingPayloadRequestDecoderCode},
		{"request-decoder-payload-user-type-with-streaming-payload", testdata.BidirectionalStreamingRPCWithPayloadDSL, testdata.PayloadUserTypeWithStreamingPayloadRequestDecoderCode},
		{"request-decoder-payload-with-metadata", testdata.MessageWithMetadataDSL, testdata.PayloadWithMetadataRequestDecoderCode},
		{"request-decoder-payload-with-validate", testdata.MessageWithValidateDSL, testdata.PayloadWithValidateRequestDecoderCode},
		{"request-decoder-payload-with-security-attributes", testdata.MessageWithSecurityAttrsDSL, testdata.PayloadWithSecurityAttrsRequestDecoderCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			RunGRPCDSL(t, c.DSL)
			fs := ServerFiles("", expr.Root)
			if len(fs) != 2 {
				t.Fatalf("got %d files, expected two", len(fs))
			}
			sections := fs[1].Section("request-decoder")
			if len(sections) == 0 {
				t.Fatalf("got zero sections, expected at least one")
			}
			code := codegen.SectionsCode(t, sections)
			if code != c.Code {
				t.Errorf("%s: got\n%s\ngot vs. expected:\n%s", c.Name, code, codegen.Diff(t, code, c.Code))
			}
		})
	}
}

func TestResponseEncoder(t *testing.T) {
	cases := []struct {
		Name string
		DSL  func()
		Code string
	}{
		{"response-encoder-empty-result", testdata.UnaryRPCNoResultDSL, testdata.EmptyResultResponseEncoderCode},
		{"response-encoder-result-with-views", testdata.MessageResultTypeWithViewsDSL, testdata.ResultWithViewsResponseEncoderCode},
		{"response-encoder-result-with-explicit-view", testdata.MessageResultTypeWithExplicitViewDSL, testdata.ResultWithExplicitViewResponseEncoderCode},
		{"response-encoder-result-array", testdata.MessageArrayDSL, testdata.ResultArrayResponseEncoderCode},
		{"response-encoder-result-primitive", testdata.UnaryRPCNoPayloadDSL, testdata.ResultPrimitiveResponseEncoderCode},
		{"response-encoder-result-with-metadata", testdata.MessageWithMetadataDSL, testdata.ResultWithMetadataResponseEncoderCode},
		{"response-encoder-result-with-validate", testdata.MessageWithValidateDSL, testdata.ResultWithValidateResponseEncoderCode},
		{"response-encoder-result-collection", testdata.MessageResultTypeCollectionDSL, testdata.ResultCollectionResponseEncoderCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			RunGRPCDSL(t, c.DSL)
			fs := ServerFiles("", expr.Root)
			if len(fs) != 2 {
				t.Fatalf("got %d files, expected two", len(fs))
			}
			sections := fs[1].Section("response-encoder")
			if len(sections) == 0 {
				t.Fatalf("got zero sections, expected at least one")
			}
			code := codegen.SectionsCode(t, sections)
			if code != c.Code {
				t.Errorf("%s: got\n%s\ngot vs. expected:\n%s", c.Name, code, codegen.Diff(t, code, c.Code))
			}
		})
	}
}
