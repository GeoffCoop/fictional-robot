pipeline {
//   agent any
  agent { docker { image 'golang:1.19' } }
  stages {
    stage('test') {
      steps {
        sh 'go version'
        sh 'ls -la'
        sh 'GOCACHE=off make test'
      }
    }

    stage('build') {
      steps {
        sh 'make'
      }
    }

  }
}
