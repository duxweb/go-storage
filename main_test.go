package storage

import (
	"context"
	"testing"
)

func Test(t *testing.T) {
	domain := "http://dux.test"
	s3, err := New("local", map[string]string{
		// Local
		"root":   ".",
		"path":   "upload",
		"domain": domain,

		// S3
		//"region":    "",
		//"endpoint":  "",
		//"bucket":    "",
		//"accessKey": "",
		//"secretKey": "",
	}, nil)
	if err != nil {
		t.Error("local init failed:" + err.Error())
	}

	content := "test"

	err = s3.Write(context.Background(), "test.txt", content)
	if err != nil {
		t.Error("local write failed:" + err.Error())
	}

	str, err := s3.Read(context.Background(), "test.txt")
	if err != nil {
		t.Error("local read failed:" + err.Error())
	}
	if str != content {
		t.Error("local read failed: inconsistency of data")
	}

	publicUrl, err := s3.PublicUrl(context.Background(), "test.txt")
	if err != nil {
		t.Error("publicUrl failed:" + err.Error())
	}
	t.Log("publicUrl: " + publicUrl)

	if domain+"/test.txt" != publicUrl {
		t.Error("publicUrl failed: Link inconsistency")
	}

	privateUrl, err := s3.PrivateUrl(context.Background(), "test.txt")
	if err != nil {
		t.Error("privateUrl failed:" + err.Error())
	}
	t.Log("privateUrl: " + privateUrl)

	url, params, err := s3.SignPostUrl(context.Background(), "test2.txt")
	if err != nil {
		t.Error("privateUrl failed:" + err.Error())
	}
	t.Log("postUrl: " + url)
	t.Log("postParams", params)

	url, err = s3.SignPutUrl(context.Background(), "test3.txt")
	if err != nil {
		t.Error("privateUrl failed:" + err.Error())
	}
	t.Log("putUrl: " + url)

	err = s3.Delete(context.Background(), "test.txt")
	if err != nil {
		t.Error("local write failed:" + err.Error())
	}
}
