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

package aop

import (
	"reflect"

	"github.com/petermattis/goid"
)

type InvocationContext struct {
	ProxyServicePtr interface{}
	SDID            string
	MethodName      string
	MethodFullName  string
	Params          []reflect.Value
	ReturnValues    []reflect.Value
	GrID            int64
}

func (c *InvocationContext) SetReturnValues(returnValues []reflect.Value) {
	c.ReturnValues = returnValues
}

func NewInvocationContext(proxyServicePtr interface{}, sdid, methodName, methodFullName string, params []reflect.Value) *InvocationContext {
	return &InvocationContext{
		ProxyServicePtr: proxyServicePtr,
		SDID:            sdid,
		MethodName:      methodName,
		Params:          params,
		GrID:            goid.Get(),
		MethodFullName:  methodFullName,
	}
}
