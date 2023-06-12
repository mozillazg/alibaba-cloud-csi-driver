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

// ListFileSystems invokes the dfs.ListFileSystems API synchronously
func (client *Client) ListFileSystems(request *ListFileSystemsRequest) (response *ListFileSystemsResponse, err error) {
	response = CreateListFileSystemsResponse()
	err = client.DoAction(request, response)
	return
}

// ListFileSystemsWithChan invokes the dfs.ListFileSystems API asynchronously
func (client *Client) ListFileSystemsWithChan(request *ListFileSystemsRequest) (<-chan *ListFileSystemsResponse, <-chan error) {
	responseChan := make(chan *ListFileSystemsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFileSystems(request)
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

// ListFileSystemsWithCallback invokes the dfs.ListFileSystems API asynchronously
func (client *Client) ListFileSystemsWithCallback(request *ListFileSystemsRequest, callback func(response *ListFileSystemsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFileSystemsResponse
		var err error
		defer close(result)
		response, err = client.ListFileSystems(request)
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

// ListFileSystemsRequest is the request struct for api ListFileSystems
type ListFileSystemsRequest struct {
	*requests.RpcRequest
	InputRegionId string           `position:"Query" name:"InputRegionId"`
	Limit         requests.Integer `position:"Query" name:"Limit"`
	OrderBy       string           `position:"Query" name:"OrderBy"`
	StartOffset   requests.Integer `position:"Query" name:"StartOffset"`
	OrderType     string           `position:"Query" name:"OrderType"`
}

// ListFileSystemsResponse is the response struct for api ListFileSystems
type ListFileSystemsResponse struct {
	*responses.BaseResponse
	TotalCount  int          `json:"TotalCount" xml:"TotalCount"`
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	FileSystems []FileSystem `json:"FileSystems" xml:"FileSystems"`
}

// CreateListFileSystemsRequest creates a request to invoke ListFileSystems API
func CreateListFileSystemsRequest() (request *ListFileSystemsRequest) {
	request = &ListFileSystemsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DFS", "2018-06-20", "ListFileSystems", "alidfs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListFileSystemsResponse creates a response to parse from ListFileSystems response
func CreateListFileSystemsResponse() (response *ListFileSystemsResponse) {
	response = &ListFileSystemsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
