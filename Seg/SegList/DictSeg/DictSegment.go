package DictSeg

import (
	"fmt"
	"github.com/zhmc/PowerNLP/Config"
	"github.com/zhmc/PowerNLP/Seg/Collections"
	"time"
)

var (
	MapTrieSeg *Collections.MapTrie = nil
)

func init() {
	MapTrieSeg = Collections.NewMapTrie()
	MapTrieSeg.LoadDict("data/dict.txt")
	time.Sleep( time.Second *2 )
	fmt.Println("字典加载完毕")
	MapTrieSeg.LoadDict(Config.DictPath)
}
