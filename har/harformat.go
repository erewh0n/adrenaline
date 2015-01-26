package har

type Creator struct {
	Name    string `name`
	Version string `version`
	Comment string `comment`
}

type Browser struct {
	Name    string `name`
	Version string `version`
	Comment string `comment`
}

type Page struct {
	StartedDateTime string       `startDateTime`
	Id              string       `id`
	Title           string       `title`
	PageTimings     []PageTiming `pageTimings`
	Comment         string       `comment`
}

type PageTiming struct {
	OnContentLoad int32  `onContentLoad`
	OnLoad        int32  `onLoad`
	Comment       string `comment`
}

type Entry struct {
	PageRef         string   `pageRef`
	StartedDateTime string   `startedDateTime`
	Time            int32    `time`
	Request         Request  `request`
	Response        Response `response`
	Cache           Cache    `cache`
	Timings         Timings  `timings`
	ServerIPAddress string   `serverIPAddress`
	Connection      string   `connection`
	Comment         string   `comment`
}

type Request struct {
	Method      string       `method`
	Url         string       `url`
	HttpVersion string       `httpVersion`
	Cookies     []Cookie     `cookies`
	Headers     []Header     `headers`
	QueryString []QueryParam `queryString`
	PostData    PostData     `postData`
	HeaderSize  int32        `headerSize`
	BodySize    int32        `bodySize`
	Comment     string       `comment`
}

type Response struct {
	Status      int32    `status`
	StatusText  string   `statusText`
	HttpVersion string   `httpVersion`
	Cookies     []Cookie `cookies`
	Headers     []Header `headers`
	Content     Content  `content`
	RedirectURL string   `redirectURL`
	HeaderSize  int32    `headerSize`
	BodySize    int32    `bodySize`
	Comment     string   `comment`
}

type Cookie struct {
	Name     string `name`
	Value    string `value`
	Path     string `path`
	Domain   string `domain`
	Expires  string `expires`
	HttpOnly bool   `httpOnly`
	Secure   bool   `secure`
	Comment  string `comment`
}

type Header struct {
	Name    string `name`
	Value   string `value`
	Comment string `comment`
}

type QueryParam struct {
	Name    string `name`
	Value   string `value`
	Comment string `comment`
}

type PostData struct {
	MimeType string      `mimeType`
	Params   []PostParam `params`
	Text     string      `text`
	Comment  string      `comment`
}

type PostParam struct {
	Name        string `name`
	Value       string `value`
	FileName    string `fileName`
	ContentType string `contentType`
	Comment     string `comment`
}

type Content struct {
	Size        int32  `size`
	Compression int32  `compression`
	MimeType    string `mimeType`
	Text        string `text`
	Encoding    string `encoding`
	Comment     string `comment`
}

type Cache struct {
	BeforeRequest CacheInfo `beforeRequest`
	AfterRequest  CacheInfo `afterRequest`
}

type CacheInfo struct {
	Expires    string `expires`
	LastAccess string `lastAccess`
	ETag       string `eTag`
	HitCount   int32  `hitCount`
	Comment    string `comment`
}

type Timings struct {
	Blocked int32  `blocked`
	DNS     int32  `dns`
	Connect int32  `connect`
	Send    int32  `send`
	Wait    int32  `wait`
	Receive int32  `receive`
	SSL     int32  `ssl`
	Comment string `comment`
}

type Log struct {
	Version string  `version`
	Creator Creator `creator`
	Browser Browser `browser`
	Pages   []Page  `pages`
	Entries []Entry `entries`
	Comment string  `comment`
}

type HARLog struct {
	Log Log `log`
}
