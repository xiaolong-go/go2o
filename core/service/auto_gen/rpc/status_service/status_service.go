// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package status_service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"go2o/core/service/auto_gen/rpc/ttype"
	"reflect"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

var _ = ttype.GoUnusedProtection__

type StatusService interface { //状态服务

	// 尝试连接
	Ping(ctx context.Context) (r string, err error)
}

//状态服务
type StatusServiceClient struct {
	c thrift.TClient
}

func NewStatusServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *StatusServiceClient {
	return &StatusServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewStatusServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *StatusServiceClient {
	return &StatusServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewStatusServiceClient(c thrift.TClient) *StatusServiceClient {
	return &StatusServiceClient{
		c: c,
	}
}

func (p *StatusServiceClient) Client_() thrift.TClient {
	return p.c
}

// 尝试连接
func (p *StatusServiceClient) Ping(ctx context.Context) (r string, err error) {
	var _args0 StatusServicePingArgs
	var _result1 StatusServicePingResult
	if err = p.Client_().Call(ctx, "Ping", &_args0, &_result1); err != nil {
		return
	}
	return _result1.GetSuccess(), nil
}

type StatusServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      StatusService
}

func (p *StatusServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *StatusServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *StatusServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewStatusServiceProcessor(handler StatusService) *StatusServiceProcessor {

	self2 := &StatusServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["Ping"] = &statusServiceProcessorPing{handler: handler}
	return self2
}

func (p *StatusServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x3.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x3

}

type statusServiceProcessorPing struct {
	handler StatusService
}

func (p *statusServiceProcessorPing) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := StatusServicePingArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Ping", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	result := StatusServicePingResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.Ping(ctx); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Ping: "+err2.Error())
		oprot.WriteMessageBegin("Ping", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("Ping", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type StatusServicePingArgs struct {
}

func NewStatusServicePingArgs() *StatusServicePingArgs {
	return &StatusServicePingArgs{}
}

func (p *StatusServicePingArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *StatusServicePingArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Ping_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *StatusServicePingArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StatusServicePingArgs(%+v)", *p)
}

// Attributes:
//  - Success
type StatusServicePingResult struct {
	Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewStatusServicePingResult() *StatusServicePingResult {
	return &StatusServicePingResult{}
}

var StatusServicePingResult_Success_DEFAULT string

func (p *StatusServicePingResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return StatusServicePingResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *StatusServicePingResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *StatusServicePingResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField0(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *StatusServicePingResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *StatusServicePingResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Ping_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *StatusServicePingResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *StatusServicePingResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StatusServicePingResult(%+v)", *p)
}