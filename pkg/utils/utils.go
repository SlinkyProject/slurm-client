// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"encoding/json"
	"strconv"

	"golang.org/x/exp/constraints"
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

type number interface {
	constraints.Integer | constraints.Float
}

func NumberToString[T number](t T) string {
	return strconv.FormatInt(int64(t), 10)
}
