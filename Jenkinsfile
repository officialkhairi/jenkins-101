pipeline {
  agent { docker { image 'golang:1.23' } }
  stages {
    stage('Deps')  { steps { sh 'go mod download' } }
    stage('Vet')   { steps { sh 'go vet ./...' } }
    stage('Test')  { steps { sh 'go test -v ./...' } }
    stage('Build') { steps { sh 'go build -o app .' } }
  }
  post {
    success { archiveArtifacts artifacts: 'app' }
    failure { echo 'Build failed' }
  }
}