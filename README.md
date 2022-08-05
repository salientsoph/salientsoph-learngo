
# Learning Go lang (2022.08. ~ ) 

# 0. INTRODUCTION
## 0.1 to start go, initial setting 
1. download go 
2. go to `'go'` directory > `src` > mkdir "github.com" 
3. go to `'github.com'` directory > mkdir "github_username"
4. go to `'github_username'` directory > touch "main.go"
5. open vscode and set all the popups 
6. open folder

## 0.2 in the new file, initial setting
1. package ____   (no needs to write "from ____" (ex.java) )
2. import ____ 
3. fmt.**P**rintln: use uppercase (if it is, it's not a private function. lowercase: private function.) 


# 1. THEORY
## 1.1 




# PATH setting 
윈도우 기준으로 go 설치시 
1. 자동으로 환경변수에 GOROOT, GOPATH가 세팅된다. 
  - GOROOT: go가 설치된 디렉토리. go 표준 패키지 사용 용도. 
  - GOPATH: 사용자가 작업하게 되는 루트. github에서 third party 패키지를 사용하게 될 경우 저장되는 곳. 
2. Path에 GOROOT\bin 이 설정된다. 

### github import error 
내 경우, go는 C:\Go 에 바로 설치했기 때문에(go 설치시 default는 C:\user\users(이름)\Go 였음) 이를 GOROOT로 잡아줬고(GOROOT가 왜인지 자동으로 생성 안되어있었음)
GOPATH는 C:\user\users\go 파일로 잡아줬음(왜 각각으로 되어있는지 모르겠으나, src 파일 많아지는 것보단 나은듯)
