pipeline {
  agent any
  stages {
    stage('test') {
      steps {
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