package splunkrest

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Index struct {
	AssureUTF8               bool    `json:"assureUTF8"`
	BlockSignSize            int     `json:"blockSignSize"`
	BlockSignatureDatabase   string  `json:"blockSignatureDatabase"`
	BloomfilterTotalSizeKB   int     `json:"bloomfilterTotalSizeKB"`
	ColdPath                 string  `json:"coldPath"`
	ColdToFrozenDir          string  `json:"coldToFrozenDir"`
	ColdToFrozenScript       string  `json:"coldToFrozenScript"`
	CompressRawdata          bool    `json:"compressRawdata"`
	CurrentDBSizeMB          string  `json:"currentDBSizeMB"`
	DefaultDatabase          string  `json:"defaultDatabase"`
	Disabled                 bool    `json:"disabled"`
	EnableRealtimeSearch     bool    `json:enableRealtimeSearch""`
	FrozenTimePeriodInSecs   int     `json:"frozenTimePeriodInSecs"`
	HomePath                 string  `json:"homePath`
	HomePathExpanded         string  `json:"homePath_expanded"`
	IndexThreads             string  `json:"indexThreads"`
	IsInternal               bool    `json:"isInternal"`
	LastInitTime             float32 `json:"lastInitTime"`
	MaxConcurrentOptimizes   int     `json:"maxConcurrentOptimizes"`
	MaxDataSize              string  `json:"maxDataSize"`
	MaxHotBuckets            int     `json:"maxHotBuckets"`
	MaxHotIdleSecs           int     `json:"maxHotIdleSecs"`
	MaxHotSpanSecs           int     `json:"maxHotSpanSecs"`
	MaxMemMB                 int     `json:"maxMemMB"`
	MaxMetaEntries           int     `json:"maxMetaEntries"`
	MaxRunningProcessGroups  int     `json:"maxRunningProcessGroups"`
	MaxTime                  string  `json:"maxTime"`
	MaxTotalDataSizeMB       int     `json:"maxTotalDataSizeMB"`
	MaxWarmDBCount           int     `json:"maxWarmDBCount"`
	MemPoolMB                string  `json:"memPoolMB"`
	MinRawFileSyncSecs       string  `json:"minRawFileSyncSecs"`
	MinTime                  string  `json:"minTime"`
	NumBloomfilters          int     `json:"numBloomfilters"`
	NumHotBuckets            int     `json:"numHotBuckets"`
	NumWarmBuckets           int     `json:"numWarmBuckets"`
	PartialServiceMetaPeriod int     `json:"partialServiceMetaPeriod"`
	QuarantineFutureSecs     int     `json:"quarantineFutureSecs"`
	QuarantinePastSecs       int     `json:"quarantinePastSecs"`
	RawChunkSizeBytes        int     `json:"rawChunkSizeBytes"`
	RotatePeriodInSecs       int     `json:"rotatePeriodInSecs"`
	ServiceMetaPeriod        int     `json:"serviceMetaPeriod"`
	SuppressBannerList       string  `json:"suppressBannerList"`
	Sync                     int     `json:"sync"`
	SyncMeta                 bool    `json:"syncMeta"`
	ThawedPath               string  `json:"thawedPath"`
	ThawedPathExpanded       string  `json:"thawedPath_expanded"`
	ThrottleCheckPeriod      int     `json:"throttleCheckPeriod"`
	TotalEventCount          int     `json:"totalEventCount"`
}

type IndexResponse struct {
	Entry []IndexEntry `json:"entry"`
}

type IndexEntry struct {
	Content Index `json:"content"`
}

func (conn SplunkConnection) GetIndex(name string) (Index, error) {
	data := make(url.Values)
	data.Add("name", name)
	data.Add("output_mode", "json")
	response, err := conn.HTTPGet(fmt.Sprintf("%s/services/data/indexes/%s", conn.BaseURL, name), &data)

	if err != nil {
		return Index{}, err
	}

	bytes := []byte(response)
	var indexResponse IndexResponse
	err = json.Unmarshal(bytes, &indexResponse)
	return indexResponse.Entry[0].Content, err
}
