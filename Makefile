pipeline {
    agent any

    environment {
        TARGET_GO = tool name: 'Go.1.10.3', type: 'go'
        GOBIN = "${TARGET_GO}/bin"
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/GOPATH"
        PATH= "${GOBIN}:$PATH"
    }

    stages {
        stage("Prepare") {
            steps {
                echo "TARGET_GO:${TARGET_GO}"
                echo "GOPATH:${GOPATH}"
                echo "PATH:${PATH}"
                sh 'go get -u github.com/golang/dep/...'
            }
        }

        stage("Checkout") {
            steps {
                checkout scm
            }
        }


        stage('Build') {
            steps {
                sh "make"
            }
        }
    }
}