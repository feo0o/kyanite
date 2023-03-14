APP_NAME		=	Kyanite
MAJOR_VERSION	=	0
MINOR_VERSION	=	2
PATCH_VERSION	=	3
GIT_COMMIT		=	$(shell git rev-parse --short HEAD)
ARTIFACT_NAME	=	$(shell echo $(APP_NAME) | tr A-Z a-z)

IMPORT_VARS		=	-X github.com/feo0o/kyanite/app.Name=$(APP_NAME) \
					-X github.com/feo0o/kyanite/app.majorVer=$(MAJOR_VERSION) \
					-X github.com/feo0o/kyanite/app.minorVer=$(MINOR_VERSION) \
					-X github.com/feo0o/kyanite/app.patchVer=$(PATCH_VERSION) \
					-X github.com/feo0o/kyanite/app.gitCommit=$(GIT_COMMIT)

# ENV_WINDOWS_X64	=	GOOS=windows GOARCH=amd64
ENV_LINUX_X64	=	GOOS=linux GOARCH=amd64
# ENV_DARWIN_X64	=	GOOS=darwin GOARCH=amd64

BUILD_RELEASE	=	CGO_ENABLED=0 go build -trimpath \
					-gcflags="all=-trimpath=$(PWD)" \
					-asmflags="all=-trimpath=$(PWD)" \
					-ldflags '-extldflags "-static" $(IMPORT_VARS)'

release:main.go
	swag init
#	$(ENV_WINDOWS_X64) $(BUILD_RELEASE) -o $(RELEASE_DIR)/$(ARTIFACT_NAME)_windows_amd64.exe main.go
	$(ENV_LINUX_X64) $(BUILD_RELEASE) -o $(ARTIFACT_NAME) main.go
#	$(ENV_DARWIN_X64) $(BUILD_RELEASE) -o $(RELEASE_DIR)/$(ARTIFACT_NAME)_darwin_amd64 main.go

.PHONY:release