package githubapi

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

type OAuthServer struct {
	server     *http.Server
	listener   net.Listener
	mu         sync.Mutex
	callbackFn func(code, state, uri string)
	state      string

	// stoppedCh is closed when the server has fully stopped
	stoppedCh chan struct{}
	// shutdownTimeout is used when shutting down the server
	shutdownTimeout time.Duration
}

func NewOAuthServer(callback func(code, state, uri string)) *OAuthServer {
	return &OAuthServer{
		callbackFn:      callback,
		state:           uuid.New().String(),
		shutdownTimeout: 5 * time.Second, // shutdown 超时时间，可调整
	}
}

func (o *OAuthServer) StartOAuthServer() (callbackURL string, err error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if o.server != nil {
		return "", fmt.Errorf("OAuth server already running")
	}

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", err
	}
	o.listener = listener

	// 确保 stoppedCh 被初始化
	if o.stoppedCh == nil {
		o.stoppedCh = make(chan struct{})
	} else {
		// 如果 channel 已存在，确保它是开放的
		select {
		case <-o.stoppedCh:
			// 如果已经关闭，创建新的
			o.stoppedCh = make(chan struct{})
		default:
			// channel 仍在开放状态，无需处理
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")

		if state != o.state {
			fmt.Fprintln(w, "State不匹配，授权失败!")
			return
		}

		// notify caller but do NOT stop server here (由上层决定)
		if code != "" && o.callbackFn != nil {
			o.callbackFn(code, state, r.URL.String())
		}
		fmt.Fprintln(w, "授权成功! 你可以关闭此页面了。")
	})

	o.server = &http.Server{Handler: mux}

	// create fresh stoppedCh so each run has its own close signal
	//o.stoppedCh = make(chan struct{})

	// Serve in background
	go func() {
		// Serve will return when listener is closed or server.Shutdown called.
		_ = o.server.Serve(o.listener)
		// Ensure stoppedCh closed after Serve returns (defensive: protect double-close)
		o.mu.Lock()
		if o.stoppedCh != nil { // <-- 增加 nil 检查
			select {
			case <-o.stoppedCh:
				// already closed
			default:
				close(o.stoppedCh)
			}
		}
		o.mu.Unlock()
	}()

	return fmt.Sprintf("http://%s/callback", listener.Addr().String()), nil
}

// StopOAuthServer attempts graceful shutdown and closes stoppedCh when done.
// It is safe to call multiple times.
func (o *OAuthServer) StopOAuthServer() error {
	o.mu.Lock()

	// 检查 stoppedCh 是否为 nil
	if o.stoppedCh == nil {
		o.mu.Unlock()
		return nil // 或者返回一个适当的错误
	}
	// capture local references
	srv := o.server
	lis := o.listener
	stopped := o.stoppedCh
	timeout := o.shutdownTimeout
	// mark fields nil to indicate stopped (pre-empt other callers)
	o.server = nil
	o.listener = nil
	o.stoppedCh = nil
	o.mu.Unlock()

	if srv == nil {
		// already stopped
		return nil
	}

	// graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	_ = srv.Shutdown(ctx)

	// close listener to ensure Serve returns
	if lis != nil {
		_ = lis.Close()
	}

	// wait for Serve goroutine to acknowledge stopped; if stoppedCh exists, wait a short time
	if stopped != nil {
		select {
		case <-stopped:
			// normal
		case <-time.After(timeout):
			// timeout waiting for serve to finish
		}
	}

	// ensure stoppedCh closed (defensive)
	select {
	case <-stopped:
		// already closed
	default:
		// try to close safely: acquire lock and close if still set
		o.mu.Lock()
		if stopped != nil {
			// cannot close a nil channel; since we moved stopped out, we must avoid double-close
			// but here it's safe to close local stopped if not already closed
			close(stopped)
		}
		o.mu.Unlock()
	}

	return nil
}
