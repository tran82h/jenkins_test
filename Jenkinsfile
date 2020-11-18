pipeline {
    agent any
    stages {
        stage ('Run tests') {
            sh('cd ./test/')
            sh('go test ./...')
        }
    }
}