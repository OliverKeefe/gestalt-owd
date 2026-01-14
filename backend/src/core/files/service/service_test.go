package files

// TODO: Finish test, verify upload works correctly.
//func UploadTest_TestFileUpload(t *testing.T) {
//	fileContents := []byte("This is a test file")
//	fileName := "testfile.txt"
//
//	var requestBody bytes.Buffer
//	writer := multipart.NewWriter(&requestBody)
//
//	part, err := writer.CreateFormFile("file", fileName)
//	if err != nil {
//		t.Fatalf("failed to create form file: %v", err)
//	}
//
//	Writer.Close()
//
//	req := httptest.NewRequest(rest.MethodPut, "files/upload", &requestBody)
//}
