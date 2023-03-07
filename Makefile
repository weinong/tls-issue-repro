TARGET     := tls-issue-repro
OS         := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH       := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))
BIN         = bin/$(OS)_$(ARCH)/$(TARGET)
ifeq ($(OS),windows)
  BIN = bin/$(OS)_$(ARCH)/$(TARGET).exe
endif

all: $(TARGET)

$(TARGET): clean
	CGO_ENABLED=0 go build -o $(BIN)

clean:
	-rm -f $(BIN)

