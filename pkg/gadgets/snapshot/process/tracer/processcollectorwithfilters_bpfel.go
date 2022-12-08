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

type ProcessCollectorWithFiltersContainer struct {
	ContainerId [256]int8
	Namespace   [256]int8
	Pod         [256]int8
	Container   [256]int8
}

// LoadProcessCollectorWithFilters returns the embedded CollectionSpec for ProcessCollectorWithFilters.
func LoadProcessCollectorWithFilters() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_ProcessCollectorWithFiltersBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load ProcessCollectorWithFilters: %w", err)
	}

	return spec, err
}

// LoadProcessCollectorWithFiltersObjects loads ProcessCollectorWithFilters and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*ProcessCollectorWithFiltersObjects
//	*ProcessCollectorWithFiltersPrograms
//	*ProcessCollectorWithFiltersMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadProcessCollectorWithFiltersObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadProcessCollectorWithFilters()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// ProcessCollectorWithFiltersSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type ProcessCollectorWithFiltersSpecs struct {
	ProcessCollectorWithFiltersProgramSpecs
	ProcessCollectorWithFiltersMapSpecs
}

// ProcessCollectorWithFiltersSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type ProcessCollectorWithFiltersProgramSpecs struct {
	IgSnapProc *ebpf.ProgramSpec `ebpf:"ig_snap_proc"`
}

// ProcessCollectorWithFiltersMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type ProcessCollectorWithFiltersMapSpecs struct {
	Containers    *ebpf.MapSpec `ebpf:"containers"`
	MountNsFilter *ebpf.MapSpec `ebpf:"mount_ns_filter"`
}

// ProcessCollectorWithFiltersObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadProcessCollectorWithFiltersObjects or ebpf.CollectionSpec.LoadAndAssign.
type ProcessCollectorWithFiltersObjects struct {
	ProcessCollectorWithFiltersPrograms
	ProcessCollectorWithFiltersMaps
}

func (o *ProcessCollectorWithFiltersObjects) Close() error {
	return _ProcessCollectorWithFiltersClose(
		&o.ProcessCollectorWithFiltersPrograms,
		&o.ProcessCollectorWithFiltersMaps,
	)
}

// ProcessCollectorWithFiltersMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadProcessCollectorWithFiltersObjects or ebpf.CollectionSpec.LoadAndAssign.
type ProcessCollectorWithFiltersMaps struct {
	Containers    *ebpf.Map `ebpf:"containers"`
	MountNsFilter *ebpf.Map `ebpf:"mount_ns_filter"`
}

func (m *ProcessCollectorWithFiltersMaps) Close() error {
	return _ProcessCollectorWithFiltersClose(
		m.Containers,
		m.MountNsFilter,
	)
}

// ProcessCollectorWithFiltersPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadProcessCollectorWithFiltersObjects or ebpf.CollectionSpec.LoadAndAssign.
type ProcessCollectorWithFiltersPrograms struct {
	IgSnapProc *ebpf.Program `ebpf:"ig_snap_proc"`
}

func (p *ProcessCollectorWithFiltersPrograms) Close() error {
	return _ProcessCollectorWithFiltersClose(
		p.IgSnapProc,
	)
}

func _ProcessCollectorWithFiltersClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed processcollectorwithfilters_bpfel.o
var _ProcessCollectorWithFiltersBytes []byte
