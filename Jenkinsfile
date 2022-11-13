pipeline {
//   agent any
agent { 
  docker { image 'golang:1.19.1-alpine' }
  stages {
    stage('test') {
      steps {
        sh 'go version'
        sh 'ls -la'
        sh 'make test'
      }
    }

    stage('build') {
      steps {
        sh 'make'
      }
    }

  }
}
