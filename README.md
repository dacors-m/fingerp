```
     _____ ___ _   _  ____ _____ ____  ____
    |  ___|_ _| \ | |/ ___| ____|  _ \|  _ \
    | |_   | ||  \| | |  _|  _| | |_) | |_) |
    |  _|  | || |\  | |_| | |___|  _ <|  __/
    |_|   |___|_| \_|\____|_____|_| \_\_|

```
## Minimalistic and really usefull cli password manager for your personal use!
This cli application allows you to create password based on different parameters
like length, ammount of characters and amount of number. Fingerp autogenerates these
passwords using a random algorithm.

## Installation
```
go install github.com/dacors-m/fingerp@latest
```

## Password Gen

If you want to generate a password, run the command ```fingerp pgen```, it will
generate you a password with the deffault values:
- Length 8
- 2 characters 
- 2 numbers
