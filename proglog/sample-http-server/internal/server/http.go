// START: newhttpserver
package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHTTPServer(addr string) *http.Server {
	/*
		実行するサーバのアドレスを受け取り、*http.serverを返す
		muxライブラリ使用(rest endpointのため)
	*/
	httpsrv := newHTTPServer()
	r := mux.NewRouter()
	r.HandleFunc("/", httpsrv.handleProduce).Methods("POST")
	r.HandleFunc("/", httpsrv.handleConsume).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

// END: newhttpserver

// START: types
type httpServer struct {
	Log *Log
}

func newHTTPServer() *httpServer {
	return &httpServer{
		Log: NewLog(),
	}
}

type ProduceRequest struct {
	Record Record `json:"record"`
}

type ProduceResponse struct {
	Offset uint64 `json:"offset"`
}

type ConsumeRequest struct {
	Offset uint64 `json:"offset"`
}

type ConsumeResponse struct {
	Record Record `json:"record"`
}

// ハンドラを追加する必要がある
// START:produce
func (s *httpServer) handleProduce(w http.ResponseWriter, r *http.Request) {
	//リクエストを構造体をアンマーシャル(*1)する
	// *1 JSON形式で受け取った値を指定したポインタに保存
	var req ProduceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// その構造体を使って、ログにレコードを保存し、ログがレコードを保存したオフセットを取得
	off, err := s.Log.Append(req.Record) // offが保存した時のindex
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// その結果をマーシャルしてレスポンスに書き込む
	res := ProduceResponse{Offset: off}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// START:consume
func (s *httpServer) handleConsume(w http.ResponseWriter, r *http.Request) {
	var req ConsumeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	record, err := s.Log.Read(req.Offset)
	if err == ErrOffsetNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := ConsumeResponse{Record: record}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// END:consume
