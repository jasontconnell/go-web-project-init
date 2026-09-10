package process

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/jasontconnell/go-web-project-init/data"
)

func ParseKeyValue(filename string, varwrap string) ([]data.KeyValue, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("can't read file %s %w", filename, err)
	}

	kv := []data.KeyValue{}

	buf := bytes.NewBuffer(b)
	scn := bufio.NewScanner(buf)
	for scn.Scan() {
		line := scn.Text()
		sp := strings.Split(line, "=")
		if len(sp) != 2 {
			continue
		}

		k, v := sp[0], sp[1]
		kv = append(kv, data.KeyValue{Key: varwrap + k + varwrap, Value: v})
	}

	return kv, nil
}
