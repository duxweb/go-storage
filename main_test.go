package storage

import (
	"context"
	"testing"
)

func Test(t *testing.T) {
	domain := ""
	s3, err := New("s3", map[string]string{
		"region":    "",
		"endpoint":  "",
		"bucket":    "",
		"accessKey": "",
		"secretKey": "",
		"domain":    domain,
	}, nil)
	if err != nil {
		t.Error("local init failed:" + err.Error())
	}

	content := "test"

	err = s3.Write(context.Background(), "test/test.txt", content)
	if err != nil {
		t.Error("local write failed:" + err.Error())
	}

	str, err := s3.Read(context.Background(), "test/test.txt")
	if err != nil {
		t.Error("local read failed:" + err.Error())
	}
	if str != content {
		t.Error("local read failed: inconsistency of data")
	}

	publicUrl, err := s3.PublicUrl(context.Background(), "test/test.txt")
	if err != nil {
		t.Error("publicUrl failed:" + err.Error())
	}
	t.Log("publicUrl: " + publicUrl)

	if domain+"/test/test.txt" != publicUrl {
		t.Error("publicUrl failed: Link inconsistency")
	}

	privateUrl, err := s3.PrivateUrl(context.Background(), "test/test.txt")
	if err != nil {
		t.Error("privateUrl failed:" + err.Error())
	}
	t.Log("privateUrl: " + privateUrl)

	url, params, err := s3.SignPostUrl(context.Background(), "test/test2.txt")
	if err != nil {
		t.Error("privateUrl failed:" + err.Error())
	}
	t.Log("postUrl: " + url)
	t.Log("postParams", params)

	url, err = s3.SignPutUrl(context.Background(), "test/test3.txt")
	if err != nil {
		t.Error("privateUrl failed:" + err.Error())
	}
	t.Log("putUrl: " + url)

	err = s3.Delete(context.Background(), "test/test.txt")
	if err != nil {
		t.Error("local write failed:" + err.Error())
	}
}
