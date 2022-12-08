// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || amd64p32 || arm || arm64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64
// +build 386 amd64 amd64p32 arm arm64 mips64le mips64p32le mipsle ppc64le riscv64

package tracer

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// LoadIterTCPv4 returns the embedded CollectionSpec for IterTCPv4.
func LoadIterTCPv4() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_IterTCPv4Bytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load IterTCPv4: %w", err)
	}

	return spec, err
}

// LoadIterTCPv4Objects loads IterTCPv4 and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*IterTCPv4Objects
//	*IterTCPv4Programs
//	*IterTCPv4Maps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadIterTCPv4Objects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadIterTCPv4()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// IterTCPv4Specs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type IterTCPv4Specs struct {
	IterTCPv4ProgramSpecs
	IterTCPv4MapSpecs
}

// IterTCPv4Specs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type IterTCPv4ProgramSpecs struct {
	IgSnapTcp4 *ebpf.ProgramSpec `ebpf:"ig_snap_tcp4"`
}

// IterTCPv4MapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type IterTCPv4MapSpecs struct {
}

// IterTCPv4Objects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadIterTCPv4Objects or ebpf.CollectionSpec.LoadAndAssign.
type IterTCPv4Objects struct {
	IterTCPv4Programs
	IterTCPv4Maps
}

func (o *IterTCPv4Objects) Close() error {
	return _IterTCPv4Close(
		&o.IterTCPv4Programs,
		&o.IterTCPv4Maps,
	)
}

// IterTCPv4Maps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadIterTCPv4Objects or ebpf.CollectionSpec.LoadAndAssign.
type IterTCPv4Maps struct {
}

func (m *IterTCPv4Maps) Close() error {
	return _IterTCPv4Close()
}

// IterTCPv4Programs contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadIterTCPv4Objects or ebpf.CollectionSpec.LoadAndAssign.
type IterTCPv4Programs struct {
	IgSnapTcp4 *ebpf.Program `ebpf:"ig_snap_tcp4"`
}

func (p *IterTCPv4Programs) Close() error {
	return _IterTCPv4Close(
		p.IgSnapTcp4,
	)
}

func _IterTCPv4Close(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed itertcpv4_bpfel.o
var _IterTCPv4Bytes []byte
