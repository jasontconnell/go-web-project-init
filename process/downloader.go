package process

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func DownloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error in do %s %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("not 200 status code %s %d %s", url, resp.StatusCode, resp.Status)
	}

	content := []byte{}

	buf := bytes.NewBuffer(content)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read %s %w", url, err)
	}

	return buf.Bytes(), nil
}
