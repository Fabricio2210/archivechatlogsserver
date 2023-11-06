pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: 'main']], userRemoteConfigs: [[url: 'https://github.com/Fabricio2210/archivechatlogsserver.git']]])
            }
        }

        stage('Build') {
            steps {
                sh 'go get -v'
                sh 'go build -o gofiber'
            }
        }

        stage('Print Workspace') {
            steps {
                echo "Workspace directory: ${WORKSPACE}"
            }
        }

        stage('Deploy') {
            steps {
                echo 'Deploying the application...'
                script {
                    sh 'sudo systemctl restart gofiber'
                }
            }
        }
    }
}
