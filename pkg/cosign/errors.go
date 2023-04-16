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

// VerificationFailure is temporary until the error types for the case
// where this has been thrown has been defined.
type VerificationFailure struct {
	err error
}

func (e *VerificationFailure) Error() string {
	return e.err.Error()
}

type ErrNoMatchingSignatures struct {
	err error
}

func (e *ErrNoMatchingSignatures) Error() string {
	return e.err.Error()
}

type ErrImageTagNotFound struct {
	err error
}

func (e *ErrImageTagNotFound) Error() string {
	return e.err.Error()
}

type ErrNoSignaturesFound struct {
	err error
}

func (e *ErrNoSignaturesFound) Error() string {
	return e.err.Error()
}

type ErrNoMatchingAttestations struct {
	err error
}

func (e *ErrNoMatchingAttestations) Error() string {
	return e.err.Error()
}

func ThrowError(err interface{ error }) error {
	return err
}
