##bcrypt-tool


###Installation
    go get "github.com/shoenig/bcrypt-tool"

###Usage
    bcrypt-tool [action] parameter ...

###Examples

####Generate Hash from a Password
    bcrypt-tool hash p4ssw0rd
    
####Generate Hash from a Password with Cost
    bcrypt-tool hash p4ssw0rd 31
    
####Determine if Password matches Hash
    bcrypt-tool match p4ssw0rd $2a$10$nWFwjoFo4zhyVosdYMb6XOxZqlVB9Bk0TzOvmuo16oIwMZJXkpanW
> note: depending on your shell, you may need to escape the $ characters

####Determine Cost of Hash
    bcrypt-tool cost $2a$10$nWFwjoFo4zhyVosdYMb6XOxZqlVB9Bk0TzOvmuo16oIwMZJXkpanW
> note: depending on your shell, you may need to escape the $ characters
