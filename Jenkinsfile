pipeline {
    agent any

    environment {
        GOPATH = "${env.HOME}/go"
        PATH = "${env.PATH}:${env.HOME}/go/bin"
    }

    stages {
        stage('Checkout') {
            agent { label 'agent3' }  // Definir el agente para esta etapa
            steps {
                git branch: 'main', url: 'https://github.com/Forez802/go-app.git'
            }
        }

        stage('Verify Go') {
            agent { label 'agent3' }  // Definir el agente para esta etapa
            steps {
                sh 'which go || { echo "Go not found"; exit 1; }'
                sh 'go version'
            }
        }

        stage('Install Dependencies') {
            agent { label 'agent3' }  // Definir el agente para esta etapa
            steps {
                sh 'go install github.com/jstemmer/go-junit-report/v2@latest'
                sh 'which go-junit-report || { echo "go-junit-report not found"; exit 1; }'
            }
        }

        stage('Test') {
            agent { label 'agent3' }  // Definir el agente para esta etapa
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
            script {
                if (fileExists('test-reports/report.xml')) {
                    junit 'test-reports/report.xml'
                } else {
                    echo "No test report found to publish."
                }
            }
        }
    }
}
