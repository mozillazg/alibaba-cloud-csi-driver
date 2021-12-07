package nas

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

// ListRecycledDirectoriesAndFiles invokes the nas.ListRecycledDirectoriesAndFiles API synchronously
func (client *Client) ListRecycledDirectoriesAndFiles(request *ListRecycledDirectoriesAndFilesRequest) (response *ListRecycledDirectoriesAndFilesResponse, err error) {
	response = CreateListRecycledDirectoriesAndFilesResponse()
	err = client.DoAction(request, response)
	return
}

// ListRecycledDirectoriesAndFilesWithChan invokes the nas.ListRecycledDirectoriesAndFiles API asynchronously
func (client *Client) ListRecycledDirectoriesAndFilesWithChan(request *ListRecycledDirectoriesAndFilesRequest) (<-chan *ListRecycledDirectoriesAndFilesResponse, <-chan error) {
	responseChan := make(chan *ListRecycledDirectoriesAndFilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRecycledDirectoriesAndFiles(request)
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

// ListRecycledDirectoriesAndFilesWithCallback invokes the nas.ListRecycledDirectoriesAndFiles API asynchronously
func (client *Client) ListRecycledDirectoriesAndFilesWithCallback(request *ListRecycledDirectoriesAndFilesRequest, callback func(response *ListRecycledDirectoriesAndFilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRecycledDirectoriesAndFilesResponse
		var err error
		defer close(result)
		response, err = client.ListRecycledDirectoriesAndFiles(request)
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

// ListRecycledDirectoriesAndFilesRequest is the request struct for api ListRecycledDirectoriesAndFiles
type ListRecycledDirectoriesAndFilesRequest struct {
	*requests.RpcRequest
	NextToken    string           `position:"Query" name:"NextToken"`
	FileSystemId string           `position:"Query" name:"FileSystemId"`
	FileId       string           `position:"Query" name:"FileId"`
	MaxResults   requests.Integer `position:"Query" name:"MaxResults"`
}

// ListRecycledDirectoriesAndFilesResponse is the response struct for api ListRecycledDirectoriesAndFiles
type ListRecycledDirectoriesAndFilesResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	NextToken string  `json:"NextToken" xml:"NextToken"`
	Entries   []Entry `json:"Entries" xml:"Entries"`
}

// CreateListRecycledDirectoriesAndFilesRequest creates a request to invoke ListRecycledDirectoriesAndFiles API
func CreateListRecycledDirectoriesAndFilesRequest() (request *ListRecycledDirectoriesAndFilesRequest) {
	request = &ListRecycledDirectoriesAndFilesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "ListRecycledDirectoriesAndFiles", "", "")
	request.Method = requests.GET
	return
}

// CreateListRecycledDirectoriesAndFilesResponse creates a response to parse from ListRecycledDirectoriesAndFiles response
func CreateListRecycledDirectoriesAndFilesResponse() (response *ListRecycledDirectoriesAndFilesResponse) {
	response = &ListRecycledDirectoriesAndFilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}