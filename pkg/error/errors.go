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

package errors

import (
	"fmt"
)

var (
	// ErrNoMatchingSignatures is the error returned when there are no matching
	// signatures during verification.
	ErrNoMatchingSignatures = &CosignError{"no matching signatures", UnmatchingSignature}

	// ErrNoMatchingAttestations is the error returned when there are no
	// matching attestations during verification.
	// TODO: don't forget to implement exit code for no matching attestation
	ErrNoMatchingAttestations = &CosignError{"no matching attestations", 1}
)

// VerificationError is the type of Go error that is used by cosign to surface
// errors actually related to verification (vs. transient, misconfiguration,
// transport, or authentication related issues).
type CosignError struct {
	Message string
	Code    int
}

// NewVerificationError constructs a new VerificationError in a manner similar
// to fmt.Errorf
func NewVerificationError(msg string, args ...interface{}) error {
	return &CosignError{
		Message: fmt.Sprintf(msg, args...),
		Code:    12,
	}
}

func NewError(cosignError CosignError) error {
	return &CosignError{
		Message: cosignError.Message,
		Code:    cosignError.Code,
	}
	// return &VerificationError{
	// 	message: fmt.Sprintf(msg, args...),
	// 	code:    12,
	// }
}

// Assert that we implement error at build time.
var _ error = (*CosignError)(nil)

// Error implements error
func (ve *CosignError) Error() string {
	return ve.Message
}

func (e *CosignError) ExitCode() int {
	return e.Code
}
