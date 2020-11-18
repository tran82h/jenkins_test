pipeline {
    agent any
    // node {
    //         // Install the desired Go version
    //     def root = tool name: 'Go 1.8', type: 'go'

    //     // Export environment variables pointing to the directory where Go was installed
    //     withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
    //         sh 'go version'
    //     }
    tools {
        go 'go 1.15'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
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
      // }
    // environment {
    //     GO111MODULE = 'on'
    // }

}


// node {
//             // Install the desired Go version
//     def root = tool name: 'go1.2', type: 'go'

//       // Export environment variables pointing to the directory where Go was installed
//       withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
//           sh 'go version'
//       }
//       stages {
//           stage ('Run tests') {
//             steps{
//               script{
//               sh('cd ./test/')
//               sh('go test ./...')
//               }
//             }
//           }
//         }
//       }