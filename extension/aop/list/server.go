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

package list

import (
	"context"
	"sort"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/alibaba/ioc-golang/aop"
	"github.com/alibaba/ioc-golang/aop/common"
	"github.com/alibaba/ioc-golang/extension/aop/list/api/ioc_golang/aop/list"
)

type listServiceImpl struct {
	list.UnimplementedListServiceServer
	allInterfaceMetadataMap common.AllInterfaceMetadata
}

func (l *listServiceImpl) List(context.Context, *emptypb.Empty) (*list.ListServiceResponse, error) {
	structsMetadatas := make(metadataSorter, 0)
	for key, v := range l.allInterfaceMetadataMap {
		methods := make([]string, 0)
		for key := range v.MethodMetadata {
			methods = append(methods, key)
		}

		structsMetadatas = append(structsMetadatas, &list.ServiceMetadata{
			Methods:            methods,
			InterfaceName:      key,
			ImplementationName: key,
		})
	}
	sort.Sort(structsMetadatas)

	return &list.ListServiceResponse{
		ServiceMetadata: structsMetadatas,
	}, nil
}

func getListServiceImpl() *listServiceImpl {
	return &listServiceImpl{
		allInterfaceMetadataMap: aop.GetAllInterfaceMetadata(),
	}
}
