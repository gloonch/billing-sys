pipeline {
    agent any
        environment {
            DOCKER_IMAGE = "go_app:${BUILD_NUMBER}"
        }
    stages {
        stage('Checkout') {
            steps {
                checkout([
                    $class: 'GitSCM',
                    branches: [[name: '*/main']],
                    userRemoteConfigs: [[url: 'https://github.com/gloonch/billing-sys.git']]
                ])
            }
        }
        stage('Build Docker Image') {
            steps {
                script {
                    sh 'docker build -t $DOCKER_IMAGE .'
                }
            }
        }
        stage('Run Tests') {
            steps {
                script {
                    sh 'docker run --rm $DOCKER_IMAGE go test ./...'
                }
            }
        }
        stage('Run Application') {
            steps {
                script {
                    sh 'docker run -d -p 8000:8000 $DOCKER_IMAGE'
                }
            }
        }
    }
    post {
        always {
            echo 'Pipeline completed!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}
