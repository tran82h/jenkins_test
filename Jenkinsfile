pipeline {
    agent any
    tools {
        go 'go-1.11'
    }
    environment {
        GO111MODULE = 'on'
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