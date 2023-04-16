//
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
	"errors"

	cosignError "github.com/sigstore/cosign/v2/pkg/cosign"
)

func LookupExitCodeForError(err interface{ error }) int {
	var errNoMatchingSignatures *cosignError.ErrNoMatchingSignatures
	if errors.As(err, &errNoMatchingSignatures) {
		return NoMatchingSignature
	}

	var errImageTagNotFound *cosignError.ErrImageTagNotFound
	if errors.As(err, &errImageTagNotFound) {
		return NonExistentTag
	}

	var errNoSignaturesFound *cosignError.ErrNoSignaturesFound
	if errors.As(err, &errNoSignaturesFound) {
		return ImageWithoutSignature
	}

	// we want to return exit code = `1` at this point because there is
	// no valid exit code found for the error type passed.
	return 1
}
