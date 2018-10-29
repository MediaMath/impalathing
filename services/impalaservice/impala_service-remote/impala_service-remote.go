// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "context"
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
	"github.com/MediaMath/impalathing/services/execstats"
	"github.com/MediaMath/impalathing/services/status"
	"github.com/MediaMath/impalathing/services/types"
	"github.com/MediaMath/impalathing/services/beeswax"
	"github.com/MediaMath/impalathing/services/tcliservice"
        "github.com/MediaMath/impalathing/services/impalaservice"
)

var _ = execstats.GoUnusedProtection__
var _ = status.GoUnusedProtection__
var _ = types.GoUnusedProtection__
var _ = beeswax.GoUnusedProtection__
var _ = tcliservice.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  TStatus Cancel(QueryHandle query_id)")
  fmt.Fprintln(os.Stderr, "  TStatus ResetCatalog()")
  fmt.Fprintln(os.Stderr, "  TStatus ResetTable(TResetTableReq request)")
  fmt.Fprintln(os.Stderr, "  string GetRuntimeProfile(QueryHandle query_id)")
  fmt.Fprintln(os.Stderr, "  TInsertResult CloseInsert(QueryHandle handle)")
  fmt.Fprintln(os.Stderr, "  TPingImpalaServiceResp PingImpalaService()")
  fmt.Fprintln(os.Stderr, "  TExecSummary GetExecSummary(QueryHandle handle)")
  fmt.Fprintln(os.Stderr, "  QueryHandle query(Query query)")
  fmt.Fprintln(os.Stderr, "  QueryHandle executeAndWait(Query query, LogContextId clientCtx)")
  fmt.Fprintln(os.Stderr, "  QueryExplanation explain(Query query)")
  fmt.Fprintln(os.Stderr, "  Results fetch(QueryHandle query_id, bool start_over, i32 fetch_size)")
  fmt.Fprintln(os.Stderr, "  QueryState get_state(QueryHandle handle)")
  fmt.Fprintln(os.Stderr, "  ResultsMetadata get_results_metadata(QueryHandle handle)")
  fmt.Fprintln(os.Stderr, "  string echo(string s)")
  fmt.Fprintln(os.Stderr, "  string dump_config()")
  fmt.Fprintln(os.Stderr, "  string get_log(LogContextId context)")
  fmt.Fprintln(os.Stderr, "   get_default_configuration(bool include_hadoop)")
  fmt.Fprintln(os.Stderr, "  void close(QueryHandle handle)")
  fmt.Fprintln(os.Stderr, "  void clean(LogContextId log_context)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := impalaservice.NewImpalaServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "Cancel":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Cancel requires 1 args")
      flag.Usage()
    }
    arg17 := flag.Arg(1)
    mbTrans18 := thrift.NewTMemoryBufferLen(len(arg17))
    defer mbTrans18.Close()
    _, err19 := mbTrans18.WriteString(arg17)
    if err19 != nil {
      Usage()
      return
    }
    factory20 := thrift.NewTJSONProtocolFactory()
    jsProt21 := factory20.GetProtocol(mbTrans18)
    argvalue0 := beeswax.NewQueryHandle()
    err22 := argvalue0.Read(jsProt21)
    if err22 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Cancel(context.Background(), value0))
    fmt.Print("\n")
    break
  case "ResetCatalog":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "ResetCatalog requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.ResetCatalog(context.Background()))
    fmt.Print("\n")
    break
  case "ResetTable":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ResetTable requires 1 args")
      flag.Usage()
    }
    arg23 := flag.Arg(1)
    mbTrans24 := thrift.NewTMemoryBufferLen(len(arg23))
    defer mbTrans24.Close()
    _, err25 := mbTrans24.WriteString(arg23)
    if err25 != nil {
      Usage()
      return
    }
    factory26 := thrift.NewTJSONProtocolFactory()
    jsProt27 := factory26.GetProtocol(mbTrans24)
    argvalue0 := impalaservice.NewTResetTableReq()
    err28 := argvalue0.Read(jsProt27)
    if err28 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.ResetTable(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetRuntimeProfile":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetRuntimeProfile requires 1 args")
      flag.Usage()
    }
    arg29 := flag.Arg(1)
    mbTrans30 := thrift.NewTMemoryBufferLen(len(arg29))
    defer mbTrans30.Close()
    _, err31 := mbTrans30.WriteString(arg29)
    if err31 != nil {
      Usage()
      return
    }
    factory32 := thrift.NewTJSONProtocolFactory()
    jsProt33 := factory32.GetProtocol(mbTrans30)
    argvalue0 := beeswax.NewQueryHandle()
    err34 := argvalue0.Read(jsProt33)
    if err34 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetRuntimeProfile(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CloseInsert":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CloseInsert requires 1 args")
      flag.Usage()
    }
    arg35 := flag.Arg(1)
    mbTrans36 := thrift.NewTMemoryBufferLen(len(arg35))
    defer mbTrans36.Close()
    _, err37 := mbTrans36.WriteString(arg35)
    if err37 != nil {
      Usage()
      return
    }
    factory38 := thrift.NewTJSONProtocolFactory()
    jsProt39 := factory38.GetProtocol(mbTrans36)
    argvalue0 := beeswax.NewQueryHandle()
    err40 := argvalue0.Read(jsProt39)
    if err40 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CloseInsert(context.Background(), value0))
    fmt.Print("\n")
    break
  case "PingImpalaService":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "PingImpalaService requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.PingImpalaService(context.Background()))
    fmt.Print("\n")
    break
  case "GetExecSummary":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetExecSummary requires 1 args")
      flag.Usage()
    }
    arg41 := flag.Arg(1)
    mbTrans42 := thrift.NewTMemoryBufferLen(len(arg41))
    defer mbTrans42.Close()
    _, err43 := mbTrans42.WriteString(arg41)
    if err43 != nil {
      Usage()
      return
    }
    factory44 := thrift.NewTJSONProtocolFactory()
    jsProt45 := factory44.GetProtocol(mbTrans42)
    argvalue0 := beeswax.NewQueryHandle()
    err46 := argvalue0.Read(jsProt45)
    if err46 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetExecSummary(context.Background(), value0))
    fmt.Print("\n")
    break
  case "query":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Query requires 1 args")
      flag.Usage()
    }
    arg47 := flag.Arg(1)
    mbTrans48 := thrift.NewTMemoryBufferLen(len(arg47))
    defer mbTrans48.Close()
    _, err49 := mbTrans48.WriteString(arg47)
    if err49 != nil {
      Usage()
      return
    }
    factory50 := thrift.NewTJSONProtocolFactory()
    jsProt51 := factory50.GetProtocol(mbTrans48)
    argvalue0 := beeswax.NewQuery()
    err52 := argvalue0.Read(jsProt51)
    if err52 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Query(context.Background(), value0))
    fmt.Print("\n")
    break
  case "executeAndWait":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ExecuteAndWait requires 2 args")
      flag.Usage()
    }
    arg53 := flag.Arg(1)
    mbTrans54 := thrift.NewTMemoryBufferLen(len(arg53))
    defer mbTrans54.Close()
    _, err55 := mbTrans54.WriteString(arg53)
    if err55 != nil {
      Usage()
      return
    }
    factory56 := thrift.NewTJSONProtocolFactory()
    jsProt57 := factory56.GetProtocol(mbTrans54)
    argvalue0 := beeswax.NewQuery()
    err58 := argvalue0.Read(jsProt57)
    if err58 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := beeswax.LogContextId(argvalue1)
    fmt.Print(client.ExecuteAndWait(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "explain":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Explain requires 1 args")
      flag.Usage()
    }
    arg60 := flag.Arg(1)
    mbTrans61 := thrift.NewTMemoryBufferLen(len(arg60))
    defer mbTrans61.Close()
    _, err62 := mbTrans61.WriteString(arg60)
    if err62 != nil {
      Usage()
      return
    }
    factory63 := thrift.NewTJSONProtocolFactory()
    jsProt64 := factory63.GetProtocol(mbTrans61)
    argvalue0 := beeswax.NewQuery()
    err65 := argvalue0.Read(jsProt64)
    if err65 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Explain(context.Background(), value0))
    fmt.Print("\n")
    break
  case "fetch":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "Fetch requires 3 args")
      flag.Usage()
    }
    arg66 := flag.Arg(1)
    mbTrans67 := thrift.NewTMemoryBufferLen(len(arg66))
    defer mbTrans67.Close()
    _, err68 := mbTrans67.WriteString(arg66)
    if err68 != nil {
      Usage()
      return
    }
    factory69 := thrift.NewTJSONProtocolFactory()
    jsProt70 := factory69.GetProtocol(mbTrans67)
    argvalue0 := beeswax.NewQueryHandle()
    err71 := argvalue0.Read(jsProt70)
    if err71 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2) == "true"
    value1 := argvalue1
    tmp2, err73 := (strconv.Atoi(flag.Arg(3)))
    if err73 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.Fetch(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "get_state":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetState requires 1 args")
      flag.Usage()
    }
    arg74 := flag.Arg(1)
    mbTrans75 := thrift.NewTMemoryBufferLen(len(arg74))
    defer mbTrans75.Close()
    _, err76 := mbTrans75.WriteString(arg74)
    if err76 != nil {
      Usage()
      return
    }
    factory77 := thrift.NewTJSONProtocolFactory()
    jsProt78 := factory77.GetProtocol(mbTrans75)
    argvalue0 := beeswax.NewQueryHandle()
    err79 := argvalue0.Read(jsProt78)
    if err79 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetState(context.Background(), value0))
    fmt.Print("\n")
    break
  case "get_results_metadata":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetResultsMetadata requires 1 args")
      flag.Usage()
    }
    arg80 := flag.Arg(1)
    mbTrans81 := thrift.NewTMemoryBufferLen(len(arg80))
    defer mbTrans81.Close()
    _, err82 := mbTrans81.WriteString(arg80)
    if err82 != nil {
      Usage()
      return
    }
    factory83 := thrift.NewTJSONProtocolFactory()
    jsProt84 := factory83.GetProtocol(mbTrans81)
    argvalue0 := beeswax.NewQueryHandle()
    err85 := argvalue0.Read(jsProt84)
    if err85 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetResultsMetadata(context.Background(), value0))
    fmt.Print("\n")
    break
  case "echo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Echo requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.Echo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "dump_config":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "DumpConfig requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.DumpConfig(context.Background()))
    fmt.Print("\n")
    break
  case "get_log":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetLog requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := beeswax.LogContextId(argvalue0)
    fmt.Print(client.GetLog(context.Background(), value0))
    fmt.Print("\n")
    break
  case "get_default_configuration":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetDefaultConfiguration requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1) == "true"
    value0 := argvalue0
    fmt.Print(client.GetDefaultConfiguration(context.Background(), value0))
    fmt.Print("\n")
    break
  case "close":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Close requires 1 args")
      flag.Usage()
    }
    arg89 := flag.Arg(1)
    mbTrans90 := thrift.NewTMemoryBufferLen(len(arg89))
    defer mbTrans90.Close()
    _, err91 := mbTrans90.WriteString(arg89)
    if err91 != nil {
      Usage()
      return
    }
    factory92 := thrift.NewTJSONProtocolFactory()
    jsProt93 := factory92.GetProtocol(mbTrans90)
    argvalue0 := beeswax.NewQueryHandle()
    err94 := argvalue0.Read(jsProt93)
    if err94 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Close(context.Background(), value0))
    fmt.Print("\n")
    break
  case "clean":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Clean requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := beeswax.LogContextId(argvalue0)
    fmt.Print(client.Clean(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
