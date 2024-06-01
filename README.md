# pswd-cli

**pswd-cli** is a command-line tool built in Go using the Cobra framework. It allows users to securely manage their passwords for a variety of services from the command line. With encryption support, your passwords are stored securely in a local SQLite database.

## Features

- Secure Password Storage - Passwords are encrypted using AES-GCM(Advanced Encryption Mode in Galois/Counter Mode) encryption before storage in the database, ensuring their security.

- Easy-to-use CLI Interface - Features a command-line itnerface with clear commands and flags to add, retrieving, updating, and deleting passwords.

- Cross-Platform Support - The tool is compatible with various operating systems, including Linux, Windows, macOS.

- SQLite Database - A locally configured SQLite database is used for storing password records, ensuring fast and reliable access to your passwords.

## Installation

To use **pswd-cli**, you must have Go installed on your system. Then, you can clone the repostiroy from GitHub to your local machine using the following command:

```bash
git clone https://github.com/lokeshllkumar/pswd-cli.git
cd pswd-cli
```

Build the CLI.

```bash
go build -o pswd-cli
```

Move the executable to a directory in your local system's PATH, after creating a custom directory.

```bash
mkdir -p $HOME/bin
mv pswd-cli $HOME/bin/
```

Add the directory to your PATH by adding the folliwng line to your configuration file(.bashrc, .zrhc, etc.):

```bash
export PATH=$HOME/bin:$PATH
```

Save and reload the configuration file using the following line:

```bash
source ~/.bashrc
```

## Usage

Simply run the following command to get started

```bash
pswd-cli
```

### Adding a Password

To add a new password record, use the `add` command.
```bash
pswd-cli add --service "serviceName" --username "yourUsername" --password "yourPassword"
```

### Retrieving a Password

To retrieve certain stored password records, use the `get` command.

- To retrieve all stored passwords for a certain service:
```bash
pswd-cli get --service "serviceName"
```

- To retrieve the stored password for a certain username registered for a specific service:
```bash
pswd-cli get --service "serviceName" --username "yourUsername"
```

### Updating a Password

To update an exising password record by replacing the stored password for a certain username registered for a specific service, use the `update` command.
```bash
pswd-cli update --service "serviceName" --username "yourUsername" --newPassword "newPassword"
```

### Deleting a Password

To add a new password record, use the `delete` command.
- To either delete all stored passwords for a certain service:
```bash
pswd-cli delete --service "serviceName"
```

- To delete a password entry for a certain username registered for a specific service:
```bash
pswd-cli delete --service "serviceName" --username "yourUsername"
```

## Additional Features

- The `utils` directory also features methods to compute the hash of passwords and perform hash checks, to perform integrity checks while fetching data.
- The `utils` directory also includes a method to generate new 32-bit AES encryption keys.
- The table containing password records also stores the time fo creating of records for auditing purposes.`