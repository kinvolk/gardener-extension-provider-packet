package deprecated

type Deprecation struct {
	DeprecatedSince           int
	AlternativeAvailableSince int
}

var Stdlib = map[string]Deprecation{
	"image/jpeg.Reader": {4, 0},
	// FIXME(dh): AllowBinary isn't being detected as deprecated
	// because the comment has a newline right after "Deprecated:"
	"go/build.AllowBinary":                        {7, 7},
	"(archive/zip.FileHeader).CompressedSize":     {1, 1},
	"(archive/zip.FileHeader).UncompressedSize":   {1, 1},
	"(archive/zip.FileHeader).ModifiedTime":       {10, 10},
	"(archive/zip.FileHeader).ModifiedDate":       {10, 10},
	"(*archive/zip.FileHeader).ModTime":           {10, 10},
	"(*archive/zip.FileHeader).SetModTime":        {10, 10},
	"(go/doc.Package).Bugs":                       {1, 1},
	"os.SEEK_SET":                                 {7, 7},
	"os.SEEK_CUR":                                 {7, 7},
	"os.SEEK_END":                                 {7, 7},
	"(net.Dialer).Cancel":                         {7, 7},
	"runtime.CPUProfile":                          {9, 0},
	"compress/flate.ReadError":                    {6, 6},
	"compress/flate.WriteError":                   {6, 6},
	"path/filepath.HasPrefix":                     {0, 0},
	"(net/http.Transport).Dial":                   {7, 7},
	"(*net/http.Transport).CancelRequest":         {6, 5},
	"net/http.ErrWriteAfterFlush":                 {7, 0},
	"net/http.ErrHeaderTooLong":                   {8, 0},
	"net/http.ErrShortBody":                       {8, 0},
	"net/http.ErrMissingContentLength":            {8, 0},
	"net/http/httputil.ErrPersistEOF":             {0, 0},
	"net/http/httputil.ErrClosed":                 {0, 0},
	"net/http/httputil.ErrPipeline":               {0, 0},
	"net/http/httputil.ServerConn":                {0, 0},
	"net/http/httputil.NewServerConn":             {0, 0},
	"net/http/httputil.ClientConn":                {0, 0},
	"net/http/httputil.NewClientConn":             {0, 0},
	"net/http/httputil.NewProxyClientConn":        {0, 0},
	"(net/http.Request).Cancel":                   {7, 7},
	"(text/template/parse.PipeNode).Line":         {1, 1},
	"(text/template/parse.ActionNode).Line":       {1, 1},
	"(text/template/parse.BranchNode).Line":       {1, 1},
	"(text/template/parse.TemplateNode).Line":     {1, 1},
	"database/sql/driver.ColumnConverter":         {9, 9},
	"database/sql/driver.Execer":                  {8, 8},
	"database/sql/driver.Queryer":                 {8, 8},
	"(database/sql/driver.Conn).Begin":            {8, 8},
	"(database/sql/driver.Stmt).Exec":             {8, 8},
	"(database/sql/driver.Stmt).Query":            {8, 8},
	"syscall.StringByteSlice":                     {1, 1},
	"syscall.StringBytePtr":                       {1, 1},
	"syscall.StringSlicePtr":                      {1, 1},
	"syscall.StringToUTF16":                       {1, 1},
	"syscall.StringToUTF16Ptr":                    {1, 1},
	"(*regexp.Regexp).Copy":                       {12, 12},
	"(archive/tar.Header).Xattrs":                 {10, 10},
	"archive/tar.TypeRegA":                        {11, 1},
	"go/types.NewInterface":                       {11, 11},
	"(*go/types.Interface).Embedded":              {11, 11},
	"go/importer.For":                             {12, 12},
	"encoding/json.InvalidUTF8Error":              {2, 2},
	"encoding/json.UnmarshalFieldError":           {2, 2},
	"encoding/csv.ErrTrailingComma":               {2, 2},
	"(encoding/csv.Reader).TrailingComma":         {2, 2},
	"(net.Dialer).DualStack":                      {12, 12},
	"net/http.ErrUnexpectedTrailer":               {12, 12},
	"net/http.CloseNotifier":                      {11, 7},
	"net/http.ProtocolError":                      {8, 8},
	"(crypto/x509.CertificateRequest).Attributes": {5, 3},
	// This function has no alternative, but also no purpose.
	"(*crypto/rc4.Cipher).Reset":                     {12, 0},
	"(net/http/httptest.ResponseRecorder).HeaderMap": {11, 7},

	// All of these have been deprecated in favour of external libraries
	"syscall.AttachLsf":             {7, 0},
	"syscall.DetachLsf":             {7, 0},
	"syscall.LsfSocket":             {7, 0},
	"syscall.SetLsfPromisc":         {7, 0},
	"syscall.LsfJump":               {7, 0},
	"syscall.LsfStmt":               {7, 0},
	"syscall.BpfStmt":               {7, 0},
	"syscall.BpfJump":               {7, 0},
	"syscall.BpfBuflen":             {7, 0},
	"syscall.SetBpfBuflen":          {7, 0},
	"syscall.BpfDatalink":           {7, 0},
	"syscall.SetBpfDatalink":        {7, 0},
	"syscall.SetBpfPromisc":         {7, 0},
	"syscall.FlushBpf":              {7, 0},
	"syscall.BpfInterface":          {7, 0},
	"syscall.SetBpfInterface":       {7, 0},
	"syscall.BpfTimeout":            {7, 0},
	"syscall.SetBpfTimeout":         {7, 0},
	"syscall.BpfStats":              {7, 0},
	"syscall.SetBpfImmediate":       {7, 0},
	"syscall.SetBpf":                {7, 0},
	"syscall.CheckBpfVersion":       {7, 0},
	"syscall.BpfHeadercmpl":         {7, 0},
	"syscall.SetBpfHeadercmpl":      {7, 0},
	"syscall.RouteRIB":              {8, 0},
	"syscall.RoutingMessage":        {8, 0},
	"syscall.RouteMessage":          {8, 0},
	"syscall.InterfaceMessage":      {8, 0},
	"syscall.InterfaceAddrMessage":  {8, 0},
	"syscall.ParseRoutingMessage":   {8, 0},
	"syscall.ParseRoutingSockaddr":  {8, 0},
	"InterfaceAnnounceMessage":      {7, 0},
	"InterfaceMulticastAddrMessage": {7, 0},
	"syscall.FormatMessage":         {5, 0},
}
