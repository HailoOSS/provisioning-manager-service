// Code generated by protoc-gen-go.
// source: github.com/HailoOSS/provisioning-manager-service/proto/getprovisioner/getprovisioner.proto
// DO NOT EDIT!

/*
Package com_HailoOSS_kernel_provisioningmanager_getprovisioner is a generated protocol buffer package.

It is generated from these files:
	github.com/HailoOSS/provisioning-manager-service/proto/getprovisioner/getprovisioner.proto

It has these top-level messages:
	Request
	Response
*/
package com_HailoOSS_kernel_provisioningmanager_getprovisioner

import proto "github.com/HailoOSS/protobuf/proto"
import json "encoding/json"
import math "math"
import com_HailoOSS_kernel_provisioningmanager "github.com/HailoOSS/provisioning-manager-service/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	Hostname         *string `protobuf:"bytes,1,req,name=hostname" json:"hostname,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

type Response struct {
	Provisioner      *com_HailoOSS_kernel_provisioningmanager.Provisioner `protobuf:"bytes,1,req,name=provisioner" json:"provisioner,omitempty"`
	XXX_unrecognized []byte                                               `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetProvisioner() *com_HailoOSS_kernel_provisioningmanager.Provisioner {
	if m != nil {
		return m.Provisioner
	}
	return nil
}

func init() {
}
