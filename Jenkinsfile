pipeline {
    agent any

    tools { go '1.22.4' }

    stages {
        stage('Test Code') {
            steps {
                sh 'go test .'
            }
        }
        stage('Build Image') {
            steps {
                sh 'go build'
            }
        }
        stage('Deploy') {
            steps {
                sh 'deploy.sh'
            }
        }
    }
}
