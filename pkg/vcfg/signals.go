package vcfg

import (
	"fmt"
	"strings"
	"syscall"
)

/**
 * SPDX-License-Identifier: Apache-2.0
 * Copyright 2020 vorteil.io Pty Ltd
 */

//TerminateSignal : The Signal to send to a program on termination
//	Additional information can be found @ https://support.vorteil.io/docs/VCFG-Reference/program/terminate
type TerminateSignal string

// DefaultTerminateSignal : Default Terminate Signal to be used on programs
const DefaultTerminateSignal TerminateSignal = "SIGTERM"

// TerminateSignals : Supported Signals
var TerminateSignals = map[TerminateSignal]syscall.Signal{
	"SIGINT":  syscall.Signal(0x2),  // Term    Interrupt from keyboard
	"SIGKILL": syscall.Signal(0x9),  // Term    Kill signal
	"SIGQUIT": syscall.Signal(0x3),  // Core    Quit from keyboard
	"SIGPWR":  syscall.Signal(0x1e), // Term	Power Failure (System V)
	"SIGSTOP": syscall.Signal(0x13), // Stop    Stop process
	"SIGTERM": syscall.Signal(0xf),  // Term    Termination signal
	"SIGUSR1": syscall.Signal(0xa),  // User-defined signal 1
	"SIGUSR2": syscall.Signal(0xc),  // User-defined signal 2
}

// Validate : Check if TerminateSignal is a supported signal
func (tSig *TerminateSignal) Validate() (err error) {
	validSignals := make([]string, 0)

	if _, ok := TerminateSignals[*tSig]; ok {
		return // Valid Signal
	}

	for strSig := range TerminateSignals {
		validSignals = append(validSignals, string(strSig))
	}

	return fmt.Errorf("terminate signal '%s' is not supported. Supported Signals: %s", *tSig, strings.Join(validSignals, ", "))
}

// Signal : Return syscall Signal
func (tSig *TerminateSignal) Signal() (syscall.Signal, error) {
	if err := tSig.Validate(); err != nil {
		return syscall.Signal(0), err
	}

	return TerminateSignals[*tSig], nil
}
