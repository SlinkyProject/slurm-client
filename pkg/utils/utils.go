// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"encoding/json"
	"errors"
	"regexp"
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

func ParseNodeName(nodeConf string) (string, error) {
	re := regexp.MustCompile(`(?i)NodeName=([^\s]+)`)
	matches := re.FindStringSubmatch(nodeConf)
	if len(matches) < 2 {
		return "", errors.New("NodeName not found in node configuration string")
	}
	return matches[1], nil
}
