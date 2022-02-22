
# https://gist.github.com/sighingnow/deee806603ec9274fd47
# https://stackoverflow.com/questions/714100/os-detecting-makefile

ifeq ($(OS),Windows_NT)
    OS_ARCH += WIN32
    ifeq ($(PROCESSOR_ARCHITEW6432),AMD64)
        OS_ARCH += AMD64
    else
        ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
            OS_ARCH += AMD64
        endif
        ifeq ($(PROCESSOR_ARCHITECTURE),x86)
            OS_ARCH += IA32
        endif
    endif
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        OS_ARCH += LINUX
    endif
    ifeq ($(UNAME_S),Darwin)
        OS_ARCH += OSX
    endif
    UNAME_P := $(shell uname -p)
    ifeq ($(UNAME_P),x86_64)
        OS_ARCH += AMD64
    endif
    ifneq ($(filter %86,$(UNAME_P)),)
        OS_ARCH += IA32
    endif
    ifneq ($(filter arm%,$(UNAME_P)),)
        OS_ARCH += ARM
    endif
endif

GOPATH=$(shell go env GOPATH)

FINAL_CMD=

ifeq ($(OS_ARCH),LINUX AMD)
	FINAL_CMD=linux amd command
endif
ifeq ($(OS_ARCH),LINUX ARM)
	FINAL_CMD=linux arm command
endif
ifeq ($(OS_ARCH),OSX AMD)
	FINAL_CMD=osx x86_64 command
endif
ifeq ($(OS_ARCH),OSX IA32)
	FINAL_CMD=osx x86_64 command
endif
ifeq ($(OS_ARCH),OSX ARM)
	FINAL_CMD=osx arm command
endif

all:
	@echo $(OS_ARCH)
	@echo $(GOPATH)
	@echo $(FINAL_CMD)

