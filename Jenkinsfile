pipeline {
  // Runs on the Jenkins container itself, which has the Docker CLI
  agent any

  // Check GitHub for new commits about every 2 minutes; builds only if something changed
  triggers {
    pollSCM('H/2 * * * *')
  }

  environment {
    APP_NAME = 'hello-app'
    APP_PORT = '9090'
  }

  stages {
    stage('Go') {
      // Only these stages run inside the Go container
      agent {
        docker {
          image 'golang:1.23'
          reuseNode true
        }
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
    }

    stage('Image') {
      steps {
        sh 'docker build -f Dockerfile.app -t $APP_NAME:$BUILD_NUMBER -t $APP_NAME:latest .'
      }
    }

    stage('Deploy') {
      steps {
        sh '''
          docker rm -f $APP_NAME || true
          docker run -d --name $APP_NAME --restart unless-stopped \
            -p $APP_PORT:9090 $APP_NAME:$BUILD_NUMBER
        '''
      }
    }

    stage('Verify') {
      steps {
        sh '''
          for i in $(seq 1 10); do
            if curl -fsS http://host.docker.internal:$APP_PORT/health; then
              curl -fsS http://host.docker.internal:$APP_PORT/
              exit 0
            fi
            sleep 1
          done
          echo "App did not become healthy"
          docker logs $APP_NAME
          exit 1
        '''
      }
    }
  }

  post {
    success { archiveArtifacts artifacts: 'app' }
    failure { echo 'Build failed' }
  }
}
