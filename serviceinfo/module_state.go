// SPDX-FileCopyrightText: (C) 2024 Intel Corporation
// SPDX-License-Identifier: Apache 2.0

package serviceinfo

import (
	"context"
	"io"
	"iter"
)

// IteratorModuleStateMachine adapts an owner module iterator to a module state machine.
type IteratorModuleStateMachine struct {
	Modules iter.Seq2[string, OwnerModule]

	currentName string
	current     OwnerModule
	next        func() (string, OwnerModule, bool)
	stop        func()
}

// Module returns the current module.
func (m *IteratorModuleStateMachine) Module(context.Context) (string, OwnerModule, error) {
	if m.current != nil {
		return m.currentName, m.current, nil
	}
	if ok, err := m.NextModule(context.Background()); err != nil {
		return "", nil, err
	} else if !ok {
		return "", nil, io.EOF
	}
	return m.currentName, m.current, nil
}

// NextModule advances to the next module.
func (m *IteratorModuleStateMachine) NextModule(context.Context) (bool, error) {
	if m.next == nil {
		m.next, m.stop = iter.Pull2(m.Modules)
	}
	name, mod, ok := m.next()
	if !ok {
		m.currentName = ""
		m.current = nil
		return false, nil
	}
	m.currentName = name
	m.current = mod
	return true, nil
}

// CleanupModules releases iterator state.
func (m *IteratorModuleStateMachine) CleanupModules(context.Context) {
	if m.stop != nil {
		m.stop()
	}
}
