pipeline {
    agent { docker { image 'golang' } }
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