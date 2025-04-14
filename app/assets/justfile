run *ARGS:
    @echo -e "\e[1;32mrunning project...\033[0m"
    @go run ./app {{ARGS}}

build:
    @echo -e "\e[1;32mbuilding project...\033[0m"
    @rm -rf ./.build
    @go build -o ./.build/bin ./app
    @echo -e "\e[1;32mbuild successful\033[0m"

test:
    @echo -e "\e[1;32mtesting project...\033[0m"
    @go test ./...
    @echo -e "\e[1;32mtests passed\033[0m"