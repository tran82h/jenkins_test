#!groovy

// def installGo = "curl -SLO https://storage.googleapis.com/golang/go1.9.3.linux-amd64.tar.gz && tar -C \$HOME -xf go1.9.3.linux-amd64.tar.gz && rm -rf go1.9.3.linux-amd64.tar.gz"
pipeline {
    agent any
    // node {
    //         // Install the desired Go version
    //     def root = tool name: 'Go 1.8', type: 'go'

    //     // Export environment variables pointing to the directory where Go was installed
    //     withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
    //         sh 'go version'
    //     }
    // tools {
    //     go 'go1.14'
    // }
    // environment {
    //     GO114MODULE = 'on'
    //     CGO_ENABLED = 0
    //     GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    // }
      stages {
          stage ('Run tests') {
              // steps{
              //   sh script: """
              //       [ "\$(go version)" == 'go version go1.9.3 linux/amd64' ] || (${installGo})
              //   """
                script{
                sh('cd ./test/')
                sh('go test ./...')
                }
              }
          }
    // }
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