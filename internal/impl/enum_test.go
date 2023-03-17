// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl_test

import (
	"testing"

	"github.com/bitquery/protobuf-sql/reflect/protoreflect"

	testpb "github.com/bitquery/protobuf-sql/internal/testprotos/test"
)

func TestEnum(t *testing.T) {
	et := testpb.ForeignEnum_FOREIGN_FOO.Type()
	if got, want := et.New(protoreflect.EnumNumber(testpb.ForeignEnum_FOREIGN_FOO)), protoreflect.Enum(testpb.ForeignEnum_FOREIGN_FOO); got != want {
		t.Errorf("testpb.ForeignEnum_FOREIGN_FOO.Type().New() = %[1]T(%[1]v), want %[2]T(%[2]v)", got, want)
	}
}
