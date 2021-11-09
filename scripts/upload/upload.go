package main

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type Api struct {
	c *s3.Client
}

func main() {
	a, err := New()
	if err != nil {
		panic(err)
	}
	err = a.Sync(os.Args[1], os.Args[2], msn, mpo)
	if err != nil {
		panic(err)
	}
}

var msnMap = make(map[string]bool)

var msn = func(src string) string {
	if strings.HasSuffix(src, ".html") && src != "index.html" && !strings.HasPrefix(src, "markup") {
		wasHtml := strings.TrimSuffix(src, ".html")
		msnMap[wasHtml] = true
		return wasHtml
	}
	return src
}

var mpo = func(poi *s3.PutObjectInput) error {
	if msnMap[*poi.Key] {
		poi.ContentType = aws.String("text/html")
	}
	return nil
}

func New() (a *Api, err error) {
	var cfg aws.Config
	if cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1")); err == nil {
		a = &Api{c: s3.NewFromConfig(cfg)}
	}
	return
}

type SourceFile struct {
	Info    fs.FileInfo
	SrcPath string
	RelPath string
}

func (sf *SourceFile) Open() (io.Reader, error) {
	return os.Open(filepath.Join(sf.SrcPath, sf.RelPath))
}

func getSourceFiles(src string) (result []*SourceFile, err error) {
	err = filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			var rp string
			rp, err = filepath.Rel(src, path)
			if err != nil {
				return err
			}
			result = append(result, &SourceFile{Info: info, SrcPath: src, RelPath: rp})
		}
		return nil
	})
	return
}

func (a *Api) getDestFiles(bucket string) (result []types.Object, err error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	var ct *string
	for {
		var resp *s3.ListObjectsV2Output
		input := &s3.ListObjectsV2Input{Bucket: &bucket, ContinuationToken: ct}
		if resp, err = a.c.ListObjectsV2(ctx, input); err != nil {
			return
		}
		for _, item := range resp.Contents {
			result = append(result, item)
		}
		if !resp.IsTruncated {
			break
		}
		ct = resp.NextContinuationToken
	}
	return
}

type FileCompare struct {
	Source     *SourceFile
	Dest       *types.Object
	SourceTime *time.Time
	DestTime   *time.Time
}

func (fc *FileCompare) extractTimes() {
	cleanupTime := func(t time.Time) *time.Time {
		r := t.Local().Truncate(time.Second)
		return &r
	}
	if fc.Source != nil {
		fc.SourceTime = cleanupTime(fc.Source.Info.ModTime())
	}
	if fc.Dest != nil {
		fc.DestTime = cleanupTime(*fc.Dest.LastModified)
	}
}

func (fc *FileCompare) DestNewer() bool {
	return fc.SourceTime != nil &&
		fc.DestTime != nil &&
		fc.SourceTime.Before(*fc.DestTime)
}

func (fc *FileCompare) SourceMissing() bool {
	return fc.SourceTime == nil && fc.DestTime != nil
}

// ModifyPutObject can be used to modify attributes of the object when uploaded.
type ModifyPutObject func(poi *s3.PutObjectInput) error

// ModifySourceName can be used to modify the file name that is uploaded to s3
type ModifySourceName func(src string) string

type FileCompareMap map[string]*FileCompare

func (a *Api) Sync(src, dest string, msn ModifySourceName, mpo ModifyPutObject) error {
	fm := make(FileCompareMap)
	if files, err := getSourceFiles(src); err != nil {
		return err
	} else {
		for _, file := range files {
			fm[msn(file.RelPath)] = &FileCompare{Source: file}
		}
	}

	if files, err := a.getDestFiles(dest); err != nil {
		return err
	} else {
		for _, f := range files {
			file := f
			key := *file.Key
			if fc, ok := fm[key]; ok {
				fc.Dest = &file
			} else {
				fm[key] = &FileCompare{Dest: &file}
			}
		}
	}

	for k, c := range fm {
		key := k
		compare := c
		compare.extractTimes()

		if compare.DestNewer() {
			continue
		}
		if compare.SourceMissing() {
			fmt.Println("deleting", *compare.Dest.Key)

			_, err := a.c.DeleteObject(context.TODO(), &s3.DeleteObjectInput{Bucket: &dest, Key: compare.Dest.Key})
			if err != nil {
				return err
			}

			continue
		}
		contentType := mime.TypeByExtension(filepath.Ext(key))

		open, err := compare.Source.Open()
		if err != nil {
			return err
		}

		params := &s3.PutObjectInput{Bucket: &dest, Key: &key, ContentType: &contentType, Body: open}
		err = mpo(params)
		if err != nil {
			return err
		}
		fmt.Println("uploading", key)
		_, err = a.c.PutObject(context.Background(), params)
		if err != nil {
			return err
		}
	}

	return nil

}
