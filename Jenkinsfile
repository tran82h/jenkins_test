pipeline {
    agent {
        node {
            label 'linux'
        }
    }
    stages {
        stage ('Run tests') {
          steps{
            script{
            sh('cd ./test/')
            sh('go test ./...')
            }
          }
        }
    }
}