// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"encoding/json"
)

func Remarshal(in any, out any) error {
	data, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, out)
}

func RemarshalOrDie(in any, out any) {
	if err := Remarshal(in, out); err != nil {
		panic(err)
	}
}
