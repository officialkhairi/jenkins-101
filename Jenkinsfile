pipeline {
  agent {
    docker { image 'golang:1.23' }
  }
  environment {
    GOCACHE    = "${WORKSPACE}/.cache/go-build"
    GOMODCACHE = "${WORKSPACE}/.cache/go-mod"
  }
  stages {
    stage('Deps')  { steps { sh 'go mod download' } }
    stage('Vet')   { steps { sh 'go vet ./...' } }
    stage('Test')  { steps { sh 'go test -v ./...' } }
    stage('Build') { steps { sh 'CGO_ENABLED=0 go build -o app .' } }
  }
  post {
    success { archiveArtifacts artifacts: 'app' }
    failure { echo 'Build failed' }
  }
}