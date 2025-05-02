pipeline {
    agent any

    environment {
        GOPATH = "${env.HOME}/go"
        PATH = "${env.PATH}:${env.HOME}/go/bin"
    }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/Forez802/go-app.git'
            }
        }

        stage('Verify Go') {
            steps {
                sh 'which go'
                sh 'go version'
            }
        }

        stage('Install Dependencies') {
            steps {
                sh 'go install github.com/jstemmer/go-junit-report/v2@latest'
                sh 'which go-junit-report || { echo "go-junit-report not found!"; exit 1; }'
            }
        }

        stage('Test') {
            steps {
                sh '''
                    mkdir -p test-reports
                    go test -v ./... | go-junit-report -set-exit-code > test-reports/report.xml
                '''
            }
        }
    }

    post {
        always {
            junit 'test-reports/report.xml'
        }
    }
}
