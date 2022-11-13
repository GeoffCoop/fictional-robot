pipeline {
//   agent any
  agent any
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
