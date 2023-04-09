// Copyright 2022 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cosign

import (
	"errors"
	"fmt"
)

var (
	ErrNoMatchingAttestations = errors.New("no matching attestations")
	ErrNoMatchingSignatures   = errors.New("no matching signatures")
	ErrImageTagNotFound       = errors.New("image tag not found")
	ErrNoSignaturesFound      = errors.New("no signatures found for image")
)

// VerificationError is the type of Go error that is used by cosign to surface
// errors actually related to verification (vs. transient, misconfiguration,
// transport, or authentication related issues).
type VerificationError struct {
	message string
	err     error
}

// NewVerificationError constructs a new VerificationError in a manner similar
// to fmt.Errorf
func NewVerificationError(msg string, args ...interface{}) error {
	return &VerificationError{
		message: fmt.Sprintf(msg, args...),
	}
}

// this is temporary for now whilst we build the exit codes for the different scenarios
func NewVerificationErrorJustMessage(msg string) error {
	return &VerificationError{
		message: msg,
		err:     errors.New(msg),
	}
}

func NewVerificationErrorWrapped(err error, msg string) error {
	return &VerificationError{
		message: msg,
		err:     err,
	}
}

// Assert that we implement error at build time.
var _ error = (*VerificationError)(nil)

// Error implements error
func (ve *VerificationError) Message() string {
	return ve.message
}

// Error implements error
func (ve *VerificationError) Error() string {
	if ve.err != nil {
		return ve.err.Error()
	}
	return ve.message
}

func (e *VerificationError) Unwrap() error { return e.err }
