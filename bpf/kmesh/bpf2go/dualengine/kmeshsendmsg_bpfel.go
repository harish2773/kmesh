// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64

package dualengine

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type KmeshSendmsgBpfSockTuple struct {
	Ipv4 struct {
		Saddr uint32
		Daddr uint32
		Sport uint16
		Dport uint16
	}
	_ [24]byte
}

type KmeshSendmsgBuf struct{ Data [40]int8 }

// LoadKmeshSendmsg returns the embedded CollectionSpec for KmeshSendmsg.
func LoadKmeshSendmsg() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_KmeshSendmsgBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load KmeshSendmsg: %w", err)
	}

	return spec, err
}

// LoadKmeshSendmsgObjects loads KmeshSendmsg and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*KmeshSendmsgObjects
//	*KmeshSendmsgPrograms
//	*KmeshSendmsgMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadKmeshSendmsgObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadKmeshSendmsg()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// KmeshSendmsgSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshSendmsgSpecs struct {
	KmeshSendmsgProgramSpecs
	KmeshSendmsgMapSpecs
	KmeshSendmsgVariableSpecs
}

// KmeshSendmsgProgramSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshSendmsgProgramSpecs struct {
	SendmsgProg *ebpf.ProgramSpec `ebpf:"sendmsg_prog"`
}

// KmeshSendmsgMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshSendmsgMapSpecs struct {
	KmLogEvent *ebpf.MapSpec `ebpf:"km_log_event"`
	KmOrigDst  *ebpf.MapSpec `ebpf:"km_orig_dst"`
	KmTmpbuf   *ebpf.MapSpec `ebpf:"km_tmpbuf"`
}

// KmeshSendmsgVariableSpecs contains global variables before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshSendmsgVariableSpecs struct {
	BpfLogLevel *ebpf.VariableSpec `ebpf:"bpf_log_level"`
}

// KmeshSendmsgObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshSendmsgObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshSendmsgObjects struct {
	KmeshSendmsgPrograms
	KmeshSendmsgMaps
	KmeshSendmsgVariables
}

func (o *KmeshSendmsgObjects) Close() error {
	return _KmeshSendmsgClose(
		&o.KmeshSendmsgPrograms,
		&o.KmeshSendmsgMaps,
	)
}

// KmeshSendmsgMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshSendmsgObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshSendmsgMaps struct {
	KmLogEvent *ebpf.Map `ebpf:"km_log_event"`
	KmOrigDst  *ebpf.Map `ebpf:"km_orig_dst"`
	KmTmpbuf   *ebpf.Map `ebpf:"km_tmpbuf"`
}

func (m *KmeshSendmsgMaps) Close() error {
	return _KmeshSendmsgClose(
		m.KmLogEvent,
		m.KmOrigDst,
		m.KmTmpbuf,
	)
}

// KmeshSendmsgVariables contains all global variables after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshSendmsgObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshSendmsgVariables struct {
	BpfLogLevel *ebpf.Variable `ebpf:"bpf_log_level"`
}

// KmeshSendmsgPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshSendmsgObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshSendmsgPrograms struct {
	SendmsgProg *ebpf.Program `ebpf:"sendmsg_prog"`
}

func (p *KmeshSendmsgPrograms) Close() error {
	return _KmeshSendmsgClose(
		p.SendmsgProg,
	)
}

func _KmeshSendmsgClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed kmeshsendmsg_bpfel.o
var _KmeshSendmsgBytes []byte
