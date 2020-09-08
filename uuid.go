// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package uuid

import (
	"database/sql/driver"
	"github.com/google/uuid"
)

type UUID uuid.UUID

func (u *UUID) Scan(src interface{}) error {
	t := (*uuid.UUID)(u)
	return t.Scan(src)
}

func (u UUID) Value() (driver.Value, error) {
	t := [16]byte(u)
	return t[:], nil
}

func (u UUID) String() string {
	return uuid.UUID(u).String()
}

func New() UUID {
	return UUID(uuid.New())
}

func MustParse(str string) UUID {
	return UUID(uuid.MustParse(str))
}

func Parse(str string) (UUID, error) {
	u, err := uuid.Parse(str)
	return UUID(u), err
}

func (u UUID) MarshalText() ([]byte, error) {
	return uuid.UUID(u).MarshalText()
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (u *UUID) UnmarshalText(data []byte) error {
	t := (*uuid.UUID)(u)
	return t.UnmarshalText(data)
}
