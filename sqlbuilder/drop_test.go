// Copyright 2018 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package sqlbuilder

var (
	_ SQLer  = &UpdateStmt{}
	_ execer = &UpdateStmt{}
)
