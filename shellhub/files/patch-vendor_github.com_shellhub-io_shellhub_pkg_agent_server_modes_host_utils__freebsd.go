--- vendor/github.com/shellhub-io/shellhub/pkg/agent/server/modes/host/utils_freebsd.go.orig	2024-11-05 21:16:27 UTC
+++ vendor/github.com/shellhub-io/shellhub/pkg/agent/server/modes/host/utils_freebsd.go
@@ -6,11 +6,15 @@ import (
 	"os"
 	"os/exec"
 
+	gliderssh "github.com/gliderlabs/ssh"
 	"github.com/shellhub-io/shellhub/pkg/agent/pkg/osauth"
 	"github.com/shellhub-io/shellhub/pkg/agent/server/modes/host/command"
 )
 
-func newShellCmd(deviceName string, username string, term string, envs []string) *exec.Cmd {
+func generateShellCmd(deviceName string, session gliderssh.Session, term string) *exec.Cmd {
+	username := session.User()
+	envs := session.Environ()
+
 	shell := os.Getenv("SHELL")
 
 	user, err := osauth.LookupUser(username)
