lambda_func := mommyskill-go

build:
	GOOS=linux go build -o main

package: build
	zip -r main.zip main index.js

push: package
	aws lambda update-function-code --function-name gotsport-lambda --zip-file fileb://main.zip

clean:
	rm index.js
	rm main
	rm main.zip
