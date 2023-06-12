package dfs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListAccessRules invokes the dfs.ListAccessRules API synchronously
func (client *Client) ListAccessRules(request *ListAccessRulesRequest) (response *ListAccessRulesResponse, err error) {
	response = CreateListAccessRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ListAccessRulesWithChan invokes the dfs.ListAccessRules API asynchronously
func (client *Client) ListAccessRulesWithChan(request *ListAccessRulesRequest) (<-chan *ListAccessRulesResponse, <-chan error) {
	responseChan := make(chan *ListAccessRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAccessRules(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListAccessRulesWithCallback invokes the dfs.ListAccessRules API asynchronously
func (client *Client) ListAccessRulesWithCallback(request *ListAccessRulesRequest, callback func(response *ListAccessRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAccessRulesResponse
		var err error
		defer close(result)
		response, err = client.ListAccessRules(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListAccessRulesRequest is the request struct for api ListAccessRules
type ListAccessRulesRequest struct {
	*requests.RpcRequest
	AccessGroupId string           `position:"Query" name:"AccessGroupId"`
	InputRegionId string           `position:"Query" name:"InputRegionId"`
	Limit         requests.Integer `position:"Query" name:"Limit"`
	OrderBy       string           `position:"Query" name:"OrderBy"`
	StartOffset   requests.Integer `position:"Query" name:"StartOffset"`
	OrderType     string           `position:"Query" name:"OrderType"`
}

// ListAccessRulesResponse is the response struct for api ListAccessRules
type ListAccessRulesResponse struct {
	*responses.BaseResponse
	TotalCount  int          `json:"TotalCount" xml:"TotalCount"`
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	AccessRules []AccessRule `json:"AccessRules" xml:"AccessRules"`
}

// CreateListAccessRulesRequest creates a request to invoke ListAccessRules API
func CreateListAccessRulesRequest() (request *ListAccessRulesRequest) {
	request = &ListAccessRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DFS", "2018-06-20", "ListAccessRules", "alidfs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAccessRulesResponse creates a response to parse from ListAccessRules response
func CreateListAccessRulesResponse() (response *ListAccessRulesResponse) {
	response = &ListAccessRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
