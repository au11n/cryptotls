// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build boringcrypto

// This package only exists with GOEXPERIMENT=boringcrypto.
// It provides the actual syso file.
//
// FORK: .syso-бинарники BoringCrypto (~4.5 МБ) не включены - форк собирается
// без GOEXPERIMENT=boringcrypto. Пакет оставлен, чтобы импорт из boring.go
// резолвился при анализе всех build-конфигураций (go mod tidy, go vet --tags).
package syso
