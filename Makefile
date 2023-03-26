BINAY_NAME=goldWatcher
APP_NAME=goldWatcher
VERSION=1.0.0
BUILD_NO=1

build:
	rm -f ${BINAY_NAME}
	fyne package -appVersion ${VERSION} -appBuild ${BUILD_NO} -name ${APP_NAME} -release
	rm -f go-for-goldWatcher

run:
	env DB_PATH="./sql.db" go run .

clean:
	@echo "Cleaing..."
	@go clean
	@rm -rf ${BINAY_NAME}
	@echo "Cleaned!"

test:
	go test -v ./...