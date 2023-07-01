
## IIS Short Name Scanner

The IIS Short Name Scanner is a set of scripts written in Go that allows you to scan a target website for the presence of short names in IIS (Internet Information Services) URLs. It can help identify potential security vulnerabilities related to IIS short name disclosure.

/ _ () () | | | | | |
/ /\ _ __ _ _ __ __ _ ___ | | | | __ _ | |
| _ | ' | | ' \ / |/ _ \| __| | __/ _ / __| __|
| | | | |) | | | | | (| | () | |_ | | | (| _ \ |_
_| || .__/||| ||_, |___/ _| || _,|/_|
| | __/ |
|| |_/


### Features

- **Path Enumeration**: The `path_enumeration.go` script generates a list of paths containing common IIS short names and saves them to a file (`paths.txt` by default).
- **Short Name Scanning**: The `shortname_scanner.go` script reads the generated paths from `paths.txt` and scans each URL for the presence of existing resources.
- **HTTP and HTTPS Support**: Both scripts support scanning websites using either HTTP or HTTPS protocols.
- **Concurrency**: The scripts use concurrent scanning to speed up the scanning process by processing multiple paths simultaneously.
- **Timeouts**: Request timeouts are implemented to prevent the scanner from getting stuck on unresponsive or slow resources.
- **Error Logging**: Errors encountered during scanning are logged to help identify and troubleshoot any issues.
- **Progress Indicator**: The scripts display a progress bar indicating the scanning progress.
- **Exportable Results**: The `shortname_scanner.go` script prints the found paths to the console and also exports them to `output.txt` for further analysis.

### Usage

1. Update the `targetURL` variable in both scripts to the target website's URL.
2. Run the `path_enumeration.go` script to generate the list of paths using the command: `go run path_enumeration.go`.
3. Run the `shortname_scanner.go` script to scan the target website for existing paths using the command: `go run shortname_scanner.go`.
4. Monitor the console output for any found paths and check the `output.txt` file for the exported results.

### Requirements

- Go programming language (https://golang.org/)

### Disclaimer

Please use this tool responsibly and only on websites you have permission to scan. Unauthorized scanning of websites can violate applicable laws and regulations. The authors assume no liability for any unauthorized or misuse of this tool.

