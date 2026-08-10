// SPDX-FileCopyrightText: (C) 2024 Intel Corporation
// SPDX-License-Identifier: Apache 2.0

package serviceinfo

import (
	"context"
	"io"
	"unicode"
)

// DevmodOwnerModule receives the device devmod module list.
type DevmodOwnerModule struct {
	Devmod  Devmod
	Modules []string
}

// HandleInfo implements OwnerModule.
func (d *DevmodOwnerModule) HandleInfo(_ context.Context, messageName string, messageBody io.Reader) error {
	if messageName != "modules" {
		_, err := io.Copy(io.Discard, messageBody)
		return err
	}
	body, err := io.ReadAll(messageBody)
	if err != nil {
		return err
	}
	for i := 0; i < len(body); i++ {
		if !isModuleNameByte(body[i]) {
			continue
		}
		j := i + 1
		for j < len(body) && isModuleNameByte(body[j]) {
			j++
		}
		if j-i > 1 {
			d.Modules = append(d.Modules, string(body[i:j]))
		}
		i = j
	}
	return nil
}

// ProduceInfo implements OwnerModule.
func (d *DevmodOwnerModule) ProduceInfo(context.Context, *Producer) (bool, bool, error) {
	return false, true, nil
}

func isModuleNameByte(b byte) bool {
	r := rune(b)
	return unicode.IsLetter(r) || unicode.IsDigit(r) || b == '.' || b == '-' || b == '_'
}
