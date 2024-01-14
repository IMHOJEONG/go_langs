
- go build에서 오류 발생 시,

GO111MODULE이 설정되어 있지 않아, 문제가 발생한듯 

go env
set GO111MODULE=

==> go env -w 
GO111MODULE=auto 명령어를 통해 하위 처럼 추가된 것을 볼 수 있음 

go env    
set GO111MODULE=auto

- gopls was not able to find modules in your workspace.
When outside of GOPATH, gopls needs to know which modules you are working on.
You can fix this by opening your workspace to a folder inside a Go module, or
by using a go.work file to specify multiple modules.
See the documentation for more information on setting up your workspace:
https://github.com/golang/tools/blob/master/gopls/doc/workspace.md.go list



- expected boolean or range expression, found assignment (missing parentheses around composite literal?)syntax

- defer, panic, recover 

    - https://go.dev/tour/flowcontrol/12


CodinGame: https://www.codingame.com/training
CodeGym: https://codegym.cc/ko/quests
CSS Diner: https://flukeout.github.io/

1. CodinGame
2. CodeGym
3. CSS Diner


https://go101.org/article/nil.html

http://golang.site/go/article/11-Go-%ED%81%B4%EB%A1%9C%EC%A0%80

https://go.dev/tour/methods/1

