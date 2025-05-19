package coupons

import (
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"tkh/logger"

	"os"
	"path/filepath"
	"strings"
)

const MatchThreshold = 2

var files = []string{"couponbase1.gz", "couponbase2.gz", "couponbase3.gz"}

func FindCoupon(ctx context.Context, couponCode string) (bool, error) {
	found := make(chan bool, 1)
	for _, val := range files {
		go func(fileName string) {
			err := findCodeInFile(context.Background(), fileName, couponCode, found)
			if err != nil {
				logger.Println(err)
				return
			}
		}(val)
	}
	count := 0
	total := 0
	for ok := range found {
		total++
		if ok {
			count++
		}
		if count >= MatchThreshold {
			return true, nil
		}
		if total == len(files) {
			break
		}
	}
	return false, nil
}

func findCodeInFile(ctx context.Context, filePath string, couponCode string, ch chan bool) error {
	file, err := os.Open(filepath.Join(os.Getenv("coupon_base_path"), filePath))
	if err != nil {
		logger.Println(err)
		ch <- false
		return err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		logger.Println(err)
		ch <- false
		return err
	}
	defer gzipReader.Close()

	b := make([]byte, 11)
	tempArr := make([]string, 0, 2)
	temp := ""
	for {
		_, err := gzipReader.Read(b)
		if err != nil && !errors.Is(err, io.EOF) {
			logger.Println("Error reading gzip file:", err)
			break
		}
		tempArr = strings.SplitN(string(b), "\n", 2)
		if couponCode == fmt.Sprintf("%s%s", temp, tempArr[0]) {
			ch <- true
			break
		}
		temp = ""
		if len(tempArr) == 1 {
			continue
		}
		if strings.Contains(tempArr[1], "\n") {
			tempArr = strings.SplitN(tempArr[1], "\n", 2)
			if couponCode == tempArr[0] {
				ch <- true
				break
			}
		}
		temp = tempArr[1]
		if errors.Is(err, io.EOF) {
			break
		}
	}
	ch <- false
	return nil
}
