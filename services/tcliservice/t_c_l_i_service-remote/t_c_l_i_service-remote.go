// Autogenerated by Thrift Compiler (0.11.0)
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
        "github.com/MediaMath/gudu/services/tcliservice"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  TOpenSessionResp OpenSession(TOpenSessionReq req)")
  fmt.Fprintln(os.Stderr, "  TCloseSessionResp CloseSession(TCloseSessionReq req)")
  fmt.Fprintln(os.Stderr, "  TGetInfoResp GetInfo(TGetInfoReq req)")
  fmt.Fprintln(os.Stderr, "  TExecuteStatementResp ExecuteStatement(TExecuteStatementReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTypeInfoResp GetTypeInfo(TGetTypeInfoReq req)")
  fmt.Fprintln(os.Stderr, "  TGetCatalogsResp GetCatalogs(TGetCatalogsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetSchemasResp GetSchemas(TGetSchemasReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTablesResp GetTables(TGetTablesReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTableTypesResp GetTableTypes(TGetTableTypesReq req)")
  fmt.Fprintln(os.Stderr, "  TGetColumnsResp GetColumns(TGetColumnsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetFunctionsResp GetFunctions(TGetFunctionsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetOperationStatusResp GetOperationStatus(TGetOperationStatusReq req)")
  fmt.Fprintln(os.Stderr, "  TCancelOperationResp CancelOperation(TCancelOperationReq req)")
  fmt.Fprintln(os.Stderr, "  TCloseOperationResp CloseOperation(TCloseOperationReq req)")
  fmt.Fprintln(os.Stderr, "  TGetResultSetMetadataResp GetResultSetMetadata(TGetResultSetMetadataReq req)")
  fmt.Fprintln(os.Stderr, "  TFetchResultsResp FetchResults(TFetchResultsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetDelegationTokenResp GetDelegationToken(TGetDelegationTokenReq req)")
  fmt.Fprintln(os.Stderr, "  TCancelDelegationTokenResp CancelDelegationToken(TCancelDelegationTokenReq req)")
  fmt.Fprintln(os.Stderr, "  TRenewDelegationTokenResp RenewDelegationToken(TRenewDelegationTokenReq req)")
  fmt.Fprintln(os.Stderr, "  TGetLogResp GetLog(TGetLogReq req)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
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
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
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
  client := tcliservice.NewTCLIServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "OpenSession":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "OpenSession requires 1 args")
      flag.Usage()
    }
    arg69 := flag.Arg(1)
    mbTrans70 := thrift.NewTMemoryBufferLen(len(arg69))
    defer mbTrans70.Close()
    _, err71 := mbTrans70.WriteString(arg69)
    if err71 != nil {
      Usage()
      return
    }
    factory72 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt73 := factory72.GetProtocol(mbTrans70)
    argvalue0 := tcliservice.NewTOpenSessionReq()
    err74 := argvalue0.Read(jsProt73)
    if err74 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.OpenSession(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CloseSession":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CloseSession requires 1 args")
      flag.Usage()
    }
    arg75 := flag.Arg(1)
    mbTrans76 := thrift.NewTMemoryBufferLen(len(arg75))
    defer mbTrans76.Close()
    _, err77 := mbTrans76.WriteString(arg75)
    if err77 != nil {
      Usage()
      return
    }
    factory78 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt79 := factory78.GetProtocol(mbTrans76)
    argvalue0 := tcliservice.NewTCloseSessionReq()
    err80 := argvalue0.Read(jsProt79)
    if err80 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CloseSession(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetInfo requires 1 args")
      flag.Usage()
    }
    arg81 := flag.Arg(1)
    mbTrans82 := thrift.NewTMemoryBufferLen(len(arg81))
    defer mbTrans82.Close()
    _, err83 := mbTrans82.WriteString(arg81)
    if err83 != nil {
      Usage()
      return
    }
    factory84 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt85 := factory84.GetProtocol(mbTrans82)
    argvalue0 := tcliservice.NewTGetInfoReq()
    err86 := argvalue0.Read(jsProt85)
    if err86 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "ExecuteStatement":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ExecuteStatement requires 1 args")
      flag.Usage()
    }
    arg87 := flag.Arg(1)
    mbTrans88 := thrift.NewTMemoryBufferLen(len(arg87))
    defer mbTrans88.Close()
    _, err89 := mbTrans88.WriteString(arg87)
    if err89 != nil {
      Usage()
      return
    }
    factory90 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt91 := factory90.GetProtocol(mbTrans88)
    argvalue0 := tcliservice.NewTExecuteStatementReq()
    err92 := argvalue0.Read(jsProt91)
    if err92 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.ExecuteStatement(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetTypeInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTypeInfo requires 1 args")
      flag.Usage()
    }
    arg93 := flag.Arg(1)
    mbTrans94 := thrift.NewTMemoryBufferLen(len(arg93))
    defer mbTrans94.Close()
    _, err95 := mbTrans94.WriteString(arg93)
    if err95 != nil {
      Usage()
      return
    }
    factory96 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt97 := factory96.GetProtocol(mbTrans94)
    argvalue0 := tcliservice.NewTGetTypeInfoReq()
    err98 := argvalue0.Read(jsProt97)
    if err98 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTypeInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetCatalogs":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetCatalogs requires 1 args")
      flag.Usage()
    }
    arg99 := flag.Arg(1)
    mbTrans100 := thrift.NewTMemoryBufferLen(len(arg99))
    defer mbTrans100.Close()
    _, err101 := mbTrans100.WriteString(arg99)
    if err101 != nil {
      Usage()
      return
    }
    factory102 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt103 := factory102.GetProtocol(mbTrans100)
    argvalue0 := tcliservice.NewTGetCatalogsReq()
    err104 := argvalue0.Read(jsProt103)
    if err104 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetCatalogs(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetSchemas":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetSchemas requires 1 args")
      flag.Usage()
    }
    arg105 := flag.Arg(1)
    mbTrans106 := thrift.NewTMemoryBufferLen(len(arg105))
    defer mbTrans106.Close()
    _, err107 := mbTrans106.WriteString(arg105)
    if err107 != nil {
      Usage()
      return
    }
    factory108 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt109 := factory108.GetProtocol(mbTrans106)
    argvalue0 := tcliservice.NewTGetSchemasReq()
    err110 := argvalue0.Read(jsProt109)
    if err110 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetSchemas(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetTables":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTables requires 1 args")
      flag.Usage()
    }
    arg111 := flag.Arg(1)
    mbTrans112 := thrift.NewTMemoryBufferLen(len(arg111))
    defer mbTrans112.Close()
    _, err113 := mbTrans112.WriteString(arg111)
    if err113 != nil {
      Usage()
      return
    }
    factory114 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt115 := factory114.GetProtocol(mbTrans112)
    argvalue0 := tcliservice.NewTGetTablesReq()
    err116 := argvalue0.Read(jsProt115)
    if err116 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTables(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetTableTypes":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTableTypes requires 1 args")
      flag.Usage()
    }
    arg117 := flag.Arg(1)
    mbTrans118 := thrift.NewTMemoryBufferLen(len(arg117))
    defer mbTrans118.Close()
    _, err119 := mbTrans118.WriteString(arg117)
    if err119 != nil {
      Usage()
      return
    }
    factory120 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt121 := factory120.GetProtocol(mbTrans118)
    argvalue0 := tcliservice.NewTGetTableTypesReq()
    err122 := argvalue0.Read(jsProt121)
    if err122 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTableTypes(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetColumns":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetColumns requires 1 args")
      flag.Usage()
    }
    arg123 := flag.Arg(1)
    mbTrans124 := thrift.NewTMemoryBufferLen(len(arg123))
    defer mbTrans124.Close()
    _, err125 := mbTrans124.WriteString(arg123)
    if err125 != nil {
      Usage()
      return
    }
    factory126 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt127 := factory126.GetProtocol(mbTrans124)
    argvalue0 := tcliservice.NewTGetColumnsReq()
    err128 := argvalue0.Read(jsProt127)
    if err128 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetColumns(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetFunctions":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetFunctions requires 1 args")
      flag.Usage()
    }
    arg129 := flag.Arg(1)
    mbTrans130 := thrift.NewTMemoryBufferLen(len(arg129))
    defer mbTrans130.Close()
    _, err131 := mbTrans130.WriteString(arg129)
    if err131 != nil {
      Usage()
      return
    }
    factory132 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt133 := factory132.GetProtocol(mbTrans130)
    argvalue0 := tcliservice.NewTGetFunctionsReq()
    err134 := argvalue0.Read(jsProt133)
    if err134 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetFunctions(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetOperationStatus":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetOperationStatus requires 1 args")
      flag.Usage()
    }
    arg135 := flag.Arg(1)
    mbTrans136 := thrift.NewTMemoryBufferLen(len(arg135))
    defer mbTrans136.Close()
    _, err137 := mbTrans136.WriteString(arg135)
    if err137 != nil {
      Usage()
      return
    }
    factory138 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt139 := factory138.GetProtocol(mbTrans136)
    argvalue0 := tcliservice.NewTGetOperationStatusReq()
    err140 := argvalue0.Read(jsProt139)
    if err140 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetOperationStatus(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CancelOperation":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CancelOperation requires 1 args")
      flag.Usage()
    }
    arg141 := flag.Arg(1)
    mbTrans142 := thrift.NewTMemoryBufferLen(len(arg141))
    defer mbTrans142.Close()
    _, err143 := mbTrans142.WriteString(arg141)
    if err143 != nil {
      Usage()
      return
    }
    factory144 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt145 := factory144.GetProtocol(mbTrans142)
    argvalue0 := tcliservice.NewTCancelOperationReq()
    err146 := argvalue0.Read(jsProt145)
    if err146 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CancelOperation(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CloseOperation":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CloseOperation requires 1 args")
      flag.Usage()
    }
    arg147 := flag.Arg(1)
    mbTrans148 := thrift.NewTMemoryBufferLen(len(arg147))
    defer mbTrans148.Close()
    _, err149 := mbTrans148.WriteString(arg147)
    if err149 != nil {
      Usage()
      return
    }
    factory150 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt151 := factory150.GetProtocol(mbTrans148)
    argvalue0 := tcliservice.NewTCloseOperationReq()
    err152 := argvalue0.Read(jsProt151)
    if err152 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CloseOperation(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetResultSetMetadata":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetResultSetMetadata requires 1 args")
      flag.Usage()
    }
    arg153 := flag.Arg(1)
    mbTrans154 := thrift.NewTMemoryBufferLen(len(arg153))
    defer mbTrans154.Close()
    _, err155 := mbTrans154.WriteString(arg153)
    if err155 != nil {
      Usage()
      return
    }
    factory156 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt157 := factory156.GetProtocol(mbTrans154)
    argvalue0 := tcliservice.NewTGetResultSetMetadataReq()
    err158 := argvalue0.Read(jsProt157)
    if err158 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetResultSetMetadata(context.Background(), value0))
    fmt.Print("\n")
    break
  case "FetchResults":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "FetchResults requires 1 args")
      flag.Usage()
    }
    arg159 := flag.Arg(1)
    mbTrans160 := thrift.NewTMemoryBufferLen(len(arg159))
    defer mbTrans160.Close()
    _, err161 := mbTrans160.WriteString(arg159)
    if err161 != nil {
      Usage()
      return
    }
    factory162 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt163 := factory162.GetProtocol(mbTrans160)
    argvalue0 := tcliservice.NewTFetchResultsReq()
    err164 := argvalue0.Read(jsProt163)
    if err164 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.FetchResults(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetDelegationToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetDelegationToken requires 1 args")
      flag.Usage()
    }
    arg165 := flag.Arg(1)
    mbTrans166 := thrift.NewTMemoryBufferLen(len(arg165))
    defer mbTrans166.Close()
    _, err167 := mbTrans166.WriteString(arg165)
    if err167 != nil {
      Usage()
      return
    }
    factory168 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt169 := factory168.GetProtocol(mbTrans166)
    argvalue0 := tcliservice.NewTGetDelegationTokenReq()
    err170 := argvalue0.Read(jsProt169)
    if err170 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetDelegationToken(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CancelDelegationToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CancelDelegationToken requires 1 args")
      flag.Usage()
    }
    arg171 := flag.Arg(1)
    mbTrans172 := thrift.NewTMemoryBufferLen(len(arg171))
    defer mbTrans172.Close()
    _, err173 := mbTrans172.WriteString(arg171)
    if err173 != nil {
      Usage()
      return
    }
    factory174 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt175 := factory174.GetProtocol(mbTrans172)
    argvalue0 := tcliservice.NewTCancelDelegationTokenReq()
    err176 := argvalue0.Read(jsProt175)
    if err176 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CancelDelegationToken(context.Background(), value0))
    fmt.Print("\n")
    break
  case "RenewDelegationToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RenewDelegationToken requires 1 args")
      flag.Usage()
    }
    arg177 := flag.Arg(1)
    mbTrans178 := thrift.NewTMemoryBufferLen(len(arg177))
    defer mbTrans178.Close()
    _, err179 := mbTrans178.WriteString(arg177)
    if err179 != nil {
      Usage()
      return
    }
    factory180 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt181 := factory180.GetProtocol(mbTrans178)
    argvalue0 := tcliservice.NewTRenewDelegationTokenReq()
    err182 := argvalue0.Read(jsProt181)
    if err182 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.RenewDelegationToken(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetLog":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetLog requires 1 args")
      flag.Usage()
    }
    arg183 := flag.Arg(1)
    mbTrans184 := thrift.NewTMemoryBufferLen(len(arg183))
    defer mbTrans184.Close()
    _, err185 := mbTrans184.WriteString(arg183)
    if err185 != nil {
      Usage()
      return
    }
    factory186 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt187 := factory186.GetProtocol(mbTrans184)
    argvalue0 := tcliservice.NewTGetLogReq()
    err188 := argvalue0.Read(jsProt187)
    if err188 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetLog(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
