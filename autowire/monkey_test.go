/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package autowire

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockMonkeyFunction = func(i interface{}, s string) {

}

func TestRegisterMonkeyFunction(t *testing.T) {
	type args struct {
		f func(interface{}, string)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test register mock monkey function",
			args: args{
				f: mockMonkeyFunction,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterMonkeyFunction(tt.args.f)
			defer RegisterMonkeyFunction(nil)
			assert.Equal(t, fmt.Sprintf("%p", mockMonkeyFunction), fmt.Sprintf("%p", mf))
		})
	}
}
