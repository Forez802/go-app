pipeline {
    agent { label 'agent3' }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/Forez802/go-app.git'
            }
        }

        stage('Verify Go') {
            steps {
                sh 'echo $PATH'
                sh 'which go'
                sh 'go version'
            }
        }

        stage('Test') {
            steps {
                sh '''
                    mkdir -p test-reports
                    export PATH=$PATH:$HOME/go/bin
                    go install github.com/jstemmer/go-junit-report/v2@latest
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
