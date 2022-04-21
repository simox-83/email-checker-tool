# Simple email checker tool

This program takes as input a domain name address from the command line and it checks if the domain is valid.

The core logic is in the checkDomain function, where I use the net package to lookup the MX and TXT records.

At the end the program outputs the domain name together with the associated records, if they're present.

TODO: the program waits for the user input indefinitely - a sort of exit should be implemented
      the program checks the domain name and the associated records, doesn't check if an email address exists or not