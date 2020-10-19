package expkeyutil

import (
	"fmt"
	"github.com/pantraeu/pantra_server/pkg/pantra_server/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetBinFiles(root string) []string {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if _, err := os.Stat(path); !os.IsNotExist(err) && strings.HasSuffix(path, ".bin") {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return files
}

func GetTmpKey(binFile string) (*pb.TemporaryExposureKeyExport, error) {

	in, err := ioutil.ReadFile(binFile)

	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	keys := &pb.TemporaryExposureKeyExport{}

	if err := proto.Unmarshal(in[16:], keys); err != nil {
		log.Fatalln("Failed to parse TemporaryExposureKey:", err)
	} else {
		return keys, nil
	}
	return nil, fmt.Errorf("")
}
