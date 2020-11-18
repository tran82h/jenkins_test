pipeline {
    agent any
    node {
            // Install the desired Go version
        def root = tool name: 'Go 1.8', type: 'go'

        // Export environment variables pointing to the directory where Go was installed
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
            sh 'go version'
        }
    }

    }
    // environment {
    //     GO111MODULE = 'on'
    // }
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