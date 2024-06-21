pipeline {
    agent any

    stages {
        stage('Checkout Code') {
            steps {
                git branch: 'main'
            }
        }
        stage('Test Code') {
            steps {
                sh 'go test .'
            }
        }
        stage('Build Image') {
            steps {
                sh 'go build -o '
            }
        }
        stage('Deploy') {
            steps {
                sh 'docker-compose down'
                sh 'docker-compose up -d'  
            }
        }
    }
}
