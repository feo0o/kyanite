APP_NAME		=	Kyanite
MAJOR_VERSION	=	0
MINOR_VERSION	=	3
PATCH_VERSION	=	0
VERSION			=	v$(MAJOR_VERSION).$(MINOR_VERSION).$(PATCH_VERSION)
GIT_COMMIT		=	$(shell git rev-parse --short HEAD)
ARTIFACT_NAME	=	$(shell echo $(APP_NAME) | tr A-Z a-z)
RELEASE_DIR		=	release

IMPORT_VARS		=	-X github.com/feo0o/kyanite/app.Name=$(APP_NAME) \
					-X github.com/feo0o/kyanite/app.majorVer=$(MAJOR_VERSION) \
					-X github.com/feo0o/kyanite/app.minorVer=$(MINOR_VERSION) \
					-X github.com/feo0o/kyanite/app.patchVer=$(PATCH_VERSION) \
					-X github.com/feo0o/kyanite/app.gitCommit=$(GIT_COMMIT)

ENV_LINUX_X64	=	GOOS=linux GOARCH=amd64

BUILD_RELEASE	=	CGO_ENABLED=0 go build -trimpath \
					-gcflags="all=-trimpath=$(PWD)" \
					-asmflags="all=-trimpath=$(PWD)" \
					-ldflags '-extldflags "-static" $(IMPORT_VARS)'

PKG_NAME_LINUX_AMD64	=	$(ARTIFACT_NAME)_$(VERSION)_linux_amd64

release:binary container clean

binary:main.go
	go install -v github.com/swaggo/swag/cmd/swag@latest
	go get -u -v github.com/swaggo/swag
	go mod download
	go mod verify
	swag init
	$(ENV_LINUX_X64) $(BUILD_RELEASE) -o $(RELEASE_DIR)/$(PKG_NAME_LINUX_AMD64) main.go

container:Dockerfile $(RELEASE_DIR)/$(PKG_NAME_LINUX_AMD64)
	cp $(RELEASE_DIR)/$(PKG_NAME_LINUX_AMD64) $(RELEASE_DIR)/kyanite
	docker build -t $(ARTIFACT_NAME):$(VERSION) -f Dockerfile .

clean:$(RELEASE_DIR)/kyanite
	rm -f $(RELEASE_DIR)/kyanite

.PHONY:release binary container clean