pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                echo 'Codigo clonado'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./... -v -json > report.json'
            }
        }
        stage('Build') {
            steps {
                sh 'go build -o app'
            }
        }
    }
}
