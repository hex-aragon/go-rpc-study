package log

import (
	"io/ioutil"
	"path"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Log struct {
	mu sync.RWMutex 
	Dir string 
	Config Config 

	activeSegment *segment 
	segments []*segment
}

func NewLog(dir string, c Config) (*Log, error){
	if c.Segment.MaxStoreBytes == 0 {
		c.Segment.MaxStoreBytes = 1024 
	}

	if c.Segment.MaxIndexBytes == 0 {
		c.Segment.MaxIndexBytes = 1024 
	}

	l := &Log {
		Dir: dir, 
		Config: c, 
	}

	return l, l.setup()
}

func (l *Log) setup() error {
	files, err := ioutil.ReadDir(l.Dir)
	if err != nil {
		return err 
	}

	var baseOffsets []uint64 
	for _, file := range files {
		offStr := strings.TrimSuffix(
			file.Name(),
			path.Ext(file.Name()), 
		)
		off , _ := strconv.ParseUint(offStr, 10, 0 )
		baseOffsets  = append(baseOffsets, off)
	}

	sort.Slice(baseOffsets, func(i, j int) bool {
		return baseOffsets[i] < baseOffsets[j]
	})

	for i := 0; i < len(baseOffsets); i++{
		if err = l.newSegment(baseOffsets[i]); err != nil {
			return err 
		}
		// 베이스 오프셋은 index와 store 두 파일을 중복해서 담고 있기에 
		// 같은 값이 하나 더 있다. 그래서 한 번 건너뛴다.  
		i++
	}

	if l.segments == nil {
		if err = l.newSegment(
			l.Config.Segment.InitialOffset,
		); err != nil {
			return err 
		}
	}

	return nil 
}