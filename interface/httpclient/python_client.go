package httpclient

import (
	"Backend-Warehouse/interface/dto"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"time"
)

type PythonClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewPythonClient(baseURL string) *PythonClient {
	return &PythonClient{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *PythonClient) UploadAvatar(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*dto.UploadAvatarResponse, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreatePart(textproto.MIMEHeader{
		"Content-Disposition": []string{
			fmt.Sprintf(`form-data; name="file"; filename="%s"`, header.Filename),
		},
		"Content-Type": []string{header.Header.Get("Content-Type")},
	})
	if err != nil {
		return nil, err
	}
	io.Copy(part, file)
	writer.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.baseURL+"/upload-avatar",
		&buf,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("python service error: %d", resp.StatusCode)
	}

	var result dto.UploadAvatarResponse
	json.NewDecoder(resp.Body).Decode(&result)
	return &result, nil
}

func (c *PythonClient) GenerateSignature(ctx context.Context, file multipart.File, header *multipart.FileHeader, qrData dto.QrDataRequest) (*dto.SignatureResponse, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	log.Printf("header:%v", header)
	log.Printf("qrdata:%s", qrData)

	part, err := writer.CreatePart(textproto.MIMEHeader{
		"Content-Disposition": []string{
			fmt.Sprintf(`form-data; name="signature"; filename="%s"`, header.Filename),
		},
		"Content-Type": []string{header.Header.Get("Content-Type")},
	})
	if err != nil {
		return nil, err
	}
	io.Copy(part, file)

	log.Printf("part: %s", part)
	log.Printf("file:%s", file)

	qrDataJSON, _ := json.Marshal(qrData)
	writer.WriteField("qr_data", string(qrDataJSON))
	writer.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.baseURL+"/generate-signature",
		&buf,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("python service error: %d", resp.StatusCode)
	}

	var result dto.SignatureResponse
	json.NewDecoder(resp.Body).Decode(&result)
	return &result, nil
}

func (c *PythonClient) DeleteFiles(ctx context.Context, urls ...string) error {
	// call Python endpoint untuk delete file di MinIO
	payload, _ := json.Marshal(map[string][]string{"urls": urls})

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete,
		c.baseURL+"/delete-files",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}