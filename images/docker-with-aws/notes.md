# Notes

[Login to Amazon ECR registry](https://docs.aws.amazon.com/cli/latest/reference/ecr/get-login-password.html)

Login to docker with ecr token:
`aws ecr get-login-password | docker login --username AWS --password-stdin 929684159033.dkr.ecr.ap-southeast-1.amazonaws.com/simple-bank`

Pull docker image from ecr:
`docker pull 929684159033.dkr.ecr.ap-southeast-1.amazonaws.com/simple-bank:d42fca92d89d280346acf9ca1f3f76cf6f68b96f`

Create and run container:
`docker run -p 8080:8080 929684159033.dkr.ecr.ap-southeast-1.amazonaws.com/simple-bank:d42fca92d89d280346acf9ca1f3f76cf6f68b96f`
