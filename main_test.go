package storage

import (
	"bytes"
	"context"
	"io"
	"testing"
)

func TestStorageCos(t *testing.T) {
	domain := ""
	tLocal := New("cos", map[string]string{
		"bucket":    "",
		"secretId":  "",
		"secretKey": "",
		"region":    "",
		"domain":    domain,
	})
	var err error
	content := "test"

	err = tLocal.Write(context.Background(), "test.txt", content, map[string]any{})
	if err != nil {
		t.Error("local write failed:" + err.Error())
	}

	str, err := tLocal.Read(context.Background(), "test.txt")
	if err != nil {
		t.Error("local read failed:" + err.Error())
	}
	if str != content {
		t.Error("local read failed: inconsistency of data")
	}

	publicUrl, err := tLocal.PublicUrl(context.Background(), "test.txt")
	if err != nil {
		t.Error("publicUrl failed:" + err.Error())
	}
	if domain+"/test.txt" != publicUrl {
		t.Error("publicUrl failed: Link inconsistency")
	}

	err = tLocal.Delete(context.Background(), "test.txt")
	if err != nil {
		t.Error("local write failed:" + err.Error())
	}
}

func TestStorageLocal(t *testing.T) {
	domain := "http://0.0.0.0:8080/uploads"
	tLocal := New("local", map[string]string{
		"path":   "./tmp",
		"domain": domain,
	})
	var err error
	content := "test"
	contentByte := []byte(content)

	err = tLocal.Write(context.Background(), "test.txt", content, map[string]any{})
	if err != nil {
		t.Error("write failed:" + err.Error())
	}

	err = tLocal.WriteStream(context.Background(), "test.txt", bytes.NewBuffer(contentByte), map[string]any{})
	if err != nil {
		t.Error("write stream failed:" + err.Error())
	}

	str, err := tLocal.Read(context.Background(), "test.txt")
	if err != nil {
		t.Error("read failed:" + err.Error())
	}
	if str != content {
		t.Error("local read failed: inconsistency of data")
	}

	reader, err := tLocal.ReadStream(context.Background(), "test.txt")
	if err != nil {
		t.Error("read failed:" + err.Error())
	}
	readerByte, err := io.ReadAll(reader)
	if err != nil {
		t.Error("read failed:" + err.Error())
	}
	if content != string(readerByte) {
		t.Error("read failed: inconsistency of data")
	}

	publicUrl, err := tLocal.PublicUrl(context.Background(), "test.txt")
	if err != nil {
		t.Error("publicUrl failed:" + err.Error())
	}
	if domain+"/test.txt" != publicUrl {
		t.Error("publicUrl failed: Link inconsistency")
	}

	err = tLocal.Delete(context.Background(), "test.txt")
	if err != nil {
		t.Error("delete failed:" + err.Error())
	}
}

func TestStorageOss(t *testing.T) {
	domain := ""
	tLocal := New("oss", map[string]string{
		"bucket":       "",
		"accessId":     "",
		"accessSecret": "",
		"endpoint":     "",
		"domain":       domain,
	})
	var err error
	content := "test"

	err = tLocal.Write(context.Background(), "test.txt", content, map[string]any{})
	if err != nil {
		t.Error("local write failed:" + err.Error())
	}

	str, err := tLocal.Read(context.Background(), "test.txt")
	if err != nil {
		t.Error("local read failed:" + err.Error())
	}
	if str != content {
		t.Error("local read failed: inconsistency of data")
	}

	publicUrl, err := tLocal.PublicUrl(context.Background(), "test.txt")
	if err != nil {
		t.Error("publicUrl failed:" + err.Error())
	}
	if domain+"/test.txt" != publicUrl {
		t.Error("publicUrl failed: Link inconsistency")
	}

	err = tLocal.Delete(context.Background(), "test.txt")
	if err != nil {
		t.Error("local write failed:" + err.Error())
	}
}

func TestStorageQiniu(t *testing.T) {
	tLocal := New("qiniu", map[string]string{
		"bucket":    "",
		"accessKey": "",
		"secretKey": "",
		"domain":    "",
	})
	var err error
	content := "test"

	err = tLocal.Write(context.Background(), "test.txt", content, map[string]any{})
	if err != nil {
		t.Error("write failed:" + err.Error())
	}

	str, err := tLocal.Read(context.Background(), "test.txt")
	if err != nil {
		t.Error("read failed:" + err.Error())
	}
	if str != content {
		t.Error("read failed: inconsistency of data")
	}

	err = tLocal.Delete(context.Background(), "test.txt")
	if err != nil {
		t.Error("write failed:" + err.Error())
	}
}
