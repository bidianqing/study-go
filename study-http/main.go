package main

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path"
	path_filepath "path/filepath"
	"strings"
	"sync"
	"time"
)

func main() {
	dir := `C:\Users\bidianqing\Desktop`
	url := "https://github.com/dapr/dapr/releases/download/v1.13.1/daprd_linux_amd64.tar.gz"
	downloadFile(dir, url)
	fmt.Println("下载完成")

	untarExternalFile(dir+`\daprd_linux_amd64.tar.gz`, dir, "daprd")
	fmt.Println("解压完成")
}

func downloadFile(dir string, url string) (string, error) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]

	filepath := path.Join(dir, fileName)
	_, err := os.Stat(filepath)
	if os.IsExist(err) {
		return "", nil
	}
	client := http.Client{ //nolint:exhaustruct
		Timeout: 0,
		Transport: &http.Transport{ //nolint:exhaustruct
			Dial: (&net.Dialer{ //nolint:exhaustruct
				Timeout: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   15 * time.Second,
			ResponseHeaderTimeout: 15 * time.Second,
			Proxy:                 http.ProxyFromEnvironment,
		},
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return "", fmt.Errorf("version not found from url: %s", url)
	} else if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("download failed with %d", resp.StatusCode)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = copyWithTimeout(context.Background(), out, resp.Body)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

/*
!
See: https://github.com/microsoft/vscode-winsta11er/blob/4b42060da64aea6f47adebe1dd654980ed87a046/common/common.go
Copyright (c) Microsoft Corporation. All rights reserved. Licensed under the MIT License.
*/
func copyWithTimeout(ctx context.Context, dst io.Writer, src io.Reader) (int64, error) {
	// Every 5 seconds, ensure at least 200 bytes (40 bytes/second average) are read.
	interval := 5
	minCopyBytes := int64(200)
	prevWritten := int64(0)
	written := int64(0)

	done := make(chan error)
	mu := sync.Mutex{}
	t := time.NewTicker(time.Duration(interval) * time.Second)
	defer t.Stop()

	// Read the stream, 32KB at a time.
	go func() {
		var (
			writeErr, readErr     error
			writeBytes, readBytes int
			buf                   = make([]byte, 32<<10)
		)
		for {
			readBytes, readErr = src.Read(buf)
			if readBytes > 0 {
				// Write to disk and update the number of bytes written.
				writeBytes, writeErr = dst.Write(buf[0:readBytes])
				mu.Lock()
				written += int64(writeBytes)
				mu.Unlock()
				if writeErr != nil {
					done <- writeErr
					return
				}
			}
			if readErr != nil {
				// If error is EOF, means we read the entire file, so don't consider that as error.
				if !errors.Is(readErr, io.EOF) {
					done <- readErr
					return
				}

				// No error.
				done <- nil
				return
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return written, ctx.Err()
		case <-t.C:
			mu.Lock()
			if written < prevWritten+minCopyBytes {
				mu.Unlock()
				return written, fmt.Errorf("stream stalled: received %d bytes over the last %d seconds", written, interval)
			}
			prevWritten = written
			mu.Unlock()
		case e := <-done:
			return written, e
		}
	}
}

func untarExternalFile(filepath, dir, binaryFilePrefix string) (string, error) {
	reader, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("error open tar gz file %s: %w", filepath, err)
	}
	defer reader.Close()

	return untar(reader, dir, binaryFilePrefix)
}

func untar(reader io.Reader, targetDir string, binaryFilePrefix string) (string, error) {
	gzr, err := gzip.NewReader(reader)
	if err != nil {
		return "", err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	foundBinary := ""
	for {
		header, err := tr.Next()
		//nolint
		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		} else if header == nil {
			continue
		}

		// untar all files in archive.
		path, err := sanitizeExtractPath(targetDir, header.Name)
		if err != nil {
			return "", err
		}

		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return "", err
			}
			continue
		}

		f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
		if err != nil {
			return "", err
		}
		defer f.Close()

		// #nosec G110
		if _, err = io.Copy(f, tr); err != nil {
			return "", err
		}

		// If the found file is the binary that we want to find, save it and return later.
		if strings.HasSuffix(header.Name, binaryFilePrefix) {
			foundBinary = path
		}
	}
	return foundBinary, nil
}

// https://github.com/snyk/zip-slip-vulnerability, fixes gosec G305
func sanitizeExtractPath(destination string, filePath string) (string, error) {
	destpath := path_filepath.Join(destination, filePath)
	if !strings.HasPrefix(destpath, path_filepath.Clean(destination)+string(os.PathSeparator)) {
		return "", fmt.Errorf("%s: illegal file path", filePath)
	}
	return destpath, nil
}
