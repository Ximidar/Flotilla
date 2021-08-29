# Variable for filename for store running procees id
PID_FILE = /tmp/my-app.pid
# We can use such syntax to get main.go and other root Go files.
#GO_FILES = $(wildcard *.go)

# Common Go Files
COMMON_TOOLS = $(shell find ./CommonTools -type f \( -iname "*.go" ! -iname "*_test.go" \))
DATA_STRUCTURES = $(shell find ./DataStructures -type f \( -iname "*.go" ! -iname "*_test.go" \))

# Main Programs
#COMMANGO = $(shell find -wholename './Commango/*.go')
COMMANGO_ENTRY = $(wildcard ./Commango/*.go)
COMMANGO = $(shell find ./Commango -type f \( -iname "*.go" ! -iname "*_test.go" \))
COMMANGO_PID = /tmp/commango.pid

FLOTILLA_CLI_ENTRY = $(wildcard ./Flotilla_CLI/*.go)
FLOTILLA_CLI = $(shell find ./Flotilla_CLI -type f \( -iname "*.go" ! -iname "*_test.go" \))
FLOTILLA_CLI_PID = /tmp/flotilla_cli.pid

FLOTILLA_FILE_MANAGER_ENTRY = $(wildcard ./Flotilla_File_Manager/*.go)
FLOTILLA_FILE_MANAGER = $(shell find ./Flotilla_File_Manager -type f \( -iname "*.go" ! -iname "*_test.go" \))
FLOTILLA_FILE_MANAGER_PID = /tmp/flotilla_file_manager.pid

FLOTILLA_STATUS_ENTRY = $(wildcard ./FlotillaStatus/*.go)
FLOTILLA_STATUS = $(shell find ./FlotillaStatus -type f \( -iname "*.go" ! -iname "*_test.go" \))
FLOTILLA_STATUS_PID = /tmp/flotilla_status.pid


# go run $(GO_FILES) & echo $$! > $(PID_FILE)
# Start task performs "go run main.go" command and writes it's process id to PID_FILE.
start:
	go run $(COMMANGO_ENTRY) & echo $$! > $(COMMANGO_PID)
	
# You can also use go build command for start task
# start:
#   go build -o /bin/my-app . && \
#   /bin/my-app & echo $$! > $(PID_FILE)

# Stop task will kill process by ID stored in PID_FILE (and all child processes by pstree).  
stop:
	-kill `pstree -p \`cat $(COMMANGO_PID)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"` 
  
# Before task will only prints message. Actually, it is not necessary. You can remove it, if you want.
before:
	@echo "STOPED Commango" && printf '%*s\n' "40" '' | tr ' ' -
  
# Restart task will execute stop, before and start tasks in strict order and prints message. 
restart: stop before start
	@echo "STARTED Commango" && printf '%*s\n' "40" '' | tr ' ' -
  
# Serve task will run fswatch monitor and performs restart task if any source file changed. Before serving it will execute start task.
serve: start
	fswatch -or --event=Updated $(Commango) | \
	xargs -n1 -I {} make restart
  
# .PHONY is used for reserving tasks words
.PHONY: start before stop restart serve